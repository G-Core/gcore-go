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

func TestBaremetalFlavorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Flavors.List(context.TODO(), cloud.BaremetalFlavorListParams{
		ProjectID:               gcore.Int(0),
		RegionID:                gcore.Int(0),
		Disabled:                gcore.Bool(true),
		ExcludeLinux:            gcore.Bool(true),
		ExcludeWindows:          gcore.Bool(true),
		IncludeCapacity:         gcore.Bool(true),
		IncludePrices:           gcore.Bool(true),
		IncludeReservationStock: gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBaremetalFlavorListSuitableWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Flavors.ListSuitable(context.TODO(), cloud.BaremetalFlavorListSuitableParams{
		ProjectID:     gcore.Int(0),
		RegionID:      gcore.Int(0),
		IncludePrices: gcore.Bool(true),
		ApptemplateID: gcore.String("apptemplate_id"),
		ImageID:       gcore.String("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
