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
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// VolumeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new volume in the project and region. The volume can be created from
// scratch, from an image, or from a snapshot. Optionally attach the volume to an
// instance during creation.
func (r *VolumeService) New(ctx context.Context, params VolumeNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new volume and polls the corresponding task until it is completed.
func (r *VolumeService) NewAndPoll(ctx context.Context, params VolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams VolumeGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Volumes) != 1 {
		return nil, errors.New("expected exactly one network to be created in a task")
	}
	resourceID := task.CreatedResources.Volumes[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Rename a volume or update tags
func (r *VolumeService) Update(ctx context.Context, volumeID string, params VolumeUpdateParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve a list of volumes in the project and region. The list can be filtered
// by various parameters like bootable status, metadata/tags, attachments, instance
// ID, name, and ID.
func (r *VolumeService) List(ctx context.Context, params VolumeListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Volume], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of volumes in the project and region. The list can be filtered
// by various parameters like bootable status, metadata/tags, attachments, instance
// ID, name, and ID.
func (r *VolumeService) ListAutoPaging(ctx context.Context, params VolumeListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Volume] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a volume and all its snapshots. The volume must be in an available state
// to be deleted.
func (r *VolumeService) Delete(ctx context.Context, volumeID string, params VolumeDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// DeleteAndPoll deletes a volume and polls the corresponding task until it is completed.
// Use the [TaskService.Poll] method if you need to poll for all tasks.
func (r *VolumeService) DeleteAndPoll(ctx context.Context, volumeID string, params VolumeDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, volumeID, params, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Attach the volume to instance. Note: ultra volume can only be attached to an
// instance with shared flavor
func (r *VolumeService) AttachToInstance(ctx context.Context, volumeID string, params VolumeAttachToInstanceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/attach", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// AttachToInstanceAndPoll attaches the volume to instance and polls the corresponding task until it is completed. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *VolumeService) AttachToInstanceAndPoll(ctx context.Context, volumeID string, params VolumeAttachToInstanceParams, opts ...option.RequestOption) error {
	resource, err := r.AttachToInstance(ctx, volumeID, params, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Change the type of a volume. The volume must not have any snapshots to change
// its type.
func (r *VolumeService) ChangeType(ctx context.Context, volumeID string, params VolumeChangeTypeParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/retype", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach the volume from instance
func (r *VolumeService) DetachFromInstance(ctx context.Context, volumeID string, params VolumeDetachFromInstanceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/detach", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DetachFromInstanceAndPoll detaches the volume to instance and polls the corresponding task until it is completed.
// Use the [TaskService.Poll] method if you need to poll for all tasks.
func (r *VolumeService) DetachFromInstanceAndPoll(ctx context.Context, volumeID string, params VolumeDetachFromInstanceParams, opts ...option.RequestOption) error {
	resource, err := r.DetachFromInstance(ctx, volumeID, params, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Retrieve detailed information about a specific volume.
func (r *VolumeService) Get(ctx context.Context, volumeID string, query VolumeGetParams, opts ...option.RequestOption) (res *Volume, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Increase the size of a volume. The new size must be greater than the current
// size.
func (r *VolumeService) Resize(ctx context.Context, volumeID string, params VolumeResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/extend", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ResizeAndPoll increases the size of a volume and polls the corresponding task until it is completed. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *VolumeService) ResizeAndPoll(ctx context.Context, volumeID string, params VolumeResizeParams, opts ...option.RequestOption) (res *Volume, err error) {
	resource, err := r.Resize(ctx, volumeID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams VolumeGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	return r.Get(ctx, volumeID, getParams, opts...)
}

// Revert a volume to its last snapshot. The volume must be in an available state
// to be reverted.
func (r *VolumeService) RevertToLastSnapshot(ctx context.Context, volumeID string, body VolumeRevertToLastSnapshotParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/revert", body.ProjectID.Value, body.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
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
	//
	// Any of "attaching", "available", "awaiting-transfer", "backing-up", "creating",
	// "deleting", "detaching", "downloading", "error", "error_backing-up",
	// "error_deleting", "error_extending", "error_restoring", "extending", "in-use",
	// "maintenance", "reserved", "restoring-backup", "retyping", "reverting",
	// "uploading".
	Status VolumeStatus `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The type of volume storage.
	VolumeType string `json:"volume_type,required"`
	// List of attachments associated with the volume.
	Attachments []VolumeAttachment `json:"attachments,nullable"`
	// The ID of the task that created this volume.
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Schema representing the Quality of Service (QoS) parameters for a volume.
	LimiterStats VolumeLimiterStats `json:"limiter_stats,nullable"`
	// List of snapshot IDs associated with this volume.
	SnapshotIDs []string `json:"snapshot_ids,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable"`
	// The date and time when the volume was last updated.
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Image metadata for volumes created from an image.
	VolumeImageMetadata map[string]string `json:"volume_image_metadata,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Bootable            respjson.Field
		CreatedAt           respjson.Field
		IsRootVolume        respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		Size                respjson.Field
		Status              respjson.Field
		Tags                respjson.Field
		VolumeType          respjson.Field
		Attachments         respjson.Field
		CreatorTaskID       respjson.Field
		LimiterStats        respjson.Field
		SnapshotIDs         respjson.Field
		TaskID              respjson.Field
		UpdatedAt           respjson.Field
		VolumeImageMetadata respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	VolumeStatusReverting        VolumeStatus = "reverting"
	VolumeStatusUploading        VolumeStatus = "uploading"
)

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
	ServerID string `json:"server_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentID respjson.Field
		VolumeID     respjson.Field
		AttachedAt   respjson.Field
		Device       respjson.Field
		FlavorID     respjson.Field
		InstanceName respjson.Field
		ServerID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeAttachment) RawJSON() string { return r.JSON.raw }
func (r *VolumeAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	MBpsBurstLimit int64 `json:"MBps_burst_limit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IopsBaseLimit  respjson.Field
		IopsBurstLimit respjson.Field
		MBpsBaseLimit  respjson.Field
		MBpsBurstLimit respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeLimiterStats) RawJSON() string { return r.JSON.raw }
func (r *VolumeLimiterStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfImage *VolumeNewParamsBodyImage `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSnapshot *VolumeNewParamsBodySnapshot `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfNewVolume *VolumeNewParamsBodyNewVolume `json:",inline"`

	paramObj
}

func (u VolumeNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfSnapshot, u.OfNewVolume)
}
func (r *VolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImageID, Name, Size, Source are required.
type VolumeNewParamsBodyImage struct {
	// Image ID
	ImageID string `json:"image_id,required" format:"uuid4"`
	// Volume name
	Name string `json:"name,required"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// Block device attachment tag (not exposed in the user tags). Only used in
	// conjunction with `instance_id_to_attach_to`
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// `instance_id` to attach newly-created volume to
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// List of lifecycle policy IDs (snapshot creation schedules) to associate with the
	// volume
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type. Defaults to `standard`. If not specified for source `snapshot`,
	// volume type will be derived from the snapshot volume.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// Volume source type
	//
	// This field can be elided, and will marshal its zero value as "image".
	Source constant.Image `json:"source,required"`
	paramObj
}

func (r VolumeNewParamsBodyImage) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodyImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeNewParamsBodyImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodyImage](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Name, SnapshotID, Source are required.
type VolumeNewParamsBodySnapshot struct {
	// Volume name
	Name string `json:"name,required"`
	// Snapshot ID
	SnapshotID string `json:"snapshot_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the user tags). Only used in
	// conjunction with `instance_id_to_attach_to`
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// `instance_id` to attach newly-created volume to
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// Volume size in GiB. If specified, value must be equal to respective snapshot
	// size
	Size param.Opt[int64] `json:"size,omitzero"`
	// List of lifecycle policy IDs (snapshot creation schedules) to associate with the
	// volume
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type. Defaults to `standard`. If not specified for source `snapshot`,
	// volume type will be derived from the snapshot volume.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// Volume source type
	//
	// This field can be elided, and will marshal its zero value as "snapshot".
	Source constant.Snapshot `json:"source,required"`
	paramObj
}

func (r VolumeNewParamsBodySnapshot) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodySnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeNewParamsBodySnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodySnapshot](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Name, Size, Source are required.
type VolumeNewParamsBodyNewVolume struct {
	// Volume name
	Name string `json:"name,required"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// Block device attachment tag (not exposed in the user tags). Only used in
	// conjunction with `instance_id_to_attach_to`
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// `instance_id` to attach newly-created volume to
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// List of lifecycle policy IDs (snapshot creation schedules) to associate with the
	// volume
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type. Defaults to `standard`. If not specified for source `snapshot`,
	// volume type will be derived from the snapshot volume.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// Volume source type
	//
	// This field can be elided, and will marshal its zero value as "new-volume".
	Source constant.NewVolume `json:"source,required"`
	paramObj
}

func (r VolumeNewParamsBodyNewVolume) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodyNewVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeNewParamsBodyNewVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodyNewVolume](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

type VolumeUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//
	// - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and '`cost_center`' while removing
	//     '`deprecated_tag`', preserving other existing tags.
	//
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Filter by bootable field
	Bootable param.Opt[bool] `query:"bootable,omitzero" json:"-"`
	// Filter volumes by k8s cluster ID
	ClusterID param.Opt[string] `query:"cluster_id,omitzero" json:"-"`
	// Filter by the presence of attachments
	HasAttachments param.Opt[bool] `query:"has_attachments,omitzero" json:"-"`
	// Filter the volume list result by the ID part of the volume
	IDPart param.Opt[string] `query:"id_part,omitzero" json:"-"`
	// Filter volumes by instance ID
	InstanceID param.Opt[string] `query:"instance_id,omitzero" format:"uuid4" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter volumes by `name_part` inclusion in volume name.Any substring can be used
	// and volumes will be returned with names containing the substring.
	NamePart param.Opt[string] `query:"name_part,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VolumeListParams]'s query parameters as `url.Values`.
func (r VolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VolumeDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Comma separated list of snapshot IDs to be deleted with the volume.
	Snapshots param.Opt[string] `query:"snapshots,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VolumeDeleteParams]'s query parameters as `url.Values`.
func (r VolumeDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VolumeAttachToInstanceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Instance ID.
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the normal tags).
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	paramObj
}

func (r VolumeAttachToInstanceParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAttachToInstanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeAttachToInstanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeChangeTypeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// New volume type name
	//
	// Any of "ssd_hiiops", "standard".
	VolumeType VolumeChangeTypeParamsVolumeType `json:"volume_type,omitzero,required"`
	paramObj
}

func (r VolumeChangeTypeParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeChangeTypeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeChangeTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New volume type name
type VolumeChangeTypeParamsVolumeType string

const (
	VolumeChangeTypeParamsVolumeTypeSsdHiiops VolumeChangeTypeParamsVolumeType = "ssd_hiiops"
	VolumeChangeTypeParamsVolumeTypeStandard  VolumeChangeTypeParamsVolumeType = "standard"
)

type VolumeDetachFromInstanceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Instance ID
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	paramObj
}

func (r VolumeDetachFromInstanceParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeDetachFromInstanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeDetachFromInstanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type VolumeResizeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// New volume size in GiB
	Size int64 `json:"size,required"`
	paramObj
}

func (r VolumeResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeRevertToLastSnapshotParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
