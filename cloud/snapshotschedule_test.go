// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestSnapshotScheduleNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.New(context.TODO(), cloud.SnapshotScheduleNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Action:    cloud.SnapshotScheduleNewParamsActionVolumeSnapshot,
		Name:      "schedule_1",
		Schedules: []cloud.SnapshotScheduleNewParamsScheduleUnion{{
			OfCron: &cloud.SnapshotScheduleNewParamsScheduleCron{
				Day:                  gcore.String("5"),
				DayOfWeek:            gcore.String("fri"),
				Hour:                 gcore.String("0, 20"),
				MaxQuantity:          gcore.Int(2),
				Minute:               gcore.String("30"),
				Month:                gcore.String("1"),
				ResourceNameTemplate: gcore.String("snapshot of volume {volume_id}"),
				RetentionTime: cloud.SnapshotScheduleNewParamsScheduleCronRetentionTime{
					Days:    gcore.Int(0),
					Hours:   gcore.Int(2),
					Minutes: gcore.Int(1),
					Weeks:   gcore.Int(0),
				},
				Timezone: gcore.String("UTC"),
				Week:     gcore.String("1"),
			},
		}},
		Status:    cloud.SnapshotScheduleNewParamsStatusActive,
		VolumeIDs: []string{"3ed9e2ce-f906-47fb-ba32-c25a3f63df4f"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.Update(
		context.TODO(),
		1,
		cloud.SnapshotScheduleUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("schedule_1"),
			Status:    cloud.SnapshotScheduleUpdateParamsStatusPaused,
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.List(context.TODO(), cloud.SnapshotScheduleListParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Cloud.SnapshotSchedules.Delete(
		context.TODO(),
		1,
		cloud.SnapshotScheduleDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleAddSchedules(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.AddSchedules(
		context.TODO(),
		1,
		cloud.SnapshotScheduleAddSchedulesParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Schedules: []cloud.SnapshotScheduleAddSchedulesParamsScheduleUnion{{
				OfCron: &cloud.SnapshotScheduleAddSchedulesParamsScheduleCron{
					Day:                  gcore.String("5"),
					DayOfWeek:            gcore.String("fri, tue"),
					Hour:                 gcore.String("0, 20"),
					MaxQuantity:          gcore.Int(2),
					Minute:               gcore.String("30"),
					Month:                gcore.String("1"),
					ResourceNameTemplate: gcore.String("CRON reserve snap of the volume {volume_id}"),
					RetentionTime: cloud.SnapshotScheduleAddSchedulesParamsScheduleCronRetentionTime{
						Days:    gcore.Int(0),
						Hours:   gcore.Int(2),
						Minutes: gcore.Int(1),
						Weeks:   gcore.Int(2),
					},
					Timezone: gcore.String("UTC"),
					Week:     gcore.String("1"),
				},
			}, {
				OfInterval: &cloud.SnapshotScheduleAddSchedulesParamsScheduleInterval{
					Days:                 gcore.Int(0),
					Hours:                gcore.Int(2),
					MaxQuantity:          gcore.Int(2),
					Minutes:              gcore.Int(1),
					ResourceNameTemplate: gcore.String("INTERVAL reserve snap of the volume {volume_id}"),
					RetentionTime: cloud.SnapshotScheduleAddSchedulesParamsScheduleIntervalRetentionTime{
						Days:    gcore.Int(1),
						Hours:   gcore.Int(2),
						Minutes: gcore.Int(1),
						Weeks:   gcore.Int(0),
					},
					Weeks: gcore.Int(0),
				},
			}},
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleAddVolumes(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.AddVolumes(
		context.TODO(),
		1,
		cloud.SnapshotScheduleAddVolumesParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			VolumeIDs: []string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"},
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleEstimateMaxUsageWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.EstimateMaxUsage(context.TODO(), cloud.SnapshotScheduleEstimateMaxUsageParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Action:    cloud.SnapshotScheduleEstimateMaxUsageParamsActionVolumeSnapshot,
		Name:      "schedule_1",
		Schedules: []cloud.SnapshotScheduleEstimateMaxUsageParamsScheduleUnion{{
			OfCron: &cloud.SnapshotScheduleEstimateMaxUsageParamsScheduleCron{
				Day:                  gcore.String("5"),
				DayOfWeek:            gcore.String("fri"),
				Hour:                 gcore.String("0, 20"),
				MaxQuantity:          gcore.Int(2),
				Minute:               gcore.String("30"),
				Month:                gcore.String("1"),
				ResourceNameTemplate: gcore.String("snapshot of volume {volume_id}"),
				RetentionTime: cloud.SnapshotScheduleEstimateMaxUsageParamsScheduleCronRetentionTime{
					Days:    gcore.Int(0),
					Hours:   gcore.Int(2),
					Minutes: gcore.Int(1),
					Weeks:   gcore.Int(0),
				},
				Timezone: gcore.String("UTC"),
				Week:     gcore.String("1"),
			},
		}},
		Status:    cloud.SnapshotScheduleEstimateMaxUsageParamsStatusActive,
		VolumeIDs: []string{"3ed9e2ce-f906-47fb-ba32-c25a3f63df4f"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.Get(
		context.TODO(),
		1,
		cloud.SnapshotScheduleGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleRemoveSchedules(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.RemoveSchedules(
		context.TODO(),
		1,
		cloud.SnapshotScheduleRemoveSchedulesParams{
			ProjectID:   gcore.Int(1),
			RegionID:    gcore.Int(1),
			ScheduleIDs: []string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"},
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSnapshotScheduleRemoveVolumes(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SnapshotSchedules.RemoveVolumes(
		context.TODO(),
		1,
		cloud.SnapshotScheduleRemoveVolumesParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			VolumeIDs: []string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"},
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
