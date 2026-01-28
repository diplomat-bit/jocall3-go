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

func TestUserMePreferenceGet(t *testing.T) {
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
	)
	_, err := client.Users.Me.Preferences.Get(context.TODO())
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserMePreferenceUpdateWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Users.Me.Preferences.Update(context.TODO(), githubcomjocall3go.UserMePreferenceUpdateParams{
		AIInteractionMode:  githubcomjocall3go.F("proactive"),
		DataSharingConsent: githubcomjocall3go.F(true),
		NotificationChannels: githubcomjocall3go.F(githubcomjocall3go.UserMePreferenceUpdateParamsNotificationChannels{
			Email: githubcomjocall3go.F(true),
			InApp: githubcomjocall3go.F(true),
			Push:  githubcomjocall3go.F(true),
			SMS:   githubcomjocall3go.F(true),
		}),
		PreferredLanguage:   githubcomjocall3go.F("preferredLanguage"),
		Theme:               githubcomjocall3go.F("Dark-Quantum"),
		TransactionGrouping: githubcomjocall3go.F("transactionGrouping"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
