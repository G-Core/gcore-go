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

func TestCloudV1BminstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Bminstances.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1BminstanceNewParams{
			Flavor: gcore.F("bm1-basic-medium"),
			Interfaces: gcore.F([]gcore.CloudV1BminstanceNewParamsInterfaceUnion{gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("subnet"),
			}}),
			AppConfig:     gcore.F[any](map[string]interface{}{}),
			ApptemplateID: gcore.F("apptemplate_id"),
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
			ImageID:       gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			KeypairName:   gcore.F("keypair_name"),
			Metadata:      gcore.F[any](map[string]interface{}{}),
			NameTemplates: gcore.F([]string{"{}q}__} _-K."}),
			Names:         gcore.F([]string{"bmubuntu"}),
			Password:      gcore.F("password"),
			UserData:      gcore.F("user_data"),
			Username:      gcore.F("username"),
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

func TestCloudV1BminstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Bminstances.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1BminstanceListParams{
			ChangesBefore:           gcore.F("changes-before"),
			ChangesSince:            gcore.F("changes-since"),
			FlavorID:                gcore.F("flavor_id"),
			FlavorPrefix:            gcore.F("flavor_prefix"),
			IncludeK8s:              gcore.F(true),
			IP:                      gcore.F("ip"),
			Limit:                   gcore.F(int64(0)),
			MetadataKv:              gcore.F("metadata_kv"),
			MetadataV:               gcore.F("metadata_v"),
			Name:                    gcore.F("name"),
			Offset:                  gcore.F(int64(0)),
			OnlyIsolated:            gcore.F(true),
			OnlyWithFixedExternalIP: gcore.F(true),
			OrderBy:                 gcore.F("order_by"),
			ProfileName:             gcore.F("profile_name"),
			ProtectionStatus:        gcore.F("protection_status"),
			Status:                  gcore.F("status"),
			TypeDdosProfile:         gcore.F("type_ddos_profile"),
			Uuid:                    gcore.F("uuid"),
			WithDdos:                gcore.F(true),
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

func TestCloudV1BminstanceCheckQuotaWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Bminstances.CheckQuota(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1BminstanceCheckQuotaParams{
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

func TestCloudV1BminstanceGetAvailableFlavorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Bminstances.GetAvailableFlavors(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1BminstanceGetAvailableFlavorsParams{
			IncludePrices: gcore.F(true),
			ApptemplateID: gcore.F("apptemplate_id"),
			ImageID:       gcore.F("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
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

func TestCloudV1BminstanceRebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Bminstances.Rebuild(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
		gcore.CloudV1BminstanceRebuildParams{
			ImageID:  gcore.F("b5b4d65d-945f-4b98-ab6f-332319c724ef"),
			UserData: gcore.F("aGVsbG9fd29ybGQ="),
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
