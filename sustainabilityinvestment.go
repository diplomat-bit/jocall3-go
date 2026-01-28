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

// SustainabilityInvestmentService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityInvestmentService] method instead.
type SustainabilityInvestmentService struct {
	Options []option.RequestOption
}

// NewSustainabilityInvestmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSustainabilityInvestmentService(opts ...option.RequestOption) (r *SustainabilityInvestmentService) {
	r = &SustainabilityInvestmentService{}
	r.Options = opts
	return
}

// Provides an AI-driven analysis of the Environmental, Social, and Governance
// (ESG) impact of the user's entire investment portfolio, benchmarking against
// industry standards and suggesting more sustainable alternatives.
func (r *SustainabilityInvestmentService) AnalyzeImpact(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityInvestmentAnalyzeImpactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/investments/impact"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityInvestmentAnalyzeImpactResponse struct {
	// Breakdown of the portfolio's ESG score by individual factors.
	BreakdownByEsgFactors interface{}                                       `json:"breakdownByESGFactors,required"`
	JSON                  sustainabilityInvestmentAnalyzeImpactResponseJSON `json:"-"`
}

// sustainabilityInvestmentAnalyzeImpactResponseJSON contains the JSON metadata for
// the struct [SustainabilityInvestmentAnalyzeImpactResponse]
type sustainabilityInvestmentAnalyzeImpactResponseJSON struct {
	BreakdownByEsgFactors apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SustainabilityInvestmentAnalyzeImpactResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityInvestmentAnalyzeImpactResponseJSON) RawJSON() string {
	return r.raw
}
