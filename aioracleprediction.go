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

// AIOraclePredictionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOraclePredictionService] method instead.
type AIOraclePredictionService struct {
	Options []option.RequestOption
}

// NewAIOraclePredictionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOraclePredictionService(opts ...option.RequestOption) (r *AIOraclePredictionService) {
	r = &AIOraclePredictionService{}
	r.Options = opts
	return
}

// Get AI-Driven Inflation Forecast
func (r *AIOraclePredictionService) GetInflationForecast(ctx context.Context, query AIOraclePredictionGetInflationForecastParams, opts ...option.RequestOption) (res *AIOraclePredictionGetInflationForecastResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/predictions/inflation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Market Volatility & Crash Probability
func (r *AIOraclePredictionService) GetMarketCrashProbability(ctx context.Context, opts ...option.RequestOption) (res *AIOraclePredictionGetMarketCrashProbabilityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/predictions/market-crash-probability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIOraclePredictionGetInflationForecastResponse struct {
	ConfidenceScore int64                                              `json:"confidenceScore"`
	ForecastedCpi   float64                                            `json:"forecastedCPI"`
	Period          string                                             `json:"period"`
	JSON            aiOraclePredictionGetInflationForecastResponseJSON `json:"-"`
}

// aiOraclePredictionGetInflationForecastResponseJSON contains the JSON metadata
// for the struct [AIOraclePredictionGetInflationForecastResponse]
type aiOraclePredictionGetInflationForecastResponseJSON struct {
	ConfidenceScore apijson.Field
	ForecastedCpi   apijson.Field
	Period          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AIOraclePredictionGetInflationForecastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOraclePredictionGetInflationForecastResponseJSON) RawJSON() string {
	return r.raw
}

type AIOraclePredictionGetMarketCrashProbabilityResponse struct {
	AINarrative      string                                                  `json:"aiNarrative"`
	CrashProbability float64                                                 `json:"crashProbability"`
	RiskFactors      []string                                                `json:"riskFactors"`
	JSON             aiOraclePredictionGetMarketCrashProbabilityResponseJSON `json:"-"`
}

// aiOraclePredictionGetMarketCrashProbabilityResponseJSON contains the JSON
// metadata for the struct [AIOraclePredictionGetMarketCrashProbabilityResponse]
type aiOraclePredictionGetMarketCrashProbabilityResponseJSON struct {
	AINarrative      apijson.Field
	CrashProbability apijson.Field
	RiskFactors      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOraclePredictionGetMarketCrashProbabilityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOraclePredictionGetMarketCrashProbabilityResponseJSON) RawJSON() string {
	return r.raw
}

type AIOraclePredictionGetInflationForecastParams struct {
	Region param.Field[string] `query:"region"`
}

// URLQuery serializes [AIOraclePredictionGetInflationForecastParams]'s query
// parameters as `url.Values`.
func (r AIOraclePredictionGetInflationForecastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
