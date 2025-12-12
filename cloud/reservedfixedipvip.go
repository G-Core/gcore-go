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

// ReservedFixedIPVipService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReservedFixedIPVipService] method instead.
type ReservedFixedIPVipService struct {
	Options        []option.RequestOption
	CandidatePorts ReservedFixedIPVipCandidatePortService
	ConnectedPorts ReservedFixedIPVipConnectedPortService
}

// NewReservedFixedIPVipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReservedFixedIPVipService(opts ...option.RequestOption) (r ReservedFixedIPVipService) {
	r = ReservedFixedIPVipService{}
	r.Options = opts
	r.CandidatePorts = NewReservedFixedIPVipCandidatePortService(opts...)
	r.ConnectedPorts = NewReservedFixedIPVipConnectedPortService(opts...)
	return
}

// Update the VIP status of a reserved fixed IP.
func (r *ReservedFixedIPVipService) Toggle(ctx context.Context, portID string, params ReservedFixedIPVipToggleParams, opts ...option.RequestOption) (res *ReservedFixedIP, err error) {
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

type IPWithSubnet struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Subnet details
	Subnet Subnet `json:"subnet,required"`
	// ID of the subnet that allocated the IP
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		Subnet      respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPWithSubnet) RawJSON() string { return r.JSON.raw }
func (r *IPWithSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
func (r *ReservedFixedIPVipToggleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
