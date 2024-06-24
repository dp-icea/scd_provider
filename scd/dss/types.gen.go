// This file is auto-generated; do not change as any changes will be overwritten
package dss

// String for flight type. Accepted values are "VLOS", "EVLOS" or "BVLOS"
type FlightType string

const (
	FlightType_VLOS  FlightType = "VLOS"
	FlightType_EVLOS FlightType = "EVLOS"
	FlightType_BVLOS FlightType = "BVLOS"
)

// String whose format matches a version-4 UUID according to RFC 4122.
type UUIDv4Format string

// Identifier for an Entity communicated through the DSS.  Formatted as a UUIDv4.
type EntityID UUIDv4Format

// A token associated with a particular UTM Entity+version created by the DSS upon creation or modification of an Entity reference and provided to the client creating or modifying the Entity reference.  The EntityOVN is stored privately by the DSS and then compared against entries in a Key provided to mutate the airspace.  The EntityOVN is also provided by the client whenever that client transmits the full information of the Entity (either via GET, or via a subscription notification).
type EntityOVN string

// Identifier for a subscription communicated through the DSS.  Formatted as a UUIDv4.
type SubscriptionID UUIDv4Format

// Proof that a client has obtained the latest airspace content, used to ensure that write operations to the DSS occur only when the latest content is known (i.e. has been read). The client is required to pass a current Key constructed from information obtained with previous read operations and subsequent requests for full information, and optionally, subscription updates for deconflicted write operations to the DSS.  The contents of this data structure are generated by the client.
type Key []EntityOVN

type Time struct {
	// RFC3339-formatted time/date string.  The time zone must be 'Z'.
	Value string `json:"value"`

	Format string `json:"format"`
}

type Radius struct {
	// Distance from the centerpoint of a circular area, along the WGS84 ellipsoid.
	Value float32 `json:"value"`

	// FIXM-compatible units.  Only meters ("M") are acceptable for UTM.
	Units string `json:"units"`
}

type Altitude struct {
	// The numeric value of the altitude. Note that min and max values are added as a sanity check. As use cases evolve and more options are made available in terms of units of measure or reference systems, these bounds may be re-evaluated.
	Value float64 `json:"value"`

	// A code indicating the reference for a vertical distance. See AIXM 5.1 and FIXM 4.2.0. Currently, UTM only allows WGS84 with no immediate plans to allow other options. FIXM and AIXM allow for 'SFC' which is equivalent to AGL.
	Reference string `json:"reference"`

	// The reference quantities used to express the value of altitude. See FIXM 4.2. Currently, UTM only allows meters with no immediate plans to allow other options.
	Units string `json:"units"`
}

// Degrees of latitude north of the equator, with reference to the WGS84 ellipsoid.
type Latitude float64

// Degrees of longitude east of the Prime Meridian, with reference to the WGS84 ellipsoid.
type Longitude float64

// An enclosed area on the earth. The bounding edges of this polygon are defined to be the shortest paths between connected vertices.  This means, for instance, that the edge between two points both defined at a particular latitude is not generally contained at that latitude. The winding order must be interpreted as the order which produces the smaller area. The path between two vertices is defined to be the shortest possible path between those vertices. Edges may not cross. Vertices may not be duplicated.  In particular, the final polygon vertex must not be identical to the first vertex.
type Polygon struct {
	Vertices []LatLngPoint `json:"vertices"`
}

// Point on the earth's surface.
type LatLngPoint struct {
	Lng Longitude `json:"lng"`

	Lat Latitude `json:"lat"`
}

// A circular area on the surface of the earth.
type Circle struct {
	Center *LatLngPoint `json:"center,omitempty"`

	Radius *Radius `json:"radius,omitempty"`
}

// A three-dimensional geographic volume consisting of a vertically-extruded shape. Exactly one outline must be specified.
type Volume3D struct {
	// A circular geographic shape on the surface of the earth.
	OutlineCircle *Circle `json:"outline_circle,omitempty"`

	// A polygonal geographic shape on the surface of the earth.
	OutlinePolygon *Polygon `json:"outline_polygon,omitempty"`

	// Minimum bounding altitude of this volume. Must be less than altitude_upper, if specified.
	AltitudeLower *Altitude `json:"altitude_lower,omitempty"`

	// Maximum bounding altitude of this volume. Must be greater than altitude_lower, if specified.
	AltitudeUpper *Altitude `json:"altitude_upper,omitempty"`
}

