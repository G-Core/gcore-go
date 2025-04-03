// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LifecyclePolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LifecyclePolicyService] method instead.
type CloudV1LifecyclePolicyService struct {
	Options []option.RequestOption
}

// NewCloudV1LifecyclePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1LifecyclePolicyService(opts ...option.RequestOption) (r *CloudV1LifecyclePolicyService) {
	r = &CloudV1LifecyclePolicyService{}
	r.Options = opts
	return
}

// Create a policy
func (r *CloudV1LifecyclePolicyService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1LifecyclePolicyNewParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the policy by id
func (r *CloudV1LifecyclePolicyService) Get(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the policy
func (r *CloudV1LifecyclePolicyService) Update(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, body CloudV1LifecyclePolicyUpdateParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all policies
func (r *CloudV1LifecyclePolicyService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1LifecyclePolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the policy
func (r *CloudV1LifecyclePolicyService) Delete(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Add new schedules to the policy
func (r *CloudV1LifecyclePolicyService) AddSchedules(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, body CloudV1LifecyclePolicyAddSchedulesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/add_schedules", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add volumes to the policy if not present
func (r *CloudV1LifecyclePolicyService) AddVolumes(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, body CloudV1LifecyclePolicyAddVolumesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/add_volumes_to_policy", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get maximum usage quota of resources if all snapshots create by the policy
func (r *CloudV1LifecyclePolicyService) EstimateMaxUsage(ctx context.Context, projectID int64, regionID int64, body CloudV1LifecyclePolicyEstimateMaxUsageParams, opts ...option.RequestOption) (res *CloudV1LifecyclePolicyEstimateMaxUsageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/estimate_max_policy_usage", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove schedules from the policy
func (r *CloudV1LifecyclePolicyService) RemoveSchedules(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, body CloudV1LifecyclePolicyRemoveSchedulesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/remove_schedules", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove volumes from the policy
func (r *CloudV1LifecyclePolicyService) RemoveVolumes(ctx context.Context, projectID int64, regionID int64, lifecyclePolicyID int64, body CloudV1LifecyclePolicyRemoveVolumesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/remove_volumes_from_policy", projectID, regionID, lifecyclePolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CreateLifecyclePolicyParam struct {
	// Action that the policy will perform.
	Action param.Field[CreateLifecyclePolicyAction] `json:"action,required"`
	// Name of the lifecycle policy.
	Name param.Field[string] `json:"name,required"`
	// List of schedules associated with the policy.
	Schedules param.Field[[]CreateScheduleUnionParam] `json:"schedules"`
	// Current status of the lifecycle policy.
	Status param.Field[LifecyclePolicyStatuses] `json:"status"`
	// List of volume IDs.
	VolumeIDs param.Field[[]string] `json:"volume_ids" format:"uuid4"`
}

func (r CreateLifecyclePolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action that the policy will perform.
type CreateLifecyclePolicyAction string

const (
	CreateLifecyclePolicyActionVolumeSnapshot CreateLifecyclePolicyAction = "volume_snapshot"
)

func (r CreateLifecyclePolicyAction) IsKnown() bool {
	switch r {
	case CreateLifecyclePolicyActionVolumeSnapshot:
		return true
	}
	return false
}

// List of schedules associated with the policy. If the schedule made
// 'max_quantity' snapshots, a new snapshot will be created in time, but the oldest
// from the
// schedule`s snapshots will be deleted. If 'retention_time' was set and the schedule`s
// snapshots will be delete after the specified period of time. In
// resource_name_template you can use such forms that will then be automatically
// filled in with values: '{volume_id}', '{schedule_id}', '{lifecycle_policy_id}',
// '{datetime_utc}', '{local_datetime}'.
type CreateScheduleParam struct {
	// Schedule type
	Type param.Field[CreateScheduleType] `json:"type,required"`
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day param.Field[string] `json:"day"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek param.Field[string] `json:"day_of_week"`
	// Number of days to wait
	Days param.Field[int64] `json:"days"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour param.Field[string] `json:"hour"`
	// Number of hours to wait
	Hours param.Field[int64] `json:"hours"`
	// Number of stored resources.
	MaxQuantity param.Field[int64] `json:"max_quantity"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute param.Field[string] `json:"minute"`
	// Number of minutes to wait
	Minutes param.Field[int64] `json:"minutes"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month param.Field[string] `json:"month"`
	// Template for resource names.
	ResourceNameTemplate param.Field[string] `json:"resource_name_template"`
	// Time after which the resource will be deleted
	RetentionTime param.Field[CreateIntervalTimeParam] `json:"retention_time"`
	// A pytz timezone. Defaults to UTC.
	Timezone param.Field[string] `json:"timezone"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Field[string] `json:"week"`
	// Number of weeks to wait
	Weeks param.Field[int64] `json:"weeks"`
}

func (r CreateScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateScheduleParam) implementsCreateScheduleUnionParam() {}

// List of schedules associated with the policy. If the schedule made
// 'max_quantity' snapshots, a new snapshot will be created in time, but the oldest
// from the
// schedule`s snapshots will be deleted. If 'retention_time' was set and the schedule`s
// snapshots will be delete after the specified period of time. In
// resource_name_template you can use such forms that will then be automatically
// filled in with values: '{volume_id}', '{schedule_id}', '{lifecycle_policy_id}',
// '{datetime_utc}', '{local_datetime}'.
//
// Satisfied by [CreateScheduleCreateCronScheduleSerializerParam],
// [CreateScheduleCreateIntervalScheduleSerializerParam], [CreateScheduleParam].
type CreateScheduleUnionParam interface {
	implementsCreateScheduleUnionParam()
}

type CreateScheduleCreateCronScheduleSerializerParam struct {
	// Schedule type
	Type param.Field[CreateScheduleCreateCronScheduleSerializerType] `json:"type,required"`
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day param.Field[string] `json:"day"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek param.Field[string] `json:"day_of_week"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour param.Field[string] `json:"hour"`
	// Number of stored resources.
	MaxQuantity param.Field[int64] `json:"max_quantity"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute param.Field[string] `json:"minute"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month param.Field[string] `json:"month"`
	// Template for resource names.
	ResourceNameTemplate param.Field[string] `json:"resource_name_template"`
	// Time after which the resource will be deleted
	RetentionTime param.Field[CreateIntervalTimeParam] `json:"retention_time"`
	// A pytz timezone. Defaults to UTC.
	Timezone param.Field[string] `json:"timezone"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Field[string] `json:"week"`
}

func (r CreateScheduleCreateCronScheduleSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateScheduleCreateCronScheduleSerializerParam) implementsCreateScheduleUnionParam() {}

// Schedule type
type CreateScheduleCreateCronScheduleSerializerType string

const (
	CreateScheduleCreateCronScheduleSerializerTypeCron CreateScheduleCreateCronScheduleSerializerType = "cron"
)

func (r CreateScheduleCreateCronScheduleSerializerType) IsKnown() bool {
	switch r {
	case CreateScheduleCreateCronScheduleSerializerTypeCron:
		return true
	}
	return false
}

type CreateScheduleCreateIntervalScheduleSerializerParam struct {
	// Schedule type
	Type param.Field[CreateScheduleCreateIntervalScheduleSerializerType] `json:"type,required"`
	// Number of days to wait
	Days param.Field[int64] `json:"days"`
	// Number of hours to wait
	Hours param.Field[int64] `json:"hours"`
	// Number of stored resources.
	MaxQuantity param.Field[int64] `json:"max_quantity"`
	// Number of minutes to wait
	Minutes param.Field[int64] `json:"minutes"`
	// Template for resource names.
	ResourceNameTemplate param.Field[string] `json:"resource_name_template"`
	// Time after which the resource will be deleted
	RetentionTime param.Field[CreateIntervalTimeParam] `json:"retention_time"`
	// Number of weeks to wait
	Weeks param.Field[int64] `json:"weeks"`
}

func (r CreateScheduleCreateIntervalScheduleSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateScheduleCreateIntervalScheduleSerializerParam) implementsCreateScheduleUnionParam() {}

// Schedule type
type CreateScheduleCreateIntervalScheduleSerializerType string

const (
	CreateScheduleCreateIntervalScheduleSerializerTypeInterval CreateScheduleCreateIntervalScheduleSerializerType = "interval"
)

func (r CreateScheduleCreateIntervalScheduleSerializerType) IsKnown() bool {
	switch r {
	case CreateScheduleCreateIntervalScheduleSerializerTypeInterval:
		return true
	}
	return false
}

// Schedule type
type CreateScheduleType string

const (
	CreateScheduleTypeCron     CreateScheduleType = "cron"
	CreateScheduleTypeInterval CreateScheduleType = "interval"
)

func (r CreateScheduleType) IsKnown() bool {
	switch r {
	case CreateScheduleTypeCron, CreateScheduleTypeInterval:
		return true
	}
	return false
}

type LifecyclePolicy struct {
	// Unique identifier for the policy.
	ID int64 `json:"id,required"`
	// Action associated with the lifecycle policy.
	Action string `json:"action,required"`
	// Name of the policy.
	Name string `json:"name,required"`
	// Project ID associated with the policy.
	ProjectID int64 `json:"project_id,required"`
	// Region ID where the policy is applied.
	RegionID int64 `json:"region_id,required"`
	// List of schedules within the policy.
	Schedules []Schedule `json:"schedules,required"`
	// Status of the lifecycle policy.
	Status LifecyclePolicyStatuses `json:"status,required"`
	// User ID of the creator of the policy.
	UserID int64 `json:"user_id,required"`
	// Data of volumes that should be reserved. Displayed only when the query parameter
	// is specified.
	Volumes []LifecyclePolicyVolume `json:"volumes,required"`
	JSON    lifecyclePolicyJSON     `json:"-"`
}

// lifecyclePolicyJSON contains the JSON metadata for the struct [LifecyclePolicy]
type lifecyclePolicyJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	RegionID    apijson.Field
	Schedules   apijson.Field
	Status      apijson.Field
	UserID      apijson.Field
	Volumes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LifecyclePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lifecyclePolicyJSON) RawJSON() string {
	return r.raw
}

type LifecyclePolicyVolume struct {
	// Unique identifier of the volume.
	VolumeID string `json:"volume_id,required"`
	// Name of the volume.
	VolumeName string                    `json:"volume_name,required"`
	JSON       lifecyclePolicyVolumeJSON `json:"-"`
}

// lifecyclePolicyVolumeJSON contains the JSON metadata for the struct
// [LifecyclePolicyVolume]
type lifecyclePolicyVolumeJSON struct {
	VolumeID    apijson.Field
	VolumeName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LifecyclePolicyVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lifecyclePolicyVolumeJSON) RawJSON() string {
	return r.raw
}

type LifecyclePolicyStatuses string

const (
	LifecyclePolicyStatusesActive LifecyclePolicyStatuses = "active"
	LifecyclePolicyStatusesPaused LifecyclePolicyStatuses = "paused"
)

func (r LifecyclePolicyStatuses) IsKnown() bool {
	switch r {
	case LifecyclePolicyStatusesActive, LifecyclePolicyStatusesPaused:
		return true
	}
	return false
}

type VolumeIDsParam struct {
	// List of volume IDs.
	VolumeIDs param.Field[[]string] `json:"volume_ids,required"`
}

func (r VolumeIDsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LifecyclePolicyListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LifecyclePolicy                      `json:"results,required"`
	JSON    cloudV1LifecyclePolicyListResponseJSON `json:"-"`
}

// cloudV1LifecyclePolicyListResponseJSON contains the JSON metadata for the struct
// [CloudV1LifecyclePolicyListResponse]
type cloudV1LifecyclePolicyListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LifecyclePolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LifecyclePolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LifecyclePolicyEstimateMaxUsageResponse struct {
	// Total billed cost of all snapshots that can be created by the schedule. Cost of
	// max_volume_snapshot_count_usage snapshots.
	MaxCost ItemPriceResponse `json:"max_cost,required"`
	// Count of snapshots that can be created if the schedule creates the maximum
	// possible number of snapshots.
	MaxVolumeSnapshotCountUsage int64 `json:"max_volume_snapshot_count_usage,required"`
	// Maximum volume snapshot sequence length.
	MaxVolumeSnapshotSequenceLength int64 `json:"max_volume_snapshot_sequence_length,required"`
	// The amount of memory in GiB that snapshots will take up if the schedule creates
	// the maximum possible number of them.
	MaxVolumeSnapshotSizeUsage int64                                              `json:"max_volume_snapshot_size_usage,required"`
	JSON                       cloudV1LifecyclePolicyEstimateMaxUsageResponseJSON `json:"-"`
}

// cloudV1LifecyclePolicyEstimateMaxUsageResponseJSON contains the JSON metadata
// for the struct [CloudV1LifecyclePolicyEstimateMaxUsageResponse]
type cloudV1LifecyclePolicyEstimateMaxUsageResponseJSON struct {
	MaxCost                         apijson.Field
	MaxVolumeSnapshotCountUsage     apijson.Field
	MaxVolumeSnapshotSequenceLength apijson.Field
	MaxVolumeSnapshotSizeUsage      apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CloudV1LifecyclePolicyEstimateMaxUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LifecyclePolicyEstimateMaxUsageResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LifecyclePolicyNewParams struct {
	CreateLifecyclePolicy CreateLifecyclePolicyParam `json:"create_lifecycle_policy,required"`
}

func (r CloudV1LifecyclePolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLifecyclePolicy)
}

type CloudV1LifecyclePolicyUpdateParams struct {
	// Name of the lifecycle policy.
	Name param.Field[string] `json:"name"`
	// Status of the lifecycle policy.
	Status param.Field[LifecyclePolicyStatuses] `json:"status"`
}

func (r CloudV1LifecyclePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LifecyclePolicyAddSchedulesParams struct {
	// List of schedules associated with the policy.
	Schedules param.Field[[]CreateScheduleUnionParam] `json:"schedules,required"`
}

func (r CloudV1LifecyclePolicyAddSchedulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LifecyclePolicyAddVolumesParams struct {
	VolumeIDs VolumeIDsParam `json:"volume_ids,required"`
}

func (r CloudV1LifecyclePolicyAddVolumesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VolumeIDs)
}

type CloudV1LifecyclePolicyEstimateMaxUsageParams struct {
	CreateLifecyclePolicy CreateLifecyclePolicyParam `json:"create_lifecycle_policy,required"`
}

func (r CloudV1LifecyclePolicyEstimateMaxUsageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLifecyclePolicy)
}

type CloudV1LifecyclePolicyRemoveSchedulesParams struct {
	// List of schedule IDs.
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required"`
}

func (r CloudV1LifecyclePolicyRemoveSchedulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LifecyclePolicyRemoveVolumesParams struct {
	VolumeIDs VolumeIDsParam `json:"volume_ids,required"`
}

func (r CloudV1LifecyclePolicyRemoveVolumesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VolumeIDs)
}
