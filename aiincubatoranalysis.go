// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIIncubatorAnalysisService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorAnalysisService] method instead.
type AIIncubatorAnalysisService struct {
	Options []option.RequestOption
}

// NewAIIncubatorAnalysisService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIncubatorAnalysisService(opts ...option.RequestOption) (r *AIIncubatorAnalysisService) {
	r = &AIIncubatorAnalysisService{}
	r.Options = opts
	return
}

// Generate Automated Competitor Landscape
func (r *AIIncubatorAnalysisService) CompetitorScan(ctx context.Context, body AIIncubatorAnalysisCompetitorScanParams, opts ...option.RequestOption) (res *AIIncubatorAnalysisCompetitorScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/analysis/competitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate AI SWOT Analysis
func (r *AIIncubatorAnalysisService) Swot(ctx context.Context, body AIIncubatorAnalysisSwotParams, opts ...option.RequestOption) (res *AIIncubatorAnalysisSwotResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/analysis/swot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIIncubatorAnalysisCompetitorScanResponse struct {
	Competitors         []interface{}                                 `json:"competitors"`
	MarketShareAnalysis string                                        `json:"marketShareAnalysis"`
	JSON                aiIncubatorAnalysisCompetitorScanResponseJSON `json:"-"`
}

// aiIncubatorAnalysisCompetitorScanResponseJSON contains the JSON metadata for the
// struct [AIIncubatorAnalysisCompetitorScanResponse]
type aiIncubatorAnalysisCompetitorScanResponseJSON struct {
	Competitors         apijson.Field
	MarketShareAnalysis apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AIIncubatorAnalysisCompetitorScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorAnalysisCompetitorScanResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorAnalysisSwotResponse struct {
	Strengths  []string                            `json:"strengths"`
	Weaknesses []string                            `json:"weaknesses"`
	JSON       aiIncubatorAnalysisSwotResponseJSON `json:"-"`
}

// aiIncubatorAnalysisSwotResponseJSON contains the JSON metadata for the struct
// [AIIncubatorAnalysisSwotResponse]
type aiIncubatorAnalysisSwotResponseJSON struct {
	Strengths   apijson.Field
	Weaknesses  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorAnalysisSwotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorAnalysisSwotResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorAnalysisCompetitorScanParams struct {
	Industry param.Field[string] `json:"industry,required"`
	Niche    param.Field[string] `json:"niche,required"`
}

func (r AIIncubatorAnalysisCompetitorScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorAnalysisSwotParams struct {
	BusinessContext param.Field[string] `json:"businessContext,required"`
}

func (r AIIncubatorAnalysisSwotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
