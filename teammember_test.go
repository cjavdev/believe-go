// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/internal/testutil"
	"github.com/stainless-sdks/believe-go/option"
)

func TestTeamMemberNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.New(context.TODO(), believe.TeamMemberNewParams{
		OfPlayer: &believe.TeamMemberNewParamsBodyPlayer{
			CharacterID:   "jamie-tartt",
			JerseyNumber:  9,
			Position:      believe.PositionForward,
			TeamID:        "afc-richmond",
			YearsWithTeam: 3,
			Assists:       believe.Int(23),
			GoalsScored:   believe.Int(47),
			IsCaptain:     believe.Bool(false),
			MemberType:    "player",
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

func TestTeamMemberGet(t *testing.T) {
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
	_, err := client.TeamMembers.Get(context.TODO(), "member_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamMemberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.Update(
		context.TODO(),
		"member_id",
		believe.TeamMemberUpdateParams{
			OfPlayerUpdate: &believe.TeamMemberUpdateParamsBodyPlayerUpdate{
				Assists:       believe.Int(0),
				GoalsScored:   believe.Int(0),
				IsCaptain:     believe.Bool(true),
				JerseyNumber:  believe.Int(1),
				Position:      believe.PositionGoalkeeper,
				TeamID:        believe.String("team_id"),
				YearsWithTeam: believe.Int(0),
			},
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

func TestTeamMemberListWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.List(context.TODO(), believe.TeamMemberListParams{
		Limit:      believe.Int(10),
		MemberType: believe.TeamMemberListParamsMemberTypePlayer,
		Skip:       believe.Int(0),
		TeamID:     believe.String("team_id"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamMemberDelete(t *testing.T) {
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
	err := client.TeamMembers.Delete(context.TODO(), "member_id")
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamMemberListCoachesWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.ListCoaches(context.TODO(), believe.TeamMemberListCoachesParams{
		Limit:     believe.Int(10),
		Skip:      believe.Int(0),
		Specialty: believe.CoachSpecialtyHeadCoach,
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

func TestTeamMemberListPlayersWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.ListPlayers(context.TODO(), believe.TeamMemberListPlayersParams{
		Limit:    believe.Int(10),
		Position: believe.PositionGoalkeeper,
		Skip:     believe.Int(0),
		TeamID:   believe.String("team_id"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamMemberListStaffWithOptionalParams(t *testing.T) {
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
	_, err := client.TeamMembers.ListStaff(context.TODO(), believe.TeamMemberListStaffParams{
		Limit:  believe.Int(10),
		Skip:   believe.Int(0),
		TeamID: believe.String("team_id"),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
