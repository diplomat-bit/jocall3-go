// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIIncubatorPitchService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorPitchService] method instead.
type AIIncubatorPitchService struct {
	Options []option.RequestOption
}

// NewAIIncubatorPitchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIncubatorPitchService(opts ...option.RequestOption) (r *AIIncubatorPitchService) {
	r = &AIIncubatorPitchService{}
	r.Options = opts
	return
}

// Retrieves the granular AI-driven analysis, strategic feedback, market validation
// results, and any outstanding questions from Quantum Weaver for a specific
// business pitch.
func (r *AIIncubatorPitchService) GetDetails(ctx context.Context, pitchID string, opts ...option.RequestOption) (res *AIIncubatorPitchGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/details", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allows the entrepreneur to respond to specific questions or provide additional
// details requested by Quantum Weaver, moving the pitch forward in the incubation
// process.
func (r *AIIncubatorPitchService) SubmitFeedback(ctx context.Context, pitchID string, body AIIncubatorPitchSubmitFeedbackParams, opts ...option.RequestOption) (res *AIIncubatorPitchSubmitFeedbackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/feedback", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AIIncubatorPitchGetDetailsResponse struct {
	// AI-generated coaching plan for the entrepreneur.
	AICoachingPlan interface{} `json:"aiCoachingPlan"`
	// AI's detailed financial model analysis.
	AIFinancialModel AIIncubatorPitchGetDetailsResponseAIFinancialModel `json:"aiFinancialModel"`
	// AI's detailed market analysis.
	AIMarketAnalysis interface{} `json:"aiMarketAnalysis"`
	// AI's assessment of risks associated with the venture.
	AIRiskAssessment interface{}                            `json:"aiRiskAssessment"`
	JSON             aiIncubatorPitchGetDetailsResponseJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseJSON contains the JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponse]
type aiIncubatorPitchGetDetailsResponseJSON struct {
	AICoachingPlan   apijson.Field
	AIFinancialModel apijson.Field
	AIMarketAnalysis apijson.Field
	AIRiskAssessment apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// AI's detailed financial model analysis.
type AIIncubatorPitchGetDetailsResponseAIFinancialModel struct {
	CostStructureAnalysis interface{}                                            `json:"costStructureAnalysis"`
	RevenueBreakdown      interface{}                                            `json:"revenueBreakdown"`
	JSON                  aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIFinancialModel]
type aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON struct {
	CostStructureAnalysis apijson.Field
	RevenueBreakdown      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIFinancialModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchSubmitFeedbackResponse = interface{}

type AIIncubatorPitchSubmitFeedbackParams struct {
}

func (r AIIncubatorPitchSubmitFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
