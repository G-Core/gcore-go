// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/tidwall/gjson"
)

// CloudV1ReservationCostReportService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ReservationCostReportService] method instead.
type CloudV1ReservationCostReportService struct {
	Options []option.RequestOption
}

// NewCloudV1ReservationCostReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1ReservationCostReportService(opts ...option.RequestOption) (r *CloudV1ReservationCostReportService) {
	r = &CloudV1ReservationCostReportService{}
	r.Options = opts
	return
}

// Retrieve a detailed cost report totals for a specified month, which includes
// both commit and pay-as-you-go (overcommit) prices. Additionally, it provides the
// spent billing units (e.g., hours or GB) for resources. The "time_to" parameter
// represents all days in the specified month.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CloudV1ReservationCostReportService) GetTotals(ctx context.Context, body CloudV1ReservationCostReportGetTotalsParams, opts ...option.RequestOption) (res *CloudV1ReservationCostReportGetTotalsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reservation_cost_report/totals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV1ReservationCostReportGetTotalsResponse struct {
	// Total count of the totals
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	PriceStatus PriceDisplayStatus                                    `json:"price_status,required"`
	Results     []CloudV1ReservationCostReportGetTotalsResponseResult `json:"results,required"`
	JSON        cloudV1ReservationCostReportGetTotalsResponseJSON     `json:"-"`
}

