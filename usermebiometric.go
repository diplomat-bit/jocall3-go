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

// Retrieves the current status of biometric enrollments for the authenticated
// user.
func (r *UserMeBiometricService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *UserMeBiometricGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Performs real-time biometric verification to authorize sensitive actions or
// access protected resources, using a one-time biometric signature.
func (r *UserMeBiometricService) Verify(ctx context.Context, body UserMeBiometricVerifyParams, opts ...option.RequestOption) (res *UserMeBiometricVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Current biometric enrollment status for a user.
type UserMeBiometricGetStatusResponse struct {
	BiometricsEnrolled bool                                                `json:"biometricsEnrolled,required"`
	EnrolledBiometrics []UserMeBiometricGetStatusResponseEnrolledBiometric `json:"enrolledBiometrics,required"`
	LastUsed           time.Time                                           `json:"lastUsed" format:"date-time"`
	JSON               userMeBiometricGetStatusResponseJSON                `json:"-"`
}

// userMeBiometricGetStatusResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricGetStatusResponse]
type userMeBiometricGetStatusResponseJSON struct {
	BiometricsEnrolled apijson.Field
	EnrolledBiometrics apijson.Field
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

type UserMeBiometricGetStatusResponseEnrolledBiometric struct {
	DeviceID       string                                                `json:"deviceId"`
	EnrollmentDate time.Time                                             `json:"enrollmentDate" format:"date-time"`
	Type           string                                                `json:"type"`
	JSON           userMeBiometricGetStatusResponseEnrolledBiometricJSON `json:"-"`
}

// userMeBiometricGetStatusResponseEnrolledBiometricJSON contains the JSON metadata
// for the struct [UserMeBiometricGetStatusResponseEnrolledBiometric]
type userMeBiometricGetStatusResponseEnrolledBiometricJSON struct {
	DeviceID       apijson.Field
	EnrollmentDate apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserMeBiometricGetStatusResponseEnrolledBiometric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeBiometricGetStatusResponseEnrolledBiometricJSON) RawJSON() string {
	return r.raw
}

type UserMeBiometricVerifyResponse struct {
	Message            string                            `json:"message"`
	VerificationStatus string                            `json:"verificationStatus"`
	JSON               userMeBiometricVerifyResponseJSON `json:"-"`
}

// userMeBiometricVerifyResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricVerifyResponse]
type userMeBiometricVerifyResponseJSON struct {
	Message            apijson.Field
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

type UserMeBiometricVerifyParams struct {
	BiometricSignature param.Field[string] `json:"biometricSignature,required"`
	BiometricType      param.Field[string] `json:"biometricType,required"`
	DeviceID           param.Field[string] `json:"deviceId,required"`
}

func (r UserMeBiometricVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
