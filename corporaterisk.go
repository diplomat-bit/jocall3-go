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

// CorporateRiskService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateRiskService] method instead.
type CorporateRiskService struct {
	Options []option.RequestOption
	Fraud   *CorporateRiskFraudService
}

// NewCorporateRiskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateRiskService(opts ...option.RequestOption) (r *CorporateRiskService) {
	r = &CorporateRiskService{}
	r.Options = opts
	r.Fraud = NewCorporateRiskFraudService(opts...)
	return
}

// Get Aggregate Risk Exposure
func (r *CorporateRiskService) GetExposure(ctx context.Context, opts ...option.RequestOption) (res *CorporateRiskGetExposureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/exposure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulates extreme market conditions (e.g., 2008 crash, hyperinflation) against
// the corporate ledger.
func (r *CorporateRiskService) StressTest(ctx context.Context, body CorporateRiskStressTestParams, opts ...option.RequestOption) (res *CorporateRiskStressTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/stress-test"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateRiskGetExposureResponse struct {
	AssetConcentration interface{}                          `json:"assetConcentration"`
	CounterpartyRisk   []interface{}                        `json:"counterpartyRisk"`
	ValueAtRisk        float64                              `json:"valueAtRisk"`
	JSON               corporateRiskGetExposureResponseJSON `json:"-"`
}

// corporateRiskGetExposureResponseJSON contains the JSON metadata for the struct
// [CorporateRiskGetExposureResponse]
type corporateRiskGetExposureResponseJSON struct {
	AssetConcentration apijson.Field
	CounterpartyRisk   apijson.Field
	ValueAtRisk        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CorporateRiskGetExposureResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskGetExposureResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskStressTestResponse struct {
	AINarrative     string                              `json:"aiNarrative"`
	LiquidityGap    float64                             `json:"liquidityGap"`
	ResilienceScore float64                             `json:"resilienceScore"`
	JSON            corporateRiskStressTestResponseJSON `json:"-"`
}

// corporateRiskStressTestResponseJSON contains the JSON metadata for the struct
// [CorporateRiskStressTestResponse]
type corporateRiskStressTestResponseJSON struct {
	AINarrative     apijson.Field
	LiquidityGap    apijson.Field
	ResilienceScore apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CorporateRiskStressTestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskStressTestResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskStressTestParams struct {
	ScenarioType param.Field[CorporateRiskStressTestParamsScenarioType] `json:"scenarioType,required"`
	Intensity    param.Field[float64]                                   `json:"intensity"`
}

func (r CorporateRiskStressTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateRiskStressTestParamsScenarioType string

const (
	CorporateRiskStressTestParamsScenarioTypeBankRun         CorporateRiskStressTestParamsScenarioType = "BANK_RUN"
	CorporateRiskStressTestParamsScenarioTypeMarketCrash     CorporateRiskStressTestParamsScenarioType = "MARKET_CRASH"
	CorporateRiskStressTestParamsScenarioTypeRegulatoryShock CorporateRiskStressTestParamsScenarioType = "REGULATORY_SHOCK"
	CorporateRiskStressTestParamsScenarioTypeDepegging       CorporateRiskStressTestParamsScenarioType = "DEPEGGING"
)

func (r CorporateRiskStressTestParamsScenarioType) IsKnown() bool {
	switch r {
	case CorporateRiskStressTestParamsScenarioTypeBankRun, CorporateRiskStressTestParamsScenarioTypeMarketCrash, CorporateRiskStressTestParamsScenarioTypeRegulatoryShock, CorporateRiskStressTestParamsScenarioTypeDepegging:
		return true
	}
	return false
}
