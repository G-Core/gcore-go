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
func (r *GPUBaremetalClusterInterfaceService) List(ctx context.Context, clusterID string, params GPUBaremetalClusterInterfaceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[NetworkInterfaceUnion], err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/interfaces", params.ProjectID.Value, params.RegionID.Value, clusterID)
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

// Retrieve a list of network interfaces attached to the GPU cluster servers.
func (r *GPUBaremetalClusterInterfaceService) ListAutoPaging(ctx context.Context, clusterID string, params GPUBaremetalClusterInterfaceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[NetworkInterfaceUnion] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, clusterID, params, opts...))
}

// Attach interface to bare metal GPU cluster server.
func (r *GPUBaremetalClusterInterfaceService) Attach(ctx context.Context, instanceID string, params GPUBaremetalClusterInterfaceAttachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/attach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Detach interface from bare metal GPU cluster server.
func (r *GPUBaremetalClusterInterfaceService) Detach(ctx context.Context, instanceID string, params GPUBaremetalClusterInterfaceDetachParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type GPUBaremetalClusterInterfaceListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterInterfaceListParams]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterInterfaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterInterfaceAttachParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExternal *GPUBaremetalClusterInterfaceAttachParamsBodyExternal `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSubnet *GPUBaremetalClusterInterfaceAttachParamsBodySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAnySubnet *GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReservedFixedIP *GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP `json:",inline"`

	paramObj
}

func (u GPUBaremetalClusterInterfaceAttachParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}
func (r *GPUBaremetalClusterInterfaceAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceAttachParamsBodyExternal struct {
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup `json:"security_groups,omitzero"`
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Any of "external".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyExternal) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodyExternal](
		"type", "external",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodySubnet struct {
	// Port will get an IP address from this subnet.
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Any of "subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodySubnet](
		"type", "subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NetworkID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet struct {
	// Port will get an IP address in this network subnet.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Any of "any_subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnet](
		"type", "any_subnet",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP struct {
	// Port ID.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Interface name.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs.
	SecurityGroups []GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// Any of "reserved_fixed_ip".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIP](
		"type", "reserved_fixed_ip",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField struct {
	// ID of DDoS profile field.
	BaseField int64 `json:"base_field" api:"required"`
	// Basic type value.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value for the DDoS profile field.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceAttachParamsBodyReservedFixedIPSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterInterfaceDetachParams struct {
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

func (r GPUBaremetalClusterInterfaceDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterInterfaceDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterInterfaceDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
