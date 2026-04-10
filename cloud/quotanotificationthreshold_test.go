// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func TestQuotaNotificationThresholdUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Quotas.NotificationThreshold.Update(
		context.TODO(),
		3,
		cloud.QuotaNotificationThresholdUpdateParams{
			Threshold: 95,
			LastMessage: cloud.QuotaNotificationThresholdUpdateParamsLastMessage{
				GlobalQuotas: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotas{
					InferenceCPUMillicoreCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountLimit{
						Limit: 10,
						Usage: 8,
					},
					InferenceCPUMillicoreCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceCPUMillicoreCountUsage{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUA100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUA100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUA100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUH100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUH100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUH100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUL40sCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountLimit{
						Limit: 10,
						Usage: 8,
					},
					InferenceGPUL40sCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceGPUL40sCountUsage{
						Limit: 10,
						Usage: 8,
					},
					InferenceInstanceCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountLimit{
						Limit: 10,
						Usage: 8,
					},
					InferenceInstanceCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasInferenceInstanceCountUsage{
						Limit: 10,
						Usage: 8,
					},
					KeypairCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountLimit{
						Limit: 10,
						Usage: 8,
					},
					KeypairCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasKeypairCountUsage{
						Limit: 10,
						Usage: 8,
					},
					ProjectCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountLimit{
						Limit: 10,
						Usage: 8,
					},
					ProjectCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageGlobalQuotasProjectCountUsage{
						Limit: 10,
						Usage: 8,
					},
				},
				RegionalQuotas: []cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuota{{
					RegionID:   2,
					RegionName: "Luxembourg",
					BaremetalBasicCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalBasicCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalBasicCountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUA100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUA100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUA100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUH100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUH100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUH200CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUH200CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUH200CountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUL40sCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalGPUL40sCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalGPUL40sCountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalHfCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalHfCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalHfCountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalInfrastructureCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalInfrastructureCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalInfrastructureCountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalNetworkCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalNetworkCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalNetworkCountUsage{
						Limit: 10,
						Usage: 8,
					},
					BaremetalStorageCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountLimit{
						Limit: 10,
						Usage: 8,
					},
					BaremetalStorageCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaBaremetalStorageCountUsage{
						Limit: 10,
						Usage: 8,
					},
					CaasContainerCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountLimit{
						Limit: 10,
						Usage: 8,
					},
					CaasContainerCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasContainerCountUsage{
						Limit: 10,
						Usage: 8,
					},
					CaasCPUCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountLimit{
						Limit: 10,
						Usage: 8,
					},
					CaasCPUCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasCPUCountUsage{
						Limit: 10,
						Usage: 8,
					},
					CaasGPUCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountLimit{
						Limit: 10,
						Usage: 8,
					},
					CaasGPUCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasGPUCountUsage{
						Limit: 10,
						Usage: 8,
					},
					CaasRamSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					CaasRamSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCaasRamSizeUsage{
						Limit: 10,
						Usage: 8,
					},
					ClusterCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountLimit{
						Limit: 10,
						Usage: 8,
					},
					ClusterCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaClusterCountUsage{
						Limit: 10,
						Usage: 8,
					},
					CPUCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountLimit{
						Limit: 10,
						Usage: 8,
					},
					CPUCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaCPUCountUsage{
						Limit: 10,
						Usage: 8,
					},
					DbaasPostgresClusterCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountLimit{
						Limit: 10,
						Usage: 8,
					},
					DbaasPostgresClusterCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaDbaasPostgresClusterCountUsage{
						Limit: 10,
						Usage: 8,
					},
					ExternalIPCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountLimit{
						Limit: 10,
						Usage: 8,
					},
					ExternalIPCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaExternalIPCountUsage{
						Limit: 10,
						Usage: 8,
					},
					FaasCPUCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountLimit{
						Limit: 10,
						Usage: 8,
					},
					FaasCPUCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasCPUCountUsage{
						Limit: 10,
						Usage: 8,
					},
					FaasFunctionCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountLimit{
						Limit: 10,
						Usage: 8,
					},
					FaasFunctionCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasFunctionCountUsage{
						Limit: 10,
						Usage: 8,
					},
					FaasNamespaceCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountLimit{
						Limit: 10,
						Usage: 8,
					},
					FaasNamespaceCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasNamespaceCountUsage{
						Limit: 10,
						Usage: 8,
					},
					FaasRamSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					FaasRamSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFaasRamSizeUsage{
						Limit: 10,
						Usage: 8,
					},
					FirewallCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountLimit{
						Limit: 10,
						Usage: 8,
					},
					FirewallCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFirewallCountUsage{
						Limit: 10,
						Usage: 8,
					},
					FloatingCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountLimit{
						Limit: 10,
						Usage: 8,
					},
					FloatingCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaFloatingCountUsage{
						Limit: 10,
						Usage: 8,
					},
					GPUCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountLimit{
						Limit: 10,
						Usage: 8,
					},
					GPUCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUCountUsage{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualA100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualA100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualA100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualH100CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountLimit{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualH100CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH100CountUsage{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualH200CountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountLimit{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualH200CountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualH200CountUsage{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualL40sCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountLimit{
						Limit: 10,
						Usage: 8,
					},
					GPUVirtualL40sCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaGPUVirtualL40sCountUsage{
						Limit: 10,
						Usage: 8,
					},
					ImageCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountLimit{
						Limit: 10,
						Usage: 8,
					},
					ImageCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageCountUsage{
						Limit: 10,
						Usage: 8,
					},
					ImageSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					ImageSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaImageSizeUsage{
						Limit: 10,
						Usage: 8,
					},
					IpuCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountLimit{
						Limit: 10,
						Usage: 8,
					},
					IpuCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaIpuCountUsage{
						Limit: 10,
						Usage: 8,
					},
					LaasTopicCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountLimit{
						Limit: 10,
						Usage: 8,
					},
					LaasTopicCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLaasTopicCountUsage{
						Limit: 10,
						Usage: 8,
					},
					LoadbalancerCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountLimit{
						Limit: 10,
						Usage: 8,
					},
					LoadbalancerCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaLoadbalancerCountUsage{
						Limit: 10,
						Usage: 8,
					},
					NetworkCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountLimit{
						Limit: 10,
						Usage: 8,
					},
					NetworkCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaNetworkCountUsage{
						Limit: 10,
						Usage: 8,
					},
					RamLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamLimit{
						Limit: 10,
						Usage: 8,
					},
					RamUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRamUsage{
						Limit: 10,
						Usage: 8,
					},
					RegistryCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountLimit{
						Limit: 10,
						Usage: 8,
					},
					RegistryCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryCountUsage{
						Limit: 10,
						Usage: 8,
					},
					RegistryStorageLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageLimit{
						Limit: 10,
						Usage: 8,
					},
					RegistryStorageUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRegistryStorageUsage{
						Limit: 10,
						Usage: 8,
					},
					RouterCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountLimit{
						Limit: 10,
						Usage: 8,
					},
					RouterCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaRouterCountUsage{
						Limit: 10,
						Usage: 8,
					},
					SecretCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountLimit{
						Limit: 10,
						Usage: 8,
					},
					SecretCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSecretCountUsage{
						Limit: 10,
						Usage: 8,
					},
					ServergroupCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountLimit{
						Limit: 10,
						Usage: 8,
					},
					ServergroupCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaServergroupCountUsage{
						Limit: 10,
						Usage: 8,
					},
					SfsCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountLimit{
						Limit: 10,
						Usage: 8,
					},
					SfsCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsCountUsage{
						Limit: 10,
						Usage: 8,
					},
					SfsSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					SfsSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSfsSizeUsage{
						Limit: 10,
						Usage: 8,
					},
					SharedVmCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountLimit{
						Limit: 10,
						Usage: 8,
					},
					SharedVmCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSharedVmCountUsage{
						Limit: 10,
						Usage: 8,
					},
					SnapshotScheduleCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountLimit{
						Limit: 10,
						Usage: 8,
					},
					SnapshotScheduleCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSnapshotScheduleCountUsage{
						Limit: 10,
						Usage: 8,
					},
					SubnetCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountLimit{
						Limit: 10,
						Usage: 8,
					},
					SubnetCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaSubnetCountUsage{
						Limit: 10,
						Usage: 8,
					},
					VmCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountLimit{
						Limit: 10,
						Usage: 8,
					},
					VmCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVmCountUsage{
						Limit: 10,
						Usage: 8,
					},
					VolumeCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountLimit{
						Limit: 10,
						Usage: 8,
					},
					VolumeCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeCountUsage{
						Limit: 10,
						Usage: 8,
					},
					VolumeSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					VolumeSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSizeUsage{
						Limit: 10,
						Usage: 8,
					},
					VolumeSnapshotsCountLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountLimit{
						Limit: 10,
						Usage: 8,
					},
					VolumeSnapshotsCountUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsCountUsage{
						Limit: 10,
						Usage: 8,
					},
					VolumeSnapshotsSizeLimit: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeLimit{
						Limit: 10,
						Usage: 8,
					},
					VolumeSnapshotsSizeUsage: cloud.QuotaNotificationThresholdUpdateParamsLastMessageRegionalQuotaVolumeSnapshotsSizeUsage{
						Limit: 10,
						Usage: 8,
					},
				}},
			},
			LastSending: param.Null[time.Time](),
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

func TestQuotaNotificationThresholdDelete(t *testing.T) {
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
	err := client.Cloud.Quotas.NotificationThreshold.Delete(context.TODO(), 3)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuotaNotificationThresholdGet(t *testing.T) {
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
	_, err := client.Cloud.Quotas.NotificationThreshold.Get(context.TODO(), 3)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
