package meta

import (
	"github.com/uol/gobol/rubber"
	"github.com/uol/mycenae/lib/bcache"
	"github.com/uol/mycenae/lib/tsstats"
	"go.uber.org/zap"
)

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
