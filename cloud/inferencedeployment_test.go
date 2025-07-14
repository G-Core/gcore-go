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

func TestInferenceDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.New(context.TODO(), cloud.InferenceDeploymentNewParams{
		ProjectID: gcore.Int(1),
		Containers: []cloud.InferenceDeploymentNewParamsContainer{{
			RegionID: 1,
			Scale: cloud.InferenceDeploymentNewParamsContainerScale{
				Max:             3,
				Min:             1,
				CooldownPeriod:  gcore.Int(60),
				PollingInterval: gcore.Int(30),
				Triggers: cloud.InferenceDeploymentNewParamsContainerScaleTriggers{
					CPU: cloud.InferenceDeploymentNewParamsContainerScaleTriggersCPU{
						Threshold: 80,
					},
					GPUMemory: cloud.InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory{
						Threshold: 80,
					},
					GPUUtilization: cloud.InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization{
						Threshold: 80,
					},
					HTTP: cloud.InferenceDeploymentNewParamsContainerScaleTriggersHTTP{
						Rate:   1,
						Window: 60,
					},
					Memory: cloud.InferenceDeploymentNewParamsContainerScaleTriggersMemory{
						Threshold: 70,
					},
					Sqs: cloud.InferenceDeploymentNewParamsContainerScaleTriggersSqs{
						ActivationQueueLength: 1,
						AwsRegion:             "us-east-1",
						QueueLength:           10,
						QueueURL:              "https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue",
						SecretName:            "x",
						AwsEndpoint:           gcore.String("aws_endpoint"),
						ScaleOnDelayed:        gcore.Bool(true),
						ScaleOnFlight:         gcore.Bool(true),
					},
				},
			},
		}},
		FlavorName:      "inference-16vcpu-232gib-1xh100-80gb",
		Image:           "nginx:latest",
		ListeningPort:   80,
		Name:            "my-instance",
		APIKeys:         []string{"key1", "key2"},
		AuthEnabled:     gcore.Bool(false),
		Command:         []string{"nginx", "-g", "daemon off;"},
		CredentialsName: gcore.String("dockerhub"),
		Description:     gcore.String("My first instance"),
		Envs: map[string]string{
			"DEBUG_MODE": "False",
			"KEY":        "12345",
		},
		IngressOpts: cloud.InferenceDeploymentNewParamsIngressOpts{
			DisableResponseBuffering: gcore.Bool(true),
		},
		Logging: cloud.InferenceDeploymentNewParamsLogging{
			DestinationRegionID: gcore.Int(1),
			Enabled:             gcore.Bool(true),
			RetentionPolicy: cloud.LaasIndexRetentionPolicyParam{
				Period: gcore.Int(42),
			},
			TopicName: gcore.String("my-log-name"),
		},
		Probes: cloud.InferenceDeploymentNewParamsProbes{
			LivenessProbe: cloud.InferenceDeploymentNewParamsProbesLivenessProbe{
				Enabled: true,
				Probe: cloud.InferenceDeploymentNewParamsProbesLivenessProbeProbe{
					Exec: cloud.InferenceDeploymentNewParamsProbesLivenessProbeProbeExec{
						Command: []string{"ls", "-l"},
					},
					FailureThreshold: gcore.Int(3),
					HTTPGet: cloud.InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet{
						Port: 80,
						Headers: map[string]string{
							"Authorization": "Bearer token 123",
						},
						Host:   gcore.String("127.0.0.1"),
						Path:   gcore.String("/healthz"),
						Schema: gcore.String("HTTP"),
					},
					InitialDelaySeconds: gcore.Int(0),
					PeriodSeconds:       gcore.Int(5),
					SuccessThreshold:    gcore.Int(1),
					TcpSocket: cloud.InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket{
						Port: 80,
					},
					TimeoutSeconds: gcore.Int(1),
				},
			},
			ReadinessProbe: cloud.InferenceDeploymentNewParamsProbesReadinessProbe{
				Enabled: true,
				Probe: cloud.InferenceDeploymentNewParamsProbesReadinessProbeProbe{
					Exec: cloud.InferenceDeploymentNewParamsProbesReadinessProbeProbeExec{
						Command: []string{"ls", "-l"},
					},
					FailureThreshold: gcore.Int(3),
					HTTPGet: cloud.InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet{
						Port: 80,
						Headers: map[string]string{
							"Authorization": "Bearer token 123",
						},
						Host:   gcore.String("127.0.0.1"),
						Path:   gcore.String("/healthz"),
						Schema: gcore.String("HTTP"),
					},
					InitialDelaySeconds: gcore.Int(0),
					PeriodSeconds:       gcore.Int(5),
					SuccessThreshold:    gcore.Int(1),
					TcpSocket: cloud.InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket{
						Port: 80,
					},
					TimeoutSeconds: gcore.Int(1),
				},
			},
			StartupProbe: cloud.InferenceDeploymentNewParamsProbesStartupProbe{
				Enabled: true,
				Probe: cloud.InferenceDeploymentNewParamsProbesStartupProbeProbe{
					Exec: cloud.InferenceDeploymentNewParamsProbesStartupProbeProbeExec{
						Command: []string{"ls", "-l"},
					},
					FailureThreshold: gcore.Int(3),
					HTTPGet: cloud.InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet{
						Port: 80,
						Headers: map[string]string{
							"Authorization": "Bearer token 123",
						},
						Host:   gcore.String("127.0.0.1"),
						Path:   gcore.String("/healthz"),
						Schema: gcore.String("HTTP"),
					},
					InitialDelaySeconds: gcore.Int(0),
					PeriodSeconds:       gcore.Int(5),
					SuccessThreshold:    gcore.Int(1),
					TcpSocket: cloud.InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket{
						Port: 80,
					},
					TimeoutSeconds: gcore.Int(1),
				},
			},
		},
		Timeout: gcore.Int(120),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.Update(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentUpdateParams{
			ProjectID:   gcore.Int(1),
			APIKeys:     []string{"key1", "key2"},
			AuthEnabled: gcore.Bool(false),
			Command:     []string{"nginx", "-g", "daemon off;"},
			Containers: []cloud.InferenceDeploymentUpdateParamsContainer{{
				RegionID: gcore.Int(0),
				Scale: cloud.InferenceDeploymentUpdateParamsContainerScale{
					Max:             3,
					Min:             1,
					CooldownPeriod:  gcore.Int(60),
					PollingInterval: gcore.Int(30),
					Triggers: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggers{
						CPU: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersCPU{
							Threshold: 75,
						},
						GPUMemory: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory{
							Threshold: 80,
						},
						GPUUtilization: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization{
							Threshold: 80,
						},
						HTTP: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP{
							Rate:   1,
							Window: 60,
						},
						Memory: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersMemory{
							Threshold: 80,
						},
						Sqs: cloud.InferenceDeploymentUpdateParamsContainerScaleTriggersSqs{
							ActivationQueueLength: 1,
							AwsRegion:             "us-east-1",
							QueueLength:           10,
							QueueURL:              "https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue",
							SecretName:            "x",
							AwsEndpoint:           gcore.String("aws_endpoint"),
							ScaleOnDelayed:        gcore.Bool(true),
							ScaleOnFlight:         gcore.Bool(true),
						},
					},
				},
			}},
			CredentialsName: gcore.String("dockerhub"),
			Description:     gcore.String("My first instance"),
			Envs: map[string]string{
				"DEBUG_MODE": "False",
				"KEY":        "12345",
			},
			FlavorName: gcore.String("inference-16vcpu-232gib-1xh100-80gb"),
			Image:      gcore.String("nginx:latest"),
			IngressOpts: cloud.InferenceDeploymentUpdateParamsIngressOpts{
				DisableResponseBuffering: gcore.Bool(true),
			},
			ListeningPort: gcore.Int(80),
			Logging: cloud.InferenceDeploymentUpdateParamsLogging{
				DestinationRegionID: gcore.Int(1),
				Enabled:             gcore.Bool(true),
				RetentionPolicy: cloud.LaasIndexRetentionPolicyParam{
					Period: gcore.Int(42),
				},
				TopicName: gcore.String("my-log-name"),
			},
			Probes: cloud.InferenceDeploymentUpdateParamsProbes{
				LivenessProbe: cloud.InferenceDeploymentUpdateParamsProbesLivenessProbe{
					Enabled: gcore.Bool(true),
					Probe: cloud.InferenceDeploymentUpdateParamsProbesLivenessProbeProbe{
						Exec: cloud.InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec{
							Command: []string{"ls", "-l"},
						},
						FailureThreshold: gcore.Int(3),
						HTTPGet: cloud.InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet{
							Headers: map[string]string{
								"Authorization": "Bearer token 123",
							},
							Host:   gcore.String("127.0.0.1"),
							Path:   gcore.String("/healthz"),
							Port:   gcore.Int(80),
							Schema: gcore.String("HTTP"),
						},
						InitialDelaySeconds: gcore.Int(0),
						PeriodSeconds:       gcore.Int(5),
						SuccessThreshold:    gcore.Int(1),
						TcpSocket: cloud.InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket{
							Port: gcore.Int(80),
						},
						TimeoutSeconds: gcore.Int(1),
					},
				},
				ReadinessProbe: cloud.InferenceDeploymentUpdateParamsProbesReadinessProbe{
					Enabled: gcore.Bool(true),
					Probe: cloud.InferenceDeploymentUpdateParamsProbesReadinessProbeProbe{
						Exec: cloud.InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec{
							Command: []string{"ls", "-l"},
						},
						FailureThreshold: gcore.Int(3),
						HTTPGet: cloud.InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet{
							Headers: map[string]string{
								"Authorization": "Bearer token 123",
							},
							Host:   gcore.String("127.0.0.1"),
							Path:   gcore.String("/healthz"),
							Port:   gcore.Int(80),
							Schema: gcore.String("HTTP"),
						},
						InitialDelaySeconds: gcore.Int(0),
						PeriodSeconds:       gcore.Int(5),
						SuccessThreshold:    gcore.Int(1),
						TcpSocket: cloud.InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket{
							Port: gcore.Int(80),
						},
						TimeoutSeconds: gcore.Int(1),
					},
				},
				StartupProbe: cloud.InferenceDeploymentUpdateParamsProbesStartupProbe{
					Enabled: gcore.Bool(true),
					Probe: cloud.InferenceDeploymentUpdateParamsProbesStartupProbeProbe{
						Exec: cloud.InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec{
							Command: []string{"ls", "-l"},
						},
						FailureThreshold: gcore.Int(3),
						HTTPGet: cloud.InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet{
							Headers: map[string]string{
								"Authorization": "Bearer token 123",
							},
							Host:   gcore.String("127.0.0.1"),
							Path:   gcore.String("/healthz"),
							Port:   gcore.Int(80),
							Schema: gcore.String("HTTP"),
						},
						InitialDelaySeconds: gcore.Int(0),
						PeriodSeconds:       gcore.Int(5),
						SuccessThreshold:    gcore.Int(1),
						TcpSocket: cloud.InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket{
							Port: gcore.Int(80),
						},
						TimeoutSeconds: gcore.Int(1),
					},
				},
			},
			Timeout: gcore.Int(120),
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

func TestInferenceDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.List(context.TODO(), cloud.InferenceDeploymentListParams{
		ProjectID: gcore.Int(1),
		Limit:     gcore.Int(1000),
		Offset:    gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceDeploymentDelete(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.Delete(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentDeleteParams{
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

func TestInferenceDeploymentGet(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.Get(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentGetParams{
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

func TestInferenceDeploymentGetAPIKey(t *testing.T) {
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
	_, err := client.Cloud.Inference.Deployments.GetAPIKey(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentGetAPIKeyParams{
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

func TestInferenceDeploymentStart(t *testing.T) {
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
	err := client.Cloud.Inference.Deployments.Start(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentStartParams{
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

func TestInferenceDeploymentStop(t *testing.T) {
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
	err := client.Cloud.Inference.Deployments.Stop(
		context.TODO(),
		"my-instance",
		cloud.InferenceDeploymentStopParams{
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
