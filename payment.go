// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// PaymentService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options       []option.RequestOption
	Domestic      *PaymentDomesticService
	International *PaymentInternationalService
	Fx            *PaymentFxService
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	r.Domestic = NewPaymentDomesticService(opts...)
	r.International = NewPaymentInternationalService(opts...)
	r.Fx = NewPaymentFxService(opts...)
	return
}

// Get Payment Receipt
func (r *PaymentService) Get(ctx context.Context, paymentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return
	}
	path := fmt.Sprintf("payments/%s", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// List Payment Activity
func (r *PaymentService) List(ctx context.Context, opts ...option.RequestOption) (res *PaymentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaymentListResponse struct {
	Data []interface{}           `json:"data"`
	JSON paymentListResponseJSON `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseJSON) RawJSON() string {
	return r.raw
}
