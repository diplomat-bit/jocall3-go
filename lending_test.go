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

func TestLendingGetStatus(t *testing.T) {
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
	_, err := client.Lending.GetStatus(context.TODO(), "appId")
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLendingSubmitApplicationWithOptionalParams(t *testing.T) {
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
	_, err := client.Lending.SubmitApplication(context.TODO(), githubcomjocall3go.LendingSubmitApplicationParams{
		Amount: githubcomjocall3go.F(0.000000),
		EmploymentData: githubcomjocall3go.F(githubcomjocall3go.LendingSubmitApplicationParamsEmploymentData{
			Employer:      githubcomjocall3go.F("employer"),
			MonthlyIncome: githubcomjocall3go.F(0.000000),
			TenureMonths:  githubcomjocall3go.F(int64(0)),
		}),
		LoanType:     githubcomjocall3go.F(githubcomjocall3go.LendingSubmitApplicationParamsLoanTypeMortgage),
		TermMonths:   githubcomjocall3go.F(int64(0)),
		Assets:       githubcomjocall3go.F([]interface{}{map[string]interface{}{}}),
		CollateralID: githubcomjocall3go.F("collateralId"),
		Liabilities:  githubcomjocall3go.F([]interface{}{map[string]interface{}{}}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
