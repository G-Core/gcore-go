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

// UsageReportService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageReportService] method instead.
type UsageReportService struct {
	Options []option.RequestOption
}

// NewUsageReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageReportService(opts ...option.RequestOption) (r UsageReportService) {
	r = UsageReportService{}
	r.Options = opts
	return
}

// Data from the past hour may not reflect the full set of statistics. For the most
// complete and accurate results, we recommend accessing the data at least one hour
// after the relevant time period. Updates are generally available within a 24-hour
// window, though timing can vary. Scheduled maintenance or other exceptions may
// occasionally cause delays beyond 24 hours.
func (r *UsageReportService) Get(ctx context.Context, body UsageReportGetParams, opts ...option.RequestOption) (res *UsageReport, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/usage_report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UsageReport struct {
	// Total count of the resources
	Count     int64                      `json:"count" api:"required"`
	Resources []UsageReportResourceUnion `json:"resources" api:"required"`
	Totals    []UsageReportTotalUnion    `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Resources   respjson.Field
		Totals      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReport) RawJSON() string { return r.JSON.raw }
func (r *UsageReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UsageReportResourceUnion contains all possible properties and values from
// [UsageReportResourceAICluster], [UsageReportResourceAIVirtualCluster],
// [UsageReportResourceBaremetal], [UsageReportResourceBasicVm],
// [UsageReportResourceBackup], [UsageReportResourceContainers],
// [UsageReportResourceEgressTraffic], [UsageReportResourceExternalIP],
// [UsageReportResourceFileShare], [UsageReportResourceFloatingip],
// [UsageReportResourceFunctions], [UsageReportResourceFunctionsCalls],
// [UsageReportResourceFunctionsTraffic], [UsageReportResourceImage],
// [UsageReportResourceInference], [UsageReportResourceInstance],
// [UsageReportResourceLoadBalancer], [UsageReportResourceLogIndex],
// [UsageReportResourceSnapshot], [UsageReportResourceVolume],
// [UsageReportResourceDbaasPostgreSQLConnectionPooler],
// [UsageReportResourceDbaasPostgreSQLMemory],
// [UsageReportResourceDbaasPostgreSQLPublicNetwork],
// [UsageReportResourceDbaasPostgreSQLCPU],
// [UsageReportResourceDbaasPostgreSQLVolume].
//
// Use the [UsageReportResourceUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UsageReportResourceUnion struct {
	BillingMetricName string              `json:"billing_metric_name"`
	BillingValue      float64             `json:"billing_value"`
	BillingValueUnit  string              `json:"billing_value_unit"`
	FirstSeen         time.Time           `json:"first_seen"`
	Flavor            string              `json:"flavor"`
	LastName          string              `json:"last_name"`
	LastSeen          time.Time           `json:"last_seen"`
	ProjectID         int64               `json:"project_id"`
	Region            int64               `json:"region"`
	RegionID          int64               `json:"region_id"`
	Tags              []map[string]string `json:"tags"`
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
	// This field is from variant [UsageReportResourceBackup].
	ScheduleID       string `json:"schedule_id"`
	SourceVolumeUuid string `json:"source_volume_uuid"`
	// This field is from variant [UsageReportResourceEgressTraffic].
	InstanceName string `json:"instance_name"`
	// This field is from variant [UsageReportResourceEgressTraffic].
	InstanceType string `json:"instance_type"`
	PortID       string `json:"port_id"`
	SizeUnit     string `json:"size_unit"`
	// This field is from variant [UsageReportResourceEgressTraffic].
	VmID         string `json:"vm_id"`
	AttachedToVm string `json:"attached_to_vm"`
	IPAddress    string `json:"ip_address"`
	// This field is from variant [UsageReportResourceExternalIP].
	NetworkID string `json:"network_id"`
	// This field is from variant [UsageReportResourceExternalIP].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [UsageReportResourceFileShare].
	FileShareType string `json:"file_share_type"`
	VolumeType    string `json:"volume_type"`
	JSON          struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		LastSize          respjson.Field
		ScheduleID        respjson.Field
		SourceVolumeUuid  respjson.Field
		InstanceName      respjson.Field
		InstanceType      respjson.Field
		PortID            respjson.Field
		SizeUnit          respjson.Field
		VmID              respjson.Field
		AttachedToVm      respjson.Field
		IPAddress         respjson.Field
		NetworkID         respjson.Field
		SubnetID          respjson.Field
		FileShareType     respjson.Field
		VolumeType        respjson.Field
		raw               string
	} `json:"-"`
}

// anyUsageReportResource is implemented by each variant of
// [UsageReportResourceUnion] to add type safety for the return type of
// [UsageReportResourceUnion.AsAny]
type anyUsageReportResource interface {
	implUsageReportResourceUnion()
}

func (UsageReportResourceAICluster) implUsageReportResourceUnion()                       {}
func (UsageReportResourceAIVirtualCluster) implUsageReportResourceUnion()                {}
func (UsageReportResourceBaremetal) implUsageReportResourceUnion()                       {}
func (UsageReportResourceBasicVm) implUsageReportResourceUnion()                         {}
func (UsageReportResourceBackup) implUsageReportResourceUnion()                          {}
func (UsageReportResourceContainers) implUsageReportResourceUnion()                      {}
func (UsageReportResourceEgressTraffic) implUsageReportResourceUnion()                   {}
func (UsageReportResourceExternalIP) implUsageReportResourceUnion()                      {}
func (UsageReportResourceFileShare) implUsageReportResourceUnion()                       {}
func (UsageReportResourceFloatingip) implUsageReportResourceUnion()                      {}
func (UsageReportResourceFunctions) implUsageReportResourceUnion()                       {}
func (UsageReportResourceFunctionsCalls) implUsageReportResourceUnion()                  {}
func (UsageReportResourceFunctionsTraffic) implUsageReportResourceUnion()                {}
func (UsageReportResourceImage) implUsageReportResourceUnion()                           {}
func (UsageReportResourceInference) implUsageReportResourceUnion()                       {}
func (UsageReportResourceInstance) implUsageReportResourceUnion()                        {}
func (UsageReportResourceLoadBalancer) implUsageReportResourceUnion()                    {}
func (UsageReportResourceLogIndex) implUsageReportResourceUnion()                        {}
func (UsageReportResourceSnapshot) implUsageReportResourceUnion()                        {}
func (UsageReportResourceVolume) implUsageReportResourceUnion()                          {}
func (UsageReportResourceDbaasPostgreSQLConnectionPooler) implUsageReportResourceUnion() {}
func (UsageReportResourceDbaasPostgreSQLMemory) implUsageReportResourceUnion()           {}
func (UsageReportResourceDbaasPostgreSQLPublicNetwork) implUsageReportResourceUnion()    {}
func (UsageReportResourceDbaasPostgreSQLCPU) implUsageReportResourceUnion()              {}
func (UsageReportResourceDbaasPostgreSQLVolume) implUsageReportResourceUnion()           {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UsageReportResourceUnion.AsAny().(type) {
//	case cloud.UsageReportResourceAICluster:
//	case cloud.UsageReportResourceAIVirtualCluster:
//	case cloud.UsageReportResourceBaremetal:
//	case cloud.UsageReportResourceBasicVm:
//	case cloud.UsageReportResourceBackup:
//	case cloud.UsageReportResourceContainers:
//	case cloud.UsageReportResourceEgressTraffic:
//	case cloud.UsageReportResourceExternalIP:
//	case cloud.UsageReportResourceFileShare:
//	case cloud.UsageReportResourceFloatingip:
//	case cloud.UsageReportResourceFunctions:
//	case cloud.UsageReportResourceFunctionsCalls:
//	case cloud.UsageReportResourceFunctionsTraffic:
//	case cloud.UsageReportResourceImage:
//	case cloud.UsageReportResourceInference:
//	case cloud.UsageReportResourceInstance:
//	case cloud.UsageReportResourceLoadBalancer:
//	case cloud.UsageReportResourceLogIndex:
//	case cloud.UsageReportResourceSnapshot:
//	case cloud.UsageReportResourceVolume:
//	case cloud.UsageReportResourceDbaasPostgreSQLConnectionPooler:
//	case cloud.UsageReportResourceDbaasPostgreSQLMemory:
//	case cloud.UsageReportResourceDbaasPostgreSQLPublicNetwork:
//	case cloud.UsageReportResourceDbaasPostgreSQLCPU:
//	case cloud.UsageReportResourceDbaasPostgreSQLVolume:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UsageReportResourceUnion) AsAny() anyUsageReportResource {
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

func (u UsageReportResourceUnion) AsAICluster() (v UsageReportResourceAICluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsAIVirtualCluster() (v UsageReportResourceAIVirtualCluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsBaremetal() (v UsageReportResourceBaremetal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsBasicVm() (v UsageReportResourceBasicVm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsBackup() (v UsageReportResourceBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsContainers() (v UsageReportResourceContainers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsEgressTraffic() (v UsageReportResourceEgressTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsExternalIP() (v UsageReportResourceExternalIP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsFileShare() (v UsageReportResourceFileShare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsFloatingip() (v UsageReportResourceFloatingip) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsFunctions() (v UsageReportResourceFunctions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsFunctionsCalls() (v UsageReportResourceFunctionsCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsFunctionsTraffic() (v UsageReportResourceFunctionsTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsImage() (v UsageReportResourceImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsInference() (v UsageReportResourceInference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsInstance() (v UsageReportResourceInstance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsLoadBalancer() (v UsageReportResourceLoadBalancer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsLogIndex() (v UsageReportResourceLogIndex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsSnapshot() (v UsageReportResourceSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsVolume() (v UsageReportResourceVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsDbaasPostgreSQLConnectionPooler() (v UsageReportResourceDbaasPostgreSQLConnectionPooler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsDbaasPostgreSQLMemory() (v UsageReportResourceDbaasPostgreSQLMemory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsDbaasPostgreSQLPublicNetwork() (v UsageReportResourceDbaasPostgreSQLPublicNetwork) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsDbaasPostgreSQLCPU() (v UsageReportResourceDbaasPostgreSQLCPU) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportResourceUnion) AsDbaasPostgreSQLVolume() (v UsageReportResourceDbaasPostgreSQLVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UsageReportResourceUnion) RawJSON() string { return u.JSON.raw }

func (r *UsageReportResourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual AI clusters in all cost and usage
// reports results (but not in totals)
type UsageReportResourceAICluster struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Name of the AI cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.AICluster  `json:"type" api:"required"`
	// UUID of the Baremetal GPU cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceAICluster) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual AI Virtual clusters in all cost
// and usage reports results (but not in totals)
type UsageReportResourceAIVirtualCluster struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Name of the AI cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string       `json:"tags" api:"required"`
	Type constant.AIVirtualCluster `json:"type" api:"required"`
	// UUID of the Virtual GPU cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceAIVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual bare metal servers in all cost
// and usage reports results (but not in totals)
type UsageReportResourceBaremetal struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor" api:"required"`
	// Name of the bare metal server
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Baremetal  `json:"type" api:"required"`
	// UUID of the bare metal server
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceBaremetal) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual basic VMs in all cost and usage
// reports results (but not in totals)
type UsageReportResourceBasicVm struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the basic VM
	Flavor string `json:"flavor" api:"required"`
	// Name of the basic VM
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.BasicVm    `json:"type" api:"required"`
	// UUID of the basic VM
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceBasicVm) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual backups in all cost and usage
// reports results (but not in totals)
type UsageReportResourceBackup struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the backup
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// ID of the backup schedule
	ScheduleID string `json:"schedule_id" api:"required" format:"uuid"`
	// UUID of the source volume
	SourceVolumeUuid string          `json:"source_volume_uuid" api:"required" format:"uuid"`
	Type             constant.Backup `json:"type" api:"required"`
	// UUID of the backup
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		ScheduleID        respjson.Field
		SourceVolumeUuid  respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceBackup) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceContainers struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the container
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.Containers `json:"type" api:"required"`
	// UUID of the container
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceContainers) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual egress traffic in all cost and
// usage reports results (but not in totals)
type UsageReportResourceEgressTraffic struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the instance
	InstanceName string `json:"instance_name" api:"required"`
	// Type of the instance
	//
	// Any of "baremetal", "vm".
	InstanceType string `json:"instance_type" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the port the traffic is associated with
	PortID string `json:"port_id" api:"required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit string `json:"size_unit" api:"required"`
	// List of tags
	Tags []map[string]string    `json:"tags" api:"required"`
	Type constant.EgressTraffic `json:"type" api:"required"`
	// ID of the bare metal server the traffic is associated with
	VmID string `json:"vm_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		InstanceName      respjson.Field
		InstanceType      respjson.Field
		LastSeen          respjson.Field
		PortID            respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		VmID              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual external IPs in all cost and
// usage reports results (but not in totals)
type UsageReportResourceExternalIP struct {
	// ID of the VM the IP is attached to
	AttachedToVm string `json:"attached_to_vm" api:"required" format:"uuid"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the network the IP is attached to
	NetworkID string `json:"network_id" api:"required" format:"uuid"`
	// ID of the port the IP is associated with
	PortID string `json:"port_id" api:"required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// ID of the subnet the IP is attached to
	SubnetID string              `json:"subnet_id" api:"required" format:"uuid"`
	Type     constant.ExternalIP `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedToVm      respjson.Field
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		IPAddress         respjson.Field
		LastSeen          respjson.Field
		NetworkID         respjson.Field
		PortID            respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SubnetID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceExternalIP) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual file shares in all cost and usage
