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

// CorporateAnomalyService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateAnomalyService] method instead.
type CorporateAnomalyService struct {
	Options []option.RequestOption
}

// NewCorporateAnomalyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateAnomalyService(opts ...option.RequestOption) (r *CorporateAnomalyService) {
	r = &CorporateAnomalyService{}
	r.Options = opts
	return
}

// Retrieves a comprehensive list of AI-detected financial anomalies across
// transactions, payments, and corporate cards that require immediate review and
// potential action to mitigate risk and ensure compliance.
func (r *CorporateAnomalyService) List(ctx context.Context, query CorporateAnomalyListParams, opts ...option.RequestOption) (res *CorporateAnomalyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the review status of a specific financial anomaly, allowing compliance
// officers to mark it as dismissed, resolved, or escalate for further
// investigation after thorough AI-assisted and human review.
func (r *CorporateAnomalyService) UpdateStatus(ctx context.Context, anomalyID string, body CorporateAnomalyUpdateStatusParams, opts ...option.RequestOption) (res *CorporateAnomalyUpdateStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if anomalyID == "" {
		err = errors.New("missing required anomalyId parameter")
		return
	}
	path := fmt.Sprintf("corporate/anomalies/%s/status", anomalyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CorporateAnomalyListResponse = interface{}

type CorporateAnomalyUpdateStatusResponse = interface{}

type CorporateAnomalyListParams struct {
	// End date for filtering results (inclusive, YYYY-MM-DD).
	EndDate param.Field[string] `query:"endDate"`
	// Filter anomalies by the type of financial entity they are related to.
	EntityType param.Field[string] `query:"entityType"`
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Filter anomalies by their AI-assessed severity level.
	Severity param.Field[string] `query:"severity"`
	// Start date for filtering results (inclusive, YYYY-MM-DD).
	StartDate param.Field[string] `query:"startDate"`
	// Filter anomalies by their current review status.
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [CorporateAnomalyListParams]'s query parameters as
// `url.Values`.
func (r CorporateAnomalyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateAnomalyUpdateStatusParams struct {
}

func (r CorporateAnomalyUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
