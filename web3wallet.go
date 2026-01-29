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

// Web3WalletService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3WalletService] method instead.
type Web3WalletService struct {
	Options []option.RequestOption
}

// NewWeb3WalletService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3WalletService(opts ...option.RequestOption) (r *Web3WalletService) {
	r = &Web3WalletService{}
	r.Options = opts
	return
}

// Create Non-Custodial Wallet
func (r *Web3WalletService) New(ctx context.Context, body Web3WalletNewParams, opts ...option.RequestOption) (res *Web3WalletNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Connected Wallets
func (r *Web3WalletService) List(ctx context.Context, opts ...option.RequestOption) (res *Web3WalletListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Link External Web3 Wallet (MetaMask/Phantom)
func (r *Web3WalletService) Connect(ctx context.Context, body Web3WalletConnectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "web3/wallets/connect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get Multi-chain Token Balances
func (r *Web3WalletService) GetBalance(ctx context.Context, walletID string, opts ...option.RequestOption) (res *Web3WalletGetBalanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required walletId parameter")
		return
	}
	path := fmt.Sprintf("web3/wallets/%s/balances", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Web3WalletNewResponse struct {
	ID                string                    `json:"id,required"`
	BlockchainNetwork string                    `json:"blockchainNetwork,required"`
	Status            string                    `json:"status,required"`
	WalletAddress     string                    `json:"walletAddress,required"`
	LastSynced        time.Time                 `json:"lastSynced" format:"date-time"`
	WalletProvider    string                    `json:"walletProvider"`
	JSON              web3WalletNewResponseJSON `json:"-"`
}

// web3WalletNewResponseJSON contains the JSON metadata for the struct
// [Web3WalletNewResponse]
type web3WalletNewResponseJSON struct {
	ID                apijson.Field
	BlockchainNetwork apijson.Field
	Status            apijson.Field
	WalletAddress     apijson.Field
	LastSynced        apijson.Field
	WalletProvider    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Web3WalletNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletNewResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletListResponse struct {
	Data []Web3WalletListResponseData `json:"data"`
	JSON web3WalletListResponseJSON   `json:"-"`
}

// web3WalletListResponseJSON contains the JSON metadata for the struct
// [Web3WalletListResponse]
type web3WalletListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletListResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletListResponseData struct {
	ID                string                         `json:"id,required"`
	BlockchainNetwork string                         `json:"blockchainNetwork,required"`
	Status            string                         `json:"status,required"`
	WalletAddress     string                         `json:"walletAddress,required"`
	LastSynced        time.Time                      `json:"lastSynced" format:"date-time"`
	WalletProvider    string                         `json:"walletProvider"`
	JSON              web3WalletListResponseDataJSON `json:"-"`
}

// web3WalletListResponseDataJSON contains the JSON metadata for the struct
// [Web3WalletListResponseData]
type web3WalletListResponseDataJSON struct {
	ID                apijson.Field
	BlockchainNetwork apijson.Field
	Status            apijson.Field
	WalletAddress     apijson.Field
	LastSynced        apijson.Field
	WalletProvider    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Web3WalletListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletListResponseDataJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalanceResponse struct {
	Balances []Web3WalletGetBalanceResponseBalance `json:"balances"`
	JSON     web3WalletGetBalanceResponseJSON      `json:"-"`
}

// web3WalletGetBalanceResponseJSON contains the JSON metadata for the struct
// [Web3WalletGetBalanceResponse]
type web3WalletGetBalanceResponseJSON struct {
	Balances    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletGetBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalanceResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalanceResponseBalance struct {
	Amount string                                  `json:"amount"`
	Symbol string                                  `json:"symbol"`
	JSON   web3WalletGetBalanceResponseBalanceJSON `json:"-"`
}

// web3WalletGetBalanceResponseBalanceJSON contains the JSON metadata for the
// struct [Web3WalletGetBalanceResponseBalance]
type web3WalletGetBalanceResponseBalanceJSON struct {
	Amount      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletGetBalanceResponseBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalanceResponseBalanceJSON) RawJSON() string {
	return r.raw
}

type Web3WalletNewParams struct {
	Network param.Field[string] `json:"network,required"`
}

func (r Web3WalletNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3WalletConnectParams struct {
	Address   param.Field[string] `json:"address,required"`
	Provider  param.Field[string] `json:"provider,required"`
	Signature param.Field[string] `json:"signature,required"`
}

func (r Web3WalletConnectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
