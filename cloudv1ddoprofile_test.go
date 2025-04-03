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

func TestCloudV1DdoProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Ddos.Profiles.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1DdoProfileNewParams{
			CreateClientProfile: gcore.CreateClientProfileParam{
				Fields: gcore.F([]gcore.CreateClientProfileFieldParam{{
					BaseField:  gcore.F(int64(123)),
					FieldName:  gcore.F("field_name"),
					FieldValue: gcore.F[gcore.CreateClientProfileFieldsFieldValueUnionParam](gcore.CreateClientProfileFieldsFieldValueArrayParam([]interface{}{map[string]interface{}{}})),
					Value:      gcore.F("value"),
				}}),
				IPAddress:           gcore.F("11.111.111.1"),
				ProfileTemplate:     gcore.F(int64(123)),
				BmInstanceID:        gcore.F("bm_instance_id"),
				ProfileTemplateName: gcore.F("profile_template_name"),
				ResourceID:          gcore.F("resource_id"),
				ResourceType:        gcore.F(gcore.CreateClientProfileResourceTypeInstance),
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

func TestCloudV1DdoProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Ddos.Profiles.Update(
		context.TODO(),
		int64(0),
		int64(0),
		int64(0),
		gcore.CloudV1DdoProfileUpdateParams{
			CreateClientProfile: gcore.CreateClientProfileParam{
				Fields: gcore.F([]gcore.CreateClientProfileFieldParam{{
					BaseField:  gcore.F(int64(123)),
					FieldName:  gcore.F("field_name"),
					FieldValue: gcore.F[gcore.CreateClientProfileFieldsFieldValueUnionParam](gcore.CreateClientProfileFieldsFieldValueArrayParam([]interface{}{map[string]interface{}{}})),
					Value:      gcore.F("value"),
				}}),
				IPAddress:           gcore.F("11.111.111.1"),
				ProfileTemplate:     gcore.F(int64(123)),
				BmInstanceID:        gcore.F("bm_instance_id"),
				ProfileTemplateName: gcore.F("profile_template_name"),
				ResourceID:          gcore.F("resource_id"),
				ResourceType:        gcore.F(gcore.CreateClientProfileResourceTypeInstance),
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

func TestCloudV1DdoProfileList(t *testing.T) {
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
	_, err := client.Cloud.V1.Ddos.Profiles.List(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1DdoProfileDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Ddos.Profiles.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1DdoProfileActivateDeactivate(t *testing.T) {
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
	_, err := client.Cloud.V1.Ddos.Profiles.ActivateDeactivate(
		context.TODO(),
		int64(0),
		int64(0),
		int64(0),
		gcore.CloudV1DdoProfileActivateDeactivateParams{
			Active: gcore.F(true),
			Bgp:    gcore.F(true),
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
