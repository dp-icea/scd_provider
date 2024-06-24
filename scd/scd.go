package scd

import "icea_uss/scd/dss"

type Deconflictor struct {
}

func (d Deconflictor) Injection(request dss.PutOirRequest) (dss.OperationalIntent, error) {
	dssHandler := dss.Dss{}
	operationalIntent, err := dssHandler.PutOperationalIntent(request)
	if err != nil {
		return dss.OperationalIntent{}, err
	}
	return operationalIntent, nil
}
