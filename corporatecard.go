// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
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

type CorporateCardIssuePhysicalResponse struct {
	ID             string                                     `json:"id,required"`
	CardNumberMask string                                     `json:"cardNumberMask,required"`
	HolderName     string                                     `json:"holderName,required"`
	Status         string                                     `json:"status,required"`
	Controls       CorporateCardIssuePhysicalResponseControls `json:"controls"`
	ExpirationDate time.Time                                  `json:"expirationDate" format:"date"`
	Frozen         bool                                       `json:"frozen"`
	JSON           corporateCardIssuePhysicalResponseJSON     `json:"-"`
}

// corporateCardIssuePhysicalResponseJSON contains the JSON metadata for the struct
// [CorporateCardIssuePhysicalResponse]
type corporateCardIssuePhysicalResponseJSON struct {
	ID             apijson.Field
	CardNumberMask apijson.Field
	HolderName     apijson.Field
	Status         apijson.Field
	Controls       apijson.Field
	ExpirationDate apijson.Field
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

type CorporateCardIssuePhysicalResponseControls struct {
	Categories   []string                                       `json:"categories"`
	MonthlyLimit float64                                        `json:"monthlyLimit"`
	JSON         corporateCardIssuePhysicalResponseControlsJSON `json:"-"`
}

// corporateCardIssuePhysicalResponseControlsJSON contains the JSON metadata for
// the struct [CorporateCardIssuePhysicalResponseControls]
type corporateCardIssuePhysicalResponseControlsJSON struct {
	Categories   apijson.Field
	MonthlyLimit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CorporateCardIssuePhysicalResponseControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardIssuePhysicalResponseControlsJSON) RawJSON() string {
	return r.raw
}

type CorporateCardIssueVirtualResponse struct {
	ID             string                                    `json:"id,required"`
	CardNumberMask string                                    `json:"cardNumberMask,required"`
	HolderName     string                                    `json:"holderName,required"`
	Status         string                                    `json:"status,required"`
	Controls       CorporateCardIssueVirtualResponseControls `json:"controls"`
	ExpirationDate time.Time                                 `json:"expirationDate" format:"date"`
	Frozen         bool                                      `json:"frozen"`
	JSON           corporateCardIssueVirtualResponseJSON     `json:"-"`
}

// corporateCardIssueVirtualResponseJSON contains the JSON metadata for the struct
// [CorporateCardIssueVirtualResponse]
type corporateCardIssueVirtualResponseJSON struct {
	ID             apijson.Field
	CardNumberMask apijson.Field
	HolderName     apijson.Field
	Status         apijson.Field
	Controls       apijson.Field
	ExpirationDate apijson.Field
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

type CorporateCardIssueVirtualResponseControls struct {
	Categories   []string                                      `json:"categories"`
	MonthlyLimit float64                                       `json:"monthlyLimit"`
	JSON         corporateCardIssueVirtualResponseControlsJSON `json:"-"`
}

// corporateCardIssueVirtualResponseControlsJSON contains the JSON metadata for the
// struct [CorporateCardIssueVirtualResponseControls]
type corporateCardIssueVirtualResponseControlsJSON struct {
	Categories   apijson.Field
	MonthlyLimit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CorporateCardIssueVirtualResponseControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardIssueVirtualResponseControlsJSON) RawJSON() string {
	return r.raw
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
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
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
