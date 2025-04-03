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

func TestCloudV1PricingAIClusterGetAIClusterPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.AI.Clusters.GetAIClusterPrice(
		context.TODO(),
		int64(1),
		int64(1),
		"cluster_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1PricingAIClusterPreviewAIClusterPriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.AI.Clusters.PreviewAIClusterPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1PricingAIClusterPreviewAIClusterPriceParams{
			Flavor: gcore.F("flavor"),
			Interfaces: gcore.F([]gcore.NewVmInterfaceSerializersPydanticUnionParam{gcore.NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam{
				InterfaceName: gcore.F("eth0"),
				IPFamily:      gcore.F(gcore.InterfaceIPFamilyDual),
				PortGroup:     gcore.F(int64(1)),
				Type:          gcore.F(gcore.NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticTypeExternal),
			}}),
			Name:           gcore.F("name"),
			InstancesCount: gcore.F(int64(0)),
			Volumes: gcore.F([]gcore.InstanceVolumePricingRequestParam{{
				Source:     gcore.F(gcore.InstanceVolumePricingRequestSourceApptemplate),
				Size:       gcore.F(int64(1000)),
				SnapshotID: gcore.F("7cca40d7-a843-4e9f-ae08-62b9a394b1ab"),
				TypeName:   gcore.F(gcore.VolumeTypeNameCold),
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
