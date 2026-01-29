// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"github.com/diplomat-bit/jocall3-go/option"
)

// SystemStatusService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemStatusService] method instead.
type SystemStatusService struct {
	Options []option.RequestOption
}

// NewSystemStatusService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemStatusService(opts ...option.RequestOption) (r *SystemStatusService) {
	r = &SystemStatusService{}
	r.Options = opts
	return
}
