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

// Retrieves the user's deep personalization preferences, including AI
// customization settings, notification channel priorities, thematic choices, and
// data sharing consents.
func (r *UserMePreferenceService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserMePreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the user's deep personalization preferences, allowing dynamic control
// over AI behavior, notification delivery, thematic choices, and data privacy
// settings.
func (r *UserMePreferenceService) Update(ctx context.Context, body UserMePreferenceUpdateParams, opts ...option.RequestOption) (res *UserMePreferenceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// User's personalized preferences for the platform.
type UserMePreferenceGetResponse struct {
	AIInteractionMode  string `json:"aiInteractionMode"`
	DataSharingConsent bool   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels UserMePreferenceGetResponseNotificationChannels `json:"notificationChannels"`
	PreferredLanguage    string                                          `json:"preferredLanguage"`
	Theme                string                                          `json:"theme"`
	TransactionGrouping  string                                          `json:"transactionGrouping"`
	JSON                 userMePreferenceGetResponseJSON                 `json:"-"`
}

// userMePreferenceGetResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceGetResponse]
type userMePreferenceGetResponseJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMePreferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

// Preferred channels for receiving notifications.
type UserMePreferenceGetResponseNotificationChannels struct {
	Email bool                                                `json:"email"`
	InApp bool                                                `json:"inApp"`
	Push  bool                                                `json:"push"`
	SMS   bool                                                `json:"sms"`
	JSON  userMePreferenceGetResponseNotificationChannelsJSON `json:"-"`
}

// userMePreferenceGetResponseNotificationChannelsJSON contains the JSON metadata
// for the struct [UserMePreferenceGetResponseNotificationChannels]
type userMePreferenceGetResponseNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMePreferenceGetResponseNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceGetResponseNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

// User's personalized preferences for the platform.
type UserMePreferenceUpdateResponse struct {
	AIInteractionMode  string `json:"aiInteractionMode"`
	DataSharingConsent bool   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels UserMePreferenceUpdateResponseNotificationChannels `json:"notificationChannels"`
	PreferredLanguage    string                                             `json:"preferredLanguage"`
	Theme                string                                             `json:"theme"`
	TransactionGrouping  string                                             `json:"transactionGrouping"`
	JSON                 userMePreferenceUpdateResponseJSON                 `json:"-"`
}

// userMePreferenceUpdateResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceUpdateResponse]
type userMePreferenceUpdateResponseJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMePreferenceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Preferred channels for receiving notifications.
type UserMePreferenceUpdateResponseNotificationChannels struct {
	Email bool                                                   `json:"email"`
	InApp bool                                                   `json:"inApp"`
	Push  bool                                                   `json:"push"`
	SMS   bool                                                   `json:"sms"`
	JSON  userMePreferenceUpdateResponseNotificationChannelsJSON `json:"-"`
}

// userMePreferenceUpdateResponseNotificationChannelsJSON contains the JSON
// metadata for the struct [UserMePreferenceUpdateResponseNotificationChannels]
type userMePreferenceUpdateResponseNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMePreferenceUpdateResponseNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceUpdateResponseNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

type UserMePreferenceUpdateParams struct {
	AIInteractionMode  param.Field[string] `json:"aiInteractionMode"`
	DataSharingConsent param.Field[bool]   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels param.Field[UserMePreferenceUpdateParamsNotificationChannels] `json:"notificationChannels"`
	PreferredLanguage    param.Field[string]                                           `json:"preferredLanguage"`
	Theme                param.Field[string]                                           `json:"theme"`
	TransactionGrouping  param.Field[string]                                           `json:"transactionGrouping"`
}

func (r UserMePreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Preferred channels for receiving notifications.
type UserMePreferenceUpdateParamsNotificationChannels struct {
	Email param.Field[bool] `json:"email"`
	InApp param.Field[bool] `json:"inApp"`
	Push  param.Field[bool] `json:"push"`
	SMS   param.Field[bool] `json:"sms"`
}

func (r UserMePreferenceUpdateParamsNotificationChannels) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
