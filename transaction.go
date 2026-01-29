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

// TransactionService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options   []option.RequestOption
	Recurring *TransactionRecurringService
	Insights  *TransactionInsightService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.Recurring = NewTransactionRecurringService(opts...)
	r.Insights = NewTransactionInsightService(opts...)
	return
}

// Retrieves granular information for a single transaction by its unique ID,
// including AI categorization confidence, merchant details, and associated carbon
// footprint.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *TransactionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of the user's transactions, with extensive options
// for filtering by type, category, date range, amount, and intelligent AI-driven
// sorting and search capabilities.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows the user to add or update personal notes for a specific transaction.
func (r *TransactionService) AddNotes(ctx context.Context, transactionID string, body TransactionAddNotesParams, opts ...option.RequestOption) (res *TransactionAddNotesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/notes", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Allows the user to override or refine the AI's categorization for a transaction,
// improving future AI accuracy and personal financial reporting.
func (r *TransactionService) Categorize(ctx context.Context, transactionID string, body TransactionCategorizeParams, opts ...option.RequestOption) (res *TransactionCategorizeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/categorize", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type TransactionGetResponse struct {
	ID                   string                                `json:"id,required"`
	AccountID            string                                `json:"accountId,required"`
	Amount               float64                               `json:"amount,required"`
	Category             string                                `json:"category,required"`
	Currency             string                                `json:"currency,required"`
	Date                 time.Time                             `json:"date,required" format:"date"`
	Description          string                                `json:"description,required"`
	Type                 string                                `json:"type,required"`
	AICategoryConfidence float64                               `json:"aiCategoryConfidence"`
	CarbonFootprint      float64                               `json:"carbonFootprint"`
	DisputeStatus        string                                `json:"disputeStatus"`
	Location             TransactionGetResponseLocation        `json:"location"`
	MerchantDetails      TransactionGetResponseMerchantDetails `json:"merchantDetails"`
	Notes                string                                `json:"notes"`
	PaymentChannel       string                                `json:"paymentChannel"`
	PostedDate           time.Time                             `json:"postedDate" format:"date"`
	ReceiptURL           string                                `json:"receiptUrl"`
	Tags                 []string                              `json:"tags"`
	JSON                 transactionGetResponseJSON            `json:"-"`
}

// transactionGetResponseJSON contains the JSON metadata for the struct
// [TransactionGetResponse]
type transactionGetResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Amount               apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	Type                 apijson.Field
	AICategoryConfidence apijson.Field
	CarbonFootprint      apijson.Field
	DisputeStatus        apijson.Field
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionGetResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionGetResponseLocation struct {
	City      string                             `json:"city"`
	Latitude  float64                            `json:"latitude"`
	Longitude float64                            `json:"longitude"`
	JSON      transactionGetResponseLocationJSON `json:"-"`
}

// transactionGetResponseLocationJSON contains the JSON metadata for the struct
// [TransactionGetResponseLocation]
type transactionGetResponseLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionGetResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionGetResponseLocationJSON) RawJSON() string {
	return r.raw
}

type TransactionGetResponseMerchantDetails struct {
	Address TransactionGetResponseMerchantDetailsAddress `json:"address"`
	LogoURL string                                       `json:"logoUrl"`
	Name    string                                       `json:"name"`
	Website string                                       `json:"website"`
	JSON    transactionGetResponseMerchantDetailsJSON    `json:"-"`
}

// transactionGetResponseMerchantDetailsJSON contains the JSON metadata for the
// struct [TransactionGetResponseMerchantDetails]
type transactionGetResponseMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionGetResponseMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionGetResponseMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionGetResponseMerchantDetailsAddress struct {
	City  string                                           `json:"city"`
	State string                                           `json:"state"`
	Zip   string                                           `json:"zip"`
	JSON  transactionGetResponseMerchantDetailsAddressJSON `json:"-"`
}

// transactionGetResponseMerchantDetailsAddressJSON contains the JSON metadata for
// the struct [TransactionGetResponseMerchantDetailsAddress]
type transactionGetResponseMerchantDetailsAddressJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionGetResponseMerchantDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionGetResponseMerchantDetailsAddressJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponse struct {
	Data       []TransactionListResponseData `json:"data,required"`
	Limit      int64                         `json:"limit,required"`
	Offset     int64                         `json:"offset,required"`
	Total      int64                         `json:"total,required"`
	NextOffset int64                         `json:"nextOffset"`
	JSON       transactionListResponseJSON   `json:"-"`
}

