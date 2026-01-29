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

// Create Custom Fraud Rule
func (r *CorporateRiskFraudRuleService) New(ctx context.Context, body CorporateRiskFraudRuleNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/risk/fraud/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Update a fraud rule
func (r *CorporateRiskFraudRuleService) Update(ctx context.Context, ruleID string, body CorporateRiskFraudRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if ruleID == "" {
		err = errors.New("missing required ruleId parameter")
		return
	}
	path := fmt.Sprintf("corporate/risk/fraud/rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// List Active Fraud Rule Set
func (r *CorporateRiskFraudRuleService) List(ctx context.Context, opts ...option.RequestOption) (res *CorporateRiskFraudRuleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/risk/fraud/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CorporateRiskFraudRuleListResponse struct {
	Rules []interface{}                          `json:"rules"`
	JSON  corporateRiskFraudRuleListResponseJSON `json:"-"`
}

// corporateRiskFraudRuleListResponseJSON contains the JSON metadata for the struct
// [CorporateRiskFraudRuleListResponse]
type corporateRiskFraudRuleListResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateRiskFraudRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateRiskFraudRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateRiskFraudRuleNewParams struct {
	Logic param.Field[interface{}] `json:"logic,required"`
	Name  param.Field[string]      `json:"name,required"`
}

func (r CorporateRiskFraudRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateRiskFraudRuleUpdateParams struct {
	Action param.Field[string] `json:"action"`
	Name   param.Field[string] `json:"name"`
}

func (r CorporateRiskFraudRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
