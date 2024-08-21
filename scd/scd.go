package scd

import (
	"scd_provider/scd/dss"
	"scd_provider/store"
)

type StrategicDeconfliction interface {
	Inject(request dss.PutOirRequest) (dss.OperationalIntent, error)
	FetchOir(id string) (dss.OperationalIntent, error)
	FetchVersion(id string) Version
	FetchLog(timestamp_start string, timestamp_end string) Log
}

type InterussDeconflictor struct {
}

type Version struct {
	SystemId string `json:"system_id"`
	Version  string `json:"version"`
}

type Log struct {
	Content string `json:"content"`
}

func (d InterussDeconflictor) Inject(request dss.PutOirRequest) (dss.OperationalIntent, error) {
	dssHandler := dss.Dss{}
	operationalIntent, err := dssHandler.PutOperationalIntent(request)
	if err != nil {
		return dss.OperationalIntent{}, err
	}

	err = store.CreateOir(operationalIntent)
	if err != nil {
		return dss.OperationalIntent{}, err
	}

	return operationalIntent, nil
}

func (d InterussDeconflictor) FetchOir(id string) (dss.OperationalIntent, error) {
	oir, err := store.GetOir(id)
	if err != nil {
		return dss.OperationalIntent{}, err
	}
	return oir, err
}

func (d InterussDeconflictor) FetchVersion(id string) Version {
	return Version{id, "1.0"}
}

func (d InterussDeconflictor) FetchLog(timestamp_start string, timestamp_end string) Log {
	return Log{"Lorem ipsum dolor sit amet"}
}
