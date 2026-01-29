// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// CorporateRiskFraudService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateRiskFraudService] method instead.
type CorporateRiskFraudService struct {
	Options []option.RequestOption
	Rules   *CorporateRiskFraudRuleService
}

// NewCorporateRiskFraudService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateRiskFraudService(opts ...option.RequestOption) (r *CorporateRiskFraudService) {
	r = &CorporateRiskFraudService{}
	r.Options = opts
	r.Rules = NewCorporateRiskFraudRuleService(opts...)
	return
}

// Real-time Transaction Fraud Analysis
func (r *CorporateRiskFraudService) Analyze(ctx context.Context, body CorporateRiskFraudAnalyzeParams, opts ...option.RequestOption) (res *CorporateRiskFraudAnalyzeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/analyze"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateRiskFraudAnalyzeResponse struct {
	Decision  CorporateRiskFraudAnalyzeResponseDecision `json:"decision"`
	RiskScore int64                                     `json:"riskScore"`
	JSON      corporateRiskFraudAnalyzeResponseJSON     `json:"-"`
}

// corporateRiskFraudAnalyzeResponseJSON contains the JSON metadata for the struct
// [CorporateRiskFraudAnalyzeResponse]
type corporateRiskFraudAnalyzeResponseJSON struct {
	Decision    apijson.Field
	RiskScore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateRiskFraudAnalyzeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskFraudAnalyzeResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskFraudAnalyzeResponseDecision string

const (
	CorporateRiskFraudAnalyzeResponseDecisionApprove CorporateRiskFraudAnalyzeResponseDecision = "APPROVE"
	CorporateRiskFraudAnalyzeResponseDecisionFlag    CorporateRiskFraudAnalyzeResponseDecision = "FLAG"
	CorporateRiskFraudAnalyzeResponseDecisionBlock   CorporateRiskFraudAnalyzeResponseDecision = "BLOCK"
)

func (r CorporateRiskFraudAnalyzeResponseDecision) IsKnown() bool {
	switch r {
	case CorporateRiskFraudAnalyzeResponseDecisionApprove, CorporateRiskFraudAnalyzeResponseDecisionFlag, CorporateRiskFraudAnalyzeResponseDecisionBlock:
		return true
	}
	return false
}

type CorporateRiskFraudAnalyzeParams struct {
	TransactionID param.Field[string] `json:"transactionId,required"`
}

func (r CorporateRiskFraudAnalyzeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
