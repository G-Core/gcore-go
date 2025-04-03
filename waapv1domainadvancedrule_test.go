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

func TestWaapV1DomainAdvancedRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.AdvancedRules.New(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainAdvancedRuleNewParams{
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
			Enabled:     gcore.F(true),
			Name:        gcore.F("name"),
			Source:      gcore.F("x"),
			Description: gcore.F("description"),
			Phase:       gcore.F(gcore.WaapV1DomainAdvancedRuleNewParamsPhaseAccess),
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

func TestWaapV1DomainAdvancedRuleGet(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.AdvancedRules.Get(
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

func TestWaapV1DomainAdvancedRuleUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.V1.Domains.AdvancedRules.Update(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.WaapV1DomainAdvancedRuleUpdateParams{
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
			Description: gcore.F("description"),
			Enabled:     gcore.F(true),
			Name:        gcore.F("name"),
			Phase:       gcore.F(gcore.WaapV1DomainAdvancedRuleUpdateParamsPhaseAccess),
			Source:      gcore.F("x"),
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

func TestWaapV1DomainAdvancedRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.AdvancedRules.List(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainAdvancedRuleListParams{
			Action:      gcore.F(gcore.RuleActionTypeAllow),
			Description: gcore.F("description"),
			Enabled:     gcore.F(true),
			Limit:       gcore.F(int64(0)),
			Name:        gcore.F("name"),
			Offset:      gcore.F(int64(0)),
			Ordering:    gcore.F(gcore.WaapV1DomainAdvancedRuleListParamsOrderingID),
			Phase:       gcore.F(gcore.WaapV1DomainAdvancedRuleListParamsPhaseAccess),
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

func TestWaapV1DomainAdvancedRuleDelete(t *testing.T) {
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
	err := client.Waap.V1.Domains.AdvancedRules.Delete(
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

func TestWaapV1DomainAdvancedRuleToggle(t *testing.T) {
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
	err := client.Waap.V1.Domains.AdvancedRules.Toggle(
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
