// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
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
func (r *InstanceImageService) Update(ctx context.Context, imageID string, params InstanceImageUpdateParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
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
func (r *InstanceImageService) List(ctx context.Context, params InstanceImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
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
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
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

// Create image
func (r *InstanceImageService) NewFromVolume(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
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
func (r *InstanceImageService) Get(ctx context.Context, imageID string, params InstanceImageGetParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
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

type InstanceImageUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].patch.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].patch.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/UpdateImageSchema/properties/is_baremetal'
	// "$.components.schemas.UpdateImageSchema.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/name'
	// "$.components.schemas.UpdateImageSchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/hw_firmware_type'
	// "$.components.schemas.UpdateImageSchema.properties.hw_firmware_type"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUpdateParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/hw_machine_type'
	// "$.components.schemas.UpdateImageSchema.properties.hw_machine_type"
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageUpdateParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/metadata'
	// "$.components.schemas.UpdateImageSchema.properties.metadata"
	Metadata any `json:"metadata,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/os_type'
	// "$.components.schemas.UpdateImageSchema.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageUpdateParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/UpdateImageSchema/properties/ssh_key'
	// "$.components.schemas.UpdateImageSchema.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUpdateParamsSSHKey `json:"ssh_key,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UpdateImageSchema/properties/hw_firmware_type'
// "$.components.schemas.UpdateImageSchema.properties.hw_firmware_type"
type InstanceImageUpdateParamsHwFirmwareType string

const (
	InstanceImageUpdateParamsHwFirmwareTypeBios InstanceImageUpdateParamsHwFirmwareType = "bios"
	InstanceImageUpdateParamsHwFirmwareTypeUefi InstanceImageUpdateParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/UpdateImageSchema/properties/hw_machine_type'
// "$.components.schemas.UpdateImageSchema.properties.hw_machine_type"
type InstanceImageUpdateParamsHwMachineType string

const (
	InstanceImageUpdateParamsHwMachineTypeI440 InstanceImageUpdateParamsHwMachineType = "i440"
	InstanceImageUpdateParamsHwMachineTypeQ35  InstanceImageUpdateParamsHwMachineType = "q35"
)

// '#/components/schemas/UpdateImageSchema/properties/os_type'
// "$.components.schemas.UpdateImageSchema.properties.os_type"
type InstanceImageUpdateParamsOsType string

const (
	InstanceImageUpdateParamsOsTypeLinux   InstanceImageUpdateParamsOsType = "linux"
	InstanceImageUpdateParamsOsTypeWindows InstanceImageUpdateParamsOsType = "windows"
)

// '#/components/schemas/UpdateImageSchema/properties/ssh_key'
// "$.components.schemas.UpdateImageSchema.properties.ssh_key"
type InstanceImageUpdateParamsSSHKey string

const (
	InstanceImageUpdateParamsSSHKeyAllow    InstanceImageUpdateParamsSSHKey = "allow"
	InstanceImageUpdateParamsSSHKeyDeny     InstanceImageUpdateParamsSSHKey = "deny"
	InstanceImageUpdateParamsSSHKeyRequired InstanceImageUpdateParamsSSHKey = "required"
)

type InstanceImageListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[3]"
	MetadataK param.Opt[string] `query:"metadata_k,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[4]"
	MetadataKv param.Opt[string] `query:"metadata_kv,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].get.parameters[5]"
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
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
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}']['delete'].parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InstanceImageNewFromVolumeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].post.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}'].post.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ImageCreateSchema/properties/name'
	// "$.components.schemas.ImageCreateSchema.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageCreateSchema/properties/volume_id'
	// "$.components.schemas.ImageCreateSchema.properties.volume_id"
	VolumeID string `json:"volume_id,required"`
	// '#/components/schemas/ImageCreateSchema/properties/is_baremetal'
	// "$.components.schemas.ImageCreateSchema.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/source'
	// "$.components.schemas.ImageCreateSchema.properties.source"
	Source param.Opt[string] `json:"source,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/architecture'
	// "$.components.schemas.ImageCreateSchema.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageNewFromVolumeParamsArchitecture `json:"architecture,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/hw_firmware_type'
	// "$.components.schemas.ImageCreateSchema.properties.hw_firmware_type"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageNewFromVolumeParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/hw_machine_type'
	// "$.components.schemas.ImageCreateSchema.properties.hw_machine_type"
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageNewFromVolumeParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/metadata'
	// "$.components.schemas.ImageCreateSchema.properties.metadata"
	Metadata any `json:"metadata,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/os_type'
	// "$.components.schemas.ImageCreateSchema.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageNewFromVolumeParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/ImageCreateSchema/properties/ssh_key'
	// "$.components.schemas.ImageCreateSchema.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageNewFromVolumeParamsSSHKey `json:"ssh_key,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageNewFromVolumeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageNewFromVolumeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageNewFromVolumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ImageCreateSchema/properties/architecture'
// "$.components.schemas.ImageCreateSchema.properties.architecture"
type InstanceImageNewFromVolumeParamsArchitecture string

const (
	InstanceImageNewFromVolumeParamsArchitectureAarch64 InstanceImageNewFromVolumeParamsArchitecture = "aarch64"
	InstanceImageNewFromVolumeParamsArchitectureX86_64  InstanceImageNewFromVolumeParamsArchitecture = "x86_64"
)

// '#/components/schemas/ImageCreateSchema/properties/hw_firmware_type'
// "$.components.schemas.ImageCreateSchema.properties.hw_firmware_type"
type InstanceImageNewFromVolumeParamsHwFirmwareType string

const (
	InstanceImageNewFromVolumeParamsHwFirmwareTypeBios InstanceImageNewFromVolumeParamsHwFirmwareType = "bios"
	InstanceImageNewFromVolumeParamsHwFirmwareTypeUefi InstanceImageNewFromVolumeParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageCreateSchema/properties/hw_machine_type'
// "$.components.schemas.ImageCreateSchema.properties.hw_machine_type"
type InstanceImageNewFromVolumeParamsHwMachineType string

const (
	InstanceImageNewFromVolumeParamsHwMachineTypeI440 InstanceImageNewFromVolumeParamsHwMachineType = "i440"
	InstanceImageNewFromVolumeParamsHwMachineTypeQ35  InstanceImageNewFromVolumeParamsHwMachineType = "q35"
)

// '#/components/schemas/ImageCreateSchema/properties/os_type'
// "$.components.schemas.ImageCreateSchema.properties.os_type"
type InstanceImageNewFromVolumeParamsOsType string

const (
	InstanceImageNewFromVolumeParamsOsTypeLinux   InstanceImageNewFromVolumeParamsOsType = "linux"
	InstanceImageNewFromVolumeParamsOsTypeWindows InstanceImageNewFromVolumeParamsOsType = "windows"
)

// '#/components/schemas/ImageCreateSchema/properties/ssh_key'
// "$.components.schemas.ImageCreateSchema.properties.ssh_key"
type InstanceImageNewFromVolumeParamsSSHKey string

const (
	InstanceImageNewFromVolumeParamsSSHKeyAllow    InstanceImageNewFromVolumeParamsSSHKey = "allow"
	InstanceImageNewFromVolumeParamsSSHKeyDeny     InstanceImageNewFromVolumeParamsSSHKey = "deny"
	InstanceImageNewFromVolumeParamsSSHKeyRequired InstanceImageNewFromVolumeParamsSSHKey = "required"
)

type InstanceImageGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[3]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[4]"
	MetadataK param.Opt[string] `query:"metadata_k,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[5]"
	MetadataKv param.Opt[string] `query:"metadata_kv,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[6]"
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[7]"
	//
	// Any of "private", "public", "shared".
	Visibility InstanceImageGetParamsVisibility `query:"visibility,omitzero" json:"-"`
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

// '#/paths/%2Fcloud%2Fv1%2Fimages%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bimage_id%7D/get/parameters/7'
// "$.paths['/cloud/v1/images/{project_id}/{region_id}/{image_id}'].get.parameters[7]"
type InstanceImageGetParamsVisibility string

const (
	InstanceImageGetParamsVisibilityPrivate InstanceImageGetParamsVisibility = "private"
	InstanceImageGetParamsVisibilityPublic  InstanceImageGetParamsVisibility = "public"
	InstanceImageGetParamsVisibilityShared  InstanceImageGetParamsVisibility = "shared"
)

type InstanceImageUploadParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fdownloadimage%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/downloadimage/{project_id}/{region_id}'].post.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fdownloadimage%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/downloadimage/{project_id}/{region_id}'].post.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ImageDownloadSchema/properties/name'
	// "$.components.schemas.ImageDownloadSchema.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageDownloadSchema/properties/url'
	// "$.components.schemas.ImageDownloadSchema.properties.url"
	URL string `json:"url,required"`
	// '#/components/schemas/ImageDownloadSchema/properties/is_baremetal'
	// "$.components.schemas.ImageDownloadSchema.properties.is_baremetal"
	IsBaremetal param.Opt[bool] `json:"is_baremetal,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/cow_format'
	// "$.components.schemas.ImageDownloadSchema.properties.cow_format"
	CowFormat param.Opt[bool] `json:"cow_format,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/os_distro'
	// "$.components.schemas.ImageDownloadSchema.properties.os_distro"
	OsDistro param.Opt[string] `json:"os_distro,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/os_version'
	// "$.components.schemas.ImageDownloadSchema.properties.os_version"
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/architecture'
	// "$.components.schemas.ImageDownloadSchema.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture InstanceImageUploadParamsArchitecture `json:"architecture,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/hw_firmware_type'
	// "$.components.schemas.ImageDownloadSchema.properties.hw_firmware_type"
	//
	// Any of "bios", "uefi".
	HwFirmwareType InstanceImageUploadParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/hw_machine_type'
	// "$.components.schemas.ImageDownloadSchema.properties.hw_machine_type"
	//
	// Any of "i440", "q35".
	HwMachineType InstanceImageUploadParamsHwMachineType `json:"hw_machine_type,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/metadata'
	// "$.components.schemas.ImageDownloadSchema.properties.metadata"
	Metadata any `json:"metadata,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/os_type'
	// "$.components.schemas.ImageDownloadSchema.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType InstanceImageUploadParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/ImageDownloadSchema/properties/ssh_key'
	// "$.components.schemas.ImageDownloadSchema.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey InstanceImageUploadParamsSSHKey `json:"ssh_key,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceImageUploadParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ImageDownloadSchema/properties/architecture'
// "$.components.schemas.ImageDownloadSchema.properties.architecture"
type InstanceImageUploadParamsArchitecture string

const (
	InstanceImageUploadParamsArchitectureAarch64 InstanceImageUploadParamsArchitecture = "aarch64"
	InstanceImageUploadParamsArchitectureX86_64  InstanceImageUploadParamsArchitecture = "x86_64"
)

// '#/components/schemas/ImageDownloadSchema/properties/hw_firmware_type'
// "$.components.schemas.ImageDownloadSchema.properties.hw_firmware_type"
type InstanceImageUploadParamsHwFirmwareType string

const (
	InstanceImageUploadParamsHwFirmwareTypeBios InstanceImageUploadParamsHwFirmwareType = "bios"
	InstanceImageUploadParamsHwFirmwareTypeUefi InstanceImageUploadParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageDownloadSchema/properties/hw_machine_type'
// "$.components.schemas.ImageDownloadSchema.properties.hw_machine_type"
type InstanceImageUploadParamsHwMachineType string

const (
	InstanceImageUploadParamsHwMachineTypeI440 InstanceImageUploadParamsHwMachineType = "i440"
	InstanceImageUploadParamsHwMachineTypeQ35  InstanceImageUploadParamsHwMachineType = "q35"
)

// '#/components/schemas/ImageDownloadSchema/properties/os_type'
// "$.components.schemas.ImageDownloadSchema.properties.os_type"
type InstanceImageUploadParamsOsType string

const (
	InstanceImageUploadParamsOsTypeLinux   InstanceImageUploadParamsOsType = "linux"
	InstanceImageUploadParamsOsTypeWindows InstanceImageUploadParamsOsType = "windows"
)

// '#/components/schemas/ImageDownloadSchema/properties/ssh_key'
// "$.components.schemas.ImageDownloadSchema.properties.ssh_key"
type InstanceImageUploadParamsSSHKey string

const (
	InstanceImageUploadParamsSSHKeyAllow    InstanceImageUploadParamsSSHKey = "allow"
	InstanceImageUploadParamsSSHKeyDeny     InstanceImageUploadParamsSSHKey = "deny"
	InstanceImageUploadParamsSSHKeyRequired InstanceImageUploadParamsSSHKey = "required"
)
