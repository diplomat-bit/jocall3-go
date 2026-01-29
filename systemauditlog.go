// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// SystemAuditLogService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemAuditLogService] method instead.
type SystemAuditLogService struct {
	Options []option.RequestOption
}

// NewSystemAuditLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemAuditLogService(opts ...option.RequestOption) (r *SystemAuditLogService) {
	r = &SystemAuditLogService{}
	r.Options = opts
	return
}

// Get Immutable System Audit Trail
func (r *SystemAuditLogService) List(ctx context.Context, query SystemAuditLogListParams, opts ...option.RequestOption) (res *SystemAuditLogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "system/audit-logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SystemAuditLogListResponse struct {
	Data []SystemAuditLogListResponseData `json:"data"`
	JSON systemAuditLogListResponseJSON   `json:"-"`
}

// systemAuditLogListResponseJSON contains the JSON metadata for the struct
// [SystemAuditLogListResponse]
type systemAuditLogListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemAuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemAuditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type SystemAuditLogListResponseData struct {
	ID        string                             `json:"id"`
	Action    string                             `json:"action"`
	Actor     string                             `json:"actor"`
	Impact    string                             `json:"impact"`
	Timestamp time.Time                          `json:"timestamp" format:"date-time"`
	JSON      systemAuditLogListResponseDataJSON `json:"-"`
}

// systemAuditLogListResponseDataJSON contains the JSON metadata for the struct
// [SystemAuditLogListResponseData]
type systemAuditLogListResponseDataJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Impact      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemAuditLogListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemAuditLogListResponseDataJSON) RawJSON() string {
	return r.raw
}

type SystemAuditLogListParams struct {
	ActorID param.Field[string] `query:"actorId"`
	Limit   param.Field[int64]  `query:"limit"`
	Offset  param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [SystemAuditLogListParams]'s query parameters as
// `url.Values`.
func (r SystemAuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
