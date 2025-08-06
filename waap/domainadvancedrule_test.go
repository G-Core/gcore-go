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

func TestDomainAdvancedRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.AdvancedRules.New(
		context.TODO(),
		1,
		waap.DomainAdvancedRuleNewParams{
			Action: waap.DomainAdvancedRuleNewParamsAction{
				Allow: map[string]interface{}{},
				Block: waap.DomainAdvancedRuleNewParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
				Captcha:   map[string]interface{}{},
				Handshake: map[string]interface{}{},
				Monitor:   map[string]interface{}{},
				Tag: waap.DomainAdvancedRuleNewParamsActionTag{
					Tags: []string{"string"},
				},
			},
			Enabled:     true,
			Name:        "name",
			Source:      "request.rate_limit([], '.*events', 5, 200, [], [], '', 'ip') and not ('mb-web-ui' in request.headers['Cookie'] or 'mb-mobile-ios' in request.headers['Cookie'] or 'session-token' in request.headers['Cookie']) and not request.headers['session']",
			Description: gcore.String("description"),
			Phase:       waap.DomainAdvancedRuleNewParamsPhaseAccess,
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

func TestDomainAdvancedRuleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.Domains.AdvancedRules.Update(
		context.TODO(),
		0,
		waap.DomainAdvancedRuleUpdateParams{
			DomainID: 1,
			Action: waap.DomainAdvancedRuleUpdateParamsAction{
				Allow: map[string]interface{}{},
				Block: waap.DomainAdvancedRuleUpdateParamsActionBlock{
					ActionDuration: gcore.String("12h"),
					StatusCode:     403,
				},
				Captcha:   map[string]interface{}{},
				Handshake: map[string]interface{}{},
				Monitor:   map[string]interface{}{},
				Tag: waap.DomainAdvancedRuleUpdateParamsActionTag{
					Tags: []string{"string"},
				},
			},
			Description: gcore.String("description"),
			Enabled:     gcore.Bool(true),
			Name:        gcore.String("name"),
			Phase:       waap.DomainAdvancedRuleUpdateParamsPhaseAccess,
			Source:      gcore.String("x"),
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

func TestDomainAdvancedRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.AdvancedRules.List(
		context.TODO(),
		1,
		waap.DomainAdvancedRuleListParams{
			Action:      waap.DomainAdvancedRuleListParamsActionBlock,
			Description: gcore.String("This rule blocks all the requests coming form a specific IP address"),
			Enabled:     gcore.Bool(false),
			Limit:       gcore.Int(0),
			Name:        gcore.String("Block by specific IP rule"),
			Offset:      gcore.Int(0),
			Ordering:    waap.DomainAdvancedRuleListParamsOrderingID,
			Phase:       waap.DomainAdvancedRuleListParamsPhaseAccess,
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

func TestDomainAdvancedRuleDelete(t *testing.T) {
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
	err := client.Waap.Domains.AdvancedRules.Delete(
		context.TODO(),
		0,
		waap.DomainAdvancedRuleDeleteParams{
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

func TestDomainAdvancedRuleGet(t *testing.T) {
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
	_, err := client.Waap.Domains.AdvancedRules.Get(
		context.TODO(),
		0,
		waap.DomainAdvancedRuleGetParams{
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

func TestDomainAdvancedRuleToggle(t *testing.T) {
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
	err := client.Waap.Domains.AdvancedRules.Toggle(
		context.TODO(),
		waap.DomainAdvancedRuleToggleParamsActionEnable,
		waap.DomainAdvancedRuleToggleParams{
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
