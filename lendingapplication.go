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

// LendingApplicationService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendingApplicationService] method instead.
type LendingApplicationService struct {
	Options []option.RequestOption
}

// NewLendingApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLendingApplicationService(opts ...option.RequestOption) (r *LendingApplicationService) {
	r = &LendingApplicationService{}
	r.Options = opts
	return
}

// Track Loan Processing
func (r *LendingApplicationService) GetStatus(ctx context.Context, appID string, opts ...option.RequestOption) (res *LendingApplicationGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("lending/applications/%s/status", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LendingApplicationGetStatusResponse struct {
	Status              string                                  `json:"status"`
	UnderwriterQueuePos int64                                   `json:"underwriterQueuePos"`
	JSON                lendingApplicationGetStatusResponseJSON `json:"-"`
}

// lendingApplicationGetStatusResponseJSON contains the JSON metadata for the
// struct [LendingApplicationGetStatusResponse]
type lendingApplicationGetStatusResponseJSON struct {
	Status              apijson.Field
	UnderwriterQueuePos apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LendingApplicationGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lendingApplicationGetStatusResponseJSON) RawJSON() string {
	return r.raw
}
