// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// PaymentInternationalService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentInternationalService] method instead.
type PaymentInternationalService struct {
	Options []option.RequestOption
}

// NewPaymentInternationalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentInternationalService(opts ...option.RequestOption) (r *PaymentInternationalService) {
	r = &PaymentInternationalService{}
	r.Options = opts
	return
}

// Retrieves the current processing status and details of an initiated
// international payment.
func (r *PaymentInternationalService) GetStatus(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *PaymentInternationalGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return
	}
	path := fmt.Sprintf("payments/international/%s/status", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaymentInternationalGetStatusResponse = interface{}
