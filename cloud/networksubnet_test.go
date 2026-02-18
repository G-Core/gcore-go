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

func TestNetworkSubnetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Subnets.New(context.TODO(), cloud.NetworkSubnetNewParams{
		ProjectID:              gcore.Int(1),
		RegionID:               gcore.Int(1),
		Cidr:                   "192.168.10.0/24",
		Name:                   "my subnet",
		NetworkID:              "ee2402d0-f0cd-4503-9b75-69be1d11c5f1",
		ConnectToNetworkRouter: gcore.Bool(true),
		DNSNameservers:         []string{"8.8.4.4", "1.1.1.1"},
		EnableDhcp:             gcore.Bool(true),
		GatewayIP:              gcore.String("192.168.10.1"),
		HostRoutes: []cloud.NetworkSubnetNewParamsHostRoute{{
			Destination: "10.0.3.0/24",
			Nexthop:     "10.0.0.13",
		}},
		IPVersion:         4,
		RouterIDToConnect: gcore.String("00000000-0000-4000-8000-000000000000"),
		Tags: map[string]string{
			"my-tag": "my-tag-value",
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

func TestNetworkSubnetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Subnets.Update(
		context.TODO(),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
		cloud.NetworkSubnetUpdateParams{
			ProjectID:      gcore.Int(1),
			RegionID:       gcore.Int(1),
			DNSNameservers: []string{"8.8.4.4", "1.1.1.1"},
			EnableDhcp:     gcore.Bool(true),
			GatewayIP:      gcore.String("192.168.10.1"),
			HostRoutes: []cloud.NetworkSubnetUpdateParamsHostRoute{{
				Destination: "10.0.3.0/24",
				Nexthop:     "10.0.0.13",
			}},
			Name: gcore.String("some_name"),
			Tags: cloud.TagUpdateMap{
				"foo": "my-tag-value",
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

func TestNetworkSubnetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Networks.Subnets.List(context.TODO(), cloud.NetworkSubnetListParams{
		ProjectID:   gcore.Int(1),
		RegionID:    gcore.Int(1),
		Limit:       gcore.Int(1000),
		NetworkID:   gcore.String("b30d0de7-bca2-4c83-9c57-9e645bd2cc92"),
		Offset:      gcore.Int(0),
		OrderBy:     cloud.NetworkSubnetListParamsOrderByNameAsc,
		TagKey:      []string{"key1", "key2"},
		TagKeyValue: gcore.String("tag_key_value"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkSubnetDelete(t *testing.T) {
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
	_, err := client.Cloud.Networks.Subnets.Delete(
		context.TODO(),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
		cloud.NetworkSubnetDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestNetworkSubnetGet(t *testing.T) {
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
	_, err := client.Cloud.Networks.Subnets.Get(
		context.TODO(),
		"b39792c3-3160-4356-912e-ba396c95cdcf",
		cloud.NetworkSubnetGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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
