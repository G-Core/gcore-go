// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// GPUVirtualClusterVolumeService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualClusterVolumeService] method instead.
type GPUVirtualClusterVolumeService struct {
	Options []option.RequestOption
}

// NewGPUVirtualClusterVolumeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUVirtualClusterVolumeService(opts ...option.RequestOption) (r GPUVirtualClusterVolumeService) {
	r = GPUVirtualClusterVolumeService{}
	r.Options = opts
	return
}

// List all volumes attached to servers in a virtual GPU cluster.
func (r *GPUVirtualClusterVolumeService) List(ctx context.Context, clusterID string, query GPUVirtualClusterVolumeListParams, opts ...option.RequestOption) (res *GPUVirtualClusterVolumeList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s/volumes", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GPUVirtualClusterVolume struct {
	// Volume unique identifier
	ID string `json:"id,required" format:"uuid4"`
	// True if this is bootable volume
	Bootable bool `json:"bootable,required"`
	// Volume creation date and time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User defined name
	Name string `json:"name,required"`
	// True if this volume contains root file system
	RootFs bool `json:"root_fs,required"`
	// Server UUID
	ServerID string `json:"server_id,required" format:"uuid4"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// Current volume status
	//
	// Any of "attaching", "available", "awaiting-transfer", "backing-up", "creating",
	// "deleting", "detaching", "downloading", "error", "error_backing-up",
	// "error_deleting", "error_extending", "error_restoring", "extending", "in-use",
	// "maintenance", "reserved", "restoring-backup", "retyping", "reverting",
	// "uploading".
	Status GPUVirtualClusterVolumeStatus `json:"status,required"`
	// User defined tags
	Tags []Tag `json:"tags,required"`
	// Volume type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bootable    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RootFs      respjson.Field
		ServerID    respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterVolume) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current volume status
type GPUVirtualClusterVolumeStatus string

const (
	GPUVirtualClusterVolumeStatusAttaching        GPUVirtualClusterVolumeStatus = "attaching"
	GPUVirtualClusterVolumeStatusAvailable        GPUVirtualClusterVolumeStatus = "available"
	GPUVirtualClusterVolumeStatusAwaitingTransfer GPUVirtualClusterVolumeStatus = "awaiting-transfer"
	GPUVirtualClusterVolumeStatusBackingUp        GPUVirtualClusterVolumeStatus = "backing-up"
	GPUVirtualClusterVolumeStatusCreating         GPUVirtualClusterVolumeStatus = "creating"
	GPUVirtualClusterVolumeStatusDeleting         GPUVirtualClusterVolumeStatus = "deleting"
	GPUVirtualClusterVolumeStatusDetaching        GPUVirtualClusterVolumeStatus = "detaching"
	GPUVirtualClusterVolumeStatusDownloading      GPUVirtualClusterVolumeStatus = "downloading"
	GPUVirtualClusterVolumeStatusError            GPUVirtualClusterVolumeStatus = "error"
	GPUVirtualClusterVolumeStatusErrorBackingUp   GPUVirtualClusterVolumeStatus = "error_backing-up"
	GPUVirtualClusterVolumeStatusErrorDeleting    GPUVirtualClusterVolumeStatus = "error_deleting"
	GPUVirtualClusterVolumeStatusErrorExtending   GPUVirtualClusterVolumeStatus = "error_extending"
	GPUVirtualClusterVolumeStatusErrorRestoring   GPUVirtualClusterVolumeStatus = "error_restoring"
	GPUVirtualClusterVolumeStatusExtending        GPUVirtualClusterVolumeStatus = "extending"
	GPUVirtualClusterVolumeStatusInUse            GPUVirtualClusterVolumeStatus = "in-use"
	GPUVirtualClusterVolumeStatusMaintenance      GPUVirtualClusterVolumeStatus = "maintenance"
	GPUVirtualClusterVolumeStatusReserved         GPUVirtualClusterVolumeStatus = "reserved"
	GPUVirtualClusterVolumeStatusRestoringBackup  GPUVirtualClusterVolumeStatus = "restoring-backup"
	GPUVirtualClusterVolumeStatusRetyping         GPUVirtualClusterVolumeStatus = "retyping"
	GPUVirtualClusterVolumeStatusReverting        GPUVirtualClusterVolumeStatus = "reverting"
	GPUVirtualClusterVolumeStatusUploading        GPUVirtualClusterVolumeStatus = "uploading"
)

type GPUVirtualClusterVolumeList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUVirtualClusterVolume `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterVolumeList) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterVolumeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterVolumeListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
