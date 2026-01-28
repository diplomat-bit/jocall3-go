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

// AccountService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	Transactions *AccountTransactionService
	Statements   *AccountStatementService
	Overdraft    *AccountOverdraftService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Transactions = NewAccountTransactionService(opts...)
	r.Statements = NewAccountStatementService(opts...)
	r.Overdraft = NewAccountOverdraftService(opts...)
	return
}

// Retrieves comprehensive analytics for a specific financial account, including
// historical balance trends, projected cash flow, and AI-driven insights into
// spending patterns.
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/details", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a comprehensive, real-time list of all external financial accounts
// linked to the user's profile, including consolidated balances and institutional
// details.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Begins the secure process of linking a new external financial institution (e.g.,
// another bank, investment platform) to the user's profile, typically involving a
// third-party tokenized flow.
func (r *AccountService) Link(ctx context.Context, body AccountLinkParams, opts ...option.RequestOption) (res *AccountLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountGetResponse struct {
	ID                string                              `json:"id,required"`
	Currency          string                              `json:"currency,required"`
	CurrentBalance    float64                             `json:"currentBalance,required"`
	InstitutionName   string                              `json:"institutionName,required"`
	LastUpdated       time.Time                           `json:"lastUpdated,required" format:"date-time"`
	Name              string                              `json:"name,required"`
	Type              string                              `json:"type,required"`
	AccountHolder     string                              `json:"accountHolder"`
	AvailableBalance  float64                             `json:"availableBalance"`
	BalanceHistory    []AccountGetResponseBalanceHistory  `json:"balanceHistory"`
	ExternalID        string                              `json:"externalId"`
	InterestRate      float64                             `json:"interestRate"`
	Mask              string                              `json:"mask"`
	OpenedDate        time.Time                           `json:"openedDate" format:"date"`
	ProjectedCashFlow AccountGetResponseProjectedCashFlow `json:"projectedCashFlow"`
	Subtype           string                              `json:"subtype"`
	TransactionsCount int64                               `json:"transactionsCount"`
	JSON              accountGetResponseJSON              `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	CurrentBalance    apijson.Field
	InstitutionName   apijson.Field
	LastUpdated       apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	AccountHolder     apijson.Field
	AvailableBalance  apijson.Field
	BalanceHistory    apijson.Field
	ExternalID        apijson.Field
	InterestRate      apijson.Field
	Mask              apijson.Field
	OpenedDate        apijson.Field
	ProjectedCashFlow apijson.Field
	Subtype           apijson.Field
	TransactionsCount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseBalanceHistory struct {
	Balance float64                              `json:"balance"`
	Date    time.Time                            `json:"date" format:"date"`
	JSON    accountGetResponseBalanceHistoryJSON `json:"-"`
}

// accountGetResponseBalanceHistoryJSON contains the JSON metadata for the struct
// [AccountGetResponseBalanceHistory]
type accountGetResponseBalanceHistoryJSON struct {
	Balance     apijson.Field
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseBalanceHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseBalanceHistoryJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseProjectedCashFlow struct {
	ConfidenceScore int64                                   `json:"confidenceScore"`
	Days30          float64                                 `json:"days30"`
	Days90          float64                                 `json:"days90"`
	JSON            accountGetResponseProjectedCashFlowJSON `json:"-"`
}

// accountGetResponseProjectedCashFlowJSON contains the JSON metadata for the
// struct [AccountGetResponseProjectedCashFlow]
type accountGetResponseProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
	Days90          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountGetResponseProjectedCashFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseProjectedCashFlowJSON) RawJSON() string {
	return r.raw
}

type AccountListResponse struct {
	Data       []AccountListResponseData `json:"data,required"`
	Limit      int64                     `json:"limit,required"`
	Offset     int64                     `json:"offset,required"`
	Total      int64                     `json:"total,required"`
	NextOffset int64                     `json:"nextOffset"`
	JSON       accountListResponseJSON   `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseData struct {
	ID               string                      `json:"id"`
	AvailableBalance float64                     `json:"availableBalance"`
	Currency         string                      `json:"currency"`
	CurrentBalance   float64                     `json:"currentBalance"`
	ExternalID       string                      `json:"externalId"`
	InstitutionName  string                      `json:"institutionName"`
	LastUpdated      time.Time                   `json:"lastUpdated" format:"date-time"`
	Mask             string                      `json:"mask"`
	Name             string                      `json:"name"`
	Subtype          string                      `json:"subtype"`
	Type             string                      `json:"type"`
	JSON             accountListResponseDataJSON `json:"-"`
}

// accountListResponseDataJSON contains the JSON metadata for the struct
// [AccountListResponseData]
type accountListResponseDataJSON struct {
	ID               apijson.Field
	AvailableBalance apijson.Field
	Currency         apijson.Field
	CurrentBalance   apijson.Field
	ExternalID       apijson.Field
	InstitutionName  apijson.Field
	LastUpdated      apijson.Field
	Mask             apijson.Field
	Name             apijson.Field
	Subtype          apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountLinkResponse struct {
	AuthUri       string                  `json:"authUri,required"`
	LinkSessionID string                  `json:"linkSessionId,required"`
	Status        string                  `json:"status,required"`
	Message       string                  `json:"message"`
	JSON          accountLinkResponseJSON `json:"-"`
}

// accountLinkResponseJSON contains the JSON metadata for the struct
// [AccountLinkResponse]
type accountLinkResponseJSON struct {
	AuthUri       apijson.Field
	LinkSessionID apijson.Field
	Status        apijson.Field
	Message       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLinkParams struct {
	CountryCode     param.Field[string] `json:"countryCode,required"`
	InstitutionName param.Field[string] `json:"institutionName,required"`
}

func (r AccountLinkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
