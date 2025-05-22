package dss

import "scd_provider/config"

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
	conf := config.GetGlobalConfig()
	parameters := PutOperationalIntentReferenceParameters{
		Extents:        request.Extents,
		Key:            nil,
		State:          "Accepted",
		UssBaseUrl:     OperationalIntentUssBaseURL(conf.HostUrl),
		SubscriptionId: nil,
		NewSubscription: &ImplicitSubscriptionParameters{
			UssBaseUrl:           SubscriptionUssBaseURL(conf.HostUrl),
			NotifyForConstraints: &notifyForConstraint,
		},
		FlightType: request.FlightType,
	}

	queryConstraint, err := dss.client.QueryConstraints(parameters)
	if err != nil {
		return OperationalIntent{}, err
	}

	err = dss.handleConstraintQuery(parameters, queryConstraint)
	if err != nil {
		return OperationalIntent{}, err
	}

	queryOperationalIntent, err := dss.client.QueryOperationalIntent(parameters)
	if err != nil {
		return OperationalIntent{}, err
	}
	err = dss.handleOirQuery(parameters, queryOperationalIntent)
	if err != nil {
		return OperationalIntent{}, err
	}

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

func (dss Dss) handleConstraintQuery(param PutOperationalIntentReferenceParameters, query QueryConstraintReferencesResponse) error {
	g := GeomHelper{}
	for _, reference := range query.ConstraintReferences {
		constraint, err := dss.client.GetConstraintDetails(reference)
		if err != nil {
			return err
		}

		err = g.Intersects(param.Extents, constraint.Constraint.Details.Volumes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dss Dss) handleOirQuery(param PutOperationalIntentReferenceParameters, query QueryOperationalIntentReferenceResponse) error {
	g := GeomHelper{}
	for _, reference := range query.OperationalIntentReferences {
		oir, err := dss.client.GetOirDetails(reference)
		if err != nil {
			return err
		}

		err = g.Intersects(param.Extents, oir.OperationalIntent.Details.Volumes)
		if err != nil {
			return err
		}
	}
	return nil
}
