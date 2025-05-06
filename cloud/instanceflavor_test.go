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

func TestInstanceFlavorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Flavors.List(context.TODO(), cloud.InstanceFlavorListParams{
		ProjectID:      gcore.Int(0),
		RegionID:       gcore.Int(0),
		Disabled:       gcore.Bool(true),
		ExcludeLinux:   gcore.Bool(true),
		ExcludeWindows: gcore.Bool(true),
		IncludePrices:  gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceFlavorListForResizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Flavors.ListForResize(
		context.TODO(),
		"instance_id",
		cloud.InstanceFlavorListForResizeParams{
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

func TestInstanceFlavorListSuitableWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Flavors.ListSuitable(context.TODO(), cloud.InstanceFlavorListSuitableParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Volumes: []cloud.InstanceFlavorListSuitableParamsVolume{{
			Source:        "image",
			ApptemplateID: gcore.String("apptemplate_id"),
			BootIndex:     gcore.Int(0),
			ImageID:       gcore.String("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			Name:          gcore.String("TestVM5 Ubuntu boot image"),
			Size:          gcore.Int(10),
			SnapshotID:    gcore.String("snapshot_id"),
			TypeName:      "ssd_hiiops",
			VolumeID:      gcore.String("volume_id"),
		}},
		IncludePrices: gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
