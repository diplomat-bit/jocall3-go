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

// AIModelService contains methods and other services that help with interacting
// with the jocall3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelService] method instead.
type AIModelService struct {
	Options []option.RequestOption
}

// NewAIModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIModelService(opts ...option.RequestOption) (r *AIModelService) {
	r = &AIModelService{}
	r.Options = opts
	return
}

// List supported AI model versions
func (r *AIModelService) List(ctx context.Context, opts ...option.RequestOption) (res *AIModelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/models/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Start a model fine-tuning job
func (r *AIModelService) FineTune(ctx context.Context, body AIModelFineTuneParams, opts ...option.RequestOption) (res *AIModelFineTuneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/models/fine-tune"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIModelListResponse struct {
	Models []AIModelListResponseModel `json:"models"`
	JSON   aiModelListResponseJSON    `json:"-"`
}

// aiModelListResponseJSON contains the JSON metadata for the struct
// [AIModelListResponse]
type aiModelListResponseJSON struct {
	Models      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiModelListResponseJSON) RawJSON() string {
	return r.raw
}

type AIModelListResponseModel struct {
	ModelID string                       `json:"modelId,required"`
	Version string                       `json:"version,required"`
	JSON    aiModelListResponseModelJSON `json:"-"`
}

// aiModelListResponseModelJSON contains the JSON metadata for the struct
// [AIModelListResponseModel]
type aiModelListResponseModelJSON struct {
	ModelID     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelListResponseModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiModelListResponseModelJSON) RawJSON() string {
	return r.raw
}

type AIModelFineTuneResponse struct {
	JobID string                      `json:"job_id"`
	JSON  aiModelFineTuneResponseJSON `json:"-"`
}

// aiModelFineTuneResponseJSON contains the JSON metadata for the struct
// [AIModelFineTuneResponse]
type aiModelFineTuneResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelFineTuneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiModelFineTuneResponseJSON) RawJSON() string {
	return r.raw
}

type AIModelFineTuneParams struct {
	BaseModel       param.Field[string]      `json:"base_model,required"`
	TrainingDataURL param.Field[string]      `json:"training_data_url,required" format:"uri"`
	Hyperparameters param.Field[interface{}] `json:"hyperparameters"`
}

func (r AIModelFineTuneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
