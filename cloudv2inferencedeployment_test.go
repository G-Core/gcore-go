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

func TestCloudV2InferenceDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Inference.Deployments.New(context.TODO(), gcore.CloudV2InferenceDeploymentNewParams{
		InstanceIn: gcore.InstanceInParam{
			Containers: gcore.F([]gcore.InstanceInContainerParam{{
				RegionID: gcore.F(int64(7)),
				Scale: gcore.F(gcore.ContainerScaleDeploymentParam{
					Max:            gcore.F(int64(3)),
					Min:            gcore.F(int64(1)),
					CooldownPeriod: gcore.F(int64(60)),
					Triggers: gcore.F(gcore.ContainerScaleDeploymentTriggersParam{
						CPU: gcore.F(gcore.ThresholdParam{
							Threshold: gcore.F(int64(80)),
						}),
						GPUMemory: gcore.F(gcore.ThresholdParam{
							Threshold: gcore.F(int64(1)),
						}),
						GPUUtilization: gcore.F(gcore.ThresholdParam{
							Threshold: gcore.F(int64(1)),
						}),
						HTTP: gcore.F(gcore.ContainerScaleDeploymentTriggersHTTPParam{
							Rate:   gcore.F(int64(1)),
							Window: gcore.F(int64(1)),
						}),
						Memory: gcore.F(gcore.ThresholdParam{
							Threshold: gcore.F(int64(70)),
						}),
					}),
				}),
			}}),
			FlavorID:      gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			Image:         gcore.F("nginx:latest"),
			ListeningPort: gcore.F(int64(8080)),
			Name:          gcore.F("my-instance"),
			APIKeyIDs:     gcore.F([]string{"3fa85f64-5717-4562-b3fc-2c963f66afa6"}),
			AuthEnabled:   gcore.F(false),
			Command:       gcore.F([]string{"nginx", "-g", "daemon off;"}),
			Description:   gcore.F("My first instance"),
			Envs: gcore.F(map[string]string{
				"DEBUG_MODE": "False",
				"KEY":        "12345",
			}),
			ImageRegistryID: gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			Probes: gcore.F(gcore.InstanceProbesDeploymentParam{
				LivenessProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
						TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				ReadinessProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
						TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
				StartupProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
					Enabled: gcore.F(true),
					Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
						Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
							Command: gcore.F([]string{"ls", "-l"}),
						}),
						FailureThreshold: gcore.F(int64(3)),
						HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
						TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
							Port: gcore.F(int64(80)),
						}),
						TimeoutSeconds: gcore.F(int64(1)),
					}),
				}),
			}),
			ProjectID: gcore.F(int64(1)),
			Timeout:   gcore.F(int64(120)),
		},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentGet(t *testing.T) {
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
	_, err := client.Cloud.V2.Inference.Deployments.Get(context.TODO(), "instance_id")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Inference.Deployments.Update(
		context.TODO(),
		"instance_id",
		gcore.CloudV2InferenceDeploymentUpdateParams{
			InstanceIn: gcore.InstanceInParam{
				Containers: gcore.F([]gcore.InstanceInContainerParam{{
					RegionID: gcore.F(int64(7)),
					Scale: gcore.F(gcore.ContainerScaleDeploymentParam{
						Max:            gcore.F(int64(3)),
						Min:            gcore.F(int64(1)),
						CooldownPeriod: gcore.F(int64(60)),
						Triggers: gcore.F(gcore.ContainerScaleDeploymentTriggersParam{
							CPU: gcore.F(gcore.ThresholdParam{
								Threshold: gcore.F(int64(80)),
							}),
							GPUMemory: gcore.F(gcore.ThresholdParam{
								Threshold: gcore.F(int64(1)),
							}),
							GPUUtilization: gcore.F(gcore.ThresholdParam{
								Threshold: gcore.F(int64(1)),
							}),
							HTTP: gcore.F(gcore.ContainerScaleDeploymentTriggersHTTPParam{
								Rate:   gcore.F(int64(1)),
								Window: gcore.F(int64(1)),
							}),
							Memory: gcore.F(gcore.ThresholdParam{
								Threshold: gcore.F(int64(70)),
							}),
						}),
					}),
				}}),
				FlavorID:      gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				Image:         gcore.F("nginx:latest"),
				ListeningPort: gcore.F(int64(8080)),
				Name:          gcore.F("my-instance"),
				APIKeyIDs:     gcore.F([]string{"3fa85f64-5717-4562-b3fc-2c963f66afa6"}),
				AuthEnabled:   gcore.F(false),
				Command:       gcore.F([]string{"nginx", "-g", "daemon off;"}),
				Description:   gcore.F("My first instance"),
				Envs: gcore.F(map[string]string{
					"DEBUG_MODE": "False",
					"KEY":        "12345",
				}),
				ImageRegistryID: gcore.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				Probes: gcore.F(gcore.InstanceProbesDeploymentParam{
					LivenessProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
						Enabled: gcore.F(true),
						Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
							Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
								Command: gcore.F([]string{"ls", "-l"}),
							}),
							FailureThreshold: gcore.F(int64(3)),
							HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
							TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
								Port: gcore.F(int64(80)),
							}),
							TimeoutSeconds: gcore.F(int64(1)),
						}),
					}),
					ReadinessProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
						Enabled: gcore.F(true),
						Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
							Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
								Command: gcore.F([]string{"ls", "-l"}),
							}),
							FailureThreshold: gcore.F(int64(3)),
							HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
							TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
								Port: gcore.F(int64(80)),
							}),
							TimeoutSeconds: gcore.F(int64(1)),
						}),
					}),
					StartupProbe: gcore.F(gcore.ContainerProbeConfigurationParam{
						Enabled: gcore.F(true),
						Probe: gcore.F(gcore.ContainerProbeConfigurationProbeParam{
							Exec: gcore.F(gcore.ContainerProbeConfigurationProbeExecParam{
								Command: gcore.F([]string{"ls", "-l"}),
							}),
							FailureThreshold: gcore.F(int64(3)),
							HTTPGet: gcore.F(gcore.ContainerProbeConfigurationProbeHTTPGetParam{
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
							TcpSocket: gcore.F(gcore.ContainerProbeConfigurationProbeTcpSocketParam{
								Port: gcore.F(int64(80)),
							}),
							TimeoutSeconds: gcore.F(int64(1)),
						}),
					}),
				}),
				ProjectID: gcore.F(int64(1)),
				Timeout:   gcore.F(int64(120)),
			},
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

func TestCloudV2InferenceDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Inference.Deployments.List(context.TODO(), gcore.CloudV2InferenceDeploymentListParams{
		Limit:     gcore.F(int64(0)),
		Offset:    gcore.F(int64(0)),
		OrderBy:   gcore.F("order_by"),
		ProjectID: gcore.F(int64(0)),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentDelete(t *testing.T) {
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
	err := client.Cloud.V2.Inference.Deployments.Delete(context.TODO(), "instance_id")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentGetLogs(t *testing.T) {
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
	_, err := client.Cloud.V2.Inference.Deployments.GetLogs(context.TODO(), "instance_id")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentStart(t *testing.T) {
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
	err := client.Cloud.V2.Inference.Deployments.Start(context.TODO(), "instance_id")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2InferenceDeploymentStop(t *testing.T) {
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
	err := client.Cloud.V2.Inference.Deployments.Stop(context.TODO(), "instance_id")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
