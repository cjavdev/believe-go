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

func TestBelieveSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Believe.Submit(context.TODO(), believe.BelieveSubmitParams{
		Situation:     "I just got passed over for a promotion I've been working toward for two years.",
		SituationType: believe.BelieveSubmitParamsSituationTypeWorkChallenge,
		Context:       believe.String("I've always tried to be a team player and support my colleagues."),
		Intensity:     believe.Int(7),
	})
	if err != nil {
		var apierr *believe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
