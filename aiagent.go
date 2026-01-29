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

// AIAgentService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAgentService] method instead.
type AIAgentService struct {
	Options []option.RequestOption
}

// NewAIAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIAgentService(opts ...option.RequestOption) (r *AIAgentService) {
	r = &AIAgentService{}
	r.Options = opts
	return
}

// List Quantum Agent Capabilities
func (r *AIAgentService) GetCapabilities(ctx context.Context, opts ...option.RequestOption) (res *AIAgentGetCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/agent/capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve Current System System Prompts
func (r *AIAgentService) GetPrompts(ctx context.Context, opts ...option.RequestOption) (res *AIAgentGetPromptsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/agent/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update System Instructions for Gemini Engine
func (r *AIAgentService) UpdatePrompts(ctx context.Context, body AIAgentUpdatePromptsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "ai/agent/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type AIAgentGetCapabilitiesResponse struct {
	Data []AIAgentGetCapabilitiesResponseData `json:"data"`
	JSON aiAgentGetCapabilitiesResponseJSON   `json:"-"`
}

// aiAgentGetCapabilitiesResponseJSON contains the JSON metadata for the struct
// [AIAgentGetCapabilitiesResponse]
type aiAgentGetCapabilitiesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAgentGetCapabilitiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAgentGetCapabilitiesResponseJSON) RawJSON() string {
	return r.raw
}

type AIAgentGetCapabilitiesResponseData struct {
	Description           string                                 `json:"description"`
	Enabled               bool                                   `json:"enabled"`
	Name                  string                                 `json:"name"`
	RequiresHumanApproval bool                                   `json:"requiresHumanApproval"`
	JSON                  aiAgentGetCapabilitiesResponseDataJSON `json:"-"`
}

// aiAgentGetCapabilitiesResponseDataJSON contains the JSON metadata for the struct
// [AIAgentGetCapabilitiesResponseData]
type aiAgentGetCapabilitiesResponseDataJSON struct {
	Description           apijson.Field
	Enabled               apijson.Field
	Name                  apijson.Field
	RequiresHumanApproval apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIAgentGetCapabilitiesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAgentGetCapabilitiesResponseDataJSON) RawJSON() string {
	return r.raw
}

type AIAgentGetPromptsResponse struct {
	SystemPrompt string                        `json:"systemPrompt"`
	Version      string                        `json:"version"`
	JSON         aiAgentGetPromptsResponseJSON `json:"-"`
}

// aiAgentGetPromptsResponseJSON contains the JSON metadata for the struct
// [AIAgentGetPromptsResponse]
type aiAgentGetPromptsResponseJSON struct {
	SystemPrompt apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AIAgentGetPromptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAgentGetPromptsResponseJSON) RawJSON() string {
	return r.raw
}

type AIAgentUpdatePromptsParams struct {
	SystemPrompt param.Field[string] `json:"systemPrompt,required"`
}

func (r AIAgentUpdatePromptsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
