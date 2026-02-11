// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
	"context"
	"os"
	"testing"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/internal/testutil"
	"github.com/cjavdev/believe-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := believe.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	page, err := client.Characters.List(context.TODO(), believe.CharacterListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", page)
}
