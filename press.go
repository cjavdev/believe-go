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

// Interactive endpoints for motivation and guidance
//
// PressService contains methods and other services that help with interacting with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment automatically. You should not instantiate this service directly, and instead use the [NewPressService] method instead.
type PressService struct {
	Options []option.RequestOption
}

// NewPressService generates a new service that applies the given options to each request. These options are applied after the parent client's options (if there is one), and before any request-specific options.
func NewPressService(opts ...option.RequestOption) (r PressService) {
	r = PressService{}
	r.Options = opts
	return
}

// Get Ted's response to press conference questions.
func (r *PressService) Simulate(ctx context.Context, body PressSimulateParams, opts ...option.RequestOption) (res *PressSimulateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "press"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Ted's press conference response.
type PressSimulateResponse struct {
	// The actual wisdom beneath the humor
	ActualWisdom string `json:"actual_wisdom" api:"required"`
	// How Ted would dodge a follow-up
	FollowUpDodge string `json:"follow_up_dodge" api:"required"`
	// How reporters would react
	ReporterReaction string `json:"reporter_reaction" api:"required"`
	// Ted's press conference answer
	Response string `json:"response" api:"required"`
	// Humorous deflection if appropriate
	DeflectionHumor string `json:"deflection_humor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualWisdom     respjson.Field
		FollowUpDodge    respjson.Field
		ReporterReaction respjson.Field
		Response         respjson.Field
		DeflectionHumor  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PressSimulateResponse) RawJSON() string { return r.JSON.raw }
func (r *PressSimulateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PressSimulateParams struct {
	// The press question to answer
	Question string `json:"question" api:"required"`
	// Topic category
	Topic param.Opt[string] `json:"topic,omitzero"`
	// Is this a hostile question from Trent Crimm?
	Hostile param.Opt[bool] `json:"hostile,omitzero"`
	paramObj
}

func (r PressSimulateParams) MarshalJSON() (data []byte, err error) {
	type shadow PressSimulateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PressSimulateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
