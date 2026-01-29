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
func (r *CorporateRiskFraudService) AnalyzeTransaction(ctx context.Context, body CorporateRiskFraudAnalyzeTransactionParams, opts ...option.RequestOption) (res *CorporateRiskFraudAnalyzeTransactionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/analyze"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateRiskFraudAnalyzeTransactionResponse struct {
	Decision  CorporateRiskFraudAnalyzeTransactionResponseDecision `json:"decision"`
	RiskScore int64                                                `json:"riskScore"`
	JSON      corporateRiskFraudAnalyzeTransactionResponseJSON     `json:"-"`
}

// corporateRiskFraudAnalyzeTransactionResponseJSON contains the JSON metadata for
// the struct [CorporateRiskFraudAnalyzeTransactionResponse]
type corporateRiskFraudAnalyzeTransactionResponseJSON struct {
	Decision    apijson.Field
	RiskScore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateRiskFraudAnalyzeTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskFraudAnalyzeTransactionResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskFraudAnalyzeTransactionResponseDecision string

const (
	CorporateRiskFraudAnalyzeTransactionResponseDecisionApprove CorporateRiskFraudAnalyzeTransactionResponseDecision = "APPROVE"
	CorporateRiskFraudAnalyzeTransactionResponseDecisionFlag    CorporateRiskFraudAnalyzeTransactionResponseDecision = "FLAG"
	CorporateRiskFraudAnalyzeTransactionResponseDecisionBlock   CorporateRiskFraudAnalyzeTransactionResponseDecision = "BLOCK"
)

func (r CorporateRiskFraudAnalyzeTransactionResponseDecision) IsKnown() bool {
	switch r {
	case CorporateRiskFraudAnalyzeTransactionResponseDecisionApprove, CorporateRiskFraudAnalyzeTransactionResponseDecisionFlag, CorporateRiskFraudAnalyzeTransactionResponseDecisionBlock:
		return true
	}
	return false
}

type CorporateRiskFraudAnalyzeTransactionParams struct {
	TransactionID param.Field[string] `json:"transactionId,required"`
}

func (r CorporateRiskFraudAnalyzeTransactionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
