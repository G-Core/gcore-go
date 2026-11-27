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

func TestPricingInstanceGet(t *testing.T) {
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
	_, err := client.Cloud.Pricing.Instances.Get(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.PricingInstanceGetParams{
			ProjectID: gcore.Int(1),
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

func TestPricingInstancePreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Pricing.Instances.Preview(context.TODO(), cloud.PricingInstancePreviewParams{
		ProjectID: gcore.Int(1),
		OfInstancePricePreviewReferenceBody: &cloud.PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody{
			FlavorName:    "g2-standard-4-8",
			RegionID:      7,
			BootVolumeID:  gcore.String("boot_volume_id"),
			DataVolumeIDs: []string{"string"},
			Interfaces: []cloud.PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion{{
				OfPricePortReference: &cloud.PricePortReferenceParam{
					PortID: "port_id",
					FloatingIP: cloud.PriceFloatingIPParam{
						ID: gcore.String("id"),
					},
				},
			}},
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
