// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// CostReportService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCostReportService] method instead.
type CostReportService struct {
	Options []option.RequestOption
}

// NewCostReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCostReportService(opts ...option.RequestOption) (r CostReportService) {
	r = CostReportService{}
	r.Options = opts
	return
}

// Retrieve a detailed cost report totals for a specified month, which includes
// both commit and pay-as-you-go (overcommit) prices. Additionally, it provides the
// spent billing units (e.g., hours or GB) for resources. The "time_to" parameter
// represents all days in the specified month.
//
// Data from the past hour may not reflect the full set of statistics. For the most
// complete and accurate results, we recommend accessing the data at least one hour
// after the relevant time period. Updates are generally available within a 24-hour
// window, though timing can vary. Scheduled maintenance or other exceptions may
// occasionally cause delays beyond 24 hours.
func (r *CostReportService) GetAggregatedMonthly(ctx context.Context, body CostReportGetAggregatedMonthlyParams, opts ...option.RequestOption) (res *CostReportAggregatedMonthly, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/reservation_cost_report/totals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get daily costs for a given period, summed down to one dimension: region
// (default), project, resource type or flavor. Requested period should not exceed
// 31 days. Costs and usages are returned at full precision.
//
// The response can have more than one row for one day and one `group_by` value
// (for example, one region on one day). Usage with a different unit (`usage.unit`,
// for example `minutes` or `bytes`) is in a different row. Usage without a price
// (`error` is not null) is also in a different row.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/`reservation_cost_report`/totals, as the results from this report will not
// be accurate.
//
// Data from the past hour may not reflect the full set of statistics. For the most
// complete and accurate results, we recommend accessing the data at least one hour
// after the relevant time period. Updates are generally available within a 24-hour
// window, though timing can vary. Scheduled maintenance or other exceptions may
// occasionally cause delays beyond 24 hours.
func (r *CostReportService) GetBreakdown(ctx context.Context, body CostReportGetBreakdownParams, opts ...option.RequestOption) (res *CostReportBreakdown, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v2/cost_report/breakdown"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a detailed cost report for a given period and specific resources, sorted by
// cost by default. Requested period should not exceed 31 days. Costs and usages
// are always returned at full precision.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/`reservation_cost_report`/totals, as the results from this report will not
// be accurate.
//
// Data from the past hour may not reflect the full set of statistics. For the most
// complete and accurate results, we recommend accessing the data at least one hour
// after the relevant time period. Updates are generally available within a 24-hour
// window, though timing can vary. Scheduled maintenance or other exceptions may
// occasionally cause delays beyond 24 hours.
func (r *CostReportService) GetDetailed(ctx context.Context, body CostReportGetDetailedParams, opts ...option.RequestOption) (res *CostReportDetailed, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v2/cost_report/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CostReportAggregatedMonthly struct {
	// Total count of the totals
	Count int64 `json:"count" api:"required"`
	// Price status for the UI, type: string
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportAggregatedMonthlyPriceStatus   `json:"price_status" api:"required"`
	Results     []CostReportAggregatedMonthlyResultUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		PriceStatus respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthly) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthly) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI, type: string
type CostReportAggregatedMonthlyPriceStatus string

const (
	CostReportAggregatedMonthlyPriceStatusError CostReportAggregatedMonthlyPriceStatus = "error"
	CostReportAggregatedMonthlyPriceStatusHide  CostReportAggregatedMonthlyPriceStatus = "hide"
	CostReportAggregatedMonthlyPriceStatusShow  CostReportAggregatedMonthlyPriceStatus = "show"
)

// CostReportAggregatedMonthlyResultUnion contains all possible properties and
// values from [CostReportAggregatedMonthlyResultAICluster],
// [CostReportAggregatedMonthlyResultAIVirtualCluster],
// [CostReportAggregatedMonthlyResultBaremetal],
// [CostReportAggregatedMonthlyResultBasicVm],
// [CostReportAggregatedMonthlyResultBackup],
// [CostReportAggregatedMonthlyResultContainers],
// [CostReportAggregatedMonthlyResultEgressTraffic],
// [CostReportAggregatedMonthlyResultExternalIP],
// [CostReportAggregatedMonthlyResultFileShare],
// [CostReportAggregatedMonthlyResultFloatingip],
// [CostReportAggregatedMonthlyResultFunctions],
// [CostReportAggregatedMonthlyResultFunctionsCalls],
// [CostReportAggregatedMonthlyResultFunctionsTraffic],
// [CostReportAggregatedMonthlyResultImage],
// [CostReportAggregatedMonthlyResultInference],
// [CostReportAggregatedMonthlyResultInstance],
// [CostReportAggregatedMonthlyResultLoadBalancer],
// [CostReportAggregatedMonthlyResultLogIndex],
// [CostReportAggregatedMonthlyResultSnapshot],
// [CostReportAggregatedMonthlyResultVolume],
// [CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler],
// [CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory],
// [CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork],
// [CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU],
// [CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume].
//
// Use the [CostReportAggregatedMonthlyResultUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CostReportAggregatedMonthlyResultUnion struct {
	BillingFeatureName string  `json:"billing_feature_name"`
	BillingMetricName  string  `json:"billing_metric_name"`
	BillingValue       float64 `json:"billing_value"`
	BillingValueUnit   string  `json:"billing_value_unit"`
	Cost               float64 `json:"cost"`
	Currency           string  `json:"currency"`
	Err                string  `json:"err"`
	Flavor             string  `json:"flavor"`
	Region             int64   `json:"region"`
	RegionID           int64   `json:"region_id"`
	// Any of "ai_cluster", "ai_virtual_cluster", "baremetal", "basic_vm", "backup",
	// "containers", "egress_traffic", "external_ip", "file_share", "floatingip",
	// "functions", "functions_calls", "functions_traffic", "image", "inference",
	// "instance", "load_balancer", "log_index", "snapshot", "volume",
	// "dbaas_postgresql_connection_pooler", "dbaas_postgresql_memory",
	// "dbaas_postgresql_public_network", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_volume".
	Type string `json:"type"`
	// This field is from variant [CostReportAggregatedMonthlyResultBackup].
	LastSize int64 `json:"last_size"`
	// This field is from variant [CostReportAggregatedMonthlyResultEgressTraffic].
	InstanceType string `json:"instance_type"`
	// This field is from variant [CostReportAggregatedMonthlyResultFileShare].
	FileShareType string `json:"file_share_type"`
	VolumeType    string `json:"volume_type"`
	JSON          struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		LastSize           respjson.Field
		InstanceType       respjson.Field
		FileShareType      respjson.Field
		VolumeType         respjson.Field
		raw                string
	} `json:"-"`
}

