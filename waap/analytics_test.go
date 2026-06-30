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

func TestAnalyticsGetFiltersWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Analytics.GetFilters(
		context.TODO(),
		waap.AnalyticsGetFiltersParamsTypeUserAgentClients,
		waap.AnalyticsGetFiltersParams{
			Start:    "2024-04-13T00:00:00+01:00",
			Domains:  []int64{1, 2, 3},
			End:      gcore.String("2024-04-14T12:00:00Z"),
			Limit:    gcore.Int(50),
			Name:     gcore.String("name"),
			Offset:   gcore.Int(0),
			Ordering: waap.AnalyticsGetFiltersParamsOrderingCount,
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
		ExcludeDecision:          []string{"allowed", "blocked"},
		ExcludeDomains:           []int64{0},
		ExcludeHTTPMethods:       []string{"POST", "PUT"},
		ExcludeIPs:               []string{"1.2.3.4", "2001:678:194::3c25:ddad"},
		ExcludeJa3:               []string{"e7d705a3286e19ea42f587b344ee6865"},
		ExcludeJa4:               []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
		ExcludeOptionalAction:    []string{"captcha", "challenge"},
		ExcludeOrganizations:     []string{"Acme corp"},
		ExcludePath:              []string{"/home", "/users/*/profile"},
		ExcludeReferenceIDs:      []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeRequestIDs:        []string{"96763b8fb655e9f18a2e04097b704e39-458959"},
		ExcludeSecurityRuleNames: []string{"SQL injection"},
		ExcludeSessionIDs:        []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeStatusCodes:       []int64{100},
		ExcludeTags:              []string{"mousedevice"},
		ExcludeUserAgent:         []string{"python-requests", "bot"},
		ExcludeUserAgentClients:  []string{"OpenAI GPTBot", "Uptimerobot"},
		ExcludeUserAgentDevices:  []string{"SMART TV"},
		HTTPMethods:              []string{"GET", "HEAD"},
		IPs:                      []string{"1.2.3.4", " 2001:678:194::3c25:ddad"},
		Ja3:                      []string{"e7d705a3286e19ea42f587b344ee6865"},
		Ja4:                      []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
		Limit:                    gcore.Int(0),
		Offset:                   gcore.Int(0),
		OptionalAction:           []string{"captcha", "challenge"},
		Ordering:                 gcore.String("userAgent"),
		Organizations:            []string{"Example Org", "Acme Corp"},
		Path:                     []string{"/home", "/users/*/profile"},
		ReferenceIDs:             []string{"210b9798eb53baa4e69d31c1071cf03d"},
		RequestIDs:               []string{"96763b8fb655e9f18a2e04097b704e39-458959"},
		SecurityRuleNames:        []string{"SQL injection"},
		SessionIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d"},
		StatusCodes:              []int64{100},
		Tags:                     []string{"evasiveclient", "injectionattack"},
		UserAgent:                []string{"Mozilla/5.0", "python-requests"},
		UserAgentClients:         []string{"Chrome", "Firefox"},
		UserAgentDevices:         []string{"mac", "Nokia", "PlayStation"},
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
		ExcludeDecision:          []string{"allowed", "blocked"},
		ExcludeDomains:           []int64{0},
		ExcludeHTTPMethods:       []string{"POST", "PUT"},
		ExcludeIPs:               []string{"1.2.3.4", "2001:678:194::3c25:ddad"},
		ExcludeJa3:               []string{"e7d705a3286e19ea42f587b344ee6865"},
		ExcludeJa4:               []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
		ExcludeOptionalAction:    []string{"captcha", "challenge"},
		ExcludeOrganizations:     []string{"Acme corp"},
		ExcludePath:              []string{"/home", "/users/*/profile"},
		ExcludeReferenceIDs:      []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeRequestIDs:        []string{"96763b8fb655e9f18a2e04097b704e39-458959"},
		ExcludeSecurityRuleNames: []string{"SQL injection"},
		ExcludeSessionIDs:        []string{"210b9798eb53baa4e69d31c1071cf03d"},
		ExcludeStatusCodes:       []int64{100},
		ExcludeTags:              []string{"mousedevice"},
		ExcludeUserAgent:         []string{"python-requests", "bot"},
		ExcludeUserAgentClients:  []string{"OpenAI GPTBot", "Uptimerobot"},
		ExcludeUserAgentDevices:  []string{"SMART TV"},
		HTTPMethods:              []string{"GET", "HEAD"},
		IPs:                      []string{"1.2.3.4", " 2001:678:194::3c25:ddad"},
		Ja3:                      []string{"e7d705a3286e19ea42f587b344ee6865"},
		Ja4:                      []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
		OptionalAction:           []string{"captcha", "challenge"},
		Organizations:            []string{"Example Org", "Acme Corp"},
		Path:                     []string{"/home", "/users/*/profile"},
		ReferenceIDs:             []string{"210b9798eb53baa4e69d31c1071cf03d"},
		RequestIDs:               []string{"96763b8fb655e9f18a2e04097b704e39-458959"},
		SecurityRuleNames:        []string{"SQL injection"},
		SessionIDs:               []string{"210b9798eb53baa4e69d31c1071cf03d"},
		StatusCodes:              []int64{100},
		Tags:                     []string{"evasiveclient", "injectionattack"},
		UserAgent:                []string{"Mozilla/5.0", "python-requests"},
		UserAgentClients:         []string{"Chrome", "Firefox"},
		UserAgentDevices:         []string{"mac", "Nokia", "PlayStation"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
