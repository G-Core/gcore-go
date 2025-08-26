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

func TestDomainCustomRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.CustomRules.New(
		context.TODO(),
		1,
		waap.DomainCustomRuleNewParams{
			Action: waap.DomainCustomRuleNewParamsAction{
				Allow: map[string]interface{}{},
				Block: waap.DomainCustomRuleNewParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
				Captcha:   map[string]interface{}{},
				Handshake: map[string]interface{}{},
				Monitor:   map[string]interface{}{},
				Tag: waap.DomainCustomRuleNewParamsActionTag{
					Tags: []string{"string"},
				},
			},
			Conditions: []waap.DomainCustomRuleNewParamsCondition{{
				ContentType: waap.DomainCustomRuleNewParamsConditionContentType{
					ContentType: []string{"application/xml"},
					Negation:    gcore.Bool(true),
				},
				Country: waap.DomainCustomRuleNewParamsConditionCountry{
					CountryCode: []string{"Mv"},
					Negation:    gcore.Bool(true),
				},
				FileExtension: waap.DomainCustomRuleNewParamsConditionFileExtension{
					FileExtension: []string{"pdf"},
					Negation:      gcore.Bool(true),
				},
				Header: waap.DomainCustomRuleNewParamsConditionHeader{
					Header:    "Origin",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				HeaderExists: waap.DomainCustomRuleNewParamsConditionHeaderExists{
					Header:   "Origin",
					Negation: gcore.Bool(true),
				},
				HTTPMethod: waap.DomainCustomRuleNewParamsConditionHTTPMethod{
					HTTPMethod: "CONNECT",
					Negation:   gcore.Bool(true),
				},
				IP: waap.DomainCustomRuleNewParamsConditionIP{
					IPAddress: "192.168.1.1",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainCustomRuleNewParamsConditionIPRange{
					LowerBound: "192.168.1.1",
					UpperBound: "192.168.1.1",
					Negation:   gcore.Bool(true),
				},
				Organization: waap.DomainCustomRuleNewParamsConditionOrganization{
					Organization: "UptimeRobot s.r.o",
					Negation:     gcore.Bool(true),
				},
				OwnerTypes: waap.DomainCustomRuleNewParamsConditionOwnerTypes{
					Negation:   gcore.Bool(true),
					OwnerTypes: []string{"COMMERCIAL"},
				},
				RequestRate: waap.DomainCustomRuleNewParamsConditionRequestRate{
					PathPattern:    "/",
					Requests:       20,
					Time:           1,
					HTTPMethods:    []string{"CONNECT"},
					IPs:            []string{"192.168.1.1"},
					UserDefinedTag: gcore.String("SQfNklznVLBBpr"),
				},
				ResponseHeader: waap.DomainCustomRuleNewParamsConditionResponseHeader{
					Header:    "header",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				ResponseHeaderExists: waap.DomainCustomRuleNewParamsConditionResponseHeaderExists{
					Header:   "header",
					Negation: gcore.Bool(true),
				},
				SessionRequestCount: waap.DomainCustomRuleNewParamsConditionSessionRequestCount{
					RequestCount: 1,
					Negation:     gcore.Bool(true),
				},
				Tags: waap.DomainCustomRuleNewParamsConditionTags{
					Tags:     []string{"string"},
					Negation: gcore.Bool(true),
				},
				URL: waap.DomainCustomRuleNewParamsConditionURL{
					URL:       "/wp-admin/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserAgent: waap.DomainCustomRuleNewParamsConditionUserAgent{
					UserAgent: "curl/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserDefinedTags: waap.DomainCustomRuleNewParamsConditionUserDefinedTags{
					Tags:     []string{"SQfNklznVLBBpr"},
					Negation: gcore.Bool(true),
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

func TestDomainCustomRuleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.Domains.CustomRules.Update(
		context.TODO(),
		0,
		waap.DomainCustomRuleUpdateParams{
			DomainID: 1,
			Action: waap.DomainCustomRuleUpdateParamsAction{
				Allow: map[string]interface{}{},
				Block: waap.DomainCustomRuleUpdateParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
				Captcha:   map[string]interface{}{},
				Handshake: map[string]interface{}{},
				Monitor:   map[string]interface{}{},
				Tag: waap.DomainCustomRuleUpdateParamsActionTag{
					Tags: []string{"string"},
				},
			},
			Conditions: []waap.DomainCustomRuleUpdateParamsCondition{{
				ContentType: waap.DomainCustomRuleUpdateParamsConditionContentType{
					ContentType: []string{"application/xml"},
					Negation:    gcore.Bool(true),
				},
				Country: waap.DomainCustomRuleUpdateParamsConditionCountry{
					CountryCode: []string{"Mv"},
					Negation:    gcore.Bool(true),
				},
				FileExtension: waap.DomainCustomRuleUpdateParamsConditionFileExtension{
					FileExtension: []string{"pdf"},
					Negation:      gcore.Bool(true),
				},
				Header: waap.DomainCustomRuleUpdateParamsConditionHeader{
					Header:    "Origin",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				HeaderExists: waap.DomainCustomRuleUpdateParamsConditionHeaderExists{
					Header:   "Origin",
					Negation: gcore.Bool(true),
				},
				HTTPMethod: waap.DomainCustomRuleUpdateParamsConditionHTTPMethod{
					HTTPMethod: "CONNECT",
					Negation:   gcore.Bool(true),
				},
				IP: waap.DomainCustomRuleUpdateParamsConditionIP{
					IPAddress: "192.168.1.1",
					Negation:  gcore.Bool(true),
				},
				IPRange: waap.DomainCustomRuleUpdateParamsConditionIPRange{
					LowerBound: "192.168.1.1",
					UpperBound: "192.168.1.1",
					Negation:   gcore.Bool(true),
				},
				Organization: waap.DomainCustomRuleUpdateParamsConditionOrganization{
					Organization: "UptimeRobot s.r.o",
					Negation:     gcore.Bool(true),
				},
				OwnerTypes: waap.DomainCustomRuleUpdateParamsConditionOwnerTypes{
					Negation:   gcore.Bool(true),
					OwnerTypes: []string{"COMMERCIAL"},
				},
				RequestRate: waap.DomainCustomRuleUpdateParamsConditionRequestRate{
					PathPattern:    "/",
					Requests:       20,
					Time:           1,
					HTTPMethods:    []string{"CONNECT"},
					IPs:            []string{"192.168.1.1"},
					UserDefinedTag: gcore.String("SQfNklznVLBBpr"),
				},
				ResponseHeader: waap.DomainCustomRuleUpdateParamsConditionResponseHeader{
					Header:    "header",
					Value:     "value",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				ResponseHeaderExists: waap.DomainCustomRuleUpdateParamsConditionResponseHeaderExists{
					Header:   "header",
					Negation: gcore.Bool(true),
				},
				SessionRequestCount: waap.DomainCustomRuleUpdateParamsConditionSessionRequestCount{
					RequestCount: 1,
					Negation:     gcore.Bool(true),
				},
				Tags: waap.DomainCustomRuleUpdateParamsConditionTags{
					Tags:     []string{"string"},
					Negation: gcore.Bool(true),
				},
				URL: waap.DomainCustomRuleUpdateParamsConditionURL{
					URL:       "/wp-admin/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserAgent: waap.DomainCustomRuleUpdateParamsConditionUserAgent{
					UserAgent: "curl/",
					MatchType: "Exact",
					Negation:  gcore.Bool(true),
				},
				UserDefinedTags: waap.DomainCustomRuleUpdateParamsConditionUserDefinedTags{
					Tags:     []string{"SQfNklznVLBBpr"},
					Negation: gcore.Bool(true),
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

func TestDomainCustomRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.CustomRules.List(
		context.TODO(),
		1,
		waap.DomainCustomRuleListParams{
			Action:      waap.DomainCustomRuleListParamsActionBlock,
			Description: gcore.String("This rule blocks all the requests coming form a specific IP address."),
			Enabled:     gcore.Bool(false),
			Limit:       gcore.Int(0),
			Name:        gcore.String("Block by specific IP rule."),
			Offset:      gcore.Int(0),
			Ordering:    waap.DomainCustomRuleListParamsOrderingMinusID,
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

func TestDomainCustomRuleDelete(t *testing.T) {
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
	err := client.Waap.Domains.CustomRules.Delete(
		context.TODO(),
		0,
		waap.DomainCustomRuleDeleteParams{
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

func TestDomainCustomRuleDeleteMultiple(t *testing.T) {
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
	err := client.Waap.Domains.CustomRules.DeleteMultiple(
		context.TODO(),
		1,
		waap.DomainCustomRuleDeleteMultipleParams{
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

func TestDomainCustomRuleGet(t *testing.T) {
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
	_, err := client.Waap.Domains.CustomRules.Get(
		context.TODO(),
		0,
		waap.DomainCustomRuleGetParams{
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

func TestDomainCustomRuleToggle(t *testing.T) {
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
	err := client.Waap.Domains.CustomRules.Toggle(
		context.TODO(),
		waap.DomainCustomRuleToggleParamsActionEnable,
		waap.DomainCustomRuleToggleParams{
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
