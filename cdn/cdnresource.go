// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
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

// CDNResourceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCDNResourceService] method instead.
type CDNResourceService struct {
	Options []option.RequestOption
	Shield  CDNResourceShieldService
	Rules   CDNResourceRuleService
}

// NewCDNResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCDNResourceService(opts ...option.RequestOption) (r CDNResourceService) {
	r = CDNResourceService{}
	r.Options = opts
	r.Shield = NewCDNResourceShieldService(opts...)
	r.Rules = NewCDNResourceRuleService(opts...)
	return
}

// Create CDN resource
func (r *CDNResourceService) New(ctx context.Context, body CDNResourceNewParams, opts ...option.RequestOption) (res *CDNResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change CDN resource
func (r *CDNResourceService) Update(ctx context.Context, resourceID int64, body CDNResourceUpdateParams, opts ...option.RequestOption) (res *CDNResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get information about all CDN resources in your account.
func (r *CDNResourceService) List(ctx context.Context, query CDNResourceListParams, opts ...option.RequestOption) (res *CDNResourceList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the CDN resource from the system permanently.
//
// Notes:
//
//   - **Deactivation Requirement**: Set the `active` attribute to `false` before
//     deletion.
//   - **Statistics Availability**: Statistics will be available for **365 days**
//     after deletion through the
//     [statistics endpoints](/docs/api-reference/cdn/cdn-statistics/cdn-resource-statistics).
//   - **Irreversibility**: This action is irreversible. Once deleted, the CDN
//     resource cannot be recovered.
func (r *CDNResourceService) Delete(ctx context.Context, resourceID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get CDN resource details
func (r *CDNResourceService) Get(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *CDNResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Pre-populate files to a CDN cache before users requests. Prefetch is recommended
// only for files that **more than 200 MB** and **less than 5 GB**.
//
// You can make one prefetch request for a CDN resource per minute. One request for
// prefetch may content only up to 100 paths to files.
//
// The time of procedure depends on the number and size of the files.
//
// If you need to update files stored in the CDN, first purge these files and then
// prefetch.
func (r *CDNResourceService) Prefetch(ctx context.Context, resourceID int64, body CDNResourcePrefetchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/prefetch", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Check whether a Let's Encrypt certificate can be issued for the CDN resource.
func (r *CDNResourceService) PrevalidateSslLeCertificate(ctx context.Context, resourceID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/ssl/le/pre-validate", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Delete cache from CDN servers. This is necessary to update CDN content.
//
// We have different limits for different purge types:
//
//   - **Purge all cache** - One purge request for a CDN resource per minute.
//   - **Purge by URL** - Two purge requests for a CDN resource per minute. One purge
//     request is limited to 100 URLs.
//   - **Purge by pattern** - One purge request for a CDN resource per minute. One
//     purge request is limited to 10 patterns.
func (r *CDNResourceService) Purge(ctx context.Context, resourceID int64, body CDNResourcePurgeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/purge", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Change CDN resource
func (r *CDNResourceService) Replace(ctx context.Context, resourceID int64, body CDNResourceReplaceParams, opts ...option.RequestOption) (res *CDNResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CDNResource struct {
	// CDN resource ID.
	ID int64 `json:"id"`
	// Enables or disables a CDN resource.
	//
	// Possible values:
	//
	// - **true** - CDN resource is active. Content is being delivered.
	// - **false** - CDN resource is deactivated. Content is not being delivered.
	Active bool `json:"active"`
	// Defines whether the CDN resource can be used for purge by URLs feature.
	//
	// It's available only in case the CDN resource has enabled `ignore_vary_header`
	// option.
	CanPurgeByURLs bool `json:"can_purge_by_urls"`
	// ID of an account to which the CDN resource belongs.
	Client int64 `json:"client"`
	// Delivery domains that will be used for content delivery through a CDN.
	//
	// Delivery domains should be added to your DNS settings.
	Cname string `json:"cname"`
	// Date of CDN resource creation.
	Created string `json:"created"`
	// Defines whether CDN resource has been deleted.
	//
	// Possible values:
	//
	// - **true** - CDN resource is deleted.
	// - **false** - CDN resource is not deleted.
	Deleted bool `json:"deleted"`
	// Optional comment describing the CDN resource.
	Description string `json:"description"`
	// Enables or disables a CDN resource change by a user.
	//
	// Possible values:
	//
	//   - **true** - CDN resource is enabled and can be changed. Content can be
	//     delivered.
	//   - **false** - CDN resource is disabled and cannot be changed. Content can not be
	//     delivered.
	Enabled bool `json:"enabled"`
	// Defines whether the CDN resource has a custom configuration.
	//
	// Possible values:
	//
	//   - **true** - CDN resource has a custom configuration. You cannot change resource
	//     settings, except for the SSL certificate. To change other settings, contact
	//     technical support.
	//   - **false** - CDN resource has a regular configuration. You can change CDN
	//     resource settings.
	FullCustomEnabled bool `json:"full_custom_enabled"`
	// Defines whether a CDN resource has a cache zone shared with other CDN resources.
	//
	// Possible values:
	//
	//   - **true** - CDN resource is main and has a shared caching zone with other CDN
	//     resources, which are called reserve.
	//   - **false** - CDN resource is reserve and it has a shared caching zone with the
	//     main CDN resource. You cannot change some options, create rules, set up origin
	//     shielding and use the reserve resource for Streaming.
	//   - **null** - CDN resource does not have a shared cache zone.
	//
	// The main CDN resource is specified in the `primary_resource` field. It cannot be
	// suspended unless all related reserve CDN resources are suspended.
	IsPrimary bool `json:"is_primary" api:"nullable"`
	// CDN resource name.
	Name string `json:"name" api:"nullable"`
	// List of options that can be configured for the CDN resource.
	//
	// In case of `null` value the option is not added to the CDN resource. Option may
	// inherit its value from the global account settings.
	Options CDNResourceOptions `json:"options"`
	// Origin group ID with which the CDN resource is associated.
	OriginGroup int64 `json:"originGroup"`
	// Origin group name.
	OriginGroupName string `json:"originGroup_name"`
	// Protocol used by CDN servers to request content from an origin source.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
	//   - **HTTP** - CDN servers will connect to the origin via HTTP.
	//   - **MATCH** - connection protocol will be chosen automatically (content on the
	//     origin source should be available for the CDN both through HTTP and HTTPS).
	//
	// If protocol is not specified, HTTP is used to connect to an origin server.
	//
	// Any of "HTTP", "HTTPS", "MATCH".
	OriginProtocol CDNResourceOriginProtocol `json:"originProtocol"`
	// Defines whether the CDN resource has a preset applied.
	//
	// Possible values:
	//
	//   - **true** - CDN resource has a preset applied. CDN resource options included in
	//     the preset cannot be edited.
	//   - **false** - CDN resource does not have a preset applied.
	PresetApplied bool `json:"preset_applied"`
	// ID of the main CDN resource which has a shared caching zone with a reserve CDN
	// resource.
	//
	// If the parameter is not empty, then the current CDN resource is the reserve. You
	// cannot change some options, create rules, set up origin shielding, or use the
	// reserve CDN resource for Streaming.
	PrimaryResource int64 `json:"primary_resource" api:"nullable"`
	// ID of the trusted CA certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslCa int64 `json:"proxy_ssl_ca" api:"nullable"`
	// ID of the SSL certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslData int64 `json:"proxy_ssl_data" api:"nullable"`
	// Enables or disables SSL certificate validation of the origin server before
	// completing any connection.
	//
	// Possible values:
	//
	// - **true** - Origin SSL certificate validation is enabled.
	// - **false** - Origin SSL certificate validation is disabled.
	ProxySslEnabled bool `json:"proxy_ssl_enabled"`
	// Rules configured for the CDN resource.
	Rules []any `json:"rules"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames" format:"domain"`
	// Name of the origin shielding location data center.
	//
	// Parameter returns **null** if origin shielding is disabled.
	ShieldDc string `json:"shield_dc" api:"nullable"`
	// Defines whether origin shield is active and working for the CDN resource.
	//
	// Possible values:
	//
	// - **true** - Origin shield is active.
	// - **false** - Origin shield is not active.
	ShieldEnabled bool `json:"shield_enabled"`
	// Defines whether the origin shield with a dynamic location is enabled for the CDN
	// resource.
	//
	// To manage origin shielding, you must contact customer support.
	ShieldRoutingMap int64 `json:"shield_routing_map" api:"nullable"`
	// Defines whether origin shielding feature is enabled for the resource.
	//
	// Possible values:
	//
	// - **true** - Origin shielding is enabled.
	// - **false** - Origin shielding is disabled.
	Shielded bool `json:"shielded"`
	// ID of the SSL certificate linked to the CDN resource.
	//
	// Can be used only with `"sslEnabled": true`.
	SslData int64 `json:"sslData" api:"nullable"`
	// Defines whether the HTTPS protocol enabled for content delivery.
	//
	// Possible values:
	//
	// - **true** - HTTPS is enabled.
	// - **false** - HTTPS is disabled.
	SslEnabled bool `json:"sslEnabled"`
	// CDN resource status.
	//
	// Possible values:
	//
	//   - **active** - CDN resource is active. Content is available to users.
	//   - **suspended** - CDN resource is suspended. Content is not available to users.
	//   - **processed** - CDN resource has recently been created and is currently being
	//     processed. It will take about fifteen minutes to propagate it to all
	//     locations.
	//   - **deleted** - CDN resource is deleted.
	//
	// Any of "active", "suspended", "processed", "deleted".
	Status CDNResourceStatus `json:"status"`
	// Date when the CDN resource was suspended automatically if there is no traffic on
	// it for 90 days.
	//
	// Not specified if the resource was not stopped due to lack of traffic.
	SuspendDate string `json:"suspend_date" api:"nullable"`
	// Defines whether the CDN resource has been automatically suspended because there
	// was no traffic on it for 90 days.
	//
	// Possible values:
	//
	// - **true** - CDN resource is currently automatically suspended.
	// - **false** - CDN resource is not automatically suspended.
	//
	// You can enable CDN resource using the `active` field. If there is no traffic on
	// the CDN resource within seven days following activation, it will be suspended
	// again.
	//
	// To avoid CDN resource suspension due to no traffic, contact technical support.
	Suspended bool `json:"suspended"`
	// Date of the last CDN resource update.
	Updated string `json:"updated"`
	// Defines whether the CDN resource is integrated with the Streaming Platform.
	//
	// Possible values:
	//
	//   - **true** - CDN resource is configured for Streaming Platform. Changing
	//     resource settings can affect its operation.
	//   - **false** - CDN resource is not configured for Streaming Platform.
	VpEnabled bool `json:"vp_enabled"`
	// The ID of the associated WAAP domain.
	WaapDomainID string `json:"waap_domain_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Active             respjson.Field
		CanPurgeByURLs     respjson.Field
		Client             respjson.Field
		Cname              respjson.Field
		Created            respjson.Field
		Deleted            respjson.Field
		Description        respjson.Field
		Enabled            respjson.Field
		FullCustomEnabled  respjson.Field
		IsPrimary          respjson.Field
		Name               respjson.Field
		Options            respjson.Field
		OriginGroup        respjson.Field
		OriginGroupName    respjson.Field
		OriginProtocol     respjson.Field
		PresetApplied      respjson.Field
		PrimaryResource    respjson.Field
		ProxySslCa         respjson.Field
		ProxySslData       respjson.Field
		ProxySslEnabled    respjson.Field
		Rules              respjson.Field
		SecondaryHostnames respjson.Field
		ShieldDc           respjson.Field
		ShieldEnabled      respjson.Field
		ShieldRoutingMap   respjson.Field
		Shielded           respjson.Field
		SslData            respjson.Field
		SslEnabled         respjson.Field
		Status             respjson.Field
		SuspendDate        respjson.Field
		Suspended          respjson.Field
		Updated            respjson.Field
		VpEnabled          respjson.Field
		WaapDomainID       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResource) RawJSON() string { return r.JSON.raw }
func (r *CDNResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type CDNResourceOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceOptionsAllowedHTTPMethods `json:"allowedHttpMethods" api:"nullable"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceOptionsBotProtection `json:"bot_protection" api:"nullable"`
	// Compresses content with Brotli on the CDN side. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
	//     activated.
	//  2. If a precache server is not active for a CDN resource, no compression occurs,
	//     even if the option is enabled.
	//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  4. `fetch_compressed` option in CDN resource settings overrides
	//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
	//     resource and want to enable `brotli_compression` in a rule, you must specify
	//     `fetch_compressed:false` in the rule.
	BrotliCompression CDNResourceOptionsBrotliCompression `json:"brotli_compression" api:"nullable"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceOptionsBrowserCacheSettings `json:"browser_cache_settings" api:"nullable"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceOptionsCacheHTTPHeaders `json:"cache_http_headers" api:"nullable"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceOptionsCors `json:"cors" api:"nullable"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceOptionsCountryACL `json:"country_acl" api:"nullable"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceOptionsDisableCache `json:"disable_cache" api:"nullable"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges" api:"nullable"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceOptionsEdgeCacheSettings `json:"edge_cache_settings" api:"nullable"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceOptionsFastedge `json:"fastedge" api:"nullable"`
	// Makes the CDN request compressed content from the origin.
	//
	// The origin server should support compression. CDN servers will not decompress
	// your content even if a user browser does not accept compression.
	//
	// Notes:
	//
	//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
	//     `slice` options enabled.
	//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
	//     you enable it in CDN resource and want to use `gzipON` and
	//     `brotli_compression` in a rule, you have to specify
	//     `"fetch_compressed": false` in the rule.
	FetchCompressed CDNResourceOptionsFetchCompressed `json:"fetch_compressed" api:"nullable"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceOptionsFollowOriginRedirect `json:"follow_origin_redirect" api:"nullable"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceOptionsForceReturn `json:"force_return" api:"nullable"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceOptionsForwardHostHeader `json:"forward_host_header" api:"nullable"`
	// Enables gRPC pass-through for the CDN resource.
	GrpcPassthrough CDNResourceOptionsGrpcPassthrough `json:"grpc_passthrough" api:"nullable"`
	// Compresses content with gzip on the CDN end. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
	//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
	//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
	GzipOn CDNResourceOptionsGzipOn `json:"gzipOn" api:"nullable"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceOptionsHostHeader `json:"hostHeader" api:"nullable"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled CDNResourceOptionsHttp3Enabled `json:"http3_enabled" api:"nullable"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceOptionsIgnoreCookie `json:"ignore_cookie" api:"nullable"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceOptionsIgnoreQueryString `json:"ignoreQueryString" api:"nullable"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceOptionsImageStack `json:"image_stack" api:"nullable"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceOptionsIPAddressACL `json:"ip_address_acl" api:"nullable"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceOptionsLimitBandwidth `json:"limit_bandwidth" api:"nullable"`
	// Allows you to modify your cache key. If omitted, the default value is
	// `$request_uri`.
	//
	// Combine the specified variables to create a key for caching.
	//
	// - **$`request_uri`**
	// - **$scheme**
	// - **$uri**
	//
	// **Warning**: Enabling and changing this option can invalidate your current cache
	// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
	// not work.
	ProxyCacheKey CDNResourceOptionsProxyCacheKey `json:"proxy_cache_key" api:"nullable"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set" api:"nullable"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceOptionsProxyConnectTimeout `json:"proxy_connect_timeout" api:"nullable"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceOptionsProxyReadTimeout `json:"proxy_read_timeout" api:"nullable"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceOptionsQueryParamsBlacklist `json:"query_params_blacklist" api:"nullable"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceOptionsQueryParamsWhitelist `json:"query_params_whitelist" api:"nullable"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceOptionsQueryStringForwarding `json:"query_string_forwarding" api:"nullable"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https" api:"nullable"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http" api:"nullable"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceOptionsReferrerACL `json:"referrer_acl" api:"nullable"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceOptionsRequestLimiter `json:"request_limiter" api:"nullable"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy" api:"nullable"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceOptionsRewrite `json:"rewrite" api:"nullable"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceOptionsSecureKey `json:"secure_key" api:"nullable"`
	// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
	// part.) This reduces time to first byte.
	//
	// The option is based on the
	// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
	//
	// Notes:
	//
	//  1. Origin must support HTTP Range requests.
	//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
	//     options enabled.
	Slice CDNResourceOptionsSlice `json:"slice" api:"nullable"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceOptionsSni `json:"sni" api:"nullable"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceOptionsStale `json:"stale" api:"nullable"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceOptionsStaticResponseHeaders `json:"static_response_headers" api:"nullable"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceOptionsStaticHeaders `json:"staticHeaders" api:"nullable"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceOptionsStaticRequestHeaders `json:"staticRequestHeaders" api:"nullable"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions CDNResourceOptionsTlsVersions `json:"tls_versions" api:"nullable"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain CDNResourceOptionsUseDefaultLeChain `json:"use_default_le_chain" api:"nullable"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge CDNResourceOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge" api:"nullable"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert CDNResourceOptionsUseRsaLeCert `json:"use_rsa_le_cert" api:"nullable"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceOptionsUserAgentACL `json:"user_agent_acl" api:"nullable"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceOptionsWaap `json:"waap" api:"nullable"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceOptionsWebsockets `json:"websockets" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedHTTPMethods          respjson.Field
		BotProtection               respjson.Field
		BrotliCompression           respjson.Field
		BrowserCacheSettings        respjson.Field
		CacheHTTPHeaders            respjson.Field
		Cors                        respjson.Field
		CountryACL                  respjson.Field
		DisableCache                respjson.Field
		DisableProxyForceRanges     respjson.Field
		EdgeCacheSettings           respjson.Field
		Fastedge                    respjson.Field
		FetchCompressed             respjson.Field
		FollowOriginRedirect        respjson.Field
		ForceReturn                 respjson.Field
		ForwardHostHeader           respjson.Field
		GrpcPassthrough             respjson.Field
		GzipOn                      respjson.Field
		HostHeader                  respjson.Field
		Http3Enabled                respjson.Field
		IgnoreCookie                respjson.Field
		IgnoreQueryString           respjson.Field
		ImageStack                  respjson.Field
		IPAddressACL                respjson.Field
		LimitBandwidth              respjson.Field
		ProxyCacheKey               respjson.Field
		ProxyCacheMethodsSet        respjson.Field
		ProxyConnectTimeout         respjson.Field
		ProxyReadTimeout            respjson.Field
		QueryParamsBlacklist        respjson.Field
		QueryParamsWhitelist        respjson.Field
		QueryStringForwarding       respjson.Field
		RedirectHTTPToHTTPS         respjson.Field
		RedirectHTTPSToHTTP         respjson.Field
		ReferrerACL                 respjson.Field
		RequestLimiter              respjson.Field
		ResponseHeadersHidingPolicy respjson.Field
		Rewrite                     respjson.Field
		SecureKey                   respjson.Field
		Slice                       respjson.Field
		Sni                         respjson.Field
		Stale                       respjson.Field
		StaticResponseHeaders       respjson.Field
		StaticHeaders               respjson.Field
		StaticRequestHeaders        respjson.Field
		TlsVersions                 respjson.Field
		UseDefaultLeChain           respjson.Field
		UseDns01LeChallenge         respjson.Field
		UseRsaLeCert                respjson.Field
		UserAgentACL                respjson.Field
		Waap                        respjson.Field
		Websockets                  respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptions) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
type CDNResourceOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsAllowedHTTPMethods) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
type CDNResourceOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceOptionsBotProtectionBotChallenge `json:"bot_challenge" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BotChallenge respjson.Field
		Enabled      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsBotProtection) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsBotProtectionBotChallenge) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with Brotli on the CDN side. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
//     activated.
//  2. If a precache server is not active for a CDN resource, no compression occurs,
//     even if the option is enabled.
//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  4. `fetch_compressed` option in CDN resource settings overrides
//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
//     resource and want to enable `brotli_compression` in a rule, you must specify
//     `fetch_compressed:false` in the rule.
type CDNResourceOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsBrotliCompression) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
type CDNResourceOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value" api:"required" format:"nginx time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsBrowserCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
type CDNResourceOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled" api:"required"`
	Value   []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsCacheHTTPHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
type CDNResourceOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value" api:"required"`
	// Defines whether the Access-Control-Allow-Origin header should be added to a
	// response from CDN regardless of response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response regardless of response code.
	//   - **false** - Header will only be added to responses with codes: 200, 201, 204,
	//     206, 301, 302, 303, 304, 307, 308.
	Always bool `json:"always"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		Always      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsCors) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
type CDNResourceOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values" api:"required" format:"country-code"`
	// Defines the type of CDN resource access policy.
	//
	// Possible values:
	//
	//   - **allow** - Access is allowed for all the countries except for those specified
	//     in `excepted_values` field.
	//   - **deny** - Access is denied for all the countries except for those specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled        respjson.Field
		ExceptedValues respjson.Field
		PolicyType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsCountryACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
type CDNResourceOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsDisableCache) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
type CDNResourceOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsDisableProxyForceRanges) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
type CDNResourceOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP object representing the caching time in seconds for a response with a
	// specific response code.
	//
	// These settings have a higher priority than the `value` field.
	//
	// - Use `any` key to specify caching time for all response codes.
	// - Use `0s` value to disable caching for a specific response code.
	CustomValues map[string]string `json:"custom_values" format:"nginx time"`
	// Enables content caching according to the origin cache settings.
	//
	// The value is applied to the following response codes 200, 201, 204, 206, 301,
	// 302, 303, 304, 307, 308, if an origin server does not have caching HTTP headers.
	//
	// Responses with other codes will not be cached.
	//
	// The maximum duration is any equivalent to `1y`.
	Default string `json:"default" format:"nginx time"`
	// Caching time.
	//
	// The value is applied to the following response codes: 200, 206, 301, 302.
	// Responses with codes 4xx, 5xx will not be cached.
	//
	// Use `0s` to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value" format:"nginx time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled      respjson.Field
		CustomValues respjson.Field
		Default      respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsEdgeCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
type CDNResourceOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceOptionsFastedgeOnRequestBody `json:"on_request_body"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceOptionsFastedgeOnRequestHeaders `json:"on_request_headers"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceOptionsFastedgeOnResponseBody `json:"on_response_body"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceOptionsFastedgeOnResponseHeaders `json:"on_response_headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled           respjson.Field
		OnRequestBody     respjson.Field
		OnRequestHeaders  respjson.Field
		OnResponseBody    respjson.Field
		OnResponseHeaders respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFastedge) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
type CDNResourceOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled bool `json:"enabled"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge bool `json:"execute_on_edge"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield bool `json:"execute_on_shield"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError bool `json:"interrupt_on_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		Enabled          respjson.Field
		ExecuteOnEdge    respjson.Field
		ExecuteOnShield  respjson.Field
		InterruptOnError respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFastedgeOnRequestBody) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
type CDNResourceOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled bool `json:"enabled"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge bool `json:"execute_on_edge"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield bool `json:"execute_on_shield"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError bool `json:"interrupt_on_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		Enabled          respjson.Field
		ExecuteOnEdge    respjson.Field
		ExecuteOnShield  respjson.Field
		InterruptOnError respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFastedgeOnRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
type CDNResourceOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled bool `json:"enabled"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge bool `json:"execute_on_edge"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield bool `json:"execute_on_shield"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError bool `json:"interrupt_on_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		Enabled          respjson.Field
		ExecuteOnEdge    respjson.Field
		ExecuteOnShield  respjson.Field
		InterruptOnError respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFastedgeOnResponseBody) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
type CDNResourceOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled bool `json:"enabled"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge bool `json:"execute_on_edge"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield bool `json:"execute_on_shield"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError bool `json:"interrupt_on_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		Enabled          respjson.Field
		ExecuteOnEdge    respjson.Field
		ExecuteOnShield  respjson.Field
		InterruptOnError respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFastedgeOnResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes the CDN request compressed content from the origin.
//
// The origin server should support compression. CDN servers will not decompress
// your content even if a user browser does not accept compression.
//
// Notes:
//
//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
//     `slice` options enabled.
//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
//     you enable it in CDN resource and want to use `gzipON` and
//     `brotli_compression` in a rule, you have to specify
//     `"fetch_compressed": false` in the rule.
type CDNResourceOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFetchCompressed) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
type CDNResourceOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsFollowOriginRedirect) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
type CDNResourceOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body" api:"required"`
	// Status code value.
	Code int64 `json:"code" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval CDNResourceOptionsForceReturnTimeInterval `json:"time_interval" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body         respjson.Field
		Code         respjson.Field
		Enabled      respjson.Field
		TimeInterval respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsForceReturn) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
type CDNResourceOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time" api:"required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time" api:"required"`
	// Time zone used to calculate time.
	TimeZone string `json:"time_zone" format:"timezone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsForceReturnTimeInterval) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CDNResourceOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsForwardHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables gRPC pass-through for the CDN resource.
type CDNResourceOptionsGrpcPassthrough struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsGrpcPassthrough) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsGrpcPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with gzip on the CDN end. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
type CDNResourceOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsGzipOn) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CDNResourceOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Host Header value.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
type CDNResourceOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsHttp3Enabled) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
type CDNResourceOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsIgnoreCookie) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsIgnoreQueryString) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
type CDNResourceOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables or disables automatic conversion of JPEG and PNG images to AVI format.
	AvifEnabled bool `json:"avif_enabled"`
	// Enables or disables compression without quality loss for PNG format.
	PngLossless bool `json:"png_lossless"`
	// Defines quality settings for JPG and PNG images. The higher the value, the
	// better the image quality, and the larger the file size after conversion.
	Quality int64 `json:"quality"`
	// Enables or disables automatic conversion of JPEG and PNG images to WebP format.
	WebpEnabled bool `json:"webp_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		AvifEnabled respjson.Field
		PngLossless respjson.Field
		Quality     respjson.Field
		WebpEnabled respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsImageStack) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
type CDNResourceOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of IP addresses with a subnet mask.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of IP addresses for which access is prohibited.
	// - **deny** - List of IP addresses for which access is allowed.
	//
	// Examples:
	//
	// - `192.168.3.2/32`
	// - `2a03:d000:2980:7::8/128`
	ExceptedValues []string `json:"excepted_values" api:"required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled        respjson.Field
		ExceptedValues respjson.Field
		PolicyType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsIPAddressACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to control the download speed per connection.
type CDNResourceOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Method of controlling the download speed per connection.
	//
	// Possible values:
	//
	//   - **static** - Use speed and buffer fields to set the download speed limit.
	//   - **dynamic** - Use query strings **speed** and **buffer** to set the download
	//     speed limit.
	//
	// # For example, when requesting content at the link
	//
	// ```
	// http://cdn.example.com/video.mp4?speed=50k&buffer=500k
	// ```
	//
	// the download speed will be limited to 50kB/s after 500 kB.
	//
	// Any of "static", "dynamic".
	LimitType string `json:"limit_type" api:"required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer int64 `json:"buffer"`
	// Maximum download speed per connection.
	Speed int64 `json:"speed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		LimitType   respjson.Field
		Buffer      respjson.Field
		Speed       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsLimitBandwidth) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to modify your cache key. If omitted, the default value is
// `$request_uri`.
//
// Combine the specified variables to create a key for caching.
//
// - **$`request_uri`**
// - **$scheme**
// - **$uri**
//
// **Warning**: Enabling and changing this option can invalidate your current cache
// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
// not work.
type CDNResourceOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Key for caching.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsProxyCacheKey) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
type CDNResourceOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsProxyCacheMethodsSet) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
type CDNResourceOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsProxyConnectTimeout) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
type CDNResourceOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsProxyReadTimeout) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsQueryParamsBlacklist) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsQueryParamsWhitelist) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
type CDNResourceOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types" api:"required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types" api:"required"`
	// The `forward_except_keys` field provides a mechanism to exclude specific
	// parameters from being forwarded from playlist files to media chunk files. By
	// listing certain keys in this field, you can ensure that these parameters are
	// omitted during the forwarding process. This is particularly useful for
	// preventing sensitive or irrelevant information from being included in requests
	// for media chunks, thereby enhancing security and optimizing performance.
	ForwardExceptKeys []string `json:"forward_except_keys"`
	// The `forward_only_keys` field allows for granular control over which specific
	// parameters are forwarded from playlist files to media chunk files. By specifying
	// certain keys, only those parameters will be propagated, ensuring that only
	// relevant information is passed along. This is particularly useful for security
	// and performance optimization, as it prevents unnecessary or sensitive data from
	// being included in requests for media chunks.
	ForwardOnlyKeys []string `json:"forward_only_keys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled              respjson.Field
		ForwardFromFileTypes respjson.Field
		ForwardToFileTypes   respjson.Field
		ForwardExceptKeys    respjson.Field
		ForwardOnlyKeys      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsQueryStringForwarding) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CDNResourceOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsRedirectHTTPToHTTPS) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CDNResourceOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsRedirectHTTPSToHTTP) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
type CDNResourceOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of domain names or wildcard domains (without protocol: `http://` or
	// `https://`.)
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of domain names for which access is prohibited.
	// - **deny** - List of IP domain names for which access is allowed.
	//
	// Examples:
	//
	// - `example.com`
	// - `*.example.com`
	ExceptedValues []string `json:"excepted_values" api:"required" format:"domain or wildcard"`
	// Policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//   - **deny** - Deny access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled        respjson.Field
		ExceptedValues respjson.Field
		PolicyType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsReferrerACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option allows to limit the amount of HTTP requests.
type CDNResourceOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Maximum request rate.
	Rate  int64 `json:"rate" api:"required"`
	Burst int64 `json:"burst"`
	Delay int64 `json:"delay"`
	// Units of measurement for the `rate` field.
	//
	// Possible values:
	//
	// - **r/s** - Requests per second.
	// - **r/m** - Requests per minute.
	//
	// If the rate is less than one request per second, it is specified in request per
	// minute (r/m.)
	//
	// Any of "r/s", "r/m".
	RateUnit string `json:"rate_unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Rate        respjson.Field
		Burst       respjson.Field
		Delay       respjson.Field
		RateUnit    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsRequestLimiter) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hides HTTP headers from an origin server in the CDN response.
type CDNResourceOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of HTTP headers.
	//
	// Parameter meaning depends on the value of the `mode` field:
	//
	//   - **show** - List of HTTP headers to hide from response.
	//   - **hide** - List of HTTP headers to include in response. Other HTTP headers
	//     will be hidden.
	//
	// The following headers are required and cannot be hidden from response:
	//
	// - `Connection`
	// - `Content-Length`
	// - `Content-Type`
	// - `Date`
	// - `Server`
	Excepted []string `json:"excepted" api:"required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Excepted    respjson.Field
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsResponseHeadersHidingPolicy) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
type CDNResourceOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Flag for the Rewrite option.
	//
	// Possible values:
	//
	//   - **last** - Stop processing the current set of `ngx_http_rewrite_module`
	//     directives and start a search for a new location matching changed URI.
	//   - **break** - Stop processing the current set of the Rewrite option.
	//   - **redirect** - Return a temporary redirect with the 302 code; used when a
	//     replacement string does not start with `http://`, `https://`, or `$scheme`.
	//   - **permanent** - Return a permanent redirect with the 301 code.
	//
	// Any of "break", "last", "redirect", "permanent".
	Flag string `json:"flag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Enabled     respjson.Field
		Flag        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsRewrite) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
type CDNResourceOptionsSecureKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Key generated on your side that will be used for URL signing.
	Key string `json:"key" api:"required"`
	// Type of URL signing.
	//
	// Possible types:
	//
	// - **Type 0** - Includes end user IP to secure token generation.
	// - **Type 2** - Excludes end user IP from secure token generation.
	//
	// Any of 0, 2.
	Type int64 `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Key         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsSecureKey) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
// part.) This reduces time to first byte.
//
// The option is based on the
// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
//
// Notes:
//
//  1. Origin must support HTTP Range requests.
//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
//     options enabled.
type CDNResourceOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsSlice) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsSlice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hostname that is added to SNI requests from CDN servers to the origin server
// via HTTPS.
//
// SNI is generally only required if your origin uses shared hosting or does not
// have a dedicated IP address. If the origin server presents multiple
// certificates, SNI allows the origin server to know which certificate to use for
// the connection.
//
// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
type CDNResourceOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname" api:"required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// SNI (Server Name Indication) type.
	//
	// Possible values:
	//
	//   - **dynamic** - SNI hostname depends on `hostHeader` and `forward_host_header`
	//     options. It has several possible combinations:
	//   - If the `hostHeader` option is enabled and specified, SNI hostname matches the
	//     Host header.
	//   - If the `forward_host_header` option is enabled and has true value, SNI
	//     hostname matches the Host header used in the request made to a CDN.
	//   - If the `hostHeader` and `forward_host_header` options are disabled, SNI
	//     hostname matches the primary CNAME.
	//   - **custom** - custom SNI hostname is in use.
	//
	// Any of "dynamic", "custom".
	SniType string `json:"sni_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomHostname respjson.Field
		Enabled        respjson.Field
		SniType        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsSni) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serves stale cached content in case of origin unavailability.
type CDNResourceOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsStale) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
type CDNResourceOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                           `json:"enabled" api:"required"`
	Value   []CDNResourceOptionsStaticResponseHeadersValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsStaticResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourceOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name" api:"required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value" api:"required"`
	// Defines whether the header will be added to a response from CDN regardless of
	// response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response from CDN regardless of response
	//     code.
	//   - **false** - Header will be added only to the following response codes: 200,
	//     201, 204, 206, 301, 302, 303, 304, 307, 308.
	Always bool `json:"always"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		Always      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsStaticResponseHeadersValue) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
type CDNResourceOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsStaticHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
type CDNResourceOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsStaticRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
type CDNResourceOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsTlsVersions) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
type CDNResourceOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsUseDefaultLeChain) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
type CDNResourceOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsUseDns01LeChallenge) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
type CDNResourceOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsUseRsaLeCert) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
type CDNResourceOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values" api:"required" format:"user_agent"`
	// User-Agents policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access for all User-Agents except specified in
	//     `excepted_values` field.
	//   - **deny** - Deny access for all User-Agents except specified in
	//     `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled        respjson.Field
		ExceptedValues respjson.Field
		PolicyType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsUserAgentACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to enable WAAP (Web Application and API Protection).
type CDNResourceOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsWaap) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
type CDNResourceOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceOptionsWebsockets) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol used by CDN servers to request content from an origin source.
//
// Possible values:
//
//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
//   - **HTTP** - CDN servers will connect to the origin via HTTP.
//   - **MATCH** - connection protocol will be chosen automatically (content on the
//     origin source should be available for the CDN both through HTTP and HTTPS).
//
// If protocol is not specified, HTTP is used to connect to an origin server.
type CDNResourceOriginProtocol string

const (
	CDNResourceOriginProtocolHTTP  CDNResourceOriginProtocol = "HTTP"
	CDNResourceOriginProtocolHTTPS CDNResourceOriginProtocol = "HTTPS"
	CDNResourceOriginProtocolMatch CDNResourceOriginProtocol = "MATCH"
)

