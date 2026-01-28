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

type BudgetGetResponse struct {
	ID                string                              `json:"id,required"`
	AlertThreshold    int64                               `json:"alertThreshold,required"`
	Categories        []BudgetGetResponseCategory         `json:"categories,required"`
	EndDate           time.Time                           `json:"endDate,required" format:"date"`
	Name              string                              `json:"name,required"`
	Period            string                              `json:"period,required"`
	RemainingAmount   float64                             `json:"remainingAmount,required"`
	SpentAmount       float64                             `json:"spentAmount,required"`
	StartDate         time.Time                           `json:"startDate,required" format:"date"`
	Status            string                              `json:"status,required"`
	TotalAmount       float64                             `json:"totalAmount,required"`
	AIRecommendations []BudgetGetResponseAIRecommendation `json:"aiRecommendations"`
	JSON              budgetGetResponseJSON               `json:"-"`
}

// budgetGetResponseJSON contains the JSON metadata for the struct
// [BudgetGetResponse]
type budgetGetResponseJSON struct {
	ID                apijson.Field
	AlertThreshold    apijson.Field
	Categories        apijson.Field
	EndDate           apijson.Field
	Name              apijson.Field
	Period            apijson.Field
	RemainingAmount   apijson.Field
	SpentAmount       apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	TotalAmount       apijson.Field
	AIRecommendations apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BudgetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetGetResponseJSON) RawJSON() string {
	return r.raw
}

type BudgetGetResponseCategory struct {
	Allocated float64                       `json:"allocated"`
	Name      string                        `json:"name"`
	Remaining float64                       `json:"remaining"`
	Spent     float64                       `json:"spent"`
	JSON      budgetGetResponseCategoryJSON `json:"-"`
}

// budgetGetResponseCategoryJSON contains the JSON metadata for the struct
// [BudgetGetResponseCategory]
type budgetGetResponseCategoryJSON struct {
	Allocated   apijson.Field
	Name        apijson.Field
	Remaining   apijson.Field
	Spent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetGetResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetGetResponseCategoryJSON) RawJSON() string {
	return r.raw
}

type BudgetGetResponseAIRecommendation struct {
	ID                       string                                `json:"id"`
	ActionableRecommendation string                                `json:"actionableRecommendation"`
	Category                 string                                `json:"category"`
	Description              string                                `json:"description"`
	Severity                 string                                `json:"severity"`
	Timestamp                time.Time                             `json:"timestamp" format:"date-time"`
	Title                    string                                `json:"title"`
	JSON                     budgetGetResponseAIRecommendationJSON `json:"-"`
}

// budgetGetResponseAIRecommendationJSON contains the JSON metadata for the struct
// [BudgetGetResponseAIRecommendation]
type budgetGetResponseAIRecommendationJSON struct {
	ID                       apijson.Field
	ActionableRecommendation apijson.Field
	Category                 apijson.Field
	Description              apijson.Field
	Severity                 apijson.Field
	Timestamp                apijson.Field
	Title                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *BudgetGetResponseAIRecommendation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetGetResponseAIRecommendationJSON) RawJSON() string {
	return r.raw
}

type BudgetUpdateResponse struct {
	ID              string                         `json:"id,required"`
	AlertThreshold  int64                          `json:"alertThreshold,required"`
	Categories      []BudgetUpdateResponseCategory `json:"categories,required"`
	EndDate         time.Time                      `json:"endDate,required" format:"date"`
	Name            string                         `json:"name,required"`
	Period          string                         `json:"period,required"`
	RemainingAmount float64                        `json:"remainingAmount,required"`
	SpentAmount     float64                        `json:"spentAmount,required"`
	StartDate       time.Time                      `json:"startDate,required" format:"date"`
	Status          string                         `json:"status,required"`
	TotalAmount     float64                        `json:"totalAmount,required"`
	JSON            budgetUpdateResponseJSON       `json:"-"`
}

