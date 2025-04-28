// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// NetworkRouterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkRouterService] method instead.
type NetworkRouterService struct {
	Options []option.RequestOption
}

// NewNetworkRouterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouterService(opts ...option.RequestOption) (r NetworkRouterService) {
	r = NetworkRouterService{}
	r.Options = opts
	return
}

// Create router
func (r *NetworkRouterService) New(ctx context.Context, params NetworkRouterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/routers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update router
func (r *NetworkRouterService) Update(ctx context.Context, routerID string, params NetworkRouterUpdateParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List routers
func (r *NetworkRouterService) List(ctx context.Context, params NetworkRouterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Router], err error) {
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
	path := fmt.Sprintf("cloud/v1/routers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List routers
func (r *NetworkRouterService) ListAutoPaging(ctx context.Context, params NetworkRouterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Router] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete router
func (r *NetworkRouterService) Delete(ctx context.Context, routerID string, body NetworkRouterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Attach subnet to router
func (r *NetworkRouterService) AttachSubnet(ctx context.Context, routerID string, params NetworkRouterAttachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/attach", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach subnet from router
func (r *NetworkRouterService) DetachSubnet(ctx context.Context, routerID string, params NetworkRouterDetachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/detach", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get specific router
func (r *NetworkRouterService) Get(ctx context.Context, routerID string, query NetworkRouterGetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/RouterSerializer' "$.components.schemas.RouterSerializer"
type Router struct {
	// '#/components/schemas/RouterSerializer/properties/id'
	// "$.components.schemas.RouterSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/RouterSerializer/properties/created_at'
	// "$.components.schemas.RouterSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RouterSerializer/properties/distributed'
	// "$.components.schemas.RouterSerializer.properties.distributed"
	Distributed bool `json:"distributed,required"`
	// '#/components/schemas/RouterSerializer/properties/interfaces'
	// "$.components.schemas.RouterSerializer.properties.interfaces"
	Interfaces []RouterInterface `json:"interfaces,required"`
	// '#/components/schemas/RouterSerializer/properties/name'
	// "$.components.schemas.RouterSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RouterSerializer/properties/project_id'
	// "$.components.schemas.RouterSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/RouterSerializer/properties/region'
	// "$.components.schemas.RouterSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/RouterSerializer/properties/region_id'
	// "$.components.schemas.RouterSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/RouterSerializer/properties/routes'
	// "$.components.schemas.RouterSerializer.properties.routes"
	Routes []NeutronRoute `json:"routes,required"`
	// '#/components/schemas/RouterSerializer/properties/status'
	// "$.components.schemas.RouterSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/RouterSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.RouterSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required" format:"uuid4"`
	// '#/components/schemas/RouterSerializer/properties/updated_at'
	// "$.components.schemas.RouterSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/RouterSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.RouterSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/RouterSerializer/properties/external_gateway_info/anyOf/0'
	// "$.components.schemas.RouterSerializer.properties.external_gateway_info.anyOf[0]"
	ExternalGatewayInfo RouterExternalGatewayInfo `json:"external_gateway_info,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		CreatedAt           resp.Field
		Distributed         resp.Field
		Interfaces          resp.Field
		Name                resp.Field
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Routes              resp.Field
		Status              resp.Field
		TaskID              resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		ExternalGatewayInfo resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Router) RawJSON() string { return r.JSON.raw }
func (r *Router) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RouterSerializer/properties/interfaces/items'
// "$.components.schemas.RouterSerializer.properties.interfaces.items"
type RouterInterface struct {
	// '#/components/schemas/PortSerializer/properties/ip_assignments'
	// "$.components.schemas.PortSerializer.properties.ip_assignments"
	IPAssignments []RouterInterfaceIPAssignment `json:"ip_assignments,required"`
	// '#/components/schemas/PortSerializer/properties/network_id'
	// "$.components.schemas.PortSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/PortSerializer/properties/port_id'
	// "$.components.schemas.PortSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/PortSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.PortSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAssignments resp.Field
		NetworkID     resp.Field
		PortID        resp.Field
		MacAddress    resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterInterface) RawJSON() string { return r.JSON.raw }
func (r *RouterInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/PortSerializer/properties/ip_assignments/items'
// "$.components.schemas.PortSerializer.properties.ip_assignments.items"
type RouterInterfaceIPAssignment struct {
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/ip_address'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *RouterInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RouterSerializer/properties/external_gateway_info/anyOf/0'
// "$.components.schemas.RouterSerializer.properties.external_gateway_info.anyOf[0]"
type RouterExternalGatewayInfo struct {
	// '#/components/schemas/ExternalGatewaySerializer/properties/enable_snat'
	// "$.components.schemas.ExternalGatewaySerializer.properties.enable_snat"
	EnableSnat bool `json:"enable_snat,required"`
	// '#/components/schemas/ExternalGatewaySerializer/properties/external_fixed_ips'
	// "$.components.schemas.ExternalGatewaySerializer.properties.external_fixed_ips"
	ExternalFixedIPs []RouterExternalGatewayInfoExternalFixedIP `json:"external_fixed_ips,required"`
	// '#/components/schemas/ExternalGatewaySerializer/properties/network_id'
	// "$.components.schemas.ExternalGatewaySerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EnableSnat       resp.Field
		ExternalFixedIPs resp.Field
		NetworkID        resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterExternalGatewayInfo) RawJSON() string { return r.JSON.raw }
func (r *RouterExternalGatewayInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ExternalGatewaySerializer/properties/external_fixed_ips/items'
// "$.components.schemas.ExternalGatewaySerializer.properties.external_fixed_ips.items"
type RouterExternalGatewayInfoExternalFixedIP struct {
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/ip_address'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterExternalGatewayInfoExternalFixedIP) RawJSON() string { return r.JSON.raw }
func (r *RouterExternalGatewayInfoExternalFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RouterSerializerList'
// "$.components.schemas.RouterSerializerList"
type RouterList struct {
	// '#/components/schemas/RouterSerializerList/properties/count'
	// "$.components.schemas.RouterSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RouterSerializerList/properties/results'
	// "$.components.schemas.RouterSerializerList.properties.results"
	Results []Router `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterList) RawJSON() string { return r.JSON.raw }
func (r *RouterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SubnetIdSerializer'
// "$.components.schemas.SubnetIdSerializer"
//
// The property SubnetID is required.
type SubnetIDParam struct {
	// '#/components/schemas/SubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.SubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SubnetIDParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r SubnetIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SubnetIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkRouterNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateRouterSerializer/properties/name'
	// "$.components.schemas.CreateRouterSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateRouterSerializer/properties/external_gateway_info'
	// "$.components.schemas.CreateRouterSerializer.properties.external_gateway_info"
	ExternalGatewayInfo NetworkRouterNewParamsExternalGatewayInfoUnion `json:"external_gateway_info,omitzero"`
	// '#/components/schemas/CreateRouterSerializer/properties/interfaces/anyOf/0'
	// "$.components.schemas.CreateRouterSerializer.properties.interfaces.anyOf[0]"
	Interfaces []NetworkRouterNewParamsInterface `json:"interfaces,omitzero"`
	// '#/components/schemas/CreateRouterSerializer/properties/routes/anyOf/0'
	// "$.components.schemas.CreateRouterSerializer.properties.routes.anyOf[0]"
	Routes []NeutronRouteParam `json:"routes,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkRouterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NetworkRouterNewParamsExternalGatewayInfoUnion struct {
	OfRouterExternalManualGwSerializer  *NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer  `json:",omitzero,inline"`
	OfRouterExternalDefaultGwSerializer *NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u NetworkRouterNewParamsExternalGatewayInfoUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u NetworkRouterNewParamsExternalGatewayInfoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[NetworkRouterNewParamsExternalGatewayInfoUnion](u.OfRouterExternalManualGwSerializer, u.OfRouterExternalDefaultGwSerializer)
}

func (u *NetworkRouterNewParamsExternalGatewayInfoUnion) asAny() any {
	if !param.IsOmitted(u.OfRouterExternalManualGwSerializer) {
		return u.OfRouterExternalManualGwSerializer
	} else if !param.IsOmitted(u.OfRouterExternalDefaultGwSerializer) {
		return u.OfRouterExternalDefaultGwSerializer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NetworkRouterNewParamsExternalGatewayInfoUnion) GetNetworkID() *string {
	if vt := u.OfRouterExternalManualGwSerializer; vt != nil {
		return &vt.NetworkID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NetworkRouterNewParamsExternalGatewayInfoUnion) GetEnableSnat() *bool {
	if vt := u.OfRouterExternalManualGwSerializer; vt != nil && vt.EnableSnat.IsPresent() {
		return &vt.EnableSnat.Value
	} else if vt := u.OfRouterExternalDefaultGwSerializer; vt != nil && vt.EnableSnat.IsPresent() {
		return &vt.EnableSnat.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u NetworkRouterNewParamsExternalGatewayInfoUnion) GetType() *string {
	if vt := u.OfRouterExternalManualGwSerializer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRouterExternalDefaultGwSerializer; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// '#/components/schemas/CreateRouterSerializer/properties/external_gateway_info/anyOf/0'
// "$.components.schemas.CreateRouterSerializer.properties.external_gateway_info.anyOf[0]"
//
// The property NetworkID is required.
type NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer struct {
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/network_id'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/enable_snat'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.enable_snat"
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/type'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.type"
	//
	// Any of "manual".
	Type string `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer](
		"Type", false, "manual",
	)
}

// '#/components/schemas/CreateRouterSerializer/properties/external_gateway_info/anyOf/1'
// "$.components.schemas.CreateRouterSerializer.properties.external_gateway_info.anyOf[1]"
type NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer struct {
	// '#/components/schemas/RouterExternalDefaultGwSerializer/properties/enable_snat'
	// "$.components.schemas.RouterExternalDefaultGwSerializer.properties.enable_snat"
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// '#/components/schemas/RouterExternalDefaultGwSerializer/properties/type'
	// "$.components.schemas.RouterExternalDefaultGwSerializer.properties.type"
	//
	// Any of "default".
	Type string `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer](
		"Type", false, "default",
	)
}

// '#/components/schemas/CreateRouterSerializer/properties/interfaces/anyOf/0/items'
// "$.components.schemas.CreateRouterSerializer.properties.interfaces.anyOf[0].items"
//
// The property SubnetID is required.
type NetworkRouterNewParamsInterface struct {
	// '#/components/schemas/CreateRouterInterfaceSubnetSerializer/properties/subnet_id'
	// "$.components.schemas.CreateRouterInterfaceSubnetSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/CreateRouterInterfaceSubnetSerializer/properties/type'
	// "$.components.schemas.CreateRouterInterfaceSubnetSerializer.properties.type"
	//
	// Any of "subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterNewParamsInterface) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r NetworkRouterNewParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsInterface](
		"Type", false, "subnet",
	)
}

type NetworkRouterUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PatchRouterSerializer/properties/name/anyOf/0'
	// "$.components.schemas.PatchRouterSerializer.properties.name.anyOf[0]"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/PatchRouterSerializer/properties/external_gateway_info/anyOf/0'
	// "$.components.schemas.PatchRouterSerializer.properties.external_gateway_info.anyOf[0]"
	ExternalGatewayInfo NetworkRouterUpdateParamsExternalGatewayInfo `json:"external_gateway_info,omitzero"`
	// '#/components/schemas/PatchRouterSerializer/properties/routes/anyOf/0'
	// "$.components.schemas.PatchRouterSerializer.properties.routes.anyOf[0]"
	Routes []NeutronRouteParam `json:"routes,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkRouterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchRouterSerializer/properties/external_gateway_info/anyOf/0'
// "$.components.schemas.PatchRouterSerializer.properties.external_gateway_info.anyOf[0]"
//
// The property NetworkID is required.
type NetworkRouterUpdateParamsExternalGatewayInfo struct {
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/network_id'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/enable_snat'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.enable_snat"
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// '#/components/schemas/RouterExternalManualGwSerializer/properties/type'
	// "$.components.schemas.RouterExternalManualGwSerializer.properties.type"
	//
	// Any of "manual".
	Type string `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterUpdateParamsExternalGatewayInfo) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r NetworkRouterUpdateParamsExternalGatewayInfo) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParamsExternalGatewayInfo
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterUpdateParamsExternalGatewayInfo](
		"Type", false, "manual",
	)
}

type NetworkRouterListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}'].get.parameters[3]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [NetworkRouterListParams]'s query parameters as
// `url.Values`.
func (r NetworkRouterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkRouterDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type NetworkRouterAttachSubnetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fattach/post/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/attach'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fattach/post/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/attach'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fattach/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/attach'].post.requestBody.content['application/json'].schema"
	SubnetID SubnetIDParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterAttachSubnetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkRouterAttachSubnetParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.SubnetID)
}

type NetworkRouterDetachSubnetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fdetach/post/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/detach'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fdetach/post/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/detach'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D%2Fdetach/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}/detach'].post.requestBody.content['application/json'].schema"
	SubnetID SubnetIDParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterDetachSubnetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkRouterDetachSubnetParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.SubnetID)
}

type NetworkRouterGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Frouters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brouter_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/routers/{project_id}/{region_id}/{router_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkRouterGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
