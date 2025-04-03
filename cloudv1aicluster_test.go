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

func TestCloudV1AIClusterAddSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.V1.AI.Clusters.AddSecurityGroup(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_id",
		gcore.CloudV1AIClusterAddSecurityGroupParams{
			InstancePortsSecurityGroups: gcore.InstancePortsSecurityGroupsParam{
				Name: gcore.F("some_name"),
				PortsSecurityGroupNames: gcore.F([]gcore.InstancePortsSecurityGroupsPortsSecurityGroupNameParam{{
					PortID:             gcore.Null[string](),
					SecurityGroupNames: gcore.F([]string{"some_name"}),
				}, {
					PortID:             gcore.F("ee2402d0-f0cd-4503-9b75-69be1d11c5f1"),
					SecurityGroupNames: gcore.F([]string{"name1", "name2"}),
				}}),
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

func TestCloudV1AIClusterAttachInterfaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.AttachInterface(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
		gcore.CloudV1AIClusterAttachInterfaceParams{
			Body: gcore.NewInterfaceExternalExtendWithDdosParam{
				DdosProfile: gcore.F(gcore.DeprecatedCreateDdosProfileParam{
					ProfileTemplate: gcore.F(int64(29)),
					Fields: gcore.F([]gcore.DeprecatedCreateDdosProfileFieldParam{{
						BaseField: gcore.F(int64(10)),
						FieldName: gcore.F("field_name"),
						FieldValue: gcore.F[any](map[string]interface{}{
							"0": 45046,
							"1": 45047,
						}),
						Value: gcore.Null[string](),
					}}),
					ProfileTemplateName: gcore.F("profile_template_name"),
				}),
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalExtendWithDdosIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("4536dba1-93b1-492e-b3df-270b6b9f3650"),
				}, {
					ID: gcore.F("cee2ca1f-507a-4a31-b714-f6c1ffb4bdfa"),
				}}),
				Type: gcore.F("external"),
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

func TestCloudV1AIClusterCheckQuotaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.CheckQuota(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1AIClusterCheckQuotaParams{
			CheckQuotaBeforeBmCreation: gcore.CheckQuotaBeforeBmCreationParam{
				Flavor: gcore.F("bm1-basic-medium"),
				Interfaces: gcore.F([]gcore.CheckQuotaBeforeBmCreationInterfacesUnionParam{gcore.NewInterfaceExternalParam{
					InterfaceName: gcore.F("interface_name"),
					IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
					PortGroup:     gcore.F(int64(0)),
					Type:          gcore.F("subnet"),
				}}),
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

func TestCloudV1AIClusterDetachInterface(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.DetachInterface(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
		gcore.CloudV1AIClusterDetachInterfaceParams{
			PortIDWithIP: gcore.PortIDWithIPParam{
				IPAddress: gcore.F("192.168.123.20"),
				PortID:    gcore.F("351b0dd7-ca09-431c-be53-935db3785067"),
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

func TestCloudV1AIClusterGetConsole(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.GetConsole(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1AIClusterListInterfaces(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.ListInterfaces(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1AIClusterListPorts(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.ListPorts(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1AIClusterPowercycle(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.Powercycle(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1AIClusterReboot(t *testing.T) {
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
	_, err := client.Cloud.V1.AI.Clusters.Reboot(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1AIClusterRemoveSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.V1.AI.Clusters.RemoveSecurityGroup(
		context.TODO(),
		int64(0),
		int64(0),
		"cluster_id",
		gcore.CloudV1AIClusterRemoveSecurityGroupParams{
			InstancePortsSecurityGroups: gcore.InstancePortsSecurityGroupsParam{
				Name: gcore.F("some_name"),
				PortsSecurityGroupNames: gcore.F([]gcore.InstancePortsSecurityGroupsPortsSecurityGroupNameParam{{
					PortID:             gcore.Null[string](),
					SecurityGroupNames: gcore.F([]string{"some_name"}),
				}, {
					PortID:             gcore.F("ee2402d0-f0cd-4503-9b75-69be1d11c5f1"),
					SecurityGroupNames: gcore.F([]string{"name1", "name2"}),
				}}),
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
