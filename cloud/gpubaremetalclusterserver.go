// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// GPUBaremetalClusterServerService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterServerService] method instead.
type GPUBaremetalClusterServerService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterServerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterServerService(opts ...option.RequestOption) (r GPUBaremetalClusterServerService) {
	r = GPUBaremetalClusterServerService{}
	r.Options = opts
	return
}

// Delete bare metal GPU server from cluster
func (r *GPUBaremetalClusterServerService) Delete(ctx context.Context, instanceID string, params GPUBaremetalClusterServerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/node/%s", params.ProjectID.Value, params.RegionID.Value, params.ClusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Attach interface to bare metal GPU cluster server
func (r *GPUBaremetalClusterServerService) AttachInterface(ctx context.Context, instanceID string, params GPUBaremetalClusterServerAttachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
func (r *GPUBaremetalClusterServerService) DetachInterface(ctx context.Context, instanceID string, params GPUBaremetalClusterServerDetachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// Get bare metal GPU cluster server console URL
func (r *GPUBaremetalClusterServerService) GetConsole(ctx context.Context, instanceID string, query GPUBaremetalClusterServerGetConsoleParams, opts ...option.RequestOption) (res *Console, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/get_console", query.ProjectID.Value, query.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stops and then starts the server, effectively performing a hard reboot.
func (r *GPUBaremetalClusterServerService) Powercycle(ctx context.Context, instanceID string, body GPUBaremetalClusterServerPowercycleParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServer, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/powercycle", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot one bare metal GPU cluster server
func (r *GPUBaremetalClusterServerService) Reboot(ctx context.Context, instanceID string, body GPUBaremetalClusterServerRebootParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServer, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/reboot", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type GPUBaremetalClusterServerDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterID string           `path:"cluster_id,required" json:"-"`
	// Set False if you do not want to delete assigned floating IPs. By default, it's
	// True.
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerDeleteParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [GPUBaremetalClusterServerDeleteParams]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterServerDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterServerAttachInterfaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to default external network
	OfNewInterfaceExternalExtendSchemaWithDDOS *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to specified subnet
	OfNewInterfaceSpecificSubnetSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the network subnet with the largest count of
	// available ips
	OfNewInterfaceAnySubnetSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the given port. Floating IP will be created and
	// attached to that IP
	OfNewInterfaceReservedFixedIPSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema `json:",inline"`

	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (u GPUBaremetalClusterServerAttachInterfaceParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GPUBaremetalClusterServerAttachInterfaceParams](u.OfNewInterfaceExternalExtendSchemaWithDDOS, u.OfNewInterfaceSpecificSubnetSchema, u.OfNewInterfaceAnySubnetSchema, u.OfNewInterfaceReservedFixedIPSchema)
}

// Instance will be attached to default external network
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'external'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS](
		"IPFamily", false, "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to specified subnet
//
// The property SubnetID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema struct {
	// Port will get an IP address from this subnet
	SubnetID string `json:"subnet_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to the network subnet with the largest count of
// available ips
//
// The property NetworkID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema struct {
	// Port will get an IP address in this network subnet
	NetworkID string `json:"network_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'any_subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema](
		"IPFamily", false, "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
//
// The property PortID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema struct {
	// Port ID
	PortID string `json:"port_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'reserved_fixed_ip'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerDetachInterfaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// IP address
	IPAddress string `json:"ip_address,required"`
	// ID of the port
	PortID string `json:"port_id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerDetachInterfaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r GPUBaremetalClusterServerDetachInterfaceParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerDetachInterfaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerGetConsoleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerGetConsoleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterServerPowercycleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerPowercycleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterServerRebootParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerRebootParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
