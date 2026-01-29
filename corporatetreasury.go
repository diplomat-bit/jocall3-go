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

// Execute bulk payouts
func (r *CorporateTreasuryService) ExecuteBulkPayouts(ctx context.Context, body CorporateTreasuryExecuteBulkPayoutsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/treasury/bulk-payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// AI Liquidity Optimization Engine
func (r *CorporateTreasuryService) OptimizeLiquidity(ctx context.Context, body CorporateTreasuryOptimizeLiquidityParams, opts ...option.RequestOption) (res *CorporateTreasuryOptimizeLiquidityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/liquidity/optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Corporate Cash Flow Projection
func (r *CorporateTreasuryService) GetCashFlowForecast(ctx context.Context, query CorporateTreasuryGetCashFlowForecastParams, opts ...option.RequestOption) (res *CorporateTreasuryGetCashFlowForecastResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/cash-flow/forecast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get current liquidity positions
func (r *CorporateTreasuryService) GetLiquidityPositions(ctx context.Context, opts ...option.RequestOption) (res *CorporateTreasuryGetLiquidityPositionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/treasury/liquidity-positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CorporateTreasuryOptimizeLiquidityResponse struct {
	ProjectedYield float64                                        `json:"projectedYield"`
	StrategyID     string                                         `json:"strategyId"`
	JSON           corporateTreasuryOptimizeLiquidityResponseJSON `json:"-"`
}

// corporateTreasuryOptimizeLiquidityResponseJSON contains the JSON metadata for
// the struct [CorporateTreasuryOptimizeLiquidityResponse]
type corporateTreasuryOptimizeLiquidityResponseJSON struct {
	ProjectedYield apijson.Field
	StrategyID     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateTreasuryOptimizeLiquidityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryOptimizeLiquidityResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryGetCashFlowForecastResponse struct {
	AIRecommendations []string                                         `json:"aiRecommendations"`
	ForecastID        string                                           `json:"forecastId"`
	ProjectedRunway   int64                                            `json:"projectedRunway"`
	JSON              corporateTreasuryGetCashFlowForecastResponseJSON `json:"-"`
}

// corporateTreasuryGetCashFlowForecastResponseJSON contains the JSON metadata for
// the struct [CorporateTreasuryGetCashFlowForecastResponse]
type corporateTreasuryGetCashFlowForecastResponseJSON struct {
	AIRecommendations apijson.Field
	ForecastID        apijson.Field
	ProjectedRunway   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CorporateTreasuryGetCashFlowForecastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetCashFlowForecastResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryGetLiquidityPositionsResponse struct {
	Positions      []interface{}                                      `json:"positions"`
	TotalLiquidity float64                                            `json:"total_liquidity"`
	JSON           corporateTreasuryGetLiquidityPositionsResponseJSON `json:"-"`
}

// corporateTreasuryGetLiquidityPositionsResponseJSON contains the JSON metadata
// for the struct [CorporateTreasuryGetLiquidityPositionsResponse]
type corporateTreasuryGetLiquidityPositionsResponseJSON struct {
	Positions      apijson.Field
	TotalLiquidity apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateTreasuryGetLiquidityPositionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateTreasuryGetLiquidityPositionsResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateTreasuryExecuteBulkPayoutsParams struct {
	Payouts param.Field[[]CorporateTreasuryExecuteBulkPayoutsParamsPayout] `json:"payouts,required"`
}

func (r CorporateTreasuryExecuteBulkPayoutsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateTreasuryExecuteBulkPayoutsParamsPayout struct {
	Amount      param.Field[float64] `json:"amount"`
	RecipientID param.Field[string]  `json:"recipient_id"`
}

func (r CorporateTreasuryExecuteBulkPayoutsParamsPayout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateTreasuryOptimizeLiquidityParams struct {
	SweepExcess   param.Field[bool]    `json:"sweepExcess"`
	TargetReserve param.Field[float64] `json:"targetReserve"`
}

func (r CorporateTreasuryOptimizeLiquidityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateTreasuryGetCashFlowForecastParams struct {
	HorizonDays param.Field[int64] `query:"horizonDays"`
}

// URLQuery serializes [CorporateTreasuryGetCashFlowForecastParams]'s query
// parameters as `url.Values`.
func (r CorporateTreasuryGetCashFlowForecastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
