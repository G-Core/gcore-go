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

func TestLogsUploaderPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.New(context.TODO(), cdn.LogsUploaderPolicyNewParams{
		DateFormat:              gcore.String("[02/Jan/2006:15:04:05 -0700]"),
		Description:             gcore.String("New policy"),
		EscapeSpecialCharacters: gcore.Bool(true),
		FieldDelimiter:          gcore.String(","),
		FieldSeparator:          gcore.String(";"),
		Fields:                  []string{"remote_addr", "status"},
		FileNameTemplate:        gcore.String("{{YYYY}}_{{MM}}_{{DD}}_{{HH}}_{{mm}}_{{ss}}_access.log.gz"),
		FormatType:              cdn.LogsUploaderPolicyNewParamsFormatTypeJson,
		IncludeEmptyLogs:        gcore.Bool(true),
		IncludeShieldLogs:       gcore.Bool(true),
		LogSampleRate:           gcore.Float(1),
		Name:                    gcore.String("Policy"),
		RetryIntervalMinutes:    gcore.Int(32),
		RotateIntervalMinutes:   gcore.Int(32),
		RotateThresholdLines:    gcore.Int(5000),
		RotateThresholdMB:       gcore.Int(252),
		Tags:                    map[string]string{},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.Update(
		context.TODO(),
		0,
		cdn.LogsUploaderPolicyUpdateParams{
			DateFormat:              gcore.String("[02/Jan/2006:15:04:05 -0700]"),
			Description:             gcore.String("New policy"),
			EscapeSpecialCharacters: gcore.Bool(true),
			FieldDelimiter:          gcore.String(","),
			FieldSeparator:          gcore.String(";"),
			Fields:                  []string{"remote_addr", "status"},
			FileNameTemplate:        gcore.String("{{YYYY}}_{{MM}}_{{DD}}_{{HH}}_{{mm}}_{{ss}}_access.log.gz"),
			FormatType:              cdn.LogsUploaderPolicyUpdateParamsFormatTypeJson,
			IncludeEmptyLogs:        gcore.Bool(true),
			IncludeShieldLogs:       gcore.Bool(true),
			LogSampleRate:           gcore.Float(0.5),
			Name:                    gcore.String("Policy"),
			RetryIntervalMinutes:    gcore.Int(32),
			RotateIntervalMinutes:   gcore.Int(32),
			RotateThresholdLines:    gcore.Int(5000),
			RotateThresholdMB:       gcore.Int(252),
			Tags:                    map[string]string{},
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

func TestLogsUploaderPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.List(context.TODO(), cdn.LogsUploaderPolicyListParams{
		ConfigIDs: []int64{0},
		Search:    gcore.String("search"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderPolicyDelete(t *testing.T) {
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
	err := client.CDN.LogsUploader.Policies.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderPolicyGet(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderPolicyListFields(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.ListFields(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderPolicyReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.LogsUploader.Policies.Replace(
		context.TODO(),
		0,
		cdn.LogsUploaderPolicyReplaceParams{
			DateFormat:              gcore.String("[02/Jan/2006:15:04:05 -0700]"),
			Description:             gcore.String("New policy"),
			EscapeSpecialCharacters: gcore.Bool(true),
			FieldDelimiter:          gcore.String(","),
			FieldSeparator:          gcore.String(";"),
			Fields:                  []string{"remote_addr", "status"},
			FileNameTemplate:        gcore.String("{{YYYY}}_{{MM}}_{{DD}}_{{HH}}_{{mm}}_{{ss}}_access.log.gz"),
			FormatType:              cdn.LogsUploaderPolicyReplaceParamsFormatTypeJson,
			IncludeEmptyLogs:        gcore.Bool(true),
			IncludeShieldLogs:       gcore.Bool(true),
			LogSampleRate:           gcore.Float(1),
			Name:                    gcore.String("Policy"),
			RetryIntervalMinutes:    gcore.Int(32),
			RotateIntervalMinutes:   gcore.Int(32),
			RotateThresholdLines:    gcore.Int(5000),
			RotateThresholdMB:       gcore.Int(252),
			Tags:                    map[string]string{},
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
