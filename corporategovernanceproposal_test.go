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

func TestCorporateGovernanceProposalNewWithOptionalParams(t *testing.T) {
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
	err := client.Corporate.Governance.Proposals.New(context.TODO(), githubcomjocall3go.CorporateGovernanceProposalNewParams{
		ActionType:        githubcomjocall3go.F(githubcomjocall3go.CorporateGovernanceProposalNewParamsActionTypeTransferLimitChange),
		Payload:           githubcomjocall3go.F[any](map[string]interface{}{}),
		Title:             githubcomjocall3go.F("title"),
		Description:       githubcomjocall3go.F("description"),
		VotingPeriodHours: githubcomjocall3go.F(int64(0)),
	})
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateGovernanceProposalList(t *testing.T) {
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
	_, err := client.Corporate.Governance.Proposals.List(context.TODO())
	if err != nil {
		var apierr *githubcomjocall3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateGovernanceProposalVoteWithOptionalParams(t *testing.T) {
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
	err := client.Corporate.Governance.Proposals.Vote(
		context.TODO(),
		"proposalId",
		githubcomjocall3go.CorporateGovernanceProposalVoteParams{
			Decision:  githubcomjocall3go.F(githubcomjocall3go.CorporateGovernanceProposalVoteParamsDecisionApprove),
			Comment:   githubcomjocall3go.F("comment"),
			Signature: githubcomjocall3go.F("signature"),
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
