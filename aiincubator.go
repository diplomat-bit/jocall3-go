// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AIIncubatorService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorService] method instead.
type AIIncubatorService struct {
	Options []option.RequestOption
	Pitch   *AIIncubatorPitchService
}

// NewAIIncubatorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIIncubatorService(opts ...option.RequestOption) (r *AIIncubatorService) {
	r = &AIIncubatorService{}
	r.Options = opts
	r.Pitch = NewAIIncubatorPitchService(opts...)
	return
}

// Submits a detailed business plan to the Quantum Weaver AI for rigorous analysis,
// market validation, and seed funding consideration. This initiates the AI-driven
// incubation journey, aiming to transform innovative ideas into commercially
// successful ventures.
func (r *AIIncubatorService) GeneratePitch(ctx context.Context, body AIIncubatorGeneratePitchParams, opts ...option.RequestOption) (res *AIIncubatorGeneratePitchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a summary list of all business pitches submitted by the authenticated
// user to Quantum Weaver.
func (r *AIIncubatorService) ListPitches(ctx context.Context, query AIIncubatorListPitchesParams, opts ...option.RequestOption) (res *AIIncubatorListPitchesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AIIncubatorGeneratePitchResponse = interface{}

type AIIncubatorListPitchesResponse = interface{}

type AIIncubatorGeneratePitchParams struct {
	// Key financial metrics and projections for the next 3-5 years.
	FinancialProjections param.Field[interface{}] `json:"financialProjections,required"`
}

func (r AIIncubatorGeneratePitchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorListPitchesParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Filter pitches by their current stage.
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [AIIncubatorListPitchesParams]'s query parameters as
// `url.Values`.
func (r AIIncubatorListPitchesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
