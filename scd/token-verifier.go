package scd

import (
	"io"
	"log"
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
	GetVersion            AuthScope = "interuss.versioning.read_system_versions"
)

func (jwt JwtTokenVerifier) Verify(token string, expectedScope AuthScope) (bool, error) {
	conf := config.GetGlobalConfig()
	url := conf.TokenUrl + "?access_token=" + token + "&required_scope=" + string(expectedScope) + "&expected_audience=icea"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer res.Body.Close()
	
	if res.StatusCode == http.StatusOK {
		return true, nil
	}
	
	// Read the response body to log the actual content
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Token verification failed with status code:", res.StatusCode)
		log.Println("Failed to read response body:", err)
	} else {
		log.Println("Token verification failed with status code:", res.StatusCode)
		log.Println("Token verification response:", string(body))
	}
	return false, nil
}
