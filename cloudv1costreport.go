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
	"github.com/stainless-sdks/gcore-go/shared"
	"github.com/tidwall/gjson"
)

// CloudV1CostReportService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1CostReportService] method instead.
type CloudV1CostReportService struct {
	Options []option.RequestOption
}

// NewCloudV1CostReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1CostReportService(opts ...option.RequestOption) (r *CloudV1CostReportService) {
	r = &CloudV1CostReportService{}
	r.Options = opts
	return
}

// Get a detailed cost report for a given period and specific resources. Requested
// period should not exceed 31 days.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/reservation_cost_report/totals, as the results from this report will not be
// accurate.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CloudV1CostReportService) GetResources(ctx context.Context, body CloudV1CostReportGetResourcesParams, opts ...option.RequestOption) (res *CloudV1CostReportGetResourcesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/cost_report/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get cost report totals (aggregated costs) for a given period. Requested period
// should not exceed 31 days.
//
// Note: This report assumes there are no active commit features in the billing
// plan. If there are active commit features (pre-paid resources) in your plan, use
// /v1/reservation_cost_report/totals, as the results from this report will not be
// accurate.
//
// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CloudV1CostReportService) GetTotals(ctx context.Context, body CloudV1CostReportGetTotalsParams, opts ...option.RequestOption) (res *CloudV1CostReportGetTotalsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/cost_report/totals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EgressTrafficInstanceType string

const (
	EgressTrafficInstanceTypeBaremetal EgressTrafficInstanceType = "baremetal"
	EgressTrafficInstanceTypeVm        EgressTrafficInstanceType = "vm"
)

func (r EgressTrafficInstanceType) IsKnown() bool {
	switch r {
	case EgressTrafficInstanceTypeBaremetal, EgressTrafficInstanceTypeVm:
		return true
	}
	return false
}

// Resource types for prebilling report
type PrebillingResourceTypes string

const (
	PrebillingResourceTypesAICluster                       PrebillingResourceTypes = "ai_cluster"
	PrebillingResourceTypesBackup                          PrebillingResourceTypes = "backup"
	PrebillingResourceTypesBaremetal                       PrebillingResourceTypes = "baremetal"
	PrebillingResourceTypesBasicVm                         PrebillingResourceTypes = "basic_vm"
	PrebillingResourceTypesContainers                      PrebillingResourceTypes = "containers"
	PrebillingResourceTypesDbaasPostgresqlConnectionPooler PrebillingResourceTypes = "dbaas_postgresql_connection_pooler"
	PrebillingResourceTypesDbaasPostgresqlCPU              PrebillingResourceTypes = "dbaas_postgresql_cpu"
	PrebillingResourceTypesDbaasPostgresqlMemory           PrebillingResourceTypes = "dbaas_postgresql_memory"
	PrebillingResourceTypesDbaasPostgresqlPublicNetwork    PrebillingResourceTypes = "dbaas_postgresql_public_network"
	PrebillingResourceTypesDbaasPostgresqlVolume           PrebillingResourceTypes = "dbaas_postgresql_volume"
	PrebillingResourceTypesEgressTraffic                   PrebillingResourceTypes = "egress_traffic"
	PrebillingResourceTypesExternalIP                      PrebillingResourceTypes = "external_ip"
	PrebillingResourceTypesFileShare                       PrebillingResourceTypes = "file_share"
	PrebillingResourceTypesFloatingip                      PrebillingResourceTypes = "floatingip"
	PrebillingResourceTypesFunctions                       PrebillingResourceTypes = "functions"
	PrebillingResourceTypesFunctionsCalls                  PrebillingResourceTypes = "functions_calls"
	PrebillingResourceTypesFunctionsTraffic                PrebillingResourceTypes = "functions_traffic"
	PrebillingResourceTypesImage                           PrebillingResourceTypes = "image"
	PrebillingResourceTypesInference                       PrebillingResourceTypes = "inference"
	PrebillingResourceTypesInstance                        PrebillingResourceTypes = "instance"
	PrebillingResourceTypesLoadBalancer                    PrebillingResourceTypes = "load_balancer"
	PrebillingResourceTypesLogIndex                        PrebillingResourceTypes = "log_index"
	PrebillingResourceTypesSnapshot                        PrebillingResourceTypes = "snapshot"
	PrebillingResourceTypesVolume                          PrebillingResourceTypes = "volume"
)

func (r PrebillingResourceTypes) IsKnown() bool {
	switch r {
	case PrebillingResourceTypesAICluster, PrebillingResourceTypesBackup, PrebillingResourceTypesBaremetal, PrebillingResourceTypesBasicVm, PrebillingResourceTypesContainers, PrebillingResourceTypesDbaasPostgresqlConnectionPooler, PrebillingResourceTypesDbaasPostgresqlCPU, PrebillingResourceTypesDbaasPostgresqlMemory, PrebillingResourceTypesDbaasPostgresqlPublicNetwork, PrebillingResourceTypesDbaasPostgresqlVolume, PrebillingResourceTypesEgressTraffic, PrebillingResourceTypesExternalIP, PrebillingResourceTypesFileShare, PrebillingResourceTypesFloatingip, PrebillingResourceTypesFunctions, PrebillingResourceTypesFunctionsCalls, PrebillingResourceTypesFunctionsTraffic, PrebillingResourceTypesImage, PrebillingResourceTypesInference, PrebillingResourceTypesInstance, PrebillingResourceTypesLoadBalancer, PrebillingResourceTypesLogIndex, PrebillingResourceTypesSnapshot, PrebillingResourceTypesVolume:
		return true
	}
	return false
}

type SchemaFilterAIClusterParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterAIClusterField] `json:"field,required"`
	Type  param.Field[SchemaFilterAIClusterType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterAIClusterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterAIClusterParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterAIClusterParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterAIClusterParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterAIClusterParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterAIClusterField string

const (
	SchemaFilterAIClusterFieldFlavor   SchemaFilterAIClusterField = "flavor"
	SchemaFilterAIClusterFieldLastName SchemaFilterAIClusterField = "last_name"
	SchemaFilterAIClusterFieldType     SchemaFilterAIClusterField = "type"
	SchemaFilterAIClusterFieldUuid     SchemaFilterAIClusterField = "uuid"
)

func (r SchemaFilterAIClusterField) IsKnown() bool {
	switch r {
	case SchemaFilterAIClusterFieldFlavor, SchemaFilterAIClusterFieldLastName, SchemaFilterAIClusterFieldType, SchemaFilterAIClusterFieldUuid:
		return true
	}
	return false
}

type SchemaFilterAIClusterType string

const (
	SchemaFilterAIClusterTypeAICluster SchemaFilterAIClusterType = "ai_cluster"
)

func (r SchemaFilterAIClusterType) IsKnown() bool {
	switch r {
	case SchemaFilterAIClusterTypeAICluster:
		return true
	}
	return false
}

type SchemaFilterBackupParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterBackupField] `json:"field,required"`
	Type  param.Field[SchemaFilterBackupType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterBackupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterBackupParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterBackupParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterBackupParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterBackupParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterBackupField string

const (
	SchemaFilterBackupFieldLastName         SchemaFilterBackupField = "last_name"
	SchemaFilterBackupFieldLastSize         SchemaFilterBackupField = "last_size"
	SchemaFilterBackupFieldScheduleID       SchemaFilterBackupField = "schedule_id"
	SchemaFilterBackupFieldSourceVolumeUuid SchemaFilterBackupField = "source_volume_uuid"
	SchemaFilterBackupFieldType             SchemaFilterBackupField = "type"
	SchemaFilterBackupFieldUuid             SchemaFilterBackupField = "uuid"
)

func (r SchemaFilterBackupField) IsKnown() bool {
	switch r {
	case SchemaFilterBackupFieldLastName, SchemaFilterBackupFieldLastSize, SchemaFilterBackupFieldScheduleID, SchemaFilterBackupFieldSourceVolumeUuid, SchemaFilterBackupFieldType, SchemaFilterBackupFieldUuid:
		return true
	}
	return false
}

type SchemaFilterBackupType string

const (
	SchemaFilterBackupTypeBackup SchemaFilterBackupType = "backup"
)

func (r SchemaFilterBackupType) IsKnown() bool {
	switch r {
	case SchemaFilterBackupTypeBackup:
		return true
	}
	return false
}

type SchemaFilterBaremetalParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterInstanceEnum]  `json:"field,required"`
	Type  param.Field[SchemaFilterBaremetalType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterBaremetalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterBaremetalParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterBaremetalParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterBaremetalParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterBaremetalParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterBaremetalType string

const (
	SchemaFilterBaremetalTypeBaremetal SchemaFilterBaremetalType = "baremetal"
)

func (r SchemaFilterBaremetalType) IsKnown() bool {
	switch r {
	case SchemaFilterBaremetalTypeBaremetal:
		return true
	}
	return false
}

type SchemaFilterBasicVmParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterInstanceEnum] `json:"field,required"`
	Type  param.Field[SchemaFilterBasicVmType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterBasicVmParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterBasicVmParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterBasicVmParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterBasicVmParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterBasicVmParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterBasicVmType string

const (
	SchemaFilterBasicVmTypeBasicVm SchemaFilterBasicVmType = "basic_vm"
)

func (r SchemaFilterBasicVmType) IsKnown() bool {
	switch r {
	case SchemaFilterBasicVmTypeBasicVm:
		return true
	}
	return false
}

type SchemaFilterContainersParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterContainersField] `json:"field,required"`
	Type  param.Field[SchemaFilterContainersType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterContainersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterContainersParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterContainersParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterContainersParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterContainersParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterContainersField string

const (
	SchemaFilterContainersFieldLastName SchemaFilterContainersField = "last_name"
	SchemaFilterContainersFieldType     SchemaFilterContainersField = "type"
	SchemaFilterContainersFieldUuid     SchemaFilterContainersField = "uuid"
)

func (r SchemaFilterContainersField) IsKnown() bool {
	switch r {
	case SchemaFilterContainersFieldLastName, SchemaFilterContainersFieldType, SchemaFilterContainersFieldUuid:
		return true
	}
	return false
}

type SchemaFilterContainersType string

const (
	SchemaFilterContainersTypeContainers SchemaFilterContainersType = "containers"
)

func (r SchemaFilterContainersType) IsKnown() bool {
	switch r {
	case SchemaFilterContainersTypeContainers:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlCPUParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterDbaasPostgresqlEnum]    `json:"field,required"`
	Type  param.Field[SchemaFilterDbaasPostgresqlCPUType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterDbaasPostgresqlCPUParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterDbaasPostgresqlCPUParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlCPUParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlCPUParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlCPUParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterDbaasPostgresqlCPUType string

const (
	SchemaFilterDbaasPostgresqlCPUTypeDbaasPostgresqlCPU SchemaFilterDbaasPostgresqlCPUType = "dbaas_postgresql_cpu"
)

func (r SchemaFilterDbaasPostgresqlCPUType) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlCPUTypeDbaasPostgresqlCPU:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlEnum string

const (
	SchemaFilterDbaasPostgresqlEnumLastName SchemaFilterDbaasPostgresqlEnum = "last_name"
	SchemaFilterDbaasPostgresqlEnumType     SchemaFilterDbaasPostgresqlEnum = "type"
	SchemaFilterDbaasPostgresqlEnumUuid     SchemaFilterDbaasPostgresqlEnum = "uuid"
)

func (r SchemaFilterDbaasPostgresqlEnum) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlEnumLastName, SchemaFilterDbaasPostgresqlEnumType, SchemaFilterDbaasPostgresqlEnumUuid:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlMemoryParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterDbaasPostgresqlEnum]       `json:"field,required"`
	Type  param.Field[SchemaFilterDbaasPostgresqlMemoryType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterDbaasPostgresqlMemoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterDbaasPostgresqlMemoryParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlMemoryParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlMemoryParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlMemoryParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterDbaasPostgresqlMemoryType string

const (
	SchemaFilterDbaasPostgresqlMemoryTypeDbaasPostgresqlMemory SchemaFilterDbaasPostgresqlMemoryType = "dbaas_postgresql_memory"
)

func (r SchemaFilterDbaasPostgresqlMemoryType) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlMemoryTypeDbaasPostgresqlMemory:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlPoolerParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterDbaasPostgresqlEnum]       `json:"field,required"`
	Type  param.Field[SchemaFilterDbaasPostgresqlPoolerType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterDbaasPostgresqlPoolerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterDbaasPostgresqlPoolerParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPoolerParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPoolerParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPoolerParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterDbaasPostgresqlPoolerType string

const (
	SchemaFilterDbaasPostgresqlPoolerTypeDbaasPostgresqlConnectionPooler SchemaFilterDbaasPostgresqlPoolerType = "dbaas_postgresql_connection_pooler"
)

func (r SchemaFilterDbaasPostgresqlPoolerType) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlPoolerTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlPublicNetworkParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterDbaasPostgresqlEnum]              `json:"field,required"`
	Type  param.Field[SchemaFilterDbaasPostgresqlPublicNetworkType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterDbaasPostgresqlPublicNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterDbaasPostgresqlPublicNetworkParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPublicNetworkParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPublicNetworkParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlPublicNetworkParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterDbaasPostgresqlPublicNetworkType string

const (
	SchemaFilterDbaasPostgresqlPublicNetworkTypeDbaasPostgresqlPublicNetwork SchemaFilterDbaasPostgresqlPublicNetworkType = "dbaas_postgresql_public_network"
)

func (r SchemaFilterDbaasPostgresqlPublicNetworkType) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlPublicNetworkTypeDbaasPostgresqlPublicNetwork:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlVolumeParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterDbaasPostgresqlVolumeField] `json:"field,required"`
	Type  param.Field[SchemaFilterDbaasPostgresqlVolumeType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterDbaasPostgresqlVolumeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterDbaasPostgresqlVolumeParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlVolumeParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlVolumeParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterDbaasPostgresqlVolumeParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterDbaasPostgresqlVolumeField string

const (
	SchemaFilterDbaasPostgresqlVolumeFieldLastName   SchemaFilterDbaasPostgresqlVolumeField = "last_name"
	SchemaFilterDbaasPostgresqlVolumeFieldType       SchemaFilterDbaasPostgresqlVolumeField = "type"
	SchemaFilterDbaasPostgresqlVolumeFieldUuid       SchemaFilterDbaasPostgresqlVolumeField = "uuid"
	SchemaFilterDbaasPostgresqlVolumeFieldVolumeType SchemaFilterDbaasPostgresqlVolumeField = "volume_type"
)

func (r SchemaFilterDbaasPostgresqlVolumeField) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlVolumeFieldLastName, SchemaFilterDbaasPostgresqlVolumeFieldType, SchemaFilterDbaasPostgresqlVolumeFieldUuid, SchemaFilterDbaasPostgresqlVolumeFieldVolumeType:
		return true
	}
	return false
}

type SchemaFilterDbaasPostgresqlVolumeType string

const (
	SchemaFilterDbaasPostgresqlVolumeTypeDbaasPostgresqlVolume SchemaFilterDbaasPostgresqlVolumeType = "dbaas_postgresql_volume"
)

func (r SchemaFilterDbaasPostgresqlVolumeType) IsKnown() bool {
	switch r {
	case SchemaFilterDbaasPostgresqlVolumeTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type SchemaFilterEgressTrafficParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterEgressTrafficField] `json:"field,required"`
	Type  param.Field[SchemaFilterEgressTrafficType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterEgressTrafficParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterEgressTrafficParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterEgressTrafficParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterEgressTrafficParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterEgressTrafficParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterEgressTrafficField string

const (
	SchemaFilterEgressTrafficFieldInstanceName SchemaFilterEgressTrafficField = "instance_name"
	SchemaFilterEgressTrafficFieldInstanceType SchemaFilterEgressTrafficField = "instance_type"
	SchemaFilterEgressTrafficFieldPortID       SchemaFilterEgressTrafficField = "port_id"
	SchemaFilterEgressTrafficFieldType         SchemaFilterEgressTrafficField = "type"
	SchemaFilterEgressTrafficFieldVmID         SchemaFilterEgressTrafficField = "vm_id"
)

func (r SchemaFilterEgressTrafficField) IsKnown() bool {
	switch r {
	case SchemaFilterEgressTrafficFieldInstanceName, SchemaFilterEgressTrafficFieldInstanceType, SchemaFilterEgressTrafficFieldPortID, SchemaFilterEgressTrafficFieldType, SchemaFilterEgressTrafficFieldVmID:
		return true
	}
	return false
}

type SchemaFilterEgressTrafficType string

const (
	SchemaFilterEgressTrafficTypeEgressTraffic SchemaFilterEgressTrafficType = "egress_traffic"
)

func (r SchemaFilterEgressTrafficType) IsKnown() bool {
	switch r {
	case SchemaFilterEgressTrafficTypeEgressTraffic:
		return true
	}
	return false
}

type SchemaFilterExternalIPParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterExternalIPField] `json:"field,required"`
	Type  param.Field[SchemaFilterExternalIPType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterExternalIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterExternalIPParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterExternalIPParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterExternalIPParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterExternalIPParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterExternalIPField string

const (
	SchemaFilterExternalIPFieldAttachedToVm SchemaFilterExternalIPField = "attached_to_vm"
	SchemaFilterExternalIPFieldIPAddress    SchemaFilterExternalIPField = "ip_address"
	SchemaFilterExternalIPFieldNetworkID    SchemaFilterExternalIPField = "network_id"
	SchemaFilterExternalIPFieldPortID       SchemaFilterExternalIPField = "port_id"
	SchemaFilterExternalIPFieldSubnetID     SchemaFilterExternalIPField = "subnet_id"
	SchemaFilterExternalIPFieldType         SchemaFilterExternalIPField = "type"
)

func (r SchemaFilterExternalIPField) IsKnown() bool {
	switch r {
	case SchemaFilterExternalIPFieldAttachedToVm, SchemaFilterExternalIPFieldIPAddress, SchemaFilterExternalIPFieldNetworkID, SchemaFilterExternalIPFieldPortID, SchemaFilterExternalIPFieldSubnetID, SchemaFilterExternalIPFieldType:
		return true
	}
	return false
}

type SchemaFilterExternalIPType string

const (
	SchemaFilterExternalIPTypeExternalIP SchemaFilterExternalIPType = "external_ip"
)

func (r SchemaFilterExternalIPType) IsKnown() bool {
	switch r {
	case SchemaFilterExternalIPTypeExternalIP:
		return true
	}
	return false
}

type SchemaFilterFileShareParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterFileShareField] `json:"field,required"`
	Type  param.Field[SchemaFilterFileShareType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterFileShareParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterFileShareParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterFileShareParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterFileShareParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterFileShareParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterFileShareField string

const (
	SchemaFilterFileShareFieldFileShareType SchemaFilterFileShareField = "file_share_type"
	SchemaFilterFileShareFieldLastName      SchemaFilterFileShareField = "last_name"
	SchemaFilterFileShareFieldLastSize      SchemaFilterFileShareField = "last_size"
	SchemaFilterFileShareFieldType          SchemaFilterFileShareField = "type"
	SchemaFilterFileShareFieldUuid          SchemaFilterFileShareField = "uuid"
)

func (r SchemaFilterFileShareField) IsKnown() bool {
	switch r {
	case SchemaFilterFileShareFieldFileShareType, SchemaFilterFileShareFieldLastName, SchemaFilterFileShareFieldLastSize, SchemaFilterFileShareFieldType, SchemaFilterFileShareFieldUuid:
		return true
	}
	return false
}

type SchemaFilterFileShareType string

const (
	SchemaFilterFileShareTypeFileShare SchemaFilterFileShareType = "file_share"
)

func (r SchemaFilterFileShareType) IsKnown() bool {
	switch r {
	case SchemaFilterFileShareTypeFileShare:
		return true
	}
	return false
}

type SchemaFilterFloatingIPParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterFloatingIPField] `json:"field,required"`
	Type  param.Field[SchemaFilterFloatingIPType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterFloatingIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterFloatingIPParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterFloatingIPParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterFloatingIPParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterFloatingIPParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterFloatingIPField string

const (
	SchemaFilterFloatingIPFieldIPAddress SchemaFilterFloatingIPField = "ip_address"
	SchemaFilterFloatingIPFieldLastName  SchemaFilterFloatingIPField = "last_name"
	SchemaFilterFloatingIPFieldType      SchemaFilterFloatingIPField = "type"
	SchemaFilterFloatingIPFieldUuid      SchemaFilterFloatingIPField = "uuid"
)

func (r SchemaFilterFloatingIPField) IsKnown() bool {
	switch r {
	case SchemaFilterFloatingIPFieldIPAddress, SchemaFilterFloatingIPFieldLastName, SchemaFilterFloatingIPFieldType, SchemaFilterFloatingIPFieldUuid:
		return true
	}
	return false
}

type SchemaFilterFloatingIPType string

const (
	SchemaFilterFloatingIPTypeFloatingip SchemaFilterFloatingIPType = "floatingip"
)

func (r SchemaFilterFloatingIPType) IsKnown() bool {
	switch r {
	case SchemaFilterFloatingIPTypeFloatingip:
		return true
	}
	return false
}

type SchemaFilterFunctionsParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterFunctionsField] `json:"field,required"`
	Type  param.Field[SchemaFilterFunctionsType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterFunctionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterFunctionsParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterFunctionsParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterFunctionsParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterFunctionsParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterFunctionsField string

const (
	SchemaFilterFunctionsFieldLastName SchemaFilterFunctionsField = "last_name"
	SchemaFilterFunctionsFieldType     SchemaFilterFunctionsField = "type"
	SchemaFilterFunctionsFieldUuid     SchemaFilterFunctionsField = "uuid"
)

func (r SchemaFilterFunctionsField) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsFieldLastName, SchemaFilterFunctionsFieldType, SchemaFilterFunctionsFieldUuid:
		return true
	}
	return false
}

type SchemaFilterFunctionsType string

const (
	SchemaFilterFunctionsTypeFunctions SchemaFilterFunctionsType = "functions"
)

func (r SchemaFilterFunctionsType) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsTypeFunctions:
		return true
	}
	return false
}

type SchemaFilterFunctionsCallsParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterFunctionsCallsField] `json:"field,required"`
	Type  param.Field[SchemaFilterFunctionsCallsType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterFunctionsCallsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterFunctionsCallsParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterFunctionsCallsParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterFunctionsCallsParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterFunctionsCallsParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterFunctionsCallsField string

const (
	SchemaFilterFunctionsCallsFieldLastName SchemaFilterFunctionsCallsField = "last_name"
	SchemaFilterFunctionsCallsFieldType     SchemaFilterFunctionsCallsField = "type"
	SchemaFilterFunctionsCallsFieldUuid     SchemaFilterFunctionsCallsField = "uuid"
)

func (r SchemaFilterFunctionsCallsField) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsCallsFieldLastName, SchemaFilterFunctionsCallsFieldType, SchemaFilterFunctionsCallsFieldUuid:
		return true
	}
	return false
}

type SchemaFilterFunctionsCallsType string

const (
	SchemaFilterFunctionsCallsTypeFunctionsCalls SchemaFilterFunctionsCallsType = "functions_calls"
)

func (r SchemaFilterFunctionsCallsType) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsCallsTypeFunctionsCalls:
		return true
	}
	return false
}

type SchemaFilterFunctionsTrafficParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterFunctionsTrafficField] `json:"field,required"`
	Type  param.Field[SchemaFilterFunctionsTrafficType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterFunctionsTrafficParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterFunctionsTrafficParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterFunctionsTrafficParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterFunctionsTrafficParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterFunctionsTrafficParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterFunctionsTrafficField string

const (
	SchemaFilterFunctionsTrafficFieldLastName SchemaFilterFunctionsTrafficField = "last_name"
	SchemaFilterFunctionsTrafficFieldType     SchemaFilterFunctionsTrafficField = "type"
	SchemaFilterFunctionsTrafficFieldUuid     SchemaFilterFunctionsTrafficField = "uuid"
)

func (r SchemaFilterFunctionsTrafficField) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsTrafficFieldLastName, SchemaFilterFunctionsTrafficFieldType, SchemaFilterFunctionsTrafficFieldUuid:
		return true
	}
	return false
}

type SchemaFilterFunctionsTrafficType string

const (
	SchemaFilterFunctionsTrafficTypeFunctionsTraffic SchemaFilterFunctionsTrafficType = "functions_traffic"
)

func (r SchemaFilterFunctionsTrafficType) IsKnown() bool {
	switch r {
	case SchemaFilterFunctionsTrafficTypeFunctionsTraffic:
		return true
	}
	return false
}

type SchemaFilterImageParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterImageField] `json:"field,required"`
	Type  param.Field[SchemaFilterImageType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterImageParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterImageParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterImageParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterImageParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterImageField string

const (
	SchemaFilterImageFieldLastName SchemaFilterImageField = "last_name"
	SchemaFilterImageFieldLastSize SchemaFilterImageField = "last_size"
	SchemaFilterImageFieldType     SchemaFilterImageField = "type"
	SchemaFilterImageFieldUuid     SchemaFilterImageField = "uuid"
)

func (r SchemaFilterImageField) IsKnown() bool {
	switch r {
	case SchemaFilterImageFieldLastName, SchemaFilterImageFieldLastSize, SchemaFilterImageFieldType, SchemaFilterImageFieldUuid:
		return true
	}
	return false
}

type SchemaFilterImageType string

const (
	SchemaFilterImageTypeImage SchemaFilterImageType = "image"
)

func (r SchemaFilterImageType) IsKnown() bool {
	switch r {
	case SchemaFilterImageTypeImage:
		return true
	}
	return false
}

type SchemaFilterInferenceParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterInferenceField] `json:"field,required"`
	Type  param.Field[SchemaFilterInferenceType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterInferenceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterInferenceParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterInferenceParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterInferenceParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterInferenceParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterInferenceField string

const (
	SchemaFilterInferenceFieldFlavor   SchemaFilterInferenceField = "flavor"
	SchemaFilterInferenceFieldLastName SchemaFilterInferenceField = "last_name"
	SchemaFilterInferenceFieldType     SchemaFilterInferenceField = "type"
	SchemaFilterInferenceFieldUuid     SchemaFilterInferenceField = "uuid"
)

func (r SchemaFilterInferenceField) IsKnown() bool {
	switch r {
	case SchemaFilterInferenceFieldFlavor, SchemaFilterInferenceFieldLastName, SchemaFilterInferenceFieldType, SchemaFilterInferenceFieldUuid:
		return true
	}
	return false
}

type SchemaFilterInferenceType string

const (
	SchemaFilterInferenceTypeInference SchemaFilterInferenceType = "inference"
)

func (r SchemaFilterInferenceType) IsKnown() bool {
	switch r {
	case SchemaFilterInferenceTypeInference:
		return true
	}
	return false
}

type SchemaFilterInstanceParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterInstanceEnum] `json:"field,required"`
	Type  param.Field[SchemaFilterInstanceType] `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterInstanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterInstanceParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterInstanceParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterInstanceParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterInstanceParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

type SchemaFilterInstanceType string

const (
	SchemaFilterInstanceTypeInstance SchemaFilterInstanceType = "instance"
)

func (r SchemaFilterInstanceType) IsKnown() bool {
	switch r {
	case SchemaFilterInstanceTypeInstance:
		return true
	}
	return false
}

type SchemaFilterInstanceEnum string

const (
	SchemaFilterInstanceEnumFlavor   SchemaFilterInstanceEnum = "flavor"
	SchemaFilterInstanceEnumLastName SchemaFilterInstanceEnum = "last_name"
	SchemaFilterInstanceEnumType     SchemaFilterInstanceEnum = "type"
	SchemaFilterInstanceEnumUuid     SchemaFilterInstanceEnum = "uuid"
)

func (r SchemaFilterInstanceEnum) IsKnown() bool {
	switch r {
	case SchemaFilterInstanceEnumFlavor, SchemaFilterInstanceEnumLastName, SchemaFilterInstanceEnumType, SchemaFilterInstanceEnumUuid:
		return true
	}
	return false
}

type SchemaFilterLoadBalancerParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterLoadBalancerField] `json:"field,required"`
	Type  param.Field[SchemaFilterLoadBalancerType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterLoadBalancerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterLoadBalancerParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterLoadBalancerParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
}

func (r SchemaFilterLoadBalancerParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
}

func (r SchemaFilterLoadBalancerParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterLoadBalancerField string

const (
	SchemaFilterLoadBalancerFieldFlavor   SchemaFilterLoadBalancerField = "flavor"
	SchemaFilterLoadBalancerFieldLastName SchemaFilterLoadBalancerField = "last_name"
	SchemaFilterLoadBalancerFieldType     SchemaFilterLoadBalancerField = "type"
	SchemaFilterLoadBalancerFieldUuid     SchemaFilterLoadBalancerField = "uuid"
)

func (r SchemaFilterLoadBalancerField) IsKnown() bool {
	switch r {
	case SchemaFilterLoadBalancerFieldFlavor, SchemaFilterLoadBalancerFieldLastName, SchemaFilterLoadBalancerFieldType, SchemaFilterLoadBalancerFieldUuid:
		return true
	}
	return false
}

type SchemaFilterLoadBalancerType string

const (
	SchemaFilterLoadBalancerTypeLoadBalancer SchemaFilterLoadBalancerType = "load_balancer"
)

func (r SchemaFilterLoadBalancerType) IsKnown() bool {
	switch r {
	case SchemaFilterLoadBalancerTypeLoadBalancer:
		return true
	}
	return false
}

type SchemaFilterLogIndexParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterLogIndexField] `json:"field,required"`
	Type  param.Field[SchemaFilterLogIndexType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterLogIndexParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterLogIndexParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterLogIndexParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterLogIndexParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterLogIndexParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterLogIndexField string

const (
	SchemaFilterLogIndexFieldLastName SchemaFilterLogIndexField = "last_name"
	SchemaFilterLogIndexFieldLastSize SchemaFilterLogIndexField = "last_size"
	SchemaFilterLogIndexFieldType     SchemaFilterLogIndexField = "type"
	SchemaFilterLogIndexFieldUuid     SchemaFilterLogIndexField = "uuid"
)

func (r SchemaFilterLogIndexField) IsKnown() bool {
	switch r {
	case SchemaFilterLogIndexFieldLastName, SchemaFilterLogIndexFieldLastSize, SchemaFilterLogIndexFieldType, SchemaFilterLogIndexFieldUuid:
		return true
	}
	return false
}

type SchemaFilterLogIndexType string

const (
	SchemaFilterLogIndexTypeLogIndex SchemaFilterLogIndexType = "log_index"
)

func (r SchemaFilterLogIndexType) IsKnown() bool {
	switch r {
	case SchemaFilterLogIndexTypeLogIndex:
		return true
	}
	return false
}

type SchemaFilterSnapshotParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterSnapshotField] `json:"field,required"`
	Type  param.Field[SchemaFilterSnapshotType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterSnapshotParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterSnapshotParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterSnapshotParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterSnapshotParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterSnapshotParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterSnapshotField string

const (
	SchemaFilterSnapshotFieldLastName         SchemaFilterSnapshotField = "last_name"
	SchemaFilterSnapshotFieldLastSize         SchemaFilterSnapshotField = "last_size"
	SchemaFilterSnapshotFieldSourceVolumeUuid SchemaFilterSnapshotField = "source_volume_uuid"
	SchemaFilterSnapshotFieldType             SchemaFilterSnapshotField = "type"
	SchemaFilterSnapshotFieldUuid             SchemaFilterSnapshotField = "uuid"
	SchemaFilterSnapshotFieldVolumeType       SchemaFilterSnapshotField = "volume_type"
)

func (r SchemaFilterSnapshotField) IsKnown() bool {
	switch r {
	case SchemaFilterSnapshotFieldLastName, SchemaFilterSnapshotFieldLastSize, SchemaFilterSnapshotFieldSourceVolumeUuid, SchemaFilterSnapshotFieldType, SchemaFilterSnapshotFieldUuid, SchemaFilterSnapshotFieldVolumeType:
		return true
	}
	return false
}

type SchemaFilterSnapshotType string

const (
	SchemaFilterSnapshotTypeSnapshot SchemaFilterSnapshotType = "snapshot"
)

func (r SchemaFilterSnapshotType) IsKnown() bool {
	switch r {
	case SchemaFilterSnapshotTypeSnapshot:
		return true
	}
	return false
}

type SchemaFilterVolumeParam struct {
	// Field name to filter by
	Field param.Field[SchemaFilterVolumeField] `json:"field,required"`
	Type  param.Field[SchemaFilterVolumeType]  `json:"type,required"`
	// List of field values to filter
	Values param.Field[[]string] `json:"values,required"`
}

func (r SchemaFilterVolumeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SchemaFilterVolumeParam) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {}

func (r SchemaFilterVolumeParam) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {}

func (r SchemaFilterVolumeParam) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {}

func (r SchemaFilterVolumeParam) implementsCloudV1ReservationCostReportGetTotalsParamsSchemaFilterUnion() {
}

// Field name to filter by
type SchemaFilterVolumeField string

const (
	SchemaFilterVolumeFieldAttachedToVm SchemaFilterVolumeField = "attached_to_vm"
	SchemaFilterVolumeFieldLastName     SchemaFilterVolumeField = "last_name"
	SchemaFilterVolumeFieldLastSize     SchemaFilterVolumeField = "last_size"
	SchemaFilterVolumeFieldType         SchemaFilterVolumeField = "type"
	SchemaFilterVolumeFieldUuid         SchemaFilterVolumeField = "uuid"
	SchemaFilterVolumeFieldVolumeType   SchemaFilterVolumeField = "volume_type"
)

func (r SchemaFilterVolumeField) IsKnown() bool {
	switch r {
	case SchemaFilterVolumeFieldAttachedToVm, SchemaFilterVolumeFieldLastName, SchemaFilterVolumeFieldLastSize, SchemaFilterVolumeFieldType, SchemaFilterVolumeFieldUuid, SchemaFilterVolumeFieldVolumeType:
		return true
	}
	return false
}

type SchemaFilterVolumeType string

const (
	SchemaFilterVolumeTypeVolume SchemaFilterVolumeType = "volume"
)

func (r SchemaFilterVolumeType) IsKnown() bool {
	switch r {
	case SchemaFilterVolumeTypeVolume:
		return true
	}
	return false
}

type SortingDirections string

const (
	SortingDirectionsAsc  SortingDirections = "asc"
	SortingDirectionsDesc SortingDirections = "desc"
)

func (r SortingDirections) IsKnown() bool {
	switch r {
	case SortingDirectionsAsc, SortingDirectionsDesc:
		return true
	}
	return false
}

// Represents a filter consisting of multiple tag filtering conditions.
type TagsFilterParam struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions param.Field[[]TagsFilterConditionParam] `json:"conditions,required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	ConditionType param.Field[TagsFilterConditionType] `json:"condition_type"`
}

