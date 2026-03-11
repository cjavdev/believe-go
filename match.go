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
	"github.com/cjavdev/believe-go/internal/apiquery"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/pagination"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// MatchService contains methods and other services that help with interacting with
// the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatchService] method instead.
type MatchService struct {
	Options []option.RequestOption
	// Server-Sent Events (SSE) streaming endpoints
	Commentary MatchCommentaryService
}

// NewMatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMatchService(opts ...option.RequestOption) (r MatchService) {
	r = MatchService{}
	r.Options = opts
	r.Commentary = NewMatchCommentaryService(opts...)
	return
}

// Schedule a new match.
func (r *MatchService) New(ctx context.Context, body MatchNewParams, opts ...option.RequestOption) (res *Match, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "matches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific match.
func (r *MatchService) Get(ctx context.Context, matchID string, opts ...option.RequestOption) (res *Match, err error) {
	opts = slices.Concat(r.Options, opts)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matches/%s", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update specific fields of an existing match (e.g., update score).
func (r *MatchService) Update(ctx context.Context, matchID string, body MatchUpdateParams, opts ...option.RequestOption) (res *Match, err error) {
	opts = slices.Concat(r.Options, opts)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matches/%s", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a paginated list of all matches with optional filtering.
func (r *MatchService) List(ctx context.Context, query MatchListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Match], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "matches"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a paginated list of all matches with optional filtering.
func (r *MatchService) ListAutoPaging(ctx context.Context, query MatchListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Match] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a match from the database.
func (r *MatchService) Delete(ctx context.Context, matchID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return err
	}
	path := fmt.Sprintf("matches/%s", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get the life lesson learned from a specific match.
func (r *MatchService) GetLesson(ctx context.Context, matchID string, opts ...option.RequestOption) (res *MatchGetLessonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matches/%s/lesson", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get all turning points from a specific match.
func (r *MatchService) GetTurningPoints(ctx context.Context, matchID string, opts ...option.RequestOption) (res *[]MatchGetTurningPointsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matches/%s/turning-points", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// WebSocket endpoint for real-time live match simulation.
//
// Connect to receive a stream of match events as they happen in a simulated
// football match.
//
// ## Connection
//
// Connect via WebSocket with optional query parameters to customize the
// simulation.
//
// ## Example WebSocket URL
//
// ```
// ws://localhost:8000/matches/live?home_team=AFC%20Richmond&away_team=Manchester%20City&speed=2.0&excitement_level=7
// ```
//
// ## Server Messages
//
// The server sends JSON messages with these types:
//
// - `match_start` - When the match begins
// - `match_event` - For each match event (goals, fouls, cards, etc.)
// - `match_end` - When the match concludes
// - `error` - If an error occurs
// - `pong` - Response to client ping
//
// ## Client Messages
//
// Send JSON to control the simulation:
//
// - `{"action": "ping"}` - Keep-alive, server responds with `{"type": "pong"}`
// - `{"action": "pause"}` - Pause the simulation
// - `{"action": "resume"}` - Resume a paused simulation
// - `{"action": "set_speed", "speed": 2.0}` - Change playback speed (0.1-10.0)
// - `{"action": "get_status"}` - Request current match status
func (r *MatchService) StreamLive(ctx context.Context, query MatchStreamLiveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matches/live"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Full match model with ID.
type Match struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Away team ID
	AwayTeamID string `json:"away_team_id" api:"required"`
	// Match date and time
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Home team ID
	HomeTeamID string `json:"home_team_id" api:"required"`
	// Type of match
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType MatchType `json:"match_type" api:"required"`
	// Match attendance
	Attendance int64 `json:"attendance" api:"nullable"`
	// Away team score
	AwayScore int64 `json:"away_score"`
	// Episode ID where this match is featured
	EpisodeID string `json:"episode_id" api:"nullable"`
	// Home team score
	HomeScore int64 `json:"home_score"`
	// The life lesson learned from this match
	LessonLearned string `json:"lesson_learned" api:"nullable"`
	// Home team possession percentage
	PossessionPercentage float64 `json:"possession_percentage" api:"nullable"`
	// Match result from home team perspective
	//
	// Any of "win", "loss", "draw", "pending".
	Result MatchResult `json:"result"`
	// Ted's inspirational halftime speech
	TedHalftimeSpeech string `json:"ted_halftime_speech" api:"nullable"`
	// Total ticket revenue in GBP
	TicketRevenueGbp string `json:"ticket_revenue_gbp" api:"nullable"`
	// Key moments that changed the match
	TurningPoints []TurningPoint `json:"turning_points"`
	// Temperature at kickoff in Celsius
	WeatherTempCelsius float64 `json:"weather_temp_celsius" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AwayTeamID           respjson.Field
		Date                 respjson.Field
		HomeTeamID           respjson.Field
		MatchType            respjson.Field
		Attendance           respjson.Field
		AwayScore            respjson.Field
		EpisodeID            respjson.Field
		HomeScore            respjson.Field
		LessonLearned        respjson.Field
		PossessionPercentage respjson.Field
		Result               respjson.Field
		TedHalftimeSpeech    respjson.Field
		TicketRevenueGbp     respjson.Field
		TurningPoints        respjson.Field
		WeatherTempCelsius   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Match) RawJSON() string { return r.JSON.raw }
func (r *Match) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match result types.
type MatchResult string

const (
	MatchResultWin     MatchResult = "win"
	MatchResultLoss    MatchResult = "loss"
	MatchResultDraw    MatchResult = "draw"
	MatchResultPending MatchResult = "pending"
)

// Types of matches.
type MatchType string

const (
	MatchTypeLeague   MatchType = "league"
	MatchTypeCup      MatchType = "cup"
	MatchTypeFriendly MatchType = "friendly"
	MatchTypePlayoff  MatchType = "playoff"
	MatchTypeFinal    MatchType = "final"
)

// A pivotal moment in a match.
type TurningPoint struct {
	// What happened
	Description string `json:"description" api:"required"`
	// How this affected the team emotionally
	EmotionalImpact string `json:"emotional_impact" api:"required"`
	// Minute of the match
	Minute int64 `json:"minute" api:"required"`
	// Character ID who was central to this moment
	CharacterInvolved string `json:"character_involved" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description       respjson.Field
		EmotionalImpact   respjson.Field
		Minute            respjson.Field
		CharacterInvolved respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurningPoint) RawJSON() string { return r.JSON.raw }
func (r *TurningPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TurningPoint to a TurningPointParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TurningPointParam.Overrides()
func (r TurningPoint) ToParam() TurningPointParam {
	return param.Override[TurningPointParam](json.RawMessage(r.RawJSON()))
}

// A pivotal moment in a match.
//
// The properties Description, EmotionalImpact, Minute are required.
type TurningPointParam struct {
	// What happened
	Description string `json:"description" api:"required"`
	// How this affected the team emotionally
	EmotionalImpact string `json:"emotional_impact" api:"required"`
	// Minute of the match
	Minute int64 `json:"minute" api:"required"`
	// Character ID who was central to this moment
	CharacterInvolved param.Opt[string] `json:"character_involved,omitzero"`
	paramObj
}

func (r TurningPointParam) MarshalJSON() (data []byte, err error) {
	type shadow TurningPointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TurningPointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatchGetLessonResponse map[string]any

type MatchGetTurningPointsResponse map[string]any

type MatchNewParams struct {
	// Away team ID
	AwayTeamID string `json:"away_team_id" api:"required"`
	// Match date and time
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Home team ID
	HomeTeamID string `json:"home_team_id" api:"required"`
	// Type of match
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType MatchType `json:"match_type,omitzero" api:"required"`
	// Match attendance
	Attendance param.Opt[int64] `json:"attendance,omitzero"`
	// Episode ID where this match is featured
	EpisodeID param.Opt[string] `json:"episode_id,omitzero"`
	// The life lesson learned from this match
	LessonLearned param.Opt[string] `json:"lesson_learned,omitzero"`
	// Home team possession percentage
	PossessionPercentage param.Opt[float64] `json:"possession_percentage,omitzero"`
	// Ted's inspirational halftime speech
	TedHalftimeSpeech param.Opt[string] `json:"ted_halftime_speech,omitzero"`
	// Temperature at kickoff in Celsius
	WeatherTempCelsius param.Opt[float64] `json:"weather_temp_celsius,omitzero"`
	// Away team score
	AwayScore param.Opt[int64] `json:"away_score,omitzero"`
	// Home team score
	HomeScore param.Opt[int64] `json:"home_score,omitzero"`
	// Total ticket revenue in GBP
	TicketRevenueGbp MatchNewParamsTicketRevenueGbpUnion `json:"ticket_revenue_gbp,omitzero"`
	// Match result from home team perspective
	//
	// Any of "win", "loss", "draw", "pending".
	Result MatchResult `json:"result,omitzero"`
	// Key moments that changed the match
	TurningPoints []TurningPointParam `json:"turning_points,omitzero"`
	paramObj
}

func (r MatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MatchNewParamsTicketRevenueGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u MatchNewParamsTicketRevenueGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *MatchNewParamsTicketRevenueGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MatchUpdateParams struct {
	Attendance           param.Opt[int64]                       `json:"attendance,omitzero"`
	AwayScore            param.Opt[int64]                       `json:"away_score,omitzero"`
	AwayTeamID           param.Opt[string]                      `json:"away_team_id,omitzero"`
	Date                 param.Opt[time.Time]                   `json:"date,omitzero" format:"date-time"`
	EpisodeID            param.Opt[string]                      `json:"episode_id,omitzero"`
	HomeScore            param.Opt[int64]                       `json:"home_score,omitzero"`
	HomeTeamID           param.Opt[string]                      `json:"home_team_id,omitzero"`
	LessonLearned        param.Opt[string]                      `json:"lesson_learned,omitzero"`
	PossessionPercentage param.Opt[float64]                     `json:"possession_percentage,omitzero"`
	TedHalftimeSpeech    param.Opt[string]                      `json:"ted_halftime_speech,omitzero"`
	WeatherTempCelsius   param.Opt[float64]                     `json:"weather_temp_celsius,omitzero"`
	TicketRevenueGbp     MatchUpdateParamsTicketRevenueGbpUnion `json:"ticket_revenue_gbp,omitzero"`
	TurningPoints        []TurningPointParam                    `json:"turning_points,omitzero"`
	// Types of matches.
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType MatchType `json:"match_type,omitzero"`
	// Match result types.
	//
	// Any of "win", "loss", "draw", "pending".
	Result MatchResult `json:"result,omitzero"`
	paramObj
}

func (r MatchUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MatchUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MatchUpdateParamsTicketRevenueGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u MatchUpdateParamsTicketRevenueGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *MatchUpdateParamsTicketRevenueGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MatchListParams struct {
	// Filter by team (home or away)
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter by match type
	//
	// Any of "league", "cup", "friendly", "playoff", "final".
	MatchType MatchType `query:"match_type,omitzero" json:"-"`
	// Filter by result
	//
	// Any of "win", "loss", "draw", "pending".
	Result MatchResult `query:"result,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatchListParams]'s query parameters as `url.Values`.
func (r MatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatchStreamLiveParams struct {
	// Away team name
	AwayTeam param.Opt[string] `query:"away_team,omitzero" json:"-"`
	// How eventful the match should be (1=boring, 10=chaos)
	ExcitementLevel param.Opt[int64] `query:"excitement_level,omitzero" json:"-"`
	// Home team name
	HomeTeam param.Opt[string] `query:"home_team,omitzero" json:"-"`
	// Simulation speed multiplier (1.0 = real-time)
	Speed param.Opt[float64] `query:"speed,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatchStreamLiveParams]'s query parameters as `url.Values`.
func (r MatchStreamLiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
