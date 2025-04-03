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

func TestCloudV1FaaNamespaceFunctionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.New(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		gcore.CloudV1FaaNamespaceFunctionNewParams{
			Autoscaling: gcore.F(gcore.FaasAutoscalingParam{
				MaxInstances: gcore.F(int64(5)),
				MinInstances: gcore.F(int64(1)),
			}),
			CodeText:     gcore.F("def main():\n    return 'Hello World'"),
			Flavor:       gcore.F("64m-64MB"),
			MainMethod:   gcore.F("handler"),
			Name:         gcore.F("function-name"),
			Runtime:      gcore.F("python3.7.12"),
			Timeout:      gcore.F(int64(5)),
			Dependencies: gcore.F("numpy==1.21.0\npandas==1.3.0"),
			Description:  gcore.F("This is a sample function."),
			EnableAPIKey: gcore.F(true),
			Envs: gcore.F[any](map[string]interface{}{
				"ENV_VAR": "value",
			}),
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

func TestCloudV1FaaNamespaceFunctionGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.Get(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		"function-name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1FaaNamespaceFunctionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.Update(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		"function-name",
		gcore.CloudV1FaaNamespaceFunctionUpdateParams{
			Autoscaling: gcore.F(gcore.FaasAutoscalingParam{
				MaxInstances: gcore.F(int64(5)),
				MinInstances: gcore.F(int64(1)),
			}),
			CodeText:     gcore.F("def main():\n    return 'Hello World'"),
			Dependencies: gcore.F("numpy==1.21.0\npandas==1.3.0"),
			Description:  gcore.F("This is a sample function."),
			Disabled:     gcore.F(false),
			EnableAPIKey: gcore.F(true),
			Envs: gcore.F[any](map[string]interface{}{
				"ENV_VAR": "value",
			}),
			Flavor:     gcore.F("128m-128MB"),
			Keys:       gcore.F([]string{"key1", "key2"}),
			MainMethod: gcore.F("run"),
			Timeout:    gcore.F(int64(60)),
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

func TestCloudV1FaaNamespaceFunctionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.List(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		gcore.CloudV1FaaNamespaceFunctionListParams{
			Limit:   gcore.F(int64(1000)),
			Offset:  gcore.F(int64(0)),
			OrderBy: gcore.F([]gcore.FunctionsOrderByChoices{gcore.FunctionsOrderByChoicesCreatedAtAsc}),
			Search:  gcore.F("search-value"),
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

func TestCloudV1FaaNamespaceFunctionDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.Delete(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		"function-name",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1FaaNamespaceFunctionGetLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Faas.Namespaces.Functions.GetLogs(
		context.TODO(),
		int64(1),
		int64(1),
		"namespace-name",
		"function-name",
		gcore.CloudV1FaaNamespaceFunctionGetLogsParams{
			Limit: gcore.F(int64(10)),
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
