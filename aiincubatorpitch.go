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

// AIIncubatorPitchService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorPitchService] method instead.
type AIIncubatorPitchService struct {
	Options []option.RequestOption
}

// NewAIIncubatorPitchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIncubatorPitchService(opts ...option.RequestOption) (r *AIIncubatorPitchService) {
	r = &AIIncubatorPitchService{}
	r.Options = opts
	return
}

// Get Full Pitch AI Deep Dive
func (r *AIIncubatorPitchService) GetDetails(ctx context.Context, pitchID string, opts ...option.RequestOption) (res *AIIncubatorPitchGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/details", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submit Answers to AI Follow-up Questions
func (r *AIIncubatorPitchService) SubmitFeedback(ctx context.Context, pitchID string, body AIIncubatorPitchSubmitFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/feedback", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type AIIncubatorPitchGetDetailsResponse struct {
	AIFeedback         string                                 `json:"aiFeedback"`
	FundingEligibility bool                                   `json:"fundingEligibility"`
	JSON               aiIncubatorPitchGetDetailsResponseJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseJSON contains the JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponse]
type aiIncubatorPitchGetDetailsResponseJSON struct {
	AIFeedback         apijson.Field
	FundingEligibility apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchSubmitFeedbackParams struct {
	Answers param.Field[[]interface{}] `json:"answers,required"`
}

func (r AIIncubatorPitchSubmitFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
