// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIIncubatorService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorService] method instead.
type AIIncubatorService struct {
	Options []option.RequestOption
	Pitch   *AIIncubatorPitchService
}

// NewAIIncubatorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIIncubatorService(opts ...option.RequestOption) (r *AIIncubatorService) {
	r = &AIIncubatorService{}
	r.Options = opts
	r.Pitch = NewAIIncubatorPitchService(opts...)
	return
}
