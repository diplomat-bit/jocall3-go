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

// Retrieves the current overdraft protection settings for a specific account.
func (r *AccountOverdraftService) GetSettings(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountOverdraftGetSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the overdraft protection settings for a specific account, enabling or
// disabling protection and configuring preferences.
func (r *AccountOverdraftService) UpdateSettings(ctx context.Context, accountID string, body AccountOverdraftUpdateSettingsParams, opts ...option.RequestOption) (res *AccountOverdraftUpdateSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountOverdraftGetSettingsResponse struct {
	AccountID              string                                  `json:"accountId,required"`
	Enabled                bool                                    `json:"enabled,required"`
	FeePreference          string                                  `json:"feePreference,required"`
	LinkedSavingsAccountID string                                  `json:"linkedSavingsAccountId"`
	LinkToSavings          bool                                    `json:"linkToSavings"`
	ProtectionLimit        float64                                 `json:"protectionLimit"`
	JSON                   accountOverdraftGetSettingsResponseJSON `json:"-"`
}

// accountOverdraftGetSettingsResponseJSON contains the JSON metadata for the
// struct [AccountOverdraftGetSettingsResponse]
type accountOverdraftGetSettingsResponseJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountOverdraftGetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOverdraftGetSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOverdraftUpdateSettingsResponse struct {
	AccountID              string                                     `json:"accountId,required"`
	Enabled                bool                                       `json:"enabled,required"`
	FeePreference          string                                     `json:"feePreference,required"`
	LinkedSavingsAccountID string                                     `json:"linkedSavingsAccountId"`
	LinkToSavings          bool                                       `json:"linkToSavings"`
	ProtectionLimit        float64                                    `json:"protectionLimit"`
	JSON                   accountOverdraftUpdateSettingsResponseJSON `json:"-"`
}

// accountOverdraftUpdateSettingsResponseJSON contains the JSON metadata for the
// struct [AccountOverdraftUpdateSettingsResponse]
type accountOverdraftUpdateSettingsResponseJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountOverdraftUpdateSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOverdraftUpdateSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOverdraftUpdateSettingsParams struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	FeePreference param.Field[string] `json:"feePreference"`
	LinkToSavings param.Field[bool]   `json:"linkToSavings"`
}

func (r AccountOverdraftUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
