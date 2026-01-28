// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/jocall3-go"
	"github.com/stainless-sdks/jocall3-go/internal/testutil"
	"github.com/stainless-sdks/jocall3-go/option"
)

func TestCorporateRiskFraudListRulesWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomjocall3go.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithGeminiAPIKey("My Gemini API Key"),
	)
	_, err := client.Corporate.Risk.Fraud.ListRules(context.TODO(), githubcomjocall3go.CorporateRiskFraudListRulesParams{
		Limit:  githubcomjocall3go.F(int64(0)),
		Offset: githubcomjocall3go.F(int64(0)),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
