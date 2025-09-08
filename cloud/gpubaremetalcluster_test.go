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

func TestGPUBaremetalClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.New(context.TODO(), cloud.GPUBaremetalClusterNewParams{
		ProjectID:    gcore.Int(1),
		RegionID:     gcore.Int(7),
		Flavor:       "g3-ai-32-192-1500-l40s-48-1",
		ImageID:      "3793c250-0b3b-4678-bab3-e11afbc29657",
		Name:         "gpu-cluster-1",
		ServersCount: 3,
		ServersSettings: cloud.GPUBaremetalClusterNewParamsServersSettings{
			Interfaces: []cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion{{
				OfExternal: &cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal{
					IPFamily: "ipv4",
					Name:     gcore.String("eth0"),
				},
			}},
			Credentials: cloud.GPUBaremetalClusterNewParamsServersSettingsCredentials{
				Password:   gcore.String("securepassword"),
				SSHKeyName: gcore.String("my-ssh-key"),
				Username:   gcore.String("admin"),
			},
			SecurityGroups: []cloud.GPUBaremetalClusterNewParamsServersSettingsSecurityGroup{{
				ID: "b4849ffa-89f2-45a1-951f-0ae5b7809d98",
			}},
			UserData: gcore.String("eyJ0ZXN0IjogImRhdGEifQ=="),
		},
		Tags: map[string]string{
			"my-tag": "my-tag-value",
		},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGPUBaremetalClusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.List(context.TODO(), cloud.GPUBaremetalClusterListParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(7),
		Limit:     gcore.Int(10),
		ManagedBy: []string{"k8s"},
		Offset:    gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGPUBaremetalClusterDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Delete(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUBaremetalClusterDeleteParams{
			ProjectID:           gcore.Int(1),
			RegionID:            gcore.Int(7),
			AllFloatingIPs:      gcore.Bool(true),
			AllReservedFixedIPs: gcore.Bool(true),
			FloatingIPIDs:       []string{"e4a01208-d6ac-4304-bf86-3028154b070a"},
			ReservedFixedIPIDs:  []string{"a29b8e1e-08d3-4cec-91fb-06e81e5f46d5"},
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

func TestGPUBaremetalClusterGet(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Get(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUBaremetalClusterGetParams{
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

func TestGPUBaremetalClusterPowercycleAllServers(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.PowercycleAllServers(
		context.TODO(),
		"cluster_id",
		cloud.GPUBaremetalClusterPowercycleAllServersParams{
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

func TestGPUBaremetalClusterRebootAllServers(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.RebootAllServers(
		context.TODO(),
		"cluster_id",
		cloud.GPUBaremetalClusterRebootAllServersParams{
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

func TestGPUBaremetalClusterRebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Rebuild(
		context.TODO(),
		"cluster_id",
		cloud.GPUBaremetalClusterRebuildParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Nodes:     []string{"string"},
			ImageID:   gcore.String("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			UserData:  gcore.String("user_data"),
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

func TestGPUBaremetalClusterResize(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Resize(
		context.TODO(),
		"cluster_id",
		cloud.GPUBaremetalClusterResizeParams{
			ProjectID:      gcore.Int(0),
			RegionID:       gcore.Int(0),
			InstancesCount: 1,
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
