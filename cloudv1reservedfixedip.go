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

// CloudV1ReservedFixedIPService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ReservedFixedIPService] method instead.
type CloudV1ReservedFixedIPService struct {
	Options          []option.RequestOption
	ConnectedDevices *CloudV1ReservedFixedIPConnectedDeviceService
}

// NewCloudV1ReservedFixedIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ReservedFixedIPService(opts ...option.RequestOption) (r *CloudV1ReservedFixedIPService) {
	r = &CloudV1ReservedFixedIPService{}
	r.Options = opts
	r.ConnectedDevices = NewCloudV1ReservedFixedIPConnectedDeviceService(opts...)
	return
}

// Create reserved fixed IP
func (r *CloudV1ReservedFixedIPService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1ReservedFixedIPNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get reserved fixed IP
func (r *CloudV1ReservedFixedIPService) Get(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List reserved fixed IPs
func (r *CloudV1ReservedFixedIPService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1ReservedFixedIPListParams, opts ...option.RequestOption) (res *CloudV1ReservedFixedIPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete reserved fixed ip
func (r *CloudV1ReservedFixedIPService) Delete(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List instance ports that are available for connecting to VIP
func (r *CloudV1ReservedFixedIPService) ListAvailableDevices(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *CloudV1ReservedFixedIPListAvailableDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/available_devices", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Switch VIP status of reserved fixed IP
func (r *CloudV1ReservedFixedIPService) SwitchVipStatus(ctx context.Context, projectID int64, regionID int64, portID string, body CloudV1ReservedFixedIPSwitchVipStatusParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type PortIPWithSubnet struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Subnet details
	Subnet Subnet `json:"subnet,required"`
	// ID of the subnet that allocated the IP
	SubnetID string               `json:"subnet_id,required" format:"uuid4"`
	JSON     portIPWithSubnetJSON `json:"-"`
}

// portIPWithSubnetJSON contains the JSON metadata for the struct
// [PortIPWithSubnet]
type portIPWithSubnetJSON struct {
	IPAddress   apijson.Field
	Subnet      apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PortIPWithSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portIPWithSubnetJSON) RawJSON() string {
	return r.raw
}

type ReservedFixedIP struct {
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs,required"`
	// Reserved fixed IP attachment entities
	Attachments []ReservedFixedIPAttachment `json:"attachments,required"`
	// Datetime when the reserved fixed IP was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If reserved fixed IP belongs to a public network
	IsExternal bool `json:"is_external,required"`
	// If reserved fixed IP is a VIP
	IsVip bool `json:"is_vip,required"`
	// Reserved fixed IP name
	Name string `json:"name,required"`
	// Network details
	Network Network `json:"network,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of the port underlying the reserved fixed IP
	PortID string `json:"port_id,required" format:"uuid4"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Reserved fixed IP status with resource type and ID it is attached to
	Reservation ReservedFixedIPReservation `json:"reservation,required"`
	// Underlying port status
	Status string `json:"status,required"`
	// Datetime when the reserved fixed IP was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// IPv4 address of the reserved fixed IP
	FixedIPAddress string `json:"fixed_ip_address,nullable" format:"ipv4"`
	// IPv6 address of the reserved fixed IP
	FixedIpv6Address string `json:"fixed_ipv6_address,nullable" format:"ipv6"`
	// Project ID
	ProjectID int64 `json:"project_id,nullable"`
	// ID of the subnet that owns the IP address
	SubnetID string `json:"subnet_id,nullable" format:"uuid4"`
	// ID of the subnet that owns the IPv6 address
	SubnetV6ID string `json:"subnet_v6_id,nullable" format:"uuid4"`
	// Active task. If null, action has been performed immediately in the request
	// itself.
	TaskID string              `json:"task_id,nullable" format:"uuid4"`
	JSON   reservedFixedIPJSON `json:"-"`
}

// reservedFixedIPJSON contains the JSON metadata for the struct [ReservedFixedIP]
type reservedFixedIPJSON struct {
	AllowedAddressPairs apijson.Field
	Attachments         apijson.Field
	CreatedAt           apijson.Field
	IsExternal          apijson.Field
	IsVip               apijson.Field
	Name                apijson.Field
	Network             apijson.Field
	NetworkID           apijson.Field
	PortID              apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Reservation         apijson.Field
	Status              apijson.Field
	UpdatedAt           apijson.Field
	CreatorTaskID       apijson.Field
	FixedIPAddress      apijson.Field
	FixedIpv6Address    apijson.Field
	ProjectID           apijson.Field
	SubnetID            apijson.Field
	SubnetV6ID          apijson.Field
	TaskID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ReservedFixedIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reservedFixedIPJSON) RawJSON() string {
	return r.raw
}

type ReservedFixedIPAttachment struct {
	// Resource ID
	ResourceID string `json:"resource_id,nullable"`
	// Resource type
	ResourceType string                        `json:"resource_type,nullable"`
	JSON         reservedFixedIPAttachmentJSON `json:"-"`
}

// reservedFixedIPAttachmentJSON contains the JSON metadata for the struct
// [ReservedFixedIPAttachment]
type reservedFixedIPAttachmentJSON struct {
	ResourceID   apijson.Field
	ResourceType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ReservedFixedIPAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reservedFixedIPAttachmentJSON) RawJSON() string {
	return r.raw
}

// Reserved fixed IP status with resource type and ID it is attached to
type ReservedFixedIPReservation struct {
	// ID of the instance or load balancer the IP is attached to
	ResourceID string `json:"resource_id,nullable" format:"uuid4"`
	// Resource type of the resource the IP is attached to
	ResourceType string `json:"resource_type,nullable"`
	// IP reservation status
	Status string                         `json:"status,nullable"`
	JSON   reservedFixedIPReservationJSON `json:"-"`
}

// reservedFixedIPReservationJSON contains the JSON metadata for the struct
// [ReservedFixedIPReservation]
type reservedFixedIPReservationJSON struct {
	ResourceID   apijson.Field
	ResourceType apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ReservedFixedIPReservation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reservedFixedIPReservationJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservedFixedIPListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ReservedFixedIP                      `json:"results,required"`
	JSON    cloudV1ReservedFixedIPListResponseJSON `json:"-"`
}

// cloudV1ReservedFixedIPListResponseJSON contains the JSON metadata for the struct
// [CloudV1ReservedFixedIPListResponse]
type cloudV1ReservedFixedIPListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ReservedFixedIPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ReservedFixedIPListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservedFixedIPListAvailableDevicesResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1ReservedFixedIPListAvailableDevicesResponseResult `json:"results,required"`
	JSON    cloudV1ReservedFixedIPListAvailableDevicesResponseJSON     `json:"-"`
}

// cloudV1ReservedFixedIPListAvailableDevicesResponseJSON contains the JSON
// metadata for the struct [CloudV1ReservedFixedIPListAvailableDevicesResponse]
type cloudV1ReservedFixedIPListAvailableDevicesResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ReservedFixedIPListAvailableDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ReservedFixedIPListAvailableDevicesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservedFixedIPListAvailableDevicesResponseResult struct {
	// ID of the instance that owns the port
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Name of the instance that owns the port
	InstanceName string `json:"instance_name,required"`
	// IP addresses assigned to this port
	IPAssignments []PortIPWithSubnet `json:"ip_assignments,required"`
	// Network details
	Network Network `json:"network,required"`
	// Port ID that shares VIP
	PortID string                                                       `json:"port_id,required" format:"uuid4"`
	JSON   cloudV1ReservedFixedIPListAvailableDevicesResponseResultJSON `json:"-"`
}

// cloudV1ReservedFixedIPListAvailableDevicesResponseResultJSON contains the JSON
// metadata for the struct
// [CloudV1ReservedFixedIPListAvailableDevicesResponseResult]
type cloudV1ReservedFixedIPListAvailableDevicesResponseResultJSON struct {
	InstanceID    apijson.Field
	InstanceName  apijson.Field
	IPAssignments apijson.Field
	Network       apijson.Field
	PortID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1ReservedFixedIPListAvailableDevicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ReservedFixedIPListAvailableDevicesResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservedFixedIPNewParams struct {
	CreateReservedFixedIP CreateReservedFixedIPUnionParam `json:"create_reserved_fixed_ip"`
}

func (r CloudV1ReservedFixedIPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateReservedFixedIP)
}

type CloudV1ReservedFixedIPListParams struct {
	// Set to true if the response should only list IP addresses that are not attached
	// to any instance
	AvailableOnly param.Field[bool] `query:"available_only"`
	// Filter IPs by device ID it is attached to
	DeviceID param.Field[string] `query:"device_id"`
	// Set to true if the response should only list public IP addresses
	ExternalOnly param.Field[bool] `query:"external_only"`
	// Set to true if the response should only list private IP addresses
	InternalOnly param.Field[bool] `query:"internal_only"`
	// An IPv4 address to filter results by. Regular expression allowed
	IPAddress param.Field[string] `query:"ip_address"`
	// Limit the number of returned IPs
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Ordering reserved fixed IP list result by name, status, updated_at, created_at
	// or fixed_ip_address fields of the reserved fixed IP and directions (status.asc),
	// default is "fixed_ip_address.asc"
	OrderBy param.Field[string] `query:"order_by"`
	// Set to true if the response should only list VIPs
	VipOnly param.Field[bool] `query:"vip_only"`
}

// URLQuery serializes [CloudV1ReservedFixedIPListParams]'s query parameters as
// `url.Values`.
func (r CloudV1ReservedFixedIPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1ReservedFixedIPSwitchVipStatusParams struct {
	// If reserved fixed IP should be a VIP
	IsVip param.Field[bool] `json:"is_vip,required"`
}

func (r CloudV1ReservedFixedIPSwitchVipStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
