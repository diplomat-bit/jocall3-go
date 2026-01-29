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

// SustainabilityImpactService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityImpactService] method instead.
type SustainabilityImpactService struct {
	Options []option.RequestOption
}

// NewSustainabilityImpactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSustainabilityImpactService(opts ...option.RequestOption) (r *SustainabilityImpactService) {
	r = &SustainabilityImpactService{}
	r.Options = opts
	return
}

// Search Global Green Projects
func (r *SustainabilityImpactService) ListGreenProjects(ctx context.Context, query SustainabilityImpactListGreenProjectsParams, opts ...option.RequestOption) (res *SustainabilityImpactListGreenProjectsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/impact/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// ESG Portfolio Impact Analysis
func (r *SustainabilityImpactService) GetPortfolioAnalysis(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityImpactGetPortfolioAnalysisResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/impact/portfolio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityImpactListGreenProjectsResponse struct {
	Data []interface{}                                     `json:"data"`
	JSON sustainabilityImpactListGreenProjectsResponseJSON `json:"-"`
}

// sustainabilityImpactListGreenProjectsResponseJSON contains the JSON metadata for
// the struct [SustainabilityImpactListGreenProjectsResponse]
type sustainabilityImpactListGreenProjectsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SustainabilityImpactListGreenProjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityImpactListGreenProjectsResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityImpactGetPortfolioAnalysisResponse struct {
	EsgScore                int64                                                `json:"esgScore"`
	FossilFuelExposure      float64                                              `json:"fossilFuelExposure"`
	GreenProjectInvolvement []string                                             `json:"greenProjectInvolvement"`
	SocialJusticeRating     string                                               `json:"socialJusticeRating"`
	JSON                    sustainabilityImpactGetPortfolioAnalysisResponseJSON `json:"-"`
}

// sustainabilityImpactGetPortfolioAnalysisResponseJSON contains the JSON metadata
// for the struct [SustainabilityImpactGetPortfolioAnalysisResponse]
type sustainabilityImpactGetPortfolioAnalysisResponseJSON struct {
	EsgScore                apijson.Field
	FossilFuelExposure      apijson.Field
	GreenProjectInvolvement apijson.Field
	SocialJusticeRating     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SustainabilityImpactGetPortfolioAnalysisResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityImpactGetPortfolioAnalysisResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityImpactListGreenProjectsParams struct {
	Continent param.Field[string] `query:"continent"`
}

// URLQuery serializes [SustainabilityImpactListGreenProjectsParams]'s query
// parameters as `url.Values`.
func (r SustainabilityImpactListGreenProjectsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
