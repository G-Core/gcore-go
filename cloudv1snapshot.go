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

// CloudV1SnapshotService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SnapshotService] method instead.
type CloudV1SnapshotService struct {
	Options []option.RequestOption
}

// NewCloudV1SnapshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1SnapshotService(opts ...option.RequestOption) (r *CloudV1SnapshotService) {
	r = &CloudV1SnapshotService{}
	r.Options = opts
	return
}

// Create snapshot
func (r *CloudV1SnapshotService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1SnapshotNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get snapshot
func (r *CloudV1SnapshotService) Get(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *Snapshot, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List snapshots
func (r *CloudV1SnapshotService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1SnapshotListParams, opts ...option.RequestOption) (res *CloudV1SnapshotListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete snapshot
func (r *CloudV1SnapshotService) Delete(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request. To remove metadata key send body without this key.
func (r *CloudV1SnapshotService) ReplaceMetadata(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SnapshotReplaceMetadataParams, opts ...option.RequestOption) (res *Snapshot, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/snapshots/%v/%v/%s/metadata", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Snapshot object
type Snapshot struct {
	// Snapshot ID
	ID string `json:"id"`
	// Datetime when the volume was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	// Snapshot description
	Description string `json:"description"`
	// Metadata
	Metadata interface{} `json:"metadata"`
	// Snapshot name
	Name string `json:"name"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Snapshot size, GiB
	Size int64 `json:"size"`
	// Snapshot status
	Status string `json:"status"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the volume was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// ID of the volume this snapshot was made from
	VolumeID string       `json:"volume_id"`
	JSON     snapshotJSON `json:"-"`
}

// snapshotJSON contains the JSON metadata for the struct [Snapshot]
type snapshotJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	CreatorTaskID apijson.Field
	Description   apijson.Field
	Metadata      apijson.Field
	Name          apijson.Field
	ProjectID     apijson.Field
	Region        apijson.Field
	RegionID      apijson.Field
	Size          apijson.Field
	Status        apijson.Field
	TaskID        apijson.Field
	UpdatedAt     apijson.Field
	VolumeID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Snapshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotJSON) RawJSON() string {
	return r.raw
}

type CloudV1SnapshotListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []Snapshot                      `json:"results"`
	JSON    cloudV1SnapshotListResponseJSON `json:"-"`
}

// cloudV1SnapshotListResponseJSON contains the JSON metadata for the struct
// [CloudV1SnapshotListResponse]
type cloudV1SnapshotListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1SnapshotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1SnapshotListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1SnapshotNewParams struct {
	// Volume name
	Name param.Field[string] `json:"name,required"`
	// Volume ID to make snapshot of
	VolumeID    param.Field[string] `json:"volume_id,required"`
	Description param.Field[string] `json:"description"`
	// Snapshot metadata
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r CloudV1SnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1SnapshotListParams struct {
	// Can be used to list snapshots of any volume in a specific server instance
	InstanceID param.Field[string] `query:"instance_id"`
	// Can be used to list snapshots by lifecycle policy id
	LifecyclePolicyID param.Field[int64] `query:"lifecycle_policy_id"`
	// Limit the number of returned snapshots
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Can be used to list snapshots by schedule id
	ScheduleID param.Field[string] `query:"schedule_id"`
	// Can be used to list snapshots of a single volume
	VolumeID param.Field[string] `query:"volume_id"`
}

// URLQuery serializes [CloudV1SnapshotListParams]'s query parameters as
// `url.Values`.
func (r CloudV1SnapshotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1SnapshotReplaceMetadataParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1SnapshotReplaceMetadataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
