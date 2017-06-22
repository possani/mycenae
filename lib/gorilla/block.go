package gorilla

import (
	"io"
	"sync"

	"go.uber.org/zap"

	tsz "github.com/uol/go-tsz"
	pb "github.com/uol/mycenae/lib/proto"
)

// block contains compressed points
type block struct {
	mtx      sync.RWMutex
	points   []byte
	id       int64
	prevDate int64
	enc      *tsz.Encoder
}

func (b *block) reset(id int64) {
	log := gblog.With(
		zap.String("package", "storage/block"),
		zap.String("func", "Add"),
		zap.Int64("blkid", b.id),
		zap.Int64("newBlkid", id),
	)

	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.id = id
	b.points = nil
	b.prevDate = 0
	if b.enc != nil {
		log.Debug("resetting block with open encoding")
		p, err := b.enc.Close()
		if err != nil {
			log.Error(
				"problem to close tsz",
				zap.Error(err),
			)
		}
		log.Debug(
			"encoding closed",
			zap.Int("ptsSize", len(p)),
		)
	}
	b.enc = nil
	return
}

func (b *block) add(date int64, value float32) {
	log := gblog.With(
		zap.String("package", "storage/block"),
		zap.String("func", "Add"),
		zap.Int64("blkid", b.id),
	)

	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.enc == nil {
		b.enc = tsz.NewEncoder(b.id)
		b.enc.Encode(date, value)
		b.prevDate = date
		return
	}

	if date > b.prevDate {
		b.enc.Encode(date, value)
		b.prevDate = date
		return
	}

	p, err := b.enc.Close()
	if err != nil {
		log.Error(
			err.Error(),
			zap.Error(err),
			zap.Int("blockSize", len(p)),
		)
		b.enc = nil
		return
	}

	err = b.newEncode(p, date, value)
	if err != nil {
		log.Error(
			err.Error(),
			zap.Error(err),
			zap.Int("blockSize", len(p)),
		)
	}
	return
}

func (b *block) close() []byte {
	log := gblog.With(
		zap.String("package", "storage/block"),
		zap.String("func", "close"),
		zap.Int64("blkid", b.id),
	)
	b.mtx.Lock()
	defer b.mtx.Unlock()
	if b.enc != nil {
		pts, err := b.enc.Close()
		if err != nil {
			log.Error(
				err.Error(),
				zap.Error(err),
				zap.Int("blockSize", len(pts)),
			)
		}
		b.points = pts
		b.enc = nil
		return pts
	}

	return nil
}

func (b *block) newEncode(pByte []byte, date int64, value float32) error {
	log := gblog.With(
		zap.String("package", "storage/block"),
		zap.String("func", "newEncode"),
		zap.Int64("blkid", b.id),
	)

	points, count, err := b.decode(pByte)
	if err != nil {
		log.Error(
			err.Error(),
			zap.Error(err),
		)
		return err
	}

	delta := int(date - b.id)

	log.Debug(
		"point delta",
		zap.Int("delta", delta),
	)

	if points[delta] == nil {
		count++
	}

	points[delta] = &pb.Point{Date: date, Value: value}

	enc := tsz.NewEncoder(b.id)
	for _, p := range points {
		if p != nil {
			enc.Encode(p.Date, p.Value)
		}
	}

	b.enc = enc

	return nil
}

func (b *block) decode(points []byte) ([bucketSize]*pb.Point, int, error) {
	id := b.id
	dec := tsz.NewDecoder(points)

	var pts [bucketSize]*pb.Point
	var d int64
	var v float32

	var count int

	log := gblog.With(
		zap.String("package", "storage"),
		zap.String("func", "block/decode"),
		zap.Int64("blkid", id),
		zap.Int("blockSize", len(points)),
	)

	for dec.Scan(&d, &v) {
		delta := d - id
		if delta >= 0 && delta < bucketSize {
			pts[delta] = &pb.Point{Date: d, Value: v}
			count++
		}
	}

	if err := dec.Close(); err != nil {
		log.Error(
			err.Error(),
			zap.Error(err),
			zap.Int("count", count),
		)
		return [bucketSize]*pb.Point{}, 0, err
	}

	log.Debug(
		"finished tsz decoding",
		zap.Int("count", count),
	)

	return pts, count, nil

}

func (b *block) SetPoints(pts []byte) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if len(pts) >= headerSize {
		b.points = pts
	}
}

func (b *block) GetPoints() []byte {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	return b.points
}

func (b *block) rangePoints(id int, start, end int64, queryCh chan query) {
	b.mtx.Lock()
	points := b.points
	if b.enc != nil {
		points = b.enc.Get()
	}
	bktid := b.id
	b.mtx.Unlock()

	var pts []*pb.Point
	if len(points) >= headerSize {
		if end >= bktid || start >= bktid {

			dec := tsz.NewDecoder(points)

			var d int64
			var v float32
			for dec.Scan(&d, &v) {
				if d >= start && d <= end {
					pts = append(pts, &pb.Point{
						Date:  d,
						Value: v,
					})
				}
			}

			err := dec.Close()
			if err != io.EOF && err != nil {
				gblog.Error("", zap.Error(err))
			}

		}
	}

	queryCh <- query{
		id:  id,
		pts: pts,
	}

}
