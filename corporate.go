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

// CorporateService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateService] method instead.
type CorporateService struct {
	Options    []option.RequestOption
	Compliance *CorporateComplianceService
	Treasury   *CorporateTreasuryService
	Cards      *CorporateCardService
	Risk       *CorporateRiskService
	Governance *CorporateGovernanceService
	Anomalies  *CorporateAnomalyService
}

// NewCorporateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateService(opts ...option.RequestOption) (r *CorporateService) {
	r = &CorporateService{}
	r.Options = opts
	r.Compliance = NewCorporateComplianceService(opts...)
	r.Treasury = NewCorporateTreasuryService(opts...)
	r.Cards = NewCorporateCardService(opts...)
	r.Risk = NewCorporateRiskService(opts...)
	r.Governance = NewCorporateGovernanceService(opts...)
	r.Anomalies = NewCorporateAnomalyService(opts...)
	return
}

// Onboard a New Corporate Entity
func (r *CorporateService) OnboardEntity(ctx context.Context, body CorporateOnboardEntityParams, opts ...option.RequestOption) (res *CorporateOnboardEntityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/onboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporateOnboardEntityResponse struct {
	CorporateID string                             `json:"corporateId"`
	Status      string                             `json:"status"`
	JSON        corporateOnboardEntityResponseJSON `json:"-"`
}

// corporateOnboardEntityResponseJSON contains the JSON metadata for the struct
// [CorporateOnboardEntityResponse]
type corporateOnboardEntityResponseJSON struct {
	CorporateID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateOnboardEntityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateOnboardEntityResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateOnboardEntityParams struct {
	EntityType   param.Field[CorporateOnboardEntityParamsEntityType] `json:"entityType,required"`
	Jurisdiction param.Field[string]                                 `json:"jurisdiction,required"`
	// Registered business name
	LegalName param.Field[string] `json:"legalName,required"`
	// EIN, VAT, or local tax ID
	TaxID            param.Field[string]                                        `json:"taxId,required"`
	BeneficialOwners param.Field[[]CorporateOnboardEntityParamsBeneficialOwner] `json:"beneficialOwners"`
}

func (r CorporateOnboardEntityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateOnboardEntityParamsEntityType string

const (
	CorporateOnboardEntityParamsEntityTypeLlc         CorporateOnboardEntityParamsEntityType = "LLC"
	CorporateOnboardEntityParamsEntityTypeCorp        CorporateOnboardEntityParamsEntityType = "CORP"
	CorporateOnboardEntityParamsEntityTypeNgo         CorporateOnboardEntityParamsEntityType = "NGO"
	CorporateOnboardEntityParamsEntityTypePartnership CorporateOnboardEntityParamsEntityType = "PARTNERSHIP"
)

func (r CorporateOnboardEntityParamsEntityType) IsKnown() bool {
	switch r {
	case CorporateOnboardEntityParamsEntityTypeLlc, CorporateOnboardEntityParamsEntityTypeCorp, CorporateOnboardEntityParamsEntityTypeNgo, CorporateOnboardEntityParamsEntityTypePartnership:
		return true
	}
	return false
}

type CorporateOnboardEntityParamsBeneficialOwner struct {
	ID               param.Field[string]                                                     `json:"id,required"`
	Email            param.Field[string]                                                     `json:"email,required" format:"email"`
	IdentityVerified param.Field[bool]                                                       `json:"identityVerified,required"`
	Name             param.Field[string]                                                     `json:"name,required"`
	Address          param.Field[CorporateOnboardEntityParamsBeneficialOwnersAddress]        `json:"address"`
	Phone            param.Field[string]                                                     `json:"phone"`
	Preferences      param.Field[CorporateOnboardEntityParamsBeneficialOwnersPreferences]    `json:"preferences"`
	SecurityStatus   param.Field[CorporateOnboardEntityParamsBeneficialOwnersSecurityStatus] `json:"securityStatus"`
}

func (r CorporateOnboardEntityParamsBeneficialOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateOnboardEntityParamsBeneficialOwnersAddress struct {
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
	Zip     param.Field[string] `json:"zip"`
}

func (r CorporateOnboardEntityParamsBeneficialOwnersAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateOnboardEntityParamsBeneficialOwnersPreferences struct {
	NotificationChannels param.Field[interface{}] `json:"notificationChannels"`
	Theme                param.Field[string]      `json:"theme"`
}

func (r CorporateOnboardEntityParamsBeneficialOwnersPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateOnboardEntityParamsBeneficialOwnersSecurityStatus struct {
	LastLogin        param.Field[time.Time] `json:"lastLogin" format:"date-time"`
	TwoFactorEnabled param.Field[bool]      `json:"twoFactorEnabled"`
}

func (r CorporateOnboardEntityParamsBeneficialOwnersSecurityStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
