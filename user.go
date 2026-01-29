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

// UserService contains methods and other services that help with interacting with
// the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	PasswordReset *UserPasswordResetService
	Me            *UserMeService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.PasswordReset = NewUserPasswordResetService(opts...)
	r.Me = NewUserMeService(opts...)
	return
}

// User Login and Session Creation
func (r *UserService) Login(ctx context.Context, body UserLoginParams, opts ...option.RequestOption) (res *UserLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Terminate User Session
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Initiates the onboarding process. Requires unique email and name.
func (r *UserService) Register(ctx context.Context, body UserRegisterParams, opts ...option.RequestOption) (res *UserRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserLoginResponse struct {
	// JWT for bearerAuth
	AccessToken string `json:"accessToken,required"`
	// Seconds until expiry
	ExpiresIn    int64                 `json:"expiresIn,required"`
	RefreshToken string                `json:"refreshToken,required"`
	TokenType    string                `json:"tokenType,required"`
	JSON         userLoginResponseJSON `json:"-"`
}

// userLoginResponseJSON contains the JSON metadata for the struct
// [UserLoginResponse]
type userLoginResponseJSON struct {
	AccessToken  apijson.Field
	ExpiresIn    apijson.Field
	RefreshToken apijson.Field
	TokenType    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserLoginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoginResponseJSON) RawJSON() string {
	return r.raw
}

type UserRegisterResponse struct {
	ID               string                             `json:"id,required"`
	Email            string                             `json:"email,required" format:"email"`
	IdentityVerified bool                               `json:"identityVerified,required"`
	Name             string                             `json:"name,required"`
	Address          UserRegisterResponseAddress        `json:"address"`
	Phone            string                             `json:"phone"`
	Preferences      UserRegisterResponsePreferences    `json:"preferences"`
	SecurityStatus   UserRegisterResponseSecurityStatus `json:"securityStatus"`
	JSON             userRegisterResponseJSON           `json:"-"`
}

// userRegisterResponseJSON contains the JSON metadata for the struct
// [UserRegisterResponse]
type userRegisterResponseJSON struct {
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

func (r *UserRegisterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRegisterResponseJSON) RawJSON() string {
	return r.raw
}

type UserRegisterResponseAddress struct {
	City    string                          `json:"city"`
	Country string                          `json:"country"`
	State   string                          `json:"state"`
	Street  string                          `json:"street"`
	Zip     string                          `json:"zip"`
	JSON    userRegisterResponseAddressJSON `json:"-"`
}

// userRegisterResponseAddressJSON contains the JSON metadata for the struct
// [UserRegisterResponseAddress]
type userRegisterResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserRegisterResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRegisterResponseAddressJSON) RawJSON() string {
	return r.raw
}

type UserRegisterResponsePreferences struct {
	NotificationChannels interface{}                         `json:"notificationChannels"`
	Theme                string                              `json:"theme"`
	JSON                 userRegisterResponsePreferencesJSON `json:"-"`
}

// userRegisterResponsePreferencesJSON contains the JSON metadata for the struct
// [UserRegisterResponsePreferences]
type userRegisterResponsePreferencesJSON struct {
	NotificationChannels apijson.Field
	Theme                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserRegisterResponsePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRegisterResponsePreferencesJSON) RawJSON() string {
	return r.raw
}

type UserRegisterResponseSecurityStatus struct {
	LastLogin        time.Time                              `json:"lastLogin" format:"date-time"`
	TwoFactorEnabled bool                                   `json:"twoFactorEnabled"`
	JSON             userRegisterResponseSecurityStatusJSON `json:"-"`
}

// userRegisterResponseSecurityStatusJSON contains the JSON metadata for the struct
// [UserRegisterResponseSecurityStatus]
type userRegisterResponseSecurityStatusJSON struct {
	LastLogin        apijson.Field
	TwoFactorEnabled apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserRegisterResponseSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRegisterResponseSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserLoginParams struct {
	Email    param.Field[string] `json:"email,required" format:"email"`
	Password param.Field[string] `json:"password,required" format:"password"`
}

func (r UserLoginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRegisterParams struct {
	// Primary login email
	Email param.Field[string] `json:"email,required" format:"email"`
	// Full legal name
	Name param.Field[string] `json:"name,required"`
	// Secure hashable string
	Password param.Field[string]                    `json:"password,required" format:"password"`
	Address  param.Field[UserRegisterParamsAddress] `json:"address"`
	// International format phone number
	Phone param.Field[string] `json:"phone"`
}

func (r UserRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRegisterParamsAddress struct {
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
	Zip     param.Field[string] `json:"zip"`
}

func (r UserRegisterParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
