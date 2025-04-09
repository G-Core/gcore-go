// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// QuotaGlobalService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaGlobalService] method instead.
type QuotaGlobalService struct {
	Options []option.RequestOption
}

// NewQuotaGlobalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuotaGlobalService(opts ...option.RequestOption) (r QuotaGlobalService) {
	r = QuotaGlobalService{}
	r.Options = opts
	return
}

// Get global quota
func (r *QuotaGlobalService) Get(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *GlobalQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/global_quotas/%v", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GlobalQuotas struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage int64 `json:"inference_cpu_millicore_count_usage"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage int64 `json:"inference_gpu_a100_count_usage"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage int64 `json:"inference_gpu_h100_count_usage"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage int64 `json:"inference_gpu_l40s_count_usage"`
	// Inference instance count limit
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// Inference instance count usage
	InferenceInstanceCountUsage int64 `json:"inference_instance_count_usage"`
	// SSH Keys Count limit
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// SSH Keys Count usage
	KeypairCountUsage int64 `json:"keypair_count_usage"`
	// Projects Count limit
	ProjectCountLimit int64 `json:"project_count_limit"`
	// Projects Count usage
	ProjectCountUsage int64 `json:"project_count_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InferenceCPUMillicoreCountLimit resp.Field
		InferenceCPUMillicoreCountUsage resp.Field
		InferenceGPUA100CountLimit      resp.Field
		InferenceGPUA100CountUsage      resp.Field
		InferenceGPUH100CountLimit      resp.Field
		InferenceGPUH100CountUsage      resp.Field
		InferenceGPUL40sCountLimit      resp.Field
		InferenceGPUL40sCountUsage      resp.Field
		InferenceInstanceCountLimit     resp.Field
		InferenceInstanceCountUsage     resp.Field
		KeypairCountLimit               resp.Field
		KeypairCountUsage               resp.Field
		ProjectCountLimit               resp.Field
		ProjectCountUsage               resp.Field
		ExtraFields                     map[string]resp.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalQuotas) RawJSON() string { return r.JSON.raw }
func (r *GlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
