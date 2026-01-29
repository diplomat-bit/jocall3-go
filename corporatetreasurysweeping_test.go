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

func TestCorporateTreasurySweepingConfigureRulesWithOptionalParams(t *testing.T) {
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
	err := client.Corporate.Treasury.Sweeping.ConfigureRules(context.TODO(), githubcomjocall3go.CorporateTreasurySweepingConfigureRulesParams{
		SourceAccount: githubcomjocall3go.F("sourceAccount"),
		TargetAccount: githubcomjocall3go.F("targetAccount"),
		Threshold:     githubcomjocall3go.F(0.000000),
		Frequency:     githubcomjocall3go.F(githubcomjocall3go.CorporateTreasurySweepingConfigureRulesParamsFrequencyDaily),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateTreasurySweepingExecute(t *testing.T) {
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
	err := client.Corporate.Treasury.Sweeping.Execute(context.TODO(), githubcomjocall3go.CorporateTreasurySweepingExecuteParams{
		RuleID: githubcomjocall3go.F("ruleId"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
