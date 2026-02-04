// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/believe-go/internal/apijson"
	"github.com/stainless-sdks/believe-go/internal/apiquery"
	"github.com/stainless-sdks/believe-go/internal/requestconfig"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-go/packages/respjson"
)

// BiscuitService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBiscuitService] method instead.
type BiscuitService struct {
	Options []option.RequestOption
}

// NewBiscuitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBiscuitService(opts ...option.RequestOption) (r BiscuitService) {
	r = BiscuitService{}
	r.Options = opts
	return
}

// Get a specific type of biscuit by ID.
func (r *BiscuitService) Get(ctx context.Context, biscuitID string, opts ...option.RequestOption) (res *Biscuit, err error) {
	opts = slices.Concat(r.Options, opts)
	if biscuitID == "" {
		err = errors.New("missing required biscuit_id parameter")
		return
	}
	path := fmt.Sprintf("biscuits/%s", url.PathEscape(biscuitID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a paginated list of Ted's famous homemade biscuits! Each comes with a
// heartwarming message.
func (r *BiscuitService) List(ctx context.Context, query BiscuitListParams, opts ...option.RequestOption) (res *BiscuitListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "biscuits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a single fresh biscuit with a personalized message from Ted.
func (r *BiscuitService) GetFresh(ctx context.Context, opts ...option.RequestOption) (res *Biscuit, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "biscuits/fresh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A biscuit from Ted.
type Biscuit struct {
	// Biscuit identifier
	ID string `json:"id,required"`
	// Message that comes with the biscuit
	Message string `json:"message,required"`
	// What this biscuit pairs well with
	PairsWellWith string `json:"pairs_well_with,required"`
	// A handwritten note from Ted
	TedNote string `json:"ted_note,required"`
	// Type of biscuit
	//
	// Any of "classic", "shortbread", "chocolate_chip", "oatmeal_raisin".
	Type BiscuitType `json:"type,required"`
	// How warm and fresh (1-10)
	WarmthLevel int64 `json:"warmth_level,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Message       respjson.Field
		PairsWellWith respjson.Field
		TedNote       respjson.Field
		Type          respjson.Field
		WarmthLevel   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Biscuit) RawJSON() string { return r.JSON.raw }
func (r *Biscuit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of biscuit
type BiscuitType string

const (
	BiscuitTypeClassic       BiscuitType = "classic"
	BiscuitTypeShortbread    BiscuitType = "shortbread"
	BiscuitTypeChocolateChip BiscuitType = "chocolate_chip"
	BiscuitTypeOatmealRaisin BiscuitType = "oatmeal_raisin"
)

type BiscuitListResponse struct {
	Data []Biscuit `json:"data,required"`
	// Whether there are more items after this page.
	HasMore bool  `json:"has_more,required"`
	Limit   int64 `json:"limit,required"`
	// Current page number (1-indexed, for display purposes).
	Page int64 `json:"page,required"`
	// Total number of pages.
	Pages int64 `json:"pages,required"`
	Skip  int64 `json:"skip,required"`
	Total int64 `json:"total,required"`
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
func (r BiscuitListResponse) RawJSON() string { return r.JSON.raw }
func (r *BiscuitListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BiscuitListParams struct {
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BiscuitListParams]'s query parameters as `url.Values`.
func (r BiscuitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
