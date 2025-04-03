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

func TestCloudV2ClientQuotaNotificationThresholdGet(t *testing.T) {
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
	_, err := client.Cloud.V2.ClientQuotas.NotificationThreshold.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2ClientQuotaNotificationThresholdDelete(t *testing.T) {
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
	err := client.Cloud.V2.ClientQuotas.NotificationThreshold.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2ClientQuotaNotificationThresholdUpdateOrNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.ClientQuotas.NotificationThreshold.UpdateOrNew(
		context.TODO(),
		int64(0),
		gcore.CloudV2ClientQuotaNotificationThresholdUpdateOrNewParams{
			Threshold: gcore.F(int64(95)),
			LastMessage: gcore.F(gcore.QuotasThresholdParam{
				GlobalQuotas: gcore.F(gcore.QuotasThresholdGlobalQuotasParam{
					InferenceCPUMillicoreCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceCPUMillicoreCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUA100CountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUA100CountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUH100CountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUH100CountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUL40sCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceGPUL40sCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceInstanceCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					InferenceInstanceCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					KeypairCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					KeypairCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ProjectCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ProjectCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
				}),
				RegionalQuotas: gcore.F([]gcore.QuotasThresholdRegionalQuotaParam{{
					RegionID:   gcore.F(int64(2)),
					RegionName: gcore.F("Luxembourg"),
					BaremetalBasicCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalBasicCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalGPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalGPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalHfCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalHfCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalInfrastructureCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalInfrastructureCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalNetworkCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalNetworkCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalStorageCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					BaremetalStorageCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasContainerCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasContainerCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasCPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasCPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasGPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasGPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasRamSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CaasRamSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ClusterCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ClusterCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					CPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					DbaasPostgresClusterCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					DbaasPostgresClusterCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ExternalIPCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ExternalIPCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasCPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasCPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasFunctionCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasFunctionCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasNamespaceCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasNamespaceCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasRamSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FaasRamSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FirewallCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FirewallCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FloatingCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					FloatingCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					GPUCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					GPUCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ImageCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ImageCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ImageSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ImageSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					IpuCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					IpuCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					LaasTopicCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					LaasTopicCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					LoadbalancerCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					LoadbalancerCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					NetworkCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					NetworkCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RamLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RamUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RegistryCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RegistryCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RegistryStorageLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RegistryStorageUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RouterCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					RouterCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SecretCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SecretCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ServergroupCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					ServergroupCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SfsCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SfsCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SfsSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SfsSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SharedVmCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SharedVmCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SnapshotScheduleCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SnapshotScheduleCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SubnetCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					SubnetCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VmCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VmCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSnapshotsCountLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSnapshotsCountUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSnapshotsSizeLimit: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
					VolumeSnapshotsSizeUsage: gcore.F(gcore.QuotaCountParam{
						Limit: gcore.F(int64(10)),
						Usage: gcore.F(int64(8)),
					}),
				}}),
			}),
			LastSending: gcore.F(time.Now()),
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