// CDN resource status.
//
// Possible values:
//
//   - **active** - CDN resource is active. Content is available to users.
//   - **suspended** - CDN resource is suspended. Content is not available to users.
//   - **processed** - CDN resource has recently been created and is currently being
//     processed. It will take about fifteen minutes to propagate it to all
//     locations.
//   - **deleted** - CDN resource is deleted.
type CDNResourceStatus string

const (
	CDNResourceStatusActive    CDNResourceStatus = "active"
	CDNResourceStatusSuspended CDNResourceStatus = "suspended"
	CDNResourceStatusProcessed CDNResourceStatus = "processed"
	CDNResourceStatusDeleted   CDNResourceStatus = "deleted"
)

type CDNResourceList []CDNResource

type CDNResourceNewParams struct {
	// Delivery domains that will be used for content delivery through a CDN.
	//
	// Delivery domains should be added to your DNS settings.
	Cname string `json:"cname" api:"required"`
	// CDN resource name.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the main CDN resource which has a shared caching zone with a reserve CDN
	// resource.
	//
	// If the parameter is not empty, then the current CDN resource is the reserve. You
	// cannot change some options, create rules, set up origin shielding, or use the
	// reserve CDN resource for Streaming.
	PrimaryResource param.Opt[int64] `json:"primary_resource,omitzero"`
	// ID of the trusted CA certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslCa param.Opt[int64] `json:"proxy_ssl_ca,omitzero"`
	// ID of the SSL certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslData param.Opt[int64] `json:"proxy_ssl_data,omitzero"`
	// ID of the SSL certificate linked to the CDN resource.
	//
	// Can be used only with `"sslEnabled": true`.
	SslData param.Opt[int64] `json:"sslData,omitzero"`
	// Enables or disables a CDN resource.
	//
	// Possible values:
	//
	// - **true** - CDN resource is active. Content is being delivered.
	// - **false** - CDN resource is deactivated. Content is not being delivered.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional comment describing the CDN resource.
	Description param.Opt[string] `json:"description,omitzero"`
	// IP address or domain name of the origin and the port, if custom port is used.
	//
	// Exactly one of `origin` or `originGroup` must be provided during resource
	// creation.
	Origin param.Opt[string] `json:"origin,omitzero"`
	// Origin group ID with which the CDN resource is associated.
	//
	// Exactly one of `origin` or `originGroup` must be provided during resource
	// creation.
	OriginGroup param.Opt[int64] `json:"originGroup,omitzero"`
	// Enables or disables SSL certificate validation of the origin server before
	// completing any connection.
	//
	// Possible values:
	//
	// - **true** - Origin SSL certificate validation is enabled.
	// - **false** - Origin SSL certificate validation is disabled.
	ProxySslEnabled param.Opt[bool] `json:"proxy_ssl_enabled,omitzero"`
	// Defines whether the HTTPS protocol enabled for content delivery.
	//
	// Possible values:
	//
	// - **true** - HTTPS is enabled.
	// - **false** - HTTPS is disabled.
	SslEnabled param.Opt[bool] `json:"sslEnabled,omitzero"`
	// Defines whether the associated WAAP Domain is identified as an API Domain.
	//
	// Possible values:
	//
	// - **true** - The associated WAAP Domain is designated as an API Domain.
	// - **false** - The associated WAAP Domain is not designated as an API Domain.
	WaapAPIDomainEnabled param.Opt[bool] `json:"waap_api_domain_enabled,omitzero"`
	// List of options that can be configured for the CDN resource.
	//
	// In case of `null` value the option is not added to the CDN resource. Option may
	// inherit its value from the global account settings.
	Options CDNResourceNewParamsOptions `json:"options,omitzero"`
	// Protocol used by CDN servers to request content from an origin source.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
	//   - **HTTP** - CDN servers will connect to the origin via HTTP.
	//   - **MATCH** - connection protocol will be chosen automatically (content on the
	//     origin source should be available for the CDN both through HTTP and HTTPS).
	//
	// If protocol is not specified, HTTP is used to connect to an origin server.
	//
	// Any of "HTTP", "HTTPS", "MATCH".
	OriginProtocol CDNResourceNewParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r CDNResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type CDNResourceNewParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceNewParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceNewParamsOptionsBotProtection `json:"bot_protection,omitzero"`
	// Compresses content with Brotli on the CDN side. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
	//     activated.
	//  2. If a precache server is not active for a CDN resource, no compression occurs,
	//     even if the option is enabled.
	//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  4. `fetch_compressed` option in CDN resource settings overrides
	//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
	//     resource and want to enable `brotli_compression` in a rule, you must specify
	//     `fetch_compressed:false` in the rule.
	BrotliCompression CDNResourceNewParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceNewParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceNewParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceNewParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceNewParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceNewParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceNewParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceNewParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceNewParamsOptionsFastedge `json:"fastedge,omitzero"`
	// Makes the CDN request compressed content from the origin.
	//
	// The origin server should support compression. CDN servers will not decompress
	// your content even if a user browser does not accept compression.
	//
	// Notes:
	//
	//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
	//     `slice` options enabled.
	//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
	//     you enable it in CDN resource and want to use `gzipON` and
	//     `brotli_compression` in a rule, you have to specify
	//     `"fetch_compressed": false` in the rule.
	FetchCompressed CDNResourceNewParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceNewParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceNewParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceNewParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
	// Enables gRPC pass-through for the CDN resource.
	GrpcPassthrough CDNResourceNewParamsOptionsGrpcPassthrough `json:"grpc_passthrough,omitzero"`
	// Compresses content with gzip on the CDN end. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
	//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
	//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
	GzipOn CDNResourceNewParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceNewParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled CDNResourceNewParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceNewParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceNewParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceNewParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceNewParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceNewParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
	// Allows you to modify your cache key. If omitted, the default value is
	// `$request_uri`.
	//
	// Combine the specified variables to create a key for caching.
	//
	// - **$`request_uri`**
	// - **$scheme**
	// - **$uri**
	//
	// **Warning**: Enabling and changing this option can invalidate your current cache
	// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
	// not work.
	ProxyCacheKey CDNResourceNewParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceNewParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceNewParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceNewParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceNewParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceNewParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceNewParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceNewParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceNewParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceNewParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceNewParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceNewParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceNewParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceNewParamsOptionsSecureKey `json:"secure_key,omitzero"`
	// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
	// part.) This reduces time to first byte.
	//
	// The option is based on the
	// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
	//
	// Notes:
	//
	//  1. Origin must support HTTP Range requests.
	//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
	//     options enabled.
	Slice CDNResourceNewParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceNewParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceNewParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceNewParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceNewParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceNewParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions CDNResourceNewParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain CDNResourceNewParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge CDNResourceNewParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert CDNResourceNewParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceNewParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceNewParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceNewParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceNewParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceNewParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceNewParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with Brotli on the CDN side. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
//     activated.
//  2. If a precache server is not active for a CDN resource, no compression occurs,
//     even if the option is enabled.
//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  4. `fetch_compressed` option in CDN resource settings overrides
//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
//     resource and want to enable `brotli_compression` in a rule, you must specify
//     `fetch_compressed:false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value" api:"required" format:"nginx time"`
	paramObj
}

func (r CDNResourceNewParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled" api:"required"`
	Value   []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the Access-Control-Allow-Origin header should be added to a
	// response from CDN regardless of response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response regardless of response code.
	//   - **false** - Header will only be added to responses with codes: 200, 201, 204,
	//     206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceNewParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"country-code"`
	// Defines the type of CDN resource access policy.
	//
	// Possible values:
	//
	//   - **allow** - Access is allowed for all the countries except for those specified
	//     in `excepted_values` field.
	//   - **deny** - Access is denied for all the countries except for those specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsCountryACL](
		"policy_type", "allow", "deny",
	)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceNewParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables content caching according to the origin cache settings.
	//
	// The value is applied to the following response codes 200, 201, 204, 206, 301,
	// 302, 303, 304, 307, 308, if an origin server does not have caching HTTP headers.
	//
	// Responses with other codes will not be cached.
	//
	// The maximum duration is any equivalent to `1y`.
	Default param.Opt[string] `json:"default,omitzero" format:"nginx time"`
	// Caching time.
	//
	// The value is applied to the following response codes: 200, 206, 301, 302.
	// Responses with codes 4xx, 5xx will not be cached.
	//
	// Use `0s` to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value param.Opt[string] `json:"value,omitzero" format:"nginx time"`
	// A MAP object representing the caching time in seconds for a response with a
	// specific response code.
	//
	// These settings have a higher priority than the `value` field.
	//
	// - Use `any` key to specify caching time for all response codes.
	// - Use `0s` value to disable caching for a specific response code.
	CustomValues map[string]string `json:"custom_values,omitzero" format:"nginx time"`
	paramObj
}

