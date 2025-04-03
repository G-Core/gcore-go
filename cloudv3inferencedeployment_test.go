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

func TestCloudV3InferenceDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.New(
		context.TODO(),
		int64(1),
		gcore.CloudV3InferenceDeploymentNewParams{
			Containers: gcore.F([]gcore.CloudV3InferenceDeploymentNewParamsContainer{{
				RegionID: gcore.F(int64(1)),
				Scale: gcore.F(gcore.ContainerScaleDeploymentV3Param{
					Max:             gcore.F(int64(3)),
					Min:             gcore.F(int64(1)),
					CooldownPeriod:  gcore.F(int64(60)),
					PollingInterval: gcore.F(int64(30)),
					Triggers: gcore.F(gcore.ContainerScaleDeploymentV3TriggersParam{
						CPU: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(80)),
						}),
						GPUMemory: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(75)),
						}),
						GPUUtilization: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(75)),
						}),
						HTTP: gcore.F(gcore.ContainerScaleDeploymentV3TriggersHTTPParam{
							Rate:   gcore.F(int64(1)),
							Window: gcore.F(int64(60)),
						}),
						Memory: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(70)),
						}),
						Sqs: gcore.F(gcore.ContainerScaleDeploymentV3TriggersSqsParam{
							ActivationQueueLength: gcore.F(int64(1)),
							AwsRegion:             gcore.F("us-east-1"),
							QueueLength:           gcore.F(int64(10)),
							QueueURL:              gcore.F("https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue"),
							SecretName:            gcore.F("x"),
							AwsEndpoint:           gcore.F("aws_endpoint"),
							ScaleOnDelayed:        gcore.F(true),
							ScaleOnFlight:         gcore.F(true),
						}),
					}),
				}),
			}}),
			FlavorName:      gcore.F("inference-16vcpu-232gib-1xh100-80gb"),
			Image:           gcore.F("nginx:latest"),
			ListeningPort:   gcore.F(int64(80)),
			Name:            gcore.F("my-instance"),
			AuthEnabled:     gcore.F(false),
			Command:         gcore.F([]string{"nginx", "-g", "daemon off;"}),
			CredentialsName: gcore.F("dockerhub"),
			Description:     gcore.F("My first instance"),
			Envs: gcore.F(map[string]string{
				"DEBUG_MODE": "False",
				"KEY":        "12345",
			}),
			IngressOpts: gcore.F(gcore.IngressOptsParam{
				DisableResponseBuffering: gcore.F(true),
			}),
			Logging: gcore.F(gcore.LoggingInParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			Probes: gcore.F(gcore.InstanceProbesDeploymentV2Param{
				LivenessProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				ReadinessProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				StartupProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
			}),
			Timeout: gcore.F(int64(120)),
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

func TestCloudV3InferenceDeploymentGet(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.Get(
		context.TODO(),
		int64(1),
		"my-instance",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3InferenceDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.Update(
		context.TODO(),
		int64(1),
		"my-instance",
		gcore.CloudV3InferenceDeploymentUpdateParams{
			AuthEnabled: gcore.F(false),
			Command:     gcore.F([]string{"nginx", "-g", "daemon off;"}),
			Containers: gcore.F([]gcore.CloudV3InferenceDeploymentUpdateParamsContainer{{
				RegionID: gcore.F(int64(0)),
				Scale: gcore.F(gcore.ContainerScaleDeploymentV3Param{
					Max:             gcore.F(int64(3)),
					Min:             gcore.F(int64(1)),
					CooldownPeriod:  gcore.F(int64(60)),
					PollingInterval: gcore.F(int64(30)),
					Triggers: gcore.F(gcore.ContainerScaleDeploymentV3TriggersParam{
						CPU: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(80)),
						}),
						GPUMemory: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(75)),
						}),
						GPUUtilization: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(75)),
						}),
						HTTP: gcore.F(gcore.ContainerScaleDeploymentV3TriggersHTTPParam{
							Rate:   gcore.F(int64(1)),
							Window: gcore.F(int64(60)),
						}),
						Memory: gcore.F(gcore.ContainerScaleTriggersThresholdParam{
							Threshold: gcore.F(int64(70)),
						}),
						Sqs: gcore.F(gcore.ContainerScaleDeploymentV3TriggersSqsParam{
							ActivationQueueLength: gcore.F(int64(1)),
							AwsRegion:             gcore.F("us-east-1"),
							QueueLength:           gcore.F(int64(10)),
							QueueURL:              gcore.F("https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue"),
							SecretName:            gcore.F("x"),
							AwsEndpoint:           gcore.F("aws_endpoint"),
							ScaleOnDelayed:        gcore.F(true),
							ScaleOnFlight:         gcore.F(true),
						}),
					}),
				}),
			}}),
			CredentialsName: gcore.F("dockerhub"),
			Description:     gcore.F("My first instance"),
			Envs: gcore.F(map[string]string{
				"DEBUG_MODE": "False",
				"KEY":        "12345",
			}),
			FlavorName: gcore.F("inference-16vcpu-232gib-1xh100-80gb"),
			Image:      gcore.F("nginx:latest"),
			IngressOpts: gcore.F(gcore.IngressOptsParam{
				DisableResponseBuffering: gcore.F(true),
			}),
			ListeningPort: gcore.F(int64(80)),
			Logging: gcore.F(gcore.LoggingInParam{
				DestinationRegionID: gcore.F(int64(1)),
				Enabled:             gcore.F(true),
				RetentionPolicy: gcore.F(gcore.LaasIndexRetentionPolicyPydanticParam{
					Period: gcore.F(int64(45)),
				}),
				TopicName: gcore.F("some-topic"),
			}),
			Probes: gcore.F(gcore.InstanceProbesDeploymentV2Param{
				LivenessProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				ReadinessProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				StartupProbe: gcore.F(gcore.InstanceContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.InstanceContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.InstanceContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.InstanceContainerProbeConfigurationProbeHTTPGetParam{
							Port: gcore.F(int64(80)),
							Headers: gcore.F(map[string]string{
								"Authorization": "Bearer token 123",
							}),
							Host:   gcore.F("127.0.0.1"),
							Path:   gcore.F("/healthz"),
							Schema: gcore.F("HTTP"),
						}),
						InitialDelaySeconds: gcore.F(int64(0)),
						PeriodSeconds:       gcore.F(int64(5)),
						SuccessThreshold:    gcore.F(int64(1)),
						TcpSocket: gcore.F(gcore.InstanceContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
			}),
			Timeout: gcore.F(int64(120)),
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

func TestCloudV3InferenceDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.List(
		context.TODO(),
		int64(1),
		gcore.CloudV3InferenceDeploymentListParams{
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

func TestCloudV3InferenceDeploymentDelete(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.Delete(
		context.TODO(),
		int64(1),
		"my-instance",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3InferenceDeploymentGetApikey(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.GetApikey(
		context.TODO(),
		int64(1),
		"my-instance",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3InferenceDeploymentGetLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V3.Inference.Deployments.GetLogs(
		context.TODO(),
		int64(1),
		"my-instance",
		gcore.CloudV3InferenceDeploymentGetLogsParams{
			Limit:    gcore.F(int64(1000)),
			Offset:   gcore.F(int64(0)),
			OrderBy:  gcore.F(gcore.CloudV3InferenceDeploymentGetLogsParamsOrderByTimeAsc),
			RegionID: gcore.F(int64(1)),
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

func TestCloudV3InferenceDeploymentStart(t *testing.T) {
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
	err := client.Cloud.V3.Inference.Deployments.Start(
		context.TODO(),
		int64(1),
		"my-instance",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV3InferenceDeploymentStop(t *testing.T) {
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
	err := client.Cloud.V3.Inference.Deployments.Stop(
		context.TODO(),
		int64(1),
		"my-instance",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
