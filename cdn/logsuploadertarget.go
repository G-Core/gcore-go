// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// Logs uploader targets define destinations for CDN log delivery, such as S3
// buckets or SFTP servers, with associated authentication and configuration
// settings.
//
// LogsUploaderTargetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogsUploaderTargetService] method instead.
type LogsUploaderTargetService struct {
	Options []option.RequestOption
}

// NewLogsUploaderTargetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogsUploaderTargetService(opts ...option.RequestOption) (r LogsUploaderTargetService) {
	r = LogsUploaderTargetService{}
	r.Options = opts
	return
}

// Create logs uploader target.
func (r *LogsUploaderTargetService) New(ctx context.Context, body LogsUploaderTargetNewParams, opts ...option.RequestOption) (res *LogsUploaderTargetUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/targets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change logs uploader target partially.
func (r *LogsUploaderTargetService) Update(ctx context.Context, id int64, body LogsUploaderTargetUpdateParams, opts ...option.RequestOption) (res *LogsUploaderTargetUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get list of logs uploader targets.
func (r *LogsUploaderTargetService) List(ctx context.Context, query LogsUploaderTargetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogsUploaderTargetUnion], err error) {
	// CUSTOM CODE: CDN API only envelopes when limit>=1; default so List always paginates.
	if !query.Limit.Valid() {
		query.Limit = param.NewOpt[int64](1000)
	}
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/logs_uploader/targets"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get list of logs uploader targets.
func (r *LogsUploaderTargetService) ListAutoPaging(ctx context.Context, query LogsUploaderTargetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogsUploaderTargetUnion] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete the logs uploader target from the system permanently.
//
// Notes:
//
//   - **Irreversibility**: This action is irreversible. Once deleted, the logs
//     uploader target cannot be recovered.
func (r *LogsUploaderTargetService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get information about logs uploader target.
func (r *LogsUploaderTargetService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderTargetUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Change logs uploader target.
func (r *LogsUploaderTargetService) Replace(ctx context.Context, id int64, body LogsUploaderTargetReplaceParams, opts ...option.RequestOption) (res *LogsUploaderTargetUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Validate logs uploader target.
func (r *LogsUploaderTargetService) Validate(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v/validate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// LogsUploaderTargetUnion contains all possible properties and values from
// [LogsUploaderTargetS3Gcore], [LogsUploaderTargetS3Amazon],
// [LogsUploaderTargetS3Oss], [LogsUploaderTargetS3Other],
// [LogsUploaderTargetS3V1], [LogsUploaderTargetFtp], [LogsUploaderTargetSftp],
// [LogsUploaderTargetHTTP], [LogsUploaderTargetAzureBlob],
// [LogsUploaderTargetSls].
//
// Use the [LogsUploaderTargetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LogsUploaderTargetUnion struct {
	// Any of "s3_gcore", "s3_amazon", "s3_oss", "s3_other", "s3_v1", "ftp", "sftp",
	// "http", "azure_blob", "sls".
	StorageType string `json:"storage_type"`
	ID          int64  `json:"id"`
	ClientID    int64  `json:"client_id"`
	// This field is a union of [LogsUploaderTargetS3GcoreConfigUnion],
	// [LogsUploaderTargetS3AmazonConfig], [LogsUploaderTargetS3OssConfig],
	// [LogsUploaderTargetS3OtherConfig], [LogsUploaderTargetS3V1Config],
	// [LogsUploaderTargetFtpConfig], [LogsUploaderTargetSftpConfig],
	// [LogsUploaderTargetHTTPConfig], [LogsUploaderTargetAzureBlobConfig],
	// [LogsUploaderTargetSlsConfig]
	Config                 LogsUploaderTargetUnionConfig `json:"config"`
	Created                time.Time                     `json:"created"`
	Description            string                        `json:"description"`
	Name                   string                        `json:"name"`
	RelatedUploaderConfigs []int64                       `json:"related_uploader_configs"`
	// This field is a union of [LogsUploaderTargetS3GcoreStatus],
	// [LogsUploaderTargetS3AmazonStatus], [LogsUploaderTargetS3OssStatus],
	// [LogsUploaderTargetS3OtherStatus], [LogsUploaderTargetS3V1Status],
	// [LogsUploaderTargetFtpStatus], [LogsUploaderTargetSftpStatus],
	// [LogsUploaderTargetHTTPStatus], [LogsUploaderTargetAzureBlobStatus],
	// [LogsUploaderTargetSlsStatus]
	Status  LogsUploaderTargetUnionStatus `json:"status"`
	Updated time.Time                     `json:"updated"`
	JSON    struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		raw                    string
	} `json:"-"`
}

// anyLogsUploaderTarget is implemented by each variant of
// [LogsUploaderTargetUnion] to add type safety for the return type of
// [LogsUploaderTargetUnion.AsAny]
type anyLogsUploaderTarget interface {
	implLogsUploaderTargetUnion()
}

func (LogsUploaderTargetS3Gcore) implLogsUploaderTargetUnion()   {}
func (LogsUploaderTargetS3Amazon) implLogsUploaderTargetUnion()  {}
func (LogsUploaderTargetS3Oss) implLogsUploaderTargetUnion()     {}
func (LogsUploaderTargetS3Other) implLogsUploaderTargetUnion()   {}
func (LogsUploaderTargetS3V1) implLogsUploaderTargetUnion()      {}
func (LogsUploaderTargetFtp) implLogsUploaderTargetUnion()       {}
func (LogsUploaderTargetSftp) implLogsUploaderTargetUnion()      {}
func (LogsUploaderTargetHTTP) implLogsUploaderTargetUnion()      {}
func (LogsUploaderTargetAzureBlob) implLogsUploaderTargetUnion() {}
func (LogsUploaderTargetSls) implLogsUploaderTargetUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := LogsUploaderTargetUnion.AsAny().(type) {
//	case cdn.LogsUploaderTargetS3Gcore:
//	case cdn.LogsUploaderTargetS3Amazon:
//	case cdn.LogsUploaderTargetS3Oss:
//	case cdn.LogsUploaderTargetS3Other:
//	case cdn.LogsUploaderTargetS3V1:
//	case cdn.LogsUploaderTargetFtp:
//	case cdn.LogsUploaderTargetSftp:
//	case cdn.LogsUploaderTargetHTTP:
//	case cdn.LogsUploaderTargetAzureBlob:
//	case cdn.LogsUploaderTargetSls:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LogsUploaderTargetUnion) AsAny() anyLogsUploaderTarget {
	switch u.StorageType {
	case "s3_gcore":
		return u.AsS3Gcore()
	case "s3_amazon":
		return u.AsS3Amazon()
	case "s3_oss":
		return u.AsS3Oss()
	case "s3_other":
		return u.AsS3Other()
	case "s3_v1":
		return u.AsS3V1()
	case "ftp":
		return u.AsFtp()
	case "sftp":
		return u.AsSftp()
	case "http":
		return u.AsHTTP()
	case "azure_blob":
		return u.AsAzureBlob()
	case "sls":
		return u.AsSls()
	}
	return nil
}

func (u LogsUploaderTargetUnion) AsS3Gcore() (v LogsUploaderTargetS3Gcore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsS3Amazon() (v LogsUploaderTargetS3Amazon) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsS3Oss() (v LogsUploaderTargetS3Oss) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsS3Other() (v LogsUploaderTargetS3Other) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsS3V1() (v LogsUploaderTargetS3V1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsFtp() (v LogsUploaderTargetFtp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsSftp() (v LogsUploaderTargetSftp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsHTTP() (v LogsUploaderTargetHTTP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsAzureBlob() (v LogsUploaderTargetAzureBlob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetUnion) AsSls() (v LogsUploaderTargetSls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogsUploaderTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *LogsUploaderTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetUnionConfig is an implicit subunion of
// [LogsUploaderTargetUnion]. LogsUploaderTargetUnionConfig provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LogsUploaderTargetUnion].
type LogsUploaderTargetUnionConfig struct {
	AccessKeyID  string `json:"access_key_id"`
	BucketName   string `json:"bucket_name"`
	Directory    string `json:"directory"`
	Endpoint     string `json:"endpoint"`
	Region       string `json:"region"`
	UsePathStyle bool   `json:"use_path_style"`
	// This field is from variant [LogsUploaderTargetS3GcoreConfigUnion].
	StorageID      int64  `json:"storage_id"`
	Hostname       string `json:"hostname"`
	TimeoutSeconds int64  `json:"timeout_seconds"`
	User           string `json:"user"`
	// This field is from variant [LogsUploaderTargetSftpConfig].
	KeyPassphrase string `json:"key_passphrase"`
	// This field is from variant [LogsUploaderTargetSftpConfig].
	Password string `json:"password"`
	// This field is from variant [LogsUploaderTargetSftpConfig].
	PrivateKey string `json:"private_key"`
	// This field is from variant [LogsUploaderTargetHTTPConfig].
	Append LogsUploaderTargetHTTPConfigAppend `json:"append"`
	// This field is a union of [LogsUploaderTargetHTTPConfigAuth],
	// [LogsUploaderTargetAzureBlobConfigAuth], [LogsUploaderTargetSlsConfigAuth]
	Auth LogsUploaderTargetUnionConfigAuth `json:"auth"`
	// This field is from variant [LogsUploaderTargetHTTPConfig].
	ContentType string `json:"content_type"`
	// This field is from variant [LogsUploaderTargetHTTPConfig].
	Retry LogsUploaderTargetHTTPConfigRetry `json:"retry"`
	// This field is from variant [LogsUploaderTargetHTTPConfig].
	Upload LogsUploaderTargetHTTPConfigUpload `json:"upload"`
	// This field is from variant [LogsUploaderTargetAzureBlobConfig].
	AccountName string `json:"account_name"`
	// This field is from variant [LogsUploaderTargetAzureBlobConfig].
	ContainerName string `json:"container_name"`
	// This field is from variant [LogsUploaderTargetSlsConfig].
	LogStore string `json:"log_store"`
	// This field is from variant [LogsUploaderTargetSlsConfig].
	Project string `json:"project"`
	// This field is from variant [LogsUploaderTargetSlsConfig].
	Topic string `json:"topic"`
	JSON  struct {
		AccessKeyID    respjson.Field
		BucketName     respjson.Field
		Directory      respjson.Field
		Endpoint       respjson.Field
		Region         respjson.Field
		UsePathStyle   respjson.Field
		StorageID      respjson.Field
		Hostname       respjson.Field
		TimeoutSeconds respjson.Field
		User           respjson.Field
		KeyPassphrase  respjson.Field
		Password       respjson.Field
		PrivateKey     respjson.Field
		Append         respjson.Field
		Auth           respjson.Field
		ContentType    respjson.Field
		Retry          respjson.Field
		Upload         respjson.Field
		AccountName    respjson.Field
		ContainerName  respjson.Field
		LogStore       respjson.Field
		Project        respjson.Field
		Topic          respjson.Field
		raw            string
	} `json:"-"`
}

func (r *LogsUploaderTargetUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetUnionConfigAuth is an implicit subunion of
// [LogsUploaderTargetUnion]. LogsUploaderTargetUnionConfigAuth provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LogsUploaderTargetUnion].
type LogsUploaderTargetUnionConfigAuth struct {
	// This field is a union of [LogsUploaderTargetHTTPConfigAuthConfig],
	// [LogsUploaderTargetAzureBlobConfigAuthConfigUnion],
	// [LogsUploaderTargetSlsConfigAuthConfig]
	Config LogsUploaderTargetUnionConfigAuthConfig `json:"config"`
	Type   string                                  `json:"type"`
	JSON   struct {
		Config respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (r *LogsUploaderTargetUnionConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetUnionConfigAuthConfig is an implicit subunion of
// [LogsUploaderTargetUnion]. LogsUploaderTargetUnionConfigAuthConfig provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LogsUploaderTargetUnion].
type LogsUploaderTargetUnionConfigAuthConfig struct {
	Token string `json:"token"`
	// This field is from variant [LogsUploaderTargetHTTPConfigAuthConfig].
	HeaderName string `json:"header_name"`
	// This field is from variant [LogsUploaderTargetAzureBlobConfigAuthConfigUnion].
	AccountKey string `json:"account_key"`
	// This field is from variant [LogsUploaderTargetSlsConfigAuthConfig].
	AccessKeyID string `json:"access_key_id"`
	// This field is from variant [LogsUploaderTargetSlsConfigAuthConfig].
	SecretAccessKey string `json:"secret_access_key"`
	JSON            struct {
		Token           respjson.Field
		HeaderName      respjson.Field
		AccountKey      respjson.Field
		AccessKeyID     respjson.Field
		SecretAccessKey respjson.Field
		raw             string
	} `json:"-"`
}

func (r *LogsUploaderTargetUnionConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetUnionStatus is an implicit subunion of
// [LogsUploaderTargetUnion]. LogsUploaderTargetUnionStatus provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LogsUploaderTargetUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLogsUploaderTargetSlsStatus]
type LogsUploaderTargetUnionStatus struct {
	// This field will be present if the value is a [any] instead of an object.
	OfLogsUploaderTargetSlsStatus any `json:",inline"`
	// This field is from variant [LogsUploaderTargetS3GcoreStatus],
	// [LogsUploaderTargetS3AmazonStatus], [LogsUploaderTargetS3OssStatus],
	// [LogsUploaderTargetS3OtherStatus], [LogsUploaderTargetS3V1Status],
	// [LogsUploaderTargetFtpStatus], [LogsUploaderTargetSftpStatus],
	// [LogsUploaderTargetHTTPStatus], [LogsUploaderTargetAzureBlobStatus],
	// [LogsUploaderTargetSlsStatus].
	Code int64 `json:"code"`
	// This field is from variant [LogsUploaderTargetS3GcoreStatus],
	// [LogsUploaderTargetS3AmazonStatus], [LogsUploaderTargetS3OssStatus],
	// [LogsUploaderTargetS3OtherStatus], [LogsUploaderTargetS3V1Status],
	// [LogsUploaderTargetFtpStatus], [LogsUploaderTargetSftpStatus],
	// [LogsUploaderTargetHTTPStatus], [LogsUploaderTargetAzureBlobStatus],
	// [LogsUploaderTargetSlsStatus].
	Details string `json:"details"`
	// This field is from variant [LogsUploaderTargetS3GcoreStatus],
	// [LogsUploaderTargetS3AmazonStatus], [LogsUploaderTargetS3OssStatus],
	// [LogsUploaderTargetS3OtherStatus], [LogsUploaderTargetS3V1Status],
	// [LogsUploaderTargetFtpStatus], [LogsUploaderTargetSftpStatus],
	// [LogsUploaderTargetHTTPStatus], [LogsUploaderTargetAzureBlobStatus],
	// [LogsUploaderTargetSlsStatus].
	Status LogsUploaderValidationStatus `json:"status"`
	// This field is from variant [LogsUploaderTargetS3GcoreStatus],
	// [LogsUploaderTargetS3AmazonStatus], [LogsUploaderTargetS3OssStatus],
	// [LogsUploaderTargetS3OtherStatus], [LogsUploaderTargetS3V1Status],
	// [LogsUploaderTargetFtpStatus], [LogsUploaderTargetSftpStatus],
	// [LogsUploaderTargetHTTPStatus], [LogsUploaderTargetAzureBlobStatus],
	// [LogsUploaderTargetSlsStatus].
	Updated time.Time `json:"updated"`
	JSON    struct {
		OfLogsUploaderTargetSlsStatus respjson.Field
		Code                          respjson.Field
		Details                       respjson.Field
		Status                        respjson.Field
		Updated                       respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *LogsUploaderTargetUnionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3Gcore struct {
	// Type of storage for logs.
	StorageType constant.S3Gcore `json:"storage_type" default:"s3_gcore"`
	ID          int64            `json:"id"`
	// Client that owns the target.
	ClientID int64 `json:"client_id"`
	// Config for the `s3_gcore` storage type.
	Config LogsUploaderTargetS3GcoreConfigUnion `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetS3GcoreStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3Gcore) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3Gcore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetS3GcoreConfigUnion contains all possible properties and values
// from [LogsUploaderTargetS3GcoreConfigS3GcoreConfig],
// [LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LogsUploaderTargetS3GcoreConfigUnion struct {
	// This field is from variant [LogsUploaderTargetS3GcoreConfigS3GcoreConfig].
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory"`
	// This field is from variant [LogsUploaderTargetS3GcoreConfigS3GcoreConfig].
	Endpoint string `json:"endpoint"`
	// This field is from variant [LogsUploaderTargetS3GcoreConfigS3GcoreConfig].
	Region string `json:"region"`
	// This field is from variant [LogsUploaderTargetS3GcoreConfigS3GcoreConfig].
	UsePathStyle bool `json:"use_path_style"`
	// This field is from variant [LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig].
	StorageID int64 `json:"storage_id"`
	JSON      struct {
		AccessKeyID  respjson.Field
		BucketName   respjson.Field
		Directory    respjson.Field
		Endpoint     respjson.Field
		Region       respjson.Field
		UsePathStyle respjson.Field
		StorageID    respjson.Field
		raw          string
	} `json:"-"`
}

func (u LogsUploaderTargetS3GcoreConfigUnion) AsS3GcoreConfig() (v LogsUploaderTargetS3GcoreConfigS3GcoreConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetS3GcoreConfigUnion) AsS3GcoreBoundConfig() (v LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogsUploaderTargetS3GcoreConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *LogsUploaderTargetS3GcoreConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3GcoreConfigS3GcoreConfig struct {
	AccessKeyID  string `json:"access_key_id" api:"required"`
	BucketName   string `json:"bucket_name"`
	Directory    string `json:"directory" api:"nullable"`
	Endpoint     string `json:"endpoint"`
	Region       string `json:"region"`
	UsePathStyle bool   `json:"use_path_style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID  respjson.Field
		BucketName   respjson.Field
		Directory    respjson.Field
		Endpoint     respjson.Field
		Region       respjson.Field
		UsePathStyle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3GcoreConfigS3GcoreConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3GcoreConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A target bound to your Gcore Object Storage. The CDN manages the connection, so
// its settings are not returned.
type LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig struct {
	BucketName string `json:"bucket_name" api:"required"`
	// ID of your Gcore Object Storage.
	StorageID int64  `json:"storage_id" api:"required"`
	Directory string `json:"directory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketName  respjson.Field
		StorageID   respjson.Field
		Directory   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3GcoreConfigS3GcoreBoundConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetS3GcoreStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3GcoreStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3GcoreStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3Amazon struct {
	// Type of storage for logs.
	StorageType constant.S3Amazon `json:"storage_type" default:"s3_amazon"`
	ID          int64             `json:"id"`
	// Client that owns the target.
	ClientID int64                            `json:"client_id"`
	Config   LogsUploaderTargetS3AmazonConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetS3AmazonStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3Amazon) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3Amazon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3AmazonConfig struct {
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory" api:"nullable"`
	Region      string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID respjson.Field
		BucketName  respjson.Field
		Directory   respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3AmazonConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetS3AmazonStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3AmazonStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3AmazonStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3Oss struct {
	// Type of storage for logs.
	StorageType constant.S3Oss `json:"storage_type" default:"s3_oss"`
	ID          int64          `json:"id"`
	// Client that owns the target.
	ClientID int64                         `json:"client_id"`
	Config   LogsUploaderTargetS3OssConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetS3OssStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3Oss) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3Oss) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3OssConfig struct {
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory" api:"nullable"`
	Endpoint    string `json:"endpoint" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID respjson.Field
		BucketName  respjson.Field
		Directory   respjson.Field
		Endpoint    respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3OssConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetS3OssStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3OssStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3OssStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3Other struct {
	// Type of storage for logs.
	StorageType constant.S3Other `json:"storage_type" default:"s3_other"`
	ID          int64            `json:"id"`
	// Client that owns the target.
	ClientID int64                           `json:"client_id"`
	Config   LogsUploaderTargetS3OtherConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetS3OtherStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3Other) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3Other) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3OtherConfig struct {
	AccessKeyID  string `json:"access_key_id" api:"required"`
	BucketName   string `json:"bucket_name"`
	Directory    string `json:"directory" api:"nullable"`
	Endpoint     string `json:"endpoint"`
	Region       string `json:"region"`
	UsePathStyle bool   `json:"use_path_style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID  respjson.Field
		BucketName   respjson.Field
		Directory    respjson.Field
		Endpoint     respjson.Field
		Region       respjson.Field
		UsePathStyle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3OtherConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetS3OtherStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3OtherStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3OtherStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3V1 struct {
	// Type of storage for logs.
	StorageType constant.S3V1 `json:"storage_type" default:"s3_v1"`
	ID          int64         `json:"id"`
	// Client that owns the target.
	ClientID int64                        `json:"client_id"`
	Config   LogsUploaderTargetS3V1Config `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetS3V1Status `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3V1) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3V1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetS3V1Config struct {
	AccessKeyID  string `json:"access_key_id" api:"required"`
	BucketName   string `json:"bucket_name"`
	Directory    string `json:"directory" api:"nullable"`
	Endpoint     string `json:"endpoint"`
	Region       string `json:"region"`
	UsePathStyle bool   `json:"use_path_style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID  respjson.Field
		BucketName   respjson.Field
		Directory    respjson.Field
		Endpoint     respjson.Field
		Region       respjson.Field
		UsePathStyle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3V1Config) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetS3V1Status struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetS3V1Status) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetS3V1Status) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetFtp struct {
	// Type of storage for logs.
	StorageType constant.Ftp `json:"storage_type" default:"ftp"`
	ID          int64        `json:"id"`
	// Client that owns the target.
	ClientID int64                       `json:"client_id"`
	Config   LogsUploaderTargetFtpConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetFtpStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetFtp) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetFtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetFtpConfig struct {
	Directory      string `json:"directory" api:"nullable"`
	Hostname       string `json:"hostname"`
	TimeoutSeconds int64  `json:"timeout_seconds"`
	User           string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Directory      respjson.Field
		Hostname       respjson.Field
		TimeoutSeconds respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetFtpConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetFtpStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetFtpStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetFtpStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetSftp struct {
	// Type of storage for logs.
	StorageType constant.Sftp `json:"storage_type" default:"sftp"`
	ID          int64         `json:"id"`
	// Client that owns the target.
	ClientID int64                        `json:"client_id"`
	Config   LogsUploaderTargetSftpConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetSftpStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSftp) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSftp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetSftpConfig struct {
	Directory      string `json:"directory" api:"nullable"`
	Hostname       string `json:"hostname"`
	KeyPassphrase  string `json:"key_passphrase" api:"nullable"`
	Password       string `json:"password" api:"nullable"`
	PrivateKey     string `json:"private_key" api:"nullable"`
	TimeoutSeconds int64  `json:"timeout_seconds"`
	User           string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Directory      respjson.Field
		Hostname       respjson.Field
		KeyPassphrase  respjson.Field
		Password       respjson.Field
		PrivateKey     respjson.Field
		TimeoutSeconds respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSftpConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetSftpStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSftpStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSftpStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTP struct {
	// Type of storage for logs.
	StorageType constant.HTTP `json:"storage_type" default:"http"`
	ID          int64         `json:"id"`
	// Client that owns the target.
	ClientID int64                        `json:"client_id"`
	Config   LogsUploaderTargetHTTPConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetHTTPStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTP) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfig struct {
	Append LogsUploaderTargetHTTPConfigAppend `json:"append" api:"nullable"`
	Auth   LogsUploaderTargetHTTPConfigAuth   `json:"auth" api:"nullable"`
	// Any of "json", "text".
	ContentType string                             `json:"content_type"`
	Retry       LogsUploaderTargetHTTPConfigRetry  `json:"retry" api:"nullable"`
	Upload      LogsUploaderTargetHTTPConfigUpload `json:"upload"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Append      respjson.Field
		Auth        respjson.Field
		ContentType respjson.Field
		Retry       respjson.Field
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigAppend struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                             `json:"method"`
	ResponseActions []LogsUploaderTargetHTTPConfigAppendResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                              `json:"timeout_seconds"`
	URL             string                                             `json:"url"`
	UseCompression  bool                                               `json:"use_compression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers         respjson.Field
		Method          respjson.Field
		ResponseActions respjson.Field
		TimeoutSeconds  respjson.Field
		URL             respjson.Field
		UseCompression  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigAppend) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string `json:"action"`
	Description     string `json:"description"`
	MatchPayload    string `json:"match_payload"`
	MatchStatusCode int64  `json:"match_status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		Description     respjson.Field
		MatchPayload    respjson.Field
		MatchStatusCode respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigAppendResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigAuth struct {
	Config LogsUploaderTargetHTTPConfigAuthConfig `json:"config"`
	// Any of "token".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigAuth) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigAuthConfig struct {
	Token      string `json:"token"`
	HeaderName string `json:"header_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		HeaderName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigAuthConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigRetry struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                            `json:"method"`
	ResponseActions []LogsUploaderTargetHTTPConfigRetryResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                             `json:"timeout_seconds"`
	URL             string                                            `json:"url"`
	UseCompression  bool                                              `json:"use_compression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers         respjson.Field
		Method          respjson.Field
		ResponseActions respjson.Field
		TimeoutSeconds  respjson.Field
		URL             respjson.Field
		UseCompression  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigRetry) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string `json:"action"`
	Description     string `json:"description"`
	MatchPayload    string `json:"match_payload"`
	MatchStatusCode int64  `json:"match_status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		Description     respjson.Field
		MatchPayload    respjson.Field
		MatchStatusCode respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigRetryResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigUpload struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                             `json:"method"`
	ResponseActions []LogsUploaderTargetHTTPConfigUploadResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                              `json:"timeout_seconds"`
	URL             string                                             `json:"url"`
	UseCompression  bool                                               `json:"use_compression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers         respjson.Field
		Method          respjson.Field
		ResponseActions respjson.Field
		TimeoutSeconds  respjson.Field
		URL             respjson.Field
		UseCompression  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigUpload) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string `json:"action"`
	Description     string `json:"description"`
	MatchPayload    string `json:"match_payload"`
	MatchStatusCode int64  `json:"match_status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		Description     respjson.Field
		MatchPayload    respjson.Field
		MatchStatusCode respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPConfigUploadResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetHTTPStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetHTTPStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetHTTPStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetAzureBlob struct {
	// Type of storage for logs.
	StorageType constant.AzureBlob `json:"storage_type" default:"azure_blob"`
	ID          int64              `json:"id"`
	// Client that owns the target.
	ClientID int64                             `json:"client_id"`
	Config   LogsUploaderTargetAzureBlobConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetAzureBlobStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetAzureBlobConfig struct {
	// Azure Blob Storage account name.
	AccountName string                                `json:"account_name"`
	Auth        LogsUploaderTargetAzureBlobConfigAuth `json:"auth"`
	// Azure Blob Storage container name.
	ContainerName string `json:"container_name"`
	// Directory path within the container.
	Directory string `json:"directory" api:"nullable"`
	// Custom Azure Blob Storage endpoint URL.
	Endpoint string `json:"endpoint" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		Auth          respjson.Field
		ContainerName respjson.Field
		Directory     respjson.Field
		Endpoint      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlobConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlobConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetAzureBlobConfigAuth struct {
	// Authentication credentials. Secret fields are masked with '**\***'.
	Config LogsUploaderTargetAzureBlobConfigAuthConfigUnion `json:"config" api:"required"`
	// Authentication type.
	//
	// Any of "shared_key", "sas_token".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlobConfigAuth) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlobConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetAzureBlobConfigAuthConfigUnion contains all possible
// properties and values from
// [LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey],
// [LogsUploaderTargetAzureBlobConfigAuthConfigToken].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LogsUploaderTargetAzureBlobConfigAuthConfigUnion struct {
	// This field is from variant
	// [LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey].
	AccountKey string `json:"account_key"`
	// This field is from variant [LogsUploaderTargetAzureBlobConfigAuthConfigToken].
	Token string `json:"token"`
	JSON  struct {
		AccountKey respjson.Field
		Token      respjson.Field
		raw        string
	} `json:"-"`
}

func (u LogsUploaderTargetAzureBlobConfigAuthConfigUnion) AsLogsUploaderTargetAzureBlobConfigAuthConfigAccountKey() (v LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetAzureBlobConfigAuthConfigUnion) AsLogsUploaderTargetAzureBlobConfigAuthConfigToken() (v LogsUploaderTargetAzureBlobConfigAuthConfigToken) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogsUploaderTargetAzureBlobConfigAuthConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *LogsUploaderTargetAzureBlobConfigAuthConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey struct {
	// Masked secret value.
	AccountKey string `json:"account_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountKey  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlobConfigAuthConfigAccountKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetAzureBlobConfigAuthConfigToken struct {
	// Masked secret value.
	Token string `json:"token" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlobConfigAuthConfigToken) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlobConfigAuthConfigToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetAzureBlobStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetAzureBlobStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetAzureBlobStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetSls struct {
	// Type of storage for logs.
	StorageType constant.Sls `json:"storage_type" default:"sls"`
	ID          int64        `json:"id"`
	// Client that owns the target.
	ClientID int64                       `json:"client_id"`
	Config   LogsUploaderTargetSlsConfig `json:"config"`
	// Time when logs uploader target was created.
	Created time.Time `json:"created" format:"date-time"`
	// Description of the target.
	Description string `json:"description"`
	// Name of the target.
	Name string `json:"name"`
	// List of logs uploader configs that use this target.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Validation status of the logs uploader target. Informs if the specified target
	// is reachable.
	Status LogsUploaderTargetSlsStatus `json:"status"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageType            respjson.Field
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSls) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetSlsConfig struct {
	Auth LogsUploaderTargetSlsConfigAuth `json:"auth"`
	// SLS endpoint. Optional — derived from the region as `{region}.log.aliyuncs.com`
	// when omitted.
	Endpoint string `json:"endpoint" api:"nullable"`
	// SLS logstore name. 3-36 characters; lowercase letters, digits, hyphens, and
	// underscores.
	LogStore string `json:"log_store"`
	// SLS project name. 3-63 characters; lowercase letters, digits, and hyphens.
	Project string `json:"project"`
	// SLS region (e.g. `eu-central-1`).
	Region string `json:"region"`
	// Optional SLS topic (0-128 characters).
	Topic string `json:"topic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Auth        respjson.Field
		Endpoint    respjson.Field
		LogStore    respjson.Field
		Project     respjson.Field
		Region      respjson.Field
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSlsConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSlsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetSlsConfigAuth struct {
	// Authentication credentials. Secret fields are masked with '**\***'.
	Config LogsUploaderTargetSlsConfigAuthConfig `json:"config" api:"required"`
	// Authentication type.
	//
	// Any of "ak_sk".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSlsConfigAuth) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSlsConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authentication credentials. Secret fields are masked with '**\***'.
type LogsUploaderTargetSlsConfigAuthConfig struct {
	// Alibaba access key ID.
	AccessKeyID string `json:"access_key_id"`
	// Masked secret value.
	SecretAccessKey string `json:"secret_access_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKeyID     respjson.Field
		SecretAccessKey respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSlsConfigAuthConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSlsConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetSlsStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetSlsStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetSlsStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetList struct {
	// Total number of items.
	Count int64 `json:"count" api:"required"`
	// URL to the next page of results. Null if current page is the last one.
	Next string `json:"next" api:"required"`
	// URL to the previous page of results. Null if current page is the first one.
	Previous string                    `json:"previous" api:"required"`
	Results  []LogsUploaderTargetUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetList) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3Gcore *LogsUploaderTargetNewParamsBodyS3Gcore `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Amazon *LogsUploaderTargetNewParamsBodyS3Amazon `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Oss *LogsUploaderTargetNewParamsBodyS3Oss `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Other *LogsUploaderTargetNewParamsBodyS3Other `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3V1 *LogsUploaderTargetNewParamsBodyS3V1 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfFtp *LogsUploaderTargetNewParamsBodyFtp `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSftp *LogsUploaderTargetNewParamsBodySftp `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfHTTP *LogsUploaderTargetNewParamsBodyHTTP `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *LogsUploaderTargetNewParamsBodyAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSls *LogsUploaderTargetNewParamsBodySls `json:",inline"`

	paramObj
}

func (u LogsUploaderTargetNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3Gcore,
		u.OfS3Amazon,
		u.OfS3Oss,
		u.OfS3Other,
		u.OfS3V1,
		u.OfFtp,
		u.OfSftp,
		u.OfHTTP,
		u.OfAzureBlob,
		u.OfSls)
}
func (r *LogsUploaderTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyS3Gcore struct {
	// Config for the `s3_gcore` storage type.
	Config LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_gcore".
	StorageType constant.S3Gcore `json:"storage_type" default:"s3_gcore"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3Gcore) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3Gcore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3Gcore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion struct {
	OfS3GcoreConfig      *LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreConfig      `json:",omitzero,inline"`
	OfS3GcoreBoundConfig *LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreBoundConfig `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig, u.OfS3GcoreBoundConfig)
}
func (u *LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
	} else if !param.IsOmitted(u.OfS3GcoreBoundConfig) {
		return u.OfS3GcoreBoundConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetUsePathStyle() *bool {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.UsePathStyle.Valid() {
		return &vt.UsePathStyle.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetStorageID() *int64 {
	if vt := u.OfS3GcoreBoundConfig; vt != nil {
		return &vt.StorageID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil {
		return (*string)(&vt.BucketName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsBodyS3GcoreConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	}
	return nil
}

type LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deliver logs to a bucket in your Gcore Object Storage. Select the storage and
// bucket, and the CDN configures the connection for you.
//
// The properties BucketName, StorageID are required.
type LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreBoundConfig struct {
	// Name of a bucket in the selected storage.
	BucketName string `json:"bucket_name" api:"required"`
	// ID of your Gcore Object Storage. The storage must be active and of the Standard
	// type.
	StorageID int64             `json:"storage_id" api:"required"`
	Directory param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreBoundConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreBoundConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3GcoreConfigS3GcoreBoundConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyS3Amazon struct {
	Config LogsUploaderTargetNewParamsBodyS3AmazonConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_amazon".
	StorageType constant.S3Amazon `json:"storage_type" default:"s3_amazon"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3Amazon) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3Amazon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3Amazon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetNewParamsBodyS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	Region          string            `json:"region" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3AmazonConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3AmazonConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyS3Oss struct {
	Config LogsUploaderTargetNewParamsBodyS3OssConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_oss".
	StorageType constant.S3Oss `json:"storage_type" default:"s3_oss"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3Oss) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3Oss
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3Oss) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, SecretAccessKey are required.
type LogsUploaderTargetNewParamsBodyS3OssConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	Region   param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyS3Other struct {
	Config LogsUploaderTargetNewParamsBodyS3OtherConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_other".
	StorageType constant.S3Other `json:"storage_type" default:"s3_other"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3Other) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3Other
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3Other) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodyS3OtherConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3OtherConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3OtherConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyS3V1 struct {
	Config LogsUploaderTargetNewParamsBodyS3V1Config `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_v1".
	StorageType constant.S3V1 `json:"storage_type" default:"s3_v1"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3V1) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3V1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3V1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodyS3V1Config struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyS3V1Config) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyS3V1Config
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyFtp struct {
	Config LogsUploaderTargetNewParamsBodyFtpConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "ftp".
	StorageType constant.Ftp `json:"storage_type" default:"ftp"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyFtp) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyFtp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyFtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodyFtpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodySftp struct {
	Config LogsUploaderTargetNewParamsBodySftpConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "sftp".
	StorageType constant.Sftp `json:"storage_type" default:"sftp"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySftp) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySftp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySftp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodySftpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySftpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySftpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyHTTP struct {
	Config LogsUploaderTargetNewParamsBodyHTTPConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "http".
	StorageType constant.HTTP `json:"storage_type" default:"http"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTP) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Upload is required.
type LogsUploaderTargetNewParamsBodyHTTPConfig struct {
	Upload LogsUploaderTargetNewParamsBodyHTTPConfigUpload `json:"upload,omitzero" api:"required"`
	Append LogsUploaderTargetNewParamsBodyHTTPConfigAppend `json:"append,omitzero"`
	Auth   LogsUploaderTargetNewParamsBodyHTTPConfigAuth   `json:"auth,omitzero"`
	// Any of "json", "text".
	ContentType string                                         `json:"content_type,omitzero"`
	Retry       LogsUploaderTargetNewParamsBodyHTTPConfigRetry `json:"retry,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfig](
		"content_type", "json", "text",
	)
}

// The property URL is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigUpload struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                          `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigUpload) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigUpload](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigUploadResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The property URL is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigAppend struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                          `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigAppend) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigAppend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigAppend](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigAppendResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, Type are required.
type LogsUploaderTargetNewParamsBodyHTTPConfigAuth struct {
	Config LogsUploaderTargetNewParamsBodyHTTPConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Any of "token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigAuth](
		"type", "token",
	)
}

// The properties Token, HeaderName are required.
type LogsUploaderTargetNewParamsBodyHTTPConfigAuthConfig struct {
	Token      string `json:"token" api:"required"`
	HeaderName string `json:"header_name" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigRetry struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                         `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigRetry) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigRetry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigRetry](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyHTTPConfigRetryResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodyAzureBlob struct {
	Config LogsUploaderTargetNewParamsBodyAzureBlobConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "azure_blob".
	StorageType constant.AzureBlob `json:"storage_type" default:"azure_blob"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodyAzureBlobConfig struct {
	// Directory path within the container.
	Directory param.Opt[string] `json:"directory,omitzero"`
	// Custom Azure Blob Storage endpoint URL.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Azure Blob Storage account name.
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// Azure Blob Storage container name.
	ContainerName param.Opt[string]                                  `json:"container_name,omitzero"`
	Auth          LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyAzureBlobConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyAzureBlobConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyAzureBlobConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigUnion `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "shared_key", "sas_token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodyAzureBlobConfigAuth](
		"type", "shared_key", "sas_token",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigUnion struct {
	OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigAccountKey *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigAccountKey `json:",omitzero,inline"`
	OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigToken      *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigToken      `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigAccountKey, u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigToken)
}
func (u *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigAccountKey) {
		return u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigAccountKey
	} else if !param.IsOmitted(u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigToken) {
		return u.OfLogsUploaderTargetNewsBodyAzureBlobConfigAuthConfigToken
	}
	return nil
}

// The property AccountKey is required.
type LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigAccountKey struct {
	// Azure Blob Storage account key.
	AccountKey string `json:"account_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigAccountKey) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigAccountKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigAccountKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Token is required.
type LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigToken struct {
	// Azure Blob Storage SAS token.
	Token string `json:"token" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigToken) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodyAzureBlobConfigAuthConfigToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetNewParamsBodySls struct {
	Config LogsUploaderTargetNewParamsBodySlsConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "sls".
	StorageType constant.Sls `json:"storage_type" default:"sls"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySls) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetNewParamsBodySlsConfig struct {
	// SLS endpoint. Optional — derived from the region as `{region}.log.aliyuncs.com`
	// when omitted.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Optional SLS topic (0-128 characters).
	Topic param.Opt[string] `json:"topic,omitzero"`
	// SLS logstore name. 3-36 characters; lowercase letters, digits, hyphens, and
	// underscores.
	LogStore param.Opt[string] `json:"log_store,omitzero"`
	// SLS project name. 3-63 characters; lowercase letters, digits, and hyphens.
	Project param.Opt[string] `json:"project,omitzero"`
	// SLS region (e.g. `eu-central-1`).
	Region param.Opt[string]                            `json:"region,omitzero"`
	Auth   LogsUploaderTargetNewParamsBodySlsConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySlsConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySlsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySlsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetNewParamsBodySlsConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetNewParamsBodySlsConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "ak_sk".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySlsConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySlsConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySlsConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsBodySlsConfigAuth](
		"type", "ak_sk",
	)
}

// Authentication credentials.
//
// The properties AccessKeyID, SecretAccessKey are required.
type LogsUploaderTargetNewParamsBodySlsConfigAuthConfig struct {
	// Alibaba access key ID.
	AccessKeyID string `json:"access_key_id" api:"required"`
	// Alibaba secret access key.
	SecretAccessKey string `json:"secret_access_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsBodySlsConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsBodySlsConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsBodySlsConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParams struct {
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Config for specific storage type.
	Config LogsUploaderTargetUpdateParamsConfigUnion `json:"config,omitzero"`
	// Type of storage for logs.
	//
	// Any of "s3_gcore", "s3_amazon", "s3_oss", "s3_other", "s3_v1", "ftp", "sftp",
	// "http", "azure_blob", "sls".
	StorageType LogsUploaderTargetUpdateParamsStorageType `json:"storage_type,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetUpdateParamsConfigUnion struct {
	OfS3GcoreConfig      *LogsUploaderTargetUpdateParamsConfigS3GcoreConfig      `json:",omitzero,inline"`
	OfS3GcoreBoundConfig *LogsUploaderTargetUpdateParamsConfigS3GcoreBoundConfig `json:",omitzero,inline"`
	OfS3AmazonConfig     *LogsUploaderTargetUpdateParamsConfigS3AmazonConfig     `json:",omitzero,inline"`
	OfS3OssConfig        *LogsUploaderTargetUpdateParamsConfigS3OssConfig        `json:",omitzero,inline"`
	OfS3OtherConfig      *LogsUploaderTargetUpdateParamsConfigS3OtherConfig      `json:",omitzero,inline"`
	OfS3V1Config         *LogsUploaderTargetUpdateParamsConfigS3V1Config         `json:",omitzero,inline"`
	OfFtpConfig          *LogsUploaderTargetUpdateParamsConfigFtpConfig          `json:",omitzero,inline"`
	OfSftpConfig         *LogsUploaderTargetUpdateParamsConfigSftpConfig         `json:",omitzero,inline"`
	OfHTTPConfig         *LogsUploaderTargetUpdateParamsConfigHTTPConfig         `json:",omitzero,inline"`
	OfAzureBlobConfig    *LogsUploaderTargetUpdateParamsConfigAzureBlobConfig    `json:",omitzero,inline"`
	OfSlsConfig          *LogsUploaderTargetUpdateParamsConfigSlsConfig          `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetUpdateParamsConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig,
		u.OfS3GcoreBoundConfig,
		u.OfS3AmazonConfig,
		u.OfS3OssConfig,
		u.OfS3OtherConfig,
		u.OfS3V1Config,
		u.OfFtpConfig,
		u.OfSftpConfig,
		u.OfHTTPConfig,
		u.OfAzureBlobConfig,
		u.OfSlsConfig)
}
func (u *LogsUploaderTargetUpdateParamsConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetUpdateParamsConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
	} else if !param.IsOmitted(u.OfS3GcoreBoundConfig) {
		return u.OfS3GcoreBoundConfig
	} else if !param.IsOmitted(u.OfS3AmazonConfig) {
		return u.OfS3AmazonConfig
	} else if !param.IsOmitted(u.OfS3OssConfig) {
		return u.OfS3OssConfig
	} else if !param.IsOmitted(u.OfS3OtherConfig) {
		return u.OfS3OtherConfig
	} else if !param.IsOmitted(u.OfS3V1Config) {
		return u.OfS3V1Config
	} else if !param.IsOmitted(u.OfFtpConfig) {
		return u.OfFtpConfig
	} else if !param.IsOmitted(u.OfSftpConfig) {
		return u.OfSftpConfig
	} else if !param.IsOmitted(u.OfHTTPConfig) {
		return u.OfHTTPConfig
	} else if !param.IsOmitted(u.OfAzureBlobConfig) {
		return u.OfAzureBlobConfig
	} else if !param.IsOmitted(u.OfSlsConfig) {
		return u.OfSlsConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetStorageID() *int64 {
	if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.StorageID.Valid() {
		return &vt.StorageID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetKeyPassphrase() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.KeyPassphrase.Valid() {
		return &vt.KeyPassphrase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetPrivateKey() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.PrivateKey.Valid() {
		return &vt.PrivateKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetUpload() *LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Upload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAppend() *LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Append
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetContentType() *string {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.ContentType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetRetry() *LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Retry
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAccountName() *string {
	if vt := u.OfAzureBlobConfig; vt != nil && vt.AccountName.Valid() {
		return &vt.AccountName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetContainerName() *string {
	if vt := u.OfAzureBlobConfig; vt != nil && vt.ContainerName.Valid() {
		return &vt.ContainerName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetLogStore() *string {
	if vt := u.OfSlsConfig; vt != nil && vt.LogStore.Valid() {
		return &vt.LogStore.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetProject() *string {
	if vt := u.OfSlsConfig; vt != nil && vt.Project.Valid() {
		return &vt.Project.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetTopic() *string {
	if vt := u.OfSlsConfig; vt != nil && vt.Topic.Valid() {
		return &vt.Topic.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3AmazonConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfFtpConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfAzureBlobConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	} else if vt := u.OfAzureBlobConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	} else if vt := u.OfSlsConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfSlsConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetUsePathStyle() *bool {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.UsePathStyle.Valid() {
		return &vt.UsePathStyle.Value
	} else if vt := u.OfS3OtherConfig; vt != nil && vt.UsePathStyle.Valid() {
		return &vt.UsePathStyle.Value
	} else if vt := u.OfS3V1Config; vt != nil && vt.UsePathStyle.Valid() {
		return &vt.UsePathStyle.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetHostname() *string {
	if vt := u.OfFtpConfig; vt != nil && vt.Hostname.Valid() {
		return &vt.Hostname.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.Hostname.Valid() {
		return &vt.Hostname.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetPassword() *string {
	if vt := u.OfFtpConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetTimeoutSeconds() *int64 {
	if vt := u.OfFtpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetUser() *string {
	if vt := u.OfFtpConfig; vt != nil && vt.User.Valid() {
		return &vt.User.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.User.Valid() {
		return &vt.User.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAuth() (res logsUploaderTargetUpdateParamsConfigUnionAuth) {
	if vt := u.OfHTTPConfig; vt != nil {
		res.any = &vt.Auth
	} else if vt := u.OfAzureBlobConfig; vt != nil {
		res.any = &vt.Auth
	} else if vt := u.OfSlsConfig; vt != nil {
		res.any = &vt.Auth
	}
	return
}

// Can have the runtime types
// [*LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth],
// [*LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth],
// [*LogsUploaderTargetUpdateParamsConfigSlsConfigAuth]
type logsUploaderTargetUpdateParamsConfigUnionAuth struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cdn.LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth:
//	case *cdn.LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth:
//	case *cdn.LogsUploaderTargetUpdateParamsConfigSlsConfigAuth:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u logsUploaderTargetUpdateParamsConfigUnionAuth) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuth) GetType() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth:
		return (*string)(&vt.Type)
	case *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth:
		return (*string)(&vt.Type)
	case *LogsUploaderTargetUpdateParamsConfigSlsConfigAuth:
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u logsUploaderTargetUpdateParamsConfigUnionAuth) GetConfig() (res logsUploaderTargetUpdateParamsConfigUnionAuthConfig) {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth:
		res.any = &vt.Config
	case *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth:
		res.any = vt.Config
	case *LogsUploaderTargetUpdateParamsConfigSlsConfigAuth:
		res.any = &vt.Config
	}
	return res
}

// Can have the runtime types
// [*LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig],
// [*LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey],
// [*LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken],
// [*LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig]
type logsUploaderTargetUpdateParamsConfigUnionAuthConfig struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cdn.LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig:
//	case *cdn.LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey:
//	case *cdn.LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken:
//	case *cdn.LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) GetHeaderName() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig:
		return &vt.HeaderName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) GetAccountKey() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion:
		return vt.GetAccountKey()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) GetAccessKeyID() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig:
		return &vt.AccessKeyID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) GetSecretAccessKey() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig:
		return &vt.SecretAccessKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u logsUploaderTargetUpdateParamsConfigUnionAuthConfig) GetToken() *string {
	switch vt := u.any.(type) {
	case *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig:
		return (*string)(&vt.Token)
	case *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion:
		return vt.GetToken()
	}
	return nil
}

type LogsUploaderTargetUpdateParamsConfigS3GcoreConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3GcoreConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3GcoreConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bind a target to your Gcore Object Storage, or change its bucket or storage. To
// bind a target that is not bound yet, send both `storage_id` and `bucket_name`,
// and set `storage_type` to `s3_gcore` if the target uses another storage type.
type LogsUploaderTargetUpdateParamsConfigS3GcoreBoundConfig struct {
	Directory param.Opt[string] `json:"directory,omitzero"`
	// Name of a bucket in the selected storage.
	BucketName param.Opt[string] `json:"bucket_name,omitzero"`
	// ID of your Gcore Object Storage. The storage must be active and of the Standard
	// type.
	StorageID param.Opt[int64] `json:"storage_id,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3GcoreBoundConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3GcoreBoundConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3GcoreBoundConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetUpdateParamsConfigS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	Region          string            `json:"region" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3AmazonConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3AmazonConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, SecretAccessKey are required.
type LogsUploaderTargetUpdateParamsConfigS3OssConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	Region   param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParamsConfigS3OtherConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3OtherConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3OtherConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParamsConfigS3V1Config struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3V1Config) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3V1Config
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParamsConfigFtpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParamsConfigSftpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigSftpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigSftpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigSftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Upload is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfig struct {
	Upload LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload `json:"upload,omitzero" api:"required"`
	Append LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend `json:"append,omitzero"`
	Auth   LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth   `json:"auth,omitzero"`
	// Any of "json", "text".
	ContentType string                                              `json:"content_type,omitzero"`
	Retry       LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry `json:"retry,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfig](
		"content_type", "json", "text",
	)
}

// The property URL is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                               `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigUploadResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The property URL is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                               `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigAppend](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigAppendResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, Type are required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth struct {
	Config LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Any of "token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth](
		"type", "token",
	)
}

// The properties Token, HeaderName are required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig struct {
	Token      string `json:"token" api:"required"`
	HeaderName string `json:"header_name" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                              `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigRetry](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigHTTPConfigRetryResponseAction](
		"action", "drop", "retry", "append",
	)
}

type LogsUploaderTargetUpdateParamsConfigAzureBlobConfig struct {
	// Directory path within the container.
	Directory param.Opt[string] `json:"directory,omitzero"`
	// Custom Azure Blob Storage endpoint URL.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Azure Blob Storage account name.
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// Azure Blob Storage container name.
	ContainerName param.Opt[string]                                       `json:"container_name,omitzero"`
	Auth          LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigAzureBlobConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigAzureBlobConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigAzureBlobConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "shared_key", "sas_token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuth](
		"type", "shared_key", "sas_token",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion struct {
	OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey `json:",omitzero,inline"`
	OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken      *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken      `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey, u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken)
}
func (u *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey) {
		return u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigAccountKey
	} else if !param.IsOmitted(u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken) {
		return u.OfLogsUploaderTargetUpdatesConfigAzureBlobConfigAuthConfigToken
	}
	return nil
}

// The property AccountKey is required.
type LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey struct {
	// Azure Blob Storage account key.
	AccountKey string `json:"account_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigAccountKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Token is required.
type LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken struct {
	// Azure Blob Storage SAS token.
	Token string `json:"token" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigAzureBlobConfigAuthConfigToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetUpdateParamsConfigSlsConfig struct {
	// SLS endpoint. Optional — derived from the region as `{region}.log.aliyuncs.com`
	// when omitted.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Optional SLS topic (0-128 characters).
	Topic param.Opt[string] `json:"topic,omitzero"`
	// SLS logstore name. 3-36 characters; lowercase letters, digits, hyphens, and
	// underscores.
	LogStore param.Opt[string] `json:"log_store,omitzero"`
	// SLS project name. 3-63 characters; lowercase letters, digits, and hyphens.
	Project param.Opt[string] `json:"project,omitzero"`
	// SLS region (e.g. `eu-central-1`).
	Region param.Opt[string]                                 `json:"region,omitzero"`
	Auth   LogsUploaderTargetUpdateParamsConfigSlsConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigSlsConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigSlsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigSlsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetUpdateParamsConfigSlsConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "ak_sk".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigSlsConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigSlsConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigSlsConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetUpdateParamsConfigSlsConfigAuth](
		"type", "ak_sk",
	)
}

// Authentication credentials.
//
// The properties AccessKeyID, SecretAccessKey are required.
type LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig struct {
	// Alibaba access key ID.
	AccessKeyID string `json:"access_key_id" api:"required"`
	// Alibaba secret access key.
	SecretAccessKey string `json:"secret_access_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigSlsConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of storage for logs.
type LogsUploaderTargetUpdateParamsStorageType string

const (
	LogsUploaderTargetUpdateParamsStorageTypeS3Gcore   LogsUploaderTargetUpdateParamsStorageType = "s3_gcore"
	LogsUploaderTargetUpdateParamsStorageTypeS3Amazon  LogsUploaderTargetUpdateParamsStorageType = "s3_amazon"
	LogsUploaderTargetUpdateParamsStorageTypeS3Oss     LogsUploaderTargetUpdateParamsStorageType = "s3_oss"
	LogsUploaderTargetUpdateParamsStorageTypeS3Other   LogsUploaderTargetUpdateParamsStorageType = "s3_other"
	LogsUploaderTargetUpdateParamsStorageTypeS3V1      LogsUploaderTargetUpdateParamsStorageType = "s3_v1"
	LogsUploaderTargetUpdateParamsStorageTypeFtp       LogsUploaderTargetUpdateParamsStorageType = "ftp"
	LogsUploaderTargetUpdateParamsStorageTypeSftp      LogsUploaderTargetUpdateParamsStorageType = "sftp"
	LogsUploaderTargetUpdateParamsStorageTypeHTTP      LogsUploaderTargetUpdateParamsStorageType = "http"
	LogsUploaderTargetUpdateParamsStorageTypeAzureBlob LogsUploaderTargetUpdateParamsStorageType = "azure_blob"
	LogsUploaderTargetUpdateParamsStorageTypeSls       LogsUploaderTargetUpdateParamsStorageType = "sls"
)

type LogsUploaderTargetListParams struct {
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search by target name or id.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by ids of related logs uploader configs that use given target.
	ConfigIDs []int64 `query:"config_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogsUploaderTargetListParams]'s query parameters as
// `url.Values`.
func (r LogsUploaderTargetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogsUploaderTargetReplaceParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3Gcore *LogsUploaderTargetReplaceParamsBodyS3Gcore `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Amazon *LogsUploaderTargetReplaceParamsBodyS3Amazon `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Oss *LogsUploaderTargetReplaceParamsBodyS3Oss `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Other *LogsUploaderTargetReplaceParamsBodyS3Other `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3V1 *LogsUploaderTargetReplaceParamsBodyS3V1 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfFtp *LogsUploaderTargetReplaceParamsBodyFtp `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSftp *LogsUploaderTargetReplaceParamsBodySftp `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfHTTP *LogsUploaderTargetReplaceParamsBodyHTTP `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *LogsUploaderTargetReplaceParamsBodyAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSls *LogsUploaderTargetReplaceParamsBodySls `json:",inline"`

	paramObj
}

func (u LogsUploaderTargetReplaceParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3Gcore,
		u.OfS3Amazon,
		u.OfS3Oss,
		u.OfS3Other,
		u.OfS3V1,
		u.OfFtp,
		u.OfSftp,
		u.OfHTTP,
		u.OfAzureBlob,
		u.OfSls)
}
func (r *LogsUploaderTargetReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyS3Gcore struct {
	// Config for the `s3_gcore` storage type.
	Config LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_gcore".
	StorageType constant.S3Gcore `json:"storage_type" default:"s3_gcore"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3Gcore) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3Gcore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3Gcore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion struct {
	OfS3GcoreConfig      *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreConfig      `json:",omitzero,inline"`
	OfS3GcoreBoundConfig *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreBoundConfig `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig, u.OfS3GcoreBoundConfig)
}
func (u *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
	} else if !param.IsOmitted(u.OfS3GcoreBoundConfig) {
		return u.OfS3GcoreBoundConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Endpoint.Valid() {
		return &vt.Endpoint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetUsePathStyle() *bool {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.UsePathStyle.Valid() {
		return &vt.UsePathStyle.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetStorageID() *int64 {
	if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.StorageID.Valid() {
		return &vt.StorageID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.BucketName.Valid() {
		return &vt.BucketName.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil {
		return (*string)(&vt.BucketName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsBodyS3GcoreConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	} else if vt := u.OfS3GcoreBoundConfig; vt != nil && vt.Directory.Valid() {
		return &vt.Directory.Value
	}
	return nil
}

type LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `storage_id` keeps its current value when omitted. To bind a target that is not
// bound yet, or one that uses another storage type, send both `storage_id` and
// `bucket_name`.
//
// The property BucketName is required.
type LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreBoundConfig struct {
	// Name of a bucket in the selected storage.
	BucketName string            `json:"bucket_name" api:"required"`
	Directory  param.Opt[string] `json:"directory,omitzero"`
	// ID of your Gcore Object Storage. The storage must be active and of the Standard
	// type.
	StorageID param.Opt[int64] `json:"storage_id,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreBoundConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreBoundConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3GcoreConfigS3GcoreBoundConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyS3Amazon struct {
	Config LogsUploaderTargetReplaceParamsBodyS3AmazonConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_amazon".
	StorageType constant.S3Amazon `json:"storage_type" default:"s3_amazon"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3Amazon) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3Amazon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3Amazon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetReplaceParamsBodyS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	Region          string            `json:"region" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3AmazonConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3AmazonConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyS3Oss struct {
	Config LogsUploaderTargetReplaceParamsBodyS3OssConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_oss".
	StorageType constant.S3Oss `json:"storage_type" default:"s3_oss"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3Oss) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3Oss
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3Oss) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, SecretAccessKey are required.
type LogsUploaderTargetReplaceParamsBodyS3OssConfig struct {
	AccessKeyID     string            `json:"access_key_id" api:"required"`
	BucketName      string            `json:"bucket_name" api:"required"`
	SecretAccessKey string            `json:"secret_access_key" api:"required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	Region   param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyS3Other struct {
	Config LogsUploaderTargetReplaceParamsBodyS3OtherConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_other".
	StorageType constant.S3Other `json:"storage_type" default:"s3_other"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3Other) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3Other
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3Other) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodyS3OtherConfig struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3OtherConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3OtherConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyS3V1 struct {
	Config LogsUploaderTargetReplaceParamsBodyS3V1Config `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "s3_v1".
	StorageType constant.S3V1 `json:"storage_type" default:"s3_v1"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3V1) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3V1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3V1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodyS3V1Config struct {
	Directory   param.Opt[string] `json:"directory,omitzero"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	BucketName  param.Opt[string] `json:"bucket_name,omitzero"`
	// S3 storage endpoint URL. Must be a public address: private and reserved IP
	// addresses, `localhost`, and hostnames that resolve to them are rejected.
	Endpoint        param.Opt[string] `json:"endpoint,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyS3V1Config) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyS3V1Config
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyFtp struct {
	Config LogsUploaderTargetReplaceParamsBodyFtpConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "ftp".
	StorageType constant.Ftp `json:"storage_type" default:"ftp"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyFtp) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyFtp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyFtp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodyFtpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodySftp struct {
	Config LogsUploaderTargetReplaceParamsBodySftpConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "sftp".
	StorageType constant.Sftp `json:"storage_type" default:"sftp"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySftp) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySftp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySftp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodySftpConfig struct {
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	Hostname       param.Opt[string] `json:"hostname,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySftpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySftpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyHTTP struct {
	Config LogsUploaderTargetReplaceParamsBodyHTTPConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "http".
	StorageType constant.HTTP `json:"storage_type" default:"http"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTP) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Upload is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfig struct {
	Upload LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload `json:"upload,omitzero" api:"required"`
	Append LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend `json:"append,omitzero"`
	Auth   LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth   `json:"auth,omitzero"`
	// Any of "json", "text".
	ContentType string                                             `json:"content_type,omitzero"`
	Retry       LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry `json:"retry,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfig](
		"content_type", "json", "text",
	)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                              `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigUpload](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigUploadResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                              `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigAppend](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigAppendResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, Type are required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth struct {
	Config LogsUploaderTargetReplaceParamsBodyHTTPConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Any of "token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigAuth](
		"type", "token",
	)
}

// The properties Token, HeaderName are required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigAuthConfig struct {
	Token      string `json:"token" api:"required"`
	HeaderName string `json:"header_name" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry struct {
	URL            string            `json:"url" api:"required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                             `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigRetry](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyHTTPConfigRetryResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodyAzureBlob struct {
	Config LogsUploaderTargetReplaceParamsBodyAzureBlobConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "azure_blob".
	StorageType constant.AzureBlob `json:"storage_type" default:"azure_blob"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodyAzureBlobConfig struct {
	// Directory path within the container.
	Directory param.Opt[string] `json:"directory,omitzero"`
	// Custom Azure Blob Storage endpoint URL.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Azure Blob Storage account name.
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// Azure Blob Storage container name.
	ContainerName param.Opt[string]                                      `json:"container_name,omitzero"`
	Auth          LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyAzureBlobConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyAzureBlobConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyAzureBlobConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigUnion `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "shared_key", "sas_token".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuth](
		"type", "shared_key", "sas_token",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigUnion struct {
	OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigAccountKey *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigAccountKey `json:",omitzero,inline"`
	OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigToken      *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigToken      `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigAccountKey, u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigToken)
}
func (u *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigAccountKey) {
		return u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigAccountKey
	} else if !param.IsOmitted(u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigToken) {
		return u.OfLogsUploaderTargetReplacesBodyAzureBlobConfigAuthConfigToken
	}
	return nil
}

// The property AccountKey is required.
type LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigAccountKey struct {
	// Azure Blob Storage account key.
	AccountKey string `json:"account_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigAccountKey) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigAccountKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigAccountKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Token is required.
type LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigToken struct {
	// Azure Blob Storage SAS token.
	Token string `json:"token" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigToken) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodyAzureBlobConfigAuthConfigToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, StorageType are required.
type LogsUploaderTargetReplaceParamsBodySls struct {
	Config LogsUploaderTargetReplaceParamsBodySlsConfig `json:"config,omitzero" api:"required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of storage for logs.
	//
	// This field can be elided, and will marshal its zero value as "sls".
	StorageType constant.Sls `json:"storage_type" default:"sls"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySls) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySls
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetReplaceParamsBodySlsConfig struct {
	// SLS endpoint. Optional — derived from the region as `{region}.log.aliyuncs.com`
	// when omitted.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Optional SLS topic (0-128 characters).
	Topic param.Opt[string] `json:"topic,omitzero"`
	// SLS logstore name. 3-36 characters; lowercase letters, digits, hyphens, and
	// underscores.
	LogStore param.Opt[string] `json:"log_store,omitzero"`
	// SLS project name. 3-63 characters; lowercase letters, digits, and hyphens.
	Project param.Opt[string] `json:"project,omitzero"`
	// SLS region (e.g. `eu-central-1`).
	Region param.Opt[string]                                `json:"region,omitzero"`
	Auth   LogsUploaderTargetReplaceParamsBodySlsConfigAuth `json:"auth,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySlsConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySlsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySlsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Type are required.
type LogsUploaderTargetReplaceParamsBodySlsConfigAuth struct {
	// Authentication credentials.
	Config LogsUploaderTargetReplaceParamsBodySlsConfigAuthConfig `json:"config,omitzero" api:"required"`
	// Authentication type.
	//
	// Any of "ak_sk".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySlsConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySlsConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySlsConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsBodySlsConfigAuth](
		"type", "ak_sk",
	)
}

// Authentication credentials.
//
// The properties AccessKeyID, SecretAccessKey are required.
type LogsUploaderTargetReplaceParamsBodySlsConfigAuthConfig struct {
	// Alibaba access key ID.
	AccessKeyID string `json:"access_key_id" api:"required"`
	// Alibaba secret access key.
	SecretAccessKey string `json:"secret_access_key" api:"required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsBodySlsConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsBodySlsConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsBodySlsConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
