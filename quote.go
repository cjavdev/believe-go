// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
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

// Memorable quotes from the show
//
// QuoteService contains methods and other services that help with interacting with
// the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuoteService] method instead.
type QuoteService struct {
	Options []option.RequestOption
}

// NewQuoteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuoteService(opts ...option.RequestOption) (r QuoteService) {
	r = QuoteService{}
	r.Options = opts
	return
}

// Add a new memorable quote to the collection.
func (r *QuoteService) New(ctx context.Context, body QuoteNewParams, opts ...option.RequestOption) (res *Quote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "quotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific quote by its ID.
func (r *QuoteService) Get(ctx context.Context, quoteID string, opts ...option.RequestOption) (res *Quote, err error) {
	opts = slices.Concat(r.Options, opts)
	if quoteID == "" {
		err = errors.New("missing required quote_id parameter")
		return
	}
	path := fmt.Sprintf("quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update specific fields of an existing quote.
func (r *QuoteService) Update(ctx context.Context, quoteID string, body QuoteUpdateParams, opts ...option.RequestOption) (res *Quote, err error) {
	opts = slices.Concat(r.Options, opts)
	if quoteID == "" {
		err = errors.New("missing required quote_id parameter")
		return
	}
	path := fmt.Sprintf("quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of all memorable Ted Lasso quotes with optional filtering.
func (r *QuoteService) List(ctx context.Context, query QuoteListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Quote], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "quotes"
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

// Get a paginated list of all memorable Ted Lasso quotes with optional filtering.
func (r *QuoteService) ListAutoPaging(ctx context.Context, query QuoteListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Quote] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a quote from the collection.
func (r *QuoteService) Delete(ctx context.Context, quoteID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if quoteID == "" {
		err = errors.New("missing required quote_id parameter")
		return
	}
	path := fmt.Sprintf("quotes/%s", url.PathEscape(quoteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a random Ted Lasso quote, optionally filtered.
func (r *QuoteService) GetRandom(ctx context.Context, query QuoteGetRandomParams, opts ...option.RequestOption) (res *Quote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "quotes/random"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a paginated list of quotes from a specific character.
func (r *QuoteService) ListByCharacter(ctx context.Context, characterID string, query QuoteListByCharacterParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Quote], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if characterID == "" {
		err = errors.New("missing required character_id parameter")
		return
	}
	path := fmt.Sprintf("quotes/characters/%s", url.PathEscape(characterID))
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

// Get a paginated list of quotes from a specific character.
func (r *QuoteService) ListByCharacterAutoPaging(ctx context.Context, characterID string, query QuoteListByCharacterParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Quote] {
	return pagination.NewSkipLimitPageAutoPager(r.ListByCharacter(ctx, characterID, query, opts...))
}

// Get a paginated list of quotes related to a specific theme.
func (r *QuoteService) ListByTheme(ctx context.Context, theme QuoteTheme, query QuoteListByThemeParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Quote], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("quotes/themes/%v", theme)
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

// Get a paginated list of quotes related to a specific theme.
func (r *QuoteService) ListByThemeAutoPaging(ctx context.Context, theme QuoteTheme, query QuoteListByThemeParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Quote] {
	return pagination.NewSkipLimitPageAutoPager(r.ListByTheme(ctx, theme, query, opts...))
}

type PaginatedResponseQuote struct {
	Data []Quote `json:"data" api:"required"`
	// Whether there are more items after this page.
	HasMore bool  `json:"has_more" api:"required"`
	Limit   int64 `json:"limit" api:"required"`
	// Current page number (1-indexed, for display purposes).
	Page int64 `json:"page" api:"required"`
	// Total number of pages.
	Pages int64 `json:"pages" api:"required"`
	Skip  int64 `json:"skip" api:"required"`
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Skip        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedResponseQuote) RawJSON() string { return r.JSON.raw }
func (r *PaginatedResponseQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full quote model with ID.
type Quote struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// ID of the character who said it
	CharacterID string `json:"character_id" api:"required"`
	// Context in which the quote was said
	Context string `json:"context" api:"required"`
	// Type of moment when the quote was said
	//
	// Any of "halftime_speech", "press_conference", "locker_room", "training",
	// "biscuits_with_boss", "pub", "one_on_one", "celebration", "crisis", "casual",
	// "confrontation".
	MomentType QuoteMoment `json:"moment_type" api:"required"`
	// The quote text
	Text string `json:"text" api:"required"`
	// Primary theme of the quote
	//
	// Any of "belief", "teamwork", "curiosity", "kindness", "resilience",
	// "vulnerability", "growth", "humor", "wisdom", "leadership", "love",
	// "forgiveness", "philosophy", "romance", "cultural-pride",
	// "cultural-differences", "antagonism", "celebration", "identity", "isolation",
	// "power", "sacrifice", "standards", "confidence", "conflict", "honesty",
	// "integrity", "intimidation", "ambition", "narcissism", "maturity".
	Theme QuoteTheme `json:"theme" api:"required"`
	// Episode where the quote appears
	EpisodeID string `json:"episode_id" api:"nullable"`
	// Whether this quote is humorous
	IsFunny bool `json:"is_funny"`
	// Whether this quote is inspirational
	IsInspirational bool `json:"is_inspirational"`
	// Popularity/virality score (0-100)
	PopularityScore float64 `json:"popularity_score" api:"nullable"`
	// Additional themes
	SecondaryThemes []QuoteTheme `json:"secondary_themes"`
	// Number of times shared on social media
	TimesShared int64 `json:"times_shared" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CharacterID     respjson.Field
		Context         respjson.Field
		MomentType      respjson.Field
		Text            respjson.Field
		Theme           respjson.Field
		EpisodeID       respjson.Field
		IsFunny         respjson.Field
		IsInspirational respjson.Field
		PopularityScore respjson.Field
		SecondaryThemes respjson.Field
		TimesShared     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quote) RawJSON() string { return r.JSON.raw }
func (r *Quote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Types of moments when quotes occur.
type QuoteMoment string

const (
	QuoteMomentHalftimeSpeech   QuoteMoment = "halftime_speech"
	QuoteMomentPressConference  QuoteMoment = "press_conference"
	QuoteMomentLockerRoom       QuoteMoment = "locker_room"
	QuoteMomentTraining         QuoteMoment = "training"
	QuoteMomentBiscuitsWithBoss QuoteMoment = "biscuits_with_boss"
	QuoteMomentPub              QuoteMoment = "pub"
	QuoteMomentOneOnOne         QuoteMoment = "one_on_one"
	QuoteMomentCelebration      QuoteMoment = "celebration"
	QuoteMomentCrisis           QuoteMoment = "crisis"
	QuoteMomentCasual           QuoteMoment = "casual"
	QuoteMomentConfrontation    QuoteMoment = "confrontation"
)

// Themes that quotes can be categorized under.
type QuoteTheme string

const (
	QuoteThemeBelief              QuoteTheme = "belief"
	QuoteThemeTeamwork            QuoteTheme = "teamwork"
	QuoteThemeCuriosity           QuoteTheme = "curiosity"
	QuoteThemeKindness            QuoteTheme = "kindness"
	QuoteThemeResilience          QuoteTheme = "resilience"
	QuoteThemeVulnerability       QuoteTheme = "vulnerability"
	QuoteThemeGrowth              QuoteTheme = "growth"
	QuoteThemeHumor               QuoteTheme = "humor"
	QuoteThemeWisdom              QuoteTheme = "wisdom"
	QuoteThemeLeadership          QuoteTheme = "leadership"
	QuoteThemeLove                QuoteTheme = "love"
	QuoteThemeForgiveness         QuoteTheme = "forgiveness"
	QuoteThemePhilosophy          QuoteTheme = "philosophy"
	QuoteThemeRomance             QuoteTheme = "romance"
	QuoteThemeCulturalPride       QuoteTheme = "cultural-pride"
	QuoteThemeCulturalDifferences QuoteTheme = "cultural-differences"
	QuoteThemeAntagonism          QuoteTheme = "antagonism"
	QuoteThemeCelebration         QuoteTheme = "celebration"
	QuoteThemeIdentity            QuoteTheme = "identity"
	QuoteThemeIsolation           QuoteTheme = "isolation"
	QuoteThemePower               QuoteTheme = "power"
	QuoteThemeSacrifice           QuoteTheme = "sacrifice"
	QuoteThemeStandards           QuoteTheme = "standards"
	QuoteThemeConfidence          QuoteTheme = "confidence"
	QuoteThemeConflict            QuoteTheme = "conflict"
	QuoteThemeHonesty             QuoteTheme = "honesty"
	QuoteThemeIntegrity           QuoteTheme = "integrity"
	QuoteThemeIntimidation        QuoteTheme = "intimidation"
	QuoteThemeAmbition            QuoteTheme = "ambition"
	QuoteThemeNarcissism          QuoteTheme = "narcissism"
	QuoteThemeMaturity            QuoteTheme = "maturity"
)

type QuoteNewParams struct {
	// ID of the character who said it
	CharacterID string `json:"character_id" api:"required"`
	// Context in which the quote was said
	Context string `json:"context" api:"required"`
	// Type of moment when the quote was said
	//
	// Any of "halftime_speech", "press_conference", "locker_room", "training",
	// "biscuits_with_boss", "pub", "one_on_one", "celebration", "crisis", "casual",
	// "confrontation".
	MomentType QuoteMoment `json:"moment_type,omitzero" api:"required"`
	// The quote text
	Text string `json:"text" api:"required"`
	// Primary theme of the quote
	//
	// Any of "belief", "teamwork", "curiosity", "kindness", "resilience",
	// "vulnerability", "growth", "humor", "wisdom", "leadership", "love",
	// "forgiveness", "philosophy", "romance", "cultural-pride",
	// "cultural-differences", "antagonism", "celebration", "identity", "isolation",
	// "power", "sacrifice", "standards", "confidence", "conflict", "honesty",
	// "integrity", "intimidation", "ambition", "narcissism", "maturity".
	Theme QuoteTheme `json:"theme,omitzero" api:"required"`
	// Episode where the quote appears
	EpisodeID param.Opt[string] `json:"episode_id,omitzero"`
	// Popularity/virality score (0-100)
	PopularityScore param.Opt[float64] `json:"popularity_score,omitzero"`
	// Number of times shared on social media
	TimesShared param.Opt[int64] `json:"times_shared,omitzero"`
	// Whether this quote is humorous
	IsFunny param.Opt[bool] `json:"is_funny,omitzero"`
	// Whether this quote is inspirational
	IsInspirational param.Opt[bool] `json:"is_inspirational,omitzero"`
	// Additional themes
	SecondaryThemes []QuoteTheme `json:"secondary_themes,omitzero"`
	paramObj
}

func (r QuoteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QuoteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuoteUpdateParams struct {
	CharacterID     param.Opt[string]  `json:"character_id,omitzero"`
	Context         param.Opt[string]  `json:"context,omitzero"`
	EpisodeID       param.Opt[string]  `json:"episode_id,omitzero"`
	IsFunny         param.Opt[bool]    `json:"is_funny,omitzero"`
	IsInspirational param.Opt[bool]    `json:"is_inspirational,omitzero"`
	PopularityScore param.Opt[float64] `json:"popularity_score,omitzero"`
	Text            param.Opt[string]  `json:"text,omitzero"`
	TimesShared     param.Opt[int64]   `json:"times_shared,omitzero"`
	SecondaryThemes []QuoteTheme       `json:"secondary_themes,omitzero"`
	// Types of moments when quotes occur.
	//
	// Any of "halftime_speech", "press_conference", "locker_room", "training",
	// "biscuits_with_boss", "pub", "one_on_one", "celebration", "crisis", "casual",
	// "confrontation".
	MomentType QuoteMoment `json:"moment_type,omitzero"`
	// Themes that quotes can be categorized under.
	//
	// Any of "belief", "teamwork", "curiosity", "kindness", "resilience",
	// "vulnerability", "growth", "humor", "wisdom", "leadership", "love",
	// "forgiveness", "philosophy", "romance", "cultural-pride",
	// "cultural-differences", "antagonism", "celebration", "identity", "isolation",
	// "power", "sacrifice", "standards", "confidence", "conflict", "honesty",
	// "integrity", "intimidation", "ambition", "narcissism", "maturity".
	Theme QuoteTheme `json:"theme,omitzero"`
	paramObj
}

func (r QuoteUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow QuoteUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuoteListParams struct {
	// Filter by character
	CharacterID param.Opt[string] `query:"character_id,omitzero" json:"-"`
	// Filter funny quotes
	Funny param.Opt[bool] `query:"funny,omitzero" json:"-"`
	// Filter inspirational quotes
	Inspirational param.Opt[bool] `query:"inspirational,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter by moment type
	//
	// Any of "halftime_speech", "press_conference", "locker_room", "training",
	// "biscuits_with_boss", "pub", "one_on_one", "celebration", "crisis", "casual",
	// "confrontation".
	MomentType QuoteMoment `query:"moment_type,omitzero" json:"-"`
	// Filter by theme
	//
	// Any of "belief", "teamwork", "curiosity", "kindness", "resilience",
	// "vulnerability", "growth", "humor", "wisdom", "leadership", "love",
	// "forgiveness", "philosophy", "romance", "cultural-pride",
	// "cultural-differences", "antagonism", "celebration", "identity", "isolation",
	// "power", "sacrifice", "standards", "confidence", "conflict", "honesty",
	// "integrity", "intimidation", "ambition", "narcissism", "maturity".
	Theme QuoteTheme `query:"theme,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuoteListParams]'s query parameters as `url.Values`.
func (r QuoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuoteGetRandomParams struct {
	// Filter by character
	CharacterID param.Opt[string] `query:"character_id,omitzero" json:"-"`
	// Filter inspirational quotes
	Inspirational param.Opt[bool] `query:"inspirational,omitzero" json:"-"`
	// Filter by theme
	//
	// Any of "belief", "teamwork", "curiosity", "kindness", "resilience",
	// "vulnerability", "growth", "humor", "wisdom", "leadership", "love",
	// "forgiveness", "philosophy", "romance", "cultural-pride",
	// "cultural-differences", "antagonism", "celebration", "identity", "isolation",
	// "power", "sacrifice", "standards", "confidence", "conflict", "honesty",
	// "integrity", "intimidation", "ambition", "narcissism", "maturity".
	Theme QuoteTheme `query:"theme,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuoteGetRandomParams]'s query parameters as `url.Values`.
func (r QuoteGetRandomParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuoteListByCharacterParams struct {
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuoteListByCharacterParams]'s query parameters as
// `url.Values`.
func (r QuoteListByCharacterParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuoteListByThemeParams struct {
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuoteListByThemeParams]'s query parameters as `url.Values`.
func (r QuoteListByThemeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
