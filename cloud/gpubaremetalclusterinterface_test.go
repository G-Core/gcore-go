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

func TestGPUBaremetalClusterInterfaceListWithOptionalParams(t *testing.T) {
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
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUBaremetalClusterInterfaceListParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(7),
			Limit:     gcore.Int(1000),
			Offset:    gcore.Int(0),
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
		"faab46fd-26fd-4321-9876-abcdef012345",
		cloud.GPUBaremetalClusterInterfaceAttachParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			OfExternal: &cloud.GPUBaremetalClusterInterfaceAttachParamsBodyExternal{
				DDOSProfile: cloud.GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfile{
					ProfileTemplate: 0,
					Fields: []cloud.GPUBaremetalClusterInterfaceAttachParamsBodyExternalDDOSProfileField{{
						BaseField:  0,
						FieldValue: map[string]any{},
						Value:      gcore.String("value"),
					}},
					ProfileTemplateName: gcore.String("profile_template_name"),
				},
				InterfaceName: gcore.String("interface_name"),
				IPFamily:      cloud.InterfaceIPFamilyDual,
				PortGroup:     gcore.Int(0),
				SecurityGroups: []cloud.GPUBaremetalClusterInterfaceAttachParamsBodyExternalSecurityGroup{{
					ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
				}},
				Type: "external",
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
		"faab46fd-26fd-4321-9876-abcdef012345",
		cloud.GPUBaremetalClusterInterfaceDetachParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			IPAddress: "ip_address",
			PortID:    "port_id",
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
