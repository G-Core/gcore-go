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

func TestCloudV2PricingInferencePreviewDeploymentPrice(t *testing.T) {
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
	_, err := client.Cloud.V2.Pricing.Inference.PreviewDeploymentPrice(context.TODO(), gcore.CloudV2PricingInferencePreviewDeploymentPriceParams{
		Containers: gcore.F([]gcore.CloudV2PricingInferencePreviewDeploymentPriceParamsContainer{{
			RegionID: gcore.F(int64(7)),
			Scale: gcore.F(gcore.CloudV2PricingInferencePreviewDeploymentPriceParamsContainersScale{
				Max: gcore.F(int64(3)),
			}),
		}}),
		FlavorName: gcore.F("flavor_name"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
