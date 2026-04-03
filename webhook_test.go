// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
  "context"
  "errors"
  "os"
  "testing"
  "time"

  "github.com/cjavdev/believe-go"
  "github.com/cjavdev/believe-go/internal/testutil"
  "github.com/cjavdev/believe-go/option"
)

func TestWebhookNewWithOptionalParams(t *testing.T) {
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
  _, err := client.Webhooks.New(context.TODO(), believe.WebhookNewParams{
    URL: "https://example.com/webhooks",
    Description: believe.String("Production webhook for match notifications"),
    EventTypes: []string{"match.completed", "team_member.transferred"},
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestWebhookGet(t *testing.T) {
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
  _, err := client.Webhooks.Get(context.TODO(), "webhook_id")
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestWebhookList(t *testing.T) {
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
  _, err := client.Webhooks.List(context.TODO())
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestWebhookDelete(t *testing.T) {
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
  _, err := client.Webhooks.Delete(context.TODO(), "webhook_id")
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestWebhookTriggerEventWithOptionalParams(t *testing.T) {
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
  _, err := client.Webhooks.TriggerEvent(context.TODO(), believe.WebhookTriggerEventParams{
    EventType: believe.WebhookTriggerEventParamsEventTypeMatchCompleted,
    Payload: believe.WebhookTriggerEventParamsPayloadUnion{
      OfMatchCompleted: &believe.WebhookTriggerEventParamsPayloadMatchCompleted{
        Data: believe.WebhookTriggerEventParamsPayloadMatchCompletedData{
          AwayScore: 0,
          AwayTeamID: "away_team_id",
          CompletedAt: time.Now(),
          HomeScore: 0,
          HomeTeamID: "home_team_id",
          MatchID: "match_id",
          MatchType: "league",
          Result: "home_win",
          TedPostMatchQuote: "ted_post_match_quote",
          LessonLearned: believe.String("lesson_learned"),
          ManOfTheMatch: believe.String("man_of_the_match"),
        },
        EventType: "match.completed",
      },
    },
  })
  if err != nil {
    var apierr *believe.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
