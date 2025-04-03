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

func TestCloudV1AIClusterGPUNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.GPU.New(
		context.TODO(),
		"project_id",
		"region_id",
		gcore.CloudV1AIClusterGPUNewParams{
			Flavor:  gcore.F("bm3-ai-18xlarge-a100-40-4"),
			ImageID: gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			Interfaces: gcore.F([]gcore.CloudV1AIClusterGPUNewParamsInterfaceUnion{gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("subnet"),
			}}),
			Name:           gcore.F("gpu-ubuntu"),
			InstancesCount: gcore.F(int64(1)),
			KeypairName:    gcore.F("keypair_name"),
			Metadata:       gcore.F[any](map[string]interface{}{}),
			Password:       gcore.F("password"),
			SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
				ID: gcore.F("4536dba1-93b1-492e-b3df-270b6b9f3650"),
			}, {
				ID: gcore.F("cee2ca1f-507a-4a31-b714-f6c1ffb4bdfa"),
			}}),
			UserData: gcore.F("user_data"),
			Username: gcore.F("username"),
			Volumes: gcore.F([]gcore.CreateInstanceVolumeParam{{
				Source:              gcore.F(gcore.CreateInstanceVolumeSourceApptemplate),
				ApptemplateID:       gcore.F("apptemplate_id"),
				AttachmentTag:       gcore.F("root"),
				BootIndex:           gcore.F(int64(0)),
				DeleteOnTermination: gcore.F(false),
				ImageID:             gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
				Metadata: gcore.F[any](map[string]interface{}{
					"key": "value",
				}),
				Name:       gcore.F("TestVM5 Ubuntu boot image"),
				Size:       gcore.F(int64(10)),
				SnapshotID: gcore.F("snapshot_id"),
				TypeName:   gcore.F(gcore.CreateInstanceVolumeTypeNameCold),
				VolumeID:   gcore.F("volume_id"),
			}}),
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

func TestCloudV1AIClusterGPURebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.GPU.Rebuild(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_id",
		gcore.CloudV1AIClusterGPURebuildParams{
			Nodes:    gcore.F([]string{"string"}),
			ImageID:  gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			UserData: gcore.F("user_data"),
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

func TestCloudV1AIClusterGPURemoveNodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.GPU.RemoveNode(
		context.TODO(),
		"project_id",
		"region_id",
		"cluster_id",
		"instance_id",
		gcore.CloudV1AIClusterGPURemoveNodeParams{
			DeleteFloatings: gcore.F(true),
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

func TestCloudV1AIClusterGPUResize(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.GPU.Resize(
		context.TODO(),
		"project_id",
		"region_id",
		"cluster_id",
		gcore.CloudV1AIClusterGPUResizeParams{
			InstancesCount: gcore.F(int64(1)),
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
