// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CDN origin groups aggregate one or more origin servers with failover and load
// balancing for content delivery.
//
// OriginGroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginGroupService] method instead.
type OriginGroupService struct {
	Options []option.RequestOption
}

// NewOriginGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOriginGroupService(opts ...option.RequestOption) (r OriginGroupService) {
	r = OriginGroupService{}
	r.Options = opts
	return
}

// Create an origin group with one or more origin sources.
func (r *OriginGroupService) New(ctx context.Context, body OriginGroupNewParams, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/origin_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change origin group
func (r *OriginGroupService) Update(ctx context.Context, originGroupID int64, body OriginGroupUpdateParams, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get all origin groups and related origin sources.
func (r *OriginGroupService) List(ctx context.Context, query OriginGroupListParams, opts ...option.RequestOption) (res *OriginGroupsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/origin_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete origin group
func (r *OriginGroupService) Delete(ctx context.Context, originGroupID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get origin group details
func (r *OriginGroupService) Get(ctx context.Context, originGroupID int64, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Change origin group
func (r *OriginGroupService) Replace(ctx context.Context, originGroupID int64, body OriginGroupReplaceParams, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// OriginGroupsUnion contains all possible properties and values from
// [OriginGroupsNoneAuth], [OriginGroupsAwsSignatureV4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OriginGroupsUnion struct {
	ID                  int64    `json:"id"`
	AuthType            string   `json:"auth_type"`
	HasRelatedResources bool     `json:"has_related_resources"`
	Name                string   `json:"name"`
	Path                string   `json:"path"`
	ProxyNextUpstream   []string `json:"proxy_next_upstream"`
	// This field is from variant [OriginGroupsNoneAuth].
	Sources []OriginGroupsNoneAuthSourceUnion `json:"sources"`
	UseNext bool                              `json:"use_next"`
	// This field is from variant [OriginGroupsAwsSignatureV4].
	Auth OriginGroupsAwsSignatureV4Auth `json:"auth"`
	JSON struct {
		ID                  respjson.Field
		AuthType            respjson.Field
		HasRelatedResources respjson.Field
		Name                respjson.Field
		Path                respjson.Field
		ProxyNextUpstream   respjson.Field
		Sources             respjson.Field
		UseNext             respjson.Field
		Auth                respjson.Field
		raw                 string
	} `json:"-"`
}

func (u OriginGroupsUnion) AsNoneAuth() (v OriginGroupsNoneAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginGroupsUnion) AsAwsSignatureV4() (v OriginGroupsAwsSignatureV4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OriginGroupsUnion) RawJSON() string { return u.JSON.raw }

func (r *OriginGroupsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin group with host origins, or mixed host and S3 origins using inline
// `origin_type` and `config` in sources.
type OriginGroupsNoneAuth struct {
	// Origin group ID.
	ID int64 `json:"id"`
	// Origin authentication type.
	//
	// Possible values:
	//
	// - **none** - Used for public origins.
	// - **awsSignatureV4** - Used for S3 storage.
	AuthType string `json:"auth_type"`
	// Defines whether the origin group has related CDN resources.
	//
	// Possible values:
	//
	// - **true** - Origin group has related CDN resources.
	// - **false** - Origin group does not have related CDN resources.
	HasRelatedResources bool `json:"has_related_resources"`
	// Origin group name.
	Name string `json:"name"`
	// Parameter is **deprecated**.
	Path string `json:"path"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream"`
	// List of origin sources in the origin group. Each entry can be a host origin or
	// an S3 origin.
	//
	// Host origins have a `source` field with the hostname or IP. S3 origins have
	// `origin_type: s3` and a `config` object with S3 credentials. Both types can be
	// mixed in the same origin group.
	Sources []OriginGroupsNoneAuthSourceUnion `json:"sources"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AuthType            respjson.Field
		HasRelatedResources respjson.Field
		Name                respjson.Field
		Path                respjson.Field
		ProxyNextUpstream   respjson.Field
		Sources             respjson.Field
		UseNext             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsNoneAuth) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsNoneAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OriginGroupsNoneAuthSourceUnion contains all possible properties and values from
// [OriginGroupsNoneAuthSourceHostSource], [OriginGroupsNoneAuthSourceS3Source].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OriginGroupsNoneAuthSourceUnion struct {
	// This field is from variant [OriginGroupsNoneAuthSourceHostSource].
	Source             string `json:"source"`
	Backup             bool   `json:"backup"`
	Enabled            bool   `json:"enabled"`
	HostHeaderOverride string `json:"host_header_override"`
	Tag                string `json:"tag"`
	// This field is from variant [OriginGroupsNoneAuthSourceS3Source].
	Config OriginGroupsNoneAuthSourceS3SourceConfig `json:"config"`
	// This field is from variant [OriginGroupsNoneAuthSourceS3Source].
	OriginType string `json:"origin_type"`
	JSON       struct {
		Source             respjson.Field
		Backup             respjson.Field
		Enabled            respjson.Field
		HostHeaderOverride respjson.Field
		Tag                respjson.Field
		Config             respjson.Field
		OriginType         respjson.Field
		raw                string
	} `json:"-"`
}

func (u OriginGroupsNoneAuthSourceUnion) AsHostSource() (v OriginGroupsNoneAuthSourceHostSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginGroupsNoneAuthSourceUnion) AsS3Source() (v OriginGroupsNoneAuthSourceS3Source) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OriginGroupsNoneAuthSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *OriginGroupsNoneAuthSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A host origin source.
type OriginGroupsNoneAuthSourceHostSource struct {
	// IP address or domain name of the origin and the port, if custom port is used.
	Source string `json:"source" api:"required"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup bool `json:"backup"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled bool `json:"enabled"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride string `json:"host_header_override" api:"nullable"`
	// Tag for the origin source.
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source             respjson.Field
		Backup             respjson.Field
		Enabled            respjson.Field
		HostHeaderOverride respjson.Field
		Tag                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsNoneAuthSourceHostSource) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsNoneAuthSourceHostSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An S3 origin source.
type OriginGroupsNoneAuthSourceS3Source struct {
	// S3 storage configuration. Required when `origin_type` is `s3`.
	Config OriginGroupsNoneAuthSourceS3SourceConfig `json:"config" api:"required"`
	// Origin type. Present in responses only for S3 sources.
	//
	// Possible values:
	//
	// - **host** - A source server or endpoint from which content is fetched.
	// - **s3** - S3 storage with either AWS v4 authentication or public access.
	//
	// Any of "host", "s3".
	OriginType string `json:"origin_type" api:"required"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup bool `json:"backup"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled bool `json:"enabled"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride string `json:"host_header_override" api:"nullable"`
	// Tag for the origin source.
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config             respjson.Field
		OriginType         respjson.Field
		Backup             respjson.Field
		Enabled            respjson.Field
		HostHeaderOverride respjson.Field
		Tag                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsNoneAuthSourceS3Source) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsNoneAuthSourceS3Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// S3 storage configuration. Required when `origin_type` is `s3`.
type OriginGroupsNoneAuthSourceS3SourceConfig struct {
	// Access key ID for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 4 to 255 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - From 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** - AWS S3 storage.
	// - **other** - Other (not AWS) S3 compatible storage.
	//
	// Any of "amazon", "other".
	S3Type string `json:"s3_type" api:"required"`
	// S3 authentication type.
	S3AuthType string `json:"s3_auth_type"`
	// S3 storage region.
	//
	// The parameter is required if `s3_type` is `amazon`.
	S3Region string `json:"s3_region" api:"nullable"`
	// S3 storage hostname.
	//
	// The parameter is required if `s3_type` is `other`.
	S3StorageHostname string `json:"s3_storage_hostname" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3AccessKeyID     respjson.Field
		S3BucketName      respjson.Field
		S3SecretAccessKey respjson.Field
		S3Type            respjson.Field
		S3AuthType        respjson.Field
		S3Region          respjson.Field
		S3StorageHostname respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsNoneAuthSourceS3SourceConfig) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsNoneAuthSourceS3SourceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGroupsAwsSignatureV4 struct {
	// Origin group ID.
	ID int64 `json:"id"`
	// Credentials to access the private bucket.
	Auth OriginGroupsAwsSignatureV4Auth `json:"auth"`
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType string `json:"auth_type"`
	// Defines whether the origin group has related CDN resources.
	//
	// Possible values:
	//
	// - **true** - Origin group has related CDN resources.
	// - **false** - Origin group does not have related CDN resources.
	HasRelatedResources bool `json:"has_related_resources"`
	// Origin group name.
	Name string `json:"name"`
	// Parameter is **deprecated**.
	Path string `json:"path"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Auth                respjson.Field
		AuthType            respjson.Field
		HasRelatedResources respjson.Field
		Name                respjson.Field
		Path                respjson.Field
		ProxyNextUpstream   respjson.Field
		UseNext             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsAwsSignatureV4) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsAwsSignatureV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credentials to access the private bucket.
type OriginGroupsAwsSignatureV4Auth struct {
	// Access key ID for the S3 account.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 3 to 512 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "s3_type": amazon, length should be 40 characters.
	//   - If "s3_type": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type" api:"required"`
	// S3 storage region.
	//
	// The parameter is required, if "s3_type": amazon.
	S3Region string `json:"s3_region"`
	// S3 storage hostname.
	//
	// The parameter is required, if "s3_type": other.
	S3StorageHostname string `json:"s3_storage_hostname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3AccessKeyID     respjson.Field
		S3BucketName      respjson.Field
		S3SecretAccessKey respjson.Field
		S3Type            respjson.Field
		S3Region          respjson.Field
		S3StorageHostname respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsAwsSignatureV4Auth) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsAwsSignatureV4Auth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGroupsList []OriginGroupsUnion

type OriginGroupNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// an origin group with host origins, or mixed host and S3 origins.
	OfNoneAuth *OriginGroupNewParamsBodyNoneAuth `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAwsSignatureV4 *OriginGroupNewParamsBodyAwsSignatureV4 `json:",inline"`

	paramObj
}

func (u OriginGroupNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoneAuth, u.OfAwsSignatureV4)
}
func (r *OriginGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create an origin group with host origins, or mixed host and S3 origins.
//
// The properties Name, Sources are required.
type OriginGroupNewParamsBodyNoneAuth struct {
	// Origin group name.
	Name    string                                        `json:"name" api:"required"`
	Sources []OriginGroupNewParamsBodyNoneAuthSourceUnion `json:"sources,omitzero" api:"required"`
	// Origin authentication type.
	//
	// Possible values:
	//
	// - **none** - Used for public origins.
	// - **awsSignatureV4** - Used for S3 storage.
	AuthType param.Opt[string] `json:"auth_type,omitzero"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext param.Opt[bool] `json:"use_next,omitzero"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyNoneAuth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyNoneAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyNoneAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OriginGroupNewParamsBodyNoneAuthSourceUnion struct {
	OfHostSource *OriginGroupNewParamsBodyNoneAuthSourceHostSource `json:",omitzero,inline"`
	OfS3Source   *OriginGroupNewParamsBodyNoneAuthSourceS3Source   `json:",omitzero,inline"`
	paramUnion
}

func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHostSource, u.OfS3Source)
}
func (u *OriginGroupNewParamsBodyNoneAuthSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OriginGroupNewParamsBodyNoneAuthSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfHostSource) {
		return u.OfHostSource
	} else if !param.IsOmitted(u.OfS3Source) {
		return u.OfS3Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetSource() *string {
	if vt := u.OfHostSource; vt != nil {
		return &vt.Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetConfig() *OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig {
	if vt := u.OfS3Source; vt != nil {
		return &vt.Config
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetOriginType() *string {
	if vt := u.OfS3Source; vt != nil {
		return &vt.OriginType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetBackup() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetEnabled() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetHostHeaderOverride() *string {
	if vt := u.OfHostSource; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	} else if vt := u.OfS3Source; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupNewParamsBodyNoneAuthSourceUnion) GetTag() *string {
	if vt := u.OfHostSource; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	}
	return nil
}

// The property Source is required.
type OriginGroupNewParamsBodyNoneAuthSourceHostSource struct {
	// IP address or domain name of the origin and the port, if custom port is used.
	Source string `json:"source" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyNoneAuthSourceHostSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyNoneAuthSourceHostSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyNoneAuthSourceHostSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Config is required.
type OriginGroupNewParamsBodyNoneAuthSourceS3Source struct {
	// S3 storage configuration. Required when `origin_type` is `s3`.
	Config OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig `json:"config,omitzero" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Origin type. Present in responses only for S3 sources.
	//
	// Possible values:
	//
	// - **host** - A source server or endpoint from which content is fetched.
	// - **s3** - S3 storage with either AWS v4 authentication or public access.
	//
	// Any of "host", "s3".
	OriginType string `json:"origin_type,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyNoneAuthSourceS3Source) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyNoneAuthSourceS3Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyNoneAuthSourceS3Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupNewParamsBodyNoneAuthSourceS3Source](
		"origin_type", "host", "s3",
	)
}

// S3 storage configuration. Required when `origin_type` is `s3`.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig struct {
	// Access key ID for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 4 to 255 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - From 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** - AWS S3 storage.
	// - **other** - Other (not AWS) S3 compatible storage.
	//
	// Any of "amazon", "other".
	S3Type string `json:"s3_type,omitzero" api:"required"`
	// S3 storage region.
	//
	// The parameter is required if `s3_type` is `amazon`.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required if `s3_type` is `other`.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	// S3 authentication type.
	S3AuthType param.Opt[string] `json:"s3_auth_type,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupNewParamsBodyNoneAuthSourceS3SourceConfig](
		"s3_type", "amazon", "other",
	)
}

// The properties Auth, AuthType, Name are required.
type OriginGroupNewParamsBodyAwsSignatureV4 struct {
	// Credentials to access the private bucket.
	Auth OriginGroupNewParamsBodyAwsSignatureV4Auth `json:"auth,omitzero" api:"required"`
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType string `json:"auth_type" api:"required"`
	// Origin group name.
	Name string `json:"name" api:"required"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext param.Opt[bool] `json:"use_next,omitzero"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyAwsSignatureV4) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyAwsSignatureV4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyAwsSignatureV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credentials to access the private bucket.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupNewParamsBodyAwsSignatureV4Auth struct {
	// Access key ID for the S3 account.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 3 to 512 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "s3_type": amazon, length should be 40 characters.
	//   - If "s3_type": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type" api:"required"`
	// S3 storage region.
	//
	// The parameter is required, if "s3_type": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "s3_type": other.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyAwsSignatureV4Auth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyAwsSignatureV4Auth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyAwsSignatureV4Auth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGroupUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfNoneAuth *OriginGroupUpdateParamsBodyNoneAuth `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAwsSignatureV4 *OriginGroupUpdateParamsBodyAwsSignatureV4 `json:",inline"`

	paramObj
}

func (u OriginGroupUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoneAuth, u.OfAwsSignatureV4)
}
func (r *OriginGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type OriginGroupUpdateParamsBodyNoneAuth struct {
	// Origin group name.
	Name string `json:"name" api:"required"`
	// Origin authentication type.
	//
	// Possible values:
	//
	// - **none** - Used for public origins.
	// - **awsSignatureV4** - Used for S3 storage.
	AuthType param.Opt[string] `json:"auth_type,omitzero"`
	// Parameter is **deprecated**.
	Path param.Opt[string] `json:"path,omitzero"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext param.Opt[bool] `json:"use_next,omitzero"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string                                         `json:"proxy_next_upstream,omitzero"`
	Sources           []OriginGroupUpdateParamsBodyNoneAuthSourceUnion `json:"sources,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OriginGroupUpdateParamsBodyNoneAuthSourceUnion struct {
	OfHostSource *OriginGroupUpdateParamsBodyNoneAuthSourceHostSource `json:",omitzero,inline"`
	OfS3Source   *OriginGroupUpdateParamsBodyNoneAuthSourceS3Source   `json:",omitzero,inline"`
	paramUnion
}

func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHostSource, u.OfS3Source)
}
func (u *OriginGroupUpdateParamsBodyNoneAuthSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OriginGroupUpdateParamsBodyNoneAuthSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfHostSource) {
		return u.OfHostSource
	} else if !param.IsOmitted(u.OfS3Source) {
		return u.OfS3Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetSource() *string {
	if vt := u.OfHostSource; vt != nil {
		return &vt.Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetConfig() *OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig {
	if vt := u.OfS3Source; vt != nil {
		return &vt.Config
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetOriginType() *string {
	if vt := u.OfS3Source; vt != nil {
		return &vt.OriginType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetBackup() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetEnabled() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetHostHeaderOverride() *string {
	if vt := u.OfHostSource; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	} else if vt := u.OfS3Source; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupUpdateParamsBodyNoneAuthSourceUnion) GetTag() *string {
	if vt := u.OfHostSource; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	}
	return nil
}

// The property Source is required.
type OriginGroupUpdateParamsBodyNoneAuthSourceHostSource struct {
	// IP address or domain name of the origin and the port, if custom port is used.
	Source string `json:"source" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuthSourceHostSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuthSourceHostSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuthSourceHostSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Config is required.
type OriginGroupUpdateParamsBodyNoneAuthSourceS3Source struct {
	// S3 storage configuration. Required when `origin_type` is `s3`.
	Config OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig `json:"config,omitzero" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Origin type. Present in responses only for S3 sources.
	//
	// Possible values:
	//
	// - **host** - A source server or endpoint from which content is fetched.
	// - **s3** - S3 storage with either AWS v4 authentication or public access.
	//
	// Any of "host", "s3".
	OriginType string `json:"origin_type,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuthSourceS3Source) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuthSourceS3Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuthSourceS3Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupUpdateParamsBodyNoneAuthSourceS3Source](
		"origin_type", "host", "s3",
	)
}

// S3 storage configuration. Required when `origin_type` is `s3`.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig struct {
	// Access key ID for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 4 to 255 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - From 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** - AWS S3 storage.
	// - **other** - Other (not AWS) S3 compatible storage.
	//
	// Any of "amazon", "other".
	S3Type string `json:"s3_type,omitzero" api:"required"`
	// S3 storage region.
	//
	// The parameter is required if `s3_type` is `amazon`.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required if `s3_type` is `other`.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	// S3 authentication type.
	S3AuthType param.Opt[string] `json:"s3_auth_type,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupUpdateParamsBodyNoneAuthSourceS3SourceConfig](
		"s3_type", "amazon", "other",
	)
}

type OriginGroupUpdateParamsBodyAwsSignatureV4 struct {
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType param.Opt[string] `json:"auth_type,omitzero"`
	// Origin group name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Parameter is **deprecated**.
	Path param.Opt[string] `json:"path,omitzero"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext param.Opt[bool] `json:"use_next,omitzero"`
	// Credentials to access the private bucket.
	Auth OriginGroupUpdateParamsBodyAwsSignatureV4Auth `json:"auth,omitzero"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyAwsSignatureV4) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyAwsSignatureV4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyAwsSignatureV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credentials to access the private bucket.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupUpdateParamsBodyAwsSignatureV4Auth struct {
	// Access key ID for the S3 account.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 3 to 512 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "s3_type": amazon, length should be 40 characters.
	//   - If "s3_type": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type" api:"required"`
	// S3 storage region.
	//
	// The parameter is required, if "s3_type": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "s3_type": other.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyAwsSignatureV4Auth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyAwsSignatureV4Auth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyAwsSignatureV4Auth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGroupListParams struct {
	// Defines whether the origin group has related CDN resources.
	//
	// Possible values:
	//
	// - **true** – Origin group has related CDN resources.
	// - **false** – Origin group does not have related CDN resources.
	HasRelatedResources param.Opt[bool] `query:"has_related_resources,omitzero" json:"-"`
	// Origin group name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Origin sources (IP addresses or domains) in the origin group.
	Sources param.Opt[string] `query:"sources,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OriginGroupListParams]'s query parameters as `url.Values`.
func (r OriginGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OriginGroupReplaceParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfNoneAuth *OriginGroupReplaceParamsBodyNoneAuth `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAwsSignatureV4 *OriginGroupReplaceParamsBodyAwsSignatureV4 `json:",inline"`

	paramObj
}

func (u OriginGroupReplaceParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoneAuth, u.OfAwsSignatureV4)
}
func (r *OriginGroupReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AuthType, Name, Path, Sources, UseNext are required.
type OriginGroupReplaceParamsBodyNoneAuth struct {
	// Origin authentication type.
	//
	// Possible values:
	//
	// - **none** - Used for public origins.
	// - **awsSignatureV4** - Used for S3 storage.
	AuthType string `json:"auth_type" api:"required"`
	// Origin group name.
	Name string `json:"name" api:"required"`
	// Parameter is **deprecated**.
	Path    string                                            `json:"path" api:"required"`
	Sources []OriginGroupReplaceParamsBodyNoneAuthSourceUnion `json:"sources,omitzero" api:"required"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next" api:"required"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyNoneAuth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyNoneAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyNoneAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OriginGroupReplaceParamsBodyNoneAuthSourceUnion struct {
	OfHostSource *OriginGroupReplaceParamsBodyNoneAuthSourceHostSource `json:",omitzero,inline"`
	OfS3Source   *OriginGroupReplaceParamsBodyNoneAuthSourceS3Source   `json:",omitzero,inline"`
	paramUnion
}

func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHostSource, u.OfS3Source)
}
func (u *OriginGroupReplaceParamsBodyNoneAuthSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OriginGroupReplaceParamsBodyNoneAuthSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfHostSource) {
		return u.OfHostSource
	} else if !param.IsOmitted(u.OfS3Source) {
		return u.OfS3Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetSource() *string {
	if vt := u.OfHostSource; vt != nil {
		return &vt.Source
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetConfig() *OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig {
	if vt := u.OfS3Source; vt != nil {
		return &vt.Config
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetOriginType() *string {
	if vt := u.OfS3Source; vt != nil {
		return &vt.OriginType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetBackup() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Backup.Valid() {
		return &vt.Backup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetEnabled() *bool {
	if vt := u.OfHostSource; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Enabled.Valid() {
		return &vt.Enabled.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetHostHeaderOverride() *string {
	if vt := u.OfHostSource; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	} else if vt := u.OfS3Source; vt != nil && vt.HostHeaderOverride.Valid() {
		return &vt.HostHeaderOverride.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginGroupReplaceParamsBodyNoneAuthSourceUnion) GetTag() *string {
	if vt := u.OfHostSource; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	} else if vt := u.OfS3Source; vt != nil && vt.Tag.Valid() {
		return &vt.Tag.Value
	}
	return nil
}

// The property Source is required.
type OriginGroupReplaceParamsBodyNoneAuthSourceHostSource struct {
	// IP address or domain name of the origin and the port, if custom port is used.
	Source string `json:"source" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyNoneAuthSourceHostSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyNoneAuthSourceHostSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyNoneAuthSourceHostSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Config is required.
type OriginGroupReplaceParamsBodyNoneAuthSourceS3Source struct {
	// S3 storage configuration. Required when `origin_type` is `s3`.
	Config OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig `json:"config,omitzero" api:"required"`
	// Per-origin Host header override. When set, the CDN sends this value as the Host
	// header when requesting content from this origin instead of the default.
	HostHeaderOverride param.Opt[string] `json:"host_header_override,omitzero"`
	// Defines whether the origin is a backup, meaning that it will not be used until
	// one of active origins become unavailable.
	//
	// Possible values:
	//
	// - **true** - Origin is a backup.
	// - **false** - Origin is not a backup.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Enables or disables an origin source in the origin group.
	//
	// Possible values:
	//
	// - **true** - Origin is enabled and the CDN uses it to pull content.
	// - **false** - Origin is disabled and the CDN does not use it to pull content.
	//
	// Origin group must contain at least one enabled origin.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Tag for the origin source.
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Origin type. Present in responses only for S3 sources.
	//
	// Possible values:
	//
	// - **host** - A source server or endpoint from which content is fetched.
	// - **s3** - S3 storage with either AWS v4 authentication or public access.
	//
	// Any of "host", "s3".
	OriginType string `json:"origin_type,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyNoneAuthSourceS3Source) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyNoneAuthSourceS3Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyNoneAuthSourceS3Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupReplaceParamsBodyNoneAuthSourceS3Source](
		"origin_type", "host", "s3",
	)
}

// S3 storage configuration. Required when `origin_type` is `s3`.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig struct {
	// Access key ID for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 4 to 255 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account. Masked as `SECRET_VALUE` in responses.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - From 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** - AWS S3 storage.
	// - **other** - Other (not AWS) S3 compatible storage.
	//
	// Any of "amazon", "other".
	S3Type string `json:"s3_type,omitzero" api:"required"`
	// S3 storage region.
	//
	// The parameter is required if `s3_type` is `amazon`.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required if `s3_type` is `other`.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	// S3 authentication type.
	S3AuthType param.Opt[string] `json:"s3_auth_type,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OriginGroupReplaceParamsBodyNoneAuthSourceS3SourceConfig](
		"s3_type", "amazon", "other",
	)
}

// The properties Auth, AuthType, Name, Path, UseNext are required.
type OriginGroupReplaceParamsBodyAwsSignatureV4 struct {
	// Credentials to access the private bucket.
	Auth OriginGroupReplaceParamsBodyAwsSignatureV4Auth `json:"auth,omitzero" api:"required"`
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType string `json:"auth_type" api:"required"`
	// Origin group name.
	Name string `json:"name" api:"required"`
	// Parameter is **deprecated**.
	Path string `json:"path" api:"required"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next" api:"required"`
	// Defines cases when the request should be passed on to the next origin.
	//
	// Possible values:
	//
	//   - **error** - an error occurred while establishing a connection with the origin,
	//     passing a request to it, or reading the response header
	//   - **timeout** - a timeout has occurred while establishing a connection with the
	//     origin, passing a request to it, or reading the response header
	//   - **`invalid_header`** - a origin returned an empty or invalid response
	//   - **`http_403`** - a origin returned a response with the code 403
	//   - **`http_404`** - a origin returned a response with the code 404
	//   - **`http_429`** - a origin returned a response with the code 429
	//   - **`http_500`** - a origin returned a response with the code 500
	//   - **`http_502`** - a origin returned a response with the code 502
	//   - **`http_503`** - a origin returned a response with the code 503
	//   - **`http_504`** - a origin returned a response with the code 504
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyAwsSignatureV4) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyAwsSignatureV4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyAwsSignatureV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credentials to access the private bucket.
//
// The properties S3AccessKeyID, S3BucketName, S3SecretAccessKey, S3Type are
// required.
type OriginGroupReplaceParamsBodyAwsSignatureV4Auth struct {
	// Access key ID for the S3 account.
	//
	// Restrictions:
	//
	// - Latin letters (A-Z, a-z), numbers (0-9), colon, dash, and underscore.
	// - From 3 to 512 characters.
	S3AccessKeyID string `json:"s3_access_key_id" api:"required"`
	// S3 bucket name.
	S3BucketName string `json:"s3_bucket_name" api:"required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "s3_type": amazon, length should be 40 characters.
	//   - If "s3_type": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key" api:"required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type" api:"required"`
	// S3 storage region.
	//
	// The parameter is required, if "s3_type": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "s3_type": other.
	S3StorageHostname param.Opt[string] `json:"s3_storage_hostname,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyAwsSignatureV4Auth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyAwsSignatureV4Auth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyAwsSignatureV4Auth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
