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
		Prompt: githubcomjocall3go.F("Analyze the systemic risk of a 20% drop in BTC prices on my cross-chain collateralized debt positions, factoring in a simultaneous 50bps hike by the Fed and a liquidity squeeze on Aave."),
		Scenarios: githubcomjocall3go.F([]githubcomjocall3go.AIOracleSimulateRunAdvancedParamsScenario{{
			Name: githubcomjocall3go.F("Crypto Black Swan + Macro Contagion"),
		}}),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.SimulationID)
}
