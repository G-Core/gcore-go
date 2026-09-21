// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/paramutil"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// Snapshot schedule policies describe when volume snapshots are taken and which
// volumes they cover. Volume membership is owned by the policy: attach and detach
// are policy-side operations, so a volume can join or leave a policy without being
// recreated.
//
// SnapshotScheduleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapshotScheduleService] method instead.
type SnapshotScheduleService struct {
	Options []option.RequestOption
}

// NewSnapshotScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSnapshotScheduleService(opts ...option.RequestOption) (r SnapshotScheduleService) {
	r = SnapshotScheduleService{}
	r.Options = opts
	return
}

// Create a new snapshot policy with the specified configuration.
func (r *SnapshotScheduleService) New(ctx context.Context, params SnapshotScheduleNewParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update the configuration of an existing snapshot policy.
func (r *SnapshotScheduleService) Update(ctx context.Context, policyID int64, params SnapshotScheduleUpdateParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", params.ProjectID.Value, params.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List all snapshot policies in the specified project and region.
func (r *SnapshotScheduleService) List(ctx context.Context, query SnapshotScheduleListParams, opts ...option.RequestOption) (res *SnapshotScheduleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a specific snapshot policy and all its associated schedules.
func (r *SnapshotScheduleService) Delete(ctx context.Context, policyID int64, body SnapshotScheduleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return err
	}
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", body.ProjectID.Value, body.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Add new schedules to an existing snapshot policy.
func (r *SnapshotScheduleService) AddSchedules(ctx context.Context, policyID int64, params SnapshotScheduleAddSchedulesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/add_schedules", params.ProjectID.Value, params.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Add volumes to an existing snapshot policy.
func (r *SnapshotScheduleService) AddVolumes(ctx context.Context, policyID int64, params SnapshotScheduleAddVolumesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/add_volumes_to_policy", params.ProjectID.Value, params.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Calculate the maximum resource usage if all snapshots are created by the policy.
func (r *SnapshotScheduleService) EstimateMaxUsage(ctx context.Context, params SnapshotScheduleEstimateMaxUsageParams, opts ...option.RequestOption) (res *SnapshotScheduleEstimateMaxUsageResponse, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/estimate_max_policy_usage", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get detailed information about a specific snapshot policy.
func (r *SnapshotScheduleService) Get(ctx context.Context, policyID int64, query SnapshotScheduleGetParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v", query.ProjectID.Value, query.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove schedules from an existing snapshot policy.
func (r *SnapshotScheduleService) RemoveSchedules(ctx context.Context, policyID int64, params SnapshotScheduleRemoveSchedulesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/remove_schedules", params.ProjectID.Value, params.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Remove volumes from an existing snapshot policy.
func (r *SnapshotScheduleService) RemoveVolumes(ctx context.Context, policyID int64, params SnapshotScheduleRemoveVolumesParams, opts ...option.RequestOption) (res *LifecyclePolicy, err error) {
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
	path := fmt.Sprintf("cloud/v1/lifecycle_policy/%v/%v/%v/remove_volumes_from_policy", params.ProjectID.Value, params.RegionID.Value, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type LifecyclePolicy struct {
	// Unique identifier for the policy.
	ID int64 `json:"id" api:"required"`
	// Action associated with the lifecycle policy.
	Action string `json:"action" api:"required"`
	// Name of the policy.
	Name string `json:"name" api:"required"`
	// Project ID associated with the policy.
	ProjectID int64 `json:"project_id" api:"required"`
	// Region ID where the policy is applied.
	RegionID int64 `json:"region_id" api:"required"`
	// List of schedules within the policy.
	Schedules []LifecyclePolicyScheduleUnion `json:"schedules" api:"required"`
	// Status of the lifecycle policy.
	//
	// Any of "active", "paused".
	Status LifecyclePolicyStatus `json:"status" api:"required"`
	// User ID of the creator of the policy.
	UserID int64 `json:"user_id" api:"required"`
	// Data of volumes that should be reserved. Displayed only when the query parameter
	// is specified.
	Volumes []LifecyclePolicyVolume `json:"volumes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		RegionID    respjson.Field
		Schedules   respjson.Field
		Status      respjson.Field
		UserID      respjson.Field
		Volumes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicy) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LifecyclePolicyScheduleUnion contains all possible properties and values from
// [LifecyclePolicyScheduleCron], [LifecyclePolicyScheduleInterval].
//
// Use the [LifecyclePolicyScheduleUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LifecyclePolicyScheduleUnion struct {
	ID          string `json:"id"`
	MaxQuantity int64  `json:"max_quantity"`
	Owner       string `json:"owner"`
	OwnerID     int64  `json:"owner_id"`
	// This field is a union of [LifecyclePolicyScheduleCronRetentionTime],
	// [LifecyclePolicyScheduleIntervalRetentionTime]
	RetentionTime LifecyclePolicyScheduleUnionRetentionTime `json:"retention_time"`
	// Any of "cron", "interval".
	Type   string `json:"type"`
	UserID int64  `json:"user_id"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Day string `json:"day"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	DayOfWeek string `json:"day_of_week"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Hour string `json:"hour"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Minute string `json:"minute"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Month                string `json:"month"`
	ResourceNameTemplate string `json:"resource_name_template"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Timezone string `json:"timezone"`
	// This field is from variant [LifecyclePolicyScheduleCron].
	Week string `json:"week"`
	// This field is from variant [LifecyclePolicyScheduleInterval].
	Days int64 `json:"days"`
	// This field is from variant [LifecyclePolicyScheduleInterval].
	Hours int64 `json:"hours"`
	// This field is from variant [LifecyclePolicyScheduleInterval].
	Minutes int64 `json:"minutes"`
	// This field is from variant [LifecyclePolicyScheduleInterval].
	Weeks int64 `json:"weeks"`
	JSON  struct {
		ID                   respjson.Field
		MaxQuantity          respjson.Field
		Owner                respjson.Field
		OwnerID              respjson.Field
		RetentionTime        respjson.Field
		Type                 respjson.Field
		UserID               respjson.Field
		Day                  respjson.Field
		DayOfWeek            respjson.Field
		Hour                 respjson.Field
		Minute               respjson.Field
		Month                respjson.Field
		ResourceNameTemplate respjson.Field
		Timezone             respjson.Field
		Week                 respjson.Field
		Days                 respjson.Field
		Hours                respjson.Field
		Minutes              respjson.Field
		Weeks                respjson.Field
		raw                  string
	} `json:"-"`
}

// anyLifecyclePolicySchedule is implemented by each variant of
// [LifecyclePolicyScheduleUnion] to add type safety for the return type of
// [LifecyclePolicyScheduleUnion.AsAny]
type anyLifecyclePolicySchedule interface {
	implLifecyclePolicyScheduleUnion()
}

func (LifecyclePolicyScheduleCron) implLifecyclePolicyScheduleUnion()     {}
func (LifecyclePolicyScheduleInterval) implLifecyclePolicyScheduleUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := LifecyclePolicyScheduleUnion.AsAny().(type) {
//	case cloud.LifecyclePolicyScheduleCron:
//	case cloud.LifecyclePolicyScheduleInterval:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LifecyclePolicyScheduleUnion) AsAny() anyLifecyclePolicySchedule {
	switch u.Type {
	case "cron":
		return u.AsCron()
	case "interval":
		return u.AsInterval()
	}
	return nil
}

func (u LifecyclePolicyScheduleUnion) AsCron() (v LifecyclePolicyScheduleCron) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LifecyclePolicyScheduleUnion) AsInterval() (v LifecyclePolicyScheduleInterval) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LifecyclePolicyScheduleUnion) RawJSON() string { return u.JSON.raw }

func (r *LifecyclePolicyScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LifecyclePolicyScheduleUnionRetentionTime is an implicit subunion of
// [LifecyclePolicyScheduleUnion]. LifecyclePolicyScheduleUnionRetentionTime
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LifecyclePolicyScheduleUnion].
type LifecyclePolicyScheduleUnionRetentionTime struct {
	Days    int64 `json:"days"`
	Hours   int64 `json:"hours"`
	Minutes int64 `json:"minutes"`
	Weeks   int64 `json:"weeks"`
	JSON    struct {
		Days    respjson.Field
		Hours   respjson.Field
		Minutes respjson.Field
		Weeks   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *LifecyclePolicyScheduleUnionRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecyclePolicyScheduleCron struct {
	// Schedule ID
	ID string `json:"id" api:"required"`
	// Number of stored resources.
	MaxQuantity int64 `json:"max_quantity" api:"required"`
	// Schedule owner
	Owner string `json:"owner" api:"required"`
	// Owner ID
	OwnerID int64 `json:"owner_id" api:"required"`
	// Time after which the resource will be deleted
	RetentionTime LifecyclePolicyScheduleCronRetentionTime `json:"retention_time" api:"required"`
	// Schedule type
	Type constant.Cron `json:"type" default:"cron"`
	// User ID
	UserID int64 `json:"user_id" api:"required"`
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day string `json:"day" api:"nullable"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek string `json:"day_of_week" api:"nullable"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour string `json:"hour" api:"nullable"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute string `json:"minute" api:"nullable"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month string `json:"month" api:"nullable"`
	// Template for resource names
	ResourceNameTemplate string `json:"resource_name_template" api:"nullable"`
	// A pytz timezone. Defaults to UTC.
	Timezone string `json:"timezone" api:"nullable"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week string `json:"week" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		MaxQuantity          respjson.Field
		Owner                respjson.Field
		OwnerID              respjson.Field
		RetentionTime        respjson.Field
		Type                 respjson.Field
		UserID               respjson.Field
		Day                  respjson.Field
		DayOfWeek            respjson.Field
		Hour                 respjson.Field
		Minute               respjson.Field
		Month                respjson.Field
		ResourceNameTemplate respjson.Field
		Timezone             respjson.Field
		Week                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicyScheduleCron) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicyScheduleCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type LifecyclePolicyScheduleCronRetentionTime struct {
	// Number of days to wait
	Days int64 `json:"days" api:"nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours" api:"nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes" api:"nullable"`
	// Number of weeks to wait
	Weeks int64 `json:"weeks" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days        respjson.Field
		Hours       respjson.Field
		Minutes     respjson.Field
		Weeks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicyScheduleCronRetentionTime) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicyScheduleCronRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecyclePolicyScheduleInterval struct {
	// Schedule ID
	ID string `json:"id" api:"required"`
	// Number of stored resources.
	MaxQuantity int64 `json:"max_quantity" api:"required"`
	// Schedule owner
	Owner string `json:"owner" api:"required"`
	// Owner ID
	OwnerID int64 `json:"owner_id" api:"required"`
	// Time after which the resource will be deleted
	RetentionTime LifecyclePolicyScheduleIntervalRetentionTime `json:"retention_time" api:"required"`
	// Schedule type
	Type constant.Interval `json:"type" default:"interval"`
	// User ID
	UserID int64 `json:"user_id" api:"required"`
	// Number of days to wait
	Days int64 `json:"days" api:"nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours" api:"nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes" api:"nullable"`
	// Template for resource names
	ResourceNameTemplate string `json:"resource_name_template" api:"nullable"`
	// Number of weeks to wait
	Weeks int64 `json:"weeks" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		MaxQuantity          respjson.Field
		Owner                respjson.Field
		OwnerID              respjson.Field
		RetentionTime        respjson.Field
		Type                 respjson.Field
		UserID               respjson.Field
		Days                 respjson.Field
		Hours                respjson.Field
		Minutes              respjson.Field
		ResourceNameTemplate respjson.Field
		Weeks                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicyScheduleInterval) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicyScheduleInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type LifecyclePolicyScheduleIntervalRetentionTime struct {
	// Number of days to wait
	Days int64 `json:"days" api:"nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours" api:"nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes" api:"nullable"`
	// Number of weeks to wait
	Weeks int64 `json:"weeks" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days        respjson.Field
		Hours       respjson.Field
		Minutes     respjson.Field
		Weeks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicyScheduleIntervalRetentionTime) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicyScheduleIntervalRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the lifecycle policy.
type LifecyclePolicyStatus string

const (
	LifecyclePolicyStatusActive LifecyclePolicyStatus = "active"
	LifecyclePolicyStatusPaused LifecyclePolicyStatus = "paused"
)

type LifecyclePolicyVolume struct {
	// Unique identifier of the volume.
	VolumeID string `json:"volume_id" api:"required"`
	// Name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VolumeID    respjson.Field
		VolumeName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecyclePolicyVolume) RawJSON() string { return r.JSON.raw }
func (r *LifecyclePolicyVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleListResponse struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []LifecyclePolicy `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotScheduleListResponse) RawJSON() string { return r.JSON.raw }
func (r *SnapshotScheduleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleEstimateMaxUsageResponse struct {
	// Total billed cost of all snapshots that can be created by the schedule. Cost of
	// `max_volume_snapshot_count_usage` snapshots.
	MaxCost SnapshotScheduleEstimateMaxUsageResponseMaxCost `json:"max_cost" api:"required"`
	// Count of snapshots that can be created if the schedule creates the maximum
	// possible number of snapshots.
	MaxVolumeSnapshotCountUsage int64 `json:"max_volume_snapshot_count_usage" api:"required"`
	// Maximum volume snapshot sequence length.
	MaxVolumeSnapshotSequenceLength int64 `json:"max_volume_snapshot_sequence_length" api:"required"`
	// The amount of memory in GiB that snapshots will take up if the schedule creates
	// the maximum possible number of them.
	MaxVolumeSnapshotSizeUsage int64 `json:"max_volume_snapshot_size_usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxCost                         respjson.Field
		MaxVolumeSnapshotCountUsage     respjson.Field
		MaxVolumeSnapshotSequenceLength respjson.Field
		MaxVolumeSnapshotSizeUsage      respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotScheduleEstimateMaxUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *SnapshotScheduleEstimateMaxUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Total billed cost of all snapshots that can be created by the schedule. Cost of
// `max_volume_snapshot_count_usage` snapshots.
type SnapshotScheduleEstimateMaxUsageResponseMaxCost struct {
	// Currency code (3 letter code per ISO 4217)
	//
	// Any of "AZN", "EUR", "USD".
	CurrencyCode string `json:"currency_code" api:"required"`
	// Actual discount relative value
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price of the item charged per hour
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price of the item charged per month
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status" api:"required"`
	// Total price VAT inclusive per month without discount
	PriceWithoutDiscountPerMonth float64 `json:"price_without_discount_per_month" api:"required"`
	// Tax rate applied to the subtotal, represented as a percentage
	TaxPercent float64 `json:"tax_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode                 respjson.Field
		DiscountPercent              respjson.Field
		PricePerHour                 respjson.Field
		PricePerMonth                respjson.Field
		PriceStatus                  respjson.Field
		PriceWithoutDiscountPerMonth respjson.Field
		TaxPercent                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotScheduleEstimateMaxUsageResponseMaxCost) RawJSON() string { return r.JSON.raw }
func (r *SnapshotScheduleEstimateMaxUsageResponseMaxCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Action that the policy will perform.
	//
	// Any of "volume_snapshot".
	Action SnapshotScheduleNewParamsAction `json:"action,omitzero" api:"required"`
	// Name of the lifecycle policy.
	Name string `json:"name" api:"required"`
	// List of schedules associated with the policy.
	Schedules []SnapshotScheduleNewParamsScheduleUnion `json:"schedules,omitzero"`
	// Current status of the lifecycle policy.
	//
	// Any of "active", "paused".
	Status SnapshotScheduleNewParamsStatus `json:"status,omitzero"`
	// List of volume IDs.
	VolumeIDs []string `json:"volume_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r SnapshotScheduleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action that the policy will perform.
type SnapshotScheduleNewParamsAction string

const (
	SnapshotScheduleNewParamsActionVolumeSnapshot SnapshotScheduleNewParamsAction = "volume_snapshot"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SnapshotScheduleNewParamsScheduleUnion struct {
	OfCron     *SnapshotScheduleNewParamsScheduleCron     `json:",omitzero,inline"`
	OfInterval *SnapshotScheduleNewParamsScheduleInterval `json:",omitzero,inline"`
	paramUnion
}

func (u SnapshotScheduleNewParamsScheduleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCron, u.OfInterval)
}
func (u *SnapshotScheduleNewParamsScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SnapshotScheduleNewParamsScheduleUnion) asAny() any {
	if !param.IsOmitted(u.OfCron) {
		return u.OfCron
	} else if !param.IsOmitted(u.OfInterval) {
		return u.OfInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetDay() *string {
	if vt := u.OfCron; vt != nil && vt.Day.Valid() {
		return &vt.Day.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetDayOfWeek() *string {
	if vt := u.OfCron; vt != nil && vt.DayOfWeek.Valid() {
		return &vt.DayOfWeek.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetHour() *string {
	if vt := u.OfCron; vt != nil && vt.Hour.Valid() {
		return &vt.Hour.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetMinute() *string {
	if vt := u.OfCron; vt != nil && vt.Minute.Valid() {
		return &vt.Minute.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetMonth() *string {
	if vt := u.OfCron; vt != nil && vt.Month.Valid() {
		return &vt.Month.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetTimezone() *string {
	if vt := u.OfCron; vt != nil && vt.Timezone.Valid() {
		return &vt.Timezone.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetWeek() *string {
	if vt := u.OfCron; vt != nil && vt.Week.Valid() {
		return &vt.Week.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetDays() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Days.Valid() {
		return &vt.Days.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetHours() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Hours.Valid() {
		return &vt.Hours.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetMinutes() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Minutes.Valid() {
		return &vt.Minutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetWeeks() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Weeks.Valid() {
		return &vt.Weeks.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetType() *string {
	if vt := u.OfCron; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInterval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetMaxQuantity() *int64 {
	if vt := u.OfCron; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	} else if vt := u.OfInterval; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleNewParamsScheduleUnion) GetResourceNameTemplate() *string {
	if vt := u.OfCron; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	} else if vt := u.OfInterval; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SnapshotScheduleNewParamsScheduleUnion) GetRetentionTime() (res snapshotScheduleNewParamsScheduleUnionRetentionTime) {
	if vt := u.OfCron; vt != nil {
		res.any = &vt.RetentionTime
	} else if vt := u.OfInterval; vt != nil {
		res.any = &vt.RetentionTime
	}
	return
}

// Can have the runtime types
// [*SnapshotScheduleNewParamsScheduleCronRetentionTime],
// [*SnapshotScheduleNewParamsScheduleIntervalRetentionTime]
type snapshotScheduleNewParamsScheduleUnionRetentionTime struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.SnapshotScheduleNewParamsScheduleCronRetentionTime:
//	case *cloud.SnapshotScheduleNewParamsScheduleIntervalRetentionTime:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u snapshotScheduleNewParamsScheduleUnionRetentionTime) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleNewParamsScheduleUnionRetentionTime) GetDays() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleNewParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	case *SnapshotScheduleNewParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleNewParamsScheduleUnionRetentionTime) GetHours() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleNewParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	case *SnapshotScheduleNewParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleNewParamsScheduleUnionRetentionTime) GetMinutes() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleNewParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	case *SnapshotScheduleNewParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleNewParamsScheduleUnionRetentionTime) GetWeeks() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleNewParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	case *SnapshotScheduleNewParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SnapshotScheduleNewParamsScheduleUnion](
		"type",
		apijson.Discriminator[SnapshotScheduleNewParamsScheduleCron]("cron"),
		apijson.Discriminator[SnapshotScheduleNewParamsScheduleInterval]("interval"),
	)
}

// The property Type is required.
type SnapshotScheduleNewParamsScheduleCron struct {
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day param.Opt[string] `json:"day,omitzero"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek param.Opt[string] `json:"day_of_week,omitzero"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour param.Opt[string] `json:"hour,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute param.Opt[string] `json:"minute,omitzero"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month param.Opt[string] `json:"month,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// A pytz timezone. Defaults to UTC.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Opt[string] `json:"week,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleNewParamsScheduleCronRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "cron".
	Type constant.Cron `json:"type" default:"cron"`
	paramObj
}

func (r SnapshotScheduleNewParamsScheduleCron) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleNewParamsScheduleCron
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleNewParamsScheduleCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleNewParamsScheduleCronRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleNewParamsScheduleCronRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleNewParamsScheduleCronRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleNewParamsScheduleCronRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type SnapshotScheduleNewParamsScheduleInterval struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleNewParamsScheduleIntervalRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "interval".
	Type constant.Interval `json:"type" default:"interval"`
	paramObj
}

func (r SnapshotScheduleNewParamsScheduleInterval) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleNewParamsScheduleInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleNewParamsScheduleInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleNewParamsScheduleIntervalRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleNewParamsScheduleIntervalRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleNewParamsScheduleIntervalRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleNewParamsScheduleIntervalRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the lifecycle policy.
type SnapshotScheduleNewParamsStatus string

const (
	SnapshotScheduleNewParamsStatusActive SnapshotScheduleNewParamsStatus = "active"
	SnapshotScheduleNewParamsStatusPaused SnapshotScheduleNewParamsStatus = "paused"
)

type SnapshotScheduleUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Name of the lifecycle policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Status of the lifecycle policy.
	//
	// Any of "active", "paused".
	Status SnapshotScheduleUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r SnapshotScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the lifecycle policy.
type SnapshotScheduleUpdateParamsStatus string

const (
	SnapshotScheduleUpdateParamsStatusActive SnapshotScheduleUpdateParamsStatus = "active"
	SnapshotScheduleUpdateParamsStatusPaused SnapshotScheduleUpdateParamsStatus = "paused"
)

type SnapshotScheduleListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type SnapshotScheduleDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type SnapshotScheduleAddSchedulesParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// List of schedules associated with the policy.
	Schedules []SnapshotScheduleAddSchedulesParamsScheduleUnion `json:"schedules,omitzero" api:"required"`
	paramObj
}

func (r SnapshotScheduleAddSchedulesParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddSchedulesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddSchedulesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SnapshotScheduleAddSchedulesParamsScheduleUnion struct {
	OfCron     *SnapshotScheduleAddSchedulesParamsScheduleCron     `json:",omitzero,inline"`
	OfInterval *SnapshotScheduleAddSchedulesParamsScheduleInterval `json:",omitzero,inline"`
	paramUnion
}

func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCron, u.OfInterval)
}
func (u *SnapshotScheduleAddSchedulesParamsScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SnapshotScheduleAddSchedulesParamsScheduleUnion) asAny() any {
	if !param.IsOmitted(u.OfCron) {
		return u.OfCron
	} else if !param.IsOmitted(u.OfInterval) {
		return u.OfInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetDay() *string {
	if vt := u.OfCron; vt != nil && vt.Day.Valid() {
		return &vt.Day.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetDayOfWeek() *string {
	if vt := u.OfCron; vt != nil && vt.DayOfWeek.Valid() {
		return &vt.DayOfWeek.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetHour() *string {
	if vt := u.OfCron; vt != nil && vt.Hour.Valid() {
		return &vt.Hour.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetMinute() *string {
	if vt := u.OfCron; vt != nil && vt.Minute.Valid() {
		return &vt.Minute.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetMonth() *string {
	if vt := u.OfCron; vt != nil && vt.Month.Valid() {
		return &vt.Month.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetTimezone() *string {
	if vt := u.OfCron; vt != nil && vt.Timezone.Valid() {
		return &vt.Timezone.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetWeek() *string {
	if vt := u.OfCron; vt != nil && vt.Week.Valid() {
		return &vt.Week.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetDays() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Days.Valid() {
		return &vt.Days.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetHours() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Hours.Valid() {
		return &vt.Hours.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetMinutes() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Minutes.Valid() {
		return &vt.Minutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetWeeks() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Weeks.Valid() {
		return &vt.Weeks.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetType() *string {
	if vt := u.OfCron; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInterval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetMaxQuantity() *int64 {
	if vt := u.OfCron; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	} else if vt := u.OfInterval; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetResourceNameTemplate() *string {
	if vt := u.OfCron; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	} else if vt := u.OfInterval; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SnapshotScheduleAddSchedulesParamsScheduleUnion) GetRetentionTime() (res snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) {
	if vt := u.OfCron; vt != nil {
		res.any = &vt.RetentionTime
	} else if vt := u.OfInterval; vt != nil {
		res.any = &vt.RetentionTime
	}
	return
}

// Can have the runtime types
// [*SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime],
// [*SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime]
type snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime:
//	case *cloud.SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) GetDays() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	case *SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) GetHours() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	case *SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) GetMinutes() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	case *SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleAddSchedulesParamsScheduleUnionRetentionTime) GetWeeks() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	case *SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SnapshotScheduleAddSchedulesParamsScheduleUnion](
		"type",
		apijson.Discriminator[SnapshotScheduleAddSchedulesParamsScheduleCron]("cron"),
		apijson.Discriminator[SnapshotScheduleAddSchedulesParamsScheduleInterval]("interval"),
	)
}

// The property Type is required.
type SnapshotScheduleAddSchedulesParamsScheduleCron struct {
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day param.Opt[string] `json:"day,omitzero"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek param.Opt[string] `json:"day_of_week,omitzero"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour param.Opt[string] `json:"hour,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute param.Opt[string] `json:"minute,omitzero"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month param.Opt[string] `json:"month,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// A pytz timezone. Defaults to UTC.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Opt[string] `json:"week,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "cron".
	Type constant.Cron `json:"type" default:"cron"`
	paramObj
}

func (r SnapshotScheduleAddSchedulesParamsScheduleCron) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddSchedulesParamsScheduleCron
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddSchedulesParamsScheduleCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type SnapshotScheduleAddSchedulesParamsScheduleInterval struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "interval".
	Type constant.Interval `json:"type" default:"interval"`
	paramObj
}

func (r SnapshotScheduleAddSchedulesParamsScheduleInterval) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddSchedulesParamsScheduleInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddSchedulesParamsScheduleInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleAddVolumesParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// List of volume IDs.
	VolumeIDs []string `json:"volume_ids,omitzero" api:"required"`
	paramObj
}

func (r SnapshotScheduleAddVolumesParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleAddVolumesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleAddVolumesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleEstimateMaxUsageParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Action that the policy will perform.
	//
	// Any of "volume_snapshot".
	Action SnapshotScheduleEstimateMaxUsageParamsAction `json:"action,omitzero" api:"required"`
	// Name of the lifecycle policy.
	Name string `json:"name" api:"required"`
	// List of schedules associated with the policy.
	Schedules []SnapshotScheduleEstimateMaxUsageParamsScheduleUnion `json:"schedules,omitzero"`
	// Current status of the lifecycle policy.
	//
	// Any of "active", "paused".
	Status SnapshotScheduleEstimateMaxUsageParamsStatus `json:"status,omitzero"`
	// List of volume IDs.
	VolumeIDs []string `json:"volume_ids,omitzero" format:"uuid4"`
	paramObj
}

func (r SnapshotScheduleEstimateMaxUsageParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleEstimateMaxUsageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleEstimateMaxUsageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action that the policy will perform.
type SnapshotScheduleEstimateMaxUsageParamsAction string

const (
	SnapshotScheduleEstimateMaxUsageParamsActionVolumeSnapshot SnapshotScheduleEstimateMaxUsageParamsAction = "volume_snapshot"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SnapshotScheduleEstimateMaxUsageParamsScheduleUnion struct {
	OfCron     *SnapshotScheduleEstimateMaxUsageParamsScheduleCron     `json:",omitzero,inline"`
	OfInterval *SnapshotScheduleEstimateMaxUsageParamsScheduleInterval `json:",omitzero,inline"`
	paramUnion
}

func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCron, u.OfInterval)
}
func (u *SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) asAny() any {
	if !param.IsOmitted(u.OfCron) {
		return u.OfCron
	} else if !param.IsOmitted(u.OfInterval) {
		return u.OfInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetDay() *string {
	if vt := u.OfCron; vt != nil && vt.Day.Valid() {
		return &vt.Day.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetDayOfWeek() *string {
	if vt := u.OfCron; vt != nil && vt.DayOfWeek.Valid() {
		return &vt.DayOfWeek.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetHour() *string {
	if vt := u.OfCron; vt != nil && vt.Hour.Valid() {
		return &vt.Hour.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetMinute() *string {
	if vt := u.OfCron; vt != nil && vt.Minute.Valid() {
		return &vt.Minute.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetMonth() *string {
	if vt := u.OfCron; vt != nil && vt.Month.Valid() {
		return &vt.Month.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetTimezone() *string {
	if vt := u.OfCron; vt != nil && vt.Timezone.Valid() {
		return &vt.Timezone.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetWeek() *string {
	if vt := u.OfCron; vt != nil && vt.Week.Valid() {
		return &vt.Week.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetDays() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Days.Valid() {
		return &vt.Days.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetHours() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Hours.Valid() {
		return &vt.Hours.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetMinutes() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Minutes.Valid() {
		return &vt.Minutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetWeeks() *int64 {
	if vt := u.OfInterval; vt != nil && vt.Weeks.Valid() {
		return &vt.Weeks.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetType() *string {
	if vt := u.OfCron; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInterval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetMaxQuantity() *int64 {
	if vt := u.OfCron; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	} else if vt := u.OfInterval; vt != nil && vt.MaxQuantity.Valid() {
		return &vt.MaxQuantity.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetResourceNameTemplate() *string {
	if vt := u.OfCron; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	} else if vt := u.OfInterval; vt != nil && vt.ResourceNameTemplate.Valid() {
		return &vt.ResourceNameTemplate.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SnapshotScheduleEstimateMaxUsageParamsScheduleUnion) GetRetentionTime() (res snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) {
	if vt := u.OfCron; vt != nil {
		res.any = &vt.RetentionTime
	} else if vt := u.OfInterval; vt != nil {
		res.any = &vt.RetentionTime
	}
	return
}

// Can have the runtime types
// [*SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime],
// [*SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime]
type snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime:
//	case *cloud.SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) GetDays() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Days)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) GetHours() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Hours)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) GetMinutes() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Minutes)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u snapshotScheduleEstimateMaxUsageParamsScheduleUnionRetentionTime) GetWeeks() *int64 {
	switch vt := u.any.(type) {
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	case *SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime:
		return paramutil.AddrIfPresent(vt.Weeks)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SnapshotScheduleEstimateMaxUsageParamsScheduleUnion](
		"type",
		apijson.Discriminator[SnapshotScheduleEstimateMaxUsageParamsScheduleCron]("cron"),
		apijson.Discriminator[SnapshotScheduleEstimateMaxUsageParamsScheduleInterval]("interval"),
	)
}

// The property Type is required.
type SnapshotScheduleEstimateMaxUsageParamsScheduleCron struct {
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day param.Opt[string] `json:"day,omitzero"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek param.Opt[string] `json:"day_of_week,omitzero"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour param.Opt[string] `json:"hour,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute param.Opt[string] `json:"minute,omitzero"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month param.Opt[string] `json:"month,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// A pytz timezone. Defaults to UTC.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Opt[string] `json:"week,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "cron".
	Type constant.Cron `json:"type" default:"cron"`
	paramObj
}

func (r SnapshotScheduleEstimateMaxUsageParamsScheduleCron) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleEstimateMaxUsageParamsScheduleCron
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleEstimateMaxUsageParamsScheduleCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type SnapshotScheduleEstimateMaxUsageParamsScheduleInterval struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of stored resources.
	MaxQuantity param.Opt[int64] `json:"max_quantity,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Template for resource names.
	ResourceNameTemplate param.Opt[string] `json:"resource_name_template,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	// Time after which the resource will be deleted
	RetentionTime SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime `json:"retention_time,omitzero"`
	// Schedule type
	//
	// This field can be elided, and will marshal its zero value as "interval".
	Type constant.Interval `json:"type" default:"interval"`
	paramObj
}

func (r SnapshotScheduleEstimateMaxUsageParamsScheduleInterval) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleEstimateMaxUsageParamsScheduleInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleEstimateMaxUsageParamsScheduleInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time after which the resource will be deleted
type SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime struct {
	// Number of days to wait
	Days param.Opt[int64] `json:"days,omitzero"`
	// Number of hours to wait
	Hours param.Opt[int64] `json:"hours,omitzero"`
	// Number of minutes to wait
	Minutes param.Opt[int64] `json:"minutes,omitzero"`
	// Number of weeks to wait
	Weeks param.Opt[int64] `json:"weeks,omitzero"`
	paramObj
}

func (r SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleEstimateMaxUsageParamsScheduleIntervalRetentionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the lifecycle policy.
type SnapshotScheduleEstimateMaxUsageParamsStatus string

const (
	SnapshotScheduleEstimateMaxUsageParamsStatusActive SnapshotScheduleEstimateMaxUsageParamsStatus = "active"
	SnapshotScheduleEstimateMaxUsageParamsStatusPaused SnapshotScheduleEstimateMaxUsageParamsStatus = "paused"
)

type SnapshotScheduleGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type SnapshotScheduleRemoveSchedulesParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// List of schedule IDs.
	ScheduleIDs []string `json:"schedule_ids,omitzero" api:"required"`
	paramObj
}

func (r SnapshotScheduleRemoveSchedulesParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleRemoveSchedulesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleRemoveSchedulesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotScheduleRemoveVolumesParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// List of volume IDs.
	VolumeIDs []string `json:"volume_ids,omitzero" api:"required"`
	paramObj
}

func (r SnapshotScheduleRemoveVolumesParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleRemoveVolumesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleRemoveVolumesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
