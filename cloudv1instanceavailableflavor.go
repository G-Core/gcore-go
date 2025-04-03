// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1InstanceAvailableFlavorService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1InstanceAvailableFlavorService] method instead.
type CloudV1InstanceAvailableFlavorService struct {
	Options []option.RequestOption
}

// NewCloudV1InstanceAvailableFlavorService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1InstanceAvailableFlavorService(opts ...option.RequestOption) (r *CloudV1InstanceAvailableFlavorService) {
	r = &CloudV1InstanceAvailableFlavorService{}
	r.Options = opts
	return
}

// Get flavors available for a potential instance
func (r *CloudV1InstanceAvailableFlavorService) GetAvailableFlavors(ctx context.Context, projectID int64, regionID int64, params CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams, opts ...option.RequestOption) (res *FlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/available_flavors", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get flavors available for the instance to resize into
func (r *CloudV1InstanceAvailableFlavorService) GetFlavorsToResize(ctx context.Context, projectID int64, regionID int64, instanceID string, query CloudV1InstanceAvailableFlavorGetFlavorsToResizeParams, opts ...option.RequestOption) (res *FlavorList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/available_flavors", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams struct {
	// Volumes details. Non-important info such as names may be omitted.
	Volumes param.Field[[]CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolume] `json:"volumes,required"`
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

func (r CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams]'s
// query parameters as `url.Values`.
func (r CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// CreateInstanceVolume schema
type CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolume struct {
	// Volume source
	Source param.Field[CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource] `json:"source,required"`
	// App template ID. Mandatory if volume is created from marketplace template
	ApptemplateID param.Field[string] `json:"apptemplate_id"`
	// Block device attachment tag (not exposed in the metadata)
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// 0 should be set for primary boot device Unique positive values for other
	// bootable devices.Negative - boot prohibited
	BootIndex param.Field[int64] `json:"boot_index"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination param.Field[bool] `json:"delete_on_termination"`
	// Image ID. Mandatory if volume is created from image
	ImageID param.Field[string] `json:"image_id"`
	// Create one or more metadata items for a volume
	Metadata param.Field[interface{}] `json:"metadata"`
	Name     param.Field[string]      `json:"name"`
	// Volume size. Must be specified when source is 'new-volume' or 'image'. If
	// specified for source 'snapshot' or 'existing-volume', value must be equal to
	// respective snapshot or volume size
	Size param.Field[int64] `json:"size"`
	// Volume snapshot ID. Mandatory if volume is created from a snapshot
	SnapshotID param.Field[string] `json:"snapshot_id"`
	// One of 'standard', 'ssd_hiiops', 'ssd_local', 'ssd_lowlatency', 'cold', 'ultra'
	TypeName param.Field[CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName] `json:"type_name"`
	// Volume ID. Mandatory if volume is pre-existing volume
	VolumeID param.Field[string] `json:"volume_id"`
}

func (r CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Volume source
type CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource string

const (
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceApptemplate    CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource = "apptemplate"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceExistingVolume CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource = "existing-volume"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceImage          CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource = "image"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceNewVolume      CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource = "new-volume"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceSnapshot       CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource = "snapshot"
)

func (r CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSource) IsKnown() bool {
	switch r {
	case CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceApptemplate, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceExistingVolume, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceImage, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceNewVolume, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceSnapshot:
		return true
	}
	return false
}

// One of 'standard', 'ssd_hiiops', 'ssd_local', 'ssd_lowlatency', 'cold', 'ultra'
type CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName string

const (
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameCold          CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "cold"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdHiiops     CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "ssd_hiiops"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdLocal      CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "ssd_local"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdLowlatency CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "ssd_lowlatency"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameStandard      CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "standard"
	CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameUltra         CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName = "ultra"
)

func (r CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeName) IsKnown() bool {
	switch r {
	case CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameCold, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdHiiops, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdLocal, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameSsdLowlatency, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameStandard, CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameUltra:
		return true
	}
	return false
}

type CloudV1InstanceAvailableFlavorGetFlavorsToResizeParams struct {
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV1InstanceAvailableFlavorGetFlavorsToResizeParams]'s
// query parameters as `url.Values`.
func (r CloudV1InstanceAvailableFlavorGetFlavorsToResizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
