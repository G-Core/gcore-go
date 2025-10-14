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

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.New(context.TODO(), cloud.LoadBalancerNewParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		Flavor:    gcore.String("lb1-1-2"),
		FloatingIP: cloud.LoadBalancerNewParamsFloatingIPUnion{
			OfExisting: &cloud.LoadBalancerNewParamsFloatingIPExisting{
				ExistingFloatingID: "c64e5db1-5f1f-43ec-a8d9-5090df85b82d",
			},
		},
		Listeners: []cloud.LoadBalancerNewParamsListener{{
			Name:             "my_listener",
			Protocol:         cloud.LbListenerProtocolHTTP,
			ProtocolPort:     80,
			AllowedCidrs:     []string{"10.0.0.0/8"},
			ConnectionLimit:  gcore.Int(100000),
			InsertXForwarded: gcore.Bool(false),
			Pools: []cloud.LoadBalancerNewParamsListenerPool{{
				LbAlgorithm: cloud.LbAlgorithmLeastConnections,
				Name:        "pool_name",
				Protocol:    cloud.LbPoolProtocolHTTP,
				CaSecretID:  gcore.String("ca_secret_id"),
				CrlSecretID: gcore.String("crl_secret_id"),
				Healthmonitor: cloud.LoadBalancerNewParamsListenerPoolHealthmonitor{
					Delay:          10,
					MaxRetries:     3,
					Timeout:        5,
					Type:           cloud.LbHealthMonitorTypeHTTP,
					ExpectedCodes:  gcore.String("200,301,302"),
					HTTPMethod:     cloud.HTTPMethodGet,
					MaxRetriesDown: gcore.Int(3),
					URLPath:        gcore.String("/"),
				},
				ListenerID:     gcore.String("listener_id"),
				LoadBalancerID: gcore.String("bbb35f84-35cc-4b2f-84c2-a6a29bba68aa"),
				Members: []cloud.LoadBalancerNewParamsListenerPoolMember{{
					Address:        "192.168.1.101",
					ProtocolPort:   8000,
					AdminStateUp:   gcore.Bool(true),
					Backup:         gcore.Bool(true),
					InstanceID:     gcore.String("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
					MonitorAddress: gcore.String("monitor_address"),
					MonitorPort:    gcore.Int(0),
					SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
					Weight:         gcore.Int(2),
				}, {
					Address:        "192.168.1.102",
					ProtocolPort:   8000,
					AdminStateUp:   gcore.Bool(true),
					Backup:         gcore.Bool(true),
					InstanceID:     gcore.String("169942e0-9b53-42df-95ef-1a8b6525c2bd"),
					MonitorAddress: gcore.String("monitor_address"),
					MonitorPort:    gcore.Int(0),
					SubnetID:       gcore.String("32283b0b-b560-4690-810c-f672cbb2e28d"),
					Weight:         gcore.Int(4),
				}},
				SecretID: gcore.String("secret_id"),
				SessionPersistence: cloud.LoadBalancerNewParamsListenerPoolSessionPersistence{
					Type:                   cloud.LbSessionPersistenceTypeAppCookie,
					CookieName:             gcore.String("cookie_name"),
					PersistenceGranularity: gcore.String("persistence_granularity"),
					PersistenceTimeout:     gcore.Int(0),
				},
				TimeoutClientData:    gcore.Int(50000),
				TimeoutMemberConnect: gcore.Int(50000),
				TimeoutMemberData:    gcore.Int(0),
			}},
			SecretID:             gcore.String("f2e734d0-fa2b-42c2-ad33-4c6db5101e00"),
			SniSecretID:          []string{"f2e734d0-fa2b-42c2-ad33-4c6db5101e00", "eb121225-7ded-4ff3-ae1f-599e145dd7cb"},
			TimeoutClientData:    gcore.Int(50000),
			TimeoutMemberConnect: gcore.Int(50000),
			TimeoutMemberData:    param.Null[int64](),
			UserList: []cloud.LoadBalancerNewParamsListenerUserList{{
				EncryptedPassword: "$5$isRr.HJ1IrQP38.m$oViu3DJOpUG2ZsjCBtbITV3mqpxxbZfyWJojLPNSPO5",
				Username:          "admin",
			}},
		}},
		Logging: cloud.LoadBalancerNewParamsLogging{
			DestinationRegionID: gcore.Int(1),
			Enabled:             gcore.Bool(true),
			RetentionPolicy: cloud.LaasIndexRetentionPolicyParam{
				Period: gcore.Int(45),
			},
			TopicName: gcore.String("my-log-name"),
		},
		Name:                  gcore.String("new_load_balancer"),
		NameTemplate:          gcore.String("lb_name_template"),
		PreferredConnectivity: cloud.LoadBalancerMemberConnectivityL2,
		Tags: map[string]string{
			"my-tag": "my-tag-value",
		},
		VipIPFamily:  cloud.InterfaceIPFamilyDual,
		VipNetworkID: gcore.String("ac307687-31a4-4a11-a949-6bea1b2878f5"),
		VipPortID:    gcore.String("ff83e13a-b256-4be2-ba5d-028d3f0ab450"),
		VipSubnetID:  gcore.String("4e7802d3-5023-44b8-b298-7726558fddf4"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.List(context.TODO(), cloud.LoadBalancerListParams{
		ProjectID:        gcore.Int(0),
		RegionID:         gcore.Int(0),
		AssignedFloating: gcore.Bool(true),
		Limit:            gcore.Int(0),
		LoggingEnabled:   gcore.Bool(true),
		Name:             gcore.String("name"),
		Offset:           gcore.Int(0),
		OrderBy:          gcore.String("order_by"),
		ShowStats:        gcore.Bool(true),
		TagKey:           []string{"string"},
		TagKeyValue:      gcore.String("tag_key_value"),
		WithDDOS:         gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