// reports results (but not in totals)
type UsageReportResourceFileShare struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Type of the file share
	FileShareType string `json:"file_share_type" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the file share
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the file share in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit constant.GiB `json:"size_unit" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.FileShare  `json:"type" api:"required"`
	// UUID of the file share
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FileShareType     respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceFileShare) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual floating IPs in all cost and
// usage reports results (but not in totals)
type UsageReportResourceFloatingip struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Name of the floating IP
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Floatingip `json:"type" api:"required"`
	// UUID of the floating IP
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		IPAddress         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceFloatingip) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual functions in all cost and usage
// reports results (but not in totals)
type UsageReportResourceFunctions struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the function
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Functions `json:"type" api:"required"`
	// UUID of the function
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceFunctions) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual function calls in all cost and
// usage reports results (but not in totals)
type UsageReportResourceFunctionsCalls struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the function call
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                   `json:"region_id" api:"required"`
	Type     constant.FunctionsCalls `json:"type" api:"required"`
	// UUID of the function call
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceFunctionsCalls) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual function egress traffic in all
// cost and usage reports results (but not in totals)
type UsageReportResourceFunctionsTraffic struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the function egress traffic
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                     `json:"region_id" api:"required"`
	Type     constant.FunctionsTraffic `json:"type" api:"required"`
	// UUID of the function egress traffic
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceFunctionsTraffic) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual images in all cost and usage
// reports results (but not in totals)
type UsageReportResourceImage struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the image
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the image in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit constant.Bytes `json:"size_unit" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Image      `json:"type" api:"required"`
	// UUID of the image
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceImage) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceInference struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the inference deployment
	Flavor string `json:"flavor" api:"required"`
	// Name of the inference deployment
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Inference `json:"type" api:"required"`
	// UUID of the inference deployment
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceInference) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual instances in all cost and usage
// reports results (but not in totals)
type UsageReportResourceInstance struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the instance
	Flavor string `json:"flavor" api:"required"`
	// Name of the instance
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Instance   `json:"type" api:"required"`
	// UUID of the instance
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceInstance) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual load balancers in all cost and
// usage reports results (but not in totals)
type UsageReportResourceLoadBalancer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Flavor of the load balancer
	Flavor string `json:"flavor" api:"required"`
	// Name of the load balancer
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// List of tags
	Tags []map[string]string   `json:"tags" api:"required"`
	Type constant.LoadBalancer `json:"type" api:"required"`
	// UUID of the load balancer
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		Flavor            respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual log indexes in all cost and usage
// reports results (but not in totals)
type UsageReportResourceLogIndex struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the log index
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the log index in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit string            `json:"size_unit" api:"required"`
	Type     constant.LogIndex `json:"type" api:"required"`
	// UUID of the log index
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceLogIndex) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual snapshots in all cost and usage
// reports results (but not in totals)
type UsageReportResourceSnapshot struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the snapshot
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the snapshot in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit string `json:"size_unit" api:"required"`
	// UUID of the source volume
	SourceVolumeUuid string `json:"source_volume_uuid" api:"required" format:"uuid"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Snapshot   `json:"type" api:"required"`
	// UUID of the snapshot
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		SourceVolumeUuid  respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceSnapshot) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// These properties are common for all individual volumes in all cost and usage
// reports results (but not in totals)
type UsageReportResourceVolume struct {
	// ID of the VM the volume is attached to
	AttachedToVm string `json:"attached_to_vm" api:"required" format:"uuid"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the volume
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// Size of the volume in bytes
	LastSize int64 `json:"last_size" api:"required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit string `json:"size_unit" api:"required"`
	// List of tags
	Tags []map[string]string `json:"tags" api:"required"`
	Type constant.Volume     `json:"type" api:"required"`
	// UUID of the volume
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedToVm      respjson.Field
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		LastSize          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Tags              respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceVolume) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceDbaasPostgreSQLConnectionPooler struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                    `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type" api:"required"`
	// UUID of the cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceDbaasPostgreSQLConnectionPooler) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceDbaasPostgreSQLMemory struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                          `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type" api:"required"`
	// UUID of the cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceDbaasPostgreSQLMemory) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceDbaasPostgreSQLPublicNetwork struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                 `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type" api:"required"`
	// UUID of the cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceDbaasPostgreSQLPublicNetwork) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceDbaasPostgreSQLCPU struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                       `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type" api:"required"`
	// UUID of the cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceDbaasPostgreSQLCPU) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportResourceDbaasPostgreSQLVolume struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen" api:"required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name" api:"required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen" api:"required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Unit of size
	SizeUnit string                         `json:"size_unit" api:"required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type" api:"required"`
	// UUID of the cluster
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FirstSeen         respjson.Field
		LastName          respjson.Field
		LastSeen          respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		SizeUnit          respjson.Field
		Type              respjson.Field
		Uuid              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportResourceDbaasPostgreSQLVolume) RawJSON() string { return r.JSON.raw }
