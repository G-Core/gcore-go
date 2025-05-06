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

func TestNetworkRouterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.New(context.TODO(), cloud.NetworkRouterNewParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Name:      "my_wonderful_router",
		ExternalGatewayInfo: cloud.NetworkRouterNewParamsExternalGatewayInfoUnion{
			OfRouterExternalDefaultGwSerializer: &cloud.NetworkRouterNewParamsExternalGatewayInfoRouterExternalDefaultGwSerializer{
				EnableSnat: gcore.Bool(true),
				Type:       "default",
			},
		},
		Interfaces: []cloud.NetworkRouterNewParamsInterface{{
			SubnetID: "3ed9e2ce-f906-47fb-ba32-c25a3f63df4f",
			Type:     "subnet",
		}},
		Routes: []cloud.NetworkRouterNewParamsRoute{{
			Destination: "10.0.3.0/24",
			Nexthop:     "10.0.0.13",
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

func TestNetworkRouterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.Update(
		context.TODO(),
		"router_id",
		cloud.NetworkRouterUpdateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			ExternalGatewayInfo: cloud.NetworkRouterUpdateParamsExternalGatewayInfo{
				NetworkID:  "d7745dcf-b302-4795-9d61-6cc52487af48",
				EnableSnat: gcore.Bool(false),
				Type:       "manual",
			},
			Name: gcore.String("my_renamed_router"),
			Routes: []cloud.NetworkRouterUpdateParamsRoute{{
				Destination: "10.0.3.0/24",
				Nexthop:     "10.0.0.13",
			}},
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

func TestNetworkRouterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.List(context.TODO(), cloud.NetworkRouterListParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Limit:     gcore.Int(0),
		Offset:    gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterDelete(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.Delete(
		context.TODO(),
		"router_id",
		cloud.NetworkRouterDeleteParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestNetworkRouterAttachSubnetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.AttachSubnet(
		context.TODO(),
		"ccd5b925-e35c-4611-a511-67ab503104c8",
		cloud.NetworkRouterAttachSubnetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			SubnetID:  "subnet_id",
			IPAddress: gcore.String("ip_address"),
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

func TestNetworkRouterDetachSubnet(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.DetachSubnet(
		context.TODO(),
		"router_id",
		cloud.NetworkRouterDetachSubnetParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			SubnetID: cloud.SubnetIDParam{
				SubnetID: "subnet_id",
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

func TestNetworkRouterGet(t *testing.T) {
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
	_, err := client.Cloud.Networks.Routers.Get(
		context.TODO(),
		"router_id",
		cloud.NetworkRouterGetParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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
