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

// CloudV1AIClusterGPUService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIClusterGPUService] method instead.
type CloudV1AIClusterGPUService struct {
	Options []option.RequestOption
}

// NewCloudV1AIClusterGPUService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1AIClusterGPUService(opts ...option.RequestOption) (r *CloudV1AIClusterGPUService) {
	r = &CloudV1AIClusterGPUService{}
	r.Options = opts
	return
}

// Create a new AI cluster.
func (r *CloudV1AIClusterGPUService) New(ctx context.Context, projectID string, regionID string, body CloudV1AIClusterGPUNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%s/%s", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Rebuild one or many nodes from GPU cluster. All cluster nodes need to be
// provided to change the cluster image.
func (r *CloudV1AIClusterGPUService) Rebuild(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV1AIClusterGPURebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/rebuild", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove single node from GPU cluster.
func (r *CloudV1AIClusterGPUService) RemoveNode(ctx context.Context, projectID string, regionID string, clusterID string, instanceID string, body CloudV1AIClusterGPURemoveNodeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%s/%s/%s/node/%s", projectID, regionID, clusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Resize an existing AI GPU cluster.
func (r *CloudV1AIClusterGPUService) Resize(ctx context.Context, projectID string, regionID string, clusterID string, body CloudV1AIClusterGPUResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%s/%s/%s/resize", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateInstanceVolume schema
type CreateInstanceVolumeParam struct {
	// Volume source
	Source param.Field[CreateInstanceVolumeSource] `json:"source,required"`
	// App template ID. Mandatory if volume is created from marketplace template
	ApptemplateID param.Field[string] `json:"apptemplate_id"`
	// Block device attachment tag (not exposed in the metadata)
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// 0 should be set for primary boot device Unique positive values for other
	// bootable devices.Negative - boot prohibited
	BootIndex param.Field[int64] `json:"boot_index"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination param.Field[bool] `json:"delete_on_termination"`
	// Image ID. Mandatory if volume is created from image
	ImageID param.Field[string] `json:"image_id"`
	// Create one or more metadata items for a volume
	Metadata param.Field[interface{}] `json:"metadata"`
	// Volume name. Will be generated if missing
	Name param.Field[string] `json:"name"`
	// Volume size. Must be specified when source is 'new-volume' or 'image'. If
	// specified for source 'snapshot' or 'existing-volume', value must be equal to
	// respective snapshot or volume size
	Size param.Field[int64] `json:"size"`
	// Volume snapshot ID. Mandatory if volume is created from a snapshot
	SnapshotID param.Field[string] `json:"snapshot_id"`
	// One of 'standard', 'ssd_hiiops', 'ssd_local', 'ssd_lowlatency', 'cold', 'ultra'
	TypeName param.Field[CreateInstanceVolumeTypeName] `json:"type_name"`
	// Volume ID. Mandatory if volume is pre-existing volume
	VolumeID param.Field[string] `json:"volume_id"`
}

func (r CreateInstanceVolumeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Volume source
type CreateInstanceVolumeSource string

const (
	CreateInstanceVolumeSourceApptemplate    CreateInstanceVolumeSource = "apptemplate"
	CreateInstanceVolumeSourceExistingVolume CreateInstanceVolumeSource = "existing-volume"
	CreateInstanceVolumeSourceImage          CreateInstanceVolumeSource = "image"
	CreateInstanceVolumeSourceNewVolume      CreateInstanceVolumeSource = "new-volume"
	CreateInstanceVolumeSourceSnapshot       CreateInstanceVolumeSource = "snapshot"
)

func (r CreateInstanceVolumeSource) IsKnown() bool {
	switch r {
	case CreateInstanceVolumeSourceApptemplate, CreateInstanceVolumeSourceExistingVolume, CreateInstanceVolumeSourceImage, CreateInstanceVolumeSourceNewVolume, CreateInstanceVolumeSourceSnapshot:
		return true
	}
	return false
}

// One of 'standard', 'ssd_hiiops', 'ssd_local', 'ssd_lowlatency', 'cold', 'ultra'
type CreateInstanceVolumeTypeName string

const (
	CreateInstanceVolumeTypeNameCold          CreateInstanceVolumeTypeName = "cold"
	CreateInstanceVolumeTypeNameSsdHiiops     CreateInstanceVolumeTypeName = "ssd_hiiops"
	CreateInstanceVolumeTypeNameSsdLocal      CreateInstanceVolumeTypeName = "ssd_local"
	CreateInstanceVolumeTypeNameSsdLowlatency CreateInstanceVolumeTypeName = "ssd_lowlatency"
	CreateInstanceVolumeTypeNameStandard      CreateInstanceVolumeTypeName = "standard"
	CreateInstanceVolumeTypeNameUltra         CreateInstanceVolumeTypeName = "ultra"
)

func (r CreateInstanceVolumeTypeName) IsKnown() bool {
	switch r {
	case CreateInstanceVolumeTypeNameCold, CreateInstanceVolumeTypeNameSsdHiiops, CreateInstanceVolumeTypeNameSsdLocal, CreateInstanceVolumeTypeNameSsdLowlatency, CreateInstanceVolumeTypeNameStandard, CreateInstanceVolumeTypeNameUltra:
		return true
	}
	return false
}

// MandatoryIdSchema schema
type MandatoryIDParam struct {
	// Resource ID
	ID param.Field[string] `json:"id,required"`
}

func (r MandatoryIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1AIClusterGPUNewParams struct {
	// Flavor name
	Flavor param.Field[string] `json:"flavor,required"`
	// Image ID
	ImageID param.Field[string] `json:"image_id,required"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]CloudV1AIClusterGPUNewParamsInterfaceUnion] `json:"interfaces,required"`
	// GPU Cluster name
	Name param.Field[string] `json:"name,required"`
	// Number of instances to create
	InstancesCount param.Field[int64] `json:"instances_count"`
	// Keypair name to inject into new cluster(s)
	KeypairName param.Field[string] `json:"keypair_name"`
	// Create one or more metadata items for a cluster
	Metadata param.Field[interface{}] `json:"metadata"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password param.Field[string] `json:"password"`
	// Security group UUIDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Field[string] `json:"user_data"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username param.Field[string] `json:"username"`
	// List of volumes for instances
	Volumes param.Field[[]CreateInstanceVolumeParam] `json:"volumes"`
}

func (r CloudV1AIClusterGPUNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to default external network
type CloudV1AIClusterGPUNewParamsInterface struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV1AIClusterGPUNewParamsInterfacesIPFamily] `json:"ip_family"`
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

func (r CloudV1AIClusterGPUNewParamsInterface) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1AIClusterGPUNewParamsInterface) implementsCloudV1AIClusterGPUNewParamsInterfaceUnion() {
}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalParam], [NetworkSubnetFipParam],
// [NetworkAnySubnetFipParam], [ReservedFixedIPFipParam],
// [CloudV1AIClusterGPUNewParamsInterface].
type CloudV1AIClusterGPUNewParamsInterfaceUnion interface {
	implementsCloudV1AIClusterGPUNewParamsInterfaceUnion()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV1AIClusterGPUNewParamsInterfacesIPFamily string

const (
	CloudV1AIClusterGPUNewParamsInterfacesIPFamilyDual CloudV1AIClusterGPUNewParamsInterfacesIPFamily = "dual"
	CloudV1AIClusterGPUNewParamsInterfacesIPFamilyIpv4 CloudV1AIClusterGPUNewParamsInterfacesIPFamily = "ipv4"
	CloudV1AIClusterGPUNewParamsInterfacesIPFamilyIpv6 CloudV1AIClusterGPUNewParamsInterfacesIPFamily = "ipv6"
)

func (r CloudV1AIClusterGPUNewParamsInterfacesIPFamily) IsKnown() bool {
	switch r {
	case CloudV1AIClusterGPUNewParamsInterfacesIPFamilyDual, CloudV1AIClusterGPUNewParamsInterfacesIPFamilyIpv4, CloudV1AIClusterGPUNewParamsInterfacesIPFamilyIpv6:
		return true
	}
	return false
}

type CloudV1AIClusterGPURebuildParams struct {
	// List of nodes uuids to be rebuild
	Nodes param.Field[[]string] `json:"nodes,required"`
	// AI GPU image ID
	ImageID param.Field[string] `json:"image_id"`
	// String in base64 format.Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Field[string] `json:"user_data"`
}

func (r CloudV1AIClusterGPURebuildParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1AIClusterGPURemoveNodeParams struct {
	// Set False if you do not want to delete assigned floating IPs. By default, it's
	// True.
	DeleteFloatings param.Field[bool] `query:"delete_floatings"`
}

// URLQuery serializes [CloudV1AIClusterGPURemoveNodeParams]'s query parameters as
// `url.Values`.
func (r CloudV1AIClusterGPURemoveNodeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1AIClusterGPUResizeParams struct {
	// Resized (total) number of instances
	InstancesCount param.Field[int64] `json:"instances_count,required"`
}

func (r CloudV1AIClusterGPUResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
