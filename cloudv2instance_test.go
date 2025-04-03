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

func TestCloudV2InstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Instances.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV2InstanceNewParams{
			Flavor: gcore.F("g1-standard-1-2"),
			Interfaces: gcore.F([]gcore.CloudV2InstanceNewParamsInterfaceUnion{gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("4536dba1-93b1-492e-b3df-270b6b9f3650"),
				}}),
				Type: gcore.F("subnet"),
			}, gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("cee2ca1f-507a-4a31-b714-f6c1ffb4bdfa"),
				}}),
				Type: gcore.F("subnet"),
			}, gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("ae74714c-c380-48b4-87f8-758d656cdad6"),
				}}),
				Type: gcore.F("subnet"),
			}, gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("ae74714c-c380-48b4-87f8-758d656cdad6"),
				}}),
				Type: gcore.F("any_subnet"),
			}, gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("ae74714c-c380-48b4-87f8-758d656cdad6"),
				}}),
				Type: gcore.F("external"),
			}, gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchema{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.CloudV2InstanceNewParamsInterfacesNewInterfaceExternalExtendSchemaIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
					ID: gcore.F("ae74714c-c380-48b4-87f8-758d656cdad6"),
				}}),
				Type: gcore.F("reserved_fixed_ip"),
			}}),
			Volumes: gcore.F([]gcore.CreateInstanceVolumeParam{{
				Source:              gcore.F(gcore.CreateInstanceVolumeSourceApptemplate),
				ApptemplateID:       gcore.F("apptemplate_id"),
				AttachmentTag:       gcore.F("root"),
				BootIndex:           gcore.F(int64(0)),
				DeleteOnTermination: gcore.F(false),
				ImageID:             gcore.F("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
				Metadata: gcore.F[any](map[string]interface{}{
					"key": "value",
				}),
				Name:       gcore.F("TestVM5 Ubuntu boot image"),
				Size:       gcore.F(int64(10)),
				SnapshotID: gcore.F("snapshot_id"),
				TypeName:   gcore.F(gcore.CreateInstanceVolumeTypeNameCold),
				VolumeID:   gcore.F("volume_id"),
			}, {
				Source:              gcore.F(gcore.CreateInstanceVolumeSourceApptemplate),
				ApptemplateID:       gcore.F("apptemplate_id"),
				AttachmentTag:       gcore.F("attachment_tag"),
				BootIndex:           gcore.F(int64(0)),
				DeleteOnTermination: gcore.F(true),
				ImageID:             gcore.F("image_id"),
				Metadata:            gcore.F[any](map[string]interface{}{}),
				Name:                gcore.F("empty volume"),
				Size:                gcore.F(int64(5)),
				SnapshotID:          gcore.F("snapshot_id"),
				TypeName:            gcore.F(gcore.CreateInstanceVolumeTypeNameCold),
				VolumeID:            gcore.F("volume_id"),
			}}),
			AllowAppPorts: gcore.F(true),
			Configuration: gcore.F[any](map[string]interface{}{}),
			KeypairName:   gcore.F("my_secret_keypair"),
			Metadata: gcore.F[any](map[string]interface{}{
				"metadata_key": "metadata_value",
			}),
			NameTemplates: gcore.F([]string{"ed-c4-{ip_octets}"}),
			Names:         gcore.F([]string{"bubuntu"}),
			Password:      gcore.F("password"),
			SecurityGroups: gcore.F([]gcore.MandatoryIDParam{{
				ID: gcore.F("4536dba1-93b1-492e-b3df-270b6b9f3650"),
			}, {
				ID: gcore.F("cee2ca1f-507a-4a31-b714-f6c1ffb4bdfa"),
			}}),
			ServergroupID: gcore.F("servergroup_id"),
			UserData:      gcore.F("aGVsbG9fd29ybGQ="),
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

func TestCloudV2InstanceCheckLimitsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Instances.CheckLimits(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV2InstanceCheckLimitsParams{
			Flavor: gcore.F("g1-standard-1-2"),
			Interfaces: gcore.F([]gcore.CloudV2InstanceCheckLimitsParamsInterfaceUnion{gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("subnet"),
			}, gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("subnet"),
			}, gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("subnet"),
			}, gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("any_subnet"),
			}, gcore.NewInterfaceExternalParam{
				InterfaceName: gcore.F("interface_name"),
				IPFamily:      gcore.F(gcore.NewInterfaceExternalIPFamilyDual),
				PortGroup:     gcore.F(int64(0)),
				Type:          gcore.F("external"),
			}}),
			NameTemplates: gcore.F([]string{"string"}),
			Names:         gcore.F([]string{"cirroz1", "cirroz2"}),
			Volumes: gcore.F([]gcore.CloudV2InstanceCheckLimitsParamsVolume{{
				Source:     gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesSourceApptemplate),
				Size:       gcore.F(int64(10)),
				SnapshotID: gcore.F("snapshot_id"),
				TypeName:   gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesTypeNameCold),
			}, {
				Source:     gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesSourceApptemplate),
				Size:       gcore.F(int64(5)),
				SnapshotID: gcore.F("snapshot_id"),
				TypeName:   gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesTypeNameCold),
			}, {
				Source:     gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesSourceApptemplate),
				Size:       gcore.F(int64(0)),
				SnapshotID: gcore.F("7cca40d7-a843-4e9f-ae08-62b9a394b1ab"),
				TypeName:   gcore.F(gcore.CloudV2InstanceCheckLimitsParamsVolumesTypeNameCold),
			}}),
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

func TestCloudV2InstancePerformActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Instances.PerformAction(
		context.TODO(),
		int64(0),
		int64(0),
		"instance_id",
		gcore.CloudV2InstancePerformActionParams{
			Body: gcore.CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializer{
				Action:          gcore.F(gcore.CloudV2InstancePerformActionParamsBodyStartActionInstanceSerializerActionStart),
				ActivateProfile: gcore.F(true),
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
