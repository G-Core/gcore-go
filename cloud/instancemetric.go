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
	"github.com/G-Core/gcore-go/packages/pagination"
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

// Returns the CPU, memory, network and disk metrics of an instance for a time
// range. The results are in ascending order of `time`.
func (r *InstanceMetricService) List(ctx context.Context, instanceID string, params InstanceMetricListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Metrics], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s/metrics", params.ProjectID.Value, instanceID)
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

// Returns the CPU, memory, network and disk metrics of an instance for a time
// range. The results are in ascending order of `time`.
func (r *InstanceMetricService) ListAutoPaging(ctx context.Context, instanceID string, params InstanceMetricListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Metrics] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, instanceID, params, opts...))
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
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Length of the time range, in the `time_unit`. The range is 1, 6 or 24 hours, or
	// 1, 7 or 14 days.
	TimeInterval int64 `query:"time_interval" api:"required" json:"-"`
	// Unit of the `time_interval`.
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `query:"time_unit,omitzero" api:"required" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceMetricListParams]'s query parameters as
// `url.Values`.
func (r InstanceMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
