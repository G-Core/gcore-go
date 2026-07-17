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

func TestInstanceInterfaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Interfaces.List(
		context.TODO(),
		"bf325375-9af6-4c8b-a2fc-7f6f4ca02e2e",
		cloud.InstanceInterfaceListParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestInstanceInterfaceAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Interfaces.Attach(
		context.TODO(),
		"bf325375-9af6-4c8b-a2fc-7f6f4ca02e2e",
		cloud.InstanceInterfaceAttachParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			OfExternal: &cloud.InstanceInterfaceAttachParamsBodyExternal{
				DDOSProfile: cloud.InstanceInterfaceAttachParamsBodyExternalDDOSProfile{
					ProfileTemplate: 0,
					Fields: []cloud.InstanceInterfaceAttachParamsBodyExternalDDOSProfileField{{
						BaseField:  0,
						FieldValue: map[string]any{},
						Value:      gcore.String("value"),
					}},
					ProfileTemplateName: gcore.String("profile_template_name"),
				},
				InterfaceName: gcore.String("interface_name"),
				IPFamily:      cloud.InterfaceIPFamilyDual,
				PortGroup:     gcore.Int(0),
				SecurityGroups: []cloud.InstanceInterfaceAttachParamsBodyExternalSecurityGroup{{
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

func TestInstanceInterfaceDetach(t *testing.T) {
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
	_, err := client.Cloud.Instances.Interfaces.Detach(
		context.TODO(),
		"bf325375-9af6-4c8b-a2fc-7f6f4ca02e2e",
		cloud.InstanceInterfaceDetachParams{
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
