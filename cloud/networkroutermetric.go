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

// NetworkRouterMetricService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkRouterMetricService] method instead.
type NetworkRouterMetricService struct {
	Options []option.RequestOption
}

// NewNetworkRouterMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkRouterMetricService(opts ...option.RequestOption) (r NetworkRouterMetricService) {
	r = NetworkRouterMetricService{}
	r.Options = opts
	return
}

// Get the network traffic of each router interface, in bits and packets per
// second.
func (r *NetworkRouterMetricService) List(ctx context.Context, routerID string, params NetworkRouterMetricListParams, opts ...option.RequestOption) (res *RouterMetricsList, err error) {
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
	if routerID == "" {
		err = errors.New("missing required router_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/routers/%v/%v/%s/metrics", params.ProjectID.Value, params.RegionID.Value, routerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type RouterMetrics struct {
	// Interface type. The public interface includes the router gateway and the
	// floating IPs. A private interface connects the router to a subnet.
	//
	// Any of "private", "public".
	InterfaceType RouterMetricsInterfaceType `json:"interface_type" api:"required"`
	// IP addresses of the port
	IPAddresses []string `json:"ip_addresses" api:"required" format:"ipvanyaddress"`
	// Traffic metrics of the interface, in order of time
	Metrics []RouterMetricsMetric `json:"metrics" api:"required"`
	// Network ID
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Port ID
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterfaceType respjson.Field
		IPAddresses   respjson.Field
		Metrics       respjson.Field
		NetworkID     respjson.Field
		PortID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterMetrics) RawJSON() string { return r.JSON.raw }
func (r *RouterMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interface type. The public interface includes the router gateway and the
// floating IPs. A private interface connects the router to a subnet.
type RouterMetricsInterfaceType string

const (
	RouterMetricsInterfaceTypePrivate RouterMetricsInterfaceType = "private"
	RouterMetricsInterfaceTypePublic  RouterMetricsInterfaceType = "public"
)

type RouterMetricsMetric struct {
	// Network out, bits per second. Null when no data point is available.
	NetworkBpsEgress float64 `json:"network_bps_egress" api:"required"`
	// Network in, bits per second. Null when no data point is available.
	NetworkBpsIngress float64 `json:"network_bps_ingress" api:"required"`
	// Network out, packets per second. Null when no data point is available.
	NetworkPpsEgress float64 `json:"network_pps_egress" api:"required"`
	// Network in, packets per second. Null when no data point is available.
	NetworkPpsIngress float64 `json:"network_pps_ingress" api:"required"`
	// Timestamp
	Time string `json:"time" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
func (r RouterMetricsMetric) RawJSON() string { return r.JSON.raw }
func (r *RouterMetricsMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouterMetricsList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []RouterMetrics `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouterMetricsList) RawJSON() string { return r.JSON.raw }
func (r *RouterMetricsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterMetricListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Time interval
	TimeInterval int64 `json:"time_interval" api:"required"`
	// Time interval unit
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `json:"time_unit,omitzero" api:"required"`
	paramObj
}

func (r NetworkRouterMetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterMetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterMetricListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
