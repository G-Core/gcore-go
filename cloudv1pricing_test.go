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

func TestCloudV1PricingGetBillingReservationPricesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.GetBillingReservationPrices(
		context.TODO(),
		int64(1),
		gcore.CloudV1PricingGetBillingReservationPricesParams{
			Period: gcore.F(gcore.CloudV1PricingGetBillingReservationPricesParamsPeriodMonths12),
			Resources: gcore.F([]gcore.BillingReservationPricingResourceRequestParam{{
				ResourceName:  gcore.F("bm1-hf-medium"),
				ResourceType:  gcore.F(gcore.BillingReservationTypeFlavor),
				ResourceValue: gcore.F(int64(2)),
			}}),
			ClientID:   gcore.F(int64(0)),
			HasWindows: gcore.F(true),
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

func TestCloudV1PricingGetDdosPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.GetDdosPrice(context.TODO(), int64(1))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1PricingGetInstancePrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.GetInstancePrice(
		context.TODO(),
		int64(1),
		int64(1),
		"instance_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1PricingGetLaasPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.GetLaasPrice(context.TODO(), int64(1))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1PricingGetLifecyclePolicyCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.GetLifecyclePolicyCost(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingGetLifecyclePolicyCostParams{
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

func TestCloudV1PricingPreviewContainerPriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewContainerPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewContainerPriceParams{
			Flavor: gcore.F("250mCPU-512MiB"),
			Scale: gcore.F(gcore.ScaleParam{
				Max:            gcore.F(int64(1)),
				Min:            gcore.F(int64(0)),
				CooldownPeriod: gcore.F(int64(1)),
				Triggers: gcore.F(gcore.ScaleTriggersParam{
					CPU: gcore.F(gcore.ScaleTriggersCPUParam{
						Threshold: gcore.F(int64(1)),
					}),
					HTTP: gcore.F(gcore.ScaleTriggersHTTPParam{
						Rate:   gcore.F(int64(1)),
						Window: gcore.F(int64(1)),
					}),
					Memory: gcore.F(gcore.ScaleTriggersMemoryParam{
						Threshold: gcore.F(int64(1)),
					}),
				}),
			}),
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

func TestCloudV1PricingPreviewFileSharePriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewFileSharePrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewFileSharePriceParams{
			Size:     gcore.F(int64(0)),
			Source:   gcore.F(gcore.CloudV1PricingPreviewFileSharePriceParamsSourceNewFileShare),
			TypeName: gcore.F(gcore.VolumeTypeNameCold),
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

func TestCloudV1PricingPreviewFloatingIPPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewFloatingIPPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewFloatingIPPriceParams{
			Count: gcore.F(int64(2)),
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

func TestCloudV1PricingPreviewImagePrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewImagePrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewImagePriceParams{
			Size: gcore.F(int64(1000)),
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

func TestCloudV1PricingPreviewLoadBalancerPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewLoadBalancerPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewLoadBalancerPriceParams{
			Flavor: gcore.F("lb1-1-2"),
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

func TestCloudV1PricingPreviewReservedFixedIPPriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewReservedFixedIPPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewReservedFixedIPPriceParams{
			CreateReservedFixedIP: gcore.CreateReservedFixedIPNewReservedFixedIPExternalSerializerParam{
				Type:     gcore.F(gcore.CreateReservedFixedIPNewReservedFixedIPExternalSerializerTypeExternal),
				IPFamily: gcore.F(gcore.InterfaceIPFamilyDual),
				IsVip:    gcore.F(false),
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

func TestCloudV1PricingPreviewSnapshotPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewSnapshotPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewSnapshotPriceParams{
			VolumeID: gcore.F("c23fb11d-4748-489e-a659-e888e121e825"),
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

func TestCloudV1PricingPreviewVolumePriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.PreviewVolumePrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingPreviewVolumePriceParams{
			Source:     gcore.F(gcore.CloudV1PricingPreviewVolumePriceParamsSourceImage),
			Size:       gcore.F(int64(1)),
			SnapshotID: gcore.F("snapshot_id"),
			TypeName:   gcore.F(gcore.VolumeTypeNameCold),
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
