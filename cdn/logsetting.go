// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LogSettingService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogSettingService] method instead.
type LogSettingService struct {
	Options []option.RequestOption
}

// NewLogSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogSettingService(opts ...option.RequestOption) (r LogSettingService) {
	r = LogSettingService{}
	r.Options = opts
	return
}

// Setup raw logs settings
func (r *LogSettingService) New(ctx context.Context, body LogSettingNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cdn/raw_log_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// PATCH method is not allowed.
func (r *LogSettingService) Update(ctx context.Context, body LogSettingUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cdn/raw_log_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Delete the raw logs delivery configuration from the system permanently.
//
// Notes:
//
//   - **Deactivation Requirement**: Set the `enabled` attribute to `false` before
//     deletion.
//   - **Irreversibility**: This action is irreversible. Once deleted, the raw logs
//     delivery configuration cannot be recovered.
func (r *LogSettingService) Delete(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cdn/raw_log_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about raw logs feature settings.
func (r *LogSettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *LogSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/raw_log_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogSettings struct {
	// Name of the S3 bucket to which logs of all CDN resources are delivered.
	//
	// Applicable for "`storage_type`": S3.
	AllResourcesBucket string `json:"all_resources_bucket"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	//   - **s3** - Name of the S3 bucket sub-folder to which logs for all CDN resources
	//     are delivered.
	//   - **ftp/sftp** - Name of the folder (or path) to which logs for all CDN
	//     resources are delivered.
	AllResourcesFolder string `json:"all_resources_folder"`
	// The size of a single piece of the archive in MB. In case of **null** value logs
	// are delivered without slicing.
	ArchiveSizeMB int64 `json:"archive_size_mb,nullable"`
	// Client ID.
	Client int64 `json:"client"`
	// System comment on the status of settings, if they are suspended.
	Comment string `json:"comment,nullable"`
	// Enables or disables a log forwarding feature.
	//
	// Possible values:
	//
	// - **true** - log forwarding feature is active.
	// - **false** - log forwarding feature is deactivated.
	Enabled bool `json:"enabled"`
	// List of folders/buckets for receiving CDN resources logs.
	Folders []LogSettingsFolder `json:"folders"`
	// Defines whether logs of all CDN resources are delivered to one folder/bucket or
	// to separate ones.
	//
	// Possible values:
	//
	// - **true** - Logs of all CDN resources are delivered to one folder/bucket.
	// - **false** - Logs of CDN resources are delivered to separate folders/buckets.
	ForAllResources bool `json:"for_all_resources"`
	// FTP storage hostname.
	FtpHostname string `json:"ftp_hostname"`
	// FTP storage login.
	FtpLogin string `json:"ftp_login"`
	// Name of prepend FTP folder for log delivery.
	FtpPrependFolder string `json:"ftp_prepend_folder"`
	// Enables or disables the forwarding of empty logs.
	//
	// Possible values:
	//
	// - **true** - Empty logs are not sent.
	// - **false** - Empty logs are sent.
	IgnoreEmptyLogs bool `json:"ignore_empty_logs"`
	// Access key ID for the S3 account.
	//
	// Access Key ID is 20 alpha-numeric characters like 022QF06E7MXBSH9DHM02
	S3AccessKeyID string `json:"s3_access_key_id"`
	// Amazon AWS region.
	S3AwsRegion string `json:"s3_aws_region"`
	// S3 storage location.
	//
	// Restrictions:
	//
	//   - Maximum 255 symbols.
	//   - Latin letters (A-Z, a-z,) digits (0-9,) dots, colons, dashes, and underscores
	//     (.:\_-).
	S3BucketLocation string `json:"s3_bucket_location"`
	// S3 storage bucket hostname.
	//
	// Restrictions:
	//
	// - Maximum 255 symbols.
	// - Latin letters (A-Z, a-z,) digits (0-9,) dots, colons, dashes, and underscores.
	S3HostBucket string `json:"s3_host_bucket"`
	// S3 storage hostname.
	S3Hostname string `json:"s3_hostname"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type"`
	// SFTP storage hostname.
	SftpHostname string `json:"sftp_hostname"`
	// SFTP storage login.
	SftpLogin string `json:"sftp_login"`
	// Name of prepend SFTP folder for log delivery.
	SftpPrependFolder string `json:"sftp_prepend_folder"`
	// Log delivery status.
	//
	// Possible values:
	//
	// - **ok** – All/part of attempts to deliver logs were successful.
	// - **failed** – All attempts to deliver logs failed.
	// - **pending** - No logs delivery attempts yet.
	// - **disabled** - Log delivery is disabled.
	Status string `json:"status"`
	// Storage type.
	//
	// Possible values:
	//
	// - **ftp**
	// - **sftp**
	// - **s3**
	StorageType string `json:"storage_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllResourcesBucket respjson.Field
		AllResourcesFolder respjson.Field
		ArchiveSizeMB      respjson.Field
		Client             respjson.Field
		Comment            respjson.Field
		Enabled            respjson.Field
		Folders            respjson.Field
		ForAllResources    respjson.Field
		FtpHostname        respjson.Field
		FtpLogin           respjson.Field
		FtpPrependFolder   respjson.Field
		IgnoreEmptyLogs    respjson.Field
		S3AccessKeyID      respjson.Field
		S3AwsRegion        respjson.Field
		S3BucketLocation   respjson.Field
		S3HostBucket       respjson.Field
		S3Hostname         respjson.Field
		S3Type             respjson.Field
		SftpHostname       respjson.Field
		SftpLogin          respjson.Field
		SftpPrependFolder  respjson.Field
		Status             respjson.Field
		StorageType        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogSettings) RawJSON() string { return r.JSON.raw }
func (r *LogSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogSettingsFolder struct {
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket ID.
	// - **ftp/sftp** - FTP/SFTP folder ID.
	ID int64 `json:"id"`
	// S3 bucket name.
	//
	// The field is required if "`storage_type`": **s3**.
	Bucket string `json:"bucket"`
	// CDN resource ID.
	CdnResource int64 `json:"cdn_resource"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket sub-folder name (optional.)
	// - **ftp/sftp** - FTP/SFTP folder name (required.)
	Folder string `json:"folder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bucket      respjson.Field
		CdnResource respjson.Field
		Folder      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogSettingsFolder) RawJSON() string { return r.JSON.raw }
func (r *LogSettingsFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogSettingNewParams struct {
	// Name of the S3 bucket to which logs for all CDN resources are delivered.
	AllResourcesBucket string `json:"all_resources_bucket,required"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	//   - If "`storage_type`": s3 - Name of the S3 bucket sub-folder to which logs for
	//     all CDN resources are delivered.
	//   - If "`storage_type`": ftp/sftp - Name of the folder (or path) to which logs for
	//     all CDN resources are delivered.
	AllResourcesFolder string `json:"all_resources_folder,required"`
	// List of folders/buckets for receiving CDN resources logs.
	Folders []LogSettingNewParamsFolder `json:"folders,omitzero,required"`
	// Defines whether logs of all CDN resources are delivered to one folder/bucket or
	// to separate ones.
	//
	// Possible values:
	//
	//   - **true** - Logs of all CDN resources are delivered to one folder/bucket.
	//   - **false** - Logs of different CDN resources are delivered to separate
	//     folders/buckets.
	ForAllResources bool `json:"for_all_resources,required"`
	// FTP storage hostname.
	FtpHostname string `json:"ftp_hostname,required"`
	// FTP storage login.
	FtpLogin string `json:"ftp_login,required"`
	// FTP storage password.
	FtpPassword string `json:"ftp_password,required"`
	// Access key ID for the S3 account.
	//
	// Access Key ID is 20 alpha-numeric characters like 022QF06E7MXBSH9DHM02
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 storage hostname.
	//
	// It is required if "`s3_type`": other.
	S3Hostname string `json:"s3_hostname,required"`
	// Secret access key for the S3 account.
	//
	// Secret Access Key is 20-50 alpha-numeric-slash-plus characters like
	// kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct
	S3SecretKey string `json:"s3_secret_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// SFTP storage hostname.
	SftpHostname string `json:"sftp_hostname,required"`
	// SFTP storage login.
	SftpLogin string `json:"sftp_login,required"`
	// SFTP storage password.
	//
	// It should be empty if "`sftp_private_key`" is set.
	SftpPassword string `json:"sftp_password,required"`
	// Storage type.
	//
	// Possible values:
	//
	// - **ftp**
	// - **sftp**
	// - **s3**
	StorageType string `json:"storage_type,required"`
	// The size of a single piece of the archive in MB. In case of **null** value logs
	// are delivered without slicing.
	ArchiveSizeMB param.Opt[int64] `json:"archive_size_mb,omitzero"`
	// Enables or disables a log forwarding feature.
	//
	// Possible values:
	//
	// - **true** - log forwarding feature is active.
	// - **false** - log forwarding feature is deactivated.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Name of the FTP prepend folder for log delivery.
	//
	// **Null** is allowed.
	FtpPrependFolder param.Opt[string] `json:"ftp_prepend_folder,omitzero"`
	// Enables or disables the forwarding of empty logs.
	//
	// Possible values:
	//
	// - **true** - Empty logs are not sent.
	// - **false** - Empty logs are sent.
	IgnoreEmptyLogs param.Opt[bool] `json:"ignore_empty_logs,omitzero"`
	// Amazon AWS region.
	S3AwsRegion param.Opt[int64] `json:"s3_aws_region,omitzero"`
	// Location of S3 storage.
	//
	// Restrictions:
	//
	//   - Maximum of 255 symbols.
	//   - Latin letters (A-Z, a-z), digits (0-9), dots, colons, dashes, and underscores
	//     (.:\_-).
	S3BucketLocation param.Opt[string] `json:"s3_bucket_location,omitzero"`
	// S3 bucket hostname.
	//
	// Restrictions:
	//
	// - Maximum of 255 symbols.
	// - Latin letters (A-Z, a-z,) digits (0-9,) dots, colons, dashes, and underscores.
	// - Required if "`s3_type`": other.
	S3HostBucket param.Opt[string] `json:"s3_host_bucket,omitzero"`
	// Passphrase for SFTP private key.
	//
	// Restrictions:
	//
	// - Should be set if private key encoded with passphrase.
	// - Should be empty if "`sftp_password`" is set.
	SftpKeyPassphrase param.Opt[string] `json:"sftp_key_passphrase,omitzero"`
	// Name of the SFTP prepend folder for log delivery.
	//
	// **Null** is allowed.
	SftpPrependFolder param.Opt[string] `json:"sftp_prepend_folder,omitzero"`
	// Private key for SFTP authorization.
	//
	// Possible values:
	//
	// - **RSA**
	// - **ED25519**
	//
	// It should be empty if "`sftp_password`" is set.
	SftpPrivateKey param.Opt[string] `json:"sftp_private_key,omitzero"`
	paramObj
}

func (r LogSettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogSettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogSettingNewParamsFolder struct {
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket ID.
	// - **ftp/sftp** - FTP/SFTP folder ID.
	ID param.Opt[int64] `json:"id,omitzero"`
	// S3 bucket name.
	//
	// The field is required if "`storage_type`": **s3**.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// CDN resource ID.
	CdnResource param.Opt[int64] `json:"cdn_resource,omitzero"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket sub-folder name (optional.)
	// - **ftp/sftp** - FTP/SFTP folder name (required.)
	Folder param.Opt[string] `json:"folder,omitzero"`
	paramObj
}

func (r LogSettingNewParamsFolder) MarshalJSON() (data []byte, err error) {
	type shadow LogSettingNewParamsFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogSettingNewParamsFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogSettingUpdateParams struct {
	// Name of the S3 bucket to which logs for all CDN resources are delivered.
	AllResourcesBucket string `json:"all_resources_bucket,required"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	//   - If "`storage_type`": s3 - Name of the S3 bucket sub-folder to which logs for
	//     all CDN resources are delivered.
	//   - If "`storage_type`": ftp/sftp - Name of the folder (or path) to which logs for
	//     all CDN resources are delivered.
	AllResourcesFolder string `json:"all_resources_folder,required"`
	// List of folders/buckets for receiving CDN resources logs.
	Folders []LogSettingUpdateParamsFolder `json:"folders,omitzero,required"`
	// Defines whether logs of all CDN resources are delivered to one folder/bucket or
	// to separate ones.
	//
	// Possible values:
	//
	//   - **true** - Logs of all CDN resources are delivered to one folder/bucket.
	//   - **false** - Logs of different CDN resources are delivered to separate
	//     folders/buckets.
	ForAllResources bool `json:"for_all_resources,required"`
	// FTP storage hostname.
	FtpHostname string `json:"ftp_hostname,required"`
	// FTP storage login.
	FtpLogin string `json:"ftp_login,required"`
	// FTP storage password.
	FtpPassword string `json:"ftp_password,required"`
	// Access key ID for the S3 account.
	//
	// Access Key ID is 20 alpha-numeric characters like 022QF06E7MXBSH9DHM02
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 storage hostname.
	//
	// It is required if "`s3_type`": other.
	S3Hostname string `json:"s3_hostname,required"`
	// Secret access key for the S3 account.
	//
	// Secret Access Key is 20-50 alpha-numeric-slash-plus characters like
	// kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct
	S3SecretKey string `json:"s3_secret_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// SFTP storage hostname.
	SftpHostname string `json:"sftp_hostname,required"`
	// SFTP storage login.
	SftpLogin string `json:"sftp_login,required"`
	// SFTP storage password.
	//
	// It should be empty if "`sftp_private_key`" is set.
	SftpPassword string `json:"sftp_password,required"`
	// Storage type.
	//
	// Possible values:
	//
	// - **ftp**
	// - **sftp**
	// - **s3**
	StorageType string `json:"storage_type,required"`
	// The size of a single piece of the archive in MB. In case of **null** value logs
	// are delivered without slicing.
	ArchiveSizeMB param.Opt[int64] `json:"archive_size_mb,omitzero"`
	// Enables or disables a log forwarding feature.
	//
	// Possible values:
	//
	// - **true** - log forwarding feature is active.
	// - **false** - log forwarding feature is deactivated.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Name of the FTP prepend folder for log delivery.
	//
	// **Null** is allowed.
	FtpPrependFolder param.Opt[string] `json:"ftp_prepend_folder,omitzero"`
	// Enables or disables the forwarding of empty logs.
	//
	// Possible values:
	//
	// - **true** - Empty logs are not sent.
	// - **false** - Empty logs are sent.
	IgnoreEmptyLogs param.Opt[bool] `json:"ignore_empty_logs,omitzero"`
	// Amazon AWS region.
	S3AwsRegion param.Opt[int64] `json:"s3_aws_region,omitzero"`
	// Location of S3 storage.
	//
	// Restrictions:
	//
	//   - Maximum of 255 symbols.
	//   - Latin letters (A-Z, a-z), digits (0-9), dots, colons, dashes, and underscores
	//     (.:\_-).
	S3BucketLocation param.Opt[string] `json:"s3_bucket_location,omitzero"`
	// S3 bucket hostname.
	//
	// Restrictions:
	//
	// - Maximum of 255 symbols.
	// - Latin letters (A-Z, a-z,) digits (0-9,) dots, colons, dashes, and underscores.
	// - Required if "`s3_type`": other.
	S3HostBucket param.Opt[string] `json:"s3_host_bucket,omitzero"`
	// Passphrase for SFTP private key.
	//
	// Restrictions:
	//
	// - Should be set if private key encoded with passphrase.
	// - Should be empty if "`sftp_password`" is set.
	SftpKeyPassphrase param.Opt[string] `json:"sftp_key_passphrase,omitzero"`
	// Name of the SFTP prepend folder for log delivery.
	//
	// **Null** is allowed.
	SftpPrependFolder param.Opt[string] `json:"sftp_prepend_folder,omitzero"`
	// Private key for SFTP authorization.
	//
	// Possible values:
	//
	// - **RSA**
	// - **ED25519**
	//
	// It should be empty if "`sftp_password`" is set.
	SftpPrivateKey param.Opt[string] `json:"sftp_private_key,omitzero"`
	paramObj
}

func (r LogSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogSettingUpdateParamsFolder struct {
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket ID.
	// - **ftp/sftp** - FTP/SFTP folder ID.
	ID param.Opt[int64] `json:"id,omitzero"`
	// S3 bucket name.
	//
	// The field is required if "`storage_type`": **s3**.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// CDN resource ID.
	CdnResource param.Opt[int64] `json:"cdn_resource,omitzero"`
	// Parameter meaning depends on the value of the "`storage_type`" value:
	//
	// - **s3** - S3 bucket sub-folder name (optional.)
	// - **ftp/sftp** - FTP/SFTP folder name (required.)
	Folder param.Opt[string] `json:"folder,omitzero"`
	paramObj
}

func (r LogSettingUpdateParamsFolder) MarshalJSON() (data []byte, err error) {
	type shadow LogSettingUpdateParamsFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogSettingUpdateParamsFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
