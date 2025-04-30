// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
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

// Remove single node from GPU cluster.
func (r *GPUBaremetalClusterServerService) Delete(ctx context.Context, clusterID string, instanceID string, params GPUBaremetalClusterServerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if params.RegionID == "" {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%s/%s/%s/node/%s", params.ProjectID, params.RegionID, clusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Attach interface to GPU cluster node
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

// Detach interface from GPU cluster node
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

// Get GPU cluster node console URL
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

// Powercycle (stop and start) one GPU cluster node, aka hard reboot
func (r *GPUBaremetalClusterServerService) Powercycle(ctx context.Context, instanceID string, body GPUBaremetalClusterServerPowercycleParams, opts ...option.RequestOption) (res *GPUClusterServer, err error) {
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

// Reboot one GPU cluster node
func (r *GPUBaremetalClusterServerService) Reboot(ctx context.Context, instanceID string, body GPUBaremetalClusterServerRebootParams, opts ...option.RequestOption) (res *GPUClusterServer, err error) {
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
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fnode%2F%7Binstance_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/node/{instance_id}']['delete'].parameters[0].schema"
	ProjectID string `path:"project_id,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fnode%2F%7Binstance_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/node/{instance_id}']['delete'].parameters[1].schema"
	RegionID string `path:"region_id,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fnode%2F%7Binstance_id%7D/delete/parameters/4'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/node/{instance_id}']['delete'].parameters[4]"
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
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema"
	Body GPUBaremetalClusterServerAttachInterfaceParamsBodyUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerAttachInterfaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r GPUBaremetalClusterServerAttachInterfaceParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema"
//
// Satisfied by
// [GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS],
// [GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema],
// [GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema]
// and
// [GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema]
type GPUBaremetalClusterServerAttachInterfaceParamsBodyUnion interface {
	implGPUBaremetalClusterServerAttachInterfaceParamsBodyUnion()
}

func (GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) implGPUBaremetalClusterServerAttachInterfaceParamsBodyUnion() {
}
func (GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema) implGPUBaremetalClusterServerAttachInterfaceParamsBodyUnion() {
}
func (GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema) implGPUBaremetalClusterServerAttachInterfaceParamsBodyUnion() {
}
func (GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema) implGPUBaremetalClusterServerAttachInterfaceParamsBodyUnion() {
}

// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/0'
// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[0]"
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
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
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/ip_family'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.ip_family"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/security_groups'
	// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.security_groups"
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

// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
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

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
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
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceExternalExtendSchemaWithDdos/properties/security_groups/items'
// "$.components.schemas.NewInterfaceExternalExtendSchemaWithDdos.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
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

// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/1'
// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[1]"
//
// The property SubnetID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema struct {
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
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.security_groups"
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

// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
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

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
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
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceSpecificSubnetSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceSpecificSubnetSchema.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
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

// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/2'
// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[2]"
//
// The property NetworkID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema struct {
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
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/ip_family'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.ip_family"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.security_groups"
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

// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
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

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
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
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceAnySubnetSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceAnySubnetSchema.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
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

// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fattach_interface/post/requestBody/content/application%2Fjson/schema/anyOf/3'
// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/attach_interface'].post.requestBody.content['application/json'].schema.anyOf[3]"
//
// The property PortID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema struct {
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
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/security_groups'
	// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.security_groups"
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

// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/ddos_profile/allOf/0'
// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.ddos_profile.allOf[0]"
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/profile_template_name'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.profile_template_name"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields'
	// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields"
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

// '#/components/schemas/DeprecatedCreateDdosProfileSchema/properties/fields/items'
// "$.components.schemas.DeprecatedCreateDdosProfileSchema.properties.fields.items"
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
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
func (f GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceReservedFixedIpSchema/properties/security_groups/items'
// "$.components.schemas.NewInterfaceReservedFixedIpSchema.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSchema/properties/id'
	// "$.components.schemas.MandatoryIdSchema.properties.id"
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
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdetach_interface/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/detach_interface'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdetach_interface/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/detach_interface'].post.parameters[1].schema"
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
func (f GPUBaremetalClusterServerDetachInterfaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r GPUBaremetalClusterServerDetachInterfaceParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerDetachInterfaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterServerGetConsoleParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fget_console/get/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/get_console'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fget_console/get/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/get_console'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerGetConsoleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterServerPowercycleParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fpowercycle/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/powercycle'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fpowercycle/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/powercycle'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerPowercycleParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterServerRebootParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Freboot/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/reboot'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Freboot/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{instance_id}/reboot'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterServerRebootParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
