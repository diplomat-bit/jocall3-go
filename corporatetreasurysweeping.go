// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// CorporateTreasurySweepingService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasurySweepingService] method instead.
type CorporateTreasurySweepingService struct {
	Options []option.RequestOption
}

// NewCorporateTreasurySweepingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateTreasurySweepingService(opts ...option.RequestOption) (r *CorporateTreasurySweepingService) {
	r = &CorporateTreasurySweepingService{}
	r.Options = opts
	return
}

// Configure Automated Cash Sweeping
func (r *CorporateTreasurySweepingService) ConfigureRules(ctx context.Context, body CorporateTreasurySweepingConfigureRulesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/treasury/sweeping/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Manual Sweep Trigger
func (r *CorporateTreasurySweepingService) Execute(ctx context.Context, body CorporateTreasurySweepingExecuteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/treasury/sweeping/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CorporateTreasurySweepingConfigureRulesParams struct {
	SourceAccount param.Field[string]                                                 `json:"sourceAccount,required"`
	TargetAccount param.Field[string]                                                 `json:"targetAccount,required"`
	Threshold     param.Field[float64]                                                `json:"threshold,required"`
	Frequency     param.Field[CorporateTreasurySweepingConfigureRulesParamsFrequency] `json:"frequency"`
}

func (r CorporateTreasurySweepingConfigureRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateTreasurySweepingConfigureRulesParamsFrequency string

const (
	CorporateTreasurySweepingConfigureRulesParamsFrequencyDaily   CorporateTreasurySweepingConfigureRulesParamsFrequency = "daily"
	CorporateTreasurySweepingConfigureRulesParamsFrequencyWeekly  CorporateTreasurySweepingConfigureRulesParamsFrequency = "weekly"
	CorporateTreasurySweepingConfigureRulesParamsFrequencyMonthly CorporateTreasurySweepingConfigureRulesParamsFrequency = "monthly"
)

func (r CorporateTreasurySweepingConfigureRulesParamsFrequency) IsKnown() bool {
	switch r {
	case CorporateTreasurySweepingConfigureRulesParamsFrequencyDaily, CorporateTreasurySweepingConfigureRulesParamsFrequencyWeekly, CorporateTreasurySweepingConfigureRulesParamsFrequencyMonthly:
		return true
	}
	return false
}

type CorporateTreasurySweepingExecuteParams struct {
	RuleID param.Field[string] `json:"ruleId,required"`
}

func (r CorporateTreasurySweepingExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
