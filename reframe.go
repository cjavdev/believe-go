// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"net/http"
	"slices"

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// ReframeService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReframeService] method instead.
type ReframeService struct {
	Options []option.RequestOption
}

// NewReframeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReframeService(opts ...option.RequestOption) (r ReframeService) {
	r = ReframeService{}
	r.Options = opts
	return
}

// Transform negative thoughts into positive perspectives with Ted's help.
func (r *ReframeService) TransformNegativeThoughts(ctx context.Context, body ReframeTransformNegativeThoughtsParams, opts ...option.RequestOption) (res *ReframeTransformNegativeThoughtsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reframe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reframed perspective response.
type ReframeTransformNegativeThoughtsResponse struct {
	// A daily affirmation to practice
	DailyAffirmation string `json:"daily_affirmation,required"`
	// The original negative thought
	OriginalThought string `json:"original_thought,required"`
	// The thought reframed positively
	ReframedThought string `json:"reframed_thought,required"`
	// Ted's take on this thought
	TedPerspective string `json:"ted_perspective,required"`
	// Dr. Sharon's therapeutic insight
	DrSharonInsight string `json:"dr_sharon_insight,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyAffirmation respjson.Field
		OriginalThought  respjson.Field
		ReframedThought  respjson.Field
		TedPerspective   respjson.Field
		DrSharonInsight  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReframeTransformNegativeThoughtsResponse) RawJSON() string { return r.JSON.raw }
func (r *ReframeTransformNegativeThoughtsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReframeTransformNegativeThoughtsParams struct {
	// The negative thought to reframe
	NegativeThought string `json:"negative_thought,required"`
	// Is this a recurring thought?
	Recurring param.Opt[bool] `json:"recurring,omitzero"`
	paramObj
}

func (r ReframeTransformNegativeThoughtsParams) MarshalJSON() (data []byte, err error) {
	type shadow ReframeTransformNegativeThoughtsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReframeTransformNegativeThoughtsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
