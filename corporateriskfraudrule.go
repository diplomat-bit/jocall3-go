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

// CorporateRiskFraudRuleService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateRiskFraudRuleService] method instead.
type CorporateRiskFraudRuleService struct {
	Options []option.RequestOption
}

// NewCorporateRiskFraudRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateRiskFraudRuleService(opts ...option.RequestOption) (r *CorporateRiskFraudRuleService) {
	r = &CorporateRiskFraudRuleService{}
	r.Options = opts
	return
}

// Updates an existing custom AI-powered fraud detection rule, modifying its
// criteria, actions, or status.
func (r *CorporateRiskFraudRuleService) Update(ctx context.Context, ruleID string, body CorporateRiskFraudRuleUpdateParams, opts ...option.RequestOption) (res *CorporateRiskFraudRuleUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required ruleId parameter")
		return
	}
	path := fmt.Sprintf("corporate/risk/fraud/rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of AI-powered fraud detection rules currently active for the
// organization, including their parameters, thresholds, and associated actions
// (e.g., flag, block, alert).
func (r *CorporateRiskFraudRuleService) List(ctx context.Context, query CorporateRiskFraudRuleListParams, opts ...option.RequestOption) (res *CorporateRiskFraudRuleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CorporateRiskFraudRuleUpdateResponse struct {
	// Action to take when a fraud rule is triggered.
	Action interface{} `json:"action,required"`
	// Criteria that define when a fraud rule should trigger.
	Criteria interface{}                              `json:"criteria,required"`
	JSON     corporateRiskFraudRuleUpdateResponseJSON `json:"-"`
}

// corporateRiskFraudRuleUpdateResponseJSON contains the JSON metadata for the
// struct [CorporateRiskFraudRuleUpdateResponse]
type corporateRiskFraudRuleUpdateResponseJSON struct {
	Action      apijson.Field
	Criteria    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateRiskFraudRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskFraudRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskFraudRuleListResponse = interface{}

type CorporateRiskFraudRuleUpdateParams struct {
	// Action to take when a fraud rule is triggered.
	Action param.Field[interface{}] `json:"action"`
	// Criteria that define when a fraud rule should trigger.
	Criteria param.Field[interface{}] `json:"criteria"`
}

func (r CorporateRiskFraudRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateRiskFraudRuleListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CorporateRiskFraudRuleListParams]'s query parameters as
// `url.Values`.
func (r CorporateRiskFraudRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
