// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// InvestmentAssetService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentAssetService] method instead.
type InvestmentAssetService struct {
	Options []option.RequestOption
}

// NewInvestmentAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestmentAssetService(opts ...option.RequestOption) (r *InvestmentAssetService) {
	r = &InvestmentAssetService{}
	r.Options = opts
	return
}

// Global Multi-Asset Search (Equities, Crypto, ESG)
func (r *InvestmentAssetService) Search(ctx context.Context, query InvestmentAssetSearchParams, opts ...option.RequestOption) (res *InvestmentAssetSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/assets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InvestmentAssetSearchResponse struct {
	Hits []interface{}                     `json:"hits"`
	JSON investmentAssetSearchResponseJSON `json:"-"`
}

// investmentAssetSearchResponseJSON contains the JSON metadata for the struct
// [InvestmentAssetSearchResponse]
type investmentAssetSearchResponseJSON struct {
	Hits        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentAssetSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentAssetSearchResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentAssetSearchParams struct {
	Query     param.Field[string]                               `query:"query,required"`
	AssetType param.Field[InvestmentAssetSearchParamsAssetType] `query:"assetType"`
}

// URLQuery serializes [InvestmentAssetSearchParams]'s query parameters as
// `url.Values`.
func (r InvestmentAssetSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentAssetSearchParamsAssetType string

const (
	InvestmentAssetSearchParamsAssetTypeEquity InvestmentAssetSearchParamsAssetType = "EQUITY"
	InvestmentAssetSearchParamsAssetTypeCrypto InvestmentAssetSearchParamsAssetType = "CRYPTO"
	InvestmentAssetSearchParamsAssetTypeEtf    InvestmentAssetSearchParamsAssetType = "ETF"
	InvestmentAssetSearchParamsAssetTypeBond   InvestmentAssetSearchParamsAssetType = "BOND"
)

func (r InvestmentAssetSearchParamsAssetType) IsKnown() bool {
	switch r {
	case InvestmentAssetSearchParamsAssetTypeEquity, InvestmentAssetSearchParamsAssetTypeCrypto, InvestmentAssetSearchParamsAssetTypeEtf, InvestmentAssetSearchParamsAssetTypeBond:
		return true
	}
	return false
}
