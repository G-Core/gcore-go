// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
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

// InstanceMetricService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceMetricService] method instead.
type InstanceMetricService struct {
	Options []option.RequestOption
}

// NewInstanceMetricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceMetricService(opts ...option.RequestOption) (r InstanceMetricService) {
	r = InstanceMetricService{}
	r.Options = opts
	return
}

// Get instance metrics, including cpu, memory, network and disk metrics
func (r *InstanceMetricService) List(ctx context.Context, instanceID string, params InstanceMetricListParams, opts ...option.RequestOption) (res *MetricsList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metrics", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Instance metrics item
type Metrics struct {
	// CPU utilization, % (max 100% for multi-core). The value is null if there is no
	// data for this time.
	CPUUtil float64 `json:"cpu_util" api:"required"`
	// Disks metrics for each of the disks attached. The value is null if there is no
	// disk data for this time.
	Disks []MetricsDisk `json:"disks" api:"required"`
	// RAM utilization, %. The value is null if there is no data for this time.
	MemoryUtil float64 `json:"memory_util" api:"required"`
	// Network out, bits per second (the 'Bps' in the field name is historical; the
	// value is bits, not bytes). The value is null if there is no data for this time.
	NetworkBpsEgress float64 `json:"network_Bps_egress" api:"required"`
	// Network in, bits per second (the 'Bps' in the field name is historical; the
	// value is bits, not bytes). The value is null if there is no data for this time.
	NetworkBpsIngress float64 `json:"network_Bps_ingress" api:"required"`
	// Network out, packets per second. The value is null if there is no data for this
	// time.
	NetworkPpsEgress float64 `json:"network_pps_egress" api:"required"`
	// Network in, packets per second. The value is null if there is no data for this
	// time.
	NetworkPpsIngress float64 `json:"network_pps_ingress" api:"required"`
	// Timestamp
	Time string `json:"time" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPUUtil           respjson.Field
		Disks             respjson.Field
		MemoryUtil        respjson.Field
		NetworkBpsEgress  respjson.Field
		NetworkBpsIngress respjson.Field
		NetworkPpsEgress  respjson.Field
		NetworkPpsIngress respjson.Field
		Time              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metrics) RawJSON() string { return r.JSON.raw }
func (r *Metrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Disk metrics item
type MetricsDisk struct {
	// Disk read, Bytes per second. The value is null if there is no data for this
	// disk.
	DiskBpsRead float64 `json:"disk_Bps_read" api:"required"`
	// Disk write, Bytes per second. The value is null if there is no data for this
	// disk.
	DiskBpsWrite float64 `json:"disk_Bps_write" api:"required"`
	// Disk read, iops. The value is null if there is no data for this disk.
	DiskIopsRead float64 `json:"disk_iops_read" api:"required"`
	// Disk write, iops. The value is null if there is no data for this disk.
	DiskIopsWrite float64 `json:"disk_iops_write" api:"required"`
	// Disk attached slot name. The value is null if the name is unknown.
	DiskName string `json:"disk_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiskBpsRead   respjson.Field
		DiskBpsWrite  respjson.Field
		DiskIopsRead  respjson.Field
		DiskIopsWrite respjson.Field
		DiskName      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsDisk) RawJSON() string { return r.JSON.raw }
func (r *MetricsDisk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricsList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []Metrics `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsList) RawJSON() string { return r.JSON.raw }
func (r *MetricsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceMetricListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Time interval.
	TimeInterval int64 `json:"time_interval" api:"required"`
	// Time interval unit.
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `json:"time_unit,omitzero" api:"required"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

func (r InstanceMetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceMetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceMetricListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [InstanceMetricListParams]'s query parameters as
// `url.Values`.
func (r InstanceMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
