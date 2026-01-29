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

// Prepares and initiates a cryptocurrency transfer from a connected wallet to a
// specified recipient address. Requires user confirmation (e.g., via wallet
// signature).
func (r *Web3TransactionService) Initiate(ctx context.Context, body Web3TransactionInitiateParams, opts ...option.RequestOption) (res *Web3TransactionInitiateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/transactions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Web3TransactionInitiateResponse = interface{}

type Web3TransactionInitiateParams struct {
}

func (r Web3TransactionInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
