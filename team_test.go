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

func TestTeamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Teams.New(context.TODO(), believe.TeamNewParams{
		CultureScore: 70,
		FoundedYear:  1895,
		League:       believe.LeaguePremierLeague,
		Name:         "West Ham United",
		Stadium:      "London Stadium",
		Values: believe.TeamValuesParam{
			PrimaryValue:    "Pride",
			SecondaryValues: []string{"History", "Community", "Passion"},
			TeamMotto:       "Forever Blowing Bubbles",
		},
		AnnualBudgetGbp: believe.TeamNewParamsAnnualBudgetGbpUnion{
			OfString: believe.String("150000000.00"),
		},
		AverageAttendance: believe.Float(59988),
		ContactEmail:      believe.String("info@westhamunited.co.uk"),
		IsActive:          believe.Bool(true),
		Nickname:          believe.String("The Hammers"),
		PrimaryColor:      believe.String("#7A263A"),
		RivalTeams:        []string{"afc-richmond", "tottenham"},
		SecondaryColor:    believe.String("#1BB1E7"),
		StadiumLocation: believe.GeoLocationParam{
			Latitude:  51.5387,
			Longitude: -0.0166,
		},
		Website:       believe.String("https://www.whufc.com"),
		WinPercentage: believe.Float(52.3),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamGet(t *testing.T) {
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
	_, err := client.Teams.Get(context.TODO(), "team_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Teams.Update(
		context.TODO(),
		"team_id",
		believe.TeamUpdateParams{
			AnnualBudgetGbp: believe.TeamUpdateParamsAnnualBudgetGbpUnion{
				OfFloat: believe.Float(0),
			},
			AverageAttendance: believe.Float(0),
			ContactEmail:      believe.String("dev@stainless.com"),
			CultureScore:      believe.Int(0),
			FoundedYear:       believe.Int(1800),
			IsActive:          believe.Bool(true),
			League:            believe.LeaguePremierLeague,
			Name:              believe.String("x"),
			Nickname:          believe.String("nickname"),
			PrimaryColor:      believe.String("primary_color"),
			RivalTeams:        []string{"string"},
			SecondaryColor:    believe.String("secondary_color"),
			Stadium:           believe.String("stadium"),
			StadiumLocation: believe.GeoLocationParam{
				Latitude:  51.4816,
				Longitude: -0.191,
			},
			Values: believe.TeamValuesParam{
				PrimaryValue:    "Believe",
				SecondaryValues: []string{"Family", "Resilience", "Joy"},
				TeamMotto:       "Football is life!",
			},
			Website:       believe.String("https://example.com"),
			WinPercentage: believe.Float(0),
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

func TestTeamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Teams.List(context.TODO(), believe.TeamListParams{
		League:          believe.LeaguePremierLeague,
		Limit:           believe.Int(10),
		MinCultureScore: believe.Int(0),
		Skip:            believe.Int(0),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamDelete(t *testing.T) {
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
	err := client.Teams.Delete(context.TODO(), "team_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamGetCulture(t *testing.T) {
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
	_, err := client.Teams.GetCulture(context.TODO(), "team_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamGetRivals(t *testing.T) {
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
	_, err := client.Teams.GetRivals(context.TODO(), "team_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamListLogos(t *testing.T) {
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
	_, err := client.Teams.ListLogos(context.TODO(), "team_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
