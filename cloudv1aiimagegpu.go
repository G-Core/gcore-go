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

// CloudV1AIImageGPUService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIImageGPUService] method instead.
type CloudV1AIImageGPUService struct {
	Options []option.RequestOption
}

// NewCloudV1AIImageGPUService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1AIImageGPUService(opts ...option.RequestOption) (r *CloudV1AIImageGPUService) {
	r = &CloudV1AIImageGPUService{}
	r.Options = opts
	return
}

// Deprecated. Retrieve available AI GPU images.
func (r *CloudV1AIImageGPUService) List(ctx context.Context, projectID string, regionID string, query CloudV1AIImageGPUListParams, opts ...option.RequestOption) (res *CloudV1AIImageGPUListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/images/gpu/%s/%s", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV1AIImageGPUListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []CloudV1AIImageGPUListResponseResult `json:"results"`
	JSON    cloudV1AIImageGPUListResponseJSON     `json:"-"`
}

// cloudV1AIImageGPUListResponseJSON contains the JSON metadata for the struct
// [CloudV1AIImageGPUListResponse]
type cloudV1AIImageGPUListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1AIImageGPUListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIImageGPUListResponseJSON) RawJSON() string {
	return r.raw
}

// Poplar Image schema
type CloudV1AIImageGPUListResponseResult struct {
	// Image display name
	Name string `json:"name,required"`
	// Image ID
	ID string `json:"id"`
	// an image architecture type: aarch64, x86_64
	Architecture CloudV1AIImageGPUListResponseResultsArchitecture `json:"architecture"`
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
	HwFirmwareType CloudV1AIImageGPUListResponseResultsHwFirmwareType `json:"hw_firmware_type"`
	// A virtual chipset type.
	HwMachineType CloudV1AIImageGPUListResponseResultsHwMachineType `json:"hw_machine_type"`
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
	OsType CloudV1AIImageGPUListResponseResultsOsType `json:"os_type"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month"`
	// Price status for the UI
	PriceStatus CloudV1AIImageGPUListResponseResultsPriceStatus `json:"price_status"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Poplar SDK version
	SDKVersion string `json:"sdk_version"`
	// Image size in bytes
	Size int64 `json:"size"`
	// Whether the image supports SSH key or not
	SSHKey CloudV1AIImageGPUListResponseResultsSSHKey `json:"ssh_key"`
	// Image status, i.e. active
	Status string `json:"status"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// vIPU controller version
	VipuVersion string `json:"vipu_version"`
	// Image visibility. Globally visible images are public
	Visibility string                                  `json:"visibility"`
	JSON       cloudV1AIImageGPUListResponseResultJSON `json:"-"`
}

// cloudV1AIImageGPUListResponseResultJSON contains the JSON metadata for the
// struct [CloudV1AIImageGPUListResponseResult]
type cloudV1AIImageGPUListResponseResultJSON struct {
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
	SDKVersion       apijson.Field
	Size             apijson.Field
	SSHKey           apijson.Field
	Status           apijson.Field
	TaskID           apijson.Field
	UpdatedAt        apijson.Field
	VipuVersion      apijson.Field
	Visibility       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudV1AIImageGPUListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIImageGPUListResponseResultJSON) RawJSON() string {
	return r.raw
}

// an image architecture type: aarch64, x86_64
type CloudV1AIImageGPUListResponseResultsArchitecture string

const (
	CloudV1AIImageGPUListResponseResultsArchitectureAarch64 CloudV1AIImageGPUListResponseResultsArchitecture = "aarch64"
	CloudV1AIImageGPUListResponseResultsArchitectureX86_64  CloudV1AIImageGPUListResponseResultsArchitecture = "x86_64"
)

func (r CloudV1AIImageGPUListResponseResultsArchitecture) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsArchitectureAarch64, CloudV1AIImageGPUListResponseResultsArchitectureX86_64:
		return true
	}
	return false
}

// Specifies the type of firmware with which to boot the guest.
type CloudV1AIImageGPUListResponseResultsHwFirmwareType string

