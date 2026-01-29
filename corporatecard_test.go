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

func TestCorporateCardFreeze(t *testing.T) {
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
	err := client.Corporate.Cards.Freeze(
		context.TODO(),
		"cardId",
		githubcomjocall3go.CorporateCardFreezeParams{
			Frozen: githubcomjocall3go.F(true),
		},
	)
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateCardIssuePhysicalWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.IssuePhysical(context.TODO(), githubcomjocall3go.CorporateCardIssuePhysicalParams{
		HolderName: githubcomjocall3go.F("holderName"),
		ShippingAddress: githubcomjocall3go.F(githubcomjocall3go.CorporateCardIssuePhysicalParamsShippingAddress{
			City:    githubcomjocall3go.F("city"),
			Country: githubcomjocall3go.F("country"),
			State:   githubcomjocall3go.F("state"),
			Street:  githubcomjocall3go.F("street"),
			Zip:     githubcomjocall3go.F("zip"),
		}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateCardIssueVirtualWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.IssueVirtual(context.TODO(), githubcomjocall3go.CorporateCardIssueVirtualParams{
		HolderName:   githubcomjocall3go.F("holderName"),
		MonthlyLimit: githubcomjocall3go.F(0.000000),
		Purpose:      githubcomjocall3go.F("purpose"),
		Metadata:     githubcomjocall3go.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
