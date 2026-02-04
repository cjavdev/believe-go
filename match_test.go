// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/internal/testutil"
	"github.com/stainless-sdks/believe-go/option"
)

func TestMatchNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.New(context.TODO(), believe.MatchNewParams{
		AwayTeamID:           "tottenham",
		Date:                 time.Now(),
		HomeTeamID:           "afc-richmond",
		MatchType:            believe.MatchTypeCup,
		Attendance:           believe.Int(24500),
		AwayScore:            believe.Int(0),
		EpisodeID:            believe.String("s02e05"),
		HomeScore:            believe.Int(0),
		LessonLearned:        believe.String("It's not about the wins and losses, it's about helping these young fellas be the best versions of themselves."),
		PossessionPercentage: believe.Float(50),
		Result:               believe.MatchResultPending,
		TedHalftimeSpeech:    believe.String("You know what the happiest animal on Earth is? It's a goldfish. You know why? It's got a 10-second memory."),
		TicketRevenueGbp: believe.MatchNewParamsTicketRevenueGbpUnion{
			OfString: believe.String("735000.00"),
		},
		TurningPoints: []believe.TurningPointParam{{
			Description:       "description",
			EmotionalImpact:   "Galvanized the team's fighting spirit",
			Minute:            0,
			CharacterInvolved: believe.String("jamie-tartt"),
		}},
		WeatherTempCelsius: believe.Float(8.5),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.Get(context.TODO(), "match_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.Update(
		context.TODO(),
		"match_id",
		believe.MatchUpdateParams{
			Attendance:           believe.Int(0),
			AwayScore:            believe.Int(0),
			AwayTeamID:           believe.String("away_team_id"),
			Date:                 believe.Time(time.Now()),
			EpisodeID:            believe.String("episode_id"),
			HomeScore:            believe.Int(0),
			HomeTeamID:           believe.String("home_team_id"),
			LessonLearned:        believe.String("lesson_learned"),
			MatchType:            believe.MatchTypeLeague,
			PossessionPercentage: believe.Float(0),
			Result:               believe.MatchResultWin,
			TedHalftimeSpeech:    believe.String("ted_halftime_speech"),
			TicketRevenueGbp: believe.MatchUpdateParamsTicketRevenueGbpUnion{
				OfFloat: believe.Float(0),
			},
			TurningPoints: []believe.TurningPointParam{{
				Description:       "description",
				EmotionalImpact:   "Galvanized the team's fighting spirit",
				Minute:            0,
				CharacterInvolved: believe.String("jamie-tartt"),
			}},
			WeatherTempCelsius: believe.Float(-30),
		},
	)
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.List(context.TODO(), believe.MatchListParams{
		Limit:     believe.Int(10),
		MatchType: believe.MatchTypeLeague,
		Result:    believe.MatchResultWin,
		Skip:      believe.Int(0),
		TeamID:    believe.String("team_id"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.Matches.Delete(context.TODO(), "match_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchGetLesson(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.GetLesson(context.TODO(), "match_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchGetTurningPoints(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Matches.GetTurningPoints(context.TODO(), "match_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatchStreamLiveWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	err := client.Matches.StreamLive(context.TODO(), believe.MatchStreamLiveParams{
		AwayTeam:        believe.String("away_team"),
		ExcitementLevel: believe.Int(1),
		HomeTeam:        believe.String("home_team"),
		Speed:           believe.Float(0.1),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
