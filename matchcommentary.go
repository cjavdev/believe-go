// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
)

// MatchCommentaryService contains methods and other services that help with
// interacting with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatchCommentaryService] method instead.
type MatchCommentaryService struct {
	Options []option.RequestOption
}

// NewMatchCommentaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatchCommentaryService(opts ...option.RequestOption) (r MatchCommentaryService) {
	r = MatchCommentaryService{}
	r.Options = opts
	return
}

// Stream live match commentary for a specific match. Uses Server-Sent Events (SSE)
// to stream commentary events in real-time.
func (r *MatchCommentaryService) Stream(ctx context.Context, matchID string, opts ...option.RequestOption) (res *MatchCommentaryStreamResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if matchID == "" {
		err = errors.New("missing required match_id parameter")
		return
	}
	path := fmt.Sprintf("matches/%s/commentary/stream", url.PathEscape(matchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type MatchCommentaryStreamResponse = any
