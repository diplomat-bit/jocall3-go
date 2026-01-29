// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
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

// Retrieves the full, detailed results of a specific financial simulation by its
// ID.
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

// Retrieves a list of all financial simulations previously run by the user,
// including their status and summaries.
func (r *AIOracleSimulationService) List(ctx context.Context, query AIOracleSimulationListParams, opts ...option.RequestOption) (res *AIOracleSimulationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AIOracleSimulationListResponse = interface{}

type AIOracleSimulationListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AIOracleSimulationListParams]'s query parameters as
// `url.Values`.
func (r AIOracleSimulationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
