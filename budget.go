// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// BudgetService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetService] method instead.
type BudgetService struct {
	Options []option.RequestOption
}

// NewBudgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBudgetService(opts ...option.RequestOption) (r *BudgetService) {
	r = &BudgetService{}
	r.Options = opts
	return
}

// Retrieves detailed information for a specific budget, including current
// spending, remaining amounts, and AI recommendations.
func (r *BudgetService) Get(ctx context.Context, budgetID string, opts ...option.RequestOption) (res *BudgetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return
	}
	path := fmt.Sprintf("budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the parameters of an existing budget, such as total amount, dates, or
// categories.
func (r *BudgetService) Update(ctx context.Context, budgetID string, body BudgetUpdateParams, opts ...option.RequestOption) (res *BudgetUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return
	}
	path := fmt.Sprintf("budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of all active and historical budgets for the authenticated
// user.
func (r *BudgetService) List(ctx context.Context, query BudgetListParams, opts ...option.RequestOption) (res *BudgetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BudgetGetResponse = interface{}

type BudgetUpdateResponse = interface{}

type BudgetListResponse = interface{}

type BudgetUpdateParams struct {
}

func (r BudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BudgetListParams]'s query parameters as `url.Values`.
func (r BudgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
