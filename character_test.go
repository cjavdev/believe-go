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

func TestCharacterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Characters.New(context.TODO(), believe.CharacterNewParams{
		Background: "Legendary midfielder for Chelsea and AFC Richmond, now assistant coach. Known for his gruff exterior hiding a heart of gold.",
		EmotionalStats: believe.EmotionalStatsParam{
			Curiosity:     40,
			Empathy:       85,
			Optimism:      45,
			Resilience:    95,
			Vulnerability: 60,
		},
		Name:              "Roy Kent",
		PersonalityTraits: []string{"intense", "loyal", "secretly caring", "profane"},
		Role:              believe.CharacterRoleCoach,
		DateOfBirth:       believe.Time(time.Now()),
		Email:             believe.String("roy.kent@afcrichmond.com"),
		GrowthArcs: []believe.GrowthArcParam{{
			Breakthrough:  "Finding purpose beyond playing",
			Challenge:     "Accepting his body's limitations",
			EndingPoint:   "Retired but lost",
			Season:        1,
			StartingPoint: "Aging footballer facing retirement",
		}},
		HeightMeters:    believe.Float(1.78),
		ProfileImageURL: believe.String("https://afcrichmond.com/images/roy-kent.jpg"),
		SalaryGbp: believe.CharacterNewParamsSalaryGbpUnion{
			OfString: believe.String("175000.00"),
		},
		SignatureQuotes: []string{"He's here, he's there, he's every-f***ing-where, Roy Kent!", "Whistle!"},
		TeamID:          believe.String("afc-richmond"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCharacterGet(t *testing.T) {
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
	_, err := client.Characters.Get(context.TODO(), "character_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCharacterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Characters.Update(
		context.TODO(),
		"character_id",
		believe.CharacterUpdateParams{
			Background:  believe.String("background"),
			DateOfBirth: believe.Time(time.Now()),
			Email:       believe.String("dev@stainless.com"),
			EmotionalStats: believe.EmotionalStatsParam{
				Curiosity:     99,
				Empathy:       100,
				Optimism:      95,
				Resilience:    90,
				Vulnerability: 80,
			},
			GrowthArcs: []believe.GrowthArcParam{{
				Breakthrough:  "breakthrough",
				Challenge:     "challenge",
				EndingPoint:   "ending_point",
				Season:        1,
				StartingPoint: "starting_point",
			}},
			HeightMeters:      believe.Float(1),
			Name:              believe.String("x"),
			PersonalityTraits: []string{"string"},
			ProfileImageURL:   believe.String("https://example.com"),
			Role:              believe.CharacterRoleCoach,
			SalaryGbp: believe.CharacterUpdateParamsSalaryGbpUnion{
				OfFloat: believe.Float(0),
			},
			SignatureQuotes: []string{"string"},
			TeamID:          believe.String("team_id"),
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

func TestCharacterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Characters.List(context.TODO(), believe.CharacterListParams{
		Limit:       believe.Int(10),
		MinOptimism: believe.Int(0),
		Role:        believe.CharacterRoleCoach,
		Skip:        believe.Int(0),
		TeamID:      believe.String("team_id"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCharacterDelete(t *testing.T) {
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
	err := client.Characters.Delete(context.TODO(), "character_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCharacterGetQuotes(t *testing.T) {
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
	_, err := client.Characters.GetQuotes(context.TODO(), "character_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