// anyCostReportAggregatedMonthlyResult is implemented by each variant of
// [CostReportAggregatedMonthlyResultUnion] to add type safety for the return type
// of [CostReportAggregatedMonthlyResultUnion.AsAny]
type anyCostReportAggregatedMonthlyResult interface {
	implCostReportAggregatedMonthlyResultUnion()
}

func (CostReportAggregatedMonthlyResultAICluster) implCostReportAggregatedMonthlyResultUnion() {}
func (CostReportAggregatedMonthlyResultAIVirtualCluster) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultBaremetal) implCostReportAggregatedMonthlyResultUnion()      {}
func (CostReportAggregatedMonthlyResultBasicVm) implCostReportAggregatedMonthlyResultUnion()        {}
func (CostReportAggregatedMonthlyResultBackup) implCostReportAggregatedMonthlyResultUnion()         {}
func (CostReportAggregatedMonthlyResultContainers) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultEgressTraffic) implCostReportAggregatedMonthlyResultUnion()  {}
func (CostReportAggregatedMonthlyResultExternalIP) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultFileShare) implCostReportAggregatedMonthlyResultUnion()      {}
func (CostReportAggregatedMonthlyResultFloatingip) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultFunctions) implCostReportAggregatedMonthlyResultUnion()      {}
func (CostReportAggregatedMonthlyResultFunctionsCalls) implCostReportAggregatedMonthlyResultUnion() {}
func (CostReportAggregatedMonthlyResultFunctionsTraffic) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultImage) implCostReportAggregatedMonthlyResultUnion()        {}
func (CostReportAggregatedMonthlyResultInference) implCostReportAggregatedMonthlyResultUnion()    {}
func (CostReportAggregatedMonthlyResultInstance) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultLoadBalancer) implCostReportAggregatedMonthlyResultUnion() {}
func (CostReportAggregatedMonthlyResultLogIndex) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultSnapshot) implCostReportAggregatedMonthlyResultUnion()     {}
func (CostReportAggregatedMonthlyResultVolume) implCostReportAggregatedMonthlyResultUnion()       {}
func (CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU) implCostReportAggregatedMonthlyResultUnion() {
}
func (CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume) implCostReportAggregatedMonthlyResultUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := CostReportAggregatedMonthlyResultUnion.AsAny().(type) {
//	case cloud.CostReportAggregatedMonthlyResultAICluster:
//	case cloud.CostReportAggregatedMonthlyResultAIVirtualCluster:
//	case cloud.CostReportAggregatedMonthlyResultBaremetal:
//	case cloud.CostReportAggregatedMonthlyResultBasicVm:
//	case cloud.CostReportAggregatedMonthlyResultBackup:
//	case cloud.CostReportAggregatedMonthlyResultContainers:
//	case cloud.CostReportAggregatedMonthlyResultEgressTraffic:
//	case cloud.CostReportAggregatedMonthlyResultExternalIP:
//	case cloud.CostReportAggregatedMonthlyResultFileShare:
//	case cloud.CostReportAggregatedMonthlyResultFloatingip:
//	case cloud.CostReportAggregatedMonthlyResultFunctions:
//	case cloud.CostReportAggregatedMonthlyResultFunctionsCalls:
//	case cloud.CostReportAggregatedMonthlyResultFunctionsTraffic:
//	case cloud.CostReportAggregatedMonthlyResultImage:
//	case cloud.CostReportAggregatedMonthlyResultInference:
//	case cloud.CostReportAggregatedMonthlyResultInstance:
//	case cloud.CostReportAggregatedMonthlyResultLoadBalancer:
//	case cloud.CostReportAggregatedMonthlyResultLogIndex:
//	case cloud.CostReportAggregatedMonthlyResultSnapshot:
//	case cloud.CostReportAggregatedMonthlyResultVolume:
//	case cloud.CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler:
//	case cloud.CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory:
//	case cloud.CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork:
//	case cloud.CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU:
//	case cloud.CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CostReportAggregatedMonthlyResultUnion) AsAny() anyCostReportAggregatedMonthlyResult {
	switch u.Type {
	case "ai_cluster":
		return u.AsAICluster()
	case "ai_virtual_cluster":
		return u.AsAIVirtualCluster()
	case "baremetal":
		return u.AsBaremetal()
	case "basic_vm":
		return u.AsBasicVm()
	case "backup":
		return u.AsBackup()
	case "containers":
		return u.AsContainers()
	case "egress_traffic":
		return u.AsEgressTraffic()
	case "external_ip":
		return u.AsExternalIP()
	case "file_share":
		return u.AsFileShare()
	case "floatingip":
		return u.AsFloatingip()
	case "functions":
		return u.AsFunctions()
	case "functions_calls":
		return u.AsFunctionsCalls()
	case "functions_traffic":
		return u.AsFunctionsTraffic()
	case "image":
		return u.AsImage()
	case "inference":
		return u.AsInference()
	case "instance":
		return u.AsInstance()
	case "load_balancer":
		return u.AsLoadBalancer()
	case "log_index":
		return u.AsLogIndex()
	case "snapshot":
		return u.AsSnapshot()
	case "volume":
		return u.AsVolume()
	case "dbaas_postgresql_connection_pooler":
		return u.AsDbaasPostgreSQLConnectionPooler()
	case "dbaas_postgresql_memory":
		return u.AsDbaasPostgreSQLMemory()
	case "dbaas_postgresql_public_network":
		return u.AsDbaasPostgreSQLPublicNetwork()
	case "dbaas_postgresql_cpu":
		return u.AsDbaasPostgreSQLCPU()
	case "dbaas_postgresql_volume":
		return u.AsDbaasPostgreSQLVolume()
	}
	return nil
}

