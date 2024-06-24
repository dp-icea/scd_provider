package dss

type Dss struct {
	client Client
}

type PutOirRequest struct {
	Extents    []Volume4D                  `json:"extents"`
	FlightType OperationalIntentFlightType `json:"flight_type"`
	Priority   int32                       `json:"priority"`
}

type OperationalIntentDetails struct {
	Volumes           []Volume4D `json:"volumes"`
	OffNominalVolumes []Volume4D `json:"off_nominal_volumes"`
	Priority          int32      `json:"priority"`
}

type OperationalIntent struct {
	Reference OperationalIntentReference `json:"reference"`
	Details   OperationalIntentDetails   `json:"details"`
}

func (dss Dss) PutOperationalIntent(request PutOirRequest) (OperationalIntent, error) {
	notifyForConstraint := true
	parameters := PutOperationalIntentReferenceParameters{
		Extents:        request.Extents,
		Key:            nil,
		State:          "Accepted",
		UssBaseUrl:     "http://localhost:9091",
		SubscriptionId: nil,
		NewSubscription: &ImplicitSubscriptionParameters{
			UssBaseUrl:           "http://localhost:9091",
			NotifyForConstraints: &notifyForConstraint,
		},
		FlightType: request.FlightType,
	}

	//TODO: Deconfliction with Constraints and other OIRs

	putResponse, err := dss.client.PutOperationalIntent(parameters)
	if err != nil {
		return OperationalIntent{}, err
	}

	//TODO: Notify subscribers

	operationalIntent := OperationalIntent{
		Reference: putResponse.OperationalIntentReference,
		Details: OperationalIntentDetails{
			Volumes:           request.Extents,
			OffNominalVolumes: nil,
			Priority:          request.Priority,
		},
	}
	return operationalIntent, nil

}