// budgetUpdateResponseJSON contains the JSON metadata for the struct
// [BudgetUpdateResponse]
type budgetUpdateResponseJSON struct {
	ID              apijson.Field
	AlertThreshold  apijson.Field
	Categories      apijson.Field
	EndDate         apijson.Field
	Name            apijson.Field
	Period          apijson.Field
	RemainingAmount apijson.Field
	SpentAmount     apijson.Field
	StartDate       apijson.Field
	Status          apijson.Field
	TotalAmount     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BudgetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type BudgetUpdateResponseCategory struct {
	Allocated float64                          `json:"allocated"`
	Name      string                           `json:"name"`
	Remaining float64                          `json:"remaining"`
	Spent     float64                          `json:"spent"`
	JSON      budgetUpdateResponseCategoryJSON `json:"-"`
}

// budgetUpdateResponseCategoryJSON contains the JSON metadata for the struct
// [BudgetUpdateResponseCategory]
type budgetUpdateResponseCategoryJSON struct {
	Allocated   apijson.Field
	Name        apijson.Field
	Remaining   apijson.Field
	Spent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetUpdateResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetUpdateResponseCategoryJSON) RawJSON() string {
	return r.raw
}

type BudgetListResponse struct {
	Data       []BudgetListResponseData `json:"data,required"`
	Limit      int64                    `json:"limit,required"`
	Offset     int64                    `json:"offset,required"`
	Total      int64                    `json:"total,required"`
	NextOffset int64                    `json:"nextOffset"`
	JSON       budgetListResponseJSON   `json:"-"`
}

// budgetListResponseJSON contains the JSON metadata for the struct
// [BudgetListResponse]
type budgetListResponseJSON struct {
	Data        apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetListResponseJSON) RawJSON() string {
	return r.raw
}

type BudgetListResponseData struct {
	ID              string                           `json:"id"`
	AlertThreshold  int64                            `json:"alertThreshold"`
	Categories      []BudgetListResponseDataCategory `json:"categories"`
	EndDate         time.Time                        `json:"endDate" format:"date"`
	Name            string                           `json:"name"`
	Period          string                           `json:"period"`
	RemainingAmount float64                          `json:"remainingAmount"`
	SpentAmount     float64                          `json:"spentAmount"`
	StartDate       time.Time                        `json:"startDate" format:"date"`
	Status          string                           `json:"status"`
	TotalAmount     float64                          `json:"totalAmount"`
	JSON            budgetListResponseDataJSON       `json:"-"`
}

// budgetListResponseDataJSON contains the JSON metadata for the struct
// [BudgetListResponseData]
type budgetListResponseDataJSON struct {
	ID              apijson.Field
	AlertThreshold  apijson.Field
	Categories      apijson.Field
	EndDate         apijson.Field
	Name            apijson.Field
	Period          apijson.Field
	RemainingAmount apijson.Field
	SpentAmount     apijson.Field
	StartDate       apijson.Field
	Status          apijson.Field
	TotalAmount     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BudgetListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetListResponseDataJSON) RawJSON() string {
	return r.raw
}

type BudgetListResponseDataCategory struct {
	Allocated float64                            `json:"allocated"`
	Name      string                             `json:"name"`
	Remaining float64                            `json:"remaining"`
	Spent     float64                            `json:"spent"`
	JSON      budgetListResponseDataCategoryJSON `json:"-"`
}

// budgetListResponseDataCategoryJSON contains the JSON metadata for the struct
// [BudgetListResponseDataCategory]
type budgetListResponseDataCategoryJSON struct {
	Allocated   apijson.Field
	Name        apijson.Field
	Remaining   apijson.Field
	Spent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetListResponseDataCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetListResponseDataCategoryJSON) RawJSON() string {
	return r.raw
}

type BudgetUpdateParams struct {
	AlertThreshold param.Field[int64]   `json:"alertThreshold"`
	TotalAmount    param.Field[float64] `json:"totalAmount"`
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
