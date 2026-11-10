// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Quotas define the maximum amount of cloud resources (compute, storage,
// networking, GPU, and more) available to a client, both globally and per region.
//
// QuotaService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options               []option.RequestOption
	Requests              QuotaRequestService
	NotificationThreshold QuotaNotificationThresholdService
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r QuotaService) {
	r = QuotaService{}
	r.Options = opts
	r.Requests = NewQuotaRequestService(opts...)
	r.NotificationThreshold = NewQuotaNotificationThresholdService(opts...)
	return
}

// Get combined client quotas, including both regional and global quotas.
func (r *QuotaService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *Quota, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v2/client_quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get quotas for a specific region and client.
func (r *QuotaService) GetByRegion(ctx context.Context, query QuotaGetByRegionParams, opts ...option.RequestOption) (res *QuotaGetByRegionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/regional_quotas/%v/%v", query.ClientID, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get global quotas for a specific client.
func (r *QuotaService) GetGlobal(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *QuotaGetGlobalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cloud/v2/global_quotas/%v", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Quota struct {
	// Global entity quotas
	GlobalQuotas QuotaGlobalQuotas `json:"global_quotas" api:"required"`
	// Regional entity quotas. Only contains initialized quotas.
	RegionalQuotas []QuotaRegionalQuota `json:"regional_quotas" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalQuotas   respjson.Field
		RegionalQuotas respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quota) RawJSON() string { return r.JSON.raw }
func (r *Quota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global entity quotas
type QuotaGlobalQuotas struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit" api:"required"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage int64 `json:"inference_cpu_millicore_count_usage" api:"required"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit" api:"required"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage int64 `json:"inference_gpu_a100_count_usage" api:"required"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit" api:"required"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage int64 `json:"inference_gpu_h100_count_usage" api:"required"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit" api:"required"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage int64 `json:"inference_gpu_l40s_count_usage" api:"required"`
	// Inference instance count limit
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit" api:"required"`
	// Inference instance count usage
	InferenceInstanceCountUsage int64 `json:"inference_instance_count_usage" api:"required"`
	// Public model API keys count limit
	InferencePublicModelAPIKeyCountLimit int64 `json:"inference_public_model_api_key_count_limit" api:"required"`
	// Public model API keys count usage
	InferencePublicModelAPIKeyCountUsage int64 `json:"inference_public_model_api_key_count_usage" api:"required"`
	// SSH Keys Count limit
	KeypairCountLimit int64 `json:"keypair_count_limit" api:"required"`
	// SSH Keys Count usage
	KeypairCountUsage int64 `json:"keypair_count_usage" api:"required"`
	// Projects Count limit
	ProjectCountLimit int64 `json:"project_count_limit" api:"required"`
	// Projects Count usage
	ProjectCountUsage int64 `json:"project_count_usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceCPUMillicoreCountLimit      respjson.Field
		InferenceCPUMillicoreCountUsage      respjson.Field
		InferenceGPUA100CountLimit           respjson.Field
		InferenceGPUA100CountUsage           respjson.Field
		InferenceGPUH100CountLimit           respjson.Field
		InferenceGPUH100CountUsage           respjson.Field
		InferenceGPUL40sCountLimit           respjson.Field
		InferenceGPUL40sCountUsage           respjson.Field
		InferenceInstanceCountLimit          respjson.Field
		InferenceInstanceCountUsage          respjson.Field
		InferencePublicModelAPIKeyCountLimit respjson.Field
		InferencePublicModelAPIKeyCountUsage respjson.Field
		KeypairCountLimit                    respjson.Field
		KeypairCountUsage                    respjson.Field
		ProjectCountLimit                    respjson.Field
		ProjectCountUsage                    respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGlobalQuotas) RawJSON() string { return r.JSON.raw }
func (r *QuotaGlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaRegionalQuota struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit" api:"required"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage" api:"required"`
	// Bare metal A100 GPU server count limit
	BaremetalGPUA100CountLimit int64 `json:"baremetal_gpu_a100_count_limit" api:"required"`
	// Bare metal A100 GPU server count usage
	BaremetalGPUA100CountUsage int64 `json:"baremetal_gpu_a100_count_usage" api:"required"`
	// Total number of AI GPU bare metal servers. This field is deprecated and is now
	// always calculated automatically as the sum of `baremetal_gpu_a100_count_limit`,
	// `baremetal_gpu_h100_count_limit`, `baremetal_gpu_h200_count_limit`, and
	// `baremetal_gpu_l40s_count_limit`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit" api:"required"`
	// Total number of AI GPU bare metal servers in use. This field is deprecated and
	// is now always calculated automatically as the sum of
	// `baremetal_gpu_a100_count_usage`, `baremetal_gpu_h100_count_usage`,
	// `baremetal_gpu_h200_count_usage`, and `baremetal_gpu_l40s_count_usage`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage" api:"required"`
	// Bare metal H100 GPU server count limit
	BaremetalGPUH100CountLimit int64 `json:"baremetal_gpu_h100_count_limit" api:"required"`
	// Bare metal H100 GPU server count usage
	BaremetalGPUH100CountUsage int64 `json:"baremetal_gpu_h100_count_usage" api:"required"`
	// Bare metal H200 GPU server count limit
	BaremetalGPUH200CountLimit int64 `json:"baremetal_gpu_h200_count_limit" api:"required"`
	// Bare metal H200 GPU server count usage
	BaremetalGPUH200CountUsage int64 `json:"baremetal_gpu_h200_count_usage" api:"required"`
	// Bare metal L40S GPU server count limit
	BaremetalGPUL40sCountLimit int64 `json:"baremetal_gpu_l40s_count_limit" api:"required"`
	// Bare metal L40S GPU server count usage
	BaremetalGPUL40sCountUsage int64 `json:"baremetal_gpu_l40s_count_usage" api:"required"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit" api:"required"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage" api:"required"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit" api:"required"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage" api:"required"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit" api:"required"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage" api:"required"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit" api:"required"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage" api:"required"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit" api:"required"`
	// Containers count usage
	CaasContainerCountUsage int64 `json:"caas_container_count_usage" api:"required"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit" api:"required"`
	// mCPU count for containers usage
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage" api:"required"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit" api:"required"`
	// Containers gpu count usage
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage" api:"required"`
	// MiB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit" api:"required"`
	// MiB memory count for containers usage
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage" api:"required"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit" api:"required"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage" api:"required"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit" api:"required"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage" api:"required"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit" api:"required"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage" api:"required"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit" api:"required"`
	// External IP Count usage
	ExternalIPCountUsage int64 `json:"external_ip_count_usage" api:"required"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit" api:"required"`
	// mCPU count for functions usage
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage" api:"required"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit" api:"required"`
	// Functions count usage
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage" api:"required"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit" api:"required"`
	// Functions namespace count usage
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage" api:"required"`
	// MiB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit" api:"required"`
	// MiB memory count for functions usage
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage" api:"required"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit" api:"required"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage" api:"required"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit" api:"required"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage" api:"required"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit" api:"required"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage" api:"required"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit" api:"required"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage" api:"required"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit" api:"required"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage" api:"required"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit int64 `json:"gpu_virtual_h200_count_limit" api:"required"`
	// Virtual H200 GPU card count usage
	GPUVirtualH200CountUsage int64 `json:"gpu_virtual_h200_count_usage" api:"required"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit" api:"required"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage" api:"required"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit" api:"required"`
	// Images Count usage
	ImageCountUsage int64 `json:"image_count_usage" api:"required"`
	// Images Size, bytes limit
	ImageSizeLimit int64 `json:"image_size_limit" api:"required"`
	// Images Size, bytes usage
	ImageSizeUsage int64 `json:"image_size_usage" api:"required"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit" api:"required"`
	// IPU Count usage
	IpuCountUsage int64 `json:"ipu_count_usage" api:"required"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit" api:"required"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage" api:"required"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit" api:"required"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage" api:"required"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit" api:"required"`
	// Networks Count usage
	NetworkCountUsage int64 `json:"network_count_usage" api:"required"`
	// RAM Size, MiB limit
	RamLimit int64 `json:"ram_limit" api:"required"`
	// RAM Size, MiB usage
	RamUsage int64 `json:"ram_usage" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit" api:"required"`
	// Registries count usage
	RegistryCountUsage int64 `json:"registry_count_usage" api:"required"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit" api:"required"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage int64 `json:"registry_storage_usage" api:"required"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit" api:"required"`
	// Routers Count usage
	RouterCountUsage int64 `json:"router_count_usage" api:"required"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit" api:"required"`
	// Secret Count usage
	SecretCountUsage int64 `json:"secret_count_usage" api:"required"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit" api:"required"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage" api:"required"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit" api:"required"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage" api:"required"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit" api:"required"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64 `json:"sfs_size_usage" api:"required"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit" api:"required"`
	// Basic VMs Count usage
	SharedVmCountUsage int64 `json:"shared_vm_count_usage" api:"required"`
	// Slurm cluster count limit for this region. 0 disables Slurm in this region.
	SlurmClusterCountLimit int64 `json:"slurm_cluster_count_limit" api:"required"`
	// Slurm cluster count usage
	SlurmClusterCountUsage int64 `json:"slurm_cluster_count_usage" api:"required"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit" api:"required"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage" api:"required"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit" api:"required"`
	// Subnets Count usage
	SubnetCountUsage int64 `json:"subnet_count_usage" api:"required"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit" api:"required"`
	// Instances Dedicated Count usage
	VmCountUsage int64 `json:"vm_count_usage" api:"required"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit" api:"required"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage" api:"required"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit" api:"required"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage" api:"required"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit" api:"required"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage" api:"required"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit" api:"required"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaremetalBasicCountLimit          respjson.Field
		BaremetalBasicCountUsage          respjson.Field
		BaremetalGPUA100CountLimit        respjson.Field
		BaremetalGPUA100CountUsage        respjson.Field
		BaremetalGPUCountLimit            respjson.Field
		BaremetalGPUCountUsage            respjson.Field
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
		RegionID                          respjson.Field
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
		SlurmClusterCountLimit            respjson.Field
		SlurmClusterCountUsage            respjson.Field
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
func (r QuotaRegionalQuota) RawJSON() string { return r.JSON.raw }
func (r *QuotaRegionalQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetByRegionResponse struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit" api:"required"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage" api:"required"`
	// Bare metal A100 GPU server count limit
	BaremetalGPUA100CountLimit int64 `json:"baremetal_gpu_a100_count_limit" api:"required"`
	// Bare metal A100 GPU server count usage
	BaremetalGPUA100CountUsage int64 `json:"baremetal_gpu_a100_count_usage" api:"required"`
	// Total number of AI GPU bare metal servers. This field is deprecated and is now
	// always calculated automatically as the sum of `baremetal_gpu_a100_count_limit`,
	// `baremetal_gpu_h100_count_limit`, `baremetal_gpu_h200_count_limit`, and
	// `baremetal_gpu_l40s_count_limit`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit" api:"required"`
	// Total number of AI GPU bare metal servers in use. This field is deprecated and
	// is now always calculated automatically as the sum of
	// `baremetal_gpu_a100_count_usage`, `baremetal_gpu_h100_count_usage`,
	// `baremetal_gpu_h200_count_usage`, and `baremetal_gpu_l40s_count_usage`.
	//
	// Deprecated: deprecated
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage" api:"required"`
	// Bare metal H100 GPU server count limit
	BaremetalGPUH100CountLimit int64 `json:"baremetal_gpu_h100_count_limit" api:"required"`
	// Bare metal H100 GPU server count usage
	BaremetalGPUH100CountUsage int64 `json:"baremetal_gpu_h100_count_usage" api:"required"`
	// Bare metal H200 GPU server count limit
	BaremetalGPUH200CountLimit int64 `json:"baremetal_gpu_h200_count_limit" api:"required"`
	// Bare metal H200 GPU server count usage
	BaremetalGPUH200CountUsage int64 `json:"baremetal_gpu_h200_count_usage" api:"required"`
	// Bare metal L40S GPU server count limit
	BaremetalGPUL40sCountLimit int64 `json:"baremetal_gpu_l40s_count_limit" api:"required"`
	// Bare metal L40S GPU server count usage
	BaremetalGPUL40sCountUsage int64 `json:"baremetal_gpu_l40s_count_usage" api:"required"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit" api:"required"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage" api:"required"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit" api:"required"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage" api:"required"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit" api:"required"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage" api:"required"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit" api:"required"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage" api:"required"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit" api:"required"`
	// Containers count usage
	CaasContainerCountUsage int64 `json:"caas_container_count_usage" api:"required"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit" api:"required"`
	// mCPU count for containers usage
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage" api:"required"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit" api:"required"`
	// Containers gpu count usage
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage" api:"required"`
	// MiB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit" api:"required"`
	// MiB memory count for containers usage
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage" api:"required"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit" api:"required"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage" api:"required"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit" api:"required"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage" api:"required"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit" api:"required"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage" api:"required"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit" api:"required"`
	// External IP Count usage
	ExternalIPCountUsage int64 `json:"external_ip_count_usage" api:"required"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit" api:"required"`
	// mCPU count for functions usage
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage" api:"required"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit" api:"required"`
	// Functions count usage
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage" api:"required"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit" api:"required"`
	// Functions namespace count usage
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage" api:"required"`
	// MiB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit" api:"required"`
	// MiB memory count for functions usage
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage" api:"required"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit" api:"required"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage" api:"required"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit" api:"required"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage" api:"required"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit" api:"required"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage" api:"required"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit" api:"required"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage" api:"required"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit" api:"required"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage" api:"required"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit int64 `json:"gpu_virtual_h200_count_limit" api:"required"`
	// Virtual H200 GPU card count usage
	GPUVirtualH200CountUsage int64 `json:"gpu_virtual_h200_count_usage" api:"required"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit" api:"required"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage" api:"required"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit" api:"required"`
	// Images Count usage
	ImageCountUsage int64 `json:"image_count_usage" api:"required"`
	// Images Size, bytes limit
	ImageSizeLimit int64 `json:"image_size_limit" api:"required"`
	// Images Size, bytes usage
	ImageSizeUsage int64 `json:"image_size_usage" api:"required"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit" api:"required"`
	// IPU Count usage
	IpuCountUsage int64 `json:"ipu_count_usage" api:"required"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit" api:"required"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage" api:"required"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit" api:"required"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage" api:"required"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit" api:"required"`
	// Networks Count usage
	NetworkCountUsage int64 `json:"network_count_usage" api:"required"`
	// RAM Size, MiB limit
	RamLimit int64 `json:"ram_limit" api:"required"`
	// RAM Size, MiB usage
	RamUsage int64 `json:"ram_usage" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit" api:"required"`
	// Registries count usage
	RegistryCountUsage int64 `json:"registry_count_usage" api:"required"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit" api:"required"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage int64 `json:"registry_storage_usage" api:"required"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit" api:"required"`
	// Routers Count usage
	RouterCountUsage int64 `json:"router_count_usage" api:"required"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit" api:"required"`
	// Secret Count usage
	SecretCountUsage int64 `json:"secret_count_usage" api:"required"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit" api:"required"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage" api:"required"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit" api:"required"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage" api:"required"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit" api:"required"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64 `json:"sfs_size_usage" api:"required"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit" api:"required"`
	// Basic VMs Count usage
	SharedVmCountUsage int64 `json:"shared_vm_count_usage" api:"required"`
	// Slurm cluster count limit for this region. 0 disables Slurm in this region.
	SlurmClusterCountLimit int64 `json:"slurm_cluster_count_limit" api:"required"`
	// Slurm cluster count usage
	SlurmClusterCountUsage int64 `json:"slurm_cluster_count_usage" api:"required"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit" api:"required"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage" api:"required"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit" api:"required"`
	// Subnets Count usage
	SubnetCountUsage int64 `json:"subnet_count_usage" api:"required"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit" api:"required"`
	// Instances Dedicated Count usage
	VmCountUsage int64 `json:"vm_count_usage" api:"required"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit" api:"required"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage" api:"required"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit" api:"required"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage" api:"required"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit" api:"required"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage" api:"required"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit" api:"required"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaremetalBasicCountLimit          respjson.Field
		BaremetalBasicCountUsage          respjson.Field
		BaremetalGPUA100CountLimit        respjson.Field
		BaremetalGPUA100CountUsage        respjson.Field
		BaremetalGPUCountLimit            respjson.Field
		BaremetalGPUCountUsage            respjson.Field
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
		RegionID                          respjson.Field
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
		SlurmClusterCountLimit            respjson.Field
		SlurmClusterCountUsage            respjson.Field
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
func (r QuotaGetByRegionResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetByRegionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetGlobalResponse struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit" api:"required"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage int64 `json:"inference_cpu_millicore_count_usage" api:"required"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit" api:"required"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage int64 `json:"inference_gpu_a100_count_usage" api:"required"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit" api:"required"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage int64 `json:"inference_gpu_h100_count_usage" api:"required"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit" api:"required"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage int64 `json:"inference_gpu_l40s_count_usage" api:"required"`
	// Inference instance count limit
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit" api:"required"`
	// Inference instance count usage
	InferenceInstanceCountUsage int64 `json:"inference_instance_count_usage" api:"required"`
	// Public model API keys count limit
	InferencePublicModelAPIKeyCountLimit int64 `json:"inference_public_model_api_key_count_limit" api:"required"`
	// Public model API keys count usage
	InferencePublicModelAPIKeyCountUsage int64 `json:"inference_public_model_api_key_count_usage" api:"required"`
	// SSH Keys Count limit
	KeypairCountLimit int64 `json:"keypair_count_limit" api:"required"`
	// SSH Keys Count usage
	KeypairCountUsage int64 `json:"keypair_count_usage" api:"required"`
	// Projects Count limit
	ProjectCountLimit int64 `json:"project_count_limit" api:"required"`
	// Projects Count usage
	ProjectCountUsage int64 `json:"project_count_usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceCPUMillicoreCountLimit      respjson.Field
		InferenceCPUMillicoreCountUsage      respjson.Field
		InferenceGPUA100CountLimit           respjson.Field
		InferenceGPUA100CountUsage           respjson.Field
		InferenceGPUH100CountLimit           respjson.Field
		InferenceGPUH100CountUsage           respjson.Field
		InferenceGPUL40sCountLimit           respjson.Field
		InferenceGPUL40sCountUsage           respjson.Field
		InferenceInstanceCountLimit          respjson.Field
		InferenceInstanceCountUsage          respjson.Field
		InferencePublicModelAPIKeyCountLimit respjson.Field
		InferencePublicModelAPIKeyCountUsage respjson.Field
		KeypairCountLimit                    respjson.Field
		KeypairCountUsage                    respjson.Field
		ProjectCountLimit                    respjson.Field
		ProjectCountUsage                    respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetGlobalResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetGlobalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetByRegionParams struct {
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Client ID
	ClientID int64 `path:"client_id" api:"required" json:"-"`
	paramObj
}