func (r CDNResourceNewParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceNewParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceNewParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceNewParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceNewParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceNewParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceNewParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceNewParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceNewParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceNewParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes the CDN request compressed content from the origin.
//
// The origin server should support compression. CDN servers will not decompress
// your content even if a user browser does not accept compression.
//
// Notes:
//
//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
//     `slice` options enabled.
//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
//     you enable it in CDN resource and want to use `gzipON` and
//     `brotli_compression` in a rule, you have to specify
//     `"fetch_compressed": false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceNewParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceNewParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body" api:"required"`
	// Status code value.
	Code int64 `json:"code" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval CDNResourceNewParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceNewParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time" api:"required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time" api:"required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r CDNResourceNewParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables gRPC pass-through for the CDN resource.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsGrpcPassthrough struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsGrpcPassthrough) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsGrpcPassthrough
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsGrpcPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with gzip on the CDN end. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Host Header value.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceNewParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables or disables automatic conversion of JPEG and PNG images to AVI format.
	AvifEnabled param.Opt[bool] `json:"avif_enabled,omitzero"`
	// Enables or disables compression without quality loss for PNG format.
	PngLossless param.Opt[bool] `json:"png_lossless,omitzero"`
	// Defines quality settings for JPG and PNG images. The higher the value, the
	// better the image quality, and the larger the file size after conversion.
	Quality param.Opt[int64] `json:"quality,omitzero"`
	// Enables or disables automatic conversion of JPEG and PNG images to WebP format.
	WebpEnabled param.Opt[bool] `json:"webp_enabled,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceNewParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of IP addresses with a subnet mask.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of IP addresses for which access is prohibited.
	// - **deny** - List of IP addresses for which access is allowed.
	//
	// Examples:
	//
	// - `192.168.3.2/32`
	// - `2a03:d000:2980:7::8/128`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceNewParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Method of controlling the download speed per connection.
	//
	// Possible values:
	//
	//   - **static** - Use speed and buffer fields to set the download speed limit.
	//   - **dynamic** - Use query strings **speed** and **buffer** to set the download
	//     speed limit.
	//
	// # For example, when requesting content at the link
	//
	// ```
	// http://cdn.example.com/video.mp4?speed=50k&buffer=500k
	// ```
	//
	// the download speed will be limited to 50kB/s after 500 kB.
	//
	// Any of "static", "dynamic".
	LimitType string `json:"limit_type,omitzero" api:"required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsLimitBandwidth](
		"limit_type", "static", "dynamic",
	)
}

// Allows you to modify your cache key. If omitted, the default value is
// `$request_uri`.
//
// Combine the specified variables to create a key for caching.
//
// - **$`request_uri`**
// - **$scheme**
// - **$uri**
//
// **Warning**: Enabling and changing this option can invalidate your current cache
// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
// not work.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Key for caching.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
//
// The properties Enabled, ForwardFromFileTypes, ForwardToFileTypes are required.
type CDNResourceNewParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero" api:"required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero" api:"required"`
	// The `forward_except_keys` field provides a mechanism to exclude specific
	// parameters from being forwarded from playlist files to media chunk files. By
	// listing certain keys in this field, you can ensure that these parameters are
	// omitted during the forwarding process. This is particularly useful for
	// preventing sensitive or irrelevant information from being included in requests
	// for media chunks, thereby enhancing security and optimizing performance.
	ForwardExceptKeys []string `json:"forward_except_keys,omitzero"`
	// The `forward_only_keys` field allows for granular control over which specific
	// parameters are forwarded from playlist files to media chunk files. By specifying
	// certain keys, only those parameters will be propagated, ensuring that only
	// relevant information is passed along. This is particularly useful for security
	// and performance optimization, as it prevents unnecessary or sensitive data from
	// being included in requests for media chunks.
	ForwardOnlyKeys []string `json:"forward_only_keys,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceNewParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of domain names or wildcard domains (without protocol: `http://` or
	// `https://`.)
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of domain names for which access is prohibited.
	// - **deny** - List of IP domain names for which access is allowed.
	//
	// Examples:
	//
	// - `example.com`
	// - `*.example.com`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"domain or wildcard"`
	// Policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//   - **deny** - Deny access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceNewParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Maximum request rate.
	Rate int64 `json:"rate" api:"required"`
	// Units of measurement for the `rate` field.
	//
	// Possible values:
	//
	// - **r/s** - Requests per second.
	// - **r/m** - Requests per minute.
	//
	// If the rate is less than one request per second, it is specified in request per
	// minute (r/m.)
	//
	// Any of "r/s", "r/m".
	RateUnit string `json:"rate_unit,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceNewParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of HTTP headers.
	//
	// Parameter meaning depends on the value of the `mode` field:
	//
	//   - **show** - List of HTTP headers to hide from response.
	//   - **hide** - List of HTTP headers to include in response. Other HTTP headers
	//     will be hidden.
	//
	// The following headers are required and cannot be hidden from response:
	//
	// - `Connection`
	// - `Content-Length`
	// - `Content-Type`
	// - `Date`
	// - `Server`
	Excepted []string `json:"excepted,omitzero" api:"required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceNewParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Flag for the Rewrite option.
	//
	// Possible values:
	//
	//   - **last** - Stop processing the current set of `ngx_http_rewrite_module`
	//     directives and start a search for a new location matching changed URI.
	//   - **break** - Stop processing the current set of the Rewrite option.
	//   - **redirect** - Return a temporary redirect with the 302 code; used when a
	//     replacement string does not start with `http://`, `https://`, or `$scheme`.
	//   - **permanent** - Return a permanent redirect with the 301 code.
	//
	// Any of "break", "last", "redirect", "permanent".
	Flag string `json:"flag,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceNewParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Type of URL signing.
	//
	// Possible types:
	//
	// - **Type 0** - Includes end user IP to secure token generation.
	// - **Type 2** - Excludes end user IP from secure token generation.
	//
	// Any of 0, 2.
	Type int64 `json:"type,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsSecureKey](
		"type", 0, 2,
	)
}

// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
// part.) This reduces time to first byte.
//
// The option is based on the
// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
//
// Notes:
//
//  1. Origin must support HTTP Range requests.
//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
//     options enabled.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsSlice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hostname that is added to SNI requests from CDN servers to the origin server
// via HTTPS.
//
// SNI is generally only required if your origin uses shared hosting or does not
// have a dedicated IP address. If the origin server presents multiple
// certificates, SNI allows the origin server to know which certificate to use for
// the connection.
//
// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
//
// The properties CustomHostname, Enabled are required.
type CDNResourceNewParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname" api:"required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// SNI (Server Name Indication) type.
	//
	// Possible values:
	//
	//   - **dynamic** - SNI hostname depends on `hostHeader` and `forward_host_header`
	//     options. It has several possible combinations:
	//   - If the `hostHeader` option is enabled and specified, SNI hostname matches the
	//     Host header.
	//   - If the `forward_host_header` option is enabled and has true value, SNI
	//     hostname matches the Host header used in the request made to a CDN.
	//   - If the `hostHeader` and `forward_host_header` options are disabled, SNI
	//     hostname matches the primary CNAME.
	//   - **custom** - custom SNI hostname is in use.
	//
	// Any of "dynamic", "custom".
	SniType string `json:"sni_type,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                    `json:"enabled" api:"required"`
	Value   []CDNResourceNewParamsOptionsStaticResponseHeadersValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceNewParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name" api:"required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the header will be added to a response from CDN regardless of
	// response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response from CDN regardless of response
	//     code.
	//   - **false** - Header will be added only to the following response codes: 200,
	//     201, 204, 206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceNewParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceNewParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"user_agent"`
	// User-Agents policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access for all User-Agents except specified in
	//     `excepted_values` field.
	//   - **deny** - Deny access for all User-Agents except specified in
	//     `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceNewParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceNewParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceNewParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceNewParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceNewParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol used by CDN servers to request content from an origin source.
//
// Possible values:
//
//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
//   - **HTTP** - CDN servers will connect to the origin via HTTP.
//   - **MATCH** - connection protocol will be chosen automatically (content on the
//     origin source should be available for the CDN both through HTTP and HTTPS).
//
// If protocol is not specified, HTTP is used to connect to an origin server.
type CDNResourceNewParamsOriginProtocol string

const (
	CDNResourceNewParamsOriginProtocolHTTP  CDNResourceNewParamsOriginProtocol = "HTTP"
	CDNResourceNewParamsOriginProtocolHTTPS CDNResourceNewParamsOriginProtocol = "HTTPS"
	CDNResourceNewParamsOriginProtocolMatch CDNResourceNewParamsOriginProtocol = "MATCH"
)

