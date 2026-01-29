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

// Retrieves detailed information for a specific investment portfolio, including
// holdings, performance, and AI insights.
func (r *InvestmentPortfolioService) Get(ctx context.Context, portfolioID string, opts ...option.RequestOption) (res *InvestmentPortfolioGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates high-level details of an investment portfolio, such as name or risk
// tolerance.
func (r *InvestmentPortfolioService) Update(ctx context.Context, portfolioID string, body InvestmentPortfolioUpdateParams, opts ...option.RequestOption) (res *InvestmentPortfolioUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a summary of all investment portfolios linked to the user's account.
func (r *InvestmentPortfolioService) List(ctx context.Context, query InvestmentPortfolioListParams, opts ...option.RequestOption) (res *InvestmentPortfolioListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Triggers an AI-driven rebalancing process for a specific investment portfolio
// based on a target risk tolerance or strategy.
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

type InvestmentPortfolioGetResponse = interface{}

type InvestmentPortfolioUpdateResponse = interface{}

type InvestmentPortfolioListResponse = interface{}

type InvestmentPortfolioRebalanceResponse = interface{}

type InvestmentPortfolioUpdateParams struct {
}

func (r InvestmentPortfolioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
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
}

func (r InvestmentPortfolioRebalanceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
