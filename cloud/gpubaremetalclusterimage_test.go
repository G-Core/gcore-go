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

func TestGPUBaremetalClusterImageList(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Images.List(context.TODO(), cloud.GPUBaremetalClusterImageListParams{
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

func TestGPUBaremetalClusterImageDelete(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Images.Delete(
		context.TODO(),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
		cloud.GPUBaremetalClusterImageDeleteParams{
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

func TestGPUBaremetalClusterImageGet(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Images.Get(
		context.TODO(),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
		cloud.GPUBaremetalClusterImageGetParams{
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

func TestGPUBaremetalClusterImageUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Images.Upload(context.TODO(), cloud.GPUBaremetalClusterImageUploadParams{
		ProjectID:      gcore.Int(1),
		RegionID:       gcore.Int(7),
		Name:           "ubuntu-23.10-x64",
		URL:            "http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img",
		Architecture:   cloud.GPUBaremetalClusterImageUploadParamsArchitectureX86_64,
		CowFormat:      gcore.Bool(true),
		HwFirmwareType: cloud.GPUBaremetalClusterImageUploadParamsHwFirmwareTypeBios,
		OsDistro:       gcore.String("os_distro"),
		OsType:         cloud.GPUBaremetalClusterImageUploadParamsOsTypeLinux,
		OsVersion:      gcore.String("19.04"),
		SSHKey:         cloud.GPUBaremetalClusterImageUploadParamsSSHKeyAllow,
		Tags: cloud.TagUpdateList{
			"foo": "my-tag-value",
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
