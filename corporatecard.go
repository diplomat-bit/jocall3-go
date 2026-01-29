// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
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

// CorporateCardService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateCardService] method instead.
type CorporateCardService struct {
	Options  []option.RequestOption
	Controls *CorporateCardControlService
}

// NewCorporateCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateCardService(opts ...option.RequestOption) (r *CorporateCardService) {
	r = &CorporateCardService{}
	r.Options = opts
	r.Controls = NewCorporateCardControlService(opts...)
	return
}

// List all corporate cards
func (r *CorporateCardService) List(ctx context.Context, query CorporateCardListParams, opts ...option.RequestOption) (res *CorporateCardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Toggle Card Lock
func (r *CorporateCardService) Freeze(ctx context.Context, cardID string, body CorporateCardFreezeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/freeze", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Request Physical Corporate Card
func (r *CorporateCardService) IssuePhysical(ctx context.Context, body CorporateCardIssuePhysicalParams, opts ...option.RequestOption) (res *CorporateCardIssuePhysicalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards/physical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Issue Corporate Virtual Card
func (r *CorporateCardService) IssueVirtual(ctx context.Context, body CorporateCardIssueVirtualParams, opts ...option.RequestOption) (res *CorporateCardIssueVirtualResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards/virtual"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get card transactions
func (r *CorporateCardService) ListTransactions(ctx context.Context, cardID string, opts ...option.RequestOption) (res *CorporateCardListTransactionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/transactions", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CorporateCardListResponse struct {
	Data  []CorporateCardListResponseData `json:"data"`
	Total int64                           `json:"total"`
	JSON  corporateCardListResponseJSON   `json:"-"`
}

// corporateCardListResponseJSON contains the JSON metadata for the struct
// [CorporateCardListResponse]
type corporateCardListResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListResponseData struct {
	ID             string                            `json:"id,required"`
	CardNumberMask string                            `json:"cardNumberMask,required"`
	HolderName     string                            `json:"holderName,required"`
	Status         string                            `json:"status,required"`
	Controls       map[string]interface{}            `json:"controls"`
	Frozen         bool                              `json:"frozen"`
	JSON           corporateCardListResponseDataJSON `json:"-"`
}

// corporateCardListResponseDataJSON contains the JSON metadata for the struct
// [CorporateCardListResponseData]
type corporateCardListResponseDataJSON struct {
	ID             apijson.Field
	CardNumberMask apijson.Field
	HolderName     apijson.Field
	Status         apijson.Field
	Controls       apijson.Field
	Frozen         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateCardListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CorporateCardIssuePhysicalResponse struct {
	ID             string                                 `json:"id,required"`
	CardNumberMask string                                 `json:"cardNumberMask,required"`
	HolderName     string                                 `json:"holderName,required"`
	Status         string                                 `json:"status,required"`
	Controls       map[string]interface{}                 `json:"controls"`
	Frozen         bool                                   `json:"frozen"`
	JSON           corporateCardIssuePhysicalResponseJSON `json:"-"`
}

// corporateCardIssuePhysicalResponseJSON contains the JSON metadata for the struct
// [CorporateCardIssuePhysicalResponse]
type corporateCardIssuePhysicalResponseJSON struct {
	ID             apijson.Field
	CardNumberMask apijson.Field
	HolderName     apijson.Field
	Status         apijson.Field
	Controls       apijson.Field
	Frozen         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateCardIssuePhysicalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardIssuePhysicalResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardIssueVirtualResponse struct {
	ID             string                                `json:"id,required"`
	CardNumberMask string                                `json:"cardNumberMask,required"`
	HolderName     string                                `json:"holderName,required"`
	Status         string                                `json:"status,required"`
	Controls       map[string]interface{}                `json:"controls"`
	Frozen         bool                                  `json:"frozen"`
	JSON           corporateCardIssueVirtualResponseJSON `json:"-"`
}

// corporateCardIssueVirtualResponseJSON contains the JSON metadata for the struct
// [CorporateCardIssueVirtualResponse]
type corporateCardIssueVirtualResponseJSON struct {
	ID             apijson.Field
	CardNumberMask apijson.Field
	HolderName     apijson.Field
	Status         apijson.Field
	Controls       apijson.Field
	Frozen         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CorporateCardIssueVirtualResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardIssueVirtualResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListTransactionsResponse struct {
	Data []CorporateCardListTransactionsResponseData `json:"data"`
	JSON corporateCardListTransactionsResponseJSON   `json:"-"`
}

// corporateCardListTransactionsResponseJSON contains the JSON metadata for the
// struct [CorporateCardListTransactionsResponse]
type corporateCardListTransactionsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateCardListTransactionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListTransactionsResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListTransactionsResponseData struct {
	ID          string                                        `json:"id,required"`
	Amount      float64                                       `json:"amount,required"`
	Currency    string                                        `json:"currency,required"`
	Date        time.Time                                     `json:"date,required" format:"date"`
	Description string                                        `json:"description,required"`
	Category    string                                        `json:"category"`
	Notes       string                                        `json:"notes"`
	JSON        corporateCardListTransactionsResponseDataJSON `json:"-"`
}

// corporateCardListTransactionsResponseDataJSON contains the JSON metadata for the
// struct [CorporateCardListTransactionsResponseData]
type corporateCardListTransactionsResponseDataJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Date        apijson.Field
	Description apijson.Field
	Category    apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateCardListTransactionsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListTransactionsResponseDataJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListParams struct {
	Limit  param.Field[int64] `query:"limit"`
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CorporateCardListParams]'s query parameters as
// `url.Values`.
func (r CorporateCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateCardFreezeParams struct {
	Frozen param.Field[bool] `json:"frozen,required"`
}

func (r CorporateCardFreezeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardIssuePhysicalParams struct {
	HolderName      param.Field[string]                                          `json:"holderName,required"`
	ShippingAddress param.Field[CorporateCardIssuePhysicalParamsShippingAddress] `json:"shippingAddress,required"`
}

func (r CorporateCardIssuePhysicalParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardIssuePhysicalParamsShippingAddress struct {
	City    param.Field[string] `json:"city,required"`
	Country param.Field[string] `json:"country,required"`
	Street  param.Field[string] `json:"street,required"`
	State   param.Field[string] `json:"state"`
	Zip     param.Field[string] `json:"zip"`
}

func (r CorporateCardIssuePhysicalParamsShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardIssueVirtualParams struct {
	HolderName   param.Field[string]      `json:"holderName,required"`
	MonthlyLimit param.Field[float64]     `json:"monthlyLimit,required"`
	Purpose      param.Field[string]      `json:"purpose,required"`
	Metadata     param.Field[interface{}] `json:"metadata"`
}

func (r CorporateCardIssueVirtualParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
