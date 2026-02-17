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

func TestGPUVirtualClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.New(context.TODO(), cloud.GPUVirtualClusterNewParams{
		ProjectID:    gcore.Int(1),
		RegionID:     gcore.Int(7),
		Flavor:       "g3-ai-32-192-1500-l40s-48-1",
		Name:         "gpu-cluster-1",
		ServersCount: 3,
		ServersSettings: cloud.GPUVirtualClusterNewParamsServersSettings{
			Interfaces: []cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceUnion{{
				OfExternal: &cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceExternal{
					IPFamily: "ipv4",
					Name:     gcore.String("eth0"),
				},
			}},
			Volumes: []cloud.GPUVirtualClusterNewParamsServersSettingsVolumeUnion{{
				OfNew: &cloud.GPUVirtualClusterNewParamsServersSettingsVolumeNew{
					BootIndex:           1,
					Name:                "my-data-disk",
					Size:                100,
					Type:                "cold",
					DeleteOnTermination: gcore.Bool(true),
					Tags: map[string]string{
						"key1": "value1",
					},
				},
			}},
			Credentials: cloud.GPUVirtualClusterNewParamsServersSettingsCredentials{
				Password:   gcore.String("securepassword"),
				SSHKeyName: gcore.String("my-ssh-key"),
				Username:   gcore.String("admin"),
			},
			FileShares: []cloud.GPUVirtualClusterNewParamsServersSettingsFileShare{{
				ID:        "a3f2d1b8-45e6-4f8a-bb5d-19dbf2cd7e9a",
				MountPath: "/mnt/vast",
			}},
			SecurityGroups: []cloud.GPUVirtualClusterNewParamsServersSettingsSecurityGroup{{
				ID: "b4849ffa-89f2-45a1-951f-0ae5b7809d98",
			}},
			UserData: gcore.String("eyJ0ZXN0IjogImRhdGEifQ=="),
		},
		Tags: map[string]any{
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

func TestGPUVirtualClusterUpdate(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Update(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUVirtualClusterUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			Name:      "gpu-cluster-1",
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

func TestGPUVirtualClusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.List(context.TODO(), cloud.GPUVirtualClusterListParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(7),
		Limit:     gcore.Int(10),
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

func TestGPUVirtualClusterDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Delete(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUVirtualClusterDeleteParams{
			ProjectID:           gcore.Int(1),
			RegionID:            gcore.Int(7),
			AllFloatingIPs:      gcore.Bool(true),
			AllReservedFixedIPs: gcore.Bool(true),
			AllVolumes:          gcore.Bool(true),
			FloatingIPIDs:       []string{"e4a01208-d6ac-4304-bf86-3028154b070a"},
			ReservedFixedIPIDs:  []string{"a29b8e1e-08d3-4cec-91fb-06e81e5f46d5"},
			VolumeIDs:           []string{"1333c684-c3da-4b91-ac9e-a92706aa2824"},
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

func TestGPUVirtualClusterAction(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Action(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUVirtualClusterActionParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			OfStart:   &cloud.GPUVirtualClusterActionParamsBodyStart{},
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

func TestGPUVirtualClusterGet(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Get(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUVirtualClusterGetParams{
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
