// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"

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

// Searches for available investment assets (stocks, ETFs, mutual funds) and
// returns their ESG impact scores.
func (r *InvestmentAssetService) Search(ctx context.Context, query InvestmentAssetSearchParams, opts ...option.RequestOption) (res *InvestmentAssetSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/assets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InvestmentAssetSearchResponse = interface{}

type InvestmentAssetSearchParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Minimum desired ESG score (0-10).
	MinEsgScore param.Field[int64] `query:"minESGScore"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Search query for asset name or symbol.
	Query param.Field[string] `query:"query"`
}

// URLQuery serializes [InvestmentAssetSearchParams]'s query parameters as
// `url.Values`.
func (r InvestmentAssetSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
