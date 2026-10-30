// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestTrustedCaCertificateNew(t *testing.T) {
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
	_, err := client.CDN.TrustedCaCertificates.New(context.TODO(), cdn.TrustedCaCertificateNewParams{
		Name:           "Example CA cert",
		SslCertificate: "-----BEGIN CERTIFICATE-----\nMIIBbTCCARKgAwIBAgIULDeBmErpf2HbicPq8M/WK6fUZxEwCgYIKoZIzj0EAwIw\nKjETMBEGA1UECgwKRXhhbXBsZSBDQTETMBEGA1UEAwwKRXhhbXBsZSBDQTAeFw0y\nNjAxMDEwMDAwMDBaFw0zNjAxMDEwMDAwMDBaMCoxEzARBgNVBAoMCkV4YW1wbGUg\nQ0ExEzARBgNVBAMMCkV4YW1wbGUgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAATdtaj2FA1AUBTNxVPfCeuuJWY6D1SQTtCNT3vbakavfJu+QISYQcc0Tbtudl5s\nv6mQPTiVlKOzi5NSWYn9F8YwoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMAoGCCqG\nSM49BAMCA0kAMEYCIQD4GWDCyoYDx48b49W/R6ox5K2Dx37otAqQCSeB96moHAIh\nAKAUFyIRFL4HQ1Y3jEELs8TvC2554gcPQKitr0S0wM1g\n-----END CERTIFICATE-----\n",
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrustedCaCertificateListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.TrustedCaCertificates.List(context.TODO(), cdn.TrustedCaCertificateListParams{
		Automated:           gcore.Bool(true),
		Limit:               gcore.Int(1),
		Offset:              gcore.Int(0),
		ResourceID:          gcore.Int(0),
		ValidityNotAfterLte: gcore.Time(time.Now()),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrustedCaCertificateDelete(t *testing.T) {
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
	err := client.CDN.TrustedCaCertificates.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrustedCaCertificateGet(t *testing.T) {
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
	_, err := client.CDN.TrustedCaCertificates.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTrustedCaCertificateReplace(t *testing.T) {
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
	_, err := client.CDN.TrustedCaCertificates.Replace(
		context.TODO(),
		0,
		cdn.TrustedCaCertificateReplaceParams{
			Name: "Example CA cert 2",
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
