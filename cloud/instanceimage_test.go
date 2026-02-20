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

func TestInstanceImageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.Update(
		context.TODO(),
		"image_id",
		cloud.InstanceImageUpdateParams{
			ProjectID:      gcore.Int(0),
			RegionID:       gcore.Int(0),
			HwFirmwareType: cloud.InstanceImageUpdateParamsHwFirmwareTypeBios,
			HwMachineType:  cloud.InstanceImageUpdateParamsHwMachineTypeQ35,
			IsBaremetal:    gcore.Bool(false),
			Name:           gcore.String("my-image"),
			OsType:         cloud.InstanceImageUpdateParamsOsTypeLinux,
			SSHKey:         cloud.InstanceImageUpdateParamsSSHKeyAllow,
			Tags:           map[string]any{},
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

func TestInstanceImageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.List(context.TODO(), cloud.InstanceImageListParams{
		ProjectID:     gcore.Int(0),
		RegionID:      gcore.Int(0),
		IncludePrices: gcore.Bool(true),
		Private:       gcore.String("private"),
		TagKey:        []string{"string"},
		TagKeyValue:   gcore.String("tag_key_value"),
		Visibility:    cloud.InstanceImageListParamsVisibilityPrivate,
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceImageDelete(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.Delete(
		context.TODO(),
		"image_id",
		cloud.InstanceImageDeleteParams{
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

func TestInstanceImageNewFromVolumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.NewFromVolume(context.TODO(), cloud.InstanceImageNewFromVolumeParams{
		ProjectID:      gcore.Int(0),
		RegionID:       gcore.Int(0),
		Name:           "my-image",
		VolumeID:       "d478ae29-dedc-4869-82f0-96104425f565",
		Architecture:   cloud.InstanceImageNewFromVolumeParamsArchitectureX86_64,
		HwFirmwareType: cloud.InstanceImageNewFromVolumeParamsHwFirmwareTypeBios,
		HwMachineType:  cloud.InstanceImageNewFromVolumeParamsHwMachineTypeQ35,
		IsBaremetal:    gcore.Bool(false),
		OsType:         cloud.InstanceImageNewFromVolumeParamsOsTypeLinux,
		Source:         cloud.InstanceImageNewFromVolumeParamsSourceVolume,
		SSHKey:         cloud.InstanceImageNewFromVolumeParamsSSHKeyAllow,
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

func TestInstanceImageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.Get(
		context.TODO(),
		"image_id",
		cloud.InstanceImageGetParams{
			ProjectID:     gcore.Int(0),
			RegionID:      gcore.Int(0),
			IncludePrices: gcore.Bool(true),
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

func TestInstanceImageUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Images.Upload(context.TODO(), cloud.InstanceImageUploadParams{
		ProjectID:      gcore.Int(0),
		RegionID:       gcore.Int(0),
		Name:           "my-image",
		URL:            "http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img",
		Architecture:   cloud.InstanceImageUploadParamsArchitectureX86_64,
		CowFormat:      gcore.Bool(false),
		HwFirmwareType: cloud.InstanceImageUploadParamsHwFirmwareTypeBios,
		HwMachineType:  cloud.InstanceImageUploadParamsHwMachineTypeQ35,
		IsBaremetal:    gcore.Bool(false),
		OsDistro:       gcore.String("ubuntu"),
		OsType:         cloud.InstanceImageUploadParamsOsTypeLinux,
		OsVersion:      gcore.String("22.04"),
		SSHKey:         cloud.InstanceImageUploadParamsSSHKeyAllow,
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
