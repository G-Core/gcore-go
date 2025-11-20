// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestLoadBalancerL7PolicyRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Rules.New(
		context.TODO(),
		"023f2e34-7806-443b-bfae-16c324569a3d",
		cloud.LoadBalancerL7PolicyRuleNewParams{
			ProjectID:   gcore.Int(1),
			RegionID:    gcore.Int(1),
			CompareType: cloud.LoadBalancerL7PolicyRuleNewParamsCompareTypeRegex,
			Type:        cloud.LoadBalancerL7PolicyRuleNewParamsTypePath,
			Value:       "/images*",
			Invert:      gcore.Bool(true),
			Key:         gcore.String("the name of the cookie to evaluate."),
			Tags:        []string{"test_tag_1", "test_tag_2"},
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

func TestLoadBalancerL7PolicyRuleList(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Rules.List(
		context.TODO(),
		"023f2e34-7806-443b-bfae-16c324569a3d",
		cloud.LoadBalancerL7PolicyRuleListParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestLoadBalancerL7PolicyRuleDelete(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Rules.Delete(
		context.TODO(),
		"023f2e34-7806-443b-bfae-16c324569a3d",
		cloud.LoadBalancerL7PolicyRuleDeleteParams{
			ProjectID:  gcore.Int(1),
			RegionID:   gcore.Int(1),
			L7policyID: "023f2e34-7806-443b-bfae-16c324569a3d",
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

func TestLoadBalancerL7PolicyRuleGet(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Rules.Get(
		context.TODO(),
		"023f2e34-7806-443b-bfae-16c324569a3d",
		cloud.LoadBalancerL7PolicyRuleGetParams{
			ProjectID:  gcore.Int(1),
			RegionID:   gcore.Int(1),
			L7policyID: "023f2e34-7806-443b-bfae-16c324569a3d",
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

func TestLoadBalancerL7PolicyRuleReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Rules.Replace(
		context.TODO(),
		"023f2e34-7806-443b-bfae-16c324569a3d",
		cloud.LoadBalancerL7PolicyRuleReplaceParams{
			ProjectID:   gcore.Int(1),
			RegionID:    gcore.Int(1),
			L7policyID:  "023f2e34-7806-443b-bfae-16c324569a3d",
			CompareType: cloud.LoadBalancerL7PolicyRuleReplaceParamsCompareTypeRegex,
			Invert:      gcore.Bool(true),
			Key:         gcore.String("the name of the cookie to evaluate."),
			Tags:        []string{"test_tag_1", "test_tag_2"},
			Type:        cloud.LoadBalancerL7PolicyRuleReplaceParamsTypePath,
			Value:       gcore.String("/images*"),
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
