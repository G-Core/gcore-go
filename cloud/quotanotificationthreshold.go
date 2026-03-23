// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// QuotaNotificationThresholdService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaNotificationThresholdService] method instead.
type QuotaNotificationThresholdService struct {
	Options []option.RequestOption
}

// NewQuotaNotificationThresholdService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewQuotaNotificationThresholdService(opts ...option.RequestOption) (r QuotaNotificationThresholdService) {
	r = QuotaNotificationThresholdService{}
	r.Options = opts
	return
}

// Update or create a client's quota notification threshold. This threshold is used
// to send warning notifications to the client when their quota usage reaches the
// specified percentage. Defaults to 80%.
func (r *QuotaNotificationThresholdService) Update(ctx context.Context, clientID int64, body QuotaNotificationThresholdUpdateParams, opts ...option.RequestOption) (res *NotificationThreshold, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Delete a client's quota notification threshold. After deletion, the default
// threshold of 80% will be used.
func (r *QuotaNotificationThresholdService) Delete(ctx context.Context, clientID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a client's quota notification threshold. This threshold is used to send
// warning notifications to the client when their quota usage reaches the specified
// percentage. Defaults to 80% if not set.
func (r *QuotaNotificationThresholdService) Get(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *NotificationThreshold, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type NotificationThreshold struct {
	// Client id
	ClientID int64 `json:"client_id" api:"required"`
	// A message data
	LastMessage NotificationThresholdLastMessage `json:"last_message" api:"required"`
	// Time of last successful email sending
	LastSending time.Time `json:"last_sending" api:"required" format:"date-time"`
	// Quota notification threshold in percentage
	Threshold int64 `json:"threshold" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		LastMessage respjson.Field
		LastSending respjson.Field
		Threshold   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThreshold) RawJSON() string { return r.JSON.raw }
func (r *NotificationThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message data
type NotificationThresholdLastMessage struct {
	// Global quota that exceed the threshold
	GlobalQuotas NotificationThresholdLastMessageGlobalQuotas `json:"global_quotas" api:"required"`
	// Regional quota that exceed the threshold
	RegionalQuotas []NotificationThresholdLastMessageRegionalQuota `json:"regional_quotas" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalQuotas   respjson.Field
		RegionalQuotas respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessage) RawJSON() string { return r.JSON.raw }
func (r *NotificationThresholdLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global quota that exceed the threshold
type NotificationThresholdLastMessageGlobalQuotas struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit `json:"inference_cpu_millicore_count_limit"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage `json:"inference_cpu_millicore_count_usage"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountLimit `json:"inference_gpu_a100_count_limit"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountUsage `json:"inference_gpu_a100_count_usage"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountLimit `json:"inference_gpu_h100_count_limit"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountUsage `json:"inference_gpu_h100_count_usage"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountLimit `json:"inference_gpu_l40s_count_limit"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountUsage `json:"inference_gpu_l40s_count_usage"`
	// Inference instance count limit
	InferenceInstanceCountLimit NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountLimit `json:"inference_instance_count_limit"`
	// Inference instance count usage
	InferenceInstanceCountUsage NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountUsage `json:"inference_instance_count_usage"`
	// SSH Keys Count limit
	KeypairCountLimit NotificationThresholdLastMessageGlobalQuotasKeypairCountLimit `json:"keypair_count_limit"`
	// SSH Keys Count usage
	KeypairCountUsage NotificationThresholdLastMessageGlobalQuotasKeypairCountUsage `json:"keypair_count_usage"`
	// Projects Count limit
	ProjectCountLimit NotificationThresholdLastMessageGlobalQuotasProjectCountLimit `json:"project_count_limit"`
	// Projects Count usage
	ProjectCountUsage NotificationThresholdLastMessageGlobalQuotasProjectCountUsage `json:"project_count_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceCPUMillicoreCountLimit respjson.Field
		InferenceCPUMillicoreCountUsage respjson.Field
		InferenceGPUA100CountLimit      respjson.Field
		InferenceGPUA100CountUsage      respjson.Field
		InferenceGPUH100CountLimit      respjson.Field
		InferenceGPUH100CountUsage      respjson.Field
		InferenceGPUL40sCountLimit      respjson.Field
		InferenceGPUL40sCountUsage      respjson.Field
		InferenceInstanceCountLimit     respjson.Field
		InferenceInstanceCountUsage     respjson.Field
		KeypairCountLimit               respjson.Field
		KeypairCountUsage               respjson.Field
		ProjectCountLimit               respjson.Field
		ProjectCountUsage               respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotas) RawJSON() string { return r.JSON.raw }
func (r *NotificationThresholdLastMessageGlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference CPU millicore count limit
type NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference CPU millicore count usage
type NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU A100 Count limit
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU A100 Count usage
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU H100 Count limit
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU H100 Count usage
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU L40s Count limit
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU L40s Count usage
type NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceGPUL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference instance count limit
type NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference instance count usage
type NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasInferenceInstanceCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH Keys Count limit
type NotificationThresholdLastMessageGlobalQuotasKeypairCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasKeypairCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasKeypairCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH Keys Count usage
type NotificationThresholdLastMessageGlobalQuotasKeypairCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasKeypairCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasKeypairCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Projects Count limit
type NotificationThresholdLastMessageGlobalQuotasProjectCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasProjectCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasProjectCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Projects Count usage
type NotificationThresholdLastMessageGlobalQuotasProjectCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageGlobalQuotasProjectCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageGlobalQuotasProjectCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationThresholdLastMessageRegionalQuota struct {
	// Region id
	RegionID int64 `json:"region_id" api:"required"`
	// Region name
	RegionName string `json:"region_name" api:"required"`
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountLimit `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountUsage `json:"baremetal_basic_count_usage"`
	// Bare metal A100 GPU server count limit
	BaremetalGPUA100CountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountLimit `json:"baremetal_gpu_a100_count_limit"`
	// Bare metal A100 GPU server count usage
	BaremetalGPUA100CountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountUsage `json:"baremetal_gpu_a100_count_usage"`
	// Bare metal H100 GPU server count limit
	BaremetalGPUH100CountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountLimit `json:"baremetal_gpu_h100_count_limit"`
	// Bare metal H100 GPU server count usage
	BaremetalGPUH100CountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountUsage `json:"baremetal_gpu_h100_count_usage"`
	// Bare metal H200 GPU server count limit
	BaremetalGPUH200CountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountLimit `json:"baremetal_gpu_h200_count_limit"`
	// Bare metal H200 GPU server count usage
	BaremetalGPUH200CountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountUsage `json:"baremetal_gpu_h200_count_usage"`
	// Bare metal L40S GPU server count limit
	BaremetalGPUL40sCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountLimit `json:"baremetal_gpu_l40s_count_limit"`
	// Bare metal L40S GPU server count usage
	BaremetalGPUL40sCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountUsage `json:"baremetal_gpu_l40s_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountLimit `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountUsage `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountLimit `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountUsage `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountLimit `json:"baremetal_network_count_limit"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountUsage `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountLimit `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountUsage `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit NotificationThresholdLastMessageRegionalQuotaCaasContainerCountLimit `json:"caas_container_count_limit"`
	// Containers count usage
	CaasContainerCountUsage NotificationThresholdLastMessageRegionalQuotaCaasContainerCountUsage `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit NotificationThresholdLastMessageRegionalQuotaCaasCPUCountLimit `json:"caas_cpu_count_limit"`
	// mCPU count for containers usage
	CaasCPUCountUsage NotificationThresholdLastMessageRegionalQuotaCaasCPUCountUsage `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit NotificationThresholdLastMessageRegionalQuotaCaasGPUCountLimit `json:"caas_gpu_count_limit"`
	// Containers gpu count usage
	CaasGPUCountUsage NotificationThresholdLastMessageRegionalQuotaCaasGPUCountUsage `json:"caas_gpu_count_usage"`
	// MiB memory count for containers limit
	CaasRamSizeLimit NotificationThresholdLastMessageRegionalQuotaCaasRamSizeLimit `json:"caas_ram_size_limit"`
	// MiB memory count for containers usage
	CaasRamSizeUsage NotificationThresholdLastMessageRegionalQuotaCaasRamSizeUsage `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit NotificationThresholdLastMessageRegionalQuotaClusterCountLimit `json:"cluster_count_limit"`
	// K8s clusters count usage
	ClusterCountUsage NotificationThresholdLastMessageRegionalQuotaClusterCountUsage `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit NotificationThresholdLastMessageRegionalQuotaCPUCountLimit `json:"cpu_count_limit"`
	// vCPU Count usage
	CPUCountUsage NotificationThresholdLastMessageRegionalQuotaCPUCountUsage `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountLimit `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountUsage `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit NotificationThresholdLastMessageRegionalQuotaExternalIPCountLimit `json:"external_ip_count_limit"`
	// External IP Count usage
	ExternalIPCountUsage NotificationThresholdLastMessageRegionalQuotaExternalIPCountUsage `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit NotificationThresholdLastMessageRegionalQuotaFaasCPUCountLimit `json:"faas_cpu_count_limit"`
	// mCPU count for functions usage
	FaasCPUCountUsage NotificationThresholdLastMessageRegionalQuotaFaasCPUCountUsage `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountLimit `json:"faas_function_count_limit"`
	// Functions count usage
	FaasFunctionCountUsage NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountUsage `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountLimit `json:"faas_namespace_count_limit"`
	// Functions namespace count usage
	FaasNamespaceCountUsage NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountUsage `json:"faas_namespace_count_usage"`
	// MiB memory count for functions limit
	FaasRamSizeLimit NotificationThresholdLastMessageRegionalQuotaFaasRamSizeLimit `json:"faas_ram_size_limit"`
	// MiB memory count for functions usage
	FaasRamSizeUsage NotificationThresholdLastMessageRegionalQuotaFaasRamSizeUsage `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit NotificationThresholdLastMessageRegionalQuotaFirewallCountLimit `json:"firewall_count_limit"`
	// Firewalls Count usage
	FirewallCountUsage NotificationThresholdLastMessageRegionalQuotaFirewallCountUsage `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit NotificationThresholdLastMessageRegionalQuotaFloatingCountLimit `json:"floating_count_limit"`
	// Floating IP Count usage
	FloatingCountUsage NotificationThresholdLastMessageRegionalQuotaFloatingCountUsage `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit NotificationThresholdLastMessageRegionalQuotaGPUCountLimit `json:"gpu_count_limit"`
	// GPU Count usage
	GPUCountUsage NotificationThresholdLastMessageRegionalQuotaGPUCountUsage `json:"gpu_count_usage"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountLimit `json:"gpu_virtual_a100_count_limit"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountUsage `json:"gpu_virtual_a100_count_usage"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountLimit `json:"gpu_virtual_h100_count_limit"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountUsage `json:"gpu_virtual_h100_count_usage"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountLimit `json:"gpu_virtual_h200_count_limit"`
	// Virtual H200 GPU card count usage
	GPUVirtualH200CountUsage NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountUsage `json:"gpu_virtual_h200_count_usage"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountLimit `json:"gpu_virtual_l40s_count_limit"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountUsage `json:"gpu_virtual_l40s_count_usage"`
	// Images Count limit
	ImageCountLimit NotificationThresholdLastMessageRegionalQuotaImageCountLimit `json:"image_count_limit"`
	// Images Count usage
	ImageCountUsage NotificationThresholdLastMessageRegionalQuotaImageCountUsage `json:"image_count_usage"`
	// Images Size, bytes limit
	ImageSizeLimit NotificationThresholdLastMessageRegionalQuotaImageSizeLimit `json:"image_size_limit"`
	// Images Size, bytes usage
	ImageSizeUsage NotificationThresholdLastMessageRegionalQuotaImageSizeUsage `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit NotificationThresholdLastMessageRegionalQuotaIpuCountLimit `json:"ipu_count_limit"`
	// IPU Count usage
	IpuCountUsage NotificationThresholdLastMessageRegionalQuotaIpuCountUsage `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit NotificationThresholdLastMessageRegionalQuotaLaasTopicCountLimit `json:"laas_topic_count_limit"`
	// LaaS Topics Count usage
	LaasTopicCountUsage NotificationThresholdLastMessageRegionalQuotaLaasTopicCountUsage `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountLimit `json:"loadbalancer_count_limit"`
	// Load Balancers Count usage
	LoadbalancerCountUsage NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountUsage `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit NotificationThresholdLastMessageRegionalQuotaNetworkCountLimit `json:"network_count_limit"`
	// Networks Count usage
	NetworkCountUsage NotificationThresholdLastMessageRegionalQuotaNetworkCountUsage `json:"network_count_usage"`
	// RAM Size, MiB limit
	RamLimit NotificationThresholdLastMessageRegionalQuotaRamLimit `json:"ram_limit"`
	// RAM Size, MiB usage
	RamUsage NotificationThresholdLastMessageRegionalQuotaRamUsage `json:"ram_usage"`
	// Registries count limit
	RegistryCountLimit NotificationThresholdLastMessageRegionalQuotaRegistryCountLimit `json:"registry_count_limit"`
	// Registries count usage
	RegistryCountUsage NotificationThresholdLastMessageRegionalQuotaRegistryCountUsage `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit NotificationThresholdLastMessageRegionalQuotaRegistryStorageLimit `json:"registry_storage_limit"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage NotificationThresholdLastMessageRegionalQuotaRegistryStorageUsage `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit NotificationThresholdLastMessageRegionalQuotaRouterCountLimit `json:"router_count_limit"`
	// Routers Count usage
	RouterCountUsage NotificationThresholdLastMessageRegionalQuotaRouterCountUsage `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit NotificationThresholdLastMessageRegionalQuotaSecretCountLimit `json:"secret_count_limit"`
	// Secret Count usage
	SecretCountUsage NotificationThresholdLastMessageRegionalQuotaSecretCountUsage `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit NotificationThresholdLastMessageRegionalQuotaServergroupCountLimit `json:"servergroup_count_limit"`
	// Placement Group Count usage
	ServergroupCountUsage NotificationThresholdLastMessageRegionalQuotaServergroupCountUsage `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit NotificationThresholdLastMessageRegionalQuotaSfsCountLimit `json:"sfs_count_limit"`
	// Shared file system Count usage
	SfsCountUsage NotificationThresholdLastMessageRegionalQuotaSfsCountUsage `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit NotificationThresholdLastMessageRegionalQuotaSfsSizeLimit `json:"sfs_size_limit"`
	// Shared file system Size, GiB usage
	SfsSizeUsage NotificationThresholdLastMessageRegionalQuotaSfsSizeUsage `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit NotificationThresholdLastMessageRegionalQuotaSharedVmCountLimit `json:"shared_vm_count_limit"`
	// Basic VMs Count usage
	SharedVmCountUsage NotificationThresholdLastMessageRegionalQuotaSharedVmCountUsage `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountLimit `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountUsage `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit NotificationThresholdLastMessageRegionalQuotaSubnetCountLimit `json:"subnet_count_limit"`
	// Subnets Count usage
	SubnetCountUsage NotificationThresholdLastMessageRegionalQuotaSubnetCountUsage `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit NotificationThresholdLastMessageRegionalQuotaVmCountLimit `json:"vm_count_limit"`
	// Instances Dedicated Count usage
	VmCountUsage NotificationThresholdLastMessageRegionalQuotaVmCountUsage `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit NotificationThresholdLastMessageRegionalQuotaVolumeCountLimit `json:"volume_count_limit"`
	// Volumes Count usage
	VolumeCountUsage NotificationThresholdLastMessageRegionalQuotaVolumeCountUsage `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit NotificationThresholdLastMessageRegionalQuotaVolumeSizeLimit `json:"volume_size_limit"`
	// Volumes Size, GiB usage
	VolumeSizeUsage NotificationThresholdLastMessageRegionalQuotaVolumeSizeUsage `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountLimit `json:"volume_snapshots_count_limit"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountUsage `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeLimit `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeUsage `json:"volume_snapshots_size_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionID                          respjson.Field
		RegionName                        respjson.Field
		BaremetalBasicCountLimit          respjson.Field
		BaremetalBasicCountUsage          respjson.Field
		BaremetalGPUA100CountLimit        respjson.Field
		BaremetalGPUA100CountUsage        respjson.Field
		BaremetalGPUH100CountLimit        respjson.Field
		BaremetalGPUH100CountUsage        respjson.Field
		BaremetalGPUH200CountLimit        respjson.Field
		BaremetalGPUH200CountUsage        respjson.Field
		BaremetalGPUL40sCountLimit        respjson.Field
		BaremetalGPUL40sCountUsage        respjson.Field
		BaremetalHfCountLimit             respjson.Field
		BaremetalHfCountUsage             respjson.Field
		BaremetalInfrastructureCountLimit respjson.Field
		BaremetalInfrastructureCountUsage respjson.Field
		BaremetalNetworkCountLimit        respjson.Field
		BaremetalNetworkCountUsage        respjson.Field
		BaremetalStorageCountLimit        respjson.Field
		BaremetalStorageCountUsage        respjson.Field
		CaasContainerCountLimit           respjson.Field
		CaasContainerCountUsage           respjson.Field
		CaasCPUCountLimit                 respjson.Field
		CaasCPUCountUsage                 respjson.Field
		CaasGPUCountLimit                 respjson.Field
		CaasGPUCountUsage                 respjson.Field
		CaasRamSizeLimit                  respjson.Field
		CaasRamSizeUsage                  respjson.Field
		ClusterCountLimit                 respjson.Field
		ClusterCountUsage                 respjson.Field
		CPUCountLimit                     respjson.Field
		CPUCountUsage                     respjson.Field
		DbaasPostgresClusterCountLimit    respjson.Field
		DbaasPostgresClusterCountUsage    respjson.Field
		ExternalIPCountLimit              respjson.Field
		ExternalIPCountUsage              respjson.Field
		FaasCPUCountLimit                 respjson.Field
		FaasCPUCountUsage                 respjson.Field
		FaasFunctionCountLimit            respjson.Field
		FaasFunctionCountUsage            respjson.Field
		FaasNamespaceCountLimit           respjson.Field
		FaasNamespaceCountUsage           respjson.Field
		FaasRamSizeLimit                  respjson.Field
		FaasRamSizeUsage                  respjson.Field
		FirewallCountLimit                respjson.Field
		FirewallCountUsage                respjson.Field
		FloatingCountLimit                respjson.Field
		FloatingCountUsage                respjson.Field
		GPUCountLimit                     respjson.Field
		GPUCountUsage                     respjson.Field
		GPUVirtualA100CountLimit          respjson.Field
		GPUVirtualA100CountUsage          respjson.Field
		GPUVirtualH100CountLimit          respjson.Field
		GPUVirtualH100CountUsage          respjson.Field
		GPUVirtualH200CountLimit          respjson.Field
		GPUVirtualH200CountUsage          respjson.Field
		GPUVirtualL40sCountLimit          respjson.Field
		GPUVirtualL40sCountUsage          respjson.Field
		ImageCountLimit                   respjson.Field
		ImageCountUsage                   respjson.Field
		ImageSizeLimit                    respjson.Field
		ImageSizeUsage                    respjson.Field
		IpuCountLimit                     respjson.Field
		IpuCountUsage                     respjson.Field
		LaasTopicCountLimit               respjson.Field
		LaasTopicCountUsage               respjson.Field
		LoadbalancerCountLimit            respjson.Field
		LoadbalancerCountUsage            respjson.Field
		NetworkCountLimit                 respjson.Field
		NetworkCountUsage                 respjson.Field
		RamLimit                          respjson.Field
		RamUsage                          respjson.Field
		RegistryCountLimit                respjson.Field
		RegistryCountUsage                respjson.Field
		RegistryStorageLimit              respjson.Field
		RegistryStorageUsage              respjson.Field
		RouterCountLimit                  respjson.Field
		RouterCountUsage                  respjson.Field
		SecretCountLimit                  respjson.Field
		SecretCountUsage                  respjson.Field
		ServergroupCountLimit             respjson.Field
		ServergroupCountUsage             respjson.Field
		SfsCountLimit                     respjson.Field
		SfsCountUsage                     respjson.Field
		SfsSizeLimit                      respjson.Field
		SfsSizeUsage                      respjson.Field
		SharedVmCountLimit                respjson.Field
		SharedVmCountUsage                respjson.Field
		SnapshotScheduleCountLimit        respjson.Field
		SnapshotScheduleCountUsage        respjson.Field
		SubnetCountLimit                  respjson.Field
		SubnetCountUsage                  respjson.Field
		VmCountLimit                      respjson.Field
		VmCountUsage                      respjson.Field
		VolumeCountLimit                  respjson.Field
		VolumeCountUsage                  respjson.Field
		VolumeSizeLimit                   respjson.Field
		VolumeSizeUsage                   respjson.Field
		VolumeSnapshotsCountLimit         respjson.Field
		VolumeSnapshotsCountUsage         respjson.Field
		VolumeSnapshotsSizeLimit          respjson.Field
		VolumeSnapshotsSizeUsage          respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuota) RawJSON() string { return r.JSON.raw }
func (r *NotificationThresholdLastMessageRegionalQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic bare metal servers count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic bare metal servers count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalBasicCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal A100 GPU server count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal A100 GPU server count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H100 GPU server count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H100 GPU server count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H200 GPU server count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H200 GPU server count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUH200CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal L40S GPU server count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal L40S GPU server count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalGPUL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High-frequency bare metal servers count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High-frequency bare metal servers count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalHfCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure bare metal servers count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure bare metal servers count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalInfrastructureCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal Network Count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal Network Count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalNetworkCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage bare metal servers count limit
type NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage bare metal servers count usage
type NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaBaremetalStorageCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers count limit
type NotificationThresholdLastMessageRegionalQuotaCaasContainerCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasContainerCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasContainerCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers count usage
type NotificationThresholdLastMessageRegionalQuotaCaasContainerCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasContainerCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasContainerCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for containers limit
type NotificationThresholdLastMessageRegionalQuotaCaasCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasCPUCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for containers usage
type NotificationThresholdLastMessageRegionalQuotaCaasCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasCPUCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers gpu count limit
type NotificationThresholdLastMessageRegionalQuotaCaasGPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasGPUCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasGPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers gpu count usage
type NotificationThresholdLastMessageRegionalQuotaCaasGPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasGPUCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasGPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for containers limit
type NotificationThresholdLastMessageRegionalQuotaCaasRamSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasRamSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasRamSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for containers usage
type NotificationThresholdLastMessageRegionalQuotaCaasRamSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCaasRamSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCaasRamSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// K8s clusters count limit
type NotificationThresholdLastMessageRegionalQuotaClusterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaClusterCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaClusterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// K8s clusters count usage
type NotificationThresholdLastMessageRegionalQuotaClusterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaClusterCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaClusterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// vCPU Count limit
type NotificationThresholdLastMessageRegionalQuotaCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCPUCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// vCPU Count usage
type NotificationThresholdLastMessageRegionalQuotaCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaCPUCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBaaS cluster count limit
type NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBaaS cluster count usage
type NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaDbaasPostgresClusterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External IP Count limit
type NotificationThresholdLastMessageRegionalQuotaExternalIPCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaExternalIPCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaExternalIPCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External IP Count usage
type NotificationThresholdLastMessageRegionalQuotaExternalIPCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaExternalIPCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaExternalIPCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for functions limit
type NotificationThresholdLastMessageRegionalQuotaFaasCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasCPUCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for functions usage
type NotificationThresholdLastMessageRegionalQuotaFaasCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasCPUCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions count limit
type NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions count usage
type NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasFunctionCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions namespace count limit
type NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions namespace count usage
type NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasNamespaceCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for functions limit
type NotificationThresholdLastMessageRegionalQuotaFaasRamSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasRamSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasRamSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for functions usage
type NotificationThresholdLastMessageRegionalQuotaFaasRamSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFaasRamSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFaasRamSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Firewalls Count limit
type NotificationThresholdLastMessageRegionalQuotaFirewallCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFirewallCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFirewallCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Firewalls Count usage
type NotificationThresholdLastMessageRegionalQuotaFirewallCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFirewallCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFirewallCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP Count limit
type NotificationThresholdLastMessageRegionalQuotaFloatingCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFloatingCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFloatingCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP Count usage
type NotificationThresholdLastMessageRegionalQuotaFloatingCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaFloatingCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaFloatingCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU Count limit
type NotificationThresholdLastMessageRegionalQuotaGPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU Count usage
type NotificationThresholdLastMessageRegionalQuotaGPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual A100 GPU card count limit
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual A100 GPU card count usage
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H100 GPU card count limit
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H100 GPU card count usage
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H200 GPU card count limit
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H200 GPU card count usage
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualH200CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual L40S GPU card count limit
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual L40S GPU card count usage
type NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaGPUVirtualL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Count limit
type NotificationThresholdLastMessageRegionalQuotaImageCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaImageCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaImageCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Count usage
type NotificationThresholdLastMessageRegionalQuotaImageCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaImageCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaImageCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Size, bytes limit
type NotificationThresholdLastMessageRegionalQuotaImageSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaImageSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaImageSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Size, bytes usage
type NotificationThresholdLastMessageRegionalQuotaImageSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaImageSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaImageSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IPU Count limit
type NotificationThresholdLastMessageRegionalQuotaIpuCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaIpuCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaIpuCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IPU Count usage
type NotificationThresholdLastMessageRegionalQuotaIpuCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaIpuCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaIpuCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LaaS Topics Count limit
type NotificationThresholdLastMessageRegionalQuotaLaasTopicCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaLaasTopicCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaLaasTopicCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LaaS Topics Count usage
type NotificationThresholdLastMessageRegionalQuotaLaasTopicCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaLaasTopicCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaLaasTopicCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Load Balancers Count limit
type NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Load Balancers Count usage
type NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaLoadbalancerCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networks Count limit
type NotificationThresholdLastMessageRegionalQuotaNetworkCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaNetworkCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaNetworkCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networks Count usage
type NotificationThresholdLastMessageRegionalQuotaNetworkCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaNetworkCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaNetworkCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RAM Size, MiB limit
type NotificationThresholdLastMessageRegionalQuotaRamLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRamLimit) RawJSON() string { return r.JSON.raw }
func (r *NotificationThresholdLastMessageRegionalQuotaRamLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RAM Size, MiB usage
type NotificationThresholdLastMessageRegionalQuotaRamUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRamUsage) RawJSON() string { return r.JSON.raw }
func (r *NotificationThresholdLastMessageRegionalQuotaRamUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries count limit
type NotificationThresholdLastMessageRegionalQuotaRegistryCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRegistryCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRegistryCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries count usage
type NotificationThresholdLastMessageRegionalQuotaRegistryCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRegistryCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRegistryCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries volume usage, GiB limit
type NotificationThresholdLastMessageRegionalQuotaRegistryStorageLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRegistryStorageLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRegistryStorageLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries volume usage, GiB usage
type NotificationThresholdLastMessageRegionalQuotaRegistryStorageUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRegistryStorageUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRegistryStorageUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routers Count limit
type NotificationThresholdLastMessageRegionalQuotaRouterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRouterCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRouterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routers Count usage
type NotificationThresholdLastMessageRegionalQuotaRouterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaRouterCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaRouterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret Count limit
type NotificationThresholdLastMessageRegionalQuotaSecretCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSecretCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSecretCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret Count usage
type NotificationThresholdLastMessageRegionalQuotaSecretCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSecretCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSecretCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Placement Group Count limit
type NotificationThresholdLastMessageRegionalQuotaServergroupCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaServergroupCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaServergroupCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Placement Group Count usage
type NotificationThresholdLastMessageRegionalQuotaServergroupCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaServergroupCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaServergroupCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Count limit
type NotificationThresholdLastMessageRegionalQuotaSfsCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSfsCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSfsCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Count usage
type NotificationThresholdLastMessageRegionalQuotaSfsCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSfsCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSfsCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Size, GiB limit
type NotificationThresholdLastMessageRegionalQuotaSfsSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSfsSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSfsSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Size, GiB usage
type NotificationThresholdLastMessageRegionalQuotaSfsSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSfsSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSfsSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic VMs Count limit
type NotificationThresholdLastMessageRegionalQuotaSharedVmCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSharedVmCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSharedVmCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic VMs Count usage
type NotificationThresholdLastMessageRegionalQuotaSharedVmCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSharedVmCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSharedVmCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot Schedules Count limit
type NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot Schedules Count usage
type NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSnapshotScheduleCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subnets Count limit
type NotificationThresholdLastMessageRegionalQuotaSubnetCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSubnetCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSubnetCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subnets Count usage
type NotificationThresholdLastMessageRegionalQuotaSubnetCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaSubnetCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaSubnetCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances Dedicated Count limit
type NotificationThresholdLastMessageRegionalQuotaVmCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVmCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVmCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances Dedicated Count usage
type NotificationThresholdLastMessageRegionalQuotaVmCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVmCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVmCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Count limit
type NotificationThresholdLastMessageRegionalQuotaVolumeCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Count usage
type NotificationThresholdLastMessageRegionalQuotaVolumeCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Size, GiB limit
type NotificationThresholdLastMessageRegionalQuotaVolumeSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Size, GiB usage
type NotificationThresholdLastMessageRegionalQuotaVolumeSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Count limit
type NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Count usage
type NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Size, GiB limit
type NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Size, GiB usage
type NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationThresholdLastMessageRegionalQuotaVolumeSnapshotsSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaNotificationThresholdUpdateParams struct {
	// Quota notification threshold in percentage
	Threshold int64 `json:"threshold" api:"required"`
	// Time of last successful email sending
	LastSending param.Opt[time.Time] `json:"last_sending,omitzero" format:"date-time"`
	// A message data
	LastMessage QuotaNotificationThresholdUpdateParamsLastMessage `json:"last_message,omitzero"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message data
//
// The properties GlobalQuotas, RegionalQuotas are required.
type QuotaNotificationThresholdUpdateParamsLastMessage struct {
	// Global quota that exceed the threshold
	GlobalQuotas QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas `json:"global_quotas,omitzero" api:"required"`
	// Regional quota that exceed the threshold
	RegionalQuotas []QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota `json:"regional_quotas,omitzero" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global quota that exceed the threshold
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit `json:"inference_cpu_millicore_count_limit,omitzero"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage `json:"inference_cpu_millicore_count_usage,omitzero"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit `json:"inference_gpu_a100_count_limit,omitzero"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage `json:"inference_gpu_a100_count_usage,omitzero"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit `json:"inference_gpu_h100_count_limit,omitzero"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage `json:"inference_gpu_h100_count_usage,omitzero"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit `json:"inference_gpu_l40s_count_limit,omitzero"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage `json:"inference_gpu_l40s_count_usage,omitzero"`
	// Inference instance count limit
	InferenceInstanceCountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit `json:"inference_instance_count_limit,omitzero"`
	// Inference instance count usage
	InferenceInstanceCountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage `json:"inference_instance_count_usage,omitzero"`
	// SSH Keys Count limit
	KeypairCountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit `json:"keypair_count_limit,omitzero"`
	// SSH Keys Count usage
	KeypairCountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage `json:"keypair_count_usage,omitzero"`
	// Projects Count limit
	ProjectCountLimit QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit `json:"project_count_limit,omitzero"`
	// Projects Count usage
	ProjectCountUsage QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage `json:"project_count_usage,omitzero"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference CPU millicore count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference CPU millicore count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU A100 Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU A100 Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU H100 Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU H100 Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU L40s Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference GPU L40s Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference instance count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference instance count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH Keys Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH Keys Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Projects Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Projects Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RegionID, RegionName are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota struct {
	// Region id
	RegionID int64 `json:"region_id" api:"required"`
	// Region name
	RegionName string `json:"region_name" api:"required"`
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit `json:"baremetal_basic_count_limit,omitzero"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage `json:"baremetal_basic_count_usage,omitzero"`
	// Bare metal A100 GPU server count limit
	BaremetalGPUA100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit `json:"baremetal_gpu_a100_count_limit,omitzero"`
	// Bare metal A100 GPU server count usage
	BaremetalGPUA100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage `json:"baremetal_gpu_a100_count_usage,omitzero"`
	// Bare metal H100 GPU server count limit
	BaremetalGPUH100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit `json:"baremetal_gpu_h100_count_limit,omitzero"`
	// Bare metal H100 GPU server count usage
	BaremetalGPUH100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage `json:"baremetal_gpu_h100_count_usage,omitzero"`
	// Bare metal H200 GPU server count limit
	BaremetalGPUH200CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit `json:"baremetal_gpu_h200_count_limit,omitzero"`
	// Bare metal H200 GPU server count usage
	BaremetalGPUH200CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage `json:"baremetal_gpu_h200_count_usage,omitzero"`
	// Bare metal L40S GPU server count limit
	BaremetalGPUL40sCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit `json:"baremetal_gpu_l40s_count_limit,omitzero"`
	// Bare metal L40S GPU server count usage
	BaremetalGPUL40sCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage `json:"baremetal_gpu_l40s_count_usage,omitzero"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit `json:"baremetal_hf_count_limit,omitzero"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage `json:"baremetal_hf_count_usage,omitzero"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit `json:"baremetal_infrastructure_count_limit,omitzero"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage `json:"baremetal_infrastructure_count_usage,omitzero"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit `json:"baremetal_network_count_limit,omitzero"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage `json:"baremetal_network_count_usage,omitzero"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit `json:"baremetal_storage_count_limit,omitzero"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage `json:"baremetal_storage_count_usage,omitzero"`
	// Containers count limit
	CaasContainerCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit `json:"caas_container_count_limit,omitzero"`
	// Containers count usage
	CaasContainerCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage `json:"caas_container_count_usage,omitzero"`
	// mCPU count for containers limit
	CaasCPUCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit `json:"caas_cpu_count_limit,omitzero"`
	// mCPU count for containers usage
	CaasCPUCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage `json:"caas_cpu_count_usage,omitzero"`
	// Containers gpu count limit
	CaasGPUCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit `json:"caas_gpu_count_limit,omitzero"`
	// Containers gpu count usage
	CaasGPUCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage `json:"caas_gpu_count_usage,omitzero"`
	// MiB memory count for containers limit
	CaasRamSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit `json:"caas_ram_size_limit,omitzero"`
	// MiB memory count for containers usage
	CaasRamSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage `json:"caas_ram_size_usage,omitzero"`
	// K8s clusters count limit
	ClusterCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit `json:"cluster_count_limit,omitzero"`
	// K8s clusters count usage
	ClusterCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage `json:"cluster_count_usage,omitzero"`
	// vCPU Count limit
	CPUCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit `json:"cpu_count_limit,omitzero"`
	// vCPU Count usage
	CPUCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage `json:"cpu_count_usage,omitzero"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit `json:"dbaas_postgres_cluster_count_limit,omitzero"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage `json:"dbaas_postgres_cluster_count_usage,omitzero"`
	// External IP Count limit
	ExternalIPCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit `json:"external_ip_count_limit,omitzero"`
	// External IP Count usage
	ExternalIPCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage `json:"external_ip_count_usage,omitzero"`
	// mCPU count for functions limit
	FaasCPUCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit `json:"faas_cpu_count_limit,omitzero"`
	// mCPU count for functions usage
	FaasCPUCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage `json:"faas_cpu_count_usage,omitzero"`
	// Functions count limit
	FaasFunctionCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit `json:"faas_function_count_limit,omitzero"`
	// Functions count usage
	FaasFunctionCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage `json:"faas_function_count_usage,omitzero"`
	// Functions namespace count limit
	FaasNamespaceCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit `json:"faas_namespace_count_limit,omitzero"`
	// Functions namespace count usage
	FaasNamespaceCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage `json:"faas_namespace_count_usage,omitzero"`
	// MiB memory count for functions limit
	FaasRamSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit `json:"faas_ram_size_limit,omitzero"`
	// MiB memory count for functions usage
	FaasRamSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage `json:"faas_ram_size_usage,omitzero"`
	// Firewalls Count limit
	FirewallCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit `json:"firewall_count_limit,omitzero"`
	// Firewalls Count usage
	FirewallCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage `json:"firewall_count_usage,omitzero"`
	// Floating IP Count limit
	FloatingCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit `json:"floating_count_limit,omitzero"`
	// Floating IP Count usage
	FloatingCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage `json:"floating_count_usage,omitzero"`
	// GPU Count limit
	GPUCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit `json:"gpu_count_limit,omitzero"`
	// GPU Count usage
	GPUCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage `json:"gpu_count_usage,omitzero"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit `json:"gpu_virtual_a100_count_limit,omitzero"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage `json:"gpu_virtual_a100_count_usage,omitzero"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit `json:"gpu_virtual_h100_count_limit,omitzero"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage `json:"gpu_virtual_h100_count_usage,omitzero"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit `json:"gpu_virtual_h200_count_limit,omitzero"`
	// Virtual H200 GPU card count usage
	GPUVirtualH200CountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage `json:"gpu_virtual_h200_count_usage,omitzero"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit `json:"gpu_virtual_l40s_count_limit,omitzero"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage `json:"gpu_virtual_l40s_count_usage,omitzero"`
	// Images Count limit
	ImageCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit `json:"image_count_limit,omitzero"`
	// Images Count usage
	ImageCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage `json:"image_count_usage,omitzero"`
	// Images Size, bytes limit
	ImageSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit `json:"image_size_limit,omitzero"`
	// Images Size, bytes usage
	ImageSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage `json:"image_size_usage,omitzero"`
	// IPU Count limit
	IpuCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit `json:"ipu_count_limit,omitzero"`
	// IPU Count usage
	IpuCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage `json:"ipu_count_usage,omitzero"`
	// LaaS Topics Count limit
	LaasTopicCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit `json:"laas_topic_count_limit,omitzero"`
	// LaaS Topics Count usage
	LaasTopicCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage `json:"laas_topic_count_usage,omitzero"`
	// Load Balancers Count limit
	LoadbalancerCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit `json:"loadbalancer_count_limit,omitzero"`
	// Load Balancers Count usage
	LoadbalancerCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage `json:"loadbalancer_count_usage,omitzero"`
	// Networks Count limit
	NetworkCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit `json:"network_count_limit,omitzero"`
	// Networks Count usage
	NetworkCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage `json:"network_count_usage,omitzero"`
	// RAM Size, MiB limit
	RamLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit `json:"ram_limit,omitzero"`
	// RAM Size, MiB usage
	RamUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage `json:"ram_usage,omitzero"`
	// Registries count limit
	RegistryCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit `json:"registry_count_limit,omitzero"`
	// Registries count usage
	RegistryCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage `json:"registry_count_usage,omitzero"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit `json:"registry_storage_limit,omitzero"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage `json:"registry_storage_usage,omitzero"`
	// Routers Count limit
	RouterCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit `json:"router_count_limit,omitzero"`
	// Routers Count usage
	RouterCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage `json:"router_count_usage,omitzero"`
	// Secret Count limit
	SecretCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit `json:"secret_count_limit,omitzero"`
	// Secret Count usage
	SecretCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage `json:"secret_count_usage,omitzero"`
	// Placement Group Count limit
	ServergroupCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit `json:"servergroup_count_limit,omitzero"`
	// Placement Group Count usage
	ServergroupCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage `json:"servergroup_count_usage,omitzero"`
	// Shared file system Count limit
	SfsCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit `json:"sfs_count_limit,omitzero"`
	// Shared file system Count usage
	SfsCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage `json:"sfs_count_usage,omitzero"`
	// Shared file system Size, GiB limit
	SfsSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit `json:"sfs_size_limit,omitzero"`
	// Shared file system Size, GiB usage
	SfsSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage `json:"sfs_size_usage,omitzero"`
	// Basic VMs Count limit
	SharedVmCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit `json:"shared_vm_count_limit,omitzero"`
	// Basic VMs Count usage
	SharedVmCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage `json:"shared_vm_count_usage,omitzero"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit `json:"snapshot_schedule_count_limit,omitzero"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage `json:"snapshot_schedule_count_usage,omitzero"`
	// Subnets Count limit
	SubnetCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit `json:"subnet_count_limit,omitzero"`
	// Subnets Count usage
	SubnetCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage `json:"subnet_count_usage,omitzero"`
	// Instances Dedicated Count limit
	VmCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit `json:"vm_count_limit,omitzero"`
	// Instances Dedicated Count usage
	VmCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage `json:"vm_count_usage,omitzero"`
	// Volumes Count limit
	VolumeCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit `json:"volume_count_limit,omitzero"`
	// Volumes Count usage
	VolumeCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage `json:"volume_count_usage,omitzero"`
	// Volumes Size, GiB limit
	VolumeSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit `json:"volume_size_limit,omitzero"`
	// Volumes Size, GiB usage
	VolumeSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage `json:"volume_size_usage,omitzero"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit `json:"volume_snapshots_count_limit,omitzero"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage `json:"volume_snapshots_count_usage,omitzero"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit `json:"volume_snapshots_size_limit,omitzero"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage `json:"volume_snapshots_size_usage,omitzero"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic bare metal servers count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic bare metal servers count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal A100 GPU server count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal A100 GPU server count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H100 GPU server count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H100 GPU server count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H200 GPU server count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal H200 GPU server count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal L40S GPU server count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal L40S GPU server count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High-frequency bare metal servers count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High-frequency bare metal servers count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure bare metal servers count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure bare metal servers count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal Network Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal Network Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage bare metal servers count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage bare metal servers count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for containers limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for containers usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers gpu count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Containers gpu count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for containers limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for containers usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// K8s clusters count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// K8s clusters count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// vCPU Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// vCPU Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBaaS cluster count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBaaS cluster count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External IP Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External IP Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for functions limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// mCPU count for functions usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions namespace count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Functions namespace count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for functions limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MiB memory count for functions usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Firewalls Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Firewalls Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual A100 GPU card count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual A100 GPU card count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H100 GPU card count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H100 GPU card count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H200 GPU card count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual H200 GPU card count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual L40S GPU card count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual L40S GPU card count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Size, bytes limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Images Size, bytes usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IPU Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IPU Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LaaS Topics Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LaaS Topics Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Load Balancers Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Load Balancers Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networks Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networks Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RAM Size, MiB limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RAM Size, MiB usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries volume usage, GiB limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registries volume usage, GiB usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routers Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routers Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Placement Group Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Placement Group Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Size, GiB limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared file system Size, GiB usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic VMs Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic VMs Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot Schedules Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot Schedules Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subnets Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subnets Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances Dedicated Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances Dedicated Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Size, GiB limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volumes Size, GiB usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Count limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Count usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Size, GiB limit
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshots Size, GiB usage
//
// The properties Limit, Usage are required.
type QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage struct {
	// Сurrent quota limit
	Limit int64 `json:"limit" api:"required"`
	// Сurrent amount of resource used
	Usage int64 `json:"usage" api:"required"`
	paramObj
}

func (r QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage) MarshalJSON() (data []byte, err error) {
	type shadow QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
