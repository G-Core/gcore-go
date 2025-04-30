// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// NetworkSubnetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetService] method instead.
type NetworkSubnetService struct {
	Options []option.RequestOption
}

// NewNetworkSubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkSubnetService(opts ...option.RequestOption) (r NetworkSubnetService) {
	r = NetworkSubnetService{}
	r.Options = opts
	return
}

// Create subnet
func (r *NetworkSubnetService) New(ctx context.Context, params NetworkSubnetNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change subnet properties
func (r *NetworkSubnetService) Update(ctx context.Context, subnetID string, params NetworkSubnetUpdateParams, opts ...option.RequestOption) (res *Subnet, err error) {
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List subnets
func (r *NetworkSubnetService) List(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Subnet], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List subnets
func (r *NetworkSubnetService) ListAutoPaging(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Subnet] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete subnet
func (r *NetworkSubnetService) Delete(ctx context.Context, subnetID string, body NetworkSubnetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get subnet
func (r *NetworkSubnetService) Get(ctx context.Context, subnetID string, query NetworkSubnetGetParams, opts ...option.RequestOption) (res *Subnet, err error) {
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NetworkSubnetNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSubnetSerializer/properties/cidr'
	// "$.components.schemas.CreateSubnetSerializer.properties.cidr"
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// '#/components/schemas/CreateSubnetSerializer/properties/name'
	// "$.components.schemas.CreateSubnetSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateSubnetSerializer/properties/network_id'
	// "$.components.schemas.CreateSubnetSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/CreateSubnetSerializer/properties/gateway_ip/anyOf/0'
	// "$.components.schemas.CreateSubnetSerializer.properties.gateway_ip.anyOf[0]"
	GatewayIP param.Opt[string] `json:"gateway_ip,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateSubnetSerializer/properties/router_id_to_connect/anyOf/0'
	// "$.components.schemas.CreateSubnetSerializer.properties.router_id_to_connect.anyOf[0]"
	RouterIDToConnect param.Opt[string] `json:"router_id_to_connect,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateSubnetSerializer/properties/connect_to_network_router'
	// "$.components.schemas.CreateSubnetSerializer.properties.connect_to_network_router"
	ConnectToNetworkRouter param.Opt[bool] `json:"connect_to_network_router,omitzero"`
	// '#/components/schemas/CreateSubnetSerializer/properties/enable_dhcp'
	// "$.components.schemas.CreateSubnetSerializer.properties.enable_dhcp"
	EnableDhcp param.Opt[bool] `json:"enable_dhcp,omitzero"`
	// '#/components/schemas/CreateSubnetSerializer/properties/dns_nameservers/anyOf/0'
	// "$.components.schemas.CreateSubnetSerializer.properties.dns_nameservers.anyOf[0]"
	DNSNameservers []string `json:"dns_nameservers,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateSubnetSerializer/properties/host_routes/anyOf/0'
	// "$.components.schemas.CreateSubnetSerializer.properties.host_routes.anyOf[0]"
	HostRoutes []NetworkSubnetNewParamsHostRoute `json:"host_routes,omitzero"`
	// '#/components/schemas/CreateSubnetSerializer/properties/ip_version'
	// "$.components.schemas.CreateSubnetSerializer.properties.ip_version"
	//
	// Any of 4, 6.
	IPVersion int64 `json:"ip_version,omitzero"`
	// '#/components/schemas/CreateSubnetSerializer/properties/tags'
	// "$.components.schemas.CreateSubnetSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkSubnetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateSubnetSerializer/properties/host_routes/anyOf/0/items'
// "$.components.schemas.CreateSubnetSerializer.properties.host_routes.anyOf[0].items"
//
// The properties Destination, Nexthop are required.
type NetworkSubnetNewParamsHostRoute struct {
	// '#/components/schemas/RouteInSerializer/properties/destination'
	// "$.components.schemas.RouteInSerializer.properties.destination"
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// '#/components/schemas/RouteInSerializer/properties/nexthop'
	// "$.components.schemas.RouteInSerializer.properties.nexthop"
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetNewParamsHostRoute) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r NetworkSubnetNewParamsHostRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetNewParamsHostRoute
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkSubnetUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PatchSubnetSerializer/properties/enable_dhcp/anyOf/0'
	// "$.components.schemas.PatchSubnetSerializer.properties.enable_dhcp.anyOf[0]"
	EnableDhcp param.Opt[bool] `json:"enable_dhcp,omitzero"`
	// '#/components/schemas/PatchSubnetSerializer/properties/gateway_ip/anyOf/0'
	// "$.components.schemas.PatchSubnetSerializer.properties.gateway_ip.anyOf[0]"
	GatewayIP param.Opt[string] `json:"gateway_ip,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/PatchSubnetSerializer/properties/name/anyOf/0'
	// "$.components.schemas.PatchSubnetSerializer.properties.name.anyOf[0]"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/PatchSubnetSerializer/properties/dns_nameservers/anyOf/0'
	// "$.components.schemas.PatchSubnetSerializer.properties.dns_nameservers.anyOf[0]"
	DNSNameservers []string `json:"dns_nameservers,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/PatchSubnetSerializer/properties/host_routes/anyOf/0'
	// "$.components.schemas.PatchSubnetSerializer.properties.host_routes.anyOf[0]"
	HostRoutes []NetworkSubnetUpdateParamsHostRoute `json:"host_routes,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkSubnetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchSubnetSerializer/properties/host_routes/anyOf/0/items'
// "$.components.schemas.PatchSubnetSerializer.properties.host_routes.anyOf[0].items"
//
// The properties Destination, Nexthop are required.
type NetworkSubnetUpdateParamsHostRoute struct {
	// '#/components/schemas/RouteInSerializer/properties/destination'
	// "$.components.schemas.RouteInSerializer.properties.destination"
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// '#/components/schemas/RouteInSerializer/properties/nexthop'
	// "$.components.schemas.RouteInSerializer.properties.nexthop"
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetUpdateParamsHostRoute) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NetworkSubnetUpdateParamsHostRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetUpdateParamsHostRoute
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkSubnetListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[3]"
	NetworkID param.Opt[string] `query:"network_id,omitzero" format:"uuid4" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[4]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[7]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[5]"
	//
	// Any of "available_ips.asc", "available_ips.desc", "cidr.asc", "cidr.desc",
	// "created_at.asc", "created_at.desc", "name.asc", "name.desc", "total_ips.asc",
	// "total_ips.desc", "updated_at.asc", "updated_at.desc".
	OrderBy NetworkSubnetListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[6]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [NetworkSubnetListParams]'s query parameters as
// `url.Values`.
func (r NetworkSubnetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}'].get.parameters[5]"
type NetworkSubnetListParamsOrderBy string

const (
	NetworkSubnetListParamsOrderByAvailableIPsAsc  NetworkSubnetListParamsOrderBy = "available_ips.asc"
	NetworkSubnetListParamsOrderByAvailableIPsDesc NetworkSubnetListParamsOrderBy = "available_ips.desc"
	NetworkSubnetListParamsOrderByCidrAsc          NetworkSubnetListParamsOrderBy = "cidr.asc"
	NetworkSubnetListParamsOrderByCidrDesc         NetworkSubnetListParamsOrderBy = "cidr.desc"
	NetworkSubnetListParamsOrderByCreatedAtAsc     NetworkSubnetListParamsOrderBy = "created_at.asc"
	NetworkSubnetListParamsOrderByCreatedAtDesc    NetworkSubnetListParamsOrderBy = "created_at.desc"
	NetworkSubnetListParamsOrderByNameAsc          NetworkSubnetListParamsOrderBy = "name.asc"
	NetworkSubnetListParamsOrderByNameDesc         NetworkSubnetListParamsOrderBy = "name.desc"
	NetworkSubnetListParamsOrderByTotalIPsAsc      NetworkSubnetListParamsOrderBy = "total_ips.asc"
	NetworkSubnetListParamsOrderByTotalIPsDesc     NetworkSubnetListParamsOrderBy = "total_ips.desc"
	NetworkSubnetListParamsOrderByUpdatedAtAsc     NetworkSubnetListParamsOrderBy = "updated_at.asc"
	NetworkSubnetListParamsOrderByUpdatedAtDesc    NetworkSubnetListParamsOrderBy = "updated_at.desc"
)

type NetworkSubnetDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type NetworkSubnetGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsubnets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsubnet_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/subnets/{project_id}/{region_id}/{subnet_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkSubnetGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
