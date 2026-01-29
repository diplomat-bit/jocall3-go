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
	"github.com/diplomat-bit/jocall3-go/shared"
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

func (r *UserService) Login(ctx context.Context, body UserLoginParams, opts ...option.RequestOption) (res *UserLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

func (r *UserService) Register(ctx context.Context, body UserRegisterParams, opts ...option.RequestOption) (res *UserRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserLoginResponse struct {
	AccessToken  string                `json:"accessToken,required"`
	ExpiresIn    int64                 `json:"expiresIn"`
	RefreshToken string                `json:"refreshToken"`
	TokenType    string                `json:"tokenType"`
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
	Address          shared.Address                     `json:"address"`
	Preferences      map[string]interface{}             `json:"preferences"`
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
	Email    param.Field[string] `json:"email,required"`
	Password param.Field[string] `json:"password,required"`
}

func (r UserLoginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRegisterParams struct {
	Email    param.Field[string] `json:"email,required"`
	Name     param.Field[string] `json:"name,required"`
	Password param.Field[string] `json:"password,required"`
}

func (r UserRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
