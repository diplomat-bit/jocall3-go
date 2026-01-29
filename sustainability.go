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

// SustainabilityService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityService] method instead.
type SustainabilityService struct {
	Options []option.RequestOption
	Offsets *SustainabilityOffsetService
	Impact  *SustainabilityImpactService
}

// NewSustainabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSustainabilityService(opts ...option.RequestOption) (r *SustainabilityService) {
	r = &SustainabilityService{}
	r.Options = opts
	r.Offsets = NewSustainabilityOffsetService(opts...)
	r.Impact = NewSustainabilityImpactService(opts...)
	return
}

// Analysis of ledger data through Gemini to estimate CO2e output.
func (r *SustainabilityService) GetCarbonFootprint(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityGetCarbonFootprintResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-footprint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityGetCarbonFootprintResponse struct {
	Period            string                                              `json:"period,required"`
	Status            SustainabilityGetCarbonFootprintResponseStatus      `json:"status,required"`
	TotalKgCo2e       float64                                             `json:"totalKgCO2e,required"`
	AIRecommendations []string                                            `json:"aiRecommendations"`
	Breakdown         []SustainabilityGetCarbonFootprintResponseBreakdown `json:"breakdown"`
	JSON              sustainabilityGetCarbonFootprintResponseJSON        `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseJSON contains the JSON metadata for the
// struct [SustainabilityGetCarbonFootprintResponse]
type sustainabilityGetCarbonFootprintResponseJSON struct {
	Period            apijson.Field
	Status            apijson.Field
	TotalKgCo2e       apijson.Field
	AIRecommendations apijson.Field
	Breakdown         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseStatus string

const (
	SustainabilityGetCarbonFootprintResponseStatusOptimal    SustainabilityGetCarbonFootprintResponseStatus = "OPTIMAL"
	SustainabilityGetCarbonFootprintResponseStatusHighOutput SustainabilityGetCarbonFootprintResponseStatus = "HIGH_OUTPUT"
	SustainabilityGetCarbonFootprintResponseStatusCritical   SustainabilityGetCarbonFootprintResponseStatus = "CRITICAL"
)

func (r SustainabilityGetCarbonFootprintResponseStatus) IsKnown() bool {
	switch r {
	case SustainabilityGetCarbonFootprintResponseStatusOptimal, SustainabilityGetCarbonFootprintResponseStatusHighOutput, SustainabilityGetCarbonFootprintResponseStatusCritical:
		return true
	}
	return false
}

type SustainabilityGetCarbonFootprintResponseBreakdown struct {
	Category string                                                `json:"category"`
	Value    float64                                               `json:"value"`
	JSON     sustainabilityGetCarbonFootprintResponseBreakdownJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseBreakdownJSON contains the JSON metadata
// for the struct [SustainabilityGetCarbonFootprintResponseBreakdown]
type sustainabilityGetCarbonFootprintResponseBreakdownJSON struct {
	Category    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseBreakdownJSON) RawJSON() string {
	return r.raw
}
