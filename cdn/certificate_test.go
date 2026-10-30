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

func TestCertificateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Certificates.New(context.TODO(), cdn.CertificateNewParams{
		OfOwnCertificate: &cdn.CertificateNewParamsBodyOwnCertificate{
			Name:           "Example 1",
			SslCertificate: "-----BEGIN CERTIFICATE-----\nMIIBbDCCAROgAwIBAgIUQJOOtM8WkOwi9fIMzL0RQttr/JQwCgYIKoZIzj0EAwIw\nKjETMBEGA1UECgwKRXhhbXBsZSBDQTETMBEGA1UEAwwKRXhhbXBsZSBDQTAeFw0y\nNjAxMDEwMDAwMDBaFw0yNzAxMDEwMDAwMDBaMBYxFDASBgNVBAMMC2V4YW1wbGUu\nY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFPjFz3rezjBUAKZXDcGcZKwl\npcCQmREWxCDWDh1D1PnWoEGmg/eaTpi0j2XsyLeLZYqwkxwcXobnic3cHo4kRKMr\nMCkwJwYDVR0RBCAwHoILZXhhbXBsZS5jb22CD3d3dy5leGFtcGxlLmNvbTAKBggq\nhkjOPQQDAgNHADBEAiALUtTb/Vd2mJb0Ig36VkIpWej3Ufux/EBCY7BNdDNQ4gIg\nfxZsfN96JxEXETedb1Grr1f/6GYPxUnGkKB4Rjj9aTc=\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIBbTCCARKgAwIBAgIULDeBmErpf2HbicPq8M/WK6fUZxEwCgYIKoZIzj0EAwIw\nKjETMBEGA1UECgwKRXhhbXBsZSBDQTETMBEGA1UEAwwKRXhhbXBsZSBDQTAeFw0y\nNjAxMDEwMDAwMDBaFw0zNjAxMDEwMDAwMDBaMCoxEzARBgNVBAoMCkV4YW1wbGUg\nQ0ExEzARBgNVBAMMCkV4YW1wbGUgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAATdtaj2FA1AUBTNxVPfCeuuJWY6D1SQTtCNT3vbakavfJu+QISYQcc0Tbtudl5s\nv6mQPTiVlKOzi5NSWYn9F8YwoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMAoGCCqG\nSM49BAMCA0kAMEYCIQD4GWDCyoYDx48b49W/R6ox5K2Dx37otAqQCSeB96moHAIh\nAKAUFyIRFL4HQ1Y3jEELs8TvC2554gcPQKitr0S0wM1g\n-----END CERTIFICATE-----\n",
			SslPrivateKey:  "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQguJC0fwVmXvhkj9fY\nCTAWhTxcAv4/tcUV+7Qhbvac5WGhRANCAAQU+MXPet7OMFQAplcNwZxkrCWlwJCZ\nERbEINYOHUPU+dagQaaD95pOmLSPZezIt4tlirCTHBxehueJzdwejiRE\n-----END PRIVATE KEY-----\n",
			Automated:      false,
			ValidateRootCa: gcore.Bool(true),
		},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Certificates.List(context.TODO(), cdn.CertificateListParams{
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

func TestCertificateDelete(t *testing.T) {
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
	err := client.CDN.Certificates.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateForceRetry(t *testing.T) {
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
	err := client.CDN.Certificates.ForceRetry(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateGet(t *testing.T) {
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
	_, err := client.CDN.Certificates.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateGetStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Certificates.GetStatus(
		context.TODO(),
		0,
		cdn.CertificateGetStatusParams{
			Exclude: []string{"string"},
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

func TestCertificateGetUsage(t *testing.T) {
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
	_, err := client.CDN.Certificates.GetUsage(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateRenew(t *testing.T) {
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
	err := client.CDN.Certificates.Renew(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCertificateReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Certificates.Replace(
		context.TODO(),
		0,
		cdn.CertificateReplaceParams{
			Name:           "Example 1",
			SslCertificate: "-----BEGIN CERTIFICATE-----\nMIIBbDCCAROgAwIBAgIUQJOOtM8WkOwi9fIMzL0RQttr/JQwCgYIKoZIzj0EAwIw\nKjETMBEGA1UECgwKRXhhbXBsZSBDQTETMBEGA1UEAwwKRXhhbXBsZSBDQTAeFw0y\nNjAxMDEwMDAwMDBaFw0yNzAxMDEwMDAwMDBaMBYxFDASBgNVBAMMC2V4YW1wbGUu\nY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFPjFz3rezjBUAKZXDcGcZKwl\npcCQmREWxCDWDh1D1PnWoEGmg/eaTpi0j2XsyLeLZYqwkxwcXobnic3cHo4kRKMr\nMCkwJwYDVR0RBCAwHoILZXhhbXBsZS5jb22CD3d3dy5leGFtcGxlLmNvbTAKBggq\nhkjOPQQDAgNHADBEAiALUtTb/Vd2mJb0Ig36VkIpWej3Ufux/EBCY7BNdDNQ4gIg\nfxZsfN96JxEXETedb1Grr1f/6GYPxUnGkKB4Rjj9aTc=\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIBbTCCARKgAwIBAgIULDeBmErpf2HbicPq8M/WK6fUZxEwCgYIKoZIzj0EAwIw\nKjETMBEGA1UECgwKRXhhbXBsZSBDQTETMBEGA1UEAwwKRXhhbXBsZSBDQTAeFw0y\nNjAxMDEwMDAwMDBaFw0zNjAxMDEwMDAwMDBaMCoxEzARBgNVBAoMCkV4YW1wbGUg\nQ0ExEzARBgNVBAMMCkV4YW1wbGUgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAATdtaj2FA1AUBTNxVPfCeuuJWY6D1SQTtCNT3vbakavfJu+QISYQcc0Tbtudl5s\nv6mQPTiVlKOzi5NSWYn9F8YwoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMAoGCCqG\nSM49BAMCA0kAMEYCIQD4GWDCyoYDx48b49W/R6ox5K2Dx37otAqQCSeB96moHAIh\nAKAUFyIRFL4HQ1Y3jEELs8TvC2554gcPQKitr0S0wM1g\n-----END CERTIFICATE-----\n",
			SslPrivateKey:  "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQguJC0fwVmXvhkj9fY\nCTAWhTxcAv4/tcUV+7Qhbvac5WGhRANCAAQU+MXPet7OMFQAplcNwZxkrCWlwJCZ\nERbEINYOHUPU+dagQaaD95pOmLSPZezIt4tlirCTHBxehueJzdwejiRE\n-----END PRIVATE KEY-----\n",
			Automated:      false,
			ValidateRootCa: gcore.Bool(true),
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
