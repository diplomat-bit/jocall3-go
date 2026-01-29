// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apiform"
	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// SystemVerificationService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemVerificationService] method instead.
type SystemVerificationService struct {
	Options []option.RequestOption
}

// NewSystemVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSystemVerificationService(opts ...option.RequestOption) (r *SystemVerificationService) {
	r = &SystemVerificationService{}
	r.Options = opts
	return
}

// Compare biometric samples
func (r *SystemVerificationService) CompareBiometrics(ctx context.Context, body SystemVerificationCompareBiometricsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "system/verification/biometric-comparison"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Verify identity document
func (r *SystemVerificationService) VerifyDocument(ctx context.Context, body SystemVerificationVerifyDocumentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "system/verification/document"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SystemVerificationCompareBiometricsParams struct {
	SampleA param.Field[string] `json:"sample_a,required"`
	SampleB param.Field[string] `json:"sample_b,required"`
}

func (r SystemVerificationCompareBiometricsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SystemVerificationVerifyDocumentParams struct {
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	Type param.Field[string]    `json:"type"`
}

func (r SystemVerificationVerifyDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
