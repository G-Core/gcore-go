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

type Region struct {
	// Region ID
	ID int64 `json:"id,required"`
	// The access level of the region.
	//
	// Any of "core", "edge".
	AccessLevel RegionAccessLevel `json:"access_level,required"`
	// AI service API endpoint ID
	AIServiceEndpointID int64 `json:"ai_service_endpoint_id,required"`
	// List of available volume types, 'standard', 'ssd_hiiops', 'cold'].
	AvailableVolumeTypes []string `json:"available_volume_types,required"`
	// Coordinates of the region
	Coordinates RegionCoordinates `json:"coordinates,required"`
	// Country
	Country string `json:"country,required"`
	// Region creation date and time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field is deprecated. Use `created_at` instead.
	//
	// Deprecated: deprecated
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DDoS endpoint ID
	DDOSEndpointID int64 `json:"ddos_endpoint_id,required"`
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
	// Region has managed kubernetes capability
	HasK8s bool `json:"has_k8s,required"`
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
	// Region has DBAAS service
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

// The access level of the region.
type RegionAccessLevel string

const (
	RegionAccessLevelCore RegionAccessLevel = "core"
	RegionAccessLevelEdge RegionAccessLevel = "edge"
)

// Coordinates of the region
type RegionCoordinates struct {
	Latitude  float64 `json:"latitude,required"`
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

type RegionGetParams struct {
	// Region ID
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// If true, null `available_volume_type` is replaced with a list of available
	// volume types.
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
