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

// Add ports that share VIP
func (r *ReservedFixedIPVipService) UpdateConnectedPorts(ctx context.Context, portID string, params ReservedFixedIPVipUpdateConnectedPortsParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type CandidatePort struct {
	// ID of the instance that owns the port
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Name of the instance that owns the port
	InstanceName string `json:"instance_name,required"`
	// IP addresses assigned to this port
	IPAssignments []IPWithSubnet `json:"ip_assignments,required"`
	// Network details
	Network Network `json:"network,required"`
	// Port ID that shares VIP
	PortID string `json:"port_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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

type CandidatePortList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CandidatePort `json:"results,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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

type ConnectedPort struct {
	// ID of the instance that owns the port
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Name of the instance that owns the port
	InstanceName string `json:"instance_name,required"`
	// IP addresses assigned to this port
	IPAssignments []IPWithSubnet `json:"ip_assignments,required"`
	// Network details
	Network Network `json:"network,required"`
	// Port ID that shares VIP
	PortID string `json:"port_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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

type ConnectedPortList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ConnectedPort `json:"results,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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

type IPWithSubnet struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Subnet details
	Subnet Subnet `json:"subnet,required"`
	// ID of the subnet that allocated the IP
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type ReservedFixedIPVipListConnectedPortsParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type ReservedFixedIPVipReplaceConnectedPortsParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of port IDs that will share one VIP
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r ReservedFixedIPVipReplaceConnectedPortsParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipReplaceConnectedPortsParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ReservedFixedIPVipToggleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// If reserved fixed IP should be a VIP
	IsVip bool `json:"is_vip,required"`
	paramObj
}

func (r ReservedFixedIPVipToggleParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipToggleParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ReservedFixedIPVipUpdateConnectedPortsParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of port IDs that will share one VIP
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r ReservedFixedIPVipUpdateConnectedPortsParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipUpdateConnectedPortsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
