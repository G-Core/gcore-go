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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// '#/components/schemas/InstanceMetricsSerializer'
// "$.components.schemas.InstanceMetricsSerializer"
type Metrics struct {
	// '#/components/schemas/InstanceMetricsSerializer/properties/time'
	// "$.components.schemas.InstanceMetricsSerializer.properties.time"
	Time string `json:"time,required"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/cpu_util/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.cpu_util.anyOf[0]"
	CPUUtil float64 `json:"cpu_util,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/disks/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.disks.anyOf[0]"
	Disks []MetricsDisk `json:"disks,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/memory_util/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.memory_util.anyOf[0]"
	MemoryUtil float64 `json:"memory_util,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/network_Bps_egress/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.network_Bps_egress.anyOf[0]"
	NetworkBpsEgress float64 `json:"network_Bps_egress,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/network_Bps_ingress/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.network_Bps_ingress.anyOf[0]"
	NetworkBpsIngress float64 `json:"network_Bps_ingress,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/network_pps_egress/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.network_pps_egress.anyOf[0]"
	NetworkPpsEgress float64 `json:"network_pps_egress,nullable"`
	// '#/components/schemas/InstanceMetricsSerializer/properties/network_pps_ingress/anyOf/0'
	// "$.components.schemas.InstanceMetricsSerializer.properties.network_pps_ingress.anyOf[0]"
	NetworkPpsIngress float64 `json:"network_pps_ingress,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Time              resp.Field
		CPUUtil           resp.Field
		Disks             resp.Field
		MemoryUtil        resp.Field
		NetworkBpsEgress  resp.Field
		NetworkBpsIngress resp.Field
		NetworkPpsEgress  resp.Field
		NetworkPpsIngress resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metrics) RawJSON() string { return r.JSON.raw }
func (r *Metrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceMetricsSerializer/properties/disks/anyOf/0/items'
// "$.components.schemas.InstanceMetricsSerializer.properties.disks.anyOf[0].items"
type MetricsDisk struct {
	// '#/components/schemas/DisksMetrics/properties/disk_Bps_read/anyOf/0'
	// "$.components.schemas.DisksMetrics.properties.disk_Bps_read.anyOf[0]"
	DiskBpsRead float64 `json:"disk_Bps_read,nullable"`
	// '#/components/schemas/DisksMetrics/properties/disk_Bps_write/anyOf/0'
	// "$.components.schemas.DisksMetrics.properties.disk_Bps_write.anyOf[0]"
	DiskBpsWrite float64 `json:"disk_Bps_write,nullable"`
	// '#/components/schemas/DisksMetrics/properties/disk_iops_read/anyOf/0'
	// "$.components.schemas.DisksMetrics.properties.disk_iops_read.anyOf[0]"
	DiskIopsRead float64 `json:"disk_iops_read,nullable"`
	// '#/components/schemas/DisksMetrics/properties/disk_iops_write/anyOf/0'
	// "$.components.schemas.DisksMetrics.properties.disk_iops_write.anyOf[0]"
	DiskIopsWrite float64 `json:"disk_iops_write,nullable"`
	// '#/components/schemas/DisksMetrics/properties/disk_name/anyOf/0'
	// "$.components.schemas.DisksMetrics.properties.disk_name.anyOf[0]"
	DiskName string `json:"disk_name,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		DiskBpsRead   resp.Field
		DiskBpsWrite  resp.Field
		DiskIopsRead  resp.Field
		DiskIopsWrite resp.Field
		DiskName      resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsDisk) RawJSON() string { return r.JSON.raw }
func (r *MetricsDisk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceMetricsListSerializer'
// "$.components.schemas.InstanceMetricsListSerializer"
type MetricsList struct {
	// '#/components/schemas/InstanceMetricsListSerializer/properties/count'
	// "$.components.schemas.InstanceMetricsListSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/InstanceMetricsListSerializer/properties/results'
	// "$.components.schemas.InstanceMetricsListSerializer.properties.results"
	Results []Metrics `json:"results,required"`
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
func (r MetricsList) RawJSON() string { return r.JSON.raw }
func (r *MetricsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceMetricListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fmetrics/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/metrics'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fmetrics/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/metrics'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/InstanceMetricsRequestSerializer/properties/time_interval'
	// "$.components.schemas.InstanceMetricsRequestSerializer.properties.time_interval"
	TimeInterval int64 `json:"time_interval,required"`
	// '#/components/schemas/InstanceMetricsRequestSerializer/properties/time_unit'
	// "$.components.schemas.InstanceMetricsRequestSerializer.properties.time_unit"
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `json:"time_unit,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceMetricListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceMetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceMetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
