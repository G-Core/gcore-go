// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2InstanceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InstanceService] method instead.
type CloudV2InstanceService struct {
	Options      []option.RequestOption
	MetadataItem *CloudV2InstanceMetadataItemService
}

// NewCloudV2InstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2InstanceService(opts ...option.RequestOption) (r *CloudV2InstanceService) {
	r = &CloudV2InstanceService{}
	r.Options = opts
	r.MetadataItem = NewCloudV2InstanceMetadataItemService(opts...)
	return
}

// Create one or many instances or basic VMs. For Linux instances, use the
// 'username' and 'password' to create a new user. When only 'password' is
// provided, it is set as the password for the default user of the image. The
// 'user_data' is ignored when the 'password' is specified. Use the 'user_data'
// field to provide a cloud-init script in base64 to apply configurations to the
// instance. For Windows instances, the 'username' cannot be specified in the
// request. Use the 'password' field to set the password for the 'Admin' user on
// Windows. Use the 'user_data' field to provide a cloudbase-init script in base64
// to create new users on Windows. The password of the Admin user cannot be updated
// via 'user_data'.
func (r *CloudV2InstanceService) New(ctx context.Context, projectID int64, regionID int64, body CloudV2InstanceNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/instances/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified instance will be calculated
func (r *CloudV2InstanceService) CheckLimits(ctx context.Context, projectID int64, regionID int64, body CloudV2InstanceCheckLimitsParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The action can be one of: start, stop, reboot, powercycle, suspend or resume.
// Suspend and resume are not available for baremetal instances.
func (r *CloudV2InstanceService) PerformAction(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV2InstancePerformActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/%s/action", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV2InstanceNewParams struct {
	// Flavor ID
	Flavor param.Field[string] `json:"flavor,required"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]CloudV2InstanceNewParamsInterfaceUnion] `json:"interfaces,required"`
	// List of volumes for instances
	Volumes param.Field[[]CreateInstanceVolumeParam] `json:"volumes,required"`
	// If true, application ports will be allowed in the security group for instances
	// created from the marketplace application template
	AllowAppPorts param.Field[bool] `json:"allow_app_ports"`
	// Parameters for the application template from the marketplace
	Configuration param.Field[interface{}] `json:"configuration"`
	// Keypair name to inject into new instance(s)
	KeypairName param.Field[string] `json:"keypair_name"`
	// Create one or more metadata items for an instance
	Metadata param.Field[interface{}] `json:"metadata"`
	// List of instance names which will be changed by template. You can use forms
	// 'ip_octets', 'two_ip_octets', 'one_ip_octet'.
	NameTemplates param.Field[[]string] `json:"name_templates"`
	// List of instance names
	Names param.Field[[]string] `json:"names"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// When only 'password' is provided, it is set as the password for the default user
	// of the image. 'user_data' is ignored when 'password' is specified. For Windows
	// instances, 'username' cannot be specified. Use the 'password' field to set the
	// password for the 'Admin' user on Windows. Use the 'user_data' field to provide a
	// script to create new users on Windows. The password of the Admin user cannot be
	// updated via 'user_data'.
	Password param.Field[string] `json:"password"`
	// Security group UUIDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Anti-affinity or affinity or soft-anti-affinity server group ID.
	ServergroupID param.Field[string] `json:"servergroup_id"`
	// String in base64 format. For Linux instances, 'user_data' is ignored when
	// 'password' field is provided. For Windows instances, Admin user password is set
	// by 'password' field and cannot be updated via 'user_data'. Examples of the
	// user_data: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Field[string] `json:"user_data"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// For Windows instances, 'username' cannot be specified. Use 'password' field to
	// set the password for the 'Admin' user on Windows.
	Username param.Field[string] `json:"username"`
}

func (r CloudV2InstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to default external network
type CloudV2InstanceNewParamsInterface struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV2InstanceNewParamsInterfacesIPFamily] `json:"ip_family"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID         param.Field[string]      `json:"port_id"`
	SecurityGroups param.Field[interface{}] `json:"security_groups"`
	// Port is assinged an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r CloudV2InstanceNewParamsInterface) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceNewParamsInterface) implementsCloudV2InstanceNewParamsInterfaceUnion() {}

// Instance will be attached to default external network
//
// Satisfied by
// [CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema],
// [CloudV2InstanceNewParamsInterfacesNewInterfaceSpecificSubnetFipExtendSchema],
// [CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchema],
// [CloudV2InstanceNewParamsInterfacesNewInterfaceReservedFixedIPFipExtendSchema],
// [CloudV2InstanceNewParamsInterface].
type CloudV2InstanceNewParamsInterfaceUnion interface {
	implementsCloudV2InstanceNewParamsInterfaceUnion()
}

// Instance will be attached to default external network
type CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema struct {
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema) implementsCloudV2InstanceNewParamsInterfaceUnion() {
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily string

const (
	CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily = "dual"
	CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyIpv4 CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily = "ipv4"
	CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyIpv6 CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily = "ipv6"
)

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamily) IsKnown() bool {
	switch r {
	case CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual, CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyIpv4, CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyIpv6:
		return true
	}
	return false
}

// Instance will be attached to specified subnet, and floating IP can be created
// and attached to this IP
type CloudV2InstanceNewParamsInterfacesNewInterfaceSpecificSubnetFipExtendSchema struct {
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
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// 'subnet'. Instance will be attached to specified subnet, and floating IP can be
	// created and attached to this IP
	Type param.Field[string] `json:"type"`
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceSpecificSubnetFipExtendSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceSpecificSubnetFipExtendSchema) implementsCloudV2InstanceNewParamsInterfaceUnion() {
}

// Instance will be attached to a subnet with the largest count of free ips.
// Floating IP will be created and attached to that IP
type CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchema struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID param.Field[string] `json:"network_id,required"`
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// 'any_subnet'. Instance will be attached to a subnet with the largest count of
	// free ips. Floating IP will be created and attached to that IP
	Type param.Field[string] `json:"type"`
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchema) implementsCloudV2InstanceNewParamsInterfaceUnion() {
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily string

const (
	CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyDual CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily = "dual"
	CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyIpv4 CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily = "ipv4"
	CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyIpv6 CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily = "ipv6"
)

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamily) IsKnown() bool {
	switch r {
	case CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyDual, CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyIpv4, CloudV2InstanceNewParamsInterfacesNewInterfaceAnySubnetFipExtendSchemaIPFamilyIpv6:
		return true
	}
	return false
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
type CloudV2InstanceNewParamsInterfacesNewInterfaceReservedFixedIPFipExtendSchema struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID param.Field[string] `json:"port_id,required"`
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// 'reserved_fixed_ip'. Instance will be attached to the given port. Floating IP
	// will be created and attached to that IP
	Type param.Field[string] `json:"type"`
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceReservedFixedIPFipExtendSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceNewParamsInterfacesNewInterfaceReservedFixedIPFipExtendSchema) implementsCloudV2InstanceNewParamsInterfaceUnion() {
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV2InstanceNewParamsInterfacesIPFamily string

const (
	CloudV2InstanceNewParamsInterfacesIPFamilyDual CloudV2InstanceNewParamsInterfacesIPFamily = "dual"
	CloudV2InstanceNewParamsInterfacesIPFamilyIpv4 CloudV2InstanceNewParamsInterfacesIPFamily = "ipv4"
	CloudV2InstanceNewParamsInterfacesIPFamilyIpv6 CloudV2InstanceNewParamsInterfacesIPFamily = "ipv6"
)

func (r CloudV2InstanceNewParamsInterfacesIPFamily) IsKnown() bool {
	switch r {
	case CloudV2InstanceNewParamsInterfacesIPFamilyDual, CloudV2InstanceNewParamsInterfacesIPFamilyIpv4, CloudV2InstanceNewParamsInterfacesIPFamilyIpv6:
		return true
	}
	return false
}

type CloudV2InstanceCheckLimitsParams struct {
	Flavor param.Field[string] `json:"flavor"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]CloudV2InstanceCheckLimitsParamsInterfaceUnion] `json:"interfaces"`
	// List of instance name templates. Either this or names must be specified.
	NameTemplates param.Field[[]string] `json:"name_templates"`
	// List of instance names. Either this or name_templates must be specified.
	Names   param.Field[[]string]                                 `json:"names"`
	Volumes param.Field[[]CloudV2InstanceCheckLimitsParamsVolume] `json:"volumes"`
}

