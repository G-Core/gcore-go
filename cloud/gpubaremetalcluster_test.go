// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
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
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Flavor:    "bm3-ai-1xlarge-h100-80-8",
		ImageID:   "f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1",
		Interfaces: []cloud.GPUBaremetalClusterNewParamsInterfaceUnion{{
			OfSubnet: &cloud.GPUBaremetalClusterNewParamsInterfaceSubnet{
				NetworkID: "024a29e9-b4b7-4c91-9a46-505be123d9f8",
				SubnetID:  "91200a6c-07e0-42aa-98da-32d1f6545ae7",
				FloatingIP: cloud.GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion{
					OfNew: &cloud.GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew{},
				},
				InterfaceName: gcore.String("interface_name"),
				PortGroup:     gcore.Int(0),
				SecurityGroups: []cloud.GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup{{
					ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
				}},
			},
		}},
		Name:           "my-gpu-cluster",
		InstancesCount: gcore.Int(1),
		Password:       gcore.String("password"),
		SSHKeyName:     gcore.String("my-ssh-key"),
		Tags: cloud.TagUpdateList{
			"foo": "my-tag-value",
		},
		UserData: gcore.String("user_data"),
		Username: gcore.String("username"),
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
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Limit:     gcore.Int(0),
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
		"cluster_id",
		cloud.GPUBaremetalClusterDeleteParams{
			ProjectID:        gcore.Int(0),
			RegionID:         gcore.Int(0),
			DeleteFloatings:  gcore.Bool(true),
			Floatings:        gcore.String("floatings"),
			ReservedFixedIPs: gcore.String("reserved_fixed_ips"),
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
		"cluster_id",
		cloud.GPUBaremetalClusterGetParams{
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
