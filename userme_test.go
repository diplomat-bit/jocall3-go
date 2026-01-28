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

func TestUserMeGet(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithGeminiAPIKey("My Gemini API Key"),
	)
	_, err := client.Users.Me.Get(context.TODO())
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserMeUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithGeminiAPIKey("My Gemini API Key"),
	)
	_, err := client.Users.Me.Update(context.TODO(), githubcomjocall3go.UserMeUpdateParams{
		Address: githubcomjocall3go.F(githubcomjocall3go.UserMeUpdateParamsAddress{
			City:    githubcomjocall3go.F("city"),
			Country: githubcomjocall3go.F("country"),
			State:   githubcomjocall3go.F("state"),
			Street:  githubcomjocall3go.F("street"),
			Zip:     githubcomjocall3go.F("zip"),
		}),
		Name:  githubcomjocall3go.F("Quantum Visionary Pro"),
		Phone: githubcomjocall3go.F("+1-555-999-0000"),
		Preferences: githubcomjocall3go.F(githubcomjocall3go.UserMeUpdateParamsPreferences{
			AIInteractionMode:  githubcomjocall3go.F("aiInteractionMode"),
			DataSharingConsent: githubcomjocall3go.F(true),
			NotificationChannels: githubcomjocall3go.F(githubcomjocall3go.UserMeUpdateParamsPreferencesNotificationChannels{
				Email: githubcomjocall3go.F(true),
				InApp: githubcomjocall3go.F(true),
				Push:  githubcomjocall3go.F(true),
				SMS:   githubcomjocall3go.F(true),
			}),
			PreferredLanguage:   githubcomjocall3go.F("preferredLanguage"),
			Theme:               githubcomjocall3go.F("theme"),
			TransactionGrouping: githubcomjocall3go.F("transactionGrouping"),
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
