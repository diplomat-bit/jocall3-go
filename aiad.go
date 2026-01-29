// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIAdService contains methods and other services that help with interacting with
// the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdService] method instead.
type AIAdService struct {
	Options []option.RequestOption
}

// NewAIAdService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIAdService(opts ...option.RequestOption) (r *AIAdService) {
	r = &AIAdService{}
	r.Options = opts
	return
}

// List All Generated Ad Assets
func (r *AIAdService) List(ctx context.Context, opts ...option.RequestOption) (res *AIAdListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate High-Conversion Ad Copy
func (r *AIAdService) GenerateCopy(ctx context.Context, body AIAdGenerateCopyParams, opts ...option.RequestOption) (res *AIAdGenerateCopyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a Standard Video Ad with Veo 2.0
func (r *AIAdService) GenerateVideo(ctx context.Context, body AIAdGenerateVideoParams, opts ...option.RequestOption) (res *AIAdGenerateVideoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate/video"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Poll for Video Gen Status
func (r *AIAdService) GetOperation(ctx context.Context, operationID string, opts ...option.RequestOption) (res *AIAdGetOperationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operationId parameter")
		return
	}
	path := fmt.Sprintf("ai/ads/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// AI Campaign Efficiency Optimizer
func (r *AIAdService) OptimizeCampaign(ctx context.Context, body AIAdOptimizeCampaignParams, opts ...option.RequestOption) (res *AIAdOptimizeCampaignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIAdListResponse struct {
	Data []interface{}        `json:"data"`
	JSON aiAdListResponseJSON `json:"-"`
}

// aiAdListResponseJSON contains the JSON metadata for the struct
// [AIAdListResponse]
type aiAdListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdListResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateCopyResponse struct {
	BodyText  string                       `json:"bodyText"`
	Headlines []string                     `json:"headlines"`
	JSON      aiAdGenerateCopyResponseJSON `json:"-"`
}

// aiAdGenerateCopyResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateCopyResponse]
type aiAdGenerateCopyResponseJSON struct {
	BodyText    apijson.Field
	Headlines   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdGenerateCopyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateCopyResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateVideoResponse struct {
	OperationID string                        `json:"operationId"`
	JSON        aiAdGenerateVideoResponseJSON `json:"-"`
}

// aiAdGenerateVideoResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateVideoResponse]
type aiAdGenerateVideoResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdGenerateVideoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateVideoResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGetOperationResponse struct {
	Progress int64                        `json:"progress"`
	Status   string                       `json:"status"`
	VideoUri string                       `json:"videoUri"`
	JSON     aiAdGetOperationResponseJSON `json:"-"`
}

// aiAdGetOperationResponseJSON contains the JSON metadata for the struct
// [AIAdGetOperationResponse]
type aiAdGetOperationResponseJSON struct {
	Progress    apijson.Field
	Status      apijson.Field
	VideoUri    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdGetOperationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGetOperationResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdOptimizeCampaignResponse struct {
	SuggestedChanges []string                         `json:"suggestedChanges"`
	JSON             aiAdOptimizeCampaignResponseJSON `json:"-"`
}

// aiAdOptimizeCampaignResponseJSON contains the JSON metadata for the struct
// [AIAdOptimizeCampaignResponse]
type aiAdOptimizeCampaignResponseJSON struct {
	SuggestedChanges apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIAdOptimizeCampaignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdOptimizeCampaignResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateCopyParams struct {
	ProductDescription param.Field[string] `json:"productDescription,required"`
	TargetAudience     param.Field[string] `json:"targetAudience,required"`
}

func (r AIAdGenerateCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdGenerateVideoParams struct {
	LengthSeconds param.Field[AIAdGenerateVideoParamsLengthSeconds] `json:"lengthSeconds,required"`
	// Visual description
	Prompt param.Field[string]                       `json:"prompt,required"`
	Style  param.Field[AIAdGenerateVideoParamsStyle] `json:"style,required"`
}

func (r AIAdGenerateVideoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdGenerateVideoParamsLengthSeconds int64

const (
	AIAdGenerateVideoParamsLengthSeconds15 AIAdGenerateVideoParamsLengthSeconds = 15
	AIAdGenerateVideoParamsLengthSeconds30 AIAdGenerateVideoParamsLengthSeconds = 30
	AIAdGenerateVideoParamsLengthSeconds60 AIAdGenerateVideoParamsLengthSeconds = 60
)

func (r AIAdGenerateVideoParamsLengthSeconds) IsKnown() bool {
	switch r {
	case AIAdGenerateVideoParamsLengthSeconds15, AIAdGenerateVideoParamsLengthSeconds30, AIAdGenerateVideoParamsLengthSeconds60:
		return true
	}
	return false
}

type AIAdGenerateVideoParamsStyle string

const (
	AIAdGenerateVideoParamsStyleCinematic    AIAdGenerateVideoParamsStyle = "Cinematic"
	AIAdGenerateVideoParamsStyleMinimalist   AIAdGenerateVideoParamsStyle = "Minimalist"
	AIAdGenerateVideoParamsStyleCyberpunk    AIAdGenerateVideoParamsStyle = "Cyberpunk"
	AIAdGenerateVideoParamsStyleProfessional AIAdGenerateVideoParamsStyle = "Professional"
)

func (r AIAdGenerateVideoParamsStyle) IsKnown() bool {
	switch r {
	case AIAdGenerateVideoParamsStyleCinematic, AIAdGenerateVideoParamsStyleMinimalist, AIAdGenerateVideoParamsStyleCyberpunk, AIAdGenerateVideoParamsStyleProfessional:
		return true
	}
	return false
}

type AIAdOptimizeCampaignParams struct {
	CampaignData param.Field[interface{}] `json:"campaignData,required"`
}

func (r AIAdOptimizeCampaignParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
