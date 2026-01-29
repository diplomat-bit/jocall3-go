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

// AIIncubatorService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorService] method instead.
type AIIncubatorService struct {
	Options  []option.RequestOption
	Analysis *AIIncubatorAnalysisService
	Pitch    *AIIncubatorPitchService
}

// NewAIIncubatorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIIncubatorService(opts ...option.RequestOption) (r *AIIncubatorService) {
	r = &AIIncubatorService{}
	r.Options = opts
	r.Analysis = NewAIIncubatorAnalysisService(opts...)
	r.Pitch = NewAIIncubatorPitchService(opts...)
	return
}

// List All User Business Pitches
func (r *AIIncubatorService) ListPitches(ctx context.Context, opts ...option.RequestOption) (res *AIIncubatorListPitchesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submit a High-Potential Business Plan
func (r *AIIncubatorService) SubmitPitch(ctx context.Context, body AIIncubatorSubmitPitchParams, opts ...option.RequestOption) (res *AIIncubatorSubmitPitchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Rapid Idea Validation Engine
func (r *AIIncubatorService) ValidateIdea(ctx context.Context, body AIIncubatorValidateIdeaParams, opts ...option.RequestOption) (res *AIIncubatorValidateIdeaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIIncubatorListPitchesResponse struct {
	Data []interface{}                      `json:"data"`
	JSON aiIncubatorListPitchesResponseJSON `json:"-"`
}

// aiIncubatorListPitchesResponseJSON contains the JSON metadata for the struct
// [AIIncubatorListPitchesResponse]
type aiIncubatorListPitchesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorListPitchesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorListPitchesResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorSubmitPitchResponse struct {
	PitchID string                             `json:"pitchId"`
	Status  string                             `json:"status"`
	JSON    aiIncubatorSubmitPitchResponseJSON `json:"-"`
}

// aiIncubatorSubmitPitchResponseJSON contains the JSON metadata for the struct
// [AIIncubatorSubmitPitchResponse]
type aiIncubatorSubmitPitchResponseJSON struct {
	PitchID     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorSubmitPitchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorSubmitPitchResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorValidateIdeaResponse struct {
	CriticalFlaws    []string                            `json:"criticalFlaws"`
	FeasibilityScore float64                             `json:"feasibilityScore"`
	JSON             aiIncubatorValidateIdeaResponseJSON `json:"-"`
}

// aiIncubatorValidateIdeaResponseJSON contains the JSON metadata for the struct
// [AIIncubatorValidateIdeaResponse]
type aiIncubatorValidateIdeaResponseJSON struct {
	CriticalFlaws    apijson.Field
	FeasibilityScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIIncubatorValidateIdeaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorValidateIdeaResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorSubmitPitchParams struct {
	// Full text of the concept
	BusinessPlan         param.Field[string]        `json:"businessPlan,required"`
	FinancialProjections param.Field[interface{}]   `json:"financialProjections,required"`
	FoundingTeam         param.Field[[]interface{}] `json:"foundingTeam,required"`
	MarketOpportunity    param.Field[string]        `json:"marketOpportunity,required"`
}

func (r AIIncubatorSubmitPitchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorValidateIdeaParams struct {
	Concept param.Field[string] `json:"concept,required"`
}

func (r AIIncubatorValidateIdeaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
