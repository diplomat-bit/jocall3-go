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

func TestAIIncubatorAnalysisCompetitorScan(t *testing.T) {
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
	)
	_, err := client.AI.Incubator.Analysis.CompetitorScan(context.TODO(), githubcomjocall3go.AIIncubatorAnalysisCompetitorScanParams{
		Industry: githubcomjocall3go.F("industry"),
		Niche:    githubcomjocall3go.F("niche"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIIncubatorAnalysisSwot(t *testing.T) {
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
	)
	_, err := client.AI.Incubator.Analysis.Swot(context.TODO(), githubcomjocall3go.AIIncubatorAnalysisSwotParams{
		BusinessContext: githubcomjocall3go.F("businessContext"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
