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

func TestCloudV1PricingInferencePreviewDeploymentPrice(t *testing.T) {
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
	_, err := client.Cloud.V1.Pricing.Inference.PreviewDeploymentPrice(context.TODO(), gcore.CloudV1PricingInferencePreviewDeploymentPriceParams{
		Containers: gcore.F([]gcore.CloudV1PricingInferencePreviewDeploymentPriceParamsContainer{{
			RegionID: gcore.F(int64(7)),
			Scale: gcore.F(gcore.CloudV1PricingInferencePreviewDeploymentPriceParamsContainersScale{
				Max: gcore.F(int64(3)),
			}),
		}}),
		FlavorID: gcore.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
