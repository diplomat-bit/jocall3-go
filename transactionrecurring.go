// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
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

// TransactionRecurringService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionRecurringService] method instead.
type TransactionRecurringService struct {
	Options []option.RequestOption
}

// NewTransactionRecurringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionRecurringService(opts ...option.RequestOption) (r *TransactionRecurringService) {
	r = &TransactionRecurringService{}
	r.Options = opts
	return
}

// Retrieves a list of all detected or user-defined recurring transactions, useful
// for budget tracking and subscription management.
func (r *TransactionRecurringService) List(ctx context.Context, query TransactionRecurringListParams, opts ...option.RequestOption) (res *TransactionRecurringListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TransactionRecurringListResponse struct {
	Data       []TransactionRecurringListResponseData `json:"data,required"`
	Limit      int64                                  `json:"limit,required"`
	Offset     int64                                  `json:"offset,required"`
	Total      int64                                  `json:"total,required"`
	NextOffset int64                                  `json:"nextOffset"`
	JSON       transactionRecurringListResponseJSON   `json:"-"`
}

// transactionRecurringListResponseJSON contains the JSON metadata for the struct
// [TransactionRecurringListResponse]
type transactionRecurringListResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionRecurringListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionRecurringListResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionRecurringListResponseData struct {
	ID                string                                   `json:"id"`
	AIConfidenceScore float64                                  `json:"aiConfidenceScore"`
	Amount            float64                                  `json:"amount"`
	Category          string                                   `json:"category"`
	Currency          string                                   `json:"currency"`
	Description       string                                   `json:"description"`
	Frequency         string                                   `json:"frequency"`
	LastPaidDate      time.Time                                `json:"lastPaidDate" format:"date"`
	LinkedAccountID   string                                   `json:"linkedAccountId"`
	NextDueDate       time.Time                                `json:"nextDueDate" format:"date"`
	Status            string                                   `json:"status"`
	JSON              transactionRecurringListResponseDataJSON `json:"-"`
}

// transactionRecurringListResponseDataJSON contains the JSON metadata for the
// struct [TransactionRecurringListResponseData]
type transactionRecurringListResponseDataJSON struct {
	ID                apijson.Field
	AIConfidenceScore apijson.Field
	Amount            apijson.Field
	Category          apijson.Field
	Currency          apijson.Field
	Description       apijson.Field
	Frequency         apijson.Field
	LastPaidDate      apijson.Field
	LinkedAccountID   apijson.Field
	NextDueDate       apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionRecurringListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionRecurringListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TransactionRecurringListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [TransactionRecurringListParams]'s query parameters as
// `url.Values`.
func (r TransactionRecurringListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
