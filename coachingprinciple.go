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

// Interactive endpoints for motivation and guidance
//
// CoachingPrincipleService contains methods and other services that help with
// interacting with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoachingPrincipleService] method instead.
type CoachingPrincipleService struct {
	Options []option.RequestOption
}

// NewCoachingPrincipleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCoachingPrincipleService(opts ...option.RequestOption) (r CoachingPrincipleService) {
	r = CoachingPrincipleService{}
	r.Options = opts
	return
}

// Get details about a specific coaching principle.
func (r *CoachingPrincipleService) Get(ctx context.Context, principleID string, opts ...option.RequestOption) (res *CoachingPrinciple, err error) {
	opts = slices.Concat(r.Options, opts)
	if principleID == "" {
		err = errors.New("missing required principle_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("coaching/principles/%s", url.PathEscape(principleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a paginated list of Ted Lasso's core coaching principles and philosophy.
func (r *CoachingPrincipleService) List(ctx context.Context, query CoachingPrincipleListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[CoachingPrinciple], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "coaching/principles"
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

// Get a paginated list of Ted Lasso's core coaching principles and philosophy.
func (r *CoachingPrincipleService) ListAutoPaging(ctx context.Context, query CoachingPrincipleListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[CoachingPrinciple] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Get a random coaching principle to inspire your day.
func (r *CoachingPrincipleService) GetRandom(ctx context.Context, opts ...option.RequestOption) (res *CoachingPrinciple, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "coaching/principles/random"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A Ted Lasso coaching principle.
type CoachingPrinciple struct {
	// Principle identifier
	ID string `json:"id" api:"required"`
	// How to apply this principle
	Application string `json:"application" api:"required"`
	// Example from the show
	ExampleFromShow string `json:"example_from_show" api:"required"`
	// What this principle means
	Explanation string `json:"explanation" api:"required"`
	// The coaching principle
	Principle string `json:"principle" api:"required"`
	// Related Ted quote
	TedQuote string `json:"ted_quote" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Application     respjson.Field
		ExampleFromShow respjson.Field
		Explanation     respjson.Field
		Principle       respjson.Field
		TedQuote        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CoachingPrinciple) RawJSON() string { return r.JSON.raw }
func (r *CoachingPrinciple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoachingPrincipleListParams struct {
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoachingPrincipleListParams]'s query parameters as
// `url.Values`.
func (r CoachingPrincipleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
