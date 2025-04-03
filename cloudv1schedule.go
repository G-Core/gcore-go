// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/tidwall/gjson"
)

// CloudV1ScheduleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ScheduleService] method instead.
type CloudV1ScheduleService struct {
	Options []option.RequestOption
}

// NewCloudV1ScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1ScheduleService(opts ...option.RequestOption) (r *CloudV1ScheduleService) {
	r = &CloudV1ScheduleService{}
	r.Options = opts
	return
}

// Get a schedule by id
func (r *CloudV1ScheduleService) Get(ctx context.Context, projectID int64, regionID int64, scheduleID string, opts ...option.RequestOption) (res *Schedule, err error) {
	opts = append(r.Options[:], opts...)
	if scheduleID == "" {
		err = errors.New("missing required schedule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/schedule/%v/%v/%s", projectID, regionID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch the schedule
func (r *CloudV1ScheduleService) Update(ctx context.Context, projectID int64, regionID int64, scheduleID string, body CloudV1ScheduleUpdateParams, opts ...option.RequestOption) (res *Schedule, err error) {
	opts = append(r.Options[:], opts...)
	if scheduleID == "" {
		err = errors.New("missing required schedule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/schedule/%v/%v/%s", projectID, regionID, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CreateIntervalTimeParam struct {
	// Number of days to wait
	Days param.Field[int64] `json:"days"`
	// Number of hours to wait
	Hours param.Field[int64] `json:"hours"`
	// Number of minutes to wait
	Minutes param.Field[int64] `json:"minutes"`
	// Number of weeks to wait
	Weeks param.Field[int64] `json:"weeks"`
}

func (r CreateIntervalTimeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RetentionTime struct {
	// Number of days to wait
	Days int64 `json:"days,nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours,nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes,nullable"`
	// Number of weeks to wait
	Weeks int64             `json:"weeks,nullable"`
	JSON  retentionTimeJSON `json:"-"`
}

// retentionTimeJSON contains the JSON metadata for the struct [RetentionTime]
type retentionTimeJSON struct {
	Days        apijson.Field
	Hours       apijson.Field
	Minutes     apijson.Field
	Weeks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RetentionTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r retentionTimeJSON) RawJSON() string {
	return r.raw
}

type Schedule struct {
	// Schedule ID
	ID string `json:"id,required"`
	// Number of stored resources.
	MaxQuantity int64 `json:"max_quantity,required"`
	// Schedule owner
	Owner string `json:"owner,required"`
	// Owner ID
	OwnerID int64 `json:"owner_id,required"`
	// Time after which the resource will be deleted
	RetentionTime RetentionTime `json:"retention_time,required,nullable"`
	// Schedule type
	Type ScheduleType `json:"type,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day string `json:"day,nullable"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek string `json:"day_of_week,nullable"`
	// Number of days to wait
	Days int64 `json:"days,nullable"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour string `json:"hour,nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours,nullable"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute string `json:"minute,nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes,nullable"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month string `json:"month,nullable"`
	// Template for resource names
	ResourceNameTemplate string `json:"resource_name_template,nullable"`
	// A pytz timezone. Defaults to UTC.
	Timezone string `json:"timezone,nullable"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week string `json:"week,nullable"`
	// Number of weeks to wait
	Weeks int64        `json:"weeks,nullable"`
	JSON  scheduleJSON `json:"-"`
	union ScheduleUnion
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	ID                   apijson.Field
	MaxQuantity          apijson.Field
	Owner                apijson.Field
	OwnerID              apijson.Field
	RetentionTime        apijson.Field
	Type                 apijson.Field
	UserID               apijson.Field
	Day                  apijson.Field
	DayOfWeek            apijson.Field
	Days                 apijson.Field
	Hour                 apijson.Field
	Hours                apijson.Field
	Minute               apijson.Field
	Minutes              apijson.Field
	Month                apijson.Field
	ResourceNameTemplate apijson.Field
	Timezone             apijson.Field
	Week                 apijson.Field
	Weeks                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	*r = Schedule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ScheduleGetCronScheduleSerializer],
// [ScheduleGetIntervalScheduleSerializer].
func (r Schedule) AsUnion() ScheduleUnion {
	return r.union
}

// Union satisfied by [ScheduleGetCronScheduleSerializer] or
// [ScheduleGetIntervalScheduleSerializer].
type ScheduleUnion interface {
	implementsSchedule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScheduleGetCronScheduleSerializer{}),
			DiscriminatorValue: "cron",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScheduleGetIntervalScheduleSerializer{}),
			DiscriminatorValue: "interval",
		},
	)
}

type ScheduleGetCronScheduleSerializer struct {
	// Schedule ID
	ID string `json:"id,required"`
	// Number of stored resources.
	MaxQuantity int64 `json:"max_quantity,required"`
	// Schedule owner
	Owner string `json:"owner,required"`
	// Owner ID
	OwnerID int64 `json:"owner_id,required"`
	// Time after which the resource will be deleted
	RetentionTime RetentionTime `json:"retention_time,required,nullable"`
	// Schedule type
	Type ScheduleGetCronScheduleSerializerType `json:"type,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Day of the month (1-31, '\*') or a comma-separated list of days
	Day string `json:"day,nullable"`
	// Weekday or a comma-separated list of weekdays (mon,tue,wed,thu,fri,sat,sun,\*)
	DayOfWeek string `json:"day_of_week,nullable"`
	// Hour (0-23, '\*') or a comma-separated list of hours
	Hour string `json:"hour,nullable"`
	// Minute (0-59, '\*') or a comma-separated list of minutes
	Minute string `json:"minute,nullable"`
	// Month (1-12, '\*') or a comma-separated list of months
	Month string `json:"month,nullable"`
	// Template for resource names
	ResourceNameTemplate string `json:"resource_name_template,nullable"`
	// A pytz timezone. Defaults to UTC.
	Timezone string `json:"timezone,nullable"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week string                                `json:"week,nullable"`
	JSON scheduleGetCronScheduleSerializerJSON `json:"-"`
}

// scheduleGetCronScheduleSerializerJSON contains the JSON metadata for the struct
// [ScheduleGetCronScheduleSerializer]
type scheduleGetCronScheduleSerializerJSON struct {
	ID                   apijson.Field
	MaxQuantity          apijson.Field
	Owner                apijson.Field
	OwnerID              apijson.Field
	RetentionTime        apijson.Field
	Type                 apijson.Field
	UserID               apijson.Field
	Day                  apijson.Field
	DayOfWeek            apijson.Field
	Hour                 apijson.Field
	Minute               apijson.Field
	Month                apijson.Field
	ResourceNameTemplate apijson.Field
	Timezone             apijson.Field
	Week                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ScheduleGetCronScheduleSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleGetCronScheduleSerializerJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleGetCronScheduleSerializer) implementsSchedule() {}

// Schedule type
type ScheduleGetCronScheduleSerializerType string

const (
	ScheduleGetCronScheduleSerializerTypeCron ScheduleGetCronScheduleSerializerType = "cron"
)

func (r ScheduleGetCronScheduleSerializerType) IsKnown() bool {
	switch r {
	case ScheduleGetCronScheduleSerializerTypeCron:
		return true
	}
	return false
}

type ScheduleGetIntervalScheduleSerializer struct {
	// Schedule ID
	ID string `json:"id,required"`
	// Number of stored resources.
	MaxQuantity int64 `json:"max_quantity,required"`
	// Schedule owner
	Owner string `json:"owner,required"`
	// Owner ID
	OwnerID int64 `json:"owner_id,required"`
	// Time after which the resource will be deleted
	RetentionTime RetentionTime `json:"retention_time,required,nullable"`
	// Schedule type
	Type ScheduleGetIntervalScheduleSerializerType `json:"type,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Number of days to wait
	Days int64 `json:"days,nullable"`
	// Number of hours to wait
	Hours int64 `json:"hours,nullable"`
	// Number of minutes to wait
	Minutes int64 `json:"minutes,nullable"`
	// Template for resource names
	ResourceNameTemplate string `json:"resource_name_template,nullable"`
	// Number of weeks to wait
	Weeks int64                                     `json:"weeks,nullable"`
	JSON  scheduleGetIntervalScheduleSerializerJSON `json:"-"`
}

// scheduleGetIntervalScheduleSerializerJSON contains the JSON metadata for the
// struct [ScheduleGetIntervalScheduleSerializer]
type scheduleGetIntervalScheduleSerializerJSON struct {
	ID                   apijson.Field
	MaxQuantity          apijson.Field
	Owner                apijson.Field
	OwnerID              apijson.Field
	RetentionTime        apijson.Field
	Type                 apijson.Field
	UserID               apijson.Field
	Days                 apijson.Field
	Hours                apijson.Field
	Minutes              apijson.Field
	ResourceNameTemplate apijson.Field
	Weeks                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ScheduleGetIntervalScheduleSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleGetIntervalScheduleSerializerJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleGetIntervalScheduleSerializer) implementsSchedule() {}

// Schedule type
type ScheduleGetIntervalScheduleSerializerType string

const (
	ScheduleGetIntervalScheduleSerializerTypeInterval ScheduleGetIntervalScheduleSerializerType = "interval"
)

func (r ScheduleGetIntervalScheduleSerializerType) IsKnown() bool {
	switch r {
	case ScheduleGetIntervalScheduleSerializerTypeInterval:
		return true
	}
	return false
}

// Schedule type
type ScheduleType string

const (
	ScheduleTypeCron     ScheduleType = "cron"
	ScheduleTypeInterval ScheduleType = "interval"
)

func (r ScheduleType) IsKnown() bool {
	switch r {
	case ScheduleTypeCron, ScheduleTypeInterval:
		return true
	}
	return false
}

type CloudV1ScheduleUpdateParams struct {
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
	// Type of the schedule.
	Type param.Field[CloudV1ScheduleUpdateParamsType] `json:"type"`
	// ISO week (1-53, '\*') or a comma-separated list of weeks
	Week param.Field[string] `json:"week"`
	// Number of weeks to wait
	Weeks param.Field[int64] `json:"weeks"`
}

func (r CloudV1ScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the schedule.
type CloudV1ScheduleUpdateParamsType string

const (
	CloudV1ScheduleUpdateParamsTypeCron     CloudV1ScheduleUpdateParamsType = "cron"
	CloudV1ScheduleUpdateParamsTypeInterval CloudV1ScheduleUpdateParamsType = "interval"
)

func (r CloudV1ScheduleUpdateParamsType) IsKnown() bool {
	switch r {
	case CloudV1ScheduleUpdateParamsTypeCron, CloudV1ScheduleUpdateParamsTypeInterval:
		return true
	}
	return false
}