// Contiguous block of geographic spacetime.
type Volume4D struct {
	Volume Volume3D `json:"volume"`

	// Beginning time of this volume. Must be before time_end.
	TimeStart *Time `json:"time_start,omitempty"`

	// End time of this volume. Must be after time_start.
	TimeEnd *Time `json:"time_end,omitempty"`
}

// Human-readable string returned when an error occurs as a result of a USS - DSS transaction.
type ErrorResponse struct {
	// Human-readable message indicating what error occurred and/or why.
	Message *string `json:"message,omitempty"`
}

// State of subscription which is causing a notification to be sent.
type SubscriptionState struct {
	SubscriptionId SubscriptionID `json:"subscription_id"`

	NotificationIndex SubscriptionNotificationIndex `json:"notification_index"`
}

// Subscriber to notify of a change in the airspace.  This is provided by the DSS to a client changing the airspace, and it is the responsibility of that client to send a notification to the specified USS according to the change made to the airspace.
type SubscriberToNotify struct {
	// Subscription(s) prompting this notification.
	Subscriptions []SubscriptionState `json:"subscriptions"`

	UssBaseUrl SubscriptionUssBaseURL `json:"uss_base_url"`
}

// Specification of a geographic area that a client is interested in on an ongoing basis (e.g., "planning area").
type Subscription struct {
	Id SubscriptionID `json:"id"`

	// Version of the subscription that the DSS changes every time a USS changes the subscription.  The DSS incrementing the notification_index does not constitute a change that triggers a new version.  A USS must specify this version when modifying an existing subscription to ensure consistency in read-modify-write operations and distributed systems.
	Version string `json:"version"`

	NotificationIndex SubscriptionNotificationIndex `json:"notification_index"`

	// If set, this subscription will not receive notifications involving airspace changes entirely before this time.
	TimeStart *Time `json:"time_start,omitempty"`

	// If set, this subscription will not receive notifications involving airspace changes entirely after this time.
	TimeEnd *Time `json:"time_end,omitempty"`

	UssBaseUrl SubscriptionUssBaseURL `json:"uss_base_url"`

	// If true, trigger notifications when operational intents are created, updated, or deleted.  Otherwise, changes in operational intents should not trigger notifications.  The scope utm.strategic_coordination is required to set this flag true.
	NotifyForOperationalIntents *bool `json:"notify_for_operational_intents,omitempty"`

	// If true, trigger notifications when constraints are created, updated, or deleted.  Otherwise, changes in constraints should not trigger notifications.  The scope utm.constraint_processing is required to set this flag true.
	NotifyForConstraints *bool `json:"notify_for_constraints,omitempty"`

	// True if this subscription was implicitly created by the DSS via the creation of an operational intent, and should therefore be deleted by the DSS when that operational intent is deleted.
	ImplicitSubscription *bool `json:"implicit_subscription,omitempty"`

	// List of IDs for operational intents that are dependent on this subscription.
	DependentOperationalIntents *[]EntityID `json:"dependent_operational_intents,omitempty"`
}

// Tracks the notifications sent for a subscription so the subscriber can detect missed notifications more easily.
type SubscriptionNotificationIndex int32

// Parameters for a request to find subscriptions matching the provided criteria.
type QuerySubscriptionParameters struct {
	AreaOfInterest *Volume4D `json:"area_of_interest,omitempty"`
}

// Response to DSS query for subscriptions in a particular geographic area.
type QuerySubscriptionsResponse struct {
	// Subscriptions that overlap the specified geographic area.
	Subscriptions []Subscription `json:"subscriptions"`
}

