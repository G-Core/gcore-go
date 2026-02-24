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

// ReservedFixedIPVipCandidatePortService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReservedFixedIPVipCandidatePortService] method instead.
type ReservedFixedIPVipCandidatePortService struct {
	Options []option.RequestOption
}

// NewReservedFixedIPVipCandidatePortService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewReservedFixedIPVipCandidatePortService(opts ...option.RequestOption) (r ReservedFixedIPVipCandidatePortService) {
	r = ReservedFixedIPVipCandidatePortService{}
	r.Options = opts
	return
}

// List all instance ports that are available for connecting to a VIP.
func (r *ReservedFixedIPVipCandidatePortService) List(ctx context.Context, portID string, query ReservedFixedIPVipCandidatePortListParams, opts ...option.RequestOption) (res *CandidatePortList, err error) {
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
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/available_devices", query.ProjectID.Value, query.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CandidatePort struct {
	// ID of the instance that owns the port
	InstanceID string `json:"instance_id" api:"required" format:"uuid4"`
	// Name of the instance that owns the port
	InstanceName string `json:"instance_name" api:"required"`
	// IP addresses assigned to this port
	IPAssignments []IPWithSubnet `json:"ip_assignments" api:"required"`
	// Network details
	Network Network `json:"network" api:"required"`
	// Port ID that shares VIP
	PortID string `json:"port_id" api:"required" format:"uuid4"`
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
func (r CandidatePort) RawJSON() string { return r.JSON.raw }
func (r *CandidatePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CandidatePortList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []CandidatePort `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CandidatePortList) RawJSON() string { return r.JSON.raw }
func (r *CandidatePortList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReservedFixedIPVipCandidatePortListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
