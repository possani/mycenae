package meta

import (
	"net/http"
	"time"

	"github.com/uol/gobol"
	"github.com/uol/gobol/rubber"
	"github.com/uol/mycenae/lib/tsstats"
)

type persistence struct {
	esearch *rubber.Elastic
	stats   *tsstats.StatsTS
}

func (persist *persistence) HeadMetaFromES(esindex, eType, id string) (bool, gobol.Error) {
	start := time.Now()
	respCode, err := persist.esearch.GetHead(esindex, eType, id)
	if err != nil {
		statsIndexError(persist.stats, esindex, eType, "head")
		return false, errPersist("HeadMetaFromES", err)
	}
	statsIndex(persist.stats, esindex, eType, "head", time.Since(start))
	return respCode == http.StatusOK, nil
}
