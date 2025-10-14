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

func TestInferenceApplicationDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Inference.Applications.Deployments.New(context.TODO(), cloud.InferenceApplicationDeploymentNewParams{
		ProjectID:       gcore.Int(1),
		ApplicationName: "demo-app",
		ComponentsConfiguration: map[string]cloud.InferenceApplicationDeploymentNewParamsComponentsConfiguration{
			"model": {
				Exposed: true,
				Flavor:  "inference-16vcpu-232gib-1xh100-80gb",
				Scale: cloud.InferenceApplicationDeploymentNewParamsComponentsConfigurationScale{
					Max: 1,
					Min: 1,
				},
				ParameterOverrides: map[string]cloud.InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride{
					"foo": {
						Value: "value",
					},
				},
			},
		},
		Name:    "name",
		Regions: []int64{1, 2},
		APIKeys: []string{"key1", "key2"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceApplicationDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Inference.Applications.Deployments.Update(
		context.TODO(),
		"deployment_name",
		cloud.InferenceApplicationDeploymentUpdateParams{
			ProjectID: gcore.Int(1),
			APIKeys:   []string{"key1", "key2"},
			ComponentsConfiguration: map[string]cloud.InferenceApplicationDeploymentUpdateParamsComponentsConfiguration{
				"model": {
					Exposed: gcore.Bool(true),
					Flavor:  gcore.String("flavor"),
					ParameterOverrides: map[string]cloud.InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride{
						"foo": {
							Value: "value",
						},
					},
					Scale: cloud.InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale{
						Max: gcore.Int(2),
						Min: gcore.Int(0),
					},
				},
			},
			Regions: []int64{1, 2},
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

func TestInferenceApplicationDeploymentList(t *testing.T) {
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
	_, err := client.Cloud.Inference.Applications.Deployments.List(context.TODO(), cloud.InferenceApplicationDeploymentListParams{
		ProjectID: gcore.Int(1),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceApplicationDeploymentDelete(t *testing.T) {
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
	_, err := client.Cloud.Inference.Applications.Deployments.Delete(
		context.TODO(),
		"deployment_name",
		cloud.InferenceApplicationDeploymentDeleteParams{
			ProjectID: gcore.Int(1),
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

func TestInferenceApplicationDeploymentGet(t *testing.T) {
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
	_, err := client.Cloud.Inference.Applications.Deployments.Get(
		context.TODO(),
		"deployment_name",
		cloud.InferenceApplicationDeploymentGetParams{
			ProjectID: gcore.Int(1),
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
