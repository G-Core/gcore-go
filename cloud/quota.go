// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// QuotaService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options  []option.RequestOption
	Requests QuotaRequestService
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r QuotaService) {
	r = QuotaService{}
	r.Options = opts
	r.Requests = NewQuotaRequestService(opts...)
	return
}

// Get combined client quotas, regional and global.
func (r *QuotaService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *QuotaGetAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/client_quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a quota by region
func (r *QuotaService) GetByRegion(ctx context.Context, query QuotaGetByRegionParams, opts ...option.RequestOption) (res *QuotaGetByRegionResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/regional_quotas/%v/%v", query.ClientID, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get global quota
func (r *QuotaService) GetGlobal(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *QuotaGetGlobalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/global_quotas/%v", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type QuotaGetAllResponse struct {
	// Global entity quotas
	GlobalQuotas QuotaGetAllResponseGlobalQuotas `json:"global_quotas"`
	// Regional entity quotas. Only contains initialized quotas.
	RegionalQuotas []QuotaGetAllResponseRegionalQuota `json:"regional_quotas"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GlobalQuotas   resp.Field
		RegionalQuotas resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetAllResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global entity quotas
type QuotaGetAllResponseGlobalQuotas struct {
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
func (r QuotaGetAllResponseGlobalQuotas) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponseGlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetAllResponseRegionalQuota struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// AI GPU bare metal servers count usage
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// Containers count usage
	CaasContainerCountUsage int64 `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// mCPU count for containers usage
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// Containers gpu count usage
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage"`
	// MB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// MB memory count for containers usage
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// External IP Count usage
	ExternalIPCountUsage int64 `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// mCPU count for functions usage
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// Functions count usage
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// Functions namespace count usage
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage"`
	// MB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// MB memory count for functions usage
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit"`
	// Images Count usage
	ImageCountUsage int64 `json:"image_count_usage"`
	// Images Size, GiB limit
	ImageSizeLimit int64 `json:"image_size_limit"`
	// Images Size, GiB usage
	ImageSizeUsage int64 `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// IPU Count usage
	IpuCountUsage int64 `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit"`
	// Networks Count usage
	NetworkCountUsage int64 `json:"network_count_usage"`
	// RAM Size, GiB limit
	RamLimit int64 `json:"ram_limit"`
	// RAM Size, GiB usage
	RamUsage int64 `json:"ram_usage"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// Registries count usage
	RegistryCountUsage int64 `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage int64 `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit"`
	// Routers Count usage
	RouterCountUsage int64 `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit"`
	// Secret Count usage
	SecretCountUsage int64 `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64 `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// Basic VMs Count usage
	SharedVmCountUsage int64 `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// Subnets Count usage
	SubnetCountUsage int64 `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// Instances Dedicated Count usage
	VmCountUsage int64 `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaremetalBasicCountLimit          resp.Field
		BaremetalBasicCountUsage          resp.Field
		BaremetalGPUCountLimit            resp.Field
		BaremetalGPUCountUsage            resp.Field
		BaremetalHfCountLimit             resp.Field
		BaremetalHfCountUsage             resp.Field
		BaremetalInfrastructureCountLimit resp.Field
		BaremetalInfrastructureCountUsage resp.Field
		BaremetalNetworkCountLimit        resp.Field
		BaremetalNetworkCountUsage        resp.Field
		BaremetalStorageCountLimit        resp.Field
		BaremetalStorageCountUsage        resp.Field
		CaasContainerCountLimit           resp.Field
		CaasContainerCountUsage           resp.Field
		CaasCPUCountLimit                 resp.Field
		CaasCPUCountUsage                 resp.Field
		CaasGPUCountLimit                 resp.Field
		CaasGPUCountUsage                 resp.Field
		CaasRamSizeLimit                  resp.Field
		CaasRamSizeUsage                  resp.Field
		ClusterCountLimit                 resp.Field
		ClusterCountUsage                 resp.Field
		CPUCountLimit                     resp.Field
		CPUCountUsage                     resp.Field
		DbaasPostgresClusterCountLimit    resp.Field
		DbaasPostgresClusterCountUsage    resp.Field
		ExternalIPCountLimit              resp.Field
		ExternalIPCountUsage              resp.Field
		FaasCPUCountLimit                 resp.Field
		FaasCPUCountUsage                 resp.Field
		FaasFunctionCountLimit            resp.Field
		FaasFunctionCountUsage            resp.Field
		FaasNamespaceCountLimit           resp.Field
		FaasNamespaceCountUsage           resp.Field
		FaasRamSizeLimit                  resp.Field
		FaasRamSizeUsage                  resp.Field
		FirewallCountLimit                resp.Field
		FirewallCountUsage                resp.Field
		FloatingCountLimit                resp.Field
		FloatingCountUsage                resp.Field
		GPUCountLimit                     resp.Field
		GPUCountUsage                     resp.Field
		GPUVirtualA100CountLimit          resp.Field
		GPUVirtualA100CountUsage          resp.Field
		GPUVirtualH100CountLimit          resp.Field
		GPUVirtualH100CountUsage          resp.Field
		GPUVirtualL40sCountLimit          resp.Field
		GPUVirtualL40sCountUsage          resp.Field
		ImageCountLimit                   resp.Field
		ImageCountUsage                   resp.Field
		ImageSizeLimit                    resp.Field
		ImageSizeUsage                    resp.Field
		IpuCountLimit                     resp.Field
		IpuCountUsage                     resp.Field
		LaasTopicCountLimit               resp.Field
		LaasTopicCountUsage               resp.Field
		LoadbalancerCountLimit            resp.Field
		LoadbalancerCountUsage            resp.Field
		NetworkCountLimit                 resp.Field
		NetworkCountUsage                 resp.Field
		RamLimit                          resp.Field
		RamUsage                          resp.Field
		RegionID                          resp.Field
		RegistryCountLimit                resp.Field
		RegistryCountUsage                resp.Field
		RegistryStorageLimit              resp.Field
		RegistryStorageUsage              resp.Field
		RouterCountLimit                  resp.Field
		RouterCountUsage                  resp.Field
		SecretCountLimit                  resp.Field
		SecretCountUsage                  resp.Field
		ServergroupCountLimit             resp.Field
		ServergroupCountUsage             resp.Field
		SfsCountLimit                     resp.Field
		SfsCountUsage                     resp.Field
		SfsSizeLimit                      resp.Field
		SfsSizeUsage                      resp.Field
		SharedVmCountLimit                resp.Field
		SharedVmCountUsage                resp.Field
		SnapshotScheduleCountLimit        resp.Field
		SnapshotScheduleCountUsage        resp.Field
		SubnetCountLimit                  resp.Field
		SubnetCountUsage                  resp.Field
		VmCountLimit                      resp.Field
		VmCountUsage                      resp.Field
		VolumeCountLimit                  resp.Field
		VolumeCountUsage                  resp.Field
		VolumeSizeLimit                   resp.Field
		VolumeSizeUsage                   resp.Field
		VolumeSnapshotsCountLimit         resp.Field
		VolumeSnapshotsCountUsage         resp.Field
		VolumeSnapshotsSizeLimit          resp.Field
		VolumeSnapshotsSizeUsage          resp.Field
		ExtraFields                       map[string]resp.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetAllResponseRegionalQuota) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponseRegionalQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetByRegionResponse struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// AI GPU bare metal servers count usage
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// Containers count usage
	CaasContainerCountUsage int64 `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// mCPU count for containers usage
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// Containers gpu count usage
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage"`
	// MB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// MB memory count for containers usage
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// External IP Count usage
	ExternalIPCountUsage int64 `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// mCPU count for functions usage
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// Functions count usage
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// Functions namespace count usage
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage"`
	// MB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// MB memory count for functions usage
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit"`
	// Images Count usage
	ImageCountUsage int64 `json:"image_count_usage"`
	// Images Size, GiB limit
	ImageSizeLimit int64 `json:"image_size_limit"`
	// Images Size, GiB usage
	ImageSizeUsage int64 `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// IPU Count usage
	IpuCountUsage int64 `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit"`
	// Networks Count usage
	NetworkCountUsage int64 `json:"network_count_usage"`
	// RAM Size, GiB limit
	RamLimit int64 `json:"ram_limit"`
	// RAM Size, GiB usage
	RamUsage int64 `json:"ram_usage"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// Registries count usage
	RegistryCountUsage int64 `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage int64 `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit"`
	// Routers Count usage
	RouterCountUsage int64 `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit"`
	// Secret Count usage
	SecretCountUsage int64 `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64 `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// Basic VMs Count usage
	SharedVmCountUsage int64 `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// Subnets Count usage
	SubnetCountUsage int64 `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// Instances Dedicated Count usage
	VmCountUsage int64 `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaremetalBasicCountLimit          resp.Field
		BaremetalBasicCountUsage          resp.Field
		BaremetalGPUCountLimit            resp.Field
		BaremetalGPUCountUsage            resp.Field
		BaremetalHfCountLimit             resp.Field
		BaremetalHfCountUsage             resp.Field
		BaremetalInfrastructureCountLimit resp.Field
		BaremetalInfrastructureCountUsage resp.Field
		BaremetalNetworkCountLimit        resp.Field
		BaremetalNetworkCountUsage        resp.Field
		BaremetalStorageCountLimit        resp.Field
		BaremetalStorageCountUsage        resp.Field
		CaasContainerCountLimit           resp.Field
		CaasContainerCountUsage           resp.Field
		CaasCPUCountLimit                 resp.Field
		CaasCPUCountUsage                 resp.Field
		CaasGPUCountLimit                 resp.Field
		CaasGPUCountUsage                 resp.Field
		CaasRamSizeLimit                  resp.Field
		CaasRamSizeUsage                  resp.Field
		ClusterCountLimit                 resp.Field
		ClusterCountUsage                 resp.Field
		CPUCountLimit                     resp.Field
		CPUCountUsage                     resp.Field
		DbaasPostgresClusterCountLimit    resp.Field
		DbaasPostgresClusterCountUsage    resp.Field
		ExternalIPCountLimit              resp.Field
		ExternalIPCountUsage              resp.Field
		FaasCPUCountLimit                 resp.Field
		FaasCPUCountUsage                 resp.Field
		FaasFunctionCountLimit            resp.Field
		FaasFunctionCountUsage            resp.Field
		FaasNamespaceCountLimit           resp.Field
		FaasNamespaceCountUsage           resp.Field
		FaasRamSizeLimit                  resp.Field
		FaasRamSizeUsage                  resp.Field
		FirewallCountLimit                resp.Field
		FirewallCountUsage                resp.Field
		FloatingCountLimit                resp.Field
		FloatingCountUsage                resp.Field
		GPUCountLimit                     resp.Field
		GPUCountUsage                     resp.Field
		GPUVirtualA100CountLimit          resp.Field
		GPUVirtualA100CountUsage          resp.Field
		GPUVirtualH100CountLimit          resp.Field
		GPUVirtualH100CountUsage          resp.Field
		GPUVirtualL40sCountLimit          resp.Field
		GPUVirtualL40sCountUsage          resp.Field
		ImageCountLimit                   resp.Field
		ImageCountUsage                   resp.Field
		ImageSizeLimit                    resp.Field
		ImageSizeUsage                    resp.Field
		IpuCountLimit                     resp.Field
		IpuCountUsage                     resp.Field
		LaasTopicCountLimit               resp.Field
		LaasTopicCountUsage               resp.Field
		LoadbalancerCountLimit            resp.Field
		LoadbalancerCountUsage            resp.Field
		NetworkCountLimit                 resp.Field
		NetworkCountUsage                 resp.Field
		RamLimit                          resp.Field
		RamUsage                          resp.Field
		RegionID                          resp.Field
		RegistryCountLimit                resp.Field
		RegistryCountUsage                resp.Field
		RegistryStorageLimit              resp.Field
		RegistryStorageUsage              resp.Field
		RouterCountLimit                  resp.Field
		RouterCountUsage                  resp.Field
		SecretCountLimit                  resp.Field
		SecretCountUsage                  resp.Field
		ServergroupCountLimit             resp.Field
		ServergroupCountUsage             resp.Field
		SfsCountLimit                     resp.Field
		SfsCountUsage                     resp.Field
		SfsSizeLimit                      resp.Field
		SfsSizeUsage                      resp.Field
		SharedVmCountLimit                resp.Field
		SharedVmCountUsage                resp.Field
		SnapshotScheduleCountLimit        resp.Field
		SnapshotScheduleCountUsage        resp.Field
		SubnetCountLimit                  resp.Field
		SubnetCountUsage                  resp.Field
		VmCountLimit                      resp.Field
		VmCountUsage                      resp.Field
		VolumeCountLimit                  resp.Field
		VolumeCountUsage                  resp.Field
		VolumeSizeLimit                   resp.Field
		VolumeSizeUsage                   resp.Field
		VolumeSnapshotsCountLimit         resp.Field
		VolumeSnapshotsCountUsage         resp.Field
		VolumeSnapshotsSizeLimit          resp.Field
		VolumeSnapshotsSizeUsage          resp.Field
		ExtraFields                       map[string]resp.Field
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
func (r QuotaGetGlobalResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetGlobalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetByRegionParams struct {
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Client ID
	ClientID int64 `path:"client_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaGetByRegionParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
