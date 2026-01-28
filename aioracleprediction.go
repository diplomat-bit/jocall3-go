// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"github.com/stainless-sdks/jocall3-go/option"
)

// AIOraclePredictionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOraclePredictionService] method instead.
type AIOraclePredictionService struct {
	Options []option.RequestOption
}

// NewAIOraclePredictionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOraclePredictionService(opts ...option.RequestOption) (r *AIOraclePredictionService) {
	r = &AIOraclePredictionService{}
	r.Options = opts
	return
}
