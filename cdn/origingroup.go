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
	return
}

// Change origin group
func (r *OriginGroupService) Update(ctx context.Context, originGroupID int64, body OriginGroupUpdateParams, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all origin groups and related origin sources.
func (r *OriginGroupService) List(ctx context.Context, query OriginGroupListParams, opts ...option.RequestOption) (res *OriginGroupsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/origin_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete origin group
func (r *OriginGroupService) Delete(ctx context.Context, originGroupID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get origin group details
func (r *OriginGroupService) Get(ctx context.Context, originGroupID int64, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change origin group
func (r *OriginGroupService) Replace(ctx context.Context, originGroupID int64, body OriginGroupReplaceParams, opts ...option.RequestOption) (res *OriginGroupsUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/origin_groups/%v", originGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
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
	Sources []OriginGroupsNoneAuthSource `json:"sources"`
	UseNext bool                         `json:"use_next"`
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
	// List of origin sources in the origin group.
	Sources []OriginGroupsNoneAuthSource `json:"sources"`
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

type OriginGroupsNoneAuthSource struct {
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
	// IP address or domain name of the origin and the port, if custom port is used.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Backup      respjson.Field
		Enabled     respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGroupsNoneAuthSource) RawJSON() string { return r.JSON.raw }
func (r *OriginGroupsNoneAuthSource) UnmarshalJSON(data []byte) error {
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
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 bucket name.
	//
	// Restrictions:
	//
	// - Maximum 128 characters.
	S3BucketName string `json:"s3_bucket_name,required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "`s3_type`": amazon, length should be 40 characters.
	//   - If "`s3_type`": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// S3 storage region.
	//
	// The parameter is required, if "`s3_type`": amazon.
	S3Region string `json:"s3_region"`
	// S3 storage hostname.
	//
	// The parameter is required, if "`s3_type`": other.
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

	// This field is a request body variant, only one variant field can be set.
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

// The properties Name, Sources are required.
type OriginGroupNewParamsBodyNoneAuth struct {
	// Origin group name.
	Name string `json:"name,required"`
	// List of origin sources in the origin group.
	Sources []OriginGroupNewParamsBodyNoneAuthSource `json:"sources,omitzero,required"`
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

type OriginGroupNewParamsBodyNoneAuthSource struct {
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
	// IP address or domain name of the origin and the port, if custom port is used.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r OriginGroupNewParamsBodyNoneAuthSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupNewParamsBodyNoneAuthSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupNewParamsBodyNoneAuthSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Auth, AuthType, Name are required.
type OriginGroupNewParamsBodyAwsSignatureV4 struct {
	// Credentials to access the private bucket.
	Auth OriginGroupNewParamsBodyAwsSignatureV4Auth `json:"auth,omitzero,required"`
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType string `json:"auth_type,required"`
	// Origin group name.
	Name string `json:"name,required"`
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
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 bucket name.
	//
	// Restrictions:
	//
	// - Maximum 128 characters.
	S3BucketName string `json:"s3_bucket_name,required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "`s3_type`": amazon, length should be 40 characters.
	//   - If "`s3_type`": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// S3 storage region.
	//
	// The parameter is required, if "`s3_type`": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "`s3_type`": other.
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
	Name string `json:"name,required"`
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
	ProxyNextUpstream []string `json:"proxy_next_upstream,omitzero"`
	// List of origin sources in the origin group.
	Sources []OriginGroupUpdateParamsBodyNoneAuthSource `json:"sources,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuth) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGroupUpdateParamsBodyNoneAuthSource struct {
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
	// IP address or domain name of the origin and the port, if custom port is used.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r OriginGroupUpdateParamsBodyNoneAuthSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupUpdateParamsBodyNoneAuthSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupUpdateParamsBodyNoneAuthSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 bucket name.
	//
	// Restrictions:
	//
	// - Maximum 128 characters.
	S3BucketName string `json:"s3_bucket_name,required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "`s3_type`": amazon, length should be 40 characters.
	//   - If "`s3_type`": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// S3 storage region.
	//
	// The parameter is required, if "`s3_type`": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "`s3_type`": other.
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
	AuthType string `json:"auth_type,required"`
	// Origin group name.
	Name string `json:"name,required"`
	// Parameter is **deprecated**.
	Path string `json:"path,required"`
	// List of origin sources in the origin group.
	Sources []OriginGroupReplaceParamsBodyNoneAuthSource `json:"sources,omitzero,required"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next,required"`
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

type OriginGroupReplaceParamsBodyNoneAuthSource struct {
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
	// IP address or domain name of the origin and the port, if custom port is used.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r OriginGroupReplaceParamsBodyNoneAuthSource) MarshalJSON() (data []byte, err error) {
	type shadow OriginGroupReplaceParamsBodyNoneAuthSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGroupReplaceParamsBodyNoneAuthSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Auth, AuthType, Name, Path, UseNext are required.
type OriginGroupReplaceParamsBodyAwsSignatureV4 struct {
	// Credentials to access the private bucket.
	Auth OriginGroupReplaceParamsBodyAwsSignatureV4Auth `json:"auth,omitzero,required"`
	// Authentication type.
	//
	// **awsSignatureV4** value is used for S3 storage.
	AuthType string `json:"auth_type,required"`
	// Origin group name.
	Name string `json:"name,required"`
	// Parameter is **deprecated**.
	Path string `json:"path,required"`
	// Defines whether to use the next origin from the origin group if origin responds
	// with the cases specified in `proxy_next_upstream`. If you enable it, you must
	// specify cases in `proxy_next_upstream`.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	UseNext bool `json:"use_next,required"`
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
	S3AccessKeyID string `json:"s3_access_key_id,required"`
	// S3 bucket name.
	//
	// Restrictions:
	//
	// - Maximum 128 characters.
	S3BucketName string `json:"s3_bucket_name,required"`
	// Secret access key for the S3 account.
	//
	// Restrictions:
	//
	//   - Latin letters (A-Z, a-z), numbers (0-9), pluses, slashes, dashes, colons and
	//     underscores.
	//   - If "`s3_type`": amazon, length should be 40 characters.
	//   - If "`s3_type`": other, length should be from 16 to 255 characters.
	S3SecretAccessKey string `json:"s3_secret_access_key,required"`
	// Storage type compatible with S3.
	//
	// Possible values:
	//
	// - **amazon** – AWS S3 storage.
	// - **other** – Other (not AWS) S3 compatible storage.
	S3Type string `json:"s3_type,required"`
	// S3 storage region.
	//
	// The parameter is required, if "`s3_type`": amazon.
	S3Region param.Opt[string] `json:"s3_region,omitzero"`
	// S3 storage hostname.
	//
	// The parameter is required, if "`s3_type`": other.
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
