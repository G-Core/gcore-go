// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestLoadBalancerL7PolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.New(context.TODO(), cloud.LoadBalancerL7PolicyNewParams{
		ProjectID:        gcore.Int(0),
		RegionID:         gcore.Int(0),
		Action:           cloud.LoadBalancerL7PolicyNewParamsActionRedirectToURL,
		ListenerID:       "023f2e34-7806-443b-bfae-16c324569a3d",
		Name:             gcore.String("redirect-example.com"),
		Position:         gcore.Int(1),
		RedirectHTTPCode: gcore.Int(301),
		RedirectPoolID:   gcore.String("redirect_pool_id"),
		RedirectPrefix:   gcore.String("redirect_prefix"),
		RedirectURL:      gcore.String("http://www.example.com"),
		Tags:             []string{"test_tag"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerL7PolicyList(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.List(context.TODO(), cloud.LoadBalancerL7PolicyListParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerL7PolicyDelete(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Delete(
		context.TODO(),
		"l7policy_id",
		cloud.LoadBalancerL7PolicyDeleteParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestLoadBalancerL7PolicyGet(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Get(
		context.TODO(),
		"l7policy_id",
		cloud.LoadBalancerL7PolicyGetParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestLoadBalancerL7PolicyReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.LoadBalancers.L7Policies.Replace(
		context.TODO(),
		"l7policy_id",
		cloud.LoadBalancerL7PolicyReplaceParams{
			ProjectID:        gcore.Int(0),
			RegionID:         gcore.Int(0),
			Action:           cloud.LoadBalancerL7PolicyReplaceParamsActionRedirectToURL,
			Name:             gcore.String("redirect-images.example.com"),
			Position:         gcore.Int(1),
			RedirectHTTPCode: gcore.Int(301),
			RedirectPoolID:   gcore.String("redirect_pool_id"),
			RedirectPrefix:   gcore.String("redirect_prefix"),
			RedirectURL:      gcore.String("http://images.example.com"),
			Tags:             []string{"updated_tag"},
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
