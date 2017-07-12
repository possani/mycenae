package meta

import (
	"io"
	"net/http"
	"time"

	"github.com/uol/gobol"
	"github.com/uol/gobol/rubber"
	index "github.com/uol/mycenae/lib/ts-index"
)

type persistence struct {
	esearch *rubber.Elastic
	index   *index.Set
}

func (persist *persistence) HeadMetaFromES(esindex, eType, id string) (bool, gobol.Error) {
	if persist.index != nil {
		exists, err := persist.index.Get(esindex, eType).Exists(index.ParseID(id))
		if err != nil {
			return false, errPersist("HeadMetaFromES", err)
		}
		return exists, nil
	}

	start := time.Now()
	respCode, err := persist.esearch.GetHead(esindex, eType, id)
	if err != nil {
		statsIndexError(esindex, eType, "head")
		return false, errPersist("HeadMetaFromES", err)
	}
	statsIndex(esindex, eType, "head", time.Since(start))
	return respCode == http.StatusOK, nil
}

func (persist *persistence) SaveBulkES(body io.Reader) gobol.Error {
	start := time.Now()
	_, err := persist.esearch.PostBulk(body)
	if err != nil {
		statsIndexError("", "", "bulk")
		return errPersist("SaveBulkES", err)
	}
	statsIndex("", "", "bulk", time.Since(start))
	return nil
}
