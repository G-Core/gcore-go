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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

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
func (r *LogsUploaderTargetService) New(ctx context.Context, body LogsUploaderTargetNewParams, opts ...option.RequestOption) (res *LogsUploaderTarget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/targets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change logs uploader target partially.
func (r *LogsUploaderTargetService) Update(ctx context.Context, id int64, body LogsUploaderTargetUpdateParams, opts ...option.RequestOption) (res *LogsUploaderTarget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get list of logs uploader targets.
func (r *LogsUploaderTargetService) List(ctx context.Context, query LogsUploaderTargetListParams, opts ...option.RequestOption) (res *LogsUploaderTargetList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/targets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	return
}

// Get information about logs uploader target.
func (r *LogsUploaderTargetService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderTarget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change logs uploader target.
func (r *LogsUploaderTargetService) Replace(ctx context.Context, id int64, body LogsUploaderTargetReplaceParams, opts ...option.RequestOption) (res *LogsUploaderTarget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Validate logs uploader target.
func (r *LogsUploaderTargetService) Validate(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/targets/%v/validate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type LogsUploaderTarget struct {
	ID int64 `json:"id"`
	// Client that owns the target.
	ClientID int64 `json:"client_id"`
	// Config for specific storage type.
	Config LogsUploaderTargetConfigUnion `json:"config"`
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
	Status LogsUploaderTargetStatus `json:"status"`
	// Type of storage for logs.
	//
	// Any of "s3_gcore", "s3_amazon", "s3_oss", "s3_other", "s3_v1", "ftp", "sftp",
	// "http".
	StorageType LogsUploaderTargetStorageType `json:"storage_type"`
	// Time when logs uploader target was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ClientID               respjson.Field
		Config                 respjson.Field
		Created                respjson.Field
		Description            respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		Status                 respjson.Field
		StorageType            respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTarget) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderTargetConfigUnion contains all possible properties and values from
// [LogsUploaderTargetConfigS3GcoreConfig],
// [LogsUploaderTargetConfigS3AmazonConfig], [LogsUploaderTargetConfigObject],
// [LogsUploaderTargetConfigS3GcoreConfig],
// [LogsUploaderTargetConfigS3GcoreConfig], [LogsUploaderTargetConfigFtpConfig],
// [LogsUploaderTargetConfigSftpConfig], [LogsUploaderTargetConfigHTTPConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LogsUploaderTargetConfigUnion struct {
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory"`
	// This field is from variant [LogsUploaderTargetConfigS3GcoreConfig].
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
	// This field is from variant [LogsUploaderTargetConfigS3GcoreConfig].
	UsePathStyle   bool   `json:"use_path_style"`
	Hostname       string `json:"hostname"`
	TimeoutSeconds int64  `json:"timeout_seconds"`
	User           string `json:"user"`
	// This field is from variant [LogsUploaderTargetConfigSftpConfig].
	KeyPassphrase string `json:"key_passphrase"`
	// This field is from variant [LogsUploaderTargetConfigSftpConfig].
	Password string `json:"password"`
	// This field is from variant [LogsUploaderTargetConfigSftpConfig].
	PrivateKey string `json:"private_key"`
	// This field is from variant [LogsUploaderTargetConfigHTTPConfig].
	Append LogsUploaderTargetConfigHTTPConfigAppend `json:"append"`
	// This field is from variant [LogsUploaderTargetConfigHTTPConfig].
	Auth LogsUploaderTargetConfigHTTPConfigAuth `json:"auth"`
	// This field is from variant [LogsUploaderTargetConfigHTTPConfig].
	ContentType string `json:"content_type"`
	// This field is from variant [LogsUploaderTargetConfigHTTPConfig].
	Retry LogsUploaderTargetConfigHTTPConfigRetry `json:"retry"`
	// This field is from variant [LogsUploaderTargetConfigHTTPConfig].
	Upload LogsUploaderTargetConfigHTTPConfigUpload `json:"upload"`
	JSON   struct {
		AccessKeyID    respjson.Field
		BucketName     respjson.Field
		Directory      respjson.Field
		Endpoint       respjson.Field
		Region         respjson.Field
		UsePathStyle   respjson.Field
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
		raw            string
	} `json:"-"`
}

func (u LogsUploaderTargetConfigUnion) AsS3GcoreConfig() (v LogsUploaderTargetConfigS3GcoreConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsS3AmazonConfig() (v LogsUploaderTargetConfigS3AmazonConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsLogsUploaderTargetConfigObject() (v LogsUploaderTargetConfigObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsLogsUploaderTargetConfigS3GcoreConfig() (v LogsUploaderTargetConfigS3GcoreConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsVariant2() (v LogsUploaderTargetConfigS3GcoreConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsFtpConfig() (v LogsUploaderTargetConfigFtpConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsSftpConfig() (v LogsUploaderTargetConfigSftpConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderTargetConfigUnion) AsHTTPConfig() (v LogsUploaderTargetConfigHTTPConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogsUploaderTargetConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *LogsUploaderTargetConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigS3GcoreConfig struct {
	AccessKeyID  string `json:"access_key_id"`
	BucketName   string `json:"bucket_name"`
	Directory    string `json:"directory,nullable"`
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
func (r LogsUploaderTargetConfigS3GcoreConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigS3AmazonConfig struct {
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory,nullable"`
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
func (r LogsUploaderTargetConfigS3AmazonConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigObject struct {
	AccessKeyID string `json:"access_key_id"`
	BucketName  string `json:"bucket_name"`
	Directory   string `json:"directory,nullable"`
	Region      string `json:"region,nullable"`
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
func (r LogsUploaderTargetConfigObject) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigFtpConfig struct {
	Directory      string `json:"directory,nullable"`
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
func (r LogsUploaderTargetConfigFtpConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigSftpConfig struct {
	Hostname       string `json:"hostname,required"`
	User           string `json:"user,required"`
	Directory      string `json:"directory,nullable"`
	KeyPassphrase  string `json:"key_passphrase,nullable"`
	Password       string `json:"password,nullable"`
	PrivateKey     string `json:"private_key,nullable"`
	TimeoutSeconds int64  `json:"timeout_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hostname       respjson.Field
		User           respjson.Field
		Directory      respjson.Field
		KeyPassphrase  respjson.Field
		Password       respjson.Field
		PrivateKey     respjson.Field
		TimeoutSeconds respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetConfigSftpConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigSftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfig struct {
	Append LogsUploaderTargetConfigHTTPConfigAppend `json:"append"`
	Auth   LogsUploaderTargetConfigHTTPConfigAuth   `json:"auth"`
	// Any of "json", "text".
	ContentType string                                   `json:"content_type"`
	Retry       LogsUploaderTargetConfigHTTPConfigRetry  `json:"retry"`
	Upload      LogsUploaderTargetConfigHTTPConfigUpload `json:"upload"`
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
func (r LogsUploaderTargetConfigHTTPConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigAppend struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                                   `json:"method"`
	ResponseActions []LogsUploaderTargetConfigHTTPConfigAppendResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                                    `json:"timeout_seconds"`
	URL             string                                                   `json:"url"`
	UseCompression  bool                                                     `json:"use_compression"`
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
func (r LogsUploaderTargetConfigHTTPConfigAppend) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigAppendResponseAction struct {
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
func (r LogsUploaderTargetConfigHTTPConfigAppendResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigAuth struct {
	Config LogsUploaderTargetConfigHTTPConfigAuthConfig `json:"config"`
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
func (r LogsUploaderTargetConfigHTTPConfigAuth) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigAuthConfig struct {
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
func (r LogsUploaderTargetConfigHTTPConfigAuthConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigRetry struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                                  `json:"method"`
	ResponseActions []LogsUploaderTargetConfigHTTPConfigRetryResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                                   `json:"timeout_seconds"`
	URL             string                                                  `json:"url"`
	UseCompression  bool                                                    `json:"use_compression"`
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
func (r LogsUploaderTargetConfigHTTPConfigRetry) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigRetryResponseAction struct {
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
func (r LogsUploaderTargetConfigHTTPConfigRetryResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigUpload struct {
	Headers map[string]string `json:"headers"`
	// Any of "POST", "PUT".
	Method          string                                                   `json:"method"`
	ResponseActions []LogsUploaderTargetConfigHTTPConfigUploadResponseAction `json:"response_actions"`
	TimeoutSeconds  int64                                                    `json:"timeout_seconds"`
	URL             string                                                   `json:"url"`
	UseCompression  bool                                                     `json:"use_compression"`
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
func (r LogsUploaderTargetConfigHTTPConfigUpload) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderTargetConfigHTTPConfigUploadResponseAction struct {
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
func (r LogsUploaderTargetConfigHTTPConfigUploadResponseAction) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetConfigHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader target. Informs if the specified target
// is reachable.
type LogsUploaderTargetStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderTargetStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderTargetStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of storage for logs.
type LogsUploaderTargetStorageType string

const (
	LogsUploaderTargetStorageTypeS3Gcore  LogsUploaderTargetStorageType = "s3_gcore"
	LogsUploaderTargetStorageTypeS3Amazon LogsUploaderTargetStorageType = "s3_amazon"
	LogsUploaderTargetStorageTypeS3Oss    LogsUploaderTargetStorageType = "s3_oss"
	LogsUploaderTargetStorageTypeS3Other  LogsUploaderTargetStorageType = "s3_other"
	LogsUploaderTargetStorageTypeS3V1     LogsUploaderTargetStorageType = "s3_v1"
	LogsUploaderTargetStorageTypeFtp      LogsUploaderTargetStorageType = "ftp"
	LogsUploaderTargetStorageTypeSftp     LogsUploaderTargetStorageType = "sftp"
	LogsUploaderTargetStorageTypeHTTP     LogsUploaderTargetStorageType = "http"
)

type LogsUploaderTargetList []LogsUploaderTarget

type LogsUploaderTargetNewParams struct {
	// Config for specific storage type.
	Config LogsUploaderTargetNewParamsConfigUnion `json:"config,omitzero,required"`
	// Type of storage for logs.
	//
	// Any of "s3_gcore", "s3_amazon", "s3_oss", "s3_other", "s3_v1", "ftp", "sftp",
	// "http".
	StorageType LogsUploaderTargetNewParamsStorageType `json:"storage_type,omitzero,required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetNewParamsConfigUnion struct {
	OfS3GcoreConfig  *LogsUploaderTargetNewParamsConfigS3GcoreConfig  `json:",omitzero,inline"`
	OfS3AmazonConfig *LogsUploaderTargetNewParamsConfigS3AmazonConfig `json:",omitzero,inline"`
	OfS3OssConfig    *LogsUploaderTargetNewParamsConfigS3OssConfig    `json:",omitzero,inline"`
	OfS3OtherConfig  *LogsUploaderTargetNewParamsConfigS3OtherConfig  `json:",omitzero,inline"`
	OfS3V1Config     *LogsUploaderTargetNewParamsConfigS3V1Config     `json:",omitzero,inline"`
	OfFtpConfig      *LogsUploaderTargetNewParamsConfigFtpConfig      `json:",omitzero,inline"`
	OfSftpConfig     *LogsUploaderTargetNewParamsConfigSftpConfig     `json:",omitzero,inline"`
	OfHTTPConfig     *LogsUploaderTargetNewParamsConfigHTTPConfig     `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetNewParamsConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig,
		u.OfS3AmazonConfig,
		u.OfS3OssConfig,
		u.OfS3OtherConfig,
		u.OfS3V1Config,
		u.OfFtpConfig,
		u.OfSftpConfig,
		u.OfHTTPConfig)
}
func (u *LogsUploaderTargetNewParamsConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetNewParamsConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetKeyPassphrase() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.KeyPassphrase.Valid() {
		return &vt.KeyPassphrase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetPrivateKey() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.PrivateKey.Valid() {
		return &vt.PrivateKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetUpload() *LogsUploaderTargetNewParamsConfigHTTPConfigUpload {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Upload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetAppend() *LogsUploaderTargetNewParamsConfigHTTPConfigAppend {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Append
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetAuth() *LogsUploaderTargetNewParamsConfigHTTPConfigAuth {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Auth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetContentType() *string {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.ContentType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetRetry() *LogsUploaderTargetNewParamsConfigHTTPConfigRetry {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Retry
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.AccessKeyID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.BucketName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Endpoint)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Region)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetUsePathStyle() *bool {
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
func (u LogsUploaderTargetNewParamsConfigUnion) GetHostname() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetPassword() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Password)
	} else if vt := u.OfSftpConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetUser() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.User)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.User)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetNewParamsConfigUnion) GetTimeoutSeconds() *int64 {
	if vt := u.OfFtpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	}
	return nil
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetNewParamsConfigS3GcoreConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigS3GcoreConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigS3GcoreConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetNewParamsConfigS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigS3AmazonConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigS3AmazonConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, SecretAccessKey are required.
type LogsUploaderTargetNewParamsConfigS3OssConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetNewParamsConfigS3OtherConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigS3OtherConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigS3OtherConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetNewParamsConfigS3V1Config struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigS3V1Config) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigS3V1Config
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hostname, Password, User are required.
type LogsUploaderTargetNewParamsConfigFtpConfig struct {
	Hostname       string            `json:"hostname,required"`
	Password       string            `json:"password,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hostname, User are required.
type LogsUploaderTargetNewParamsConfigSftpConfig struct {
	Hostname       string            `json:"hostname,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigSftpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigSftpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigSftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Upload is required.
type LogsUploaderTargetNewParamsConfigHTTPConfig struct {
	Upload LogsUploaderTargetNewParamsConfigHTTPConfigUpload `json:"upload,omitzero,required"`
	Append LogsUploaderTargetNewParamsConfigHTTPConfigAppend `json:"append,omitzero"`
	Auth   LogsUploaderTargetNewParamsConfigHTTPConfigAuth   `json:"auth,omitzero"`
	// Any of "json", "text".
	ContentType string                                           `json:"content_type,omitzero"`
	Retry       LogsUploaderTargetNewParamsConfigHTTPConfigRetry `json:"retry,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfig](
		"content_type", "json", "text",
	)
}

// The property URL is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigUpload struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                            `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigUpload) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigUpload](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigUploadResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The property URL is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigAppend struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                            `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigAppend) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigAppend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigAppend](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigAppendResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, Type are required.
type LogsUploaderTargetNewParamsConfigHTTPConfigAuth struct {
	Config LogsUploaderTargetNewParamsConfigHTTPConfigAuthConfig `json:"config,omitzero,required"`
	// Any of "token".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigAuth](
		"type", "token",
	)
}

// The properties Token, HeaderName are required.
type LogsUploaderTargetNewParamsConfigHTTPConfigAuthConfig struct {
	Token      string `json:"token,required"`
	HeaderName string `json:"header_name,required"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigRetry struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                           `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigRetry) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigRetry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigRetry](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetNewParamsConfigHTTPConfigRetryResponseAction](
		"action", "drop", "retry", "append",
	)
}

// Type of storage for logs.
type LogsUploaderTargetNewParamsStorageType string

const (
	LogsUploaderTargetNewParamsStorageTypeS3Gcore  LogsUploaderTargetNewParamsStorageType = "s3_gcore"
	LogsUploaderTargetNewParamsStorageTypeS3Amazon LogsUploaderTargetNewParamsStorageType = "s3_amazon"
	LogsUploaderTargetNewParamsStorageTypeS3Oss    LogsUploaderTargetNewParamsStorageType = "s3_oss"
	LogsUploaderTargetNewParamsStorageTypeS3Other  LogsUploaderTargetNewParamsStorageType = "s3_other"
	LogsUploaderTargetNewParamsStorageTypeS3V1     LogsUploaderTargetNewParamsStorageType = "s3_v1"
	LogsUploaderTargetNewParamsStorageTypeFtp      LogsUploaderTargetNewParamsStorageType = "ftp"
	LogsUploaderTargetNewParamsStorageTypeSftp     LogsUploaderTargetNewParamsStorageType = "sftp"
	LogsUploaderTargetNewParamsStorageTypeHTTP     LogsUploaderTargetNewParamsStorageType = "http"
)

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
	// "http".
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
	OfS3GcoreConfig  *LogsUploaderTargetUpdateParamsConfigS3GcoreConfig  `json:",omitzero,inline"`
	OfS3AmazonConfig *LogsUploaderTargetUpdateParamsConfigS3AmazonConfig `json:",omitzero,inline"`
	OfS3OssConfig    *LogsUploaderTargetUpdateParamsConfigS3OssConfig    `json:",omitzero,inline"`
	OfS3OtherConfig  *LogsUploaderTargetUpdateParamsConfigS3OtherConfig  `json:",omitzero,inline"`
	OfS3V1Config     *LogsUploaderTargetUpdateParamsConfigS3V1Config     `json:",omitzero,inline"`
	OfFtpConfig      *LogsUploaderTargetUpdateParamsConfigFtpConfig      `json:",omitzero,inline"`
	OfSftpConfig     *LogsUploaderTargetUpdateParamsConfigSftpConfig     `json:",omitzero,inline"`
	OfHTTPConfig     *LogsUploaderTargetUpdateParamsConfigHTTPConfig     `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetUpdateParamsConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig,
		u.OfS3AmazonConfig,
		u.OfS3OssConfig,
		u.OfS3OtherConfig,
		u.OfS3V1Config,
		u.OfFtpConfig,
		u.OfSftpConfig,
		u.OfHTTPConfig)
}
func (u *LogsUploaderTargetUpdateParamsConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetUpdateParamsConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
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
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAuth() *LogsUploaderTargetUpdateParamsConfigHTTPConfigAuth {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Auth
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
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.AccessKeyID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.BucketName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Endpoint)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Region)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
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
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetPassword() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Password)
	} else if vt := u.OfSftpConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetUpdateParamsConfigUnion) GetUser() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.User)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.User)
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

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetUpdateParamsConfigS3GcoreConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
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

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetUpdateParamsConfigS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
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
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetUpdateParamsConfigS3OtherConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
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

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetUpdateParamsConfigS3V1Config struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
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

// The properties Hostname, Password, User are required.
type LogsUploaderTargetUpdateParamsConfigFtpConfig struct {
	Hostname       string            `json:"hostname,required"`
	Password       string            `json:"password,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r LogsUploaderTargetUpdateParamsConfigFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetUpdateParamsConfigFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetUpdateParamsConfigFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hostname, User are required.
type LogsUploaderTargetUpdateParamsConfigSftpConfig struct {
	Hostname       string            `json:"hostname,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
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
	Upload LogsUploaderTargetUpdateParamsConfigHTTPConfigUpload `json:"upload,omitzero,required"`
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
	URL            string            `json:"url,required"`
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
	Action          string            `json:"action,omitzero,required"`
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
	URL            string            `json:"url,required"`
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
	Action          string            `json:"action,omitzero,required"`
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
	Config LogsUploaderTargetUpdateParamsConfigHTTPConfigAuthConfig `json:"config,omitzero,required"`
	// Any of "token".
	Type string `json:"type,omitzero,required"`
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
	Token      string `json:"token,required"`
	HeaderName string `json:"header_name,required"`
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
	URL            string            `json:"url,required"`
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
	Action          string            `json:"action,omitzero,required"`
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

// Type of storage for logs.
type LogsUploaderTargetUpdateParamsStorageType string

const (
	LogsUploaderTargetUpdateParamsStorageTypeS3Gcore  LogsUploaderTargetUpdateParamsStorageType = "s3_gcore"
	LogsUploaderTargetUpdateParamsStorageTypeS3Amazon LogsUploaderTargetUpdateParamsStorageType = "s3_amazon"
	LogsUploaderTargetUpdateParamsStorageTypeS3Oss    LogsUploaderTargetUpdateParamsStorageType = "s3_oss"
	LogsUploaderTargetUpdateParamsStorageTypeS3Other  LogsUploaderTargetUpdateParamsStorageType = "s3_other"
	LogsUploaderTargetUpdateParamsStorageTypeS3V1     LogsUploaderTargetUpdateParamsStorageType = "s3_v1"
	LogsUploaderTargetUpdateParamsStorageTypeFtp      LogsUploaderTargetUpdateParamsStorageType = "ftp"
	LogsUploaderTargetUpdateParamsStorageTypeSftp     LogsUploaderTargetUpdateParamsStorageType = "sftp"
	LogsUploaderTargetUpdateParamsStorageTypeHTTP     LogsUploaderTargetUpdateParamsStorageType = "http"
)

type LogsUploaderTargetListParams struct {
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
	// Config for specific storage type.
	Config LogsUploaderTargetReplaceParamsConfigUnion `json:"config,omitzero,required"`
	// Type of storage for logs.
	//
	// Any of "s3_gcore", "s3_amazon", "s3_oss", "s3_other", "s3_v1", "ftp", "sftp",
	// "http".
	StorageType LogsUploaderTargetReplaceParamsStorageType `json:"storage_type,omitzero,required"`
	// Description of the target.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderTargetReplaceParamsConfigUnion struct {
	OfS3GcoreConfig  *LogsUploaderTargetReplaceParamsConfigS3GcoreConfig  `json:",omitzero,inline"`
	OfS3AmazonConfig *LogsUploaderTargetReplaceParamsConfigS3AmazonConfig `json:",omitzero,inline"`
	OfS3OssConfig    *LogsUploaderTargetReplaceParamsConfigS3OssConfig    `json:",omitzero,inline"`
	OfS3OtherConfig  *LogsUploaderTargetReplaceParamsConfigS3OtherConfig  `json:",omitzero,inline"`
	OfS3V1Config     *LogsUploaderTargetReplaceParamsConfigS3V1Config     `json:",omitzero,inline"`
	OfFtpConfig      *LogsUploaderTargetReplaceParamsConfigFtpConfig      `json:",omitzero,inline"`
	OfSftpConfig     *LogsUploaderTargetReplaceParamsConfigSftpConfig     `json:",omitzero,inline"`
	OfHTTPConfig     *LogsUploaderTargetReplaceParamsConfigHTTPConfig     `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderTargetReplaceParamsConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3GcoreConfig,
		u.OfS3AmazonConfig,
		u.OfS3OssConfig,
		u.OfS3OtherConfig,
		u.OfS3V1Config,
		u.OfFtpConfig,
		u.OfSftpConfig,
		u.OfHTTPConfig)
}
func (u *LogsUploaderTargetReplaceParamsConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderTargetReplaceParamsConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfS3GcoreConfig) {
		return u.OfS3GcoreConfig
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetKeyPassphrase() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.KeyPassphrase.Valid() {
		return &vt.KeyPassphrase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetPrivateKey() *string {
	if vt := u.OfSftpConfig; vt != nil && vt.PrivateKey.Valid() {
		return &vt.PrivateKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetUpload() *LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Upload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetAppend() *LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Append
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetAuth() *LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Auth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetContentType() *string {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.ContentType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetRetry() *LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry {
	if vt := u.OfHTTPConfig; vt != nil {
		return &vt.Retry
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetAccessKeyID() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.AccessKeyID)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.AccessKeyID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetBucketName() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.BucketName)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.BucketName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetEndpoint() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Endpoint)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Endpoint)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetRegion() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3OssConfig; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.Region)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.Region)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetSecretAccessKey() *string {
	if vt := u.OfS3GcoreConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3AmazonConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OssConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3OtherConfig; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	} else if vt := u.OfS3V1Config; vt != nil {
		return (*string)(&vt.SecretAccessKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetDirectory() *string {
	if vt := u.OfS3GcoreConfig; vt != nil && vt.Directory.Valid() {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetUsePathStyle() *bool {
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
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetHostname() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.Hostname)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetPassword() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.Password)
	} else if vt := u.OfSftpConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetUser() *string {
	if vt := u.OfFtpConfig; vt != nil {
		return (*string)(&vt.User)
	} else if vt := u.OfSftpConfig; vt != nil {
		return (*string)(&vt.User)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderTargetReplaceParamsConfigUnion) GetTimeoutSeconds() *int64 {
	if vt := u.OfFtpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	} else if vt := u.OfSftpConfig; vt != nil && vt.TimeoutSeconds.Valid() {
		return &vt.TimeoutSeconds.Value
	}
	return nil
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetReplaceParamsConfigS3GcoreConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigS3GcoreConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigS3GcoreConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigS3GcoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Region, SecretAccessKey are required.
type LogsUploaderTargetReplaceParamsConfigS3AmazonConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigS3AmazonConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigS3AmazonConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigS3AmazonConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, SecretAccessKey are required.
type LogsUploaderTargetReplaceParamsConfigS3OssConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	Region          param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigS3OssConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigS3OssConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigS3OssConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetReplaceParamsConfigS3OtherConfig struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigS3OtherConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigS3OtherConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigS3OtherConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKeyID, BucketName, Endpoint, Region, SecretAccessKey are
// required.
type LogsUploaderTargetReplaceParamsConfigS3V1Config struct {
	AccessKeyID     string            `json:"access_key_id,required"`
	BucketName      string            `json:"bucket_name,required"`
	Endpoint        string            `json:"endpoint,required"`
	Region          string            `json:"region,required"`
	SecretAccessKey string            `json:"secret_access_key,required"`
	Directory       param.Opt[string] `json:"directory,omitzero"`
	UsePathStyle    param.Opt[bool]   `json:"use_path_style,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigS3V1Config) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigS3V1Config
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigS3V1Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hostname, Password, User are required.
type LogsUploaderTargetReplaceParamsConfigFtpConfig struct {
	Hostname       string            `json:"hostname,required"`
	Password       string            `json:"password,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigFtpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigFtpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigFtpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Hostname, User are required.
type LogsUploaderTargetReplaceParamsConfigSftpConfig struct {
	Hostname       string            `json:"hostname,required"`
	User           string            `json:"user,required"`
	Directory      param.Opt[string] `json:"directory,omitzero"`
	KeyPassphrase  param.Opt[string] `json:"key_passphrase,omitzero"`
	Password       param.Opt[string] `json:"password,omitzero"`
	PrivateKey     param.Opt[string] `json:"private_key,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigSftpConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigSftpConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigSftpConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Upload is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfig struct {
	Upload LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload `json:"upload,omitzero,required"`
	Append LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend `json:"append,omitzero"`
	Auth   LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth   `json:"auth,omitzero"`
	// Any of "json", "text".
	ContentType string                                               `json:"content_type,omitzero"`
	Retry       LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry `json:"retry,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfig](
		"content_type", "json", "text",
	)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                                `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigUpload](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigUploadResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                                `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigAppend](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigAppendResponseAction](
		"action", "drop", "retry", "append",
	)
}

// The properties Config, Type are required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth struct {
	Config LogsUploaderTargetReplaceParamsConfigHTTPConfigAuthConfig `json:"config,omitzero,required"`
	// Any of "token".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigAuth](
		"type", "token",
	)
}

// The properties Token, HeaderName are required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigAuthConfig struct {
	Token      string `json:"token,required"`
	HeaderName string `json:"header_name,required"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigAuthConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigAuthConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry struct {
	URL            string            `json:"url,required"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	UseCompression param.Opt[bool]   `json:"use_compression,omitzero"`
	Headers        map[string]string `json:"headers,omitzero"`
	// Any of "POST", "PUT".
	Method          string                                                               `json:"method,omitzero"`
	ResponseActions []LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction `json:"response_actions,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigRetry](
		"method", "POST", "PUT",
	)
}

// The property Action is required.
type LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction struct {
	// Any of "drop", "retry", "append".
	Action          string            `json:"action,omitzero,required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	MatchPayload    param.Opt[string] `json:"match_payload,omitzero"`
	MatchStatusCode param.Opt[int64]  `json:"match_status_code,omitzero"`
	paramObj
}

func (r LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderTargetReplaceParamsConfigHTTPConfigRetryResponseAction](
		"action", "drop", "retry", "append",
	)
}

// Type of storage for logs.
type LogsUploaderTargetReplaceParamsStorageType string

const (
	LogsUploaderTargetReplaceParamsStorageTypeS3Gcore  LogsUploaderTargetReplaceParamsStorageType = "s3_gcore"
	LogsUploaderTargetReplaceParamsStorageTypeS3Amazon LogsUploaderTargetReplaceParamsStorageType = "s3_amazon"
	LogsUploaderTargetReplaceParamsStorageTypeS3Oss    LogsUploaderTargetReplaceParamsStorageType = "s3_oss"
	LogsUploaderTargetReplaceParamsStorageTypeS3Other  LogsUploaderTargetReplaceParamsStorageType = "s3_other"
	LogsUploaderTargetReplaceParamsStorageTypeS3V1     LogsUploaderTargetReplaceParamsStorageType = "s3_v1"
	LogsUploaderTargetReplaceParamsStorageTypeFtp      LogsUploaderTargetReplaceParamsStorageType = "ftp"
	LogsUploaderTargetReplaceParamsStorageTypeSftp     LogsUploaderTargetReplaceParamsStorageType = "sftp"
	LogsUploaderTargetReplaceParamsStorageTypeHTTP     LogsUploaderTargetReplaceParamsStorageType = "http"
)
