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

func TestCloudV1RouterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1RouterNewParams{
			Name: gcore.F("my_wonderful_router"),
			ExternalGatewayInfo: gcore.F[gcore.CloudV1RouterNewParamsExternalGatewayInfoUnion](gcore.CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer{
				EnableSnat: gcore.F(true),
				Type:       gcore.F(gcore.CloudV1RouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializerTypeDefault),
			}),
			Interfaces: gcore.F([]gcore.CloudV1RouterNewParamsInterface{{
				SubnetID: gcore.F("3ed9e2ce-f906-47fb-ba32-c25a3f63df4f"),
				Type:     gcore.F(gcore.CloudV1RouterNewParamsInterfacesTypeSubnet),
			}}),
			Routes: gcore.F([]gcore.NeutronRouteParam{{
				Destination: gcore.F("10.0.3.0/24"),
				Nexthop:     gcore.F("10.0.0.13"),
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

func TestCloudV1RouterGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"router_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1RouterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"router_id",
		gcore.CloudV1RouterUpdateParams{
			ExternalGatewayInfo: gcore.F(gcore.RouterExternalManualGwParam{
				NetworkID:  gcore.F("3d1e9779-b8ad-4c8e-a1a4-0a03671f07f7"),
				EnableSnat: gcore.F(true),
				Type:       gcore.F(gcore.RouterExternalManualGwTypeManual),
			}),
			Name: gcore.F("my_renamed_router"),
			Routes: gcore.F([]gcore.NeutronRouteParam{{
				Destination: gcore.F("10.0.3.0/24"),
				Nexthop:     gcore.F("10.0.0.13"),
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

func TestCloudV1RouterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1RouterListParams{
			Limit:  gcore.F(int64(0)),
			Offset: gcore.F(int64(0)),
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

func TestCloudV1RouterDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"router_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1RouterAttachSubnet(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.AttachSubnet(
		context.TODO(),
		int64(0),
		int64(0),
		"router_id",
		gcore.CloudV1RouterAttachSubnetParams{
			SubnetID: gcore.SubnetIDParam{
				SubnetID: gcore.F("subnet_id"),
			},
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

func TestCloudV1RouterDetachSubnet(t *testing.T) {
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
	_, err := client.Cloud.V1.Routers.DetachSubnet(
		context.TODO(),
		int64(0),
		int64(0),
		"router_id",
		gcore.CloudV1RouterDetachSubnetParams{
			SubnetID: gcore.SubnetIDParam{
				SubnetID: gcore.F("subnet_id"),
			},
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
