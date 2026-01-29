// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AccountOverdraftService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOverdraftService] method instead.
type AccountOverdraftService struct {
	Options []option.RequestOption
}

// NewAccountOverdraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountOverdraftService(opts ...option.RequestOption) (r *AccountOverdraftService) {
	r = &AccountOverdraftService{}
	r.Options = opts
	return
}

// Update Overdraft Settings
func (r *AccountOverdraftService) Update(ctx context.Context, accountID string, body AccountOverdraftUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get Overdraft Settings
func (r *AccountOverdraftService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountOverdraftGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountOverdraftGetResponse struct {
	Enabled       bool                            `json:"enabled"`
	FeePreference string                          `json:"feePreference"`
	Limit         float64                         `json:"limit"`
	JSON          accountOverdraftGetResponseJSON `json:"-"`
}

// accountOverdraftGetResponseJSON contains the JSON metadata for the struct
// [AccountOverdraftGetResponse]
type accountOverdraftGetResponseJSON struct {
	Enabled       apijson.Field
	FeePreference apijson.Field
	Limit         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountOverdraftGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOverdraftGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOverdraftUpdateParams struct {
	Enabled param.Field[bool]    `json:"enabled"`
	Limit   param.Field[float64] `json:"limit"`
}

func (r AccountOverdraftUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
