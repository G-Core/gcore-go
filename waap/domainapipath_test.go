// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/waap"
)

func TestDomainAPIPathNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.APIPaths.New(
		context.TODO(),
		1,
		waap.DomainAPIPathNewParams{
			HTTPScheme: waap.DomainAPIPathNewParamsHTTPSchemeHTTP,
			Method:     waap.DomainAPIPathNewParamsMethodGet,
			Path:       "/api/v1/paths/{path_id}",
			APIGroups:  []string{"accounts", "internal"},
			APIVersion: gcore.String("v1"),
			Tags:       []string{"sensitivedataurl", "highriskurl"},
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

func TestDomainAPIPathUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.Domains.APIPaths.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		waap.DomainAPIPathUpdateParams{
			DomainID:  1,
			APIGroups: []string{"accounts", "internal"},
			Path:      gcore.String("/api/v1/paths/{path_id}"),
			Status:    waap.DomainAPIPathUpdateParamsStatusConfirmedAPI,
			Tags:      []string{"sensitivedataurl", "highriskurl"},
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

func TestDomainAPIPathListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.APIPaths.List(
		context.TODO(),
		1,
		waap.DomainAPIPathListParams{
			APIGroup:   gcore.String("api_group"),
			APIVersion: gcore.String("api_version"),
			HTTPScheme: waap.DomainAPIPathListParamsHTTPSchemeHTTP,
			IDs:        []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Limit:      gcore.Int(0),
			Method:     waap.DomainAPIPathListParamsMethodGet,
			Offset:     gcore.Int(0),
			Ordering:   waap.DomainAPIPathListParamsOrderingID,
			Path:       gcore.String("path"),
			Source:     waap.DomainAPIPathListParamsSourceAPIDescriptionFile,
			Status:     []string{"CONFIRMED_API", "POTENTIAL_API"},
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

func TestDomainAPIPathDelete(t *testing.T) {
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
	err := client.Waap.Domains.APIPaths.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		waap.DomainAPIPathDeleteParams{
			DomainID: 1,
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

func TestDomainAPIPathGet(t *testing.T) {
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
	_, err := client.Waap.Domains.APIPaths.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		waap.DomainAPIPathGetParams{
			DomainID: 1,
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
