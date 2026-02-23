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

func TestVolumeSnapshotNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.VolumeSnapshots.New(context.TODO(), cloud.VolumeSnapshotNewParams{
		ProjectID:   gcore.Int(0),
		RegionID:    gcore.Int(0),
		Name:        "my-snapshot",
		VolumeID:    "67baa7d1-08ea-4fc5-bef2-6b2465b7d227",
		Description: gcore.String("Snapshot description"),
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

func TestVolumeSnapshotUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.VolumeSnapshots.Update(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeSnapshotUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("my-backup-snapshot"),
			Tags: cloud.TagUpdateMap{
				"foo": "string",
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

func TestVolumeSnapshotDelete(t *testing.T) {
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
	_, err := client.Cloud.VolumeSnapshots.Delete(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeSnapshotDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestVolumeSnapshotGet(t *testing.T) {
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
	_, err := client.Cloud.VolumeSnapshots.Get(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeSnapshotGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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
