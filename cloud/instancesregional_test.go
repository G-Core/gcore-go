// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func TestInstancesRegionalNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.New(context.TODO(), cloud.InstancesRegionalNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Flavor:    "g2-standard-32-64",
		Interfaces: []cloud.InstancesRegionalNewParamsInterfaceUnion{{
			OfExternal: &cloud.InstancesRegionalNewParamsInterfaceExternal{
				InterfaceName: gcore.String("eth0"),
				IPFamily:      cloud.InterfaceIPFamilyIpv4,
				SecurityGroups: []cloud.InstancesRegionalNewParamsInterfaceExternalSecurityGroup{{
					ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
				}},
			},
		}},
		Volumes: []cloud.InstancesRegionalNewParamsVolumeUnion{{
			OfImage: &cloud.InstancesRegionalNewParamsVolumeImage{
				ImageID:             "e460e48c-6836-447e-bc9c-16fc4225d318",
				AttachmentTag:       gcore.String("boot"),
				BootIndex:           gcore.Int(0),
				DeleteOnTermination: gcore.Bool(false),
				Name:                gcore.String("boot-volume"),
				Size:                gcore.Int(50),
				Tags: map[string]string{
					"my-tag": "my-tag-value",
				},
				TypeName: "ssd_hiiops",
			},
		}},
		AllowAppPorts: gcore.Bool(true),
		Configuration: map[string]any{
			"foo": "bar",
		},
		Name:         gcore.String("my-instance"),
		NameTemplate: gcore.String("name_template"),
		Password:     gcore.String("password"),
		SecurityGroups: []cloud.InstancesRegionalNewParamsSecurityGroup{{
			ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
		}},
		ServergroupID: gcore.String("servergroup_id"),
		SSHKeyName:    gcore.String("my-ssh-key"),
		Tags: map[string]string{
			"my-tag": "my-tag-value",
		},
		UserData: gcore.String("user_data"),
		Username: gcore.String("username"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstancesRegionalUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.Update(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalUpdateParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("instance_name"),
			Tags: cloud.TagUpdateMap{
				"my-tag":           gcore.Ptr("my-tag-value"),
				"my-tag-to-remove": nil,
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

func TestInstancesRegionalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.List(context.TODO(), cloud.InstancesRegionalListParams{
		ProjectID:               gcore.Int(1),
		RegionID:                gcore.Int(1),
		AvailableFloating:       gcore.Bool(true),
		ChangesBefore:           gcore.Time(time.Now()),
		ChangesSince:            gcore.Time(time.Now()),
		ExcludeFlavorPrefix:     gcore.String("g1-"),
		ExcludeSecgroup:         gcore.String("secgroup_name"),
		FlavorID:                gcore.String("g2-standard-32-64"),
		FlavorPrefix:            gcore.String("g2-"),
		IncludeAI:               gcore.Bool(false),
		IncludeBaremetal:        gcore.Bool(false),
		IncludeK8S:              gcore.Bool(true),
		IP:                      gcore.String("192.168.0.1"),
		Limit:                   gcore.Int(1000),
		Name:                    gcore.String("name"),
		Offset:                  gcore.Int(0),
		OnlyIsolated:            gcore.Bool(true),
		OnlyWithFixedExternalIP: gcore.Bool(true),
		OrderBy:                 cloud.InstancesRegionalListParamsOrderByNameAsc,
		ProfileName:             gcore.String("profile_name"),
		ProtectionStatus:        cloud.InstancesRegionalListParamsProtectionStatusActive,
		Status:                  cloud.InstancesRegionalListParamsStatusActive,
		TagKeyValue:             gcore.String("tag_key_value"),
		TagValue:                []string{"value1", "value2"},
		TypeDDOSProfile:         cloud.InstancesRegionalListParamsTypeDDOSProfileAdvanced,
		Uuid:                    gcore.String("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
		WithDDOS:                gcore.Bool(true),
		WithInterfacesName:      gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstancesRegionalDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.Delete(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalDeleteParams{
			ProjectID:        gcore.Int(1),
			RegionID:         gcore.Int(1),
			DeleteFloatings:  gcore.Bool(true),
			Floatings:        gcore.String("floatings"),
			ReservedFixedIPs: gcore.String("reserved_fixed_ips"),
			Volumes:          gcore.String("volume_id_1,volume_id_2"),
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

func TestInstancesRegionalActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.Action(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalActionParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			OfStartActionInstanceSerializer: &cloud.InstancesRegionalActionParamsBodyStartActionInstanceSerializer{
				ActivateProfile: gcore.Bool(true),
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

func TestInstancesRegionalAddToPlacementGroup(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.AddToPlacementGroup(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalAddToPlacementGroupParams{
			ProjectID:     gcore.Int(1),
			RegionID:      gcore.Int(1),
			ServergroupID: "47003067-550a-6f17-93b6-81ee16ba061e",
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

func TestInstancesRegionalAssignSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.InstancesRegional.AssignSecurityGroup(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalAssignSecurityGroupParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("some_name"),
			PortsSecurityGroupNames: []cloud.InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName{{
				PortID:             param.Null[string](),
				SecurityGroupNames: []string{"some_name"},
			}, {
				PortID:             gcore.String("ee2402d0-f0cd-4503-9b75-69be1d11c5f1"),
				SecurityGroupNames: []string{"name1", "name2"},
			}},
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

func TestInstancesRegionalDisablePortSecurity(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.DisablePortSecurity(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.InstancesRegionalDisablePortSecurityParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestInstancesRegionalEnablePortSecurity(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.EnablePortSecurity(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.InstancesRegionalEnablePortSecurityParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestInstancesRegionalGet(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.Get(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestInstancesRegionalGetConsoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.GetConsole(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalGetConsoleParams{
			ProjectID:   gcore.Int(1),
			RegionID:    gcore.Int(1),
			ConsoleType: cloud.InstancesRegionalGetConsoleParamsConsoleTypeNovnc,
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

func TestInstancesRegionalRemoveFromPlacementGroup(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.RemoveFromPlacementGroup(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalRemoveFromPlacementGroupParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestInstancesRegionalResize(t *testing.T) {
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
	_, err := client.Cloud.InstancesRegional.Resize(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalResizeParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			FlavorID:  "g1s-shared-1-0.5",
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

func TestInstancesRegionalUnassignSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.InstancesRegional.UnassignSecurityGroup(
		context.TODO(),
		"b5b4d65d-945f-4b98-ab6f-332319c724ef",
		cloud.InstancesRegionalUnassignSecurityGroupParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      gcore.String("some_name"),
			PortsSecurityGroupNames: []cloud.InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName{{
				PortID:             param.Null[string](),
				SecurityGroupNames: []string{"some_name"},
			}, {
				PortID:             gcore.String("ee2402d0-f0cd-4503-9b75-69be1d11c5f1"),
				SecurityGroupNames: []string{"name1", "name2"},
			}},
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
