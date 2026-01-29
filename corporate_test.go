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
	"github.com/diplomat-bit/jocall3-go/shared"
)

func TestCorporateOnboardEntityWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.OnboardEntity(context.TODO(), githubcomjocall3go.CorporateOnboardEntityParams{
		EntityType:   githubcomjocall3go.F(githubcomjocall3go.CorporateOnboardEntityParamsEntityTypeLlc),
		Jurisdiction: githubcomjocall3go.F("DE"),
		LegalName:    githubcomjocall3go.F("legalName"),
		TaxID:        githubcomjocall3go.F("taxId"),
		BeneficialOwners: githubcomjocall3go.F([]githubcomjocall3go.CorporateOnboardEntityParamsBeneficialOwner{{
			ID:               githubcomjocall3go.F("id"),
			Email:            githubcomjocall3go.F("dev@stainless.com"),
			IdentityVerified: githubcomjocall3go.F(true),
			Name:             githubcomjocall3go.F("name"),
			Address: githubcomjocall3go.F(shared.AddressParam{
				City:    githubcomjocall3go.F("city"),
				Country: githubcomjocall3go.F("country"),
				Street:  githubcomjocall3go.F("street"),
				State:   githubcomjocall3go.F("state"),
				Zip:     githubcomjocall3go.F("zip"),
			}),
			Preferences: githubcomjocall3go.F(map[string]interface{}{
				"foo": "bar",
			}),
			SecurityStatus: githubcomjocall3go.F(githubcomjocall3go.CorporateOnboardEntityParamsBeneficialOwnersSecurityStatus{
				LastLogin:        githubcomjocall3go.F(time.Now()),
				TwoFactorEnabled: githubcomjocall3go.F(true),
			}),
		}}),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
