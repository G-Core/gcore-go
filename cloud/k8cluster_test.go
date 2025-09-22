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

func TestK8ClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.New(context.TODO(), cloud.K8ClusterNewParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Keypair:   "some_keypair",
		Name:      "string",
		Pools: []cloud.K8ClusterNewParamsPool{{
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
		Authentication: cloud.K8ClusterNewParamsAuthentication{
			Oidc: cloud.K8ClusterNewParamsAuthenticationOidc{
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
		Cni: cloud.K8ClusterNewParamsCni{
			Cilium: cloud.K8ClusterNewParamsCniCilium{
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
		Csi: cloud.K8ClusterNewParamsCsi{
			Nfs: cloud.K8ClusterNewParamsCsiNfs{
				VastEnabled: gcore.Bool(true),
			},
		},
		DDOSProfile: cloud.K8ClusterNewParamsDDOSProfile{
			Enabled: true,
			Fields: []cloud.K8ClusterNewParamsDDOSProfileField{{
				BaseField: 10,
				FieldValue: map[string]interface{}{
					"0": 45046,
					"1": 45047,
				},
				Value: param.Null[string](),
			}},
			ProfileTemplate:     gcore.Int(29),
			ProfileTemplateName: gcore.String("profile_template_name"),
		},
		FixedNetwork: gcore.String("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		FixedSubnet:  gcore.String("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		IsIpv6:       gcore.Bool(true),
		Logging: cloud.K8ClusterNewParamsLogging{
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

func TestK8ClusterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Update(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterUpdateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Authentication: cloud.K8ClusterUpdateParamsAuthentication{
				Oidc: cloud.K8ClusterUpdateParamsAuthenticationOidc{
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
			Cni: cloud.K8ClusterUpdateParamsCni{
				Cilium: cloud.K8ClusterUpdateParamsCniCilium{
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
			DDOSProfile: cloud.K8ClusterUpdateParamsDDOSProfile{
				Enabled: true,
				Fields: []cloud.K8ClusterUpdateParamsDDOSProfileField{{
					BaseField: 10,
					FieldValue: map[string]interface{}{
						"0": 45046,
						"1": 45047,
					},
					Value: param.Null[string](),
				}},
				ProfileTemplate:     gcore.Int(29),
				ProfileTemplateName: gcore.String("profile_template_name"),
			},
			Logging: cloud.K8ClusterUpdateParamsLogging{
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

func TestK8ClusterList(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.List(context.TODO(), cloud.K8ClusterListParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestK8ClusterDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Delete(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterDeleteParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Volumes:   gcore.String("volumes"),
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

func TestK8ClusterGet(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Get(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterGetParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestK8ClusterGetCertificate(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.GetCertificate(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterGetCertificateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestK8ClusterGetKubeconfig(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.GetKubeconfig(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterGetKubeconfigParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestK8ClusterListVersionsForUpgrade(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.ListVersionsForUpgrade(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterListVersionsForUpgradeParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestK8ClusterUpgrade(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Upgrade(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterUpgradeParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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
