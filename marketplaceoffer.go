// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// MarketplaceOfferService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketplaceOfferService] method instead.
type MarketplaceOfferService struct {
	Options []option.RequestOption
}

// NewMarketplaceOfferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketplaceOfferService(opts ...option.RequestOption) (r *MarketplaceOfferService) {
	r = &MarketplaceOfferService{}
	r.Options = opts
	return
}

// List AI-Targeted Loyalty Offers
func (r *MarketplaceOfferService) List(ctx context.Context, opts ...option.RequestOption) (res *MarketplaceOfferListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketplace/offers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Redeem Marketplace Reward
func (r *MarketplaceOfferService) Redeem(ctx context.Context, offerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if offerID == "" {
		err = errors.New("missing required offerId parameter")
		return
	}
	path := fmt.Sprintf("marketplace/offers/%s/redeem", offerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type MarketplaceOfferListResponse struct {
	Data []interface{}                    `json:"data"`
	JSON marketplaceOfferListResponseJSON `json:"-"`
}

// marketplaceOfferListResponseJSON contains the JSON metadata for the struct
// [MarketplaceOfferListResponse]
type marketplaceOfferListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketplaceOfferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketplaceOfferListResponseJSON) RawJSON() string {
	return r.raw
}
