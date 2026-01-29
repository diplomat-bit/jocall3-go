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

// AccountTransactionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTransactionService] method instead.
type AccountTransactionService struct {
	Options []option.RequestOption
}

// NewAccountTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTransactionService(opts ...option.RequestOption) (r *AccountTransactionService) {
	r = &AccountTransactionService{}
	r.Options = opts
	return
}

// Get Historical Ledger Archive
func (r *AccountTransactionService) ListArchived(ctx context.Context, accountID string, query AccountTransactionListArchivedParams, opts ...option.RequestOption) (res *AccountTransactionListArchivedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/transactions/archived", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Pending Ledger Entries
func (r *AccountTransactionService) ListPending(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountTransactionListPendingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/transactions/pending", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountTransactionListArchivedResponse struct {
	Data       []AccountTransactionListArchivedResponseData `json:"data"`
	NextOffset int64                                        `json:"nextOffset"`
	Total      int64                                        `json:"total"`
	JSON       accountTransactionListArchivedResponseJSON   `json:"-"`
}

// accountTransactionListArchivedResponseJSON contains the JSON metadata for the
// struct [AccountTransactionListArchivedResponse]
type accountTransactionListArchivedResponseJSON struct {
	Data        apijson.Field
	NextOffset  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransactionListArchivedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListArchivedResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListArchivedResponseData struct {
	ID              string                                                    `json:"id,required"`
	Amount          float64                                                   `json:"amount,required"`
	Currency        string                                                    `json:"currency,required"`
	Date            time.Time                                                 `json:"date,required" format:"date"`
	Description     string                                                    `json:"description,required"`
	AccountID       string                                                    `json:"accountId"`
	CarbonFootprint float64                                                   `json:"carbonFootprint"`
	Category        string                                                    `json:"category"`
	MerchantDetails AccountTransactionListArchivedResponseDataMerchantDetails `json:"merchantDetails"`
	JSON            accountTransactionListArchivedResponseDataJSON            `json:"-"`
}

// accountTransactionListArchivedResponseDataJSON contains the JSON metadata for
// the struct [AccountTransactionListArchivedResponseData]
type accountTransactionListArchivedResponseDataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	Currency        apijson.Field
	Date            apijson.Field
	Description     apijson.Field
	AccountID       apijson.Field
	CarbonFootprint apijson.Field
	Category        apijson.Field
	MerchantDetails apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountTransactionListArchivedResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListArchivedResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListArchivedResponseDataMerchantDetails struct {
	LogoURL string                                                        `json:"logoUrl"`
	Name    string                                                        `json:"name"`
	JSON    accountTransactionListArchivedResponseDataMerchantDetailsJSON `json:"-"`
}

// accountTransactionListArchivedResponseDataMerchantDetailsJSON contains the JSON
// metadata for the struct
// [AccountTransactionListArchivedResponseDataMerchantDetails]
type accountTransactionListArchivedResponseDataMerchantDetailsJSON struct {
	LogoURL     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransactionListArchivedResponseDataMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListArchivedResponseDataMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListPendingResponse struct {
	Data       []AccountTransactionListPendingResponseData `json:"data"`
	NextOffset int64                                       `json:"nextOffset"`
	Total      int64                                       `json:"total"`
	JSON       accountTransactionListPendingResponseJSON   `json:"-"`
}

// accountTransactionListPendingResponseJSON contains the JSON metadata for the
// struct [AccountTransactionListPendingResponse]
type accountTransactionListPendingResponseJSON struct {
	Data        apijson.Field
	NextOffset  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransactionListPendingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListPendingResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListPendingResponseData struct {
	ID              string                                                   `json:"id,required"`
	Amount          float64                                                  `json:"amount,required"`
	Currency        string                                                   `json:"currency,required"`
	Date            time.Time                                                `json:"date,required" format:"date"`
	Description     string                                                   `json:"description,required"`
	AccountID       string                                                   `json:"accountId"`
	CarbonFootprint float64                                                  `json:"carbonFootprint"`
	Category        string                                                   `json:"category"`
	MerchantDetails AccountTransactionListPendingResponseDataMerchantDetails `json:"merchantDetails"`
	JSON            accountTransactionListPendingResponseDataJSON            `json:"-"`
}

// accountTransactionListPendingResponseDataJSON contains the JSON metadata for the
// struct [AccountTransactionListPendingResponseData]
type accountTransactionListPendingResponseDataJSON struct {
	ID              apijson.Field
	Amount          apijson.Field
	Currency        apijson.Field
	Date            apijson.Field
	Description     apijson.Field
	AccountID       apijson.Field
	CarbonFootprint apijson.Field
	Category        apijson.Field
	MerchantDetails apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountTransactionListPendingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListPendingResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListPendingResponseDataMerchantDetails struct {
	LogoURL string                                                       `json:"logoUrl"`
	Name    string                                                       `json:"name"`
	JSON    accountTransactionListPendingResponseDataMerchantDetailsJSON `json:"-"`
}

// accountTransactionListPendingResponseDataMerchantDetailsJSON contains the JSON
// metadata for the struct
// [AccountTransactionListPendingResponseDataMerchantDetails]
type accountTransactionListPendingResponseDataMerchantDetailsJSON struct {
	LogoURL     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTransactionListPendingResponseDataMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListPendingResponseDataMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListArchivedParams struct {
	Year param.Field[int64] `query:"year"`
}

// URLQuery serializes [AccountTransactionListArchivedParams]'s query parameters as
// `url.Values`.
func (r AccountTransactionListArchivedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
