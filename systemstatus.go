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

// SystemStatusService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemStatusService] method instead.
type SystemStatusService struct {
	Options []option.RequestOption
}

// NewSystemStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemStatusService(opts ...option.RequestOption) (r *SystemStatusService) {
	r = &SystemStatusService{}
	r.Options = opts
	return
}

// Get Global Infrastructure Health
func (r *SystemStatusService) Get(ctx context.Context, opts ...option.RequestOption) (res *SystemStatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "system/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SystemStatusGetResponse struct {
	ActiveNodes       int64                       `json:"activeNodes"`
	APIStatus         string                      `json:"apiStatus"`
	GeminiUptime      string                      `json:"geminiUptime"`
	MockServerLatency int64                       `json:"mockServerLatency"`
	JSON              systemStatusGetResponseJSON `json:"-"`
}

// systemStatusGetResponseJSON contains the JSON metadata for the struct
// [SystemStatusGetResponse]
type systemStatusGetResponseJSON struct {
	ActiveNodes       apijson.Field
	APIStatus         apijson.Field
	GeminiUptime      apijson.Field
	MockServerLatency apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SystemStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemStatusGetResponseJSON) RawJSON() string {
	return r.raw
}
