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

// UserMeSecurityService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeSecurityService] method instead.
type UserMeSecurityService struct {
	Options []option.RequestOption
}

// NewUserMeSecurityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeSecurityService(opts ...option.RequestOption) (r *UserMeSecurityService) {
	r = &UserMeSecurityService{}
	r.Options = opts
	return
}

// Retrieve Security Access Logs
func (r *UserMeSecurityService) GetLog(ctx context.Context, query UserMeSecurityGetLogParams, opts ...option.RequestOption) (res *UserMeSecurityGetLogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/security/log"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Rotate API/Access Keys
func (r *UserMeSecurityService) RotateKeys(ctx context.Context, opts ...option.RequestOption) (res *UserMeSecurityRotateKeysResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/security/rotate-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type UserMeSecurityGetLogResponse struct {
	Data []UserMeSecurityGetLogResponseData `json:"data"`
	JSON userMeSecurityGetLogResponseJSON   `json:"-"`
}

// userMeSecurityGetLogResponseJSON contains the JSON metadata for the struct
// [UserMeSecurityGetLogResponse]
type userMeSecurityGetLogResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeSecurityGetLogResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeSecurityGetLogResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeSecurityGetLogResponseData struct {
	Event     string                               `json:"event"`
	IPAddress string                               `json:"ipAddress"`
	Location  interface{}                          `json:"location"`
	Timestamp time.Time                            `json:"timestamp" format:"date-time"`
	JSON      userMeSecurityGetLogResponseDataJSON `json:"-"`
}

// userMeSecurityGetLogResponseDataJSON contains the JSON metadata for the struct
// [UserMeSecurityGetLogResponseData]
type userMeSecurityGetLogResponseDataJSON struct {
	Event       apijson.Field
	IPAddress   apijson.Field
	Location    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeSecurityGetLogResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeSecurityGetLogResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserMeSecurityRotateKeysResponse struct {
	NewExpiry time.Time                            `json:"newExpiry" format:"date-time"`
	Status    string                               `json:"status"`
	JSON      userMeSecurityRotateKeysResponseJSON `json:"-"`
}

// userMeSecurityRotateKeysResponseJSON contains the JSON metadata for the struct
// [UserMeSecurityRotateKeysResponse]
type userMeSecurityRotateKeysResponseJSON struct {
	NewExpiry   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeSecurityRotateKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeSecurityRotateKeysResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeSecurityGetLogParams struct {
	Limit  param.Field[int64] `query:"limit"`
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [UserMeSecurityGetLogParams]'s query parameters as
// `url.Values`.
func (r UserMeSecurityGetLogParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
