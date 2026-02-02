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
	"github.com/G-Core/gcore-go/packages/param"
)

func TestK8SClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.New(context.TODO(), cloud.K8SClusterNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(7),
		Keypair:   "some_keypair",
		Name:      "string",
		Pools: []cloud.K8SClusterNewParamsPool{{
			FlavorID:           "g1-standard-1-2",
			MinNodeCount:       3,
			Name:               "my-pool",
			AutoHealingEnabled: gcore.Bool(true),
			BootVolumeSize:     gcore.Int(50),
			BootVolumeType:     "ssd_hiiops",
			CrioConfig: map[string]string{
				"default-ulimits": "nofile=1024:2048",
			},
			IsPublicIpv4: gcore.Bool(true),
			KubeletConfig: map[string]string{
				"podMaxPids": "4096",
			},
			Labels: map[string]string{
				"my-label": "foo",
			},
			MaxNodeCount:      gcore.Int(5),
			ServergroupPolicy: "affinity",
			Taints: map[string]string{
				"my-taint": "bar:NoSchedule",
			},
		}},
		Version: "1.28.1",
		AddOns: cloud.K8SClusterNewParamsAddOns{
			Slurm: cloud.K8SClusterNewParamsAddOnsSlurm{
				Enabled:     true,
				FileShareID: "cbc94d0e-06c6-4d12-9e86-9782ba14fc8c",
				SSHKeyIDs:   []string{"25735292-bd97-44b0-a1af-d7eab876261d", "efc01f3a-35b9-4385-89f9-e38439093ee7"},
				WorkerCount: 2,
			},
		},
		Authentication: cloud.K8SClusterNewParamsAuthentication{
			Oidc: cloud.K8SClusterNewParamsAuthenticationOidc{
				ClientID:     gcore.String("kubernetes"),
				GroupsClaim:  gcore.String("groups"),
				GroupsPrefix: gcore.String("oidc:"),
				IssuerURL:    gcore.String("https://accounts.provider.example"),
				RequiredClaims: map[string]string{
					"claim": "value",
				},
				SigningAlgs:    []string{"RS256", "RS512"},
				UsernameClaim:  gcore.String("sub"),
				UsernamePrefix: gcore.String("oidc:"),
			},
		},
		AutoscalerConfig: map[string]string{
			"scale-down-unneeded-time": "5m",
		},
		Cni: cloud.K8SClusterNewParamsCni{
			Cilium: cloud.K8SClusterNewParamsCniCilium{
				Encryption:     gcore.Bool(true),
				HubbleRelay:    gcore.Bool(true),
				HubbleUi:       gcore.Bool(true),
				LbAcceleration: gcore.Bool(true),
				LbMode:         "snat",
				MaskSize:       gcore.Int(24),
				MaskSizeV6:     gcore.Int(120),
				RoutingMode:    "tunnel",
				Tunnel:         "geneve",
			},
			Provider: "cilium",
		},
		Csi: cloud.K8SClusterNewParamsCsi{
			Nfs: cloud.K8SClusterNewParamsCsiNfs{
				VastEnabled: gcore.Bool(true),
			},
		},
		DDOSProfile: cloud.K8SClusterNewParamsDDOSProfile{
			Enabled: true,
			Fields: []cloud.K8SClusterNewParamsDDOSProfileField{{
				BaseField: 10,
				FieldValue: []float64{
					45046,
					45047,
				},
				Value: param.Null[string](),
			}},
			ProfileTemplate:     gcore.Int(29),
			ProfileTemplateName: gcore.String("profile_template_name"),
		},
		FixedNetwork: gcore.String("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		FixedSubnet:  gcore.String("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		IsIpv6:       gcore.Bool(true),
		Logging: cloud.K8SClusterNewParamsLogging{
			DestinationRegionID: gcore.Int(1),
			Enabled:             gcore.Bool(true),
			RetentionPolicy: cloud.LaasIndexRetentionPolicyParam{
				Period: gcore.Int(45),
			},
			TopicName: gcore.String("my-log-name"),
		},
		PodsIPPool:       gcore.String("172.16.0.0/18"),
		PodsIpv6Pool:     gcore.String("2a03:90c0:88:393::/64"),
		ServicesIPPool:   gcore.String("172.24.0.0/18"),
		ServicesIpv6Pool: gcore.String("2a03:90c0:88:381::/108"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestK8SClusterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Update(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			AddOns: cloud.K8SClusterUpdateParamsAddOns{
				Slurm: cloud.K8SClusterUpdateParamsAddOnsSlurmUnion{
					OfK8SClusterSlurmAddonEnableV2Serializer: &cloud.K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer{
						Enabled:     true,
						FileShareID: "cbc94d0e-06c6-4d12-9e86-9782ba14fc8c",
						SSHKeyIDs:   []string{"25735292-bd97-44b0-a1af-d7eab876261d", "efc01f3a-35b9-4385-89f9-e38439093ee7"},
						WorkerCount: 2,
					},
				},
			},
			Authentication: cloud.K8SClusterUpdateParamsAuthentication{
				Oidc: cloud.K8SClusterUpdateParamsAuthenticationOidc{
					ClientID:     gcore.String("kubernetes"),
					GroupsClaim:  gcore.String("groups"),
					GroupsPrefix: gcore.String("oidc:"),
					IssuerURL:    gcore.String("https://accounts.provider.example"),
					RequiredClaims: map[string]string{
						"claim": "value",
					},
					SigningAlgs:    []string{"RS256", "RS512"},
					UsernameClaim:  gcore.String("sub"),
					UsernamePrefix: gcore.String("oidc:"),
				},
			},
			AutoscalerConfig: map[string]string{
				"scale-down-unneeded-time": "5m",
			},
			Cni: cloud.K8SClusterUpdateParamsCni{
				Cilium: cloud.K8SClusterUpdateParamsCniCilium{
					Encryption:     gcore.Bool(true),
					HubbleRelay:    gcore.Bool(true),
					HubbleUi:       gcore.Bool(true),
					LbAcceleration: gcore.Bool(true),
					LbMode:         "snat",
					MaskSize:       gcore.Int(24),
					MaskSizeV6:     gcore.Int(120),
					RoutingMode:    "tunnel",
					Tunnel:         "geneve",
				},
				Provider: "cilium",
			},
			DDOSProfile: cloud.K8SClusterUpdateParamsDDOSProfile{
				Enabled: true,
				Fields: []cloud.K8SClusterUpdateParamsDDOSProfileField{{
					BaseField: 10,
					FieldValue: []float64{
						45046,
						45047,
					},
					Value: param.Null[string](),
				}},
				ProfileTemplate:     gcore.Int(29),
				ProfileTemplateName: gcore.String("profile_template_name"),
			},
			Logging: cloud.K8SClusterUpdateParamsLogging{
				DestinationRegionID: gcore.Int(1),
				Enabled:             gcore.Bool(true),
				RetentionPolicy: cloud.LaasIndexRetentionPolicyParam{
					Period: gcore.Int(45),
				},
				TopicName: gcore.String("my-log-name"),
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

func TestK8SClusterList(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.List(context.TODO(), cloud.K8SClusterListParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(7),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestK8SClusterDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Delete(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			Volumes:   gcore.String("550e8400-e29b-41d4-a716-446655440000"),
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

func TestK8SClusterGet(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Get(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
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

func TestK8SClusterGetCertificate(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.GetCertificate(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterGetCertificateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
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

func TestK8SClusterGetKubeconfig(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.GetKubeconfig(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterGetKubeconfigParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
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

func TestK8SClusterListVersionsForUpgrade(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.ListVersionsForUpgrade(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterListVersionsForUpgradeParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
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

func TestK8SClusterUpgrade(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Upgrade(
		context.TODO(),
		"my-cluster",
		cloud.K8SClusterUpgradeParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			Version:   "v1.28.1",
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
