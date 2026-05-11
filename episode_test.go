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

func TestEpisodeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Episodes.New(context.TODO(), believe.EpisodeNewParams{
		AirDate:                time.Now(),
		CharacterFocus:         []string{"ted-lasso", "coach-beard", "higgins", "nate"},
		Director:               "MJ Delaney",
		EpisodeNumber:          8,
		MainTheme:              "The power of vulnerability and male friendship",
		RuntimeMinutes:         29,
		Season:                 1,
		Synopsis:               "Ted creates a support group for the coaching staff while Rebecca faces a difficult decision about her future.",
		TedWisdom:              "There's two buttons I never like to hit: that's panic and snooze.",
		Title:                  "The Diamond Dogs",
		Writer:                 "Jason Sudeikis, Brendan Hunt, Joe Kelly",
		BiscuitsWithBossMoment: believe.String("Ted and Rebecca have an honest conversation about trust."),
		MemorableMoments:       []string{"First Diamond Dogs meeting", "The famous dart scene with Rupert", "Be curious, not judgmental speech"},
		UsViewersMillions:      believe.Float(1.42),
		ViewerRating:           believe.Float(9.1),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEpisodeGet(t *testing.T) {
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
	_, err := client.Episodes.Get(context.TODO(), "episode_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEpisodeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Episodes.Update(
		context.TODO(),
		"episode_id",
		believe.EpisodeUpdateParams{
			AirDate:                believe.Time(time.Now()),
			BiscuitsWithBossMoment: believe.String("biscuits_with_boss_moment"),
			CharacterFocus:         []string{"string"},
			Director:               believe.String("director"),
			EpisodeNumber:          believe.Int(1),
			MainTheme:              believe.String("main_theme"),
			MemorableMoments:       []string{"string"},
			RuntimeMinutes:         believe.Int(20),
			Season:                 believe.Int(1),
			Synopsis:               believe.String("synopsis"),
			TedWisdom:              believe.String("ted_wisdom"),
			Title:                  believe.String("x"),
			UsViewersMillions:      believe.Float(0),
			ViewerRating:           believe.Float(0),
			Writer:                 believe.String("writer"),
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

func TestEpisodeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Episodes.List(context.TODO(), believe.EpisodeListParams{
		CharacterFocus: believe.String("character_focus"),
		Limit:          believe.Int(10),
		Season:         believe.Int(1),
		Skip:           believe.Int(0),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEpisodeDelete(t *testing.T) {
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
	err := client.Episodes.Delete(context.TODO(), "episode_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEpisodeGetWisdom(t *testing.T) {
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
	_, err := client.Episodes.GetWisdom(context.TODO(), "episode_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
