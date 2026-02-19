// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"github.com/cjavdev/believe-go/option"
)

// ClientWService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientWService] method instead.
type ClientWService struct {
	Options []option.RequestOption
}

// NewClientWService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClientWService(opts ...option.RequestOption) (r ClientWService) {
	r = ClientWService{}
	r.Options = opts
	return
}
