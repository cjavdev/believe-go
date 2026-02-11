// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/apiquery"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// PepTalkService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPepTalkService] method instead.
type PepTalkService struct {
	Options []option.RequestOption
}

// NewPepTalkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPepTalkService(opts ...option.RequestOption) (r PepTalkService) {
	r = PepTalkService{}
	r.Options = opts
	return
}

// Get a motivational pep talk from Ted Lasso himself. By default returns the
// complete pep talk. Add `?stream=true` to get Server-Sent Events (SSE) streaming
// the pep talk chunk by chunk.
func (r *PepTalkService) Get(ctx context.Context, query PepTalkGetParams, opts ...option.RequestOption) (res *PepTalkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pep-talk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A complete pep talk response.
type PepTalkGetResponse struct {
	// Individual chunks of the pep talk
	Chunks []PepTalkGetResponseChunk `json:"chunks,required"`
	// The full pep talk text
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chunks      respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PepTalkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PepTalkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk of a streaming pep talk from Ted.
type PepTalkGetResponseChunk struct {
	// Chunk sequence number
	ChunkID int64 `json:"chunk_id,required"`
	// Is this the final chunk
	IsFinal bool `json:"is_final,required"`
	// The text of this chunk
	Text string `json:"text,required"`
	// The emotional purpose of this chunk (e.g., greeting, acknowledgment, wisdom,
	// affirmation, encouragement)
	EmotionalBeat string `json:"emotional_beat,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID       respjson.Field
		IsFinal       respjson.Field
		Text          respjson.Field
		EmotionalBeat respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PepTalkGetResponseChunk) RawJSON() string { return r.JSON.raw }
func (r *PepTalkGetResponseChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PepTalkGetParams struct {
	// If true, returns SSE stream instead of full response
	Stream param.Opt[bool] `query:"stream,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PepTalkGetParams]'s query parameters as `url.Values`.
func (r PepTalkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