type CDNResourceUpdateParams struct {
	// CDN resource name.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the trusted CA certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslCa param.Opt[int64] `json:"proxy_ssl_ca,omitzero"`
	// ID of the SSL certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslData param.Opt[int64] `json:"proxy_ssl_data,omitzero"`
	// ID of the SSL certificate linked to the CDN resource.
	//
	// Can be used only with `"sslEnabled": true`.
	SslData param.Opt[int64] `json:"sslData,omitzero"`
	// Enables or disables a CDN resource.
	//
	// Possible values:
	//
	// - **true** - CDN resource is active. Content is being delivered.
	// - **false** - CDN resource is deactivated. Content is not being delivered.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional comment describing the CDN resource.
	Description param.Opt[string] `json:"description,omitzero"`
	// Origin group ID with which the CDN resource is associated.
	OriginGroup param.Opt[int64] `json:"originGroup,omitzero"`
	// Enables or disables SSL certificate validation of the origin server before
	// completing any connection.
	//
	// Possible values:
	//
	// - **true** - Origin SSL certificate validation is enabled.
	// - **false** - Origin SSL certificate validation is disabled.
	ProxySslEnabled param.Opt[bool] `json:"proxy_ssl_enabled,omitzero"`
	// Defines whether the HTTPS protocol enabled for content delivery.
	//
	// Possible values:
	//
	// - **true** - HTTPS is enabled.
	// - **false** - HTTPS is disabled.
	SslEnabled param.Opt[bool] `json:"sslEnabled,omitzero"`
	// List of options that can be configured for the CDN resource.
	//
	// In case of `null` value the option is not added to the CDN resource. Option may
	// inherit its value from the global account settings.
	Options CDNResourceUpdateParamsOptions `json:"options,omitzero"`
	// Protocol used by CDN servers to request content from an origin source.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
	//   - **HTTP** - CDN servers will connect to the origin via HTTP.
	//   - **MATCH** - connection protocol will be chosen automatically (content on the
	//     origin source should be available for the CDN both through HTTP and HTTPS).
	//
	// If protocol is not specified, HTTP is used to connect to an origin server.
	//
	// Any of "HTTP", "HTTPS", "MATCH".
	OriginProtocol CDNResourceUpdateParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r CDNResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type CDNResourceUpdateParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceUpdateParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceUpdateParamsOptionsBotProtection `json:"bot_protection,omitzero"`
	// Compresses content with Brotli on the CDN side. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
	//     activated.
	//  2. If a precache server is not active for a CDN resource, no compression occurs,
	//     even if the option is enabled.
	//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  4. `fetch_compressed` option in CDN resource settings overrides
	//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
	//     resource and want to enable `brotli_compression` in a rule, you must specify
	//     `fetch_compressed:false` in the rule.
	BrotliCompression CDNResourceUpdateParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceUpdateParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceUpdateParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceUpdateParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceUpdateParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceUpdateParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceUpdateParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceUpdateParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceUpdateParamsOptionsFastedge `json:"fastedge,omitzero"`
	// Makes the CDN request compressed content from the origin.
	//
	// The origin server should support compression. CDN servers will not decompress
	// your content even if a user browser does not accept compression.
	//
	// Notes:
	//
	//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
	//     `slice` options enabled.
	//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
	//     you enable it in CDN resource and want to use `gzipON` and
	//     `brotli_compression` in a rule, you have to specify
	//     `"fetch_compressed": false` in the rule.
	FetchCompressed CDNResourceUpdateParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceUpdateParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceUpdateParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceUpdateParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
	// Enables gRPC pass-through for the CDN resource.
	GrpcPassthrough CDNResourceUpdateParamsOptionsGrpcPassthrough `json:"grpc_passthrough,omitzero"`
	// Compresses content with gzip on the CDN end. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
	//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
	//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
	GzipOn CDNResourceUpdateParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceUpdateParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled CDNResourceUpdateParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceUpdateParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceUpdateParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceUpdateParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceUpdateParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceUpdateParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
	// Allows you to modify your cache key. If omitted, the default value is
	// `$request_uri`.
	//
	// Combine the specified variables to create a key for caching.
	//
	// - **$`request_uri`**
	// - **$scheme**
	// - **$uri**
	//
	// **Warning**: Enabling and changing this option can invalidate your current cache
	// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
	// not work.
	ProxyCacheKey CDNResourceUpdateParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceUpdateParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceUpdateParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceUpdateParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceUpdateParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceUpdateParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceUpdateParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceUpdateParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceUpdateParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceUpdateParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceUpdateParamsOptionsSecureKey `json:"secure_key,omitzero"`
	// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
	// part.) This reduces time to first byte.
	//
	// The option is based on the
	// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
	//
	// Notes:
	//
	//  1. Origin must support HTTP Range requests.
	//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
	//     options enabled.
	Slice CDNResourceUpdateParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceUpdateParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceUpdateParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceUpdateParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceUpdateParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceUpdateParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions CDNResourceUpdateParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain CDNResourceUpdateParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge CDNResourceUpdateParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert CDNResourceUpdateParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceUpdateParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceUpdateParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceUpdateParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceUpdateParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceUpdateParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceUpdateParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with Brotli on the CDN side. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
//     activated.
//  2. If a precache server is not active for a CDN resource, no compression occurs,
//     even if the option is enabled.
//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  4. `fetch_compressed` option in CDN resource settings overrides
//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
//     resource and want to enable `brotli_compression` in a rule, you must specify
//     `fetch_compressed:false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value" api:"required" format:"nginx time"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled" api:"required"`
	Value   []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the Access-Control-Allow-Origin header should be added to a
	// response from CDN regardless of response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response regardless of response code.
	//   - **false** - Header will only be added to responses with codes: 200, 201, 204,
	//     206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceUpdateParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"country-code"`
	// Defines the type of CDN resource access policy.
	//
	// Possible values:
	//
	//   - **allow** - Access is allowed for all the countries except for those specified
	//     in `excepted_values` field.
	//   - **deny** - Access is denied for all the countries except for those specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsCountryACL](
		"policy_type", "allow", "deny",
	)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceUpdateParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables content caching according to the origin cache settings.
	//
	// The value is applied to the following response codes 200, 201, 204, 206, 301,
	// 302, 303, 304, 307, 308, if an origin server does not have caching HTTP headers.
	//
	// Responses with other codes will not be cached.
	//
	// The maximum duration is any equivalent to `1y`.
	Default param.Opt[string] `json:"default,omitzero" format:"nginx time"`
	// Caching time.
	//
	// The value is applied to the following response codes: 200, 206, 301, 302.
	// Responses with codes 4xx, 5xx will not be cached.
	//
	// Use `0s` to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value param.Opt[string] `json:"value,omitzero" format:"nginx time"`
	// A MAP object representing the caching time in seconds for a response with a
	// specific response code.
	//
	// These settings have a higher priority than the `value` field.
	//
	// - Use `any` key to specify caching time for all response codes.
	// - Use `0s` value to disable caching for a specific response code.
	CustomValues map[string]string `json:"custom_values,omitzero" format:"nginx time"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceUpdateParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceUpdateParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceUpdateParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceUpdateParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceUpdateParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes the CDN request compressed content from the origin.
//
// The origin server should support compression. CDN servers will not decompress
// your content even if a user browser does not accept compression.
//
// Notes:
//
//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
//     `slice` options enabled.
//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
//     you enable it in CDN resource and want to use `gzipON` and
//     `brotli_compression` in a rule, you have to specify
//     `"fetch_compressed": false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceUpdateParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceUpdateParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body" api:"required"`
	// Status code value.
	Code int64 `json:"code" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval CDNResourceUpdateParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceUpdateParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time" api:"required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time" api:"required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables gRPC pass-through for the CDN resource.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsGrpcPassthrough struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsGrpcPassthrough) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsGrpcPassthrough
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsGrpcPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with gzip on the CDN end. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Host Header value.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceUpdateParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables or disables automatic conversion of JPEG and PNG images to AVI format.
	AvifEnabled param.Opt[bool] `json:"avif_enabled,omitzero"`
	// Enables or disables compression without quality loss for PNG format.
	PngLossless param.Opt[bool] `json:"png_lossless,omitzero"`
	// Defines quality settings for JPG and PNG images. The higher the value, the
	// better the image quality, and the larger the file size after conversion.
	Quality param.Opt[int64] `json:"quality,omitzero"`
	// Enables or disables automatic conversion of JPEG and PNG images to WebP format.
	WebpEnabled param.Opt[bool] `json:"webp_enabled,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceUpdateParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of IP addresses with a subnet mask.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of IP addresses for which access is prohibited.
	// - **deny** - List of IP addresses for which access is allowed.
	//
	// Examples:
	//
	// - `192.168.3.2/32`
	// - `2a03:d000:2980:7::8/128`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceUpdateParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Method of controlling the download speed per connection.
	//
	// Possible values:
	//
	//   - **static** - Use speed and buffer fields to set the download speed limit.
	//   - **dynamic** - Use query strings **speed** and **buffer** to set the download
	//     speed limit.
	//
	// # For example, when requesting content at the link
	//
	// ```
	// http://cdn.example.com/video.mp4?speed=50k&buffer=500k
	// ```
	//
	// the download speed will be limited to 50kB/s after 500 kB.
	//
	// Any of "static", "dynamic".
	LimitType string `json:"limit_type,omitzero" api:"required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsLimitBandwidth](
		"limit_type", "static", "dynamic",
	)
}

// Allows you to modify your cache key. If omitted, the default value is
// `$request_uri`.
//
// Combine the specified variables to create a key for caching.
//
// - **$`request_uri`**
// - **$scheme**
// - **$uri**
//
// **Warning**: Enabling and changing this option can invalidate your current cache
// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
// not work.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Key for caching.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
//
// The properties Enabled, ForwardFromFileTypes, ForwardToFileTypes are required.
type CDNResourceUpdateParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero" api:"required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero" api:"required"`
	// The `forward_except_keys` field provides a mechanism to exclude specific
	// parameters from being forwarded from playlist files to media chunk files. By
	// listing certain keys in this field, you can ensure that these parameters are
	// omitted during the forwarding process. This is particularly useful for
	// preventing sensitive or irrelevant information from being included in requests
	// for media chunks, thereby enhancing security and optimizing performance.
	ForwardExceptKeys []string `json:"forward_except_keys,omitzero"`
	// The `forward_only_keys` field allows for granular control over which specific
	// parameters are forwarded from playlist files to media chunk files. By specifying
	// certain keys, only those parameters will be propagated, ensuring that only
	// relevant information is passed along. This is particularly useful for security
	// and performance optimization, as it prevents unnecessary or sensitive data from
	// being included in requests for media chunks.
	ForwardOnlyKeys []string `json:"forward_only_keys,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceUpdateParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of domain names or wildcard domains (without protocol: `http://` or
	// `https://`.)
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of domain names for which access is prohibited.
	// - **deny** - List of IP domain names for which access is allowed.
	//
	// Examples:
	//
	// - `example.com`
	// - `*.example.com`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"domain or wildcard"`
	// Policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//   - **deny** - Deny access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceUpdateParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Maximum request rate.
	Rate int64 `json:"rate" api:"required"`
	// Units of measurement for the `rate` field.
	//
	// Possible values:
	//
	// - **r/s** - Requests per second.
	// - **r/m** - Requests per minute.
	//
	// If the rate is less than one request per second, it is specified in request per
	// minute (r/m.)
	//
	// Any of "r/s", "r/m".
	RateUnit string `json:"rate_unit,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of HTTP headers.
	//
	// Parameter meaning depends on the value of the `mode` field:
	//
	//   - **show** - List of HTTP headers to hide from response.
	//   - **hide** - List of HTTP headers to include in response. Other HTTP headers
	//     will be hidden.
	//
	// The following headers are required and cannot be hidden from response:
	//
	// - `Connection`
	// - `Content-Length`
	// - `Content-Type`
	// - `Date`
	// - `Server`
	Excepted []string `json:"excepted,omitzero" api:"required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceUpdateParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Flag for the Rewrite option.
	//
	// Possible values:
	//
	//   - **last** - Stop processing the current set of `ngx_http_rewrite_module`
	//     directives and start a search for a new location matching changed URI.
	//   - **break** - Stop processing the current set of the Rewrite option.
	//   - **redirect** - Return a temporary redirect with the 302 code; used when a
	//     replacement string does not start with `http://`, `https://`, or `$scheme`.
	//   - **permanent** - Return a permanent redirect with the 301 code.
	//
	// Any of "break", "last", "redirect", "permanent".
	Flag string `json:"flag,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceUpdateParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Type of URL signing.
	//
	// Possible types:
	//
	// - **Type 0** - Includes end user IP to secure token generation.
	// - **Type 2** - Excludes end user IP from secure token generation.
	//
	// Any of 0, 2.
	Type int64 `json:"type,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsSecureKey](
		"type", 0, 2,
	)
}

// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
// part.) This reduces time to first byte.
//
// The option is based on the
// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
//
// Notes:
//
//  1. Origin must support HTTP Range requests.
//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
//     options enabled.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsSlice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hostname that is added to SNI requests from CDN servers to the origin server
// via HTTPS.
//
// SNI is generally only required if your origin uses shared hosting or does not
// have a dedicated IP address. If the origin server presents multiple
// certificates, SNI allows the origin server to know which certificate to use for
// the connection.
//
// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
//
// The properties CustomHostname, Enabled are required.
type CDNResourceUpdateParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname" api:"required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// SNI (Server Name Indication) type.
	//
	// Possible values:
	//
	//   - **dynamic** - SNI hostname depends on `hostHeader` and `forward_host_header`
	//     options. It has several possible combinations:
	//   - If the `hostHeader` option is enabled and specified, SNI hostname matches the
	//     Host header.
	//   - If the `forward_host_header` option is enabled and has true value, SNI
	//     hostname matches the Host header used in the request made to a CDN.
	//   - If the `hostHeader` and `forward_host_header` options are disabled, SNI
	//     hostname matches the primary CNAME.
	//   - **custom** - custom SNI hostname is in use.
	//
	// Any of "dynamic", "custom".
	SniType string `json:"sni_type,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                       `json:"enabled" api:"required"`
	Value   []CDNResourceUpdateParamsOptionsStaticResponseHeadersValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceUpdateParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name" api:"required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the header will be added to a response from CDN regardless of
	// response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response from CDN regardless of response
	//     code.
	//   - **false** - Header will be added only to the following response codes: 200,
	//     201, 204, 206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceUpdateParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"user_agent"`
	// User-Agents policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access for all User-Agents except specified in
	//     `excepted_values` field.
	//   - **deny** - Deny access for all User-Agents except specified in
	//     `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceUpdateParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceUpdateParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceUpdateParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceUpdateParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceUpdateParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol used by CDN servers to request content from an origin source.
//
// Possible values:
//
//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
//   - **HTTP** - CDN servers will connect to the origin via HTTP.
//   - **MATCH** - connection protocol will be chosen automatically (content on the
//     origin source should be available for the CDN both through HTTP and HTTPS).
//
// If protocol is not specified, HTTP is used to connect to an origin server.
type CDNResourceUpdateParamsOriginProtocol string

const (
	CDNResourceUpdateParamsOriginProtocolHTTP  CDNResourceUpdateParamsOriginProtocol = "HTTP"
	CDNResourceUpdateParamsOriginProtocolHTTPS CDNResourceUpdateParamsOriginProtocol = "HTTPS"
	CDNResourceUpdateParamsOriginProtocolMatch CDNResourceUpdateParamsOriginProtocol = "MATCH"
)

type CDNResourceListParams struct {
	// Delivery domain (CNAME) of the CDN resource.
	Cname param.Opt[string] `query:"cname,omitzero" json:"-"`
	// Defines whether a CDN resource has been deleted.
	//
	// Possible values:
	//
	// - **true** - CDN resource has been deleted.
	// - **false** - CDN resource has not been deleted.
	Deleted param.Opt[bool] `query:"deleted,omitzero" json:"-"`
	// Enables or disables a CDN resource change by a user.
	//
	// Possible values:
	//
	// - **true** - CDN resource is enabled.
	// - **false** - CDN resource is disabled.
	Enabled param.Opt[bool] `query:"enabled,omitzero" json:"-"`
	// Most recent date of CDN resource creation for which CDN resources should be
	// returned (ISO 8601/RFC 3339 format, UTC.)
	MaxCreated param.Opt[string] `query:"max_created,omitzero" json:"-"`
	// Earliest date of CDN resource creation for which CDN resources should be
	// returned (ISO 8601/RFC 3339 format, UTC.)
	MinCreated param.Opt[string] `query:"min_created,omitzero" json:"-"`
	// Origin group ID.
	OriginGroup param.Opt[int64] `query:"originGroup,omitzero" json:"-"`
	// Rule name or pattern.
	Rules param.Opt[string] `query:"rules,omitzero" json:"-"`
	// Additional delivery domains (CNAMEs) of the CDN resource.
	SecondaryHostnames param.Opt[string] `query:"secondaryHostnames,omitzero" json:"-"`
	// Name of the origin shielding data center location.
	ShieldDc param.Opt[string] `query:"shield_dc,omitzero" json:"-"`
	// Defines whether origin shielding is enabled for the CDN resource.
	//
	// Possible values:
	//
	// - **true** - Origin shielding is enabled for the CDN resource.
	// - **false** - Origin shielding is disabled for the CDN resource.
	Shielded param.Opt[bool] `query:"shielded,omitzero" json:"-"`
	// SSL certificate ID.
	SslData param.Opt[int64] `query:"sslData,omitzero" json:"-"`
	// SSL certificates IDs.
	//
	// Example:
	//
	// - ?`sslData_in`=1643,1644,1652
	SslDataIn param.Opt[int64] `query:"sslData_in,omitzero" json:"-"`
	// Defines whether the HTTPS protocol is enabled for content delivery.
	//
	// Possible values:
	//
	// - **true** - HTTPS protocol is enabled for CDN resource.
	// - **false** - HTTPS protocol is disabled for CDN resource.
	SslEnabled param.Opt[bool] `query:"sslEnabled,omitzero" json:"-"`
	// Defines whether the CDN resource was automatically suspended by the system.
	//
	// Possible values:
	//
	//   - **true** - CDN resource is selected for automatic suspension in the next 7
	//     days.
	//   - **false** - CDN resource is not selected for automatic suspension.
	Suspend param.Opt[bool] `query:"suspend,omitzero" json:"-"`
	// Defines whether the CDN resource is integrated with the Streaming platform.
	//
	// Possible values:
	//
	// - **true** - CDN resource is used for Streaming platform.
	// - **false** - CDN resource is not used for Streaming platform.
	VpEnabled param.Opt[bool] `query:"vp_enabled,omitzero" json:"-"`
	// CDN resource status.
	//
	// Any of "active", "processed", "suspended", "deleted".
	Status CDNResourceListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CDNResourceListParams]'s query parameters as `url.Values`.
