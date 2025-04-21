// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// QuotaRequestService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaRequestService] method instead.
type QuotaRequestService struct {
	Options []option.RequestOption
}

// NewQuotaRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuotaRequestService(opts ...option.RequestOption) (r QuotaRequestService) {
	r = QuotaRequestService{}
	r.Options = opts
	return
}

// Create request to change quotas
func (r *QuotaRequestService) New(ctx context.Context, body QuotaRequestNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v2/limits_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Returns a list of sent requests to change current quotas and their statuses
func (r *QuotaRequestService) List(ctx context.Context, query QuotaRequestListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v2/limits_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Delete request to change quotas
func (r *QuotaRequestService) Delete(ctx context.Context, requestID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/limits_request/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get request to change quota limits.
func (r *QuotaRequestService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *QuotaRequestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/limits_request/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/paths/%2Fcloud%2Fv2%2Flimits_request%2F%7Brequest_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v2/limits_request/{request_id}'].get.responses[200].content['application/json'].schema"
type QuotaRequestGetResponse struct {
	// '#/components/schemas/LimitsRequestSerializer/properties/id'
	// "$.components.schemas.LimitsRequestSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/LimitsRequestSerializer/properties/client_id'
	// "$.components.schemas.LimitsRequestSerializer.properties.client_id"
	ClientID int64 `json:"client_id,required"`
	// '#/components/schemas/LimitsRequestSerializer/properties/requested_limits'
	// "$.components.schemas.LimitsRequestSerializer.properties.requested_limits"
	RequestedLimits QuotaRequestGetResponseRequestedLimits `json:"requested_limits,required"`
	// '#/components/schemas/LimitsRequestSerializer/properties/status'
	// "$.components.schemas.LimitsRequestSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/LimitsRequestSerializer/properties/created_at'
	// "$.components.schemas.LimitsRequestSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// '#/components/schemas/LimitsRequestSerializer/properties/description/anyOf/0'
	// "$.components.schemas.LimitsRequestSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/LimitsRequestSerializer/properties/updated_at/anyOf/0'
	// "$.components.schemas.LimitsRequestSerializer.properties.updated_at.anyOf[0]"
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID              resp.Field
		ClientID        resp.Field
		RequestedLimits resp.Field
		Status          resp.Field
		CreatedAt       resp.Field
		Description     resp.Field
		UpdatedAt       resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LimitsRequestSerializer/properties/requested_limits'
// "$.components.schemas.LimitsRequestSerializer.properties.requested_limits"
type QuotaRequestGetResponseRequestedLimits struct {
	// '#/components/schemas/AllClientQuotasLimitsSerializer/properties/global_limits'
	// "$.components.schemas.AllClientQuotasLimitsSerializer.properties.global_limits"
	GlobalLimits QuotaRequestGetResponseRequestedLimitsGlobalLimits `json:"global_limits"`
	// '#/components/schemas/AllClientQuotasLimitsSerializer/properties/regional_limits'
	// "$.components.schemas.AllClientQuotasLimitsSerializer.properties.regional_limits"
	RegionalLimits []QuotaRequestGetResponseRequestedLimitsRegionalLimit `json:"regional_limits"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GlobalLimits   resp.Field
		RegionalLimits resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/AllClientQuotasLimitsSerializer/properties/global_limits'
// "$.components.schemas.AllClientQuotasLimitsSerializer.properties.global_limits"
type QuotaRequestGetResponseRequestedLimitsGlobalLimits struct {
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_cpu_millicore_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_cpu_millicore_count_limit"
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_a100_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_a100_count_limit"
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_h100_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_h100_count_limit"
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_l40s_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_l40s_count_limit"
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_instance_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_instance_count_limit"
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/keypair_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.keypair_count_limit"
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/project_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.project_count_limit"
	ProjectCountLimit int64 `json:"project_count_limit"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InferenceCPUMillicoreCountLimit resp.Field
		InferenceGPUA100CountLimit      resp.Field
		InferenceGPUH100CountLimit      resp.Field
		InferenceGPUL40sCountLimit      resp.Field
		InferenceInstanceCountLimit     resp.Field
		KeypairCountLimit               resp.Field
		ProjectCountLimit               resp.Field
		ExtraFields                     map[string]resp.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimitsGlobalLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimitsGlobalLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/AllClientQuotasLimitsSerializer/properties/regional_limits/items'
// "$.components.schemas.AllClientQuotasLimitsSerializer.properties.regional_limits.items"
type QuotaRequestGetResponseRequestedLimitsRegionalLimit struct {
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_basic_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_basic_count_limit"
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_gpu_count_limit"
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_hf_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_hf_count_limit"
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_infrastructure_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_infrastructure_count_limit"
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_network_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_network_count_limit"
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_storage_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_storage_count_limit"
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_container_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_container_count_limit"
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_cpu_count_limit"
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_gpu_count_limit"
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_ram_size_limit"
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/cluster_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.cluster_count_limit"
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.cpu_count_limit"
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/dbaas_postgres_cluster_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.dbaas_postgres_cluster_count_limit"
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/external_ip_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.external_ip_count_limit"
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_cpu_count_limit"
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_function_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_function_count_limit"
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_namespace_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_namespace_count_limit"
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_ram_size_limit"
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/firewall_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.firewall_count_limit"
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/floating_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.floating_count_limit"
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_count_limit"
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_a100_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_a100_count_limit"
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_h100_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_h100_count_limit"
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_l40s_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_l40s_count_limit"
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/image_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.image_count_limit"
	ImageCountLimit int64 `json:"image_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/image_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.image_size_limit"
	ImageSizeLimit int64 `json:"image_size_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/ipu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.ipu_count_limit"
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/laas_topic_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.laas_topic_count_limit"
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/loadbalancer_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.loadbalancer_count_limit"
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/network_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.network_count_limit"
	NetworkCountLimit int64 `json:"network_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/ram_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.ram_limit"
	RamLimit int64 `json:"ram_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/region_id'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.region_id"
	RegionID int64 `json:"region_id"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/registry_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.registry_count_limit"
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/registry_storage_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.registry_storage_limit"
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/router_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.router_count_limit"
	RouterCountLimit int64 `json:"router_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/secret_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.secret_count_limit"
	SecretCountLimit int64 `json:"secret_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/servergroup_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.servergroup_count_limit"
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/sfs_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.sfs_count_limit"
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/sfs_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.sfs_size_limit"
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/shared_vm_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.shared_vm_count_limit"
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/snapshot_schedule_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.snapshot_schedule_count_limit"
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/subnet_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.subnet_count_limit"
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/vm_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.vm_count_limit"
	VmCountLimit int64 `json:"vm_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_count_limit"
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_size_limit"
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_snapshots_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_snapshots_count_limit"
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_snapshots_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_snapshots_size_limit"
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaremetalBasicCountLimit          resp.Field
		BaremetalGPUCountLimit            resp.Field
		BaremetalHfCountLimit             resp.Field
		BaremetalInfrastructureCountLimit resp.Field
		BaremetalNetworkCountLimit        resp.Field
		BaremetalStorageCountLimit        resp.Field
		CaasContainerCountLimit           resp.Field
		CaasCPUCountLimit                 resp.Field
		CaasGPUCountLimit                 resp.Field
		CaasRamSizeLimit                  resp.Field
		ClusterCountLimit                 resp.Field
		CPUCountLimit                     resp.Field
		DbaasPostgresClusterCountLimit    resp.Field
		ExternalIPCountLimit              resp.Field
		FaasCPUCountLimit                 resp.Field
		FaasFunctionCountLimit            resp.Field
		FaasNamespaceCountLimit           resp.Field
		FaasRamSizeLimit                  resp.Field
		FirewallCountLimit                resp.Field
		FloatingCountLimit                resp.Field
		GPUCountLimit                     resp.Field
		GPUVirtualA100CountLimit          resp.Field
		GPUVirtualH100CountLimit          resp.Field
		GPUVirtualL40sCountLimit          resp.Field
		ImageCountLimit                   resp.Field
		ImageSizeLimit                    resp.Field
		IpuCountLimit                     resp.Field
		LaasTopicCountLimit               resp.Field
		LoadbalancerCountLimit            resp.Field
		NetworkCountLimit                 resp.Field
		RamLimit                          resp.Field
		RegionID                          resp.Field
		RegistryCountLimit                resp.Field
		RegistryStorageLimit              resp.Field
		RouterCountLimit                  resp.Field
		SecretCountLimit                  resp.Field
		ServergroupCountLimit             resp.Field
		SfsCountLimit                     resp.Field
		SfsSizeLimit                      resp.Field
		SharedVmCountLimit                resp.Field
		SnapshotScheduleCountLimit        resp.Field
		SubnetCountLimit                  resp.Field
		VmCountLimit                      resp.Field
		VolumeCountLimit                  resp.Field
		VolumeSizeLimit                   resp.Field
		VolumeSnapshotsCountLimit         resp.Field
		VolumeSnapshotsSizeLimit          resp.Field
		ExtraFields                       map[string]resp.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimitsRegionalLimit) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimitsRegionalLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestNewParams struct {
	// '#/components/schemas/LimitsRequestCreateSerializer/properties/description'
	// "$.components.schemas.LimitsRequestCreateSerializer.properties.description"
	Description string `json:"description,required"`
	// '#/components/schemas/LimitsRequestCreateSerializer/properties/requested_limits'
	// "$.components.schemas.LimitsRequestCreateSerializer.properties.requested_limits"
	RequestedLimits QuotaRequestNewParamsRequestedLimits `json:"requested_limits,omitzero,required"`
	// '#/components/schemas/LimitsRequestCreateSerializer/properties/client_id'
	// "$.components.schemas.LimitsRequestCreateSerializer.properties.client_id"
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaRequestNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r QuotaRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/LimitsRequestCreateSerializer/properties/requested_limits'
// "$.components.schemas.LimitsRequestCreateSerializer.properties.requested_limits"
type QuotaRequestNewParamsRequestedLimits struct {
	// '#/components/schemas/ClientMixedQuotasLimitsSerializer/properties/global_limits'
	// "$.components.schemas.ClientMixedQuotasLimitsSerializer.properties.global_limits"
	GlobalLimits QuotaRequestNewParamsRequestedLimitsGlobalLimits `json:"global_limits,omitzero"`
	// '#/components/schemas/ClientMixedQuotasLimitsSerializer/properties/regional_limits'
	// "$.components.schemas.ClientMixedQuotasLimitsSerializer.properties.regional_limits"
	RegionalLimits []QuotaRequestNewParamsRequestedLimitsRegionalLimit `json:"regional_limits,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaRequestNewParamsRequestedLimits) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r QuotaRequestNewParamsRequestedLimits) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimits
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ClientMixedQuotasLimitsSerializer/properties/global_limits'
// "$.components.schemas.ClientMixedQuotasLimitsSerializer.properties.global_limits"
type QuotaRequestNewParamsRequestedLimitsGlobalLimits struct {
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_cpu_millicore_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_cpu_millicore_count_limit"
	InferenceCPUMillicoreCountLimit param.Opt[int64] `json:"inference_cpu_millicore_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_a100_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_a100_count_limit"
	InferenceGPUA100CountLimit param.Opt[int64] `json:"inference_gpu_a100_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_h100_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_h100_count_limit"
	InferenceGPUH100CountLimit param.Opt[int64] `json:"inference_gpu_h100_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_gpu_l40s_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_gpu_l40s_count_limit"
	InferenceGPUL40sCountLimit param.Opt[int64] `json:"inference_gpu_l40s_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/inference_instance_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.inference_instance_count_limit"
	InferenceInstanceCountLimit param.Opt[int64] `json:"inference_instance_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/keypair_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.keypair_count_limit"
	KeypairCountLimit param.Opt[int64] `json:"keypair_count_limit,omitzero"`
	// '#/components/schemas/CreateGlobalQuotasLimitsSerializer/properties/project_count_limit'
	// "$.components.schemas.CreateGlobalQuotasLimitsSerializer.properties.project_count_limit"
	ProjectCountLimit param.Opt[int64] `json:"project_count_limit,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaRequestNewParamsRequestedLimitsGlobalLimits) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r QuotaRequestNewParamsRequestedLimitsGlobalLimits) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimitsGlobalLimits
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ClientMixedQuotasLimitsSerializer/properties/regional_limits/items'
// "$.components.schemas.ClientMixedQuotasLimitsSerializer.properties.regional_limits.items"
type QuotaRequestNewParamsRequestedLimitsRegionalLimit struct {
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_basic_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_basic_count_limit"
	BaremetalBasicCountLimit param.Opt[int64] `json:"baremetal_basic_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_gpu_count_limit"
	BaremetalGPUCountLimit param.Opt[int64] `json:"baremetal_gpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_hf_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_hf_count_limit"
	BaremetalHfCountLimit param.Opt[int64] `json:"baremetal_hf_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_infrastructure_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_infrastructure_count_limit"
	BaremetalInfrastructureCountLimit param.Opt[int64] `json:"baremetal_infrastructure_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_network_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_network_count_limit"
	BaremetalNetworkCountLimit param.Opt[int64] `json:"baremetal_network_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/baremetal_storage_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.baremetal_storage_count_limit"
	BaremetalStorageCountLimit param.Opt[int64] `json:"baremetal_storage_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_container_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_container_count_limit"
	CaasContainerCountLimit param.Opt[int64] `json:"caas_container_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_cpu_count_limit"
	CaasCPUCountLimit param.Opt[int64] `json:"caas_cpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_gpu_count_limit"
	CaasGPUCountLimit param.Opt[int64] `json:"caas_gpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/caas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.caas_ram_size_limit"
	CaasRamSizeLimit param.Opt[int64] `json:"caas_ram_size_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/cluster_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.cluster_count_limit"
	ClusterCountLimit param.Opt[int64] `json:"cluster_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.cpu_count_limit"
	CPUCountLimit param.Opt[int64] `json:"cpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/dbaas_postgres_cluster_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.dbaas_postgres_cluster_count_limit"
	DbaasPostgresClusterCountLimit param.Opt[int64] `json:"dbaas_postgres_cluster_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/external_ip_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.external_ip_count_limit"
	ExternalIPCountLimit param.Opt[int64] `json:"external_ip_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_cpu_count_limit"
	FaasCPUCountLimit param.Opt[int64] `json:"faas_cpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_function_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_function_count_limit"
	FaasFunctionCountLimit param.Opt[int64] `json:"faas_function_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_namespace_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_namespace_count_limit"
	FaasNamespaceCountLimit param.Opt[int64] `json:"faas_namespace_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/faas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.faas_ram_size_limit"
	FaasRamSizeLimit param.Opt[int64] `json:"faas_ram_size_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/firewall_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.firewall_count_limit"
	FirewallCountLimit param.Opt[int64] `json:"firewall_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/floating_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.floating_count_limit"
	FloatingCountLimit param.Opt[int64] `json:"floating_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_count_limit"
	GPUCountLimit param.Opt[int64] `json:"gpu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_a100_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_a100_count_limit"
	GPUVirtualA100CountLimit param.Opt[int64] `json:"gpu_virtual_a100_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_h100_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_h100_count_limit"
	GPUVirtualH100CountLimit param.Opt[int64] `json:"gpu_virtual_h100_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/gpu_virtual_l40s_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.gpu_virtual_l40s_count_limit"
	GPUVirtualL40sCountLimit param.Opt[int64] `json:"gpu_virtual_l40s_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/image_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.image_count_limit"
	ImageCountLimit param.Opt[int64] `json:"image_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/image_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.image_size_limit"
	ImageSizeLimit param.Opt[int64] `json:"image_size_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/ipu_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.ipu_count_limit"
	IpuCountLimit param.Opt[int64] `json:"ipu_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/laas_topic_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.laas_topic_count_limit"
	LaasTopicCountLimit param.Opt[int64] `json:"laas_topic_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/loadbalancer_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.loadbalancer_count_limit"
	LoadbalancerCountLimit param.Opt[int64] `json:"loadbalancer_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/network_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.network_count_limit"
	NetworkCountLimit param.Opt[int64] `json:"network_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/ram_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.ram_limit"
	RamLimit param.Opt[int64] `json:"ram_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/region_id'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.region_id"
	RegionID param.Opt[int64] `json:"region_id,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/registry_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.registry_count_limit"
	RegistryCountLimit param.Opt[int64] `json:"registry_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/registry_storage_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.registry_storage_limit"
	RegistryStorageLimit param.Opt[int64] `json:"registry_storage_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/router_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.router_count_limit"
	RouterCountLimit param.Opt[int64] `json:"router_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/secret_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.secret_count_limit"
	SecretCountLimit param.Opt[int64] `json:"secret_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/servergroup_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.servergroup_count_limit"
	ServergroupCountLimit param.Opt[int64] `json:"servergroup_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/sfs_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.sfs_count_limit"
	SfsCountLimit param.Opt[int64] `json:"sfs_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/sfs_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.sfs_size_limit"
	SfsSizeLimit param.Opt[int64] `json:"sfs_size_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/shared_vm_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.shared_vm_count_limit"
	SharedVmCountLimit param.Opt[int64] `json:"shared_vm_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/snapshot_schedule_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.snapshot_schedule_count_limit"
	SnapshotScheduleCountLimit param.Opt[int64] `json:"snapshot_schedule_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/subnet_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.subnet_count_limit"
	SubnetCountLimit param.Opt[int64] `json:"subnet_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/vm_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.vm_count_limit"
	VmCountLimit param.Opt[int64] `json:"vm_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_count_limit"
	VolumeCountLimit param.Opt[int64] `json:"volume_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_size_limit"
	VolumeSizeLimit param.Opt[int64] `json:"volume_size_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_snapshots_count_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_snapshots_count_limit"
	VolumeSnapshotsCountLimit param.Opt[int64] `json:"volume_snapshots_count_limit,omitzero"`
	// '#/components/schemas/RegionalQuotasLimitsSerializer/properties/volume_snapshots_size_limit'
	// "$.components.schemas.RegionalQuotasLimitsSerializer.properties.volume_snapshots_size_limit"
	VolumeSnapshotsSizeLimit param.Opt[int64] `json:"volume_snapshots_size_limit,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaRequestNewParamsRequestedLimitsRegionalLimit) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r QuotaRequestNewParamsRequestedLimitsRegionalLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimitsRegionalLimit
	return param.MarshalObject(r, (*shadow)(&r))
}

type QuotaRequestListParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Flimits_request/get/parameters/0'
	// "$.paths['/cloud/v2/limits_request'].get.parameters[0]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Flimits_request/get/parameters/1'
	// "$.paths['/cloud/v2/limits_request'].get.parameters[1]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Flimits_request/get/parameters/2/schema/anyOf/0'
	// "$.paths['/cloud/v2/limits_request'].get.parameters[2].schema.anyOf[0]"
	//
	// Any of "done", "in progress", "rejected".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaRequestListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [QuotaRequestListParams]'s query parameters as `url.Values`.
func (r QuotaRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
