package collector

import (
	"fmt"
	"time"

	"github.com/uol/gobol"
	"github.com/uol/mycenae/lib/gorilla"
)

func (collector *Collector) makePacket(packet *gorilla.Point, rcvMsg gorilla.TSDBpoint, number bool) gobol.Error {

	if number {
		if rcvMsg.Value == nil {
			return errValidation(`Wrong Format: Field "value" is required. NO information will be saved`)
		}
	} else {
		if rcvMsg.Text == "" {
			return errValidation(`Wrong Format: Field "text" is required. NO information will be saved`)
		}

		if len(rcvMsg.Text) > 10000 {
			return errValidation(`Wrong Format: Field "text" can not have more than 10k`)
		}
	}

	lt := len(rcvMsg.Tags)

	if lt == 0 {
		return errValidation(`Wrong Format: At least one tag is required. NO information will be saved`)
	}

	if !collector.validKey.MatchString(rcvMsg.Metric) {
		return errValidation(
			fmt.Sprintf(
				`Wrong Format: Field "metric" (%s) is not well formed. NO information will be saved`,
				rcvMsg.Metric,
			),
		)
	}

	if ksid, ok := rcvMsg.Tags["ksid"]; !ok {
		return errValidation(`Wrong Format: Tag "ksid" is required. NO information will be saved`)
	} else if ksid == collector.settings.Cassandra.Keyspace {
		return errValidation(
			fmt.Sprintf(
				`Wrong Format: Keyspace "%s" can not be used. NO information will be saved`,
				collector.settings.Cassandra.Keyspace,
			),
		)
	} else {
		packet.KsID = ksid
	}

	if lt == 1 {
		return errValidation(`Wrong Format: At least one tag other than "ksid" is required. NO information will be saved`)
	}

	if lt == 2 {
		if _, ok := rcvMsg.Tags["ttl"]; ok {
			return errValidation(`Wrong Format: At least one tag other than "ksid" and "ttl" is required. NO information will be saved`)
		}
	}

	for k, v := range rcvMsg.Tags {
		if !collector.validKey.MatchString(k) {
			return errValidation(
				fmt.Sprintf(
					`Wrong Format: Tag key (%s) is not well formed. NO information will be saved`,
					k,
				),
			)
		}
		if !collector.validKey.MatchString(v) {
			return errValidation(
				fmt.Sprintf(
					`Wrong Format: Tag value (%s) is not well formed. NO information will be saved`,
					v,
				),
			)
		}
	}

	_, found, gerr := collector.boltc.GetKeyspace(packet.KsID)
	if !found {
		return errValidation(`Keyspace not found`)
	}
	if gerr != nil {
		return gerr
	}

	if rcvMsg.Timestamp == 0 {
		packet.Timestamp = time.Now().Unix()
	} else {
		t, err := gorilla.MilliToSeconds(rcvMsg.Timestamp)
		if gerr != nil {
			return errValidation(err.Error())
		}
		packet.Timestamp = t
	}

	packet.Number = number

	packet.Message = rcvMsg
	packet.ID = GenerateID(rcvMsg)
	if !number {
		packet.ID = fmt.Sprintf("T%v", packet.ID)
	}

	return nil
}
