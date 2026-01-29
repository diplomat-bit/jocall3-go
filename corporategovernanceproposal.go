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

// CorporateGovernanceProposalService contains methods and other services that help
// with interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateGovernanceProposalService] method instead.
type CorporateGovernanceProposalService struct {
	Options []option.RequestOption
}

// NewCorporateGovernanceProposalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCorporateGovernanceProposalService(opts ...option.RequestOption) (r *CorporateGovernanceProposalService) {
	r = &CorporateGovernanceProposalService{}
	r.Options = opts
	return
}

// Create New Multi-sig Financial Proposal
func (r *CorporateGovernanceProposalService) New(ctx context.Context, body CorporateGovernanceProposalNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "corporate/governance/proposals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List Active Governance Proposals
func (r *CorporateGovernanceProposalService) List(ctx context.Context, opts ...option.RequestOption) (res *CorporateGovernanceProposalListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/governance/proposals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cast Vote or Sign Transaction
func (r *CorporateGovernanceProposalService) Vote(ctx context.Context, proposalID string, body CorporateGovernanceProposalVoteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if proposalID == "" {
		err = errors.New("missing required proposalId parameter")
		return
	}
	path := fmt.Sprintf("corporate/governance/proposals/%s/vote", proposalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CorporateGovernanceProposalListResponse struct {
	Data []CorporateGovernanceProposalListResponseData `json:"data,required"`
	JSON corporateGovernanceProposalListResponseJSON   `json:"-"`
}

// corporateGovernanceProposalListResponseJSON contains the JSON metadata for the
// struct [CorporateGovernanceProposalListResponse]
type corporateGovernanceProposalListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateGovernanceProposalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateGovernanceProposalListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateGovernanceProposalListResponseData struct {
	ID                string                                            `json:"id,required"`
	RequiredApprovals int64                                             `json:"requiredApprovals,required"`
	Status            CorporateGovernanceProposalListResponseDataStatus `json:"status,required"`
	Title             string                                            `json:"title,required"`
	CurrentApprovals  int64                                             `json:"currentApprovals"`
	Description       string                                            `json:"description"`
	ExpiresAt         time.Time                                         `json:"expiresAt" format:"date-time"`
	JSON              corporateGovernanceProposalListResponseDataJSON   `json:"-"`
}

// corporateGovernanceProposalListResponseDataJSON contains the JSON metadata for
// the struct [CorporateGovernanceProposalListResponseData]
type corporateGovernanceProposalListResponseDataJSON struct {
	ID                apijson.Field
	RequiredApprovals apijson.Field
	Status            apijson.Field
	Title             apijson.Field
	CurrentApprovals  apijson.Field
	Description       apijson.Field
	ExpiresAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CorporateGovernanceProposalListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateGovernanceProposalListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CorporateGovernanceProposalListResponseDataStatus string

const (
	CorporateGovernanceProposalListResponseDataStatusPending  CorporateGovernanceProposalListResponseDataStatus = "PENDING"
	CorporateGovernanceProposalListResponseDataStatusApproved CorporateGovernanceProposalListResponseDataStatus = "APPROVED"
	CorporateGovernanceProposalListResponseDataStatusExecuted CorporateGovernanceProposalListResponseDataStatus = "EXECUTED"
	CorporateGovernanceProposalListResponseDataStatusRejected CorporateGovernanceProposalListResponseDataStatus = "REJECTED"
)

func (r CorporateGovernanceProposalListResponseDataStatus) IsKnown() bool {
	switch r {
	case CorporateGovernanceProposalListResponseDataStatusPending, CorporateGovernanceProposalListResponseDataStatusApproved, CorporateGovernanceProposalListResponseDataStatusExecuted, CorporateGovernanceProposalListResponseDataStatusRejected:
		return true
	}
	return false
}

type CorporateGovernanceProposalNewParams struct {
	ActionType param.Field[CorporateGovernanceProposalNewParamsActionType] `json:"actionType,required"`
	// The raw action data to be executed upon approval
	Payload           param.Field[interface{}] `json:"payload,required"`
	Title             param.Field[string]      `json:"title,required"`
	Description       param.Field[string]      `json:"description"`
	VotingPeriodHours param.Field[int64]       `json:"votingPeriodHours"`
}

func (r CorporateGovernanceProposalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateGovernanceProposalNewParamsActionType string

const (
	CorporateGovernanceProposalNewParamsActionTypeTransferLimitChange CorporateGovernanceProposalNewParamsActionType = "TRANSFER_LIMIT_CHANGE"
	CorporateGovernanceProposalNewParamsActionTypeNewAdmin            CorporateGovernanceProposalNewParamsActionType = "NEW_ADMIN"
	CorporateGovernanceProposalNewParamsActionTypeLargePayment        CorporateGovernanceProposalNewParamsActionType = "LARGE_PAYMENT"
)

func (r CorporateGovernanceProposalNewParamsActionType) IsKnown() bool {
	switch r {
	case CorporateGovernanceProposalNewParamsActionTypeTransferLimitChange, CorporateGovernanceProposalNewParamsActionTypeNewAdmin, CorporateGovernanceProposalNewParamsActionTypeLargePayment:
		return true
	}
	return false
}

type CorporateGovernanceProposalVoteParams struct {
	Decision param.Field[CorporateGovernanceProposalVoteParamsDecision] `json:"decision,required"`
	Comment  param.Field[string]                                        `json:"comment"`
	// Cryptographic signature if required
	Signature param.Field[string] `json:"signature"`
}

func (r CorporateGovernanceProposalVoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateGovernanceProposalVoteParamsDecision string

const (
	CorporateGovernanceProposalVoteParamsDecisionApprove CorporateGovernanceProposalVoteParamsDecision = "APPROVE"
	CorporateGovernanceProposalVoteParamsDecisionReject  CorporateGovernanceProposalVoteParamsDecision = "REJECT"
)

func (r CorporateGovernanceProposalVoteParamsDecision) IsKnown() bool {
	switch r {
	case CorporateGovernanceProposalVoteParamsDecisionApprove, CorporateGovernanceProposalVoteParamsDecisionReject:
		return true
	}
	return false
}
