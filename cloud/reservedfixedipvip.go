// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// ReservedFixedIPVipService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReservedFixedIPVipService] method instead.
type ReservedFixedIPVipService struct {
	Options []option.RequestOption
}

// NewReservedFixedIPVipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReservedFixedIPVipService(opts ...option.RequestOption) (r ReservedFixedIPVipService) {
	r = ReservedFixedIPVipService{}
	r.Options = opts
	return
}

// List instance ports that are available for connecting to VIP
func (r *ReservedFixedIPVipService) ListCandidatePorts(ctx context.Context, portID string, query ReservedFixedIPVipListCandidatePortsParams, opts ...option.RequestOption) (res *CandidatePortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/available_devices", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List instance ports that share VIP
func (r *ReservedFixedIPVipService) ListConnectedPorts(ctx context.Context, portID string, query ReservedFixedIPVipListConnectedPortsParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace ports that share VIP
func (r *ReservedFixedIPVipService) ReplaceConnectedPorts(ctx context.Context, portID string, params ReservedFixedIPVipReplaceConnectedPortsParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Switch VIP status of reserved fixed IP
func (r *ReservedFixedIPVipService) Toggle(ctx context.Context, portID string, params ReservedFixedIPVipToggleParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Add ports that share VIP
func (r *ReservedFixedIPVipService) UpdateConnectedPorts(ctx context.Context, portID string, params ReservedFixedIPVipUpdateConnectedPortsParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// '#/components/schemas/VIPAttachCandidateSerializer'
// "$.components.schemas.VIPAttachCandidateSerializer"
type CandidatePort struct {
	// '#/components/schemas/VIPAttachCandidateSerializer/properties/instance_id'
	// "$.components.schemas.VIPAttachCandidateSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/VIPAttachCandidateSerializer/properties/instance_name'
	// "$.components.schemas.VIPAttachCandidateSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/VIPAttachCandidateSerializer/properties/ip_assignments'
	// "$.components.schemas.VIPAttachCandidateSerializer.properties.ip_assignments"
	IPAssignments []IPWithSubnet `json:"ip_assignments,required"`
	// '#/components/schemas/VIPAttachCandidateSerializer/properties/network'
	// "$.components.schemas.VIPAttachCandidateSerializer.properties.network"
	Network Network `json:"network,required"`
	// '#/components/schemas/VIPAttachCandidateSerializer/properties/port_id'
	// "$.components.schemas.VIPAttachCandidateSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InstanceID    resp.Field
		InstanceName  resp.Field
		IPAssignments resp.Field
		Network       resp.Field
		PortID        resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CandidatePort) RawJSON() string { return r.JSON.raw }
func (r *CandidatePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/VIPAttachCandidateSerializerList'
// "$.components.schemas.VIPAttachCandidateSerializerList"
type CandidatePortList struct {
	// '#/components/schemas/VIPAttachCandidateSerializerList/properties/count'
	// "$.components.schemas.VIPAttachCandidateSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/VIPAttachCandidateSerializerList/properties/results'
	// "$.components.schemas.VIPAttachCandidateSerializerList.properties.results"
	Results []CandidatePort `json:"results,required"`
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
func (r CandidatePortList) RawJSON() string { return r.JSON.raw }
func (r *CandidatePortList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ConnectedDevicesVIPSerializer'
// "$.components.schemas.ConnectedDevicesVIPSerializer"
type ConnectedPort struct {
	// '#/components/schemas/ConnectedDevicesVIPSerializer/properties/instance_id'
	// "$.components.schemas.ConnectedDevicesVIPSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/ConnectedDevicesVIPSerializer/properties/instance_name'
	// "$.components.schemas.ConnectedDevicesVIPSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/ConnectedDevicesVIPSerializer/properties/ip_assignments'
	// "$.components.schemas.ConnectedDevicesVIPSerializer.properties.ip_assignments"
	IPAssignments []IPWithSubnet `json:"ip_assignments,required"`
	// '#/components/schemas/ConnectedDevicesVIPSerializer/properties/network'
	// "$.components.schemas.ConnectedDevicesVIPSerializer.properties.network"
	Network Network `json:"network,required"`
	// '#/components/schemas/ConnectedDevicesVIPSerializer/properties/port_id'
	// "$.components.schemas.ConnectedDevicesVIPSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InstanceID    resp.Field
		InstanceName  resp.Field
		IPAssignments resp.Field
		Network       resp.Field
		PortID        resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectedPort) RawJSON() string { return r.JSON.raw }
func (r *ConnectedPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ConnectedDevicesVIPSerializerList'
// "$.components.schemas.ConnectedDevicesVIPSerializerList"
type ConnectedPortList struct {
	// '#/components/schemas/ConnectedDevicesVIPSerializerList/properties/count'
	// "$.components.schemas.ConnectedDevicesVIPSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ConnectedDevicesVIPSerializerList/properties/results'
	// "$.components.schemas.ConnectedDevicesVIPSerializerList.properties.results"
	Results []ConnectedPort `json:"results,required"`
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
func (r ConnectedPortList) RawJSON() string { return r.JSON.raw }
func (r *ConnectedPortList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/PortIpWithSubnetSerializer'
// "$.components.schemas.PortIpWithSubnetSerializer"
type IPWithSubnet struct {
	// '#/components/schemas/PortIpWithSubnetSerializer/properties/ip_address'
	// "$.components.schemas.PortIpWithSubnetSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpWithSubnetSerializer/properties/subnet'
	// "$.components.schemas.PortIpWithSubnetSerializer.properties.subnet"
	Subnet Subnet `json:"subnet,required"`
	// '#/components/schemas/PortIpWithSubnetSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpWithSubnetSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		Subnet      resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPWithSubnet) RawJSON() string { return r.JSON.raw }
func (r *IPWithSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPVipListCandidatePortsParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Favailable_devices/get/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/available_devices'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Favailable_devices/get/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/available_devices'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPVipListCandidatePortsParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type ReservedFixedIPVipListConnectedPortsParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/get/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/get/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPVipListConnectedPortsParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type ReservedFixedIPVipReplaceConnectedPortsParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/put/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/put/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].put.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PortIDsForVIPSerializer/properties/port_ids'
	// "$.components.schemas.PortIDsForVIPSerializer.properties.port_ids"
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPVipReplaceConnectedPortsParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r ReservedFixedIPVipReplaceConnectedPortsParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipReplaceConnectedPortsParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ReservedFixedIPVipToggleParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PatchReservedFixedIPSerializer/properties/is_vip'
	// "$.components.schemas.PatchReservedFixedIPSerializer.properties.is_vip"
	IsVip bool `json:"is_vip,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPVipToggleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ReservedFixedIPVipToggleParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipToggleParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ReservedFixedIPVipUpdateConnectedPortsParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freserved_fixed_ips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fconnected_devices/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/reserved_fixed_ips/{project_id}/{region_id}/{port_id}/connected_devices'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PortIDsForVIPSerializer/properties/port_ids'
	// "$.components.schemas.PortIDsForVIPSerializer.properties.port_ids"
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ReservedFixedIPVipUpdateConnectedPortsParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r ReservedFixedIPVipUpdateConnectedPortsParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipUpdateConnectedPortsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
