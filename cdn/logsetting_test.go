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

func TestLogSettingNewWithOptionalParams(t *testing.T) {
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
	err := client.Cdn.Logs.Settings.New(context.TODO(), cdn.LogSettingNewParams{
		AllResourcesBucket: "all_resources_bucket",
		AllResourcesFolder: "all_resources_folder",
		Folders: []cdn.LogSettingNewParamsFolder{{
			ID:          gcore.Int(0),
			Bucket:      gcore.String("bucket"),
			CdnResource: gcore.Int(0),
			Folder:      gcore.String("folder"),
		}},
		ForAllResources:   true,
		FtpHostname:       "ftp_hostname",
		FtpLogin:          "ftp_login",
		FtpPassword:       "ftp_password",
		S3AccessKeyID:     "s3_access_key_id",
		S3Hostname:        "s3_hostname",
		S3SecretKey:       "s3_secret_key",
		S3Type:            "s3_type",
		SftpHostname:      "sftp_hostname",
		SftpLogin:         "sftp_login",
		SftpPassword:      "sftp_password",
		StorageType:       "storage_type",
		ArchiveSizeMB:     gcore.Int(500),
		Enabled:           gcore.Bool(true),
		FtpPrependFolder:  gcore.String("ftp_prepend_folder"),
		IgnoreEmptyLogs:   gcore.Bool(true),
		S3AwsRegion:       gcore.Int(0),
		S3BucketLocation:  gcore.String("s3_bucket_location"),
		S3HostBucket:      gcore.String("s3_host_bucket"),
		SftpKeyPassphrase: gcore.String("sftp_key_passphrase"),
		SftpPrependFolder: gcore.String("sftp_prepend_folder"),
		SftpPrivateKey:    gcore.String("sftp_private_key"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogSettingUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Cdn.Logs.Settings.Update(context.TODO(), cdn.LogSettingUpdateParams{
		AllResourcesBucket: "all_resources_bucket",
		AllResourcesFolder: "all_resources_folder",
		Folders: []cdn.LogSettingUpdateParamsFolder{{
			ID:          gcore.Int(0),
			Bucket:      gcore.String("bucket"),
			CdnResource: gcore.Int(0),
			Folder:      gcore.String("folder"),
		}},
		ForAllResources:   true,
		FtpHostname:       "ftp_hostname",
		FtpLogin:          "ftp_login",
		FtpPassword:       "ftp_password",
		S3AccessKeyID:     "s3_access_key_id",
		S3Hostname:        "s3_hostname",
		S3SecretKey:       "s3_secret_key",
		S3Type:            "s3_type",
		SftpHostname:      "sftp_hostname",
		SftpLogin:         "sftp_login",
		SftpPassword:      "sftp_password",
		StorageType:       "storage_type",
		ArchiveSizeMB:     gcore.Int(500),
		Enabled:           gcore.Bool(true),
		FtpPrependFolder:  gcore.String("ftp_prepend_folder"),
		IgnoreEmptyLogs:   gcore.Bool(true),
		S3AwsRegion:       gcore.Int(0),
		S3BucketLocation:  gcore.String("s3_bucket_location"),
		S3HostBucket:      gcore.String("s3_host_bucket"),
		SftpKeyPassphrase: gcore.String("sftp_key_passphrase"),
		SftpPrependFolder: gcore.String("sftp_prepend_folder"),
		SftpPrivateKey:    gcore.String("sftp_private_key"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogSettingDelete(t *testing.T) {
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
	err := client.Cdn.Logs.Settings.Delete(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogSettingGet(t *testing.T) {
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
	_, err := client.Cdn.Logs.Settings.Get(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
