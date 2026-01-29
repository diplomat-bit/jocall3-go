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

// Web3Service contains methods and other services that help with interacting with
// the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3Service] method instead.
type Web3Service struct {
	Options        []option.RequestOption
	Wallets        *Web3WalletService
	Transactions   *Web3TransactionService
	NFTs           *Web3NFTService
	SmartContracts *Web3SmartContractService
}

// NewWeb3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWeb3Service(opts ...option.RequestOption) (r *Web3Service) {
	r = &Web3Service{}
	r.Options = opts
	r.Wallets = NewWeb3WalletService(opts...)
	r.Transactions = NewWeb3TransactionService(opts...)
	r.NFTs = NewWeb3NFTService(opts...)
	r.SmartContracts = NewWeb3SmartContractService(opts...)
	return
}

// Get Blockchain Network Health
func (r *Web3Service) GetNetworkStatus(ctx context.Context, opts ...option.RequestOption) (res *Web3GetNetworkStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/network/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Web3GetNetworkStatusResponse struct {
	Ethereum interface{}                      `json:"ethereum"`
	Polygon  interface{}                      `json:"polygon"`
	Solana   interface{}                      `json:"solana"`
	JSON     web3GetNetworkStatusResponseJSON `json:"-"`
}

// web3GetNetworkStatusResponseJSON contains the JSON metadata for the struct
// [Web3GetNetworkStatusResponse]
type web3GetNetworkStatusResponseJSON struct {
	Ethereum    apijson.Field
	Polygon     apijson.Field
	Solana      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3GetNetworkStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3GetNetworkStatusResponseJSON) RawJSON() string {
	return r.raw
}
