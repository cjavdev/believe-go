// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/believe-go/internal/apijson"
	"github.com/stainless-sdks/believe-go/internal/requestconfig"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-go/packages/respjson"
)

// ConflictService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConflictService] method instead.
type ConflictService struct {
	Options []option.RequestOption
}

// NewConflictService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConflictService(opts ...option.RequestOption) (r ConflictService) {
	r = ConflictService{}
	r.Options = opts
	return
}

// Get Ted Lasso-style advice for resolving conflicts.
func (r *ConflictService) Resolve(ctx context.Context, body ConflictResolveParams, opts ...option.RequestOption) (res *ConflictResolveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "conflicts/resolve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Conflict resolution response.
type ConflictResolveResponse struct {
	// A folksy metaphor to remember
	BarbecueSauceWisdom string `json:"barbecue_sauce_wisdom,required"`
	// Understanding the root cause
	Diagnosis string `json:"diagnosis,required"`
	// Advice from the Diamond Dogs support group
	DiamondDogsAdvice string `json:"diamond_dogs_advice,required"`
	// What resolution could look like
	PotentialOutcome string `json:"potential_outcome,required"`
	// Concrete steps to resolve the conflict
	StepsToResolution []string `json:"steps_to_resolution,required"`
	// How Ted would handle this
	TedApproach string `json:"ted_approach,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BarbecueSauceWisdom respjson.Field
		Diagnosis           respjson.Field
		DiamondDogsAdvice   respjson.Field
		PotentialOutcome    respjson.Field
		StepsToResolution   respjson.Field
		TedApproach         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConflictResolveResponse) RawJSON() string { return r.JSON.raw }
func (r *ConflictResolveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConflictResolveParams struct {
	// Type of conflict
	//
	// Any of "interpersonal", "team_dynamics", "leadership", "ego",
	// "miscommunication", "competition".
	ConflictType ConflictResolveParamsConflictType `json:"conflict_type,omitzero,required"`
	// Describe the conflict
	Description string `json:"description,required"`
	// Who is involved in the conflict
	PartiesInvolved []string `json:"parties_involved,omitzero,required"`
	// What you've already tried
	AttemptsMade []string `json:"attempts_made,omitzero"`
	paramObj
}

func (r ConflictResolveParams) MarshalJSON() (data []byte, err error) {
	type shadow ConflictResolveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConflictResolveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of conflict
type ConflictResolveParamsConflictType string

const (
	ConflictResolveParamsConflictTypeInterpersonal    ConflictResolveParamsConflictType = "interpersonal"
	ConflictResolveParamsConflictTypeTeamDynamics     ConflictResolveParamsConflictType = "team_dynamics"
	ConflictResolveParamsConflictTypeLeadership       ConflictResolveParamsConflictType = "leadership"
	ConflictResolveParamsConflictTypeEgo              ConflictResolveParamsConflictType = "ego"
	ConflictResolveParamsConflictTypeMiscommunication ConflictResolveParamsConflictType = "miscommunication"
	ConflictResolveParamsConflictTypeCompetition      ConflictResolveParamsConflictType = "competition"
)
