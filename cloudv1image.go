// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ImageService] method instead.
type CloudV1ImageService struct {
	Options      []option.RequestOption
	Metadata     *CloudV1ImageMetadataService
	MetadataItem *CloudV1ImageMetadataItemService
}

// NewCloudV1ImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1ImageService(opts ...option.RequestOption) (r *CloudV1ImageService) {
	r = &CloudV1ImageService{}
	r.Options = opts
	r.Metadata = NewCloudV1ImageMetadataService(opts...)
	r.MetadataItem = NewCloudV1ImageMetadataItemService(opts...)
	return
}

// Create image
func (r *CloudV1ImageService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1ImageNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/images/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get image
func (r *CloudV1ImageService) Get(ctx context.Context, projectID int64, regionID int64, imageID string, query CloudV1ImageGetParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update image fields
func (r *CloudV1ImageService) Update(ctx context.Context, projectID int64, regionID int64, imageID string, body CloudV1ImageUpdateParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve an available images list. Returned entities owned by the project and
// public OR shared with the client
func (r *CloudV1ImageService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1ImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/images/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the image
func (r *CloudV1ImageService) Delete(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Image schema
type Image struct {
	// Image display name
	Name string `json:"name,required"`
	// Image ID
	ID string `json:"id"`
	// an image architecture type: aarch64, x86_64
	Architecture ImageArchitecture `json:"architecture"`
	// Datetime when the image was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code"`
	// Image description
	Description string `json:"description"`
	// Disk format
	DiskFormat   string `json:"disk_format"`
	DisplayOrder int64  `json:"display_order"`
	// Specifies the type of firmware with which to boot the guest.
	HwFirmwareType ImageHwFirmwareType `json:"hw_firmware_type"`
	// A virtual chipset type.
	HwMachineType ImageHwMachineType `json:"hw_machine_type"`
	// Set to true if image has been internally shared.
	InternallyShared bool `json:"internally_shared"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal bool `json:"is_baremetal"`
	// Metadata
	Metadata interface{} `json:"metadata"`
	// Image metadata
	MetadataDetailed interface{} `json:"metadata_detailed"`
	// Minimal boot volume required
	MinDisk int64 `json:"min_disk"`
	// Minimal VM RAM required
	MinRam int64 `json:"min_ram"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro string `json:"os_distro"`
	// The operating system installed on the image.
	OsType ImageOsType `json:"os_type"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month"`
	// Price status for the UI
	PriceStatus ImagePriceStatus `json:"price_status"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Image size in bytes
	Size int64 `json:"size"`
	// Whether the image supports SSH key or not
	SSHKey ImageSSHKey `json:"ssh_key"`
	// Image status, i.e. active
	Status string `json:"status"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Image visibility. Globally visible images are public
	Visibility string    `json:"visibility"`
	JSON       imageJSON `json:"-"`
}

// imageJSON contains the JSON metadata for the struct [Image]
type imageJSON struct {
	Name             apijson.Field
	ID               apijson.Field
	Architecture     apijson.Field
	CreatedAt        apijson.Field
	CreatorTaskID    apijson.Field
	CurrencyCode     apijson.Field
	Description      apijson.Field
	DiskFormat       apijson.Field
	DisplayOrder     apijson.Field
	HwFirmwareType   apijson.Field
	HwMachineType    apijson.Field
	InternallyShared apijson.Field
	IsBaremetal      apijson.Field
	Metadata         apijson.Field
	MetadataDetailed apijson.Field
	MinDisk          apijson.Field
	MinRam           apijson.Field
	OsDistro         apijson.Field
	OsType           apijson.Field
	OsVersion        apijson.Field
	PricePerHour     apijson.Field
	PricePerMonth    apijson.Field
	PriceStatus      apijson.Field
	ProjectID        apijson.Field
	Region           apijson.Field
	RegionID         apijson.Field
	Size             apijson.Field
	SSHKey           apijson.Field
	Status           apijson.Field
	TaskID           apijson.Field
	UpdatedAt        apijson.Field
	Visibility       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Image) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageJSON) RawJSON() string {
	return r.raw
}

// an image architecture type: aarch64, x86_64
type ImageArchitecture string

const (
	ImageArchitectureAarch64 ImageArchitecture = "aarch64"
	ImageArchitectureX86_64  ImageArchitecture = "x86_64"
)

func (r ImageArchitecture) IsKnown() bool {
	switch r {
	case ImageArchitectureAarch64, ImageArchitectureX86_64:
		return true
	}
	return false
}

// Specifies the type of firmware with which to boot the guest.
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

// A virtual chipset type.
type ImageHwMachineType string

const (
	ImageHwMachineTypeI440 ImageHwMachineType = "i440"
	ImageHwMachineTypeQ35  ImageHwMachineType = "q35"
)

func (r ImageHwMachineType) IsKnown() bool {
	switch r {
	case ImageHwMachineTypeI440, ImageHwMachineTypeQ35:
		return true
	}
	return false
}

// The operating system installed on the image.
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

// Price status for the UI
type ImagePriceStatus string

const (
	ImagePriceStatusError ImagePriceStatus = "error"
	ImagePriceStatusHide  ImagePriceStatus = "hide"
	ImagePriceStatusShow  ImagePriceStatus = "show"
)

func (r ImagePriceStatus) IsKnown() bool {
	switch r {
	case ImagePriceStatusError, ImagePriceStatusHide, ImagePriceStatusShow:
		return true
	}
	return false
}

// Whether the image supports SSH key or not
type ImageSSHKey string

const (
	ImageSSHKeyAllow    ImageSSHKey = "allow"
	ImageSSHKeyDeny     ImageSSHKey = "deny"
	ImageSSHKeyRequired ImageSSHKey = "required"
)

func (r ImageSSHKey) IsKnown() bool {
	switch r {
	case ImageSSHKeyAllow, ImageSSHKeyDeny, ImageSSHKeyRequired:
		return true
	}
	return false
}

type CloudV1ImageNewParams struct {
	// Image schema when it should be created from a volume
	ImageCreate ImageCreateParam `json:"image_create,required"`
}

func (r CloudV1ImageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ImageCreate)
}

type CloudV1ImageGetParams struct {
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
	Visibility param.Field[CloudV1ImageGetParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [CloudV1ImageGetParams]'s query parameters as `url.Values`.
func (r CloudV1ImageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Image visibility. Globally visible images are public
type CloudV1ImageGetParamsVisibility string

const (
	CloudV1ImageGetParamsVisibilityPrivate CloudV1ImageGetParamsVisibility = "private"
	CloudV1ImageGetParamsVisibilityPublic  CloudV1ImageGetParamsVisibility = "public"
	CloudV1ImageGetParamsVisibilityShared  CloudV1ImageGetParamsVisibility = "shared"
)

func (r CloudV1ImageGetParamsVisibility) IsKnown() bool {
	switch r {
	case CloudV1ImageGetParamsVisibilityPrivate, CloudV1ImageGetParamsVisibilityPublic, CloudV1ImageGetParamsVisibilityShared:
		return true
	}
	return false
}

type CloudV1ImageUpdateParams struct {
	// Specifies the type of firmware with which to boot the guest.
	HwFirmwareType param.Field[CloudV1ImageUpdateParamsHwFirmwareType] `json:"hw_firmware_type"`
	// A virtual chipset type.
	HwMachineType param.Field[CloudV1ImageUpdateParamsHwMachineType] `json:"hw_machine_type"`
	// Set to true if the image will be used by bare metal servers. Set to None if you
	// want to delete the property.
	IsBaremetal param.Field[bool] `json:"is_baremetal"`
	// Create one or more metadata items for an image
	Metadata param.Field[interface{}] `json:"metadata"`
	// Image display name
	Name param.Field[string] `json:"name"`
	// The operating system installed on the image.
	OsType param.Field[CloudV1ImageUpdateParamsOsType] `json:"os_type"`
	// Whether the image supports SSH key or not
	SSHKey param.Field[CloudV1ImageUpdateParamsSSHKey] `json:"ssh_key"`
}

func (r CloudV1ImageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of firmware with which to boot the guest.
type CloudV1ImageUpdateParamsHwFirmwareType string

const (
	CloudV1ImageUpdateParamsHwFirmwareTypeBios CloudV1ImageUpdateParamsHwFirmwareType = "bios"
	CloudV1ImageUpdateParamsHwFirmwareTypeUefi CloudV1ImageUpdateParamsHwFirmwareType = "uefi"
)

func (r CloudV1ImageUpdateParamsHwFirmwareType) IsKnown() bool {
	switch r {
	case CloudV1ImageUpdateParamsHwFirmwareTypeBios, CloudV1ImageUpdateParamsHwFirmwareTypeUefi:
		return true
	}
	return false
}

// A virtual chipset type.
type CloudV1ImageUpdateParamsHwMachineType string

const (
	CloudV1ImageUpdateParamsHwMachineTypeI440 CloudV1ImageUpdateParamsHwMachineType = "i440"
	CloudV1ImageUpdateParamsHwMachineTypeQ35  CloudV1ImageUpdateParamsHwMachineType = "q35"
)

func (r CloudV1ImageUpdateParamsHwMachineType) IsKnown() bool {
	switch r {
	case CloudV1ImageUpdateParamsHwMachineTypeI440, CloudV1ImageUpdateParamsHwMachineTypeQ35:
		return true
	}
	return false
}

// The operating system installed on the image.
type CloudV1ImageUpdateParamsOsType string

const (
	CloudV1ImageUpdateParamsOsTypeLinux   CloudV1ImageUpdateParamsOsType = "linux"
	CloudV1ImageUpdateParamsOsTypeWindows CloudV1ImageUpdateParamsOsType = "windows"
)

func (r CloudV1ImageUpdateParamsOsType) IsKnown() bool {
	switch r {
	case CloudV1ImageUpdateParamsOsTypeLinux, CloudV1ImageUpdateParamsOsTypeWindows:
		return true
	}
	return false
}

// Whether the image supports SSH key or not
type CloudV1ImageUpdateParamsSSHKey string

const (
	CloudV1ImageUpdateParamsSSHKeyAllow    CloudV1ImageUpdateParamsSSHKey = "allow"
	CloudV1ImageUpdateParamsSSHKeyDeny     CloudV1ImageUpdateParamsSSHKey = "deny"
	CloudV1ImageUpdateParamsSSHKeyRequired CloudV1ImageUpdateParamsSSHKey = "required"
)

func (r CloudV1ImageUpdateParamsSSHKey) IsKnown() bool {
	switch r {
	case CloudV1ImageUpdateParamsSSHKeyAllow, CloudV1ImageUpdateParamsSSHKeyDeny, CloudV1ImageUpdateParamsSSHKeyRequired:
		return true
	}
	return false
}

type CloudV1ImageListParams struct {
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
	Visibility param.Field[CloudV1ImageListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [CloudV1ImageListParams]'s query parameters as `url.Values`.
func (r CloudV1ImageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Image visibility. Globally visible images are public
type CloudV1ImageListParamsVisibility string

const (
	CloudV1ImageListParamsVisibilityPrivate CloudV1ImageListParamsVisibility = "private"
	CloudV1ImageListParamsVisibilityPublic  CloudV1ImageListParamsVisibility = "public"
	CloudV1ImageListParamsVisibilityShared  CloudV1ImageListParamsVisibility = "shared"
)

func (r CloudV1ImageListParamsVisibility) IsKnown() bool {
	switch r {
	case CloudV1ImageListParamsVisibilityPrivate, CloudV1ImageListParamsVisibilityPublic, CloudV1ImageListParamsVisibilityShared:
		return true
	}
	return false
}
