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

func TestLogsUploaderTargetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.New(context.TODO(), cdn.LogsUploaderTargetNewParams{
		Config: cdn.LogsUploaderTargetNewParamsConfigUnion{
			OfS3GcoreConfig: &cdn.LogsUploaderTargetNewParamsConfigS3GcoreConfig{
				AccessKeyID:     "access_key_id",
				BucketName:      "bucket_name",
				Endpoint:        "endpoint",
				Region:          "region",
				SecretAccessKey: "secret_access_key",
				Directory:       gcore.String("directory"),
				UsePathStyle:    gcore.Bool(true),
			},
		},
		StorageType: cdn.LogsUploaderTargetNewParamsStorageTypeS3Gcore,
		Description: gcore.String("description"),
		Name:        gcore.String("name"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderTargetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.Update(
		context.TODO(),
		0,
		cdn.LogsUploaderTargetUpdateParams{
			Config: cdn.LogsUploaderTargetUpdateParamsConfigUnion{
				OfS3GcoreConfig: &cdn.LogsUploaderTargetUpdateParamsConfigS3GcoreConfig{
					AccessKeyID:     "access_key_id",
					BucketName:      "bucket_name",
					Endpoint:        "endpoint",
					Region:          "region",
					SecretAccessKey: "secret_access_key",
					Directory:       gcore.String("directory"),
					UsePathStyle:    gcore.Bool(true),
				},
			},
			Description: gcore.String("description"),
			Name:        gcore.String("name"),
			StorageType: cdn.LogsUploaderTargetUpdateParamsStorageTypeS3Gcore,
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

func TestLogsUploaderTargetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.List(context.TODO(), cdn.LogsUploaderTargetListParams{
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

func TestLogsUploaderTargetDelete(t *testing.T) {
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
	err := client.Cdn.LogsUploader.Targets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderTargetGet(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogsUploaderTargetReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.Replace(
		context.TODO(),
		0,
		cdn.LogsUploaderTargetReplaceParams{
			Config: cdn.LogsUploaderTargetReplaceParamsConfigUnion{
				OfS3GcoreConfig: &cdn.LogsUploaderTargetReplaceParamsConfigS3GcoreConfig{
					AccessKeyID:     "access_key_id",
					BucketName:      "bucket_name",
					Endpoint:        "endpoint",
					Region:          "region",
					SecretAccessKey: "secret_access_key",
					Directory:       gcore.String("directory"),
					UsePathStyle:    gcore.Bool(true),
				},
			},
			StorageType: cdn.LogsUploaderTargetReplaceParamsStorageTypeS3Gcore,
			Description: gcore.String("description"),
			Name:        gcore.String("name"),
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

func TestLogsUploaderTargetValidate(t *testing.T) {
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
	_, err := client.Cdn.LogsUploader.Targets.Validate(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
