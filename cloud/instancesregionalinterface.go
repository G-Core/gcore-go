// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// InstancesRegionalInterfaceService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstancesRegionalInterfaceService] method instead.
type InstancesRegionalInterfaceService struct {
	Options []option.RequestOption
}

// NewInstancesRegionalInterfaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInstancesRegionalInterfaceService(opts ...option.RequestOption) (r InstancesRegionalInterfaceService) {
	r = InstancesRegionalInterfaceService{}
	r.Options = opts
	return
}

// Use
// [GET /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.get)
// instead. Read `interfaces`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// List all network interfaces attached to the specified instance.
func (r *InstancesRegionalInterfaceService) List(ctx context.Context, instanceID string, params InstancesRegionalInterfaceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[NetworkInterfaceUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/interfaces", params.ProjectID.Value, params.RegionID.Value, instanceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Use
// [GET /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.get)
// instead. Read `interfaces`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// List all network interfaces attached to the specified instance.
func (r *InstancesRegionalInterfaceService) ListAutoPaging(ctx context.Context, instanceID string, params InstancesRegionalInterfaceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[NetworkInterfaceUnion] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, instanceID, params, opts...))
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. Send `interfaces`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Attach interface to instance
func (r *InstancesRegionalInterfaceService) Attach(ctx context.Context, instanceID string, params InstancesRegionalInterfaceAttachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/attach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. Send `interfaces`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Detach interface from instance
func (r *InstancesRegionalInterfaceService) Detach(ctx context.Context, instanceID string, params InstancesRegionalInterfaceDetachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type InstancesRegionalInterfaceListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstancesRegionalInterfaceListParams]'s query parameters as
// `url.Values`.
func (r InstancesRegionalInterfaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstancesRegionalInterfaceAttachParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExternal *InstancesRegionalInterfaceAttachParamsBodyExternal `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSubnet *InstancesRegionalInterfaceAttachParamsBodySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAnySubnet *InstancesRegionalInterfaceAttachParamsBodyAnySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReservedFixedIP *InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP `json:",inline"`

	paramObj
}

func (u InstancesRegionalInterfaceAttachParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}
func (r *InstancesRegionalInterfaceAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalInterfaceAttachParamsBodyExternal struct {
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstancesRegionalInterfaceAttachParamsBodyExternalSecurityGroup `json:"security_groups,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both. If omitted, the API selects
	// `ipv4` when the network has an IPv4 subnet, `ipv6` otherwise.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Port will get an IP address in a subnet of the external network with the largest
	// count of free IPs. If the instance already has an IP address in a subnet of the
	// external network with the same IP family, the API tries to reuse that subnet.
	//
	// Any of "external".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyExternal) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalInterfaceAttachParamsBodyExternal](
		"type", "external",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyExternalDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalInterfaceAttachParamsBodyExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type InstancesRegionalInterfaceAttachParamsBodySubnet struct {
	// Port will get an IP address from this subnet.
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstancesRegionalInterfaceAttachParamsBodySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Any of "subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalInterfaceAttachParamsBodySubnet](
		"type", "subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalInterfaceAttachParamsBodySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NetworkID is required.
type InstancesRegionalInterfaceAttachParamsBodyAnySubnet struct {
	// Port will get an IP address in this network subnet.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstancesRegionalInterfaceAttachParamsBodyAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both. If omitted, the API selects
	// `ipv4` when the network has an IPv4 subnet, `ipv6` otherwise.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Port will get an IP address in the subnet with the largest count of free IPs. If
	// the instance already has an IP address in a subnet of `network_id` with the same
	// IP family, the API tries to reuse that subnet.
	//
	// Any of "any_subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalInterfaceAttachParamsBodyAnySubnet](
		"type", "any_subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyAnySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalInterfaceAttachParamsBodyAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortID is required.
type InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP struct {
	// Port ID.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// Any of "reserved_fixed_ip".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalInterfaceAttachParamsBodyReservedFixedIP](
		"type", "reserved_fixed_ip",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalInterfaceDetachParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// IP address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// ID of the port
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalInterfaceDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
