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

// InvestmentPerformanceService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentPerformanceService] method instead.
type InvestmentPerformanceService struct {
	Options []option.RequestOption
}

// NewInvestmentPerformanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestmentPerformanceService(opts ...option.RequestOption) (r *InvestmentPerformanceService) {
	r = &InvestmentPerformanceService{}
	r.Options = opts
	return
}

// Get Historical Performance Curves
func (r *InvestmentPerformanceService) GetHistorical(ctx context.Context, query InvestmentPerformanceGetHistoricalParams, opts ...option.RequestOption) (res *InvestmentPerformanceGetHistoricalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/performance/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InvestmentPerformanceGetHistoricalResponse struct {
	BenchmarkComparison float64                                        `json:"benchmarkComparison"`
	Points              []interface{}                                  `json:"points"`
	JSON                investmentPerformanceGetHistoricalResponseJSON `json:"-"`
}

// investmentPerformanceGetHistoricalResponseJSON contains the JSON metadata for
// the struct [InvestmentPerformanceGetHistoricalResponse]
type investmentPerformanceGetHistoricalResponseJSON struct {
	BenchmarkComparison apijson.Field
	Points              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestmentPerformanceGetHistoricalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPerformanceGetHistoricalResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPerformanceGetHistoricalParams struct {
	Range param.Field[InvestmentPerformanceGetHistoricalParamsRange] `query:"range"`
}

// URLQuery serializes [InvestmentPerformanceGetHistoricalParams]'s query
// parameters as `url.Values`.
func (r InvestmentPerformanceGetHistoricalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPerformanceGetHistoricalParamsRange string

const (
	InvestmentPerformanceGetHistoricalParamsRange1m  InvestmentPerformanceGetHistoricalParamsRange = "1m"
	InvestmentPerformanceGetHistoricalParamsRange3m  InvestmentPerformanceGetHistoricalParamsRange = "3m"
	InvestmentPerformanceGetHistoricalParamsRange1y  InvestmentPerformanceGetHistoricalParamsRange = "1y"
	InvestmentPerformanceGetHistoricalParamsRange5y  InvestmentPerformanceGetHistoricalParamsRange = "5y"
	InvestmentPerformanceGetHistoricalParamsRangeMax InvestmentPerformanceGetHistoricalParamsRange = "max"
)

func (r InvestmentPerformanceGetHistoricalParamsRange) IsKnown() bool {
	switch r {
	case InvestmentPerformanceGetHistoricalParamsRange1m, InvestmentPerformanceGetHistoricalParamsRange3m, InvestmentPerformanceGetHistoricalParamsRange1y, InvestmentPerformanceGetHistoricalParamsRange5y, InvestmentPerformanceGetHistoricalParamsRangeMax:
		return true
	}
	return false
}
