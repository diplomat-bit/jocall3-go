// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/diplomat-bit/jocall3-go"
	"github.com/diplomat-bit/jocall3-go/internal/testutil"
	"github.com/diplomat-bit/jocall3-go/option"
)

func TestCorporateAnomalyListWithOptionalParams(t *testing.T) {
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
		option.WithGeminiAPIKey("My Gemini API Key"),
	)
	_, err := client.Corporate.Anomalies.List(context.TODO(), githubcomjocall3go.CorporateAnomalyListParams{
		EndDate:    githubcomjocall3go.F("endDate"),
		EntityType: githubcomjocall3go.F("entityType"),
		Limit:      githubcomjocall3go.F(int64(0)),
		Offset:     githubcomjocall3go.F(int64(0)),
		Severity:   githubcomjocall3go.F("severity"),
		StartDate:  githubcomjocall3go.F("startDate"),
		Status:     githubcomjocall3go.F("status"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateAnomalyUpdateStatus(t *testing.T) {
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
		option.WithGeminiAPIKey("My Gemini API Key"),
	)
	_, err := client.Corporate.Anomalies.UpdateStatus(
		context.TODO(),
		"anom_risk-2024-07-21-D1E2F3",
		githubcomjocall3go.CorporateAnomalyUpdateStatusParams{},
	)
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
