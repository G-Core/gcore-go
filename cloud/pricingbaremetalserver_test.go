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

func TestPricingBaremetalServerGet(t *testing.T) {
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
	_, err := client.Cloud.Pricing.Baremetal.Servers.Get(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.PricingBaremetalServerGetParams{
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

func TestPricingBaremetalServerPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Pricing.Baremetal.Servers.Preview(context.TODO(), cloud.PricingBaremetalServerPreviewParams{
		ProjectID:  gcore.Int(1),
		FlavorName: "g2-standard-4-8",
		RegionID:   7,
		Interfaces: []cloud.PricingBaremetalServerPreviewParamsInterfaceUnion{{
			OfPricePortReference: &cloud.PricePortReferenceParam{
				PortID: "port_id",
				FloatingIP: cloud.PriceFloatingIPParam{
					FloatingIPID: gcore.String("floating_ip_id"),
				},
			},
		}},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
