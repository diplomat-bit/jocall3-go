// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// LendingService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendingService] method instead.
type LendingService struct {
	Options   []option.RequestOption
	Decisions *LendingDecisionService
}

// NewLendingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLendingService(opts ...option.RequestOption) (r *LendingService) {
	r = &LendingService{}
	r.Options = opts
	r.Decisions = NewLendingDecisionService(opts...)
	return
}

// Track Loan Processing
func (r *LendingService) GetStatus(ctx context.Context, appID string, opts ...option.RequestOption) (res *LendingGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("lending/applications/%s/status", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submit Advanced Credit Application
func (r *LendingService) SubmitApplication(ctx context.Context, body LendingSubmitApplicationParams, opts ...option.RequestOption) (res *LendingSubmitApplicationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lending/applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendingGetStatusResponse struct {
	Status              string                       `json:"status"`
	UnderwriterQueuePos int64                        `json:"underwriterQueuePos"`
	JSON                lendingGetStatusResponseJSON `json:"-"`
}

// lendingGetStatusResponseJSON contains the JSON metadata for the struct
// [LendingGetStatusResponse]
type lendingGetStatusResponseJSON struct {
	Status              apijson.Field
	UnderwriterQueuePos apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LendingGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lendingGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type LendingSubmitApplicationResponse struct {
	ApplicationID string                               `json:"applicationId"`
	Status        string                               `json:"status"`
	JSON          lendingSubmitApplicationResponseJSON `json:"-"`
}

// lendingSubmitApplicationResponseJSON contains the JSON metadata for the struct
// [LendingSubmitApplicationResponse]
type lendingSubmitApplicationResponseJSON struct {
	ApplicationID apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LendingSubmitApplicationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lendingSubmitApplicationResponseJSON) RawJSON() string {
	return r.raw
}

type LendingSubmitApplicationParams struct {
	Amount         param.Field[float64]                                      `json:"amount,required"`
	EmploymentData param.Field[LendingSubmitApplicationParamsEmploymentData] `json:"employmentData,required"`
	LoanType       param.Field[LendingSubmitApplicationParamsLoanType]       `json:"loanType,required"`
	TermMonths     param.Field[int64]                                        `json:"termMonths,required"`
	Assets         param.Field[[]interface{}]                                `json:"assets"`
	CollateralID   param.Field[string]                                       `json:"collateralId"`
	Liabilities    param.Field[[]interface{}]                                `json:"liabilities"`
}

func (r LendingSubmitApplicationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LendingSubmitApplicationParamsEmploymentData struct {
	Employer      param.Field[string]  `json:"employer,required"`
	MonthlyIncome param.Field[float64] `json:"monthlyIncome,required"`
	TenureMonths  param.Field[int64]   `json:"tenureMonths"`
}

func (r LendingSubmitApplicationParamsEmploymentData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LendingSubmitApplicationParamsLoanType string

const (
	LendingSubmitApplicationParamsLoanTypeMortgage          LendingSubmitApplicationParamsLoanType = "MORTGAGE"
	LendingSubmitApplicationParamsLoanTypePersonal          LendingSubmitApplicationParamsLoanType = "PERSONAL"
	LendingSubmitApplicationParamsLoanTypeAuto              LendingSubmitApplicationParamsLoanType = "AUTO"
	LendingSubmitApplicationParamsLoanTypeBusinessExpansion LendingSubmitApplicationParamsLoanType = "BUSINESS_EXPANSION"
)

func (r LendingSubmitApplicationParamsLoanType) IsKnown() bool {
	switch r {
	case LendingSubmitApplicationParamsLoanTypeMortgage, LendingSubmitApplicationParamsLoanTypePersonal, LendingSubmitApplicationParamsLoanTypeAuto, LendingSubmitApplicationParamsLoanTypeBusinessExpansion:
		return true
	}
	return false
}
