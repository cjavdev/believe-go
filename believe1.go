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
// BelieveService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBelieveService] method instead.
type BelieveService struct {
Options []option.RequestOption
}

// NewBelieveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBelieveService(opts ...option.RequestOption) (r BelieveService) {
  r = BelieveService{}
  r.Options = opts
  return
}

// Submit your situation and receive Ted Lasso-style motivational guidance.
func (r *BelieveService) Submit(ctx context.Context, body BelieveSubmitParams, opts ...option.RequestOption) (res *BelieveSubmitResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "believe"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return res, err
}

// Response from the Believe Engine.
type BelieveSubmitResponse struct {
// Suggested action to take
ActionSuggestion string `json:"action_suggestion" api:"required"`
// Your current believe-o-meter score
BelieveScore int64 `json:"believe_score" api:"required"`
// A reminder to have a goldfish memory when needed
GoldfishWisdom string `json:"goldfish_wisdom" api:"required"`
// A relevant Ted Lasso quote
RelevantQuote string `json:"relevant_quote" api:"required"`
// Ted's motivational response
TedResponse string `json:"ted_response" api:"required"`
// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
JSON struct {
              ActionSuggestion respjson.Field
              BelieveScore respjson.Field
              GoldfishWisdom respjson.Field
              RelevantQuote respjson.Field
              TedResponse respjson.Field
              ExtraFields map[string]respjson.Field
              raw string
            } `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BelieveSubmitResponse) RawJSON() (string) { return r.JSON.raw }
func (r *BelieveSubmitResponse) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

type BelieveSubmitParams struct {
// Describe your situation
Situation string `json:"situation" api:"required"`
// Type of situation
//
// Any of "work_challenge", "personal_setback", "team_conflict", "self_doubt",
// "big_decision", "failure", "new_beginning", "relationship".
SituationType BelieveSubmitParamsSituationType `json:"situation_type,omitzero" api:"required"`
// Additional context
Context param.Opt[string] `json:"context,omitzero"`
// How intense is the response needed (1=gentle, 10=full Ted)
Intensity param.Opt[int64] `json:"intensity,omitzero"`
paramObj
}

func (r BelieveSubmitParams) MarshalJSON() (data []byte, err error) {
  type shadow BelieveSubmitParams
  return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BelieveSubmitParams) UnmarshalJSON(data []byte) (error) {
  return apijson.UnmarshalRoot(data, r)
}

// Type of situation
type BelieveSubmitParamsSituationType string

const (
    BelieveSubmitParamsSituationTypeWorkChallenge BelieveSubmitParamsSituationType = "work_challenge"
    BelieveSubmitParamsSituationTypePersonalSetback BelieveSubmitParamsSituationType = "personal_setback"
    BelieveSubmitParamsSituationTypeTeamConflict BelieveSubmitParamsSituationType = "team_conflict"
    BelieveSubmitParamsSituationTypeSelfDoubt BelieveSubmitParamsSituationType = "self_doubt"
    BelieveSubmitParamsSituationTypeBigDecision BelieveSubmitParamsSituationType = "big_decision"
    BelieveSubmitParamsSituationTypeFailure BelieveSubmitParamsSituationType = "failure"
    BelieveSubmitParamsSituationTypeNewBeginning BelieveSubmitParamsSituationType = "new_beginning"
    BelieveSubmitParamsSituationTypeRelationship BelieveSubmitParamsSituationType = "relationship"
  )
