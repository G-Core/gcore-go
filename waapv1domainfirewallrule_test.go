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

func TestWaapV1DomainFirewallRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.FirewallRules.New(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainFirewallRuleNewParams{
			Action: gcore.F(gcore.WaapV1DomainFirewallRuleNewParamsAction{
				Allow: gcore.F[any](map[string]interface{}{}),
				Block: gcore.F(gcore.RuleBlockActionParam{
					ActionDuration: gcore.F("12h"),
					StatusCode:     gcore.F(gcore.RuleBlockActionStatusCode403),
				}),
			}),
			Conditions: gcore.F([]gcore.FirewallRuleConditionParam{{
				IP: gcore.F(gcore.IPConditionParam{
					IPAddress: gcore.F("ip_address"),
					Negation:  gcore.F(true),
				}),
				IPRange: gcore.F(gcore.IPRangeConditionParam{
					LowerBound: gcore.F("lower_bound"),
					UpperBound: gcore.F("upper_bound"),
					Negation:   gcore.F(true),
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

func TestWaapV1DomainFirewallRuleGet(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.FirewallRules.Get(
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

func TestWaapV1DomainFirewallRuleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.V1.Domains.FirewallRules.Update(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.WaapV1DomainFirewallRuleUpdateParams{
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
			Conditions: gcore.F([]gcore.FirewallRuleConditionParam{{
				IP: gcore.F(gcore.IPConditionParam{
					IPAddress: gcore.F("ip_address"),
					Negation:  gcore.F(true),
				}),
				IPRange: gcore.F(gcore.IPRangeConditionParam{
					LowerBound: gcore.F("lower_bound"),
					UpperBound: gcore.F("upper_bound"),
					Negation:   gcore.F(true),
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

func TestWaapV1DomainFirewallRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.FirewallRules.List(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainFirewallRuleListParams{
			Action:      gcore.F(gcore.RuleActionTypeAllow),
			Description: gcore.F("description"),
			Enabled:     gcore.F(true),
			Limit:       gcore.F(int64(0)),
			Name:        gcore.F("name"),
			Offset:      gcore.F(int64(0)),
			Ordering:    gcore.F(gcore.WaapV1DomainFirewallRuleListParamsOrderingID),
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

func TestWaapV1DomainFirewallRuleDelete(t *testing.T) {
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
	err := client.Waap.V1.Domains.FirewallRules.Delete(
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

func TestWaapV1DomainFirewallRuleBulkDelete(t *testing.T) {
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
	err := client.Waap.V1.Domains.FirewallRules.BulkDelete(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainFirewallRuleBulkDeleteParams{
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

func TestWaapV1DomainFirewallRuleToggle(t *testing.T) {
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
	err := client.Waap.V1.Domains.FirewallRules.Toggle(
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