// Response to DSS request for the subscription with the given id.
type GetSubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`
}

// Parameters for a request to create/update a subscription in the DSS.  At least one form of notifications must be requested.
type PutSubscriptionParameters struct {
	// Spacetime extents of the volume to subscribe to.
	// This subscription will automatically be deleted after its end time if it has not been refreshed by then. If end time is not specified, the value will be chosen automatically by the DSS. If start time is not specified, it will default to the time the request is processed. The end time may not be in the past.
	// Note that some Entities triggering notifications may lie entirely outside the requested area.
	Extents Volume4D `json:"extents"`

	UssBaseUrl SubscriptionUssBaseURL `json:"uss_base_url"`

	// If true, trigger notifications when operational intents are created, updated, or deleted.  Otherwise, changes in operational intents should not trigger notifications.  The scope utm.strategic_coordination is required to set this flag true.
	NotifyForOperationalIntents *bool `json:"notify_for_operational_intents,omitempty"`

	// If true, trigger notifications when constraints are created, updated, or deleted.  Otherwise, changes in constraints should not trigger notifications.  The scope utm.constraint_processing is required to set this flag true.
	NotifyForConstraints *bool `json:"notify_for_constraints,omitempty"`
}

// The base URL of a USS implementation of the parts of the USS-USS API necessary for receiving the notifications requested by this subscription.
type SubscriptionUssBaseURL UssBaseURL

// Response for a request to create or update a subscription.
type PutSubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`

	// Operational intents in or near the subscription area at the time of creation/update, if `notify_for_operational_intents` is true.
	OperationalIntentReferences *[]OperationalIntentReference `json:"operational_intent_references,omitempty"`

	// Constraints in or near the subscription area at the time of creation/update, if `notify_for_constraints` is true.
	ConstraintReferences *[]ConstraintReference `json:"constraint_references,omitempty"`
}

