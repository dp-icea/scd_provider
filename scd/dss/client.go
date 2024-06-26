package dss

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/antonholmquist/jason"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"scd_provider/config"
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

	if resp.StatusCode == http.StatusForbidden {
		log.Println("Authentication Failed")
		log.Println(body)
		return "", errors.New("authentication failed")
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

func (c Client) QueryOperationalIntent(param PutOperationalIntentReferenceParameters) (QueryOperationalIntentReferenceResponse, error) {
	conf := config.GetGlobalConfig()
	fullResponse := QueryOperationalIntentReferenceResponse{
		OperationalIntentReferences: make([]OperationalIntentReference, 0),
	}

	token, err := c.auth("utm.strategic_coordination", "localhost")
	if err != nil {
		return fullResponse, err
	}

	for _, volume := range param.Extents {
		body := QueryOperationalIntentReferenceParameters{
			AreaOfInterest: &volume,
		}

		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return fullResponse, err
		}

		url := conf.DssUrl + "/dss/v1/operational_intent_references/query"

		req, err := http.NewRequest(http.MethodPost,
			url,
			bytes.NewBuffer(bodyJSON))

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fullResponse, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Failed to query OIR")
			return fullResponse, errors.New("failed to query OIR")
		}

		respBody, err := io.ReadAll(resp.Body)

		var response QueryOperationalIntentReferenceResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return fullResponse, err
		}

		fullResponse.OperationalIntentReferences = append(fullResponse.OperationalIntentReferences, response.OperationalIntentReferences...)
	}

	return fullResponse, nil
}

func (c Client) QueryConstraints(param PutOperationalIntentReferenceParameters) (QueryConstraintReferencesResponse, error) {
	conf := config.GetGlobalConfig()
	fullResponse := QueryConstraintReferencesResponse{
		ConstraintReferences: make([]ConstraintReference, 0),
	}

	token, err := c.auth("utm.constraint_management", "localhost")
	if err != nil {
		return fullResponse, err
	}

	for _, volume := range param.Extents {
		body := QueryConstraintReferenceParameters{
			AreaOfInterest: &volume,
		}

		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return fullResponse, err
		}

		url := conf.DssUrl + "/dss/v1/constraint_references/query"

		req, err := http.NewRequest(http.MethodPost,
			url,
			bytes.NewBuffer(bodyJSON))

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fullResponse, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Failed to query Constraint")
			return fullResponse, errors.New("failed to query Constraint")
		}

		respBody, err := io.ReadAll(resp.Body)

		var response QueryConstraintReferencesResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return fullResponse, err
		}

		fullResponse.ConstraintReferences = append(fullResponse.ConstraintReferences, response.ConstraintReferences...)
	}

	return fullResponse, nil
}
