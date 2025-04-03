// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1BminstanceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1BminstanceService] method instead.
type CloudV1BminstanceService struct {
	Options []option.RequestOption
}

// NewCloudV1BminstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1BminstanceService(opts ...option.RequestOption) (r *CloudV1BminstanceService) {
	r = &CloudV1BminstanceService{}
	r.Options = opts
	return
}

// Create a new bare metal server or multiple servers
func (r *CloudV1BminstanceService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1BminstanceNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The date and time stamp format in changes-since or changes-before should be ISO
// 8601: CCYY-MM-DDThh:mm:ss±hh:mm For example, 2015-08-27T09:49:58-05:00. Values
// must be urlencoded. If the time zone is omitted, the UTC time zone is assumed.
// When both changes-since and changes-before are specified, the value of the
// changes-since must be earlier than or equal to the value of the changes-before.
func (r *CloudV1BminstanceService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1BminstanceListParams, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified bare metal server will be calculated
func (r *CloudV1BminstanceService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1BminstanceCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get flavors available for a potential bare metal server
func (r *CloudV1BminstanceService) GetAvailableFlavors(ctx context.Context, projectID int64, regionID int64, params CloudV1BminstanceGetAvailableFlavorsParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/available_flavors", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rebuild bare metal server
func (r *CloudV1BminstanceService) Rebuild(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1BminstanceRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s/rebuild", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check quota before bare metal servers will be created.
type CheckQuotaBeforeBmCreationParam struct {
	// Bare metal flavor name.
	Flavor param.Field[string] `json:"flavor"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]CheckQuotaBeforeBmCreationInterfacesUnionParam] `json:"interfaces"`
}

func (r CheckQuotaBeforeBmCreationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to default external network
type CheckQuotaBeforeBmCreationInterfaceParam struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CheckQuotaBeforeBmCreationInterfacesIPFamily] `json:"ip_family"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID param.Field[string] `json:"port_id"`
	// Port is assinged an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r CheckQuotaBeforeBmCreationInterfaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CheckQuotaBeforeBmCreationInterfaceParam) implementsCheckQuotaBeforeBmCreationInterfacesUnionParam() {
}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalParam], [NetworkSubnetFipParam],
// [NetworkAnySubnetFipParam], [ReservedFixedIPFipParam],
// [CheckQuotaBeforeBmCreationInterfaceParam].
type CheckQuotaBeforeBmCreationInterfacesUnionParam interface {
	implementsCheckQuotaBeforeBmCreationInterfacesUnionParam()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CheckQuotaBeforeBmCreationInterfacesIPFamily string

const (
	CheckQuotaBeforeBmCreationInterfacesIPFamilyDual CheckQuotaBeforeBmCreationInterfacesIPFamily = "dual"
	CheckQuotaBeforeBmCreationInterfacesIPFamilyIpv4 CheckQuotaBeforeBmCreationInterfacesIPFamily = "ipv4"
	CheckQuotaBeforeBmCreationInterfacesIPFamilyIpv6 CheckQuotaBeforeBmCreationInterfacesIPFamily = "ipv6"
)

func (r CheckQuotaBeforeBmCreationInterfacesIPFamily) IsKnown() bool {
	switch r {
	case CheckQuotaBeforeBmCreationInterfacesIPFamilyDual, CheckQuotaBeforeBmCreationInterfacesIPFamilyIpv4, CheckQuotaBeforeBmCreationInterfacesIPFamilyIpv6:
		return true
	}
	return false
}

type DeprecatedCreateDdosProfileParam struct {
	// DDoS profile template ID.
	ProfileTemplate param.Field[int64] `json:"profile_template,required"`
	// Protection parameters.
	Fields param.Field[[]DeprecatedCreateDdosProfileFieldParam] `json:"fields"`
	// DDoS profile template name.
	ProfileTemplateName param.Field[string] `json:"profile_template_name"`
}

func (r DeprecatedCreateDdosProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeprecatedCreateDdosProfileFieldParam struct {
	// ID of DDoS profile field
	BaseField param.Field[int64] `json:"base_field"`
	// Name of DDoS profile field
	FieldName param.Field[string] `json:"field_name"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue param.Field[interface{}] `json:"field_value"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Field[string] `json:"value"`
}

func (r DeprecatedCreateDdosProfileFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeprecatedInstanceList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []DeprecatedInstance       `json:"results"`
	JSON    deprecatedInstanceListJSON `json:"-"`
}

// deprecatedInstanceListJSON contains the JSON metadata for the struct
// [DeprecatedInstanceList]
type deprecatedInstanceListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedInstanceList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceListJSON) RawJSON() string {
	return r.raw
}

// Instance will be attached to a subnet with the largest count of free ips.
// Floating IP will be created and attached to that IP
type NetworkAnySubnetFipParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id,required"`
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[NetworkAnySubnetFipIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// 'any_subnet'. Instance will be attached to a subnet with the largest count of
	// free ips. Floating IP will be created and attached to that IP
	Type param.Field[string] `json:"type"`
}

func (r NetworkAnySubnetFipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NetworkAnySubnetFipParam) implementsCheckQuotaBeforeBmCreationInterfacesUnionParam() {}

func (r NetworkAnySubnetFipParam) implementsCloudV1BminstanceNewParamsInterfaceUnion() {}

func (r NetworkAnySubnetFipParam) implementsCloudV1AIClusterGPUNewParamsInterfaceUnion() {}

func (r NetworkAnySubnetFipParam) implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion() {}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type NetworkAnySubnetFipIPFamily string

const (
	NetworkAnySubnetFipIPFamilyDual NetworkAnySubnetFipIPFamily = "dual"
	NetworkAnySubnetFipIPFamilyIpv4 NetworkAnySubnetFipIPFamily = "ipv4"
	NetworkAnySubnetFipIPFamilyIpv6 NetworkAnySubnetFipIPFamily = "ipv6"
)

func (r NetworkAnySubnetFipIPFamily) IsKnown() bool {
	switch r {
	case NetworkAnySubnetFipIPFamilyDual, NetworkAnySubnetFipIPFamilyIpv4, NetworkAnySubnetFipIPFamilyIpv6:
		return true
	}
	return false
}

// Instance will be attached to specified subnet, and floating IP can be created
// and attached to this IP
type NetworkSubnetFipParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id,required"`
	// Port is assinged an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id,required"`
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// 'subnet'. Instance will be attached to specified subnet, and floating IP can be
	// created and attached to this IP
	Type param.Field[string] `json:"type"`
}

func (r NetworkSubnetFipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NetworkSubnetFipParam) implementsCheckQuotaBeforeBmCreationInterfacesUnionParam() {}

func (r NetworkSubnetFipParam) implementsCloudV1BminstanceNewParamsInterfaceUnion() {}

func (r NetworkSubnetFipParam) implementsCloudV1AIClusterGPUNewParamsInterfaceUnion() {}

func (r NetworkSubnetFipParam) implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion() {}

// Instance will be attached to default external network
type NewInterfaceExternalParam struct {
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[NewInterfaceExternalIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r NewInterfaceExternalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewInterfaceExternalParam) implementsCheckQuotaBeforeBmCreationInterfacesUnionParam() {}

func (r NewInterfaceExternalParam) implementsCloudV1BminstanceNewParamsInterfaceUnion() {}

func (r NewInterfaceExternalParam) implementsCloudV1AIClusterGPUNewParamsInterfaceUnion() {}

func (r NewInterfaceExternalParam) implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion() {}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type NewInterfaceExternalIPFamily string

const (
	NewInterfaceExternalIPFamilyDual NewInterfaceExternalIPFamily = "dual"
	NewInterfaceExternalIPFamilyIpv4 NewInterfaceExternalIPFamily = "ipv4"
	NewInterfaceExternalIPFamilyIpv6 NewInterfaceExternalIPFamily = "ipv6"
)

func (r NewInterfaceExternalIPFamily) IsKnown() bool {
	switch r {
	case NewInterfaceExternalIPFamilyDual, NewInterfaceExternalIPFamilyIpv4, NewInterfaceExternalIPFamilyIpv6:
		return true
	}
	return false
}

type RegionalDiffQuotas struct {
	// Basic bare metal servers count limit
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// Basic bare metal servers count requested
	BaremetalBasicCountRequested int64 `json:"baremetal_basic_count_requested"`
	// Basic bare metal servers count usage
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage"`
	// AI GPU bare metal servers count limit
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// AI GPU bare metal servers count requested
	BaremetalGPUCountRequested int64 `json:"baremetal_gpu_count_requested"`
	// AI GPU bare metal servers count usage
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count requested
	BaremetalHfCountRequested int64 `json:"baremetal_hf_count_requested"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// Infrastructure bare metal servers count limit
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// Infrastructure bare metal servers count requested
	BaremetalInfrastructureCountRequested int64 `json:"baremetal_infrastructure_count_requested"`
	// Infrastructure bare metal servers count usage
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage"`
	// Bare metal Network Count limit
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// Bare metal Network Count requested
	BaremetalNetworkCountRequested int64 `json:"baremetal_network_count_requested"`
	// Bare metal Network Count usage
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage"`
	// Storage bare metal servers count limit
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// Storage bare metal servers count requested
	BaremetalStorageCountRequested int64 `json:"baremetal_storage_count_requested"`
	// Storage bare metal servers count usage
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage"`
	// Containers count limit
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// Containers count requested
	CaasContainerCountRequested int64 `json:"caas_container_count_requested"`
	// Containers count usage
	CaasContainerCountUsage int64 `json:"caas_container_count_usage"`
	// mCPU count for containers limit
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// mCPU count for containers requested
	CaasCPUCountRequested int64 `json:"caas_cpu_count_requested"`
	// mCPU count for containers usage
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage"`
	// Containers gpu count limit
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// Containers gpu count requested
	CaasGPUCountRequested int64 `json:"caas_gpu_count_requested"`
	// Containers gpu count usage
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage"`
	// MB memory count for containers limit
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// MB memory count for containers requested
	CaasRamSizeRequested int64 `json:"caas_ram_size_requested"`
	// MB memory count for containers usage
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// K8s clusters count requested
	ClusterCountRequested int64 `json:"cluster_count_requested"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// vCPU Count requested
	CPUCountRequested int64 `json:"cpu_count_requested"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// DBaaS cluster count limit
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// DBaaS cluster count requested
	DbaasPostgresClusterCountRequested int64 `json:"dbaas_postgres_cluster_count_requested"`
	// DBaaS cluster count usage
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage"`
	// External IP Count limit
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// External IP Count requested
	ExternalIPCountRequested int64 `json:"external_ip_count_requested"`
	// External IP Count usage
	ExternalIPCountUsage int64 `json:"external_ip_count_usage"`
	// mCPU count for functions limit
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// mCPU count for functions requested
	FaasCPUCountRequested int64 `json:"faas_cpu_count_requested"`
	// mCPU count for functions usage
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage"`
	// Functions count limit
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// Functions count requested
	FaasFunctionCountRequested int64 `json:"faas_function_count_requested"`
	// Functions count usage
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage"`
	// Functions namespace count limit
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// Functions namespace count requested
	FaasNamespaceCountRequested int64 `json:"faas_namespace_count_requested"`
	// Functions namespace count usage
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage"`
	// MB memory count for functions limit
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// MB memory count for functions requested
	FaasRamSizeRequested int64 `json:"faas_ram_size_requested"`
	// MB memory count for functions usage
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Firewalls Count requested
	FirewallCountRequested int64 `json:"firewall_count_requested"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// Floating IP Count requested
	FloatingCountRequested int64 `json:"floating_count_requested"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// GPU Count requested
	GPUCountRequested int64 `json:"gpu_count_requested"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// Images Count limit
	ImageCountLimit int64 `json:"image_count_limit"`
	// Images Count requested
	ImageCountRequested int64 `json:"image_count_requested"`
	// Images Count usage
	ImageCountUsage int64 `json:"image_count_usage"`
	// Images Size, GiB limit
	ImageSizeLimit int64 `json:"image_size_limit"`
	// Images Size, GiB requested
	ImageSizeRequested int64 `json:"image_size_requested"`
	// Images Size, GiB usage
	ImageSizeUsage int64 `json:"image_size_usage"`
	// IPU Count limit
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// IPU Count requested
	IpuCountRequested int64 `json:"ipu_count_requested"`
	// IPU Count usage
	IpuCountUsage int64 `json:"ipu_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// LaaS Topics Count requested
	LaasTopicCountRequested int64 `json:"laas_topic_count_requested"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Load Balancers Count requested
	LoadbalancerCountRequested int64 `json:"loadbalancer_count_requested"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// Networks Count limit
	NetworkCountLimit int64 `json:"network_count_limit"`
	// Networks Count requested
	NetworkCountRequested int64 `json:"network_count_requested"`
	// Networks Count usage
	NetworkCountUsage int64 `json:"network_count_usage"`
	// RAM Size, GiB limit
	RamLimit int64 `json:"ram_limit"`
	// RAM Size, GiB requested
	RamRequested int64 `json:"ram_requested"`
	// RAM Size, GiB usage
	RamUsage int64 `json:"ram_usage"`
	// Registries count limit
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// Registries count requested
	RegistryCountRequested int64 `json:"registry_count_requested"`
	// Registries count usage
	RegistryCountUsage int64 `json:"registry_count_usage"`
	// Registries volume usage, GiB limit
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// Registries volume usage, GiB requested
	RegistryStorageRequested int64 `json:"registry_storage_requested"`
	// Registries volume usage, GiB usage
	RegistryStorageUsage int64 `json:"registry_storage_usage"`
	// Routers Count limit
	RouterCountLimit int64 `json:"router_count_limit"`
	// Routers Count requested
	RouterCountRequested int64 `json:"router_count_requested"`
	// Routers Count usage
	RouterCountUsage int64 `json:"router_count_usage"`
	// Secret Count limit
	SecretCountLimit int64 `json:"secret_count_limit"`
	// Secret Count requested
	SecretCountRequested int64 `json:"secret_count_requested"`
	// Secret Count usage
	SecretCountUsage int64 `json:"secret_count_usage"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Placement Group Count requested
	ServergroupCountRequested int64 `json:"servergroup_count_requested"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Count requested
	SfsCountRequested int64 `json:"sfs_count_requested"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Shared file system Size, GiB requested
	SfsSizeRequested int64 `json:"sfs_size_requested"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64 `json:"sfs_size_usage"`
	// Basic VMs Count limit
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// Basic VMs Count requested
	SharedVmCountRequested int64 `json:"shared_vm_count_requested"`
	// Basic VMs Count usage
	SharedVmCountUsage int64 `json:"shared_vm_count_usage"`
	// Snapshot Schedules Count limit
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// Snapshot Schedules Count requested
	SnapshotScheduleCountRequested int64 `json:"snapshot_schedule_count_requested"`
	// Snapshot Schedules Count usage
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage"`
	// Subnets Count limit
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// Subnets Count requested
	SubnetCountRequested int64 `json:"subnet_count_requested"`
	// Subnets Count usage
	SubnetCountUsage int64 `json:"subnet_count_usage"`
	// Instances Dedicated Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// Instances Dedicated Count requested
	VmCountRequested int64 `json:"vm_count_requested"`
	// Instances Dedicated Count usage
	VmCountUsage int64 `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Count requested
	VolumeCountRequested int64 `json:"volume_count_requested"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Volumes Size, GiB requested
	VolumeSizeRequested int64 `json:"volume_size_requested"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// Snapshots Count limit
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// Snapshots Count requested
	VolumeSnapshotsCountRequested int64 `json:"volume_snapshots_count_requested"`
	// Snapshots Count usage
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage"`
	// Snapshots Size, GiB limit
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// Snapshots Size, GiB requested
	VolumeSnapshotsSizeRequested int64 `json:"volume_snapshots_size_requested"`
	// Snapshots Size, GiB usage
	VolumeSnapshotsSizeUsage int64                  `json:"volume_snapshots_size_usage"`
	JSON                     regionalDiffQuotasJSON `json:"-"`
}

// regionalDiffQuotasJSON contains the JSON metadata for the struct
// [RegionalDiffQuotas]
type regionalDiffQuotasJSON struct {
	BaremetalBasicCountLimit              apijson.Field
	BaremetalBasicCountRequested          apijson.Field
	BaremetalBasicCountUsage              apijson.Field
	BaremetalGPUCountLimit                apijson.Field
	BaremetalGPUCountRequested            apijson.Field
	BaremetalGPUCountUsage                apijson.Field
	BaremetalHfCountLimit                 apijson.Field
	BaremetalHfCountRequested             apijson.Field
	BaremetalHfCountUsage                 apijson.Field
	BaremetalInfrastructureCountLimit     apijson.Field
	BaremetalInfrastructureCountRequested apijson.Field
	BaremetalInfrastructureCountUsage     apijson.Field
	BaremetalNetworkCountLimit            apijson.Field
	BaremetalNetworkCountRequested        apijson.Field
	BaremetalNetworkCountUsage            apijson.Field
	BaremetalStorageCountLimit            apijson.Field
	BaremetalStorageCountRequested        apijson.Field
	BaremetalStorageCountUsage            apijson.Field
	CaasContainerCountLimit               apijson.Field
	CaasContainerCountRequested           apijson.Field
	CaasContainerCountUsage               apijson.Field
	CaasCPUCountLimit                     apijson.Field
	CaasCPUCountRequested                 apijson.Field
	CaasCPUCountUsage                     apijson.Field
	CaasGPUCountLimit                     apijson.Field
	CaasGPUCountRequested                 apijson.Field
	CaasGPUCountUsage                     apijson.Field
	CaasRamSizeLimit                      apijson.Field
	CaasRamSizeRequested                  apijson.Field
	CaasRamSizeUsage                      apijson.Field
	ClusterCountLimit                     apijson.Field
	ClusterCountRequested                 apijson.Field
	ClusterCountUsage                     apijson.Field
	CPUCountLimit                         apijson.Field
	CPUCountRequested                     apijson.Field
	CPUCountUsage                         apijson.Field
	DbaasPostgresClusterCountLimit        apijson.Field
	DbaasPostgresClusterCountRequested    apijson.Field
	DbaasPostgresClusterCountUsage        apijson.Field
	ExternalIPCountLimit                  apijson.Field
	ExternalIPCountRequested              apijson.Field
	ExternalIPCountUsage                  apijson.Field
	FaasCPUCountLimit                     apijson.Field
	FaasCPUCountRequested                 apijson.Field
	FaasCPUCountUsage                     apijson.Field
	FaasFunctionCountLimit                apijson.Field
	FaasFunctionCountRequested            apijson.Field
	FaasFunctionCountUsage                apijson.Field
	FaasNamespaceCountLimit               apijson.Field
	FaasNamespaceCountRequested           apijson.Field
	FaasNamespaceCountUsage               apijson.Field
	FaasRamSizeLimit                      apijson.Field
	FaasRamSizeRequested                  apijson.Field
	FaasRamSizeUsage                      apijson.Field
	FirewallCountLimit                    apijson.Field
	FirewallCountRequested                apijson.Field
	FirewallCountUsage                    apijson.Field
	FloatingCountLimit                    apijson.Field
	FloatingCountRequested                apijson.Field
	FloatingCountUsage                    apijson.Field
	GPUCountLimit                         apijson.Field
	GPUCountRequested                     apijson.Field
	GPUCountUsage                         apijson.Field
	ImageCountLimit                       apijson.Field
	ImageCountRequested                   apijson.Field
	ImageCountUsage                       apijson.Field
	ImageSizeLimit                        apijson.Field
	ImageSizeRequested                    apijson.Field
	ImageSizeUsage                        apijson.Field
	IpuCountLimit                         apijson.Field
	IpuCountRequested                     apijson.Field
	IpuCountUsage                         apijson.Field
	LaasTopicCountLimit                   apijson.Field
	LaasTopicCountRequested               apijson.Field
	LaasTopicCountUsage                   apijson.Field
	LoadbalancerCountLimit                apijson.Field
	LoadbalancerCountRequested            apijson.Field
	LoadbalancerCountUsage                apijson.Field
	NetworkCountLimit                     apijson.Field
	NetworkCountRequested                 apijson.Field
	NetworkCountUsage                     apijson.Field
	RamLimit                              apijson.Field
	RamRequested                          apijson.Field
	RamUsage                              apijson.Field
	RegistryCountLimit                    apijson.Field
	RegistryCountRequested                apijson.Field
	RegistryCountUsage                    apijson.Field
	RegistryStorageLimit                  apijson.Field
	RegistryStorageRequested              apijson.Field
	RegistryStorageUsage                  apijson.Field
	RouterCountLimit                      apijson.Field
	RouterCountRequested                  apijson.Field
	RouterCountUsage                      apijson.Field
	SecretCountLimit                      apijson.Field
	SecretCountRequested                  apijson.Field
	SecretCountUsage                      apijson.Field
	ServergroupCountLimit                 apijson.Field
	ServergroupCountRequested             apijson.Field
	ServergroupCountUsage                 apijson.Field
	SfsCountLimit                         apijson.Field
	SfsCountRequested                     apijson.Field
	SfsCountUsage                         apijson.Field
	SfsSizeLimit                          apijson.Field
	SfsSizeRequested                      apijson.Field
	SfsSizeUsage                          apijson.Field
	SharedVmCountLimit                    apijson.Field
	SharedVmCountRequested                apijson.Field
	SharedVmCountUsage                    apijson.Field
	SnapshotScheduleCountLimit            apijson.Field
	SnapshotScheduleCountRequested        apijson.Field
	SnapshotScheduleCountUsage            apijson.Field
	SubnetCountLimit                      apijson.Field
	SubnetCountRequested                  apijson.Field
	SubnetCountUsage                      apijson.Field
	VmCountLimit                          apijson.Field
	VmCountRequested                      apijson.Field
	VmCountUsage                          apijson.Field
	VolumeCountLimit                      apijson.Field
	VolumeCountRequested                  apijson.Field
	VolumeCountUsage                      apijson.Field
	VolumeSizeLimit                       apijson.Field
	VolumeSizeRequested                   apijson.Field
	VolumeSizeUsage                       apijson.Field
	VolumeSnapshotsCountLimit             apijson.Field
	VolumeSnapshotsCountRequested         apijson.Field
	VolumeSnapshotsCountUsage             apijson.Field
	VolumeSnapshotsSizeLimit              apijson.Field
	VolumeSnapshotsSizeRequested          apijson.Field
	VolumeSnapshotsSizeUsage              apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *RegionalDiffQuotas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalDiffQuotasJSON) RawJSON() string {
	return r.raw
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
type ReservedFixedIPFipParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID param.Field[string] `json:"port_id,required"`
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// 'reserved_fixed_ip'. Instance will be attached to the given port. Floating IP
	// will be created and attached to that IP
	Type param.Field[string] `json:"type"`
}

func (r ReservedFixedIPFipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReservedFixedIPFipParam) implementsCheckQuotaBeforeBmCreationInterfacesUnionParam() {}

func (r ReservedFixedIPFipParam) implementsCloudV1BminstanceNewParamsInterfaceUnion() {}

func (r ReservedFixedIPFipParam) implementsCloudV1AIClusterGPUNewParamsInterfaceUnion() {}

func (r ReservedFixedIPFipParam) implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion() {}

// Task ID list object
type TaskIDList struct {
	// Task list
	Tasks []string       `json:"tasks"`
	JSON  taskIDListJSON `json:"-"`
}

// taskIDListJSON contains the JSON metadata for the struct [TaskIDList]
type taskIDListJSON struct {
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskIDList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskIDListJSON) RawJSON() string {
	return r.raw
}

type CloudV1BminstanceNewParams struct {
	// Flavor ID
	Flavor param.Field[string] `json:"flavor,required"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]CloudV1BminstanceNewParamsInterfaceUnion] `json:"interfaces,required"`
	// Parameters for the application template from the marketplace
	AppConfig param.Field[interface{}] `json:"app_config"`
	// Apptemplate ID. Either image_id or apptemplate_id is required.
	ApptemplateID param.Field[string] `json:"apptemplate_id"`
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Image ID. Either image_id or apptemplate_id is required.
	ImageID param.Field[string] `json:"image_id"`
	// Keypair name to inject into new instance(s)
	KeypairName param.Field[string] `json:"keypair_name"`
	// Create one or more metadata items for an instance
	Metadata param.Field[interface{}] `json:"metadata"`
	// List of instance names which will be changed by template. You can use forms
	// 'ip_octets', 'two_ip_octets', 'one_ip_octet'.
	NameTemplates param.Field[[]string] `json:"name_templates"`
	// List of instance names
	Names param.Field[[]string] `json:"names"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password param.Field[string] `json:"password"`
	// String in base64 format. For Linux instances, 'user_data' is ignored when
	// 'password' field is provided. For Windows instances, Admin user password is set
	// by 'password' field and cannot be updated via 'user_data'. Examples of the
	// user_data: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Field[string] `json:"user_data"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username param.Field[string] `json:"username"`
}

func (r CloudV1BminstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to default external network
type CloudV1BminstanceNewParamsInterface struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV1BminstanceNewParamsInterfacesIPFamily] `json:"ip_family"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID param.Field[string] `json:"port_id"`
	// Port is assinged an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r CloudV1BminstanceNewParamsInterface) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1BminstanceNewParamsInterface) implementsCloudV1BminstanceNewParamsInterfaceUnion() {}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalParam], [NetworkSubnetFipParam],
// [NetworkAnySubnetFipParam], [ReservedFixedIPFipParam],
// [CloudV1BminstanceNewParamsInterface].
type CloudV1BminstanceNewParamsInterfaceUnion interface {
	implementsCloudV1BminstanceNewParamsInterfaceUnion()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV1BminstanceNewParamsInterfacesIPFamily string

const (
	CloudV1BminstanceNewParamsInterfacesIPFamilyDual CloudV1BminstanceNewParamsInterfacesIPFamily = "dual"
	CloudV1BminstanceNewParamsInterfacesIPFamilyIpv4 CloudV1BminstanceNewParamsInterfacesIPFamily = "ipv4"
	CloudV1BminstanceNewParamsInterfacesIPFamilyIpv6 CloudV1BminstanceNewParamsInterfacesIPFamily = "ipv6"
)

func (r CloudV1BminstanceNewParamsInterfacesIPFamily) IsKnown() bool {
	switch r {
	case CloudV1BminstanceNewParamsInterfacesIPFamilyDual, CloudV1BminstanceNewParamsInterfacesIPFamilyIpv4, CloudV1BminstanceNewParamsInterfacesIPFamilyIpv6:
		return true
	}
	return false
}

type CloudV1BminstanceListParams struct {
	// Filters the instances by a date and time stamp when the instances last changed.
	// Those instances that changed before or equal to the specified date and time
	// stamp are returned.
	ChangesBefore param.Field[string] `query:"changes-before"`
	// Filters the instances by a date and time stamp when the instances last changed
	// status.
	ChangesSince param.Field[string] `query:"changes-since"`
	// Filter out instances by flavor_id. Flavor id must match exactly. Example:
	// "g1-standard-1-2"
	FlavorID param.Field[string] `query:"flavor_id"`
	// Filter out instances by flavor_prefix. Example: "g1-standard" or "g1-"
	FlavorPrefix param.Field[string] `query:"flavor_prefix"`
	// Include k8s instances. default True.
	IncludeK8s param.Field[bool] `query:"include_k8s"`
	// An IPv4 address to filter results by. Regular expression allowed
	IP param.Field[string] `query:"ip"`
	// Limit the number of returned instances
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/instances/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Filter by metadata values. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_v=["value", "sense"]" --url "http://localhost:1111/v1/instances/1/1"
	MetadataV param.Field[string] `query:"metadata_v"`
	// Filter out instances by name. Use MySQL regular expression. Example:
	// "^.est*....*[0-9]$". Also, any substring can be used and instances will be
	// returned with names containing the substring. Example: "test".
	Name param.Field[string] `query:"name"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Include only isolated instances. default False.
	OnlyIsolated param.Field[bool] `query:"only_isolated"`
	// Return bare metals only with external fixed IP addresses.
	OnlyWithFixedExternalIP param.Field[bool] `query:"only_with_fixed_external_ip"`
	// Order instances by transmitted fields and directions (name.asc).
	OrderBy param.Field[string] `query:"order_by"`
	// Filter result by profile name. Effective only with with_ddos set to true.
	ProfileName param.Field[string] `query:"profile_name"`
	// Filter result by DDoS protection_status. if parameter is provided. Effective
	// only with with_ddos set to true. (Active, Queued or Error)
	ProtectionStatus param.Field[string] `query:"protection_status"`
	// Filters instances by a server status, as a string. Possible statuses are ACTIVE,
	// ERROR, SHUTOFF, REBOOT, PAUSED, etc.
	Status param.Field[string] `query:"status"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with with_ddos set to true. (advanced or basic)
	TypeDdosProfile param.Field[string] `query:"type_ddos_profile"`
	// Filter the server list result by the UUID of the server. Allowed UUID part
	Uuid param.Field[string] `query:"uuid"`
	// DDoS profile information will be included to Bare metal servers
	WithDdos param.Field[bool] `query:"with_ddos"`
}

// URLQuery serializes [CloudV1BminstanceListParams]'s query parameters as
// `url.Values`.
func (r CloudV1BminstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1BminstanceCheckQuotaParams struct {
	// Check quota before bare metal servers will be created.
	CheckQuotaBeforeBmCreation CheckQuotaBeforeBmCreationParam `json:"check_quota_before_bm_creation"`
}

func (r CloudV1BminstanceCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckQuotaBeforeBmCreation)
}

type CloudV1BminstanceGetAvailableFlavorsParams struct {
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
	// Apptemplate ID
	ApptemplateID param.Field[string] `json:"apptemplate_id"`
	// Image ID
	ImageID param.Field[string] `json:"image_id"`
}

func (r CloudV1BminstanceGetAvailableFlavorsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CloudV1BminstanceGetAvailableFlavorsParams]'s query
// parameters as `url.Values`.
func (r CloudV1BminstanceGetAvailableFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1BminstanceRebuildParams struct {
	// Image ID
	ImageID param.Field[string] `json:"image_id"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Field[string] `json:"user_data"`
}

func (r CloudV1BminstanceRebuildParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
