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

// AIAdvisorService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorService] method instead.
type AIAdvisorService struct {
	Options []option.RequestOption
	Tools   *AIAdvisorToolService
}

// NewAIAdvisorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorService(opts ...option.RequestOption) (r *AIAdvisorService) {
	r = &AIAdvisorService{}
	r.Options = opts
	r.Tools = NewAIAdvisorToolService(opts...)
	return
}

// The primary orchestration point. Connects Postman Data to Gemini Logic.
func (r *AIAdvisorService) Chat(ctx context.Context, body AIAdvisorChatParams, opts ...option.RequestOption) (res *AIAdvisorChatResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Full Chat Transcript
func (r *AIAdvisorService) GetHistory(ctx context.Context, opts ...option.RequestOption) (res *AIAdvisorGetHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIAdvisorChatResponse struct {
	Reply            string                    `json:"reply"`
	SessionID        string                    `json:"sessionId"`
	SuggestedActions []interface{}             `json:"suggestedActions"`
	JSON             aiAdvisorChatResponseJSON `json:"-"`
}

// aiAdvisorChatResponseJSON contains the JSON metadata for the struct
// [AIAdvisorChatResponse]
type aiAdvisorChatResponseJSON struct {
	Reply            apijson.Field
	SessionID        apijson.Field
	SuggestedActions apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIAdvisorChatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorGetHistoryResponse struct {
	Messages []interface{}                   `json:"messages"`
	JSON     aiAdvisorGetHistoryResponseJSON `json:"-"`
}

// aiAdvisorGetHistoryResponseJSON contains the JSON metadata for the struct
// [AIAdvisorGetHistoryResponse]
type aiAdvisorGetHistoryResponseJSON struct {
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorGetHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorGetHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatParams struct {
	Message           param.Field[string]   `json:"message,required"`
	ContextAccountIDs param.Field[[]string] `json:"contextAccountIds"`
	Mode              param.Field[string]   `json:"mode"`
	Stream            param.Field[bool]     `json:"stream"`
}

func (r AIAdvisorChatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