func (r *UsageReportResourceDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UsageReportTotalUnion contains all possible properties and values from
// [UsageReportTotalAICluster], [UsageReportTotalAIVirtualCluster],
// [UsageReportTotalBaremetal], [UsageReportTotalBasicVm],
// [UsageReportTotalContainers], [UsageReportTotalEgressTraffic],
// [UsageReportTotalExternalIP], [UsageReportTotalFileShare],
// [UsageReportTotalFloatingip], [UsageReportTotalFunctions],
// [UsageReportTotalFunctionsCalls], [UsageReportTotalFunctionsTraffic],
// [UsageReportTotalImage], [UsageReportTotalInference],
// [UsageReportTotalInstance], [UsageReportTotalLoadBalancer],
// [UsageReportTotalLogIndex], [UsageReportTotalSnapshot],
// [UsageReportTotalVolume], [UsageReportTotalDbaasPostgreSQLConnectionPooler],
// [UsageReportTotalDbaasPostgreSQLMemory],
// [UsageReportTotalDbaasPostgreSQLPublicNetwork],
// [UsageReportTotalDbaasPostgreSQLCPU], [UsageReportTotalDbaasPostgreSQLVolume].
//
// Use the [UsageReportTotalUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UsageReportTotalUnion struct {
	BillingMetricName string  `json:"billing_metric_name"`
	BillingValue      float64 `json:"billing_value"`
	BillingValueUnit  string  `json:"billing_value_unit"`
	Flavor            string  `json:"flavor"`
	Region            int64   `json:"region"`
	RegionID          int64   `json:"region_id"`
	// Any of "ai_cluster", "ai_virtual_cluster", "baremetal", "basic_vm",
	// "containers", "egress_traffic", "external_ip", "file_share", "floatingip",
	// "functions", "functions_calls", "functions_traffic", "image", "inference",
	// "instance", "load_balancer", "log_index", "snapshot", "volume",
	// "dbaas_postgresql_connection_pooler", "dbaas_postgresql_memory",
	// "dbaas_postgresql_public_network", "dbaas_postgresql_cpu",
	// "dbaas_postgresql_volume".
	Type string `json:"type"`
	// This field is from variant [UsageReportTotalEgressTraffic].
	InstanceType string `json:"instance_type"`
	// This field is from variant [UsageReportTotalFileShare].
	FileShareType string `json:"file_share_type"`
	VolumeType    string `json:"volume_type"`
	JSON          struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		InstanceType      respjson.Field
		FileShareType     respjson.Field
		VolumeType        respjson.Field
		raw               string
	} `json:"-"`
}

