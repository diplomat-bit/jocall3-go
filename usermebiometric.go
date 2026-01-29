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

// UserMeBiometricService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeBiometricService] method instead.
type UserMeBiometricService struct {
	Options []option.RequestOption
}

// NewUserMeBiometricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeBiometricService(opts ...option.RequestOption) (r *UserMeBiometricService) {
	r = &UserMeBiometricService{}
	r.Options = opts
	return
}

// Remove All Biometric Data
func (r *UserMeBiometricService) Delete(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Enroll New Biometric Signature
func (r *UserMeBiometricService) Enroll(ctx context.Context, body UserMeBiometricEnrollParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me/biometrics/enroll"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get Biometric Enrollment Status
func (r *UserMeBiometricService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *UserMeBiometricGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Verify Biometric Data for Sensitive Operations
func (r *UserMeBiometricService) Verify(ctx context.Context, body UserMeBiometricVerifyParams, opts ...option.RequestOption) (res *UserMeBiometricVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserMeBiometricGetStatusResponse struct {
	BiometricsEnrolled bool                                 `json:"biometricsEnrolled"`
	LastUsed           time.Time                            `json:"lastUsed" format:"date-time"`
	JSON               userMeBiometricGetStatusResponseJSON `json:"-"`
}

// userMeBiometricGetStatusResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricGetStatusResponse]
type userMeBiometricGetStatusResponseJSON struct {
	BiometricsEnrolled apijson.Field
	LastUsed           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeBiometricGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeBiometricGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeBiometricVerifyResponse struct {
	VerificationStatus string                            `json:"verificationStatus"`
	JSON               userMeBiometricVerifyResponseJSON `json:"-"`
}

// userMeBiometricVerifyResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricVerifyResponse]
type userMeBiometricVerifyResponseJSON struct {
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeBiometricVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeBiometricVerifyResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeBiometricEnrollParams struct {
	BiometricType param.Field[UserMeBiometricEnrollParamsBiometricType] `json:"biometricType,required"`
	// Public key or hash of signature
	Signature param.Field[string] `json:"signature,required"`
}

func (r UserMeBiometricEnrollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeBiometricEnrollParamsBiometricType string

const (
	UserMeBiometricEnrollParamsBiometricTypeFingerprint       UserMeBiometricEnrollParamsBiometricType = "fingerprint"
	UserMeBiometricEnrollParamsBiometricTypeFacialRecognition UserMeBiometricEnrollParamsBiometricType = "facial_recognition"
)

func (r UserMeBiometricEnrollParamsBiometricType) IsKnown() bool {
	switch r {
	case UserMeBiometricEnrollParamsBiometricTypeFingerprint, UserMeBiometricEnrollParamsBiometricTypeFacialRecognition:
		return true
	}
	return false
}

type UserMeBiometricVerifyParams struct {
	BiometricSignature param.Field[string] `json:"biometricSignature,required"`
}

func (r UserMeBiometricVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
