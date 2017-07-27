package meta

import (
	"strings"

	"github.com/uol/gobol"
	pb "github.com/uol/mycenae/lib/proto"
	index "github.com/uol/mycenae/lib/ts-index"
)

type localMeta struct {
	index *index.Set
}

// NewLocal creates a local file meta store
func NewLocal() Backend {
	return &localMeta{
		index: index.CreateSet(),
	}
}

func (m *localMeta) Handle(pkt *pb.Meta) bool {
	found, err := m.index.Get(pkt.GetKsid(), "meta").Exists(index.ParseID(pkt.GetTsid()))
	return (err != nil || !found)
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
