// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/shared"
	"github.com/tidwall/gjson"
)

// CloudV1RegionService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RegionService] method instead.
type CloudV1RegionService struct {
	Options []option.RequestOption
}

// NewCloudV1RegionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1RegionService(opts ...option.RequestOption) (r *CloudV1RegionService) {
	r = &CloudV1RegionService{}
	r.Options = opts
	return
}

// Get region
func (r *CloudV1RegionService) Get(ctx context.Context, regionID int64, query CloudV1RegionGetParams, opts ...option.RequestOption) (res *Region, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/regions/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List regions
func (r *CloudV1RegionService) List(ctx context.Context, query CloudV1RegionListParams, opts ...option.RequestOption) (res *CloudV1RegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Region struct {
	// Region ID. It automatically is being generated when it is created
	ID int64 `json:"id,required"`
	// The access level of the region.
	AccessLevel RegionAccessType `json:"access_level,required"`
	// AI service API endpoint ID
	AIServiceEndpointID int64 `json:"ai_service_endpoint_id,required,nullable"`
	// List of available volume types, 'standard', 'ssd_hiiops', 'cold'].
	AvailableVolumeTypes []string `json:"available_volume_types,required,nullable"`
	// Coordinates of the region
	Coordinates RegionCoordinatesUnion `json:"coordinates,required,nullable"`
	// Country
	Country string `json:"country,required,nullable"`
	// Datetime object. It automatically is being generated when it is created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DDoS endpoint ID
	DdosEndpointID int64 `json:"ddos_endpoint_id,required,nullable"`
	// Unique constraint. Display name
	DisplayName string `json:"display_name,required"`
	// Endpoint type
	EndpointType RegionEndpointType `json:"endpoint_type,required"`
	// External network ID for Neutron
	ExternalNetworkID string `json:"external_network_id,required,nullable"`
	// Region has AI capability
	HasAI bool `json:"has_ai,required"`
	// Region has AI GPU capability
	HasAIGPU bool `json:"has_ai_gpu,required"`
	// Region has bare metal capability
	HasBaremetal bool `json:"has_baremetal,required"`
	// Region has basic vm capability
	HasBasicVm bool `json:"has_basic_vm,required"`
	// Region has managed kubernetes capability
	HasK8s bool `json:"has_k8s,required"`
	// Region has KVM virtualization capability
	HasKvm bool `json:"has_kvm,required"`
	// Region has SFS capability
	HasSfs bool `json:"has_sfs,required"`
	// Foreign key to Keystone entity
	KeystoneID int64 `json:"keystone_id,required"`
	// Region name exactly as stated in keystone
	KeystoneName string `json:"keystone_name,required"`
	// Foreign key to Metrics database entity
	MetricsDatabaseID int64 `json:"metrics_database_id,required,nullable"`
	// Region state
	State RegionState `json:"state,required"`
	// UUID. Id of the Task entity that is handling region's state transition
	TaskID string `json:"task_id,required,nullable"`
	// Physical network name to create vlan networks
	VlanPhysicalNetwork string `json:"vlan_physical_network,required"`
	// Zone
	Zone Zone `json:"zone,required,nullable"`
	// Region has DBAAS service
	HasDbaas bool       `json:"has_dbaas"`
	JSON     regionJSON `json:"-"`
}

// regionJSON contains the JSON metadata for the struct [Region]
type regionJSON struct {
	ID                   apijson.Field
	AccessLevel          apijson.Field
	AIServiceEndpointID  apijson.Field
	AvailableVolumeTypes apijson.Field
	Coordinates          apijson.Field
	Country              apijson.Field
	CreatedOn            apijson.Field
	DdosEndpointID       apijson.Field
	DisplayName          apijson.Field
	EndpointType         apijson.Field
	ExternalNetworkID    apijson.Field
	HasAI                apijson.Field
	HasAIGPU             apijson.Field
	HasBaremetal         apijson.Field
	HasBasicVm           apijson.Field
	HasK8s               apijson.Field
	HasKvm               apijson.Field
	HasSfs               apijson.Field
	KeystoneID           apijson.Field
	KeystoneName         apijson.Field
	MetricsDatabaseID    apijson.Field
	State                apijson.Field
	TaskID               apijson.Field
	VlanPhysicalNetwork  apijson.Field
	Zone                 apijson.Field
	HasDbaas             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Region) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionJSON) RawJSON() string {
	return r.raw
}

// Coordinates of the region
//
// Union satisfied by [RegionCoordinatesCoordinate], [RegionCoordinatesArray] or
// [shared.UnionString].
type RegionCoordinatesUnion interface {
	ImplementsRegionCoordinatesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RegionCoordinatesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RegionCoordinatesCoordinate{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RegionCoordinatesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RegionCoordinatesCoordinate struct {
	Latitude  float64                         `json:"latitude,required"`
	Longitude float64                         `json:"longitude,required"`
	JSON      regionCoordinatesCoordinateJSON `json:"-"`
}

// regionCoordinatesCoordinateJSON contains the JSON metadata for the struct
// [RegionCoordinatesCoordinate]
type regionCoordinatesCoordinateJSON struct {
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionCoordinatesCoordinate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionCoordinatesCoordinateJSON) RawJSON() string {
	return r.raw
}

func (r RegionCoordinatesCoordinate) ImplementsRegionCoordinatesUnion() {}

type RegionCoordinatesArray []interface{}

func (r RegionCoordinatesArray) ImplementsRegionCoordinatesUnion() {}

type RegionAccessType string

const (
	RegionAccessTypeCore RegionAccessType = "core"
	RegionAccessTypeEdge RegionAccessType = "edge"
)

func (r RegionAccessType) IsKnown() bool {
	switch r {
	case RegionAccessTypeCore, RegionAccessTypeEdge:
		return true
	}
	return false
}

type RegionEndpointType string

const (
	RegionEndpointTypeAdmin    RegionEndpointType = "admin"
	RegionEndpointTypeInternal RegionEndpointType = "internal"
	RegionEndpointTypePublic   RegionEndpointType = "public"
)

func (r RegionEndpointType) IsKnown() bool {
	switch r {
	case RegionEndpointTypeAdmin, RegionEndpointTypeInternal, RegionEndpointTypePublic:
		return true
	}
	return false
}

type RegionState string

const (
	RegionStateActive         RegionState = "ACTIVE"
	RegionStateDeleted        RegionState = "DELETED"
	RegionStateDeleting       RegionState = "DELETING"
	RegionStateDeletionFailed RegionState = "DELETION_FAILED"
	RegionStateInactive       RegionState = "INACTIVE"
	RegionStateMaintenance    RegionState = "MAINTENANCE"
	RegionStateNew            RegionState = "NEW"
)

func (r RegionState) IsKnown() bool {
	switch r {
	case RegionStateActive, RegionStateDeleted, RegionStateDeleting, RegionStateDeletionFailed, RegionStateInactive, RegionStateMaintenance, RegionStateNew:
		return true
	}
	return false
}

type Zone string

const (
	ZoneAmericas     Zone = "AMERICAS"
	ZoneApac         Zone = "APAC"
	ZoneEmea         Zone = "EMEA"
	ZoneRussiaAndCis Zone = "RUSSIA_AND_CIS"
)

func (r Zone) IsKnown() bool {
	switch r {
	case ZoneAmericas, ZoneApac, ZoneEmea, ZoneRussiaAndCis:
		return true
	}
	return false
}

type CloudV1RegionListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Region                      `json:"results,required"`
	JSON    cloudV1RegionListResponseJSON `json:"-"`
}

// cloudV1RegionListResponseJSON contains the JSON metadata for the struct
// [CloudV1RegionListResponse]
type cloudV1RegionListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1RegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RegionListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1RegionGetParams struct {
	// If true, null `available_volume_type` is replaced with a list of available
	// volume types.
	ShowVolumeTypes param.Field[bool] `query:"show_volume_types"`
}

// URLQuery serializes [CloudV1RegionGetParams]'s query parameters as `url.Values`.
func (r CloudV1RegionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1RegionListParams struct {
	// Limit the number of returned regions. Falls back to default of 100 if not
	// specified. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Currently supported products: 'containers' and 'inference'. If defined then
	// return only regions with defined product.
	Product param.Field[string] `query:"product"`
	// If true, null `available_volume_type` is replaced with a list of available
	// volume types.
	ShowVolumeTypes param.Field[bool] `query:"show_volume_types"`
}

// URLQuery serializes [CloudV1RegionListParams]'s query parameters as
// `url.Values`.
func (r CloudV1RegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
