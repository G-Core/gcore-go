// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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

// Create a request to change current quotas.
func (r *QuotaRequestService) New(ctx context.Context, body QuotaRequestNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v2/limits_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a list of sent requests to change current quotas and their statuses.
func (r *QuotaRequestService) List(ctx context.Context, query QuotaRequestListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[QuotaRequestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v2/limits_request"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of sent requests to change current quotas and their statuses.
func (r *QuotaRequestService) ListAutoPaging(ctx context.Context, query QuotaRequestListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[QuotaRequestListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a specific quota limit request.
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

// Get detailed information about a specific quota limit request.
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

type QuotaRequestListResponse struct {
	// Request ID
	ID int64 `json:"id,required"`
	// Client ID
	ClientID int64 `json:"client_id,required"`
	// Requested limits.
	RequestedLimits QuotaRequestListResponseRequestedLimits `json:"requested_limits,required"`
	// Request status
	Status string `json:"status,required"`
	// Datetime when the request was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Describe the reason, in general terms.
	Description string `json:"description,nullable"`
	// Datetime when the request was updated.
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClientID        respjson.Field
		RequestedLimits respjson.Field
		Status          respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requested limits.
type QuotaRequestListResponseRequestedLimits struct {
	// Global entity quota limits
	GlobalLimits QuotaRequestListResponseRequestedLimitsGlobalLimits `json:"global_limits"`
	// Regions and their quota limits
	RegionalLimits []QuotaRequestListResponseRequestedLimitsRegionalLimit `json:"regional_limits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalLimits   respjson.Field
		RegionalLimits respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestListResponseRequestedLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestListResponseRequestedLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global entity quota limits
type QuotaRequestListResponseRequestedLimitsGlobalLimits struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// Inference instance count limit
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// SSH Keys Count limit
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// Projects Count limit
	ProjectCountLimit int64 `json:"project_count_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceCPUMillicoreCountLimit respjson.Field
		InferenceGPUA100CountLimit      respjson.Field
		InferenceGPUH100CountLimit      respjson.Field
		InferenceGPUL40sCountLimit      respjson.Field
		InferenceInstanceCountLimit     respjson.Field
		KeypairCountLimit               respjson.Field
		ProjectCountLimit               respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestListResponseRequestedLimitsGlobalLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestListResponseRequestedLimitsGlobalLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestListResponseRequestedLimitsRegionalLimit struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// Baremetal A100 GPU card count limit
	BaremetalGPUA100CountLimit int64 `json:"baremetal_gpu_a100_count_limit"`
	// Total number of AI GPU bare metal servers. This field is deprecated and is now
	// always calculated automatically as the sum of `baremetal_gpu_a100_count_limit`,
	// `baremetal_gpu_h100_count_limit`, `baremetal_gpu_h200_count_limit`, and
	// `baremetal_gpu_l40s_count_limit`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// Baremetal H100 GPU card count limit
	BaremetalGPUH100CountLimit int64 `json:"baremetal_gpu_h100_count_limit"`
	// Baremetal H200 GPU card count limit
	BaremetalGPUH200CountLimit int64 `json:"baremetal_gpu_h200_count_limit"`
	// Baremetal L40S GPU card count limit
	BaremetalGPUL40sCountLimit int64 `json:"baremetal_gpu_l40s_count_limit"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// MB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// MB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit int64 `json:"gpu_virtual_h200_count_limit"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit"`
	// Images Size, GiB limit
	ImageSizeLimit int64 `json:"image_size_limit"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit"`
	// RAM Size, GiB limit
	RamLimit int64 `json:"ram_limit"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaremetalBasicCountLimit          respjson.Field
		BaremetalGPUA100CountLimit        respjson.Field
		BaremetalGPUCountLimit            respjson.Field
		BaremetalGPUH100CountLimit        respjson.Field
		BaremetalGPUH200CountLimit        respjson.Field
		BaremetalGPUL40sCountLimit        respjson.Field
		BaremetalHfCountLimit             respjson.Field
		BaremetalInfrastructureCountLimit respjson.Field
		BaremetalNetworkCountLimit        respjson.Field
		BaremetalStorageCountLimit        respjson.Field
		CaasContainerCountLimit           respjson.Field
		CaasCPUCountLimit                 respjson.Field
		CaasGPUCountLimit                 respjson.Field
		CaasRamSizeLimit                  respjson.Field
		ClusterCountLimit                 respjson.Field
		CPUCountLimit                     respjson.Field
		DbaasPostgresClusterCountLimit    respjson.Field
		ExternalIPCountLimit              respjson.Field
		FaasCPUCountLimit                 respjson.Field
		FaasFunctionCountLimit            respjson.Field
		FaasNamespaceCountLimit           respjson.Field
		FaasRamSizeLimit                  respjson.Field
		FirewallCountLimit                respjson.Field
		FloatingCountLimit                respjson.Field
		GPUCountLimit                     respjson.Field
		GPUVirtualA100CountLimit          respjson.Field
		GPUVirtualH100CountLimit          respjson.Field
		GPUVirtualH200CountLimit          respjson.Field
		GPUVirtualL40sCountLimit          respjson.Field
		ImageCountLimit                   respjson.Field
		ImageSizeLimit                    respjson.Field
		IpuCountLimit                     respjson.Field
		LaasTopicCountLimit               respjson.Field
		LoadbalancerCountLimit            respjson.Field
		NetworkCountLimit                 respjson.Field
		RamLimit                          respjson.Field
		RegionID                          respjson.Field
		RegistryCountLimit                respjson.Field
		RegistryStorageLimit              respjson.Field
		RouterCountLimit                  respjson.Field
		SecretCountLimit                  respjson.Field
		ServergroupCountLimit             respjson.Field
		SfsCountLimit                     respjson.Field
		SfsSizeLimit                      respjson.Field
		SharedVmCountLimit                respjson.Field
		SnapshotScheduleCountLimit        respjson.Field
		SubnetCountLimit                  respjson.Field
		VmCountLimit                      respjson.Field
		VolumeCountLimit                  respjson.Field
		VolumeSizeLimit                   respjson.Field
		VolumeSnapshotsCountLimit         respjson.Field
		VolumeSnapshotsSizeLimit          respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestListResponseRequestedLimitsRegionalLimit) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestListResponseRequestedLimitsRegionalLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestGetResponse struct {
	// Request ID
	ID int64 `json:"id,required"`
	// Client ID
	ClientID int64 `json:"client_id,required"`
	// Requested limits.
	RequestedLimits QuotaRequestGetResponseRequestedLimits `json:"requested_limits,required"`
	// Request status
	Status string `json:"status,required"`
	// Datetime when the request was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Describe the reason, in general terms.
	Description string `json:"description,nullable"`
	// Datetime when the request was updated.
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClientID        respjson.Field
		RequestedLimits respjson.Field
		Status          respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requested limits.
type QuotaRequestGetResponseRequestedLimits struct {
	// Global entity quota limits
	GlobalLimits QuotaRequestGetResponseRequestedLimitsGlobalLimits `json:"global_limits"`
	// Regions and their quota limits
	RegionalLimits []QuotaRequestGetResponseRequestedLimitsRegionalLimit `json:"regional_limits"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalLimits   respjson.Field
		RegionalLimits respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global entity quota limits
type QuotaRequestGetResponseRequestedLimitsGlobalLimits struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// Inference instance count limit
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// SSH Keys Count limit
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// Projects Count limit
	ProjectCountLimit int64 `json:"project_count_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceCPUMillicoreCountLimit respjson.Field
		InferenceGPUA100CountLimit      respjson.Field
		InferenceGPUH100CountLimit      respjson.Field
		InferenceGPUL40sCountLimit      respjson.Field
		InferenceInstanceCountLimit     respjson.Field
		KeypairCountLimit               respjson.Field
		ProjectCountLimit               respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimitsGlobalLimits) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimitsGlobalLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestGetResponseRequestedLimitsRegionalLimit struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// Baremetal A100 GPU card count limit
	BaremetalGPUA100CountLimit int64 `json:"baremetal_gpu_a100_count_limit"`
	// Total number of AI GPU bare metal servers. This field is deprecated and is now
	// always calculated automatically as the sum of `baremetal_gpu_a100_count_limit`,
	// `baremetal_gpu_h100_count_limit`, `baremetal_gpu_h200_count_limit`, and
	// `baremetal_gpu_l40s_count_limit`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// Baremetal H100 GPU card count limit
	BaremetalGPUH100CountLimit int64 `json:"baremetal_gpu_h100_count_limit"`
	// Baremetal H200 GPU card count limit
	BaremetalGPUH200CountLimit int64 `json:"baremetal_gpu_h200_count_limit"`
	// Baremetal L40S GPU card count limit
	BaremetalGPUL40sCountLimit int64 `json:"baremetal_gpu_l40s_count_limit"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// MB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// MB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit int64 `json:"gpu_virtual_h200_count_limit"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit"`
	// Images Size, GiB limit
	ImageSizeLimit int64 `json:"image_size_limit"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit"`
	// RAM Size, GiB limit
	RamLimit int64 `json:"ram_limit"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaremetalBasicCountLimit          respjson.Field
		BaremetalGPUA100CountLimit        respjson.Field
		BaremetalGPUCountLimit            respjson.Field
		BaremetalGPUH100CountLimit        respjson.Field
		BaremetalGPUH200CountLimit        respjson.Field
		BaremetalGPUL40sCountLimit        respjson.Field
		BaremetalHfCountLimit             respjson.Field
		BaremetalInfrastructureCountLimit respjson.Field
		BaremetalNetworkCountLimit        respjson.Field
		BaremetalStorageCountLimit        respjson.Field
		CaasContainerCountLimit           respjson.Field
		CaasCPUCountLimit                 respjson.Field
		CaasGPUCountLimit                 respjson.Field
		CaasRamSizeLimit                  respjson.Field
		ClusterCountLimit                 respjson.Field
		CPUCountLimit                     respjson.Field
		DbaasPostgresClusterCountLimit    respjson.Field
		ExternalIPCountLimit              respjson.Field
		FaasCPUCountLimit                 respjson.Field
		FaasFunctionCountLimit            respjson.Field
		FaasNamespaceCountLimit           respjson.Field
		FaasRamSizeLimit                  respjson.Field
		FirewallCountLimit                respjson.Field
		FloatingCountLimit                respjson.Field
		GPUCountLimit                     respjson.Field
		GPUVirtualA100CountLimit          respjson.Field
		GPUVirtualH100CountLimit          respjson.Field
		GPUVirtualH200CountLimit          respjson.Field
		GPUVirtualL40sCountLimit          respjson.Field
		ImageCountLimit                   respjson.Field
		ImageSizeLimit                    respjson.Field
		IpuCountLimit                     respjson.Field
		LaasTopicCountLimit               respjson.Field
		LoadbalancerCountLimit            respjson.Field
		NetworkCountLimit                 respjson.Field
		RamLimit                          respjson.Field
		RegionID                          respjson.Field
		RegistryCountLimit                respjson.Field
		RegistryStorageLimit              respjson.Field
		RouterCountLimit                  respjson.Field
		SecretCountLimit                  respjson.Field
		ServergroupCountLimit             respjson.Field
		SfsCountLimit                     respjson.Field
		SfsSizeLimit                      respjson.Field
		SharedVmCountLimit                respjson.Field
		SnapshotScheduleCountLimit        respjson.Field
		SubnetCountLimit                  respjson.Field
		VmCountLimit                      respjson.Field
		VolumeCountLimit                  respjson.Field
		VolumeSizeLimit                   respjson.Field
		VolumeSnapshotsCountLimit         respjson.Field
		VolumeSnapshotsSizeLimit          respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaRequestGetResponseRequestedLimitsRegionalLimit) RawJSON() string { return r.JSON.raw }
func (r *QuotaRequestGetResponseRequestedLimitsRegionalLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestNewParams struct {
	// Describe the reason, in general terms.
	Description string `json:"description,required"`
	// Limits you want to increase.
	RequestedLimits QuotaRequestNewParamsRequestedLimits `json:"requested_limits,omitzero,required"`
	// Client ID that requests the limit increase.
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	paramObj
}

func (r QuotaRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaRequestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Limits you want to increase.
type QuotaRequestNewParamsRequestedLimits struct {
	// Global entity quota limits
	GlobalLimits QuotaRequestNewParamsRequestedLimitsGlobalLimits `json:"global_limits,omitzero"`
	// Regions and their quota limits
	RegionalLimits []QuotaRequestNewParamsRequestedLimitsRegionalLimit `json:"regional_limits,omitzero"`
	paramObj
}

func (r QuotaRequestNewParamsRequestedLimits) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimits
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaRequestNewParamsRequestedLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global entity quota limits
type QuotaRequestNewParamsRequestedLimitsGlobalLimits struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit param.Opt[int64] `json:"inference_cpu_millicore_count_limit,omitzero"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit param.Opt[int64] `json:"inference_gpu_a100_count_limit,omitzero"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit param.Opt[int64] `json:"inference_gpu_h100_count_limit,omitzero"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit param.Opt[int64] `json:"inference_gpu_l40s_count_limit,omitzero"`
	// Inference instance count limit
	InferenceInstanceCountLimit param.Opt[int64] `json:"inference_instance_count_limit,omitzero"`
	// SSH Keys Count limit
	KeypairCountLimit param.Opt[int64] `json:"keypair_count_limit,omitzero"`
	// Projects Count limit
	ProjectCountLimit param.Opt[int64] `json:"project_count_limit,omitzero"`
	paramObj
}

func (r QuotaRequestNewParamsRequestedLimitsGlobalLimits) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimitsGlobalLimits
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaRequestNewParamsRequestedLimitsGlobalLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestNewParamsRequestedLimitsRegionalLimit struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit param.Opt[int64] `json:"baremetal_basic_count_limit,omitzero"`
	// Baremetal A100 GPU card count limit
	BaremetalGPUA100CountLimit param.Opt[int64] `json:"baremetal_gpu_a100_count_limit,omitzero"`
	// Total number of AI GPU bare metal servers. This field is deprecated and is now
	// always calculated automatically as the sum of `baremetal_gpu_a100_count_limit`,
	// `baremetal_gpu_h100_count_limit`, `baremetal_gpu_h200_count_limit`, and
	// `baremetal_gpu_l40s_count_limit`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountLimit param.Opt[int64] `json:"baremetal_gpu_count_limit,omitzero"`
	// Baremetal H100 GPU card count limit
	BaremetalGPUH100CountLimit param.Opt[int64] `json:"baremetal_gpu_h100_count_limit,omitzero"`
	// Baremetal H200 GPU card count limit
	BaremetalGPUH200CountLimit param.Opt[int64] `json:"baremetal_gpu_h200_count_limit,omitzero"`
	// Baremetal L40S GPU card count limit
	BaremetalGPUL40sCountLimit param.Opt[int64] `json:"baremetal_gpu_l40s_count_limit,omitzero"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit param.Opt[int64] `json:"baremetal_hf_count_limit,omitzero"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit param.Opt[int64] `json:"baremetal_infrastructure_count_limit,omitzero"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit param.Opt[int64] `json:"baremetal_network_count_limit,omitzero"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit param.Opt[int64] `json:"baremetal_storage_count_limit,omitzero"`
	// Containers count limit
	CaasContainerCountLimit param.Opt[int64] `json:"caas_container_count_limit,omitzero"`
	// mCPU count for containers limit
	CaasCPUCountLimit param.Opt[int64] `json:"caas_cpu_count_limit,omitzero"`
	// Containers gpu count limit
	CaasGPUCountLimit param.Opt[int64] `json:"caas_gpu_count_limit,omitzero"`
	// MB memory count for containers limit
	CaasRamSizeLimit param.Opt[int64] `json:"caas_ram_size_limit,omitzero"`
	// K8s clusters count limit
	ClusterCountLimit param.Opt[int64] `json:"cluster_count_limit,omitzero"`
	// vCPU Count limit
	CPUCountLimit param.Opt[int64] `json:"cpu_count_limit,omitzero"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit param.Opt[int64] `json:"dbaas_postgres_cluster_count_limit,omitzero"`
	// External IP Count limit
	ExternalIPCountLimit param.Opt[int64] `json:"external_ip_count_limit,omitzero"`
	// mCPU count for functions limit
	FaasCPUCountLimit param.Opt[int64] `json:"faas_cpu_count_limit,omitzero"`
	// Functions count limit
	FaasFunctionCountLimit param.Opt[int64] `json:"faas_function_count_limit,omitzero"`
	// Functions namespace count limit
	FaasNamespaceCountLimit param.Opt[int64] `json:"faas_namespace_count_limit,omitzero"`
	// MB memory count for functions limit
	FaasRamSizeLimit param.Opt[int64] `json:"faas_ram_size_limit,omitzero"`
	// Firewalls Count limit
	FirewallCountLimit param.Opt[int64] `json:"firewall_count_limit,omitzero"`
	// Floating IP Count limit
	FloatingCountLimit param.Opt[int64] `json:"floating_count_limit,omitzero"`
	// GPU Count limit
	GPUCountLimit param.Opt[int64] `json:"gpu_count_limit,omitzero"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit param.Opt[int64] `json:"gpu_virtual_a100_count_limit,omitzero"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit param.Opt[int64] `json:"gpu_virtual_h100_count_limit,omitzero"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit param.Opt[int64] `json:"gpu_virtual_h200_count_limit,omitzero"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit param.Opt[int64] `json:"gpu_virtual_l40s_count_limit,omitzero"`
	// Images Count limit
	ImageCountLimit param.Opt[int64] `json:"image_count_limit,omitzero"`
	// Images Size, GiB limit
	ImageSizeLimit param.Opt[int64] `json:"image_size_limit,omitzero"`
	// IPU Count limit
	IpuCountLimit param.Opt[int64] `json:"ipu_count_limit,omitzero"`
	// LaaS Topics Count limit
	LaasTopicCountLimit param.Opt[int64] `json:"laas_topic_count_limit,omitzero"`
	// Load Balancers Count limit
	LoadbalancerCountLimit param.Opt[int64] `json:"loadbalancer_count_limit,omitzero"`
	// Networks Count limit
	NetworkCountLimit param.Opt[int64] `json:"network_count_limit,omitzero"`
	// RAM Size, GiB limit
	RamLimit param.Opt[int64] `json:"ram_limit,omitzero"`
	// Region ID
	RegionID param.Opt[int64] `json:"region_id,omitzero"`
	// Registries count limit
	RegistryCountLimit param.Opt[int64] `json:"registry_count_limit,omitzero"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit param.Opt[int64] `json:"registry_storage_limit,omitzero"`
	// Routers Count limit
	RouterCountLimit param.Opt[int64] `json:"router_count_limit,omitzero"`
	// Secret Count limit
	SecretCountLimit param.Opt[int64] `json:"secret_count_limit,omitzero"`
	// Placement Group Count limit
	ServergroupCountLimit param.Opt[int64] `json:"servergroup_count_limit,omitzero"`
	// Shared file system Count limit
	SfsCountLimit param.Opt[int64] `json:"sfs_count_limit,omitzero"`
	// Shared file system Size, GiB limit
	SfsSizeLimit param.Opt[int64] `json:"sfs_size_limit,omitzero"`
	// Basic VMs Count limit
	SharedVmCountLimit param.Opt[int64] `json:"shared_vm_count_limit,omitzero"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit param.Opt[int64] `json:"snapshot_schedule_count_limit,omitzero"`
	// Subnets Count limit
	SubnetCountLimit param.Opt[int64] `json:"subnet_count_limit,omitzero"`
	// Instances Dedicated Count limit
	VmCountLimit param.Opt[int64] `json:"vm_count_limit,omitzero"`
	// Volumes Count limit
	VolumeCountLimit param.Opt[int64] `json:"volume_count_limit,omitzero"`
	// Volumes Size, GiB limit
	VolumeSizeLimit param.Opt[int64] `json:"volume_size_limit,omitzero"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit param.Opt[int64] `json:"volume_snapshots_count_limit,omitzero"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit param.Opt[int64] `json:"volume_snapshots_size_limit,omitzero"`
	paramObj
}

func (r QuotaRequestNewParamsRequestedLimitsRegionalLimit) MarshalJSON() (data []byte, err error) {
	type shadow QuotaRequestNewParamsRequestedLimitsRegionalLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuotaRequestNewParamsRequestedLimitsRegionalLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRequestListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// List of limit requests statuses for filtering
	//
	// Any of "done", "in progress", "rejected".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuotaRequestListParams]'s query parameters as `url.Values`.
func (r QuotaRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
