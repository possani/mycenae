package collector

import (
	"time"

	"github.com/uol/gobol"
	"github.com/uol/gobol/rubber"
	"github.com/uol/mycenae/lib/cluster"
	"github.com/uol/mycenae/lib/depot"
	"github.com/uol/mycenae/lib/gorilla"

	pb "github.com/uol/mycenae/lib/proto"
)

type persistence struct {
	cluster *cluster.Cluster
	esearch *rubber.Elastic
	cass    *depot.Cassandra
}

func (persist *persistence) InsertPoint(packet *gorilla.Point) gobol.Error {
	p := &pb.TSPoint{
		Ksid:  packet.KsID,
		Tsid:  packet.ID,
		Date:  packet.Timestamp,
		Value: *packet.Message.Value,
	}
	return persist.cluster.Write([]*pb.TSPoint{p})
}

func (persist *persistence) InsertText(ksid, tsid string, timestamp int64, text string) gobol.Error {
	return persist.cass.InsertText(ksid, tsid, timestamp, text)
}

func (persist *persistence) InsertError(id, msg, errMsg string, date time.Time) gobol.Error {
	return persist.cass.InsertError(id, msg, errMsg, date)
}

func (persist *persistence) SendErrorToES(index, eType, id string, doc StructV2Error) gobol.Error {
	start := time.Now()
	_, err := persist.esearch.Put(index, eType, id, doc)
	if err != nil {
		statsIndexError(index, eType, "put")
		return errPersist("SendErrorToES", err)
	}
	statsIndex(index, eType, "PUT", time.Since(start))
	return nil
}
