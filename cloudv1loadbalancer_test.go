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

func TestCloudV1LoadbalancerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1LoadbalancerNewParams{
			Flavor: gcore.F("lb1-1-2"),
			FloatingIP: gcore.F(gcore.NewInterfaceFloatingIPPydanticParam{
				Source:             gcore.F(gcore.NewInterfaceFloatingIPPydanticSourceExisting),
				ExistingFloatingID: gcore.F("e3c6ee77-48cb-416b-b204-11b492cc776e3"),
			}),
			Listeners: gcore.F([]gcore.CloudV1LoadbalancerNewParamsListener{{
				Name:             gcore.F("my_listener"),
				Protocol:         gcore.F(gcore.LbListenerProtocolHTTP),
				ProtocolPort:     gcore.F(int64(80)),
				AllowedCidrs:     gcore.F([]string{"10.0.0.0/8"}),
				ConnectionLimit:  gcore.F(int64(100000)),
				InsertXForwarded: gcore.F(false),
				Pools: gcore.F([]gcore.CreateLbPoolParam{{
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
				}}),
				SecretID:             gcore.F("f2e734d0-fa2b-42c2-ad33-4c6db5101e00"),
				SniSecretID:          gcore.F([]string{"f2e734d0-fa2b-42c2-ad33-4c6db5101e00", "eb121225-7ded-4ff3-ae1f-599e145dd7cb"}),
				TimeoutClientData:    gcore.F(int64(50000)),
				TimeoutMemberConnect: gcore.F(int64(50000)),
				TimeoutMemberData:    gcore.F(int64(0)),
				UserList: gcore.F([]gcore.UserListItemParam{{
					EncryptedPassword: gcore.F("$5$isRr.HJ1IrQP38.m$oViu3DJOpUG2ZsjCBtbITV3mqpxxbZfyWJojLPNSPO5"),
					Username:          gcore.F("admin"),
				}}),
			}}),
			Logging: gcore.F(gcore.LoadbalancerLoggingParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			Metadata: gcore.F[any](map[string]interface{}{
				"type": "standard",
			}),
			Name:                  gcore.F("new_load_balancer"),
			NameTemplate:          gcore.F("lb_name_template"),
			PreferredConnectivity: gcore.F(gcore.MemberConnectivityL2),
			Tag:                   gcore.F([]string{"k8s"}),
			VipIPFamily:           gcore.F(gcore.InterfaceIPFamilyDual),
			VipNetworkID:          gcore.F("ac307687-31a4-4a11-a949-6bea1b2878f5"),
			VipPortID:             gcore.F("ff83e13a-b256-4be2-ba5d-028d3f0ab450"),
			VipSubnetID:           gcore.F("4e7802d3-5023-44b8-b298-7726558fddf4"),
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

func TestCloudV1LoadbalancerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
		gcore.CloudV1LoadbalancerGetParams{
			ShowStats: gcore.F(true),
			WithDdos:  gcore.F(true),
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

func TestCloudV1LoadbalancerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
		gcore.CloudV1LoadbalancerUpdateParams{
			Logging: gcore.F(gcore.LoadbalancerLoggingParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			Name:                  gcore.F("some_name"),
			PreferredConnectivity: gcore.F(gcore.MemberConnectivityL2),
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

func TestCloudV1LoadbalancerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1LoadbalancerListParams{
			AssignedFloating: gcore.F(true),
			Limit:            gcore.F(int64(0)),
			LoggingEnabled:   gcore.F(true),
			MetadataK:        gcore.F("metadata_k"),
			MetadataKv:       gcore.F("metadata_kv"),
			Name:             gcore.F("name"),
			Offset:           gcore.F(int64(0)),
			OrderBy:          gcore.F("order_by"),
			ShowStats:        gcore.F(true),
			WithDdos:         gcore.F(true),
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

func TestCloudV1LoadbalancerDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1LoadbalancerCheckQuotaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.CheckQuota(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1LoadbalancerCheckQuotaParams{
			FloatingIP: gcore.F(gcore.NewInterfaceFloatingIPParam{
				Source:             gcore.F(gcore.NewInterfaceFloatingIPSourceExisting),
				ExistingFloatingID: gcore.F("existing_floating_id"),
			}),
			Logging: gcore.F(gcore.CloudV1LoadbalancerCheckQuotaParamsLogging{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(false),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some_topic_name"),
			}),
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

func TestCloudV1LoadbalancerFailoverWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.Failover(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
		gcore.CloudV1LoadbalancerFailoverParams{
			Force: gcore.F(true),
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

func TestCloudV1LoadbalancerGetMetrics(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.GetMetrics(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
		gcore.CloudV1LoadbalancerGetMetricsParams{
			TimeInterval: gcore.F(int64(6)),
			TimeUnit:     gcore.F(gcore.InstanceMetricsTimeUnitDay),
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

func TestCloudV1LoadbalancerResize(t *testing.T) {
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
	_, err := client.Cloud.V1.Loadbalancers.Resize(
		context.TODO(),
		int64(0),
		int64(0),
		"loadbalancer_id",
		gcore.CloudV1LoadbalancerResizeParams{
			Flavor: gcore.F("lb1-2-4"),
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
