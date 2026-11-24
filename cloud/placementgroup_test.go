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

func TestPlacementGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.PlacementGroups.New(context.TODO(), cloud.PlacementGroupNewParams{
		ProjectID: gcore.Int(1),
		Name:      "web-affinity",
		Policy:    cloud.PlacementGroupNewParamsPolicyAntiAffinity,
		RegionID:  5,
		DryRun:    gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlacementGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.PlacementGroups.List(context.TODO(), cloud.PlacementGroupListParams{
		ProjectID: gcore.Int(1),
		IDs:       []string{"47003067-550a-4f17-93b6-81ee16ba061e"},
		Limit:     gcore.Int(100),
		Managed:   gcore.Bool(false),
		ManagedBy: cloud.PlacementGroupListParamsManagedBySystem,
		MemberIDs: []string{"6c6aa80f-7836-4dc2-a2ae-125e248be476"},
		Name: cloud.PlacementGroupListParamsName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		Offset:    gcore.Int(0),
		OrderBy:   cloud.PlacementGroupListParamsOrderByCreatedAtDesc,
		Policies:  []string{"anti-affinity"},
		RegionIDs: []int64{1, 5},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlacementGroupDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.PlacementGroups.Delete(
		context.TODO(),
		"47003067-550a-4f17-93b6-81ee16ba061e",
		cloud.PlacementGroupDeleteParams{
			ProjectID: gcore.Int(1),
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

func TestPlacementGroupGet(t *testing.T) {
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
	_, err := client.Cloud.PlacementGroups.Get(
		context.TODO(),
		"47003067-550a-4f17-93b6-81ee16ba061e",
		cloud.PlacementGroupGetParams{
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
