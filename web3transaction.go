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

// Web3TransactionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3TransactionService] method instead.
type Web3TransactionService struct {
	Options []option.RequestOption
}

// NewWeb3TransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3TransactionService(opts ...option.RequestOption) (r *Web3TransactionService) {
	r = &Web3TransactionService{}
	r.Options = opts
	return
}

// Cross-chain Asset Bridge
func (r *Web3TransactionService) BridgeChain(ctx context.Context, body Web3TransactionBridgeChainParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "web3/transactions/bridge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Initiate a Web3 transaction
func (r *Web3TransactionService) Initiate(ctx context.Context, body Web3TransactionInitiateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "web3/transactions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Initiate On-chain Transfer
func (r *Web3TransactionService) Send(ctx context.Context, body Web3TransactionSendParams, opts ...option.RequestOption) (res *Web3TransactionSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/transactions/send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute Multi-chain Token Swap
func (r *Web3TransactionService) SwapTokens(ctx context.Context, body Web3TransactionSwapTokensParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "web3/transactions/swap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type Web3TransactionSendResponse struct {
	TxHash string                          `json:"txHash"`
	JSON   web3TransactionSendResponseJSON `json:"-"`
}

// web3TransactionSendResponseJSON contains the JSON metadata for the struct
// [Web3TransactionSendResponse]
type web3TransactionSendResponseJSON struct {
	TxHash      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3TransactionSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3TransactionSendResponseJSON) RawJSON() string {
	return r.raw
}

type Web3TransactionBridgeChainParams struct {
	Token       param.Field[string] `json:"token,required"`
	Amount      param.Field[string] `json:"amount,required"`
	DestChain   param.Field[string] `json:"destChain,required"`
	SourceChain param.Field[string] `json:"sourceChain,required"`
}

func (r Web3TransactionBridgeChainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3TransactionInitiateParams struct {
	Amount   param.Field[float64] `json:"amount,required"`
	Asset    param.Field[string]  `json:"asset,required"`
	WalletID param.Field[string]  `json:"wallet_id,required"`
}

func (r Web3TransactionInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3TransactionSendParams struct {
	Token  param.Field[string] `json:"token,required"`
	Amount param.Field[string] `json:"amount,required"`
	To     param.Field[string] `json:"to,required"`
}

func (r Web3TransactionSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3TransactionSwapTokensParams struct {
	Amount    param.Field[string] `json:"amount,required"`
	FromToken param.Field[string] `json:"fromToken,required"`
	ToToken   param.Field[string] `json:"toToken,required"`
}

func (r Web3TransactionSwapTokensParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
