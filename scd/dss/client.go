package dss

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/antonholmquist/jason"
	"github.com/google/uuid"
	"icea_uss/config"
	"io"
	"log"
	"net/http"
)

type Client struct {
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (c Client) auth(scope string, audience string) (string, error) {
	conf := config.GetGlobalConfig()

	url := conf.AuthUrl + "?" +
		"apikey=" + conf.ApiKey +
		"&intended_audience=" + audience +
		"&scope=" + scope +
		"&grant_type=client_credentials"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := jason.NewObjectFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	token, err := body.GetString("access_token")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c Client) PutOperationalIntent(param PutOperationalIntentReferenceParameters) (ChangeOperationalIntentReferenceResponse, error) {

	var response ChangeOperationalIntentReferenceResponse
	conf := config.GetGlobalConfig()
	id := uuid.New()

	token, err := c.auth("utm.strategic_coordination", "localhost")
	if err != nil {
		return response, err
	}

	parametersJSON, err := json.Marshal(param)
	if err != nil {
		return response, err
	}
	req, err := http.NewRequest(http.MethodPut,
		conf.DssUrl+"/dss/v1/operational_intent_references/"+id.String(),
		bytes.NewBuffer(parametersJSON))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		//TODO: Log when DSS refuses the creation of OIR
		return response, errors.New("failed to create OIR")
	}

	body, err := io.ReadAll(resp.Body)
	print(body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
