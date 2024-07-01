package scd

import (
	"net/http"
	"scd_provider/config"
)

type TokenVerifier interface {
	Verify(token string, expectedScope AuthScope) (bool, error)
}

type JwtTokenVerifier struct {
}

type AuthScope string

const (
	StrategicCoordination AuthScope = "utm.strategic_coordination"
)

func (jwt JwtTokenVerifier) Verify(token string, expectedScope AuthScope) (bool, error) {
	conf := config.GetGlobalConfig()
	url := conf.TokenUrl + "?access_token=" + token + "&required_scope=" + string(expectedScope) + "&expected_audience=icea"
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	if res.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}
