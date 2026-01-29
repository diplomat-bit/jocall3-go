// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// PaymentInternationalService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentInternationalService] method instead.
type PaymentInternationalService struct {
	Options []option.RequestOption
}

// NewPaymentInternationalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentInternationalService(opts ...option.RequestOption) (r *PaymentInternationalService) {
	r = &PaymentInternationalService{}
	r.Options = opts
	return
}

// EU SEPA Credit Transfer
func (r *PaymentInternationalService) Sepa(ctx context.Context, body PaymentInternationalSepaParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/international/sepa"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Global SWIFT Transaction
func (r *PaymentInternationalService) Swift(ctx context.Context, body PaymentInternationalSwiftParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/international/swift"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PaymentInternationalSepaParams struct {
	Amount param.Field[float64] `json:"amount,required"`
	Iban   param.Field[string]  `json:"iban,required"`
}

func (r PaymentInternationalSepaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentInternationalSwiftParams struct {
	Amount   param.Field[float64] `json:"amount,required"`
	Bic      param.Field[string]  `json:"bic,required"`
	Currency param.Field[string]  `json:"currency,required"`
	Iban     param.Field[string]  `json:"iban,required"`
}

func (r PaymentInternationalSwiftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
