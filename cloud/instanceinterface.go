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

// List all network interfaces attached to the specified instance.
func (r *InstanceInterfaceService) List(ctx context.Context, instanceID string, params InstanceInterfaceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[NetworkInterfaceUnion], err error) {
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

// List all network interfaces attached to the specified instance.
func (r *InstanceInterfaceService) ListAutoPaging(ctx context.Context, instanceID string, params InstanceInterfaceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[NetworkInterfaceUnion] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, instanceID, params, opts...))
}

// Attach interface to instance
func (r *InstanceInterfaceService) Attach(ctx context.Context, instanceID string, params InstanceInterfaceAttachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// Detach interface from instance
func (r *InstanceInterfaceService) Detach(ctx context.Context, instanceID string, params InstanceInterfaceDetachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

type InstanceInterfaceListParams struct {
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

// URLQuery serializes [InstanceInterfaceListParams]'s query parameters as
// `url.Values`.
func (r InstanceInterfaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceInterfaceAttachParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExternal *InstanceInterfaceAttachParamsBodyExternal `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSubnet *InstanceInterfaceAttachParamsBodySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAnySubnet *InstanceInterfaceAttachParamsBodyAnySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReservedFixedIP *InstanceInterfaceAttachParamsBodyReservedFixedIP `json:",inline"`

	paramObj
}

func (u InstanceInterfaceAttachParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}
func (r *InstanceInterfaceAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceInterfaceAttachParamsBodyExternal struct {
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyExternalDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstanceInterfaceAttachParamsBodyExternalSecurityGroup `json:"security_groups,omitzero"`
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

func (r InstanceInterfaceAttachParamsBodyExternal) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodyExternal](
		"type", "external",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyExternalDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyExternalDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyExternalDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyExternalDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyExternalDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstanceInterfaceAttachParamsBodyExternalDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyExternalDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyExternalDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyExternalDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstanceInterfaceAttachParamsBodyExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type InstanceInterfaceAttachParamsBodySubnet struct {
	// Port will get an IP address from this subnet.
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstanceInterfaceAttachParamsBodySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Any of "subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodySubnet](
		"type", "subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstanceInterfaceAttachParamsBodySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstanceInterfaceAttachParamsBodySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NetworkID is required.
type InstanceInterfaceAttachParamsBodyAnySubnet struct {
	// Port will get an IP address in this network subnet.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstanceInterfaceAttachParamsBodyAnySubnetSecurityGroup `json:"security_groups,omitzero"`
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

func (r InstanceInterfaceAttachParamsBodyAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodyAnySubnet](
		"type", "any_subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyAnySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstanceInterfaceAttachParamsBodyAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortID is required.
type InstanceInterfaceAttachParamsBodyReservedFixedIP struct {
	// Port ID.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []InstanceInterfaceAttachParamsBodyReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// Any of "reserved_fixed_ip".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceInterfaceAttachParamsBodyReservedFixedIP](
		"type", "reserved_fixed_ip",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstanceInterfaceAttachParamsBodyReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceAttachParamsBodyReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceInterfaceDetachParams struct {
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

func (r InstanceInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceInterfaceDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
