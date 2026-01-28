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

// Updates the overdraft protection settings for a specific account, enabling or
// disabling protection and configuring preferences.
func (r *AccountOverdraftService) Update(ctx context.Context, accountID string, body AccountOverdraftUpdateParams, opts ...option.RequestOption) (res *AccountOverdraftUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves the current overdraft protection settings for a specific account.
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

type AccountOverdraftUpdateResponse struct {
	AccountID              string                             `json:"accountId,required"`
	Enabled                bool                               `json:"enabled,required"`
	FeePreference          string                             `json:"feePreference,required"`
	LinkedSavingsAccountID string                             `json:"linkedSavingsAccountId"`
	LinkToSavings          bool                               `json:"linkToSavings"`
	ProtectionLimit        float64                            `json:"protectionLimit"`
	JSON                   accountOverdraftUpdateResponseJSON `json:"-"`
}

// accountOverdraftUpdateResponseJSON contains the JSON metadata for the struct
// [AccountOverdraftUpdateResponse]
type accountOverdraftUpdateResponseJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountOverdraftUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOverdraftUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOverdraftGetResponse struct {
	AccountID              string                          `json:"accountId,required"`
	Enabled                bool                            `json:"enabled,required"`
	FeePreference          string                          `json:"feePreference,required"`
	LinkedSavingsAccountID string                          `json:"linkedSavingsAccountId"`
	LinkToSavings          bool                            `json:"linkToSavings"`
	ProtectionLimit        float64                         `json:"protectionLimit"`
	JSON                   accountOverdraftGetResponseJSON `json:"-"`
}

// accountOverdraftGetResponseJSON contains the JSON metadata for the struct
// [AccountOverdraftGetResponse]
type accountOverdraftGetResponseJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountOverdraftGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOverdraftGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOverdraftUpdateParams struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	FeePreference param.Field[string] `json:"feePreference"`
	LinkToSavings param.Field[bool]   `json:"linkToSavings"`
}

func (r AccountOverdraftUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
