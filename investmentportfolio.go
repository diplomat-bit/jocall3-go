// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// InvestmentPortfolioService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentPortfolioService] method instead.
type InvestmentPortfolioService struct {
	Options []option.RequestOption
}

// NewInvestmentPortfolioService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestmentPortfolioService(opts ...option.RequestOption) (r *InvestmentPortfolioService) {
	r = &InvestmentPortfolioService{}
	r.Options = opts
	return
}

// Create Strategic Portfolio
func (r *InvestmentPortfolioService) New(ctx context.Context, body InvestmentPortfolioNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get Full Portfolio Performance
func (r *InvestmentPortfolioService) Get(ctx context.Context, portfolioID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Update Portfolio Strategy
func (r *InvestmentPortfolioService) Update(ctx context.Context, portfolioID string, body InvestmentPortfolioUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// List All Investment Portfolios
func (r *InvestmentPortfolioService) List(ctx context.Context, query InvestmentPortfolioListParams, opts ...option.RequestOption) (res *InvestmentPortfolioListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Trigger Gemini AI Rebalancing
func (r *InvestmentPortfolioService) Rebalance(ctx context.Context, portfolioID string, body InvestmentPortfolioRebalanceParams, opts ...option.RequestOption) (res *InvestmentPortfolioRebalanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s/rebalance", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvestmentPortfolioListResponse struct {
	Data []InvestmentPortfolioListResponseData `json:"data"`
	JSON investmentPortfolioListResponseJSON   `json:"-"`
}

// investmentPortfolioListResponseJSON contains the JSON metadata for the struct
// [InvestmentPortfolioListResponse]
type investmentPortfolioListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentPortfolioListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioListResponseData struct {
	ID         string                                  `json:"id"`
	Name       string                                  `json:"name"`
	TotalValue float64                                 `json:"totalValue"`
	JSON       investmentPortfolioListResponseDataJSON `json:"-"`
}

// investmentPortfolioListResponseDataJSON contains the JSON metadata for the
// struct [InvestmentPortfolioListResponseData]
type investmentPortfolioListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	TotalValue  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentPortfolioListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioListResponseDataJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRebalanceResponse struct {
	ImpactSummary string                                   `json:"impactSummary"`
	RebalanceID   string                                   `json:"rebalanceId"`
	JSON          investmentPortfolioRebalanceResponseJSON `json:"-"`
}

// investmentPortfolioRebalanceResponseJSON contains the JSON metadata for the
// struct [InvestmentPortfolioRebalanceResponse]
type investmentPortfolioRebalanceResponseJSON struct {
	ImpactSummary apijson.Field
	RebalanceID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvestmentPortfolioRebalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioRebalanceResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioNewParams struct {
	Name              param.Field[string]                               `json:"name,required"`
	Strategy          param.Field[InvestmentPortfolioNewParamsStrategy] `json:"strategy,required"`
	InitialAllocation param.Field[interface{}]                          `json:"initialAllocation"`
}

func (r InvestmentPortfolioNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioNewParamsStrategy string

const (
	InvestmentPortfolioNewParamsStrategyGrowth     InvestmentPortfolioNewParamsStrategy = "GROWTH"
	InvestmentPortfolioNewParamsStrategyBalanced   InvestmentPortfolioNewParamsStrategy = "BALANCED"
	InvestmentPortfolioNewParamsStrategyIncome     InvestmentPortfolioNewParamsStrategy = "INCOME"
	InvestmentPortfolioNewParamsStrategyEsgFocused InvestmentPortfolioNewParamsStrategy = "ESG_FOCUSED"
)

func (r InvestmentPortfolioNewParamsStrategy) IsKnown() bool {
	switch r {
	case InvestmentPortfolioNewParamsStrategyGrowth, InvestmentPortfolioNewParamsStrategyBalanced, InvestmentPortfolioNewParamsStrategyIncome, InvestmentPortfolioNewParamsStrategyEsgFocused:
		return true
	}
	return false
}

type InvestmentPortfolioUpdateParams struct {
	RiskTolerance param.Field[int64]  `json:"riskTolerance"`
	Strategy      param.Field[string] `json:"strategy"`
}

func (r InvestmentPortfolioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioListParams struct {
	Limit  param.Field[int64] `query:"limit"`
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [InvestmentPortfolioListParams]'s query parameters as
// `url.Values`.
func (r InvestmentPortfolioListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPortfolioRebalanceParams struct {
	ExecutionMode param.Field[InvestmentPortfolioRebalanceParamsExecutionMode] `json:"executionMode"`
}

func (r InvestmentPortfolioRebalanceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioRebalanceParamsExecutionMode string

const (
	InvestmentPortfolioRebalanceParamsExecutionModeAuto        InvestmentPortfolioRebalanceParamsExecutionMode = "AUTO"
	InvestmentPortfolioRebalanceParamsExecutionModeConfirmOnly InvestmentPortfolioRebalanceParamsExecutionMode = "CONFIRM_ONLY"
)

func (r InvestmentPortfolioRebalanceParamsExecutionMode) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceParamsExecutionModeAuto, InvestmentPortfolioRebalanceParamsExecutionModeConfirmOnly:
		return true
	}
	return false
}
