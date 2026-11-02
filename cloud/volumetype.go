// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Volume types describe the operations each type supports and its QoS profile,
// optionally limited to one region.
//
// VolumeTypeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeTypeService] method instead.
type VolumeTypeService struct {
	Options []option.RequestOption
}

// NewVolumeTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVolumeTypeService(opts ...option.RequestOption) (r VolumeTypeService) {
	r = VolumeTypeService{}
	r.Options = opts
	return
}

// Returns the capabilities and the QoS profile of each volume type, sorted by
// name. Use the capabilities to find which operations a volume supports before you
// send a request. Set `region_id` to get the capabilities in that region. Without
// it, the capabilities ignore region limits.
func (r *VolumeTypeService) List(ctx context.Context, query VolumeTypeListParams, opts ...option.RequestOption) (res *VolumeTypeList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/volumes/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VolumeType struct {
	// Operations that a volume of this type supports.
	Capabilities VolumeTypeCapabilities `json:"capabilities" api:"required"`
	// Volume type name.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	Name VolumeTypeName `json:"name" api:"required"`
	// QoS profile of a volume type.
	//
	// The base limit of a volume is the integer part of its size in GiB multiplied by
	// the per-GiB rate. The minimum and maximum values limit the result. The burst
	// limit is the base limit multiplied by `burst_ratio`. The maximum value limits
	// the result.
	Qos VolumeTypeQos `json:"qos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		Name         respjson.Field
		Qos          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeType) RawJSON() string { return r.JSON.raw }
func (r *VolumeType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operations that a volume of this type supports.
type VolumeTypeCapabilities struct {
	// A volume of this type can be the boot volume of an instance. The instance flavor
	// can add more limits.
	AttachableToBoot bool `json:"attachable_to_boot" api:"required"`
	// A volume of this type can increase in size. Without `region_id`, the value
	// ignores region limits. A region can set it to false.
	Resizable bool `json:"resizable" api:"required"`
	// A volume of this type can change to a different retypable type, and the reverse.
	Retypable bool `json:"retypable" api:"required"`
	// A volume of this type can have snapshots.
	Snapshottable bool `json:"snapshottable" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachableToBoot respjson.Field
		Resizable        respjson.Field
		Retypable        respjson.Field
		Snapshottable    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTypeCapabilities) RawJSON() string { return r.JSON.raw }
func (r *VolumeTypeCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume type name.
type VolumeTypeName string

const (
	VolumeTypeNameCold          VolumeTypeName = "cold"
	VolumeTypeNameSsdHiiops     VolumeTypeName = "ssd_hiiops"
	VolumeTypeNameSsdLocal      VolumeTypeName = "ssd_local"
	VolumeTypeNameSsdLowlatency VolumeTypeName = "ssd_lowlatency"
	VolumeTypeNameStandard      VolumeTypeName = "standard"
	VolumeTypeNameUltra         VolumeTypeName = "ultra"
)

// QoS profile of a volume type.
//
// The base limit of a volume is the integer part of its size in GiB multiplied by
// the per-GiB rate. The minimum and maximum values limit the result. The burst
// limit is the base limit multiplied by `burst_ratio`. The maximum value limits
// the result.
type VolumeTypeQos struct {
	// Highest base and burst bandwidth limit in megabytes per second.
	BandwidthMaxMBPerSec int64 `json:"bandwidth_max_mb_per_sec" api:"required"`
	// Base bandwidth in megabytes per second for each GiB of volume size.
	BandwidthMBPerSecPerGib float64 `json:"bandwidth_mb_per_sec_per_gib" api:"required"`
	// Lowest base bandwidth limit in megabytes per second.
	BandwidthMinMBPerSec int64 `json:"bandwidth_min_mb_per_sec" api:"required"`
	// Burst limit divided by base limit.
	BurstRatio float64 `json:"burst_ratio" api:"required"`
	// Highest base and burst IOPS limit.
	IopsMax int64 `json:"iops_max" api:"required"`
	// Lowest base IOPS limit.
	IopsMin int64 `json:"iops_min" api:"required"`
	// Base IOPS for each GiB of volume size.
	IopsPerGib float64 `json:"iops_per_gib" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthMaxMBPerSec    respjson.Field
		BandwidthMBPerSecPerGib respjson.Field
		BandwidthMinMBPerSec    respjson.Field
		BurstRatio              respjson.Field
		IopsMax                 respjson.Field
		IopsMin                 respjson.Field
		IopsPerGib              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTypeQos) RawJSON() string { return r.JSON.raw }
func (r *VolumeTypeQos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeTypeList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []VolumeType `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTypeList) RawJSON() string { return r.JSON.raw }
func (r *VolumeTypeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeTypeListParams struct {
	// Region ID. If you set it, the list shows only the types that the region
	// supports. The capabilities then also include the limits of that region.
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VolumeTypeListParams]'s query parameters as `url.Values`.
func (r VolumeTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