const (
	CloudV1AIImageGPUListResponseResultsHwFirmwareTypeBios CloudV1AIImageGPUListResponseResultsHwFirmwareType = "bios"
	CloudV1AIImageGPUListResponseResultsHwFirmwareTypeUefi CloudV1AIImageGPUListResponseResultsHwFirmwareType = "uefi"
)

func (r CloudV1AIImageGPUListResponseResultsHwFirmwareType) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsHwFirmwareTypeBios, CloudV1AIImageGPUListResponseResultsHwFirmwareTypeUefi:
		return true
	}
	return false
}

// A virtual chipset type.
type CloudV1AIImageGPUListResponseResultsHwMachineType string

const (
	CloudV1AIImageGPUListResponseResultsHwMachineTypeI440 CloudV1AIImageGPUListResponseResultsHwMachineType = "i440"
	CloudV1AIImageGPUListResponseResultsHwMachineTypeQ35  CloudV1AIImageGPUListResponseResultsHwMachineType = "q35"
)

func (r CloudV1AIImageGPUListResponseResultsHwMachineType) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsHwMachineTypeI440, CloudV1AIImageGPUListResponseResultsHwMachineTypeQ35:
		return true
	}
	return false
}

// The operating system installed on the image.
type CloudV1AIImageGPUListResponseResultsOsType string

const (
	CloudV1AIImageGPUListResponseResultsOsTypeLinux   CloudV1AIImageGPUListResponseResultsOsType = "linux"
	CloudV1AIImageGPUListResponseResultsOsTypeWindows CloudV1AIImageGPUListResponseResultsOsType = "windows"
)

func (r CloudV1AIImageGPUListResponseResultsOsType) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsOsTypeLinux, CloudV1AIImageGPUListResponseResultsOsTypeWindows:
		return true
	}
	return false
}

// Price status for the UI
type CloudV1AIImageGPUListResponseResultsPriceStatus string

const (
	CloudV1AIImageGPUListResponseResultsPriceStatusError CloudV1AIImageGPUListResponseResultsPriceStatus = "error"
	CloudV1AIImageGPUListResponseResultsPriceStatusHide  CloudV1AIImageGPUListResponseResultsPriceStatus = "hide"
	CloudV1AIImageGPUListResponseResultsPriceStatusShow  CloudV1AIImageGPUListResponseResultsPriceStatus = "show"
)

func (r CloudV1AIImageGPUListResponseResultsPriceStatus) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsPriceStatusError, CloudV1AIImageGPUListResponseResultsPriceStatusHide, CloudV1AIImageGPUListResponseResultsPriceStatusShow:
		return true
	}
	return false
}

// Whether the image supports SSH key or not
type CloudV1AIImageGPUListResponseResultsSSHKey string

const (
	CloudV1AIImageGPUListResponseResultsSSHKeyAllow    CloudV1AIImageGPUListResponseResultsSSHKey = "allow"
	CloudV1AIImageGPUListResponseResultsSSHKeyDeny     CloudV1AIImageGPUListResponseResultsSSHKey = "deny"
	CloudV1AIImageGPUListResponseResultsSSHKeyRequired CloudV1AIImageGPUListResponseResultsSSHKey = "required"
)

func (r CloudV1AIImageGPUListResponseResultsSSHKey) IsKnown() bool {
	switch r {
	case CloudV1AIImageGPUListResponseResultsSSHKeyAllow, CloudV1AIImageGPUListResponseResultsSSHKeyDeny, CloudV1AIImageGPUListResponseResultsSSHKeyRequired:
		return true
	}
	return false
}

type CloudV1AIImageGPUListParams struct {
	// Any value to show private images
	Private param.Field[string] `query:"private"`
	// Search only public, shared or private images
	Visibility param.Field[string] `query:"visibility"`
}

// URLQuery serializes [CloudV1AIImageGPUListParams]'s query parameters as
// `url.Values`.
func (r CloudV1AIImageGPUListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
