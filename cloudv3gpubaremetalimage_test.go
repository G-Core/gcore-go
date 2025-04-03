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

func TestCloudV3GPUBaremetalImageGet(t *testing.T) {
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
	_, err := client.Cloud.V3.GPU.Baremetal.Images.Get(
		context.TODO(),
		int64(1),
		int64(7),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3GPUBaremetalImageList(t *testing.T) {
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
	_, err := client.Cloud.V3.GPU.Baremetal.Images.List(
		context.TODO(),
		int64(1),
		int64(7),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3GPUBaremetalImageDelete(t *testing.T) {
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
	_, err := client.Cloud.V3.GPU.Baremetal.Images.Delete(
		context.TODO(),
		int64(1),
		int64(7),
		"8cab6f28-09ca-4201-b3f7-23c7893f4bd6",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3GPUBaremetalImageUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V3.GPU.Baremetal.Images.Upload(
		context.TODO(),
		int64(1),
		int64(7),
		gcore.CloudV3GPUBaremetalImageUploadParams{
			Name:           gcore.F("ubuntu-23.10-x64"),
			URL:            gcore.F("http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img"),
			Architecture:   gcore.F(gcore.ImageArchitectureTypeAarch64),
			CowFormat:      gcore.F(true),
			HwFirmwareType: gcore.F(gcore.ImageHwFirmwareTypeBios),
			Metadata: gcore.F[any](map[string]interface{}{
				"key": "value",
			}),
			OsDistro:  gcore.F("os_distro"),
			OsType:    gcore.F(gcore.ImageOsTypeLinux),
			OsVersion: gcore.F("19.04"),
			SSHKey:    gcore.F(gcore.SSHKeyAllow),
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
