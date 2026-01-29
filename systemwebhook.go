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

// SystemWebhookService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemWebhookService] method instead.
type SystemWebhookService struct {
	Options []option.RequestOption
}

// NewSystemWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemWebhookService(opts ...option.RequestOption) (r *SystemWebhookService) {
	r = &SystemWebhookService{}
	r.Options = opts
	return
}

// Register Real-time Event Hook
func (r *SystemWebhookService) New(ctx context.Context, body SystemWebhookNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "system/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List Registered Webhooks
func (r *SystemWebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *SystemWebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "system/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Webhook
func (r *SystemWebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("system/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SystemWebhookListResponse struct {
	Data []interface{}                 `json:"data"`
	JSON systemWebhookListResponseJSON `json:"-"`
}

// systemWebhookListResponseJSON contains the JSON metadata for the struct
// [SystemWebhookListResponse]
type systemWebhookListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

type SystemWebhookNewParams struct {
	Events param.Field[[]string] `json:"events,required"`
	URL    param.Field[string]   `json:"url,required" format:"uri"`
	// HMAC signing secret
	Secret param.Field[string] `json:"secret"`
}

func (r SystemWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