func (u CostReportAggregatedMonthlyResultUnion) AsAICluster() (v CostReportAggregatedMonthlyResultAICluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsAIVirtualCluster() (v CostReportAggregatedMonthlyResultAIVirtualCluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsBaremetal() (v CostReportAggregatedMonthlyResultBaremetal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsBasicVm() (v CostReportAggregatedMonthlyResultBasicVm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsBackup() (v CostReportAggregatedMonthlyResultBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsContainers() (v CostReportAggregatedMonthlyResultContainers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsEgressTraffic() (v CostReportAggregatedMonthlyResultEgressTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsExternalIP() (v CostReportAggregatedMonthlyResultExternalIP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsFileShare() (v CostReportAggregatedMonthlyResultFileShare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsFloatingip() (v CostReportAggregatedMonthlyResultFloatingip) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsFunctions() (v CostReportAggregatedMonthlyResultFunctions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsFunctionsCalls() (v CostReportAggregatedMonthlyResultFunctionsCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsFunctionsTraffic() (v CostReportAggregatedMonthlyResultFunctionsTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsImage() (v CostReportAggregatedMonthlyResultImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsInference() (v CostReportAggregatedMonthlyResultInference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsInstance() (v CostReportAggregatedMonthlyResultInstance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsLoadBalancer() (v CostReportAggregatedMonthlyResultLoadBalancer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsLogIndex() (v CostReportAggregatedMonthlyResultLogIndex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsSnapshot() (v CostReportAggregatedMonthlyResultSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsVolume() (v CostReportAggregatedMonthlyResultVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsDbaasPostgreSQLConnectionPooler() (v CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsDbaasPostgreSQLMemory() (v CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsDbaasPostgreSQLPublicNetwork() (v CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsDbaasPostgreSQLCPU() (v CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedMonthlyResultUnion) AsDbaasPostgreSQLVolume() (v CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CostReportAggregatedMonthlyResultUnion) RawJSON() string { return u.JSON.raw }

func (r *CostReportAggregatedMonthlyResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultAICluster struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.AICluster `json:"type" default:"ai_cluster"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultAICluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultAIVirtualCluster struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                     `json:"region_id" api:"required"`
	Type     constant.AIVirtualCluster `json:"type" default:"ai_virtual_cluster"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultAIVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultBaremetal struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Baremetal `json:"type" default:"baremetal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultBaremetal) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultBasicVm struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64            `json:"region_id" api:"required"`
	Type     constant.BasicVm `json:"type" default:"basic_vm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultBasicVm) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultBackup struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64           `json:"region_id" api:"required"`
	Type     constant.Backup `json:"type" default:"backup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		LastSize           respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultBackup) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultContainers struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" default:"GBS"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.Containers `json:"type" default:"containers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultContainers) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultEgressTraffic struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit" default:"bytes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Type of the instance
	//
	// Any of "baremetal", "router", "vm".
	InstanceType string `json:"instance_type" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                  `json:"region_id" api:"required"`
	Type     constant.EgressTraffic `json:"type" default:"egress_traffic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		InstanceType       respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultExternalIP struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.ExternalIP `json:"type" default:"external_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultExternalIP) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultFileShare struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Type of the file share
	FileShareType string `json:"file_share_type" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.FileShare `json:"type" default:"file_share"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FileShareType      respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultFileShare) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultFloatingip struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.Floatingip `json:"type" default:"floatingip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultFloatingip) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultFunctions struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" default:"GBS"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Functions `json:"type" default:"functions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultFunctions) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultFunctionsCalls struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit" default:"MLS"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                   `json:"region_id" api:"required"`
	Type     constant.FunctionsCalls `json:"type" default:"functions_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultFunctionsCalls) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultFunctionsTraffic struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit" default:"GB"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                     `json:"region_id" api:"required"`
	Type     constant.FunctionsTraffic `json:"type" default:"functions_traffic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultFunctionsTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultImage struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64          `json:"region_id" api:"required"`
	Type     constant.Image `json:"type" default:"image"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultImage) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultInference struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit" api:"required"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Inference `json:"type" default:"inference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultInference) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultInstance struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the instance
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.Instance `json:"type" default:"instance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultInstance) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultLoadBalancer struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                 `json:"region_id" api:"required"`
	Type     constant.LoadBalancer `json:"type" default:"load_balancer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Flavor             respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultLogIndex struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.LogIndex `json:"type" default:"log_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultLogIndex) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultSnapshot struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.Snapshot `json:"type" default:"snapshot"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		VolumeType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultSnapshot) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultVolume struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64           `json:"region_id" api:"required"`
	Type     constant.Volume `json:"type" default:"volume"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		VolumeType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                    `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type" default:"dbaas_postgresql_connection_pooler"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler) RawJSON() string {
	return r.JSON.raw
}
func (r *CostReportAggregatedMonthlyResultDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                          `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type" default:"dbaas_postgresql_memory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                 `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type" default:"dbaas_postgresql_public_network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork) RawJSON() string {
	return r.JSON.raw
}
func (r *CostReportAggregatedMonthlyResultDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" default:"minutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                       `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type" default:"dbaas_postgresql_cpu"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume struct {
	BillingFeatureName string `json:"billing_feature_name" api:"required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" default:"gbminutes"`
	// Cost for requested period
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost
	Currency string `json:"currency" api:"required"`
	// Error message
	Err string `json:"err" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                          `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type" default:"dbaas_postgresql_volume"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		VolumeType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedMonthlyResultDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportBreakdown struct {
	// Count of returned rows
	Count int64 `json:"count" api:"required"`
	// 'show' when the client may see prices, 'hide' for trial clients, clients with no
	// product status, and clients of a reseller that does not expose prices. This
	// report never returns 'error'.
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportBreakdownPriceStatus `json:"price_status" api:"required"`
	Results     []CostReportBreakdownResult    `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		PriceStatus respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportBreakdown) RawJSON() string { return r.JSON.raw }
func (r *CostReportBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'show' when the client may see prices, 'hide' for trial clients, clients with no
// product status, and clients of a reseller that does not expose prices. This
// report never returns 'error'.
type CostReportBreakdownPriceStatus string

const (
	CostReportBreakdownPriceStatusError CostReportBreakdownPriceStatus = "error"
	CostReportBreakdownPriceStatusHide  CostReportBreakdownPriceStatus = "hide"
	CostReportBreakdownPriceStatusShow  CostReportBreakdownPriceStatus = "show"
)

// One day of usage for one `group_by` value.
type CostReportBreakdownResult struct {
	// Day the usage was billed for
	Date time.Time `json:"date" api:"required" format:"date"`
	// Why a row has no cost. Values are the plain strings v1 has always returned.
	//
	// Any of "Multiple subscriptions", "No tariff".
	Error string `json:"error" api:"required"`
	// Flavor name, set when `group_by` is 'flavor'
	Flavor string `json:"flavor" api:"required"`
	// Project ID, set when `group_by` is 'project'
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID, set when `group_by` is 'region'. 0 means global usage that has no
	// region.
	RegionID int64 `json:"region_id" api:"required"`
	// Resource types for prebilling report
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "inference_public_models", "instance", "load_balancer",
	// "log_index", "snapshot", "volume".
	Type string `json:"type" api:"required"`
	// How much the group was used that day, and what that costs.
	Usage CostReportBreakdownResultUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Error       respjson.Field
		Flavor      respjson.Field
		ProjectID   respjson.Field
		RegionID    respjson.Field
		Type        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportBreakdownResult) RawJSON() string { return r.JSON.raw }
func (r *CostReportBreakdownResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How much the group was used that day, and what that costs.
type CostReportBreakdownResultUsage struct {
	// Cost for the day, never rounded. Null if there is no price.
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost, null if there is no cost
	Currency string `json:"currency" api:"required"`
	// Unit of the billing value
	Unit string `json:"unit" api:"required"`
	// Usage of the billing metric for the day, in `unit`, kept at full precision.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Currency    respjson.Field
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportBreakdownResultUsage) RawJSON() string { return r.JSON.raw }
func (r *CostReportBreakdownResultUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailed struct {
	// Count of all the resources
	Count int64 `json:"count" api:"required"`
	// 'show' when the client may see prices, 'hide' for trial clients, clients with no
	// product status, and clients of a reseller that does not expose prices.
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportDetailedPriceStatus `json:"price_status" api:"required"`
	Results     []CostReportDetailedResult    `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		PriceStatus respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailed) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'show' when the client may see prices, 'hide' for trial clients, clients with no
// product status, and clients of a reseller that does not expose prices.
type CostReportDetailedPriceStatus string

const (
	CostReportDetailedPriceStatusError CostReportDetailedPriceStatus = "error"
	CostReportDetailedPriceStatusHide  CostReportDetailedPriceStatus = "hide"
	CostReportDetailedPriceStatusShow  CostReportDetailedPriceStatus = "show"
)

type CostReportDetailedResult struct {
	// Why a row has no cost. Values are the plain strings v1 has always returned.
	//
	// Any of "Multiple subscriptions", "No tariff".
	Error string `json:"error" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID. 0 means global usage that has no region.
	RegionID int64 `json:"region_id" api:"required"`
	// What the resource is.
	Resource CostReportDetailedResultResource `json:"resource" api:"required"`
	// How much it was used, and what that costs.
	Usage CostReportDetailedResultUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ProjectID   respjson.Field
		RegionID    respjson.Field
		Resource    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResult) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What the resource is.
type CostReportDetailedResultResource struct {
	// Flavor name
	Flavor string `json:"flavor" api:"required"`
	// Name of the resource
	Name string `json:"name" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	// Resource type
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "inference_public_models", "instance", "load_balancer",
	// "log_index", "snapshot", "volume".
	Type string `json:"type" api:"required"`
	// UUID of the resource
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flavor      respjson.Field
		Name        respjson.Field
		Tags        respjson.Field
		Type        respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultResource) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How much it was used, and what that costs.
type CostReportDetailedResultUsage struct {
	// Cost for the requested period, never rounded. Null if there is no price.
	Cost float64 `json:"cost" api:"required"`
	// Currency of the cost, null if there is no cost
	Currency string `json:"currency" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Last time the resource was seen in the given period. Earlier than 'time_to'
	// means the resource stopped before the period ended.
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Unit of the billing value
	Unit string `json:"unit" api:"required"`
	// Usage of the billing metric for the period, in `unit`, kept at full precision.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Currency    respjson.Field
		FirstSeen   respjson.Field
		LastSeen    respjson.Field
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultUsage) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportGetAggregatedMonthlyParams struct {
	// Round cost values to 5 decimal places. When false, returns full precision.
	Rounding param.Opt[bool] `json:"rounding,omitzero"`
	// Deprecated. Use `year_month` instead. Beginning of the period: YYYY-mm
	TimeFrom param.Opt[time.Time] `json:"time_from,omitzero" format:"date-time"`
	// Deprecated. Use `year_month` instead. End of the period: YYYY-mm
	TimeTo param.Opt[time.Time] `json:"time_to,omitzero" format:"date-time"`
	// Year and month in the format YYYY-MM
	YearMonth param.Opt[string] `json:"year_month,omitzero"`
	// List of region IDs.
	Regions []int64 `json:"regions,omitzero"`
	// Format of the response (`csv_totals` or json).
	//
	// Any of "csv_totals", "json".
	ResponseFormat CostReportGetAggregatedMonthlyParamsResponseFormat `json:"response_format,omitzero"`
	// Extended filter for field filtering.
	SchemaFilter CostReportGetAggregatedMonthlyParamsSchemaFilterUnion `json:"schema_filter,omitzero"`
	// Filter by tags
	Tags CostReportGetAggregatedMonthlyParamsTags `json:"tags,omitzero"`
	// List of resource types to be filtered in the report.
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "inference_public_models", "instance", "load_balancer",
	// "log_index", "snapshot", "volume".
	Types []string `json:"types,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParams) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format of the response (`csv_totals` or json).
type CostReportGetAggregatedMonthlyParamsResponseFormat string

const (
	CostReportGetAggregatedMonthlyParamsResponseFormatCsvTotals CostReportGetAggregatedMonthlyParamsResponseFormat = "csv_totals"
	CostReportGetAggregatedMonthlyParamsResponseFormatJson      CostReportGetAggregatedMonthlyParamsResponseFormat = "json"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CostReportGetAggregatedMonthlyParamsSchemaFilterUnion struct {
	OfSnapshot                        *CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot                        `json:",omitzero,inline"`
	OfInstance                        *CostReportGetAggregatedMonthlyParamsSchemaFilterInstance                        `json:",omitzero,inline"`
	OfAICluster                       *CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster                       `json:",omitzero,inline"`
	OfAIVirtualCluster                *CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster                `json:",omitzero,inline"`
	OfBasicVm                         *CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm                         `json:",omitzero,inline"`
	OfBaremetal                       *CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal                       `json:",omitzero,inline"`
	OfVolume                          *CostReportGetAggregatedMonthlyParamsSchemaFilterVolume                          `json:",omitzero,inline"`
	OfFileShare                       *CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare                       `json:",omitzero,inline"`
	OfImage                           *CostReportGetAggregatedMonthlyParamsSchemaFilterImage                           `json:",omitzero,inline"`
	OfFloatingip                      *CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip                      `json:",omitzero,inline"`
	OfEgressTraffic                   *CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic                   `json:",omitzero,inline"`
	OfLoadBalancer                    *CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer                    `json:",omitzero,inline"`
	OfExternalIP                      *CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP                      `json:",omitzero,inline"`
	OfBackup                          *CostReportGetAggregatedMonthlyParamsSchemaFilterBackup                          `json:",omitzero,inline"`
	OfLogIndex                        *CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex                        `json:",omitzero,inline"`
	OfFunctions                       *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions                       `json:",omitzero,inline"`
	OfFunctionsCalls                  *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls                  `json:",omitzero,inline"`
	OfFunctionsTraffic                *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic                `json:",omitzero,inline"`
	OfContainers                      *CostReportGetAggregatedMonthlyParamsSchemaFilterContainers                      `json:",omitzero,inline"`
	OfInference                       *CostReportGetAggregatedMonthlyParamsSchemaFilterInference                       `json:",omitzero,inline"`
	OfDbaasPostgreSQLVolume           *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume           `json:",omitzero,inline"`
	OfDbaasPostgreSQLPublicNetwork    *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork    `json:",omitzero,inline"`
	OfDbaasPostgreSQLCPU              *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU              `json:",omitzero,inline"`
	OfDbaasPostgreSQLMemory           *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory           `json:",omitzero,inline"`
	OfDbaasPostgreSQLConnectionPooler *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler `json:",omitzero,inline"`
	paramUnion
}

func (u CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSnapshot,
		u.OfInstance,
		u.OfAICluster,
		u.OfAIVirtualCluster,
		u.OfBasicVm,
		u.OfBaremetal,
		u.OfVolume,
		u.OfFileShare,
		u.OfImage,
		u.OfFloatingip,
		u.OfEgressTraffic,
		u.OfLoadBalancer,
		u.OfExternalIP,
		u.OfBackup,
		u.OfLogIndex,
		u.OfFunctions,
		u.OfFunctionsCalls,
		u.OfFunctionsTraffic,
		u.OfContainers,
		u.OfInference,
		u.OfDbaasPostgreSQLVolume,
		u.OfDbaasPostgreSQLPublicNetwork,
		u.OfDbaasPostgreSQLCPU,
		u.OfDbaasPostgreSQLMemory,
		u.OfDbaasPostgreSQLConnectionPooler)
}
func (u *CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) asAny() any {
	if !param.IsOmitted(u.OfSnapshot) {
		return u.OfSnapshot
	} else if !param.IsOmitted(u.OfInstance) {
		return u.OfInstance
	} else if !param.IsOmitted(u.OfAICluster) {
		return u.OfAICluster
	} else if !param.IsOmitted(u.OfAIVirtualCluster) {
		return u.OfAIVirtualCluster
	} else if !param.IsOmitted(u.OfBasicVm) {
		return u.OfBasicVm
	} else if !param.IsOmitted(u.OfBaremetal) {
		return u.OfBaremetal
	} else if !param.IsOmitted(u.OfVolume) {
		return u.OfVolume
	} else if !param.IsOmitted(u.OfFileShare) {
		return u.OfFileShare
	} else if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfFloatingip) {
		return u.OfFloatingip
	} else if !param.IsOmitted(u.OfEgressTraffic) {
		return u.OfEgressTraffic
	} else if !param.IsOmitted(u.OfLoadBalancer) {
		return u.OfLoadBalancer
	} else if !param.IsOmitted(u.OfExternalIP) {
		return u.OfExternalIP
	} else if !param.IsOmitted(u.OfBackup) {
		return u.OfBackup
	} else if !param.IsOmitted(u.OfLogIndex) {
		return u.OfLogIndex
	} else if !param.IsOmitted(u.OfFunctions) {
		return u.OfFunctions
	} else if !param.IsOmitted(u.OfFunctionsCalls) {
		return u.OfFunctionsCalls
	} else if !param.IsOmitted(u.OfFunctionsTraffic) {
		return u.OfFunctionsTraffic
	} else if !param.IsOmitted(u.OfContainers) {
		return u.OfContainers
	} else if !param.IsOmitted(u.OfInference) {
		return u.OfInference
	} else if !param.IsOmitted(u.OfDbaasPostgreSQLVolume) {
		return u.OfDbaasPostgreSQLVolume
	} else if !param.IsOmitted(u.OfDbaasPostgreSQLPublicNetwork) {
		return u.OfDbaasPostgreSQLPublicNetwork
	} else if !param.IsOmitted(u.OfDbaasPostgreSQLCPU) {
		return u.OfDbaasPostgreSQLCPU
	} else if !param.IsOmitted(u.OfDbaasPostgreSQLMemory) {
		return u.OfDbaasPostgreSQLMemory
	} else if !param.IsOmitted(u.OfDbaasPostgreSQLConnectionPooler) {
		return u.OfDbaasPostgreSQLConnectionPooler
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) GetField() *string {
	if vt := u.OfSnapshot; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfInstance; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfAICluster; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfAIVirtualCluster; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfBasicVm; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfBaremetal; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfVolume; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfFileShare; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfFloatingip; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEgressTraffic; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfLoadBalancer; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfExternalIP; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfBackup; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfLogIndex; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfFunctions; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfFunctionsCalls; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfFunctionsTraffic; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfContainers; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfInference; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfDbaasPostgreSQLVolume; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfDbaasPostgreSQLPublicNetwork; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfDbaasPostgreSQLCPU; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfDbaasPostgreSQLMemory; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfDbaasPostgreSQLConnectionPooler; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) GetType() *string {
	if vt := u.OfSnapshot; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInstance; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAICluster; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAIVirtualCluster; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBasicVm; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBaremetal; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVolume; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileShare; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFloatingip; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEgressTraffic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLoadBalancer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExternalIP; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBackup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLogIndex; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctions; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctionsCalls; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctionsTraffic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContainers; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInference; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDbaasPostgreSQLVolume; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDbaasPostgreSQLPublicNetwork; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDbaasPostgreSQLCPU; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDbaasPostgreSQLMemory; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDbaasPostgreSQLConnectionPooler; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Values property, if present.
func (u CostReportGetAggregatedMonthlyParamsSchemaFilterUnion) GetValues() []string {
	if vt := u.OfSnapshot; vt != nil {
		return vt.Values
	} else if vt := u.OfInstance; vt != nil {
		return vt.Values
	} else if vt := u.OfAICluster; vt != nil {
		return vt.Values
	} else if vt := u.OfAIVirtualCluster; vt != nil {
		return vt.Values
	} else if vt := u.OfBasicVm; vt != nil {
		return vt.Values
	} else if vt := u.OfBaremetal; vt != nil {
		return vt.Values
	} else if vt := u.OfVolume; vt != nil {
		return vt.Values
	} else if vt := u.OfFileShare; vt != nil {
		return vt.Values
	} else if vt := u.OfImage; vt != nil {
		return vt.Values
	} else if vt := u.OfFloatingip; vt != nil {
		return vt.Values
	} else if vt := u.OfEgressTraffic; vt != nil {
		return vt.Values
	} else if vt := u.OfLoadBalancer; vt != nil {
		return vt.Values
	} else if vt := u.OfExternalIP; vt != nil {
		return vt.Values
	} else if vt := u.OfBackup; vt != nil {
		return vt.Values
	} else if vt := u.OfLogIndex; vt != nil {
		return vt.Values
	} else if vt := u.OfFunctions; vt != nil {
		return vt.Values
	} else if vt := u.OfFunctionsCalls; vt != nil {
		return vt.Values
	} else if vt := u.OfFunctionsTraffic; vt != nil {
		return vt.Values
	} else if vt := u.OfContainers; vt != nil {
		return vt.Values
	} else if vt := u.OfInference; vt != nil {
		return vt.Values
	} else if vt := u.OfDbaasPostgreSQLVolume; vt != nil {
		return vt.Values
	} else if vt := u.OfDbaasPostgreSQLPublicNetwork; vt != nil {
		return vt.Values
	} else if vt := u.OfDbaasPostgreSQLCPU; vt != nil {
		return vt.Values
	} else if vt := u.OfDbaasPostgreSQLMemory; vt != nil {
		return vt.Values
	} else if vt := u.OfDbaasPostgreSQLConnectionPooler; vt != nil {
		return vt.Values
	}
	return nil
}

func init() {
	apijson.RegisterUnion[CostReportGetAggregatedMonthlyParamsSchemaFilterUnion](
		"type",
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot]("snapshot"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterInstance]("instance"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster]("ai_cluster"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster]("ai_virtual_cluster"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm]("basic_vm"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal]("baremetal"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterVolume]("volume"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare]("file_share"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterImage]("image"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip]("floatingip"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic]("egress_traffic"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer]("load_balancer"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP]("external_ip"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterBackup]("backup"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex]("log_index"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions]("functions"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls]("functions_calls"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic]("functions_traffic"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterContainers]("containers"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterInference]("inference"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume]("dbaas_postgresql_volume"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork]("dbaas_postgresql_public_network"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU]("dbaas_postgresql_cpu"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory]("dbaas_postgresql_memory"),
		apijson.Discriminator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler]("dbaas_postgresql_connection_pooler"),
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "source_volume_uuid", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "snapshot".
	Type constant.Snapshot `json:"type" default:"snapshot"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterSnapshot](
		"field", "last_name", "last_size", "source_volume_uuid", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterInstance struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "instance".
	Type constant.Instance `json:"type" default:"instance"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterInstance) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterInstance](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "ai_cluster".
	Type constant.AICluster `json:"type" default:"ai_cluster"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterAICluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ai_virtual_cluster".
	Type constant.AIVirtualCluster `json:"type" default:"ai_virtual_cluster"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterAIVirtualCluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "basic_vm".
	Type constant.BasicVm `json:"type" default:"basic_vm"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterBasicVm](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "baremetal".
	Type constant.Baremetal `json:"type" default:"baremetal"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterBaremetal](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterVolume struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "last_name", "last_size", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "volume".
	Type constant.Volume `json:"type" default:"volume"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterVolume](
		"field", "attached_to_vm", "last_name", "last_size", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare struct {
	// Field name to filter by
	//
	// Any of "file_share_type", "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "file_share".
	Type constant.FileShare `json:"type" default:"file_share"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterFileShare](
		"field", "file_share_type", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterImage struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type" default:"image"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterImage) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterImage](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip struct {
	// Field name to filter by
	//
	// Any of "ip_address", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "floatingip".
	Type constant.Floatingip `json:"type" default:"floatingip"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterFloatingip](
		"field", "ip_address", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic struct {
	// Field name to filter by
	//
	// Any of "instance_name", "instance_type", "port_id", "type", "vm_id".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "egress_traffic".
	Type constant.EgressTraffic `json:"type" default:"egress_traffic"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterEgressTraffic](
		"field", "instance_name", "instance_type", "port_id", "type", "vm_id",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "load_balancer".
	Type constant.LoadBalancer `json:"type" default:"load_balancer"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterLoadBalancer](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id",
	// "type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "external_ip".
	Type constant.ExternalIP `json:"type" default:"external_ip"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterExternalIP](
		"field", "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id", "type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterBackup struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "schedule_id", "source_volume_uuid", "type",
	// "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "backup".
	Type constant.Backup `json:"type" default:"backup"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterBackup) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterBackup](
		"field", "last_name", "last_size", "schedule_id", "source_volume_uuid", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "log_index".
	Type constant.LogIndex `json:"type" default:"log_index"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterLogIndex](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "functions".
	Type constant.Functions `json:"type" default:"functions"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctions](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "functions_calls".
	Type constant.FunctionsCalls `json:"type" default:"functions_calls"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsCalls](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "functions_traffic".
	Type constant.FunctionsTraffic `json:"type" default:"functions_traffic"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterFunctionsTraffic](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterContainers struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "containers".
	Type constant.Containers `json:"type" default:"containers"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterContainers) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterContainers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterContainers](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterInference struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "inference".
	Type constant.Inference `json:"type" default:"inference"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterInference) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterInference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterInference](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid", "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_volume".
	Type constant.DbaasPostgreSQLVolume `json:"type" default:"dbaas_postgresql_volume"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLVolume](
		"field", "last_name", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_public_network".
	Type constant.DbaasPostgreSQLPublicNetwork `json:"type" default:"dbaas_postgresql_public_network"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLPublicNetwork](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_cpu".
	Type constant.DbaasPostgreSQLCPU `json:"type" default:"dbaas_postgresql_cpu"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLCPU](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_memory".
	Type constant.DbaasPostgreSQLMemory `json:"type" default:"dbaas_postgresql_memory"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLMemory](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_connection_pooler".
	Type constant.DbaasPostgreSQLConnectionPooler `json:"type" default:"dbaas_postgresql_connection_pooler"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsSchemaFilterDbaasPostgreSQLConnectionPooler](
		"field", "last_name", "type", "uuid",
	)
}

// Filter by tags
//
// The property Conditions is required.
type CostReportGetAggregatedMonthlyParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []CostReportGetAggregatedMonthlyParamsTagsCondition `json:"conditions,omitzero" api:"required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	//
	// Any of "AND", "OR".
	ConditionType string `json:"condition_type,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedMonthlyParamsTags](
		"condition_type", "AND", "OR",
	)
}

type CostReportGetAggregatedMonthlyParamsTagsCondition struct {
	// The name of the tag to filter (e.g., 'os_version').
	Key param.Opt[string] `json:"key,omitzero"`
	// Determines how strictly the tag value must match the specified value. If true,
	// the tag value must exactly match the given value. If false, a less strict match
	// (e.g., partial or case-insensitive match) may be applied.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The value of the tag to filter (e.g., '22.04').
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedMonthlyParamsTagsCondition) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedMonthlyParamsTagsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedMonthlyParamsTagsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportGetBreakdownParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day in UTC.
	TimeFrom time.Time `json:"time_from" api:"required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day in UTC.
	TimeTo time.Time `json:"time_to" api:"required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Opt[bool] `json:"enable_last_day,omitzero"`
	// Dimension the daily rows are summed down to. Only one grouping per request.
	//
	// Any of "flavor", "project", "region", "type".
	GroupBy CostReportGetBreakdownParamsGroupBy `json:"group_by,omitzero"`
	// List of project IDs. Omit for all accessible projects.
	ProjectIDs []int64 `json:"project_ids,omitzero"`
	// List of region IDs. Omit for all regions and global usage; an empty list is
	// rejected. Use 0 for global usage that has no region.
	RegionIDs []int64 `json:"region_ids,omitzero"`
	// Format of the response: json (default) or `csv_records`.
	//
	// Any of "csv_records", "json".
	ResponseFormat CostReportGetBreakdownParamsResponseFormat `json:"response_format,omitzero"`
	// Filter by tags
	Tags CostReportGetBreakdownParamsTags `json:"tags,omitzero"`
	// List of resource types. Omit for all types.
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "inference_public_models", "instance", "load_balancer",
	// "log_index", "snapshot", "volume".
	Types []string `json:"types,omitzero"`
	paramObj
}

func (r CostReportGetBreakdownParams) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetBreakdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetBreakdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dimension the daily rows are summed down to. Only one grouping per request.
type CostReportGetBreakdownParamsGroupBy string

const (
	CostReportGetBreakdownParamsGroupByFlavor  CostReportGetBreakdownParamsGroupBy = "flavor"
	CostReportGetBreakdownParamsGroupByProject CostReportGetBreakdownParamsGroupBy = "project"
	CostReportGetBreakdownParamsGroupByRegion  CostReportGetBreakdownParamsGroupBy = "region"
	CostReportGetBreakdownParamsGroupByType    CostReportGetBreakdownParamsGroupBy = "type"
)

// Format of the response: json (default) or `csv_records`.
type CostReportGetBreakdownParamsResponseFormat string

const (
	CostReportGetBreakdownParamsResponseFormatCsvRecords CostReportGetBreakdownParamsResponseFormat = "csv_records"
	CostReportGetBreakdownParamsResponseFormatJson       CostReportGetBreakdownParamsResponseFormat = "json"
)

// Filter by tags
//
// The property Conditions is required.
type CostReportGetBreakdownParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []CostReportGetBreakdownParamsTagsCondition `json:"conditions,omitzero" api:"required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	//
	// Any of "AND", "OR".
	ConditionType string `json:"condition_type,omitzero"`
	paramObj
}

func (r CostReportGetBreakdownParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetBreakdownParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetBreakdownParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetBreakdownParamsTags](
		"condition_type", "AND", "OR",
	)
}

type CostReportGetBreakdownParamsTagsCondition struct {
	// The name of the tag to filter (e.g., 'os_version').
	Key param.Opt[string] `json:"key,omitzero"`
	// Determines how strictly the tag value must match the specified value. If true,
	// the tag value must exactly match the given value. If false, a less strict match
	// (e.g., partial or case-insensitive match) may be applied.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The value of the tag to filter (e.g., '22.04').
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r CostReportGetBreakdownParamsTagsCondition) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetBreakdownParamsTagsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetBreakdownParamsTagsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportGetDetailedParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day in UTC.
	TimeFrom time.Time `json:"time_from" api:"required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day in UTC.
	TimeTo time.Time `json:"time_to" api:"required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Opt[bool] `json:"enable_last_day,omitzero"`
	// The response resources limit. Defaults to 10. Maximum 1000 for the json format,
	// 10000 for the `csv_records` format.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The response resources offset.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// Partial, case-insensitive search over region names, project names, resource
	// names, flavors and tags.
	Search param.Opt[string] `json:"search,omitzero"`
	// List of sorting criteria in 'field.direction' format.
	//
	// Any of "cost.asc", "cost.desc", "first_seen.asc", "first_seen.desc",
	// "flavor.asc", "flavor.desc", "last_seen.asc", "last_seen.desc",
	// "project_id.asc", "project_id.desc", "resource_name.asc", "resource_name.desc",
	// "usage.asc", "usage.desc".
	OrderBy []string `json:"order_by,omitzero"`
	// List of project IDs. Omit for all accessible projects; an empty list is
	// rejected.
	ProjectIDs []int64 `json:"project_ids,omitzero"`
	// List of region IDs. Omit for all regions and global usage; an empty list is
	// rejected. Use 0 for global usage that has no region.
	RegionIDs []int64 `json:"region_ids,omitzero"`
	// Format of the response: json (default) or `csv_records`.
	//
	// Any of "csv_records", "json".
	ResponseFormat CostReportGetDetailedParamsResponseFormat `json:"response_format,omitzero"`
	// Filter by tags
	Tags CostReportGetDetailedParamsTags `json:"tags,omitzero"`
	// List of resource types. Omit for all types; an empty list is rejected.
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "inference_public_models", "instance", "load_balancer",
	// "log_index", "snapshot", "volume".
	Types []string `json:"types,omitzero"`
	paramObj
}

func (r CostReportGetDetailedParams) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format of the response: json (default) or `csv_records`.
type CostReportGetDetailedParamsResponseFormat string

const (
	CostReportGetDetailedParamsResponseFormatCsvRecords CostReportGetDetailedParamsResponseFormat = "csv_records"
	CostReportGetDetailedParamsResponseFormatJson       CostReportGetDetailedParamsResponseFormat = "json"
)

// Filter by tags
//
// The property Conditions is required.
type CostReportGetDetailedParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []CostReportGetDetailedParamsTagsCondition `json:"conditions,omitzero" api:"required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	//
	// Any of "AND", "OR".
	ConditionType string `json:"condition_type,omitzero"`
	paramObj
}

func (r CostReportGetDetailedParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsTags](
		"condition_type", "AND", "OR",
	)
}

type CostReportGetDetailedParamsTagsCondition struct {
	// The name of the tag to filter (e.g., 'os_version').
	Key param.Opt[string] `json:"key,omitzero"`
	// Determines how strictly the tag value must match the specified value. If true,
	// the tag value must exactly match the given value. If false, a less strict match
	// (e.g., partial or case-insensitive match) may be applied.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The value of the tag to filter (e.g., '22.04').
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r CostReportGetDetailedParamsTagsCondition) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsTagsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsTagsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
