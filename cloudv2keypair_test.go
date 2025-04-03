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

func TestCloudV2KeypairNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Keypairs.New(context.TODO(), gcore.CloudV2KeypairNewParams{
		ProjectID:       gcore.F(int64(1)),
		SshkeyName:      gcore.F("alice"),
		PublicKey:       gcore.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFHrnwGVBZs6q6vmTBzQFfzdRLQW8N6Rd0ogGe3h8tm83ZJLTTsF+1H4JcOvwI5ETkHMaFIWd2U15nHU5M7plE6UPRKfzy4rq6yI6cE4tojd3A9attMpbEEX7EbGKrbrb4AsjzxHKAVaREAb31ZplJkUlsiees25hTQXBcWQnOESlc9RCxZ/QQgNUUgqm7QGg7CNkL8Mpq9V4YaOhcFGWj0jXP1CL3g6Xe3xJo1CmUbkIOGUyAmrSfLEiy2O91iOUhbmYQyXksznNrT9O6uLkijf6syLZOdyAuUd/Z86eYXej4/YsvIA5eIFU4B6y9zOXEO2A81txPYMRAytYt7+e7 alice@alice"),
		SharedInProject: gcore.F(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2KeypairGet(t *testing.T) {
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
	_, err := client.Cloud.V2.Keypairs.Get(context.TODO(), "36a7a97a-0672-4911-8f2b-92cd4e5b0d91")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2KeypairListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Keypairs.List(context.TODO(), gcore.CloudV2KeypairListParams{
		ProjectID: gcore.F(int64(100)),
		UserID:    gcore.F(int64(5)),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV2KeypairDelete(t *testing.T) {
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
	err := client.Cloud.V2.Keypairs.Delete(context.TODO(), "36a7a97a-0672-4911-8f2b-92cd4e5b0d91")
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
