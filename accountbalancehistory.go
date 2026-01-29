// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AccountBalanceHistoryService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBalanceHistoryService] method instead.
type AccountBalanceHistoryService struct {
	Options []option.RequestOption
}

// NewAccountBalanceHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBalanceHistoryService(opts ...option.RequestOption) (r *AccountBalanceHistoryService) {
	r = &AccountBalanceHistoryService{}
	r.Options = opts
	return
}

// Get Historical Balance Snapshots
func (r *AccountBalanceHistoryService) Get(ctx context.Context, accountID string, query AccountBalanceHistoryGetParams, opts ...option.RequestOption) (res *AccountBalanceHistoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/balance-history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountBalanceHistoryGetResponse struct {
	History []AccountBalanceHistoryGetResponseHistory `json:"history"`
	JSON    accountBalanceHistoryGetResponseJSON      `json:"-"`
}

// accountBalanceHistoryGetResponseJSON contains the JSON metadata for the struct
// [AccountBalanceHistoryGetResponse]
type accountBalanceHistoryGetResponseJSON struct {
	History     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBalanceHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBalanceHistoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBalanceHistoryGetResponseHistory struct {
	Balance   float64                                     `json:"balance"`
	Timestamp time.Time                                   `json:"timestamp" format:"date-time"`
	JSON      accountBalanceHistoryGetResponseHistoryJSON `json:"-"`
}

// accountBalanceHistoryGetResponseHistoryJSON contains the JSON metadata for the
// struct [AccountBalanceHistoryGetResponseHistory]
type accountBalanceHistoryGetResponseHistoryJSON struct {
	Balance     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBalanceHistoryGetResponseHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBalanceHistoryGetResponseHistoryJSON) RawJSON() string {
	return r.raw
}

type AccountBalanceHistoryGetParams struct {
	Period param.Field[AccountBalanceHistoryGetParamsPeriod] `query:"period"`
}

// URLQuery serializes [AccountBalanceHistoryGetParams]'s query parameters as
// `url.Values`.
func (r AccountBalanceHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBalanceHistoryGetParamsPeriod string

const (
	AccountBalanceHistoryGetParamsPeriod1d  AccountBalanceHistoryGetParamsPeriod = "1d"
	AccountBalanceHistoryGetParamsPeriod7d  AccountBalanceHistoryGetParamsPeriod = "7d"
	AccountBalanceHistoryGetParamsPeriod30d AccountBalanceHistoryGetParamsPeriod = "30d"
	AccountBalanceHistoryGetParamsPeriod1y  AccountBalanceHistoryGetParamsPeriod = "1y"
	AccountBalanceHistoryGetParamsPeriodAll AccountBalanceHistoryGetParamsPeriod = "all"
)

func (r AccountBalanceHistoryGetParamsPeriod) IsKnown() bool {
	switch r {
	case AccountBalanceHistoryGetParamsPeriod1d, AccountBalanceHistoryGetParamsPeriod7d, AccountBalanceHistoryGetParamsPeriod30d, AccountBalanceHistoryGetParamsPeriod1y, AccountBalanceHistoryGetParamsPeriodAll:
		return true
	}
	return false
}
