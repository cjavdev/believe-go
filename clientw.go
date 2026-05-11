// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"net/http"
	"slices"

	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
)

// WebSocket endpoints for real-time bidirectional communication - Live match
// simulation
//
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

// Simple WebSocket test endpoint for connectivity testing.
//
// Connect to test WebSocket functionality. The server will:
//
// 1. Send a welcome message on connection
// 2. Echo back any message you send
//
// ## Example
//
// ```javascript
// const ws = new WebSocket("ws://localhost:8000/ws/test");
// ws.onmessage = (event) => console.log(event.data);
// ws.send("Hello!"); // Server responds with echo
// ```
func (r *ClientWService) Test(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "ws/test"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}
