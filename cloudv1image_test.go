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

func TestCloudV1ImageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Images.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1ImageNewParams{
			ImageCreate: gcore.ImageCreateParam{
				Name:           gcore.F("test_image"),
				VolumeID:       gcore.F("d478ae29-dedc-4869-82f0-96104425f565"),
				Architecture:   gcore.F(gcore.ImageCreateArchitectureAarch64),
				HwFirmwareType: gcore.F(gcore.ImageCreateHwFirmwareTypeBios),
				HwMachineType:  gcore.F(gcore.ImageCreateHwMachineTypeI440),
				IsBaremetal:    gcore.F(false),
				Metadata: gcore.F[any](map[string]interface{}{
					"key": "value",
				}),
				OsType: gcore.F(gcore.ImageCreateOsTypeLinux),
				Source: gcore.F("volume"),
				SSHKey: gcore.F(gcore.ImageCreateSSHKeyAllow),
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

func TestCloudV1ImageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Images.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"image_id",
		gcore.CloudV1ImageGetParams{
			IncludePrices: gcore.F(true),
			MetadataK:     gcore.F("metadata_k"),
			MetadataKv:    gcore.F("metadata_kv"),
			Private:       gcore.F("private"),
			Visibility:    gcore.F(gcore.CloudV1ImageGetParamsVisibilityPrivate),
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

func TestCloudV1ImageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Images.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"image_id",
		gcore.CloudV1ImageUpdateParams{
			HwFirmwareType: gcore.F(gcore.CloudV1ImageUpdateParamsHwFirmwareTypeBios),
			HwMachineType:  gcore.F(gcore.CloudV1ImageUpdateParamsHwMachineTypeI440),
			IsBaremetal:    gcore.F(true),
			Metadata:       gcore.F[any](map[string]interface{}{}),
			Name:           gcore.F("string"),
			OsType:         gcore.F(gcore.CloudV1ImageUpdateParamsOsTypeLinux),
			SSHKey:         gcore.F(gcore.CloudV1ImageUpdateParamsSSHKeyAllow),
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

func TestCloudV1ImageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Images.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1ImageListParams{
			IncludePrices: gcore.F(true),
			MetadataK:     gcore.F("metadata_k"),
			MetadataKv:    gcore.F("metadata_kv"),
			Private:       gcore.F("private"),
			Visibility:    gcore.F(gcore.CloudV1ImageListParamsVisibilityPrivate),
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

func TestCloudV1ImageDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Images.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"image_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
