// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// CorporateComplianceService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateComplianceService] method instead.
type CorporateComplianceService struct {
	Options []option.RequestOption
	Audits  *CorporateComplianceAuditService
}

// NewCorporateComplianceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateComplianceService(opts ...option.RequestOption) (r *CorporateComplianceService) {
	r = &CorporateComplianceService{}
	r.Options = opts
	r.Audits = NewCorporateComplianceAuditService(opts...)
	return
}

// Adverse Media Sentiment Screening
func (r *CorporateComplianceService) ScreenAdverseMedia(ctx context.Context, body CorporateComplianceScreenAdverseMediaParams, opts ...option.RequestOption) (res *CorporateComplianceScreenAdverseMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/compliance/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Politically Exposed Person (PEP) Screening
func (r *CorporateComplianceService) ScreenPep(ctx context.Context, body CorporateComplianceScreenPepParams, opts ...option.RequestOption) (res *CorporateComplianceScreenPepResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/compliance/pep"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Enhanced Global Sanctions Screening
func (r *CorporateComplianceService) ScreenSanctions(ctx context.Context, body CorporateComplianceScreenSanctionsParams, opts ...option.RequestOption) (res *CorporateComplianceScreenSanctionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/compliance/sanctions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateComplianceScreenAdverseMediaResponse struct {
	NegativeNewsLinks []string                                          `json:"negativeNewsLinks"`
	SentimentScore    float64                                           `json:"sentimentScore"`
	JSON              corporateComplianceScreenAdverseMediaResponseJSON `json:"-"`
}

// corporateComplianceScreenAdverseMediaResponseJSON contains the JSON metadata for
// the struct [CorporateComplianceScreenAdverseMediaResponse]
type corporateComplianceScreenAdverseMediaResponseJSON struct {
	NegativeNewsLinks apijson.Field
	SentimentScore    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CorporateComplianceScreenAdverseMediaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateComplianceScreenAdverseMediaResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateComplianceScreenPepResponse struct {
	Details string                                   `json:"details"`
	IsMatch bool                                     `json:"isMatch"`
	JSON    corporateComplianceScreenPepResponseJSON `json:"-"`
}

// corporateComplianceScreenPepResponseJSON contains the JSON metadata for the
// struct [CorporateComplianceScreenPepResponse]
type corporateComplianceScreenPepResponseJSON struct {
	Details     apijson.Field
	IsMatch     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateComplianceScreenPepResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateComplianceScreenPepResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateComplianceScreenSanctionsResponse struct {
	Hits      []interface{}                                       `json:"hits"`
	RiskLevel CorporateComplianceScreenSanctionsResponseRiskLevel `json:"riskLevel"`
	JSON      corporateComplianceScreenSanctionsResponseJSON      `json:"-"`
}

// corporateComplianceScreenSanctionsResponseJSON contains the JSON metadata for
// the struct [CorporateComplianceScreenSanctionsResponse]
type corporateComplianceScreenSanctionsResponseJSON struct {
	Hits        apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateComplianceScreenSanctionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateComplianceScreenSanctionsResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateComplianceScreenSanctionsResponseRiskLevel string

const (
	CorporateComplianceScreenSanctionsResponseRiskLevelLow      CorporateComplianceScreenSanctionsResponseRiskLevel = "Low"
	CorporateComplianceScreenSanctionsResponseRiskLevelMedium   CorporateComplianceScreenSanctionsResponseRiskLevel = "Medium"
	CorporateComplianceScreenSanctionsResponseRiskLevelHigh     CorporateComplianceScreenSanctionsResponseRiskLevel = "High"
	CorporateComplianceScreenSanctionsResponseRiskLevelCritical CorporateComplianceScreenSanctionsResponseRiskLevel = "Critical"
)

func (r CorporateComplianceScreenSanctionsResponseRiskLevel) IsKnown() bool {
	switch r {
	case CorporateComplianceScreenSanctionsResponseRiskLevelLow, CorporateComplianceScreenSanctionsResponseRiskLevelMedium, CorporateComplianceScreenSanctionsResponseRiskLevelHigh, CorporateComplianceScreenSanctionsResponseRiskLevelCritical:
		return true
	}
	return false
}

type CorporateComplianceScreenAdverseMediaParams struct {
	Query param.Field[string]                                           `json:"query,required"`
	Depth param.Field[CorporateComplianceScreenAdverseMediaParamsDepth] `json:"depth"`
}

func (r CorporateComplianceScreenAdverseMediaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateComplianceScreenAdverseMediaParamsDepth string

const (
	CorporateComplianceScreenAdverseMediaParamsDepthShallow CorporateComplianceScreenAdverseMediaParamsDepth = "shallow"
	CorporateComplianceScreenAdverseMediaParamsDepthDeep    CorporateComplianceScreenAdverseMediaParamsDepth = "deep"
)

func (r CorporateComplianceScreenAdverseMediaParamsDepth) IsKnown() bool {
	switch r {
	case CorporateComplianceScreenAdverseMediaParamsDepthShallow, CorporateComplianceScreenAdverseMediaParamsDepthDeep:
		return true
	}
	return false
}

type CorporateComplianceScreenPepParams struct {
	FullName param.Field[string]    `json:"fullName,required"`
	Dob      param.Field[time.Time] `json:"dob" format:"date"`
}

func (r CorporateComplianceScreenPepParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateComplianceScreenSanctionsParams struct {
	Entities  param.Field[[]CorporateComplianceScreenSanctionsParamsEntity]  `json:"entities,required"`
	CheckType param.Field[CorporateComplianceScreenSanctionsParamsCheckType] `json:"checkType"`
}

func (r CorporateComplianceScreenSanctionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateComplianceScreenSanctionsParamsEntity struct {
	Country param.Field[string] `json:"country"`
	Name    param.Field[string] `json:"name"`
}

func (r CorporateComplianceScreenSanctionsParamsEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateComplianceScreenSanctionsParamsCheckType string

const (
	CorporateComplianceScreenSanctionsParamsCheckTypeStandard             CorporateComplianceScreenSanctionsParamsCheckType = "standard"
	CorporateComplianceScreenSanctionsParamsCheckTypeEnhancedDueDiligence CorporateComplianceScreenSanctionsParamsCheckType = "enhanced_due_diligence"
)

func (r CorporateComplianceScreenSanctionsParamsCheckType) IsKnown() bool {
	switch r {
	case CorporateComplianceScreenSanctionsParamsCheckTypeStandard, CorporateComplianceScreenSanctionsParamsCheckTypeEnhancedDueDiligence:
		return true
	}
	return false
}
