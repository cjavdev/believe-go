// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"github.com/cjavdev/believe-go/option"
)

// CoachingService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoachingService] method instead.
type CoachingService struct {
	Options    []option.RequestOption
	Principles CoachingPrincipleService
}

// NewCoachingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoachingService(opts ...option.RequestOption) (r CoachingService) {
	r = CoachingService{}
	r.Options = opts
	r.Principles = NewCoachingPrincipleService(opts...)
	return
}
