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

// VolumeSnapshotService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeSnapshotService] method instead.
type VolumeSnapshotService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewVolumeSnapshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVolumeSnapshotService(opts ...option.RequestOption) (r VolumeSnapshotService) {
	r = VolumeSnapshotService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new snapshot from a volume.
func (r *VolumeSnapshotService) New(ctx context.Context, params VolumeSnapshotNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new snapshot from a volume and polls for completion
func (r *VolumeSnapshotService) NewAndPoll(ctx context.Context, params VolumeSnapshotNewParams, opts ...option.RequestOption) (v *Snapshot, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams VolumeSnapshotGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Snapshots) != 1 {
		return nil, errors.New("expected exactly one snapshot to be created in a task")
	}
	resourceID := task.CreatedResources.Snapshots[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Rename snapshot or update tags.
func (r *VolumeSnapshotService) Update(ctx context.Context, snapshotID string, params VolumeSnapshotUpdateParams, opts ...option.RequestOption) (res *Snapshot, err error) {
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
	if snapshotID == "" {
		err = errors.New("missing required snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Delete a specific snapshot.
func (r *VolumeSnapshotService) Delete(ctx context.Context, snapshotID string, body VolumeSnapshotDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if snapshotID == "" {
		err = errors.New("missing required snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a specific snapshot and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *VolumeSnapshotService) DeleteAndPoll(ctx context.Context, snapshotID string, body VolumeSnapshotDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, snapshotID, body, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Get detailed information about a specific snapshot.
func (r *VolumeSnapshotService) Get(ctx context.Context, snapshotID string, query VolumeSnapshotGetParams, opts ...option.RequestOption) (res *Snapshot, err error) {
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
	if snapshotID == "" {
		err = errors.New("missing required snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Snapshot struct {
	// Snapshot ID
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the snapshot was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id" api:"required" format:"uuid4"`
	// Snapshot description
	Description string `json:"description" api:"required"`
	// Snapshot name
	Name string `json:"name" api:"required"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Snapshot size, GiB
	Size int64 `json:"size" api:"required"`
	// Snapshot status
	//
	// Any of "available", "backing-up", "creating", "deleted", "deleting", "error",
	// "error_deleting", "restoring", "unmanaging".
	Status SnapshotStatus `json:"status" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags" api:"required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id" api:"required" format:"uuid4"`
	// Datetime when the snapshot was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the volume this snapshot was made from
	VolumeID string `json:"volume_id" api:"required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CreatorTaskID respjson.Field
		Description   respjson.Field
		Name          respjson.Field
		ProjectID     respjson.Field
		Region        respjson.Field
		RegionID      respjson.Field
		Size          respjson.Field
		Status        respjson.Field
		Tags          respjson.Field
		TaskID        respjson.Field
		UpdatedAt     respjson.Field
		VolumeID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Snapshot) RawJSON() string { return r.JSON.raw }
func (r *Snapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot status
type SnapshotStatus string

const (
	SnapshotStatusAvailable     SnapshotStatus = "available"
	SnapshotStatusBackingUp     SnapshotStatus = "backing-up"
	SnapshotStatusCreating      SnapshotStatus = "creating"
	SnapshotStatusDeleted       SnapshotStatus = "deleted"
	SnapshotStatusDeleting      SnapshotStatus = "deleting"
	SnapshotStatusError         SnapshotStatus = "error"
	SnapshotStatusErrorDeleting SnapshotStatus = "error_deleting"
	SnapshotStatusRestoring     SnapshotStatus = "restoring"
	SnapshotStatusUnmanaging    SnapshotStatus = "unmanaging"
)

type VolumeSnapshotNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Snapshot name
	Name string `json:"name" api:"required"`
	// Volume ID to make snapshot of
	VolumeID string `json:"volume_id" api:"required" format:"uuid4"`
	// Snapshot description
	Description param.Opt[string] `json:"description,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r VolumeSnapshotNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeSnapshotNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeSnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeSnapshotUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Display name for the snapshot (3-63 chars). Used in customer portal and API.
	// Does not affect snapshot data.
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
	//   - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and 'cost_center' while removing 'deprecated_tag',
	//     preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r VolumeSnapshotUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeSnapshotUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeSnapshotUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeSnapshotDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type VolumeSnapshotGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
