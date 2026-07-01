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

func TestReservedFixedIPNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.ReservedFixedIPs.New(context.TODO(), cloud.ReservedFixedIPNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(4),
		OfExternal: &cloud.ReservedFixedIPNewParamsBodyExternal{
			IPFamily: cloud.InterfaceIPFamilyDual,
			IsVip:    gcore.Bool(false),
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

func TestReservedFixedIPUpdate(t *testing.T) {
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
	_, err := client.Cloud.ReservedFixedIPs.Update(
		context.TODO(),
		"ac177f1f-eb04-42c4-9864-e7d6486813af",
		cloud.ReservedFixedIPUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(4),
			IsVip:     true,
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

func TestReservedFixedIPListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.ReservedFixedIPs.List(context.TODO(), cloud.ReservedFixedIPListParams{
		ProjectID:     gcore.Int(1),
		RegionID:      gcore.Int(4),
		AvailableOnly: gcore.Bool(true),
		DeviceID:      gcore.String("device_id"),
		ExternalOnly:  gcore.Bool(true),
		InternalOnly:  gcore.Bool(true),
		IPAddress:     gcore.String("ip_address"),
		Limit:         gcore.Int(1000),
		Offset:        gcore.Int(0),
		OrderBy:       cloud.ReservedFixedIPListParamsOrderByCreatedAtAsc,
		VipOnly:       gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReservedFixedIPDelete(t *testing.T) {
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
	_, err := client.Cloud.ReservedFixedIPs.Delete(
		context.TODO(),
		"ac177f1f-eb04-42c4-9864-e7d6486813af",
		cloud.ReservedFixedIPDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(4),
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

func TestReservedFixedIPGet(t *testing.T) {
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
	_, err := client.Cloud.ReservedFixedIPs.Get(
		context.TODO(),
		"ac177f1f-eb04-42c4-9864-e7d6486813af",
		cloud.ReservedFixedIPGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(4),
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
