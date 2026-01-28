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
	// Preferred channels for receiving notifications.
	NotificationChannels interface{}                     `json:"notificationChannels"`
	JSON                 userMePreferenceGetResponseJSON `json:"-"`
}

// userMePreferenceGetResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceGetResponse]
type userMePreferenceGetResponseJSON struct {
	NotificationChannels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMePreferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

// User's personalized preferences for the platform.
type UserMePreferenceUpdateResponse struct {
	// Preferred channels for receiving notifications.
	NotificationChannels interface{}                        `json:"notificationChannels"`
	JSON                 userMePreferenceUpdateResponseJSON `json:"-"`
}

// userMePreferenceUpdateResponseJSON contains the JSON metadata for the struct
// [UserMePreferenceUpdateResponse]
type userMePreferenceUpdateResponseJSON struct {
	NotificationChannels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMePreferenceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMePreferenceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserMePreferenceUpdateParams struct {
	// Preferred channels for receiving notifications.
	NotificationChannels param.Field[interface{}] `json:"notificationChannels"`
}

func (r UserMePreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
