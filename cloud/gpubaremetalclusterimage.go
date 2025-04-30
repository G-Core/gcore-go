// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// GPUBaremetalClusterImageService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterImageService] method instead.
type GPUBaremetalClusterImageService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterImageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterImageService(opts ...option.RequestOption) (r GPUBaremetalClusterImageService) {
	r = GPUBaremetalClusterImageService{}
	r.Options = opts
	return
}

// List bare metal GPU images
func (r *GPUBaremetalClusterImageService) List(ctx context.Context, query GPUBaremetalClusterImageListParams, opts ...option.RequestOption) (res *GPUImageList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete bare metal GPU image by ID
func (r *GPUBaremetalClusterImageService) Delete(ctx context.Context, imageID string, body GPUBaremetalClusterImageDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images/%s", body.ProjectID.Value, body.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get bare metal GPU image by ID
func (r *GPUBaremetalClusterImageService) Get(ctx context.Context, imageID string, query GPUBaremetalClusterImageGetParams, opts ...option.RequestOption) (res *GPUImage, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images/%s", query.ProjectID.Value, query.RegionID.Value, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload new bare metal GPU image
func (r *GPUBaremetalClusterImageService) Upload(ctx context.Context, params GPUBaremetalClusterImageUploadParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GPUBaremetalClusterImageListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages/get/parameters/0/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages/get/parameters/1/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterImageListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterImageDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages%2F%7Bimage_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images/{image_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages%2F%7Bimage_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images/{image_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterImageDeleteParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterImageGetParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages%2F%7Bimage_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images/{image_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages%2F%7Bimage_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images/{image_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterImageGetParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterImageUploadParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages/post/parameters/0/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Fgpu%2Fbaremetal%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fimages/post/parameters/1/schema'
	// "$.paths['/cloud/v3/gpu/baremetal/{project_id}/{region_id}/images'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/name'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/url/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.url.anyOf[0]"
	URL string `json:"url,required" format:"uri"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/os_distro/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.os_distro.anyOf[0]"
	OsDistro param.Opt[string] `json:"os_distro,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/os_version/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.os_version.anyOf[0]"
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/cow_format'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.cow_format"
	CowFormat param.Opt[bool] `json:"cow_format,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/architecture/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.architecture.anyOf[0]"
	//
	// Any of "aarch64", "x86_64".
	Architecture GPUBaremetalClusterImageUploadParamsArchitecture `json:"architecture,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/hw_firmware_type/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.hw_firmware_type.anyOf[0]"
	//
	// Any of "bios", "uefi".
	HwFirmwareType GPUBaremetalClusterImageUploadParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/os_type/anyOf/0'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.os_type.anyOf[0]"
	//
	// Any of "linux", "windows".
	OsType GPUBaremetalClusterImageUploadParamsOsType `json:"os_type,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/ssh_key'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.ssh_key"
	//
	// Any of "allow", "deny", "required".
	SSHKey GPUBaremetalClusterImageUploadParamsSSHKey `json:"ssh_key,omitzero"`
	// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/tags'
	// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterImageUploadParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r GPUBaremetalClusterImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/architecture/anyOf/0'
// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.architecture.anyOf[0]"
type GPUBaremetalClusterImageUploadParamsArchitecture string

const (
	GPUBaremetalClusterImageUploadParamsArchitectureAarch64 GPUBaremetalClusterImageUploadParamsArchitecture = "aarch64"
	GPUBaremetalClusterImageUploadParamsArchitectureX86_64  GPUBaremetalClusterImageUploadParamsArchitecture = "x86_64"
)

// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.hw_firmware_type.anyOf[0]"
type GPUBaremetalClusterImageUploadParamsHwFirmwareType string

const (
	GPUBaremetalClusterImageUploadParamsHwFirmwareTypeBios GPUBaremetalClusterImageUploadParamsHwFirmwareType = "bios"
	GPUBaremetalClusterImageUploadParamsHwFirmwareTypeUefi GPUBaremetalClusterImageUploadParamsHwFirmwareType = "uefi"
)

// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/os_type/anyOf/0'
// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.os_type.anyOf[0]"
type GPUBaremetalClusterImageUploadParamsOsType string

const (
	GPUBaremetalClusterImageUploadParamsOsTypeLinux   GPUBaremetalClusterImageUploadParamsOsType = "linux"
	GPUBaremetalClusterImageUploadParamsOsTypeWindows GPUBaremetalClusterImageUploadParamsOsType = "windows"
)

// '#/components/schemas/UploadBaremetalGpuImageSerializer/properties/ssh_key'
// "$.components.schemas.UploadBaremetalGpuImageSerializer.properties.ssh_key"
type GPUBaremetalClusterImageUploadParamsSSHKey string

const (
	GPUBaremetalClusterImageUploadParamsSSHKeyAllow    GPUBaremetalClusterImageUploadParamsSSHKey = "allow"
	GPUBaremetalClusterImageUploadParamsSSHKeyDeny     GPUBaremetalClusterImageUploadParamsSSHKey = "deny"
	GPUBaremetalClusterImageUploadParamsSSHKeyRequired GPUBaremetalClusterImageUploadParamsSSHKey = "required"
)
