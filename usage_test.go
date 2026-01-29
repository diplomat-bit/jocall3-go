// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go_test

import (
	"context"
	"os"
	"testing"

	"github.com/diplomat-bit/jocall3-go"
	"github.com/diplomat-bit/jocall3-go/internal/testutil"
	"github.com/diplomat-bit/jocall3-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.AI.Oracle.Simulate.RunAdvanced(context.TODO(), githubcomjocall3go.AIOracleSimulateRunAdvancedParams{
		Prompt: githubcomjocall3go.F("Analyze systemic risk of a 20% BTC drop."),
		Scenarios: githubcomjocall3go.F([]githubcomjocall3go.AIOracleSimulateRunAdvancedParamsScenario{{
			Name:        githubcomjocall3go.F("Crypto Black Swan"),
			Description: githubcomjocall3go.F("Extreme market volatility scenario."),
		}}),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.SimulationID)
}
