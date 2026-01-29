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
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

// AccountStatementService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStatementService] method instead.
type AccountStatementService struct {
	Options []option.RequestOption
}

// NewAccountStatementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStatementService(opts ...option.RequestOption) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Options = opts
	return
}

// List Available Statements
func (r *AccountStatementService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountStatementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/statements", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Download Statement PDF
func (r *AccountStatementService) Download(ctx context.Context, accountID string, statementID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if statementID == "" {
		err = errors.New("missing required statementId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/statements/%s/pdf", accountID, statementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStatementListResponse struct {
	Data []AccountStatementListResponseData `json:"data"`
	JSON accountStatementListResponseJSON   `json:"-"`
}

// accountStatementListResponseJSON contains the JSON metadata for the struct
// [AccountStatementListResponse]
type accountStatementListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStatementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStatementListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStatementListResponseData struct {
	ID        string                               `json:"id"`
	IssueDate time.Time                            `json:"issueDate" format:"date"`
	Period    string                               `json:"period"`
	JSON      accountStatementListResponseDataJSON `json:"-"`
}

// accountStatementListResponseDataJSON contains the JSON metadata for the struct
// [AccountStatementListResponseData]
type accountStatementListResponseDataJSON struct {
	ID          apijson.Field
	IssueDate   apijson.Field
	Period      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStatementListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStatementListResponseDataJSON) RawJSON() string {
	return r.raw
}
