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

func TestWaapV1DomainAPIPathNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.APIPaths.New(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainAPIPathNewParams{
			HTTPScheme: gcore.F(gcore.HTTPSchemeHTTP),
			Method:     gcore.F(gcore.MethodGet),
			Path:       gcore.F("/api/v1/paths/{path_id}"),
			APIGroups:  gcore.F([]string{"accounts", "internal"}),
			APIVersion: gcore.F("v1"),
			Tags:       gcore.F([]string{"sensitivedataurl", "highriskurl"}),
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

func TestWaapV1DomainAPIPathGet(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.APIPaths.Get(
		context.TODO(),
		int64(0),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1DomainAPIPathUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.V1.Domains.APIPaths.Update(
		context.TODO(),
		int64(0),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		gcore.WaapV1DomainAPIPathUpdateParams{
			APIGroups: gcore.F([]string{"accounts", "internal"}),
			Path:      gcore.F("/api/v1/paths/{path_id}"),
			Status:    gcore.F(gcore.APIPathStatusConfirmedAPI),
			Tags:      gcore.F([]string{"sensitivedataurl", "highriskurl"}),
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

func TestWaapV1DomainAPIPathListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.APIPaths.List(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainAPIPathListParams{
			APIGroup:   gcore.F("api_group"),
			APIVersion: gcore.F("api_version"),
			HTTPScheme: gcore.F(gcore.HTTPSchemeHTTP),
			IDs:        gcore.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Limit:      gcore.F(int64(0)),
			Method:     gcore.F(gcore.MethodGet),
			Offset:     gcore.F(int64(0)),
			Ordering:   gcore.F(gcore.WaapV1DomainAPIPathListParamsOrderingID),
			Path:       gcore.F("path"),
			Source:     gcore.F(gcore.SourceAPIDescriptionFile),
			Status:     gcore.F([]gcore.APIPathStatus{gcore.APIPathStatusConfirmedAPI}),
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

func TestWaapV1DomainAPIPathDelete(t *testing.T) {
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
	err := client.Waap.V1.Domains.APIPaths.Delete(
		context.TODO(),
		int64(0),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
