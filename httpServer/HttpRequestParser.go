package httpServer

import (
	"encoding/json"
	"icea_uss/scd/dss"
	"io"
	"net/http"
)

type HttpRequestParser struct{}

func (p HttpRequestParser) ParseInjection(r *http.Request) (dss.PutOirRequest, error) {
	var request dss.PutOirRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return request, err
	}

	print(request.FlightType)

	return request, nil

}
