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

// Get Transaction Deep Metadata
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

// Global Transaction Search & Filter
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Attach Manual Notes to Transaction
func (r *TransactionService) AddNotes(ctx context.Context, transactionID string, body TransactionAddNotesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/notes", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Override AI Categorization
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

// Initiate Transaction Dispute
func (r *TransactionService) Dispute(ctx context.Context, transactionID string, body TransactionDisputeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/dispute", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Split Transaction Across Multiple Categories
func (r *TransactionService) Split(ctx context.Context, transactionID string, body TransactionSplitParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/split", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type TransactionGetResponse struct {
	ID          string                     `json:"id,required"`
	Amount      float64                    `json:"amount,required"`
	Currency    string                     `json:"currency,required"`
	Date        time.Time                  `json:"date,required" format:"date"`
	Description string                     `json:"description,required"`
	Category    string                     `json:"category"`
	Notes       string                     `json:"notes"`
	JSON        transactionGetResponseJSON `json:"-"`
}

// transactionGetResponseJSON contains the JSON metadata for the struct
// [TransactionGetResponse]
type transactionGetResponseJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Date        apijson.Field
	Description apijson.Field
	Category    apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionGetResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponse struct {
	Data       []TransactionListResponseData `json:"data,required"`
	Total      int64                         `json:"total,required"`
	NextOffset int64                         `json:"nextOffset"`
	JSON       transactionListResponseJSON   `json:"-"`
}

// transactionListResponseJSON contains the JSON metadata for the struct
// [TransactionListResponse]
type transactionListResponseJSON struct {
	Data        apijson.Field
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
	ID          string                          `json:"id,required"`
	Amount      float64                         `json:"amount,required"`
	Currency    string                          `json:"currency,required"`
	Date        time.Time                       `json:"date,required" format:"date"`
	Description string                          `json:"description,required"`
	Category    string                          `json:"category"`
	Notes       string                          `json:"notes"`
	JSON        transactionListResponseDataJSON `json:"-"`
}

// transactionListResponseDataJSON contains the JSON metadata for the struct
// [TransactionListResponseData]
type transactionListResponseDataJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Date        apijson.Field
	Description apijson.Field
	Category    apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TransactionCategorizeResponse struct {
	ID          string                            `json:"id,required"`
	Amount      float64                           `json:"amount,required"`
	Currency    string                            `json:"currency,required"`
	Date        time.Time                         `json:"date,required" format:"date"`
	Description string                            `json:"description,required"`
	Category    string                            `json:"category"`
	Notes       string                            `json:"notes"`
	JSON        transactionCategorizeResponseJSON `json:"-"`
}

// transactionCategorizeResponseJSON contains the JSON metadata for the struct
// [TransactionCategorizeResponse]
type transactionCategorizeResponseJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Date        apijson.Field
	Description apijson.Field
	Category    apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionCategorizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionCategorizeResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionListParams struct {
	Limit     param.Field[int64]   `query:"limit"`
	MaxAmount param.Field[float64] `query:"maxAmount"`
	MinAmount param.Field[float64] `query:"minAmount"`
	Offset    param.Field[int64]   `query:"offset"`
	Type      param.Field[string]  `query:"type"`
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
}

func (r TransactionCategorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDisputeParams struct {
	Reason param.Field[TransactionDisputeParamsReason] `json:"reason,required"`
	// URIs to evidence
	EvidenceFiles param.Field[[]string] `json:"evidenceFiles"`
}

func (r TransactionDisputeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDisputeParamsReason string

const (
	TransactionDisputeParamsReasonFraudulent         TransactionDisputeParamsReason = "fraudulent"
	TransactionDisputeParamsReasonDuplicate          TransactionDisputeParamsReason = "duplicate"
	TransactionDisputeParamsReasonIncorrectAmount    TransactionDisputeParamsReason = "incorrect_amount"
	TransactionDisputeParamsReasonServiceNotRendered TransactionDisputeParamsReason = "service_not_rendered"
)

func (r TransactionDisputeParamsReason) IsKnown() bool {
	switch r {
	case TransactionDisputeParamsReasonFraudulent, TransactionDisputeParamsReasonDuplicate, TransactionDisputeParamsReasonIncorrectAmount, TransactionDisputeParamsReasonServiceNotRendered:
		return true
	}
	return false
}

type TransactionSplitParams struct {
	Splits param.Field[[]TransactionSplitParamsSplit] `json:"splits,required"`
}

func (r TransactionSplitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionSplitParamsSplit struct {
	Amount   param.Field[float64] `json:"amount"`
	Category param.Field[string]  `json:"category"`
}

func (r TransactionSplitParamsSplit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
