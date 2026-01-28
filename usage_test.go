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
	)
	t.Skip("Prism tests are disabled")
	response, err := client.Users.Register(context.TODO(), githubcomjocall3go.UserRegisterParams{
		Email:    githubcomjocall3go.F("executive@corp.com"),
		Name:     githubcomjocall3go.F("Alice Wonderland"),
		Password: githubcomjocall3go.F("ComplexPassword99!"),
		Phone:    githubcomjocall3go.F("+1-555-0199"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.ID)
}
