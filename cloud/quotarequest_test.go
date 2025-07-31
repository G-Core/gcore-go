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

func TestQuotaRequestNewWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.Quotas.Requests.New(context.TODO(), cloud.QuotaRequestNewParams{
		Description: "Scale up mysql cluster",
		RequestedLimits: cloud.QuotaRequestNewParamsRequestedLimits{
			GlobalLimits: cloud.QuotaRequestNewParamsRequestedLimitsGlobalLimits{
				InferenceCPUMillicoreCountLimit: gcore.Int(0),
				InferenceGPUA100CountLimit:      gcore.Int(0),
				InferenceGPUH100CountLimit:      gcore.Int(0),
				InferenceGPUL40sCountLimit:      gcore.Int(0),
				InferenceInstanceCountLimit:     gcore.Int(0),
				KeypairCountLimit:               gcore.Int(100),
				ProjectCountLimit:               gcore.Int(2),
			},
			RegionalLimits: []cloud.QuotaRequestNewParamsRequestedLimitsRegionalLimit{{
				BaremetalBasicCountLimit:          gcore.Int(0),
				BaremetalGPUA100CountLimit:        gcore.Int(0),
				BaremetalGPUCountLimit:            gcore.Int(0),
				BaremetalGPUH100CountLimit:        gcore.Int(0),
				BaremetalGPUH200CountLimit:        gcore.Int(0),
				BaremetalGPUL40sCountLimit:        gcore.Int(0),
				BaremetalHfCountLimit:             gcore.Int(0),
				BaremetalInfrastructureCountLimit: gcore.Int(0),
				BaremetalNetworkCountLimit:        gcore.Int(0),
				BaremetalStorageCountLimit:        gcore.Int(0),
				CaasContainerCountLimit:           gcore.Int(0),
				CaasCPUCountLimit:                 gcore.Int(0),
				CaasGPUCountLimit:                 gcore.Int(0),
				CaasRamSizeLimit:                  gcore.Int(0),
				ClusterCountLimit:                 gcore.Int(0),
				CPUCountLimit:                     gcore.Int(0),
				DbaasPostgresClusterCountLimit:    gcore.Int(0),
				ExternalIPCountLimit:              gcore.Int(0),
				FaasCPUCountLimit:                 gcore.Int(0),
				FaasFunctionCountLimit:            gcore.Int(0),
				FaasNamespaceCountLimit:           gcore.Int(0),
				FaasRamSizeLimit:                  gcore.Int(0),
				FirewallCountLimit:                gcore.Int(0),
				FloatingCountLimit:                gcore.Int(0),
				GPUCountLimit:                     gcore.Int(0),
				GPUVirtualA100CountLimit:          gcore.Int(0),
				GPUVirtualH100CountLimit:          gcore.Int(0),
				GPUVirtualH200CountLimit:          gcore.Int(0),
				GPUVirtualL40sCountLimit:          gcore.Int(0),
				ImageCountLimit:                   gcore.Int(0),
				ImageSizeLimit:                    gcore.Int(0),
				IpuCountLimit:                     gcore.Int(0),
				LaasTopicCountLimit:               gcore.Int(0),
				LoadbalancerCountLimit:            gcore.Int(0),
				NetworkCountLimit:                 gcore.Int(0),
				RamLimit:                          gcore.Int(0),
				RegionID:                          gcore.Int(1),
				RegistryCountLimit:                gcore.Int(0),
				RegistryStorageLimit:              gcore.Int(0),
				RouterCountLimit:                  gcore.Int(0),
				SecretCountLimit:                  gcore.Int(0),
				ServergroupCountLimit:             gcore.Int(0),
				SfsCountLimit:                     gcore.Int(0),
				SfsSizeLimit:                      gcore.Int(0),
				SharedVmCountLimit:                gcore.Int(0),
				SnapshotScheduleCountLimit:        gcore.Int(0),
				SubnetCountLimit:                  gcore.Int(0),
				VmCountLimit:                      gcore.Int(0),
				VolumeCountLimit:                  gcore.Int(0),
				VolumeSizeLimit:                   gcore.Int(0),
				VolumeSnapshotsCountLimit:         gcore.Int(0),
				VolumeSnapshotsSizeLimit:          gcore.Int(0),
			}},
		},
		ClientID: gcore.Int(1),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Quotas.Requests.List(context.TODO(), cloud.QuotaRequestListParams{
		Limit:  gcore.Int(1000),
		Offset: gcore.Int(0),
		Status: []string{"done", "in progress"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaRequestDelete(t *testing.T) {
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
	err := client.Cloud.Quotas.Requests.Delete(context.TODO(), 3)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaRequestGet(t *testing.T) {
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
	_, err := client.Cloud.Quotas.Requests.Get(context.TODO(), 3)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