func (r CloudV2InstanceCheckLimitsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to default external network
type CloudV2InstanceCheckLimitsParamsInterface struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV2InstanceCheckLimitsParamsInterfacesIPFamily] `json:"ip_family"`
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

func (r CloudV2InstanceCheckLimitsParamsInterface) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstanceCheckLimitsParamsInterface) implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion() {
}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalParam], [NetworkSubnetFipParam],
// [NetworkAnySubnetFipParam], [ReservedFixedIPFipParam],
// [CloudV2InstanceCheckLimitsParamsInterface].
type CloudV2InstanceCheckLimitsParamsInterfaceUnion interface {
	implementsCloudV2InstanceCheckLimitsParamsInterfaceUnion()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV2InstanceCheckLimitsParamsInterfacesIPFamily string

const (
	CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyDual CloudV2InstanceCheckLimitsParamsInterfacesIPFamily = "dual"
	CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyIpv4 CloudV2InstanceCheckLimitsParamsInterfacesIPFamily = "ipv4"
	CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyIpv6 CloudV2InstanceCheckLimitsParamsInterfacesIPFamily = "ipv6"
)

func (r CloudV2InstanceCheckLimitsParamsInterfacesIPFamily) IsKnown() bool {
	switch r {
	case CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyDual, CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyIpv4, CloudV2InstanceCheckLimitsParamsInterfacesIPFamilyIpv6:
		return true
	}
	return false
}

// Schema lists only the fields required to calculate price
type CloudV2InstanceCheckLimitsParamsVolume struct {
	// One of: 'new-volume', 'image', 'snapshot', 'apptemplate'
	Source param.Field[CloudV2InstanceCheckLimitsParamsVolumesSource] `json:"source,required"`
	// Volume size in GiB
	Size param.Field[int64] `json:"size"`
	// Snapshot ID. Mandatory if volume is created from a snapshot
	SnapshotID param.Field[string] `json:"snapshot_id"`
	// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'
	TypeName param.Field[CloudV2InstanceCheckLimitsParamsVolumesTypeName] `json:"type_name"`
}

func (r CloudV2InstanceCheckLimitsParamsVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of: 'new-volume', 'image', 'snapshot', 'apptemplate'
type CloudV2InstanceCheckLimitsParamsVolumesSource string

const (
	CloudV2InstanceCheckLimitsParamsVolumesSourceApptemplate    CloudV2InstanceCheckLimitsParamsVolumesSource = "apptemplate"
	CloudV2InstanceCheckLimitsParamsVolumesSourceExistingVolume CloudV2InstanceCheckLimitsParamsVolumesSource = "existing-volume"
	CloudV2InstanceCheckLimitsParamsVolumesSourceImage          CloudV2InstanceCheckLimitsParamsVolumesSource = "image"
	CloudV2InstanceCheckLimitsParamsVolumesSourceNewVolume      CloudV2InstanceCheckLimitsParamsVolumesSource = "new-volume"
	CloudV2InstanceCheckLimitsParamsVolumesSourceSnapshot       CloudV2InstanceCheckLimitsParamsVolumesSource = "snapshot"
)

func (r CloudV2InstanceCheckLimitsParamsVolumesSource) IsKnown() bool {
	switch r {
	case CloudV2InstanceCheckLimitsParamsVolumesSourceApptemplate, CloudV2InstanceCheckLimitsParamsVolumesSourceExistingVolume, CloudV2InstanceCheckLimitsParamsVolumesSourceImage, CloudV2InstanceCheckLimitsParamsVolumesSourceNewVolume, CloudV2InstanceCheckLimitsParamsVolumesSourceSnapshot:
		return true
	}
	return false
}

// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'
type CloudV2InstanceCheckLimitsParamsVolumesTypeName string

const (
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameCold          CloudV2InstanceCheckLimitsParamsVolumesTypeName = "cold"
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdHiiops     CloudV2InstanceCheckLimitsParamsVolumesTypeName = "ssd_hiiops"
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdLocal      CloudV2InstanceCheckLimitsParamsVolumesTypeName = "ssd_local"
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdLowlatency CloudV2InstanceCheckLimitsParamsVolumesTypeName = "ssd_lowlatency"
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameStandard      CloudV2InstanceCheckLimitsParamsVolumesTypeName = "standard"
	CloudV2InstanceCheckLimitsParamsVolumesTypeNameUltra         CloudV2InstanceCheckLimitsParamsVolumesTypeName = "ultra"
)

func (r CloudV2InstanceCheckLimitsParamsVolumesTypeName) IsKnown() bool {
	switch r {
	case CloudV2InstanceCheckLimitsParamsVolumesTypeNameCold, CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdHiiops, CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdLocal, CloudV2InstanceCheckLimitsParamsVolumesTypeNameSsdLowlatency, CloudV2InstanceCheckLimitsParamsVolumesTypeNameStandard, CloudV2InstanceCheckLimitsParamsVolumesTypeNameUltra:
		return true
	}
	return false
}

type CloudV2InstancePerformActionParams struct {
	Body CloudV2InstancePerformActionParamsBodyUnion `json:"body"`
}

func (r CloudV2InstancePerformActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CloudV2InstancePerformActionParamsBody struct {
	// Instance action name
	Action param.Field[CloudV2InstancePerformActionParamsBodyAction] `json:"action,required"`
	// Used on start instance to activate Advanced DDoS profile
	ActivateProfile param.Field[bool] `json:"activate_profile"`
}

func (r CloudV2InstancePerformActionParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstancePerformActionParamsBody) implementsCloudV2InstancePerformActionParamsBodyUnion() {
}

// Satisfied by
// [CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializer],
// [CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializer],
// [CloudV2InstancePerformActionParamsBody].
type CloudV2InstancePerformActionParamsBodyUnion interface {
	implementsCloudV2InstancePerformActionParamsBodyUnion()
}

type CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializer struct {
	// Instance action name
	Action param.Field[CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerAction] `json:"action,required"`
	// Used on start instance to activate Advanced DDoS profile
	ActivateProfile param.Field[bool] `json:"activate_profile"`
}

func (r CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializer) implementsCloudV2InstancePerformActionParamsBodyUnion() {
}

// Instance action name
type CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerAction string

const (
	CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerActionStart CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerAction = "start"
)

func (r CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerAction) IsKnown() bool {
	switch r {
	case CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerActionStart:
		return true
	}
	return false
}

type CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializer struct {
	// Instance action name
	Action param.Field[CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction] `json:"action,required"`
}

func (r CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializer) implementsCloudV2InstancePerformActionParamsBodyUnion() {
}

// Instance action name
type CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction string

const (
	CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionReboot     CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction = "reboot"
	CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionRebootHard CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction = "reboot_hard"
	CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionResume     CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction = "resume"
	CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionStop       CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction = "stop"
	CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionSuspend    CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction = "suspend"
)

func (r CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerAction) IsKnown() bool {
	switch r {
	case CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionReboot, CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionRebootHard, CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionResume, CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionStop, CloudV2InstancePerformActionParamsBodyBasicActionInstanceSerializerActionSuspend:
		return true
	}
	return false
}

// Instance action name
type CloudV2InstancePerformActionParamsBodyAction string

const (
	CloudV2InstancePerformActionParamsBodyActionStart      CloudV2InstancePerformActionParamsBodyAction = "start"
	CloudV2InstancePerformActionParamsBodyActionReboot     CloudV2InstancePerformActionParamsBodyAction = "reboot"
	CloudV2InstancePerformActionParamsBodyActionRebootHard CloudV2InstancePerformActionParamsBodyAction = "reboot_hard"
	CloudV2InstancePerformActionParamsBodyActionResume     CloudV2InstancePerformActionParamsBodyAction = "resume"
	CloudV2InstancePerformActionParamsBodyActionStop       CloudV2InstancePerformActionParamsBodyAction = "stop"
	CloudV2InstancePerformActionParamsBodyActionSuspend    CloudV2InstancePerformActionParamsBodyAction = "suspend"
)

func (r CloudV2InstancePerformActionParamsBodyAction) IsKnown() bool {
	switch r {
	case CloudV2InstancePerformActionParamsBodyActionStart, CloudV2InstancePerformActionParamsBodyActionReboot, CloudV2InstancePerformActionParamsBodyActionRebootHard, CloudV2InstancePerformActionParamsBodyActionResume, CloudV2InstancePerformActionParamsBodyActionStop, CloudV2InstancePerformActionParamsBodyActionSuspend:
		return true
	}
	return false
}
