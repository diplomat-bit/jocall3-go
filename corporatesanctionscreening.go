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

// CorporateSanctionScreeningService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateSanctionScreeningService] method instead.
type CorporateSanctionScreeningService struct {
	Options []option.RequestOption
}

// NewCorporateSanctionScreeningService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateSanctionScreeningService(opts ...option.RequestOption) (r *CorporateSanctionScreeningService) {
	r = &CorporateSanctionScreeningService{}
	r.Options = opts
	return
}

// Executes a real-time screening of an individual or entity against global
// sanction lists and watchlists.
func (r *CorporateSanctionScreeningService) Screen(ctx context.Context, body CorporateSanctionScreeningScreenParams, opts ...option.RequestOption) (res *CorporateSanctionScreeningScreenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/sanction-screening"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateSanctionScreeningScreenResponse = interface{}

type CorporateSanctionScreeningScreenParams struct {
	Address param.Field[interface{}] `json:"address"`
}

func (r CorporateSanctionScreeningScreenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
