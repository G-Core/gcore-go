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
	"github.com/stainless-sdks/gcore-go/packages/respjson"
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metrics", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Instance metrics item
type Metrics struct {
	// Timestamp
	Time string `json:"time,required"`
	// CPU utilization, % (max 100% for multi-core)
	CPUUtil float64 `json:"cpu_util,nullable"`
	// Disks metrics for each of the disks attached
	Disks []MetricsDisk `json:"disks,nullable"`
	// RAM utilization, %
	MemoryUtil float64 `json:"memory_util,nullable"`
	// Network out, bytes per second
	NetworkBpsEgress float64 `json:"network_Bps_egress,nullable"`
	// Network in, bytes per second
	NetworkBpsIngress float64 `json:"network_Bps_ingress,nullable"`
	// Network out, packets per second
	NetworkPpsEgress float64 `json:"network_pps_egress,nullable"`
	// Network in, packets per second
	NetworkPpsIngress float64 `json:"network_pps_ingress,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Time              respjson.Field
		CPUUtil           respjson.Field
		Disks             respjson.Field
		MemoryUtil        respjson.Field
		NetworkBpsEgress  respjson.Field
		NetworkBpsIngress respjson.Field
		NetworkPpsEgress  respjson.Field
		NetworkPpsIngress respjson.Field
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
	// Disk read, Bytes per second
	DiskBpsRead float64 `json:"disk_Bps_read,nullable"`
	// Disk write, Bytes per second
	DiskBpsWrite float64 `json:"disk_Bps_write,nullable"`
	// Disk read, iops
	DiskIopsRead float64 `json:"disk_iops_read,nullable"`
	// Disk write, iops
	DiskIopsWrite float64 `json:"disk_iops_write,nullable"`
	// Disk attached slot name
	DiskName string `json:"disk_name,nullable"`
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
	Count int64 `json:"count,required"`
	// Objects
	Results []Metrics `json:"results,required"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Time interval.
	TimeInterval int64 `json:"time_interval,required"`
	// Time interval unit.
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `json:"time_unit,omitzero,required"`
	paramObj
}

func (r InstanceMetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceMetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
