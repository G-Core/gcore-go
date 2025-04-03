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

func TestCloudV1ScheduleGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Schedule.Get(
		context.TODO(),
		int64(1),
		int64(1),
		"1488e2ce-f906-47fb-ba32-c25a3f63df4f",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1ScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Schedule.Update(
		context.TODO(),
		int64(1),
		int64(1),
		"1488e2ce-f906-47fb-ba32-c25a3f63df4f",
		gcore.CloudV1ScheduleUpdateParams{
			Day:                  gcore.F("5"),
			DayOfWeek:            gcore.F("mon,fri"),
			Days:                 gcore.F(int64(0)),
			Hour:                 gcore.F("0, 20"),
			Hours:                gcore.F(int64(2)),
			MaxQuantity:          gcore.F(int64(2)),
			Minute:               gcore.F("30"),
			Minutes:              gcore.F(int64(1)),
			Month:                gcore.F("1,6"),
			ResourceNameTemplate: gcore.F("Snapshot of volume {volume_id} created by policy {lifecycle_policy_id}"),
			RetentionTime: gcore.F(gcore.CreateIntervalTimeParam{
				Days:    gcore.F(int64(0)),
				Hours:   gcore.F(int64(2)),
				Minutes: gcore.F(int64(1)),
				Weeks:   gcore.F(int64(0)),
			}),
			Timezone: gcore.F("UTC"),
			Type:     gcore.F(gcore.CloudV1ScheduleUpdateParamsTypeCron),
			Week:     gcore.F("1"),
			Weeks:    gcore.F(int64(0)),
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
