// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

func TestInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.New(context.TODO(), cloud.InstanceNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Flavor:    "g2-standard-32-64",
		Interfaces: []cloud.InstanceNewParamsInterfaceUnion{{
			OfExternal: &cloud.InstanceNewParamsInterfaceExternal{
				InterfaceName: gcore.String("eth0"),
				IPFamily:      cloud.InterfaceIPFamilyIpv4,
				SecurityGroups: []cloud.InstanceNewParamsInterfaceExternalSecurityGroup{{
					ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
				}},
			},
		}},
		Volumes: []cloud.InstanceNewParamsVolumeUnion{{
			OfNewVolume: &cloud.InstanceNewParamsVolumeNewVolume{
				Size:                20,
				AttachmentTag:       gcore.String("boot"),
				DeleteOnTermination: gcore.Bool(false),
				Name:                gcore.String("boot-volume"),
				Tags: cloud.TagUpdateList{
					"foo": "my-tag-value",
				},
				TypeName: "ssd_hiiops",
			},
		}},
		AllowAppPorts: gcore.Bool(true),
		Configuration: map[string]interface{}{},
		Name:          gcore.String("my-instance"),
		NameTemplate:  gcore.String("name_template"),
		Password:      gcore.String("password"),
		SecurityGroups: []cloud.InstanceNewParamsSecurityGroup{{
			ID: "ae74714c-c380-48b4-87f8-758d656cdad6",
		}},
		ServergroupID: gcore.String("servergroup_id"),
		SSHKeyName:    gcore.String("my-ssh-key"),
		Tags: cloud.TagUpdateList{
			"foo": "my-tag-value",
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

func TestInstanceUpdate(t *testing.T) {
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
	_, err := client.Cloud.Instances.Update(
		context.TODO(),
		"instance_id",
		cloud.InstanceUpdateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Name:      "my-resource",
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

func TestInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.List(context.TODO(), cloud.InstanceListParams{
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
		IncludeK8s:              gcore.Bool(true),
		IP:                      gcore.String("192.168.0.1"),
		Limit:                   gcore.Int(1000),
		Name:                    gcore.String("name"),
		Offset:                  gcore.Int(0),
		OnlyIsolated:            gcore.Bool(true),
		OnlyWithFixedExternalIP: gcore.Bool(true),
		OrderBy:                 cloud.InstanceListParamsOrderByNameAsc,
		ProfileName:             gcore.String("profile_name"),
		ProtectionStatus:        cloud.InstanceListParamsProtectionStatusActive,
		Status:                  cloud.InstanceListParamsStatusActive,
		TagKeyValue:             gcore.String("tag_key_value"),
		TagValue:                []string{"value1", "value2"},
		TypeDDOSProfile:         cloud.InstanceListParamsTypeDDOSProfileAdvanced,
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

func TestInstanceDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Delete(
		context.TODO(),
		"instance_id",
		cloud.InstanceDeleteParams{
			ProjectID:        gcore.Int(0),
			RegionID:         gcore.Int(0),
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

func TestInstanceActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Action(
		context.TODO(),
		"instance_id",
		cloud.InstanceActionParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			OfStartActionInstanceSerializer: &cloud.InstanceActionParamsBodyStartActionInstanceSerializer{
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

func TestInstanceAddToPlacementGroup(t *testing.T) {
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
	_, err := client.Cloud.Instances.AddToPlacementGroup(
		context.TODO(),
		"instance_id",
		cloud.InstanceAddToPlacementGroupParams{
			ProjectID:     gcore.Int(0),
			RegionID:      gcore.Int(0),
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

func TestInstanceAssignSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.Instances.AssignSecurityGroup(
		context.TODO(),
		"instance_id",
		cloud.InstanceAssignSecurityGroupParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Name:      gcore.String("some_name"),
			PortsSecurityGroupNames: []cloud.InstanceAssignSecurityGroupParamsPortsSecurityGroupName{{
				PortID:             param.NullOpt[string](),
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

func TestInstanceDisablePortSecurity(t *testing.T) {
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
	_, err := client.Cloud.Instances.DisablePortSecurity(
		context.TODO(),
		"port_id",
		cloud.InstanceDisablePortSecurityParams{
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

func TestInstanceEnablePortSecurity(t *testing.T) {
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
	_, err := client.Cloud.Instances.EnablePortSecurity(
		context.TODO(),
		"port_id",
		cloud.InstanceEnablePortSecurityParams{
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

func TestInstanceGet(t *testing.T) {
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
	_, err := client.Cloud.Instances.Get(
		context.TODO(),
		"instance_id",
		cloud.InstanceGetParams{
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

func TestInstanceGetConsoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.GetConsole(
		context.TODO(),
		"instance_id",
		cloud.InstanceGetConsoleParams{
			ProjectID:   gcore.Int(0),
			RegionID:    gcore.Int(0),
			ConsoleType: gcore.String("console_type"),
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

func TestInstanceRemoveFromPlacementGroup(t *testing.T) {
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
	_, err := client.Cloud.Instances.RemoveFromPlacementGroup(
		context.TODO(),
		"instance_id",
		cloud.InstanceRemoveFromPlacementGroupParams{
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

func TestInstanceResize(t *testing.T) {
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
	_, err := client.Cloud.Instances.Resize(
		context.TODO(),
		"instance_id",
		cloud.InstanceResizeParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
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

func TestInstanceUnassignSecurityGroupWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.Instances.UnassignSecurityGroup(
		context.TODO(),
		"instance_id",
		cloud.InstanceUnassignSecurityGroupParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Name:      gcore.String("some_name"),
			PortsSecurityGroupNames: []cloud.InstanceUnassignSecurityGroupParamsPortsSecurityGroupName{{
				PortID:             param.NullOpt[string](),
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
