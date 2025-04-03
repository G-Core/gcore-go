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

func TestCloudV1DbaaPostgreCheckQuotaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Dbaas.Postgres.CheckQuota(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1DbaaPostgreCheckQuotaParams{
			PgClusterCreate: gcore.PgClusterCreateParam{
				ClusterName: gcore.F("3"),
				Flavor: gcore.F(gcore.FlavorDatabaseParam{
					CPU:       gcore.F(int64(1)),
					MemoryGib: gcore.F(int64(1)),
				}),
				HighAvailability: gcore.F(gcore.HighAvailabilityOptionsParam{
					ReplicationMode: gcore.F(gcore.HighAvailabilityOptionsReplicationModeAsync),
				}),
				Network: gcore.F(gcore.PublicNetworkParam{
					ACL:         gcore.F([]string{"string"}),
					NetworkType: gcore.F(gcore.PublicNetworkNetworkTypePublic),
				}),
				PgServerConfiguration: gcore.F(gcore.PostgresqlServerConfigParam{
					PgConf:  gcore.F("pg_conf"),
					Version: gcore.F("version"),
					Pooler: gcore.F(gcore.PoolerConfigurationParam{
						Mode: gcore.F(gcore.PoolerConfigurationModeSession),
						Type: gcore.F(gcore.PoolerConfigurationTypePgbouncer),
					}),
				}),
				Storage: gcore.F(gcore.StorageParam{
					SizeGib: gcore.F(int64(1)),
					Type:    gcore.F("type"),
				}),
				Databases: gcore.F([]gcore.DatabaseParam{{
					Name:  gcore.F("name"),
					Owner: gcore.F("owner"),
				}}),
				Users: gcore.F([]gcore.UserParam{{
					Name:           gcore.F("name"),
					RoleAttributes: gcore.F([]gcore.UserRoleAttribute{gcore.UserRoleAttributeBypassrls}),
				}}),
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

func TestCloudV1DbaaPostgreGetConfigurations(t *testing.T) {
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
	_, err := client.Cloud.V1.Dbaas.Postgres.GetConfigurations(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1DbaaPostgreValidateConfiguration(t *testing.T) {
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
	_, err := client.Cloud.V1.Dbaas.Postgres.ValidateConfiguration(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1DbaaPostgreValidateConfigurationParams{
			PgConf:  gcore.F("pg_conf"),
			Version: gcore.F("version"),
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
