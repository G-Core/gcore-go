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

func TestLifecyclePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.New(context.TODO(), cloud.LifecyclePolicyNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Action:    cloud.LifecyclePolicyNewParamsActionVolumeSnapshot,
		Name:      "schedule_1",
		Schedules: []cloud.LifecyclePolicyNewParamsScheduleUnion{{
			OfCron: &cloud.LifecyclePolicyNewParamsScheduleCron{
				Day:                  gcore.String("5"),
				DayOfWeek:            gcore.String("fri"),
				Hour:                 gcore.String("0, 20"),
				MaxQuantity:          gcore.Int(2),
				Minute:               gcore.String("30"),
				Month:                gcore.String("1"),
				ResourceNameTemplate: gcore.String("snapshot of volume {volume_id}"),
				RetentionTime: cloud.LifecyclePolicyNewParamsScheduleCronRetentionTime{
					Days:    gcore.Int(0),
					Hours:   gcore.Int(2),
					Minutes: gcore.Int(1),
					Weeks:   gcore.Int(0),
				},
				Timezone: gcore.String("UTC"),
				Week:     gcore.String("1"),
			},
		}},
		Status:    cloud.LifecyclePolicyNewParamsStatusActive,
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

func TestLifecyclePolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.Update(
		context.TODO(),
		1,
		cloud.LifecyclePolicyUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("schedule_1"),
			Status:    cloud.LifecyclePolicyUpdateParamsStatusPaused,
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

func TestLifecyclePolicyList(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.List(context.TODO(), cloud.LifecyclePolicyListParams{
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

func TestLifecyclePolicyDelete(t *testing.T) {
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
	err := client.Cloud.LifecyclePolicies.Delete(
		context.TODO(),
		1,
		cloud.LifecyclePolicyDeleteParams{
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

func TestLifecyclePolicyAddSchedules(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.AddSchedules(
		context.TODO(),
		1,
		cloud.LifecyclePolicyAddSchedulesParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Schedules: []cloud.LifecyclePolicyAddSchedulesParamsScheduleUnion{{
				OfCron: &cloud.LifecyclePolicyAddSchedulesParamsScheduleCron{
					Day:                  gcore.String("5"),
					DayOfWeek:            gcore.String("fri, tue"),
					Hour:                 gcore.String("0, 20"),
					MaxQuantity:          gcore.Int(2),
					Minute:               gcore.String("30"),
					Month:                gcore.String("1"),
					ResourceNameTemplate: gcore.String("CRON reserve snap of the volume {volume_id}"),
					RetentionTime: cloud.LifecyclePolicyAddSchedulesParamsScheduleCronRetentionTime{
						Days:    gcore.Int(0),
						Hours:   gcore.Int(2),
						Minutes: gcore.Int(1),
						Weeks:   gcore.Int(2),
					},
					Timezone: gcore.String("UTC"),
					Week:     gcore.String("1"),
				},
			}, {
				OfInterval: &cloud.LifecyclePolicyAddSchedulesParamsScheduleInterval{
					Days:                 gcore.Int(0),
					Hours:                gcore.Int(2),
					MaxQuantity:          gcore.Int(2),
					Minutes:              gcore.Int(1),
					ResourceNameTemplate: gcore.String("INTERVAL reserve snap of the volume {volume_id}"),
					RetentionTime: cloud.LifecyclePolicyAddSchedulesParamsScheduleIntervalRetentionTime{
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

func TestLifecyclePolicyAddVolumes(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.AddVolumes(
		context.TODO(),
		1,
		cloud.LifecyclePolicyAddVolumesParams{
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

func TestLifecyclePolicyEstimateMaxUsageWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.EstimateMaxUsage(context.TODO(), cloud.LifecyclePolicyEstimateMaxUsageParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Action:    cloud.LifecyclePolicyEstimateMaxUsageParamsActionVolumeSnapshot,
		Name:      "schedule_1",
		Schedules: []cloud.LifecyclePolicyEstimateMaxUsageParamsScheduleUnion{{
			OfCron: &cloud.LifecyclePolicyEstimateMaxUsageParamsScheduleCron{
				Day:                  gcore.String("5"),
				DayOfWeek:            gcore.String("fri"),
				Hour:                 gcore.String("0, 20"),
				MaxQuantity:          gcore.Int(2),
				Minute:               gcore.String("30"),
				Month:                gcore.String("1"),
				ResourceNameTemplate: gcore.String("snapshot of volume {volume_id}"),
				RetentionTime: cloud.LifecyclePolicyEstimateMaxUsageParamsScheduleCronRetentionTime{
					Days:    gcore.Int(0),
					Hours:   gcore.Int(2),
					Minutes: gcore.Int(1),
					Weeks:   gcore.Int(0),
				},
				Timezone: gcore.String("UTC"),
				Week:     gcore.String("1"),
			},
		}},
		Status:    cloud.LifecyclePolicyEstimateMaxUsageParamsStatusActive,
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

func TestLifecyclePolicyGet(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.Get(
		context.TODO(),
		1,
		cloud.LifecyclePolicyGetParams{
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

func TestLifecyclePolicyRemoveSchedules(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.RemoveSchedules(
		context.TODO(),
		1,
		cloud.LifecyclePolicyRemoveSchedulesParams{
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

func TestLifecyclePolicyRemoveVolumes(t *testing.T) {
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
	_, err := client.Cloud.LifecyclePolicies.RemoveVolumes(
		context.TODO(),
		1,
		cloud.LifecyclePolicyRemoveVolumesParams{
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
