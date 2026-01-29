// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/diplomat-bit/jocall3-go"
	"github.com/diplomat-bit/jocall3-go/internal/testutil"
	"github.com/diplomat-bit/jocall3-go/option"
)

func TestCorporateComplianceScreenAdverseMediaWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Compliance.ScreenAdverseMedia(context.TODO(), githubcomjocall3go.CorporateComplianceScreenAdverseMediaParams{
		Query: githubcomjocall3go.F("query"),
		Depth: githubcomjocall3go.F(githubcomjocall3go.CorporateComplianceScreenAdverseMediaParamsDepthShallow),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateComplianceScreenPepWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Compliance.ScreenPep(context.TODO(), githubcomjocall3go.CorporateComplianceScreenPepParams{
		FullName: githubcomjocall3go.F("fullName"),
		Dob:      githubcomjocall3go.F(time.Now()),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateComplianceScreenSanctionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Compliance.ScreenSanctions(context.TODO(), githubcomjocall3go.CorporateComplianceScreenSanctionsParams{
		Entities: githubcomjocall3go.F([]githubcomjocall3go.CorporateComplianceScreenSanctionsParamsEntity{{
			Country: githubcomjocall3go.F("country"),
			Name:    githubcomjocall3go.F("name"),
		}}),
		CheckType: githubcomjocall3go.F(githubcomjocall3go.CorporateComplianceScreenSanctionsParamsCheckTypeStandard),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
