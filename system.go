// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"github.com/diplomat-bit/jocall3-go/option"
)

// SystemService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemService] method instead.
type SystemService struct {
	Options       []option.RequestOption
	Status        *SystemStatusService
	Webhooks      *SystemWebhookService
	AuditLogs     *SystemAuditLogService
	Sandbox       *SystemSandboxService
	Verification  *SystemVerificationService
	Notifications *SystemNotificationService
}

// NewSystemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSystemService(opts ...option.RequestOption) (r *SystemService) {
	r = &SystemService{}
	r.Options = opts
	r.Status = NewSystemStatusService(opts...)
	r.Webhooks = NewSystemWebhookService(opts...)
	r.AuditLogs = NewSystemAuditLogService(opts...)
	r.Sandbox = NewSystemSandboxService(opts...)
	r.Verification = NewSystemVerificationService(opts...)
	r.Notifications = NewSystemNotificationService(opts...)
	return
}
