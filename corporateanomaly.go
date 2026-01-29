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

// List detected anomalies
func (r *CorporateAnomalyService) List(ctx context.Context, opts ...option.RequestOption) (res *CorporateAnomalyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update anomaly status
func (r *CorporateAnomalyService) UpdateStatus(ctx context.Context, anomalyID string, body CorporateAnomalyUpdateStatusParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if anomalyID == "" {
		err = errors.New("missing required anomalyId parameter")
		return
	}
	path := fmt.Sprintf("corporate/anomalies/%s/status", anomalyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CorporateAnomalyListResponse struct {
	Data []CorporateAnomalyListResponseData `json:"data"`
	JSON corporateAnomalyListResponseJSON   `json:"-"`
}

// corporateAnomalyListResponseJSON contains the JSON metadata for the struct
// [CorporateAnomalyListResponse]
type corporateAnomalyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateAnomalyListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateAnomalyListResponseData struct {
	ID          string                                   `json:"id,required"`
	Severity    CorporateAnomalyListResponseDataSeverity `json:"severity,required"`
	Type        string                                   `json:"type,required"`
	Description string                                   `json:"description"`
	DetectedAt  time.Time                                `json:"detectedAt" format:"date-time"`
	JSON        corporateAnomalyListResponseDataJSON     `json:"-"`
}

// corporateAnomalyListResponseDataJSON contains the JSON metadata for the struct
// [CorporateAnomalyListResponseData]
type corporateAnomalyListResponseDataJSON struct {
	ID          apijson.Field
	Severity    apijson.Field
	Type        apijson.Field
	Description apijson.Field
	DetectedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateAnomalyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateAnomalyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CorporateAnomalyListResponseDataSeverity string

const (
	CorporateAnomalyListResponseDataSeverityLow      CorporateAnomalyListResponseDataSeverity = "low"
	CorporateAnomalyListResponseDataSeverityMedium   CorporateAnomalyListResponseDataSeverity = "medium"
	CorporateAnomalyListResponseDataSeverityHigh     CorporateAnomalyListResponseDataSeverity = "high"
	CorporateAnomalyListResponseDataSeverityCritical CorporateAnomalyListResponseDataSeverity = "critical"
)

func (r CorporateAnomalyListResponseDataSeverity) IsKnown() bool {
	switch r {
	case CorporateAnomalyListResponseDataSeverityLow, CorporateAnomalyListResponseDataSeverityMedium, CorporateAnomalyListResponseDataSeverityHigh, CorporateAnomalyListResponseDataSeverityCritical:
		return true
	}
	return false
}

type CorporateAnomalyUpdateStatusParams struct {
	Status param.Field[CorporateAnomalyUpdateStatusParamsStatus] `json:"status,required"`
}

func (r CorporateAnomalyUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateAnomalyUpdateStatusParamsStatus string

const (
	CorporateAnomalyUpdateStatusParamsStatusDismissed     CorporateAnomalyUpdateStatusParamsStatus = "dismissed"
	CorporateAnomalyUpdateStatusParamsStatusInvestigating CorporateAnomalyUpdateStatusParamsStatus = "investigating"
	CorporateAnomalyUpdateStatusParamsStatusResolved      CorporateAnomalyUpdateStatusParamsStatus = "resolved"
)

func (r CorporateAnomalyUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case CorporateAnomalyUpdateStatusParamsStatusDismissed, CorporateAnomalyUpdateStatusParamsStatusInvestigating, CorporateAnomalyUpdateStatusParamsStatusResolved:
		return true
	}
	return false
}
