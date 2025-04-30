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
	"github.com/stainless-sdks/gcore-go/packages/param"
)

func TestGPUBaremetalClusterServerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.Delete(
		context.TODO(),
		"cluster_id",
		"instance_id",
		cloud.GPUBaremetalClusterServerDeleteParams{
			ProjectID:       gcore.Int(0),
			RegionID:        gcore.Int(0),
			DeleteFloatings: gcore.Bool(true),
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

func TestGPUBaremetalClusterServerAttachInterfaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.AttachInterface(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterServerAttachInterfaceParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Body: cloud.GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS{
				DDOSProfile: cloud.GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile{
					ProfileTemplate: 29,
					Fields: []cloud.GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField{{
						BaseField: gcore.Int(10),
						FieldName: gcore.String("field_name"),
						FieldValue: map[string]interface{}{
							"0": 45046,
							"1": 45047,
						},
						Value: param.NullOpt[string](),
					}},
					ProfileTemplateName: gcore.String("profile_template_name"),
				},
				InterfaceName: gcore.String("interface_name"),
				IPFamily:      "dual",
				PortGroup:     gcore.Int(0),
				SecurityGroups: []cloud.GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup{{
					ID: "4536dba1-93b1-492e-b3df-270b6b9f3650",
				}, {
					ID: "cee2ca1f-507a-4a31-b714-f6c1ffb4bdfa",
				}},
				Type: gcore.String("external"),
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

func TestGPUBaremetalClusterServerDetachInterface(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.DetachInterface(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterServerDetachInterfaceParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			IPAddress: "192.168.123.20",
			PortID:    "351b0dd7-ca09-431c-be53-935db3785067",
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

func TestGPUBaremetalClusterServerGetConsole(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.GetConsole(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterServerGetConsoleParams{
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

func TestGPUBaremetalClusterServerPowercycle(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.Powercycle(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterServerPowercycleParams{
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

func TestGPUBaremetalClusterServerReboot(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetalClusters.Servers.Reboot(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterServerRebootParams{
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
