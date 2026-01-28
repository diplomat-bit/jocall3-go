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

func TestUserLogin(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), githubcomjocall3go.UserLoginParams{
		Email:    githubcomjocall3go.F("quantum.visionary@demobank.com"),
		Password: githubcomjocall3go.F("YourSecurePassword123"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserRegisterWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Register(context.TODO(), githubcomjocall3go.UserRegisterParams{
		Email:    githubcomjocall3go.F("alice.w@example.com"),
		Name:     githubcomjocall3go.F("Alice Wonderland"),
		Password: githubcomjocall3go.F("SecureP@ssw0rd2024!"),
		Address: githubcomjocall3go.F(githubcomjocall3go.UserRegisterParamsAddress{
			City:    githubcomjocall3go.F("city"),
			Country: githubcomjocall3go.F("country"),
			State:   githubcomjocall3go.F("state"),
			Street:  githubcomjocall3go.F("street"),
			Zip:     githubcomjocall3go.F("zip"),
		}),
		Phone: githubcomjocall3go.F("+1-555-987-6543"),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
