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
func (r *TransactionInsightService) GetSpendingTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetSpendingTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionInsightGetSpendingTrendsResponse struct {
	AIInsights            []TransactionInsightGetSpendingTrendsResponseAIInsight             `json:"aiInsights,required"`
	ForecastNextMonth     float64                                                            `json:"forecastNextMonth,required"`
	OverallTrend          string                                                             `json:"overallTrend,required"`
	PercentageChange      float64                                                            `json:"percentageChange,required"`
	Period                string                                                             `json:"period,required"`
	TopCategoriesByChange []TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange `json:"topCategoriesByChange,required"`
	JSON                  transactionInsightGetSpendingTrendsResponseJSON                    `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseJSON contains the JSON metadata for
// the struct [TransactionInsightGetSpendingTrendsResponse]
type transactionInsightGetSpendingTrendsResponseJSON struct {
	AIInsights            apijson.Field
	ForecastNextMonth     apijson.Field
	OverallTrend          apijson.Field
	PercentageChange      apijson.Field
	Period                apijson.Field
	TopCategoriesByChange apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetSpendingTrendsResponseAIInsight struct {
	ID                       string                                                   `json:"id"`
	ActionableRecommendation string                                                   `json:"actionableRecommendation"`
	Category                 string                                                   `json:"category"`
	Description              string                                                   `json:"description"`
	Severity                 string                                                   `json:"severity"`
	Timestamp                time.Time                                                `json:"timestamp" format:"date-time"`
	Title                    string                                                   `json:"title"`
	JSON                     transactionInsightGetSpendingTrendsResponseAIInsightJSON `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseAIInsightJSON contains the JSON
// metadata for the struct [TransactionInsightGetSpendingTrendsResponseAIInsight]
type transactionInsightGetSpendingTrendsResponseAIInsightJSON struct {
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

func (r *TransactionInsightGetSpendingTrendsResponseAIInsight) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseAIInsightJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange struct {
	AbsoluteChange   float64                                                              `json:"absoluteChange"`
	Category         string                                                               `json:"category"`
	PercentageChange float64                                                              `json:"percentageChange"`
	JSON             transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON contains
// the JSON metadata for the struct
// [TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange]
type transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON struct {
	AbsoluteChange   apijson.Field
	Category         apijson.Field
	PercentageChange apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON) RawJSON() string {
	return r.raw
}
