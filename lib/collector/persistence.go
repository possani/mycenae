package collector

import (
	"time"

	"github.com/uol/gobol"
	"github.com/uol/mycenae/lib/cluster"
	"github.com/uol/mycenae/lib/depot"
	"github.com/uol/mycenae/lib/gorilla"
	"github.com/uol/mycenae/lib/meta"

	pb "github.com/uol/mycenae/lib/proto"
)

type persistence struct {
	cluster *cluster.Cluster
	cass    *depot.Cassandra
	meta    *meta.Meta
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

func (persist *persistence) SendErrorToES(index, eType, id string, doc meta.ErrorData) gobol.Error {
	return persist.meta.SendError(index, eType, id, doc)
}
