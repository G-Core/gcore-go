// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1BmimageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1BmimageService] method instead.
type CloudV1BmimageService struct {
	Options []option.RequestOption
}

// NewCloudV1BmimageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1BmimageService(opts ...option.RequestOption) (r *CloudV1BmimageService) {
	r = &CloudV1BmimageService{}
	r.Options = opts
	return
}

// Create image
func (r *CloudV1BmimageService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1BmimageNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the available images list for bare metal servers. Returned entities may
// or may not be owned by the project
func (r *CloudV1BmimageService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1BmimageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Image schema when it should be created from a volume
type ImageCreateParam struct {
	// Image name
	Name param.Field[string] `json:"name,required"`
	// Required if source is volume. Volume id
	VolumeID param.Field[string] `json:"volume_id,required"`
	// an image architecture type: aarch64, x86_64
	Architecture param.Field[ImageCreateArchitecture] `json:"architecture"`
	// Specifies the type of firmware with which to boot the guest.
	HwFirmwareType param.Field[ImageCreateHwFirmwareType] `json:"hw_firmware_type"`
	// A virtual chipset type.
	HwMachineType param.Field[ImageCreateHwMachineType] `json:"hw_machine_type"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal param.Field[bool] `json:"is_baremetal"`
	// Create one or more metadata items for an image
	Metadata param.Field[interface{}] `json:"metadata"`
	// The operating system installed on the image.
	OsType param.Field[ImageCreateOsType] `json:"os_type"`
	// Image source
	Source param.Field[string] `json:"source"`
	// Permission to use a ssh key in instances
	SSHKey param.Field[ImageCreateSSHKey] `json:"ssh_key"`
}

func (r ImageCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// an image architecture type: aarch64, x86_64
type ImageCreateArchitecture string

const (
	ImageCreateArchitectureAarch64 ImageCreateArchitecture = "aarch64"
	ImageCreateArchitectureX86_64  ImageCreateArchitecture = "x86_64"
)

func (r ImageCreateArchitecture) IsKnown() bool {
	switch r {
	case ImageCreateArchitectureAarch64, ImageCreateArchitectureX86_64:
		return true
	}
	return false
}

// Specifies the type of firmware with which to boot the guest.
type ImageCreateHwFirmwareType string

const (
	ImageCreateHwFirmwareTypeBios ImageCreateHwFirmwareType = "bios"
	ImageCreateHwFirmwareTypeUefi ImageCreateHwFirmwareType = "uefi"
)

func (r ImageCreateHwFirmwareType) IsKnown() bool {
	switch r {
	case ImageCreateHwFirmwareTypeBios, ImageCreateHwFirmwareTypeUefi:
		return true
	}
	return false
}

// A virtual chipset type.
type ImageCreateHwMachineType string

const (
	ImageCreateHwMachineTypeI440 ImageCreateHwMachineType = "i440"
	ImageCreateHwMachineTypeQ35  ImageCreateHwMachineType = "q35"
)

func (r ImageCreateHwMachineType) IsKnown() bool {
	switch r {
	case ImageCreateHwMachineTypeI440, ImageCreateHwMachineTypeQ35:
		return true
	}
	return false
}

// The operating system installed on the image.
type ImageCreateOsType string

const (
	ImageCreateOsTypeLinux   ImageCreateOsType = "linux"
	ImageCreateOsTypeWindows ImageCreateOsType = "windows"
)

func (r ImageCreateOsType) IsKnown() bool {
	switch r {
	case ImageCreateOsTypeLinux, ImageCreateOsTypeWindows:
		return true
	}
	return false
}

// Permission to use a ssh key in instances
type ImageCreateSSHKey string

const (
	ImageCreateSSHKeyAllow    ImageCreateSSHKey = "allow"
	ImageCreateSSHKeyDeny     ImageCreateSSHKey = "deny"
	ImageCreateSSHKeyRequired ImageCreateSSHKey = "required"
)

func (r ImageCreateSSHKey) IsKnown() bool {
	switch r {
	case ImageCreateSSHKeyAllow, ImageCreateSSHKeyDeny, ImageCreateSSHKeyRequired:
		return true
	}
	return false
}

type ImageList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []Image       `json:"results"`
	JSON    imageListJSON `json:"-"`
}

// imageListJSON contains the JSON metadata for the struct [ImageList]
type imageListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageListJSON) RawJSON() string {
	return r.raw
}

type CloudV1BmimageNewParams struct {
	// Image schema when it should be created from a volume
	ImageCreate ImageCreateParam `json:"image_create,required"`
}

func (r CloudV1BmimageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ImageCreate)
}

type CloudV1BmimageListParams struct {
	// Show price
	IncludePrices param.Field[bool] `query:"include_prices"`
	// Filter by metadata keys. Must be a valid JSON string. ' curl -G --data-urlencode
	// 'metadata_k=["key1", "key2"]' --url 'http://localhost:1111/v1/images/1/1'
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. 'curl -G
	// --data-urlencode 'metadata_kv={"key": "value"}' --url
	// 'http://localhost:1111/v1/images/1/1'"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Any value to show private images
	Private param.Field[string] `query:"private"`
	// Image visibility. Globally visible images are public
	Visibility param.Field[CloudV1BmimageListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [CloudV1BmimageListParams]'s query parameters as
// `url.Values`.
func (r CloudV1BmimageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Image visibility. Globally visible images are public
type CloudV1BmimageListParamsVisibility string

const (
	CloudV1BmimageListParamsVisibilityPrivate CloudV1BmimageListParamsVisibility = "private"
	CloudV1BmimageListParamsVisibilityPublic  CloudV1BmimageListParamsVisibility = "public"
	CloudV1BmimageListParamsVisibilityShared  CloudV1BmimageListParamsVisibility = "shared"
)

func (r CloudV1BmimageListParamsVisibility) IsKnown() bool {
	switch r {
	case CloudV1BmimageListParamsVisibilityPrivate, CloudV1BmimageListParamsVisibilityPublic, CloudV1BmimageListParamsVisibilityShared:
		return true
	}
	return false
}
