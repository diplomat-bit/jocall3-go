// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// MarketplaceService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketplaceService] method instead.
type MarketplaceService struct {
	Options []option.RequestOption
	Offers  *MarketplaceOfferService
}

// NewMarketplaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketplaceService(opts ...option.RequestOption) (r *MarketplaceService) {
	r = &MarketplaceService{}
	r.Options = opts
	r.Offers = NewMarketplaceOfferService(opts...)
	return
}

// List Financial Products & Add-ons
func (r *MarketplaceService) ListProducts(ctx context.Context, opts ...option.RequestOption) (res *MarketplaceListProductsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketplace/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MarketplaceListProductsResponse struct {
	Data []interface{}                       `json:"data"`
	JSON marketplaceListProductsResponseJSON `json:"-"`
}

// marketplaceListProductsResponseJSON contains the JSON metadata for the struct
// [MarketplaceListProductsResponse]
type marketplaceListProductsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketplaceListProductsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceListProductsResponseJSON) RawJSON() string {
	return r.raw
}
