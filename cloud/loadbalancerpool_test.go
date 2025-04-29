// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestLoadBalancerPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.New(context.TODO(), cloud.LoadBalancerPoolNewParams{
		ProjectID:   gcore.Int(0),
		RegionID:    gcore.Int(0),
		LbAlgorithm: cloud.LbAlgorithmLeastConnections,
		Name:        "pool_name",
		Protocol:    cloud.LbPoolProtocolHTTP,
		CaSecretID:  gcore.String("ca_secret_id"),
		CrlSecretID: gcore.String("crl_secret_id"),
		Healthmonitor: cloud.LoadBalancerPoolNewParamsHealthmonitor{
			Delay:          10,
			MaxRetries:     3,
			Timeout:        5,
			Type:           cloud.HealthMonitorTypeHTTP,
			ExpectedCodes:  gcore.String("200,301,302"),
			HTTPMethod:     cloud.HTTPMethodGet,
			MaxRetriesDown: gcore.Int(3),
			URLPath:        gcore.String("/"),
		},
		ListenerID:     gcore.String("listener_id"),
		LoadbalancerID: gcore.String("bbb35f84-35cc-4b2f-84c2-a6a29bba68aa"),
		Members: []cloud.LoadBalancerPoolNewParamsMember{{
			Address:        "192.168.1.101",
			ProtocolPort:   8000,
			AdminStateUp:   gcore.Bool(false),
			InstanceID:     gcore.String("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
			MonitorAddress: gcore.String("monitor_address"),
			MonitorPort:    gcore.Int(0),
			SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
			Weight:         gcore.Int(2),
		}, {
			Address:        "192.168.1.102",
			ProtocolPort:   8000,
			AdminStateUp:   gcore.Bool(false),
			InstanceID:     gcore.String("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
			MonitorAddress: gcore.String("monitor_address"),
			MonitorPort:    gcore.Int(0),
			SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
			Weight:         gcore.Int(4),
		}},
		SecretID: gcore.String("secret_id"),
		SessionPersistence: cloud.LoadBalancerPoolNewParamsSessionPersistence{
			Type:                   cloud.SessionPersistenceTypeAppCookie,
			CookieName:             gcore.String("cookie_name"),
			PersistenceGranularity: gcore.String("persistence_granularity"),
			PersistenceTimeout:     gcore.Int(0),
		},
		TimeoutClientData:    gcore.Int(50000),
		TimeoutMemberConnect: gcore.Int(50000),
		TimeoutMemberData:    gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.Update(
		context.TODO(),
		"pool_id",
		cloud.LoadBalancerPoolUpdateParams{
			ProjectID:   gcore.Int(0),
			RegionID:    gcore.Int(0),
			CaSecretID:  gcore.String("ca_secret_id"),
			CrlSecretID: gcore.String("crl_secret_id"),
			Healthmonitor: cloud.LoadBalancerPoolUpdateParamsHealthmonitor{
				Delay:          10,
				MaxRetries:     2,
				Timeout:        5,
				ExpectedCodes:  gcore.String("200,301,302"),
				HTTPMethod:     cloud.HTTPMethodConnect,
				MaxRetriesDown: gcore.Int(2),
				Type:           cloud.HealthMonitorTypeHTTP,
				URLPath:        gcore.String("/"),
			},
			LbAlgorithm: cloud.LbAlgorithmLeastConnections,
			Members: []cloud.LoadBalancerPoolUpdateParamsMember{{
				Address:        "192.168.40.33",
				ProtocolPort:   80,
				AdminStateUp:   gcore.Bool(false),
				InstanceID:     gcore.String("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
				MonitorAddress: gcore.String("monitor_address"),
				MonitorPort:    gcore.Int(0),
				SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
				Weight:         gcore.Int(1),
			}},
			Name:     gcore.String("new_pool_name"),
			Protocol: cloud.LbPoolProtocolHTTP,
			SecretID: gcore.String("secret_id"),
			SessionPersistence: cloud.LoadBalancerPoolUpdateParamsSessionPersistence{
				Type:                   cloud.SessionPersistenceTypeAppCookie,
				CookieName:             gcore.String("cookie_name"),
				PersistenceGranularity: gcore.String("persistence_granularity"),
				PersistenceTimeout:     gcore.Int(0),
			},
			TimeoutClientData:    gcore.Int(50000),
			TimeoutMemberConnect: gcore.Int(50000),
			TimeoutMemberData:    gcore.Int(0),
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

func TestLoadBalancerPoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.List(context.TODO(), cloud.LoadBalancerPoolListParams{
		ProjectID:      gcore.Int(0),
		RegionID:       gcore.Int(0),
		Details:        gcore.Bool(true),
		ListenerID:     gcore.String("listener_id"),
		LoadbalancerID: gcore.String("loadbalancer_id"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerPoolDelete(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.Delete(
		context.TODO(),
		"pool_id",
		cloud.LoadBalancerPoolDeleteParams{
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

func TestLoadBalancerPoolGet(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.Pools.Get(
		context.TODO(),
		"pool_id",
		cloud.LoadBalancerPoolGetParams{
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
