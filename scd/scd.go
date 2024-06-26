package scd

import (
	"scd_provider/scd/dss"
	"scd_provider/store"
)

type StrategicDeconfliction interface {
	Inject(request dss.PutOirRequest) (dss.OperationalIntent, error)
	FetchOir(id string) (dss.OperationalIntent, error)
}

type InterussDeconflictor struct {
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
