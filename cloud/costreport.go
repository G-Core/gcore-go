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

// Get cost report totals (aggregated costs) for a given period. Requested period
// should not exceed 31 days.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/`reservation_cost_report`/totals, as the results from this report will not
// be accurate.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CostReportService) GetAggregated(ctx context.Context, body CostReportGetAggregatedParams, opts ...option.RequestOption) (res *CostReportAggregated, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/cost_report/totals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a detailed cost report totals for a specified month, which includes
// both commit and pay-as-you-go (overcommit) prices. Additionally, it provides the
// spent billing units (e.g., hours or GB) for resources. The "`time_to`" parameter
// represents all days in the specified month.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CostReportService) GetAggregatedMonthly(ctx context.Context, body CostReportGetAggregatedMonthlyParams, opts ...option.RequestOption) (res *CostReportAggregatedMonthly, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/reservation_cost_report/totals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a detailed cost report for a given period and specific resources. Requested
// period should not exceed 31 days.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/`reservation_cost_report`/totals, as the results from this report will not
// be accurate.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CostReportService) GetDetailed(ctx context.Context, body CostReportGetDetailedParams, opts ...option.RequestOption) (res *CostReportDetailed, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/cost_report/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CostReportAggregated struct {
	// Count of returned totals
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportAggregatedPriceStatus   `json:"price_status,required"`
	Results     []CostReportAggregatedResultUnion `json:"results,required"`
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
func (r CostReportAggregated) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI, type: string
type CostReportAggregatedPriceStatus string

const (
	CostReportAggregatedPriceStatusError CostReportAggregatedPriceStatus = "error"
	CostReportAggregatedPriceStatusHide  CostReportAggregatedPriceStatus = "hide"
	CostReportAggregatedPriceStatusShow  CostReportAggregatedPriceStatus = "show"
)

// CostReportAggregatedResultUnion contains all possible properties and values from
// [CostReportAggregatedResultAICluster],
// [CostReportAggregatedResultAIVirtualCluster],
// [CostReportAggregatedResultBaremetal], [CostReportAggregatedResultBasicVm],
// [CostReportAggregatedResultBackup], [CostReportAggregatedResultContainers],
// [CostReportAggregatedResultEgressTraffic],
// [CostReportAggregatedResultExternalIP], [CostReportAggregatedResultFileShare],
// [CostReportAggregatedResultFloatingip], [CostReportAggregatedResultFunctions],
// [CostReportAggregatedResultFunctionsCalls],
// [CostReportAggregatedResultFunctionsTraffic], [CostReportAggregatedResultImage],
// [CostReportAggregatedResultInference], [CostReportAggregatedResultInstance],
// [CostReportAggregatedResultLoadBalancer], [CostReportAggregatedResultLogIndex],
// [CostReportAggregatedResultSnapshot], [CostReportAggregatedResultVolume],
// [CostReportAggregatedResultDbaasPostgreSQLConnectionPooler],
// [CostReportAggregatedResultDbaasPostgreSQLMemory],
// [CostReportAggregatedResultDbaasPostgreSQLPublicNetwork],
// [CostReportAggregatedResultDbaasPostgreSQLCPU],
// [CostReportAggregatedResultDbaasPostgreSQLVolume].
//
// Use the [CostReportAggregatedResultUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CostReportAggregatedResultUnion struct {
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
	// This field is from variant [CostReportAggregatedResultBackup].
	LastSize int64 `json:"last_size"`
	// This field is from variant [CostReportAggregatedResultEgressTraffic].
	InstanceType string `json:"instance_type"`
	// This field is from variant [CostReportAggregatedResultFileShare].
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

// anyCostReportAggregatedResult is implemented by each variant of
// [CostReportAggregatedResultUnion] to add type safety for the return type of
// [CostReportAggregatedResultUnion.AsAny]
type anyCostReportAggregatedResult interface {
	implCostReportAggregatedResultUnion()
}

