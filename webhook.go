// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
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

func (r *WebhookService) Unwrap(payload []byte, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	res := &UnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// A registered webhook endpoint.
type RegisteredWebhook struct {
	// Unique webhook identifier
	ID string `json:"id" api:"required"`
	// When the webhook was registered
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// List of event types this webhook is subscribed to
	//
	// Any of "match.completed", "team_member.transferred".
	EventTypes []string `json:"event_types" api:"required"`
	// The secret key for verifying webhook signatures (base64 encoded)
	Secret string `json:"secret" api:"required"`
	// The URL to send webhook events to
	URL string `json:"url" api:"required" format:"uri"`
	// Optional description for this webhook
	Description string `json:"description" api:"nullable"`
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
	Webhook RegisteredWebhook `json:"webhook" api:"required"`
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
	Deliveries []WebhookTriggerEventResponseDelivery `json:"deliveries" api:"required"`
	// Unique event identifier
	EventID string `json:"event_id" api:"required"`
	// The type of event triggered
	//
	// Any of "match.completed", "team_member.transferred".
	EventType WebhookTriggerEventResponseEventType `json:"event_type" api:"required"`
	// Number of successful deliveries
	SuccessfulDeliveries int64 `json:"successful_deliveries" api:"required"`
	// Ted's reaction
	TedSays string `json:"ted_says" api:"required"`
	// Total number of webhooks that received this event
	TotalWebhooks int64 `json:"total_webhooks" api:"required"`
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
	Success bool `json:"success" api:"required"`
	// URL the webhook was sent to
	URL string `json:"url" api:"required"`
	// ID of the webhook
	WebhookID string `json:"webhook_id" api:"required"`
	// Error message if delivery failed
	Error string `json:"error" api:"nullable"`
	// HTTP status code from the endpoint
	StatusCode int64 `json:"status_code" api:"nullable"`
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

// Webhook event sent when a match completes.
type MatchCompletedWebhookEvent struct {
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Data payload for a match completed event.
	Data MatchCompletedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// The type of webhook event
	//
	// Any of "match.completed".
	EventType MatchCompletedWebhookEventEventType `json:"event_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		EventID     respjson.Field
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatchCompletedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MatchCompletedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data payload for a match completed event.
type MatchCompletedWebhookEventData struct {
	// Final away team score
	AwayScore int64 `json:"away_score" api:"required"`
	// Away team ID
	AwayTeamID string `json:"away_team_id" api:"required"`
	// When the match completed
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Final home team score
	HomeScore int64 `json:"home_score" api:"required"`
	// Home team ID
	HomeTeamID string `json:"home_team_id" api:"required"`
	// Unique match identifier
	MatchID string `json:"match_id" api:"required"`
	// Type of match
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType string `json:"match_type" api:"required"`
	// Match result from home team perspective
	//
	// Any of "home_win", "away_win", "draw".
	Result string `json:"result" api:"required"`
	// Ted's post-match wisdom
	TedPostMatchQuote string `json:"ted_post_match_quote" api:"required"`
	// Ted's lesson from the match
	LessonLearned string `json:"lesson_learned" api:"nullable"`
	// Player of the match (if awarded)
	ManOfTheMatch string `json:"man_of_the_match" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwayScore         respjson.Field
		AwayTeamID        respjson.Field
		CompletedAt       respjson.Field
		HomeScore         respjson.Field
		HomeTeamID        respjson.Field
		MatchID           respjson.Field
		MatchType         respjson.Field
		Result            respjson.Field
		TedPostMatchQuote respjson.Field
		LessonLearned     respjson.Field
		ManOfTheMatch     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatchCompletedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MatchCompletedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event
type MatchCompletedWebhookEventEventType string

const (
	MatchCompletedWebhookEventEventTypeMatchCompleted MatchCompletedWebhookEventEventType = "match.completed"
)

// Webhook event sent when a team member (player, coach, staff) transfers between
// teams.
type TeamMemberTransferredWebhookEvent struct {
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Data payload for a team member transfer event.
	Data TeamMemberTransferredWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// The type of webhook event
	//
	// Any of "team_member.transferred".
	EventType TeamMemberTransferredWebhookEventEventType `json:"event_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		EventID     respjson.Field
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamMemberTransferredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *TeamMemberTransferredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data payload for a team member transfer event.
type TeamMemberTransferredWebhookEventData struct {
	// ID of the character (links to /characters)
	CharacterID string `json:"character_id" api:"required"`
	// Name of the character
	CharacterName string `json:"character_name" api:"required"`
	// Type of team member
	//
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type" api:"required"`
	// ID of the team involved
	TeamID string `json:"team_id" api:"required"`
	// ID of the team member
	TeamMemberID string `json:"team_member_id" api:"required"`
	// Name of the team involved
	TeamName string `json:"team_name" api:"required"`
	// Ted's reaction to the transfer
	TedReaction string `json:"ted_reaction" api:"required"`
	// Whether the member joined or departed
	//
	// Any of "joined", "departed".
	TransferType string `json:"transfer_type" api:"required"`
	// Previous team ID (for joins from another team)
	PreviousTeamID string `json:"previous_team_id" api:"nullable"`
	// Previous team name (for joins from another team)
	PreviousTeamName string `json:"previous_team_name" api:"nullable"`
	// Transfer fee in GBP (for players)
	TransferFeeGbp string `json:"transfer_fee_gbp" api:"nullable"`
	// Years spent with previous team
	YearsWithPreviousTeam int64 `json:"years_with_previous_team" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterID           respjson.Field
		CharacterName         respjson.Field
		MemberType            respjson.Field
		TeamID                respjson.Field
		TeamMemberID          respjson.Field
		TeamName              respjson.Field
		TedReaction           respjson.Field
		TransferType          respjson.Field
		PreviousTeamID        respjson.Field
		PreviousTeamName      respjson.Field
		TransferFeeGbp        respjson.Field
		YearsWithPreviousTeam respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamMemberTransferredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *TeamMemberTransferredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event
type TeamMemberTransferredWebhookEventEventType string

const (
	TeamMemberTransferredWebhookEventEventTypeTeamMemberTransferred TeamMemberTransferredWebhookEventEventType = "team_member.transferred"
)

// UnwrapWebhookEventUnion contains all possible properties and values from
// [MatchCompletedWebhookEvent], [TeamMemberTransferredWebhookEvent].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [MatchCompletedWebhookEventData],
	// [TeamMemberTransferredWebhookEventData]
	Data    UnwrapWebhookEventUnionData `json:"data"`
	EventID string                      `json:"event_id"`
	// Any of "match.completed", "team_member.transferred".
	EventType string `json:"event_type"`
	JSON      struct {
		CreatedAt respjson.Field
		Data      respjson.Field
		EventID   respjson.Field
		EventType respjson.Field
		raw       string
	} `json:"-"`
}

// anyUnwrapWebhookEvent is implemented by each variant of
// [UnwrapWebhookEventUnion] to add type safety for the return type of
// [UnwrapWebhookEventUnion.AsAny]
type anyUnwrapWebhookEvent interface {
	implUnwrapWebhookEventUnion()
}

func (MatchCompletedWebhookEvent) implUnwrapWebhookEventUnion()        {}
func (TeamMemberTransferredWebhookEvent) implUnwrapWebhookEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnwrapWebhookEventUnion.AsAny().(type) {
//	case believe.MatchCompletedWebhookEvent:
//	case believe.TeamMemberTransferredWebhookEvent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnwrapWebhookEventUnion) AsAny() anyUnwrapWebhookEvent {
	switch u.EventType {
	case "match.completed":
		return u.AsMatchCompleted()
	case "team_member.transferred":
		return u.AsTeamMemberTransferred()
	}
	return nil
}

func (u UnwrapWebhookEventUnion) AsMatchCompleted() (v MatchCompletedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsTeamMemberTransferred() (v TeamMemberTransferredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	// This field is from variant [MatchCompletedWebhookEventData].
	AwayScore int64 `json:"away_score"`
	// This field is from variant [MatchCompletedWebhookEventData].
	AwayTeamID string `json:"away_team_id"`
	// This field is from variant [MatchCompletedWebhookEventData].
	CompletedAt time.Time `json:"completed_at"`
	// This field is from variant [MatchCompletedWebhookEventData].
	HomeScore int64 `json:"home_score"`
	// This field is from variant [MatchCompletedWebhookEventData].
	HomeTeamID string `json:"home_team_id"`
	// This field is from variant [MatchCompletedWebhookEventData].
	MatchID string `json:"match_id"`
	// This field is from variant [MatchCompletedWebhookEventData].
	MatchType string `json:"match_type"`
	// This field is from variant [MatchCompletedWebhookEventData].
	Result string `json:"result"`
	// This field is from variant [MatchCompletedWebhookEventData].
	TedPostMatchQuote string `json:"ted_post_match_quote"`
	// This field is from variant [MatchCompletedWebhookEventData].
	LessonLearned string `json:"lesson_learned"`
	// This field is from variant [MatchCompletedWebhookEventData].
	ManOfTheMatch string `json:"man_of_the_match"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	CharacterID string `json:"character_id"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	CharacterName string `json:"character_name"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	MemberType string `json:"member_type"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TeamID string `json:"team_id"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TeamMemberID string `json:"team_member_id"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TeamName string `json:"team_name"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TedReaction string `json:"ted_reaction"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TransferType string `json:"transfer_type"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	PreviousTeamID string `json:"previous_team_id"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	PreviousTeamName string `json:"previous_team_name"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	TransferFeeGbp string `json:"transfer_fee_gbp"`
	// This field is from variant [TeamMemberTransferredWebhookEventData].
	YearsWithPreviousTeam int64 `json:"years_with_previous_team"`
	JSON                  struct {
		AwayScore             respjson.Field
		AwayTeamID            respjson.Field
		CompletedAt           respjson.Field
		HomeScore             respjson.Field
		HomeTeamID            respjson.Field
		MatchID               respjson.Field
		MatchType             respjson.Field
		Result                respjson.Field
		TedPostMatchQuote     respjson.Field
		LessonLearned         respjson.Field
		ManOfTheMatch         respjson.Field
		CharacterID           respjson.Field
		CharacterName         respjson.Field
		MemberType            respjson.Field
		TeamID                respjson.Field
		TeamMemberID          respjson.Field
		TeamName              respjson.Field
		TedReaction           respjson.Field
		TransferType          respjson.Field
		PreviousTeamID        respjson.Field
		PreviousTeamName      respjson.Field
		TransferFeeGbp        respjson.Field
		YearsWithPreviousTeam respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// The URL to send webhook events to
	URL string `json:"url" api:"required" format:"uri"`
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
	EventType WebhookTriggerEventParamsEventType `json:"event_type,omitzero" api:"required"`
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
	Data WebhookTriggerEventParamsPayloadMatchCompletedData `json:"data,omitzero" api:"required"`
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
	AwayScore int64 `json:"away_score" api:"required"`
	// Away team ID
	AwayTeamID string `json:"away_team_id" api:"required"`
	// When the match completed
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Final home team score
	HomeScore int64 `json:"home_score" api:"required"`
	// Home team ID
	HomeTeamID string `json:"home_team_id" api:"required"`
	// Unique match identifier
	MatchID string `json:"match_id" api:"required"`
	// Type of match
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType string `json:"match_type,omitzero" api:"required"`
	// Match result from home team perspective
	//
	// Any of "home_win", "away_win", "draw".
	Result string `json:"result,omitzero" api:"required"`
	// Ted's post-match wisdom
	TedPostMatchQuote string `json:"ted_post_match_quote" api:"required"`
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
	Data WebhookTriggerEventParamsPayloadTeamMemberTransferredData `json:"data,omitzero" api:"required"`
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
	CharacterID string `json:"character_id" api:"required"`
	// Name of the character
	CharacterName string `json:"character_name" api:"required"`
	// Type of team member
	//
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type,omitzero" api:"required"`
	// ID of the team involved
	TeamID string `json:"team_id" api:"required"`
	// ID of the team member
	TeamMemberID string `json:"team_member_id" api:"required"`
	// Name of the team involved
	TeamName string `json:"team_name" api:"required"`
	// Ted's reaction to the transfer
	TedReaction string `json:"ted_reaction" api:"required"`
	// Whether the member joined or departed
	//
	// Any of "joined", "departed".
	TransferType string `json:"transfer_type,omitzero" api:"required"`
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
