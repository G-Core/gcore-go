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

func TestCloudV2LimitsRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.LimitsRequest.New(context.TODO(), gcore.CloudV2LimitsRequestNewParams{
		Description: gcore.F("Scale up mysql cluster"),
		RequestedLimits: gcore.F(gcore.CloudV2LimitsRequestNewParamsRequestedLimits{
			GlobalLimits: gcore.F(gcore.GlobalQuotasLimitsParam{
				InferenceCPUMillicoreCountLimit: gcore.F(int64(0)),
				InferenceGPUA100CountLimit:      gcore.F(int64(0)),
				InferenceGPUH100CountLimit:      gcore.F(int64(0)),
				InferenceGPUL40sCountLimit:      gcore.F(int64(0)),
				InferenceInstanceCountLimit:     gcore.F(int64(0)),
				KeypairCountLimit:               gcore.F(int64(100)),
				ProjectCountLimit:               gcore.F(int64(2)),
			}),
			RegionalLimits: gcore.F([]gcore.RegionalQuotasLimitsParam{{
				BaremetalBasicCountLimit:          gcore.F(int64(0)),
				BaremetalGPUCountLimit:            gcore.F(int64(0)),
				BaremetalHfCountLimit:             gcore.F(int64(0)),
				BaremetalInfrastructureCountLimit: gcore.F(int64(0)),
				BaremetalNetworkCountLimit:        gcore.F(int64(0)),
				BaremetalStorageCountLimit:        gcore.F(int64(0)),
				CaasContainerCountLimit:           gcore.F(int64(0)),
				CaasCPUCountLimit:                 gcore.F(int64(0)),
				CaasGPUCountLimit:                 gcore.F(int64(0)),
				CaasRamSizeLimit:                  gcore.F(int64(0)),
				ClusterCountLimit:                 gcore.F(int64(0)),
				CPUCountLimit:                     gcore.F(int64(0)),
				DbaasPostgresClusterCountLimit:    gcore.F(int64(0)),
				ExternalIPCountLimit:              gcore.F(int64(0)),
				FaasCPUCountLimit:                 gcore.F(int64(0)),
				FaasFunctionCountLimit:            gcore.F(int64(0)),
				FaasNamespaceCountLimit:           gcore.F(int64(0)),
				FaasRamSizeLimit:                  gcore.F(int64(0)),
				FirewallCountLimit:                gcore.F(int64(0)),
				FloatingCountLimit:                gcore.F(int64(0)),
				GPUCountLimit:                     gcore.F(int64(0)),
				ImageCountLimit:                   gcore.F(int64(0)),
				ImageSizeLimit:                    gcore.F(int64(0)),
				IpuCountLimit:                     gcore.F(int64(0)),
				LaasTopicCountLimit:               gcore.F(int64(0)),
				LoadbalancerCountLimit:            gcore.F(int64(0)),
				NetworkCountLimit:                 gcore.F(int64(0)),
				RamLimit:                          gcore.F(int64(0)),
				RegionID:                          gcore.F(int64(1)),
				RegistryCountLimit:                gcore.F(int64(0)),
				RegistryStorageLimit:              gcore.F(int64(0)),
				RouterCountLimit:                  gcore.F(int64(0)),
				SecretCountLimit:                  gcore.F(int64(0)),
				ServergroupCountLimit:             gcore.F(int64(0)),
				SfsCountLimit:                     gcore.F(int64(0)),
				SfsSizeLimit:                      gcore.F(int64(0)),
				SharedVmCountLimit:                gcore.F(int64(0)),
				SnapshotScheduleCountLimit:        gcore.F(int64(0)),
				SubnetCountLimit:                  gcore.F(int64(0)),
				VmCountLimit:                      gcore.F(int64(0)),
				VolumeCountLimit:                  gcore.F(int64(0)),
				VolumeSizeLimit:                   gcore.F(int64(0)),
				VolumeSnapshotsCountLimit:         gcore.F(int64(0)),
				VolumeSnapshotsSizeLimit:          gcore.F(int64(0)),
			}}),
		}),
		ClientID: gcore.F(int64(1)),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2LimitsRequestGet(t *testing.T) {
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
	_, err := client.Cloud.V2.LimitsRequest.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2LimitsRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.LimitsRequest.List(context.TODO(), gcore.CloudV2LimitsRequestListParams{
		ClientID: gcore.F(int64(0)),
		Limit:    gcore.F(int64(0)),
		Offset:   gcore.F(int64(0)),
		Status:   gcore.F(gcore.CloudV2LimitsRequestListParamsStatusDone),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2LimitsRequestDelete(t *testing.T) {
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
	err := client.Cloud.V2.LimitsRequest.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
