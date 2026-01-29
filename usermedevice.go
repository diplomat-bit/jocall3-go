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

// UserMeDeviceService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeDeviceService] method instead.
type UserMeDeviceService struct {
	Options []option.RequestOption
}

// NewUserMeDeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeDeviceService(opts ...option.RequestOption) (r *UserMeDeviceService) {
	r = &UserMeDeviceService{}
	r.Options = opts
	return
}

// Retrieves a list of all devices linked to the user's account, including mobile
// phones, tablets, and desktops, indicating their last active status and security
// posture.
func (r *UserMeDeviceService) List(ctx context.Context, query UserMeDeviceListParams, opts ...option.RequestOption) (res *UserMeDeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserMeDeviceListResponse struct {
	Data       []UserMeDeviceListResponseData `json:"data,required"`
	Limit      int64                          `json:"limit,required"`
	Offset     int64                          `json:"offset,required"`
	Total      int64                          `json:"total,required"`
	NextOffset int64                          `json:"nextOffset"`
	JSON       userMeDeviceListResponseJSON   `json:"-"`
}

// userMeDeviceListResponseJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponse]
type userMeDeviceListResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListResponseData struct {
	ID         string                           `json:"id"`
	IPAddress  string                           `json:"ipAddress"`
	LastActive time.Time                        `json:"lastActive" format:"date-time"`
	Model      string                           `json:"model"`
	Os         string                           `json:"os"`
	PushToken  string                           `json:"pushToken"`
	TrustLevel string                           `json:"trustLevel"`
	Type       string                           `json:"type"`
	JSON       userMeDeviceListResponseDataJSON `json:"-"`
}

// userMeDeviceListResponseDataJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponseData]
type userMeDeviceListResponseDataJSON struct {
	ID          apijson.Field
	IPAddress   apijson.Field
	LastActive  apijson.Field
	Model       apijson.Field
	Os          apijson.Field
	PushToken   apijson.Field
	TrustLevel  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeDeviceListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeDeviceListResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [UserMeDeviceListParams]'s query parameters as `url.Values`.
func (r UserMeDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
