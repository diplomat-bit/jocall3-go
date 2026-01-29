// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/param"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// UserMePreferenceService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMePreferenceService] method instead.
type UserMePreferenceService struct {
	Options []option.RequestOption
}

// NewUserMePreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMePreferenceService(opts ...option.RequestOption) (r *UserMePreferenceService) {
	r = &UserMePreferenceService{}
	r.Options = opts
	return
}

// Get User Personalization Preferences
func (r *UserMePreferenceService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserMePreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update User Personalization Preferences
func (r *UserMePreferenceService) Update(ctx context.Context, body UserMePreferenceUpdateParams, opts ...option.RequestOption) (res *UserMePreferenceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UserMePreferenceGetResponse struct {
	AIInteractionMode  UserMePreferenceGetResponseAIInteractionMode `json:"aiInteractionMode"`
	DataSharingConsent bool                                         `json:"dataSharingConsent"`
	PreferredLanguage  string                                       `json:"preferredLanguage"`
	Theme              string                                       `json:"theme"`
	JSON               userMePreferenceGetResponseJSON              `json:"-"`
}

// userMePreferenceGetResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceGetResponse]
type userMePreferenceGetResponseJSON struct {
	AIInteractionMode  apijson.Field
	DataSharingConsent apijson.Field
	PreferredLanguage  apijson.Field
	Theme              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMePreferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserMePreferenceGetResponseAIInteractionMode string

const (
	UserMePreferenceGetResponseAIInteractionModeProactive UserMePreferenceGetResponseAIInteractionMode = "proactive"
	UserMePreferenceGetResponseAIInteractionModeReactive  UserMePreferenceGetResponseAIInteractionMode = "reactive"
	UserMePreferenceGetResponseAIInteractionModeSilent    UserMePreferenceGetResponseAIInteractionMode = "silent"
)

func (r UserMePreferenceGetResponseAIInteractionMode) IsKnown() bool {
	switch r {
	case UserMePreferenceGetResponseAIInteractionModeProactive, UserMePreferenceGetResponseAIInteractionModeReactive, UserMePreferenceGetResponseAIInteractionModeSilent:
		return true
	}
	return false
}

type UserMePreferenceUpdateResponse struct {
	AIInteractionMode string                             `json:"aiInteractionMode"`
	Theme             string                             `json:"theme"`
	JSON              userMePreferenceUpdateResponseJSON `json:"-"`
}

// userMePreferenceUpdateResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceUpdateResponse]
type userMePreferenceUpdateResponseJSON struct {
	AIInteractionMode apijson.Field
	Theme             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserMePreferenceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserMePreferenceUpdateParams struct {
	AIInteractionMode param.Field[string] `json:"aiInteractionMode"`
	Theme             param.Field[string] `json:"theme"`
}

func (r UserMePreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
