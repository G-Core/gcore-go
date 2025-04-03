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

func TestCloudV2PricingGetK8sClusterPrice(t *testing.T) {
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
	_, err := client.Cloud.V2.Pricing.GetK8sClusterPrice(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2PricingPreviewInstancePriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Pricing.PreviewInstancePrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV2PricingPreviewInstancePriceParams{
			Flavor: gcore.F("g1-standard-1-2"),
			Interfaces: gcore.F([]gcore.NewVmInterfaceSerializersPydanticUnionParam{gcore.NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam{
				InterfaceName: gcore.F("eth0"),
				IPFamily:      gcore.F(gcore.InterfaceIPFamilyDual),
				PortGroup:     gcore.F(int64(1)),
				Type:          gcore.F(gcore.NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticTypeExternal),
			}}),
			NameTemplates: gcore.F([]string{"string"}),
			Names:         gcore.F([]string{"string"}),
			Volumes: gcore.F([]gcore.InstanceVolumePricingRequestParam{{
				Source:     gcore.F(gcore.InstanceVolumePricingRequestSourceApptemplate),
				Size:       gcore.F(int64(10)),
				SnapshotID: gcore.F("7cca40d7-a843-4e9f-ae08-62b9a394b1ab"),
				TypeName:   gcore.F(gcore.VolumeTypeNameCold),
			}, {
				Source:     gcore.F(gcore.InstanceVolumePricingRequestSourceApptemplate),
				Size:       gcore.F(int64(5)),
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

func TestCloudV2PricingPreviewK8sClusterPrice(t *testing.T) {
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
	_, err := client.Cloud.V2.Pricing.PreviewK8sClusterPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV2PricingPreviewK8sClusterPriceParams{
			Pools: gcore.F([]gcore.CloudV2PricingPreviewK8sClusterPriceParamsPool{{
				FlavorID:       gcore.F("g0-standard-1-2"),
				BootVolumeSize: gcore.F(int64(10)),
				BootVolumeType: gcore.F(gcore.CloudV2PricingPreviewK8sClusterPriceParamsPoolsBootVolumeTypeCold),
				IsPublicIpv4:   gcore.F(true),
				NodeCount:      gcore.F(int64(1)),
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

func TestCloudV2PricingPreviewLoadBalancerPriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Pricing.PreviewLoadBalancerPrice(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV2PricingPreviewLoadBalancerPriceParams{
			Flavor: gcore.F("lb1-1-2"),
			FloatingIP: gcore.F(gcore.NewInterfaceFloatingIPPydanticParam{
				Source:             gcore.F(gcore.NewInterfaceFloatingIPPydanticSourceExisting),
				ExistingFloatingID: gcore.F("e3c6ee77-48cb-416b-b204-11b492cc776e3"),
			}),
			VipIPFamily:  gcore.F(gcore.InterfaceIPFamilyDual),
			VipNetworkID: gcore.F("ddc28e44-2acb-499b-985b-831f29432e1c"),
			VipPortID:    gcore.F("57be69f6-6f6a-4f03-a4ad-8eb86c69ec0a"),
			VipSubnetID:  gcore.F("2731a56e-a5c9-44be-80c2-02c57c594573"),
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
