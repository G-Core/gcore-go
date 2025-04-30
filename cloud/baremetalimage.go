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
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// BaremetalImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalImageService] method instead.
type BaremetalImageService struct {
	Options []option.RequestOption
}

// NewBaremetalImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalImageService(opts ...option.RequestOption) (r BaremetalImageService) {
	r = BaremetalImageService{}
	r.Options = opts
	return
}

// Retrieve the available images list for bare metal servers. Returned entities may
// or may not be owned by the project
func (r *BaremetalImageService) List(ctx context.Context, params BaremetalImageListParams, opts ...option.RequestOption) (res *BaremetalImageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.responses[200].content['application/json'].schema"
type BaremetalImageListResponse struct {
	// '#/components/schemas/ImageCollectionSerializer/properties/count'
	// "$.components.schemas.ImageCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ImageCollectionSerializer/properties/results'
	// "$.components.schemas.ImageCollectionSerializer.properties.results"
	Results []BaremetalImageListResponseResult `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalImageListResponse) RawJSON() string { return r.JSON.raw }
func (r *BaremetalImageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageCollectionSerializer/properties/results/items'
// "$.components.schemas.ImageCollectionSerializer.properties.results.items"
type BaremetalImageListResponseResult struct {
	// '#/components/schemas/ImageSerializer/properties/id'
	// "$.components.schemas.ImageSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/ImageSerializer/properties/created_at'
	// "$.components.schemas.ImageSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/ImageSerializer/properties/disk_format'
	// "$.components.schemas.ImageSerializer.properties.disk_format"
	DiskFormat string `json:"disk_format,required"`
	// '#/components/schemas/ImageSerializer/properties/min_disk'
	// "$.components.schemas.ImageSerializer.properties.min_disk"
	MinDisk int64 `json:"min_disk,required"`
	// '#/components/schemas/ImageSerializer/properties/min_ram'
	// "$.components.schemas.ImageSerializer.properties.min_ram"
	MinRam int64 `json:"min_ram,required"`
	// '#/components/schemas/ImageSerializer/properties/name'
	// "$.components.schemas.ImageSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageSerializer/properties/os_distro'
	// "$.components.schemas.ImageSerializer.properties.os_distro"
	OsDistro string `json:"os_distro,required"`
	// '#/components/schemas/ImageSerializer/properties/os_type'
	// "$.components.schemas.ImageSerializer.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType string `json:"os_type,required"`
	// '#/components/schemas/ImageSerializer/properties/os_version'
	// "$.components.schemas.ImageSerializer.properties.os_version"
	OsVersion string `json:"os_version,required"`
	// '#/components/schemas/ImageSerializer/properties/project_id'
	// "$.components.schemas.ImageSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/ImageSerializer/properties/region'
	// "$.components.schemas.ImageSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/ImageSerializer/properties/region_id'
	// "$.components.schemas.ImageSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ImageSerializer/properties/size'
	// "$.components.schemas.ImageSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/ImageSerializer/properties/status'
	// "$.components.schemas.ImageSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/ImageSerializer/properties/tags'
	// "$.components.schemas.ImageSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/ImageSerializer/properties/updated_at'
	// "$.components.schemas.ImageSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/ImageSerializer/properties/visibility'
	// "$.components.schemas.ImageSerializer.properties.visibility"
	Visibility string `json:"visibility,required"`
	// '#/components/schemas/ImageSerializer/properties/architecture'
	// "$.components.schemas.ImageSerializer.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture string `json:"architecture"`
	// '#/components/schemas/ImageSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// '#/components/schemas/ImageSerializer/properties/description/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/ImageSerializer/properties/display_order/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.display_order.anyOf[0]"
	DisplayOrder int64 `json:"display_order,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_firmware_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_firmware_type.anyOf[0]"
	//
	// Any of "bios", "uefi".
	HwFirmwareType string `json:"hw_firmware_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType string `json:"hw_machine_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/is_baremetal/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.is_baremetal.anyOf[0]"
	IsBaremetal bool `json:"is_baremetal,nullable"`
	// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
	//
	// Any of "allow", "deny", "required".
	SSHKey string `json:"ssh_key,nullable"`
	// '#/components/schemas/ImageSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		CreatedAt      resp.Field
		DiskFormat     resp.Field
		MinDisk        resp.Field
		MinRam         resp.Field
		Name           resp.Field
		OsDistro       resp.Field
		OsType         resp.Field
		OsVersion      resp.Field
		ProjectID      resp.Field
		Region         resp.Field
		RegionID       resp.Field
		Size           resp.Field
		Status         resp.Field
		Tags           resp.Field
		UpdatedAt      resp.Field
		Visibility     resp.Field
		Architecture   resp.Field
		CreatorTaskID  resp.Field
		Description    resp.Field
		DisplayOrder   resp.Field
		HwFirmwareType resp.Field
		HwMachineType  resp.Field
		IsBaremetal    resp.Field
		SSHKey         resp.Field
		TaskID         resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalImageListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *BaremetalImageListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalImageListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[3]"
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[5]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[4]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[6]"
	//
	// Any of "private", "public", "shared".
	Visibility BaremetalImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalImageListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BaremetalImageListParams]'s query parameters as
// `url.Values`.
func (r BaremetalImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fbmimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
// "$.paths['/cloud/v1/bmimages/{project_id}/{region_id}'].get.parameters[6]"
type BaremetalImageListParamsVisibility string

const (
	BaremetalImageListParamsVisibilityPrivate BaremetalImageListParamsVisibility = "private"
	BaremetalImageListParamsVisibilityPublic  BaremetalImageListParamsVisibility = "public"
	BaremetalImageListParamsVisibilityShared  BaremetalImageListParamsVisibility = "shared"
)
