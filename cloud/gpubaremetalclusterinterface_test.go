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
	"github.com/G-Core/gcore-go/packages/param"
)

func TestGPUBaremetalClusterInterfaceList(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetal.Clusters.Interfaces.List(
		context.TODO(),
		"cluster_id",
		cloud.GPUBaremetalClusterInterfaceListParams{
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

func TestGPUBaremetalClusterInterfaceAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetal.Clusters.Interfaces.Attach(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterInterfaceAttachParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			OfNewInterfaceExternalExtendSchemaWithDDOS: &cloud.GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOS{
				DDOSProfile: cloud.GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile{
					ProfileTemplate: 29,
					Fields: []cloud.GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField{{
						BaseField: gcore.Int(10),
						FieldName: gcore.String("field_name"),
						FieldValue: []float64{
							45046,
							45047,
						},
						Value: param.Null[string](),
					}},
					ProfileTemplateName: gcore.String("profile_template_name"),
				},
				InterfaceName: gcore.String("interface_name"),
				IPFamily:      "dual",
				PortGroup:     gcore.Int(0),
				SecurityGroups: []cloud.GPUBaremetalClusterInterfaceAttachParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup{{
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

func TestGPUBaremetalClusterInterfaceDetach(t *testing.T) {
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
	_, err := client.Cloud.GPUBaremetal.Clusters.Interfaces.Detach(
		context.TODO(),
		"instance_id",
		cloud.GPUBaremetalClusterInterfaceDetachParams{
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