// anyUsageReportTotal is implemented by each variant of [UsageReportTotalUnion] to
// add type safety for the return type of [UsageReportTotalUnion.AsAny]
type anyUsageReportTotal interface {
	implUsageReportTotalUnion()
}

func (UsageReportTotalAICluster) implUsageReportTotalUnion()                       {}
func (UsageReportTotalAIVirtualCluster) implUsageReportTotalUnion()                {}
func (UsageReportTotalBaremetal) implUsageReportTotalUnion()                       {}
func (UsageReportTotalBasicVm) implUsageReportTotalUnion()                         {}
func (UsageReportTotalContainers) implUsageReportTotalUnion()                      {}
func (UsageReportTotalEgressTraffic) implUsageReportTotalUnion()                   {}
func (UsageReportTotalExternalIP) implUsageReportTotalUnion()                      {}
func (UsageReportTotalFileShare) implUsageReportTotalUnion()                       {}
func (UsageReportTotalFloatingip) implUsageReportTotalUnion()                      {}
func (UsageReportTotalFunctions) implUsageReportTotalUnion()                       {}
func (UsageReportTotalFunctionsCalls) implUsageReportTotalUnion()                  {}
func (UsageReportTotalFunctionsTraffic) implUsageReportTotalUnion()                {}
func (UsageReportTotalImage) implUsageReportTotalUnion()                           {}
func (UsageReportTotalInference) implUsageReportTotalUnion()                       {}
func (UsageReportTotalInstance) implUsageReportTotalUnion()                        {}
func (UsageReportTotalLoadBalancer) implUsageReportTotalUnion()                    {}
func (UsageReportTotalLogIndex) implUsageReportTotalUnion()                        {}
func (UsageReportTotalSnapshot) implUsageReportTotalUnion()                        {}
func (UsageReportTotalVolume) implUsageReportTotalUnion()                          {}
func (UsageReportTotalDbaasPostgreSQLConnectionPooler) implUsageReportTotalUnion() {}
func (UsageReportTotalDbaasPostgreSQLMemory) implUsageReportTotalUnion()           {}
func (UsageReportTotalDbaasPostgreSQLPublicNetwork) implUsageReportTotalUnion()    {}
func (UsageReportTotalDbaasPostgreSQLCPU) implUsageReportTotalUnion()              {}
func (UsageReportTotalDbaasPostgreSQLVolume) implUsageReportTotalUnion()           {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UsageReportTotalUnion.AsAny().(type) {
//	case cloud.UsageReportTotalAICluster:
//	case cloud.UsageReportTotalAIVirtualCluster:
//	case cloud.UsageReportTotalBaremetal:
//	case cloud.UsageReportTotalBasicVm:
//	case cloud.UsageReportTotalContainers:
//	case cloud.UsageReportTotalEgressTraffic:
//	case cloud.UsageReportTotalExternalIP:
//	case cloud.UsageReportTotalFileShare:
//	case cloud.UsageReportTotalFloatingip:
//	case cloud.UsageReportTotalFunctions:
//	case cloud.UsageReportTotalFunctionsCalls:
//	case cloud.UsageReportTotalFunctionsTraffic:
//	case cloud.UsageReportTotalImage:
//	case cloud.UsageReportTotalInference:
//	case cloud.UsageReportTotalInstance:
//	case cloud.UsageReportTotalLoadBalancer:
//	case cloud.UsageReportTotalLogIndex:
//	case cloud.UsageReportTotalSnapshot:
//	case cloud.UsageReportTotalVolume:
//	case cloud.UsageReportTotalDbaasPostgreSQLConnectionPooler:
//	case cloud.UsageReportTotalDbaasPostgreSQLMemory:
//	case cloud.UsageReportTotalDbaasPostgreSQLPublicNetwork:
//	case cloud.UsageReportTotalDbaasPostgreSQLCPU:
//	case cloud.UsageReportTotalDbaasPostgreSQLVolume:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UsageReportTotalUnion) AsAny() anyUsageReportTotal {
	switch u.Type {
	case "ai_cluster":
		return u.AsAICluster()
	case "ai_virtual_cluster":
		return u.AsAIVirtualCluster()
	case "baremetal":
		return u.AsBaremetal()
	case "basic_vm":
		return u.AsBasicVm()
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

func (u UsageReportTotalUnion) AsAICluster() (v UsageReportTotalAICluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsAIVirtualCluster() (v UsageReportTotalAIVirtualCluster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsBaremetal() (v UsageReportTotalBaremetal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsBasicVm() (v UsageReportTotalBasicVm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsContainers() (v UsageReportTotalContainers) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsEgressTraffic() (v UsageReportTotalEgressTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsExternalIP() (v UsageReportTotalExternalIP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsFileShare() (v UsageReportTotalFileShare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsFloatingip() (v UsageReportTotalFloatingip) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsFunctions() (v UsageReportTotalFunctions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsFunctionsCalls() (v UsageReportTotalFunctionsCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsFunctionsTraffic() (v UsageReportTotalFunctionsTraffic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsImage() (v UsageReportTotalImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsInference() (v UsageReportTotalInference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsInstance() (v UsageReportTotalInstance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsLoadBalancer() (v UsageReportTotalLoadBalancer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsLogIndex() (v UsageReportTotalLogIndex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsSnapshot() (v UsageReportTotalSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsVolume() (v UsageReportTotalVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsDbaasPostgreSQLConnectionPooler() (v UsageReportTotalDbaasPostgreSQLConnectionPooler) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsDbaasPostgreSQLMemory() (v UsageReportTotalDbaasPostgreSQLMemory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsDbaasPostgreSQLPublicNetwork() (v UsageReportTotalDbaasPostgreSQLPublicNetwork) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsDbaasPostgreSQLCPU() (v UsageReportTotalDbaasPostgreSQLCPU) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UsageReportTotalUnion) AsDbaasPostgreSQLVolume() (v UsageReportTotalDbaasPostgreSQLVolume) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UsageReportTotalUnion) RawJSON() string { return u.JSON.raw }

func (r *UsageReportTotalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalAICluster struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the Baremetal GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.AICluster `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalAICluster) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalAIVirtualCluster struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the Virtual GPU cluster
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                     `json:"region_id" api:"required"`
	Type     constant.AIVirtualCluster `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalAIVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalBaremetal struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Baremetal `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalBaremetal) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalBasicVm struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64            `json:"region_id" api:"required"`
	Type     constant.BasicVm `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalBasicVm) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalContainers struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.Containers `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalContainers) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalEgressTraffic struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Bytes `json:"billing_value_unit" api:"required"`
	// Type of the instance
	//
	// Any of "baremetal", "vm".
	InstanceType string `json:"instance_type" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                  `json:"region_id" api:"required"`
	Type     constant.EgressTraffic `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		InstanceType      respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalExternalIP struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.ExternalIP `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalExternalIP) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalFileShare struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Type of the file share
	FileShareType string `json:"file_share_type" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.FileShare `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		FileShareType     respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalFileShare) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalFloatingip struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64               `json:"region_id" api:"required"`
	Type     constant.Floatingip `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalFloatingip) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalFunctions struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GBs `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Functions `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalFunctions) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalFunctionsCalls struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Mls `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                   `json:"region_id" api:"required"`
	Type     constant.FunctionsCalls `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalFunctionsCalls) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalFunctionsTraffic struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.GB `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                     `json:"region_id" api:"required"`
	Type     constant.FunctionsTraffic `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalFunctionsTraffic) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalImage struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64          `json:"region_id" api:"required"`
	Type     constant.Image `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalImage) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalInference struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64              `json:"region_id" api:"required"`
	Type     constant.Inference `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalInference) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalInstance struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the instance
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.Instance `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalInstance) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalLoadBalancer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                 `json:"region_id" api:"required"`
	Type     constant.LoadBalancer `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Flavor            respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalLogIndex struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.LogIndex `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalLogIndex) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalSnapshot struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64             `json:"region_id" api:"required"`
	Type     constant.Snapshot `json:"type" api:"required"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalSnapshot) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalVolume struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64           `json:"region_id" api:"required"`
	Type     constant.Volume `json:"type" api:"required"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalVolume) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalDbaasPostgreSQLConnectionPooler struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                    `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLConnectionPooler `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalDbaasPostgreSQLConnectionPooler) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalDbaasPostgreSQLMemory struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                          `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLMemory `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalDbaasPostgreSQLMemory) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalDbaasPostgreSQLPublicNetwork struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                                 `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLPublicNetwork `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalDbaasPostgreSQLPublicNetwork) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalDbaasPostgreSQLCPU struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Minutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                       `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLCPU `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalDbaasPostgreSQLCPU) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportTotalDbaasPostgreSQLVolume struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name" api:"required"`
	// Value of the billing metric
	BillingValue float64 `json:"billing_value" api:"required"`
	// Unit of billing value
	BillingValueUnit constant.Gbminutes `json:"billing_value_unit" api:"required"`
	// Region ID
	Region int64 `json:"region" api:"required"`
	// Region ID
	RegionID int64                          `json:"region_id" api:"required"`
	Type     constant.DbaasPostgreSQLVolume `json:"type" api:"required"`
	// Type of the volume
	VolumeType string `json:"volume_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingMetricName respjson.Field
		BillingValue      respjson.Field
		BillingValueUnit  respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		Type              respjson.Field
		VolumeType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportTotalDbaasPostgreSQLVolume) RawJSON() string { return r.JSON.raw }
func (r *UsageReportTotalDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportGetParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day in UTC.
	TimeFrom time.Time `json:"time_from" api:"required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day in UTC.
	TimeTo time.Time `json:"time_to" api:"required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Opt[bool] `json:"enable_last_day,omitzero"`
	// The response resources limit. Defaults to 10.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The response resources offset.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// List of project IDs
	Projects []int64 `json:"projects,omitzero"`
	// List of region IDs.
	Regions []int64 `json:"regions,omitzero"`
	// Extended filter for field filtering.
	SchemaFilter UsageReportGetParamsSchemaFilterUnion `json:"schema_filter,omitzero"`
	// List of sorting filters (JSON objects) fields: project. directions: asc, desc.
	Sorting []UsageReportGetParamsSorting `json:"sorting,omitzero"`
	// Filter by tags
	Tags UsageReportGetParamsTags `json:"tags,omitzero"`
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

func (r UsageReportGetParams) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UsageReportGetParamsSchemaFilterUnion struct {
	OfSnapshot                        *UsageReportGetParamsSchemaFilterSnapshot                        `json:",omitzero,inline"`
	OfInstance                        *UsageReportGetParamsSchemaFilterInstance                        `json:",omitzero,inline"`
	OfAICluster                       *UsageReportGetParamsSchemaFilterAICluster                       `json:",omitzero,inline"`
	OfAIVirtualCluster                *UsageReportGetParamsSchemaFilterAIVirtualCluster                `json:",omitzero,inline"`
	OfBasicVm                         *UsageReportGetParamsSchemaFilterBasicVm                         `json:",omitzero,inline"`
	OfBaremetal                       *UsageReportGetParamsSchemaFilterBaremetal                       `json:",omitzero,inline"`
	OfVolume                          *UsageReportGetParamsSchemaFilterVolume                          `json:",omitzero,inline"`
	OfFileShare                       *UsageReportGetParamsSchemaFilterFileShare                       `json:",omitzero,inline"`
	OfImage                           *UsageReportGetParamsSchemaFilterImage                           `json:",omitzero,inline"`
	OfFloatingip                      *UsageReportGetParamsSchemaFilterFloatingip                      `json:",omitzero,inline"`
	OfEgressTraffic                   *UsageReportGetParamsSchemaFilterEgressTraffic                   `json:",omitzero,inline"`
	OfLoadBalancer                    *UsageReportGetParamsSchemaFilterLoadBalancer                    `json:",omitzero,inline"`
	OfExternalIP                      *UsageReportGetParamsSchemaFilterExternalIP                      `json:",omitzero,inline"`
	OfBackup                          *UsageReportGetParamsSchemaFilterBackup                          `json:",omitzero,inline"`
	OfLogIndex                        *UsageReportGetParamsSchemaFilterLogIndex                        `json:",omitzero,inline"`
	OfFunctions                       *UsageReportGetParamsSchemaFilterFunctions                       `json:",omitzero,inline"`
	OfFunctionsCalls                  *UsageReportGetParamsSchemaFilterFunctionsCalls                  `json:",omitzero,inline"`
	OfFunctionsTraffic                *UsageReportGetParamsSchemaFilterFunctionsTraffic                `json:",omitzero,inline"`
	OfContainers                      *UsageReportGetParamsSchemaFilterContainers                      `json:",omitzero,inline"`
	OfInference                       *UsageReportGetParamsSchemaFilterInference                       `json:",omitzero,inline"`
	OfDbaasPostgreSQLVolume           *UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume           `json:",omitzero,inline"`
	OfDbaasPostgreSQLPublicNetwork    *UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork    `json:",omitzero,inline"`
	OfDbaasPostgreSQLCPU              *UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU              `json:",omitzero,inline"`
	OfDbaasPostgreSQLMemory           *UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory           `json:",omitzero,inline"`
	OfDbaasPostgreSQLConnectionPooler *UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler `json:",omitzero,inline"`
	paramUnion
}

func (u UsageReportGetParamsSchemaFilterUnion) MarshalJSON() ([]byte, error) {
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
func (u *UsageReportGetParamsSchemaFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UsageReportGetParamsSchemaFilterUnion) asAny() any {
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
func (u UsageReportGetParamsSchemaFilterUnion) GetField() *string {
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
func (u UsageReportGetParamsSchemaFilterUnion) GetType() *string {
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
func (u UsageReportGetParamsSchemaFilterUnion) GetValues() []string {
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
	apijson.RegisterUnion[UsageReportGetParamsSchemaFilterUnion](
		"type",
		apijson.Discriminator[UsageReportGetParamsSchemaFilterSnapshot]("snapshot"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterInstance]("instance"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterAICluster]("ai_cluster"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterAIVirtualCluster]("ai_virtual_cluster"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterBasicVm]("basic_vm"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterBaremetal]("baremetal"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterVolume]("volume"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterFileShare]("file_share"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterImage]("image"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterFloatingip]("floatingip"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterEgressTraffic]("egress_traffic"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterLoadBalancer]("load_balancer"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterExternalIP]("external_ip"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterBackup]("backup"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterLogIndex]("log_index"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterFunctions]("functions"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterFunctionsCalls]("functions_calls"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterFunctionsTraffic]("functions_traffic"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterContainers]("containers"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterInference]("inference"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume]("dbaas_postgresql_volume"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork]("dbaas_postgresql_public_network"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU]("dbaas_postgresql_cpu"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory]("dbaas_postgresql_memory"),
		apijson.Discriminator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler]("dbaas_postgresql_connection_pooler"),
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterSnapshot struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "source_volume_uuid", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "snapshot".
	Type constant.Snapshot `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterSnapshot](
		"field", "last_name", "last_size", "source_volume_uuid", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterInstance struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "instance".
	Type constant.Instance `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterInstance) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterInstance](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterAICluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "ai_cluster".
	Type constant.AICluster `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterAICluster) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterAICluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterAICluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterAICluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterAIVirtualCluster struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ai_virtual_cluster".
	Type constant.AIVirtualCluster `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterAIVirtualCluster) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterAIVirtualCluster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterAIVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterAIVirtualCluster](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterBasicVm struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "basic_vm".
	Type constant.BasicVm `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterBasicVm) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterBasicVm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterBasicVm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterBasicVm](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterBaremetal struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "baremetal".
	Type constant.Baremetal `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterBaremetal) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterBaremetal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterBaremetal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterBaremetal](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterVolume struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "last_name", "last_size", "type", "uuid",
	// "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "volume".
	Type constant.Volume `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterVolume) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterVolume](
		"field", "attached_to_vm", "last_name", "last_size", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterFileShare struct {
	// Field name to filter by
	//
	// Any of "file_share_type", "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "file_share".
	Type constant.FileShare `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterFileShare) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterFileShare](
		"field", "file_share_type", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterImage struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterImage) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterImage](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterFloatingip struct {
	// Field name to filter by
	//
	// Any of "ip_address", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "floatingip".
	Type constant.Floatingip `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterFloatingip) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterFloatingip
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterFloatingip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterFloatingip](
		"field", "ip_address", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterEgressTraffic struct {
	// Field name to filter by
	//
	// Any of "instance_name", "instance_type", "port_id", "type", "vm_id".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "egress_traffic".
	Type constant.EgressTraffic `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterEgressTraffic) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterEgressTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterEgressTraffic](
		"field", "instance_name", "instance_type", "port_id", "type", "vm_id",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterLoadBalancer struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "load_balancer".
	Type constant.LoadBalancer `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterLoadBalancer) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterLoadBalancer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterLoadBalancer](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterExternalIP struct {
	// Field name to filter by
	//
	// Any of "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id",
	// "type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "external_ip".
	Type constant.ExternalIP `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterExternalIP) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterExternalIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterExternalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterExternalIP](
		"field", "attached_to_vm", "ip_address", "network_id", "port_id", "subnet_id", "type",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterBackup struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "schedule_id", "source_volume_uuid", "type",
	// "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "backup".
	Type constant.Backup `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterBackup) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterBackup](
		"field", "last_name", "last_size", "schedule_id", "source_volume_uuid", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterLogIndex struct {
	// Field name to filter by
	//
	// Any of "last_name", "last_size", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "log_index".
	Type constant.LogIndex `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterLogIndex) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterLogIndex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterLogIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterLogIndex](
		"field", "last_name", "last_size", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterFunctions struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "functions".
	Type constant.Functions `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterFunctions) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterFunctions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterFunctions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterFunctions](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterFunctionsCalls struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "functions_calls".
	Type constant.FunctionsCalls `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterFunctionsCalls) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterFunctionsCalls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterFunctionsCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterFunctionsCalls](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterFunctionsTraffic struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "functions_traffic".
	Type constant.FunctionsTraffic `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterFunctionsTraffic) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterFunctionsTraffic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterFunctionsTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterFunctionsTraffic](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterContainers struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "containers".
	Type constant.Containers `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterContainers) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterContainers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterContainers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterContainers](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterInference struct {
	// Field name to filter by
	//
	// Any of "flavor", "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "inference".
	Type constant.Inference `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterInference) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterInference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterInference](
		"field", "flavor", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid", "volume_type".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_volume".
	Type constant.DbaasPostgreSQLVolume `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLVolume](
		"field", "last_name", "type", "uuid", "volume_type",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_public_network".
	Type constant.DbaasPostgreSQLPublicNetwork `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLPublicNetwork](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_cpu".
	Type constant.DbaasPostgreSQLCPU `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLCPU](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_memory".
	Type constant.DbaasPostgreSQLMemory `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLMemory](
		"field", "last_name", "type", "uuid",
	)
}

// The properties Field, Type, Values are required.
type UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler struct {
	// Field name to filter by
	//
	// Any of "last_name", "type", "uuid".
	Field string `json:"field,omitzero" api:"required"`
	// List of field values to filter
	Values []string `json:"values,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "dbaas_postgresql_connection_pooler".
	Type constant.DbaasPostgreSQLConnectionPooler `json:"type" api:"required"`
	paramObj
}

func (r UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSchemaFilterDbaasPostgreSQLConnectionPooler](
		"field", "last_name", "type", "uuid",
	)
}

type UsageReportGetParamsSorting struct {
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

func (r UsageReportGetParamsSorting) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsSorting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsSorting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"billing_value", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"first_seen", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"last_name", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"last_seen", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"project", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"region", "asc", "desc",
	)
	apijson.RegisterFieldValidator[UsageReportGetParamsSorting](
		"type", "asc", "desc",
	)
}

// Filter by tags
//
// The property Conditions is required.
type UsageReportGetParamsTags struct {
	// A list of tag filtering conditions defining how tags should match.
	Conditions []UsageReportGetParamsTagsCondition `json:"conditions,omitzero" api:"required"`
	// Specifies whether conditions are combined using OR (default) or AND logic.
	//
	// Any of "AND", "OR".
	ConditionType string `json:"condition_type,omitzero"`
	paramObj
}

func (r UsageReportGetParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsageReportGetParamsTags](
		"condition_type", "AND", "OR",
	)
}

type UsageReportGetParamsTagsCondition struct {
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

func (r UsageReportGetParamsTagsCondition) MarshalJSON() (data []byte, err error) {
	type shadow UsageReportGetParamsTagsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageReportGetParamsTagsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
