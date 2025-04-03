// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestCloudV1GetAvailableFloatingIPs(t *testing.T) {
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
	_, err := client.Cloud.V1.GetAvailableFloatingIPs(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1GetAvailableNetworksWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetAvailableNetworks(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1GetAvailableNetworksParams{
			Limit:       gcore.F(int64(0)),
			MetadataK:   gcore.F("metadata_k"),
			MetadataKv:  gcore.F("metadata_kv"),
			NetworkID:   gcore.F("network_id"),
			NetworkType: gcore.F("network_type"),
			Offset:      gcore.F(int64(0)),
			OrderBy:     gcore.F("order_by"),
			Shared:      gcore.F(true),
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

func TestCloudV1GetBmCapacity(t *testing.T) {
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
	_, err := client.Cloud.V1.GetBmCapacity(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1GetBmReservationFlavorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetBmReservationFlavors(
		context.TODO(),
		int64(0),
		gcore.CloudV1GetBmReservationFlavorsParams{
			ClientID:     gcore.F(int64(0)),
			Disabled:     gcore.F(true),
			EnsureCached: gcore.F("ensure_cached"),
			WindowsOs:    gcore.F(true),
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

func TestCloudV1GetFlavorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetFlavors(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1GetFlavorsParams{
			Disabled:       gcore.F(true),
			ExcludeLinux:   gcore.F(true),
			ExcludeWindows: gcore.F(true),
			IncludePrices:  gcore.F(true),
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

func TestCloudV1GetLbFlavorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetLbFlavors(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1GetLbFlavorsParams{
			IncludePrices: gcore.F(true),
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

func TestCloudV1GetServiceEndpointsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetServiceEndpoints(context.TODO(), gcore.CloudV1GetServiceEndpointsParams{
		RegionID: gcore.F(int64(2)),
		Service:  gcore.F(gcore.EndpointServiceAIService),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1GetUsageReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.GetUsageReport(context.TODO(), gcore.CloudV1GetUsageReportParams{
		TimeFrom:      gcore.F(time.Now()),
		TimeTo:        gcore.F(time.Now()),
		EnableLastDay: gcore.F(false),
		Limit:         gcore.F(int64(10)),
		Offset:        gcore.F(int64(0)),
		Projects:      gcore.F([]int64{int64(16), int64(17), int64(18), int64(19), int64(20)}),
		Regions:       gcore.F([]int64{int64(1), int64(2), int64(3)}),
		SchemaFilter: gcore.F[gcore.CloudV1GetUsageReportParamsSchemaFilterUnion](gcore.SchemaFilterInstanceParam{
			Field:  gcore.F(gcore.SchemaFilterInstanceEnumFlavor),
			Type:   gcore.F(gcore.SchemaFilterInstanceTypeInstance),
			Values: gcore.F([]string{"g1-standard-1-2"}),
		}),
		Sorting: gcore.F([]gcore.CloudV1GetUsageReportParamsSorting{{
			BillingValue: gcore.F(gcore.SortingDirectionsAsc),
			FirstSeen:    gcore.F(gcore.SortingDirectionsAsc),
			LastName:     gcore.F(gcore.SortingDirectionsAsc),
			LastSeen:     gcore.F(gcore.SortingDirectionsAsc),
			Project:      gcore.F(gcore.SortingDirectionsAsc),
			Region:       gcore.F(gcore.SortingDirectionsAsc),
			Type:         gcore.F(gcore.SortingDirectionsAsc),
		}}),
		Tags: gcore.F(gcore.TagsFilterParam{
			Conditions: gcore.F([]gcore.TagsFilterConditionParam{{
				Key:    gcore.F("os_version"),
				Strict: gcore.F(true),
				Value:  gcore.F("22.04"),
			}}),
			ConditionType: gcore.F(gcore.TagsFilterConditionTypeAnd),
		}),
		Types: gcore.F([]gcore.PrebillingResourceTypes{gcore.PrebillingResourceTypesAICluster, gcore.PrebillingResourceTypesBackup}),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1ListProjectImagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.ListProjectImages(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1ListProjectImagesParams{
			IncludePrices: gcore.F(true),
			MetadataK:     gcore.F("metadata_k"),
			MetadataKv:    gcore.F("metadata_kv"),
			Private:       gcore.F("private"),
			Visibility:    gcore.F(gcore.CloudV1ListProjectImagesParamsVisibilityPrivate),
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

func TestCloudV1UploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.UploadImage(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1UploadImageParams{
			Name:           gcore.F("image_name"),
			URL:            gcore.F("http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img"),
			Architecture:   gcore.F(gcore.CloudV1UploadImageParamsArchitectureAarch64),
			CowFormat:      gcore.F(false),
			HwFirmwareType: gcore.F(gcore.CloudV1UploadImageParamsHwFirmwareTypeBios),
			HwMachineType:  gcore.F(gcore.CloudV1UploadImageParamsHwMachineTypeI440),
			IsBaremetal:    gcore.F(false),
			Metadata: gcore.F[any](map[string]interface{}{
				"key": "value",
			}),
			OsDistro:  gcore.F("os_distro"),
			OsType:    gcore.F(gcore.CloudV1UploadImageParamsOsTypeLinux),
			OsVersion: gcore.F("os_version"),
			SSHKey:    gcore.F(gcore.CloudV1UploadImageParamsSSHKeyAllow),
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
