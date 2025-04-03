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

func TestCloudV1LbpoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1LbpoolNewParams{
			CreateLbPool: gcore.CreateLbPoolParam{
				LbAlgorithm: gcore.F(gcore.LbAlgorithmLeastConnections),
				Name:        gcore.F("pool_name"),
				Protocol:    gcore.F(gcore.LbPoolProtocolHTTP),
				CaSecretID:  gcore.F("ca_secret_id"),
				CrlSecretID: gcore.F("crl_secret_id"),
				Healthmonitor: gcore.F(gcore.CreateLbHealthMonitorParam{
					Delay:          gcore.F(int64(5)),
					MaxRetries:     gcore.F(int64(1)),
					Timeout:        gcore.F(int64(30)),
					Type:           gcore.F(gcore.HealthMonitorTypeHTTP),
					ExpectedCodes:  gcore.F("200,301,302"),
					HTTPMethod:     gcore.F(gcore.HTTPMethodHealthmonitorConnect),
					MaxRetriesDown: gcore.F(int64(3)),
					URLPath:        gcore.F("/"),
				}),
				ListenerID:     gcore.F("listener_id"),
				LoadbalancerID: gcore.F("bbb35f84-35cc-4b2f-84c2-a6a29bba68aa"),
				Members: gcore.F([]gcore.CreateLbPoolMemberParam{{
					Address:        gcore.F("192.168.1.101"),
					ProtocolPort:   gcore.F(int64(8000)),
					AdminStateUp:   gcore.F(false),
					InstanceID:     gcore.F("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
					MonitorAddress: gcore.F("monitor_address"),
					MonitorPort:    gcore.F(int64(0)),
					SubnetID:       gcore.F("32283b0b-b560-4690-810c-f672cbb2e28d"),
					Weight:         gcore.F(int64(2)),
				}, {
					Address:        gcore.F("192.168.1.102"),
					ProtocolPort:   gcore.F(int64(8000)),
					AdminStateUp:   gcore.F(false),
					InstanceID:     gcore.F("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
					MonitorAddress: gcore.F("monitor_address"),
					MonitorPort:    gcore.F(int64(0)),
					SubnetID:       gcore.F("32283b0b-b560-4690-810c-f672cbb2e28d"),
					Weight:         gcore.F(int64(4)),
				}}),
				SecretID: gcore.F("secret_id"),
				SessionPersistence: gcore.F(gcore.MutateLbSessionPersistenceParam{
					Type:                   gcore.F(gcore.SessionPersistenceTypeAppCookie),
					CookieName:             gcore.F("cookie_name"),
					PersistenceGranularity: gcore.F("persistence_granularity"),
					PersistenceTimeout:     gcore.F(int64(0)),
				}),
				TimeoutClientData:    gcore.F(int64(50000)),
				TimeoutMemberConnect: gcore.F(int64(50000)),
				TimeoutMemberData:    gcore.F(int64(0)),
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

func TestCloudV1LbpoolGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1LbpoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
		gcore.CloudV1LbpoolUpdateParams{
			CaSecretID:  gcore.F("ca_secret_id"),
			CrlSecretID: gcore.F("crl_secret_id"),
			Healthmonitor: gcore.F(gcore.CloudV1LbpoolUpdateParamsHealthmonitor{
				Delay:          gcore.F(int64(5)),
				MaxRetries:     gcore.F(int64(1)),
				Timeout:        gcore.F(int64(30)),
				ExpectedCodes:  gcore.F("200,301,302"),
				HTTPMethod:     gcore.F(gcore.HTTPMethodHealthmonitorConnect),
				MaxRetriesDown: gcore.F(int64(3)),
				Type:           gcore.F(gcore.HealthMonitorTypeHTTP),
				URLPath:        gcore.F("/"),
			}),
			LbAlgorithm: gcore.F(gcore.LbAlgorithmLeastConnections),
			Members: gcore.F([]gcore.CreateLbPoolMemberParam{{
				Address:        gcore.F("192.168.40.33"),
				ProtocolPort:   gcore.F(int64(80)),
				AdminStateUp:   gcore.F(false),
				InstanceID:     gcore.F("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
				MonitorAddress: gcore.F("monitor_address"),
				MonitorPort:    gcore.F(int64(0)),
				SubnetID:       gcore.F("32283b0b-b560-4690-810c-f672cbb2e28d"),
				Weight:         gcore.F(int64(1)),
			}}),
			Name:     gcore.F("new_pool_name"),
			Protocol: gcore.F(gcore.LbPoolProtocolHTTP),
			SecretID: gcore.F("secret_id"),
			SessionPersistence: gcore.F(gcore.MutateLbSessionPersistenceParam{
				Type:                   gcore.F(gcore.SessionPersistenceTypeAppCookie),
				CookieName:             gcore.F("cookie_name"),
				PersistenceGranularity: gcore.F("persistence_granularity"),
				PersistenceTimeout:     gcore.F(int64(0)),
			}),
			TimeoutClientData:    gcore.F(int64(50000)),
			TimeoutMemberConnect: gcore.F(int64(50000)),
			TimeoutMemberData:    gcore.F(int64(0)),
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

func TestCloudV1LbpoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1LbpoolListParams{
			Details:        gcore.F(true),
			ListenerID:     gcore.F("listener_id"),
			LoadbalancerID: gcore.F("loadbalancer_id"),
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

func TestCloudV1LbpoolDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
