// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// PaymentFxService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentFxService] method instead.
type PaymentFxService struct {
	Options []option.RequestOption
}

// NewPaymentFxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentFxService(opts ...option.RequestOption) (r *PaymentFxService) {
	r = &PaymentFxService{}
	r.Options = opts
	return
}

// Book a Forward FX Deal
func (r *PaymentFxService) BookDeal(ctx context.Context, body PaymentFxBookDealParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/fx/deals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Execute Currency Conversion
func (r *PaymentFxService) ConvertCurrency(ctx context.Context, body PaymentFxConvertCurrencyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "payments/fx/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Market FX Rates
func (r *PaymentFxService) GetRates(ctx context.Context, query PaymentFxGetRatesParams, opts ...option.RequestOption) (res *PaymentFxGetRatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/fx/rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaymentFxGetRatesResponse struct {
	MidRate   float64                       `json:"midRate"`
	Timestamp time.Time                     `json:"timestamp" format:"date-time"`
	JSON      paymentFxGetRatesResponseJSON `json:"-"`
}

// paymentFxGetRatesResponseJSON contains the JSON metadata for the struct
// [PaymentFxGetRatesResponse]
type paymentFxGetRatesResponseJSON struct {
	MidRate     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentFxBookDealParams struct {
	Amount    param.Field[float64]   `json:"amount,required"`
	Pair      param.Field[string]    `json:"pair,required"`
	ValueDate param.Field[time.Time] `json:"valueDate,required" format:"date"`
}

func (r PaymentFxBookDealParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFxConvertCurrencyParams struct {
	Amount param.Field[float64] `json:"amount,required"`
	From   param.Field[string]  `json:"from,required"`
	To     param.Field[string]  `json:"to,required"`
}

func (r PaymentFxConvertCurrencyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFxGetRatesParams struct {
	Pair param.Field[string] `query:"pair,required"`
}

// URLQuery serializes [PaymentFxGetRatesParams]'s query parameters as
// `url.Values`.
func (r PaymentFxGetRatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
