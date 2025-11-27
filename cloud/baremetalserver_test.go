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

func TestBaremetalServerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.New(context.TODO(), cloud.BaremetalServerNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  gcore.Int(1),
		Flavor:    "bm2-hf-medium",
		Interfaces: []cloud.BaremetalServerNewParamsInterfaceUnion{{
			OfExternal: &cloud.BaremetalServerNewParamsInterfaceExternal{
				InterfaceName: gcore.String("eth0"),
				IPFamily:      cloud.InterfaceIPFamilyIpv4,
				PortGroup:     gcore.Int(0),
			},
		}},
		AppConfig: map[string]any{
			"foo": "bar",
		},
		ApptemplateID: gcore.String("apptemplate_id"),
		DDOSProfile: cloud.BaremetalServerNewParamsDDOSProfile{
			ProfileTemplate: 123,
			Fields: []cloud.BaremetalServerNewParamsDDOSProfileField{{
				BaseField: gcore.Int(10),
				FieldName: param.Null[string](),
				FieldValue: []float64{
					45046,
					45047,
				},
				Value: param.Null[string](),
			}},
		},
		ImageID:      gcore.String("image_id"),
		Name:         gcore.String("my-bare-metal"),
		NameTemplate: gcore.String("name_template"),
		Password:     gcore.String("password"),
		SSHKeyName:   gcore.String("my-ssh-key"),
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

func TestBaremetalServerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.List(context.TODO(), cloud.BaremetalServerListParams{
		ProjectID:               gcore.Int(1),
		RegionID:                gcore.Int(1),
		ChangesBefore:           gcore.Time(time.Now()),
		ChangesSince:            gcore.Time(time.Now()),
		FlavorID:                gcore.String("bm2-hf-small"),
		FlavorPrefix:            gcore.String("bm2-"),
		IncludeK8S:              gcore.Bool(true),
		IP:                      gcore.String("192.168.0.1"),
		Limit:                   gcore.Int(1000),
		Name:                    gcore.String("name"),
		Offset:                  gcore.Int(0),
		OnlyIsolated:            gcore.Bool(true),
		OnlyWithFixedExternalIP: gcore.Bool(true),
		OrderBy:                 cloud.BaremetalServerListParamsOrderByNameAsc,
		ProfileName:             gcore.String("profile_name"),
		ProtectionStatus:        cloud.BaremetalServerListParamsProtectionStatusActive,
		Status:                  cloud.BaremetalServerListParamsStatusActive,
		TagKeyValue:             gcore.String("tag_key_value"),
		TagValue:                []string{"value1", "value2"},
		TypeDDOSProfile:         cloud.BaremetalServerListParamsTypeDDOSProfileAdvanced,
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

func TestBaremetalServerRebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Baremetal.Servers.Rebuild(
		context.TODO(),
		"024a29e-b4b7-4c91-9a46-505be123d9f8",
		cloud.BaremetalServerRebuildParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			ImageID:   gcore.String("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
			UserData:  gcore.String("aGVsbG9fd29ybGQ="),
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
