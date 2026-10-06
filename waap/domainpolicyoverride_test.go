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

func TestDomainPolicyOverrideNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.PolicyOverrides.New(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideNewParams{
			Enabled:    true,
			Name:       "x",
			TargetIDs:  []int64{1},
			TargetType: waap.DomainPolicyOverrideNewParamsTargetTypeDetector,
			Conditions: []waap.DomainPolicyOverrideNewParamsCondition{{
				Country: waap.DomainPolicyOverrideNewParamsConditionCountry{
					CountryCode: []string{"Mv"},
					Negation:    gcore.Bool(true),
				},
				Header: waap.DomainPolicyOverrideNewParamsConditionHeader{
					Header:    "Origin",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				HeaderExists: waap.DomainPolicyOverrideNewParamsConditionHeaderExists{
					Header:   "Origin",
					Negation: gcore.Bool(true),
				},
				HTTPMethod: waap.DomainPolicyOverrideNewParamsConditionHTTPMethod{
					HTTPMethod: "CONNECT",
					Negation:   gcore.Bool(true),
				},
				IP: waap.DomainPolicyOverrideNewParamsConditionIP{
					IPAddress: "ip_address",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainPolicyOverrideNewParamsConditionIPRange{
					LowerBound: "lower_bound",
					UpperBound: "upper_bound",
					Negation:   gcore.Bool(true),
				},
				Ja3: waap.DomainPolicyOverrideNewParamsConditionJa3{
					Ja3Fingerprints: []string{"e7d705a3286e19ea42f587b344ee6865"},
					Negation:        gcore.Bool(true),
				},
				Ja4: waap.DomainPolicyOverrideNewParamsConditionJa4{
					Ja4Fingerprints: []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
					Negation:        gcore.Bool(true),
				},
				Organization: waap.DomainPolicyOverrideNewParamsConditionOrganization{
					Organization: "UptimeRobot s.r.o",
					Negation:     gcore.Bool(true),
				},
				URL: waap.DomainPolicyOverrideNewParamsConditionURL{
					URL:       "/wp-admin/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserAgent: waap.DomainPolicyOverrideNewParamsConditionUserAgent{
					UserAgent: "curl/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
			}},
			Description: gcore.String("description"),
			RequestID:   gcore.String("210b9798eb53baa4e69d31c1071cf03d-852383"),
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

func TestDomainPolicyOverrideUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.PolicyOverrides.Update(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideUpdateParams{
			DomainID: 1,
			Conditions: []waap.DomainPolicyOverrideUpdateParamsCondition{{
				Country: waap.DomainPolicyOverrideUpdateParamsConditionCountry{
					CountryCode: []string{"Mv"},
					Negation:    gcore.Bool(true),
				},
				Header: waap.DomainPolicyOverrideUpdateParamsConditionHeader{
					Header:    "Origin",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				HeaderExists: waap.DomainPolicyOverrideUpdateParamsConditionHeaderExists{
					Header:   "Origin",
					Negation: gcore.Bool(true),
				},
				HTTPMethod: waap.DomainPolicyOverrideUpdateParamsConditionHTTPMethod{
					HTTPMethod: "CONNECT",
					Negation:   gcore.Bool(true),
				},
				IP: waap.DomainPolicyOverrideUpdateParamsConditionIP{
					IPAddress: "ip_address",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainPolicyOverrideUpdateParamsConditionIPRange{
					LowerBound: "lower_bound",
					UpperBound: "upper_bound",
					Negation:   gcore.Bool(true),
				},
				Ja3: waap.DomainPolicyOverrideUpdateParamsConditionJa3{
					Ja3Fingerprints: []string{"e7d705a3286e19ea42f587b344ee6865"},
					Negation:        gcore.Bool(true),
				},
				Ja4: waap.DomainPolicyOverrideUpdateParamsConditionJa4{
					Ja4Fingerprints: []string{"t13d3113h2_e8f1e7e78f70_ce5650b735ce"},
					Negation:        gcore.Bool(true),
				},
				Organization: waap.DomainPolicyOverrideUpdateParamsConditionOrganization{
					Organization: "UptimeRobot s.r.o",
					Negation:     gcore.Bool(true),
				},
				URL: waap.DomainPolicyOverrideUpdateParamsConditionURL{
					URL:       "/wp-admin/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserAgent: waap.DomainPolicyOverrideUpdateParamsConditionUserAgent{
					UserAgent: "curl/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
			}},
			Description: gcore.String("description"),
			Enabled:     gcore.Bool(true),
			Name:        gcore.String("x"),
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

func TestDomainPolicyOverrideListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.PolicyOverrides.List(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideListParams{
			Enabled:    gcore.Bool(true),
			Limit:      gcore.Int(0),
			Offset:     gcore.Int(0),
			Ordering:   waap.DomainPolicyOverrideListParamsOrderingID,
			Search:     gcore.String("x"),
			TargetType: waap.DomainPolicyOverrideListParamsTargetTypeDetector,
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

func TestDomainPolicyOverrideDelete(t *testing.T) {
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
	err := client.Waap.Domains.PolicyOverrides.Delete(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideDeleteParams{
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

func TestDomainPolicyOverrideGet(t *testing.T) {
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
	_, err := client.Waap.Domains.PolicyOverrides.Get(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideGetParams{
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

func TestDomainPolicyOverrideListOwaspTopThreatsTargetsWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.PolicyOverrides.ListOwaspTopThreatsTargets(
		context.TODO(),
		1,
		waap.DomainPolicyOverrideListOwaspTopThreatsTargetsParams{
			Limit:  gcore.Int(0),
			Offset: gcore.Int(0),
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
