// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
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

// Manually Create Recurring Schedule
func (r *TransactionRecurringService) New(ctx context.Context, body TransactionRecurringNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List Detected Subscriptions
func (r *TransactionRecurringService) List(ctx context.Context, opts ...option.RequestOption) (res *TransactionRecurringListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel Recurring Payment Detection
func (r *TransactionRecurringService) Cancel(ctx context.Context, recurringID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if recurringID == "" {
		err = errors.New("missing required recurringId parameter")
		return
	}
	path := fmt.Sprintf("transactions/recurring/%s", recurringID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type TransactionRecurringListResponse struct {
	Data []TransactionRecurringListResponseData `json:"data"`
	JSON transactionRecurringListResponseJSON   `json:"-"`
}

// transactionRecurringListResponseJSON contains the JSON metadata for the struct
// [TransactionRecurringListResponse]
type transactionRecurringListResponseJSON struct {
	Data        apijson.Field
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
	ID               string                                   `json:"id"`
	Description      string                                   `json:"description"`
	Frequency        string                                   `json:"frequency"`
	NextExpectedDate time.Time                                `json:"nextExpectedDate" format:"date"`
	JSON             transactionRecurringListResponseDataJSON `json:"-"`
}

// transactionRecurringListResponseDataJSON contains the JSON metadata for the
// struct [TransactionRecurringListResponseData]
type transactionRecurringListResponseDataJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Frequency        apijson.Field
	NextExpectedDate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionRecurringListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionRecurringListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TransactionRecurringNewParams struct {
	Amount    param.Field[float64] `json:"amount,required"`
	Category  param.Field[string]  `json:"category,required"`
	Frequency param.Field[string]  `json:"frequency,required"`
}

func (r TransactionRecurringNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
