// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// ReservedFixedIPService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReservedFixedIPService] method instead.
type ReservedFixedIPService struct {
	Options []option.RequestOption
	Vip     ReservedFixedIPVipService
	task    TaskService
}

// NewReservedFixedIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReservedFixedIPService(opts ...option.RequestOption) (r ReservedFixedIPService) {
	r = ReservedFixedIPService{}
	r.Options = opts
	r.Vip = NewReservedFixedIPVipService(opts...)
	r.task = NewTaskService(opts...)
	return
}

// Create a new reserved fixed IP with the specified configuration.
func (r *ReservedFixedIPService) New(ctx context.Context, params ReservedFixedIPNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create reserved fixed IP and poll for the result
func (r *ReservedFixedIPService) NewAndPoll(ctx context.Context, params ReservedFixedIPNewParams, opts ...option.RequestOption) (v *ReservedFixedIP, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams ReservedFixedIPGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.task.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Ports) != 1 {
		return nil, errors.New("expected exactly one port to be created in a task")
	}
	resourceID := task.CreatedResources.Ports[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Update the VIP status of a reserved fixed IP.
func (r *ReservedFixedIPService) Update(ctx context.Context, portID string, params ReservedFixedIPUpdateParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all reserved fixed IPs in the specified project and region.
func (r *ReservedFixedIPService) List(ctx context.Context, params ReservedFixedIPListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ReservedFixedIP], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List all reserved fixed IPs in the specified project and region.
func (r *ReservedFixedIPService) ListAutoPaging(ctx context.Context, params ReservedFixedIPListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ReservedFixedIP] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific reserved fixed IP and all its associated resources.
func (r *ReservedFixedIPService) Delete(ctx context.Context, portID string, body ReservedFixedIPDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a reserved fixed IP and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *ReservedFixedIPService) DeleteAndPoll(ctx context.Context, portID string, body ReservedFixedIPDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, portID, body, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.task.Poll(ctx, taskID, opts...)
	return err
}

// Get detailed information about a specific reserved fixed IP.
func (r *ReservedFixedIPService) Get(ctx context.Context, portID string, query ReservedFixedIPGetParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		Attachments         respjson.Field
		CreatedAt           respjson.Field
		IsExternal          respjson.Field
		IsVip               respjson.Field
		Name                respjson.Field
		Network             respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		Reservation         respjson.Field
		Status              respjson.Field
		UpdatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		FixedIPAddress      respjson.Field
		FixedIpv6Address    respjson.Field
		ProjectID           respjson.Field
		SubnetID            respjson.Field
		SubnetV6ID          respjson.Field
		TaskID              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIP) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPAttachment struct {
	// Resource ID
	ResourceID string `json:"resource_id,nullable"`
	// Resource type
	ResourceType string `json:"resource_type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIPAttachment) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIPAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reserved fixed IP status with resource type and ID it is attached to
type ReservedFixedIPReservation struct {
	// ID of the instance or load balancer the IP is attached to
	ResourceID string `json:"resource_id,nullable" format:"uuid4"`
	// Resource type of the resource the IP is attached to
	ResourceType string `json:"resource_type,nullable"`
	// IP reservation status
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceType respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIPReservation) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIPReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExternal *ReservedFixedIPNewParamsBodyExternal `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSubnet *ReservedFixedIPNewParamsBodySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAnySubnet *ReservedFixedIPNewParamsBodyAnySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfIPAddress *ReservedFixedIPNewParamsBodyIPAddress `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPort *ReservedFixedIPNewParamsBodyPort `json:",inline"`

	paramObj
}

func (u ReservedFixedIPNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal,
		u.OfSubnet,
		u.OfAnySubnet,
		u.OfIPAddress,
		u.OfPort)
}
func (r *ReservedFixedIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ReservedFixedIPNewParamsBodyExternal struct {
	// If reserved fixed IP is a VIP
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Must be 'external'
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

func (r ReservedFixedIPNewParamsBodyExternal) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPNewParamsBodyExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SubnetID, Type are required.
type ReservedFixedIPNewParamsBodySubnet struct {
	// Reserved fixed IP will be allocated in this subnet
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// If reserved fixed IP is a VIP
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// Must be 'subnet'.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

func (r ReservedFixedIPNewParamsBodySubnet) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPNewParamsBodySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type ReservedFixedIPNewParamsBodyAnySubnet struct {
	// Reserved fixed IP will be allocated in a subnet of this network
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// If reserved fixed IP is a VIP
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Must be '`any_subnet`'.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

func (r ReservedFixedIPNewParamsBodyAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPNewParamsBodyAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IPAddress, NetworkID, Type are required.
type ReservedFixedIPNewParamsBodyIPAddress struct {
	// Reserved fixed IP will be allocated the given IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Reserved fixed IP will be allocated in a subnet of this network
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// If reserved fixed IP is a VIP
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// Must be '`ip_address`'.
	//
	// This field can be elided, and will marshal its zero value as "ip_address".
	Type constant.IPAddress `json:"type,required"`
	paramObj
}

func (r ReservedFixedIPNewParamsBodyIPAddress) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyIPAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPNewParamsBodyIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortID, Type are required.
type ReservedFixedIPNewParamsBodyPort struct {
	// Port ID to make a reserved fixed IP (for example, `vip_port_id` of the Load
	// Balancer entity).
	PortID string `json:"port_id,required" format:"uuid4"`
	// Must be 'port'.
	//
	// This field can be elided, and will marshal its zero value as "port".
	Type constant.Port `json:"type,required"`
	paramObj
}

func (r ReservedFixedIPNewParamsBodyPort) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyPort
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPNewParamsBodyPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// If reserved fixed IP should be a VIP
	IsVip bool `json:"is_vip,required"`
	paramObj
}

func (r ReservedFixedIPUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to true if the response should only list IP addresses that are not attached
	// to any instance
	AvailableOnly param.Opt[bool] `query:"available_only,omitzero" json:"-"`
	// Filter IPs by device ID it is attached to
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	// Set to true if the response should only list public IP addresses
	ExternalOnly param.Opt[bool] `query:"external_only,omitzero" json:"-"`
	// Set to true if the response should only list private IP addresses
	InternalOnly param.Opt[bool] `query:"internal_only,omitzero" json:"-"`
	// An IPv4 address to filter results by. Regular expression allowed
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// Limit the number of returned IPs
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Ordering reserved fixed IP list result by name, status, `updated_at`,
	// `created_at` or `fixed_ip_address` fields and directions (status.asc), default
	// is "`fixed_ip_address`.asc"
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Set to true if the response should only list VIPs
	VipOnly param.Opt[bool] `query:"vip_only,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReservedFixedIPListParams]'s query parameters as
// `url.Values`.
func (r ReservedFixedIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ReservedFixedIPDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type ReservedFixedIPGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
