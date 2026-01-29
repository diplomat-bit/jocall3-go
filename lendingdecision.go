// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// LendingDecisionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendingDecisionService] method instead.
type LendingDecisionService struct {
	Options []option.RequestOption
}

// NewLendingDecisionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLendingDecisionService(opts ...option.RequestOption) (r *LendingDecisionService) {
	r = &LendingDecisionService{}
	r.Options = opts
	return
}

// Fetches the deep neural logic behind why a loan was approved or denied.
func (r *LendingDecisionService) GetRationale(ctx context.Context, decisionID string, opts ...option.RequestOption) (res *LendingDecisionGetRationaleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if decisionID == "" {
		err = errors.New("missing required decisionId parameter")
		return
	}
	path := fmt.Sprintf("lending/decisions/%s/rationale", decisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LendingDecisionGetRationaleResponse struct {
	Approved           bool                                    `json:"approved"`
	NextSteps          string                                  `json:"nextSteps"`
	ReasoningNodes     []string                                `json:"reasoningNodes"`
	RiskVectorAnalysis interface{}                             `json:"riskVectorAnalysis"`
	JSON               lendingDecisionGetRationaleResponseJSON `json:"-"`
}

// lendingDecisionGetRationaleResponseJSON contains the JSON metadata for the
// struct [LendingDecisionGetRationaleResponse]
type lendingDecisionGetRationaleResponseJSON struct {
	Approved           apijson.Field
	NextSteps          apijson.Field
	ReasoningNodes     apijson.Field
	RiskVectorAnalysis apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LendingDecisionGetRationaleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lendingDecisionGetRationaleResponseJSON) RawJSON() string {
	return r.raw
}