// transactionListResponseJSON contains the JSON metadata for the struct
// [TransactionListResponse]
type transactionListResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponseData struct {
	ID                   string                                     `json:"id"`
	AccountID            string                                     `json:"accountId"`
	AICategoryConfidence float64                                    `json:"aiCategoryConfidence"`
	Amount               float64                                    `json:"amount"`
	CarbonFootprint      float64                                    `json:"carbonFootprint"`
	Category             string                                     `json:"category"`
	Currency             string                                     `json:"currency"`
	Date                 time.Time                                  `json:"date" format:"date"`
	Description          string                                     `json:"description"`
	DisputeStatus        string                                     `json:"disputeStatus"`
	Location             TransactionListResponseDataLocation        `json:"location"`
	MerchantDetails      TransactionListResponseDataMerchantDetails `json:"merchantDetails"`
	Notes                string                                     `json:"notes"`
	PaymentChannel       string                                     `json:"paymentChannel"`
	PostedDate           time.Time                                  `json:"postedDate" format:"date"`
	ReceiptURL           string                                     `json:"receiptUrl"`
	Tags                 []string                                   `json:"tags"`
	Type                 string                                     `json:"type"`
	JSON                 transactionListResponseDataJSON            `json:"-"`
}

// transactionListResponseDataJSON contains the JSON metadata for the struct
// [TransactionListResponseData]
type transactionListResponseDataJSON struct {
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
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponseDataLocation struct {
	City      string                                  `json:"city"`
	Latitude  float64                                 `json:"latitude"`
	Longitude float64                                 `json:"longitude"`
	JSON      transactionListResponseDataLocationJSON `json:"-"`
}

// transactionListResponseDataLocationJSON contains the JSON metadata for the
// struct [TransactionListResponseDataLocation]
type transactionListResponseDataLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponseDataLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseDataLocationJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponseDataMerchantDetails struct {
	Address TransactionListResponseDataMerchantDetailsAddress `json:"address"`
	LogoURL string                                            `json:"logoUrl"`
	Name    string                                            `json:"name"`
	Website string                                            `json:"website"`
	JSON    transactionListResponseDataMerchantDetailsJSON    `json:"-"`
}

// transactionListResponseDataMerchantDetailsJSON contains the JSON metadata for
// the struct [TransactionListResponseDataMerchantDetails]
type transactionListResponseDataMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponseDataMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseDataMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponseDataMerchantDetailsAddress struct {
	City  string                                                `json:"city"`
	State string                                                `json:"state"`
	Zip   string                                                `json:"zip"`
	JSON  transactionListResponseDataMerchantDetailsAddressJSON `json:"-"`
}

// transactionListResponseDataMerchantDetailsAddressJSON contains the JSON metadata
// for the struct [TransactionListResponseDataMerchantDetailsAddress]
type transactionListResponseDataMerchantDetailsAddressJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponseDataMerchantDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseDataMerchantDetailsAddressJSON) RawJSON() string {
	return r.raw
}

type TransactionAddNotesResponse struct {
	ID                   string                                     `json:"id,required"`
	AccountID            string                                     `json:"accountId,required"`
	Amount               float64                                    `json:"amount,required"`
	Category             string                                     `json:"category,required"`
	Currency             string                                     `json:"currency,required"`
	Date                 time.Time                                  `json:"date,required" format:"date"`
	Description          string                                     `json:"description,required"`
	Type                 string                                     `json:"type,required"`
	AICategoryConfidence float64                                    `json:"aiCategoryConfidence"`
	CarbonFootprint      float64                                    `json:"carbonFootprint"`
	DisputeStatus        string                                     `json:"disputeStatus"`
	Location             TransactionAddNotesResponseLocation        `json:"location"`
	MerchantDetails      TransactionAddNotesResponseMerchantDetails `json:"merchantDetails"`
	Notes                string                                     `json:"notes"`
	PaymentChannel       string                                     `json:"paymentChannel"`
	PostedDate           time.Time                                  `json:"postedDate" format:"date"`
	ReceiptURL           string                                     `json:"receiptUrl"`
	Tags                 []string                                   `json:"tags"`
	JSON                 transactionAddNotesResponseJSON            `json:"-"`
}

// transactionAddNotesResponseJSON contains the JSON metadata for the struct
// [TransactionAddNotesResponse]
type transactionAddNotesResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Amount               apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	Type                 apijson.Field
	AICategoryConfidence apijson.Field
	CarbonFootprint      apijson.Field
	DisputeStatus        apijson.Field
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionAddNotesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAddNotesResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionAddNotesResponseLocation struct {
	City      string                                  `json:"city"`
	Latitude  float64                                 `json:"latitude"`
	Longitude float64                                 `json:"longitude"`
	JSON      transactionAddNotesResponseLocationJSON `json:"-"`
}

// transactionAddNotesResponseLocationJSON contains the JSON metadata for the
// struct [TransactionAddNotesResponseLocation]
type transactionAddNotesResponseLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAddNotesResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAddNotesResponseLocationJSON) RawJSON() string {
	return r.raw
}

type TransactionAddNotesResponseMerchantDetails struct {
	Address TransactionAddNotesResponseMerchantDetailsAddress `json:"address"`
	LogoURL string                                            `json:"logoUrl"`
	Name    string                                            `json:"name"`
	Website string                                            `json:"website"`
	JSON    transactionAddNotesResponseMerchantDetailsJSON    `json:"-"`
}

