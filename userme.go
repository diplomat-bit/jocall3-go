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

// UserMeService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeService] method instead.
type UserMeService struct {
	Options     []option.RequestOption
	Preferences *UserMePreferenceService
	Security    *UserMeSecurityService
	Devices     *UserMeDeviceService
	Biometrics  *UserMeBiometricService
}

// NewUserMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserMeService(opts ...option.RequestOption) (r *UserMeService) {
	r = &UserMeService{}
	r.Options = opts
	r.Preferences = NewUserMePreferenceService(opts...)
	r.Security = NewUserMeSecurityService(opts...)
	r.Devices = NewUserMeDeviceService(opts...)
	r.Biometrics = NewUserMeBiometricService(opts...)
	return
}

// Retrieve Comprehensive Current User Profile
func (r *UserMeService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserMeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Current User Profile
func (r *UserMeService) Update(ctx context.Context, body UserMeUpdateParams, opts ...option.RequestOption) (res *UserMeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete User Account
func (r *UserMeService) Delete(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type UserMeGetResponse struct {
	ID               string                          `json:"id,required"`
	Email            string                          `json:"email,required" format:"email"`
	IdentityVerified bool                            `json:"identityVerified,required"`
	Name             string                          `json:"name,required"`
	Address          UserMeGetResponseAddress        `json:"address"`
	Phone            string                          `json:"phone"`
	Preferences      UserMeGetResponsePreferences    `json:"preferences"`
	SecurityStatus   UserMeGetResponseSecurityStatus `json:"securityStatus"`
	JSON             userMeGetResponseJSON           `json:"-"`
}

// userMeGetResponseJSON contains the JSON metadata for the struct
// [UserMeGetResponse]
type userMeGetResponseJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	IdentityVerified apijson.Field
	Name             apijson.Field
	Address          apijson.Field
	Phone            apijson.Field
	Preferences      apijson.Field
	SecurityStatus   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserMeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeGetResponseAddress struct {
	City    string                       `json:"city"`
	Country string                       `json:"country"`
	State   string                       `json:"state"`
	Street  string                       `json:"street"`
	Zip     string                       `json:"zip"`
	JSON    userMeGetResponseAddressJSON `json:"-"`
}

// userMeGetResponseAddressJSON contains the JSON metadata for the struct
// [UserMeGetResponseAddress]
type userMeGetResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeGetResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseAddressJSON) RawJSON() string {
	return r.raw
}

type UserMeGetResponsePreferences struct {
	NotificationChannels interface{}                      `json:"notificationChannels"`
	Theme                string                           `json:"theme"`
	JSON                 userMeGetResponsePreferencesJSON `json:"-"`
}

// userMeGetResponsePreferencesJSON contains the JSON metadata for the struct
// [UserMeGetResponsePreferences]
type userMeGetResponsePreferencesJSON struct {
	NotificationChannels apijson.Field
	Theme                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMeGetResponsePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponsePreferencesJSON) RawJSON() string {
	return r.raw
}

type UserMeGetResponseSecurityStatus struct {
	LastLogin        time.Time                           `json:"lastLogin" format:"date-time"`
	TwoFactorEnabled bool                                `json:"twoFactorEnabled"`
	JSON             userMeGetResponseSecurityStatusJSON `json:"-"`
}

// userMeGetResponseSecurityStatusJSON contains the JSON metadata for the struct
// [UserMeGetResponseSecurityStatus]
type userMeGetResponseSecurityStatusJSON struct {
	LastLogin        apijson.Field
	TwoFactorEnabled apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserMeGetResponseSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponse struct {
	ID               string                             `json:"id,required"`
	Email            string                             `json:"email,required" format:"email"`
	IdentityVerified bool                               `json:"identityVerified,required"`
	Name             string                             `json:"name,required"`
	Address          UserMeUpdateResponseAddress        `json:"address"`
	Phone            string                             `json:"phone"`
	Preferences      UserMeUpdateResponsePreferences    `json:"preferences"`
	SecurityStatus   UserMeUpdateResponseSecurityStatus `json:"securityStatus"`
	JSON             userMeUpdateResponseJSON           `json:"-"`
}

// userMeUpdateResponseJSON contains the JSON metadata for the struct
// [UserMeUpdateResponse]
type userMeUpdateResponseJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	IdentityVerified apijson.Field
	Name             apijson.Field
	Address          apijson.Field
	Phone            apijson.Field
	Preferences      apijson.Field
	SecurityStatus   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserMeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponseAddress struct {
	City    string                          `json:"city"`
	Country string                          `json:"country"`
	State   string                          `json:"state"`
	Street  string                          `json:"street"`
	Zip     string                          `json:"zip"`
	JSON    userMeUpdateResponseAddressJSON `json:"-"`
}

// userMeUpdateResponseAddressJSON contains the JSON metadata for the struct
// [UserMeUpdateResponseAddress]
type userMeUpdateResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeUpdateResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseAddressJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponsePreferences struct {
	NotificationChannels interface{}                         `json:"notificationChannels"`
	Theme                string                              `json:"theme"`
	JSON                 userMeUpdateResponsePreferencesJSON `json:"-"`
}

// userMeUpdateResponsePreferencesJSON contains the JSON metadata for the struct
// [UserMeUpdateResponsePreferences]
type userMeUpdateResponsePreferencesJSON struct {
	NotificationChannels apijson.Field
	Theme                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMeUpdateResponsePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponsePreferencesJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponseSecurityStatus struct {
	LastLogin        time.Time                              `json:"lastLogin" format:"date-time"`
	TwoFactorEnabled bool                                   `json:"twoFactorEnabled"`
	JSON             userMeUpdateResponseSecurityStatusJSON `json:"-"`
}

// userMeUpdateResponseSecurityStatusJSON contains the JSON metadata for the struct
// [UserMeUpdateResponseSecurityStatus]
type userMeUpdateResponseSecurityStatusJSON struct {
	LastLogin        apijson.Field
	TwoFactorEnabled apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserMeUpdateResponseSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateParams struct {
	Address param.Field[UserMeUpdateParamsAddress] `json:"address"`
	Name    param.Field[string]                    `json:"name"`
	Phone   param.Field[string]                    `json:"phone"`
}

func (r UserMeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeUpdateParamsAddress struct {
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
	Zip     param.Field[string] `json:"zip"`
}

func (r UserMeUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
