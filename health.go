// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
  "context"
  "net/http"
  "slices"

  "github.com/cjavdev/believe-go/internal/requestconfig"
  "github.com/cjavdev/believe-go/option"
)

// HealthService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
  r = HealthService{}
  r.Options = opts
  return
}

// Check if the API is running and healthy.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
  opts = slices.Concat(r.Options, opts)
  path := "health"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return res, err
}

type HealthCheckResponse = any
