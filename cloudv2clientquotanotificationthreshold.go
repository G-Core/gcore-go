// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2ClientQuotaNotificationThresholdService contains methods and other
// services that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2ClientQuotaNotificationThresholdService] method instead.
type CloudV2ClientQuotaNotificationThresholdService struct {
	Options []option.RequestOption
}

// NewCloudV2ClientQuotaNotificationThresholdService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCloudV2ClientQuotaNotificationThresholdService(opts ...option.RequestOption) (r *CloudV2ClientQuotaNotificationThresholdService) {
	r = &CloudV2ClientQuotaNotificationThresholdService{}
	r.Options = opts
	return
}

// A quota notification threshold is necessary to send a notification warning to
// the client. Defaults to 80%. without their own threshold.
func (r *CloudV2ClientQuotaNotificationThresholdService) Get(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *NotificationThreshold, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete client's quota notification threshold
func (r *CloudV2ClientQuotaNotificationThresholdService) Delete(ctx context.Context, clientID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A quota notification threshold is necessary to send a notification warning to
// the client. Defaults to 80%.
func (r *CloudV2ClientQuotaNotificationThresholdService) UpdateOrNew(ctx context.Context, clientID int64, body CloudV2ClientQuotaNotificationThresholdUpdateOrNewParams, opts ...option.RequestOption) (res *NotificationThreshold, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/client_quotas/%v/notification_threshold", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type NotificationThreshold struct {
	// Client id
	ClientID int64 `json:"client_id,required"`
	// A message data
	LastMessage QuotasThreshold `json:"last_message,required,nullable"`
	// Time of last successful email sending
	LastSending time.Time `json:"last_sending,required,nullable" format:"date-time"`
	// Quota notification threshold in percentage
	Threshold int64                     `json:"threshold,required"`
	JSON      notificationThresholdJSON `json:"-"`
}

// notificationThresholdJSON contains the JSON metadata for the struct
// [NotificationThreshold]
type notificationThresholdJSON struct {
	ClientID    apijson.Field
	LastMessage apijson.Field
	LastSending apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationThresholdJSON) RawJSON() string {
	return r.raw
}

type QuotaCount struct {
	// Сurrent quota limit
	Limit int64 `json:"limit,required"`
	// Сurrent amount of resource used
	Usage int64          `json:"usage,required"`
	JSON  quotaCountJSON `json:"-"`
}

// quotaCountJSON contains the JSON metadata for the struct [QuotaCount]
type quotaCountJSON struct {
	Limit       apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaCountJSON) RawJSON() string {
	return r.raw
}

type QuotaCountParam struct {
	// Сurrent quota limit
	Limit param.Field[int64] `json:"limit,required"`
	// Сurrent amount of resource used
	Usage param.Field[int64] `json:"usage,required"`
}

func (r QuotaCountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QuotasThreshold struct {
	// Global quota that exceed the threshold
	GlobalQuotas QuotasThresholdGlobalQuotas `json:"global_quotas,required"`
	// Regional quota that exceed the threshold
	RegionalQuotas []QuotasThresholdRegionalQuota `json:"regional_quotas,required"`
	JSON           quotasThresholdJSON            `json:"-"`
}

// quotasThresholdJSON contains the JSON metadata for the struct [QuotasThreshold]
type quotasThresholdJSON struct {
	GlobalQuotas   apijson.Field
	RegionalQuotas apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QuotasThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotasThresholdJSON) RawJSON() string {
	return r.raw
}

// Global quota that exceed the threshold
type QuotasThresholdGlobalQuotas struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit QuotaCount `json:"inference_cpu_millicore_count_limit"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage QuotaCount `json:"inference_cpu_millicore_count_usage"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit QuotaCount `json:"inference_gpu_a100_count_limit"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage QuotaCount `json:"inference_gpu_a100_count_usage"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit QuotaCount `json:"inference_gpu_h100_count_limit"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage QuotaCount `json:"inference_gpu_h100_count_usage"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit QuotaCount `json:"inference_gpu_l40s_count_limit"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage QuotaCount `json:"inference_gpu_l40s_count_usage"`
	// Inference instance count limit
	InferenceInstanceCountLimit QuotaCount `json:"inference_instance_count_limit"`
	// Inference instance count usage
	InferenceInstanceCountUsage QuotaCount `json:"inference_instance_count_usage"`
	// SSH Keys Count limit
	KeypairCountLimit QuotaCount `json:"keypair_count_limit"`
	// SSH Keys Count usage
	KeypairCountUsage QuotaCount `json:"keypair_count_usage"`
	// Projects Count limit
	ProjectCountLimit QuotaCount `json:"project_count_limit"`
	// Projects Count usage
	ProjectCountUsage QuotaCount                      `json:"project_count_usage"`
	JSON              quotasThresholdGlobalQuotasJSON `json:"-"`
}

// quotasThresholdGlobalQuotasJSON contains the JSON metadata for the struct
// [QuotasThresholdGlobalQuotas]
type quotasThresholdGlobalQuotasJSON struct {
	InferenceCPUMillicoreCountLimit apijson.Field
	InferenceCPUMillicoreCountUsage apijson.Field
	InferenceGPUA100CountLimit      apijson.Field
	InferenceGPUA100CountUsage      apijson.Field
	InferenceGPUH100CountLimit      apijson.Field
	InferenceGPUH100CountUsage      apijson.Field
	InferenceGPUL40sCountLimit      apijson.Field
	InferenceGPUL40sCountUsage      apijson.Field
	InferenceInstanceCountLimit     apijson.Field
	InferenceInstanceCountUsage     apijson.Field
	KeypairCountLimit               apijson.Field
	KeypairCountUsage               apijson.Field
	ProjectCountLimit               apijson.Field
	ProjectCountUsage               apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *QuotasThresholdGlobalQuotas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotasThresholdGlobalQuotasJSON) RawJSON() string {
	return r.raw
}

type QuotasThresholdRegionalQuota struct {
	// Region id
	RegionID int64 `json:"region_id,required"`
	// Region name
	RegionName string `json:"region_name,required"`
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit QuotaCount `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage QuotaCount `json:"baremetal_basic_count_usage"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit QuotaCount `json:"baremetal_gpu_count_limit"`
	// AI GPU bare metal servers count usage
	BaremetalGPUCountUsage QuotaCount `json:"baremetal_gpu_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit QuotaCount `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage QuotaCount `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit QuotaCount `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage QuotaCount `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit QuotaCount `json:"baremetal_network_count_limit"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage QuotaCount `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit QuotaCount `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage QuotaCount `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit QuotaCount `json:"caas_container_count_limit"`
	// Containers count usage
	CaasContainerCountUsage QuotaCount `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit QuotaCount `json:"caas_cpu_count_limit"`
	// mCPU count for containers usage
	CaasCPUCountUsage QuotaCount `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit QuotaCount `json:"caas_gpu_count_limit"`
	// Containers gpu count usage
	CaasGPUCountUsage QuotaCount `json:"caas_gpu_count_usage"`
	// MB memory count for containers limit
	CaasRamSizeLimit QuotaCount `json:"caas_ram_size_limit"`
	// MB memory count for containers usage
	CaasRamSizeUsage QuotaCount `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit QuotaCount `json:"cluster_count_limit"`
	// K8s clusters count usage
	ClusterCountUsage QuotaCount `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit QuotaCount `json:"cpu_count_limit"`
	// vCPU Count usage
	CPUCountUsage QuotaCount `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit QuotaCount `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage QuotaCount `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit QuotaCount `json:"external_ip_count_limit"`
	// External IP Count usage
	ExternalIPCountUsage QuotaCount `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit QuotaCount `json:"faas_cpu_count_limit"`
	// mCPU count for functions usage
	FaasCPUCountUsage QuotaCount `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit QuotaCount `json:"faas_function_count_limit"`
	// Functions count usage
	FaasFunctionCountUsage QuotaCount `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit QuotaCount `json:"faas_namespace_count_limit"`
	// Functions namespace count usage
	FaasNamespaceCountUsage QuotaCount `json:"faas_namespace_count_usage"`
	// MB memory count for functions limit
	FaasRamSizeLimit QuotaCount `json:"faas_ram_size_limit"`
	// MB memory count for functions usage
	FaasRamSizeUsage QuotaCount `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit QuotaCount `json:"firewall_count_limit"`
	// Firewalls Count usage
	FirewallCountUsage QuotaCount `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit QuotaCount `json:"floating_count_limit"`
	// Floating IP Count usage
	FloatingCountUsage QuotaCount `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit QuotaCount `json:"gpu_count_limit"`
	// GPU Count usage
	GPUCountUsage QuotaCount `json:"gpu_count_usage"`
	// Images Count limit
	ImageCountLimit QuotaCount `json:"image_count_limit"`
	// Images Count usage
	ImageCountUsage QuotaCount `json:"image_count_usage"`
	// Images Size, GiB limit
	ImageSizeLimit QuotaCount `json:"image_size_limit"`
	// Images Size, GiB usage
	ImageSizeUsage QuotaCount `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit QuotaCount `json:"ipu_count_limit"`
	// IPU Count usage
	IpuCountUsage QuotaCount `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit QuotaCount `json:"laas_topic_count_limit"`
	// LaaS Topics Count usage
	LaasTopicCountUsage QuotaCount `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit QuotaCount `json:"loadbalancer_count_limit"`
	// Load Balancers Count usage
	LoadbalancerCountUsage QuotaCount `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit QuotaCount `json:"network_count_limit"`
	// Networks Count usage
	NetworkCountUsage QuotaCount `json:"network_count_usage"`
	// RAM Size, GiB limit
	RamLimit QuotaCount `json:"ram_limit"`
	// RAM Size, GiB usage
	RamUsage QuotaCount `json:"ram_usage"`
	// Registries count limit
	RegistryCountLimit QuotaCount `json:"registry_count_limit"`
	// Registries count usage
	RegistryCountUsage QuotaCount `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit QuotaCount `json:"registry_storage_limit"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage QuotaCount `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit QuotaCount `json:"router_count_limit"`
	// Routers Count usage
	RouterCountUsage QuotaCount `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit QuotaCount `json:"secret_count_limit"`
	// Secret Count usage
	SecretCountUsage QuotaCount `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit QuotaCount `json:"servergroup_count_limit"`
	// Placement Group Count usage
	ServergroupCountUsage QuotaCount `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit QuotaCount `json:"sfs_count_limit"`
	// Shared file system Count usage
	SfsCountUsage QuotaCount `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit QuotaCount `json:"sfs_size_limit"`
	// Shared file system Size, GiB usage
	SfsSizeUsage QuotaCount `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit QuotaCount `json:"shared_vm_count_limit"`
	// Basic VMs Count usage
	SharedVmCountUsage QuotaCount `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit QuotaCount `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage QuotaCount `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit QuotaCount `json:"subnet_count_limit"`
	// Subnets Count usage
	SubnetCountUsage QuotaCount `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit QuotaCount `json:"vm_count_limit"`
	// Instances Dedicated Count usage
	VmCountUsage QuotaCount `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit QuotaCount `json:"volume_count_limit"`
	// Volumes Count usage
	VolumeCountUsage QuotaCount `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit QuotaCount `json:"volume_size_limit"`
	// Volumes Size, GiB usage
	VolumeSizeUsage QuotaCount `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit QuotaCount `json:"volume_snapshots_count_limit"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage QuotaCount `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit QuotaCount `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage QuotaCount                       `json:"volume_snapshots_size_usage"`
	JSON                     quotasThresholdRegionalQuotaJSON `json:"-"`
}

// quotasThresholdRegionalQuotaJSON contains the JSON metadata for the struct
// [QuotasThresholdRegionalQuota]
type quotasThresholdRegionalQuotaJSON struct {
	RegionID                          apijson.Field
	RegionName                        apijson.Field
	BaremetalBasicCountLimit          apijson.Field
	BaremetalBasicCountUsage          apijson.Field
	BaremetalGPUCountLimit            apijson.Field
	BaremetalGPUCountUsage            apijson.Field
	BaremetalHfCountLimit             apijson.Field
	BaremetalHfCountUsage             apijson.Field
	BaremetalInfrastructureCountLimit apijson.Field
	BaremetalInfrastructureCountUsage apijson.Field
	BaremetalNetworkCountLimit        apijson.Field
	BaremetalNetworkCountUsage        apijson.Field
	BaremetalStorageCountLimit        apijson.Field
	BaremetalStorageCountUsage        apijson.Field
	CaasContainerCountLimit           apijson.Field
	CaasContainerCountUsage           apijson.Field
	CaasCPUCountLimit                 apijson.Field
	CaasCPUCountUsage                 apijson.Field
	CaasGPUCountLimit                 apijson.Field
	CaasGPUCountUsage                 apijson.Field
	CaasRamSizeLimit                  apijson.Field
	CaasRamSizeUsage                  apijson.Field
	ClusterCountLimit                 apijson.Field
	ClusterCountUsage                 apijson.Field
	CPUCountLimit                     apijson.Field
	CPUCountUsage                     apijson.Field
	DbaasPostgresClusterCountLimit    apijson.Field
	DbaasPostgresClusterCountUsage    apijson.Field
	ExternalIPCountLimit              apijson.Field
	ExternalIPCountUsage              apijson.Field
	FaasCPUCountLimit                 apijson.Field
	FaasCPUCountUsage                 apijson.Field
	FaasFunctionCountLimit            apijson.Field
	FaasFunctionCountUsage            apijson.Field
	FaasNamespaceCountLimit           apijson.Field
	FaasNamespaceCountUsage           apijson.Field
	FaasRamSizeLimit                  apijson.Field
	FaasRamSizeUsage                  apijson.Field
	FirewallCountLimit                apijson.Field
	FirewallCountUsage                apijson.Field
	FloatingCountLimit                apijson.Field
	FloatingCountUsage                apijson.Field
	GPUCountLimit                     apijson.Field
	GPUCountUsage                     apijson.Field
	ImageCountLimit                   apijson.Field
	ImageCountUsage                   apijson.Field
	ImageSizeLimit                    apijson.Field
	ImageSizeUsage                    apijson.Field
	IpuCountLimit                     apijson.Field
	IpuCountUsage                     apijson.Field
	LaasTopicCountLimit               apijson.Field
	LaasTopicCountUsage               apijson.Field
	LoadbalancerCountLimit            apijson.Field
	LoadbalancerCountUsage            apijson.Field
	NetworkCountLimit                 apijson.Field
	NetworkCountUsage                 apijson.Field
	RamLimit                          apijson.Field
	RamUsage                          apijson.Field
	RegistryCountLimit                apijson.Field
	RegistryCountUsage                apijson.Field
	RegistryStorageLimit              apijson.Field
	RegistryStorageUsage              apijson.Field
	RouterCountLimit                  apijson.Field
	RouterCountUsage                  apijson.Field
	SecretCountLimit                  apijson.Field
	SecretCountUsage                  apijson.Field
	ServergroupCountLimit             apijson.Field
	ServergroupCountUsage             apijson.Field
	SfsCountLimit                     apijson.Field
	SfsCountUsage                     apijson.Field
	SfsSizeLimit                      apijson.Field
	SfsSizeUsage                      apijson.Field
	SharedVmCountLimit                apijson.Field
	SharedVmCountUsage                apijson.Field
	SnapshotScheduleCountLimit        apijson.Field
	SnapshotScheduleCountUsage        apijson.Field
	SubnetCountLimit                  apijson.Field
	SubnetCountUsage                  apijson.Field
	VmCountLimit                      apijson.Field
	VmCountUsage                      apijson.Field
	VolumeCountLimit                  apijson.Field
	VolumeCountUsage                  apijson.Field
	VolumeSizeLimit                   apijson.Field
	VolumeSizeUsage                   apijson.Field
	VolumeSnapshotsCountLimit         apijson.Field
	VolumeSnapshotsCountUsage         apijson.Field
	VolumeSnapshotsSizeLimit          apijson.Field
	VolumeSnapshotsSizeUsage          apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *QuotasThresholdRegionalQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotasThresholdRegionalQuotaJSON) RawJSON() string {
	return r.raw
}

type QuotasThresholdParam struct {
	// Global quota that exceed the threshold
	GlobalQuotas param.Field[QuotasThresholdGlobalQuotasParam] `json:"global_quotas,required"`
	// Regional quota that exceed the threshold
	RegionalQuotas param.Field[[]QuotasThresholdRegionalQuotaParam] `json:"regional_quotas,required"`
}

func (r QuotasThresholdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Global quota that exceed the threshold
type QuotasThresholdGlobalQuotasParam struct {
	// Inference CPU millicore count limit
	InferenceCPUMillicoreCountLimit param.Field[QuotaCountParam] `json:"inference_cpu_millicore_count_limit"`
	// Inference CPU millicore count usage
	InferenceCPUMillicoreCountUsage param.Field[QuotaCountParam] `json:"inference_cpu_millicore_count_usage"`
	// Inference GPU A100 Count limit
	InferenceGPUA100CountLimit param.Field[QuotaCountParam] `json:"inference_gpu_a100_count_limit"`
	// Inference GPU A100 Count usage
	InferenceGPUA100CountUsage param.Field[QuotaCountParam] `json:"inference_gpu_a100_count_usage"`
	// Inference GPU H100 Count limit
	InferenceGPUH100CountLimit param.Field[QuotaCountParam] `json:"inference_gpu_h100_count_limit"`
	// Inference GPU H100 Count usage
	InferenceGPUH100CountUsage param.Field[QuotaCountParam] `json:"inference_gpu_h100_count_usage"`
	// Inference GPU L40s Count limit
	InferenceGPUL40sCountLimit param.Field[QuotaCountParam] `json:"inference_gpu_l40s_count_limit"`
	// Inference GPU L40s Count usage
	InferenceGPUL40sCountUsage param.Field[QuotaCountParam] `json:"inference_gpu_l40s_count_usage"`
	// Inference instance count limit
	InferenceInstanceCountLimit param.Field[QuotaCountParam] `json:"inference_instance_count_limit"`
	// Inference instance count usage
	InferenceInstanceCountUsage param.Field[QuotaCountParam] `json:"inference_instance_count_usage"`
	// SSH Keys Count limit
	KeypairCountLimit param.Field[QuotaCountParam] `json:"keypair_count_limit"`
	// SSH Keys Count usage
	KeypairCountUsage param.Field[QuotaCountParam] `json:"keypair_count_usage"`
	// Projects Count limit
	ProjectCountLimit param.Field[QuotaCountParam] `json:"project_count_limit"`
	// Projects Count usage
	ProjectCountUsage param.Field[QuotaCountParam] `json:"project_count_usage"`
}

func (r QuotasThresholdGlobalQuotasParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QuotasThresholdRegionalQuotaParam struct {
	// Region id
	RegionID param.Field[int64] `json:"region_id,required"`
	// Region name
	RegionName param.Field[string] `json:"region_name,required"`
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit param.Field[QuotaCountParam] `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage param.Field[QuotaCountParam] `json:"baremetal_basic_count_usage"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit param.Field[QuotaCountParam] `json:"baremetal_gpu_count_limit"`
	// AI GPU bare metal servers count usage
	BaremetalGPUCountUsage param.Field[QuotaCountParam] `json:"baremetal_gpu_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit param.Field[QuotaCountParam] `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage param.Field[QuotaCountParam] `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit param.Field[QuotaCountParam] `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage param.Field[QuotaCountParam] `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit param.Field[QuotaCountParam] `json:"baremetal_network_count_limit"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage param.Field[QuotaCountParam] `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit param.Field[QuotaCountParam] `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage param.Field[QuotaCountParam] `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit param.Field[QuotaCountParam] `json:"caas_container_count_limit"`
	// Containers count usage
	CaasContainerCountUsage param.Field[QuotaCountParam] `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit param.Field[QuotaCountParam] `json:"caas_cpu_count_limit"`
	// mCPU count for containers usage
	CaasCPUCountUsage param.Field[QuotaCountParam] `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit param.Field[QuotaCountParam] `json:"caas_gpu_count_limit"`
	// Containers gpu count usage
	CaasGPUCountUsage param.Field[QuotaCountParam] `json:"caas_gpu_count_usage"`
	// MB memory count for containers limit
	CaasRamSizeLimit param.Field[QuotaCountParam] `json:"caas_ram_size_limit"`
	// MB memory count for containers usage
	CaasRamSizeUsage param.Field[QuotaCountParam] `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit param.Field[QuotaCountParam] `json:"cluster_count_limit"`
	// K8s clusters count usage
	ClusterCountUsage param.Field[QuotaCountParam] `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit param.Field[QuotaCountParam] `json:"cpu_count_limit"`
	// vCPU Count usage
	CPUCountUsage param.Field[QuotaCountParam] `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit param.Field[QuotaCountParam] `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage param.Field[QuotaCountParam] `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit param.Field[QuotaCountParam] `json:"external_ip_count_limit"`
	// External IP Count usage
	ExternalIPCountUsage param.Field[QuotaCountParam] `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit param.Field[QuotaCountParam] `json:"faas_cpu_count_limit"`
	// mCPU count for functions usage
	FaasCPUCountUsage param.Field[QuotaCountParam] `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit param.Field[QuotaCountParam] `json:"faas_function_count_limit"`
	// Functions count usage
	FaasFunctionCountUsage param.Field[QuotaCountParam] `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit param.Field[QuotaCountParam] `json:"faas_namespace_count_limit"`
	// Functions namespace count usage
	FaasNamespaceCountUsage param.Field[QuotaCountParam] `json:"faas_namespace_count_usage"`
	// MB memory count for functions limit
	FaasRamSizeLimit param.Field[QuotaCountParam] `json:"faas_ram_size_limit"`
	// MB memory count for functions usage
	FaasRamSizeUsage param.Field[QuotaCountParam] `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit param.Field[QuotaCountParam] `json:"firewall_count_limit"`
	// Firewalls Count usage
	FirewallCountUsage param.Field[QuotaCountParam] `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit param.Field[QuotaCountParam] `json:"floating_count_limit"`
	// Floating IP Count usage
	FloatingCountUsage param.Field[QuotaCountParam] `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit param.Field[QuotaCountParam] `json:"gpu_count_limit"`
	// GPU Count usage
	GPUCountUsage param.Field[QuotaCountParam] `json:"gpu_count_usage"`
	// Images Count limit
	ImageCountLimit param.Field[QuotaCountParam] `json:"image_count_limit"`
	// Images Count usage
	ImageCountUsage param.Field[QuotaCountParam] `json:"image_count_usage"`
	// Images Size, GiB limit
	ImageSizeLimit param.Field[QuotaCountParam] `json:"image_size_limit"`
	// Images Size, GiB usage
	ImageSizeUsage param.Field[QuotaCountParam] `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit param.Field[QuotaCountParam] `json:"ipu_count_limit"`
	// IPU Count usage
	IpuCountUsage param.Field[QuotaCountParam] `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit param.Field[QuotaCountParam] `json:"laas_topic_count_limit"`
	// LaaS Topics Count usage
	LaasTopicCountUsage param.Field[QuotaCountParam] `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit param.Field[QuotaCountParam] `json:"loadbalancer_count_limit"`
	// Load Balancers Count usage
	LoadbalancerCountUsage param.Field[QuotaCountParam] `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit param.Field[QuotaCountParam] `json:"network_count_limit"`
	// Networks Count usage
	NetworkCountUsage param.Field[QuotaCountParam] `json:"network_count_usage"`
	// RAM Size, GiB limit
	RamLimit param.Field[QuotaCountParam] `json:"ram_limit"`
	// RAM Size, GiB usage
	RamUsage param.Field[QuotaCountParam] `json:"ram_usage"`
	// Registries count limit
	RegistryCountLimit param.Field[QuotaCountParam] `json:"registry_count_limit"`
	// Registries count usage
	RegistryCountUsage param.Field[QuotaCountParam] `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit param.Field[QuotaCountParam] `json:"registry_storage_limit"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage param.Field[QuotaCountParam] `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit param.Field[QuotaCountParam] `json:"router_count_limit"`
	// Routers Count usage
	RouterCountUsage param.Field[QuotaCountParam] `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit param.Field[QuotaCountParam] `json:"secret_count_limit"`
	// Secret Count usage
	SecretCountUsage param.Field[QuotaCountParam] `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit param.Field[QuotaCountParam] `json:"servergroup_count_limit"`
	// Placement Group Count usage
	ServergroupCountUsage param.Field[QuotaCountParam] `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit param.Field[QuotaCountParam] `json:"sfs_count_limit"`
	// Shared file system Count usage
	SfsCountUsage param.Field[QuotaCountParam] `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit param.Field[QuotaCountParam] `json:"sfs_size_limit"`
	// Shared file system Size, GiB usage
	SfsSizeUsage param.Field[QuotaCountParam] `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit param.Field[QuotaCountParam] `json:"shared_vm_count_limit"`
	// Basic VMs Count usage
	SharedVmCountUsage param.Field[QuotaCountParam] `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit param.Field[QuotaCountParam] `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage param.Field[QuotaCountParam] `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit param.Field[QuotaCountParam] `json:"subnet_count_limit"`
	// Subnets Count usage
	SubnetCountUsage param.Field[QuotaCountParam] `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit param.Field[QuotaCountParam] `json:"vm_count_limit"`
	// Instances Dedicated Count usage
	VmCountUsage param.Field[QuotaCountParam] `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit param.Field[QuotaCountParam] `json:"volume_count_limit"`
	// Volumes Count usage
	VolumeCountUsage param.Field[QuotaCountParam] `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit param.Field[QuotaCountParam] `json:"volume_size_limit"`
	// Volumes Size, GiB usage
	VolumeSizeUsage param.Field[QuotaCountParam] `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit param.Field[QuotaCountParam] `json:"volume_snapshots_count_limit"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage param.Field[QuotaCountParam] `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit param.Field[QuotaCountParam] `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage param.Field[QuotaCountParam] `json:"volume_snapshots_size_usage"`
}

func (r QuotasThresholdRegionalQuotaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2ClientQuotaNotificationThresholdUpdateOrNewParams struct {
	// Quota notification threshold in percentage
	Threshold param.Field[int64] `json:"threshold,required"`
	// A message data
	LastMessage param.Field[QuotasThresholdParam] `json:"last_message"`
	// Time of last successful email sending
	LastSending param.Field[time.Time] `json:"last_sending" format:"date-time"`
}

func (r CloudV2ClientQuotaNotificationThresholdUpdateOrNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
