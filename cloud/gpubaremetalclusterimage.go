// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
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
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GPUBaremetalClusterImageListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterImageDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterImageGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterImageUploadParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Image name
	Name string `json:"name,required"`
	// Image URL
	URL string `json:"url,required" format:"uri"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro param.Opt[string] `json:"os_distro,omitzero"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// When True, image cannot be deleted unless all volumes, created from it, are
	// deleted.
	CowFormat param.Opt[bool] `json:"cow_format,omitzero"`
	// Image architecture type: aarch64, `x86_64`
	//
	// Any of "aarch64", "x86_64".
	Architecture GPUBaremetalClusterImageUploadParamsArchitecture `json:"architecture,omitzero"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType GPUBaremetalClusterImageUploadParamsHwFirmwareType `json:"hw_firmware_type,omitzero"`
	// The operating system installed on the image. Linux by default
	//
	// Any of "linux", "windows".
	OsType GPUBaremetalClusterImageUploadParamsOsType `json:"os_type,omitzero"`
	// Permission to use a ssh key in instances
	//
	// Any of "allow", "deny", "required".
	SSHKey GPUBaremetalClusterImageUploadParamsSSHKey `json:"ssh_key,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterImageUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image architecture type: aarch64, `x86_64`
type GPUBaremetalClusterImageUploadParamsArchitecture string

const (
	GPUBaremetalClusterImageUploadParamsArchitectureAarch64 GPUBaremetalClusterImageUploadParamsArchitecture = "aarch64"
	GPUBaremetalClusterImageUploadParamsArchitectureX86_64  GPUBaremetalClusterImageUploadParamsArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type GPUBaremetalClusterImageUploadParamsHwFirmwareType string

const (
	GPUBaremetalClusterImageUploadParamsHwFirmwareTypeBios GPUBaremetalClusterImageUploadParamsHwFirmwareType = "bios"
	GPUBaremetalClusterImageUploadParamsHwFirmwareTypeUefi GPUBaremetalClusterImageUploadParamsHwFirmwareType = "uefi"
)

// The operating system installed on the image. Linux by default
type GPUBaremetalClusterImageUploadParamsOsType string

const (
	GPUBaremetalClusterImageUploadParamsOsTypeLinux   GPUBaremetalClusterImageUploadParamsOsType = "linux"
	GPUBaremetalClusterImageUploadParamsOsTypeWindows GPUBaremetalClusterImageUploadParamsOsType = "windows"
)

// Permission to use a ssh key in instances
type GPUBaremetalClusterImageUploadParamsSSHKey string

const (
	GPUBaremetalClusterImageUploadParamsSSHKeyAllow    GPUBaremetalClusterImageUploadParamsSSHKey = "allow"
	GPUBaremetalClusterImageUploadParamsSSHKeyDeny     GPUBaremetalClusterImageUploadParamsSSHKey = "deny"
	GPUBaremetalClusterImageUploadParamsSSHKeyRequired GPUBaremetalClusterImageUploadParamsSSHKey = "required"
)
