// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cjavdev/believe-go"
  "github.com/cjavdev/believe-go/internal/testutil"
  "github.com/cjavdev/believe-go/option"
)

func TestConflictResolveWithOptionalParams(t *testing.T) {
  t.Skip("Mock server tests are disabled")
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
  _, err := client.Conflicts.Resolve(context.TODO(), believe.ConflictResolveParams{
    ConflictType: believe.ConflictResolveParamsConflictTypeInterpersonal,
    Description: "Alex keeps taking credit for my ideas in meetings and I'm getting resentful.",
    PartiesInvolved: []string{"Me", "My teammate Alex"},
    AttemptsMade: []string{"Mentioned it casually", "Avoided them"},
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
