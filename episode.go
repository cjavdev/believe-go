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

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/apiquery"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/pagination"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// Operations related to TV episodes
//
// EpisodeService contains methods and other services that help with interacting with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment automatically. You should not instantiate this service directly, and instead use the [NewEpisodeService] method instead.
type EpisodeService struct {
	Options []option.RequestOption
}

// NewEpisodeService generates a new service that applies the given options to each request. These options are applied after the parent client's options (if there is one), and before any request-specific options.
func NewEpisodeService(opts ...option.RequestOption) (r EpisodeService) {
	r = EpisodeService{}
	r.Options = opts
	return
}

// Add a new episode to the series.
func (r *EpisodeService) New(ctx context.Context, body EpisodeNewParams, opts ...option.RequestOption) (res *Episode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "episodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific episode.
func (r *EpisodeService) Get(ctx context.Context, episodeID string, opts ...option.RequestOption) (res *Episode, err error) {
	opts = slices.Concat(r.Options, opts)
	if episodeID == "" {
		err = errors.New("missing required episode_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("episodes/%s", url.PathEscape(episodeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update specific fields of an existing episode.
func (r *EpisodeService) Update(ctx context.Context, episodeID string, body EpisodeUpdateParams, opts ...option.RequestOption) (res *Episode, err error) {
	opts = slices.Concat(r.Options, opts)
	if episodeID == "" {
		err = errors.New("missing required episode_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("episodes/%s", url.PathEscape(episodeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a paginated list of all Ted Lasso episodes with optional filtering by season.
func (r *EpisodeService) List(ctx context.Context, query EpisodeListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Episode], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "episodes"
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

// Get a paginated list of all Ted Lasso episodes with optional filtering by season.
func (r *EpisodeService) ListAutoPaging(ctx context.Context, query EpisodeListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Episode] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove an episode from the database.
func (r *EpisodeService) Delete(ctx context.Context, episodeID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if episodeID == "" {
		err = errors.New("missing required episode_id parameter")
		return err
	}
	path := fmt.Sprintf("episodes/%s", url.PathEscape(episodeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get Ted's wisdom and memorable moments from a specific episode.
func (r *EpisodeService) GetWisdom(ctx context.Context, episodeID string, opts ...option.RequestOption) (res *EpisodeGetWisdomResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if episodeID == "" {
		err = errors.New("missing required episode_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("episodes/%s/wisdom", url.PathEscape(episodeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Full episode model with ID.
type Episode struct {
	// Unique identifier (format: s##e##)
	ID string `json:"id" api:"required"`
	// Original air date
	AirDate time.Time `json:"air_date" api:"required" format:"date"`
	// Characters with significant development
	CharacterFocus []string `json:"character_focus" api:"required"`
	// Episode director
	Director string `json:"director" api:"required"`
	// Episode number within season
	EpisodeNumber int64 `json:"episode_number" api:"required"`
	// Central theme of the episode
	MainTheme string `json:"main_theme" api:"required"`
	// Episode runtime in minutes
	RuntimeMinutes int64 `json:"runtime_minutes" api:"required"`
	// Season number
	Season int64 `json:"season" api:"required"`
	// Brief plot synopsis
	Synopsis string `json:"synopsis" api:"required"`
	// Key piece of Ted wisdom from the episode
	TedWisdom string `json:"ted_wisdom" api:"required"`
	// Episode title
	Title string `json:"title" api:"required"`
	// Episode writer(s)
	Writer string `json:"writer" api:"required"`
	// Notable biscuits with the boss scene
	BiscuitsWithBossMoment string `json:"biscuits_with_boss_moment" api:"nullable"`
	// Standout moments from the episode
	MemorableMoments []string `json:"memorable_moments"`
	// US viewership in millions
	UsViewersMillions float64 `json:"us_viewers_millions" api:"nullable"`
	// Viewer rating out of 10
	ViewerRating float64 `json:"viewer_rating" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AirDate                respjson.Field
		CharacterFocus         respjson.Field
		Director               respjson.Field
		EpisodeNumber          respjson.Field
		MainTheme              respjson.Field
		RuntimeMinutes         respjson.Field
		Season                 respjson.Field
		Synopsis               respjson.Field
		TedWisdom              respjson.Field
		Title                  respjson.Field
		Writer                 respjson.Field
		BiscuitsWithBossMoment respjson.Field
		MemorableMoments       respjson.Field
		UsViewersMillions      respjson.Field
		ViewerRating           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Episode) RawJSON() string { return r.JSON.raw }
func (r *Episode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginatedResponse struct {
	Data []Episode `json:"data" api:"required"`
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
func (r PaginatedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EpisodeGetWisdomResponse map[string]any

type EpisodeNewParams struct {
	// Original air date
	AirDate time.Time `json:"air_date" api:"required" format:"date"`
	// Characters with significant development
	CharacterFocus []string `json:"character_focus,omitzero" api:"required"`
	// Episode director
	Director string `json:"director" api:"required"`
	// Episode number within season
	EpisodeNumber int64 `json:"episode_number" api:"required"`
	// Central theme of the episode
	MainTheme string `json:"main_theme" api:"required"`
	// Episode runtime in minutes
	RuntimeMinutes int64 `json:"runtime_minutes" api:"required"`
	// Season number
	Season int64 `json:"season" api:"required"`
	// Brief plot synopsis
	Synopsis string `json:"synopsis" api:"required"`
	// Key piece of Ted wisdom from the episode
	TedWisdom string `json:"ted_wisdom" api:"required"`
	// Episode title
	Title string `json:"title" api:"required"`
	// Episode writer(s)
	Writer string `json:"writer" api:"required"`
	// Notable biscuits with the boss scene
	BiscuitsWithBossMoment param.Opt[string] `json:"biscuits_with_boss_moment,omitzero"`
	// US viewership in millions
	UsViewersMillions param.Opt[float64] `json:"us_viewers_millions,omitzero"`
	// Viewer rating out of 10
	ViewerRating param.Opt[float64] `json:"viewer_rating,omitzero"`
	// Standout moments from the episode
	MemorableMoments []string `json:"memorable_moments,omitzero"`
	paramObj
}

func (r EpisodeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EpisodeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EpisodeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EpisodeUpdateParams struct {
	AirDate                param.Opt[time.Time] `json:"air_date,omitzero" format:"date"`
	BiscuitsWithBossMoment param.Opt[string]    `json:"biscuits_with_boss_moment,omitzero"`
	Director               param.Opt[string]    `json:"director,omitzero"`
	EpisodeNumber          param.Opt[int64]     `json:"episode_number,omitzero"`
	MainTheme              param.Opt[string]    `json:"main_theme,omitzero"`
	RuntimeMinutes         param.Opt[int64]     `json:"runtime_minutes,omitzero"`
	Season                 param.Opt[int64]     `json:"season,omitzero"`
	Synopsis               param.Opt[string]    `json:"synopsis,omitzero"`
	TedWisdom              param.Opt[string]    `json:"ted_wisdom,omitzero"`
	Title                  param.Opt[string]    `json:"title,omitzero"`
	UsViewersMillions      param.Opt[float64]   `json:"us_viewers_millions,omitzero"`
	ViewerRating           param.Opt[float64]   `json:"viewer_rating,omitzero"`
	Writer                 param.Opt[string]    `json:"writer,omitzero"`
	CharacterFocus         []string             `json:"character_focus,omitzero"`
	MemorableMoments       []string             `json:"memorable_moments,omitzero"`
	paramObj
}

func (r EpisodeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EpisodeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EpisodeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EpisodeListParams struct {
	// Filter by character focus (character ID)
	CharacterFocus param.Opt[string] `query:"character_focus,omitzero" json:"-"`
	// Filter by season
	Season param.Opt[int64] `query:"season,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EpisodeListParams]'s query parameters as `url.Values`.
func (r EpisodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
