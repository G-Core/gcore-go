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

func TestCloudV2K8ClusterPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.New(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		gcore.CloudV2K8ClusterPoolNewParams{
			PoolCreate: gcore.PoolCreateParam{
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

func TestCloudV2K8ClusterPoolGet(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		"pool_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		"pool_name",
		gcore.CloudV2K8ClusterPoolUpdateParams{
			AutoHealingEnabled: gcore.F(true),
			Labels: gcore.F(map[string]string{
				"my-label": "foo",
			}),
			MaxNodeCount: gcore.F(int64(3)),
			MinNodeCount: gcore.F(int64(1)),
			NodeCount:    gcore.F(int64(2)),
			Taints: gcore.F(map[string]string{
				"my-taint": "bar:NoSchedule",
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

func TestCloudV2K8ClusterPoolList(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.List(
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

func TestCloudV2K8ClusterPoolDelete(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		"pool_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2K8ClusterPoolListInstancesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.ListInstances(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		"pool_name",
		gcore.CloudV2K8ClusterPoolListInstancesParams{
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

func TestCloudV2K8ClusterPoolResize(t *testing.T) {
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
	_, err := client.Cloud.V2.K8s.Clusters.Pools.Resize(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
		"pool_name",
		gcore.CloudV2K8ClusterPoolResizeParams{
			NodeCount: gcore.F(int64(2)),
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
