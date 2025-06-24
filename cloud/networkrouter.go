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

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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

// Create a new router with the specified configuration.
func (r *NetworkRouterService) New(ctx context.Context, params NetworkRouterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/routers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the configuration of an existing router.
func (r *NetworkRouterService) Update(ctx context.Context, routerID string, params NetworkRouterUpdateParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all routers in the specified project and region.
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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

// List all routers in the specified project and region.
func (r *NetworkRouterService) ListAutoPaging(ctx context.Context, params NetworkRouterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Router] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific router and all its associated resources.
func (r *NetworkRouterService) Delete(ctx context.Context, routerID string, body NetworkRouterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
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

// Attach a subnet to an existing router.
func (r *NetworkRouterService) AttachSubnet(ctx context.Context, routerID string, params NetworkRouterAttachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/attach", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach a subnet from an existing router.
func (r *NetworkRouterService) DetachSubnet(ctx context.Context, routerID string, params NetworkRouterDetachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/detach", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get detailed information about a specific router.
func (r *NetworkRouterService) Get(ctx context.Context, routerID string, query NetworkRouterGetParams, opts ...option.RequestOption) (res *Router, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Router struct {
	// Router ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the router was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the router is distributed or centralized.
	Distributed bool `json:"distributed,required"`
	// List of router interfaces.
	Interfaces []RouterInterface `json:"interfaces,required"`
	// Router name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of custom routes.
	Routes []Route `json:"routes,required"`
	// Status of the router.
	Status string `json:"status,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Datetime when the router was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// State of this router's external gateway.
	ExternalGatewayInfo RouterExternalGatewayInfo `json:"external_gateway_info,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Distributed         respjson.Field
		Interfaces          respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		Routes              respjson.Field
		Status              respjson.Field
		TaskID              respjson.Field
		UpdatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		ExternalGatewayInfo respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Router) RawJSON() string { return r.JSON.raw }
func (r *Router) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouterInterface struct {
	// IP addresses assigned to this port
	IPAssignments []IPAssignment `json:"ip_assignments,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAssignments respjson.Field
		NetworkID     respjson.Field
		PortID        respjson.Field
		MacAddress    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterInterface) RawJSON() string { return r.JSON.raw }
func (r *RouterInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State of this router's external gateway.
type RouterExternalGatewayInfo struct {
	// Is SNAT enabled.
	EnableSnat bool `json:"enable_snat,required"`
	// List of external IPs that emit SNAT-ed traffic.
	ExternalFixedIPs []IPAssignment `json:"external_fixed_ips,required"`
	// Id of the external network.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableSnat       respjson.Field
		ExternalFixedIPs respjson.Field
		NetworkID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterExternalGatewayInfo) RawJSON() string { return r.JSON.raw }
func (r *RouterExternalGatewayInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouterList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Router `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterList) RawJSON() string { return r.JSON.raw }
func (r *RouterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type SubnetIDParam struct {
	// Target IP is identified by it's subnet
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	paramObj
}

func (r SubnetIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SubnetIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubnetIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// name of router
	Name                string                                         `json:"name,required"`
	ExternalGatewayInfo NetworkRouterNewParamsExternalGatewayInfoUnion `json:"external_gateway_info,omitzero"`
	// List of interfaces to attach to router immediately after creation.
	Interfaces []NetworkRouterNewParamsInterface `json:"interfaces,omitzero"`
	// List of custom routes.
	Routes []NetworkRouterNewParamsRoute `json:"routes,omitzero"`
	paramObj
}

func (r NetworkRouterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NetworkRouterNewParamsExternalGatewayInfoUnion struct {
	OfRouterExternalManualGwSerializer  *NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer  `json:",omitzero,inline"`
	OfRouterExternalDefaultGwSerializer *NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer `json:",omitzero,inline"`
	paramUnion
}

func (u NetworkRouterNewParamsExternalGatewayInfoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRouterExternalManualGwSerializer, u.OfRouterExternalDefaultGwSerializer)
}
func (u *NetworkRouterNewParamsExternalGatewayInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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
	if vt := u.OfRouterExternalManualGwSerializer; vt != nil && vt.EnableSnat.Valid() {
		return &vt.EnableSnat.Value
	} else if vt := u.OfRouterExternalDefaultGwSerializer; vt != nil && vt.EnableSnat.Valid() {
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

// The property NetworkID is required.
type NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer struct {
	// id of the external network.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// must be 'manual'.
	//
	// Any of "manual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsExternalGatewayInfoRouterExternalManualGwSerializer](
		"type", "manual",
	)
}

type NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer struct {
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// must be 'default'.
	//
	// Any of "default".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer](
		"type", "default",
	)
}

// The property SubnetID is required.
type NetworkRouterNewParamsInterface struct {
	// id of the subnet to attach to.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// must be 'subnet'.
	//
	// Any of "subnet".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r NetworkRouterNewParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterNewParamsInterface](
		"type", "subnet",
	)
}

// The properties Destination, Nexthop are required.
type NetworkRouterNewParamsRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

func (r NetworkRouterNewParamsRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsRoute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// New name of router
	Name param.Opt[string] `json:"name,omitzero"`
	// New external gateway.
	ExternalGatewayInfo NetworkRouterUpdateParamsExternalGatewayInfo `json:"external_gateway_info,omitzero"`
	// List of custom routes.
	Routes []NetworkRouterUpdateParamsRoute `json:"routes,omitzero"`
	paramObj
}

func (r NetworkRouterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New external gateway.
//
// The property NetworkID is required.
type NetworkRouterUpdateParamsExternalGatewayInfo struct {
	// id of the external network.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Opt[bool] `json:"enable_snat,omitzero"`
	// must be 'manual'.
	//
	// Any of "manual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r NetworkRouterUpdateParamsExternalGatewayInfo) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParamsExternalGatewayInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateParamsExternalGatewayInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouterUpdateParamsExternalGatewayInfo](
		"type", "manual",
	)
}

// The properties Destination, Nexthop are required.
type NetworkRouterUpdateParamsRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

func (r NetworkRouterUpdateParamsRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParamsRoute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateParamsRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit the number of returned routers
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkRouterListParams]'s query parameters as
// `url.Values`.
func (r NetworkRouterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkRouterDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type NetworkRouterAttachSubnetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Subnet ID on which router interface will be created
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// IP address to assign for router's interface, if not specified, address will be
	// selected automatically
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	paramObj
}

func (r NetworkRouterAttachSubnetParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterAttachSubnetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterAttachSubnetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterDetachSubnetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	SubnetID  SubnetIDParam
	paramObj
}

func (r NetworkRouterDetachSubnetParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.SubnetID)
}
func (r *NetworkRouterDetachSubnetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubnetID)
}

type NetworkRouterGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
