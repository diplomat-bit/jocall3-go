// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// CorporateTreasuryService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasuryService] method instead.
type CorporateTreasuryService struct {
	Options  []option.RequestOption
	Sweeping *CorporateTreasurySweepingService
	Pooling  *CorporateTreasuryPoolingService
}

// NewCorporateTreasuryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateTreasuryService(opts ...option.RequestOption) (r *CorporateTreasuryService) {
	r = &CorporateTreasuryService{}
	r.Options = opts
	r.Sweeping = NewCorporateTreasurySweepingService(opts...)
	r.Pooling = NewCorporateTreasuryPoolingService(opts...)
	return
}

// Corporate Cash Flow Projection
func (r *CorporateTreasuryService) ForecastCashFlow(ctx context.Context, query CorporateTreasuryForecastCashFlowParams, opts ...option.RequestOption) (res *CorporateTreasuryForecastCashFlowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/cash-flow/forecast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// AI Liquidity Optimization Engine
func (r *CorporateTreasuryService) ManageLiquidity(ctx context.Context, body CorporateTreasuryManageLiquidityParams, opts ...option.RequestOption) (res *CorporateTreasuryManageLiquidityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/liquidity/optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateTreasuryForecastCashFlowResponse struct {
	AIRecommendations []string                                      `json:"aiRecommendations"`
	ForecastID        string                                        `json:"forecastId"`
	ProjectedRunway   int64                                         `json:"projectedRunway"`
	JSON              corporateTreasuryForecastCashFlowResponseJSON `json:"-"`
}

// corporateTreasuryForecastCashFlowResponseJSON contains the JSON metadata for the
// struct [CorporateTreasuryForecastCashFlowResponse]
type corporateTreasuryForecastCashFlowResponseJSON struct {
	AIRecommendations apijson.Field
	ForecastID        apijson.Field
	ProjectedRunway   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CorporateTreasuryForecastCashFlowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryForecastCashFlowResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryManageLiquidityResponse struct {
	ProjectedYield float64                                      `json:"projectedYield"`
	StrategyID     string                                       `json:"strategyId"`
	JSON           corporateTreasuryManageLiquidityResponseJSON `json:"-"`
}

// corporateTreasuryManageLiquidityResponseJSON contains the JSON metadata for the
// struct [CorporateTreasuryManageLiquidityResponse]
type corporateTreasuryManageLiquidityResponseJSON struct {
	ProjectedYield apijson.Field
	StrategyID     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateTreasuryManageLiquidityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryManageLiquidityResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryForecastCashFlowParams struct {
	HorizonDays param.Field[int64] `query:"horizonDays"`
}

// URLQuery serializes [CorporateTreasuryForecastCashFlowParams]'s query parameters
// as `url.Values`.
func (r CorporateTreasuryForecastCashFlowParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateTreasuryManageLiquidityParams struct {
	SweepExcess   param.Field[bool]    `json:"sweepExcess"`
	TargetReserve param.Field[float64] `json:"targetReserve"`
}

func (r CorporateTreasuryManageLiquidityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
