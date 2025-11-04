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

func TestDatabasePostgresClusterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Databases.Postgres.Clusters.New(context.TODO(), cloud.DatabasePostgresClusterNewParams{
		ProjectID:   gcore.Int(0),
		RegionID:    gcore.Int(0),
		ClusterName: "3",
		Flavor: cloud.DatabasePostgresClusterNewParamsFlavor{
			CPU:       1,
			MemoryGib: 1,
		},
		HighAvailability: cloud.DatabasePostgresClusterNewParamsHighAvailability{
			ReplicationMode: "sync",
		},
		Network: cloud.DatabasePostgresClusterNewParamsNetwork{
			ACL: []string{"92.33.34.127"},
		},
		PgServerConfiguration: cloud.DatabasePostgresClusterNewParamsPgServerConfiguration{
			PgConf:  "\nlisten_addresses = 'localhost'\nport = 5432\nmax_connections = 100\nshared_buffers = 128MB\nlogging_collector = on",
			Version: "version",
			Pooler: cloud.DatabasePostgresClusterNewParamsPgServerConfigurationPooler{
				Mode: "transaction",
				Type: "pgbouncer",
			},
		},
		Storage: cloud.DatabasePostgresClusterNewParamsStorage{
			SizeGib: 100,
			Type:    "ssd-hiiops",
		},
		Databases: []cloud.DatabasePostgresClusterNewParamsDatabase{{
			Name:  "mydatabase",
			Owner: "myuser",
		}},
		Users: []cloud.DatabasePostgresClusterNewParamsUser{{
			Name:           "myuser",
			RoleAttributes: []string{"INHERIT"},
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

func TestDatabasePostgresClusterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Databases.Postgres.Clusters.Update(
		context.TODO(),
		"cluster_name",
		cloud.DatabasePostgresClusterUpdateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Databases: []cloud.DatabasePostgresClusterUpdateParamsDatabase{{
				Name:  "mydatabase",
				Owner: "myuser",
			}},
			Flavor: cloud.DatabasePostgresClusterUpdateParamsFlavor{
				CPU:       1,
				MemoryGib: 1,
			},
			HighAvailability: cloud.DatabasePostgresClusterUpdateParamsHighAvailability{
				ReplicationMode: "sync",
			},
			Network: cloud.DatabasePostgresClusterUpdateParamsNetwork{
				ACL: []string{"92.33.34.127"},
			},
			PgServerConfiguration: cloud.DatabasePostgresClusterUpdateParamsPgServerConfiguration{
				PgConf: gcore.String("\nlisten_addresses = 'localhost'\nport = 5432\nmax_connections = 100\nshared_buffers = 128MB\nlogging_collector = on"),
				Pooler: cloud.DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler{
					Mode: "transaction",
					Type: "pgbouncer",
				},
				Version: gcore.String("15"),
			},
			Storage: cloud.DatabasePostgresClusterUpdateParamsStorage{
				SizeGib: 100,
			},
			Users: []cloud.DatabasePostgresClusterUpdateParamsUser{{
				Name:           "myuser",
				RoleAttributes: []string{"INHERIT"},
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

func TestDatabasePostgresClusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Databases.Postgres.Clusters.List(context.TODO(), cloud.DatabasePostgresClusterListParams{
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

func TestDatabasePostgresClusterDelete(t *testing.T) {
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
	_, err := client.Cloud.Databases.Postgres.Clusters.Delete(
		context.TODO(),
		"cluster_name",
		cloud.DatabasePostgresClusterDeleteParams{
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

func TestDatabasePostgresClusterGet(t *testing.T) {
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
	_, err := client.Cloud.Databases.Postgres.Clusters.Get(
		context.TODO(),
		"cluster_name",
		cloud.DatabasePostgresClusterGetParams{
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
