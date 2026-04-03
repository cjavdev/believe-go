// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
  "context"
  "net/http"
  "slices"

  "github.com/cjavdev/believe-go/internal/requestconfig"
  "github.com/cjavdev/believe-go/option"
)

// Server-Sent Events (SSE) streaming endpoints
//
// StreamService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
Options []option.RequestOption
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r StreamService) {
  r = StreamService{}
  r.Options = opts
  return
}

// A simple SSE test endpoint that streams numbers 1-5.
func (r *StreamService) TestConnection(ctx context.Context, opts ...option.RequestOption) (res *StreamTestConnectionResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "stream/test"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return res, err
}

type StreamTestConnectionResponse = any
