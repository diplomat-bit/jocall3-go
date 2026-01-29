// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
	"github.com/diplomat-bit/jocall3-go/shared"
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

func (r *UserMeService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserMeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *UserMeService) Update(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

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
	Address          shared.Address                  `json:"address"`
	Preferences      map[string]interface{}          `json:"preferences"`
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
