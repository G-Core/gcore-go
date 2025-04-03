// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2LimitsRequestService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2LimitsRequestService] method instead.
type CloudV2LimitsRequestService struct {
	Options []option.RequestOption
}

// NewCloudV2LimitsRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2LimitsRequestService(opts ...option.RequestOption) (r *CloudV2LimitsRequestService) {
	r = &CloudV2LimitsRequestService{}
	r.Options = opts
	return
}

// Create request to change quotas
func (r *CloudV2LimitsRequestService) New(ctx context.Context, body CloudV2LimitsRequestNewParams, opts ...option.RequestOption) (res *LimitsRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/limits_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get request to change quota limits.
func (r *CloudV2LimitsRequestService) Get(ctx context.Context, reqID int64, opts ...option.RequestOption) (res *LimitsRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/limits_request/%v", reqID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of sent requests to change current quotas and their statuses
func (r *CloudV2LimitsRequestService) List(ctx context.Context, query CloudV2LimitsRequestListParams, opts ...option.RequestOption) (res *CloudV2LimitsRequestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/limits_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete request to change quotas
func (r *CloudV2LimitsRequestService) Delete(ctx context.Context, reqID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v2/limits_request/%v", reqID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type GlobalQuotasLimits struct {
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
	ProjectCountLimit int64                  `json:"project_count_limit"`
	JSON              globalQuotasLimitsJSON `json:"-"`
}

// globalQuotasLimitsJSON contains the JSON metadata for the struct
// [GlobalQuotasLimits]
type globalQuotasLimitsJSON struct {
	InferenceCPUMillicoreCountLimit apijson.Field
	InferenceGPUA100CountLimit      apijson.Field
	InferenceGPUH100CountLimit      apijson.Field
	InferenceGPUL40sCountLimit      apijson.Field
	InferenceInstanceCountLimit     apijson.Field
	KeypairCountLimit               apijson.Field
	ProjectCountLimit               apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *GlobalQuotasLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalQuotasLimitsJSON) RawJSON() string {
	return r.raw
}

type GlobalQuotasLimitsParam struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit param.Field[int64] `json:"inference_cpu_millicore_count_limit"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit param.Field[int64] `json:"inference_gpu_a100_count_limit"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit param.Field[int64] `json:"inference_gpu_h100_count_limit"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit param.Field[int64] `json:"inference_gpu_l40s_count_limit"`
	// Inference instance count limit
	InferenceInstanceCountLimit param.Field[int64] `json:"inference_instance_count_limit"`
	// SSH Keys Count limit
	KeypairCountLimit param.Field[int64] `json:"keypair_count_limit"`
	// Projects Count limit
	ProjectCountLimit param.Field[int64] `json:"project_count_limit"`
}

func (r GlobalQuotasLimitsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LimitsRequest struct {
	// Request ID
	ID int64 `json:"id,required"`
	// Client ID
	ClientID int64 `json:"client_id,required"`
	// Requested limits.
	RequestedLimits LimitsRequestRequestedLimits `json:"requested_limits,required"`
	// Request status
	Status string `json:"status,required"`
	// Datetime when the request was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Describe the reason, in general terms.
	Description string `json:"description,nullable"`
	// Datetime when the request was updated.
	UpdatedAt time.Time         `json:"updated_at,nullable" format:"date-time"`
	JSON      limitsRequestJSON `json:"-"`
}

// limitsRequestJSON contains the JSON metadata for the struct [LimitsRequest]
type limitsRequestJSON struct {
	ID              apijson.Field
	ClientID        apijson.Field
	RequestedLimits apijson.Field
	Status          apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LimitsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r limitsRequestJSON) RawJSON() string {
	return r.raw
}

// Requested limits.
type LimitsRequestRequestedLimits struct {
	// Global entity quota limits
	GlobalLimits GlobalQuotasLimits `json:"global_limits"`
	// Regions and their quota limits
	RegionalLimits []RegionalQuotasLimits           `json:"regional_limits"`
	JSON           limitsRequestRequestedLimitsJSON `json:"-"`
}

// limitsRequestRequestedLimitsJSON contains the JSON metadata for the struct
// [LimitsRequestRequestedLimits]
type limitsRequestRequestedLimitsJSON struct {
	GlobalLimits   apijson.Field
	RegionalLimits apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LimitsRequestRequestedLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r limitsRequestRequestedLimitsJSON) RawJSON() string {
	return r.raw
}

type RegionalQuotasLimits struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
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
	VolumeSnapshotsSizeLimit int64                    `json:"volume_snapshots_size_limit"`
	JSON                     regionalQuotasLimitsJSON `json:"-"`
}

// regionalQuotasLimitsJSON contains the JSON metadata for the struct
// [RegionalQuotasLimits]
type regionalQuotasLimitsJSON struct {
	BaremetalBasicCountLimit          apijson.Field
	BaremetalGPUCountLimit            apijson.Field
	BaremetalHfCountLimit             apijson.Field
	BaremetalInfrastructureCountLimit apijson.Field
	BaremetalNetworkCountLimit        apijson.Field
	BaremetalStorageCountLimit        apijson.Field
	CaasContainerCountLimit           apijson.Field
	CaasCPUCountLimit                 apijson.Field
	CaasGPUCountLimit                 apijson.Field
	CaasRamSizeLimit                  apijson.Field
	ClusterCountLimit                 apijson.Field
	CPUCountLimit                     apijson.Field
	DbaasPostgresClusterCountLimit    apijson.Field
	ExternalIPCountLimit              apijson.Field
	FaasCPUCountLimit                 apijson.Field
	FaasFunctionCountLimit            apijson.Field
	FaasNamespaceCountLimit           apijson.Field
	FaasRamSizeLimit                  apijson.Field
	FirewallCountLimit                apijson.Field
	FloatingCountLimit                apijson.Field
	GPUCountLimit                     apijson.Field
	ImageCountLimit                   apijson.Field
	ImageSizeLimit                    apijson.Field
	IpuCountLimit                     apijson.Field
	LaasTopicCountLimit               apijson.Field
	LoadbalancerCountLimit            apijson.Field
	NetworkCountLimit                 apijson.Field
	RamLimit                          apijson.Field
	RegionID                          apijson.Field
	RegistryCountLimit                apijson.Field
	RegistryStorageLimit              apijson.Field
	RouterCountLimit                  apijson.Field
	SecretCountLimit                  apijson.Field
	ServergroupCountLimit             apijson.Field
	SfsCountLimit                     apijson.Field
	SfsSizeLimit                      apijson.Field
	SharedVmCountLimit                apijson.Field
	SnapshotScheduleCountLimit        apijson.Field
	SubnetCountLimit                  apijson.Field
	VmCountLimit                      apijson.Field
	VolumeCountLimit                  apijson.Field
	VolumeSizeLimit                   apijson.Field
	VolumeSnapshotsCountLimit         apijson.Field
	VolumeSnapshotsSizeLimit          apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *RegionalQuotasLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalQuotasLimitsJSON) RawJSON() string {
	return r.raw
}

type RegionalQuotasLimitsParam struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit param.Field[int64] `json:"baremetal_basic_count_limit"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit param.Field[int64] `json:"baremetal_gpu_count_limit"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit param.Field[int64] `json:"baremetal_hf_count_limit"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit param.Field[int64] `json:"baremetal_infrastructure_count_limit"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit param.Field[int64] `json:"baremetal_network_count_limit"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit param.Field[int64] `json:"baremetal_storage_count_limit"`
	// Containers count limit
	CaasContainerCountLimit param.Field[int64] `json:"caas_container_count_limit"`
	// mCPU count for containers limit
	CaasCPUCountLimit param.Field[int64] `json:"caas_cpu_count_limit"`
	// Containers gpu count limit
	CaasGPUCountLimit param.Field[int64] `json:"caas_gpu_count_limit"`
	// MB memory count for containers limit
	CaasRamSizeLimit param.Field[int64] `json:"caas_ram_size_limit"`
	// K8s clusters count limit
	ClusterCountLimit param.Field[int64] `json:"cluster_count_limit"`
	// vCPU Count limit
	CPUCountLimit param.Field[int64] `json:"cpu_count_limit"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit param.Field[int64] `json:"dbaas_postgres_cluster_count_limit"`
	// External IP Count limit
	ExternalIPCountLimit param.Field[int64] `json:"external_ip_count_limit"`
	// mCPU count for functions limit
	FaasCPUCountLimit param.Field[int64] `json:"faas_cpu_count_limit"`
	// Functions count limit
	FaasFunctionCountLimit param.Field[int64] `json:"faas_function_count_limit"`
	// Functions namespace count limit
	FaasNamespaceCountLimit param.Field[int64] `json:"faas_namespace_count_limit"`
	// MB memory count for functions limit
	FaasRamSizeLimit param.Field[int64] `json:"faas_ram_size_limit"`
	// Firewalls Count limit
	FirewallCountLimit param.Field[int64] `json:"firewall_count_limit"`
	// Floating IP Count limit
	FloatingCountLimit param.Field[int64] `json:"floating_count_limit"`
	// GPU Count limit
	GPUCountLimit param.Field[int64] `json:"gpu_count_limit"`
	// Images Count limit
	ImageCountLimit param.Field[int64] `json:"image_count_limit"`
	// Images Size, GiB limit
	ImageSizeLimit param.Field[int64] `json:"image_size_limit"`
	// IPU Count limit
	IpuCountLimit param.Field[int64] `json:"ipu_count_limit"`
	// LaaS Topics Count limit
	LaasTopicCountLimit param.Field[int64] `json:"laas_topic_count_limit"`
	// Load Balancers Count limit
	LoadbalancerCountLimit param.Field[int64] `json:"loadbalancer_count_limit"`
	// Networks Count limit
	NetworkCountLimit param.Field[int64] `json:"network_count_limit"`
	// RAM Size, GiB limit
	RamLimit param.Field[int64] `json:"ram_limit"`
	// Region ID
	RegionID param.Field[int64] `json:"region_id"`
	// Registries count limit
	RegistryCountLimit param.Field[int64] `json:"registry_count_limit"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit param.Field[int64] `json:"registry_storage_limit"`
	// Routers Count limit
	RouterCountLimit param.Field[int64] `json:"router_count_limit"`
	// Secret Count limit
	SecretCountLimit param.Field[int64] `json:"secret_count_limit"`
	// Placement Group Count limit
	ServergroupCountLimit param.Field[int64] `json:"servergroup_count_limit"`
	// Shared file system Count limit
	SfsCountLimit param.Field[int64] `json:"sfs_count_limit"`
	// Shared file system Size, GiB limit
	SfsSizeLimit param.Field[int64] `json:"sfs_size_limit"`
	// Basic VMs Count limit
	SharedVmCountLimit param.Field[int64] `json:"shared_vm_count_limit"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit param.Field[int64] `json:"snapshot_schedule_count_limit"`
	// Subnets Count limit
	SubnetCountLimit param.Field[int64] `json:"subnet_count_limit"`
	// Instances Dedicated Count limit
	VmCountLimit param.Field[int64] `json:"vm_count_limit"`
	// Volumes Count limit
	VolumeCountLimit param.Field[int64] `json:"volume_count_limit"`
	// Volumes Size, GiB limit
	VolumeSizeLimit param.Field[int64] `json:"volume_size_limit"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit param.Field[int64] `json:"volume_snapshots_count_limit"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit param.Field[int64] `json:"volume_snapshots_size_limit"`
}

func (r RegionalQuotasLimitsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2LimitsRequestListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LimitsRequest                      `json:"results,required"`
	JSON    cloudV2LimitsRequestListResponseJSON `json:"-"`
}

// cloudV2LimitsRequestListResponseJSON contains the JSON metadata for the struct
// [CloudV2LimitsRequestListResponse]
type cloudV2LimitsRequestListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2LimitsRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2LimitsRequestListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2LimitsRequestNewParams struct {
	// Describe the reason, in general terms.
	Description param.Field[string] `json:"description,required"`
	// Limits you want to increase.
	RequestedLimits param.Field[CloudV2LimitsRequestNewParamsRequestedLimits] `json:"requested_limits,required"`
	// Client ID that requests the limit increase.
	ClientID param.Field[int64] `json:"client_id"`
}

func (r CloudV2LimitsRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits you want to increase.
type CloudV2LimitsRequestNewParamsRequestedLimits struct {
	// Global entity quota limits
	GlobalLimits param.Field[GlobalQuotasLimitsParam] `json:"global_limits"`
	// Regions and their quota limits
	RegionalLimits param.Field[[]RegionalQuotasLimitsParam] `json:"regional_limits"`
}

func (r CloudV2LimitsRequestNewParamsRequestedLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2LimitsRequestListParams struct {
	// Client ID to filter by. If not specified, inferred from jwt.
	ClientID param.Field[int64] `query:"client_id"`
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
	// Status of limit requests for filtering. Several options can be specified.
	Status param.Field[CloudV2LimitsRequestListParamsStatus] `query:"status"`
}

// URLQuery serializes [CloudV2LimitsRequestListParams]'s query parameters as
// `url.Values`.
func (r CloudV2LimitsRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status of limit requests for filtering. Several options can be specified.
type CloudV2LimitsRequestListParamsStatus string

const (
	CloudV2LimitsRequestListParamsStatusDone       CloudV2LimitsRequestListParamsStatus = "done"
	CloudV2LimitsRequestListParamsStatusInProgress CloudV2LimitsRequestListParamsStatus = "in progress"
	CloudV2LimitsRequestListParamsStatusRejected   CloudV2LimitsRequestListParamsStatus = "rejected"
)

func (r CloudV2LimitsRequestListParamsStatus) IsKnown() bool {
	switch r {
	case CloudV2LimitsRequestListParamsStatusDone, CloudV2LimitsRequestListParamsStatusInProgress, CloudV2LimitsRequestListParamsStatusRejected:
		return true
	}
	return false
}
