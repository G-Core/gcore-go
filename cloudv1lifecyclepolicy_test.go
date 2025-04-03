// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestCloudV1LifecyclePolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.New(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyNewParams{
			CreateLifecyclePolicy: gcore.CreateLifecyclePolicyParam{
				Action: gcore.F(gcore.CreateLifecyclePolicyActionVolumeSnapshot),
				Name:   gcore.F("schedule_1"),
				Schedules: gcore.F([]gcore.CreateScheduleUnionParam{gcore.CreateScheduleCreateCronScheduleSerializerParam{
					Type:                 gcore.F(gcore.CreateScheduleCreateCronScheduleSerializerTypeCron),
					Day:                  gcore.F("5"),
					DayOfWeek:            gcore.F("mon,fri"),
					Hour:                 gcore.F("0, 20"),
					MaxQuantity:          gcore.F(int64(2)),
					Minute:               gcore.F("30"),
					Month:                gcore.F("1,6"),
					ResourceNameTemplate: gcore.F("Snapshot of volume {volume_id} created by policy {lifecycle_policy_id}"),
					RetentionTime: gcore.F(gcore.CreateIntervalTimeParam{
						Days:    gcore.F(int64(0)),
						Hours:   gcore.F(int64(2)),
						Minutes: gcore.F(int64(1)),
						Weeks:   gcore.F(int64(0)),
					}),
					Timezone: gcore.F("UTC"),
					Week:     gcore.F("1"),
				}}),
				Status:    gcore.F(gcore.LifecyclePolicyStatusesActive),
				VolumeIDs: gcore.F([]string{"3ed9e2ce-f906-47fb-ba32-c25a3f63df4f"}),
			},
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

func TestCloudV1LifecyclePolicyGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.Get(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1LifecyclePolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.Update(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyUpdateParams{
			Name:   gcore.F("schedule_1"),
			Status: gcore.F(gcore.LifecyclePolicyStatusesActive),
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

func TestCloudV1LifecyclePolicyList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.List(
		context.TODO(),
		int64(1),
		int64(1),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1LifecyclePolicyDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Cloud.V1.LifecyclePolicy.Delete(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1LifecyclePolicyAddSchedules(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.AddSchedules(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyAddSchedulesParams{
			Schedules: gcore.F([]gcore.CreateScheduleUnionParam{gcore.CreateScheduleCreateCronScheduleSerializerParam{
				Type:                 gcore.F(gcore.CreateScheduleCreateCronScheduleSerializerTypeCron),
				Day:                  gcore.F("5"),
				DayOfWeek:            gcore.F("mon,fri"),
				Hour:                 gcore.F("0, 20"),
				MaxQuantity:          gcore.F(int64(2)),
				Minute:               gcore.F("30"),
				Month:                gcore.F("1,6"),
				ResourceNameTemplate: gcore.F("Snapshot of volume {volume_id} created by policy {lifecycle_policy_id}"),
				RetentionTime: gcore.F(gcore.CreateIntervalTimeParam{
					Days:    gcore.F(int64(0)),
					Hours:   gcore.F(int64(2)),
					Minutes: gcore.F(int64(1)),
					Weeks:   gcore.F(int64(0)),
				}),
				Timezone: gcore.F("UTC"),
				Week:     gcore.F("1"),
			}}),
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

func TestCloudV1LifecyclePolicyAddVolumes(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.AddVolumes(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyAddVolumesParams{
			VolumeIDs: gcore.VolumeIDsParam{
				VolumeIDs: gcore.F([]string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"}),
			},
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

func TestCloudV1LifecyclePolicyEstimateMaxUsageWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.EstimateMaxUsage(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyEstimateMaxUsageParams{
			CreateLifecyclePolicy: gcore.CreateLifecyclePolicyParam{
				Action: gcore.F(gcore.CreateLifecyclePolicyActionVolumeSnapshot),
				Name:   gcore.F("schedule_1"),
				Schedules: gcore.F([]gcore.CreateScheduleUnionParam{gcore.CreateScheduleCreateCronScheduleSerializerParam{
					Type:                 gcore.F(gcore.CreateScheduleCreateCronScheduleSerializerTypeCron),
					Day:                  gcore.F("5"),
					DayOfWeek:            gcore.F("mon,fri"),
					Hour:                 gcore.F("0, 20"),
					MaxQuantity:          gcore.F(int64(2)),
					Minute:               gcore.F("30"),
					Month:                gcore.F("1,6"),
					ResourceNameTemplate: gcore.F("Snapshot of volume {volume_id} created by policy {lifecycle_policy_id}"),
					RetentionTime: gcore.F(gcore.CreateIntervalTimeParam{
						Days:    gcore.F(int64(0)),
						Hours:   gcore.F(int64(2)),
						Minutes: gcore.F(int64(1)),
						Weeks:   gcore.F(int64(0)),
					}),
					Timezone: gcore.F("UTC"),
					Week:     gcore.F("1"),
				}}),
				Status:    gcore.F(gcore.LifecyclePolicyStatusesActive),
				VolumeIDs: gcore.F([]string{"3ed9e2ce-f906-47fb-ba32-c25a3f63df4f"}),
			},
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

func TestCloudV1LifecyclePolicyRemoveSchedules(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.RemoveSchedules(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyRemoveSchedulesParams{
			ScheduleIDs: gcore.F([]string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"}),
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

func TestCloudV1LifecyclePolicyRemoveVolumes(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.LifecyclePolicy.RemoveVolumes(
		context.TODO(),
		int64(1),
		int64(1),
		int64(1),
		gcore.CloudV1LifecyclePolicyRemoveVolumesParams{
			VolumeIDs: gcore.VolumeIDsParam{
				VolumeIDs: gcore.F([]string{"1488e2ce-f906-47fb-ba32-c25a3f63df4f"}),
			},
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
