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

// BaremetalImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalImageService] method instead.
type BaremetalImageService struct {
	Options []option.RequestOption
}

// NewBaremetalImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalImageService(opts ...option.RequestOption) (r BaremetalImageService) {
	r = BaremetalImageService{}
	r.Options = opts
	return
}

// Retrieve a list of available images for bare metal servers. The list can be
// filtered by visibility, tags, and other parameters. Returned entities may or may
// not be owned by the project.
func (r *BaremetalImageService) List(ctx context.Context, params BaremetalImageListParams, opts ...option.RequestOption) (res *BaremetalImageList, err error) {
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
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type BaremetalImage struct {
	// Image ID
	ID string `json:"id" api:"required"`
	// An image architecture type: aarch64, `x86_64`.
	//
	// Any of "aarch64", "x86_64".
	Architecture BaremetalImageArchitecture `json:"architecture" api:"required"`
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
	HwFirmwareType BaremetalImageHwFirmwareType `json:"hw_firmware_type" api:"required"`
	// A virtual chipset type.
	//
	// Any of "i440", "q35".
	HwMachineType BaremetalImageHwMachineType `json:"hw_machine_type" api:"required"`
	// For bare metal servers this value is always set to true
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
	OsType BaremetalImageOsType `json:"os_type" api:"required"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version" api:"required"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus BaremetalImagePriceStatus `json:"price_status" api:"required"`
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
	SSHKey BaremetalImageSSHKey `json:"ssh_key" api:"required"`
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
func (r BaremetalImage) RawJSON() string { return r.JSON.raw }
func (r *BaremetalImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image architecture type: aarch64, `x86_64`.
type BaremetalImageArchitecture string

const (
	BaremetalImageArchitectureAarch64 BaremetalImageArchitecture = "aarch64"
	BaremetalImageArchitectureX86_64  BaremetalImageArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type BaremetalImageHwFirmwareType string

const (
	BaremetalImageHwFirmwareTypeBios BaremetalImageHwFirmwareType = "bios"
	BaremetalImageHwFirmwareTypeUefi BaremetalImageHwFirmwareType = "uefi"
)

// A virtual chipset type.
type BaremetalImageHwMachineType string

const (
	BaremetalImageHwMachineTypeI440 BaremetalImageHwMachineType = "i440"
	BaremetalImageHwMachineTypeQ35  BaremetalImageHwMachineType = "q35"
)

// The operating system installed on the image.
type BaremetalImageOsType string

const (
	BaremetalImageOsTypeLinux   BaremetalImageOsType = "linux"
	BaremetalImageOsTypeWindows BaremetalImageOsType = "windows"
)

// Price status for the UI
type BaremetalImagePriceStatus string

const (
	BaremetalImagePriceStatusError BaremetalImagePriceStatus = "error"
	BaremetalImagePriceStatusHide  BaremetalImagePriceStatus = "hide"
	BaremetalImagePriceStatusShow  BaremetalImagePriceStatus = "show"
)

// Whether the image supports SSH key or not
type BaremetalImageSSHKey string

const (
	BaremetalImageSSHKeyAllow    BaremetalImageSSHKey = "allow"
	BaremetalImageSSHKeyDeny     BaremetalImageSSHKey = "deny"
	BaremetalImageSSHKeyRequired BaremetalImageSSHKey = "required"
)

type BaremetalImageList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []BaremetalImage `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalImageList) RawJSON() string { return r.JSON.raw }
func (r *BaremetalImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalImageListParams struct {
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
	Visibility BaremetalImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalImageListParams]'s query parameters as
// `url.Values`.
func (r BaremetalImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Image visibility. Globally visible images are public
type BaremetalImageListParamsVisibility string

const (
	BaremetalImageListParamsVisibilityPrivate BaremetalImageListParamsVisibility = "private"
	BaremetalImageListParamsVisibilityPublic  BaremetalImageListParamsVisibility = "public"
	BaremetalImageListParamsVisibilityShared  BaremetalImageListParamsVisibility = "shared"
)
