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

// UserMeService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeService] method instead.
type UserMeService struct {
	Options     []option.RequestOption
	Preferences *UserMePreferenceService
	Security    *UserMeSecurityService
	Devices     *UserMeDeviceService
	Biometrics  *UserMeBiometricService
}

// NewUserMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserMeService(opts ...option.RequestOption) (r *UserMeService) {
	r = &UserMeService{}
	r.Options = opts
	r.Preferences = NewUserMePreferenceService(opts...)
	r.Security = NewUserMeSecurityService(opts...)
	r.Devices = NewUserMeDeviceService(opts...)
	r.Biometrics = NewUserMeBiometricService(opts...)
	return
}

// Fetches the complete and dynamically updated profile information for the
// currently authenticated user, encompassing personal details, security status,
// gamification level, loyalty points, and linked identity attributes.
func (r *UserMeService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserMeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates selected fields of the currently authenticated user's profile
// information.
func (r *UserMeService) Update(ctx context.Context, body UserMeUpdateParams, opts ...option.RequestOption) (res *UserMeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UserMeGetResponse struct {
	ID                string                   `json:"id,required"`
	Email             string                   `json:"email,required" format:"email"`
	IdentityVerified  bool                     `json:"identityVerified,required"`
	Name              string                   `json:"name,required"`
	Address           UserMeGetResponseAddress `json:"address"`
	AIPersona         string                   `json:"aiPersona"`
	DateOfBirth       time.Time                `json:"dateOfBirth" format:"date"`
	GamificationLevel int64                    `json:"gamificationLevel"`
	LoyaltyPoints     int64                    `json:"loyaltyPoints"`
	LoyaltyTier       string                   `json:"loyaltyTier"`
	Phone             string                   `json:"phone"`
	// User's personalized preferences for the platform.
	Preferences UserMeGetResponsePreferences `json:"preferences"`
	// Security-related status for the user account.
	SecurityStatus UserMeGetResponseSecurityStatus `json:"securityStatus"`
	JSON           userMeGetResponseJSON           `json:"-"`
}

// userMeGetResponseJSON contains the JSON metadata for the struct
// [UserMeGetResponse]
type userMeGetResponseJSON struct {
	ID                apijson.Field
	Email             apijson.Field
	IdentityVerified  apijson.Field
	Name              apijson.Field
	Address           apijson.Field
	AIPersona         apijson.Field
	DateOfBirth       apijson.Field
	GamificationLevel apijson.Field
	LoyaltyPoints     apijson.Field
	LoyaltyTier       apijson.Field
	Phone             apijson.Field
	Preferences       apijson.Field
	SecurityStatus    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserMeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeGetResponseAddress struct {
	City    string                       `json:"city"`
	Country string                       `json:"country"`
	State   string                       `json:"state"`
	Street  string                       `json:"street"`
	Zip     string                       `json:"zip"`
	JSON    userMeGetResponseAddressJSON `json:"-"`
}

// userMeGetResponseAddressJSON contains the JSON metadata for the struct
// [UserMeGetResponseAddress]
type userMeGetResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeGetResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseAddressJSON) RawJSON() string {
	return r.raw
}

// User's personalized preferences for the platform.
type UserMeGetResponsePreferences struct {
	AIInteractionMode  string `json:"aiInteractionMode"`
	DataSharingConsent bool   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels UserMeGetResponsePreferencesNotificationChannels `json:"notificationChannels"`
	PreferredLanguage    string                                           `json:"preferredLanguage"`
	Theme                string                                           `json:"theme"`
	TransactionGrouping  string                                           `json:"transactionGrouping"`
	JSON                 userMeGetResponsePreferencesJSON                 `json:"-"`
}

// userMeGetResponsePreferencesJSON contains the JSON metadata for the struct
// [UserMeGetResponsePreferences]
type userMeGetResponsePreferencesJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMeGetResponsePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponsePreferencesJSON) RawJSON() string {
	return r.raw
}

// Preferred channels for receiving notifications.
type UserMeGetResponsePreferencesNotificationChannels struct {
	Email bool                                                 `json:"email"`
	InApp bool                                                 `json:"inApp"`
	Push  bool                                                 `json:"push"`
	SMS   bool                                                 `json:"sms"`
	JSON  userMeGetResponsePreferencesNotificationChannelsJSON `json:"-"`
}

// userMeGetResponsePreferencesNotificationChannelsJSON contains the JSON metadata
// for the struct [UserMeGetResponsePreferencesNotificationChannels]
type userMeGetResponsePreferencesNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeGetResponsePreferencesNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponsePreferencesNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

// Security-related status for the user account.
type UserMeGetResponseSecurityStatus struct {
	BiometricsEnrolled bool                                `json:"biometricsEnrolled"`
	LastLogin          time.Time                           `json:"lastLogin" format:"date-time"`
	LastLoginIP        string                              `json:"lastLoginIp"`
	TwoFactorEnabled   bool                                `json:"twoFactorEnabled"`
	JSON               userMeGetResponseSecurityStatusJSON `json:"-"`
}

// userMeGetResponseSecurityStatusJSON contains the JSON metadata for the struct
// [UserMeGetResponseSecurityStatus]
type userMeGetResponseSecurityStatusJSON struct {
	BiometricsEnrolled apijson.Field
	LastLogin          apijson.Field
	LastLoginIP        apijson.Field
	TwoFactorEnabled   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeGetResponseSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeGetResponseSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponse struct {
	ID                string                      `json:"id,required"`
	Email             string                      `json:"email,required" format:"email"`
	IdentityVerified  bool                        `json:"identityVerified,required"`
	Name              string                      `json:"name,required"`
	Address           UserMeUpdateResponseAddress `json:"address"`
	AIPersona         string                      `json:"aiPersona"`
	DateOfBirth       time.Time                   `json:"dateOfBirth" format:"date"`
	GamificationLevel int64                       `json:"gamificationLevel"`
	LoyaltyPoints     int64                       `json:"loyaltyPoints"`
	LoyaltyTier       string                      `json:"loyaltyTier"`
	Phone             string                      `json:"phone"`
	// User's personalized preferences for the platform.
	Preferences UserMeUpdateResponsePreferences `json:"preferences"`
	// Security-related status for the user account.
	SecurityStatus UserMeUpdateResponseSecurityStatus `json:"securityStatus"`
	JSON           userMeUpdateResponseJSON           `json:"-"`
}

// userMeUpdateResponseJSON contains the JSON metadata for the struct
// [UserMeUpdateResponse]
type userMeUpdateResponseJSON struct {
	ID                apijson.Field
	Email             apijson.Field
	IdentityVerified  apijson.Field
	Name              apijson.Field
	Address           apijson.Field
	AIPersona         apijson.Field
	DateOfBirth       apijson.Field
	GamificationLevel apijson.Field
	LoyaltyPoints     apijson.Field
	LoyaltyTier       apijson.Field
	Phone             apijson.Field
	Preferences       apijson.Field
	SecurityStatus    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserMeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateResponseAddress struct {
	City    string                          `json:"city"`
	Country string                          `json:"country"`
	State   string                          `json:"state"`
	Street  string                          `json:"street"`
	Zip     string                          `json:"zip"`
	JSON    userMeUpdateResponseAddressJSON `json:"-"`
}

// userMeUpdateResponseAddressJSON contains the JSON metadata for the struct
// [UserMeUpdateResponseAddress]
type userMeUpdateResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeUpdateResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseAddressJSON) RawJSON() string {
	return r.raw
}

// User's personalized preferences for the platform.
type UserMeUpdateResponsePreferences struct {
	AIInteractionMode  string `json:"aiInteractionMode"`
	DataSharingConsent bool   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels UserMeUpdateResponsePreferencesNotificationChannels `json:"notificationChannels"`
	PreferredLanguage    string                                              `json:"preferredLanguage"`
	Theme                string                                              `json:"theme"`
	TransactionGrouping  string                                              `json:"transactionGrouping"`
	JSON                 userMeUpdateResponsePreferencesJSON                 `json:"-"`
}

// userMeUpdateResponsePreferencesJSON contains the JSON metadata for the struct
// [UserMeUpdateResponsePreferences]
type userMeUpdateResponsePreferencesJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserMeUpdateResponsePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponsePreferencesJSON) RawJSON() string {
	return r.raw
}

// Preferred channels for receiving notifications.
type UserMeUpdateResponsePreferencesNotificationChannels struct {
	Email bool                                                    `json:"email"`
	InApp bool                                                    `json:"inApp"`
	Push  bool                                                    `json:"push"`
	SMS   bool                                                    `json:"sms"`
	JSON  userMeUpdateResponsePreferencesNotificationChannelsJSON `json:"-"`
}

// userMeUpdateResponsePreferencesNotificationChannelsJSON contains the JSON
// metadata for the struct [UserMeUpdateResponsePreferencesNotificationChannels]
type userMeUpdateResponsePreferencesNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeUpdateResponsePreferencesNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponsePreferencesNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

// Security-related status for the user account.
type UserMeUpdateResponseSecurityStatus struct {
	BiometricsEnrolled bool                                   `json:"biometricsEnrolled"`
	LastLogin          time.Time                              `json:"lastLogin" format:"date-time"`
	LastLoginIP        string                                 `json:"lastLoginIp"`
	TwoFactorEnabled   bool                                   `json:"twoFactorEnabled"`
	JSON               userMeUpdateResponseSecurityStatusJSON `json:"-"`
}

// userMeUpdateResponseSecurityStatusJSON contains the JSON metadata for the struct
// [UserMeUpdateResponseSecurityStatus]
type userMeUpdateResponseSecurityStatusJSON struct {
	BiometricsEnrolled apijson.Field
	LastLogin          apijson.Field
	LastLoginIP        apijson.Field
	TwoFactorEnabled   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeUpdateResponseSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeUpdateResponseSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserMeUpdateParams struct {
	Address param.Field[UserMeUpdateParamsAddress] `json:"address"`
	Name    param.Field[string]                    `json:"name"`
	Phone   param.Field[string]                    `json:"phone"`
	// User's personalized preferences for the platform.
	Preferences param.Field[UserMeUpdateParamsPreferences] `json:"preferences"`
}

func (r UserMeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeUpdateParamsAddress struct {
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
	Zip     param.Field[string] `json:"zip"`
}

func (r UserMeUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User's personalized preferences for the platform.
type UserMeUpdateParamsPreferences struct {
	AIInteractionMode  param.Field[string] `json:"aiInteractionMode"`
	DataSharingConsent param.Field[bool]   `json:"dataSharingConsent"`
	// Preferred channels for receiving notifications.
	NotificationChannels param.Field[UserMeUpdateParamsPreferencesNotificationChannels] `json:"notificationChannels"`
	PreferredLanguage    param.Field[string]                                            `json:"preferredLanguage"`
	Theme                param.Field[string]                                            `json:"theme"`
	TransactionGrouping  param.Field[string]                                            `json:"transactionGrouping"`
}

func (r UserMeUpdateParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Preferred channels for receiving notifications.
type UserMeUpdateParamsPreferencesNotificationChannels struct {
	Email param.Field[bool] `json:"email"`
	InApp param.Field[bool] `json:"inApp"`
	Push  param.Field[bool] `json:"push"`
	SMS   param.Field[bool] `json:"sms"`
}

func (r UserMeUpdateParamsPreferencesNotificationChannels) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
