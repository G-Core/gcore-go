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

func TestDomainFirewallRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.FirewallRules.New(
		context.TODO(),
		1,
		waap.DomainFirewallRuleNewParams{
			Action: waap.DomainFirewallRuleNewParamsAction{
				Allow: map[string]any{
					"foo": "bar",
				},
				Block: waap.DomainFirewallRuleNewParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
			},
			Conditions: []waap.DomainFirewallRuleNewParamsCondition{{
				IP: waap.DomainFirewallRuleNewParamsConditionIP{
					IPAddress: "ip_address",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainFirewallRuleNewParamsConditionIPRange{
					LowerBound: "lower_bound",
					UpperBound: "upper_bound",
					Negation:   gcore.Bool(true),
				},
			}},
			Enabled:     true,
			Name:        "Block foobar bot",
			Description: gcore.String("description"),
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

func TestDomainFirewallRuleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.Domains.FirewallRules.Update(
		context.TODO(),
		0,
		waap.DomainFirewallRuleUpdateParams{
			DomainID: 1,
			Action: waap.DomainFirewallRuleUpdateParamsAction{
				Allow: map[string]any{
					"foo": "bar",
				},
				Block: waap.DomainFirewallRuleUpdateParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
			},
			Conditions: []waap.DomainFirewallRuleUpdateParamsCondition{{
				IP: waap.DomainFirewallRuleUpdateParamsConditionIP{
					IPAddress: "ip_address",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainFirewallRuleUpdateParamsConditionIPRange{
					LowerBound: "lower_bound",
					UpperBound: "upper_bound",
					Negation:   gcore.Bool(true),
				},
			}},
			Description: gcore.String("description"),
			Enabled:     gcore.Bool(true),
			Name:        gcore.String("Block foobar bot"),
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

func TestDomainFirewallRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.FirewallRules.List(
		context.TODO(),
		1,
		waap.DomainFirewallRuleListParams{
			Action:      waap.DomainFirewallRuleListParamsActionAllow,
			Description: gcore.String("This rule blocks all the requests coming form a specific IP address."),
			Enabled:     gcore.Bool(false),
			Limit:       gcore.Int(0),
			Name:        gcore.String("Block by specific IP rule."),
			Offset:      gcore.Int(0),
			Ordering:    waap.DomainFirewallRuleListParamsOrderingMinusID,
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

func TestDomainFirewallRuleDelete(t *testing.T) {
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
	err := client.Waap.Domains.FirewallRules.Delete(
		context.TODO(),
		0,
		waap.DomainFirewallRuleDeleteParams{
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

func TestDomainFirewallRuleDeleteMultiple(t *testing.T) {
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
	err := client.Waap.Domains.FirewallRules.DeleteMultiple(
		context.TODO(),
		1,
		waap.DomainFirewallRuleDeleteMultipleParams{
			RuleIDs: []int64{0},
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

func TestDomainFirewallRuleGet(t *testing.T) {
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
	_, err := client.Waap.Domains.FirewallRules.Get(
		context.TODO(),
		0,
		waap.DomainFirewallRuleGetParams{
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

func TestDomainFirewallRuleToggle(t *testing.T) {
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
	err := client.Waap.Domains.FirewallRules.Toggle(
		context.TODO(),
		waap.DomainFirewallRuleToggleParamsActionEnable,
		waap.DomainFirewallRuleToggleParams{
			DomainID: 1,
			RuleID:   0,
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
