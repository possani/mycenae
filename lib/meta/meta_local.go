package meta

import (
	"strings"

	"github.com/uol/gobol"
	pb "github.com/uol/mycenae/lib/proto"
	index "github.com/uol/mycenae/lib/ts-index"
	"go.uber.org/zap"
)

type localMeta struct {
	index  *index.Set
	logger *zap.Logger
}

func createLocal(logger *zap.Logger) *localMeta {
	return &localMeta{
		index:  index.CreateSet(),
		logger: logger,
	}
}

// NewLocal creates a local file meta store
func NewLocal(logger *zap.Logger) Backend {
	return createLocal(logger)
}

func (m *localMeta) Handle(pkt *pb.Meta) bool {
	m.logger.Debug("Handling package",
		zap.String("function", "Handle"),
		zap.String("structurte", "localMeta"),
		zap.String("package", "meta"),
	)
	lindex := m.index.Get(pkt.GetKsid(), "meta")
	if lindex == nil {
		m.CreateIndex(pkt.GetKsid())
		lindex = m.index.Get(pkt.GetKsid(), "meta")
		m.logger.Debug("Creating index",
			zap.String("index", pkt.GetKsid()),
			zap.String("function", "Handle"),
			zap.String("structurte", "localMeta"),
			zap.String("package", "meta"),
		)
	}
	found, err := lindex.Exists(index.ParseID(pkt.GetTsid()))
	if err != nil {
		m.logger.Debug("Error getting index",
			zap.String("function", "Handle"),
			zap.String("structurte", "localMeta"),
			zap.String("package", "meta"),
		)
	}
	if !found {
		err = lindex.Add(index.Metric(pkt.GetMetric()), index.ParseTags(pkt.GetTags()), index.ParseID(pkt.GetTsid()))
		if err != nil {
			m.logger.Debug("Error indexing timeseries",
				zap.String("function", "Handle"),
				zap.String("structurte", "localMeta"),
				zap.String("package", "meta"),
			)
		}
	}
	return (err == nil && !found)
}

func (m *localMeta) SaveTxtMeta(packet *pb.Meta) {
	var (
		keyspace   = packet.GetKsid()
		timeseries = index.ParseID(packet.GetTsid())
		metric     = index.Metric(packet.GetMetric())

		tags []index.KVPair
	)
	for _, tag := range packet.GetTags() {
		tags = append(tags, index.KVPair{Key: tag.GetKey(), Value: tag.GetValue()})
	}
	m.index.Get(keyspace, "metatext").Add(metric, tags, timeseries)
}

func (m *localMeta) CheckTSID(esType, id string) (bool, gobol.Error) {
	info := strings.Split(id, "|")
	esindex, id := info[0], info[1]
	exists, err := m.index.Get(esindex, esType).Exists(index.ParseID(id))
	if err != nil {
		return false, errPersist("Exists", err)
	}
	return exists, nil
}

func (m *localMeta) CreateIndex(name string) gobol.Error {
	m.index.Add(name, "meta", index.Create())
	return nil
}

func (m *localMeta) DeleteIndex(index string) gobol.Error {
	m.index.Delete(index, "meta")
	return nil
}
