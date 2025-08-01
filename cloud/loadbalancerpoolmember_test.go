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

func TestLoadBalancerPoolMemberAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.Members.Add(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.LoadBalancerPoolMemberAddParams{
			ProjectID:      gcore.Int(1),
			RegionID:       gcore.Int(1),
			Address:        "192.168.40.33",
			ProtocolPort:   80,
			AdminStateUp:   gcore.Bool(true),
			Backup:         gcore.Bool(true),
			InstanceID:     gcore.String("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
			MonitorAddress: gcore.String("monitor_address"),
			MonitorPort:    gcore.Int(0),
			SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
			Weight:         gcore.Int(1),
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

func TestLoadBalancerPoolMemberRemove(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.Members.Remove(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.LoadBalancerPoolMemberRemoveParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			PoolID:    "00000000-0000-4000-8000-000000000000",
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
