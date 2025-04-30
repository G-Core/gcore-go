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

// InstanceImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceImageService] method instead.
type InstanceImageService struct {
	Options []option.RequestOption
}

// NewInstanceImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceImageService(opts ...option.RequestOption) (r InstanceImageService) {
	r = InstanceImageService{}
	r.Options = opts
	return
}

// Update image fields
func (r *InstanceImageService) Update(ctx context.Context, imageID string, params InstanceImageUpdateParams, opts ...option.RequestOption) (res *InstanceImageUpdateResponse, err error) {
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
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve an available images list. Returned entities owned by the project and
// public OR shared with the client
func (r *InstanceImageService) List(ctx context.Context, params InstanceImageListParams, opts ...option.RequestOption) (res *InstanceImageListResponse, err error) {
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
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete the image
func (r *InstanceImageService) Delete(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create image from volume
func (r *InstanceImageService) NewFromVolume(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get image
func (r *InstanceImageService) Get(ctx context.Context, imageID string, params InstanceImageGetParams, opts ...option.RequestOption) (res *InstanceImageGetResponse, err error) {
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
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Upload image
func (r *InstanceImageService) Upload(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/downloadimage/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/patch/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].patch.responses[200].content['application/json'].schema"
type InstanceImageUpdateResponse struct {
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
	OsType InstanceImageUpdateResponseOsType `json:"os_type,required"`
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
	Architecture InstanceImageUpdateResponseArchitecture `json:"architecture"`
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
	HwFirmwareType InstanceImageUpdateResponseHwFirmwareType `json:"hw_firmware_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType InstanceImageUpdateResponseHwMachineType `json:"hw_machine_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/is_baremetal/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.is_baremetal.anyOf[0]"
	IsBaremetal bool `json:"is_baremetal,nullable"`
	// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUpdateResponseSSHKey `json:"ssh_key,nullable"`
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
func (r InstanceImageUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *InstanceImageUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageSerializer/properties/os_type'
// "$.components.schemas.ImageSerializer.properties.os_type"
type InstanceImageUpdateResponseOsType string

const (
	InstanceImageUpdateResponseOsTypeLinux   InstanceImageUpdateResponseOsType = "linux"
	InstanceImageUpdateResponseOsTypeWindows InstanceImageUpdateResponseOsType = "windows"
)

// '#/components/schemas/ImageSerializer/properties/architecture'
// "$.components.schemas.ImageSerializer.properties.architecture"
type InstanceImageUpdateResponseArchitecture string

const (
	InstanceImageUpdateResponseArchitectureAarch64 InstanceImageUpdateResponseArchitecture = "aarch64"
	InstanceImageUpdateResponseArchitectureX86_64  InstanceImageUpdateResponseArchitecture = "x86_64"
)

// '#/components/schemas/ImageSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_firmware_type.anyOf[0]"
type InstanceImageUpdateResponseHwFirmwareType string

const (
	InstanceImageUpdateResponseHwFirmwareTypeBios InstanceImageUpdateResponseHwFirmwareType = "bios"
	InstanceImageUpdateResponseHwFirmwareTypeUefi InstanceImageUpdateResponseHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
type InstanceImageUpdateResponseHwMachineType string

const (
	InstanceImageUpdateResponseHwMachineTypePc  InstanceImageUpdateResponseHwMachineType = "pc"
	InstanceImageUpdateResponseHwMachineTypeQ35 InstanceImageUpdateResponseHwMachineType = "q35"
)

// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
type InstanceImageUpdateResponseSSHKey string

const (
	InstanceImageUpdateResponseSSHKeyAllow    InstanceImageUpdateResponseSSHKey = "allow"
	InstanceImageUpdateResponseSSHKeyDeny     InstanceImageUpdateResponseSSHKey = "deny"
	InstanceImageUpdateResponseSSHKeyRequired InstanceImageUpdateResponseSSHKey = "required"
)

// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.responses[200].content['application/json'].schema"
type InstanceImageListResponse struct {
	// '#/components/schemas/ImageCollectionSerializer/properties/count'
	// "$.components.schemas.ImageCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ImageCollectionSerializer/properties/results'
	// "$.components.schemas.ImageCollectionSerializer.properties.results"
	Results []InstanceImageListResponseResult `json:"results,required"`
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
func (r InstanceImageListResponse) RawJSON() string { return r.JSON.raw }
func (r *InstanceImageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageCollectionSerializer/properties/results/items'
// "$.components.schemas.ImageCollectionSerializer.properties.results.items"
type InstanceImageListResponseResult struct {
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
func (r InstanceImageListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *InstanceImageListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.responses[200].content['application/json'].schema"
type InstanceImageGetResponse struct {
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
	OsType InstanceImageGetResponseOsType `json:"os_type,required"`
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
	Architecture InstanceImageGetResponseArchitecture `json:"architecture"`
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
	HwFirmwareType InstanceImageGetResponseHwFirmwareType `json:"hw_firmware_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType InstanceImageGetResponseHwMachineType `json:"hw_machine_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/is_baremetal/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.is_baremetal.anyOf[0]"
	IsBaremetal bool `json:"is_baremetal,nullable"`
	// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageGetResponseSSHKey `json:"ssh_key,nullable"`
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
func (r InstanceImageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InstanceImageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageSerializer/properties/os_type'
// "$.components.schemas.ImageSerializer.properties.os_type"
type InstanceImageGetResponseOsType string

const (
	InstanceImageGetResponseOsTypeLinux   InstanceImageGetResponseOsType = "linux"
	InstanceImageGetResponseOsTypeWindows InstanceImageGetResponseOsType = "windows"
)

// '#/components/schemas/ImageSerializer/properties/architecture'
// "$.components.schemas.ImageSerializer.properties.architecture"
type InstanceImageGetResponseArchitecture string

const (
	InstanceImageGetResponseArchitectureAarch64 InstanceImageGetResponseArchitecture = "aarch64"
	InstanceImageGetResponseArchitectureX86_64  InstanceImageGetResponseArchitecture = "x86_64"
)

// '#/components/schemas/ImageSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_firmware_type.anyOf[0]"
type InstanceImageGetResponseHwFirmwareType string

const (
	InstanceImageGetResponseHwFirmwareTypeBios InstanceImageGetResponseHwFirmwareType = "bios"
	InstanceImageGetResponseHwFirmwareTypeUefi InstanceImageGetResponseHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
type InstanceImageGetResponseHwMachineType string

const (
	InstanceImageGetResponseHwMachineTypePc  InstanceImageGetResponseHwMachineType = "pc"
	InstanceImageGetResponseHwMachineTypeQ35 InstanceImageGetResponseHwMachineType = "q35"
)

// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
type InstanceImageGetResponseSSHKey string

const (
	InstanceImageGetResponseSSHKeyAllow    InstanceImageGetResponseSSHKey = "allow"
	InstanceImageGetResponseSSHKeyDeny     InstanceImageGetResponseSSHKey = "deny"
	InstanceImageGetResponseSSHKeyRequired InstanceImageGetResponseSSHKey = "required"
)

type InstanceImageUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/UpdateImageSerializer/properties/is_baremetal'
	// "$.components.schemas.UpdateImageSerializer.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/name'
	// "$.components.schemas.UpdateImageSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/hw_firmware_type'
	// "$.components.schemas.UpdateImageSerializer.properties.hw_firmware_type"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUpdateParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/hw_machine_type'
	// "$.components.schemas.UpdateImageSerializer.properties.hw_machine_type"
	//
	// Any of "pc", "q35".
	HwMachineType InstanceImageUpdateParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/os_type'
	// "$.components.schemas.UpdateImageSerializer.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageUpdateParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/ssh_key'
	// "$.components.schemas.UpdateImageSerializer.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUpdateParamsSSHKey `json:"ssh_key,omitzero"`
	// '#/components/schemas/UpdateImageSerializer/properties/tags'
	// "$.components.schemas.UpdateImageSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UpdateImageSerializer/properties/hw_firmware_type'
// "$.components.schemas.UpdateImageSerializer.properties.hw_firmware_type"
type InstanceImageUpdateParamsHwFirmwareType string

const (
	InstanceImageUpdateParamsHwFirmwareTypeBios InstanceImageUpdateParamsHwFirmwareType = "bios"
	InstanceImageUpdateParamsHwFirmwareTypeUefi InstanceImageUpdateParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/UpdateImageSerializer/properties/hw_machine_type'
// "$.components.schemas.UpdateImageSerializer.properties.hw_machine_type"
type InstanceImageUpdateParamsHwMachineType string

const (
	InstanceImageUpdateParamsHwMachineTypePc  InstanceImageUpdateParamsHwMachineType = "pc"
	InstanceImageUpdateParamsHwMachineTypeQ35 InstanceImageUpdateParamsHwMachineType = "q35"
)

// '#/components/schemas/UpdateImageSerializer/properties/os_type'
// "$.components.schemas.UpdateImageSerializer.properties.os_type"
type InstanceImageUpdateParamsOsType string

const (
	InstanceImageUpdateParamsOsTypeLinux   InstanceImageUpdateParamsOsType = "linux"
	InstanceImageUpdateParamsOsTypeWindows InstanceImageUpdateParamsOsType = "windows"
)

// '#/components/schemas/UpdateImageSerializer/properties/ssh_key'
// "$.components.schemas.UpdateImageSerializer.properties.ssh_key"
type InstanceImageUpdateParamsSSHKey string

const (
	InstanceImageUpdateParamsSSHKeyAllow    InstanceImageUpdateParamsSSHKey = "allow"
	InstanceImageUpdateParamsSSHKeyDeny     InstanceImageUpdateParamsSSHKey = "deny"
	InstanceImageUpdateParamsSSHKeyRequired InstanceImageUpdateParamsSSHKey = "required"
)

type InstanceImageListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[3]"
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[5]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[4]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[6]"
	//
	// Any of "private", "public", "shared".
	Visibility InstanceImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceImageListParams]'s query parameters as
// `url.Values`.
func (r InstanceImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[6]"
type InstanceImageListParamsVisibility string

const (
	InstanceImageListParamsVisibilityPrivate InstanceImageListParamsVisibility = "private"
	InstanceImageListParamsVisibilityPublic  InstanceImageListParamsVisibility = "public"
	InstanceImageListParamsVisibilityShared  InstanceImageListParamsVisibility = "shared"
)

type InstanceImageDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InstanceImageNewFromVolumeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/name'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/volume_id'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.volume_id"
	VolumeID string `json:"volume_id,required" format:"uuid4"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/is_baremetal'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/hw_firmware_type/anyOf/0'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.hw_firmware_type.anyOf[0]"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageNewFromVolumeParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType InstanceImageNewFromVolumeParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/architecture'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageNewFromVolumeParamsArchitecture `json:"architecture,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/os_type'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageNewFromVolumeParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/source'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.source"
	//
	// Any of "volume".
	Source InstanceImageNewFromVolumeParamsSource `json:"source,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/ssh_key'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageNewFromVolumeParamsSSHKey `json:"ssh_key,omitzero"`
	// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/tags'
	// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageNewFromVolumeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageNewFromVolumeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageNewFromVolumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/architecture'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.architecture"
type InstanceImageNewFromVolumeParamsArchitecture string

const (
	InstanceImageNewFromVolumeParamsArchitectureAarch64 InstanceImageNewFromVolumeParamsArchitecture = "aarch64"
	InstanceImageNewFromVolumeParamsArchitectureX86_64  InstanceImageNewFromVolumeParamsArchitecture = "x86_64"
)

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.hw_firmware_type.anyOf[0]"
type InstanceImageNewFromVolumeParamsHwFirmwareType string

const (
	InstanceImageNewFromVolumeParamsHwFirmwareTypeBios InstanceImageNewFromVolumeParamsHwFirmwareType = "bios"
	InstanceImageNewFromVolumeParamsHwFirmwareTypeUefi InstanceImageNewFromVolumeParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/hw_machine_type/anyOf/0'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.hw_machine_type.anyOf[0]"
type InstanceImageNewFromVolumeParamsHwMachineType string

const (
	InstanceImageNewFromVolumeParamsHwMachineTypePc  InstanceImageNewFromVolumeParamsHwMachineType = "pc"
	InstanceImageNewFromVolumeParamsHwMachineTypeQ35 InstanceImageNewFromVolumeParamsHwMachineType = "q35"
)

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/os_type'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.os_type"
type InstanceImageNewFromVolumeParamsOsType string

const (
	InstanceImageNewFromVolumeParamsOsTypeLinux   InstanceImageNewFromVolumeParamsOsType = "linux"
	InstanceImageNewFromVolumeParamsOsTypeWindows InstanceImageNewFromVolumeParamsOsType = "windows"
)

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/source'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.source"
type InstanceImageNewFromVolumeParamsSource string

const (
	InstanceImageNewFromVolumeParamsSourceVolume InstanceImageNewFromVolumeParamsSource = "volume"
)

// '#/components/schemas/ImageCreateFromVolumeSerializer/properties/ssh_key'
// "$.components.schemas.ImageCreateFromVolumeSerializer.properties.ssh_key"
type InstanceImageNewFromVolumeParamsSSHKey string

const (
	InstanceImageNewFromVolumeParamsSSHKeyAllow    InstanceImageNewFromVolumeParamsSSHKey = "allow"
	InstanceImageNewFromVolumeParamsSSHKeyDeny     InstanceImageNewFromVolumeParamsSSHKey = "deny"
	InstanceImageNewFromVolumeParamsSSHKeyRequired InstanceImageNewFromVolumeParamsSSHKey = "required"
)

type InstanceImageGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[3]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceImageGetParams]'s query parameters as `url.Values`.
func (r InstanceImageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceImageUploadParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fdownloadimage%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/downloadimage/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fdownloadimage%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/downloadimage/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ImageDownloadSerializer/properties/name'
	// "$.components.schemas.ImageDownloadSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageDownloadSerializer/properties/url'
	// "$.components.schemas.ImageDownloadSerializer.properties.url"
	URL string `json:"url,required" format:"uri"`
	// '#/components/schemas/ImageDownloadSerializer/properties/os_distro/anyOf/0'
	// "$.components.schemas.ImageDownloadSerializer.properties.os_distro.anyOf[0]"
	OsDistro param.Opt[string] `json:"os_distro,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/os_version/anyOf/0'
	// "$.components.schemas.ImageDownloadSerializer.properties.os_version.anyOf[0]"
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/cow_format'
	// "$.components.schemas.ImageDownloadSerializer.properties.cow_format"
	CowFormat param.Opt[bool] `json:"cow_format,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/is_baremetal'
	// "$.components.schemas.ImageDownloadSerializer.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/hw_firmware_type/anyOf/0'
	// "$.components.schemas.ImageDownloadSerializer.properties.hw_firmware_type.anyOf[0]"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUploadParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageDownloadSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType InstanceImageUploadParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/architecture'
	// "$.components.schemas.ImageDownloadSerializer.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageUploadParamsArchitecture `json:"architecture,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/os_type'
	// "$.components.schemas.ImageDownloadSerializer.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageUploadParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/ssh_key'
	// "$.components.schemas.ImageDownloadSerializer.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUploadParamsSSHKey `json:"ssh_key,omitzero"`
	// '#/components/schemas/ImageDownloadSerializer/properties/tags'
	// "$.components.schemas.ImageDownloadSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageUploadParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ImageDownloadSerializer/properties/architecture'
// "$.components.schemas.ImageDownloadSerializer.properties.architecture"
type InstanceImageUploadParamsArchitecture string

const (
	InstanceImageUploadParamsArchitectureAarch64 InstanceImageUploadParamsArchitecture = "aarch64"
	InstanceImageUploadParamsArchitectureX86_64  InstanceImageUploadParamsArchitecture = "x86_64"
)

// '#/components/schemas/ImageDownloadSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.ImageDownloadSerializer.properties.hw_firmware_type.anyOf[0]"
type InstanceImageUploadParamsHwFirmwareType string

const (
	InstanceImageUploadParamsHwFirmwareTypeBios InstanceImageUploadParamsHwFirmwareType = "bios"
	InstanceImageUploadParamsHwFirmwareTypeUefi InstanceImageUploadParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageDownloadSerializer/properties/hw_machine_type/anyOf/0'
// "$.components.schemas.ImageDownloadSerializer.properties.hw_machine_type.anyOf[0]"
type InstanceImageUploadParamsHwMachineType string

const (
	InstanceImageUploadParamsHwMachineTypePc  InstanceImageUploadParamsHwMachineType = "pc"
	InstanceImageUploadParamsHwMachineTypeQ35 InstanceImageUploadParamsHwMachineType = "q35"
)

// '#/components/schemas/ImageDownloadSerializer/properties/os_type'
// "$.components.schemas.ImageDownloadSerializer.properties.os_type"
type InstanceImageUploadParamsOsType string

const (
	InstanceImageUploadParamsOsTypeLinux   InstanceImageUploadParamsOsType = "linux"
	InstanceImageUploadParamsOsTypeWindows InstanceImageUploadParamsOsType = "windows"
)

// '#/components/schemas/ImageDownloadSerializer/properties/ssh_key'
// "$.components.schemas.ImageDownloadSerializer.properties.ssh_key"
type InstanceImageUploadParamsSSHKey string

const (
	InstanceImageUploadParamsSSHKeyAllow    InstanceImageUploadParamsSSHKey = "allow"
	InstanceImageUploadParamsSSHKeyDeny     InstanceImageUploadParamsSSHKey = "deny"
	InstanceImageUploadParamsSSHKeyRequired InstanceImageUploadParamsSSHKey = "required"
)
