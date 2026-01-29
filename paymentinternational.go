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

// Get international payment status
func (r *PaymentInternationalService) GetStatus(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *PaymentInternationalGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return
	}
	path := fmt.Sprintf("payments/international/%s/status", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// EU SEPA Credit Transfer
func (r *PaymentInternationalService) SendSepa(ctx context.Context, body PaymentInternationalSendSepaParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/international/sepa"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Global SWIFT Transaction
func (r *PaymentInternationalService) SendSwift(ctx context.Context, body PaymentInternationalSendSwiftParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/international/swift"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PaymentInternationalGetStatusResponse struct {
	FxRate float64                                   `json:"fx_rate"`
	Status string                                    `json:"status"`
	JSON   paymentInternationalGetStatusResponseJSON `json:"-"`
}

// paymentInternationalGetStatusResponseJSON contains the JSON metadata for the
// struct [PaymentInternationalGetStatusResponse]
type paymentInternationalGetStatusResponseJSON struct {
	FxRate      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentInternationalGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentInternationalGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentInternationalSendSepaParams struct {
	Amount param.Field[float64] `json:"amount,required"`
	Iban   param.Field[string]  `json:"iban,required"`
}

func (r PaymentInternationalSendSepaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentInternationalSendSwiftParams struct {
	Amount   param.Field[float64] `json:"amount,required"`
	Bic      param.Field[string]  `json:"bic,required"`
	Currency param.Field[string]  `json:"currency,required"`
	Iban     param.Field[string]  `json:"iban,required"`
}

func (r PaymentInternationalSendSwiftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
