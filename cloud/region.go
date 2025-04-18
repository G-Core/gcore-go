// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
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

// Get region
func (r *RegionService) Get(ctx context.Context, params RegionGetParams, opts ...option.RequestOption) (res *Region, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/regions/%v", params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List regions
func (r *RegionService) List(ctx context.Context, query RegionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Region], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// '#/components/schemas/RegionSerializer' "$.components.schemas.RegionSerializer"
type Region struct {
	// '#/components/schemas/RegionSerializer/properties/id'
	// "$.components.schemas.RegionSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegionSerializer/properties/access_level'
	// "$.components.schemas.RegionSerializer.properties.access_level"
	//
	// Any of "core", "edge".
	AccessLevel RegionAccessLevel `json:"access_level,required"`
	// '#/components/schemas/RegionSerializer/properties/ai_service_endpoint_id/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.ai_service_endpoint_id.anyOf[0]"
	AIServiceEndpointID int64 `json:"ai_service_endpoint_id,required"`
	// '#/components/schemas/RegionSerializer/properties/available_volume_types/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.available_volume_types.anyOf[0]"
	AvailableVolumeTypes []string `json:"available_volume_types,required"`
	// '#/components/schemas/RegionSerializer/properties/coordinates/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.coordinates.anyOf[0]"
	Coordinates RegionCoordinates `json:"coordinates,required"`
	// '#/components/schemas/RegionSerializer/properties/country/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.country.anyOf[0]"
	Country string `json:"country,required"`
	// '#/components/schemas/RegionSerializer/properties/created_at'
	// "$.components.schemas.RegionSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RegionSerializer/properties/created_on'
	// "$.components.schemas.RegionSerializer.properties.created_on"
	//
	// Deprecated: deprecated
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// '#/components/schemas/RegionSerializer/properties/ddos_endpoint_id/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.ddos_endpoint_id.anyOf[0]"
	DDOSEndpointID int64 `json:"ddos_endpoint_id,required"`
	// '#/components/schemas/RegionSerializer/properties/display_name'
	// "$.components.schemas.RegionSerializer.properties.display_name"
	DisplayName string `json:"display_name,required"`
	// '#/components/schemas/RegionSerializer/properties/endpoint_type'
	// "$.components.schemas.RegionSerializer.properties.endpoint_type"
	//
	// Any of "admin", "internal", "public".
	EndpointType RegionEndpointType `json:"endpoint_type,required"`
	// '#/components/schemas/RegionSerializer/properties/external_network_id/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.external_network_id.anyOf[0]"
	ExternalNetworkID string `json:"external_network_id,required"`
	// '#/components/schemas/RegionSerializer/properties/file_share_types/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.file_share_types.anyOf[0]"
	//
	// Any of "standard", "vast".
	FileShareTypes []string `json:"file_share_types,required"`
	// '#/components/schemas/RegionSerializer/properties/has_ai'
	// "$.components.schemas.RegionSerializer.properties.has_ai"
	HasAI bool `json:"has_ai,required"`
	// '#/components/schemas/RegionSerializer/properties/has_ai_gpu'
	// "$.components.schemas.RegionSerializer.properties.has_ai_gpu"
	HasAIGPU bool `json:"has_ai_gpu,required"`
	// '#/components/schemas/RegionSerializer/properties/has_baremetal'
	// "$.components.schemas.RegionSerializer.properties.has_baremetal"
	HasBaremetal bool `json:"has_baremetal,required"`
	// '#/components/schemas/RegionSerializer/properties/has_basic_vm'
	// "$.components.schemas.RegionSerializer.properties.has_basic_vm"
	HasBasicVm bool `json:"has_basic_vm,required"`
	// '#/components/schemas/RegionSerializer/properties/has_k8s'
	// "$.components.schemas.RegionSerializer.properties.has_k8s"
	HasK8s bool `json:"has_k8s,required"`
	// '#/components/schemas/RegionSerializer/properties/has_kvm'
	// "$.components.schemas.RegionSerializer.properties.has_kvm"
	HasKvm bool `json:"has_kvm,required"`
	// '#/components/schemas/RegionSerializer/properties/has_sfs'
	// "$.components.schemas.RegionSerializer.properties.has_sfs"
	HasSfs bool `json:"has_sfs,required"`
	// '#/components/schemas/RegionSerializer/properties/keystone_id'
	// "$.components.schemas.RegionSerializer.properties.keystone_id"
	KeystoneID int64 `json:"keystone_id,required"`
	// '#/components/schemas/RegionSerializer/properties/keystone_name'
	// "$.components.schemas.RegionSerializer.properties.keystone_name"
	KeystoneName string `json:"keystone_name,required"`
	// '#/components/schemas/RegionSerializer/properties/metrics_database_id/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.metrics_database_id.anyOf[0]"
	MetricsDatabaseID int64 `json:"metrics_database_id,required"`
	// '#/components/schemas/RegionSerializer/properties/state'
	// "$.components.schemas.RegionSerializer.properties.state"
	//
	// Any of "ACTIVE", "DELETED", "DELETING", "DELETION_FAILED", "INACTIVE",
	// "MAINTENANCE", "NEW".
	State RegionState `json:"state,required"`
	// '#/components/schemas/RegionSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.task_id.anyOf[0]"
	//
	// Deprecated: deprecated
	TaskID string `json:"task_id,required"`
	// '#/components/schemas/RegionSerializer/properties/vlan_physical_network'
	// "$.components.schemas.RegionSerializer.properties.vlan_physical_network"
	VlanPhysicalNetwork string `json:"vlan_physical_network,required"`
	// '#/components/schemas/RegionSerializer/properties/zone/anyOf/0'
	// "$.components.schemas.RegionSerializer.properties.zone.anyOf[0]"
	//
	// Any of "AMERICAS", "APAC", "EMEA", "RUSSIA_AND_CIS".
	Zone RegionZone `json:"zone,required"`
	// '#/components/schemas/RegionSerializer/properties/has_dbaas'
	// "$.components.schemas.RegionSerializer.properties.has_dbaas"
	HasDbaas bool `json:"has_dbaas"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                   resp.Field
		AccessLevel          resp.Field
		AIServiceEndpointID  resp.Field
		AvailableVolumeTypes resp.Field
		Coordinates          resp.Field
		Country              resp.Field
		CreatedAt            resp.Field
		CreatedOn            resp.Field
		DDOSEndpointID       resp.Field
		DisplayName          resp.Field
		EndpointType         resp.Field
		ExternalNetworkID    resp.Field
		FileShareTypes       resp.Field
		HasAI                resp.Field
		HasAIGPU             resp.Field
		HasBaremetal         resp.Field
		HasBasicVm           resp.Field
		HasK8s               resp.Field
		HasKvm               resp.Field
		HasSfs               resp.Field
		KeystoneID           resp.Field
		KeystoneName         resp.Field
		MetricsDatabaseID    resp.Field
		State                resp.Field
		TaskID               resp.Field
		VlanPhysicalNetwork  resp.Field
		Zone                 resp.Field
		HasDbaas             resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Region) RawJSON() string { return r.JSON.raw }
func (r *Region) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegionSerializer/properties/access_level'
// "$.components.schemas.RegionSerializer.properties.access_level"
type RegionAccessLevel string

const (
	RegionAccessLevelCore RegionAccessLevel = "core"
	RegionAccessLevelEdge RegionAccessLevel = "edge"
)

// '#/components/schemas/RegionSerializer/properties/coordinates/anyOf/0'
// "$.components.schemas.RegionSerializer.properties.coordinates.anyOf[0]"
type RegionCoordinates struct {
	// '#/components/schemas/Coordinate/properties/latitude'
	// "$.components.schemas.Coordinate.properties.latitude"
	Latitude float64 `json:"latitude,required"`
	// '#/components/schemas/Coordinate/properties/longitude'
	// "$.components.schemas.Coordinate.properties.longitude"
	Longitude float64 `json:"longitude,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Latitude    resp.Field
		Longitude   resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCoordinates) RawJSON() string { return r.JSON.raw }
func (r *RegionCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegionSerializer/properties/endpoint_type'
// "$.components.schemas.RegionSerializer.properties.endpoint_type"
type RegionEndpointType string

const (
	RegionEndpointTypeAdmin    RegionEndpointType = "admin"
	RegionEndpointTypeInternal RegionEndpointType = "internal"
	RegionEndpointTypePublic   RegionEndpointType = "public"
)

// '#/components/schemas/RegionSerializer/properties/state'
// "$.components.schemas.RegionSerializer.properties.state"
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

// '#/components/schemas/RegionSerializer/properties/zone/anyOf/0'
// "$.components.schemas.RegionSerializer.properties.zone.anyOf[0]"
type RegionZone string

const (
	RegionZoneAmericas     RegionZone = "AMERICAS"
	RegionZoneApac         RegionZone = "APAC"
	RegionZoneEmea         RegionZone = "EMEA"
	RegionZoneRussiaAndCis RegionZone = "RUSSIA_AND_CIS"
)

type RegionGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregions%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/regions/{region_id}'].get.parameters[0].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregions%2F%7Bregion_id%7D/get/parameters/1'
	// "$.paths['/cloud/v1/regions/{region_id}'].get.parameters[1]"
	ShowVolumeTypes param.Opt[bool] `query:"show_volume_types,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegionGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RegionGetParams]'s query parameters as `url.Values`.
func (r RegionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegionListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/0'
	// "$.paths['/cloud/v1/regions'].get.parameters[0]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/1'
	// "$.paths['/cloud/v1/regions'].get.parameters[1]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/4'
	// "$.paths['/cloud/v1/regions'].get.parameters[4]"
	ShowVolumeTypes param.Opt[bool] `query:"show_volume_types,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/2'
	// "$.paths['/cloud/v1/regions'].get.parameters[2]"
	//
	// Any of "created_at.asc", "created_at.desc", "display_name.asc",
	// "display_name.desc".
	OrderBy RegionListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/3'
	// "$.paths['/cloud/v1/regions'].get.parameters[3]"
	//
	// Any of "containers", "inference".
	Product RegionListParamsProduct `query:"product,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegionListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RegionListParams]'s query parameters as `url.Values`.
func (r RegionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/2'
// "$.paths['/cloud/v1/regions'].get.parameters[2]"
type RegionListParamsOrderBy string

const (
	RegionListParamsOrderByCreatedAtAsc    RegionListParamsOrderBy = "created_at.asc"
	RegionListParamsOrderByCreatedAtDesc   RegionListParamsOrderBy = "created_at.desc"
	RegionListParamsOrderByDisplayNameAsc  RegionListParamsOrderBy = "display_name.asc"
	RegionListParamsOrderByDisplayNameDesc RegionListParamsOrderBy = "display_name.desc"
)

// '#/paths/%2Fcloud%2Fv1%2Fregions/get/parameters/3'
// "$.paths['/cloud/v1/regions'].get.parameters[3]"
type RegionListParamsProduct string

const (
	RegionListParamsProductContainers RegionListParamsProduct = "containers"
	RegionListParamsProductInference  RegionListParamsProduct = "inference"
)
