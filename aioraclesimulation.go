// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIOracleSimulationService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleSimulationService] method instead.
type AIOracleSimulationService struct {
	Options []option.RequestOption
}

// NewAIOracleSimulationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOracleSimulationService(opts ...option.RequestOption) (r *AIOracleSimulationService) {
	r = &AIOracleSimulationService{}
	r.Options = opts
	return
}

// Get Specific Simulation Result
func (r *AIOracleSimulationService) Get(ctx context.Context, simulationID string, opts ...option.RequestOption) (res *AIOracleSimulationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if simulationID == "" {
		err = errors.New("missing required simulationId parameter")
		return
	}
	path := fmt.Sprintf("ai/oracle/simulations/%s", simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List All Past Simulations
func (r *AIOracleSimulationService) List(ctx context.Context, opts ...option.RequestOption) (res *AIOracleSimulationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIOracleSimulationGetResponse struct {
	SimulationID     string                            `json:"simulationId,required"`
	Status           string                            `json:"status,required"`
	ConfidenceScore  float64                           `json:"confidenceScore"`
	Data             []map[string]interface{}          `json:"data"`
	OutcomeNarrative string                            `json:"outcomeNarrative"`
	ProjectedValue   float64                           `json:"projectedValue"`
	JSON             aiOracleSimulationGetResponseJSON `json:"-"`
}

// aiOracleSimulationGetResponseJSON contains the JSON metadata for the struct
// [AIOracleSimulationGetResponse]
type aiOracleSimulationGetResponseJSON struct {
	SimulationID     apijson.Field
	Status           apijson.Field
	ConfidenceScore  apijson.Field
	Data             apijson.Field
	OutcomeNarrative apijson.Field
	ProjectedValue   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOracleSimulationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulationListResponse struct {
	Data []AIOracleSimulationListResponseData `json:"data"`
	JSON aiOracleSimulationListResponseJSON   `json:"-"`
}

// aiOracleSimulationListResponseJSON contains the JSON metadata for the struct
// [AIOracleSimulationListResponse]
type aiOracleSimulationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIOracleSimulationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulationListResponseJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulationListResponseData struct {
	SimulationID     string                                 `json:"simulationId,required"`
	Status           string                                 `json:"status,required"`
	ConfidenceScore  float64                                `json:"confidenceScore"`
	Data             []map[string]interface{}               `json:"data"`
	OutcomeNarrative string                                 `json:"outcomeNarrative"`
	ProjectedValue   float64                                `json:"projectedValue"`
	JSON             aiOracleSimulationListResponseDataJSON `json:"-"`
}

// aiOracleSimulationListResponseDataJSON contains the JSON metadata for the struct
// [AIOracleSimulationListResponseData]
type aiOracleSimulationListResponseDataJSON struct {
	SimulationID     apijson.Field
	Status           apijson.Field
	ConfidenceScore  apijson.Field
	Data             apijson.Field
	OutcomeNarrative apijson.Field
	ProjectedValue   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIOracleSimulationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulationListResponseDataJSON) RawJSON() string {
	return r.raw
}
