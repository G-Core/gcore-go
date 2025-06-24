// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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

// Retrieve a list of available images in the project and region. The list can be
// filtered by visibility, tags, and other parameters. Returned entities are owned
// by the project or are public/shared with the client.
func (r *InstanceImageService) List(ctx context.Context, params InstanceImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a specific image. The image cannot be deleted if it is used by protected
// snapshots.
func (r *InstanceImageService) Delete(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
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

// DeleteAndPoll delete the image and poll for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *InstanceImageService) DeleteAndPoll(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, imageID, body, opts...)
	if err != nil {
		return err
	}

	opts = append(r.Options[:], opts...)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}

	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Create a new image from a bootable volume. The volume must be bootable to create
// an image from it.
func (r *InstanceImageService) NewFromVolume(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewFromVolumeAndPoll create image from volume and poll for the result
func (r *InstanceImageService) NewFromVolumeAndPoll(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (v *Image, err error) {
	resource, err := r.NewFromVolume(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
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
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Retrieve detailed information about a specific image.
func (r *InstanceImageService) Get(ctx context.Context, imageID string, params InstanceImageGetParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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

// Upload an image from a URL. The image can be configured with various properties
// like OS type, architecture, and tags.
func (r *InstanceImageService) Upload(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/downloadimage/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UploadAndPoll upload image and poll for the task completion. Use the [TaskService.Poll] method if you need to poll
// for all tasks.
func (r *InstanceImageService) UploadAndPoll(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (v *Image, err error) {
	resource, err := r.Upload(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
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
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

type InstanceImageUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
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
	// Any of "pc", "q35".
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
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
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
	InstanceImageUpdateParamsHwMachineTypePc  InstanceImageUpdateParamsHwMachineType = "pc"
	InstanceImageUpdateParamsHwMachineTypeQ35 InstanceImageUpdateParamsHwMachineType = "q35"
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Show price
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Any value to show private images
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by tag keys.
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceImageNewFromVolumeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Image name
	Name string `json:"name,required"`
	// Required if source is volume. Volume id
	VolumeID string `json:"volume_id,required" format:"uuid4"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageNewFromVolumeParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// A virtual chipset type.
	//
	// Any of "pc", "q35".
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
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
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
	InstanceImageNewFromVolumeParamsHwMachineTypePc  InstanceImageNewFromVolumeParamsHwMachineType = "pc"
	InstanceImageNewFromVolumeParamsHwMachineTypeQ35 InstanceImageNewFromVolumeParamsHwMachineType = "q35"
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Image name
	Name string `json:"name,required"`
	// URL
	URL string `json:"url,required" format:"uri"`
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
	// Any of "pc", "q35".
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
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
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
	InstanceImageUploadParamsHwMachineTypePc  InstanceImageUploadParamsHwMachineType = "pc"
	InstanceImageUploadParamsHwMachineTypeQ35 InstanceImageUploadParamsHwMachineType = "q35"
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