func (r TagsFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TagsFilterConditionParam struct {
	// The name of the tag to filter (e.g., 'os_version').
	Key param.Field[string] `json:"key"`
	// Determines how strictly the tag value must match the specified value. If true,
	// the tag value must exactly match the given value. If false, a less strict match
	// (e.g., partial or case-insensitive match) may be applied.
	Strict param.Field[bool] `json:"strict"`
	// The value of the tag to filter (e.g., '22.04').
	Value param.Field[string] `json:"value"`
}

func (r TagsFilterConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies whether conditions are combined using OR (default) or AND logic.
type TagsFilterConditionType string

const (
	TagsFilterConditionTypeAnd TagsFilterConditionType = "AND"
	TagsFilterConditionTypeOr  TagsFilterConditionType = "OR"
)

func (r TagsFilterConditionType) IsKnown() bool {
	switch r {
	case TagsFilterConditionTypeAnd, TagsFilterConditionTypeOr:
		return true
	}
	return false
}

type TotalAIClusterWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalAIClusterWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalAIClusterWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalAIClusterWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalAIClusterWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalAIClusterWithCostJSON `json:"-"`
}

// totalAIClusterWithCostJSON contains the JSON metadata for the struct
// [TotalAIClusterWithCost]
type totalAIClusterWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Flavor             apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalAIClusterWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalAIClusterWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalAIClusterWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalAIClusterWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalAIClusterWithCostBillingValueUnion interface {
	ImplementsTotalAIClusterWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalAIClusterWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalAIClusterWithCostBillingValueUnit string

const (
	TotalAIClusterWithCostBillingValueUnitMinutes TotalAIClusterWithCostBillingValueUnit = "minutes"
)

func (r TotalAIClusterWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalAIClusterWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalAIClusterWithCostType string

const (
	TotalAIClusterWithCostTypeAICluster TotalAIClusterWithCostType = "ai_cluster"
)

func (r TotalAIClusterWithCostType) IsKnown() bool {
	switch r {
	case TotalAIClusterWithCostTypeAICluster:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalAIClusterWithCostCostUnion interface {
	ImplementsTotalAIClusterWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalAIClusterWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalBackupWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalBackupWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalBackupWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                   `json:"region_id,required"`
	Type               TotalBackupWithCostType `json:"type,required"`
	BillingFeatureName string                  `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalBackupWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                  `json:"region_name,nullable"`
	JSON       totalBackupWithCostJSON `json:"-"`
}

// totalBackupWithCostJSON contains the JSON metadata for the struct
// [TotalBackupWithCost]
type totalBackupWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	LastSize           apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalBackupWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalBackupWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalBackupWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalBackupWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBackupWithCostBillingValueUnion interface {
	ImplementsTotalBackupWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBackupWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalBackupWithCostBillingValueUnit string

const (
	TotalBackupWithCostBillingValueUnitGbminutes TotalBackupWithCostBillingValueUnit = "gbminutes"
)

func (r TotalBackupWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalBackupWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalBackupWithCostType string

const (
	TotalBackupWithCostTypeBackup TotalBackupWithCostType = "backup"
)

func (r TotalBackupWithCostType) IsKnown() bool {
	switch r {
	case TotalBackupWithCostTypeBackup:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBackupWithCostCostUnion interface {
	ImplementsTotalBackupWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBackupWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalBaremetalWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalBaremetalWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalBaremetalWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalBaremetalWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalBaremetalWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalBaremetalWithCostJSON `json:"-"`
}

// totalBaremetalWithCostJSON contains the JSON metadata for the struct
// [TotalBaremetalWithCost]
type totalBaremetalWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Flavor             apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalBaremetalWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalBaremetalWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalBaremetalWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalBaremetalWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBaremetalWithCostBillingValueUnion interface {
	ImplementsTotalBaremetalWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBaremetalWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalBaremetalWithCostBillingValueUnit string

const (
	TotalBaremetalWithCostBillingValueUnitMinutes TotalBaremetalWithCostBillingValueUnit = "minutes"
)

func (r TotalBaremetalWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalBaremetalWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalBaremetalWithCostType string

const (
	TotalBaremetalWithCostTypeBaremetal TotalBaremetalWithCostType = "baremetal"
)

func (r TotalBaremetalWithCostType) IsKnown() bool {
	switch r {
	case TotalBaremetalWithCostTypeBaremetal:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBaremetalWithCostCostUnion interface {
	ImplementsTotalBaremetalWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBaremetalWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalBasicVmWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalBasicVmWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalBasicVmWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                    `json:"region_id,required"`
	Type               TotalBasicVmWithCostType `json:"type,required"`
	BillingFeatureName string                   `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalBasicVmWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                   `json:"region_name,nullable"`
	JSON       totalBasicVmWithCostJSON `json:"-"`
}

// totalBasicVmWithCostJSON contains the JSON metadata for the struct
// [TotalBasicVmWithCost]
type totalBasicVmWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Flavor             apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalBasicVmWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalBasicVmWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalBasicVmWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalBasicVmWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBasicVmWithCostBillingValueUnion interface {
	ImplementsTotalBasicVmWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBasicVmWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalBasicVmWithCostBillingValueUnit string

const (
	TotalBasicVmWithCostBillingValueUnitMinutes TotalBasicVmWithCostBillingValueUnit = "minutes"
)

func (r TotalBasicVmWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalBasicVmWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalBasicVmWithCostType string

const (
	TotalBasicVmWithCostTypeBasicVm TotalBasicVmWithCostType = "basic_vm"
)

func (r TotalBasicVmWithCostType) IsKnown() bool {
	switch r {
	case TotalBasicVmWithCostTypeBasicVm:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalBasicVmWithCostCostUnion interface {
	ImplementsTotalBasicVmWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalBasicVmWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalContainerWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalContainerWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalContainerWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalContainerWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalContainerWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalContainerWithCostJSON `json:"-"`
}

// totalContainerWithCostJSON contains the JSON metadata for the struct
// [TotalContainerWithCost]
type totalContainerWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalContainerWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalContainerWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalContainerWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalContainerWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalContainerWithCostBillingValueUnion interface {
	ImplementsTotalContainerWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalContainerWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalContainerWithCostBillingValueUnit string

const (
	TotalContainerWithCostBillingValueUnitGBs TotalContainerWithCostBillingValueUnit = "GBS"
)

func (r TotalContainerWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalContainerWithCostBillingValueUnitGBs:
		return true
	}
	return false
}

type TotalContainerWithCostType string

const (
	TotalContainerWithCostTypeContainers TotalContainerWithCostType = "containers"
)

func (r TotalContainerWithCostType) IsKnown() bool {
	switch r {
	case TotalContainerWithCostTypeContainers:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalContainerWithCostCostUnion interface {
	ImplementsTotalContainerWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalContainerWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalDbaasPostgresqlCPUWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalDbaasPostgresqlCPUWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalDbaasPostgresqlCPUWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                               `json:"region_id,required"`
	Type               TotalDbaasPostgresqlCPUWithCostType `json:"type,required"`
	BillingFeatureName string                              `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalDbaasPostgresqlCPUWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                              `json:"region_name,nullable"`
	JSON       totalDbaasPostgresqlCPUWithCostJSON `json:"-"`
}

// totalDbaasPostgresqlCPUWithCostJSON contains the JSON metadata for the struct
// [TotalDbaasPostgresqlCPUWithCost]
type totalDbaasPostgresqlCPUWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalDbaasPostgresqlCPUWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalDbaasPostgresqlCPUWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalDbaasPostgresqlCPUWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalDbaasPostgresqlCPUWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlCPUWithCostBillingValueUnion interface {
	ImplementsTotalDbaasPostgresqlCPUWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlCPUWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalDbaasPostgresqlCPUWithCostBillingValueUnit string

const (
	TotalDbaasPostgresqlCPUWithCostBillingValueUnitMinutes TotalDbaasPostgresqlCPUWithCostBillingValueUnit = "minutes"
)

func (r TotalDbaasPostgresqlCPUWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlCPUWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalDbaasPostgresqlCPUWithCostType string

const (
	TotalDbaasPostgresqlCPUWithCostTypeDbaasPostgresqlCPU TotalDbaasPostgresqlCPUWithCostType = "dbaas_postgresql_cpu"
)

func (r TotalDbaasPostgresqlCPUWithCostType) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlCPUWithCostTypeDbaasPostgresqlCPU:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlCPUWithCostCostUnion interface {
	ImplementsTotalDbaasPostgresqlCPUWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlCPUWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalDbaasPostgresqlMemoryWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalDbaasPostgresqlMemoryWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalDbaasPostgresqlMemoryWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                                  `json:"region_id,required"`
	Type               TotalDbaasPostgresqlMemoryWithCostType `json:"type,required"`
	BillingFeatureName string                                 `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalDbaasPostgresqlMemoryWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                 `json:"region_name,nullable"`
	JSON       totalDbaasPostgresqlMemoryWithCostJSON `json:"-"`
}

// totalDbaasPostgresqlMemoryWithCostJSON contains the JSON metadata for the struct
// [TotalDbaasPostgresqlMemoryWithCost]
type totalDbaasPostgresqlMemoryWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalDbaasPostgresqlMemoryWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalDbaasPostgresqlMemoryWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalDbaasPostgresqlMemoryWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalDbaasPostgresqlMemoryWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlMemoryWithCostBillingValueUnion interface {
	ImplementsTotalDbaasPostgresqlMemoryWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlMemoryWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalDbaasPostgresqlMemoryWithCostBillingValueUnit string

const (
	TotalDbaasPostgresqlMemoryWithCostBillingValueUnitGbminutes TotalDbaasPostgresqlMemoryWithCostBillingValueUnit = "gbminutes"
)

func (r TotalDbaasPostgresqlMemoryWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlMemoryWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalDbaasPostgresqlMemoryWithCostType string

const (
	TotalDbaasPostgresqlMemoryWithCostTypeDbaasPostgresqlMemory TotalDbaasPostgresqlMemoryWithCostType = "dbaas_postgresql_memory"
)

func (r TotalDbaasPostgresqlMemoryWithCostType) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlMemoryWithCostTypeDbaasPostgresqlMemory:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlMemoryWithCostCostUnion interface {
	ImplementsTotalDbaasPostgresqlMemoryWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlMemoryWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalDbaasPostgresqlPoolerWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalDbaasPostgresqlPoolerWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalDbaasPostgresqlPoolerWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                                  `json:"region_id,required"`
	Type               TotalDbaasPostgresqlPoolerWithCostType `json:"type,required"`
	BillingFeatureName string                                 `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalDbaasPostgresqlPoolerWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                 `json:"region_name,nullable"`
	JSON       totalDbaasPostgresqlPoolerWithCostJSON `json:"-"`
}

// totalDbaasPostgresqlPoolerWithCostJSON contains the JSON metadata for the struct
// [TotalDbaasPostgresqlPoolerWithCost]
type totalDbaasPostgresqlPoolerWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalDbaasPostgresqlPoolerWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalDbaasPostgresqlPoolerWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalDbaasPostgresqlPoolerWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalDbaasPostgresqlPoolerWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlPoolerWithCostBillingValueUnion interface {
	ImplementsTotalDbaasPostgresqlPoolerWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlPoolerWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalDbaasPostgresqlPoolerWithCostBillingValueUnit string

const (
	TotalDbaasPostgresqlPoolerWithCostBillingValueUnitMinutes TotalDbaasPostgresqlPoolerWithCostBillingValueUnit = "minutes"
)

func (r TotalDbaasPostgresqlPoolerWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlPoolerWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalDbaasPostgresqlPoolerWithCostType string

const (
	TotalDbaasPostgresqlPoolerWithCostTypeDbaasPostgresqlConnectionPooler TotalDbaasPostgresqlPoolerWithCostType = "dbaas_postgresql_connection_pooler"
)

func (r TotalDbaasPostgresqlPoolerWithCostType) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlPoolerWithCostTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlPoolerWithCostCostUnion interface {
	ImplementsTotalDbaasPostgresqlPoolerWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlPoolerWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalDbaasPostgresqlPublicNetworkWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                                         `json:"region_id,required"`
	Type               TotalDbaasPostgresqlPublicNetworkWithCostType `json:"type,required"`
	BillingFeatureName string                                        `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalDbaasPostgresqlPublicNetworkWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                        `json:"region_name,nullable"`
	JSON       totalDbaasPostgresqlPublicNetworkWithCostJSON `json:"-"`
}

// totalDbaasPostgresqlPublicNetworkWithCostJSON contains the JSON metadata for the
// struct [TotalDbaasPostgresqlPublicNetworkWithCost]
type totalDbaasPostgresqlPublicNetworkWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalDbaasPostgresqlPublicNetworkWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalDbaasPostgresqlPublicNetworkWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalDbaasPostgresqlPublicNetworkWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {
}

func (r TotalDbaasPostgresqlPublicNetworkWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnion interface {
	ImplementsTotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnit string

const (
	TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnitMinutes TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnit = "minutes"
)

func (r TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlPublicNetworkWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalDbaasPostgresqlPublicNetworkWithCostType string

const (
	TotalDbaasPostgresqlPublicNetworkWithCostTypeDbaasPostgresqlPublicNetwork TotalDbaasPostgresqlPublicNetworkWithCostType = "dbaas_postgresql_public_network"
)

func (r TotalDbaasPostgresqlPublicNetworkWithCostType) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlPublicNetworkWithCostTypeDbaasPostgresqlPublicNetwork:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlPublicNetworkWithCostCostUnion interface {
	ImplementsTotalDbaasPostgresqlPublicNetworkWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlPublicNetworkWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalDbaasPostgresqlVolumeWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalDbaasPostgresqlVolumeWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalDbaasPostgresqlVolumeWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                  `json:"region_id,required"`
	Type     TotalDbaasPostgresqlVolumeWithCostType `json:"type,required"`
	// Type of the volume
	VolumeType         string `json:"volume_type,required"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalDbaasPostgresqlVolumeWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                 `json:"region_name,nullable"`
	JSON       totalDbaasPostgresqlVolumeWithCostJSON `json:"-"`
}

// totalDbaasPostgresqlVolumeWithCostJSON contains the JSON metadata for the struct
// [TotalDbaasPostgresqlVolumeWithCost]
type totalDbaasPostgresqlVolumeWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	VolumeType         apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalDbaasPostgresqlVolumeWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalDbaasPostgresqlVolumeWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalDbaasPostgresqlVolumeWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalDbaasPostgresqlVolumeWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlVolumeWithCostBillingValueUnion interface {
	ImplementsTotalDbaasPostgresqlVolumeWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlVolumeWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalDbaasPostgresqlVolumeWithCostBillingValueUnit string

const (
	TotalDbaasPostgresqlVolumeWithCostBillingValueUnitGbminutes TotalDbaasPostgresqlVolumeWithCostBillingValueUnit = "gbminutes"
)

func (r TotalDbaasPostgresqlVolumeWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlVolumeWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalDbaasPostgresqlVolumeWithCostType string

const (
	TotalDbaasPostgresqlVolumeWithCostTypeDbaasPostgresqlVolume TotalDbaasPostgresqlVolumeWithCostType = "dbaas_postgresql_volume"
)

func (r TotalDbaasPostgresqlVolumeWithCostType) IsKnown() bool {
	switch r {
	case TotalDbaasPostgresqlVolumeWithCostTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalDbaasPostgresqlVolumeWithCostCostUnion interface {
	ImplementsTotalDbaasPostgresqlVolumeWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalDbaasPostgresqlVolumeWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalEgressTrafficWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalEgressTrafficWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalEgressTrafficWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                          `json:"region_id,required"`
	Type               TotalEgressTrafficWithCostType `json:"type,required"`
	BillingFeatureName string                         `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalEgressTrafficWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                         `json:"region_name,nullable"`
	JSON       totalEgressTrafficWithCostJSON `json:"-"`
}

// totalEgressTrafficWithCostJSON contains the JSON metadata for the struct
// [TotalEgressTrafficWithCost]
type totalEgressTrafficWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	InstanceType       apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalEgressTrafficWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalEgressTrafficWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalEgressTrafficWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalEgressTrafficWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalEgressTrafficWithCostBillingValueUnion interface {
	ImplementsTotalEgressTrafficWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalEgressTrafficWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalEgressTrafficWithCostBillingValueUnit string

const (
	TotalEgressTrafficWithCostBillingValueUnitBytes TotalEgressTrafficWithCostBillingValueUnit = "bytes"
)

func (r TotalEgressTrafficWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalEgressTrafficWithCostBillingValueUnitBytes:
		return true
	}
	return false
}

type TotalEgressTrafficWithCostType string

const (
	TotalEgressTrafficWithCostTypeEgressTraffic TotalEgressTrafficWithCostType = "egress_traffic"
)

func (r TotalEgressTrafficWithCostType) IsKnown() bool {
	switch r {
	case TotalEgressTrafficWithCostTypeEgressTraffic:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalEgressTrafficWithCostCostUnion interface {
	ImplementsTotalEgressTrafficWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalEgressTrafficWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalExternalIPWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalExternalIPWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalExternalIPWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                       `json:"region_id,required"`
	Type               TotalExternalIPWithCostType `json:"type,required"`
	BillingFeatureName string                      `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalExternalIPWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                      `json:"region_name,nullable"`
	JSON       totalExternalIPWithCostJSON `json:"-"`
}

// totalExternalIPWithCostJSON contains the JSON metadata for the struct
// [TotalExternalIPWithCost]
type totalExternalIPWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalExternalIPWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalExternalIPWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalExternalIPWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalExternalIPWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalExternalIPWithCostBillingValueUnion interface {
	ImplementsTotalExternalIPWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalExternalIPWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalExternalIPWithCostBillingValueUnit string

const (
	TotalExternalIPWithCostBillingValueUnitMinutes TotalExternalIPWithCostBillingValueUnit = "minutes"
)

func (r TotalExternalIPWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalExternalIPWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalExternalIPWithCostType string

const (
	TotalExternalIPWithCostTypeExternalIP TotalExternalIPWithCostType = "external_ip"
)

func (r TotalExternalIPWithCostType) IsKnown() bool {
	switch r {
	case TotalExternalIPWithCostTypeExternalIP:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalExternalIPWithCostCostUnion interface {
	ImplementsTotalExternalIPWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalExternalIPWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalFileShareWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalFileShareWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalFileShareWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalFileShareWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalFileShareWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalFileShareWithCostJSON `json:"-"`
}

// totalFileShareWithCostJSON contains the JSON metadata for the struct
// [TotalFileShareWithCost]
type totalFileShareWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FileShareType      apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalFileShareWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalFileShareWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalFileShareWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalFileShareWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFileShareWithCostBillingValueUnion interface {
	ImplementsTotalFileShareWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFileShareWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalFileShareWithCostBillingValueUnit string

const (
	TotalFileShareWithCostBillingValueUnitGbminutes TotalFileShareWithCostBillingValueUnit = "gbminutes"
)

func (r TotalFileShareWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalFileShareWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalFileShareWithCostType string

const (
	TotalFileShareWithCostTypeFileShare TotalFileShareWithCostType = "file_share"
)

func (r TotalFileShareWithCostType) IsKnown() bool {
	switch r {
	case TotalFileShareWithCostTypeFileShare:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFileShareWithCostCostUnion interface {
	ImplementsTotalFileShareWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFileShareWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalFloatingIPWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalFloatingIPWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalFloatingIPWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                       `json:"region_id,required"`
	Type               TotalFloatingIPWithCostType `json:"type,required"`
	BillingFeatureName string                      `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalFloatingIPWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                      `json:"region_name,nullable"`
	JSON       totalFloatingIPWithCostJSON `json:"-"`
}

// totalFloatingIPWithCostJSON contains the JSON metadata for the struct
// [TotalFloatingIPWithCost]
type totalFloatingIPWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalFloatingIPWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalFloatingIPWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalFloatingIPWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalFloatingIPWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFloatingIPWithCostBillingValueUnion interface {
	ImplementsTotalFloatingIPWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFloatingIPWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalFloatingIPWithCostBillingValueUnit string

const (
	TotalFloatingIPWithCostBillingValueUnitMinutes TotalFloatingIPWithCostBillingValueUnit = "minutes"
)

func (r TotalFloatingIPWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalFloatingIPWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalFloatingIPWithCostType string

const (
	TotalFloatingIPWithCostTypeFloatingip TotalFloatingIPWithCostType = "floatingip"
)

func (r TotalFloatingIPWithCostType) IsKnown() bool {
	switch r {
	case TotalFloatingIPWithCostTypeFloatingip:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFloatingIPWithCostCostUnion interface {
	ImplementsTotalFloatingIPWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFloatingIPWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalFunctionCallsWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalFunctionCallsWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalFunctionCallsWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                          `json:"region_id,required"`
	Type               TotalFunctionCallsWithCostType `json:"type,required"`
	BillingFeatureName string                         `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalFunctionCallsWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                         `json:"region_name,nullable"`
	JSON       totalFunctionCallsWithCostJSON `json:"-"`
}

// totalFunctionCallsWithCostJSON contains the JSON metadata for the struct
// [TotalFunctionCallsWithCost]
type totalFunctionCallsWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalFunctionCallsWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalFunctionCallsWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalFunctionCallsWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalFunctionCallsWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionCallsWithCostBillingValueUnion interface {
	ImplementsTotalFunctionCallsWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionCallsWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalFunctionCallsWithCostBillingValueUnit string

const (
	TotalFunctionCallsWithCostBillingValueUnitMls TotalFunctionCallsWithCostBillingValueUnit = "MLS"
)

func (r TotalFunctionCallsWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalFunctionCallsWithCostBillingValueUnitMls:
		return true
	}
	return false
}

type TotalFunctionCallsWithCostType string

const (
	TotalFunctionCallsWithCostTypeFunctionsCalls TotalFunctionCallsWithCostType = "functions_calls"
)

func (r TotalFunctionCallsWithCostType) IsKnown() bool {
	switch r {
	case TotalFunctionCallsWithCostTypeFunctionsCalls:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionCallsWithCostCostUnion interface {
	ImplementsTotalFunctionCallsWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionCallsWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalFunctionEgressTrafficWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalFunctionEgressTrafficWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalFunctionEgressTrafficWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                                  `json:"region_id,required"`
	Type               TotalFunctionEgressTrafficWithCostType `json:"type,required"`
	BillingFeatureName string                                 `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalFunctionEgressTrafficWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                 `json:"region_name,nullable"`
	JSON       totalFunctionEgressTrafficWithCostJSON `json:"-"`
}

// totalFunctionEgressTrafficWithCostJSON contains the JSON metadata for the struct
// [TotalFunctionEgressTrafficWithCost]
type totalFunctionEgressTrafficWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalFunctionEgressTrafficWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalFunctionEgressTrafficWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalFunctionEgressTrafficWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalFunctionEgressTrafficWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionEgressTrafficWithCostBillingValueUnion interface {
	ImplementsTotalFunctionEgressTrafficWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionEgressTrafficWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalFunctionEgressTrafficWithCostBillingValueUnit string

const (
	TotalFunctionEgressTrafficWithCostBillingValueUnitGB TotalFunctionEgressTrafficWithCostBillingValueUnit = "GB"
)

func (r TotalFunctionEgressTrafficWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalFunctionEgressTrafficWithCostBillingValueUnitGB:
		return true
	}
	return false
}

type TotalFunctionEgressTrafficWithCostType string

const (
	TotalFunctionEgressTrafficWithCostTypeFunctionsTraffic TotalFunctionEgressTrafficWithCostType = "functions_traffic"
)

func (r TotalFunctionEgressTrafficWithCostType) IsKnown() bool {
	switch r {
	case TotalFunctionEgressTrafficWithCostTypeFunctionsTraffic:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionEgressTrafficWithCostCostUnion interface {
	ImplementsTotalFunctionEgressTrafficWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionEgressTrafficWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalFunctionsWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalFunctionsWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalFunctionsWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalFunctionsWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalFunctionsWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalFunctionsWithCostJSON `json:"-"`
}

// totalFunctionsWithCostJSON contains the JSON metadata for the struct
// [TotalFunctionsWithCost]
type totalFunctionsWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalFunctionsWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalFunctionsWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalFunctionsWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalFunctionsWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionsWithCostBillingValueUnion interface {
	ImplementsTotalFunctionsWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionsWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalFunctionsWithCostBillingValueUnit string

const (
	TotalFunctionsWithCostBillingValueUnitGBs TotalFunctionsWithCostBillingValueUnit = "GBS"
)

func (r TotalFunctionsWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalFunctionsWithCostBillingValueUnitGBs:
		return true
	}
	return false
}

type TotalFunctionsWithCostType string

const (
	TotalFunctionsWithCostTypeFunctions TotalFunctionsWithCostType = "functions"
)

func (r TotalFunctionsWithCostType) IsKnown() bool {
	switch r {
	case TotalFunctionsWithCostTypeFunctions:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalFunctionsWithCostCostUnion interface {
	ImplementsTotalFunctionsWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalFunctionsWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalImagesWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalImagesWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalImagesWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                   `json:"region_id,required"`
	Type               TotalImagesWithCostType `json:"type,required"`
	BillingFeatureName string                  `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalImagesWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                  `json:"region_name,nullable"`
	JSON       totalImagesWithCostJSON `json:"-"`
}

// totalImagesWithCostJSON contains the JSON metadata for the struct
// [TotalImagesWithCost]
type totalImagesWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalImagesWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalImagesWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalImagesWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalImagesWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalImagesWithCostBillingValueUnion interface {
	ImplementsTotalImagesWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalImagesWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalImagesWithCostBillingValueUnit string

const (
	TotalImagesWithCostBillingValueUnitGbminutes TotalImagesWithCostBillingValueUnit = "gbminutes"
)

func (r TotalImagesWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalImagesWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalImagesWithCostType string

const (
	TotalImagesWithCostTypeImage TotalImagesWithCostType = "image"
)

func (r TotalImagesWithCostType) IsKnown() bool {
	switch r {
	case TotalImagesWithCostTypeImage:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalImagesWithCostCostUnion interface {
	ImplementsTotalImagesWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalImagesWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalInferenceWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalInferenceWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                      `json:"region_id,required"`
	Type               TotalInferenceWithCostType `json:"type,required"`
	BillingFeatureName string                     `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalInferenceWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                     `json:"region_name,nullable"`
	JSON       totalInferenceWithCostJSON `json:"-"`
}

// totalInferenceWithCostJSON contains the JSON metadata for the struct
// [TotalInferenceWithCost]
type totalInferenceWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalInferenceWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalInferenceWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalInferenceWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalInferenceWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalInferenceWithCostBillingValueUnion interface {
	ImplementsTotalInferenceWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalInferenceWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalInferenceWithCostType string

const (
	TotalInferenceWithCostTypeInference TotalInferenceWithCostType = "inference"
)

func (r TotalInferenceWithCostType) IsKnown() bool {
	switch r {
	case TotalInferenceWithCostTypeInference:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalInferenceWithCostCostUnion interface {
	ImplementsTotalInferenceWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalInferenceWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalInstanceWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalInstanceWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalInstanceWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                     `json:"region_id,required"`
	Type               TotalInstanceWithCostType `json:"type,required"`
	BillingFeatureName string                    `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalInstanceWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                    `json:"region_name,nullable"`
	JSON       totalInstanceWithCostJSON `json:"-"`
}

// totalInstanceWithCostJSON contains the JSON metadata for the struct
// [TotalInstanceWithCost]
type totalInstanceWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Flavor             apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalInstanceWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalInstanceWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalInstanceWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalInstanceWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalInstanceWithCostBillingValueUnion interface {
	ImplementsTotalInstanceWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalInstanceWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalInstanceWithCostBillingValueUnit string

const (
	TotalInstanceWithCostBillingValueUnitMinutes TotalInstanceWithCostBillingValueUnit = "minutes"
)

func (r TotalInstanceWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalInstanceWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalInstanceWithCostType string

const (
	TotalInstanceWithCostTypeInstance TotalInstanceWithCostType = "instance"
)

func (r TotalInstanceWithCostType) IsKnown() bool {
	switch r {
	case TotalInstanceWithCostTypeInstance:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalInstanceWithCostCostUnion interface {
	ImplementsTotalInstanceWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalInstanceWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalLoadBalancerWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalLoadBalancerWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalLoadBalancerWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                         `json:"region_id,required"`
	Type               TotalLoadBalancerWithCostType `json:"type,required"`
	BillingFeatureName string                        `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalLoadBalancerWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                        `json:"region_name,nullable"`
	JSON       totalLoadBalancerWithCostJSON `json:"-"`
}

// totalLoadBalancerWithCostJSON contains the JSON metadata for the struct
// [TotalLoadBalancerWithCost]
type totalLoadBalancerWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Flavor             apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalLoadBalancerWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalLoadBalancerWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalLoadBalancerWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalLoadBalancerWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalLoadBalancerWithCostBillingValueUnion interface {
	ImplementsTotalLoadBalancerWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalLoadBalancerWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalLoadBalancerWithCostBillingValueUnit string

const (
	TotalLoadBalancerWithCostBillingValueUnitMinutes TotalLoadBalancerWithCostBillingValueUnit = "minutes"
)

func (r TotalLoadBalancerWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalLoadBalancerWithCostBillingValueUnitMinutes:
		return true
	}
	return false
}

type TotalLoadBalancerWithCostType string

const (
	TotalLoadBalancerWithCostTypeLoadBalancer TotalLoadBalancerWithCostType = "load_balancer"
)

func (r TotalLoadBalancerWithCostType) IsKnown() bool {
	switch r {
	case TotalLoadBalancerWithCostTypeLoadBalancer:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalLoadBalancerWithCostCostUnion interface {
	ImplementsTotalLoadBalancerWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalLoadBalancerWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalLogIndexWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalLogIndexWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalLogIndexWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID           int64                     `json:"region_id,required"`
	Type               TotalLogIndexWithCostType `json:"type,required"`
	BillingFeatureName string                    `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalLogIndexWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                    `json:"region_name,nullable"`
	JSON       totalLogIndexWithCostJSON `json:"-"`
}

// totalLogIndexWithCostJSON contains the JSON metadata for the struct
// [TotalLogIndexWithCost]
type totalLogIndexWithCostJSON struct {
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
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalLogIndexWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalLogIndexWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalLogIndexWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalLogIndexWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalLogIndexWithCostBillingValueUnion interface {
	ImplementsTotalLogIndexWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalLogIndexWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalLogIndexWithCostBillingValueUnit string

const (
	TotalLogIndexWithCostBillingValueUnitGbminutes TotalLogIndexWithCostBillingValueUnit = "gbminutes"
)

func (r TotalLogIndexWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalLogIndexWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalLogIndexWithCostType string

const (
	TotalLogIndexWithCostTypeLogIndex TotalLogIndexWithCostType = "log_index"
)

func (r TotalLogIndexWithCostType) IsKnown() bool {
	switch r {
	case TotalLogIndexWithCostTypeLogIndex:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalLogIndexWithCostCostUnion interface {
	ImplementsTotalLogIndexWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalLogIndexWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalSnapshotWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalSnapshotWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalSnapshotWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                     `json:"region_id,required"`
	Type     TotalSnapshotWithCostType `json:"type,required"`
	// Type of the volume
	VolumeType         string `json:"volume_type,required"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalSnapshotWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                    `json:"region_name,nullable"`
	JSON       totalSnapshotWithCostJSON `json:"-"`
}

// totalSnapshotWithCostJSON contains the JSON metadata for the struct
// [TotalSnapshotWithCost]
type totalSnapshotWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	VolumeType         apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalSnapshotWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalSnapshotWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalSnapshotWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalSnapshotWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalSnapshotWithCostBillingValueUnion interface {
	ImplementsTotalSnapshotWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalSnapshotWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalSnapshotWithCostBillingValueUnit string

const (
	TotalSnapshotWithCostBillingValueUnitGbminutes TotalSnapshotWithCostBillingValueUnit = "gbminutes"
)

func (r TotalSnapshotWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalSnapshotWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalSnapshotWithCostType string

const (
	TotalSnapshotWithCostTypeSnapshot TotalSnapshotWithCostType = "snapshot"
)

func (r TotalSnapshotWithCostType) IsKnown() bool {
	switch r {
	case TotalSnapshotWithCostTypeSnapshot:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalSnapshotWithCostCostUnion interface {
	ImplementsTotalSnapshotWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalSnapshotWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TotalVolumeWithCost struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue TotalVolumeWithCostBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit TotalVolumeWithCostBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                   `json:"region_id,required"`
	Type     TotalVolumeWithCostType `json:"type,required"`
	// Type of the volume
	VolumeType         string `json:"volume_type,required"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost TotalVolumeWithCostCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                  `json:"region_name,nullable"`
	JSON       totalVolumeWithCostJSON `json:"-"`
}

// totalVolumeWithCostJSON contains the JSON metadata for the struct
// [TotalVolumeWithCost]
type totalVolumeWithCostJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	VolumeType         apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TotalVolumeWithCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalVolumeWithCostJSON) RawJSON() string {
	return r.raw
}

func (r TotalVolumeWithCost) implementsCloudV1CostReportGetTotalsResponseResult() {}

func (r TotalVolumeWithCost) implementsCloudV1ReservationCostReportGetTotalsResponseResult() {}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalVolumeWithCostBillingValueUnion interface {
	ImplementsTotalVolumeWithCostBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalVolumeWithCostBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type TotalVolumeWithCostBillingValueUnit string

const (
	TotalVolumeWithCostBillingValueUnitGbminutes TotalVolumeWithCostBillingValueUnit = "gbminutes"
)

func (r TotalVolumeWithCostBillingValueUnit) IsKnown() bool {
	switch r {
	case TotalVolumeWithCostBillingValueUnitGbminutes:
		return true
	}
	return false
}

type TotalVolumeWithCostType string

const (
	TotalVolumeWithCostTypeVolume TotalVolumeWithCostType = "volume"
)

func (r TotalVolumeWithCostType) IsKnown() bool {
	switch r {
	case TotalVolumeWithCostTypeVolume:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type TotalVolumeWithCostCostUnion interface {
	ImplementsTotalVolumeWithCostCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TotalVolumeWithCostCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponse struct {
	// Count of all the resources
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	PriceStatus PriceDisplayStatus                            `json:"price_status,required"`
	Results     []CloudV1CostReportGetResourcesResponseResult `json:"results,required"`
	JSON        cloudV1CostReportGetResourcesResponseJSON     `json:"-"`
}

// cloudV1CostReportGetResourcesResponseJSON contains the JSON metadata for the
// struct [CloudV1CostReportGetResourcesResponse]
type cloudV1CostReportGetResourcesResponseJSON struct {
	Count       apijson.Field
	PriceStatus apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CostReportGetResourcesResponseResult struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// This field can have the runtime type of
	// [CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnion].
	BillingValue interface{} `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                            `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsType `json:"type,required"`
	// ID of the VM the IP is attached to
	AttachedToVm       string `json:"attached_to_vm,nullable" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// This field can have the runtime type of
	// [CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerCostUnion],
	// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerCostUnion].
	Cost interface{} `json:"cost"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Type of the file share
	FileShareType string `json:"file_share_type"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor"`
	// Name of the instance
	InstanceName string `json:"instance_name,nullable"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type"`
	// IP address
	IPAddress string `json:"ip_address" format:"ipvanyaddress"`
	// Name of the AI cluster
	LastName string `json:"last_name"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size"`
	// ID of the network the IP is attached to
	NetworkID string `json:"network_id" format:"uuid"`
	// ID of the port the traffic is associated with
	PortID string `json:"port_id" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// ID of the backup schedule
	ScheduleID string `json:"schedule_id" format:"uuid"`
	// Unit of size
	SizeUnit string `json:"size_unit"`
	// UUID of the source volume
	SourceVolumeUuid string `json:"source_volume_uuid" format:"uuid"`
	// ID of the subnet the IP is attached to
	SubnetID string `json:"subnet_id" format:"uuid"`
	// This field can have the runtime type of [[]interface{}].
	Tags interface{} `json:"tags"`
	// UUID of the AI cluster
	Uuid string `json:"uuid,nullable" format:"uuid"`
	// ID of the bare metal server the traffic is associated with
	VmID string `json:"vm_id" format:"uuid"`
	// Type of the volume
	VolumeType string                                          `json:"volume_type"`
	JSON       cloudV1CostReportGetResourcesResponseResultJSON `json:"-"`
	union      CloudV1CostReportGetResourcesResponseResultsUnion
}

// cloudV1CostReportGetResourcesResponseResultJSON contains the JSON metadata for
// the struct [CloudV1CostReportGetResourcesResponseResult]
type cloudV1CostReportGetResourcesResponseResultJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	AttachedToVm       apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	FileShareType      apijson.Field
	Flavor             apijson.Field
	InstanceName       apijson.Field
	InstanceType       apijson.Field
	IPAddress          apijson.Field
	LastName           apijson.Field
	LastSize           apijson.Field
	NetworkID          apijson.Field
	PortID             apijson.Field
	RegionName         apijson.Field
	ScheduleID         apijson.Field
	SizeUnit           apijson.Field
	SourceVolumeUuid   apijson.Field
	SubnetID           apijson.Field
	Tags               apijson.Field
	Uuid               apijson.Field
	VmID               apijson.Field
	VolumeType         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r cloudV1CostReportGetResourcesResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1CostReportGetResourcesResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1CostReportGetResourcesResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1CostReportGetResourcesResponseResultsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer].
func (r CloudV1CostReportGetResourcesResponseResult) AsUnion() CloudV1CostReportGetResourcesResponseResultsUnion {
	return r.union
}

// Union satisfied by
// [CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer],
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer]
// or
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer].
type CloudV1CostReportGetResourcesResponseResultsUnion interface {
	implementsCloudV1CostReportGetResourcesResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer{}),
			DiscriminatorValue: "ai_cluster",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer{}),
			DiscriminatorValue: "baremetal",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer{}),
			DiscriminatorValue: "basic_vm",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer{}),
			DiscriminatorValue: "backup",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer{}),
			DiscriminatorValue: "containers",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer{}),
			DiscriminatorValue: "egress_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer{}),
			DiscriminatorValue: "external_ip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer{}),
			DiscriminatorValue: "file_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer{}),
			DiscriminatorValue: "floatingip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer{}),
			DiscriminatorValue: "functions",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer{}),
			DiscriminatorValue: "functions_calls",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer{}),
			DiscriminatorValue: "functions_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer{}),
			DiscriminatorValue: "inference",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer{}),
			DiscriminatorValue: "instance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer{}),
			DiscriminatorValue: "load_balancer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer{}),
			DiscriminatorValue: "log_index",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer{}),
			DiscriminatorValue: "snapshot",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer{}),
			DiscriminatorValue: "volume",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_connection_pooler",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_memory",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_public_network",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_cpu",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_volume",
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the AI cluster
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
	RegionID int64                                                                               `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerType `json:"type,required"`
	// UUID of the AI cluster
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                       `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerTypeAICluster CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerType = "ai_cluster"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerTypeAICluster:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceAIClusterWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                               `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerType `json:"type,required"`
	// UUID of the bare metal server
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                       `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerTypeBaremetal CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerType = "baremetal"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerTypeBaremetal:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBaremetalWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                             `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerType `json:"type,required"`
	// UUID of the basic VM
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                     `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerTypeBasicVm CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerType = "basic_vm"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerTypeBasicVm:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBasicVmWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SourceVolumeUuid string                                                                           `json:"source_volume_uuid,required" format:"uuid"`
	Type             CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerType `json:"type,required"`
	// UUID of the backup
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                           `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	ScheduleID         apijson.Field
	SourceVolumeUuid   apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerTypeBackup CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerType = "backup"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerTypeBackup:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceBackupWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                               `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerType `json:"type,required"`
	// UUID of the container
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                              `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnitGBs CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnit = "GBS"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerTypeContainers CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerType = "containers"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerTypeContainers:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceContainerWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type,required"`
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
	SizeUnit string                                                                                  `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerType `json:"type,required"`
	// ID of the bare metal server the traffic is associated with
	VmID               string `json:"vm_id,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Name of the instance
	InstanceName string `json:"instance_name,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                           `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	InstanceType       apijson.Field
	LastSeen           apijson.Field
	PortID             apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	VmID               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	InstanceName       apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnitBytes CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnit = "bytes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerBillingValueUnitBytes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerTypeEgressTraffic CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerType = "egress_traffic"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerTypeEgressTraffic:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceEgressTrafficWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer struct {
	// ID of the VM the IP is attached to
	AttachedToVm string `json:"attached_to_vm,required,nullable" format:"uuid"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SubnetID           string                                                                               `json:"subnet_id,required" format:"uuid"`
	Type               CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerType `json:"type,required"`
	BillingFeatureName string                                                                               `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                               `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerJSON struct {
	AttachedToVm       apijson.Field
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	IPAddress          apijson.Field
	LastSeen           apijson.Field
	NetworkID          apijson.Field
	PortID             apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SubnetID           apijson.Field
	Type               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerTypeExternalIP CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerType = "external_ip"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerTypeExternalIP:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceExternalIPWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SizeUnit CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnit `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerType     `json:"type,required"`
	// UUID of the file share
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                       `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FileShareType      apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

// Unit of size
type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnitGiB CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnit = "GiB"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerSizeUnitGiB:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerTypeFileShare CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerType = "file_share"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerTypeFileShare:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFileShareWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerType `json:"type,required"`
	// UUID of the floating IP
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                        `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	IPAddress          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerTypeFloatingip CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerType = "floatingip"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerTypeFloatingip:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFloatingIPWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                               `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerType `json:"type,required"`
	// UUID of the function
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                              `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnitGBs CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnit = "GBS"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerTypeFunctions CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerType = "functions"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerTypeFunctions:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionsWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                   `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerType `json:"type,required"`
	// UUID of the function call
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                  `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnitMls CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnit = "MLS"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerBillingValueUnitMls:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerTypeFunctionsCalls CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerType = "functions_calls"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerTypeFunctionsCalls:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionCallsWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                           `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerType `json:"type,required"`
	// UUID of the function egress traffic
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                          `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnitGB CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnit = "GB"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerBillingValueUnitGB:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerTypeFunctionsTraffic CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerType = "functions_traffic"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerTypeFunctionsTraffic:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceFunctionEgressTrafficWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SizeUnit CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnit `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerType     `json:"type,required"`
	// UUID of the image
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                    `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

// Unit of size
type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnitBytes CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnit = "bytes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerSizeUnitBytes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerTypeImage CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerType = "image"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerTypeImage:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceImagesWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
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
	RegionID int64                                                                               `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerType `json:"type,required"`
	// UUID of the inference deployment
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                              `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerTypeInference CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerType = "inference"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerTypeInference:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceInferenceWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                              `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerType `json:"type,required"`
	// UUID of the instance
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                      `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerTypeInstance CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerType = "instance"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerTypeInstance:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceInstanceWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                  `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerType `json:"type,required"`
	// UUID of the load balancer
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                          `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	Flavor             apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerTypeLoadBalancer CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerType = "load_balancer"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerTypeLoadBalancer:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceLoadBalancerWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SizeUnit string                                                                             `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerType `json:"type,required"`
	// UUID of the log index
	Uuid               string `json:"uuid,required,nullable" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                             `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerTypeLogIndex CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerType = "log_index"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerTypeLogIndex:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceLogIndexWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SourceVolumeUuid string                                                                             `json:"source_volume_uuid,required" format:"uuid"`
	Type             CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerType `json:"type,required"`
	// UUID of the snapshot
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType         string `json:"volume_type,required"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                      `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	SourceVolumeUuid   apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	VolumeType         apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerTypeSnapshot CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerType = "snapshot"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerTypeSnapshot:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceSnapshotWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SizeUnit string                                                                           `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerType `json:"type,required"`
	// UUID of the volume
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// ID of the VM the volume is attached to
	AttachedToVm       string `json:"attached_to_vm,nullable" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                                    `json:"tags,nullable"`
	JSON cloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	LastSize           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	VolumeType         apijson.Field
	AttachedToVm       apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerTypeVolume CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerType = "volume"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerTypeVolume:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceVolumeWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                           `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                          `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerTypeDbaasPostgresqlConnectionPooler CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPoolerWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                           `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                          `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerTypeDbaasPostgresqlMemory CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerType = "dbaas_postgresql_memory"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerTypeDbaasPostgresqlMemory:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlMemoryWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                                  `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                                 `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerTypeDbaasPostgresqlPublicNetwork CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerType = "dbaas_postgresql_public_network"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerTypeDbaasPostgresqlPublicNetwork:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlPublicNetworkWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	RegionID int64                                                                                        `json:"region_id,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid               string `json:"uuid,required" format:"uuid"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                       `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnitMinutes CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnit = "minutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerTypeDbaasPostgresqlCPU CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerType = "dbaas_postgresql_cpu"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerTypeDbaasPostgresqlCPU:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlcpuWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnit `json:"billing_value_unit,required"`
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
	SizeUnit string                                                                                          `json:"size_unit,required"`
	Type     CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType         string `json:"volume_type,required"`
	BillingFeatureName string `json:"billing_feature_name,nullable"`
	// Cost for requested period
	Cost CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerCostUnion `json:"cost,nullable"`
	// Currency of the cost
	Currency string `json:"currency,nullable"`
	// Error message
	Err string `json:"err,nullable"`
	// Region name
	RegionName string                                                                                          `json:"region_name,nullable"`
	JSON       cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerJSON `json:"-"`
}

// cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer]
type cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerJSON struct {
	BillingMetricName  apijson.Field
	BillingValue       apijson.Field
	BillingValueUnit   apijson.Field
	FirstSeen          apijson.Field
	LastName           apijson.Field
	LastSeen           apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	SizeUnit           apijson.Field
	Type               apijson.Field
	Uuid               apijson.Field
	VolumeType         apijson.Field
	BillingFeatureName apijson.Field
	Cost               apijson.Field
	Currency           apijson.Field
	Err                apijson.Field
	RegionName         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializer) implementsCloudV1CostReportGetResourcesResponseResult() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit of billing value
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnit string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnitGbminutes CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerType string

const (
	CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerTypeDbaasPostgresqlVolume CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerType = "dbaas_postgresql_volume"
)

func (r CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

// Cost for requested period
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerCostUnion interface {
	ImplementsCloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerCostUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetResourcesResponseResultsResourceDbaasPostgreSqlVolumeWithCostSerializerCostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1CostReportGetResourcesResponseResultsType string

const (
	CloudV1CostReportGetResourcesResponseResultsTypeAICluster                       CloudV1CostReportGetResourcesResponseResultsType = "ai_cluster"
	CloudV1CostReportGetResourcesResponseResultsTypeBaremetal                       CloudV1CostReportGetResourcesResponseResultsType = "baremetal"
	CloudV1CostReportGetResourcesResponseResultsTypeBasicVm                         CloudV1CostReportGetResourcesResponseResultsType = "basic_vm"
	CloudV1CostReportGetResourcesResponseResultsTypeBackup                          CloudV1CostReportGetResourcesResponseResultsType = "backup"
	CloudV1CostReportGetResourcesResponseResultsTypeContainers                      CloudV1CostReportGetResourcesResponseResultsType = "containers"
	CloudV1CostReportGetResourcesResponseResultsTypeEgressTraffic                   CloudV1CostReportGetResourcesResponseResultsType = "egress_traffic"
	CloudV1CostReportGetResourcesResponseResultsTypeExternalIP                      CloudV1CostReportGetResourcesResponseResultsType = "external_ip"
	CloudV1CostReportGetResourcesResponseResultsTypeFileShare                       CloudV1CostReportGetResourcesResponseResultsType = "file_share"
	CloudV1CostReportGetResourcesResponseResultsTypeFloatingip                      CloudV1CostReportGetResourcesResponseResultsType = "floatingip"
	CloudV1CostReportGetResourcesResponseResultsTypeFunctions                       CloudV1CostReportGetResourcesResponseResultsType = "functions"
	CloudV1CostReportGetResourcesResponseResultsTypeFunctionsCalls                  CloudV1CostReportGetResourcesResponseResultsType = "functions_calls"
	CloudV1CostReportGetResourcesResponseResultsTypeFunctionsTraffic                CloudV1CostReportGetResourcesResponseResultsType = "functions_traffic"
	CloudV1CostReportGetResourcesResponseResultsTypeImage                           CloudV1CostReportGetResourcesResponseResultsType = "image"
	CloudV1CostReportGetResourcesResponseResultsTypeInference                       CloudV1CostReportGetResourcesResponseResultsType = "inference"
	CloudV1CostReportGetResourcesResponseResultsTypeInstance                        CloudV1CostReportGetResourcesResponseResultsType = "instance"
	CloudV1CostReportGetResourcesResponseResultsTypeLoadBalancer                    CloudV1CostReportGetResourcesResponseResultsType = "load_balancer"
	CloudV1CostReportGetResourcesResponseResultsTypeLogIndex                        CloudV1CostReportGetResourcesResponseResultsType = "log_index"
	CloudV1CostReportGetResourcesResponseResultsTypeSnapshot                        CloudV1CostReportGetResourcesResponseResultsType = "snapshot"
	CloudV1CostReportGetResourcesResponseResultsTypeVolume                          CloudV1CostReportGetResourcesResponseResultsType = "volume"
	CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlConnectionPooler CloudV1CostReportGetResourcesResponseResultsType = "dbaas_postgresql_connection_pooler"
	CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlMemory           CloudV1CostReportGetResourcesResponseResultsType = "dbaas_postgresql_memory"
	CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlPublicNetwork    CloudV1CostReportGetResourcesResponseResultsType = "dbaas_postgresql_public_network"
	CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlCPU              CloudV1CostReportGetResourcesResponseResultsType = "dbaas_postgresql_cpu"
	CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlVolume           CloudV1CostReportGetResourcesResponseResultsType = "dbaas_postgresql_volume"
)

func (r CloudV1CostReportGetResourcesResponseResultsType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesResponseResultsTypeAICluster, CloudV1CostReportGetResourcesResponseResultsTypeBaremetal, CloudV1CostReportGetResourcesResponseResultsTypeBasicVm, CloudV1CostReportGetResourcesResponseResultsTypeBackup, CloudV1CostReportGetResourcesResponseResultsTypeContainers, CloudV1CostReportGetResourcesResponseResultsTypeEgressTraffic, CloudV1CostReportGetResourcesResponseResultsTypeExternalIP, CloudV1CostReportGetResourcesResponseResultsTypeFileShare, CloudV1CostReportGetResourcesResponseResultsTypeFloatingip, CloudV1CostReportGetResourcesResponseResultsTypeFunctions, CloudV1CostReportGetResourcesResponseResultsTypeFunctionsCalls, CloudV1CostReportGetResourcesResponseResultsTypeFunctionsTraffic, CloudV1CostReportGetResourcesResponseResultsTypeImage, CloudV1CostReportGetResourcesResponseResultsTypeInference, CloudV1CostReportGetResourcesResponseResultsTypeInstance, CloudV1CostReportGetResourcesResponseResultsTypeLoadBalancer, CloudV1CostReportGetResourcesResponseResultsTypeLogIndex, CloudV1CostReportGetResourcesResponseResultsTypeSnapshot, CloudV1CostReportGetResourcesResponseResultsTypeVolume, CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlConnectionPooler, CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlMemory, CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlPublicNetwork, CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlCPU, CloudV1CostReportGetResourcesResponseResultsTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1CostReportGetTotalsResponse struct {
	// Count of returned totals
	Count int64 `json:"count,required"`
	// Price status for the UI, type: string
	PriceStatus PriceDisplayStatus                         `json:"price_status,required"`
	Results     []CloudV1CostReportGetTotalsResponseResult `json:"results,required"`
	JSON        cloudV1CostReportGetTotalsResponseJSON     `json:"-"`
}

// cloudV1CostReportGetTotalsResponseJSON contains the JSON metadata for the struct
// [CloudV1CostReportGetTotalsResponse]
type cloudV1CostReportGetTotalsResponseJSON struct {
	Count       apijson.Field
	PriceStatus apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CostReportGetTotalsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CostReportGetTotalsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CostReportGetTotalsResponseResult struct {
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
	RegionID           int64                                         `json:"region_id,required"`
	Type               CloudV1CostReportGetTotalsResponseResultsType `json:"type,required"`
	BillingFeatureName string                                        `json:"billing_feature_name,nullable"`
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
	VolumeType string                                       `json:"volume_type"`
	JSON       cloudV1CostReportGetTotalsResponseResultJSON `json:"-"`
	union      CloudV1CostReportGetTotalsResponseResultsUnion
}

// cloudV1CostReportGetTotalsResponseResultJSON contains the JSON metadata for the
// struct [CloudV1CostReportGetTotalsResponseResult]
type cloudV1CostReportGetTotalsResponseResultJSON struct {
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

func (r cloudV1CostReportGetTotalsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1CostReportGetTotalsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1CostReportGetTotalsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1CostReportGetTotalsResponseResultsUnion] interface
// which you can cast to the specific types for more type safety.
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
func (r CloudV1CostReportGetTotalsResponseResult) AsUnion() CloudV1CostReportGetTotalsResponseResultsUnion {
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
type CloudV1CostReportGetTotalsResponseResultsUnion interface {
	implementsCloudV1CostReportGetTotalsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1CostReportGetTotalsResponseResultsUnion)(nil)).Elem(),
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

type CloudV1CostReportGetTotalsResponseResultsType string

const (
	CloudV1CostReportGetTotalsResponseResultsTypeAICluster                       CloudV1CostReportGetTotalsResponseResultsType = "ai_cluster"
	CloudV1CostReportGetTotalsResponseResultsTypeBaremetal                       CloudV1CostReportGetTotalsResponseResultsType = "baremetal"
	CloudV1CostReportGetTotalsResponseResultsTypeBasicVm                         CloudV1CostReportGetTotalsResponseResultsType = "basic_vm"
	CloudV1CostReportGetTotalsResponseResultsTypeBackup                          CloudV1CostReportGetTotalsResponseResultsType = "backup"
	CloudV1CostReportGetTotalsResponseResultsTypeContainers                      CloudV1CostReportGetTotalsResponseResultsType = "containers"
	CloudV1CostReportGetTotalsResponseResultsTypeEgressTraffic                   CloudV1CostReportGetTotalsResponseResultsType = "egress_traffic"
	CloudV1CostReportGetTotalsResponseResultsTypeExternalIP                      CloudV1CostReportGetTotalsResponseResultsType = "external_ip"
	CloudV1CostReportGetTotalsResponseResultsTypeFileShare                       CloudV1CostReportGetTotalsResponseResultsType = "file_share"
	CloudV1CostReportGetTotalsResponseResultsTypeFloatingip                      CloudV1CostReportGetTotalsResponseResultsType = "floatingip"
	CloudV1CostReportGetTotalsResponseResultsTypeFunctions                       CloudV1CostReportGetTotalsResponseResultsType = "functions"
	CloudV1CostReportGetTotalsResponseResultsTypeFunctionsCalls                  CloudV1CostReportGetTotalsResponseResultsType = "functions_calls"
	CloudV1CostReportGetTotalsResponseResultsTypeFunctionsTraffic                CloudV1CostReportGetTotalsResponseResultsType = "functions_traffic"
	CloudV1CostReportGetTotalsResponseResultsTypeImage                           CloudV1CostReportGetTotalsResponseResultsType = "image"
	CloudV1CostReportGetTotalsResponseResultsTypeInference                       CloudV1CostReportGetTotalsResponseResultsType = "inference"
	CloudV1CostReportGetTotalsResponseResultsTypeInstance                        CloudV1CostReportGetTotalsResponseResultsType = "instance"
	CloudV1CostReportGetTotalsResponseResultsTypeLoadBalancer                    CloudV1CostReportGetTotalsResponseResultsType = "load_balancer"
	CloudV1CostReportGetTotalsResponseResultsTypeLogIndex                        CloudV1CostReportGetTotalsResponseResultsType = "log_index"
	CloudV1CostReportGetTotalsResponseResultsTypeSnapshot                        CloudV1CostReportGetTotalsResponseResultsType = "snapshot"
	CloudV1CostReportGetTotalsResponseResultsTypeVolume                          CloudV1CostReportGetTotalsResponseResultsType = "volume"
	CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlConnectionPooler CloudV1CostReportGetTotalsResponseResultsType = "dbaas_postgresql_connection_pooler"
	CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlMemory           CloudV1CostReportGetTotalsResponseResultsType = "dbaas_postgresql_memory"
	CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlPublicNetwork    CloudV1CostReportGetTotalsResponseResultsType = "dbaas_postgresql_public_network"
	CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlCPU              CloudV1CostReportGetTotalsResponseResultsType = "dbaas_postgresql_cpu"
	CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlVolume           CloudV1CostReportGetTotalsResponseResultsType = "dbaas_postgresql_volume"
)

func (r CloudV1CostReportGetTotalsResponseResultsType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetTotalsResponseResultsTypeAICluster, CloudV1CostReportGetTotalsResponseResultsTypeBaremetal, CloudV1CostReportGetTotalsResponseResultsTypeBasicVm, CloudV1CostReportGetTotalsResponseResultsTypeBackup, CloudV1CostReportGetTotalsResponseResultsTypeContainers, CloudV1CostReportGetTotalsResponseResultsTypeEgressTraffic, CloudV1CostReportGetTotalsResponseResultsTypeExternalIP, CloudV1CostReportGetTotalsResponseResultsTypeFileShare, CloudV1CostReportGetTotalsResponseResultsTypeFloatingip, CloudV1CostReportGetTotalsResponseResultsTypeFunctions, CloudV1CostReportGetTotalsResponseResultsTypeFunctionsCalls, CloudV1CostReportGetTotalsResponseResultsTypeFunctionsTraffic, CloudV1CostReportGetTotalsResponseResultsTypeImage, CloudV1CostReportGetTotalsResponseResultsTypeInference, CloudV1CostReportGetTotalsResponseResultsTypeInstance, CloudV1CostReportGetTotalsResponseResultsTypeLoadBalancer, CloudV1CostReportGetTotalsResponseResultsTypeLogIndex, CloudV1CostReportGetTotalsResponseResultsTypeSnapshot, CloudV1CostReportGetTotalsResponseResultsTypeVolume, CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlConnectionPooler, CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlMemory, CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlPublicNetwork, CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlCPU, CloudV1CostReportGetTotalsResponseResultsTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day.
	TimeFrom param.Field[time.Time] `json:"time_from,required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day.
	TimeTo param.Field[time.Time] `json:"time_to,required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Field[bool] `json:"enable_last_day"`
	// The response resources limit. Defaults to 10.
	Limit param.Field[int64] `json:"limit"`
	// The response resources offset.
	Offset param.Field[int64] `json:"offset"`
	// List of project IDs
	Projects param.Field[[]int64] `json:"projects"`
	// List of region IDs.
	Regions param.Field[[]int64] `json:"regions"`
	// Format of the response (csv or json).
	ResponseFormat param.Field[CloudV1CostReportGetResourcesParamsResponseFormat] `json:"response_format"`
	// Extended filter for field filtering.
	SchemaFilter param.Field[CloudV1CostReportGetResourcesParamsSchemaFilterUnion] `json:"schema_filter"`
	// List of sorting filters (JSON objects) fields: project. directions: asc, desc.
	Sorting param.Field[[]CloudV1CostReportGetResourcesParamsSorting] `json:"sorting"`
	// Filter by tags
	Tags param.Field[TagsFilterParam] `json:"tags"`
	// List of resource types to be filtered in the report.
	Types param.Field[[]PrebillingResourceTypes] `json:"types"`
}

func (r CloudV1CostReportGetResourcesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Format of the response (csv or json).
type CloudV1CostReportGetResourcesParamsResponseFormat string

const (
	CloudV1CostReportGetResourcesParamsResponseFormatCsvRecords CloudV1CostReportGetResourcesParamsResponseFormat = "csv_records"
	CloudV1CostReportGetResourcesParamsResponseFormatJson       CloudV1CostReportGetResourcesParamsResponseFormat = "json"
)

func (r CloudV1CostReportGetResourcesParamsResponseFormat) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesParamsResponseFormatCsvRecords, CloudV1CostReportGetResourcesParamsResponseFormatJson:
		return true
	}
	return false
}

// Extended filter for field filtering.
type CloudV1CostReportGetResourcesParamsSchemaFilter struct {
	// Field name to filter by
	Field  param.Field[CloudV1CostReportGetResourcesParamsSchemaFilterField] `json:"field,required"`
	Type   param.Field[CloudV1CostReportGetResourcesParamsSchemaFilterType]  `json:"type,required"`
	Values param.Field[interface{}]                                          `json:"values,required"`
}

func (r CloudV1CostReportGetResourcesParamsSchemaFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1CostReportGetResourcesParamsSchemaFilter) implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion() {
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
// [CloudV1CostReportGetResourcesParamsSchemaFilter].
type CloudV1CostReportGetResourcesParamsSchemaFilterUnion interface {
	implementsCloudV1CostReportGetResourcesParamsSchemaFilterUnion()
}

// Field name to filter by
type CloudV1CostReportGetResourcesParamsSchemaFilterField string

const (
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldLastName         CloudV1CostReportGetResourcesParamsSchemaFilterField = "last_name"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldLastSize         CloudV1CostReportGetResourcesParamsSchemaFilterField = "last_size"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldSourceVolumeUuid CloudV1CostReportGetResourcesParamsSchemaFilterField = "source_volume_uuid"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldType             CloudV1CostReportGetResourcesParamsSchemaFilterField = "type"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldUuid             CloudV1CostReportGetResourcesParamsSchemaFilterField = "uuid"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldVolumeType       CloudV1CostReportGetResourcesParamsSchemaFilterField = "volume_type"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldFlavor           CloudV1CostReportGetResourcesParamsSchemaFilterField = "flavor"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldAttachedToVm     CloudV1CostReportGetResourcesParamsSchemaFilterField = "attached_to_vm"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldFileShareType    CloudV1CostReportGetResourcesParamsSchemaFilterField = "file_share_type"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldIPAddress        CloudV1CostReportGetResourcesParamsSchemaFilterField = "ip_address"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldInstanceName     CloudV1CostReportGetResourcesParamsSchemaFilterField = "instance_name"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldInstanceType     CloudV1CostReportGetResourcesParamsSchemaFilterField = "instance_type"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldPortID           CloudV1CostReportGetResourcesParamsSchemaFilterField = "port_id"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldVmID             CloudV1CostReportGetResourcesParamsSchemaFilterField = "vm_id"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldNetworkID        CloudV1CostReportGetResourcesParamsSchemaFilterField = "network_id"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldSubnetID         CloudV1CostReportGetResourcesParamsSchemaFilterField = "subnet_id"
	CloudV1CostReportGetResourcesParamsSchemaFilterFieldScheduleID       CloudV1CostReportGetResourcesParamsSchemaFilterField = "schedule_id"
)

func (r CloudV1CostReportGetResourcesParamsSchemaFilterField) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesParamsSchemaFilterFieldLastName, CloudV1CostReportGetResourcesParamsSchemaFilterFieldLastSize, CloudV1CostReportGetResourcesParamsSchemaFilterFieldSourceVolumeUuid, CloudV1CostReportGetResourcesParamsSchemaFilterFieldType, CloudV1CostReportGetResourcesParamsSchemaFilterFieldUuid, CloudV1CostReportGetResourcesParamsSchemaFilterFieldVolumeType, CloudV1CostReportGetResourcesParamsSchemaFilterFieldFlavor, CloudV1CostReportGetResourcesParamsSchemaFilterFieldAttachedToVm, CloudV1CostReportGetResourcesParamsSchemaFilterFieldFileShareType, CloudV1CostReportGetResourcesParamsSchemaFilterFieldIPAddress, CloudV1CostReportGetResourcesParamsSchemaFilterFieldInstanceName, CloudV1CostReportGetResourcesParamsSchemaFilterFieldInstanceType, CloudV1CostReportGetResourcesParamsSchemaFilterFieldPortID, CloudV1CostReportGetResourcesParamsSchemaFilterFieldVmID, CloudV1CostReportGetResourcesParamsSchemaFilterFieldNetworkID, CloudV1CostReportGetResourcesParamsSchemaFilterFieldSubnetID, CloudV1CostReportGetResourcesParamsSchemaFilterFieldScheduleID:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesParamsSchemaFilterType string

const (
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeSnapshot                        CloudV1CostReportGetResourcesParamsSchemaFilterType = "snapshot"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeInstance                        CloudV1CostReportGetResourcesParamsSchemaFilterType = "instance"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeAICluster                       CloudV1CostReportGetResourcesParamsSchemaFilterType = "ai_cluster"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeBasicVm                         CloudV1CostReportGetResourcesParamsSchemaFilterType = "basic_vm"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeBaremetal                       CloudV1CostReportGetResourcesParamsSchemaFilterType = "baremetal"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeVolume                          CloudV1CostReportGetResourcesParamsSchemaFilterType = "volume"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeFileShare                       CloudV1CostReportGetResourcesParamsSchemaFilterType = "file_share"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeImage                           CloudV1CostReportGetResourcesParamsSchemaFilterType = "image"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeFloatingip                      CloudV1CostReportGetResourcesParamsSchemaFilterType = "floatingip"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeEgressTraffic                   CloudV1CostReportGetResourcesParamsSchemaFilterType = "egress_traffic"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeLoadBalancer                    CloudV1CostReportGetResourcesParamsSchemaFilterType = "load_balancer"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeExternalIP                      CloudV1CostReportGetResourcesParamsSchemaFilterType = "external_ip"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeBackup                          CloudV1CostReportGetResourcesParamsSchemaFilterType = "backup"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeLogIndex                        CloudV1CostReportGetResourcesParamsSchemaFilterType = "log_index"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctions                       CloudV1CostReportGetResourcesParamsSchemaFilterType = "functions"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctionsCalls                  CloudV1CostReportGetResourcesParamsSchemaFilterType = "functions_calls"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctionsTraffic                CloudV1CostReportGetResourcesParamsSchemaFilterType = "functions_traffic"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeContainers                      CloudV1CostReportGetResourcesParamsSchemaFilterType = "containers"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeInference                       CloudV1CostReportGetResourcesParamsSchemaFilterType = "inference"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlVolume           CloudV1CostReportGetResourcesParamsSchemaFilterType = "dbaas_postgresql_volume"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork    CloudV1CostReportGetResourcesParamsSchemaFilterType = "dbaas_postgresql_public_network"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlCPU              CloudV1CostReportGetResourcesParamsSchemaFilterType = "dbaas_postgresql_cpu"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlMemory           CloudV1CostReportGetResourcesParamsSchemaFilterType = "dbaas_postgresql_memory"
	CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler CloudV1CostReportGetResourcesParamsSchemaFilterType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1CostReportGetResourcesParamsSchemaFilterType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetResourcesParamsSchemaFilterTypeSnapshot, CloudV1CostReportGetResourcesParamsSchemaFilterTypeInstance, CloudV1CostReportGetResourcesParamsSchemaFilterTypeAICluster, CloudV1CostReportGetResourcesParamsSchemaFilterTypeBasicVm, CloudV1CostReportGetResourcesParamsSchemaFilterTypeBaremetal, CloudV1CostReportGetResourcesParamsSchemaFilterTypeVolume, CloudV1CostReportGetResourcesParamsSchemaFilterTypeFileShare, CloudV1CostReportGetResourcesParamsSchemaFilterTypeImage, CloudV1CostReportGetResourcesParamsSchemaFilterTypeFloatingip, CloudV1CostReportGetResourcesParamsSchemaFilterTypeEgressTraffic, CloudV1CostReportGetResourcesParamsSchemaFilterTypeLoadBalancer, CloudV1CostReportGetResourcesParamsSchemaFilterTypeExternalIP, CloudV1CostReportGetResourcesParamsSchemaFilterTypeBackup, CloudV1CostReportGetResourcesParamsSchemaFilterTypeLogIndex, CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctions, CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctionsCalls, CloudV1CostReportGetResourcesParamsSchemaFilterTypeFunctionsTraffic, CloudV1CostReportGetResourcesParamsSchemaFilterTypeContainers, CloudV1CostReportGetResourcesParamsSchemaFilterTypeInference, CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlVolume, CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork, CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlCPU, CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlMemory, CloudV1CostReportGetResourcesParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

type CloudV1CostReportGetResourcesParamsSorting struct {
	BillingValue param.Field[SortingDirections] `json:"billing_value"`
	FirstSeen    param.Field[SortingDirections] `json:"first_seen"`
	LastName     param.Field[SortingDirections] `json:"last_name"`
	LastSeen     param.Field[SortingDirections] `json:"last_seen"`
	Project      param.Field[SortingDirections] `json:"project"`
	Region       param.Field[SortingDirections] `json:"region"`
	Type         param.Field[SortingDirections] `json:"type"`
}

func (r CloudV1CostReportGetResourcesParamsSorting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CostReportGetTotalsParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day.
	TimeFrom param.Field[time.Time] `json:"time_from,required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day.
	TimeTo param.Field[time.Time] `json:"time_to,required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Field[bool] `json:"enable_last_day"`
	// List of project IDs
	Projects param.Field[[]int64] `json:"projects"`
	// List of region IDs.
	Regions param.Field[[]int64] `json:"regions"`
	// Format of the response (csv or json).
	ResponseFormat param.Field[CloudV1CostReportGetTotalsParamsResponseFormat] `json:"response_format"`
	// Extended filter for field filtering.
	SchemaFilter param.Field[CloudV1CostReportGetTotalsParamsSchemaFilterUnion] `json:"schema_filter"`
	// Filter by tags
	Tags param.Field[TagsFilterParam] `json:"tags"`
	// List of resource types to be filtered in the report.
	Types param.Field[[]PrebillingResourceTypes] `json:"types"`
}

func (r CloudV1CostReportGetTotalsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Format of the response (csv or json).
type CloudV1CostReportGetTotalsParamsResponseFormat string

const (
	CloudV1CostReportGetTotalsParamsResponseFormatCsvTotals CloudV1CostReportGetTotalsParamsResponseFormat = "csv_totals"
	CloudV1CostReportGetTotalsParamsResponseFormatJson      CloudV1CostReportGetTotalsParamsResponseFormat = "json"
)

func (r CloudV1CostReportGetTotalsParamsResponseFormat) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetTotalsParamsResponseFormatCsvTotals, CloudV1CostReportGetTotalsParamsResponseFormatJson:
		return true
	}
	return false
}

// Extended filter for field filtering.
type CloudV1CostReportGetTotalsParamsSchemaFilter struct {
	// Field name to filter by
	Field  param.Field[CloudV1CostReportGetTotalsParamsSchemaFilterField] `json:"field,required"`
	Type   param.Field[CloudV1CostReportGetTotalsParamsSchemaFilterType]  `json:"type,required"`
	Values param.Field[interface{}]                                       `json:"values,required"`
}

func (r CloudV1CostReportGetTotalsParamsSchemaFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1CostReportGetTotalsParamsSchemaFilter) implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion() {
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
// [CloudV1CostReportGetTotalsParamsSchemaFilter].
type CloudV1CostReportGetTotalsParamsSchemaFilterUnion interface {
	implementsCloudV1CostReportGetTotalsParamsSchemaFilterUnion()
}

// Field name to filter by
type CloudV1CostReportGetTotalsParamsSchemaFilterField string

const (
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldLastName         CloudV1CostReportGetTotalsParamsSchemaFilterField = "last_name"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldLastSize         CloudV1CostReportGetTotalsParamsSchemaFilterField = "last_size"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldSourceVolumeUuid CloudV1CostReportGetTotalsParamsSchemaFilterField = "source_volume_uuid"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldType             CloudV1CostReportGetTotalsParamsSchemaFilterField = "type"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldUuid             CloudV1CostReportGetTotalsParamsSchemaFilterField = "uuid"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldVolumeType       CloudV1CostReportGetTotalsParamsSchemaFilterField = "volume_type"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldFlavor           CloudV1CostReportGetTotalsParamsSchemaFilterField = "flavor"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldAttachedToVm     CloudV1CostReportGetTotalsParamsSchemaFilterField = "attached_to_vm"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldFileShareType    CloudV1CostReportGetTotalsParamsSchemaFilterField = "file_share_type"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldIPAddress        CloudV1CostReportGetTotalsParamsSchemaFilterField = "ip_address"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldInstanceName     CloudV1CostReportGetTotalsParamsSchemaFilterField = "instance_name"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldInstanceType     CloudV1CostReportGetTotalsParamsSchemaFilterField = "instance_type"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldPortID           CloudV1CostReportGetTotalsParamsSchemaFilterField = "port_id"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldVmID             CloudV1CostReportGetTotalsParamsSchemaFilterField = "vm_id"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldNetworkID        CloudV1CostReportGetTotalsParamsSchemaFilterField = "network_id"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldSubnetID         CloudV1CostReportGetTotalsParamsSchemaFilterField = "subnet_id"
	CloudV1CostReportGetTotalsParamsSchemaFilterFieldScheduleID       CloudV1CostReportGetTotalsParamsSchemaFilterField = "schedule_id"
)

func (r CloudV1CostReportGetTotalsParamsSchemaFilterField) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetTotalsParamsSchemaFilterFieldLastName, CloudV1CostReportGetTotalsParamsSchemaFilterFieldLastSize, CloudV1CostReportGetTotalsParamsSchemaFilterFieldSourceVolumeUuid, CloudV1CostReportGetTotalsParamsSchemaFilterFieldType, CloudV1CostReportGetTotalsParamsSchemaFilterFieldUuid, CloudV1CostReportGetTotalsParamsSchemaFilterFieldVolumeType, CloudV1CostReportGetTotalsParamsSchemaFilterFieldFlavor, CloudV1CostReportGetTotalsParamsSchemaFilterFieldAttachedToVm, CloudV1CostReportGetTotalsParamsSchemaFilterFieldFileShareType, CloudV1CostReportGetTotalsParamsSchemaFilterFieldIPAddress, CloudV1CostReportGetTotalsParamsSchemaFilterFieldInstanceName, CloudV1CostReportGetTotalsParamsSchemaFilterFieldInstanceType, CloudV1CostReportGetTotalsParamsSchemaFilterFieldPortID, CloudV1CostReportGetTotalsParamsSchemaFilterFieldVmID, CloudV1CostReportGetTotalsParamsSchemaFilterFieldNetworkID, CloudV1CostReportGetTotalsParamsSchemaFilterFieldSubnetID, CloudV1CostReportGetTotalsParamsSchemaFilterFieldScheduleID:
		return true
	}
	return false
}

type CloudV1CostReportGetTotalsParamsSchemaFilterType string

const (
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeSnapshot                        CloudV1CostReportGetTotalsParamsSchemaFilterType = "snapshot"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeInstance                        CloudV1CostReportGetTotalsParamsSchemaFilterType = "instance"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeAICluster                       CloudV1CostReportGetTotalsParamsSchemaFilterType = "ai_cluster"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeBasicVm                         CloudV1CostReportGetTotalsParamsSchemaFilterType = "basic_vm"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeBaremetal                       CloudV1CostReportGetTotalsParamsSchemaFilterType = "baremetal"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeVolume                          CloudV1CostReportGetTotalsParamsSchemaFilterType = "volume"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeFileShare                       CloudV1CostReportGetTotalsParamsSchemaFilterType = "file_share"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeImage                           CloudV1CostReportGetTotalsParamsSchemaFilterType = "image"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeFloatingip                      CloudV1CostReportGetTotalsParamsSchemaFilterType = "floatingip"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeEgressTraffic                   CloudV1CostReportGetTotalsParamsSchemaFilterType = "egress_traffic"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeLoadBalancer                    CloudV1CostReportGetTotalsParamsSchemaFilterType = "load_balancer"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeExternalIP                      CloudV1CostReportGetTotalsParamsSchemaFilterType = "external_ip"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeBackup                          CloudV1CostReportGetTotalsParamsSchemaFilterType = "backup"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeLogIndex                        CloudV1CostReportGetTotalsParamsSchemaFilterType = "log_index"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctions                       CloudV1CostReportGetTotalsParamsSchemaFilterType = "functions"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctionsCalls                  CloudV1CostReportGetTotalsParamsSchemaFilterType = "functions_calls"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctionsTraffic                CloudV1CostReportGetTotalsParamsSchemaFilterType = "functions_traffic"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeContainers                      CloudV1CostReportGetTotalsParamsSchemaFilterType = "containers"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeInference                       CloudV1CostReportGetTotalsParamsSchemaFilterType = "inference"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlVolume           CloudV1CostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_volume"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork    CloudV1CostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_public_network"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlCPU              CloudV1CostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_cpu"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlMemory           CloudV1CostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_memory"
	CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler CloudV1CostReportGetTotalsParamsSchemaFilterType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1CostReportGetTotalsParamsSchemaFilterType) IsKnown() bool {
	switch r {
	case CloudV1CostReportGetTotalsParamsSchemaFilterTypeSnapshot, CloudV1CostReportGetTotalsParamsSchemaFilterTypeInstance, CloudV1CostReportGetTotalsParamsSchemaFilterTypeAICluster, CloudV1CostReportGetTotalsParamsSchemaFilterTypeBasicVm, CloudV1CostReportGetTotalsParamsSchemaFilterTypeBaremetal, CloudV1CostReportGetTotalsParamsSchemaFilterTypeVolume, CloudV1CostReportGetTotalsParamsSchemaFilterTypeFileShare, CloudV1CostReportGetTotalsParamsSchemaFilterTypeImage, CloudV1CostReportGetTotalsParamsSchemaFilterTypeFloatingip, CloudV1CostReportGetTotalsParamsSchemaFilterTypeEgressTraffic, CloudV1CostReportGetTotalsParamsSchemaFilterTypeLoadBalancer, CloudV1CostReportGetTotalsParamsSchemaFilterTypeExternalIP, CloudV1CostReportGetTotalsParamsSchemaFilterTypeBackup, CloudV1CostReportGetTotalsParamsSchemaFilterTypeLogIndex, CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctions, CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctionsCalls, CloudV1CostReportGetTotalsParamsSchemaFilterTypeFunctionsTraffic, CloudV1CostReportGetTotalsParamsSchemaFilterTypeContainers, CloudV1CostReportGetTotalsParamsSchemaFilterTypeInference, CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlVolume, CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork, CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlCPU, CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlMemory, CloudV1CostReportGetTotalsParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}
