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

func TestAIOracleSimulateRunAdvancedWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Oracle.Simulate.RunAdvanced(context.TODO(), githubcomjocall3go.AIOracleSimulateRunAdvancedParams{
		Prompt: githubcomjocall3go.F("prompt"),
		Scenarios: githubcomjocall3go.F([]githubcomjocall3go.AIOracleSimulateRunAdvancedParamsScenario{{
			Name:        githubcomjocall3go.F("name"),
			Description: githubcomjocall3go.F("description"),
			Variables: githubcomjocall3go.F(map[string]interface{}{
				"foo": "bar",
			}),
		}}),
		GlobalEconomicFactors: githubcomjocall3go.F[any](map[string]interface{}{}),
		PersonalAssumptions:   githubcomjocall3go.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIOracleSimulateRunMonteCarloWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Oracle.Simulate.RunMonteCarlo(context.TODO(), githubcomjocall3go.AIOracleSimulateRunMonteCarloParams{
		Iterations:         githubcomjocall3go.F(int64(100)),
		Variables:          githubcomjocall3go.F([]string{"inflation", "oil_prices"}),
		ConfidenceInterval: githubcomjocall3go.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIOracleSimulateRunStandardWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Oracle.Simulate.RunStandard(context.TODO(), githubcomjocall3go.AIOracleSimulateRunStandardParams{
		Prompt:     githubcomjocall3go.F("prompt"),
		Parameters: githubcomjocall3go.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
