// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/dns"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestZoneRrsetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.New(
		context.TODO(),
		"rrsetType",
		dns.ZoneRrsetNewParams{
			ZoneName:  "zoneName",
			RrsetName: "rrsetName",
			ResourceRecords: []dns.ZoneRrsetNewParamsResourceRecord{{
				Content: []any{map[string]interface{}{}},
				Enabled: gcore.Bool(true),
				Meta: map[string]any{
					"foo": map[string]interface{}{},
				},
			}},
			Meta: map[string]any{},
			Pickers: []dns.ZoneRrsetNewParamsPicker{{
				Type:   "geodns",
				Limit:  gcore.Int(0),
				Strict: gcore.Bool(true),
			}},
			Ttl: gcore.Int(0),
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

func TestZoneRrsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.List(
		context.TODO(),
		"zoneName",
		dns.ZoneRrsetListParams{
			Limit:          gcore.Int(0),
			Offset:         gcore.Int(0),
			OrderBy:        gcore.String("order_by"),
			OrderDirection: dns.ZoneRrsetListParamsOrderDirectionAsc,
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

func TestZoneRrsetDelete(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.Delete(
		context.TODO(),
		"rrsetType",
		dns.ZoneRrsetDeleteParams{
			ZoneName:  "zoneName",
			RrsetName: "rrsetName",
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

func TestZoneRrsetGet(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.Get(
		context.TODO(),
		"rrsetType",
		dns.ZoneRrsetGetParams{
			ZoneName:  "zoneName",
			RrsetName: "rrsetName",
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

func TestZoneRrsetGetFailoverLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.GetFailoverLogs(
		context.TODO(),
		"rrsetType",
		dns.ZoneRrsetGetFailoverLogsParams{
			ZoneName:  "zoneName",
			RrsetName: "rrsetName",
			Limit:     gcore.Int(0),
			Offset:    gcore.Int(0),
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

func TestZoneRrsetReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Zones.Rrsets.Replace(
		context.TODO(),
		"rrsetType",
		dns.ZoneRrsetReplaceParams{
			ZoneName:  "zoneName",
			RrsetName: "rrsetName",
			ResourceRecords: []dns.ZoneRrsetReplaceParamsResourceRecord{{
				Content: []any{map[string]interface{}{}},
				Enabled: gcore.Bool(true),
				Meta: map[string]any{
					"foo": map[string]interface{}{},
				},
			}},
			Meta: map[string]any{},
			Pickers: []dns.ZoneRrsetReplaceParamsPicker{{
				Type:   "geodns",
				Limit:  gcore.Int(0),
				Strict: gcore.Bool(true),
			}},
			Ttl: gcore.Int(0),
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
