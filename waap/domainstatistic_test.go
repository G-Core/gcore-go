// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/waap"
)

func TestDomainStatisticGetDDOSAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetDDOSAttacks(
		context.TODO(),
		1,
		waap.DomainStatisticGetDDOSAttacksParams{
			EndTime:   gcore.Time(time.Now()),
			Limit:     gcore.Int(0),
			Offset:    gcore.Int(0),
			Ordering:  waap.DomainStatisticGetDDOSAttacksParamsOrderingStartTime,
			StartTime: gcore.Time(time.Now()),
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

func TestDomainStatisticGetDDOSInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetDDOSInfo(
		context.TODO(),
		1,
		waap.DomainStatisticGetDDOSInfoParams{
			GroupBy: waap.DomainStatisticGetDDOSInfoParamsGroupByURL,
			Start:   "2024-04-13T00:00:00+01:00",
			End:     gcore.String("2024-04-14T12:00:00Z"),
			Limit:   gcore.Int(0),
			Offset:  gcore.Int(0),
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

func TestDomainStatisticGetEventsAggregatedWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetEventsAggregated(
		context.TODO(),
		1,
		waap.DomainStatisticGetEventsAggregatedParams{
			Start:       "2024-04-13T00:00:00+01:00",
			Action:      []string{"allow", "block"},
			End:         gcore.String("2024-04-14T12:00:00Z"),
			IP:          []string{"string", "string"},
			ReferenceID: []string{"string", "string"},
			Result:      []string{"passed", "blocked"},
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

func TestDomainStatisticGetRequestDetails(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetRequestDetails(
		context.TODO(),
		"request_id",
		waap.DomainStatisticGetRequestDetailsParams{
			DomainID: 1,
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

func TestDomainStatisticGetRequestsSeriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetRequestsSeries(
		context.TODO(),
		1,
		waap.DomainStatisticGetRequestsSeriesParams{
			Start:            "2024-04-13T00:00:00+01:00",
			Actions:          []string{"allow"},
			Countries:        []string{"Mv"},
			End:              gcore.String("2024-04-14T12:00:00Z"),
			IP:               gcore.String(".:"),
			Limit:            gcore.Int(0),
			Offset:           gcore.Int(0),
			Ordering:         gcore.String("ordering"),
			ReferenceID:      gcore.String("ad07c06f19054e484974fa22e9fb6bb1"),
			SecurityRuleName: gcore.String("security_rule_name"),
			StatusCode:       gcore.Int(100),
			TrafficTypes:     []string{"policy_allowed"},
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

func TestDomainStatisticGetTrafficSeriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Statistics.GetTrafficSeries(
		context.TODO(),
		1,
		waap.DomainStatisticGetTrafficSeriesParams{
			Resolution: waap.DomainStatisticGetTrafficSeriesParamsResolutionDaily,
			Start:      "2024-04-13T00:00:00+01:00",
			End:        gcore.String("2024-04-14T12:00:00Z"),
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
