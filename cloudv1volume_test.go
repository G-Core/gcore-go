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

func TestCloudV1VolumeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1VolumeNewParams{
			Body: gcore.CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchema{
				ImageID:              gcore.F("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
				Name:                 gcore.F("vol3"),
				Size:                 gcore.F(int64(30)),
				AttachmentTag:        gcore.F("attachment_tag"),
				InstanceIDToAttachTo: gcore.F("instance_id_to_attach_to"),
				LifecyclePolicyIDs:   gcore.F([]int64{int64(0)}),
				Metadata:             gcore.F[any](map[string]interface{}{}),
				Source:               gcore.F("image"),
				TypeName:             gcore.F(gcore.CloudV1VolumeNewParamsBodyCreateVolumeFromImageSchemaTypeNameCold),
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

func TestCloudV1VolumeGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Get(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1VolumeUpdate(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Update(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV1VolumeUpdateParams{
			Name: gcore.NameParam{
				Name: gcore.F("my-resource"),
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

func TestCloudV1VolumeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1VolumeListParams{
			Bootable:       gcore.F(true),
			ClusterID:      gcore.F("cluster_id"),
			HasAttachments: gcore.F(true),
			IDPart:         gcore.F("id_part"),
			InstanceID:     gcore.F("instance_id"),
			Limit:          gcore.F(int64(0)),
			MetadataK:      gcore.F("metadata_k"),
			MetadataKv:     gcore.F("metadata_kv"),
			NamePart:       gcore.F("name_part"),
			Offset:         gcore.F(int64(0)),
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

func TestCloudV1VolumeDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Delete(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV1VolumeDeleteParams{
			Snapshots: gcore.F("726ecfcc-7fd0-4e30-a86e-7892524aa483,726ecfcc-7fd0-4e30-a86e-7892524aa484"),
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

func TestCloudV1VolumeAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Attach(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV1VolumeAttachParams{
			AttachVolumeToInstance: gcore.AttachVolumeToInstanceParam{
				InstanceID:    gcore.F("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
				AttachmentTag: gcore.F("root"),
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

func TestCloudV1VolumeChangeType(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.ChangeType(
		context.TODO(),
		int64(0),
		int64(0),
		"volume_id",
		gcore.CloudV1VolumeChangeTypeParams{
			VolumeType: gcore.F(gcore.CloudV1VolumeChangeTypeParamsVolumeTypeCold),
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

func TestCloudV1VolumeDetach(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Detach(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV1VolumeDetachParams{
			InstanceID: gcore.InstanceIDParam{
				InstanceID: gcore.F("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
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

func TestCloudV1VolumeExtend(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Extend(
		context.TODO(),
		int64(0),
		int64(0),
		"volume_id",
		gcore.CloudV1VolumeExtendParams{
			Size: gcore.F(int64(16)),
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

func TestCloudV1VolumeRevert(t *testing.T) {
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
	_, err := client.Cloud.V1.Volumes.Revert(
		context.TODO(),
		int64(0),
		int64(0),
		"volume_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
