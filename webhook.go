// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/believe-go/internal/apijson"
	"github.com/stainless-sdks/believe-go/internal/requestconfig"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Register a new webhook endpoint to receive event notifications.
//
// ## Event Types
//
// Available event types to subscribe to:
//
// - `match.completed` - Fired when a football match ends
// - `team_member.transferred` - Fired when a player/coach joins or leaves a team
//
// If no event types are specified, the webhook will receive all event types.
//
// ## Webhook Signatures
//
// All webhook deliveries include Standard Webhooks signature headers:
//
// - `webhook-id` - Unique message identifier
// - `webhook-timestamp` - Unix timestamp of when the webhook was sent
// - `webhook-signature` - HMAC-SHA256 signature in format `v1,{base64_signature}`
//
// Store the returned `secret` securely - you'll need it to verify webhook
// signatures.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific webhook endpoint.
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *RegisteredWebhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all registered webhook endpoints.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *[]RegisteredWebhook, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unregister a webhook endpoint. It will no longer receive events.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Trigger a webhook event and deliver it to all subscribed endpoints.
//
// This endpoint is useful for testing your webhook integration. It will:
//
// 1. Generate an event with the specified type and payload
// 2. Find all webhooks subscribed to that event type
// 3. Send a POST request to each webhook URL with signature headers
// 4. Return the delivery results
//
// ## Event Payload
//
// You can provide a custom payload, or leave it empty to use a sample payload.
//
// ## Webhook Signature Headers
//
// Each webhook delivery includes:
//
// - `webhook-id` - Unique event identifier (e.g., `evt_abc123...`)
// - `webhook-timestamp` - Unix timestamp
// - `webhook-signature` - HMAC-SHA256 signature (`v1,{base64}`)
//
// To verify signatures, compute:
//
// ```
// signature = HMAC-SHA256(
//
//	key = base64_decode(secret_without_prefix),
//	message = "{timestamp}.{raw_json_payload}"
//
// )
// ```
func (r *WebhookService) TriggerEvent(ctx context.Context, body WebhookTriggerEventParams, opts ...option.RequestOption) (res *WebhookTriggerEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks/trigger"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A registered webhook endpoint.
type RegisteredWebhook struct {
	// Unique webhook identifier
	ID string `json:"id,required"`
	// When the webhook was registered
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// List of event types this webhook is subscribed to
	//
	// Any of "match.completed", "team_member.transferred".
	EventTypes []string `json:"event_types,required"`
	// The secret key for verifying webhook signatures (base64 encoded)
	Secret string `json:"secret,required"`
	// The URL to send webhook events to
	URL string `json:"url,required" format:"uri"`
	// Optional description for this webhook
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		Secret      respjson.Field
		URL         respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegisteredWebhook) RawJSON() string { return r.JSON.raw }
func (r *RegisteredWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after registering a webhook.
type WebhookNewResponse struct {
	// The registered webhook details
	Webhook RegisteredWebhook `json:"webhook,required"`
	// Status message
	Message string `json:"message"`
	// Ted's reaction
	TedSays string `json:"ted_says"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhook     respjson.Field
		Message     respjson.Field
		TedSays     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponse map[string]any

// Response after triggering webhook events.
type WebhookTriggerEventResponse struct {
	// Results of webhook deliveries
	Deliveries []WebhookTriggerEventResponseDelivery `json:"deliveries,required"`
	// Unique event identifier
	EventID string `json:"event_id,required"`
	// The type of event triggered
	//
	// Any of "match.completed", "team_member.transferred".
	EventType WebhookTriggerEventResponseEventType `json:"event_type,required"`
	// Number of successful deliveries
	SuccessfulDeliveries int64 `json:"successful_deliveries,required"`
	// Ted's reaction
	TedSays string `json:"ted_says,required"`
	// Total number of webhooks that received this event
	TotalWebhooks int64 `json:"total_webhooks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries           respjson.Field
		EventID              respjson.Field
		EventType            respjson.Field
		SuccessfulDeliveries respjson.Field
		TedSays              respjson.Field
		TotalWebhooks        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTriggerEventResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTriggerEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of delivering a webhook to a single endpoint.
type WebhookTriggerEventResponseDelivery struct {
	// Whether delivery was successful
	Success bool `json:"success,required"`
	// URL the webhook was sent to
	URL string `json:"url,required"`
	// ID of the webhook
	WebhookID string `json:"webhook_id,required"`
	// Error message if delivery failed
	Error string `json:"error,nullable"`
	// HTTP status code from the endpoint
	StatusCode int64 `json:"status_code,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		URL         respjson.Field
		WebhookID   respjson.Field
		Error       respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTriggerEventResponseDelivery) RawJSON() string { return r.JSON.raw }
func (r *WebhookTriggerEventResponseDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event triggered
type WebhookTriggerEventResponseEventType string

const (
	WebhookTriggerEventResponseEventTypeMatchCompleted        WebhookTriggerEventResponseEventType = "match.completed"
	WebhookTriggerEventResponseEventTypeTeamMemberTransferred WebhookTriggerEventResponseEventType = "team_member.transferred"
)

type WebhookNewParams struct {
	// The URL to send webhook events to
	URL string `json:"url,required" format:"uri"`
	// Optional description for this webhook
	Description param.Opt[string] `json:"description,omitzero"`
	// List of event types to subscribe to. If not provided, subscribes to all events.
	//
	// Any of "match.completed", "team_member.transferred".
	EventTypes []string `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTriggerEventParams struct {
	// The type of event to trigger
	//
	// Any of "match.completed", "team_member.transferred".
	EventType WebhookTriggerEventParamsEventType `json:"event_type,omitzero,required"`
	// Optional event payload. If not provided, a sample payload will be generated.
	Payload WebhookTriggerEventParamsPayloadUnion `json:"payload,omitzero"`
	paramObj
}

func (r WebhookTriggerEventParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTriggerEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTriggerEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event to trigger
type WebhookTriggerEventParamsEventType string

const (
	WebhookTriggerEventParamsEventTypeMatchCompleted        WebhookTriggerEventParamsEventType = "match.completed"
	WebhookTriggerEventParamsEventTypeTeamMemberTransferred WebhookTriggerEventParamsEventType = "team_member.transferred"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookTriggerEventParamsPayloadUnion struct {
	OfMatchCompleted        *WebhookTriggerEventParamsPayloadMatchCompleted        `json:",omitzero,inline"`
	OfTeamMemberTransferred *WebhookTriggerEventParamsPayloadTeamMemberTransferred `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookTriggerEventParamsPayloadUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMatchCompleted, u.OfTeamMemberTransferred)
}
func (u *WebhookTriggerEventParamsPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WebhookTriggerEventParamsPayloadUnion](
		"event_type",
		apijson.Discriminator[WebhookTriggerEventParamsPayloadMatchCompleted]("match.completed"),
		apijson.Discriminator[WebhookTriggerEventParamsPayloadTeamMemberTransferred]("team_member.transferred"),
	)
}

// Payload for match.completed event.
//
// The property Data is required.
type WebhookTriggerEventParamsPayloadMatchCompleted struct {
	// Event data
	Data WebhookTriggerEventParamsPayloadMatchCompletedData `json:"data,omitzero,required"`
	// The type of webhook event
	//
	// Any of "match.completed".
	EventType string `json:"event_type,omitzero"`
	paramObj
}

func (r WebhookTriggerEventParamsPayloadMatchCompleted) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTriggerEventParamsPayloadMatchCompleted
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTriggerEventParamsPayloadMatchCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadMatchCompleted](
		"event_type", "match.completed",
	)
}

// Event data
//
// The properties AwayScore, AwayTeamID, CompletedAt, HomeScore, HomeTeamID,
// MatchID, MatchType, Result, TedPostMatchQuote are required.
type WebhookTriggerEventParamsPayloadMatchCompletedData struct {
	// Final away team score
	AwayScore int64 `json:"away_score,required"`
	// Away team ID
	AwayTeamID string `json:"away_team_id,required"`
	// When the match completed
	CompletedAt time.Time `json:"completed_at,required" format:"date-time"`
	// Final home team score
	HomeScore int64 `json:"home_score,required"`
	// Home team ID
	HomeTeamID string `json:"home_team_id,required"`
	// Unique match identifier
	MatchID string `json:"match_id,required"`
	// Type of match
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType string `json:"match_type,omitzero,required"`
	// Match result from home team perspective
	//
	// Any of "home_win", "away_win", "draw".
	Result string `json:"result,omitzero,required"`
	// Ted's post-match wisdom
	TedPostMatchQuote string `json:"ted_post_match_quote,required"`
	// Ted's lesson from the match
	LessonLearned param.Opt[string] `json:"lesson_learned,omitzero"`
	// Player of the match (if awarded)
	ManOfTheMatch param.Opt[string] `json:"man_of_the_match,omitzero"`
	paramObj
}

func (r WebhookTriggerEventParamsPayloadMatchCompletedData) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTriggerEventParamsPayloadMatchCompletedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTriggerEventParamsPayloadMatchCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadMatchCompletedData](
		"match_type", "league", "cup", "friendly", "playoff", "final",
	)
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadMatchCompletedData](
		"result", "home_win", "away_win", "draw",
	)
}

// Payload for team_member.transferred event.
//
// The property Data is required.
type WebhookTriggerEventParamsPayloadTeamMemberTransferred struct {
	// Event data
	Data WebhookTriggerEventParamsPayloadTeamMemberTransferredData `json:"data,omitzero,required"`
	// The type of webhook event
	//
	// Any of "team_member.transferred".
	EventType string `json:"event_type,omitzero"`
	paramObj
}

func (r WebhookTriggerEventParamsPayloadTeamMemberTransferred) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTriggerEventParamsPayloadTeamMemberTransferred
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTriggerEventParamsPayloadTeamMemberTransferred) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadTeamMemberTransferred](
		"event_type", "team_member.transferred",
	)
}

// Event data
//
// The properties CharacterID, CharacterName, MemberType, TeamID, TeamMemberID,
// TeamName, TedReaction, TransferType are required.
type WebhookTriggerEventParamsPayloadTeamMemberTransferredData struct {
	// ID of the character (links to /characters)
	CharacterID string `json:"character_id,required"`
	// Name of the character
	CharacterName string `json:"character_name,required"`
	// Type of team member
	//
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type,omitzero,required"`
	// ID of the team involved
	TeamID string `json:"team_id,required"`
	// ID of the team member
	TeamMemberID string `json:"team_member_id,required"`
	// Name of the team involved
	TeamName string `json:"team_name,required"`
	// Ted's reaction to the transfer
	TedReaction string `json:"ted_reaction,required"`
	// Whether the member joined or departed
	//
	// Any of "joined", "departed".
	TransferType string `json:"transfer_type,omitzero,required"`
	// Previous team ID (for joins from another team)
	PreviousTeamID param.Opt[string] `json:"previous_team_id,omitzero"`
	// Previous team name (for joins from another team)
	PreviousTeamName param.Opt[string] `json:"previous_team_name,omitzero"`
	// Transfer fee in GBP (for players)
	TransferFeeGbp param.Opt[string] `json:"transfer_fee_gbp,omitzero"`
	// Years spent with previous team
	YearsWithPreviousTeam param.Opt[int64] `json:"years_with_previous_team,omitzero"`
	paramObj
}

func (r WebhookTriggerEventParamsPayloadTeamMemberTransferredData) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTriggerEventParamsPayloadTeamMemberTransferredData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTriggerEventParamsPayloadTeamMemberTransferredData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadTeamMemberTransferredData](
		"member_type", "player", "coach", "medical_staff", "equipment_manager",
	)
	apijson.RegisterFieldValidator[WebhookTriggerEventParamsPayloadTeamMemberTransferredData](
		"transfer_type", "joined", "departed",
	)
}
