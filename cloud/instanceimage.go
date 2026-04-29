// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Instance images are operating system images (public, private, or shared) used to
// boot cloud instances.
//
// InstanceImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceImageService] method instead.
type InstanceImageService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewInstanceImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceImageService(opts ...option.RequestOption) (r InstanceImageService) {
	r = InstanceImageService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Update image properties and tags.
func (r *InstanceImageService) Update(ctx context.Context, imageID string, params InstanceImageUpdateParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a list of available images in the project and region. The list can be
// filtered by visibility, tags, and other parameters. Returned entities are owned
// by the project or are public/shared with the client.
func (r *InstanceImageService) List(ctx context.Context, params InstanceImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a specific image. The image cannot be deleted if it is used by protected
// snapshots.
func (r *InstanceImageService) Delete(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// DeleteAndPoll delete the image and poll for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *InstanceImageService) DeleteAndPoll(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, imageID, body, actionOpts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}

	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	return err
}

// Create a new image from a bootable volume. The volume must be bootable to create
// an image from it.
func (r *InstanceImageService) NewFromVolume(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// NewFromVolumeAndPoll create image from volume and poll for the result
func (r *InstanceImageService) NewFromVolumeAndPoll(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (v *Image, err error) {
	// Exclude WithResponseBodyInto for the action (NewFromVolume returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.NewFromVolume(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams InstanceImageGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	task, err := r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// Retrieve detailed information about a specific image.
func (r *InstanceImageService) Get(ctx context.Context, imageID string, params InstanceImageGetParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Upload an image from a URL. The image can be configured with various properties
// like OS type, architecture, and tags.
func (r *InstanceImageService) Upload(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/downloadimage/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// UploadAndPoll upload image and poll for the task completion. Use the [TaskService.Poll] method if you need to poll
// for all tasks.
func (r *InstanceImageService) UploadAndPoll(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (v *Image, err error) {
	// Exclude WithResponseBodyInto for the action (Upload returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Upload(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams InstanceImageGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	task, err := r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

type Image struct {
	// Image ID
	ID string `json:"id" api:"required"`
	// An image architecture type: aarch64, `x86_64`.
	//
	// Any of "aarch64", "x86_64".
	Architecture ImageArchitecture `json:"architecture" api:"required"`
	// Datetime when the image was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id" api:"required"`
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code" api:"required"`
	// Image description
	Description string `json:"description" api:"required"`
	// Disk format
	DiskFormat   string `json:"disk_format" api:"required"`
	DisplayOrder int64  `json:"display_order" api:"required"`
	// Name of the GPU driver vendor
	GPUDriver string `json:"gpu_driver" api:"required"`
	// Type of the GPU driver
	GPUDriverType string `json:"gpu_driver_type" api:"required"`
	// Version of the installed GPU driver
	GPUDriverVersion string `json:"gpu_driver_version" api:"required"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType ImageHwFirmwareType `json:"hw_firmware_type" api:"required"`
	// A virtual chipset type.
	//
	// Any of "i440", "q35".
	HwMachineType ImageHwMachineType `json:"hw_machine_type" api:"required"`
	// Set to true if the image will be used by bare metal servers.
	IsBaremetal bool `json:"is_baremetal" api:"required"`
	// Minimal boot volume required
	MinDisk int64 `json:"min_disk" api:"required"`
	// Minimal VM RAM required
	MinRam int64 `json:"min_ram" api:"required"`
	// Image display name
	Name string `json:"name" api:"required"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro string `json:"os_distro" api:"required"`
	// The operating system installed on the image.
	//
	// Any of "linux", "windows".
	OsType ImageOsType `json:"os_type" api:"required"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version" api:"required"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus ImagePriceStatus `json:"price_status" api:"required"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Image size in bytes
	Size int64 `json:"size" api:"required"`
	// Whether the image supports SSH key or not
	//
	// Any of "allow", "deny", "required".
	SSHKey ImageSSHKey `json:"ssh_key" api:"required"`
	// Image status, i.e. active
	Status string `json:"status" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags" api:"required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id" api:"required"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Image visibility. Globally visible images are public
	Visibility string `json:"visibility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Architecture     respjson.Field
		CreatedAt        respjson.Field
		CreatorTaskID    respjson.Field
		CurrencyCode     respjson.Field
		Description      respjson.Field
		DiskFormat       respjson.Field
		DisplayOrder     respjson.Field
		GPUDriver        respjson.Field
		GPUDriverType    respjson.Field
		GPUDriverVersion respjson.Field
		HwFirmwareType   respjson.Field
		HwMachineType    respjson.Field
		IsBaremetal      respjson.Field
		MinDisk          respjson.Field
		MinRam           respjson.Field
		Name             respjson.Field
		OsDistro         respjson.Field
		OsType           respjson.Field
		OsVersion        respjson.Field
		PricePerHour     respjson.Field
		PricePerMonth    respjson.Field
		PriceStatus      respjson.Field
		ProjectID        respjson.Field
		Region           respjson.Field
		RegionID         respjson.Field
		Size             respjson.Field
		SSHKey           respjson.Field
		Status           respjson.Field
		Tags             respjson.Field
		TaskID           respjson.Field
		UpdatedAt        respjson.Field
		Visibility       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image architecture type: aarch64, `x86_64`.
type ImageArchitecture string

const (
	ImageArchitectureAarch64 ImageArchitecture = "aarch64"
	ImageArchitectureX86_64  ImageArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type ImageHwFirmwareType string

const (
	ImageHwFirmwareTypeBios ImageHwFirmwareType = "bios"
	ImageHwFirmwareTypeUefi ImageHwFirmwareType = "uefi"
)

// A virtual chipset type.
type ImageHwMachineType string

const (
	ImageHwMachineTypeI440 ImageHwMachineType = "i440"
	ImageHwMachineTypeQ35  ImageHwMachineType = "q35"
)

// The operating system installed on the image.
type ImageOsType string

const (
	ImageOsTypeLinux   ImageOsType = "linux"
	ImageOsTypeWindows ImageOsType = "windows"
)

// Price status for the UI
type ImagePriceStatus string

const (
	ImagePriceStatusError ImagePriceStatus = "error"
	ImagePriceStatusHide  ImagePriceStatus = "hide"
	ImagePriceStatusShow  ImagePriceStatus = "show"
)

// Whether the image supports SSH key or not
type ImageSSHKey string

const (
	ImageSSHKeyAllow    ImageSSHKey = "allow"
	ImageSSHKeyDeny     ImageSSHKey = "deny"
	ImageSSHKeyRequired ImageSSHKey = "required"
)

type ImageList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []Image `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageList) RawJSON() string { return r.JSON.raw }
func (r *ImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceImageUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Set to true if the image will be used by bare metal servers.
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// Image display name
	Name param.Opt[string] `json:"name,omitzero"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUpdateParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// A virtual chipset type.
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageUpdateParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// The operating system installed on the image.
	//
	// Any of "linux", "windows".
	OsType InstanceImageUpdateParamsOsType `json:"os_type,omitzero"`
	// Whether the image supports SSH key or not
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUpdateParamsSSHKey `json:"ssh_key,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r InstanceImageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceImageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of firmware with which to boot the guest.
type InstanceImageUpdateParamsHwFirmwareType string

const (
	InstanceImageUpdateParamsHwFirmwareTypeBios InstanceImageUpdateParamsHwFirmwareType = "bios"
	InstanceImageUpdateParamsHwFirmwareTypeUefi InstanceImageUpdateParamsHwFirmwareType = "uefi"
)

// A virtual chipset type.
type InstanceImageUpdateParamsHwMachineType string

const (
	InstanceImageUpdateParamsHwMachineTypeI440 InstanceImageUpdateParamsHwMachineType = "i440"
	InstanceImageUpdateParamsHwMachineTypeQ35  InstanceImageUpdateParamsHwMachineType = "q35"
)

// The operating system installed on the image.
type InstanceImageUpdateParamsOsType string

const (
	InstanceImageUpdateParamsOsTypeLinux   InstanceImageUpdateParamsOsType = "linux"
	InstanceImageUpdateParamsOsTypeWindows InstanceImageUpdateParamsOsType = "windows"
)

// Whether the image supports SSH key or not
type InstanceImageUpdateParamsSSHKey string

const (
	InstanceImageUpdateParamsSSHKeyAllow    InstanceImageUpdateParamsSSHKey = "allow"
	InstanceImageUpdateParamsSSHKeyDeny     InstanceImageUpdateParamsSSHKey = "deny"
	InstanceImageUpdateParamsSSHKeyRequired InstanceImageUpdateParamsSSHKey = "required"
)

type InstanceImageListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Show price.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Any value to show private images
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	// Image visibility. Globally visible images are public
	//
	// Any of "private", "public", "shared".
	Visibility InstanceImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceImageListParams]'s query parameters as
// `url.Values`.
func (r InstanceImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Image visibility. Globally visible images are public
type InstanceImageListParamsVisibility string

const (
	InstanceImageListParamsVisibilityPrivate InstanceImageListParamsVisibility = "private"
	InstanceImageListParamsVisibilityPublic  InstanceImageListParamsVisibility = "public"
	InstanceImageListParamsVisibilityShared  InstanceImageListParamsVisibility = "shared"
)

type InstanceImageDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstanceImageNewFromVolumeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Image name
	Name string `json:"name" api:"required"`
	// Required if source is volume. Volume id
	VolumeID string `json:"volume_id" api:"required" format:"uuid4"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageNewFromVolumeParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// A virtual chipset type.
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageNewFromVolumeParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// Image CPU architecture type: `aarch64`, `x86_64`
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageNewFromVolumeParamsArchitecture `json:"architecture,omitzero"`
	// The operating system installed on the image.
	//
	// Any of "linux", "windows".
	OsType InstanceImageNewFromVolumeParamsOsType `json:"os_type,omitzero"`
	// Image source
	//
	// Any of "volume".
	Source InstanceImageNewFromVolumeParamsSource `json:"source,omitzero"`
	// Whether the image supports SSH key or not
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageNewFromVolumeParamsSSHKey `json:"ssh_key,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r InstanceImageNewFromVolumeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageNewFromVolumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceImageNewFromVolumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image CPU architecture type: `aarch64`, `x86_64`
type InstanceImageNewFromVolumeParamsArchitecture string

const (
	InstanceImageNewFromVolumeParamsArchitectureAarch64 InstanceImageNewFromVolumeParamsArchitecture = "aarch64"
	InstanceImageNewFromVolumeParamsArchitectureX86_64  InstanceImageNewFromVolumeParamsArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type InstanceImageNewFromVolumeParamsHwFirmwareType string

const (
	InstanceImageNewFromVolumeParamsHwFirmwareTypeBios InstanceImageNewFromVolumeParamsHwFirmwareType = "bios"
	InstanceImageNewFromVolumeParamsHwFirmwareTypeUefi InstanceImageNewFromVolumeParamsHwFirmwareType = "uefi"
)

// A virtual chipset type.
type InstanceImageNewFromVolumeParamsHwMachineType string

const (
	InstanceImageNewFromVolumeParamsHwMachineTypeI440 InstanceImageNewFromVolumeParamsHwMachineType = "i440"
	InstanceImageNewFromVolumeParamsHwMachineTypeQ35  InstanceImageNewFromVolumeParamsHwMachineType = "q35"
)

// The operating system installed on the image.
type InstanceImageNewFromVolumeParamsOsType string

const (
	InstanceImageNewFromVolumeParamsOsTypeLinux   InstanceImageNewFromVolumeParamsOsType = "linux"
	InstanceImageNewFromVolumeParamsOsTypeWindows InstanceImageNewFromVolumeParamsOsType = "windows"
)

// Image source
type InstanceImageNewFromVolumeParamsSource string

const (
	InstanceImageNewFromVolumeParamsSourceVolume InstanceImageNewFromVolumeParamsSource = "volume"
)

// Whether the image supports SSH key or not
type InstanceImageNewFromVolumeParamsSSHKey string

const (
	InstanceImageNewFromVolumeParamsSSHKeyAllow    InstanceImageNewFromVolumeParamsSSHKey = "allow"
	InstanceImageNewFromVolumeParamsSSHKeyDeny     InstanceImageNewFromVolumeParamsSSHKey = "deny"
	InstanceImageNewFromVolumeParamsSSHKeyRequired InstanceImageNewFromVolumeParamsSSHKey = "required"
)

type InstanceImageGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Show price
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceImageGetParams]'s query parameters as `url.Values`.
func (r InstanceImageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceImageUploadParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Image name
	Name string `json:"name" api:"required"`
	// URL
	URL string `json:"url" api:"required" format:"uri"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro param.Opt[string] `json:"os_distro,omitzero"`
	// OS version, i.e. 22.04 (for Ubuntu) or 9.4 for Debian
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// When True, image cannot be deleted unless all volumes, created from it, are
	// deleted.
	CowFormat param.Opt[bool] `json:"cow_format,omitzero"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUploadParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// A virtual chipset type.
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageUploadParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// Image CPU architecture type: `aarch64`, `x86_64`
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageUploadParamsArchitecture `json:"architecture,omitzero"`
	// The operating system installed on the image.
	//
	// Any of "linux", "windows".
	OsType InstanceImageUploadParamsOsType `json:"os_type,omitzero"`
	// Whether the image supports SSH key or not
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUploadParamsSSHKey `json:"ssh_key,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r InstanceImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceImageUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image CPU architecture type: `aarch64`, `x86_64`
type InstanceImageUploadParamsArchitecture string

const (
	InstanceImageUploadParamsArchitectureAarch64 InstanceImageUploadParamsArchitecture = "aarch64"
	InstanceImageUploadParamsArchitectureX86_64  InstanceImageUploadParamsArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type InstanceImageUploadParamsHwFirmwareType string

const (
	InstanceImageUploadParamsHwFirmwareTypeBios InstanceImageUploadParamsHwFirmwareType = "bios"
	InstanceImageUploadParamsHwFirmwareTypeUefi InstanceImageUploadParamsHwFirmwareType = "uefi"
)

// A virtual chipset type.
type InstanceImageUploadParamsHwMachineType string

const (
	InstanceImageUploadParamsHwMachineTypeI440 InstanceImageUploadParamsHwMachineType = "i440"
	InstanceImageUploadParamsHwMachineTypeQ35  InstanceImageUploadParamsHwMachineType = "q35"
)

// The operating system installed on the image.
type InstanceImageUploadParamsOsType string

const (
	InstanceImageUploadParamsOsTypeLinux   InstanceImageUploadParamsOsType = "linux"
	InstanceImageUploadParamsOsTypeWindows InstanceImageUploadParamsOsType = "windows"
)

// Whether the image supports SSH key or not
type InstanceImageUploadParamsSSHKey string

const (
	InstanceImageUploadParamsSSHKeyAllow    InstanceImageUploadParamsSSHKey = "allow"
	InstanceImageUploadParamsSSHKeyDeny     InstanceImageUploadParamsSSHKey = "deny"
	InstanceImageUploadParamsSSHKeyRequired InstanceImageUploadParamsSSHKey = "required"
)
