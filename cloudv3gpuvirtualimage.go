// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3GPUVirtualImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3GPUVirtualImageService] method instead.
type CloudV3GPUVirtualImageService struct {
	Options []option.RequestOption
}

// NewCloudV3GPUVirtualImageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3GPUVirtualImageService(opts ...option.RequestOption) (r *CloudV3GPUVirtualImageService) {
	r = &CloudV3GPUVirtualImageService{}
	r.Options = opts
	return
}

// Get virtual GPU image by ID
func (r *CloudV3GPUVirtualImageService) Get(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *GPUImage, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/images/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List virtual GPU images
func (r *CloudV3GPUVirtualImageService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *ListGPUImage, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/images", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete virtual GPU image by ID
func (r *CloudV3GPUVirtualImageService) Delete(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/images/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload new virtual GPU image
func (r *CloudV3GPUVirtualImageService) Upload(ctx context.Context, projectID int64, regionID int64, body CloudV3GPUVirtualImageUploadParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/images", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV3GPUVirtualImageUploadParams struct {
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

func (r CloudV3GPUVirtualImageUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
