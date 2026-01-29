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

// PaymentDomesticService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentDomesticService] method instead.
type PaymentDomesticService struct {
	Options []option.RequestOption
}

// NewPaymentDomesticService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentDomesticService(opts ...option.RequestOption) (r *PaymentDomesticService) {
	r = &PaymentDomesticService{}
	r.Options = opts
	return
}

// Execute ACH Transfer
func (r *PaymentDomesticService) ACH(ctx context.Context, body PaymentDomesticACHParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/domestic/ach"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Real-time Payment (RTP)
func (r *PaymentDomesticService) Rtp(ctx context.Context, body PaymentDomesticRtpParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/domestic/rtp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Execute Federal Wire
func (r *PaymentDomesticService) Wire(ctx context.Context, body PaymentDomesticWireParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/domestic/wire"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PaymentDomesticACHParams struct {
	Account param.Field[string]  `json:"account,required"`
	Amount  param.Field[float64] `json:"amount,required"`
	Routing param.Field[string]  `json:"routing,required"`
}

func (r PaymentDomesticACHParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentDomesticRtpParams struct {
	Amount      param.Field[float64] `json:"amount,required"`
	RecipientID param.Field[string]  `json:"recipientId,required"`
}

func (r PaymentDomesticRtpParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentDomesticWireParams struct {
	Account param.Field[string]  `json:"account,required"`
	Amount  param.Field[float64] `json:"amount,required"`
	Routing param.Field[string]  `json:"routing,required"`
}

func (r PaymentDomesticWireParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
