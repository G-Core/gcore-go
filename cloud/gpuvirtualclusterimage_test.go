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

func TestGPUVirtualClusterImageList(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtualClusters.Images.List(context.TODO(), cloud.GPUVirtualClusterImageListParams{
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

func TestGPUVirtualClusterImageDelete(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtualClusters.Images.Delete(
		context.TODO(),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
		cloud.GPUVirtualClusterImageDeleteParams{
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

func TestGPUVirtualClusterImageGet(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtualClusters.Images.Get(
		context.TODO(),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
		cloud.GPUVirtualClusterImageGetParams{
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

func TestGPUVirtualClusterImageUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtualClusters.Images.Upload(context.TODO(), cloud.GPUVirtualClusterImageUploadParams{
		ProjectID:      gcore.Int(1),
		RegionID:       gcore.Int(7),
		Name:           "ubuntu-23.10-x64",
		URL:            "http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img",
		Architecture:   cloud.GPUVirtualClusterImageUploadParamsArchitectureX86_64,
		CowFormat:      gcore.Bool(true),
		HwFirmwareType: cloud.GPUVirtualClusterImageUploadParamsHwFirmwareTypeBios,
		OsDistro:       gcore.String("os_distro"),
		OsType:         cloud.GPUVirtualClusterImageUploadParamsOsTypeLinux,
		OsVersion:      gcore.String("19.04"),
		SSHKey:         cloud.GPUVirtualClusterImageUploadParamsSSHKeyAllow,
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
