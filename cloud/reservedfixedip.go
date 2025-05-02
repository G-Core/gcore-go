// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
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
	"github.com/stainless-sdks/gcore-go/shared/constant"
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
}

// NewReservedFixedIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReservedFixedIPService(opts ...option.RequestOption) (r ReservedFixedIPService) {
	r = ReservedFixedIPService{}
	r.Options = opts
	r.Vip = NewReservedFixedIPVipService(opts...)
	return
}

// Create reserved fixed IP
func (r *ReservedFixedIPService) New(ctx context.Context, params ReservedFixedIPNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List reserved fixed IPs
func (r *ReservedFixedIPService) List(ctx context.Context, params ReservedFixedIPListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ReservedFixedIP], err error) {
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

// List reserved fixed IPs
func (r *ReservedFixedIPService) ListAutoPaging(ctx context.Context, params ReservedFixedIPListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ReservedFixedIP] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete reserved fixed ip
func (r *ReservedFixedIPService) Delete(ctx context.Context, portID string, body ReservedFixedIPDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get reserved fixed IP
func (r *ReservedFixedIPService) Get(ctx context.Context, portID string, query ReservedFixedIPGetParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/ReservedFixedIPSerializer'
// "$.components.schemas.ReservedFixedIPSerializer"
type ReservedFixedIP struct {
	// '#/components/schemas/ReservedFixedIPSerializer/properties/allowed_address_pairs'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.allowed_address_pairs"
	AllowedAddressPairs []ReservedFixedIPAllowedAddressPair `json:"allowed_address_pairs,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/attachments'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.attachments"
	Attachments []ReservedFixedIPAttachment `json:"attachments,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/created_at'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/is_external'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.is_external"
	IsExternal bool `json:"is_external,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/is_vip'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.is_vip"
	IsVip bool `json:"is_vip,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/name'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/network'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.network"
	Network Network `json:"network,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/network_id'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/port_id'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/region'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/region_id'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/reservation'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.reservation"
	Reservation ReservedFixedIPReservation `json:"reservation,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/status'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/updated_at'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/fixed_ip_address/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.fixed_ip_address.anyOf[0]"
	FixedIPAddress string `json:"fixed_ip_address,nullable" format:"ipv4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/fixed_ipv6_address/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.fixed_ipv6_address.anyOf[0]"
	FixedIpv6Address string `json:"fixed_ipv6_address,nullable" format:"ipv6"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.subnet_id.anyOf[0]"
	SubnetID string `json:"subnet_id,nullable" format:"uuid4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/subnet_v6_id/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.subnet_v6_id.anyOf[0]"
	SubnetV6ID string `json:"subnet_v6_id,nullable" format:"uuid4"`
	// '#/components/schemas/ReservedFixedIPSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.ReservedFixedIPSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AllowedAddressPairs resp.Field
		Attachments         resp.Field
		CreatedAt           resp.Field
		IsExternal          resp.Field
		IsVip               resp.Field
		Name                resp.Field
		Network             resp.Field
		NetworkID           resp.Field
		PortID              resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Reservation         resp.Field
		Status              resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		FixedIPAddress      resp.Field
		FixedIpv6Address    resp.Field
		ProjectID           resp.Field
		SubnetID            resp.Field
		SubnetV6ID          resp.Field
		TaskID              resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIP) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ReservedFixedIPSerializer/properties/allowed_address_pairs/items'
// "$.components.schemas.ReservedFixedIPSerializer.properties.allowed_address_pairs.items"
type ReservedFixedIPAllowedAddressPair struct {
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		MacAddress  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIPAllowedAddressPair) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIPAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ReservedFixedIPSerializer/properties/attachments/items'
// "$.components.schemas.ReservedFixedIPSerializer.properties.attachments.items"
type ReservedFixedIPAttachment struct {
	// '#/components/schemas/AttachmentSerializer/properties/resource_id/anyOf/0'
	// "$.components.schemas.AttachmentSerializer.properties.resource_id.anyOf[0]"
	ResourceID string `json:"resource_id,nullable"`
	// '#/components/schemas/AttachmentSerializer/properties/resource_type/anyOf/0'
	// "$.components.schemas.AttachmentSerializer.properties.resource_type.anyOf[0]"
	ResourceType string `json:"resource_type,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ResourceID   resp.Field
		ResourceType resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIPAttachment) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIPAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ReservedFixedIPSerializer/properties/reservation'
// "$.components.schemas.ReservedFixedIPSerializer.properties.reservation"
type ReservedFixedIPReservation struct {
	// '#/components/schemas/ReservationSerializer/properties/resource_id/anyOf/0'
	// "$.components.schemas.ReservationSerializer.properties.resource_id.anyOf[0]"
	ResourceID string `json:"resource_id,nullable" format:"uuid4"`
	// '#/components/schemas/ReservationSerializer/properties/resource_type/anyOf/0'
	// "$.components.schemas.ReservationSerializer.properties.resource_type.anyOf[0]"
	ResourceType string `json:"resource_type,nullable"`
	// '#/components/schemas/ReservationSerializer/properties/status/anyOf/0'
	// "$.components.schemas.ReservationSerializer.properties.status.anyOf[0]"
	Status string `json:"status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ResourceID   resp.Field
		ResourceType resp.Field
		Status       resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedFixedIPReservation) RawJSON() string { return r.JSON.raw }
func (r *ReservedFixedIPReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/0'
	// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[0]"
	OfExternal *ReservedFixedIPNewParamsBodyExternal `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/1'
	// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[1]"
	OfSubnet *ReservedFixedIPNewParamsBodySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/2'
	// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[2]"
	OfAnySubnet *ReservedFixedIPNewParamsBodyAnySubnet `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/3'
	// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[3]"
	OfIPAddress *ReservedFixedIPNewParamsBodyIPAddress `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/4'
	// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[4]"
	OfPort *ReservedFixedIPNewParamsBodyPort `json:",inline"`

	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (u ReservedFixedIPNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ReservedFixedIPNewParams](u.OfExternal,
		u.OfSubnet,
		u.OfAnySubnet,
		u.OfIPAddress,
		u.OfPort)
}

// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/0'
// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[0]"
//
// The property Type is required.
type ReservedFixedIPNewParamsBodyExternal struct {
	// '#/components/schemas/NewReservedFixedIpExternalSerializer/properties/is_vip'
	// "$.components.schemas.NewReservedFixedIpExternalSerializer.properties.is_vip"
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// '#/components/schemas/NewReservedFixedIpExternalSerializer/properties/ip_family/anyOf/0'
	// "$.components.schemas.NewReservedFixedIpExternalSerializer.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/NewReservedFixedIpExternalSerializer/properties/type'
	// "$.components.schemas.NewReservedFixedIpExternalSerializer.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParamsBodyExternal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ReservedFixedIPNewParamsBodyExternal) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyExternal
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/1'
// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[1]"
//
// The properties SubnetID, Type are required.
type ReservedFixedIPNewParamsBodySubnet struct {
	// '#/components/schemas/NewReservedFixedIpSpecificSubnetSerializer/properties/subnet_id'
	// "$.components.schemas.NewReservedFixedIpSpecificSubnetSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/NewReservedFixedIpSpecificSubnetSerializer/properties/is_vip'
	// "$.components.schemas.NewReservedFixedIpSpecificSubnetSerializer.properties.is_vip"
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// '#/components/schemas/NewReservedFixedIpSpecificSubnetSerializer/properties/type'
	// "$.components.schemas.NewReservedFixedIpSpecificSubnetSerializer.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParamsBodySubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ReservedFixedIPNewParamsBodySubnet) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/2'
// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[2]"
//
// The properties NetworkID, Type are required.
type ReservedFixedIPNewParamsBodyAnySubnet struct {
	// '#/components/schemas/NewReservedFixedIpAnySubnetSerializer/properties/network_id'
	// "$.components.schemas.NewReservedFixedIpAnySubnetSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/NewReservedFixedIpAnySubnetSerializer/properties/is_vip'
	// "$.components.schemas.NewReservedFixedIpAnySubnetSerializer.properties.is_vip"
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// '#/components/schemas/NewReservedFixedIpAnySubnetSerializer/properties/ip_family/anyOf/0'
	// "$.components.schemas.NewReservedFixedIpAnySubnetSerializer.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/NewReservedFixedIpAnySubnetSerializer/properties/type'
	// "$.components.schemas.NewReservedFixedIpAnySubnetSerializer.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParamsBodyAnySubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ReservedFixedIPNewParamsBodyAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/3'
// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[3]"
//
// The properties IPAddress, NetworkID, Type are required.
type ReservedFixedIPNewParamsBodyIPAddress struct {
	// '#/components/schemas/NewReservedFixedIpSpecificIpAddressSerializer/properties/ip_address'
	// "$.components.schemas.NewReservedFixedIpSpecificIpAddressSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/NewReservedFixedIpSpecificIpAddressSerializer/properties/network_id'
	// "$.components.schemas.NewReservedFixedIpSpecificIpAddressSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/NewReservedFixedIpSpecificIpAddressSerializer/properties/is_vip'
	// "$.components.schemas.NewReservedFixedIpSpecificIpAddressSerializer.properties.is_vip"
	IsVip param.Opt[bool] `json:"is_vip,omitzero"`
	// '#/components/schemas/NewReservedFixedIpSpecificIpAddressSerializer/properties/type'
	// "$.components.schemas.NewReservedFixedIpSpecificIpAddressSerializer.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "ip_address".
	Type constant.IPAddress `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParamsBodyIPAddress) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ReservedFixedIPNewParamsBodyIPAddress) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyIPAddress
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateReservedFixedIpSerializer/anyOf/4'
// "$.components.schemas.CreateReservedFixedIpSerializer.anyOf[4]"
//
// The properties PortID, Type are required.
type ReservedFixedIPNewParamsBodyPort struct {
	// '#/components/schemas/NewReservedFixedIpSpecificPortSerializer/properties/port_id'
	// "$.components.schemas.NewReservedFixedIpSpecificPortSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/NewReservedFixedIpSpecificPortSerializer/properties/type'
	// "$.components.schemas.NewReservedFixedIpSpecificPortSerializer.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "port".
	Type constant.Port `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPNewParamsBodyPort) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ReservedFixedIPNewParamsBodyPort) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPNewParamsBodyPort
	return param.MarshalObject(r, (*shadow)(&r))
}

type ReservedFixedIPListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[2]"
	AvailableOnly param.Opt[bool] `query:"available_only,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[3]"
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[4]"
	ExternalOnly param.Opt[bool] `query:"external_only,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[5]"
	InternalOnly param.Opt[bool] `query:"internal_only,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[6]"
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[7]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/8'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[8]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/9'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[9]"
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/10'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}'].get.parameters[10]"
	VipOnly param.Opt[bool] `query:"vip_only,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ReservedFixedIPListParams]'s query parameters as
// `url.Values`.
func (r ReservedFixedIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ReservedFixedIPDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type ReservedFixedIPGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
