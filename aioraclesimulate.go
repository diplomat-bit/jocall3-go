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

// AIOracleSimulateService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleSimulateService] method instead.
type AIOracleSimulateService struct {
	Options []option.RequestOption
}

// NewAIOracleSimulateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOracleSimulateService(opts ...option.RequestOption) (r *AIOracleSimulateService) {
	r = &AIOracleSimulateService{}
	r.Options = opts
	return
}

func (r *AIOracleSimulateService) RunAdvanced(ctx context.Context, body AIOracleSimulateRunAdvancedParams, opts ...option.RequestOption) (res *AIOracleSimulateRunAdvancedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate/advanced"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *AIOracleSimulateService) RunMonteCarlo(ctx context.Context, body AIOracleSimulateRunMonteCarloParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "ai/oracle/simulate/monte-carlo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Run a 'What-If' Financial Simulation (Standard)
func (r *AIOracleSimulateService) RunStandard(ctx context.Context, body AIOracleSimulateRunStandardParams, opts ...option.RequestOption) (res *AIOracleSimulateRunStandardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIOracleSimulateRunAdvancedResponse struct {
	SimulationID     string                                  `json:"simulationId,required"`
	Status           string                                  `json:"status,required"`
	OutcomeNarrative string                                  `json:"outcomeNarrative"`
	ProjectedValue   float64                                 `json:"projectedValue"`
	JSON             aiOracleSimulateRunAdvancedResponseJSON `json:"-"`
}

// aiOracleSimulateRunAdvancedResponseJSON contains the JSON metadata for the
// struct [AIOracleSimulateRunAdvancedResponse]
type aiOracleSimulateRunAdvancedResponseJSON struct {
	SimulationID     apijson.Field
	Status           apijson.Field
	OutcomeNarrative apijson.Field
	ProjectedValue   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOracleSimulateRunAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulateRunAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulateRunStandardResponse struct {
	SimulationID     string                                  `json:"simulationId,required"`
	Status           string                                  `json:"status,required"`
	OutcomeNarrative string                                  `json:"outcomeNarrative"`
	ProjectedValue   float64                                 `json:"projectedValue"`
	JSON             aiOracleSimulateRunStandardResponseJSON `json:"-"`
}

// aiOracleSimulateRunStandardResponseJSON contains the JSON metadata for the
// struct [AIOracleSimulateRunStandardResponse]
type aiOracleSimulateRunStandardResponseJSON struct {
	SimulationID     apijson.Field
	Status           apijson.Field
	OutcomeNarrative apijson.Field
	ProjectedValue   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOracleSimulateRunStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulateRunStandardResponseJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulateRunAdvancedParams struct {
	Prompt    param.Field[string]                                      `json:"prompt,required"`
	Scenarios param.Field[[]AIOracleSimulateRunAdvancedParamsScenario] `json:"scenarios,required"`
}

func (r AIOracleSimulateRunAdvancedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenario struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
}

func (r AIOracleSimulateRunAdvancedParamsScenario) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunMonteCarloParams struct {
	Iterations param.Field[int64]    `json:"iterations,required"`
	Variables  param.Field[[]string] `json:"variables,required"`
}

func (r AIOracleSimulateRunMonteCarloParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunStandardParams struct {
	// Describe the financial scenario
	Prompt param.Field[string] `json:"prompt,required"`
	// Key variables like duration, rate, or amount
	Parameters param.Field[interface{}] `json:"parameters"`
}

func (r AIOracleSimulateRunStandardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
