// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// CorporateComplianceAuditService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateComplianceAuditService] method instead.
type CorporateComplianceAuditService struct {
	Options []option.RequestOption
}

// NewCorporateComplianceAuditService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateComplianceAuditService(opts ...option.RequestOption) (r *CorporateComplianceAuditService) {
	r = &CorporateComplianceAuditService{}
	r.Options = opts
	return
}

// Request Real-time Compliance Audit
func (r *CorporateComplianceAuditService) Request(ctx context.Context, body CorporateComplianceAuditRequestParams, opts ...option.RequestOption) (res *CorporateComplianceAuditRequestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/compliance/audits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve AI-Generated Audit Report
func (r *CorporateComplianceAuditService) GetReport(ctx context.Context, auditID string, opts ...option.RequestOption) (res *CorporateComplianceAuditGetReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if auditID == "" {
		err = errors.New("missing required auditId parameter")
		return
	}
	path := fmt.Sprintf("corporate/compliance/audits/%s/report", auditID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CorporateComplianceAuditRequestResponse struct {
	AuditID string                                      `json:"auditId"`
	JSON    corporateComplianceAuditRequestResponseJSON `json:"-"`
}

// corporateComplianceAuditRequestResponseJSON contains the JSON metadata for the
// struct [CorporateComplianceAuditRequestResponse]
type corporateComplianceAuditRequestResponseJSON struct {
	AuditID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateComplianceAuditRequestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateComplianceAuditRequestResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateComplianceAuditGetReportResponse struct {
	AuditID                string                                        `json:"auditId,required"`
	OverallComplianceScore int64                                         `json:"overallComplianceScore,required"`
	Status                 string                                        `json:"status,required"`
	AuditDate              time.Time                                     `json:"auditDate" format:"date-time"`
	Findings               []interface{}                                 `json:"findings"`
	JSON                   corporateComplianceAuditGetReportResponseJSON `json:"-"`
}

// corporateComplianceAuditGetReportResponseJSON contains the JSON metadata for the
// struct [CorporateComplianceAuditGetReportResponse]
type corporateComplianceAuditGetReportResponseJSON struct {
	AuditID                apijson.Field
	OverallComplianceScore apijson.Field
	Status                 apijson.Field
	AuditDate              apijson.Field
	Findings               apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CorporateComplianceAuditGetReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateComplianceAuditGetReportResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateComplianceAuditRequestParams struct {
	AuditScope param.Field[string]    `json:"auditScope,required"`
	EndDate    param.Field[time.Time] `json:"endDate,required" format:"date"`
	StartDate  param.Field[time.Time] `json:"startDate,required" format:"date"`
}

func (r CorporateComplianceAuditRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