func (CostReportAggregatedResultAICluster) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultAIVirtualCluster) implCostReportAggregatedResultUnion() {}
func (CostReportAggregatedResultBaremetal) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultBasicVm) implCostReportAggregatedResultUnion()          {}
func (CostReportAggregatedResultBackup) implCostReportAggregatedResultUnion()           {}
func (CostReportAggregatedResultContainers) implCostReportAggregatedResultUnion()       {}
func (CostReportAggregatedResultEgressTraffic) implCostReportAggregatedResultUnion()    {}
func (CostReportAggregatedResultExternalIP) implCostReportAggregatedResultUnion()       {}
func (CostReportAggregatedResultFileShare) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultFloatingip) implCostReportAggregatedResultUnion()       {}
func (CostReportAggregatedResultFunctions) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultFunctionsCalls) implCostReportAggregatedResultUnion()   {}
func (CostReportAggregatedResultFunctionsTraffic) implCostReportAggregatedResultUnion() {}
func (CostReportAggregatedResultImage) implCostReportAggregatedResultUnion()            {}
func (CostReportAggregatedResultInference) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultInstance) implCostReportAggregatedResultUnion()         {}
func (CostReportAggregatedResultLoadBalancer) implCostReportAggregatedResultUnion()     {}
func (CostReportAggregatedResultLogIndex) implCostReportAggregatedResultUnion()         {}
func (CostReportAggregatedResultSnapshot) implCostReportAggregatedResultUnion()         {}
func (CostReportAggregatedResultVolume) implCostReportAggregatedResultUnion()           {}
func (CostReportAggregatedResultDbaasPostgreSQLConnectionPooler) implCostReportAggregatedResultUnion() {
}
func (CostReportAggregatedResultDbaasPostgreSQLMemory) implCostReportAggregatedResultUnion()        {}
func (CostReportAggregatedResultDbaasPostgreSQLPublicNetwork) implCostReportAggregatedResultUnion() {}
func (CostReportAggregatedResultDbaasPostgreSQLCPU) implCostReportAggregatedResultUnion()           {}
func (CostReportAggregatedResultDbaasPostgreSQLVolume) implCostReportAggregatedResultUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CostReportAggregatedResultUnion.AsAny().(type) {
//	case cloud.CostReportAggregatedResultAICluster:
//	case cloud.CostReportAggregatedResultAIVirtualCluster:
//	case cloud.CostReportAggregatedResultBaremetal:
//	case cloud.CostReportAggregatedResultBasicVm:
//	case cloud.CostReportAggregatedResultBackup:
//	case cloud.CostReportAggregatedResultContainers:
//	case cloud.CostReportAggregatedResultEgressTraffic:
//	case cloud.CostReportAggregatedResultExternalIP:
//	case cloud.CostReportAggregatedResultFileShare:
//	case cloud.CostReportAggregatedResultFloatingip:
//	case cloud.CostReportAggregatedResultFunctions:
//	case cloud.CostReportAggregatedResultFunctionsCalls:
//	case cloud.CostReportAggregatedResultFunctionsTraffic:
//	case cloud.CostReportAggregatedResultImage:
//	case cloud.CostReportAggregatedResultInference:
//	case cloud.CostReportAggregatedResultInstance:
//	case cloud.CostReportAggregatedResultLoadBalancer:
//	case cloud.CostReportAggregatedResultLogIndex:
//	case cloud.CostReportAggregatedResultSnapshot:
//	case cloud.CostReportAggregatedResultVolume:
//	case cloud.CostReportAggregatedResultDbaasPostgreSQLConnectionPooler:
//	case cloud.CostReportAggregatedResultDbaasPostgreSQLMemory:
//	case cloud.CostReportAggregatedResultDbaasPostgreSQLPublicNetwork:
//	case cloud.CostReportAggregatedResultDbaasPostgreSQLCPU:
//	case cloud.CostReportAggregatedResultDbaasPostgreSQLVolume:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CostReportAggregatedResultUnion) AsAny() anyCostReportAggregatedResult {
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

func (u CostReportAggregatedResultUnion) AsAICluster() (v CostReportAggregatedResultAICluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsAIVirtualCluster() (v CostReportAggregatedResultAIVirtualCluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsBaremetal() (v CostReportAggregatedResultBaremetal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsBasicVm() (v CostReportAggregatedResultBasicVm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsBackup() (v CostReportAggregatedResultBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsContainers() (v CostReportAggregatedResultContainers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsEgressTraffic() (v CostReportAggregatedResultEgressTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsExternalIP() (v CostReportAggregatedResultExternalIP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsFileShare() (v CostReportAggregatedResultFileShare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsFloatingip() (v CostReportAggregatedResultFloatingip) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsFunctions() (v CostReportAggregatedResultFunctions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsFunctionsCalls() (v CostReportAggregatedResultFunctionsCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsFunctionsTraffic() (v CostReportAggregatedResultFunctionsTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsImage() (v CostReportAggregatedResultImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsInference() (v CostReportAggregatedResultInference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsInstance() (v CostReportAggregatedResultInstance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsLoadBalancer() (v CostReportAggregatedResultLoadBalancer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsLogIndex() (v CostReportAggregatedResultLogIndex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsSnapshot() (v CostReportAggregatedResultSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsVolume() (v CostReportAggregatedResultVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsDbaasPostgreSQLConnectionPooler() (v CostReportAggregatedResultDbaasPostgreSQLConnectionPooler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsDbaasPostgreSQLMemory() (v CostReportAggregatedResultDbaasPostgreSQLMemory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsDbaasPostgreSQLPublicNetwork() (v CostReportAggregatedResultDbaasPostgreSQLPublicNetwork) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsDbaasPostgreSQLCPU() (v CostReportAggregatedResultDbaasPostgreSQLCPU) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportAggregatedResultUnion) AsDbaasPostgreSQLVolume() (v CostReportAggregatedResultDbaasPostgreSQLVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CostReportAggregatedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *CostReportAggregatedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultAICluster struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.AICluster `json:"type,required"`
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
func (r CostReportAggregatedResultAICluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultAIVirtualCluster struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     constant.AIVirtualCluster `json:"type,required"`
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
func (r CostReportAggregatedResultAIVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultBaremetal struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Baremetal `json:"type,required"`
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
func (r CostReportAggregatedResultBaremetal) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultBasicVm struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64            `json:"region_id,required"`
	Type     constant.BasicVm `json:"type,required"`
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
func (r CostReportAggregatedResultBasicVm) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultBackup struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64           `json:"region_id,required"`
	Type     constant.Backup `json:"type,required"`
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
func (r CostReportAggregatedResultBackup) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultContainers struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.Containers `json:"type,required"`
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
func (r CostReportAggregatedResultContainers) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultEgressTraffic struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Type of the instance
	//
	// Any of "baremetal", "vm".
	InstanceType string `json:"instance_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                  `json:"region_id,required"`
	Type     constant.EgressTraffic `json:"type,required"`
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
func (r CostReportAggregatedResultEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultExternalIP struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.ExternalIP `json:"type,required"`
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
func (r CostReportAggregatedResultExternalIP) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultFileShare struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.FileShare `json:"type,required"`
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
func (r CostReportAggregatedResultFileShare) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultFloatingip struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.Floatingip `json:"type,required"`
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
func (r CostReportAggregatedResultFloatingip) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultFunctions struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Functions `json:"type,required"`
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
func (r CostReportAggregatedResultFunctions) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultFunctionsCalls struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                   `json:"region_id,required"`
	Type     constant.FunctionsCalls `json:"type,required"`
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
func (r CostReportAggregatedResultFunctionsCalls) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultFunctionsTraffic struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     constant.FunctionsTraffic `json:"type,required"`
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
func (r CostReportAggregatedResultFunctionsTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultImage struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64          `json:"region_id,required"`
	Type     constant.Image `json:"type,required"`
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
func (r CostReportAggregatedResultImage) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultInference struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Inference `json:"type,required"`
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
func (r CostReportAggregatedResultInference) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultInstance struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.Instance `json:"type,required"`
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
func (r CostReportAggregatedResultInstance) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultLoadBalancer struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                 `json:"region_id,required"`
	Type     constant.LoadBalancer `json:"type,required"`
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
func (r CostReportAggregatedResultLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultLogIndex struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.LogIndex `json:"type,required"`
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
func (r CostReportAggregatedResultLogIndex) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultSnapshot struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.Snapshot `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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
func (r CostReportAggregatedResultSnapshot) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultVolume struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64           `json:"region_id,required"`
	Type     constant.Volume `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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
func (r CostReportAggregatedResultVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultDbaasPostgreSQLConnectionPooler struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                    `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
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
func (r CostReportAggregatedResultDbaasPostgreSQLConnectionPooler) RawJSON() string {
	return r.JSON.raw
}
func (r *CostReportAggregatedResultDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultDbaasPostgreSQLMemory struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                          `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type,required"`
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
func (r CostReportAggregatedResultDbaasPostgreSQLMemory) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultDbaasPostgreSQLPublicNetwork struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                 `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
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
func (r CostReportAggregatedResultDbaasPostgreSQLPublicNetwork) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultDbaasPostgreSQLCPU struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                       `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type,required"`
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
func (r CostReportAggregatedResultDbaasPostgreSQLCPU) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedResultDbaasPostgreSQLVolume struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                          `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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
func (r CostReportAggregatedResultDbaasPostgreSQLVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportAggregatedResultDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportAggregatedMonthly struct {
	// Total count of the totals
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportAggregatedMonthlyPriceStatus   `json:"price_status,required"`
	Results     []CostReportAggregatedMonthlyResultUnion `json:"results,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.AICluster `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     constant.AIVirtualCluster `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Baremetal `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64            `json:"region_id,required"`
	Type     constant.BasicVm `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64           `json:"region_id,required"`
	Type     constant.Backup `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.Containers `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Type of the instance
	//
	// Any of "baremetal", "vm".
	InstanceType string `json:"instance_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                  `json:"region_id,required"`
	Type     constant.EgressTraffic `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.ExternalIP `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.FileShare `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.Floatingip `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Functions `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                   `json:"region_id,required"`
	Type     constant.FunctionsCalls `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     constant.FunctionsTraffic `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64          `json:"region_id,required"`
	Type     constant.Image `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Inference `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.Instance `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                 `json:"region_id,required"`
	Type     constant.LoadBalancer `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.LogIndex `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64             `json:"region_id,required"`
	Type     constant.Snapshot `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64           `json:"region_id,required"`
	Type     constant.Volume `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                    `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                          `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                 `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                       `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type,required"`
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
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                          `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
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

type CostReportDetailed struct {
	// Count of all the resources
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	//
	// Any of "error", "hide", "show".
	PriceStatus CostReportDetailedPriceStatus   `json:"price_status,required"`
	Results     []CostReportDetailedResultUnion `json:"results,required"`
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

// Price status for the UI, type: string
type CostReportDetailedPriceStatus string

const (
	CostReportDetailedPriceStatusError CostReportDetailedPriceStatus = "error"
	CostReportDetailedPriceStatusHide  CostReportDetailedPriceStatus = "hide"
	CostReportDetailedPriceStatusShow  CostReportDetailedPriceStatus = "show"
)

// CostReportDetailedResultUnion contains all possible properties and values from
// [CostReportDetailedResultAICluster], [CostReportDetailedResultAIVirtualCluster],
// [CostReportDetailedResultBaremetal], [CostReportDetailedResultBasicVm],
// [CostReportDetailedResultBackup], [CostReportDetailedResultContainers],
// [CostReportDetailedResultEgressTraffic], [CostReportDetailedResultExternalIP],
// [CostReportDetailedResultFileShare], [CostReportDetailedResultFloatingip],
// [CostReportDetailedResultFunctions], [CostReportDetailedResultFunctionsCalls],
// [CostReportDetailedResultFunctionsTraffic], [CostReportDetailedResultImage],
// [CostReportDetailedResultInference], [CostReportDetailedResultInstance],
// [CostReportDetailedResultLoadBalancer], [CostReportDetailedResultLogIndex],
// [CostReportDetailedResultSnapshot], [CostReportDetailedResultVolume],
// [CostReportDetailedResultDbaasPostgreSQLConnectionPooler],
// [CostReportDetailedResultDbaasPostgreSQLMemory],
// [CostReportDetailedResultDbaasPostgreSQLPublicNetwork],
// [CostReportDetailedResultDbaasPostgreSQLCPU],
// [CostReportDetailedResultDbaasPostgreSQLVolume].
//
// Use the [CostReportDetailedResultUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CostReportDetailedResultUnion struct {
	BillingFeatureName string              `json:"billing_feature_name"`
	BillingMetricName  string              `json:"billing_metric_name"`
	BillingValue       float64             `json:"billing_value"`
	BillingValueUnit   string              `json:"billing_value_unit"`
	Cost               float64             `json:"cost"`
	Currency           string              `json:"currency"`
	Err                string              `json:"err"`
	FirstSeen          time.Time           `json:"first_seen"`
	Flavor             string              `json:"flavor"`
	LastName           string              `json:"last_name"`
	LastSeen           time.Time           `json:"last_seen"`
	ProjectID          int64               `json:"project_id"`
	Region             int64               `json:"region"`
	RegionID           int64               `json:"region_id"`
	Tags               []map[string]string `json:"tags"`
	// Any of "ai_cluster", "ai_virtual_cluster", "baremetal", "basic_vm", "backup",
	// "containers", "egress_traffic", "external_ip", "file_share", "floatingip",
	// "functions", "functions_calls", "functions_traffic", "image", "inference",
	// "instance", "load_balancer", "log_index", "snapshot", "volume",
	// "dbaas_postgresql_connection_pooler", "dbaas_postgresql_memory",
	// "dbaas_postgresql_public_network", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_volume".
	Type     string `json:"type"`
	Uuid     string `json:"uuid"`
	LastSize int64  `json:"last_size"`
	// This field is from variant [CostReportDetailedResultBackup].
	ScheduleID       string `json:"schedule_id"`
	SourceVolumeUuid string `json:"source_volume_uuid"`
	// This field is from variant [CostReportDetailedResultEgressTraffic].
	InstanceType string `json:"instance_type"`
	PortID       string `json:"port_id"`
	SizeUnit     string `json:"size_unit"`
	// This field is from variant [CostReportDetailedResultEgressTraffic].
	VmID string `json:"vm_id"`
	// This field is from variant [CostReportDetailedResultEgressTraffic].
	InstanceName string `json:"instance_name"`
	AttachedToVm string `json:"attached_to_vm"`
	IPAddress    string `json:"ip_address"`
	// This field is from variant [CostReportDetailedResultExternalIP].
	NetworkID string `json:"network_id"`
	// This field is from variant [CostReportDetailedResultExternalIP].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [CostReportDetailedResultFileShare].
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
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		LastSize           respjson.Field
		ScheduleID         respjson.Field
		SourceVolumeUuid   respjson.Field
		InstanceType       respjson.Field
		PortID             respjson.Field
		SizeUnit           respjson.Field
		VmID               respjson.Field
		InstanceName       respjson.Field
		AttachedToVm       respjson.Field
		IPAddress          respjson.Field
		NetworkID          respjson.Field
		SubnetID           respjson.Field
		FileShareType      respjson.Field
		VolumeType         respjson.Field
		raw                string
	} `json:"-"`
}

// anyCostReportDetailedResult is implemented by each variant of
// [CostReportDetailedResultUnion] to add type safety for the return type of
// [CostReportDetailedResultUnion.AsAny]
type anyCostReportDetailedResult interface {
	implCostReportDetailedResultUnion()
}

func (CostReportDetailedResultAICluster) implCostReportDetailedResultUnion()                       {}
func (CostReportDetailedResultAIVirtualCluster) implCostReportDetailedResultUnion()                {}
func (CostReportDetailedResultBaremetal) implCostReportDetailedResultUnion()                       {}
func (CostReportDetailedResultBasicVm) implCostReportDetailedResultUnion()                         {}
func (CostReportDetailedResultBackup) implCostReportDetailedResultUnion()                          {}
func (CostReportDetailedResultContainers) implCostReportDetailedResultUnion()                      {}
func (CostReportDetailedResultEgressTraffic) implCostReportDetailedResultUnion()                   {}
func (CostReportDetailedResultExternalIP) implCostReportDetailedResultUnion()                      {}
func (CostReportDetailedResultFileShare) implCostReportDetailedResultUnion()                       {}
func (CostReportDetailedResultFloatingip) implCostReportDetailedResultUnion()                      {}
func (CostReportDetailedResultFunctions) implCostReportDetailedResultUnion()                       {}
func (CostReportDetailedResultFunctionsCalls) implCostReportDetailedResultUnion()                  {}
func (CostReportDetailedResultFunctionsTraffic) implCostReportDetailedResultUnion()                {}
func (CostReportDetailedResultImage) implCostReportDetailedResultUnion()                           {}
func (CostReportDetailedResultInference) implCostReportDetailedResultUnion()                       {}
func (CostReportDetailedResultInstance) implCostReportDetailedResultUnion()                        {}
func (CostReportDetailedResultLoadBalancer) implCostReportDetailedResultUnion()                    {}
func (CostReportDetailedResultLogIndex) implCostReportDetailedResultUnion()                        {}
func (CostReportDetailedResultSnapshot) implCostReportDetailedResultUnion()                        {}
func (CostReportDetailedResultVolume) implCostReportDetailedResultUnion()                          {}
func (CostReportDetailedResultDbaasPostgreSQLConnectionPooler) implCostReportDetailedResultUnion() {}
func (CostReportDetailedResultDbaasPostgreSQLMemory) implCostReportDetailedResultUnion()           {}
func (CostReportDetailedResultDbaasPostgreSQLPublicNetwork) implCostReportDetailedResultUnion()    {}
func (CostReportDetailedResultDbaasPostgreSQLCPU) implCostReportDetailedResultUnion()              {}
func (CostReportDetailedResultDbaasPostgreSQLVolume) implCostReportDetailedResultUnion()           {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CostReportDetailedResultUnion.AsAny().(type) {
//	case cloud.CostReportDetailedResultAICluster:
//	case cloud.CostReportDetailedResultAIVirtualCluster:
//	case cloud.CostReportDetailedResultBaremetal:
//	case cloud.CostReportDetailedResultBasicVm:
//	case cloud.CostReportDetailedResultBackup:
//	case cloud.CostReportDetailedResultContainers:
//	case cloud.CostReportDetailedResultEgressTraffic:
//	case cloud.CostReportDetailedResultExternalIP:
//	case cloud.CostReportDetailedResultFileShare:
//	case cloud.CostReportDetailedResultFloatingip:
//	case cloud.CostReportDetailedResultFunctions:
//	case cloud.CostReportDetailedResultFunctionsCalls:
//	case cloud.CostReportDetailedResultFunctionsTraffic:
//	case cloud.CostReportDetailedResultImage:
//	case cloud.CostReportDetailedResultInference:
//	case cloud.CostReportDetailedResultInstance:
//	case cloud.CostReportDetailedResultLoadBalancer:
//	case cloud.CostReportDetailedResultLogIndex:
//	case cloud.CostReportDetailedResultSnapshot:
//	case cloud.CostReportDetailedResultVolume:
//	case cloud.CostReportDetailedResultDbaasPostgreSQLConnectionPooler:
//	case cloud.CostReportDetailedResultDbaasPostgreSQLMemory:
//	case cloud.CostReportDetailedResultDbaasPostgreSQLPublicNetwork:
//	case cloud.CostReportDetailedResultDbaasPostgreSQLCPU:
//	case cloud.CostReportDetailedResultDbaasPostgreSQLVolume:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CostReportDetailedResultUnion) AsAny() anyCostReportDetailedResult {
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

func (u CostReportDetailedResultUnion) AsAICluster() (v CostReportDetailedResultAICluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsAIVirtualCluster() (v CostReportDetailedResultAIVirtualCluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsBaremetal() (v CostReportDetailedResultBaremetal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsBasicVm() (v CostReportDetailedResultBasicVm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsBackup() (v CostReportDetailedResultBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsContainers() (v CostReportDetailedResultContainers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsEgressTraffic() (v CostReportDetailedResultEgressTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsExternalIP() (v CostReportDetailedResultExternalIP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsFileShare() (v CostReportDetailedResultFileShare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsFloatingip() (v CostReportDetailedResultFloatingip) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsFunctions() (v CostReportDetailedResultFunctions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsFunctionsCalls() (v CostReportDetailedResultFunctionsCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsFunctionsTraffic() (v CostReportDetailedResultFunctionsTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsImage() (v CostReportDetailedResultImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsInference() (v CostReportDetailedResultInference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsInstance() (v CostReportDetailedResultInstance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsLoadBalancer() (v CostReportDetailedResultLoadBalancer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsLogIndex() (v CostReportDetailedResultLogIndex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsSnapshot() (v CostReportDetailedResultSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsVolume() (v CostReportDetailedResultVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsDbaasPostgreSQLConnectionPooler() (v CostReportDetailedResultDbaasPostgreSQLConnectionPooler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsDbaasPostgreSQLMemory() (v CostReportDetailedResultDbaasPostgreSQLMemory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsDbaasPostgreSQLPublicNetwork() (v CostReportDetailedResultDbaasPostgreSQLPublicNetwork) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsDbaasPostgreSQLCPU() (v CostReportDetailedResultDbaasPostgreSQLCPU) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CostReportDetailedResultUnion) AsDbaasPostgreSQLVolume() (v CostReportDetailedResultDbaasPostgreSQLVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CostReportDetailedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *CostReportDetailedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultAICluster struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor,required"`
	// Name of the AI cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.AICluster  `json:"type,required"`
	// UUID of the Baremetal GPU cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultAICluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultAIVirtualCluster struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor,required"`
	// Name of the AI cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string       `json:"tags,required"`
	Type constant.AIVirtualCluster `json:"type,required"`
	// UUID of the Virtual GPU cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultAIVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultBaremetal struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Name of the bare metal server
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Baremetal  `json:"type,required"`
	// UUID of the bare metal server
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultBaremetal) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultBasicVm struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Name of the basic VM
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.BasicVm    `json:"type,required"`
	// UUID of the basic VM
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultBasicVm) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultBackup struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the backup
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// ID of the backup schedule
	ScheduleID string `json:"schedule_id,required" format:"uuid"`
	// UUID of the source volume
	SourceVolumeUuid string          `json:"source_volume_uuid,required" format:"uuid"`
	Type             constant.Backup `json:"type,required"`
	// UUID of the backup
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		ScheduleID         respjson.Field
		SourceVolumeUuid   respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultBackup) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultContainers struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the container
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64               `json:"region_id,required"`
	Type     constant.Containers `json:"type,required"`
	// UUID of the container
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultContainers) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultEgressTraffic struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Type of the instance
	//
	// Any of "baremetal", "vm".
	InstanceType string `json:"instance_type,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the port the traffic is associated with
	PortID string `json:"port_id,required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string `json:"size_unit,required"`
	// List of tags
	Tags []map[string]string    `json:"tags,required"`
	Type constant.EgressTraffic `json:"type,required"`
	// ID of the bare metal server the traffic is associated with
	VmID string `json:"vm_id,required" format:"uuid"`
	// Name of the instance
	InstanceName string `json:"instance_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		InstanceType       respjson.Field
		LastSeen           respjson.Field
		PortID             respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		VmID               respjson.Field
		InstanceName       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultExternalIP struct {
	// ID of the VM the IP is attached to
	AttachedToVm       string `json:"attached_to_vm,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the network the IP is attached to
	NetworkID string `json:"network_id,required" format:"uuid"`
	// ID of the port the IP is associated with
	PortID string `json:"port_id,required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// ID of the subnet the IP is attached to
	SubnetID string              `json:"subnet_id,required" format:"uuid"`
	Type     constant.ExternalIP `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedToVm       respjson.Field
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		IPAddress          respjson.Field
		LastSeen           respjson.Field
		NetworkID          respjson.Field
		PortID             respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SubnetID           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultExternalIP) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultFileShare struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the file share
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the file share in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit constant.GiB `json:"size_unit,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.FileShare  `json:"type,required"`
	// UUID of the file share
	Uuid string `json:"uuid,required" format:"uuid"`
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
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultFileShare) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultFloatingip struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Name of the floating IP
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Floatingip `json:"type,required"`
	// UUID of the floating IP
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		IPAddress          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultFloatingip) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultFunctions struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Functions `json:"type,required"`
	// UUID of the function
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultFunctions) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultFunctionsCalls struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function call
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                   `json:"region_id,required"`
	Type     constant.FunctionsCalls `json:"type,required"`
	// UUID of the function call
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultFunctionsCalls) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultFunctionsTraffic struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function egress traffic
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     constant.FunctionsTraffic `json:"type,required"`
	// UUID of the function egress traffic
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultFunctionsTraffic) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultImage struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the image
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the image in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit constant.Bytes `json:"size_unit,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Image      `json:"type,required"`
	// UUID of the image
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultImage) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultInference struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the inference deployment
	Flavor string `json:"flavor,required"`
	// Name of the inference deployment
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64              `json:"region_id,required"`
	Type     constant.Inference `json:"type,required"`
	// UUID of the inference deployment
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultInference) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultInstance struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Name of the instance
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Instance   `json:"type,required"`
	// UUID of the instance
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultInstance) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultLoadBalancer struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Name of the load balancer
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of tags
	Tags []map[string]string   `json:"tags,required"`
	Type constant.LoadBalancer `json:"type,required"`
	// UUID of the load balancer
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		Flavor             respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultLogIndex struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the log index
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the log index in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string            `json:"size_unit,required"`
	Type     constant.LogIndex `json:"type,required"`
	// UUID of the log index
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultLogIndex) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultSnapshot struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the snapshot
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the snapshot in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string `json:"size_unit,required"`
	// UUID of the source volume
	SourceVolumeUuid string `json:"source_volume_uuid,required" format:"uuid"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Snapshot   `json:"type,required"`
	// UUID of the snapshot
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		SourceVolumeUuid   respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		VolumeType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultSnapshot) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultVolume struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the volume
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the volume in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string `json:"size_unit,required"`
	// List of tags
	Tags []map[string]string `json:"tags,required"`
	Type constant.Volume     `json:"type,required"`
	// UUID of the volume
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// ID of the VM the volume is attached to
	AttachedToVm string `json:"attached_to_vm,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		LastSize           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Tags               respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		VolumeType         respjson.Field
		AttachedToVm       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultDbaasPostgreSQLConnectionPooler struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                    `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultDbaasPostgreSQLConnectionPooler) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultDbaasPostgreSQLMemory struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                          `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultDbaasPostgreSQLMemory) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultDbaasPostgreSQLPublicNetwork struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                 `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultDbaasPostgreSQLPublicNetwork) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultDbaasPostgreSQLCPU struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                       `json:"region_id,required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultDbaasPostgreSQLCPU) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportDetailedResultDbaasPostgreSQLVolume struct {
	BillingFeatureName string `json:"billing_feature_name,required"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit,required"`
	// Cost for requested period
	Cost float64 `json:"cost,required"`
	// Currency of the cost
	Currency string `json:"currency,required"`
	// Error message
	Err string `json:"err,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string                         `json:"size_unit,required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingFeatureName respjson.Field
		BillingMetricName  respjson.Field
		BillingValue       respjson.Field
		BillingValueUnit   respjson.Field
		Cost               respjson.Field
		Currency           respjson.Field
		Err                respjson.Field
		FirstSeen          respjson.Field
		LastName           respjson.Field
		LastSeen           respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SizeUnit           respjson.Field
		Type               respjson.Field
		Uuid               respjson.Field
		VolumeType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostReportDetailedResultDbaasPostgreSQLVolume) RawJSON() string { return r.JSON.raw }
func (r *CostReportDetailedResultDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CostReportGetAggregatedParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day.
	TimeFrom time.Time `json:"time_from,required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day.
	TimeTo time.Time `json:"time_to,required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Opt[bool] `json:"enable_last_day,omitzero"`
	// Round cost values to 5 decimal places. When false, returns full precision.
	Rounding param.Opt[bool] `json:"rounding,omitzero"`
	// List of project IDs
	Projects []int64 `json:"projects,omitzero"`
	// List of region IDs.
	Regions []int64 `json:"regions,omitzero"`
	// Format of the response (csv or json).
	//
	// Any of "csv_totals", "json".
	ResponseFormat CostReportGetAggregatedParamsResponseFormat `json:"response_format,omitzero"`
	// Extended filter for field filtering.
	SchemaFilter CostReportGetAggregatedParamsSchemaFilterUnion `json:"schema_filter,omitzero"`
	// Filter by tags
	Tags CostReportGetAggregatedParamsTags `json:"tags,omitzero"`
	// List of resource types to be filtered in the report.
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "instance", "load_balancer", "log_index", "snapshot", "volume".
	Types []string `json:"types,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedParams) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format of the response (csv or json).
type CostReportGetAggregatedParamsResponseFormat string

const (
	CostReportGetAggregatedParamsResponseFormatCsvTotals CostReportGetAggregatedParamsResponseFormat = "csv_totals"
	CostReportGetAggregatedParamsResponseFormatJson      CostReportGetAggregatedParamsResponseFormat = "json"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CostReportGetAggregatedParamsSchemaFilterUnion struct {
	OfSnapshot                        *CostReportGetAggregatedParamsSchemaFilterSnapshot                        `json:",omitzero,inline"`
	OfInstance                        *CostReportGetAggregatedParamsSchemaFilterInstance                        `json:",omitzero,inline"`
	OfAICluster                       *CostReportGetAggregatedParamsSchemaFilterAICluster                       `json:",omitzero,inline"`
	OfAIVirtualCluster                *CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster                `json:",omitzero,inline"`
	OfBasicVm                         *CostReportGetAggregatedParamsSchemaFilterBasicVm                         `json:",omitzero,inline"`
	OfBaremetal                       *CostReportGetAggregatedParamsSchemaFilterBaremetal                       `json:",omitzero,inline"`
	OfVolume                          *CostReportGetAggregatedParamsSchemaFilterVolume                          `json:",omitzero,inline"`
	OfFileShare                       *CostReportGetAggregatedParamsSchemaFilterFileShare                       `json:",omitzero,inline"`
	OfImage                           *CostReportGetAggregatedParamsSchemaFilterImage                           `json:",omitzero,inline"`
	OfFloatingip                      *CostReportGetAggregatedParamsSchemaFilterFloatingip                      `json:",omitzero,inline"`
	OfEgressTraffic                   *CostReportGetAggregatedParamsSchemaFilterEgressTraffic                   `json:",omitzero,inline"`
	OfLoadBalancer                    *CostReportGetAggregatedParamsSchemaFilterLoadBalancer                    `json:",omitzero,inline"`
	OfExternalIP                      *CostReportGetAggregatedParamsSchemaFilterExternalIP                      `json:",omitzero,inline"`
	OfBackup                          *CostReportGetAggregatedParamsSchemaFilterBackup                          `json:",omitzero,inline"`
	OfLogIndex                        *CostReportGetAggregatedParamsSchemaFilterLogIndex                        `json:",omitzero,inline"`
	OfFunctions                       *CostReportGetAggregatedParamsSchemaFilterFunctions                       `json:",omitzero,inline"`
	OfFunctionsCalls                  *CostReportGetAggregatedParamsSchemaFilterFunctionsCalls                  `json:",omitzero,inline"`
	OfFunctionsTraffic                *CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic                `json:",omitzero,inline"`
	OfContainers                      *CostReportGetAggregatedParamsSchemaFilterContainers                      `json:",omitzero,inline"`
	OfInference                       *CostReportGetAggregatedParamsSchemaFilterInference                       `json:",omitzero,inline"`
	OfDbaasPostgreSQLVolume           *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume           `json:",omitzero,inline"`
	OfDbaasPostgreSQLPublicNetwork    *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork    `json:",omitzero,inline"`
	OfDbaasPostgreSQLCPU              *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU              `json:",omitzero,inline"`
	OfDbaasPostgreSQLMemory           *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory           `json:",omitzero,inline"`
	OfDbaasPostgreSQLConnectionPooler *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler `json:",omitzero,inline"`
	paramUnion
}

func (u CostReportGetAggregatedParamsSchemaFilterUnion) MarshalJSON() ([]byte, error) {
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
func (u *CostReportGetAggregatedParamsSchemaFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CostReportGetAggregatedParamsSchemaFilterUnion) asAny() any {
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
func (u CostReportGetAggregatedParamsSchemaFilterUnion) GetField() *string {
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
func (u CostReportGetAggregatedParamsSchemaFilterUnion) GetType() *string {
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
func (u CostReportGetAggregatedParamsSchemaFilterUnion) GetValues() []string {
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
	apijson.RegisterUnion[CostReportGetAggregatedParamsSchemaFilterUnion](
		"type",
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterSnapshot]("snapshot"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterInstance]("instance"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterAICluster]("ai_cluster"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster]("ai_virtual_cluster"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterBasicVm]("basic_vm"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterBaremetal]("baremetal"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterVolume]("volume"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterFileShare]("file_share"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterImage]("image"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterFloatingip]("floatingip"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterEgressTraffic]("egress_traffic"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterLoadBalancer]("load_balancer"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterExternalIP]("external_ip"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterBackup]("backup"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterLogIndex]("log_index"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterFunctions]("functions"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterFunctionsCalls]("functions_calls"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic]("functions_traffic"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterContainers]("containers"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterInference]("inference"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume]("dbaas_postgresql_volume"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork]("dbaas_postgresql_public_network"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU]("dbaas_postgresql_cpu"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory]("dbaas_postgresql_memory"),
		apijson.Discriminator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler]("dbaas_postgresql_connection_pooler"),
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterSnapshot struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "source_volume_uuid", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "snapshot".
	Type constant.Snapshot `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterSnapshot](
		"field", "last_name", "last_size", "source_volume_uuid", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterInstance struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "instance".
	Type constant.Instance `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterInstance) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterInstance](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterAICluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "ai_cluster".
	Type constant.AICluster `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterAICluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterAICluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterAICluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ai_virtual_cluster".
	Type constant.AIVirtualCluster `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterAIVirtualCluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterBasicVm struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "basic_vm".
	Type constant.BasicVm `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterBasicVm) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterBasicVm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterBasicVm](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterBaremetal struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "baremetal".
	Type constant.Baremetal `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterBaremetal) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterBaremetal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterBaremetal](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterVolume struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "last_name", "last_size", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "volume".
	Type constant.Volume `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterVolume](
		"field", "attached_to_vm", "last_name", "last_size", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterFileShare struct {
	// Field name to filter by
	//
	// Any of "file_share_type", "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file_share".
	Type constant.FileShare `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterFileShare) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterFileShare](
		"field", "file_share_type", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterImage struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterImage) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterImage](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterFloatingip struct {
	// Field name to filter by
	//
	// Any of "ip_address", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "floatingip".
	Type constant.Floatingip `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterFloatingip) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterFloatingip
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterFloatingip](
		"field", "ip_address", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterEgressTraffic struct {
	// Field name to filter by
	//
	// Any of "instance_name", "instance_type", "port_id", "type", "vm_id".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "egress_traffic".
	Type constant.EgressTraffic `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterEgressTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterEgressTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterEgressTraffic](
		"field", "instance_name", "instance_type", "port_id", "type", "vm_id",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterLoadBalancer struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "load_balancer".
	Type constant.LoadBalancer `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterLoadBalancer) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterLoadBalancer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterLoadBalancer](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterExternalIP struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id",
	// "type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "external_ip".
	Type constant.ExternalIP `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterExternalIP) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterExternalIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterExternalIP](
		"field", "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id", "type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterBackup struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "schedule_id", "source_volume_uuid", "type",
	// "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "backup".
	Type constant.Backup `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterBackup) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterBackup](
		"field", "last_name", "last_size", "schedule_id", "source_volume_uuid", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterLogIndex struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "log_index".
	Type constant.LogIndex `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterLogIndex) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterLogIndex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterLogIndex](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterFunctions struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions".
	Type constant.Functions `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterFunctions) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterFunctions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterFunctions](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterFunctionsCalls struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions_calls".
	Type constant.FunctionsCalls `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterFunctionsCalls) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterFunctionsCalls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterFunctionsCalls](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "functions_traffic".
	Type constant.FunctionsTraffic `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterFunctionsTraffic](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterContainers struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "containers".
	Type constant.Containers `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterContainers) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterContainers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterContainers](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterInference struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "inference".
	Type constant.Inference `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterInference) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterInference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterInference](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid", "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_volume".
	Type constant.DbaasPostgreSQLVolume `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLVolume](
		"field", "last_name", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_public_network".
	Type constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLPublicNetwork](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_cpu".
	Type constant.DbaasPostgreSQLCPU `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLCPU](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_memory".
	Type constant.DbaasPostgreSQLMemory `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLMemory](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_connection_pooler".
	Type constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
	paramObj
}

func (r CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsSchemaFilterDbaasPostgreSQLConnectionPooler](
		"field", "last_name", "type", "uuid",
	)
}

// Filter by tags
//
// The property Conditions is required.
type CostReportGetAggregatedParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []CostReportGetAggregatedParamsTagsCondition `json:"conditions,omitzero,required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	//
	// Any of "AND", "OR".
	ConditionType string `json:"condition_type,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetAggregatedParamsTags](
		"condition_type", "AND", "OR",
	)
}

type CostReportGetAggregatedParamsTagsCondition struct {
	// The name of the tag to filter (e.g., '`os_version`').
	Key param.Opt[string] `json:"key,omitzero"`
	// Determines how strictly the tag value must match the specified value. If true,
	// the tag value must exactly match the given value. If false, a less strict match
	// (e.g., partial or case-insensitive match) may be applied.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The value of the tag to filter (e.g., '22.04').
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r CostReportGetAggregatedParamsTagsCondition) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetAggregatedParamsTagsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetAggregatedParamsTagsCondition) UnmarshalJSON(data []byte) error {
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
	// "inference", "instance", "load_balancer", "log_index", "snapshot", "volume".
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "snapshot".
	Type constant.Snapshot `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "instance".
	Type constant.Instance `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "ai_cluster".
	Type constant.AICluster `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ai_virtual_cluster".
	Type constant.AIVirtualCluster `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "basic_vm".
	Type constant.BasicVm `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "baremetal".
	Type constant.Baremetal `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "volume".
	Type constant.Volume `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file_share".
	Type constant.FileShare `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "floatingip".
	Type constant.Floatingip `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "egress_traffic".
	Type constant.EgressTraffic `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "load_balancer".
	Type constant.LoadBalancer `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "external_ip".
	Type constant.ExternalIP `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "backup".
	Type constant.Backup `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "log_index".
	Type constant.LogIndex `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions".
	Type constant.Functions `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions_calls".
	Type constant.FunctionsCalls `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "functions_traffic".
	Type constant.FunctionsTraffic `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "containers".
	Type constant.Containers `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "inference".
	Type constant.Inference `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_volume".
	Type constant.DbaasPostgreSQLVolume `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_public_network".
	Type constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_cpu".
	Type constant.DbaasPostgreSQLCPU `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_memory".
	Type constant.DbaasPostgreSQLMemory `json:"type,required"`
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
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_connection_pooler".
	Type constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
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
	Conditions []CostReportGetAggregatedMonthlyParamsTagsCondition `json:"conditions,omitzero,required"`
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
	// The name of the tag to filter (e.g., '`os_version`').
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

type CostReportGetDetailedParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day.
	TimeFrom time.Time `json:"time_from,required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day.
	TimeTo time.Time `json:"time_to,required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Opt[bool] `json:"enable_last_day,omitzero"`
	// The response resources limit. Defaults to 10.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The response resources offset.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// Round cost values to 5 decimal places. When false, returns full precision.
	Rounding param.Opt[bool] `json:"rounding,omitzero"`
	// List of project IDs
	Projects []int64 `json:"projects,omitzero"`
	// List of region IDs.
	Regions []int64 `json:"regions,omitzero"`
	// Format of the response (csv or json).
	//
	// Any of "csv_records", "json".
	ResponseFormat CostReportGetDetailedParamsResponseFormat `json:"response_format,omitzero"`
	// Extended filter for field filtering.
	SchemaFilter CostReportGetDetailedParamsSchemaFilterUnion `json:"schema_filter,omitzero"`
	// List of sorting filters (JSON objects) fields: project. directions: asc, desc.
	Sorting []CostReportGetDetailedParamsSorting `json:"sorting,omitzero"`
	// Filter by tags
	Tags CostReportGetDetailedParamsTags `json:"tags,omitzero"`
	// List of resource types to be filtered in the report.
	//
	// Any of "ai_cluster", "ai_virtual_cluster", "backup", "baremetal", "basic_vm",
	// "containers", "dbaas_postgresql_connection_pooler", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_memory", "dbaas_postgresql_public_network",
	// "dbaas_postgresql_volume", "egress_traffic", "external_ip", "file_share",
	// "floatingip", "functions", "functions_calls", "functions_traffic", "image",
	// "inference", "instance", "load_balancer", "log_index", "snapshot", "volume".
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

// Format of the response (csv or json).
type CostReportGetDetailedParamsResponseFormat string

const (
	CostReportGetDetailedParamsResponseFormatCsvRecords CostReportGetDetailedParamsResponseFormat = "csv_records"
	CostReportGetDetailedParamsResponseFormatJson       CostReportGetDetailedParamsResponseFormat = "json"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CostReportGetDetailedParamsSchemaFilterUnion struct {
	OfSnapshot                        *CostReportGetDetailedParamsSchemaFilterSnapshot                        `json:",omitzero,inline"`
	OfInstance                        *CostReportGetDetailedParamsSchemaFilterInstance                        `json:",omitzero,inline"`
	OfAICluster                       *CostReportGetDetailedParamsSchemaFilterAICluster                       `json:",omitzero,inline"`
	OfAIVirtualCluster                *CostReportGetDetailedParamsSchemaFilterAIVirtualCluster                `json:",omitzero,inline"`
	OfBasicVm                         *CostReportGetDetailedParamsSchemaFilterBasicVm                         `json:",omitzero,inline"`
	OfBaremetal                       *CostReportGetDetailedParamsSchemaFilterBaremetal                       `json:",omitzero,inline"`
	OfVolume                          *CostReportGetDetailedParamsSchemaFilterVolume                          `json:",omitzero,inline"`
	OfFileShare                       *CostReportGetDetailedParamsSchemaFilterFileShare                       `json:",omitzero,inline"`
	OfImage                           *CostReportGetDetailedParamsSchemaFilterImage                           `json:",omitzero,inline"`
	OfFloatingip                      *CostReportGetDetailedParamsSchemaFilterFloatingip                      `json:",omitzero,inline"`
	OfEgressTraffic                   *CostReportGetDetailedParamsSchemaFilterEgressTraffic                   `json:",omitzero,inline"`
	OfLoadBalancer                    *CostReportGetDetailedParamsSchemaFilterLoadBalancer                    `json:",omitzero,inline"`
	OfExternalIP                      *CostReportGetDetailedParamsSchemaFilterExternalIP                      `json:",omitzero,inline"`
	OfBackup                          *CostReportGetDetailedParamsSchemaFilterBackup                          `json:",omitzero,inline"`
	OfLogIndex                        *CostReportGetDetailedParamsSchemaFilterLogIndex                        `json:",omitzero,inline"`
	OfFunctions                       *CostReportGetDetailedParamsSchemaFilterFunctions                       `json:",omitzero,inline"`
	OfFunctionsCalls                  *CostReportGetDetailedParamsSchemaFilterFunctionsCalls                  `json:",omitzero,inline"`
	OfFunctionsTraffic                *CostReportGetDetailedParamsSchemaFilterFunctionsTraffic                `json:",omitzero,inline"`
	OfContainers                      *CostReportGetDetailedParamsSchemaFilterContainers                      `json:",omitzero,inline"`
	OfInference                       *CostReportGetDetailedParamsSchemaFilterInference                       `json:",omitzero,inline"`
	OfDbaasPostgreSQLVolume           *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume           `json:",omitzero,inline"`
	OfDbaasPostgreSQLPublicNetwork    *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork    `json:",omitzero,inline"`
	OfDbaasPostgreSQLCPU              *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU              `json:",omitzero,inline"`
	OfDbaasPostgreSQLMemory           *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory           `json:",omitzero,inline"`
	OfDbaasPostgreSQLConnectionPooler *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler `json:",omitzero,inline"`
	paramUnion
}

func (u CostReportGetDetailedParamsSchemaFilterUnion) MarshalJSON() ([]byte, error) {
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
func (u *CostReportGetDetailedParamsSchemaFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CostReportGetDetailedParamsSchemaFilterUnion) asAny() any {
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
func (u CostReportGetDetailedParamsSchemaFilterUnion) GetField() *string {
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
func (u CostReportGetDetailedParamsSchemaFilterUnion) GetType() *string {
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
func (u CostReportGetDetailedParamsSchemaFilterUnion) GetValues() []string {
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
	apijson.RegisterUnion[CostReportGetDetailedParamsSchemaFilterUnion](
		"type",
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterSnapshot]("snapshot"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterInstance]("instance"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterAICluster]("ai_cluster"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterAIVirtualCluster]("ai_virtual_cluster"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterBasicVm]("basic_vm"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterBaremetal]("baremetal"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterVolume]("volume"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterFileShare]("file_share"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterImage]("image"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterFloatingip]("floatingip"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterEgressTraffic]("egress_traffic"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterLoadBalancer]("load_balancer"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterExternalIP]("external_ip"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterBackup]("backup"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterLogIndex]("log_index"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterFunctions]("functions"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterFunctionsCalls]("functions_calls"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterFunctionsTraffic]("functions_traffic"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterContainers]("containers"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterInference]("inference"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume]("dbaas_postgresql_volume"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork]("dbaas_postgresql_public_network"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU]("dbaas_postgresql_cpu"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory]("dbaas_postgresql_memory"),
		apijson.Discriminator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler]("dbaas_postgresql_connection_pooler"),
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterSnapshot struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "source_volume_uuid", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "snapshot".
	Type constant.Snapshot `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterSnapshot](
		"field", "last_name", "last_size", "source_volume_uuid", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterInstance struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "instance".
	Type constant.Instance `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterInstance) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterInstance](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterAICluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "ai_cluster".
	Type constant.AICluster `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterAICluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterAICluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterAICluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterAIVirtualCluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ai_virtual_cluster".
	Type constant.AIVirtualCluster `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterAIVirtualCluster) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterAIVirtualCluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterAIVirtualCluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterBasicVm struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "basic_vm".
	Type constant.BasicVm `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterBasicVm) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterBasicVm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterBasicVm](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterBaremetal struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "baremetal".
	Type constant.Baremetal `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterBaremetal) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterBaremetal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterBaremetal](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterVolume struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "last_name", "last_size", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "volume".
	Type constant.Volume `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterVolume](
		"field", "attached_to_vm", "last_name", "last_size", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterFileShare struct {
	// Field name to filter by
	//
	// Any of "file_share_type", "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file_share".
	Type constant.FileShare `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterFileShare) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterFileShare](
		"field", "file_share_type", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterImage struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterImage) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterImage](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterFloatingip struct {
	// Field name to filter by
	//
	// Any of "ip_address", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "floatingip".
	Type constant.Floatingip `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterFloatingip) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterFloatingip
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterFloatingip](
		"field", "ip_address", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterEgressTraffic struct {
	// Field name to filter by
	//
	// Any of "instance_name", "instance_type", "port_id", "type", "vm_id".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "egress_traffic".
	Type constant.EgressTraffic `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterEgressTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterEgressTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterEgressTraffic](
		"field", "instance_name", "instance_type", "port_id", "type", "vm_id",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterLoadBalancer struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "load_balancer".
	Type constant.LoadBalancer `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterLoadBalancer) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterLoadBalancer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterLoadBalancer](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterExternalIP struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id",
	// "type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "external_ip".
	Type constant.ExternalIP `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterExternalIP) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterExternalIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterExternalIP](
		"field", "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id", "type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterBackup struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "schedule_id", "source_volume_uuid", "type",
	// "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "backup".
	Type constant.Backup `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterBackup) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterBackup](
		"field", "last_name", "last_size", "schedule_id", "source_volume_uuid", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterLogIndex struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "log_index".
	Type constant.LogIndex `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterLogIndex) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterLogIndex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterLogIndex](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterFunctions struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions".
	Type constant.Functions `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterFunctions) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterFunctions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterFunctions](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterFunctionsCalls struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "functions_calls".
	Type constant.FunctionsCalls `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterFunctionsCalls) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterFunctionsCalls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterFunctionsCalls](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterFunctionsTraffic struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "functions_traffic".
	Type constant.FunctionsTraffic `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterFunctionsTraffic) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterFunctionsTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterFunctionsTraffic](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterContainers struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "containers".
	Type constant.Containers `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterContainers) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterContainers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterContainers](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterInference struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "inference".
	Type constant.Inference `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterInference) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterInference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterInference](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid", "volume_type".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_volume".
	Type constant.DbaasPostgreSQLVolume `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLVolume](
		"field", "last_name", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_public_network".
	Type constant.DbaasPostgreSQLPublicNetwork `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLPublicNetwork](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_cpu".
	Type constant.DbaasPostgreSQLCPU `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLCPU](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_memory".
	Type constant.DbaasPostgreSQLMemory `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLMemory](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero,required"`
	// List of field values to filter
	Values []string `json:"values,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_connection_pooler".
	Type constant.DbaasPostgreSQLConnectionPooler `json:"type,required"`
	paramObj
}

func (r CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSchemaFilterDbaasPostgreSQLConnectionPooler](
		"field", "last_name", "type", "uuid",
	)
}

type CostReportGetDetailedParamsSorting struct {
	// Any of "asc", "desc".
	BillingValue string `json:"billing_value,omitzero"`
	// Any of "asc", "desc".
	FirstSeen string `json:"first_seen,omitzero"`
	// Any of "asc", "desc".
	LastName string `json:"last_name,omitzero"`
	// Any of "asc", "desc".
	LastSeen string `json:"last_seen,omitzero"`
	// Any of "asc", "desc".
	Project string `json:"project,omitzero"`
	// Any of "asc", "desc".
	Region string `json:"region,omitzero"`
	// Any of "asc", "desc".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r CostReportGetDetailedParamsSorting) MarshalJSON() (data []byte, err error) {
	type shadow CostReportGetDetailedParamsSorting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostReportGetDetailedParamsSorting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"billing_value", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"first_seen", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"last_name", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"last_seen", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"project", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"region", "asc", "desc",
	)
	apijson.RegisterFieldValidator[CostReportGetDetailedParamsSorting](
		"type", "asc", "desc",
	)
}

// Filter by tags
//
// The property Conditions is required.
type CostReportGetDetailedParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []CostReportGetDetailedParamsTagsCondition `json:"conditions,omitzero,required"`
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
	// The name of the tag to filter (e.g., '`os_version`').
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