// cloudV1ReservationCostReportGetTotalsResponseJSON contains the JSON metadata for
// the struct [CloudV1ReservationCostReportGetTotalsResponse]
type cloudV1ReservationCostReportGetTotalsResponseJSON struct {
	Count       apijson.Field
	PriceStatus apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ReservationCostReportGetTotalsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ReservationCostReportGetTotalsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservationCostReportGetTotalsResponseResult struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// This field can have the runtime type of
	// [TotalAIClusterWithCostBillingValueUnion],
	// [TotalBaremetalWithCostBillingValueUnion],
	// [TotalBasicVmWithCostBillingValueUnion], [TotalBackupWithCostBillingValueUnion],
	// [TotalContainerWithCostBillingValueUnion],
	// [TotalEgressTrafficWithCostBillingValueUnion],
	// [TotalExternalIPWithCostBillingValueUnion],
	// [TotalFileShareWithCostBillingValueUnion],
	// [TotalFloatingIPWithCostBillingValueUnion],
	// [TotalFunctionsWithCostBillingValueUnion],
	// [TotalFunctionCallsWithCostBillingValueUnion],
	// [TotalFunctionEgressTrafficWithCostBillingValueUnion],
	// [TotalImagesWithCostBillingValueUnion],
	// [TotalInferenceWithCostBillingValueUnion],
	// [TotalInstanceWithCostBillingValueUnion],
	// [TotalLoadBalancerWithCostBillingValueUnion],
	// [TotalLogIndexWithCostBillingValueUnion],
	// [TotalSnapshotWithCostBillingValueUnion],
	// [TotalVolumeWithCostBillingValueUnion],
	// [TotalDbaasPostgresqlPoolerWithCostBillingValueUnion],
	// [TotalDbaasPostgresqlMemoryWithCostBillingValueUnion],
	// [TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnion],
	// [TotalDbaasPostgresqlCPUWithCostBillingValueUnion],
	// [TotalDbaasPostgresqlVolumeWithCostBillingValueUnion].
	BillingValue interface{} `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                                                    `json:"region_id,required"`
	Type               CloudV1ReservationCostReportGetTotalsResponseResultsType `json:"type,required"`
	BillingFeatureName string                                                   `json:"billing_feature_name,nullable"`
	// This field can have the runtime type of [TotalAIClusterWithCostCostUnion],
	// [TotalBaremetalWithCostCostUnion], [TotalBasicVmWithCostCostUnion],
	// [TotalBackupWithCostCostUnion], [TotalContainerWithCostCostUnion],
	// [TotalEgressTrafficWithCostCostUnion], [TotalExternalIPWithCostCostUnion],
	// [TotalFileShareWithCostCostUnion], [TotalFloatingIPWithCostCostUnion],
	// [TotalFunctionsWithCostCostUnion], [TotalFunctionCallsWithCostCostUnion],
	// [TotalFunctionEgressTrafficWithCostCostUnion], [TotalImagesWithCostCostUnion],
	// [TotalInferenceWithCostCostUnion], [TotalInstanceWithCostCostUnion],
	// [TotalLoadBalancerWithCostCostUnion], [TotalLogIndexWithCostCostUnion],
	// [TotalSnapshotWithCostCostUnion], [TotalVolumeWithCostCostUnion],
	// [TotalDbaasPostgresqlPoolerWithCostCostUnion],
	// [TotalDbaasPostgresqlMemoryWithCostCostUnion],
	// [TotalDbaasPostgresqlPublicNetworkWithCostCostUnion],
	// [TotalDbaasPostgresqlCPUWithCostCostUnion],
	// [TotalDbaasPostgresqlVolumeWithCostCostUnion].
	Cost interface{} `json:"cost"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Type of the file share
	FileShareType string `json:"file_share_type"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// Type of the volume
	VolumeType string                                                  `json:"volume_type"`
	JSON       cloudV1ReservationCostReportGetTotalsResponseResultJSON `json:"-"`
	union      CloudV1ReservationCostReportGetTotalsResponseResultsUnion
}

// cloudV1ReservationCostReportGetTotalsResponseResultJSON contains the JSON
// metadata for the struct [CloudV1ReservationCostReportGetTotalsResponseResult]
type cloudV1ReservationCostReportGetTotalsResponseResultJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	FileShareType      apijson.Field
	Flavor             apijson.Field
	InstanceType       apijson.Field
	LastSize           apijson.Field
	RegionName         apijson.Field
	VolumeType         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r cloudV1ReservationCostReportGetTotalsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1ReservationCostReportGetTotalsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1ReservationCostReportGetTotalsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1ReservationCostReportGetTotalsResponseResultsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TotalAIClusterWithCost],
// [TotalBaremetalWithCost], [TotalBasicVmWithCost], [TotalBackupWithCost],
// [TotalContainerWithCost], [TotalEgressTrafficWithCost],
// [TotalExternalIPWithCost], [TotalFileShareWithCost], [TotalFloatingIPWithCost],
// [TotalFunctionsWithCost], [TotalFunctionCallsWithCost],
// [TotalFunctionEgressTrafficWithCost], [TotalImagesWithCost],
// [TotalInferenceWithCost], [TotalInstanceWithCost], [TotalLoadBalancerWithCost],
// [TotalLogIndexWithCost], [TotalSnapshotWithCost], [TotalVolumeWithCost],
// [TotalDbaasPostgresqlPoolerWithCost], [TotalDbaasPostgresqlMemoryWithCost],
// [TotalDbaasPostgresqlPublicNetworkWithCost], [TotalDbaasPostgresqlCPUWithCost],
// [TotalDbaasPostgresqlVolumeWithCost].
func (r CloudV1ReservationCostReportGetTotalsResponseResult) AsUnion() CloudV1ReservationCostReportGetTotalsResponseResultsUnion {
	return r.union
}

// Union satisfied by [TotalAIClusterWithCost], [TotalBaremetalWithCost],
// [TotalBasicVmWithCost], [TotalBackupWithCost], [TotalContainerWithCost],
// [TotalEgressTrafficWithCost], [TotalExternalIPWithCost],
// [TotalFileShareWithCost], [TotalFloatingIPWithCost], [TotalFunctionsWithCost],
// [TotalFunctionCallsWithCost], [TotalFunctionEgressTrafficWithCost],
// [TotalImagesWithCost], [TotalInferenceWithCost], [TotalInstanceWithCost],
// [TotalLoadBalancerWithCost], [TotalLogIndexWithCost], [TotalSnapshotWithCost],
// [TotalVolumeWithCost], [TotalDbaasPostgresqlPoolerWithCost],
// [TotalDbaasPostgresqlMemoryWithCost],
// [TotalDbaasPostgresqlPublicNetworkWithCost], [TotalDbaasPostgresqlCPUWithCost]
// or [TotalDbaasPostgresqlVolumeWithCost].
type CloudV1ReservationCostReportGetTotalsResponseResultsUnion interface {
	implementsCloudV1ReservationCostReportGetTotalsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1ReservationCostReportGetTotalsResponseResultsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalAIClusterWithCost{}),
			DiscriminatorValue: "ai_cluster",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalBaremetalWithCost{}),
			DiscriminatorValue: "baremetal",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalBasicVmWithCost{}),
			DiscriminatorValue: "basic_vm",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalBackupWithCost{}),
			DiscriminatorValue: "backup",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalContainerWithCost{}),
			DiscriminatorValue: "containers",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalEgressTrafficWithCost{}),
			DiscriminatorValue: "egress_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalExternalIPWithCost{}),
			DiscriminatorValue: "external_ip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalFileShareWithCost{}),
			DiscriminatorValue: "file_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalFloatingIPWithCost{}),
			DiscriminatorValue: "floatingip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalFunctionsWithCost{}),
			DiscriminatorValue: "functions",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalFunctionCallsWithCost{}),
			DiscriminatorValue: "functions_calls",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalFunctionEgressTrafficWithCost{}),
			DiscriminatorValue: "functions_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalImagesWithCost{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalInferenceWithCost{}),
			DiscriminatorValue: "inference",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalInstanceWithCost{}),
			DiscriminatorValue: "instance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalLoadBalancerWithCost{}),
			DiscriminatorValue: "load_balancer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalLogIndexWithCost{}),
			DiscriminatorValue: "log_index",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalSnapshotWithCost{}),
			DiscriminatorValue: "snapshot",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalVolumeWithCost{}),
			DiscriminatorValue: "volume",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalDbaasPostgresqlPoolerWithCost{}),
			DiscriminatorValue: "dbaas_postgresql_connection_pooler",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalDbaasPostgresqlMemoryWithCost{}),
			DiscriminatorValue: "dbaas_postgresql_memory",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalDbaasPostgresqlPublicNetworkWithCost{}),
			DiscriminatorValue: "dbaas_postgresql_public_network",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalDbaasPostgresqlCPUWithCost{}),
			DiscriminatorValue: "dbaas_postgresql_cpu",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TotalDbaasPostgresqlVolumeWithCost{}),
			DiscriminatorValue: "dbaas_postgresql_volume",
		},
	)
}

type CloudV1ReservationCostReportGetTotalsResponseResultsType string

const (
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeAICluster                       CloudV1ReservationCostReportGetTotalsResponseResultsType = "ai_cluster"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeBaremetal                       CloudV1ReservationCostReportGetTotalsResponseResultsType = "baremetal"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeBasicVm                         CloudV1ReservationCostReportGetTotalsResponseResultsType = "basic_vm"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeBackup                          CloudV1ReservationCostReportGetTotalsResponseResultsType = "backup"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeContainers                      CloudV1ReservationCostReportGetTotalsResponseResultsType = "containers"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeEgressTraffic                   CloudV1ReservationCostReportGetTotalsResponseResultsType = "egress_traffic"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeExternalIP                      CloudV1ReservationCostReportGetTotalsResponseResultsType = "external_ip"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeFileShare                       CloudV1ReservationCostReportGetTotalsResponseResultsType = "file_share"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeFloatingip                      CloudV1ReservationCostReportGetTotalsResponseResultsType = "floatingip"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctions                       CloudV1ReservationCostReportGetTotalsResponseResultsType = "functions"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctionsCalls                  CloudV1ReservationCostReportGetTotalsResponseResultsType = "functions_calls"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctionsTraffic                CloudV1ReservationCostReportGetTotalsResponseResultsType = "functions_traffic"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeImage                           CloudV1ReservationCostReportGetTotalsResponseResultsType = "image"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeInference                       CloudV1ReservationCostReportGetTotalsResponseResultsType = "inference"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeInstance                        CloudV1ReservationCostReportGetTotalsResponseResultsType = "instance"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeLoadBalancer                    CloudV1ReservationCostReportGetTotalsResponseResultsType = "load_balancer"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeLogIndex                        CloudV1ReservationCostReportGetTotalsResponseResultsType = "log_index"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeSnapshot                        CloudV1ReservationCostReportGetTotalsResponseResultsType = "snapshot"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeVolume                          CloudV1ReservationCostReportGetTotalsResponseResultsType = "volume"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlConnectionPooler CloudV1ReservationCostReportGetTotalsResponseResultsType = "dbaas_postgresql_connection_pooler"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlMemory           CloudV1ReservationCostReportGetTotalsResponseResultsType = "dbaas_postgresql_memory"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlPublicNetwork    CloudV1ReservationCostReportGetTotalsResponseResultsType = "dbaas_postgresql_public_network"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlCPU              CloudV1ReservationCostReportGetTotalsResponseResultsType = "dbaas_postgresql_cpu"
	CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlVolume           CloudV1ReservationCostReportGetTotalsResponseResultsType = "dbaas_postgresql_volume"
)

func (r CloudV1ReservationCostReportGetTotalsResponseResultsType) IsKnown() bool {
	switch r {
	case CloudV1ReservationCostReportGetTotalsResponseResultsTypeAICluster, CloudV1ReservationCostReportGetTotalsResponseResultsTypeBaremetal, CloudV1ReservationCostReportGetTotalsResponseResultsTypeBasicVm, CloudV1ReservationCostReportGetTotalsResponseResultsTypeBackup, CloudV1ReservationCostReportGetTotalsResponseResultsTypeContainers, CloudV1ReservationCostReportGetTotalsResponseResultsTypeEgressTraffic, CloudV1ReservationCostReportGetTotalsResponseResultsTypeExternalIP, CloudV1ReservationCostReportGetTotalsResponseResultsTypeFileShare, CloudV1ReservationCostReportGetTotalsResponseResultsTypeFloatingip, CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctions, CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctionsCalls, CloudV1ReservationCostReportGetTotalsResponseResultsTypeFunctionsTraffic, CloudV1ReservationCostReportGetTotalsResponseResultsTypeImage, CloudV1ReservationCostReportGetTotalsResponseResultsTypeInference, CloudV1ReservationCostReportGetTotalsResponseResultsTypeInstance, CloudV1ReservationCostReportGetTotalsResponseResultsTypeLoadBalancer, CloudV1ReservationCostReportGetTotalsResponseResultsTypeLogIndex, CloudV1ReservationCostReportGetTotalsResponseResultsTypeSnapshot, CloudV1ReservationCostReportGetTotalsResponseResultsTypeVolume, CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlConnectionPooler, CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlMemory, CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlPublicNetwork, CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlCPU, CloudV1ReservationCostReportGetTotalsResponseResultsTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1ReservationCostReportGetTotalsParams struct {
	// Beginning of the period: YYYY-mm
	TimeFrom param.Field[time.Time] `json:"time_from,required" format:"date-time"`
	// End of the period: YYYY-mm
	TimeTo param.Field[time.Time] `json:"time_to,required" format:"date-time"`
	// List of region IDs.
	Regions param.Field[[]int64] `json:"regions"`
	// Format of the response (csv_totals or json).
	ResponseFormat param.Field[CloudV1ReservationCostReportGetTotalsParamsResponseFormat] `json:"response_format"`
	// Extended filter for field filtering.
	SchemaFilter param.Field[CloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion] `json:"schema_filter"`
	// Filter by tags
	Tags param.Field[TagsFilterParam] `json:"tags"`
	// List of resource types to be filtered in the report.
	Types param.Field[[]PrebillingResourceTypes] `json:"types"`
}

func (r CloudV1ReservationCostReportGetTotalsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Format of the response (csv_totals or json).
type CloudV1ReservationCostReportGetTotalsParamsResponseFormat string

const (
	CloudV1ReservationCostReportGetTotalsParamsResponseFormatCsvTotals CloudV1ReservationCostReportGetTotalsParamsResponseFormat = "csv_totals"
	CloudV1ReservationCostReportGetTotalsParamsResponseFormatJson      CloudV1ReservationCostReportGetTotalsParamsResponseFormat = "json"
)

func (r CloudV1ReservationCostReportGetTotalsParamsResponseFormat) IsKnown() bool {
	switch r {
	case CloudV1ReservationCostReportGetTotalsParamsResponseFormatCsvTotals, CloudV1ReservationCostReportGetTotalsParamsResponseFormatJson:
		return true
	}
	return false
}

// Extended filter for field filtering.
type CloudV1ReservationCostReportGetTotalsParamsSchemaFilter struct {
	// Field name to filter by
	Field  param.Field[CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField] `json:"field,required"`
	Type   param.Field[CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType]  `json:"type,required"`
	Values param.Field[interface{}]                                                  `json:"values,required"`
}

func (r CloudV1ReservationCostReportGetTotalsParamsSchemaFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1ReservationCostReportGetTotalsParamsSchemaFilter) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Extended filter for field filtering.
//
// Satisfied by [SchemaFilterSnapshotParam], [SchemaFilterInstanceParam],
// [SchemaFilterAIClusterParam], [SchemaFilterBasicVmParam],
// [SchemaFilterBaremetalParam], [SchemaFilterVolumeParam],
// [SchemaFilterFileShareParam], [SchemaFilterImageParam],
// [SchemaFilterFloatingIPParam], [SchemaFilterEgressTrafficParam],
// [SchemaFilterLoadBalancerParam], [SchemaFilterExternalIPParam],
// [SchemaFilterBackupParam], [SchemaFilterLogIndexParam],
// [SchemaFilterFunctionsParam], [SchemaFilterFunctionsCallsParam],
// [SchemaFilterFunctionsTrafficParam], [SchemaFilterContainersParam],
// [SchemaFilterInferenceParam], [SchemaFilterDbaasPostgresqlVolumeParam],
// [SchemaFilterDbaasPostgresqlPublicNetworkParam],
// [SchemaFilterDbaasPostgresqlCPUParam], [SchemaFilterDbaasPostgresqlMemoryParam],
// [SchemaFilterDbaasPostgresqlPoolerParam],
// [CloudV1ReservationCostReportGetTotalsParamsSchemaFilter].
type CloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion interface {
	implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion()
}

// Field name to filter by
type CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField string

const (
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldLastName         CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "last_name"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldLastSize         CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "last_size"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldSourceVolumeUuid CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "source_volume_uuid"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldType             CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "type"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldUuid             CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "uuid"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldVolumeType       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "volume_type"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldFlavor           CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "flavor"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldAttachedToVm     CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "attached_to_vm"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldFileShareType    CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "file_share_type"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldIPAddress        CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "ip_address"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldInstanceName     CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "instance_name"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldInstanceType     CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "instance_type"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldPortID           CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "port_id"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldVmID             CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "vm_id"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldNetworkID        CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "network_id"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldSubnetID         CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "subnet_id"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldScheduleID       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField = "schedule_id"
)

func (r CloudV1ReservationCostReportGetTotalsParamsSchemaFilterField) IsKnown() bool {
	switch r {
	case CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldLastName, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldLastSize, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldSourceVolumeUuid, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldType, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldUuid, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldVolumeType, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldFlavor, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldAttachedToVm, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldFileShareType, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldIPAddress, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldInstanceName, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldInstanceType, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldPortID, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldVmID, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldNetworkID, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldSubnetID, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterFieldScheduleID:
		return true
	}
	return false
}

type CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType string

const (
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeSnapshot                        CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "snapshot"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeInstance                        CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "instance"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeAICluster                       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "ai_cluster"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBasicVm                         CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "basic_vm"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBaremetal                       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "baremetal"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeVolume                          CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "volume"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFileShare                       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "file_share"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeImage                           CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "image"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFloatingip                      CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "floatingip"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeEgressTraffic                   CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "egress_traffic"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeLoadBalancer                    CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "load_balancer"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeExternalIP                      CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "external_ip"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBackup                          CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "backup"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeLogIndex                        CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "log_index"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctions                       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "functions"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctionsCalls                  CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "functions_calls"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctionsTraffic                CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "functions_traffic"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeContainers                      CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "containers"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeInference                       CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "inference"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlVolume           CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_volume"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork    CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_public_network"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlCPU              CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_cpu"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlMemory           CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_memory"
	CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1ReservationCostReportGetTotalsParamsSchemaFilterType) IsKnown() bool {
	switch r {
	case CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeSnapshot, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeInstance, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeAICluster, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBasicVm, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBaremetal, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeVolume, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFileShare, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeImage, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFloatingip, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeEgressTraffic, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeLoadBalancer, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeExternalIP, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeBackup, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeLogIndex, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctions, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctionsCalls, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeFunctionsTraffic, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeContainers, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeInference, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlVolume, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlCPU, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlMemory, CloudV1ReservationCostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}
