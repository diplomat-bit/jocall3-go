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
func (r *AIOraclePredictionService) Inflation(ctx context.Context, query AIOraclePredictionInflationParams, opts ...option.RequestOption) (res *AIOraclePredictionInflationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/predictions/inflation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Market Volatility & Crash Probability
func (r *AIOraclePredictionService) MarketCrash(ctx context.Context, opts ...option.RequestOption) (res *AIOraclePredictionMarketCrashResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/predictions/market-crash-probability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIOraclePredictionInflationResponse struct {
	ConfidenceScore int64                                   `json:"confidenceScore"`
	ForecastedCpi   float64                                 `json:"forecastedCPI"`
	Period          string                                  `json:"period"`
	JSON            aiOraclePredictionInflationResponseJSON `json:"-"`
}

// aiOraclePredictionInflationResponseJSON contains the JSON metadata for the
// struct [AIOraclePredictionInflationResponse]
type aiOraclePredictionInflationResponseJSON struct {
	ConfidenceScore apijson.Field
	ForecastedCpi   apijson.Field
	Period          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AIOraclePredictionInflationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOraclePredictionInflationResponseJSON) RawJSON() string {
	return r.raw
}

type AIOraclePredictionMarketCrashResponse struct {
	AINarrative      string                                    `json:"aiNarrative"`
	CrashProbability float64                                   `json:"crashProbability"`
	RiskFactors      []string                                  `json:"riskFactors"`
	JSON             aiOraclePredictionMarketCrashResponseJSON `json:"-"`
}

// aiOraclePredictionMarketCrashResponseJSON contains the JSON metadata for the
// struct [AIOraclePredictionMarketCrashResponse]
type aiOraclePredictionMarketCrashResponseJSON struct {
	AINarrative      apijson.Field
	CrashProbability apijson.Field
	RiskFactors      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOraclePredictionMarketCrashResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOraclePredictionMarketCrashResponseJSON) RawJSON() string {
	return r.raw
}

type AIOraclePredictionInflationParams struct {
	Region param.Field[string] `query:"region"`
}

// URLQuery serializes [AIOraclePredictionInflationParams]'s query parameters as
// `url.Values`.
func (r AIOraclePredictionInflationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
