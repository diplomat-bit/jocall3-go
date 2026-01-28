// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
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
