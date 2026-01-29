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

// SystemNotificationService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemNotificationService] method instead.
type SystemNotificationService struct {
	Options []option.RequestOption
}

// NewSystemNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSystemNotificationService(opts ...option.RequestOption) (r *SystemNotificationService) {
	r = &SystemNotificationService{}
	r.Options = opts
	return
}

// List notification templates
func (r *SystemNotificationService) ListTemplates(ctx context.Context, opts ...option.RequestOption) (res *[]SystemNotificationListTemplatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "system/notifications/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Send push notification
func (r *SystemNotificationService) SendPush(ctx context.Context, body SystemNotificationSendPushParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "system/notifications/push"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SystemNotificationListTemplatesResponse struct {
	ID      string                                      `json:"id,required"`
	Channel string                                      `json:"channel,required"`
	Name    string                                      `json:"name,required"`
	JSON    systemNotificationListTemplatesResponseJSON `json:"-"`
}

// systemNotificationListTemplatesResponseJSON contains the JSON metadata for the
// struct [SystemNotificationListTemplatesResponse]
type systemNotificationListTemplatesResponseJSON struct {
	ID          apijson.Field
	Channel     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemNotificationListTemplatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemNotificationListTemplatesResponseJSON) RawJSON() string {
	return r.raw
}

type SystemNotificationSendPushParams struct {
	Body   param.Field[string] `json:"body,required"`
	Title  param.Field[string] `json:"title,required"`
	UserID param.Field[string] `json:"user_id,required"`
}

func (r SystemNotificationSendPushParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
