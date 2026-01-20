// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// RegionService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r RegionService) {
	r = RegionService{}
	r.Options = opts
	return
}

// List regions
func (r *RegionService) List(ctx context.Context, query RegionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Region], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v1/regions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List regions
func (r *RegionService) ListAutoPaging(ctx context.Context, query RegionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Region] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get region
func (r *RegionService) Get(ctx context.Context, params RegionGetParams, opts ...option.RequestOption) (res *Region, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/regions/%v", params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type Region struct {
	// Region ID
	ID int64 `json:"id,required"`
	// The access level of the region.
	//
	// Any of "core", "edge".
	AccessLevel RegionAccessLevel `json:"access_level,required"`
	// List of available volume types, 'standard', 'ssd_hiiops', 'cold'].
	AvailableVolumeTypes []string `json:"available_volume_types,required"`
	// Coordinates of the region
	Coordinates RegionCoordinates `json:"coordinates,required"`
	// Two-letter country code, ISO 3166-1 alpha-2
	Country string `json:"country,required"`
	// Region creation date and time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field is deprecated. Use `created_at` instead.
	//
	// Deprecated: deprecated
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Human-readable region name
	DisplayName string `json:"display_name,required"`
	// Endpoint type
	//
	// Any of "admin", "internal", "public".
	EndpointType RegionEndpointType `json:"endpoint_type,required"`
	// External network ID for Neutron
	ExternalNetworkID string `json:"external_network_id,required"`
	// List of available file share types
	//
	// Any of "standard", "vast".
	FileShareTypes []string `json:"file_share_types,required"`
	// Region has AI capability
	HasAI bool `json:"has_ai,required"`
	// Region has AI GPU capability
	HasAIGPU bool `json:"has_ai_gpu,required"`
	// Region has bare metal capability
	HasBaremetal bool `json:"has_baremetal,required"`
	// Region has basic vm capability
	HasBasicVm bool `json:"has_basic_vm,required"`
	// Region has DBAAS service
	HasDbaas bool `json:"has_dbaas,required"`
	// Region has Advanced DDoS Protection capability
	HasDDOS bool `json:"has_ddos,required"`
	// Region has managed kubernetes capability
	HasK8S bool `json:"has_k8s,required"`
	// Region has KVM virtualization capability
	HasKvm bool `json:"has_kvm,required"`
	// Region has SFS capability
	HasSfs bool `json:"has_sfs,required"`
	// Foreign key to Keystone entity
	KeystoneID int64 `json:"keystone_id,required"`
	// Technical region name
	KeystoneName string `json:"keystone_name,required"`
	// Foreign key to Metrics database entity
	MetricsDatabaseID int64 `json:"metrics_database_id,required"`
	// Region state
	//
	// Any of "ACTIVE", "DELETED", "DELETING", "DELETION_FAILED", "INACTIVE",
	// "MAINTENANCE", "NEW".
	State RegionState `json:"state,required"`
	// This field is deprecated and can be ignored
	//
	// Deprecated: deprecated
	TaskID string `json:"task_id,required"`
	// Physical network name to create vlan networks
	VlanPhysicalNetwork string `json:"vlan_physical_network,required"`
	// Geographical zone
	//
	// Any of "AMERICAS", "APAC", "EMEA", "RUSSIA_AND_CIS".
	Zone RegionZone `json:"zone,required"`
	// DDoS endpoint ID
	//
	// Deprecated: deprecated
	DDOSEndpointID int64 `json:"ddos_endpoint_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AccessLevel          respjson.Field
		AvailableVolumeTypes respjson.Field
		Coordinates          respjson.Field
		Country              respjson.Field
		CreatedAt            respjson.Field
		CreatedOn            respjson.Field
		DisplayName          respjson.Field
		EndpointType         respjson.Field
		ExternalNetworkID    respjson.Field
		FileShareTypes       respjson.Field
		HasAI                respjson.Field
		HasAIGPU             respjson.Field
		HasBaremetal         respjson.Field
		HasBasicVm           respjson.Field
		HasDbaas             respjson.Field
		HasDDOS              respjson.Field
		HasK8S               respjson.Field
		HasKvm               respjson.Field
		HasSfs               respjson.Field
		KeystoneID           respjson.Field
		KeystoneName         respjson.Field
		MetricsDatabaseID    respjson.Field
		State                respjson.Field
		TaskID               respjson.Field
		VlanPhysicalNetwork  respjson.Field
		Zone                 respjson.Field
		DDOSEndpointID       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Region) RawJSON() string { return r.JSON.raw }
func (r *Region) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The access level of the region.
type RegionAccessLevel string

const (
	RegionAccessLevelCore RegionAccessLevel = "core"
	RegionAccessLevelEdge RegionAccessLevel = "edge"
)

// Coordinates of the region
type RegionCoordinates struct {
	Latitude  RegionCoordinatesLatitudeUnion  `json:"latitude,required"`
	Longitude RegionCoordinatesLongitudeUnion `json:"longitude,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCoordinates) RawJSON() string { return r.JSON.raw }
func (r *RegionCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RegionCoordinatesLatitudeUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type RegionCoordinatesLatitudeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u RegionCoordinatesLatitudeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RegionCoordinatesLatitudeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RegionCoordinatesLatitudeUnion) RawJSON() string { return u.JSON.raw }

func (r *RegionCoordinatesLatitudeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RegionCoordinatesLongitudeUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type RegionCoordinatesLongitudeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u RegionCoordinatesLongitudeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RegionCoordinatesLongitudeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RegionCoordinatesLongitudeUnion) RawJSON() string { return u.JSON.raw }

func (r *RegionCoordinatesLongitudeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint type
type RegionEndpointType string

const (
	RegionEndpointTypeAdmin    RegionEndpointType = "admin"
	RegionEndpointTypeInternal RegionEndpointType = "internal"
	RegionEndpointTypePublic   RegionEndpointType = "public"
)

// Region state
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

// Geographical zone
type RegionZone string

const (
	RegionZoneAmericas     RegionZone = "AMERICAS"
	RegionZoneApac         RegionZone = "APAC"
	RegionZoneEmea         RegionZone = "EMEA"
	RegionZoneRussiaAndCis RegionZone = "RUSSIA_AND_CIS"
)

type RegionListParams struct {
	// Limit the number of returned regions. Falls back to default of 100 if not
	// specified. Limited by max limit value of 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// If true, null `available_volume_type` is replaced with a list of available
	// volume types.
	ShowVolumeTypes param.Opt[bool] `query:"show_volume_types,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "created_at.asc", "created_at.desc", "display_name.asc",
	// "display_name.desc".
	OrderBy RegionListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// If defined then return only regions that support given product.
	//
	// Any of "containers", "inference".
	Product RegionListParamsProduct `query:"product,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegionListParams]'s query parameters as `url.Values`.
func (r RegionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field and direction.
type RegionListParamsOrderBy string

const (
	RegionListParamsOrderByCreatedAtAsc    RegionListParamsOrderBy = "created_at.asc"
	RegionListParamsOrderByCreatedAtDesc   RegionListParamsOrderBy = "created_at.desc"
	RegionListParamsOrderByDisplayNameAsc  RegionListParamsOrderBy = "display_name.asc"
	RegionListParamsOrderByDisplayNameDesc RegionListParamsOrderBy = "display_name.desc"
)

// If defined then return only regions that support given product.
type RegionListParamsProduct string

const (
	RegionListParamsProductContainers RegionListParamsProduct = "containers"
	RegionListParamsProductInference  RegionListParamsProduct = "inference"
)

type RegionGetParams struct {
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// If true, null `available_volume_type` is replaced with a list of available
	// volume types.
	ShowVolumeTypes param.Opt[bool] `query:"show_volume_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegionGetParams]'s query parameters as `url.Values`.
func (r RegionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
