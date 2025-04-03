// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3GPUBaremetalImageService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3GPUBaremetalImageService] method instead.
type CloudV3GPUBaremetalImageService struct {
	Options []option.RequestOption
}

// NewCloudV3GPUBaremetalImageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV3GPUBaremetalImageService(opts ...option.RequestOption) (r *CloudV3GPUBaremetalImageService) {
	r = &CloudV3GPUBaremetalImageService{}
	r.Options = opts
	return
}

// Get bare metal GPU image by ID
func (r *CloudV3GPUBaremetalImageService) Get(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *GPUImage, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List bare metal GPU images
func (r *CloudV3GPUBaremetalImageService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *ListGPUImage, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete bare metal GPU image by ID
func (r *CloudV3GPUBaremetalImageService) Delete(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload new bare metal GPU image
func (r *CloudV3GPUBaremetalImageService) Upload(ctx context.Context, projectID int64, regionID int64, body CloudV3GPUBaremetalImageUploadParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/images", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GPUImage struct {
	// Image ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the image was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Minimal boot volume required
	MinDisk int64 `json:"min_disk,required"`
	// Minimal VM RAM required
	MinRam int64 `json:"min_ram,required"`
	// Image name
	Name string `json:"name,required"`
	// Image status
	Status string `json:"status,required"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Image visibility. Globally visible images are public
	Visibility string `json:"visibility,required"`
	// Image architecture type
	Architecture string `json:"architecture,nullable"`
	// Image metadata dictionary
	Metadata interface{} `json:"metadata"`
	// OS Distribution
	OsDistro string `json:"os_distro,nullable"`
	// The operating system installed on the image
	OsType string `json:"os_type,nullable"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version,nullable"`
	// Image size in bytes.
	Size int64 `json:"size"`
	// Whether the image supports SSH key or not
	SSHKey string `json:"ssh_key,nullable"`
	// Active task id. If None, then the action has been performed immediately in the
	// request itself.
	TaskID string       `json:"task_id,nullable"`
	JSON   gpuImageJSON `json:"-"`
}

// gpuImageJSON contains the JSON metadata for the struct [GPUImage]
type gpuImageJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	MinDisk      apijson.Field
	MinRam       apijson.Field
	Name         apijson.Field
	Status       apijson.Field
	UpdatedAt    apijson.Field
	Visibility   apijson.Field
	Architecture apijson.Field
	Metadata     apijson.Field
	OsDistro     apijson.Field
	OsType       apijson.Field
	OsVersion    apijson.Field
	Size         apijson.Field
	SSHKey       apijson.Field
	TaskID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GPUImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gpuImageJSON) RawJSON() string {
	return r.raw
}

type ImageArchitectureType string

const (
	ImageArchitectureTypeAarch64 ImageArchitectureType = "aarch64"
	ImageArchitectureTypeX86_64  ImageArchitectureType = "x86_64"
)

func (r ImageArchitectureType) IsKnown() bool {
	switch r {
	case ImageArchitectureTypeAarch64, ImageArchitectureTypeX86_64:
		return true
	}
	return false
}

type ImageHwFirmwareType string

const (
	ImageHwFirmwareTypeBios ImageHwFirmwareType = "bios"
	ImageHwFirmwareTypeUefi ImageHwFirmwareType = "uefi"
)

func (r ImageHwFirmwareType) IsKnown() bool {
	switch r {
	case ImageHwFirmwareTypeBios, ImageHwFirmwareTypeUefi:
		return true
	}
	return false
}

type ImageOsType string

const (
	ImageOsTypeLinux   ImageOsType = "linux"
	ImageOsTypeWindows ImageOsType = "windows"
)

func (r ImageOsType) IsKnown() bool {
	switch r {
	case ImageOsTypeLinux, ImageOsTypeWindows:
		return true
	}
	return false
}

type ListGPUImage struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUImage       `json:"results,required"`
	JSON    listGPUImageJSON `json:"-"`
}

// listGPUImageJSON contains the JSON metadata for the struct [ListGPUImage]
type listGPUImageJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListGPUImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGPUImageJSON) RawJSON() string {
	return r.raw
}

type SSHKey string

const (
	SSHKeyAllow    SSHKey = "allow"
	SSHKeyDeny     SSHKey = "deny"
	SSHKeyRequired SSHKey = "required"
)

func (r SSHKey) IsKnown() bool {
	switch r {
	case SSHKeyAllow, SSHKeyDeny, SSHKeyRequired:
		return true
	}
	return false
}

type CloudV3GPUBaremetalImageUploadParams struct {
	// Image name
	Name param.Field[string] `json:"name,required"`
	// Image URL
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Image architecture type: aarch64, x86_64
	Architecture param.Field[ImageArchitectureType] `json:"architecture"`
	// When True, image cannot be deleted unless all volumes, created from it, are
	// deleted.
	CowFormat param.Field[bool] `json:"cow_format"`
	// Specifies the type of firmware with which to boot the guest.
	HwFirmwareType param.Field[ImageHwFirmwareType] `json:"hw_firmware_type"`
	// Create one or more metadata items for a cluster
	Metadata param.Field[interface{}] `json:"metadata"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro param.Field[string] `json:"os_distro"`
	// The operating system installed on the image. Linux by default
	OsType param.Field[ImageOsType] `json:"os_type"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion param.Field[string] `json:"os_version"`
	// Permission to use a ssh key in instances
	SSHKey param.Field[SSHKey] `json:"ssh_key"`
}

func (r CloudV3GPUBaremetalImageUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