func (r CDNResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// CDN resource status.
type CDNResourceListParamsStatus string

const (
	CDNResourceListParamsStatusActive    CDNResourceListParamsStatus = "active"
	CDNResourceListParamsStatusProcessed CDNResourceListParamsStatus = "processed"
	CDNResourceListParamsStatusSuspended CDNResourceListParamsStatus = "suspended"
	CDNResourceListParamsStatusDeleted   CDNResourceListParamsStatus = "deleted"
)

type CDNResourcePrefetchParams struct {
	// Paths to files that should be pre-populated to the CDN.
	//
	// Paths to the files should be specified without a domain name.
	Paths []string `json:"paths,omitzero" api:"required"`
	paramObj
}

func (r CDNResourcePrefetchParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourcePrefetchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourcePrefetchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourcePurgeParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfPurgeByURL *CDNResourcePurgeParamsBodyPurgeByURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPurgeByPattern *CDNResourcePurgeParamsBodyPurgeByPattern `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPurgeAllCache *CDNResourcePurgeParamsBodyPurgeAllCache `json:",inline"`

	paramObj
}

func (u CDNResourcePurgeParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPurgeByURL, u.OfPurgeByPattern, u.OfPurgeAllCache)
}
func (r *CDNResourcePurgeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourcePurgeParamsBodyPurgeByURL struct {
	// **Purge by URL** clears the cache of a specific files. This purge type is
	// recommended.
	//
	// Specify file URLs including query strings. URLs should start with / without a
	// domain name.
	//
	// Purge by URL depends on the following CDN options:
	//
	//  1. "vary response header" is used. If your origin serves variants of the same
	//     content depending on the Vary HTTP response header, purge by URL will delete
	//     only one version of the file.
	//  2. "slice" is used. If you update several files in the origin without clearing
	//     the CDN cache, purge by URL will delete only the first slice (with bytes=0…
	//     .)
	//  3. "ignoreQueryString" is used. Don’t specify parameters in the purge request.
	//  4. "query_params_blacklist" is used. Only files with the listed in the option
	//     parameters will be cached as different objects. Files with other parameters
	//     will be cached as one object. In this case, specify the listed parameters in
	//     the Purge request. Don't specify other parameters.
	//  5. "query_params_whitelist" is used. Files with listed in the option parameters
	//     will be cached as one object. Files with other parameters will be cached as
	//     different objects. In this case, specify other parameters (if any) besides
	//     the ones listed in the purge request.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r CDNResourcePurgeParamsBodyPurgeByURL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourcePurgeParamsBodyPurgeByURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourcePurgeParamsBodyPurgeByURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourcePurgeParamsBodyPurgeByPattern struct {
	// **Purge by pattern** clears the cache that matches the pattern.
	//
	// Use _ operator, which replaces any number of symbols in your path. It's
	// important to note that wildcard usage (_) is permitted only at the end of a
	// pattern.
	//
	// Query string added to any patterns will be ignored, and purge request will be
	// processed as if there weren't any parameters.
	//
	// Purge by pattern is recursive. Both /path and /path* will result in recursive
	// purging, meaning all content under the specified path will be affected. As such,
	// using the pattern /path* is functionally equivalent to simply using /path.
	Paths []string `json:"paths,omitzero"`
	paramObj
}

func (r CDNResourcePurgeParamsBodyPurgeByPattern) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourcePurgeParamsBodyPurgeByPattern
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourcePurgeParamsBodyPurgeByPattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourcePurgeParamsBodyPurgeAllCache struct {
	// **Purge all cache** clears the entire cache for the CDN resource.
	//
	// Specify an empty array to purge all content for the resource.
	//
	// When you purge all assets, CDN servers request content from your origin server
	// and cause a high load. Therefore, we recommend to use purge by URL for large
	// content quantities.
	Paths []string `json:"paths,omitzero"`
	paramObj
}

func (r CDNResourcePurgeParamsBodyPurgeAllCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourcePurgeParamsBodyPurgeAllCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourcePurgeParamsBodyPurgeAllCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourceReplaceParams struct {
	// Origin group ID with which the CDN resource is associated.
	OriginGroup int64 `json:"originGroup" api:"required"`
	// CDN resource name.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the trusted CA certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslCa param.Opt[int64] `json:"proxy_ssl_ca,omitzero"`
	// ID of the SSL certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslData param.Opt[int64] `json:"proxy_ssl_data,omitzero"`
	// ID of the SSL certificate linked to the CDN resource.
	//
	// Can be used only with `"sslEnabled": true`.
	SslData param.Opt[int64] `json:"sslData,omitzero"`
	// Enables or disables a CDN resource.
	//
	// Possible values:
	//
	// - **true** - CDN resource is active. Content is being delivered.
	// - **false** - CDN resource is deactivated. Content is not being delivered.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional comment describing the CDN resource.
	Description param.Opt[string] `json:"description,omitzero"`
	// Enables or disables SSL certificate validation of the origin server before
	// completing any connection.
	//
	// Possible values:
	//
	// - **true** - Origin SSL certificate validation is enabled.
	// - **false** - Origin SSL certificate validation is disabled.
	ProxySslEnabled param.Opt[bool] `json:"proxy_ssl_enabled,omitzero"`
	// Defines whether the HTTPS protocol enabled for content delivery.
	//
	// Possible values:
	//
	// - **true** - HTTPS is enabled.
	// - **false** - HTTPS is disabled.
	SslEnabled param.Opt[bool] `json:"sslEnabled,omitzero"`
	// Defines whether the associated WAAP Domain is identified as an API Domain.
	//
	// Possible values:
	//
	// - **true** - The associated WAAP Domain is designated as an API Domain.
	// - **false** - The associated WAAP Domain is not designated as an API Domain.
	WaapAPIDomainEnabled param.Opt[bool] `json:"waap_api_domain_enabled,omitzero"`
	// List of options that can be configured for the CDN resource.
	//
	// In case of `null` value the option is not added to the CDN resource. Option may
	// inherit its value from the global account settings.
	Options CDNResourceReplaceParamsOptions `json:"options,omitzero"`
	// Protocol used by CDN servers to request content from an origin source.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
	//   - **HTTP** - CDN servers will connect to the origin via HTTP.
	//   - **MATCH** - connection protocol will be chosen automatically (content on the
	//     origin source should be available for the CDN both through HTTP and HTTPS).
	//
	// If protocol is not specified, HTTP is used to connect to an origin server.
	//
	// Any of "HTTP", "HTTPS", "MATCH".
	OriginProtocol CDNResourceReplaceParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r CDNResourceReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type CDNResourceReplaceParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceReplaceParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceReplaceParamsOptionsBotProtection `json:"bot_protection,omitzero"`
	// Compresses content with Brotli on the CDN side. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
	//     activated.
	//  2. If a precache server is not active for a CDN resource, no compression occurs,
	//     even if the option is enabled.
	//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  4. `fetch_compressed` option in CDN resource settings overrides
	//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
	//     resource and want to enable `brotli_compression` in a rule, you must specify
	//     `fetch_compressed:false` in the rule.
	BrotliCompression CDNResourceReplaceParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceReplaceParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceReplaceParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceReplaceParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceReplaceParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceReplaceParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceReplaceParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceReplaceParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceReplaceParamsOptionsFastedge `json:"fastedge,omitzero"`
	// Makes the CDN request compressed content from the origin.
	//
	// The origin server should support compression. CDN servers will not decompress
	// your content even if a user browser does not accept compression.
	//
	// Notes:
	//
	//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
	//     `slice` options enabled.
	//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
	//     you enable it in CDN resource and want to use `gzipON` and
	//     `brotli_compression` in a rule, you have to specify
	//     `"fetch_compressed": false` in the rule.
	FetchCompressed CDNResourceReplaceParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceReplaceParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceReplaceParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceReplaceParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
	// Enables gRPC pass-through for the CDN resource.
	GrpcPassthrough CDNResourceReplaceParamsOptionsGrpcPassthrough `json:"grpc_passthrough,omitzero"`
	// Compresses content with gzip on the CDN end. CDN servers will request only
	// uncompressed content from the origin.
	//
	// Notes:
	//
	//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
	//     options enabled.
	//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
	//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
	//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
	GzipOn CDNResourceReplaceParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceReplaceParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled CDNResourceReplaceParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceReplaceParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceReplaceParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceReplaceParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceReplaceParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceReplaceParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
	// Allows you to modify your cache key. If omitted, the default value is
	// `$request_uri`.
	//
	// Combine the specified variables to create a key for caching.
	//
	// - **$`request_uri`**
	// - **$scheme**
	// - **$uri**
	//
	// **Warning**: Enabling and changing this option can invalidate your current cache
	// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
	// not work.
	ProxyCacheKey CDNResourceReplaceParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceReplaceParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceReplaceParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceReplaceParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceReplaceParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceReplaceParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceReplaceParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceReplaceParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceReplaceParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceReplaceParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceReplaceParamsOptionsSecureKey `json:"secure_key,omitzero"`
	// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
	// part.) This reduces time to first byte.
	//
	// The option is based on the
	// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
	//
	// Notes:
	//
	//  1. Origin must support HTTP Range requests.
	//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
	//     options enabled.
	Slice CDNResourceReplaceParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceReplaceParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceReplaceParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceReplaceParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceReplaceParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceReplaceParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions CDNResourceReplaceParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain CDNResourceReplaceParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge CDNResourceReplaceParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert CDNResourceReplaceParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceReplaceParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceReplaceParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceReplaceParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceReplaceParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceReplaceParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceReplaceParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with Brotli on the CDN side. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. CDN only supports "Brotli compression" when the "origin shielding" feature is
//     activated.
//  2. If a precache server is not active for a CDN resource, no compression occurs,
//     even if the option is enabled.
//  3. `brotli_compression` is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  4. `fetch_compressed` option in CDN resource settings overrides
//     `brotli_compression` in rules. If you enabled `fetch_compressed` in CDN
//     resource and want to enable `brotli_compression` in a rule, you must specify
//     `fetch_compressed:false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value" api:"required" format:"nginx time"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled" api:"required"`
	Value   []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the Access-Control-Allow-Origin header should be added to a
	// response from CDN regardless of response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response regardless of response code.
	//   - **false** - Header will only be added to responses with codes: 200, 201, 204,
	//     206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceReplaceParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"country-code"`
	// Defines the type of CDN resource access policy.
	//
	// Possible values:
	//
	//   - **allow** - Access is allowed for all the countries except for those specified
	//     in `excepted_values` field.
	//   - **deny** - Access is denied for all the countries except for those specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsCountryACL](
		"policy_type", "allow", "deny",
	)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceReplaceParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables content caching according to the origin cache settings.
	//
	// The value is applied to the following response codes 200, 201, 204, 206, 301,
	// 302, 303, 304, 307, 308, if an origin server does not have caching HTTP headers.
	//
	// Responses with other codes will not be cached.
	//
	// The maximum duration is any equivalent to `1y`.
	Default param.Opt[string] `json:"default,omitzero" format:"nginx time"`
	// Caching time.
	//
	// The value is applied to the following response codes: 200, 206, 301, 302.
	// Responses with codes 4xx, 5xx will not be cached.
	//
	// Use `0s` to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value param.Opt[string] `json:"value,omitzero" format:"nginx time"`
	// A MAP object representing the caching time in seconds for a response with a
	// specific response code.
	//
	// These settings have a higher priority than the `value` field.
	//
	// - Use `any` key to specify caching time for all response codes.
	// - Use `0s` value to disable caching for a specific response code.
	CustomValues map[string]string `json:"custom_values,omitzero" format:"nginx time"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceReplaceParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceReplaceParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceReplaceParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceReplaceParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceReplaceParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id" api:"required"`
	// Determines if the FastEdge application should be called whenever HTTP request
	// headers are received.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Determines if the request should be executed at the edge nodes.
	ExecuteOnEdge param.Opt[bool] `json:"execute_on_edge,omitzero"`
	// Determines if the request should be executed at the shield nodes.
	ExecuteOnShield param.Opt[bool] `json:"execute_on_shield,omitzero"`
	// Determines if the request execution should be interrupted when an error occurs.
	InterruptOnError param.Opt[bool] `json:"interrupt_on_error,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes the CDN request compressed content from the origin.
//
// The origin server should support compression. CDN servers will not decompress
// your content even if a user browser does not accept compression.
//
// Notes:
//
//  1. `fetch_compressed` is not supported with `gzipON` or `brotli_compression` or
//     `slice` options enabled.
//  2. `fetch_compressed` overrides `gzipON` and `brotli_compression` in rule. If
//     you enable it in CDN resource and want to use `gzipON` and
//     `brotli_compression` in a rule, you have to specify
//     `"fetch_compressed": false` in the rule.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceReplaceParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceReplaceParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body" api:"required"`
	// Status code value.
	Code int64 `json:"code" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval CDNResourceReplaceParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceReplaceParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time" api:"required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time" api:"required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables gRPC pass-through for the CDN resource.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsGrpcPassthrough struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsGrpcPassthrough) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsGrpcPassthrough
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsGrpcPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compresses content with gzip on the CDN end. CDN servers will request only
// uncompressed content from the origin.
//
// Notes:
//
//  1. Compression with gzip is not supported with `fetch_compressed` or `slice`
//     options enabled.
//  2. `fetch_compressed` option in CDN resource settings overrides `gzipON` in
//     rules. If you enable `fetch_compressed` in CDN resource and want to enable
//     `gzipON` in rules, you need to specify `"fetch_compressed":false` for rules.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Host Header value.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceReplaceParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Enables or disables automatic conversion of JPEG and PNG images to AVI format.
	AvifEnabled param.Opt[bool] `json:"avif_enabled,omitzero"`
	// Enables or disables compression without quality loss for PNG format.
	PngLossless param.Opt[bool] `json:"png_lossless,omitzero"`
	// Defines quality settings for JPG and PNG images. The higher the value, the
	// better the image quality, and the larger the file size after conversion.
	Quality param.Opt[int64] `json:"quality,omitzero"`
	// Enables or disables automatic conversion of JPEG and PNG images to WebP format.
	WebpEnabled param.Opt[bool] `json:"webp_enabled,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceReplaceParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of IP addresses with a subnet mask.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of IP addresses for which access is prohibited.
	// - **deny** - List of IP addresses for which access is allowed.
	//
	// Examples:
	//
	// - `192.168.3.2/32`
	// - `2a03:d000:2980:7::8/128`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceReplaceParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Method of controlling the download speed per connection.
	//
	// Possible values:
	//
	//   - **static** - Use speed and buffer fields to set the download speed limit.
	//   - **dynamic** - Use query strings **speed** and **buffer** to set the download
	//     speed limit.
	//
	// # For example, when requesting content at the link
	//
	// ```
	// http://cdn.example.com/video.mp4?speed=50k&buffer=500k
	// ```
	//
	// the download speed will be limited to 50kB/s after 500 kB.
	//
	// Any of "static", "dynamic".
	LimitType string `json:"limit_type,omitzero" api:"required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsLimitBandwidth](
		"limit_type", "static", "dynamic",
	)
}

// Allows you to modify your cache key. If omitted, the default value is
// `$request_uri`.
//
// Combine the specified variables to create a key for caching.
//
// - **$`request_uri`**
// - **$scheme**
// - **$uri**
//
// **Warning**: Enabling and changing this option can invalidate your current cache
// and affect the cache hit ratio. Furthermore, the "Purge by pattern" option will
// not work.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Key for caching.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of query parameters.
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
//
// The properties Enabled, ForwardFromFileTypes, ForwardToFileTypes are required.
type CDNResourceReplaceParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero" api:"required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero" api:"required"`
	// The `forward_except_keys` field provides a mechanism to exclude specific
	// parameters from being forwarded from playlist files to media chunk files. By
	// listing certain keys in this field, you can ensure that these parameters are
	// omitted during the forwarding process. This is particularly useful for
	// preventing sensitive or irrelevant information from being included in requests
	// for media chunks, thereby enhancing security and optimizing performance.
	ForwardExceptKeys []string `json:"forward_except_keys,omitzero"`
	// The `forward_only_keys` field allows for granular control over which specific
	// parameters are forwarded from playlist files to media chunk files. By specifying
	// certain keys, only those parameters will be propagated, ensuring that only
	// relevant information is passed along. This is particularly useful for security
	// and performance optimization, as it prevents unnecessary or sensitive data from
	// being included in requests for media chunks.
	ForwardOnlyKeys []string `json:"forward_only_keys,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceReplaceParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of domain names or wildcard domains (without protocol: `http://` or
	// `https://`.)
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of domain names for which access is prohibited.
	// - **deny** - List of IP domain names for which access is allowed.
	//
	// Examples:
	//
	// - `example.com`
	// - `*.example.com`
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"domain or wildcard"`
	// Policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//   - **deny** - Deny access to all domain names except the domain names specified
	//     in `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceReplaceParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Maximum request rate.
	Rate int64 `json:"rate" api:"required"`
	// Units of measurement for the `rate` field.
	//
	// Possible values:
	//
	// - **r/s** - Requests per second.
	// - **r/m** - Requests per minute.
	//
	// If the rate is less than one request per second, it is specified in request per
	// minute (r/m.)
	//
	// Any of "r/s", "r/m".
	RateUnit string `json:"rate_unit,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of HTTP headers.
	//
	// Parameter meaning depends on the value of the `mode` field:
	//
	//   - **show** - List of HTTP headers to hide from response.
	//   - **hide** - List of HTTP headers to include in response. Other HTTP headers
	//     will be hidden.
	//
	// The following headers are required and cannot be hidden from response:
	//
	// - `Connection`
	// - `Content-Length`
	// - `Content-Type`
	// - `Date`
	// - `Server`
	Excepted []string `json:"excepted,omitzero" api:"required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceReplaceParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Flag for the Rewrite option.
	//
	// Possible values:
	//
	//   - **last** - Stop processing the current set of `ngx_http_rewrite_module`
	//     directives and start a search for a new location matching changed URI.
	//   - **break** - Stop processing the current set of the Rewrite option.
	//   - **redirect** - Return a temporary redirect with the 302 code; used when a
	//     replacement string does not start with `http://`, `https://`, or `$scheme`.
	//   - **permanent** - Return a permanent redirect with the 301 code.
	//
	// Any of "break", "last", "redirect", "permanent".
	Flag string `json:"flag,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceReplaceParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero" api:"required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Type of URL signing.
	//
	// Possible types:
	//
	// - **Type 0** - Includes end user IP to secure token generation.
	// - **Type 2** - Excludes end user IP from secure token generation.
	//
	// Any of 0, 2.
	Type int64 `json:"type,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsSecureKey](
		"type", 0, 2,
	)
}

