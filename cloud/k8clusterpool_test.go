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

func TestK8ClusterPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.New(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterPoolNewParams{
			ProjectID:          gcore.Int(0),
			RegionID:           gcore.Int(0),
			FlavorID:           "g1-standard-1-2",
			MinNodeCount:       3,
			Name:               "my-pool",
			AutoHealingEnabled: gcore.Bool(true),
			BootVolumeSize:     gcore.Int(50),
			BootVolumeType:     cloud.K8ClusterPoolNewParamsBootVolumeTypeSsdHiiops,
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
			ServergroupPolicy: cloud.K8ClusterPoolNewParamsServergroupPolicyAffinity,
			Taints: map[string]string{
				"my-taint": "bar:NoSchedule",
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

func TestK8ClusterPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.Update(
		context.TODO(),
		"pool_name",
		cloud.K8ClusterPoolUpdateParams{
			ProjectID:          gcore.Int(0),
			RegionID:           gcore.Int(0),
			ClusterName:        "cluster_name",
			AutoHealingEnabled: gcore.Bool(true),
			Labels: map[string]string{
				"my-label": "foo",
			},
			MaxNodeCount: gcore.Int(3),
			MinNodeCount: gcore.Int(1),
			NodeCount:    gcore.Int(2),
			Taints: map[string]string{
				"my-taint": "bar:NoSchedule",
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

func TestK8ClusterPoolList(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.List(
		context.TODO(),
		"cluster_name",
		cloud.K8ClusterPoolListParams{
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

func TestK8ClusterPoolDelete(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.Delete(
		context.TODO(),
		"pool_name",
		cloud.K8ClusterPoolDeleteParams{
			ProjectID:   gcore.Int(0),
			RegionID:    gcore.Int(0),
			ClusterName: "cluster_name",
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

func TestK8ClusterPoolGet(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.Get(
		context.TODO(),
		"pool_name",
		cloud.K8ClusterPoolGetParams{
			ProjectID:   gcore.Int(0),
			RegionID:    gcore.Int(0),
			ClusterName: "cluster_name",
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

func TestK8ClusterPoolResize(t *testing.T) {
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
	_, err := client.Cloud.K8s.Clusters.Pools.Resize(
		context.TODO(),
		"pool_name",
		cloud.K8ClusterPoolResizeParams{
			ProjectID:   gcore.Int(0),
			RegionID:    gcore.Int(0),
			ClusterName: "cluster_name",
			NodeCount:   2,
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
