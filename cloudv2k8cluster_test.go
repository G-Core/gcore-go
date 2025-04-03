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

func TestCloudV2K8ClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV2K8ClusterNewParams{
			Keypair: gcore.F("some_keypair"),
			Name:    gcore.F("string"),
			Pools: gcore.F([]gcore.PoolCreateParam{{
				FlavorID:           gcore.F("g1-standard-1-2"),
				MinNodeCount:       gcore.F(int64(3)),
				Name:               gcore.F("my-pool"),
				AutoHealingEnabled: gcore.F(true),
				BootVolumeSize:     gcore.F(int64(50)),
				BootVolumeType:     gcore.F(gcore.VolumeTypeNameCold),
				CrioConfig: gcore.F(map[string]string{
					"default-ulimits": "nofile=1024:2048",
				}),
				IsPublicIpv4: gcore.F(true),
				KubeletConfig: gcore.F(map[string]string{
					"podMaxPids": "4096",
				}),
				Labels: gcore.F(map[string]string{
					"my-label": "foo",
				}),
				MaxNodeCount:      gcore.F(int64(5)),
				ServergroupPolicy: gcore.F(gcore.ServerGroupPolicyAffinity),
				Taints: gcore.F(map[string]string{
					"my-taint": "bar:NoSchedule",
				}),
			}}),
			Version: gcore.F("1.28.1"),
			Authentication: gcore.F(gcore.AuthenticationCreateParam{
				Oidc: gcore.F(gcore.OidcParam{
					ClientID:     gcore.F("kubernetes"),
					GroupsClaim:  gcore.F("groups"),
					GroupsPrefix: gcore.F("oidc:"),
					IssuerURL:    gcore.F("https://accounts.provider.example"),
					RequiredClaims: gcore.F(map[string]string{
						"claim": "value",
					}),
					SigningAlgs:    gcore.F([]gcore.OidcSigningAlg{gcore.OidcSigningAlgEs256}),
					UsernameClaim:  gcore.F("sub"),
					UsernamePrefix: gcore.F("oidc:"),
				}),
			}),
			AutoscalerConfig: gcore.F(map[string]string{
				"scale-down-unneeded-time": "5m",
			}),
			Cni: gcore.F(gcore.CniCreateParam{
				Cilium: gcore.F(gcore.CiliumParam{
					Encryption:     gcore.F(true),
					HubbleRelay:    gcore.F(true),
					HubbleUi:       gcore.F(true),
					LbAcceleration: gcore.F(true),
					LbMode:         gcore.F(gcore.CiliumLbModeDsr),
					MaskSize:       gcore.F(int64(24)),
					MaskSizeV6:     gcore.F(int64(120)),
					RoutingMode:    gcore.F(gcore.CiliumRoutingModeNative),
					Tunnel:         gcore.F(gcore.CiliumTunnelEmpty),
				}),
				Provider: gcore.F(gcore.CniCalico),
			}),
			DdosProfile: gcore.F(gcore.CloudV2K8ClusterNewParamsDdosProfile{
				Enabled: gcore.F(true),
				Fields: gcore.F([]gcore.DdosProfileParam{{
					BaseField: gcore.F(int64(10)),
					FieldValue: gcore.F[any](map[string]interface{}{
						"0": 45046,
						"1": 45047,
					}),
					Value: gcore.F("value"),
				}}),
				ProfileTemplate:     gcore.F(int64(29)),
				ProfileTemplateName: gcore.F("profile_template_name"),
			}),
			FixedNetwork: gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			FixedSubnet:  gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			IsIpv6:       gcore.F(true),
			Logging: gcore.F(gcore.CloudV2K8ClusterNewParamsLogging{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			PodsIPPool:       gcore.F("172.16.0.0/18"),
			PodsIpv6Pool:     gcore.F("2a03:90c0:88:393::/64"),
			ServicesIPPool:   gcore.F("172.24.0.0/18"),
			ServicesIpv6Pool: gcore.F("2a03:90c0:88:381::/108"),
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

func TestCloudV2K8ClusterGet(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		gcore.CloudV2K8ClusterUpdateParams{
			Authentication: gcore.F(gcore.AuthenticationCreateParam{
				Oidc: gcore.F(gcore.OidcParam{
					ClientID:     gcore.F("kubernetes"),
					GroupsClaim:  gcore.F("groups"),
					GroupsPrefix: gcore.F("oidc:"),
					IssuerURL:    gcore.F("https://accounts.provider.example"),
					RequiredClaims: gcore.F(map[string]string{
						"claim": "value",
					}),
					SigningAlgs:    gcore.F([]gcore.OidcSigningAlg{gcore.OidcSigningAlgEs256}),
					UsernameClaim:  gcore.F("sub"),
					UsernamePrefix: gcore.F("oidc:"),
				}),
			}),
			AutoscalerConfig: gcore.F(map[string]string{
				"scale-down-unneeded-time": "5m",
			}),
			Cni: gcore.F(gcore.CniCreateParam{
				Cilium: gcore.F(gcore.CiliumParam{
					Encryption:     gcore.F(true),
					HubbleRelay:    gcore.F(true),
					HubbleUi:       gcore.F(true),
					LbAcceleration: gcore.F(true),
					LbMode:         gcore.F(gcore.CiliumLbModeDsr),
					MaskSize:       gcore.F(int64(24)),
					MaskSizeV6:     gcore.F(int64(120)),
					RoutingMode:    gcore.F(gcore.CiliumRoutingModeNative),
					Tunnel:         gcore.F(gcore.CiliumTunnelEmpty),
				}),
				Provider: gcore.F(gcore.CniCalico),
			}),
			DdosProfile: gcore.F(gcore.CloudV2K8ClusterUpdateParamsDdosProfile{
				Enabled: gcore.F(true),
				Fields: gcore.F([]gcore.DdosProfileParam{{
					BaseField: gcore.F(int64(10)),
					FieldValue: gcore.F[any](map[string]interface{}{
						"0": 45046,
						"1": 45047,
					}),
					Value: gcore.F("value"),
				}}),
				ProfileTemplate:     gcore.F(int64(29)),
				ProfileTemplateName: gcore.F("profile_template_name"),
			}),
			Logging: gcore.F(gcore.CloudV2K8ClusterUpdateParamsLogging{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
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

func TestCloudV2K8ClusterList(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.List(
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

func TestCloudV2K8ClusterDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		gcore.CloudV2K8ClusterDeleteParams{
			Volumes: gcore.F("volumes"),
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

func TestCloudV2K8ClusterCheckLimitsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.CheckLimits(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV2K8ClusterCheckLimitsParams{
			Logging: gcore.F(gcore.CloudV2K8ClusterCheckLimitsParamsLogging{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			Pools: gcore.F([]gcore.CloudV2K8ClusterCheckLimitsParamsPool{{
				FlavorID:          gcore.F("g1-standard-1-2"),
				BootVolumeSize:    gcore.F(int64(50)),
				MaxNodeCount:      gcore.F(int64(5)),
				MinNodeCount:      gcore.F(int64(3)),
				Name:              gcore.F("test"),
				NodeCount:         gcore.F(int64(5)),
				ServergroupPolicy: gcore.F(gcore.ServerGroupPolicyAffinity),
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

func TestCloudV2K8ClusterGetCertificates(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.GetCertificates(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterGetConfig(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.GetConfig(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterGetUpgradeVersions(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.GetUpgradeVersions(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterListInstancesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.ListInstances(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		gcore.CloudV2K8ClusterListInstancesParams{
			WithDdos: gcore.F(true),
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

func TestCloudV2K8ClusterUpgrade(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Upgrade(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		gcore.CloudV2K8ClusterUpgradeParams{
			Version: gcore.F("v1.28.1"),
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
