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
			Start:   time.Now(),
			End:     gcore.Time(time.Now()),
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
			Start:       time.Now(),
			Action:      []string{"block", "captcha"},
			End:         gcore.Time(time.Now()),
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
			Start:            time.Now(),
			Actions:          []string{"allow"},
			Countries:        []string{"Mv"},
			End:              gcore.Time(time.Now()),
			IP:               gcore.String(".:"),
			Limit:            gcore.Int(0),
			Offset:           gcore.Int(0),
			Ordering:         gcore.String("ordering"),
			ReferenceID:      gcore.String("2c02efDd09B3BA1AEaDd3dCAa7aC7A37"),
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
			Start:      time.Now(),
			End:        gcore.Time(time.Now()),
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
