// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
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
