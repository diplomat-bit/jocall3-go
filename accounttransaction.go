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

// Retrieves a list of pending transactions that have not yet cleared for a
// specific financial account.
func (r *AccountTransactionService) ListPending(ctx context.Context, accountID string, query AccountTransactionListPendingParams, opts ...option.RequestOption) (res *AccountTransactionListPendingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/transactions/pending", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTransactionListPendingResponse struct {
	Data       []AccountTransactionListPendingResponseData `json:"data,required"`
	Limit      int64                                       `json:"limit,required"`
	Offset     int64                                       `json:"offset,required"`
	Total      int64                                       `json:"total,required"`
	NextOffset int64                                       `json:"nextOffset"`
	JSON       accountTransactionListPendingResponseJSON   `json:"-"`
}

// accountTransactionListPendingResponseJSON contains the JSON metadata for the
// struct [AccountTransactionListPendingResponse]
type accountTransactionListPendingResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
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
	ID                   string                                        `json:"id"`
	AccountID            string                                        `json:"accountId"`
	AICategoryConfidence float64                                       `json:"aiCategoryConfidence"`
	Amount               float64                                       `json:"amount"`
	CarbonFootprint      float64                                       `json:"carbonFootprint"`
	Category             string                                        `json:"category"`
	Currency             string                                        `json:"currency"`
	Date                 time.Time                                     `json:"date" format:"date"`
	Description          string                                        `json:"description"`
	DisputeStatus        string                                        `json:"disputeStatus"`
	PaymentChannel       string                                        `json:"paymentChannel"`
	Type                 string                                        `json:"type"`
	JSON                 accountTransactionListPendingResponseDataJSON `json:"-"`
}

// accountTransactionListPendingResponseDataJSON contains the JSON metadata for the
// struct [AccountTransactionListPendingResponseData]
type accountTransactionListPendingResponseDataJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	AICategoryConfidence apijson.Field
	Amount               apijson.Field
	CarbonFootprint      apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	DisputeStatus        apijson.Field
	PaymentChannel       apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountTransactionListPendingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTransactionListPendingResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountTransactionListPendingParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AccountTransactionListPendingParams]'s query parameters as
// `url.Values`.
func (r AccountTransactionListPendingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
