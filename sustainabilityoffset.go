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

// SustainabilityOffsetService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityOffsetService] method instead.
type SustainabilityOffsetService struct {
	Options []option.RequestOption
}

// NewSustainabilityOffsetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSustainabilityOffsetService(opts ...option.RequestOption) (r *SustainabilityOffsetService) {
	r = &SustainabilityOffsetService{}
	r.Options = opts
	return
}

// Purchase Verified Carbon Credits
func (r *SustainabilityOffsetService) Purchase(ctx context.Context, body SustainabilityOffsetPurchaseParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "sustainability/offsets/purchase"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retire Carbon Credits (Permanent Offsetting)
func (r *SustainabilityOffsetService) Retire(ctx context.Context, body SustainabilityOffsetRetireParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "sustainability/offsets/retire"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SustainabilityOffsetPurchaseParams struct {
	ProjectID       param.Field[string]  `json:"projectId,required"`
	Tonnes          param.Field[float64] `json:"tonnes,required"`
	PaymentSourceID param.Field[string]  `json:"paymentSourceId"`
}

func (r SustainabilityOffsetPurchaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SustainabilityOffsetRetireParams struct {
	CertificateID param.Field[string] `json:"certificateId,required"`
}

func (r SustainabilityOffsetRetireParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
