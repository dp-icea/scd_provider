package scd

import (
	"icea_uss/scd/dss"
	"icea_uss/store"
)

type Deconflictor struct {
}

func (d Deconflictor) Injection(request dss.PutOirRequest) (dss.OperationalIntent, error) {
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

func (d Deconflictor) GetOir(id string) (dss.OperationalIntent, error) {
	oir, err := store.GetOir(id)
	if err != nil {
		return dss.OperationalIntent{}, err
	}

	return oir, err
}
