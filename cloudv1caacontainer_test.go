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

func TestCloudV1CaaContainerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.New(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1CaaContainerNewParams{
			Flavor:        gcore.F("250mCPU-512MiB"),
			Image:         gcore.F("nginx:latest"),
			ListeningPort: gcore.F(int64(80)),
			Name:          gcore.F("nginx"),
			Scale: gcore.F(gcore.ScaleParam{
				Max:            gcore.F(int64(1)),
				Min:            gcore.F(int64(0)),
				CooldownPeriod: gcore.F(int64(1)),
				Triggers: gcore.F(gcore.ScaleTriggersParam{
					CPU: gcore.F(gcore.ScaleTriggersCPUParam{
						Threshold: gcore.F(int64(1)),
					}),
					HTTP: gcore.F(gcore.ScaleTriggersHTTPParam{
						Rate:   gcore.F(int64(1)),
						Window: gcore.F(int64(1)),
					}),
					Memory: gcore.F(gcore.ScaleTriggersMemoryParam{
						Threshold: gcore.F(int64(1)),
					}),
				}),
			}),
			Commands:    gcore.F("python3 app.py -m test"),
			Description: gcore.F("My first container"),
			Envs: gcore.F[any](map[string]interface{}{
				"ENVIRONMENT_VARIABLE": "value 2",
				"ENV_VAR":              "value 1",
			}),
			IsAPIKeyAuth: gcore.F(false),
			IsDisabled:   gcore.F(false),
			Logging: gcore.F(gcore.LoggingParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			PullSecret: gcore.F("my-secret"),
			Timeout:    gcore.F(int64(5)),
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

func TestCloudV1CaaContainerGet(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.Get(
		context.TODO(),
		int64(1),
		int64(1),
		"my-container",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1CaaContainerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.Update(
		context.TODO(),
		int64(1),
		int64(1),
		"my-container",
		gcore.CloudV1CaaContainerUpdateParams{
			Commands:    gcore.F("python3 app.py -m test"),
			Description: gcore.F("My first container"),
			Envs: gcore.F[any](map[string]interface{}{
				"ENVIRONMENT_VARIABLE": "value 2",
				"ENV_VAR":              "value 1",
			}),
			Flavor:        gcore.F("250mCPU-512MiB"),
			Image:         gcore.F("nginx:latest"),
			IsAPIKeyAuth:  gcore.F(false),
			IsDisabled:    gcore.F(false),
			ListeningPort: gcore.F(int64(80)),
			Logging: gcore.F(gcore.LoggingParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			PullSecret: gcore.F("my-secret"),
			Scale: gcore.F(gcore.ScaleParam{
				Max:            gcore.F(int64(1)),
				Min:            gcore.F(int64(0)),
				CooldownPeriod: gcore.F(int64(1)),
				Triggers: gcore.F(gcore.ScaleTriggersParam{
					CPU: gcore.F(gcore.ScaleTriggersCPUParam{
						Threshold: gcore.F(int64(1)),
					}),
					HTTP: gcore.F(gcore.ScaleTriggersHTTPParam{
						Rate:   gcore.F(int64(1)),
						Window: gcore.F(int64(1)),
					}),
					Memory: gcore.F(gcore.ScaleTriggersMemoryParam{
						Threshold: gcore.F(int64(1)),
					}),
				}),
			}),
			Timeout: gcore.F(int64(5)),
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

func TestCloudV1CaaContainerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.List(
		context.TODO(),
		int64(1),
		int64(1),
		gcore.CloudV1CaaContainerListParams{
			Limit:  gcore.F(int64(1000)),
			Offset: gcore.F(int64(0)),
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

func TestCloudV1CaaContainerDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.Delete(
		context.TODO(),
		int64(1),
		int64(1),
		"my-container",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1CaaContainerGetLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Caas.Containers.GetLogs(
		context.TODO(),
		int64(1),
		int64(1),
		"my-container",
		gcore.CloudV1CaaContainerGetLogsParams{
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
