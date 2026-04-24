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

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/apiquery"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/pagination"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// Operations related to football teams
//
// TeamService contains methods and other services that help with interacting with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment automatically. You should not instantiate this service directly, and instead use the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
	// Operations related to football teams
	Logo TeamLogoService
}

// NewTeamService generates a new service that applies the given options to each request. These options are applied after the parent client's options (if there is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r TeamService) {
	r = TeamService{}
	r.Options = opts
	r.Logo = NewTeamLogoService(opts...)
	return
}

// Add a new team to the league.
func (r *TeamService) New(ctx context.Context, body TeamNewParams, opts ...option.RequestOption) (res *Team, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific team.
func (r *TeamService) Get(ctx context.Context, teamID string, opts ...option.RequestOption) (res *Team, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("teams/%s", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update specific fields of an existing team.
func (r *TeamService) Update(ctx context.Context, teamID string, body TeamUpdateParams, opts ...option.RequestOption) (res *Team, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("teams/%s", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a paginated list of all teams with optional filtering by league or culture score.
func (r *TeamService) List(ctx context.Context, query TeamListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Team], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "teams"
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

// Get a paginated list of all teams with optional filtering by league or culture score.
func (r *TeamService) ListAutoPaging(ctx context.Context, query TeamListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Team] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a team from the database (relegation to oblivion).
func (r *TeamService) Delete(ctx context.Context, teamID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return err
	}
	path := fmt.Sprintf("teams/%s", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get detailed culture and values information for a team.
func (r *TeamService) GetCulture(ctx context.Context, teamID string, opts ...option.RequestOption) (res *TeamGetCultureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("teams/%s/culture", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get all rival teams for a specific team.
func (r *TeamService) GetRivals(ctx context.Context, teamID string, opts ...option.RequestOption) (res *[]Team, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("teams/%s/rivals", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all uploaded logos for a team.
func (r *TeamService) ListLogos(ctx context.Context, teamID string, opts ...option.RequestOption) (res *[]FileUpload, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("teams/%s/logos", url.PathEscape(teamID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Geographic coordinates for a location.
type GeoLocation struct {
	// Latitude in degrees
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude in degrees
	Longitude float64 `json:"longitude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoLocation) RawJSON() string { return r.JSON.raw }
func (r *GeoLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GeoLocation to a GeoLocationParam.
//
// Warning: the fields of the param type will not be present.
// ToParam should only be used at the last possible moment before sending a request.
// Test for this with GeoLocationParam.Overrides()
func (r GeoLocation) ToParam() GeoLocationParam {
	return param.Override[GeoLocationParam](json.RawMessage(r.RawJSON()))
}

// Geographic coordinates for a location.
//
// The properties Latitude, Longitude are required.
type GeoLocationParam struct {
	// Latitude in degrees
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude in degrees
	Longitude float64 `json:"longitude" api:"required"`
	paramObj
}

func (r GeoLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow GeoLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeoLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Football leagues.
type League string

const (
	LeaguePremierLeague League = "Premier League"
	LeagueChampionship  League = "Championship"
	LeagueLeagueOne     League = "League One"
	LeagueLeagueTwo     League = "League Two"
	LeagueLaLiga        League = "La Liga"
	LeagueSerieA        League = "Serie A"
	LeagueBundesliga    League = "Bundesliga"
	LeagueLigue1        League = "Ligue 1"
)

// Full team model with ID.
type Team struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Team culture/morale score (0-100)
	CultureScore int64 `json:"culture_score" api:"required"`
	// Year the club was founded
	FoundedYear int64 `json:"founded_year" api:"required"`
	// Current league
	//
	// Any of "Premier League", "Championship", "League One", "League Two", "La Liga", "Serie A", "Bundesliga", "Ligue 1".
	League League `json:"league" api:"required"`
	// Team name
	Name string `json:"name" api:"required"`
	// Home stadium name
	Stadium string `json:"stadium" api:"required"`
	// Team's core values
	Values TeamValues `json:"values" api:"required"`
	// Annual budget in GBP
	AnnualBudgetGbp string `json:"annual_budget_gbp" api:"nullable"`
	// Average match attendance
	AverageAttendance float64 `json:"average_attendance" api:"nullable"`
	// Team contact email
	ContactEmail string `json:"contact_email" api:"nullable" format:"email"`
	// Whether the team is currently active
	IsActive bool `json:"is_active"`
	// Team nickname
	Nickname string `json:"nickname" api:"nullable"`
	// Primary team color (hex)
	PrimaryColor string `json:"primary_color" api:"nullable"`
	// List of rival team IDs
	RivalTeams []string `json:"rival_teams"`
	// Secondary team color (hex)
	SecondaryColor string `json:"secondary_color" api:"nullable"`
	// Geographic coordinates for a location.
	StadiumLocation GeoLocation `json:"stadium_location" api:"nullable"`
	// Official team website
	Website string `json:"website" api:"nullable" format:"uri"`
	// Season win percentage
	WinPercentage float64 `json:"win_percentage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CultureScore      respjson.Field
		FoundedYear       respjson.Field
		League            respjson.Field
		Name              respjson.Field
		Stadium           respjson.Field
		Values            respjson.Field
		AnnualBudgetGbp   respjson.Field
		AverageAttendance respjson.Field
		ContactEmail      respjson.Field
		IsActive          respjson.Field
		Nickname          respjson.Field
		PrimaryColor      respjson.Field
		RivalTeams        respjson.Field
		SecondaryColor    respjson.Field
		StadiumLocation   respjson.Field
		Website           respjson.Field
		WinPercentage     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Team) RawJSON() string { return r.JSON.raw }
func (r *Team) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Core values that define a team's culture.
type TeamValues struct {
	// The team's primary guiding value
	PrimaryValue string `json:"primary_value" api:"required"`
	// Supporting values
	SecondaryValues []string `json:"secondary_values" api:"required"`
	// Team's motivational motto
	TeamMotto string `json:"team_motto" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryValue    respjson.Field
		SecondaryValues respjson.Field
		TeamMotto       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamValues) RawJSON() string { return r.JSON.raw }
func (r *TeamValues) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TeamValues to a TeamValuesParam.
//
// Warning: the fields of the param type will not be present.
// ToParam should only be used at the last possible moment before sending a request.
// Test for this with TeamValuesParam.Overrides()
func (r TeamValues) ToParam() TeamValuesParam {
	return param.Override[TeamValuesParam](json.RawMessage(r.RawJSON()))
}

// Core values that define a team's culture.
//
// The properties PrimaryValue, SecondaryValues, TeamMotto are required.
type TeamValuesParam struct {
	// The team's primary guiding value
	PrimaryValue string `json:"primary_value" api:"required"`
	// Supporting values
	SecondaryValues []string `json:"secondary_values,omitzero" api:"required"`
	// Team's motivational motto
	TeamMotto string `json:"team_motto" api:"required"`
	paramObj
}

func (r TeamValuesParam) MarshalJSON() (data []byte, err error) {
	type shadow TeamValuesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamValuesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamGetCultureResponse map[string]any

type TeamNewParams struct {
	// Team culture/morale score (0-100)
	CultureScore int64 `json:"culture_score" api:"required"`
	// Year the club was founded
	FoundedYear int64 `json:"founded_year" api:"required"`
	// Current league
	//
	// Any of "Premier League", "Championship", "League One", "League Two", "La Liga", "Serie A", "Bundesliga", "Ligue 1".
	League League `json:"league,omitzero" api:"required"`
	// Team name
	Name string `json:"name" api:"required"`
	// Home stadium name
	Stadium string `json:"stadium" api:"required"`
	// Team's core values
	Values TeamValuesParam `json:"values,omitzero" api:"required"`
	// Average match attendance
	AverageAttendance param.Opt[float64] `json:"average_attendance,omitzero"`
	// Team contact email
	ContactEmail param.Opt[string] `json:"contact_email,omitzero" format:"email"`
	// Team nickname
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	// Primary team color (hex)
	PrimaryColor param.Opt[string] `json:"primary_color,omitzero"`
	// Secondary team color (hex)
	SecondaryColor param.Opt[string] `json:"secondary_color,omitzero"`
	// Official team website
	Website param.Opt[string] `json:"website,omitzero" format:"uri"`
	// Season win percentage
	WinPercentage param.Opt[float64] `json:"win_percentage,omitzero"`
	// Whether the team is currently active
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// Annual budget in GBP
	AnnualBudgetGbp TeamNewParamsAnnualBudgetGbpUnion `json:"annual_budget_gbp,omitzero"`
	// List of rival team IDs
	RivalTeams []string `json:"rival_teams,omitzero"`
	// Geographic coordinates for a location.
	StadiumLocation GeoLocationParam `json:"stadium_location,omitzero"`
	paramObj
}

func (r TeamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TeamNewParamsAnnualBudgetGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TeamNewParamsAnnualBudgetGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TeamNewParamsAnnualBudgetGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TeamUpdateParams struct {
	AverageAttendance param.Opt[float64]                   `json:"average_attendance,omitzero"`
	ContactEmail      param.Opt[string]                    `json:"contact_email,omitzero" format:"email"`
	CultureScore      param.Opt[int64]                     `json:"culture_score,omitzero"`
	FoundedYear       param.Opt[int64]                     `json:"founded_year,omitzero"`
	IsActive          param.Opt[bool]                      `json:"is_active,omitzero"`
	Name              param.Opt[string]                    `json:"name,omitzero"`
	Nickname          param.Opt[string]                    `json:"nickname,omitzero"`
	PrimaryColor      param.Opt[string]                    `json:"primary_color,omitzero"`
	SecondaryColor    param.Opt[string]                    `json:"secondary_color,omitzero"`
	Stadium           param.Opt[string]                    `json:"stadium,omitzero"`
	Website           param.Opt[string]                    `json:"website,omitzero" format:"uri"`
	WinPercentage     param.Opt[float64]                   `json:"win_percentage,omitzero"`
	AnnualBudgetGbp   TeamUpdateParamsAnnualBudgetGbpUnion `json:"annual_budget_gbp,omitzero"`
	RivalTeams        []string                             `json:"rival_teams,omitzero"`
	// Football leagues.
	//
	// Any of "Premier League", "Championship", "League One", "League Two", "La Liga", "Serie A", "Bundesliga", "Ligue 1".
	League League `json:"league,omitzero"`
	// Geographic coordinates for a location.
	StadiumLocation GeoLocationParam `json:"stadium_location,omitzero"`
	// Core values that define a team's culture.
	Values TeamValuesParam `json:"values,omitzero"`
	paramObj
}

func (r TeamUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TeamUpdateParamsAnnualBudgetGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TeamUpdateParamsAnnualBudgetGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TeamUpdateParamsAnnualBudgetGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TeamListParams struct {
	// Minimum culture score
	MinCultureScore param.Opt[int64] `query:"min_culture_score,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter by league
	//
	// Any of "Premier League", "Championship", "League One", "League Two", "La Liga", "Serie A", "Bundesliga", "Ligue 1".
	League League `query:"league,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamListParams]'s query parameters as `url.Values`.
func (r TeamListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
