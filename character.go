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

// Operations related to Ted Lasso characters
//
// CharacterService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCharacterService] method instead.
type CharacterService struct {
	Options []option.RequestOption
}

// NewCharacterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCharacterService(opts ...option.RequestOption) (r CharacterService) {
	r = CharacterService{}
	r.Options = opts
	return
}

// Add a new character to the Ted Lasso universe.
func (r *CharacterService) New(ctx context.Context, body CharacterNewParams, opts ...option.RequestOption) (res *Characterz, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "characters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific character.
func (r *CharacterService) Get(ctx context.Context, characterID string, opts ...option.RequestOption) (res *Characterz, err error) {
	opts = slices.Concat(r.Options, opts)
	if characterID == "" {
		err = errors.New("missing required character_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("characters/%s", url.PathEscape(characterID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update specific fields of an existing character.
func (r *CharacterService) Update(ctx context.Context, characterID string, body CharacterUpdateParams, opts ...option.RequestOption) (res *Characterz, err error) {
	opts = slices.Concat(r.Options, opts)
	if characterID == "" {
		err = errors.New("missing required character_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("characters/%s", url.PathEscape(characterID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a paginated list of Ted Lasso characters.
func (r *CharacterService) List(ctx context.Context, query CharacterListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Characterz], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "characters"
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

// Get a paginated list of Ted Lasso characters.
func (r *CharacterService) ListAutoPaging(ctx context.Context, query CharacterListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Characterz] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a character from the database.
func (r *CharacterService) Delete(ctx context.Context, characterID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if characterID == "" {
		err = errors.New("missing required character_id parameter")
		return err
	}
	path := fmt.Sprintf("characters/%s", url.PathEscape(characterID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get all signature quotes from a specific character.
func (r *CharacterService) GetQuotes(ctx context.Context, characterID string, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	if characterID == "" {
		err = errors.New("missing required character_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("characters/%s/quotes", url.PathEscape(characterID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Roles characters can have.
type CharacterRole string

const (
	CharacterRoleCoach      CharacterRole = "coach"
	CharacterRolePlayer     CharacterRole = "player"
	CharacterRoleOwner      CharacterRole = "owner"
	CharacterRoleManager    CharacterRole = "manager"
	CharacterRoleStaff      CharacterRole = "staff"
	CharacterRoleJournalist CharacterRole = "journalist"
	CharacterRoleFamily     CharacterRole = "family"
	CharacterRoleFriend     CharacterRole = "friend"
	CharacterRoleFan        CharacterRole = "fan"
	CharacterRoleOther      CharacterRole = "other"
)

// Full character model with ID.
type Characterz struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Character background and history
	Background string `json:"background" api:"required"`
	// Emotional intelligence stats
	EmotionalStats EmotionalStats `json:"emotional_stats" api:"required"`
	// Character's full name
	Name string `json:"name" api:"required"`
	// Key personality traits
	PersonalityTraits []string `json:"personality_traits" api:"required"`
	// Character's role
	//
	// Any of "coach", "player", "owner", "manager", "staff", "journalist", "family",
	// "friend", "fan", "other".
	Role CharacterRole `json:"role" api:"required"`
	// Character's date of birth
	DateOfBirth time.Time `json:"date_of_birth" api:"nullable" format:"date"`
	// Character's email address
	Email string `json:"email" api:"nullable" format:"email"`
	// Character development across seasons
	GrowthArcs []GrowthArc `json:"growth_arcs"`
	// Height in meters
	HeightMeters float64 `json:"height_meters" api:"nullable"`
	// URL to character's profile image
	ProfileImageURL string `json:"profile_image_url" api:"nullable" format:"uri"`
	// Annual salary in GBP
	SalaryGbp string `json:"salary_gbp" api:"nullable"`
	// Memorable quotes from this character
	SignatureQuotes []string `json:"signature_quotes"`
	// ID of the team they belong to
	TeamID string `json:"team_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Background        respjson.Field
		EmotionalStats    respjson.Field
		Name              respjson.Field
		PersonalityTraits respjson.Field
		Role              respjson.Field
		DateOfBirth       respjson.Field
		Email             respjson.Field
		GrowthArcs        respjson.Field
		HeightMeters      respjson.Field
		ProfileImageURL   respjson.Field
		SalaryGbp         respjson.Field
		SignatureQuotes   respjson.Field
		TeamID            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Characterz) RawJSON() string { return r.JSON.raw }
func (r *Characterz) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emotional intelligence statistics for a character.
type EmotionalStats struct {
	// Level of curiosity over judgment (0-100)
	Curiosity int64 `json:"curiosity" api:"required"`
	// Capacity for empathy (0-100)
	Empathy int64 `json:"empathy" api:"required"`
	// Level of optimism (0-100)
	Optimism int64 `json:"optimism" api:"required"`
	// Bounce-back ability (0-100)
	Resilience int64 `json:"resilience" api:"required"`
	// Willingness to be vulnerable (0-100)
	Vulnerability int64 `json:"vulnerability" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Curiosity     respjson.Field
		Empathy       respjson.Field
		Optimism      respjson.Field
		Resilience    respjson.Field
		Vulnerability respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmotionalStats) RawJSON() string { return r.JSON.raw }
func (r *EmotionalStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmotionalStats to a EmotionalStatsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmotionalStatsParam.Overrides()
func (r EmotionalStats) ToParam() EmotionalStatsParam {
	return param.Override[EmotionalStatsParam](json.RawMessage(r.RawJSON()))
}

// Emotional intelligence statistics for a character.
//
// The properties Curiosity, Empathy, Optimism, Resilience, Vulnerability are
// required.
type EmotionalStatsParam struct {
	// Level of curiosity over judgment (0-100)
	Curiosity int64 `json:"curiosity" api:"required"`
	// Capacity for empathy (0-100)
	Empathy int64 `json:"empathy" api:"required"`
	// Level of optimism (0-100)
	Optimism int64 `json:"optimism" api:"required"`
	// Bounce-back ability (0-100)
	Resilience int64 `json:"resilience" api:"required"`
	// Willingness to be vulnerable (0-100)
	Vulnerability int64 `json:"vulnerability" api:"required"`
	paramObj
}

func (r EmotionalStatsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmotionalStatsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmotionalStatsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Character development arc.
type GrowthArc struct {
	// Key breakthrough moment
	Breakthrough string `json:"breakthrough" api:"required"`
	// Main challenge faced
	Challenge string `json:"challenge" api:"required"`
	// Where the character ends up
	EndingPoint string `json:"ending_point" api:"required"`
	// Season number
	Season int64 `json:"season" api:"required"`
	// Where the character starts emotionally
	StartingPoint string `json:"starting_point" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakthrough  respjson.Field
		Challenge     respjson.Field
		EndingPoint   respjson.Field
		Season        respjson.Field
		StartingPoint respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GrowthArc) RawJSON() string { return r.JSON.raw }
func (r *GrowthArc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GrowthArc to a GrowthArcParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GrowthArcParam.Overrides()
func (r GrowthArc) ToParam() GrowthArcParam {
	return param.Override[GrowthArcParam](json.RawMessage(r.RawJSON()))
}

// Character development arc.
//
// The properties Breakthrough, Challenge, EndingPoint, Season, StartingPoint are
// required.
type GrowthArcParam struct {
	// Key breakthrough moment
	Breakthrough string `json:"breakthrough" api:"required"`
	// Main challenge faced
	Challenge string `json:"challenge" api:"required"`
	// Where the character ends up
	EndingPoint string `json:"ending_point" api:"required"`
	// Season number
	Season int64 `json:"season" api:"required"`
	// Where the character starts emotionally
	StartingPoint string `json:"starting_point" api:"required"`
	paramObj
}

func (r GrowthArcParam) MarshalJSON() (data []byte, err error) {
	type shadow GrowthArcParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GrowthArcParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CharacterNewParams struct {
	// Character background and history
	Background string `json:"background" api:"required"`
	// Emotional intelligence stats
	EmotionalStats EmotionalStatsParam `json:"emotional_stats,omitzero" api:"required"`
	// Character's full name
	Name string `json:"name" api:"required"`
	// Key personality traits
	PersonalityTraits []string `json:"personality_traits,omitzero" api:"required"`
	// Character's role
	//
	// Any of "coach", "player", "owner", "manager", "staff", "journalist", "family",
	// "friend", "fan", "other".
	Role CharacterRole `json:"role,omitzero" api:"required"`
	// Character's date of birth
	DateOfBirth param.Opt[time.Time] `json:"date_of_birth,omitzero" format:"date"`
	// Character's email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Height in meters
	HeightMeters param.Opt[float64] `json:"height_meters,omitzero"`
	// URL to character's profile image
	ProfileImageURL param.Opt[string] `json:"profile_image_url,omitzero" format:"uri"`
	// ID of the team they belong to
	TeamID param.Opt[string] `json:"team_id,omitzero"`
	// Annual salary in GBP
	SalaryGbp CharacterNewParamsSalaryGbpUnion `json:"salary_gbp,omitzero"`
	// Character development across seasons
	GrowthArcs []GrowthArcParam `json:"growth_arcs,omitzero"`
	// Memorable quotes from this character
	SignatureQuotes []string `json:"signature_quotes,omitzero"`
	paramObj
}

func (r CharacterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CharacterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CharacterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CharacterNewParamsSalaryGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CharacterNewParamsSalaryGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CharacterNewParamsSalaryGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type CharacterUpdateParams struct {
	Background        param.Opt[string]                   `json:"background,omitzero"`
	DateOfBirth       param.Opt[time.Time]                `json:"date_of_birth,omitzero" format:"date"`
	Email             param.Opt[string]                   `json:"email,omitzero" format:"email"`
	HeightMeters      param.Opt[float64]                  `json:"height_meters,omitzero"`
	Name              param.Opt[string]                   `json:"name,omitzero"`
	ProfileImageURL   param.Opt[string]                   `json:"profile_image_url,omitzero" format:"uri"`
	TeamID            param.Opt[string]                   `json:"team_id,omitzero"`
	GrowthArcs        []GrowthArcParam                    `json:"growth_arcs,omitzero"`
	PersonalityTraits []string                            `json:"personality_traits,omitzero"`
	SalaryGbp         CharacterUpdateParamsSalaryGbpUnion `json:"salary_gbp,omitzero"`
	SignatureQuotes   []string                            `json:"signature_quotes,omitzero"`
	// Emotional intelligence statistics for a character.
	EmotionalStats EmotionalStatsParam `json:"emotional_stats,omitzero"`
	// Roles characters can have.
	//
	// Any of "coach", "player", "owner", "manager", "staff", "journalist", "family",
	// "friend", "fan", "other".
	Role CharacterRole `json:"role,omitzero"`
	paramObj
}

func (r CharacterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CharacterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CharacterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CharacterUpdateParamsSalaryGbpUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CharacterUpdateParamsSalaryGbpUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CharacterUpdateParamsSalaryGbpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type CharacterListParams struct {
	// Minimum optimism score
	MinOptimism param.Opt[int64] `query:"min_optimism,omitzero" json:"-"`
	// Filter by team ID
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter by role
	//
	// Any of "coach", "player", "owner", "manager", "staff", "journalist", "family",
	// "friend", "fan", "other".
	Role CharacterRole `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CharacterListParams]'s query parameters as `url.Values`.
func (r CharacterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
