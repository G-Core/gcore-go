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

func TestVolumeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Volumes.New(context.TODO(), cloud.VolumeNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		OfImage: &cloud.VolumeNewParamsBodyImage{
			ImageID:              "169942e0-9b53-42df-95ef-1a8b6525c2bd",
			Name:                 "volume-1",
			Size:                 10,
			AttachmentTag:        gcore.String("device-tag"),
			InstanceIDToAttachTo: gcore.String("88f3e0bd-ca86-4cf7-be8b-dd2988e23c2d"),
			LifecyclePolicyIDs:   []int64{1, 2},
			Tags: cloud.TagUpdateMap{
				"foo": "my-tag-value",
			},
			TypeName: "standard",
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

func TestVolumeUpdate(t *testing.T) {
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
	_, err := client.Cloud.Volumes.Update(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      "my-resource",
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

func TestVolumeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Volumes.List(context.TODO(), cloud.VolumeListParams{
		ProjectID:      gcore.Int(1),
		RegionID:       gcore.Int(1),
		Bootable:       gcore.Bool(false),
		ClusterID:      gcore.String("t12345"),
		HasAttachments: gcore.Bool(true),
		IDPart:         gcore.String("726ecfcc-7fd0-4e30-a86e-7892524aa483"),
		InstanceID:     gcore.String("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
		Limit:          gcore.Int(1000),
		NamePart:       gcore.String("test"),
		Offset:         gcore.Int(0),
		TagKey:         []string{"key1", "key2"},
		TagKeyValue:    gcore.String("tag_key_value"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVolumeDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Volumes.Delete(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Snapshots: gcore.String("726ecfcc-7fd0-4e30-a86e-7892524aa483,726ecfcc-7fd0-4e30-a86e-7892524aa484"),
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

func TestVolumeAttachToInstanceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Volumes.AttachToInstance(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeAttachToInstanceParams{
			ProjectID:     gcore.Int(1),
			RegionID:      gcore.Int(1),
			InstanceID:    "169942e0-9b53-42df-95ef-1a8b6525c2bd",
			AttachmentTag: gcore.String("boot"),
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

func TestVolumeChangeType(t *testing.T) {
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
	_, err := client.Cloud.Volumes.ChangeType(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeChangeTypeParams{
			ProjectID:  gcore.Int(1),
			RegionID:   gcore.Int(1),
			VolumeType: cloud.VolumeChangeTypeParamsVolumeTypeSsdHiiops,
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

func TestVolumeDetachFromInstance(t *testing.T) {
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
	_, err := client.Cloud.Volumes.DetachFromInstance(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeDetachFromInstanceParams{
			ProjectID:  gcore.Int(1),
			RegionID:   gcore.Int(1),
			InstanceID: "169942e0-9b53-42df-95ef-1a8b6525c2bd",
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

func TestVolumeGet(t *testing.T) {
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
	_, err := client.Cloud.Volumes.Get(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeGetParams{
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

func TestVolumeResize(t *testing.T) {
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
	_, err := client.Cloud.Volumes.Resize(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeResizeParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Size:      100,
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

func TestVolumeRevertToLastSnapshot(t *testing.T) {
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
	err := client.Cloud.Volumes.RevertToLastSnapshot(
		context.TODO(),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		cloud.VolumeRevertToLastSnapshotParams{
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
