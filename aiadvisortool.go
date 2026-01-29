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

// AIAdvisorToolService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorToolService] method instead.
type AIAdvisorToolService struct {
	Options []option.RequestOption
}

// NewAIAdvisorToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorToolService(opts ...option.RequestOption) (r *AIAdvisorToolService) {
	r = &AIAdvisorToolService{}
	r.Options = opts
	return
}

// List AI-Executable Financial Tools
func (r *AIAdvisorToolService) List(ctx context.Context, opts ...option.RequestOption) (res *AIAdvisorToolListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Grant AI Execution Permission for Tool
func (r *AIAdvisorToolService) Enable(ctx context.Context, toolID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return
	}
	path := fmt.Sprintf("ai/advisor/tools/%s/enable", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type AIAdvisorToolListResponse struct {
	Data []interface{}                 `json:"data"`
	JSON aiAdvisorToolListResponseJSON `json:"-"`
}

// aiAdvisorToolListResponseJSON contains the JSON metadata for the struct
// [AIAdvisorToolListResponse]
type aiAdvisorToolListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorToolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorToolListResponseJSON) RawJSON() string {
	return r.raw
}
