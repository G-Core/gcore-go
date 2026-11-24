// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestBaremetalServerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.New(context.TODO(), cloud.BaremetalServerNewParams{
		ProjectID:  gcore.Int(1),
		FlavorName: "bm2-hf-medium",
		ImageID:    "2fa8c5b1-1234-4abc-9def-0123456789ab",
		Interfaces: []cloud.BaremetalServerNewParamsInterface{{
			PortID:    "75757793-5e33-4eed-ab5d-9ed5030e57f7",
			PortGroup: gcore.Int(0),
		}},
		RegionID: 7,
		DryRun:   gcore.Bool(true),
		Name:     gcore.String("bm-1"),
		Password: gcore.String("x"),
		SSHKeyID: gcore.String("21b3c2d5-efe8-4367-99ff-86e1e060d27f"),
		Tags: map[string]string{
			"my-tag": "my-tag-value",
		},
		UserData: gcore.String("x"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBaremetalServerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Update(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.BaremetalServerUpdateParams{
			ProjectID: gcore.Int(1),
			DryRun:    gcore.Bool(true),
			Interfaces: []cloud.BaremetalServerUpdateParamsInterface{{
				PortID:    "port_id",
				PortGroup: gcore.Int(0),
			}},
			Name: gcore.String("bm-1"),
			Tags: cloud.TagUpdateMap{
				"my-tag":           gcore.Ptr("my-tag-value"),
				"my-tag-to-remove": nil,
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

func TestBaremetalServerListWithOptionalParams(t *testing.T) {
	t.Skip("GCLOUD2-27612: the mock rejects arrays inside object query filters")
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
	_, err := client.Cloud.Baremetal.Servers.List(context.TODO(), cloud.BaremetalServerListParams{
		ProjectID: gcore.Int(1),
		CreatedAt: cloud.BaremetalServerListParamsCreatedAt{
			Gt:  gcore.Time(time.Now()),
			Gte: gcore.Time(time.Now()),
			Lt:  gcore.Time(time.Now()),
			Lte: gcore.Time(time.Now()),
		},
		FlavorName: cloud.BaremetalServerListParamsFlavorName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		IDs: []string{"string"},
		IPAddress: cloud.BaremetalServerListParamsIPAddress{
			Contains: []string{"string"},
			Exact:    []string{"string"},
		},
		Limit:                  gcore.Int(1),
		Managed:                gcore.Bool(true),
		ManagedBy:              cloud.BaremetalServerListParamsManagedByK8SCluster,
		ManagedByResourceIDs:   []string{"string"},
		ManagedByResourceNames: []string{"string"},
		Name: cloud.BaremetalServerListParamsName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		NetworkIDs: []string{"string"},
		Offset:     gcore.Int(0),
		OrderBy:    cloud.BaremetalServerListParamsOrderByCreatedAtAsc,
		RegionIDs:  []int64{0},
		Statuses:   []string{"active"},
		SubnetIDs:  []string{"string"},
		TagKey: cloud.BaremetalServerListParamsTagKey{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		TagValue: cloud.BaremetalServerListParamsTagValue{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		Tags: map[string][]string{
			"env": {"prod"},
		},
		UpdatedAt: cloud.BaremetalServerListParamsUpdatedAt{
			Gt:  gcore.Time(time.Now()),
			Gte: gcore.Time(time.Now()),
			Lt:  gcore.Time(time.Now()),
			Lte: gcore.Time(time.Now()),
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

func TestBaremetalServerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Delete(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.BaremetalServerDeleteParams{
			ProjectID:      gcore.Int(1),
			AllFloatingIPs: gcore.Bool(true),
			AllPorts:       gcore.Bool(true),
			DryRun:         gcore.Bool(true),
			FloatingIPIDs:  []string{"string"},
			PortIDs:        []string{"string"},
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

func TestBaremetalServerActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Action(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.BaremetalServerActionParams{
			ProjectID: gcore.Int(1),
			Action:    cloud.BaremetalServerActionParamsActionStop,
			DryRun:    gcore.Bool(true),
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

func TestBaremetalServerGet(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Get(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.BaremetalServerGetParams{
			ProjectID: gcore.Int(1),
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

func TestBaremetalServerRebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Rebuild(
		context.TODO(),
		"024a29e-b4b7-4c91-9a46-505be123d9f8",
		cloud.BaremetalServerRebuildParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			ImageID:   gcore.String("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
			UserData:  gcore.String("aGVsbG9fd29ybGQ="),
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
