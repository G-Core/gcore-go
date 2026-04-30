// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/waap"
)

func TestAnalyticsGetEventStatisticsWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Analytics.GetEventStatistics(
		context.TODO(),
		waap.AnalyticsGetEventStatisticsParamsDimensionCountry,
		waap.AnalyticsGetEventStatisticsParams{
			Start:             "2024-04-13T00:00:00+01:00",
			Domains:           []int64{1, 2, 3},
			End:               gcore.String("2024-04-14T12:00:00Z"),
			IPs:               []string{"1.2.3.4", "2001:678:194::3c25:ddad"},
			SecurityRuleNames: []string{"SQL injection"},
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

func TestAnalyticsGetRequestsWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Analytics.GetRequests(context.TODO(), waap.AnalyticsGetRequestsParams{
		Start:                    "2024-04-13T00:00:00+01:00",
		Countries:                []string{"DE", "MY"},
		Decision:                 []string{"allowed", "blocked"},
		Domains:                  []int64{1, 2, 3},
		End:                      gcore.String("2024-04-14T12:00:00Z"),
		ExcludeCountries:         []string{"DE", "MY"},
		ExcludeDomains:           []int64{0},
		ExcludeIPs:               []string{"1.2.3.4", "2001:678:194::3c25:ddad"},
		ExcludeReferenceIDs:      []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeSecurityRuleNames: []string{"SQL injection"},
		ExcludeSessionIDs:        []string{"210b9798eb53baa4e69d31c1071cf03d"},
		HTTPMethods:              []string{"GET", "HEAD"},
		IPs:                      []string{"1.2.3.4", " 2001:678:194::3c25:ddad"},
		Limit:                    gcore.Int(0),
		Offset:                   gcore.Int(0),
		OptionalAction:           []string{"captcha", "challenge"},
		Ordering:                 gcore.String("userAgent"),
		Path:                     gcore.String("/home"),
		ReferenceIDs:             []string{"210b9798eb53baa4e69d31c1071cf03d"},
		RequestIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d-852383"},
		SecurityRuleNames:        []string{"SQL injection"},
		SessionIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d"},
		StatusCodes:              []int64{100},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetTrafficWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Analytics.GetTraffic(context.TODO(), waap.AnalyticsGetTrafficParams{
		Resolution: waap.AnalyticsGetTrafficParamsResolutionDaily,
		Start:      "2024-04-13T00:00:00+01:00",
		BucketSize: 60,
		Domains:    []int64{1, 2, 3},
		End:        gcore.String("2024-04-14T12:00:00Z"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAnalyticsGetTrafficFilteredWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Analytics.GetTrafficFiltered(context.TODO(), waap.AnalyticsGetTrafficFilteredParams{
		Resolution:               waap.AnalyticsGetTrafficFilteredParamsResolutionDaily,
		Start:                    "2024-04-13T00:00:00+01:00",
		BucketSize:               60,
		Countries:                []string{"DE", "MY"},
		Decision:                 []string{"allowed", "blocked"},
		Domains:                  []int64{1, 2, 3},
		End:                      gcore.String("2024-04-14T12:00:00Z"),
		ExcludeCountries:         []string{"DE", "MY"},
		ExcludeDomains:           []int64{0},
		ExcludeIPs:               []string{"1.2.3.4", "2001:678:194::3c25:ddad"},
		ExcludeReferenceIDs:      []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeSecurityRuleNames: []string{"SQL injection"},
		ExcludeSessionIDs:        []string{"210b9798eb53baa4e69d31c1071cf03d"},
		HTTPMethods:              []string{"GET", "HEAD"},
		IPs:                      []string{"1.2.3.4", " 2001:678:194::3c25:ddad"},
		OptionalAction:           []string{"captcha", "challenge"},
		Path:                     gcore.String("/home"),
		ReferenceIDs:             []string{"210b9798eb53baa4e69d31c1071cf03d"},
		RequestIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d-852383"},
		SecurityRuleNames:        []string{"SQL injection"},
		SessionIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d"},
		StatusCodes:              []int64{100},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
