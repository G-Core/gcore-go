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

func TestCloudV1InstanceAvailableFlavorGetAvailableFlavorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Instances.AvailableFlavors.GetAvailableFlavors(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1InstanceAvailableFlavorGetAvailableFlavorsParams{
			Volumes: gcore.F([]gcore.CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolume{{
				Source:              gcore.F(gcore.CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesSourceApptemplate),
				ApptemplateID:       gcore.F("apptemplate_id"),
				AttachmentTag:       gcore.F("root"),
				BootIndex:           gcore.F(int64(0)),
				DeleteOnTermination: gcore.F(true),
				ImageID:             gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
				Metadata:            gcore.F[any](map[string]interface{}{}),
				Name:                gcore.F("TestVM5 Ubuntu boot image"),
				Size:                gcore.F(int64(10)),
				SnapshotID:          gcore.F("snapshot_id"),
				TypeName:            gcore.F(gcore.CloudV1InstanceAvailableFlavorGetAvailableFlavorsParamsVolumesTypeNameCold),
				VolumeID:            gcore.F("volume_id"),
			}}),
			IncludePrices: gcore.F(true),
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

func TestCloudV1InstanceAvailableFlavorGetFlavorsToResizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Instances.AvailableFlavors.GetFlavorsToResize(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
		gcore.CloudV1InstanceAvailableFlavorGetFlavorsToResizeParams{
			IncludePrices: gcore.F(true),
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
