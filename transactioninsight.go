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
func (r *TransactionInsightService) GetForecast(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetForecastResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/future-flow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get AISpending Trend Analysis
func (r *TransactionInsightService) GetTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionInsightGetForecastResponse struct {
	ForecastDays      int64                                     `json:"forecastDays"`
	ProjectedLowPoint float64                                   `json:"projectedLowPoint"`
	Recommendations   []string                                  `json:"recommendations"`
	JSON              transactionInsightGetForecastResponseJSON `json:"-"`
}

// transactionInsightGetForecastResponseJSON contains the JSON metadata for the
// struct [TransactionInsightGetForecastResponse]
type transactionInsightGetForecastResponseJSON struct {
	ForecastDays      apijson.Field
	ProjectedLowPoint apijson.Field
	Recommendations   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionInsightGetForecastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetForecastResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetTrendsResponse struct {
	AINarrative       string                                  `json:"aiNarrative"`
	AnomaliesDetected int64                                   `json:"anomaliesDetected"`
	OverallTrend      string                                  `json:"overallTrend"`
	JSON              transactionInsightGetTrendsResponseJSON `json:"-"`
}

// transactionInsightGetTrendsResponseJSON contains the JSON metadata for the
// struct [TransactionInsightGetTrendsResponse]
type transactionInsightGetTrendsResponseJSON struct {
	AINarrative       apijson.Field
	AnomaliesDetected apijson.Field
	OverallTrend      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionInsightGetTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetTrendsResponseJSON) RawJSON() string {
	return r.raw
}
