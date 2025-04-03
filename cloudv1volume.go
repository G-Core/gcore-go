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

// CloudV1VolumeService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1VolumeService] method instead.
type CloudV1VolumeService struct {
	Options      []option.RequestOption
	Metadata     *CloudV1VolumeMetadataService
	MetadataItem *CloudV1VolumeMetadataItemService
}

// NewCloudV1VolumeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1VolumeService(opts ...option.RequestOption) (r *CloudV1VolumeService) {
	r = &CloudV1VolumeService{}
	r.Options = opts
	r.Metadata = NewCloudV1VolumeMetadataService(opts...)
	r.MetadataItem = NewCloudV1VolumeMetadataItemService(opts...)
	return
}

// Create volume
func (r *CloudV1VolumeService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1VolumeNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get volume
func (r *CloudV1VolumeService) Get(ctx context.Context, projectID int64, regionID int64, volumeID string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rename volume
func (r *CloudV1VolumeService) Update(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeUpdateParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List volumes
func (r *CloudV1VolumeService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1VolumeListParams, opts ...option.RequestOption) (res *CloudV1VolumeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete volume
func (r *CloudV1VolumeService) Delete(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeDeleteParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Attach the volume to instance. Note: ultra volume can only be attached to an
// instance with shared flavor
func (r *CloudV1VolumeService) Attach(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeAttachParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/attach", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate volume type change
func (r *CloudV1VolumeService) ChangeType(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeChangeTypeParams, opts ...option.RequestOption) (res *DeprecatedVolume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/retype", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach the volume from instance
func (r *CloudV1VolumeService) Detach(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeDetachParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/detach", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extend volume
func (r *CloudV1VolumeService) Extend(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeExtendParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/extend", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Revert volume to it's last snapshot
func (r *CloudV1VolumeService) Revert(ctx context.Context, projectID int64, regionID int64, volumeID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/revert", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Attach volume to instance schema
type AttachVolumeToInstanceParam struct {
	// Instance ID.
	InstanceID param.Field[string] `json:"instance_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the metadata).
	AttachmentTag param.Field[string] `json:"attachment_tag"`
}

func (r AttachVolumeToInstanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Volume object
type DeprecatedVolume struct {
	// Volume ID
	ID string `json:"id"`
	// Attachment list
	Attachments []DeprecatedVolumeAttachment `json:"attachments"`
	// Bootable boolean flag
	Bootable bool `json:"bootable"`
	// Datetime when the volume was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	Device        string `json:"device"`
	InstanceID    string `json:"instance_id"`
	// Root volume flag
	IsRootVolume bool `json:"is_root_volume"`
	// Volume QoS parameters.
	LimiterStats DeprecatedVolumeLimiterStats `json:"limiter_stats"`
	// Metadata
	Metadata interface{} `json:"metadata"`
	// Create one or more metadata items for a volume
	MetadataDetailed interface{} `json:"metadata_detailed"`
	// Volume name
	Name string `json:"name,nullable"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Volume size, GiB
	Size int64 `json:"size"`
	// Snapshots of this volume
	SnapshotIDs []string `json:"snapshot_ids"`
	// Volume status
	Status string `json:"status"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the volume was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Image information for volumes that were created from image
	VolumeImageMetadata interface{} `json:"volume_image_metadata"`
	// Volume type
	VolumeType string               `json:"volume_type"`
	JSON       deprecatedVolumeJSON `json:"-"`
}

// deprecatedVolumeJSON contains the JSON metadata for the struct
// [DeprecatedVolume]
type deprecatedVolumeJSON struct {
	ID                  apijson.Field
	Attachments         apijson.Field
	Bootable            apijson.Field
	CreatedAt           apijson.Field
	CreatorTaskID       apijson.Field
	Device              apijson.Field
	InstanceID          apijson.Field
	IsRootVolume        apijson.Field
	LimiterStats        apijson.Field
	Metadata            apijson.Field
	MetadataDetailed    apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Size                apijson.Field
	SnapshotIDs         apijson.Field
	Status              apijson.Field
	TaskID              apijson.Field
	UpdatedAt           apijson.Field
	VolumeImageMetadata apijson.Field
	VolumeType          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeprecatedVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedVolumeJSON) RawJSON() string {
	return r.raw
}

// Volume attachment object
type DeprecatedVolumeAttachment struct {
	// ID of attachment object
	AttachmentID string `json:"attachment_id,required"`
	// Volume ID
	VolumeID string `json:"volume_id,required"`
	// Attachment creation datetime
	AttachedAt time.Time `json:"attached_at" format:"date-time"`
	// Block device name in guest
	Device string `json:"device"`
	// Instance flavor ID
	FlavorID string `json:"flavor_id"`
	// Instance name (if attached and server name is known)
	InstanceName string `json:"instance_name"`
	// Instance ID
	ServerID string                         `json:"server_id"`
	JSON     deprecatedVolumeAttachmentJSON `json:"-"`
}

// deprecatedVolumeAttachmentJSON contains the JSON metadata for the struct
// [DeprecatedVolumeAttachment]
type deprecatedVolumeAttachmentJSON struct {
	AttachmentID apijson.Field
	VolumeID     apijson.Field
	AttachedAt   apijson.Field
	Device       apijson.Field
	FlavorID     apijson.Field
	InstanceName apijson.Field
	ServerID     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DeprecatedVolumeAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedVolumeAttachmentJSON) RawJSON() string {
	return r.raw
}

// Volume QoS parameters.
type DeprecatedVolumeLimiterStats struct {
	// Sustained IOPS limit.
	IopsBaseLimit int64 `json:"iops_base_limit,required"`
	// Burst IOPS limit.
	IopsBurstLimit int64 `json:"iops_burst_limit,required"`
	// Sustained bandwidth limit.
	MBpsBaseLimit int64 `json:"MBps_base_limit,required"`
	// Burst bandwidth limit.
	MBpsBurstLimit int64                            `json:"MBps_burst_limit,required"`
	JSON           deprecatedVolumeLimiterStatsJSON `json:"-"`
}

// deprecatedVolumeLimiterStatsJSON contains the JSON metadata for the struct
// [DeprecatedVolumeLimiterStats]
type deprecatedVolumeLimiterStatsJSON struct {
	IopsBaseLimit  apijson.Field
	IopsBurstLimit apijson.Field
	MBpsBaseLimit  apijson.Field
	MBpsBurstLimit apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeprecatedVolumeLimiterStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedVolumeLimiterStatsJSON) RawJSON() string {
	return r.raw
}

type InstanceIDParam struct {
	// Instance ID
	InstanceID param.Field[string] `json:"instance_id,required" format:"uuid4"`
}

func (r InstanceIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Volume struct {
	// The unique identifier of the volume.
	ID string `json:"id,required"`
	// Indicates whether the volume is bootable.
	Bootable bool `json:"bootable,required"`
	// The date and time when the volume was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates whether this is a root volume.
	IsRootVolume bool `json:"is_root_volume,required"`
	// The name of the volume.
	Name string `json:"name,required"`
	// Project ID.
	ProjectID int64 `json:"project_id,required"`
	// The region where the volume is located.
	Region string `json:"region,required"`
	// The identifier of the region.
	RegionID int64 `json:"region_id,required"`
	// The size of the volume in gibibytes (GiB).
	Size int64 `json:"size,required"`
	// The current status of the volume.
	Status VolumeStatus `json:"status,required"`
	// The type of volume storage.
	VolumeType string `json:"volume_type,required"`
	// List of attachments associated with the volume.
	Attachments []VolumeAttachment `json:"attachments,nullable"`
	// The ID of the task that created this volume.
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Schema representing the Quality of Service (QoS) parameters for a volume.
	LimiterStats VolumeLimiterStats `json:"limiter_stats,nullable"`
	// Metadata associated with the volume.
	Metadata map[string]string `json:"metadata,nullable"`
	// Detailed metadata items for the volume.
	MetadataDetailed []DetailedMetadata `json:"metadata_detailed,nullable"`
	// List of snapshot IDs associated with this volume.
	SnapshotIDs []string `json:"snapshot_ids,nullable"`
	// The ID of the active task associated with the volume.
	TaskID string `json:"task_id,nullable"`
	// The date and time when the volume was last updated.
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Image metadata for volumes created from an image.
	VolumeImageMetadata map[string]string `json:"volume_image_metadata,nullable"`
	JSON                volumeJSON        `json:"-"`
}

// volumeJSON contains the JSON metadata for the struct [Volume]
type volumeJSON struct {
	ID                  apijson.Field
	Bootable            apijson.Field
	CreatedAt           apijson.Field
	IsRootVolume        apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Size                apijson.Field
	Status              apijson.Field
	VolumeType          apijson.Field
	Attachments         apijson.Field
	CreatorTaskID       apijson.Field
	LimiterStats        apijson.Field
	Metadata            apijson.Field
	MetadataDetailed    apijson.Field
	SnapshotIDs         apijson.Field
	TaskID              apijson.Field
	UpdatedAt           apijson.Field
	VolumeImageMetadata apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Volume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeJSON) RawJSON() string {
	return r.raw
}

// The current status of the volume.
type VolumeStatus string

const (
	VolumeStatusAttaching        VolumeStatus = "attaching"
	VolumeStatusAvailable        VolumeStatus = "available"
	VolumeStatusAwaitingTransfer VolumeStatus = "awaiting-transfer"
	VolumeStatusBackingUp        VolumeStatus = "backing-up"
	VolumeStatusCreating         VolumeStatus = "creating"
	VolumeStatusDeleting         VolumeStatus = "deleting"
	VolumeStatusDetaching        VolumeStatus = "detaching"
	VolumeStatusDownloading      VolumeStatus = "downloading"
	VolumeStatusError            VolumeStatus = "error"
	VolumeStatusErrorBackingUp   VolumeStatus = "error_backing-up"
	VolumeStatusErrorDeleting    VolumeStatus = "error_deleting"
	VolumeStatusErrorExtending   VolumeStatus = "error_extending"
	VolumeStatusErrorRestoring   VolumeStatus = "error_restoring"
	VolumeStatusExtending        VolumeStatus = "extending"
	VolumeStatusInUse            VolumeStatus = "in-use"
	VolumeStatusMaintenance      VolumeStatus = "maintenance"
	VolumeStatusReserved         VolumeStatus = "reserved"
	VolumeStatusRestoringBackup  VolumeStatus = "restoring-backup"
	VolumeStatusRetyping         VolumeStatus = "retyping"
	VolumeStatusUploading        VolumeStatus = "uploading"
)

func (r VolumeStatus) IsKnown() bool {
	switch r {
	case VolumeStatusAttaching, VolumeStatusAvailable, VolumeStatusAwaitingTransfer, VolumeStatusBackingUp, VolumeStatusCreating, VolumeStatusDeleting, VolumeStatusDetaching, VolumeStatusDownloading, VolumeStatusError, VolumeStatusErrorBackingUp, VolumeStatusErrorDeleting, VolumeStatusErrorExtending, VolumeStatusErrorRestoring, VolumeStatusExtending, VolumeStatusInUse, VolumeStatusMaintenance, VolumeStatusReserved, VolumeStatusRestoringBackup, VolumeStatusRetyping, VolumeStatusUploading:
		return true
	}
	return false
}

type VolumeAttachment struct {
	// The unique identifier of the attachment object.
	AttachmentID string `json:"attachment_id,required"`
	// The unique identifier of the attached volume.
	VolumeID string `json:"volume_id,required"`
	// The date and time when the attachment was created.
	AttachedAt time.Time `json:"attached_at,nullable" format:"date-time"`
	// The block device name inside the guest instance.
	Device string `json:"device,nullable"`
	// The flavor ID of the instance.
	FlavorID string `json:"flavor_id,nullable"`
	// The name of the instance if attached and the server name is known.
	InstanceName string `json:"instance_name,nullable"`
	// The unique identifier of the instance.
	ServerID string               `json:"server_id,nullable"`
	JSON     volumeAttachmentJSON `json:"-"`
}

// volumeAttachmentJSON contains the JSON metadata for the struct
// [VolumeAttachment]
type volumeAttachmentJSON struct {
	AttachmentID apijson.Field
	VolumeID     apijson.Field
	AttachedAt   apijson.Field
	Device       apijson.Field
	FlavorID     apijson.Field
	InstanceName apijson.Field
	ServerID     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VolumeAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeAttachmentJSON) RawJSON() string {
	return r.raw
}

// Schema representing the Quality of Service (QoS) parameters for a volume.
type VolumeLimiterStats struct {
	// The sustained IOPS (Input/Output Operations Per Second) limit.
	IopsBaseLimit int64 `json:"iops_base_limit,required"`
	// The burst IOPS limit.
	IopsBurstLimit int64 `json:"iops_burst_limit,required"`
	// The sustained bandwidth limit in megabytes per second (MBps).
	MBpsBaseLimit int64 `json:"MBps_base_limit,required"`
	// The burst bandwidth limit in megabytes per second (MBps).
	MBpsBurstLimit int64                  `json:"MBps_burst_limit,required"`
	JSON           volumeLimiterStatsJSON `json:"-"`
}

// volumeLimiterStatsJSON contains the JSON metadata for the struct
// [VolumeLimiterStats]
type volumeLimiterStatsJSON struct {
	IopsBaseLimit  apijson.Field
	IopsBurstLimit apijson.Field
	MBpsBaseLimit  apijson.Field
	MBpsBurstLimit apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VolumeLimiterStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeLimiterStatsJSON) RawJSON() string {
	return r.raw
}

type CloudV1VolumeListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []DeprecatedVolume            `json:"results"`
	JSON    cloudV1VolumeListResponseJSON `json:"-"`
}

// cloudV1VolumeListResponseJSON contains the JSON metadata for the struct
// [CloudV1VolumeListResponse]
type cloudV1VolumeListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1VolumeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1VolumeListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1VolumeNewParams struct {
	Body CloudV1VolumeNewParamsBodyUnion `json:"body"`
}

func (r CloudV1VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CloudV1VolumeNewParamsBody struct {
	// Volume name
	Name param.Field[string] `json:"name,required"`
	// Block device attachment tag (not exposed in the metadata). Only used in
	// conjunction with instance_id_to_attach_to
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// Image ID
	ImageID param.Field[string] `json:"image_id"`
	// VM's instance_id to attach newly-created volume to
	InstanceIDToAttachTo param.Field[string]      `json:"instance_id_to_attach_to"`
	LifecyclePolicyIDs   param.Field[interface{}] `json:"lifecycle_policy_ids"`
	Metadata             param.Field[interface{}] `json:"metadata"`
	// Volume size. Must be specified
	Size param.Field[int64] `json:"size"`
	// Snapshot ID
	SnapshotID param.Field[string] `json:"snapshot_id"`
	Source     param.Field[string] `json:"source"`
	// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
	// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
	// defaults to standard.
	TypeName param.Field[CloudV1VolumeNewParamsBodyTypeName] `json:"type_name"`
}

func (r CloudV1VolumeNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1VolumeNewParamsBody) implementsCloudV1VolumeNewParamsBodyUnion() {}

// Satisfied by [CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchema],
// [CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchema],
// [CloudV1VolumeNewParamsBodyCreateNewVolumeSchema], [CloudV1VolumeNewParamsBody].
type CloudV1VolumeNewParamsBodyUnion interface {
	implementsCloudV1VolumeNewParamsBodyUnion()
}

type CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchema struct {
	// Image ID
	ImageID param.Field[string] `json:"image_id,required"`
	// Volume name
	Name param.Field[string] `json:"name,required"`
	// Volume size. Must be specified
	Size param.Field[int64] `json:"size,required"`
	// Block device attachment tag (not exposed in the metadata). Only used in
	// conjunction with instance_id_to_attach_to
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// VM's instance_id to attach newly-created volume to
	InstanceIDToAttachTo param.Field[string]  `json:"instance_id_to_attach_to"`
	LifecyclePolicyIDs   param.Field[[]int64] `json:"lifecycle_policy_ids"`
	// Create one or more metadata items for a volume
	Metadata param.Field[interface{}] `json:"metadata"`
	Source   param.Field[string]      `json:"source"`
	// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
	// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
	// defaults to standard.
	TypeName param.Field[CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName] `json:"type_name"`
}

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchema) implementsCloudV1VolumeNewParamsBodyUnion() {
}

// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
// defaults to standard.
type CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName string

const (
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameCold          CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "cold"
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdHiiops     CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "ssd_hiiops"
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdLocal      CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "ssd_local"
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdLowlatency CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "ssd_lowlatency"
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameStandard      CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "standard"
	CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameUltra         CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName = "ultra"
)

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeName) IsKnown() bool {
	switch r {
	case CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameCold, CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdHiiops, CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdLocal, CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameSsdLowlatency, CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameStandard, CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameUltra:
		return true
	}
	return false
}

type CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchema struct {
	// Volume name
	Name param.Field[string] `json:"name,required"`
	// Snapshot ID
	SnapshotID param.Field[string] `json:"snapshot_id,required"`
	// Block device attachment tag (not exposed in the metadata). Only used in
	// conjunction with instance_id_to_attach_to
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// VM's instance_id to attach newly-created volume to
	InstanceIDToAttachTo param.Field[string]  `json:"instance_id_to_attach_to"`
	LifecyclePolicyIDs   param.Field[[]int64] `json:"lifecycle_policy_ids"`
	// Create one or more metadata items for a volume
	Metadata param.Field[interface{}] `json:"metadata"`
	// Volume size. If specified, value must be equal to respective snapshot size
	Size   param.Field[int64]  `json:"size"`
	Source param.Field[string] `json:"source"`
	// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
	// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
	// defaults to standard.
	TypeName param.Field[CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName] `json:"type_name"`
}

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchema) implementsCloudV1VolumeNewParamsBodyUnion() {
}

// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
// defaults to standard.
type CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName string

const (
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameCold          CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "cold"
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdHiiops     CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "ssd_hiiops"
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdLocal      CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "ssd_local"
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdLowlatency CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "ssd_lowlatency"
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameStandard      CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "standard"
	CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameUltra         CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName = "ultra"
)

func (r CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeName) IsKnown() bool {
	switch r {
	case CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameCold, CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdHiiops, CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdLocal, CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameSsdLowlatency, CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameStandard, CloudV1VolumeNewParamsBodyCreateVolumeFromSnapshotSchemaTypeNameUltra:
		return true
	}
	return false
}

type CloudV1VolumeNewParamsBodyCreateNewVolumeSchema struct {
	// Volume name
	Name param.Field[string] `json:"name,required"`
	// Volume size. Must be specified
	Size param.Field[int64] `json:"size,required"`
	// Block device attachment tag (not exposed in the metadata). Only used in
	// conjunction with instance_id_to_attach_to
	AttachmentTag param.Field[string] `json:"attachment_tag"`
	// VM's instance_id to attach newly-created volume to
	InstanceIDToAttachTo param.Field[string]  `json:"instance_id_to_attach_to"`
	LifecyclePolicyIDs   param.Field[[]int64] `json:"lifecycle_policy_ids"`
	// Create one or more metadata items for a volume
	Metadata param.Field[interface{}] `json:"metadata"`
	Source   param.Field[string]      `json:"source"`
	// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
	// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
	// defaults to standard.
	TypeName param.Field[CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName] `json:"type_name"`
}

func (r CloudV1VolumeNewParamsBodyCreateNewVolumeSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1VolumeNewParamsBodyCreateNewVolumeSchema) implementsCloudV1VolumeNewParamsBodyUnion() {
}

// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
// defaults to standard.
type CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName string

const (
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameCold          CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "cold"
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdHiiops     CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "ssd_hiiops"
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdLocal      CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "ssd_local"
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdLowlatency CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "ssd_lowlatency"
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameStandard      CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "standard"
	CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameUltra         CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName = "ultra"
)

func (r CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeName) IsKnown() bool {
	switch r {
	case CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameCold, CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdHiiops, CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdLocal, CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameSsdLowlatency, CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameStandard, CloudV1VolumeNewParamsBodyCreateNewVolumeSchemaTypeNameUltra:
		return true
	}
	return false
}

// One of 'standard', 'ssd_hiiops', 'cold', 'ultra'. If not specified for source
// 'snapshot', volume type will be derived from the snapshot volume. Otherwise
// defaults to standard.
type CloudV1VolumeNewParamsBodyTypeName string

const (
	CloudV1VolumeNewParamsBodyTypeNameCold          CloudV1VolumeNewParamsBodyTypeName = "cold"
	CloudV1VolumeNewParamsBodyTypeNameSsdHiiops     CloudV1VolumeNewParamsBodyTypeName = "ssd_hiiops"
	CloudV1VolumeNewParamsBodyTypeNameSsdLocal      CloudV1VolumeNewParamsBodyTypeName = "ssd_local"
	CloudV1VolumeNewParamsBodyTypeNameSsdLowlatency CloudV1VolumeNewParamsBodyTypeName = "ssd_lowlatency"
	CloudV1VolumeNewParamsBodyTypeNameStandard      CloudV1VolumeNewParamsBodyTypeName = "standard"
	CloudV1VolumeNewParamsBodyTypeNameUltra         CloudV1VolumeNewParamsBodyTypeName = "ultra"
)

func (r CloudV1VolumeNewParamsBodyTypeName) IsKnown() bool {
	switch r {
	case CloudV1VolumeNewParamsBodyTypeNameCold, CloudV1VolumeNewParamsBodyTypeNameSsdHiiops, CloudV1VolumeNewParamsBodyTypeNameSsdLocal, CloudV1VolumeNewParamsBodyTypeNameSsdLowlatency, CloudV1VolumeNewParamsBodyTypeNameStandard, CloudV1VolumeNewParamsBodyTypeNameUltra:
		return true
	}
	return false
}

type CloudV1VolumeUpdateParams struct {
	Name NameParam `json:"name,required"`
}

func (r CloudV1VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Name)
}

type CloudV1VolumeListParams struct {
	// Filter by a bootable field
	Bootable param.Field[bool] `query:"bootable"`
	// Can be used to only show volumes of a specific k8s cluster
	ClusterID param.Field[string] `query:"cluster_id"`
	// Filter by the presence of attachments
	HasAttachments param.Field[bool] `query:"has_attachments"`
	// Filter the volume list result by the ID part of the volume
	IDPart param.Field[string] `query:"id_part"`
	// Can be used to only show volumes of a specific instance
	InstanceID param.Field[string] `query:"instance_id"`
	// Limit the number of returned volumes
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata keyss. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_k=["key1", "key2"]" --url "http://localhost:1111/v1/volumes/1/1"
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/volumes/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Filter out volumes by name_part inclusion in volume name. Any substring can be
	// used and volumes will be returned with names containing the substring. Example:
	// "test".
	NamePart param.Field[string] `query:"name_part"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1VolumeListParams]'s query parameters as
// `url.Values`.
func (r CloudV1VolumeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1VolumeDeleteParams struct {
	// Comma separated list of snapshot IDs to be deleted with the volume.
	Snapshots param.Field[string] `query:"snapshots"`
}

// URLQuery serializes [CloudV1VolumeDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV1VolumeDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1VolumeAttachParams struct {
	// Attach volume to instance schema
	AttachVolumeToInstance AttachVolumeToInstanceParam `json:"attach_volume_to_instance,required"`
}

func (r CloudV1VolumeAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AttachVolumeToInstance)
}

type CloudV1VolumeChangeTypeParams struct {
	// New volume type name
	VolumeType param.Field[CloudV1VolumeChangeTypeParamsVolumeType] `json:"volume_type,required"`
}

func (r CloudV1VolumeChangeTypeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New volume type name
type CloudV1VolumeChangeTypeParamsVolumeType string

const (
	CloudV1VolumeChangeTypeParamsVolumeTypeCold          CloudV1VolumeChangeTypeParamsVolumeType = "cold"
	CloudV1VolumeChangeTypeParamsVolumeTypeSsdHiiops     CloudV1VolumeChangeTypeParamsVolumeType = "ssd_hiiops"
	CloudV1VolumeChangeTypeParamsVolumeTypeSsdLocal      CloudV1VolumeChangeTypeParamsVolumeType = "ssd_local"
	CloudV1VolumeChangeTypeParamsVolumeTypeSsdLowlatency CloudV1VolumeChangeTypeParamsVolumeType = "ssd_lowlatency"
	CloudV1VolumeChangeTypeParamsVolumeTypeStandard      CloudV1VolumeChangeTypeParamsVolumeType = "standard"
	CloudV1VolumeChangeTypeParamsVolumeTypeUltra         CloudV1VolumeChangeTypeParamsVolumeType = "ultra"
)

func (r CloudV1VolumeChangeTypeParamsVolumeType) IsKnown() bool {
	switch r {
	case CloudV1VolumeChangeTypeParamsVolumeTypeCold, CloudV1VolumeChangeTypeParamsVolumeTypeSsdHiiops, CloudV1VolumeChangeTypeParamsVolumeTypeSsdLocal, CloudV1VolumeChangeTypeParamsVolumeTypeSsdLowlatency, CloudV1VolumeChangeTypeParamsVolumeTypeStandard, CloudV1VolumeChangeTypeParamsVolumeTypeUltra:
		return true
	}
	return false
}

type CloudV1VolumeDetachParams struct {
	InstanceID InstanceIDParam `json:"instance_id,required"`
}

func (r CloudV1VolumeDetachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstanceID)
}

type CloudV1VolumeExtendParams struct {
	// Size in GiB
	Size param.Field[int64] `json:"size,required"`
}

func (r CloudV1VolumeExtendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
