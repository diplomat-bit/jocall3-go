// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomjocall3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/apiquery"
	"github.com/diplomat-bit/jocall3-go/internal/param"
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

// Fetches digital statements for a specific account, allowing filtering by date
// range and format.
func (r *AccountStatementService) List(ctx context.Context, accountID string, query AccountStatementListParams, opts ...option.RequestOption) (res *AccountStatementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/statements", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStatementListResponse struct {
	AccountID    string                                   `json:"accountId,required"`
	DownloadURLs AccountStatementListResponseDownloadURLs `json:"downloadUrls,required"`
	Period       string                                   `json:"period,required"`
	StatementID  string                                   `json:"statementId,required"`
	JSON         accountStatementListResponseJSON         `json:"-"`
}

// accountStatementListResponseJSON contains the JSON metadata for the struct
// [AccountStatementListResponse]
type accountStatementListResponseJSON struct {
	AccountID    apijson.Field
	DownloadURLs apijson.Field
	Period       apijson.Field
	StatementID  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountStatementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStatementListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStatementListResponseDownloadURLs struct {
	Csv  string                                       `json:"csv"`
	Pdf  string                                       `json:"pdf"`
	JSON accountStatementListResponseDownloadURLsJSON `json:"-"`
}

// accountStatementListResponseDownloadURLsJSON contains the JSON metadata for the
// struct [AccountStatementListResponseDownloadURLs]
type accountStatementListResponseDownloadURLsJSON struct {
	Csv         apijson.Field
	Pdf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStatementListResponseDownloadURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStatementListResponseDownloadURLsJSON) RawJSON() string {
	return r.raw
}

type AccountStatementListParams struct {
	// Desired format for the statement. Use 'application/json' Accept header for
	// download links.
	Format param.Field[string] `query:"format"`
	// Month for the statement (1-12).
	Month param.Field[int64] `query:"month"`
	// Year for the statement.
	Year param.Field[int64] `query:"year"`
}

// URLQuery serializes [AccountStatementListParams]'s query parameters as
// `url.Values`.
func (r AccountStatementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
