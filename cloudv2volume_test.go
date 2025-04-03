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

func TestCloudV2VolumeAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Volumes.Attach(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV2VolumeAttachParams{
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

func TestCloudV2VolumeDetach(t *testing.T) {
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
	_, err := client.Cloud.V2.Volumes.Detach(
		context.TODO(),
		int64(1),
		int64(1),
		"726ecfcc-7fd0-4e30-a86e-7892524aa483",
		gcore.CloudV2VolumeDetachParams{
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
