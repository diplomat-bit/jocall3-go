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
	Options        []option.RequestOption
	Transactions   *AccountTransactionService
	BalanceHistory *AccountBalanceHistoryService
	Statements     *AccountStatementService
	Overdraft      *AccountOverdraftService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Transactions = NewAccountTransactionService(opts...)
	r.BalanceHistory = NewAccountBalanceHistoryService(opts...)
	r.Statements = NewAccountStatementService(opts...)
	r.Overdraft = NewAccountOverdraftService(opts...)
	return
}

// Get Deep Account Analytics
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

// List All Linked & Internal Accounts
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Close Financial Account
func (r *AccountService) Close(ctx context.Context, accountID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Link an External Financial Institution
func (r *AccountService) Link(ctx context.Context, body AccountLinkParams, opts ...option.RequestOption) (res *AccountLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Open a New Quantum Internal Account
func (r *AccountService) Open(ctx context.Context, body AccountOpenParams, opts ...option.RequestOption) (res *AccountOpenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/open"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountGetResponse struct {
	ID                string                              `json:"id,required"`
	Currency          string                              `json:"currency,required"`
	CurrentBalance    float64                             `json:"currentBalance,required"`
	Type              string                              `json:"type,required"`
	AvailableBalance  float64                             `json:"availableBalance"`
	InstitutionName   string                              `json:"institutionName"`
	LastUpdated       time.Time                           `json:"lastUpdated" format:"date-time"`
	Name              string                              `json:"name"`
	ProjectedCashFlow AccountGetResponseProjectedCashFlow `json:"projectedCashFlow"`
	JSON              accountGetResponseJSON              `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	CurrentBalance    apijson.Field
	Type              apijson.Field
	AvailableBalance  apijson.Field
	InstitutionName   apijson.Field
	LastUpdated       apijson.Field
	Name              apijson.Field
	ProjectedCashFlow apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseProjectedCashFlow struct {
	ConfidenceScore int64                                   `json:"confidenceScore"`
	Days30          float64                                 `json:"days30"`
	JSON            accountGetResponseProjectedCashFlowJSON `json:"-"`
}

// accountGetResponseProjectedCashFlowJSON contains the JSON metadata for the
// struct [AccountGetResponseProjectedCashFlow]
type accountGetResponseProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
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
	Data       []AccountListResponseData `json:"data"`
	NextOffset int64                     `json:"nextOffset"`
	Total      int64                     `json:"total"`
	JSON       accountListResponseJSON   `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	NextOffset  apijson.Field
	Total       apijson.Field
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
	ID                string                                   `json:"id,required"`
	Currency          string                                   `json:"currency,required"`
	CurrentBalance    float64                                  `json:"currentBalance,required"`
	Type              string                                   `json:"type,required"`
	AvailableBalance  float64                                  `json:"availableBalance"`
	InstitutionName   string                                   `json:"institutionName"`
	LastUpdated       time.Time                                `json:"lastUpdated" format:"date-time"`
	Name              string                                   `json:"name"`
	ProjectedCashFlow AccountListResponseDataProjectedCashFlow `json:"projectedCashFlow"`
	JSON              accountListResponseDataJSON              `json:"-"`
}

// accountListResponseDataJSON contains the JSON metadata for the struct
// [AccountListResponseData]
type accountListResponseDataJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	CurrentBalance    apijson.Field
	Type              apijson.Field
	AvailableBalance  apijson.Field
	InstitutionName   apijson.Field
	LastUpdated       apijson.Field
	Name              apijson.Field
	ProjectedCashFlow apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseDataProjectedCashFlow struct {
	ConfidenceScore int64                                        `json:"confidenceScore"`
	Days30          float64                                      `json:"days30"`
	JSON            accountListResponseDataProjectedCashFlowJSON `json:"-"`
}

// accountListResponseDataProjectedCashFlowJSON contains the JSON metadata for the
// struct [AccountListResponseDataProjectedCashFlow]
type accountListResponseDataProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountListResponseDataProjectedCashFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseDataProjectedCashFlowJSON) RawJSON() string {
	return r.raw
}

type AccountLinkResponse struct {
	LinkSessionID string                  `json:"linkSessionId"`
	Status        string                  `json:"status"`
	JSON          accountLinkResponseJSON `json:"-"`
}

// accountLinkResponseJSON contains the JSON metadata for the struct
// [AccountLinkResponse]
type accountLinkResponseJSON struct {
	LinkSessionID apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOpenResponse struct {
	ID                string                               `json:"id,required"`
	Currency          string                               `json:"currency,required"`
	CurrentBalance    float64                              `json:"currentBalance,required"`
	Type              string                               `json:"type,required"`
	AvailableBalance  float64                              `json:"availableBalance"`
	InstitutionName   string                               `json:"institutionName"`
	LastUpdated       time.Time                            `json:"lastUpdated" format:"date-time"`
	Name              string                               `json:"name"`
	ProjectedCashFlow AccountOpenResponseProjectedCashFlow `json:"projectedCashFlow"`
	JSON              accountOpenResponseJSON              `json:"-"`
}

// accountOpenResponseJSON contains the JSON metadata for the struct
// [AccountOpenResponse]
type accountOpenResponseJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	CurrentBalance    apijson.Field
	Type              apijson.Field
	AvailableBalance  apijson.Field
	InstitutionName   apijson.Field
	LastUpdated       apijson.Field
	Name              apijson.Field
	ProjectedCashFlow apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountOpenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOpenResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOpenResponseProjectedCashFlow struct {
	ConfidenceScore int64                                    `json:"confidenceScore"`
	Days30          float64                                  `json:"days30"`
	JSON            accountOpenResponseProjectedCashFlowJSON `json:"-"`
}

// accountOpenResponseProjectedCashFlowJSON contains the JSON metadata for the
// struct [AccountOpenResponseProjectedCashFlow]
type accountOpenResponseProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountOpenResponseProjectedCashFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOpenResponseProjectedCashFlowJSON) RawJSON() string {
	return r.raw
}

type AccountListParams struct {
	Limit  param.Field[int64] `query:"limit"`
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
	InstitutionID param.Field[string] `json:"institutionId,required"`
	PublicToken   param.Field[string] `json:"publicToken,required"`
}

func (r AccountLinkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountOpenParams struct {
	Currency       param.Field[string]                       `json:"currency,required"`
	InitialDeposit param.Field[float64]                      `json:"initialDeposit,required"`
	ProductType    param.Field[AccountOpenParamsProductType] `json:"productType,required"`
	// User IDs for joint accounts
	Owners param.Field[[]string] `json:"owners"`
}

func (r AccountOpenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountOpenParamsProductType string

const (
	AccountOpenParamsProductTypeQuantumChecking AccountOpenParamsProductType = "quantum_checking"
	AccountOpenParamsProductTypeEliteSavings    AccountOpenParamsProductType = "elite_savings"
	AccountOpenParamsProductTypeHighYieldVault  AccountOpenParamsProductType = "high_yield_vault"
)

func (r AccountOpenParamsProductType) IsKnown() bool {
	switch r {
	case AccountOpenParamsProductTypeQuantumChecking, AccountOpenParamsProductTypeEliteSavings, AccountOpenParamsProductTypeHighYieldVault:
		return true
	}
	return false
}
