// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
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
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type InstanceInterfaceListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Finterfaces/get/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/interfaces'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Finterfaces/get/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/interfaces'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InstanceInterfaceAttachParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema"
	Body InstanceInterfaceAttachParamsBodyUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceInterfaceAttachParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema"
//
// Satisfied by
// [InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS],
// [InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema],
// [InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema] and
// [InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema]
type InstanceInterfaceAttachParamsBodyUnion interface {
	implInstanceInterfaceAttachParamsBodyUnion()
}

func (InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) implInstanceInterfaceAttachParamsBodyUnion() {
}
func (InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) implInstanceInterfaceAttachParamsBodyUnion() {
}
func (InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema) implInstanceInterfaceAttachParamsBodyUnion() {
}
func (InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) implInstanceInterfaceAttachParamsBodyUnion() {
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/0'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[0]"
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/interface_name'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/port_group'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/type'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.type"
	Type param.Opt[string] `json:"type,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/ddos_profile/allOf/0'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.ddos_profile.allOf[0]"
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/ip_family'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.ip_family"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/security_groups'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.security_groups"
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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

// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/base_field'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.base_field"
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_name'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_name"
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.value"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_value"
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/security_groups/items'
// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.security_groups.items"
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/1'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[1]"
//
// The property SubnetID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema struct {
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/subnet_id'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/interface_name'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/port_group'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/type'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.type"
	Type param.Opt[string] `json:"type,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/ddos_profile/allOf/0'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.ddos_profile.allOf[0]"
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.security_groups"
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/base_field'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.base_field"
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_name'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_name"
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.value"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_value"
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.security_groups.items"
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/2'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[2]"
//
// The property NetworkID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema struct {
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/network_id'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.network_id"
	NetworkID string `json:"network_id,required"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/interface_name'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/port_group'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/type'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.type"
	Type param.Opt[string] `json:"type,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/ddos_profile/allOf/0'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.ddos_profile.allOf[0]"
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/ip_family'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.ip_family"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.security_groups"
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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

// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/base_field'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.base_field"
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_name'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_name"
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.value"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_value"
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.security_groups.items"
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/3'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[3]"
//
// The property PortID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema struct {
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/port_id'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.port_id"
	PortID string `json:"port_id,required"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/interface_name'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/port_group'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/type'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.type"
	Type param.Opt[string] `json:"type,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/ddos_profile/allOf/0'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.ddos_profile.allOf[0]"
	DDOSProfile InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.security_groups"
	SecurityGroups []InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
	Fields []InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/base_field'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.base_field"
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_name'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_name"
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.value"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/DeprecatedCreateClientProfileFieldSchema/properties/field_value'
	// "$.components.schemas.DeprecatedCreateClientProfileFieldSchema.properties.field_value"
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.security_groups.items"
//
// The property ID is required.
type InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
	ID string `json:"id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceInterfaceDetachParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdetach_interface/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/detach_interface'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdetach_interface/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/detach_interface'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PortIdWithIpSchema/properties/ip_address'
	// "$.components.schemas.PortIdWithIpSchema.properties.ip_address"
	IPAddress string `json:"ip_address,required"`
	// '#/components/schemas/PortIdWithIpSchema/properties/port_id'
	// "$.components.schemas.PortIdWithIpSchema.properties.port_id"
	PortID string `json:"port_id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceInterfaceDetachParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
