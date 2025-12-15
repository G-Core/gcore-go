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

func TestK8SClusterPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.New(
		context.TODO(),
		"cluster_name",
		cloud.K8SClusterPoolNewParams{
			ProjectID:          gcore.Int(0),
			RegionID:           gcore.Int(0),
			FlavorID:           "g1-standard-1-2",
			MinNodeCount:       3,
			Name:               "my-pool",
			AutoHealingEnabled: gcore.Bool(true),
			BootVolumeSize:     gcore.Int(50),
			BootVolumeType:     cloud.K8SClusterPoolNewParamsBootVolumeTypeSsdHiiops,
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
			ServergroupPolicy: cloud.K8SClusterPoolNewParamsServergroupPolicyAffinity,
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

func TestK8SClusterPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.Update(
		context.TODO(),
		"pool_name",
		cloud.K8SClusterPoolUpdateParams{
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

func TestK8SClusterPoolList(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.List(
		context.TODO(),
		"cluster_name",
		cloud.K8SClusterPoolListParams{
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

func TestK8SClusterPoolDelete(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.Delete(
		context.TODO(),
		"pool_name",
		cloud.K8SClusterPoolDeleteParams{
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

func TestK8SClusterPoolCheckQuotaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.CheckQuota(context.TODO(), cloud.K8SClusterPoolCheckQuotaParams{
		ProjectID:         gcore.Int(1),
		RegionID:          gcore.Int(7),
		FlavorID:          "g1-standard-1-2",
		BootVolumeSize:    gcore.Int(50),
		MaxNodeCount:      gcore.Int(5),
		MinNodeCount:      gcore.Int(3),
		Name:              gcore.String("test"),
		NodeCount:         gcore.Int(5),
		ServergroupPolicy: cloud.K8SClusterPoolCheckQuotaParamsServergroupPolicyAntiAffinity,
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestK8SClusterPoolGet(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.Get(
		context.TODO(),
		"pool_name",
		cloud.K8SClusterPoolGetParams{
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

func TestK8SClusterPoolResize(t *testing.T) {
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
	_, err := client.Cloud.K8S.Clusters.Pools.Resize(
		context.TODO(),
		"pool_name",
		cloud.K8SClusterPoolResizeParams{
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
