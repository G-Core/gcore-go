// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// InstanceInterfaceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceInterfaceService] method instead.
type InstanceInterfaceService struct {
	Options []option.RequestOption
}

// NewInstanceInterfaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstanceInterfaceService(opts ...option.RequestOption) (r InstanceInterfaceService) {
	r = InstanceInterfaceService{}
	r.Options = opts
	return
}

// List network interfaces attached to the instance
func (r *InstanceInterfaceService) List(ctx context.Context, instanceID string, query InstanceInterfaceListParams, opts ...option.RequestOption) (res *NetworkInterfaceList, err error) {
	opts = append(r.Options[:], opts...)
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/interfaces", query.ProjectID.Value, query.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Attach interface to instance
func (r *InstanceInterfaceService) Attach(ctx context.Context, instanceID string, params InstanceInterfaceAttachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/attach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach interface from instance
func (r *InstanceInterfaceService) Detach(ctx context.Context, instanceID string, params InstanceInterfaceDetachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type InstanceInterfaceListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceInterfaceAttachParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to default external network
	OfNewInterfaceExternalExtendSchemaWithDDOS *InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to specified subnet
	OfNewInterfaceSpecificSubnetSchema *InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the network subnet with the largest count of
	// available ips
	OfNewInterfaceAnySubnetSchema *InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the given port. Floating IP will be created and
	// attached to that IP
	OfNewInterfaceReservedFixedIPSchema *InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema `json:",inline"`

	paramObj
}

func (u InstanceInterfaceAttachParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceInterfaceAttachParams](u.OfNewInterfaceExternalExtendSchemaWithDDOS, u.OfNewInterfaceSpecificSubnetSchema, u.OfNewInterfaceAnySubnetSchema, u.OfNewInterfaceReservedFixedIPSchema)
}

// Instance will be attached to default external network
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'external'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS](
		"IPFamily", false, "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
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

func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to specified subnet
//
// The property SubnetID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema struct {
	// Port will get an IP address from this subnet
	SubnetID string `json:"subnet_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
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

func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to the network subnet with the largest count of
// available ips
//
// The property NetworkID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema struct {
	// Port will get an IP address in this network subnet
	NetworkID string `json:"network_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'any_subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema](
		"IPFamily", false, "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
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

func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
//
// The property PortID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema struct {
	// Port ID
	PortID string `json:"port_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'reserved_fixed_ip'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
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

func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceDetachParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// IP address
	IPAddress string `json:"ip_address,required"`
	// ID of the port
	PortID string `json:"port_id,required"`
	paramObj
}

func (r InstanceInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
