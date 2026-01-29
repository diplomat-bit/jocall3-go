// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// SystemSandboxService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemSandboxService] method instead.
type SystemSandboxService struct {
	Options []option.RequestOption
}

// NewSystemSandboxService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemSandboxService(opts ...option.RequestOption) (r *SystemSandboxService) {
	r = &SystemSandboxService{}
	r.Options = opts
	return
}

// Reset Sandbox Ledger Data
func (r *SystemSandboxService) Reset(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "system/sandbox/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Force Specific API Error (For Testing)
func (r *SystemSandboxService) SimulateError(ctx context.Context, body SystemSandboxSimulateErrorParams, opts ...option.RequestOption) (res *SystemSandboxSimulateErrorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "system/sandbox/simulate-error"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SystemSandboxSimulateErrorResponse struct {
	Code      string                                 `json:"code,required"`
	Message   string                                 `json:"message,required"`
	Timestamp time.Time                              `json:"timestamp" format:"date-time"`
	JSON      systemSandboxSimulateErrorResponseJSON `json:"-"`
}

// systemSandboxSimulateErrorResponseJSON contains the JSON metadata for the struct
// [SystemSandboxSimulateErrorResponse]
type systemSandboxSimulateErrorResponseJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemSandboxSimulateErrorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemSandboxSimulateErrorResponseJSON) RawJSON() string {
	return r.raw
}

type SystemSandboxSimulateErrorParams struct {
	ErrorCode param.Field[int64] `json:"errorCode,required"`
}

func (r SystemSandboxSimulateErrorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