// Requests and caches files larger than 10 MB in parts (no larger than 10 MB per
// part.) This reduces time to first byte.
//
// The option is based on the
// [Slice](https://nginx.org/en/docs/http/ngx_http_slice_module.html) module.
//
// Notes:
//
//  1. Origin must support HTTP Range requests.
//  2. Not supported with `gzipON`, `brotli_compression` or `fetch_compressed`
//     options enabled.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsSlice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hostname that is added to SNI requests from CDN servers to the origin server
// via HTTPS.
//
// SNI is generally only required if your origin uses shared hosting or does not
// have a dedicated IP address. If the origin server presents multiple
// certificates, SNI allows the origin server to know which certificate to use for
// the connection.
//
// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
//
// The properties CustomHostname, Enabled are required.
type CDNResourceReplaceParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname" api:"required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// SNI (Server Name Indication) type.
	//
	// Possible values:
	//
	//   - **dynamic** - SNI hostname depends on `hostHeader` and `forward_host_header`
	//     options. It has several possible combinations:
	//   - If the `hostHeader` option is enabled and specified, SNI hostname matches the
	//     Host header.
	//   - If the `forward_host_header` option is enabled and has true value, SNI
	//     hostname matches the Host header used in the request made to a CDN.
	//   - If the `hostHeader` and `forward_host_header` options are disabled, SNI
	//     hostname matches the primary CNAME.
	//   - **custom** - custom SNI hostname is in use.
	//
	// Any of "dynamic", "custom".
	SniType string `json:"sni_type,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                        `json:"enabled" api:"required"`
	Value   []CDNResourceReplaceParamsOptionsStaticResponseHeadersValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceReplaceParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name" api:"required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero" api:"required"`
	// Defines whether the header will be added to a response from CDN regardless of
	// response code.
	//
	// Possible values:
	//
	//   - **true** - Header will be added to a response from CDN regardless of response
	//     code.
	//   - **false** - Header will be added only to the following response codes: 200,
	//     201, 204, 206, 301, 302, 303, 304, 307, 308.
	Always param.Opt[bool] `json:"always,omitzero"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceReplaceParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero" api:"required" format:"user_agent"`
	// User-Agents policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access for all User-Agents except specified in
	//     `excepted_values` field.
	//   - **deny** - Deny access for all User-Agents except specified in
	//     `excepted_values` field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceReplaceParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceReplaceParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled" api:"required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value" api:"required"`
	paramObj
}

func (r CDNResourceReplaceParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceReplaceParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceReplaceParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol used by CDN servers to request content from an origin source.
//
// Possible values:
//
//   - **HTTPS** - CDN servers will connect to the origin via HTTPS.
//   - **HTTP** - CDN servers will connect to the origin via HTTP.
//   - **MATCH** - connection protocol will be chosen automatically (content on the
//     origin source should be available for the CDN both through HTTP and HTTPS).
//
// If protocol is not specified, HTTP is used to connect to an origin server.
type CDNResourceReplaceParamsOriginProtocol string

const (
	CDNResourceReplaceParamsOriginProtocolHTTP  CDNResourceReplaceParamsOriginProtocol = "HTTP"
	CDNResourceReplaceParamsOriginProtocolHTTPS CDNResourceReplaceParamsOriginProtocol = "HTTPS"
	CDNResourceReplaceParamsOriginProtocolMatch CDNResourceReplaceParamsOriginProtocol = "MATCH"
)
