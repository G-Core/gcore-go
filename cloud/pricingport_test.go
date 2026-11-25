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

func TestPricingPortPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Pricing.Ports.Preview(context.TODO(), cloud.PricingPortPreviewParams{
		ProjectID: gcore.Int(1),
		RegionID:  1,
		IPAssignments: []cloud.IPAssignmentRequestParam{{
			SubnetID:  "b39792c3-3160-4356-912e-ba396c95cdcf",
			IPAddress: gcore.String("192.168.123.20"),
		}},
		NetworkID: gcore.String("bc688791-f1b0-44eb-97d4-07697294b1e1"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
