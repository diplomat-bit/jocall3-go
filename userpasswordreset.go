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

// UserPasswordResetService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPasswordResetService] method instead.
type UserPasswordResetService struct {
	Options []option.RequestOption
}

// NewUserPasswordResetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserPasswordResetService(opts ...option.RequestOption) (r *UserPasswordResetService) {
	r = &UserPasswordResetService{}
	r.Options = opts
	return
}

// Confirms the password reset using the received verification code and sets a new
// password.
func (r *UserPasswordResetService) Confirm(ctx context.Context, body UserPasswordResetConfirmParams, opts ...option.RequestOption) (res *UserPasswordResetConfirmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/password-reset/confirm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Starts the password reset flow by sending a verification code or link to the
// user's registered email or phone.
func (r *UserPasswordResetService) Initiate(ctx context.Context, body UserPasswordResetInitiateParams, opts ...option.RequestOption) (res *UserPasswordResetInitiateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/password-reset/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserPasswordResetConfirmResponse = interface{}

type UserPasswordResetInitiateResponse = interface{}

type UserPasswordResetConfirmParams struct {
}

func (r UserPasswordResetConfirmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserPasswordResetInitiateParams struct {
}

func (r UserPasswordResetInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
