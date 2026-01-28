// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"
	"time"

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

// Retrieves AI-generated insights into user spending trends over time, identifying
// patterns and anomalies.
func (r *TransactionInsightService) GetTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionInsightGetTrendsResponse struct {
	AIInsights            []TransactionInsightGetTrendsResponseAIInsight             `json:"aiInsights,required"`
	ForecastNextMonth     float64                                                    `json:"forecastNextMonth,required"`
	OverallTrend          string                                                     `json:"overallTrend,required"`
	PercentageChange      float64                                                    `json:"percentageChange,required"`
	Period                string                                                     `json:"period,required"`
	TopCategoriesByChange []TransactionInsightGetTrendsResponseTopCategoriesByChange `json:"topCategoriesByChange,required"`
	JSON                  transactionInsightGetTrendsResponseJSON                    `json:"-"`
}

// transactionInsightGetTrendsResponseJSON contains the JSON metadata for the
// struct [TransactionInsightGetTrendsResponse]
type transactionInsightGetTrendsResponseJSON struct {
	AIInsights            apijson.Field
	ForecastNextMonth     apijson.Field
	OverallTrend          apijson.Field
	PercentageChange      apijson.Field
	Period                apijson.Field
	TopCategoriesByChange apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionInsightGetTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetTrendsResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetTrendsResponseAIInsight struct {
	ID                       string                                           `json:"id"`
	ActionableRecommendation string                                           `json:"actionableRecommendation"`
	Category                 string                                           `json:"category"`
	Description              string                                           `json:"description"`
	Severity                 string                                           `json:"severity"`
	Timestamp                time.Time                                        `json:"timestamp" format:"date-time"`
	Title                    string                                           `json:"title"`
	JSON                     transactionInsightGetTrendsResponseAIInsightJSON `json:"-"`
}

// transactionInsightGetTrendsResponseAIInsightJSON contains the JSON metadata for
// the struct [TransactionInsightGetTrendsResponseAIInsight]
type transactionInsightGetTrendsResponseAIInsightJSON struct {
	ID                       apijson.Field
	ActionableRecommendation apijson.Field
	Category                 apijson.Field
	Description              apijson.Field
	Severity                 apijson.Field
	Timestamp                apijson.Field
	Title                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *TransactionInsightGetTrendsResponseAIInsight) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetTrendsResponseAIInsightJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetTrendsResponseTopCategoriesByChange struct {
	AbsoluteChange   float64                                                      `json:"absoluteChange"`
	Category         string                                                       `json:"category"`
	PercentageChange float64                                                      `json:"percentageChange"`
	JSON             transactionInsightGetTrendsResponseTopCategoriesByChangeJSON `json:"-"`
}

// transactionInsightGetTrendsResponseTopCategoriesByChangeJSON contains the JSON
// metadata for the struct
// [TransactionInsightGetTrendsResponseTopCategoriesByChange]
type transactionInsightGetTrendsResponseTopCategoriesByChangeJSON struct {
	AbsoluteChange   apijson.Field
	Category         apijson.Field
	PercentageChange apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionInsightGetTrendsResponseTopCategoriesByChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetTrendsResponseTopCategoriesByChangeJSON) RawJSON() string {
	return r.raw
}
