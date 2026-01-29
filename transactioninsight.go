// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// TransactionInsightService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionInsightService] method instead.
type TransactionInsightService struct {
	Options []option.RequestOption
}

// NewTransactionInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionInsightService(opts ...option.RequestOption) (r *TransactionInsightService) {
	r = &TransactionInsightService{}
	r.Options = opts
	return
}

// Get Cash Flow Prediction (Gemini Powered)
func (r *TransactionInsightService) GetFutureFlow(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetFutureFlowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/future-flow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get AISpending Trend Analysis
func (r *TransactionInsightService) GetSpendingTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetSpendingTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionInsightGetFutureFlowResponse struct {
	ForecastDays      int64                                       `json:"forecastDays"`
	ProjectedLowPoint float64                                     `json:"projectedLowPoint"`
	Recommendations   []string                                    `json:"recommendations"`
	JSON              transactionInsightGetFutureFlowResponseJSON `json:"-"`
}

// transactionInsightGetFutureFlowResponseJSON contains the JSON metadata for the
// struct [TransactionInsightGetFutureFlowResponse]
type transactionInsightGetFutureFlowResponseJSON struct {
	ForecastDays      apijson.Field
	ProjectedLowPoint apijson.Field
	Recommendations   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionInsightGetFutureFlowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetFutureFlowResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetSpendingTrendsResponse struct {
	AINarrative       string                                          `json:"aiNarrative"`
	AnomaliesDetected int64                                           `json:"anomaliesDetected"`
	OverallTrend      string                                          `json:"overallTrend"`
	JSON              transactionInsightGetSpendingTrendsResponseJSON `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseJSON contains the JSON metadata for
// the struct [TransactionInsightGetSpendingTrendsResponse]
type transactionInsightGetSpendingTrendsResponseJSON struct {
	AINarrative       apijson.Field
	AnomaliesDetected apijson.Field
	OverallTrend      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseJSON) RawJSON() string {
	return r.raw
}
