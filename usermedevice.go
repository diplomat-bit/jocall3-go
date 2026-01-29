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

// List Connected Devices
func (r *UserMeDeviceService) List(ctx context.Context, opts ...option.RequestOption) (res *UserMeDeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// De-register a Device
func (r *UserMeDeviceService) Deregister(ctx context.Context, deviceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return
	}
	path := fmt.Sprintf("users/me/devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Register New Trusted Device
func (r *UserMeDeviceService) Register(ctx context.Context, body UserMeDeviceRegisterParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserMeDeviceListResponse struct {
	Data []UserMeDeviceListResponseData `json:"data"`
	JSON userMeDeviceListResponseJSON   `json:"-"`
}

// userMeDeviceListResponseJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponse]
type userMeDeviceListResponseJSON struct {
	Data        apijson.Field
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
	ID         string                                 `json:"id"`
	Os         string                                 `json:"os"`
	TrustLevel UserMeDeviceListResponseDataTrustLevel `json:"trustLevel"`
	Type       string                                 `json:"type"`
	JSON       userMeDeviceListResponseDataJSON       `json:"-"`
}

// userMeDeviceListResponseDataJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponseData]
type userMeDeviceListResponseDataJSON struct {
	ID          apijson.Field
	Os          apijson.Field
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

type UserMeDeviceListResponseDataTrustLevel string

const (
	UserMeDeviceListResponseDataTrustLevelTrusted   UserMeDeviceListResponseDataTrustLevel = "trusted"
	UserMeDeviceListResponseDataTrustLevelUntrusted UserMeDeviceListResponseDataTrustLevel = "untrusted"
)

func (r UserMeDeviceListResponseDataTrustLevel) IsKnown() bool {
	switch r {
	case UserMeDeviceListResponseDataTrustLevelTrusted, UserMeDeviceListResponseDataTrustLevelUntrusted:
		return true
	}
	return false
}

type UserMeDeviceRegisterParams struct {
	DeviceID  param.Field[string] `json:"deviceId,required"`
	Type      param.Field[string] `json:"type,required"`
	PushToken param.Field[string] `json:"pushToken"`
}

func (r UserMeDeviceRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
