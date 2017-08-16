package meta

import (
	"github.com/uol/gobol"
	"github.com/uol/gobol/rubber"
	"github.com/uol/mycenae/lib/bcache"
	pb "github.com/uol/mycenae/lib/proto"
	"github.com/uol/mycenae/lib/tsstats"
	"go.uber.org/zap"
)

// Tag represents a tag key-value pair
type Tag struct {
	Key   string `json:"tagKey"`
	Value string `json:"tagValue"`
}

// ErrorData represents an error
type ErrorData struct {
	Key    string `json:"key"`
	Metric string `json:"metric"`
	Tags   []Tag  `json:"tagsError"`
}

// Backend defines the behaviour of Meta
type Backend interface {
	Handle(pkt *pb.Meta) bool
	SaveTxtMeta(packet *pb.Meta)
	CheckTSID(dtype, id string) (bool, gobol.Error)
	SendError(index, dtype, id string, doc ErrorData) gobol.Error

	CreateIndex(index string) gobol.Error
	DeleteIndex(index string) gobol.Error
}

// Meta is a wrapper around the meta backend
type Meta struct {
	Backend
}

// Create creates the wrapper given a backend
func Create(backend Backend) *Meta {
	return &Meta{
		Backend: backend,
	}
}

// New creates an elastic search meta storage
func New(log *zap.Logger, sts *tsstats.StatsTS, es *rubber.Elastic, bc *bcache.Bcache, set *Settings) (*Meta, error) {
	em, err := createElasticMeta(log, sts, es, bc, set)
	if err != nil {
		return nil, err
	}
	return Create(em), nil
}