// transactionAddNotesResponseMerchantDetailsJSON contains the JSON metadata for
// the struct [TransactionAddNotesResponseMerchantDetails]
type transactionAddNotesResponseMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAddNotesResponseMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAddNotesResponseMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionAddNotesResponseMerchantDetailsAddress struct {
	City  string                                                `json:"city"`
	State string                                                `json:"state"`
	Zip   string                                                `json:"zip"`
	JSON  transactionAddNotesResponseMerchantDetailsAddressJSON `json:"-"`
}

// transactionAddNotesResponseMerchantDetailsAddressJSON contains the JSON metadata
// for the struct [TransactionAddNotesResponseMerchantDetailsAddress]
type transactionAddNotesResponseMerchantDetailsAddressJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionAddNotesResponseMerchantDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAddNotesResponseMerchantDetailsAddressJSON) RawJSON() string {
	return r.raw
}

type TransactionCategorizeResponse struct {
	ID                   string                                       `json:"id,required"`
	AccountID            string                                       `json:"accountId,required"`
	Amount               float64                                      `json:"amount,required"`
	Category             string                                       `json:"category,required"`
	Currency             string                                       `json:"currency,required"`
	Date                 time.Time                                    `json:"date,required" format:"date"`
	Description          string                                       `json:"description,required"`
	Type                 string                                       `json:"type,required"`
	AICategoryConfidence float64                                      `json:"aiCategoryConfidence"`
	CarbonFootprint      float64                                      `json:"carbonFootprint"`
	DisputeStatus        string                                       `json:"disputeStatus"`
	Location             TransactionCategorizeResponseLocation        `json:"location"`
	MerchantDetails      TransactionCategorizeResponseMerchantDetails `json:"merchantDetails"`
	Notes                string                                       `json:"notes"`
	PaymentChannel       string                                       `json:"paymentChannel"`
	PostedDate           time.Time                                    `json:"postedDate" format:"date"`
	ReceiptURL           string                                       `json:"receiptUrl"`
	Tags                 []string                                     `json:"tags"`
	JSON                 transactionCategorizeResponseJSON            `json:"-"`
}

// transactionCategorizeResponseJSON contains the JSON metadata for the struct
// [TransactionCategorizeResponse]
type transactionCategorizeResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Amount               apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	Type                 apijson.Field
	AICategoryConfidence apijson.Field
	CarbonFootprint      apijson.Field
	DisputeStatus        apijson.Field
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionCategorizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategorizeResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionCategorizeResponseLocation struct {
	City      string                                    `json:"city"`
	Latitude  float64                                   `json:"latitude"`
	Longitude float64                                   `json:"longitude"`
	JSON      transactionCategorizeResponseLocationJSON `json:"-"`
}

// transactionCategorizeResponseLocationJSON contains the JSON metadata for the
// struct [TransactionCategorizeResponseLocation]
type transactionCategorizeResponseLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionCategorizeResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategorizeResponseLocationJSON) RawJSON() string {
	return r.raw
}

type TransactionCategorizeResponseMerchantDetails struct {
	Address TransactionCategorizeResponseMerchantDetailsAddress `json:"address"`
	LogoURL string                                              `json:"logoUrl"`
	Name    string                                              `json:"name"`
	Website string                                              `json:"website"`
	JSON    transactionCategorizeResponseMerchantDetailsJSON    `json:"-"`
}

// transactionCategorizeResponseMerchantDetailsJSON contains the JSON metadata for
// the struct [TransactionCategorizeResponseMerchantDetails]
type transactionCategorizeResponseMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionCategorizeResponseMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategorizeResponseMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionCategorizeResponseMerchantDetailsAddress struct {
	City  string                                                  `json:"city"`
	State string                                                  `json:"state"`
	Zip   string                                                  `json:"zip"`
	JSON  transactionCategorizeResponseMerchantDetailsAddressJSON `json:"-"`
}

// transactionCategorizeResponseMerchantDetailsAddressJSON contains the JSON
// metadata for the struct [TransactionCategorizeResponseMerchantDetailsAddress]
type transactionCategorizeResponseMerchantDetailsAddressJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionCategorizeResponseMerchantDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategorizeResponseMerchantDetailsAddressJSON) RawJSON() string {
	return r.raw
}

type TransactionListParams struct {
	// Filter transactions by their AI-assigned or user-defined category.
	Category param.Field[string] `query:"category"`
	// Retrieve transactions up to this date (inclusive).
	EndDate param.Field[string] `query:"endDate"`
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Filter for transactions with an amount less than or equal to this value.
	MaxAmount param.Field[int64] `query:"maxAmount"`
	// Filter for transactions with an amount greater than or equal to this value.
	MinAmount param.Field[int64] `query:"minAmount"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Free-text search across transaction descriptions, merchants, and notes.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Retrieve transactions from this date (inclusive).
	StartDate param.Field[string] `query:"startDate"`
	// Filter transactions by type (e.g., income, expense, transfer).
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionAddNotesParams struct {
	Notes param.Field[string] `json:"notes,required"`
}

func (r TransactionAddNotesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionCategorizeParams struct {
	Category      param.Field[string] `json:"category,required"`
	ApplyToFuture param.Field[bool]   `json:"applyToFuture"`
	Notes         param.Field[string] `json:"notes"`
}

func (r TransactionCategorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
