// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ReservedFixedIPVipConnectedPortService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReservedFixedIPVipConnectedPortService] method instead.
type ReservedFixedIPVipConnectedPortService struct {
	Options []option.RequestOption
}

// NewReservedFixedIPVipConnectedPortService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewReservedFixedIPVipConnectedPortService(opts ...option.RequestOption) (r ReservedFixedIPVipConnectedPortService) {
	r = ReservedFixedIPVipConnectedPortService{}
	r.Options = opts
	return
}

// List all instance ports that share a VIP.
func (r *ReservedFixedIPVipConnectedPortService) List(ctx context.Context, portID string, query ReservedFixedIPVipConnectedPortListParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add instance ports to share a VIP.
func (r *ReservedFixedIPVipConnectedPortService) Add(ctx context.Context, portID string, params ReservedFixedIPVipConnectedPortAddParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Replace the list of instance ports that share a VIP.
func (r *ReservedFixedIPVipConnectedPortService) Replace(ctx context.Context, portID string, params ReservedFixedIPVipConnectedPortReplaceParams, opts ...option.RequestOption) (res *ConnectedPortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", params.ProjectID.Value, params.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstanceID    respjson.Field
		InstanceName  respjson.Field
		IPAssignments respjson.Field
		Network       respjson.Field
		PortID        respjson.Field
		ExtraFields   map[string]respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectedPortList) RawJSON() string { return r.JSON.raw }
func (r *ConnectedPortList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPVipConnectedPortListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type ReservedFixedIPVipConnectedPortAddParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of port IDs that will share one VIP
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r ReservedFixedIPVipConnectedPortAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipConnectedPortAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPVipConnectedPortAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPVipConnectedPortReplaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of port IDs that will share one VIP
	PortIDs []string `json:"port_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r ReservedFixedIPVipConnectedPortReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ReservedFixedIPVipConnectedPortReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedFixedIPVipConnectedPortReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
