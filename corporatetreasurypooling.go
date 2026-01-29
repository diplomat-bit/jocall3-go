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

// CorporateTreasuryPoolingService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasuryPoolingService] method instead.
type CorporateTreasuryPoolingService struct {
	Options []option.RequestOption
}

// NewCorporateTreasuryPoolingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateTreasuryPoolingService(opts ...option.RequestOption) (r *CorporateTreasuryPoolingService) {
	r = &CorporateTreasuryPoolingService{}
	r.Options = opts
	return
}

// Configure liquidity pooling
func (r *CorporateTreasuryPoolingService) Configure(ctx context.Context, body CorporateTreasuryPoolingConfigureParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/treasury/liquidity/pooling"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CorporateTreasuryPoolingConfigureParams struct {
	SourceAccountIDs param.Field[[]string] `json:"source_account_ids"`
	TargetAccountID  param.Field[string]   `json:"target_account_id"`
}

func (r CorporateTreasuryPoolingConfigureParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
