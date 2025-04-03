// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestWaapV1DomainCustomRuleNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Waap.V1.Domains.CustomRules.New(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainCustomRuleNewParams{
			Action: gcore.F(gcore.CustomerRuleActionInputParam{
				Allow: gcore.F[any](map[string]interface{}{}),
				Block: gcore.F(gcore.RuleBlockActionParam{
					ActionDuration: gcore.F("12h"),
					StatusCode:     gcore.F(gcore.RuleBlockActionStatusCode403),
				}),
				Captcha:   gcore.F[any](map[string]interface{}{}),
				Handshake: gcore.F[any](map[string]interface{}{}),
				Monitor:   gcore.F[any](map[string]interface{}{}),
				Tag: gcore.F(gcore.RuleTagActionParam{
					Tags: gcore.F([]string{"string"}),
				}),
			}),
			Conditions: gcore.F([]gcore.CustomRuleConditionInputParam{{
				ContentType: gcore.F(gcore.ContentTypeConditionParam{
					ContentType: gcore.F([]string{"string"}),
					Negation:    gcore.F(true),
				}),
				Country: gcore.F(gcore.CountryConditionParam{
					CountryCode: gcore.F([]string{"Mv"}),
					Negation:    gcore.F(true),
				}),
				FileExtension: gcore.F(gcore.FileExtensionConditionParam{
					FileExtension: gcore.F([]string{"string"}),
					Negation:      gcore.F(true),
				}),
				Header: gcore.F(gcore.HeaderConditionParam{
					Header:    gcore.F("header"),
					Value:     gcore.F("value"),
					MatchType: gcore.F(gcore.HeaderConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				HeaderExists: gcore.F(gcore.HeaderExistsConditionParam{
					Header:   gcore.F("header"),
					Negation: gcore.F(true),
				}),
				HTTPMethod: gcore.F(gcore.HTTPMethodConditionParam{
					HTTPMethod: gcore.F(gcore.HTTPMethodCustomRuleConnect),
					Negation:   gcore.F(true),
				}),
				IP: gcore.F(gcore.IPConditionParam{
					IPAddress: gcore.F("ip_address"),
					Negation:  gcore.F(true),
				}),
				IPRange: gcore.F(gcore.IPRangeConditionParam{
					LowerBound: gcore.F("lower_bound"),
					UpperBound: gcore.F("upper_bound"),
					Negation:   gcore.F(true),
				}),
				Organization: gcore.F(gcore.OrganizationConditionParam{
					Organization: gcore.F("organization"),
					Negation:     gcore.F(true),
				}),
				OwnerTypes: gcore.F(gcore.OwnerTypesConditionParam{
					Negation:   gcore.F(true),
					OwnerTypes: gcore.F([]gcore.OwnerTypesConditionOwnerType{gcore.OwnerTypesConditionOwnerTypeCommercial}),
				}),
				RequestRate: gcore.F(gcore.RequestRateConditionParam{
					PathPattern:    gcore.F("/"),
					Requests:       gcore.F(int64(20)),
					Time:           gcore.F(int64(1)),
					HTTPMethods:    gcore.F([]gcore.HTTPMethodCustomRule{gcore.HTTPMethodCustomRuleConnect}),
					IPs:            gcore.F([]string{"string"}),
					Negation:       gcore.F(true),
					UserDefinedTag: gcore.F("SQfNklznVLBBpr"),
				}),
				ResponseHeader: gcore.F(gcore.ResponseHeaderConditionParam{
					Header:    gcore.F("header"),
					Value:     gcore.F("value"),
					MatchType: gcore.F(gcore.ResponseHeaderConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				ResponseHeaderExists: gcore.F(gcore.ResponseHeaderExistsConditionParam{
					Header:   gcore.F("header"),
					Negation: gcore.F(true),
				}),
				SessionRequestCount: gcore.F(gcore.SessionRequestCountConditionParam{
					RequestCount: gcore.F(int64(1)),
					Negation:     gcore.F(true),
				}),
				Tags: gcore.F(gcore.TagsConditionParam{
					Tags:     gcore.F([]string{"string"}),
					Negation: gcore.F(true),
				}),
				URL: gcore.F(gcore.URLConditionParam{
					URL:       gcore.F("url"),
					MatchType: gcore.F(gcore.URLConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				UserAgent: gcore.F(gcore.UserAgentConditionParam{
					UserAgent: gcore.F("user_agent"),
					MatchType: gcore.F(gcore.UserAgentConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				UserDefinedTags: gcore.F(gcore.UserDefinedTagsConditionParam{
					Tags:     gcore.F([]string{"string"}),
					Negation: gcore.F(true),
				}),
			}}),
			Enabled:     gcore.F(true),
			Name:        gcore.F("name"),
			Description: gcore.F("description"),
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

func TestWaapV1DomainCustomRuleGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Waap.V1.Domains.CustomRules.Get(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1DomainCustomRuleUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Waap.V1.Domains.CustomRules.Update(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.WaapV1DomainCustomRuleUpdateParams{
			Action: gcore.F(gcore.CustomerRuleActionInputParam{
				Allow: gcore.F[any](map[string]interface{}{}),
				Block: gcore.F(gcore.RuleBlockActionParam{
					ActionDuration: gcore.F("12h"),
					StatusCode:     gcore.F(gcore.RuleBlockActionStatusCode403),
				}),
				Captcha:   gcore.F[any](map[string]interface{}{}),
				Handshake: gcore.F[any](map[string]interface{}{}),
				Monitor:   gcore.F[any](map[string]interface{}{}),
				Tag: gcore.F(gcore.RuleTagActionParam{
					Tags: gcore.F([]string{"string"}),
				}),
			}),
			Conditions: gcore.F([]gcore.CustomRuleConditionInputParam{{
				ContentType: gcore.F(gcore.ContentTypeConditionParam{
					ContentType: gcore.F([]string{"string"}),
					Negation:    gcore.F(true),
				}),
				Country: gcore.F(gcore.CountryConditionParam{
					CountryCode: gcore.F([]string{"Mv"}),
					Negation:    gcore.F(true),
				}),
				FileExtension: gcore.F(gcore.FileExtensionConditionParam{
					FileExtension: gcore.F([]string{"string"}),
					Negation:      gcore.F(true),
				}),
				Header: gcore.F(gcore.HeaderConditionParam{
					Header:    gcore.F("header"),
					Value:     gcore.F("value"),
					MatchType: gcore.F(gcore.HeaderConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				HeaderExists: gcore.F(gcore.HeaderExistsConditionParam{
					Header:   gcore.F("header"),
					Negation: gcore.F(true),
				}),
				HTTPMethod: gcore.F(gcore.HTTPMethodConditionParam{
					HTTPMethod: gcore.F(gcore.HTTPMethodCustomRuleConnect),
					Negation:   gcore.F(true),
				}),
				IP: gcore.F(gcore.IPConditionParam{
					IPAddress: gcore.F("ip_address"),
					Negation:  gcore.F(true),
				}),
				IPRange: gcore.F(gcore.IPRangeConditionParam{
					LowerBound: gcore.F("lower_bound"),
					UpperBound: gcore.F("upper_bound"),
					Negation:   gcore.F(true),
				}),
				Organization: gcore.F(gcore.OrganizationConditionParam{
					Organization: gcore.F("organization"),
					Negation:     gcore.F(true),
				}),
				OwnerTypes: gcore.F(gcore.OwnerTypesConditionParam{
					Negation:   gcore.F(true),
					OwnerTypes: gcore.F([]gcore.OwnerTypesConditionOwnerType{gcore.OwnerTypesConditionOwnerTypeCommercial}),
				}),
				RequestRate: gcore.F(gcore.RequestRateConditionParam{
					PathPattern:    gcore.F("/"),
					Requests:       gcore.F(int64(20)),
					Time:           gcore.F(int64(1)),
					HTTPMethods:    gcore.F([]gcore.HTTPMethodCustomRule{gcore.HTTPMethodCustomRuleConnect}),
					IPs:            gcore.F([]string{"string"}),
					Negation:       gcore.F(true),
					UserDefinedTag: gcore.F("SQfNklznVLBBpr"),
				}),
				ResponseHeader: gcore.F(gcore.ResponseHeaderConditionParam{
					Header:    gcore.F("header"),
					Value:     gcore.F("value"),
					MatchType: gcore.F(gcore.ResponseHeaderConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				ResponseHeaderExists: gcore.F(gcore.ResponseHeaderExistsConditionParam{
					Header:   gcore.F("header"),
					Negation: gcore.F(true),
				}),
				SessionRequestCount: gcore.F(gcore.SessionRequestCountConditionParam{
					RequestCount: gcore.F(int64(1)),
					Negation:     gcore.F(true),
				}),
				Tags: gcore.F(gcore.TagsConditionParam{
					Tags:     gcore.F([]string{"string"}),
					Negation: gcore.F(true),
				}),
				URL: gcore.F(gcore.URLConditionParam{
					URL:       gcore.F("url"),
					MatchType: gcore.F(gcore.URLConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				UserAgent: gcore.F(gcore.UserAgentConditionParam{
					UserAgent: gcore.F("user_agent"),
					MatchType: gcore.F(gcore.UserAgentConditionMatchTypeExact),
					Negation:  gcore.F(true),
				}),
				UserDefinedTags: gcore.F(gcore.UserDefinedTagsConditionParam{
					Tags:     gcore.F([]string{"string"}),
					Negation: gcore.F(true),
				}),
			}}),
			Description: gcore.F("description"),
			Enabled:     gcore.F(true),
			Name:        gcore.F("name"),
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

func TestWaapV1DomainCustomRuleListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Waap.V1.Domains.CustomRules.List(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainCustomRuleListParams{
			Action:      gcore.F(gcore.RuleActionTypeAllow),
			Description: gcore.F("description"),
			Enabled:     gcore.F(true),
			Limit:       gcore.F(int64(0)),
			Name:        gcore.F("name"),
			Offset:      gcore.F(int64(0)),
			Ordering:    gcore.F(gcore.WaapV1DomainCustomRuleListParamsOrderingID),
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

func TestWaapV1DomainCustomRuleDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Waap.V1.Domains.CustomRules.Delete(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1DomainCustomRuleBulkDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Waap.V1.Domains.CustomRules.BulkDelete(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainCustomRuleBulkDeleteParams{
			RulesBulkDelete: gcore.RulesBulkDeleteParam{
				RuleIDs: gcore.F([]int64{int64(0)}),
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

func TestWaapV1DomainCustomRuleToggle(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Waap.V1.Domains.CustomRules.Toggle(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CustomerRuleStateEnable,
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
