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

func TestCloudV1SubnetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Subnets.New(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1SubnetNewParams{
			Cidr:                   gcore.F("192.168.10.0/24"),
			Name:                   gcore.F("my subnet"),
			NetworkID:              gcore.F("ee2402d0-f0cd-4503-9b75-69be1d11c5f1"),
			ConnectToNetworkRouter: gcore.F(true),
			DNSNameservers:         gcore.F([]string{"8.8.4.4", "1.1.1.1"}),
			EnableDhcp:             gcore.F(true),
			GatewayIP:              gcore.F("192.168.10.1"),
			HostRoutes: gcore.F([]gcore.NeutronRouteParam{{
				Destination: gcore.F("10.0.3.0/24"),
				Nexthop:     gcore.F("10.0.0.13"),
			}}),
			IPVersion: gcore.F(gcore.IPVersion4),
			Metadata: gcore.F[any](map[string]interface{}{
				"key": "value",
			}),
			RouterIDToConnect: gcore.F("00000000-0000-4000-8000-000000000000"),
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

func TestCloudV1SubnetGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Subnets.Get(
		context.TODO(),
		int64(1),
		int64(1),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1SubnetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Subnets.Update(
		context.TODO(),
		int64(1),
		int64(1),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
		gcore.CloudV1SubnetUpdateParams{
			DNSNameservers: gcore.F([]string{"8.8.4.4", "1.1.1.1"}),
			EnableDhcp:     gcore.F(true),
			GatewayIP:      gcore.F("192.168.10.1"),
			HostRoutes: gcore.F([]gcore.NeutronRouteParam{{
				Destination: gcore.F("10.0.3.0/24"),
				Nexthop:     gcore.F("10.0.0.13"),
			}}),
			Name: gcore.F("some_name"),
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

func TestCloudV1SubnetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Subnets.List(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1SubnetListParams{
			Limit:      gcore.F(int64(1000)),
			MetadataK:  gcore.F([]string{"key1", "key2"}),
			MetadataKv: gcore.F("metadata_kv"),
			NetworkID:  gcore.F("b30d0de7-bca2-4c83-9c57-9e645bd2cc92"),
			Offset:     gcore.F(int64(0)),
			OrderBy:    gcore.F([]gcore.CloudV1SubnetListParamsOrderBy{gcore.CloudV1SubnetListParamsOrderByAvailableIPsAsc}),
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

func TestCloudV1SubnetDelete(t *testing.T) {
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
	err := client.Cloud.V1.Subnets.Delete(
		context.TODO(),
		int64(1),
		int64(1),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
