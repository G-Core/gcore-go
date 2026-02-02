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
	"github.com/G-Core/gcore-go/packages/param"
)

func TestLoadBalancerListenerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Listeners.New(context.TODO(), cloud.LoadBalancerListenerNewParams{
		ProjectID:            gcore.Int(1),
		RegionID:             gcore.Int(1),
		LoadBalancerID:       "30f4f55b-4a7c-48e0-9954-5cddfee216e7",
		Name:                 "my_listener",
		Protocol:             cloud.LbListenerProtocolHTTP,
		ProtocolPort:         80,
		AllowedCidrs:         []string{"10.0.0.0/8"},
		ConnectionLimit:      gcore.Int(100000),
		DefaultPoolID:        gcore.String("00000000-0000-4000-8000-000000000000"),
		InsertXForwarded:     gcore.Bool(false),
		SecretID:             gcore.String("f2e734d0-fa2b-42c2-ad33-4c6db5101e00"),
		SniSecretID:          []string{"f2e734d0-fa2b-42c2-ad33-4c6db5101e00", "eb121225-7ded-4ff3-ae1f-599e145dd7cb"},
		TimeoutClientData:    gcore.Int(50000),
		TimeoutMemberConnect: gcore.Int(50000),
		TimeoutMemberData:    param.Null[int64](),
		UserList: []cloud.LoadBalancerListenerNewParamsUserList{{
			EncryptedPassword: "$5$isRr.HJ1IrQP38.m$oViu3DJOpUG2ZsjCBtbITV3mqpxxbZfyWJojLPNSPO5",
			Username:          "admin",
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

func TestLoadBalancerListenerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Listeners.Update(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.LoadBalancerListenerUpdateParams{
			ProjectID:            gcore.Int(1),
			RegionID:             gcore.Int(1),
			AdminStateUp:         gcore.Bool(true),
			AllowedCidrs:         []string{"10.0.0.0/8"},
			ConnectionLimit:      gcore.Int(100000),
			Name:                 gcore.String("new_listener_name"),
			SecretID:             gcore.String("af4a64e7-03ca-470f-9a09-b77d54c5abd8"),
			SniSecretID:          []string{"af4a64e7-03ca-470f-9a09-b77d54c5abd8", "12b43d95-d420-4c79-a883-49bf146cbdff"},
			TimeoutClientData:    gcore.Int(50000),
			TimeoutMemberConnect: gcore.Int(50000),
			TimeoutMemberData:    param.Null[int64](),
			UserList: []cloud.LoadBalancerListenerUpdateParamsUserList{{
				EncryptedPassword: "$5$isRr.HJ1IrQP38.m$oViu3DJOpUG2ZsjCBtbITV3mqpxxbZfyWJojLPNSPO5",
				Username:          "admin",
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

func TestLoadBalancerListenerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Listeners.List(context.TODO(), cloud.LoadBalancerListenerListParams{
		ProjectID:      gcore.Int(1),
		RegionID:       gcore.Int(1),
		LoadBalancerID: gcore.String("00000000-0000-4000-8000-000000000000"),
		ShowStats:      gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerListenerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Listeners.Delete(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.LoadBalancerListenerDeleteParams{
			ProjectID:         gcore.Int(1),
			RegionID:          gcore.Int(1),
			DeleteDefaultPool: gcore.Bool(false),
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

func TestLoadBalancerListenerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Listeners.Get(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.LoadBalancerListenerGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			ShowStats: gcore.Bool(true),
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
