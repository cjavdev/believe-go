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

func TestBiscuitGet(t *testing.T) {
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
  _, err := client.Biscuits.Get(context.TODO(), "biscuit_id")
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestBiscuitListWithOptionalParams(t *testing.T) {
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
  _, err := client.Biscuits.List(context.TODO(), believe.BiscuitListParams{
    Limit: believe.Int(10),
    Skip: believe.Int(0),
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestBiscuitGetFresh(t *testing.T) {
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
  _, err := client.Biscuits.GetFresh(context.TODO())
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
