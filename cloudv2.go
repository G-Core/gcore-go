// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
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

// CloudV2Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2Service] method instead.
type CloudV2Service struct {
	Options       []option.RequestOption
	AI            *CloudV2AIService
	Inference     *CloudV2InferenceService
	Pricing       *CloudV2PricingService
	Instances     *CloudV2InstanceService
	Ports         *CloudV2PortService
	Loadbalancers *CloudV2LoadbalancerService
	K8s           *CloudV2K8Service
	ClientQuotas  *CloudV2ClientQuotaService
	LimitsRequest *CloudV2LimitsRequestService
	Keypairs      *CloudV2KeypairService
	Volumes       *CloudV2VolumeService
}

// NewCloudV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudV2Service(opts ...option.RequestOption) (r *CloudV2Service) {
	r = &CloudV2Service{}
	r.Options = opts
	r.AI = NewCloudV2AIService(opts...)
	r.Inference = NewCloudV2InferenceService(opts...)
	r.Pricing = NewCloudV2PricingService(opts...)
	r.Instances = NewCloudV2InstanceService(opts...)
	r.Ports = NewCloudV2PortService(opts...)
	r.Loadbalancers = NewCloudV2LoadbalancerService(opts...)
	r.K8s = NewCloudV2K8Service(opts...)
	r.ClientQuotas = NewCloudV2ClientQuotaService(opts...)
	r.LimitsRequest = NewCloudV2LimitsRequestService(opts...)
	r.Keypairs = NewCloudV2KeypairService(opts...)
	r.Volumes = NewCloudV2VolumeService(opts...)
	return
}

// Create secret v2
func (r *CloudV2Service) NewSecret(ctx context.Context, projectID int64, regionID int64, body CloudV2NewSecretParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/secrets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get amount of available bare metal nodes without client's reservations It's
// similar v2/bmcapacity just without <project_id> in path
func (r *CloudV2Service) GetBmcapacity(ctx context.Context, regionID int64, query CloudV2GetBmcapacityParams, opts ...option.RequestOption) (res *BaremetalCapacity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/bmcapacity/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get global quota
func (r *CloudV2Service) GetGlobalQuota(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *GlobalQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/global_quotas/%v", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a quota by region
func (r *CloudV2Service) GetRegionalQuota(ctx context.Context, clientID int64, regionID int64, opts ...option.RequestOption) (res *RegionalQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/regional_quotas/%v/%v", clientID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update listener
func (r *CloudV2Service) UpdateListener(ctx context.Context, projectID int64, regionID int64, listenerID string, body CloudV2UpdateListenerParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/lblisteners/%v/%v/%s", projectID, regionID, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
	ProjectCountUsage int64            `json:"project_count_usage"`
	JSON              globalQuotasJSON `json:"-"`
}

// globalQuotasJSON contains the JSON metadata for the struct [GlobalQuotas]
type globalQuotasJSON struct {
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

func (r *GlobalQuotas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalQuotasJSON) RawJSON() string {
	return r.raw
}

type RegionalQuotas struct {
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
	VolumeSnapshotsSizeUsage int64              `json:"volume_snapshots_size_usage"`
	JSON                     regionalQuotasJSON `json:"-"`
}

// regionalQuotasJSON contains the JSON metadata for the struct [RegionalQuotas]
type regionalQuotasJSON struct {
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
	RegionID                          apijson.Field
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

func (r *RegionalQuotas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalQuotasJSON) RawJSON() string {
	return r.raw
}

type CloudV2NewSecretParams struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret payload.
	Payload param.Field[CloudV2NewSecretParamsPayload] `json:"payload,required"`
	// Datetime when the secret will expire. Defaults to None
	Expiration param.Field[time.Time] `json:"expiration" format:"date-time"`
}

func (r CloudV2NewSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Secret payload.
type CloudV2NewSecretParamsPayload struct {
	// SSL certificate in PEM format.
	Certificate param.Field[string] `json:"certificate,required" format:"password"`
	// SSL certificate chain of intermediates and root certificates in PEM format.
	CertificateChain param.Field[string] `json:"certificate_chain,required" format:"password"`
	// SSL private key in PEM format.
	PrivateKey param.Field[string] `json:"private_key,required" format:"password"`
}

func (r CloudV2NewSecretParamsPayload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2GetBmcapacityParams struct {
	// Available for roots
	ClientID param.Field[int64] `query:"client_id"`
}

// URLQuery serializes [CloudV2GetBmcapacityParams]'s query parameters as
// `url.Values`.
func (r CloudV2GetBmcapacityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2UpdateListenerParams struct {
	// Network CIDRs from which service will be accessible
	AllowedCidrs param.Field[[]string] `json:"allowed_cidrs" format:"ipvanynetwork"`
	// Limit of simultaneous connections
	ConnectionLimit param.Field[int64] `json:"connection_limit"`
	// Load balancer listener name
	Name param.Field[string] `json:"name"`
	// ID of the secret where PKCS12 file is stored for TERMINATED_HTTPS or PROMETHEUS
	// load balancer
	SecretID param.Field[string] `json:"secret_id" format:"uuid4"`
	// List of secret's ID containing PKCS12 format certificate/key bundfles for
	// TERMINATED_HTTPS or PROMETHEUS listeners
	SniSecretID param.Field[[]string] `json:"sni_secret_id" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Field[int64] `json:"timeout_client_data"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Field[int64] `json:"timeout_member_connect"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Field[int64] `json:"timeout_member_data"`
	// Load balancer listener users list
	UserList param.Field[[]UserListItemParam] `json:"user_list"`
}

func (r CloudV2UpdateListenerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
