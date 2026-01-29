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

// Web3NFTService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3NFTService] method instead.
type Web3NFTService struct {
	Options []option.RequestOption
}

// NewWeb3NFTService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWeb3NFTService(opts ...option.RequestOption) (r *Web3NFTService) {
	r = &Web3NFTService{}
	r.Options = opts
	return
}

// List NFT Collection
func (r *Web3NFTService) List(ctx context.Context, opts ...option.RequestOption) (res *Web3NFTListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/nfts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Mint Utility NFT
func (r *Web3NFTService) Mint(ctx context.Context, body Web3NFTMintParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "web3/nfts/mint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type Web3NFTListResponse struct {
	Data []interface{}           `json:"data"`
	JSON web3NFTListResponseJSON `json:"-"`
}

// web3NFTListResponseJSON contains the JSON metadata for the struct
// [Web3NFTListResponse]
type web3NFTListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3NFTListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3NFTListResponseJSON) RawJSON() string {
	return r.raw
}

type Web3NFTMintParams struct {
	MetadataUri param.Field[string] `json:"metadataUri,required"`
}

func (r Web3NFTMintParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