// Response for a successful request to delete a subscription.
type DeleteSubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`
}

// The base URL of a USS implementation of part or all of the USS-USS API. Per the USS-USS API, the full URL of a particular resource can be constructed by appending, e.g., `/uss/v1/{resource}/{id}` to this base URL. Accordingly, this URL may not have a trailing '/' character.
type UssBaseURL string

// State of an operational intent.
// 'Accepted': Operational intent is created and shared, but not yet in use; see standard text for more details.
// The create or update request for this operational intent reference must include a Key containing all OVNs for all relevant Entities.
// 'Activated': Operational intent is in active use; see standard text for more details.
// The create or update request for this operational intent reference must include a Key containing all OVNs for all relevant Entities.
// 'Nonconforming': UA is temporarily outside its volumes, but the situation is expected to be recoverable; see standard text for more details.
// In this state, the `/uss/v1/operational_intents/{entityid}/telemetry` USS-USS endpoint should respond, if available, to queries from USS peers.  The create or update request for this operational intent may omit a Key in this case because the operational intent is being adjusted as flown and cannot necessarily deconflict.
// 'Contingent': UA is considered unrecoverably unable to conform with its coordinate operational intent; see standard text for more details.
// This state must transition to Ended.  In this state, the `/uss/v1/operational_intents/{entityid}/telemetry` USS-USS endpoint should respond, if available, to queries from USS peers.  The create or update request for this operational intent may omit a Key in this case because the operational intent is being adjusted as flown and cannot necessarily deconflict.
type OperationalIntentState string

const (
	OperationalIntentState_Accepted      OperationalIntentState = "Accepted"
	OperationalIntentState_Activated     OperationalIntentState = "Activated"
	OperationalIntentState_Nonconforming OperationalIntentState = "Nonconforming"
	OperationalIntentState_Contingent    OperationalIntentState = "Contingent"
)

// The high-level information of a planned or active operational intent with the URL of a USS to query for details.  Note: 'ovn' is returned ONLY to the USS that created the operational intent but NEVER to other USS instances.
type OperationalIntentReference struct {
	Id EntityID `json:"id"`

	FlightType OperationalIntentFlightType `json:"flight_type"`

	// Created by the DSS based on creating client's ID (via access token).  Used internal to the DSS for restricting mutation and deletion operations to manager.  Used by USSs to reject operational intent update notifications originating from a USS that does not manage the operational intent.
	Manager string `json:"manager"`

	UssAvailability UssAvailabilityState `json:"uss_availability"`

	// Numeric version of this operational intent which increments upon each change in the operational intent, regardless of whether any field of the operational intent reference changes.  A USS with the details of this operational intent when it was at a particular version does not need to retrieve the details again until the version changes.
	Version int32 `json:"version"`

	State OperationalIntentState `json:"state"`

	// Opaque version number of this operational intent.  Populated only when the OperationalIntentReference is managed by the USS retrieving or providing it.  Not populated when the OperationalIntentReference is not managed by the USS retrieving or providing it (instead, the USS must obtain the OVN from the details retrieved from the managing USS).
	Ovn *EntityOVN `json:"ovn,omitempty"`

	// Beginning time of operational intent.
	TimeStart Time `json:"time_start"`

	// End time of operational intent.
	TimeEnd Time `json:"time_end"`

	UssBaseUrl OperationalIntentUssBaseURL `json:"uss_base_url"`

	// The ID of the subscription that is ensuring the operational intent manager receives relevant airspace updates.
	SubscriptionId SubscriptionID `json:"subscription_id"`
}

// Flight Type
type OperationalIntentFlightType FlightType

// The base URL of a USS implementation that implements the parts of the USS-USS API necessary for providing the details of this operational intent, and telemetry during non-conformance or contingency, if applicable.
type OperationalIntentUssBaseURL UssBaseURL

// Parameters for a request to create an OperationalIntentReference in the DSS. A subscription to changes overlapping this volume may be implicitly created, but this can be overridden by providing the (optional) 'subscription_id' to use. Note: The implicit subscription is managed by the DSS, not the USS.
type PutOperationalIntentReferenceParameters struct {
	// Spacetime extents that bound this operational intent.
	// Start and end times, as well as lower and upper altitudes, are required for each volume. The end time may not be in the past. All volumes, both nominal and off-nominal, must be encompassed in these extents. However, these extents do not need to match the precise volumes of the operational intent; a single bounding extent may be provided instead, for instance.
	Extents []Volume4D `json:"extents"`

	// Proof that the USS creating or mutating this operational intent was aware of the current state of the airspace, with the expectation that this operational intent is therefore deconflicted from all relevant features in the airspace.  This field is not required when declaring an operational intent Nonconforming or Contingent, or when there are no relevant Entities in the airspace, but is otherwise required. OVNs for constraints are required if and only if the USS managing this operational intent is performing the constraint processing role, which is indicated by whether the subscription associated with this operational intent triggers notifications for constraints.  The key does not need to contain the OVN for the operational intent being updated.
	Key *Key `json:"key,omitempty"`

	State OperationalIntentState `json:"state"`

	UssBaseUrl OperationalIntentUssBaseURL `json:"uss_base_url"`

	// The ID of an existing subscription that the USS will use to keep the operator informed about updates to relevant airspace information. If this field is not provided when the operational intent is in the Activated, Nonconforming, or Contingent state, then the `new_subscription` field must be provided in order to provide notification capability for the operational intent.  The subscription specified by this ID must cover at least the area over which this operational intent is conducted, and it must provide notifications for operational intents.
	SubscriptionId *EntityID `json:"subscription_id,omitempty"`

	// If an existing subscription is not specified in `subscription_id`, and the operational intent is in the Activated, Nonconforming, or Contingent state, then this field must be populated.  When this field is populated, an implicit subscription will be created and associated with this operational intent, and will generally be deleted automatically upon the deletion of this operational intent.
	NewSubscription *ImplicitSubscriptionParameters `json:"new_subscription,omitempty"`

	FlightType OperationalIntentFlightType `json:"flight_type"`
}

// Information necessary to create a subscription to serve a single operational intent's notification needs.
type ImplicitSubscriptionParameters struct {
	// The base URL of a USS implementation of the parts of the USS-USS API necessary for receiving the notifications that the operational intent must be aware of.  This includes, at least, notifications for relevant changes in operational intents.
	UssBaseUrl SubscriptionUssBaseURL `json:"uss_base_url"`

	// True if this operational intent's subscription should trigger notifications when constraints change. Otherwise, changes in constraints should not trigger notifications.  The scope utm.constraint_processing is required to set this flag true, and a USS performing the constraint processing role should set this flag true.
	NotifyForConstraints *bool `json:"notify_for_constraints,omitempty"`
}

// Response to DSS request for the OperationalIntentReference with the given ID.
type GetOperationalIntentReferenceResponse struct {
	OperationalIntentReference OperationalIntentReference `json:"operational_intent_reference"`
}

// Response to a request to create, update, or delete an OperationalIntentReference in the DSS.
type ChangeOperationalIntentReferenceResponse struct {
	// DSS subscribers that this client now has the obligation to notify of the operational intent changes just made.  This client must call POST for each provided URL according to the USS-USS `/uss/v1/operational_intents` path API.  The client's own subscriptions will also be included in this list.
	Subscribers []SubscriberToNotify `json:"subscribers"`

	OperationalIntentReference OperationalIntentReference `json:"operational_intent_reference"`
}

// Parameters for a request to find OperationalIntentReferences matching the provided criteria.
type QueryOperationalIntentReferenceParameters struct {
	AreaOfInterest *Volume4D `json:"area_of_interest,omitempty"`
}

// Response to DSS query for OperationalIntentReferences in an area of interest.
type QueryOperationalIntentReferenceResponse struct {
	// OperationalIntentReferences in the area of interest.
	OperationalIntentReferences []OperationalIntentReference `json:"operational_intent_references"`
}

// A ConstraintReference (area in which a constraint is present, along with other high-level information, but no details).  The DSS reports only these references and clients must exchange details and additional information peer-to-peer.
type ConstraintReference struct {
	Id EntityID `json:"id"`

	// Created by the DSS based on creating client's ID (via access token).  Used internal to the DSS for restricting mutation and deletion operations to manager.  Used by USSs to reject constraint update notifications originating from a USS that does not manage the constraint.
	Manager string `json:"manager"`

	UssAvailability UssAvailabilityState `json:"uss_availability"`

	// Numeric version of this constraint which increments upon each change in the constraint, regardless of whether any field of the constraint reference changes.  A USS with the details of this constraint when it was at a particular version does not need to retrieve the details again until the version changes.
	Version int32 `json:"version"`

	// Opaque version number of this constraint.  Populated only when the ConstraintReference is managed by the USS retrieving or providing it.  Not populated when the ConstraintReference is not managed by the USS retrieving or providing it (instead, the USS must obtain the OVN from the details retrieved from the managing USS).
	Ovn *EntityOVN `json:"ovn,omitempty"`

	TimeStart Time `json:"time_start"`

	TimeEnd Time `json:"time_end"`

	UssBaseUrl ConstraintUssBaseURL `json:"uss_base_url"`
}

// The base URL of a USS implementation that implements the parts of the USS-USS API necessary for providing the details of this constraint.
type ConstraintUssBaseURL UssBaseURL

// Parameters for a request to create/update a ConstraintReference in the DSS.
type PutConstraintReferenceParameters struct {
	// Spacetime extents that bound this constraint.
	// The end time may not be in the past.
	// All volumes of the constraint must be encompassed in these extents. However, these extents do not need to match the precise volumes of the constraint; a single bounding extent may be provided instead, for instance.
	Extents []Volume4D `json:"extents"`

	UssBaseUrl ConstraintUssBaseURL `json:"uss_base_url"`
}

// Response to DSS request for the ConstraintReference with the given ID.
type GetConstraintReferenceResponse struct {
	ConstraintReference ConstraintReference `json:"constraint_reference"`
}

// Response to a request to create, update, or delete a ConstraintReference. in the DSS.
type ChangeConstraintReferenceResponse struct {
	// DSS subscribers that this client now has the obligation to notify of the constraint changes just made.  This client must call POST for each provided URL according to the USS-USS `/uss/v1/constraints` path API.  The client's own subscriptions will also be included in this list.
	Subscribers []SubscriberToNotify `json:"subscribers"`

	ConstraintReference ConstraintReference `json:"constraint_reference"`
}

// Parameters for a request to find ConstraintReferences matching the provided criteria.
type QueryConstraintReferenceParameters struct {
	AreaOfInterest *Volume4D `json:"area_of_interest,omitempty"`
}

// Response to DSS query for ConstraintReferences in an area of interest.
type QueryConstraintReferencesResponse struct {
	// ConstraintReferences in the area of interest.
	ConstraintReferences []ConstraintReference `json:"constraint_references"`
}

// Data provided when an airspace conflict was encountered.
type AirspaceConflictResponse struct {
	// Human-readable message indicating what error occurred and/or why.
	Message *string `json:"message,omitempty"`

	// List of operational intent references for which current proof of knowledge was not provided.  If this field is present and contains elements, the calling USS should query the details URLs for these operational intents to obtain their details and correct OVNs.  The OVNs can be used to update the key, at which point the USS may retry this call.
	MissingOperationalIntents *[]OperationalIntentReference `json:"missing_operational_intents,omitempty"`

	// List of constraint references for which current proof of knowledge was not provided.  If this field is present and contains elements, the calling USS should query the details URLs for these constraints to obtain their details and correct OVNs.  The OVNs can be used to update the key, at which point the USS may retry this call.
	MissingConstraints *[]ConstraintReference `json:"missing_constraints,omitempty"`
}

type UssAvailabilityStatus struct {
	// Client ID (matching their `sub` in access tokens) of the USS to which this availability applies.
	Uss string `json:"uss"`

	Availability UssAvailabilityState `json:"availability"`
}

// A USS is presumed to be in the Unknown state absent indication otherwise by a USS with availability arbitration scope.  Upon determination via availability arbitration, a USS is Down when it does not respond appropriately, and a Down USS may not perform the following operations in the DSS:
// * Create an operational intent in the Accepted or Activated states
// * Modify an operational intent whose new or unchanged state is Accepted or Activated
// * Delete an operational intent
// A USS in the Unknown state possesses all the capabilities, within the DSS, of a USS in the Normal state.
type UssAvailabilityState string

const (
	UssAvailabilityState_Unknown UssAvailabilityState = "Unknown"
	UssAvailabilityState_Normal  UssAvailabilityState = "Normal"
	UssAvailabilityState_Down    UssAvailabilityState = "Down"
)

type SetUssAvailabilityStatusParameters struct {
	// Version of USS's availability to change, for consistent read-modify-write operations and consistent retry behavior.
	OldVersion string `json:"old_version"`

	Availability UssAvailabilityState `json:"availability"`
}

type UssAvailabilityStatusResponse struct {
	// Current version of USS's availability.  Used to change USS's availability.
	Version string `json:"version"`

	Status UssAvailabilityStatus `json:"status"`
}

// Details of a request/response data exchange.
type ExchangeRecord struct {
	// Full URL of request.
	Url string `json:"url"`

	// HTTP verb used by requestor (e.g., "PUT," "GET," etc.)
	Method string `json:"method"`

	// Set of headers associated with request or response. Requires 'Authorization:' field (at a minimum)
	Headers *[]string `json:"headers,omitempty"`

	// A coded value that indicates the role of the logging USS: 'Client' (initiating a request to a remote USS) or 'Server' (handling a request from a remote USS)
	RecorderRole string `json:"recorder_role"`

	// The time at which the request was sent/received.
	RequestTime Time `json:"request_time"`

	// Base64-encoded body content sent/received as a request.
	RequestBody *string `json:"request_body,omitempty"`

	// The time at which the response was sent/received.
	ResponseTime *Time `json:"response_time,omitempty"`

	// Base64-encoded body content sent/received in response to request.
	ResponseBody *string `json:"response_body,omitempty"`

	// HTTP response code sent/received in response to request.
	ResponseCode *int32 `json:"response_code,omitempty"`

	// 'Human-readable description of the problem with the exchange, if any.'
	Problem *string `json:"problem,omitempty"`
}

// A report informing a server of a communication problem.
type ErrorReport struct {
	// ID assigned by the server receiving the report.  Not populated when submitting a report.
	ReportId *string `json:"report_id,omitempty"`

	// The request (by this USS) and response associated with the error.
	Exchange ExchangeRecord `json:"exchange"`
}