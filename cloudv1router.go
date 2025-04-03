// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1RouterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RouterService] method instead.
type CloudV1RouterService struct {
	Options []option.RequestOption
}

// NewCloudV1RouterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1RouterService(opts ...option.RequestOption) (r *CloudV1RouterService) {
	r = &CloudV1RouterService{}
	r.Options = opts
	return
}

// Create router
func (r *CloudV1RouterService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1RouterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/routers/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get specific router
func (r *CloudV1RouterService) Get(ctx context.Context, projectID int64, regionID int64, routerID string, opts ...option.RequestOption) (res *Router, err error) {
	opts = append(r.Options[:], opts...)
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", projectID, regionID, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update router
func (r *CloudV1RouterService) Update(ctx context.Context, projectID int64, regionID int64, routerID string, body CloudV1RouterUpdateParams, opts ...option.RequestOption) (res *Router, err error) {
	opts = append(r.Options[:], opts...)
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", projectID, regionID, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List routers
func (r *CloudV1RouterService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1RouterListParams, opts ...option.RequestOption) (res *CloudV1RouterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/routers/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete router
func (r *CloudV1RouterService) Delete(ctx context.Context, projectID int64, regionID int64, routerID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s", projectID, regionID, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Attach subnet to router
func (r *CloudV1RouterService) AttachSubnet(ctx context.Context, projectID int64, regionID int64, routerID string, body CloudV1RouterAttachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
	opts = append(r.Options[:], opts...)
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/attach", projectID, regionID, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach subnet from router
func (r *CloudV1RouterService) DetachSubnet(ctx context.Context, projectID int64, regionID int64, routerID string, body CloudV1RouterDetachSubnetParams, opts ...option.RequestOption) (res *Router, err error) {
	opts = append(r.Options[:], opts...)
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/detach", projectID, regionID, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NeutronRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string           `json:"nexthop,required" format:"ipvanyaddress"`
	JSON    neutronRouteJSON `json:"-"`
}

// neutronRouteJSON contains the JSON metadata for the struct [NeutronRoute]
type neutronRouteJSON struct {
	Destination apijson.Field
	Nexthop     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NeutronRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r neutronRouteJSON) RawJSON() string {
	return r.raw
}

type NeutronRouteParam struct {
	// CIDR of destination IPv4 subnet.
	Destination param.Field[string] `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop param.Field[string] `json:"nexthop,required" format:"ipvanyaddress"`
}

func (r NeutronRouteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PortIPSubnetID struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// ID of the subnet that allocated the IP
	SubnetID string             `json:"subnet_id,required" format:"uuid4"`
	JSON     portIPSubnetIDJSON `json:"-"`
}

// portIPSubnetIDJSON contains the JSON metadata for the struct [PortIPSubnetID]
type portIPSubnetIDJSON struct {
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PortIPSubnetID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portIPSubnetIDJSON) RawJSON() string {
	return r.raw
}

type Router struct {
	// router ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the router was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// List of router interfaces.
	Interfaces []RouterInterface `json:"interfaces,required"`
	// router name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of custom routes.
	Routes []NeutronRoute `json:"routes,required"`
	// Status of the router.
	Status string `json:"status,required"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,required,nullable" format:"uuid4"`
	// Datetime when the router was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// State of this router's external gateway.
	ExternalGatewayInfo RouterExternalGatewayInfo `json:"external_gateway_info,nullable"`
	JSON                routerJSON                `json:"-"`
}

// routerJSON contains the JSON metadata for the struct [Router]
type routerJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Interfaces          apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Routes              apijson.Field
	Status              apijson.Field
	TaskID              apijson.Field
	UpdatedAt           apijson.Field
	CreatorTaskID       apijson.Field
	ExternalGatewayInfo apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Router) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routerJSON) RawJSON() string {
	return r.raw
}

type RouterInterface struct {
	// IP addresses assigned to this port
	IPAssignments []PortIPSubnetID `json:"ip_assignments,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// MAC address of the virtual port
	MacAddress string              `json:"mac_address,nullable"`
	JSON       routerInterfaceJSON `json:"-"`
}

// routerInterfaceJSON contains the JSON metadata for the struct [RouterInterface]
type routerInterfaceJSON struct {
	IPAssignments apijson.Field
	NetworkID     apijson.Field
	PortID        apijson.Field
	MacAddress    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RouterInterface) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routerInterfaceJSON) RawJSON() string {
	return r.raw
}

// State of this router's external gateway.
type RouterExternalGatewayInfo struct {
	// Is SNAT enabled.
	EnableSnat bool `json:"enable_snat,required"`
	// List of external IPs that emit SNAT-ed traffic.
	ExternalFixedIPs []PortIPSubnetID `json:"external_fixed_ips,required"`
	// Id of the external network.
	NetworkID string                        `json:"network_id,required" format:"uuid4"`
	JSON      routerExternalGatewayInfoJSON `json:"-"`
}

// routerExternalGatewayInfoJSON contains the JSON metadata for the struct
// [RouterExternalGatewayInfo]
type routerExternalGatewayInfoJSON struct {
	EnableSnat       apijson.Field
	ExternalFixedIPs apijson.Field
	NetworkID        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouterExternalGatewayInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routerExternalGatewayInfoJSON) RawJSON() string {
	return r.raw
}

type RouterExternalManualGwParam struct {
	// id of the external network.
	NetworkID param.Field[string] `json:"network_id,required" format:"uuid4"`
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Field[bool] `json:"enable_snat"`
	// must be 'manual'.
	Type param.Field[RouterExternalManualGwType] `json:"type"`
}

func (r RouterExternalManualGwParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RouterExternalManualGwParam) implementsCloudV1RouterNewParamsExternalGatewayInfoUnion() {}

// must be 'manual'.
type RouterExternalManualGwType string

const (
	RouterExternalManualGwTypeManual RouterExternalManualGwType = "manual"
)

func (r RouterExternalManualGwType) IsKnown() bool {
	switch r {
	case RouterExternalManualGwTypeManual:
		return true
	}
	return false
}

type SubnetIDParam struct {
	// Target IP is identified by it's subnet
	SubnetID param.Field[string] `json:"subnet_id,required" format:"uuid4"`
}

func (r SubnetIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RouterListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Router                      `json:"results,required"`
	JSON    cloudV1RouterListResponseJSON `json:"-"`
}

// cloudV1RouterListResponseJSON contains the JSON metadata for the struct
// [CloudV1RouterListResponse]
type cloudV1RouterListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1RouterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RouterListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1RouterNewParams struct {
	// name of router
	Name                param.Field[string]                                         `json:"name,required"`
	ExternalGatewayInfo param.Field[CloudV1RouterNewParamsExternalGatewayInfoUnion] `json:"external_gateway_info"`
	// List of interfaces to attach to router immediately after creation.
	Interfaces param.Field[[]CloudV1RouterNewParamsInterface] `json:"interfaces"`
	// List of custom routes.
	Routes param.Field[[]NeutronRouteParam] `json:"routes"`
}

func (r CloudV1RouterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RouterNewParamsExternalGatewayInfo struct {
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Field[bool] `json:"enable_snat"`
	// id of the external network.
	NetworkID param.Field[string] `json:"network_id" format:"uuid4"`
	// must be 'manual'.
	Type param.Field[CloudV1RouterNewParamsExternalGatewayInfoType] `json:"type"`
}

func (r CloudV1RouterNewParamsExternalGatewayInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1RouterNewParamsExternalGatewayInfo) implementsCloudV1RouterNewParamsExternalGatewayInfoUnion() {
}

// Satisfied by [RouterExternalManualGwParam],
// [CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer],
// [CloudV1RouterNewParamsExternalGatewayInfo].
type CloudV1RouterNewParamsExternalGatewayInfoUnion interface {
	implementsCloudV1RouterNewParamsExternalGatewayInfoUnion()
}

type CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer struct {
	// Is SNAT enabled. Defaults to true.
	EnableSnat param.Field[bool] `json:"enable_snat"`
	// must be 'default'.
	Type param.Field[CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerType] `json:"type"`
}

func (r CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer) implementsCloudV1RouterNewParamsExternalGatewayInfoUnion() {
}

// must be 'default'.
type CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerType string

const (
	CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerTypeDefault CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerType = "default"
)

func (r CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerType) IsKnown() bool {
	switch r {
	case CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerTypeDefault:
		return true
	}
	return false
}

// must be 'manual'.
type CloudV1RouterNewParamsExternalGatewayInfoType string

const (
	CloudV1RouterNewParamsExternalGatewayInfoTypeManual  CloudV1RouterNewParamsExternalGatewayInfoType = "manual"
	CloudV1RouterNewParamsExternalGatewayInfoTypeDefault CloudV1RouterNewParamsExternalGatewayInfoType = "default"
)

func (r CloudV1RouterNewParamsExternalGatewayInfoType) IsKnown() bool {
	switch r {
	case CloudV1RouterNewParamsExternalGatewayInfoTypeManual, CloudV1RouterNewParamsExternalGatewayInfoTypeDefault:
		return true
	}
	return false
}

type CloudV1RouterNewParamsInterface struct {
	// id of the subnet to attach to.
	SubnetID param.Field[string] `json:"subnet_id,required" format:"uuid4"`
	// must be 'subnet'.
	Type param.Field[CloudV1RouterNewParamsInterfacesType] `json:"type"`
}

func (r CloudV1RouterNewParamsInterface) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// must be 'subnet'.
type CloudV1RouterNewParamsInterfacesType string

const (
	CloudV1RouterNewParamsInterfacesTypeSubnet CloudV1RouterNewParamsInterfacesType = "subnet"
)

func (r CloudV1RouterNewParamsInterfacesType) IsKnown() bool {
	switch r {
	case CloudV1RouterNewParamsInterfacesTypeSubnet:
		return true
	}
	return false
}

type CloudV1RouterUpdateParams struct {
	// New external gateway.
	ExternalGatewayInfo param.Field[RouterExternalManualGwParam] `json:"external_gateway_info"`
	// New name of router
	Name param.Field[string] `json:"name"`
	// List of custom routes.
	Routes param.Field[[]NeutronRouteParam] `json:"routes"`
}

func (r CloudV1RouterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RouterListParams struct {
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1RouterListParams]'s query parameters as
// `url.Values`.
func (r CloudV1RouterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1RouterAttachSubnetParams struct {
	SubnetID SubnetIDParam `json:"subnet_id,required"`
}

func (r CloudV1RouterAttachSubnetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubnetID)
}

type CloudV1RouterDetachSubnetParams struct {
	SubnetID SubnetIDParam `json:"subnet_id,required"`
}

func (r CloudV1RouterDetachSubnetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubnetID)
}
