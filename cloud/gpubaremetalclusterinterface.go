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
)

// GPUBaremetalClusterInterfaceService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterInterfaceService] method instead.
type GPUBaremetalClusterInterfaceService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterInterfaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterInterfaceService(opts ...option.RequestOption) (r GPUBaremetalClusterInterfaceService) {
	r = GPUBaremetalClusterInterfaceService{}
	r.Options = opts
	return
}

// Retrieve a list of network interfaces attached to the GPU cluster servers.
func (r *GPUBaremetalClusterInterfaceService) List(ctx context.Context, clusterID string, query GPUBaremetalClusterInterfaceListParams, opts ...option.RequestOption) (res *NetworkInterfaceList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/interfaces", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Attach interface to bare metal GPU cluster server
func (r *GPUBaremetalClusterInterfaceService) Attach(ctx context.Context, instanceID string, params GPUBaremetalClusterInterfaceAttachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/attach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach interface from bare metal GPU cluster server
func (r *GPUBaremetalClusterInterfaceService) Detach(ctx context.Context, instanceID string, params GPUBaremetalClusterInterfaceDetachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GPUBaremetalClusterInterfaceListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterInterfaceAttachParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to default external network
	OfNewInterfaceExternalExtendSchemaWithDDOS *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to specified subnet
	OfNewInterfaceSpecificSubnetSchema *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the network subnet with the largest count of
	// available ips
	OfNewInterfaceAnySubnetSchema *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the given port. Floating IP will be created and
	// attached to that IP
	OfNewInterfaceReservedFixedIPSchema *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema `json:",inline"`

	paramObj
}

func (u GPUBaremetalClusterInterfaceAttachParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNewInterfaceExternalExtendSchemaWithDDOS, u.OfNewInterfaceSpecificSubnetSchema, u.OfNewInterfaceAnySubnetSchema, u.OfNewInterfaceReservedFixedIPSchema)
}
func (r *GPUBaremetalClusterInterfaceAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to default external network
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'external'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to specified subnet
//
// The property SubnetID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema struct {
	// Port will get an IP address from this subnet
	SubnetID string `json:"subnet_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to the network subnet with the largest count of
// available ips
//
// The property NetworkID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema struct {
	// Port will get an IP address in this network subnet
	NetworkID string `json:"network_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'any_subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
//
// The property PortID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema struct {
	// Port ID
	PortID string `json:"port_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'reserved_fixed_ip'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceDetachParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// IP address
	IPAddress string `json:"ip_address,required"`
	// ID of the port
	PortID string `json:"port_id,required"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
