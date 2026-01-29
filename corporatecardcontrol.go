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

// CorporateCardControlService contains methods and other services that help with
// interacting with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateCardControlService] method instead.
type CorporateCardControlService struct {
	Options []option.RequestOption
}

// NewCorporateCardControlService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateCardControlService(opts ...option.RequestOption) (r *CorporateCardControlService) {
	r = &CorporateCardControlService{}
	r.Options = opts
	return
}

// Update Spending Limits & MCC Controls
func (r *CorporateCardControlService) Update(ctx context.Context, cardID string, body CorporateCardControlUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/controls", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CorporateCardControlUpdateParams struct {
	AllowedCategories param.Field[[]string] `json:"allowedCategories"`
	GeoRestriction    param.Field[[]string] `json:"geoRestriction"`
	MonthlyLimit      param.Field[float64]  `json:"monthlyLimit"`
}

func (r CorporateCardControlUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
