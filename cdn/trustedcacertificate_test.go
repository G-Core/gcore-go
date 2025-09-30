// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"context"
	"errors"
	"os"
	"testing"

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
	_, err := client.Cdn.TrustedCaCertificates.New(context.TODO(), cdn.TrustedCaCertificateNewParams{
		Name:           "Example CA cert",
		SslCertificate: "-----BEGIN CERTIFICATE-----\nMIIC0zCCAbugAwIBAgICA+gwDQYJKoZIhvcNAQELBQAwFjEUMBIGA1UEAwwLZXhh\nbXBsZS5jb20wHhcNMjAwNjI2MTIwMzUzWhcNMjEwNjI2MTIwMzUzWjAWMRQwEgYD\nVQQDDAtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB\nAN4nnSfTsMEnfPgL7rkbImxZAQoND+bpPoX8q16iXZz3fFfqdRk+uEIpU3Brleeg\np0zrrT2eI3+c2h/PRod0Fam4TO6EcfwuboUFzV3j6yw6aWdfBjWZsWBR/FoqWLYq\nb3UejN7yiTYNSiIy3zVpi9pnFM8N8qT+VGBrRDGef2v9JCzhsSSU7wAYM5HKZTp+\nWHojjiyB2hOYqft7A2WlTEDmHFa5UcPHMRZKATUYI1T2TRVqLlSiE2mJ3dFRXGM2\nZAS33J0NVUjkx3w8RmJ7DNflEFJt/6IXdfaokVgfza7LFarrQFQP/YURXEeJT7jm\nDvKpZ/a8wu3ve6N4ykC+CBsCAwEAAaMrMCkwDwYDVR0TBAgwBgEB/wIBADAWBgNV\nHREEDzANggtleGFtcGxlLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAovxY5lm89Eod\nL8CH3dZzIH7nv8MXtwgpv2vth4PDq2btLS8xrqm2SsA/cV+DsbDjh5CxQLoDX+8V\ng8NtY+ipOE0hdJAUo7UVlsxuAY4frkmLL1/RwpjZg+Z2NAxpR7xGWgoMn7CH481w\nAOBypAuCxcfcyyAOttdS+YMRJnpL6z8/C3W0LGkNOs26Qhu1/U8lfz1f9F4XummD\nu2SCmJsAd1PrL1shsyh4HtmFjuY698aTjYUDUleAnx7ytrGlZuLOIeoQi7tcsLJJ\nTPMbxTLgGN2HEkdJerFRBNViuWvqioEyYlzZ3MshOCR2wsL4wrXrCF0Y3cNOYcIh\nZ8z+wUAP2g==\n-----END CERTIFICATE-----\n",
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
	_, err := client.Cdn.TrustedCaCertificates.List(context.TODO(), cdn.TrustedCaCertificateListParams{
		Automated:           gcore.Bool(true),
		ResourceID:          gcore.Int(0),
		ValidityNotAfterLte: gcore.String("validity_not_after_lte"),
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
	err := client.Cdn.TrustedCaCertificates.Delete(context.TODO(), 0)
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
	_, err := client.Cdn.TrustedCaCertificates.Get(context.TODO(), 0)
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
	_, err := client.Cdn.TrustedCaCertificates.Replace(
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
