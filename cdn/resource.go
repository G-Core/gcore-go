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

// ResourceService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceService] method instead.
type ResourceService struct {
	Options []option.RequestOption
	Shield  ResourceShieldService
	Rules   ResourceRuleService
}

// NewResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceService(opts ...option.RequestOption) (r ResourceService) {
	r = ResourceService{}
	r.Options = opts
	r.Shield = NewResourceShieldService(opts...)
	r.Rules = NewResourceRuleService(opts...)
	return
}

// Create CDN resource
func (r *ResourceService) New(ctx context.Context, body ResourceNewParams, opts ...option.RequestOption) (res *CdnResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change CDN resource
func (r *ResourceService) Update(ctx context.Context, resourceID int64, body ResourceUpdateParams, opts ...option.RequestOption) (res *CdnResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get information about all CDN resources in your account.
func (r *ResourceService) List(ctx context.Context, query ResourceListParams, opts ...option.RequestOption) (res *CdnResourceList, err error) {
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
func (r *ResourceService) Delete(ctx context.Context, resourceID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get CDN resource details
func (r *ResourceService) Get(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *CdnResource, err error) {
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
func (r *ResourceService) Prefetch(ctx context.Context, resourceID int64, body ResourcePrefetchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/prefetch", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Check whether a Let's Encrypt certificate can be issued for the CDN resource.
func (r *ResourceService) PrevalidateSslLeCertificate(ctx context.Context, resourceID int64, opts ...option.RequestOption) (err error) {
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
func (r *ResourceService) Purge(ctx context.Context, resourceID int64, body ResourcePurgeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/purge", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Change CDN resource
func (r *ResourceService) Replace(ctx context.Context, resourceID int64, body ResourceReplaceParams, opts ...option.RequestOption) (res *CdnResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CdnResource struct {
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
	IsPrimary bool `json:"is_primary,nullable"`
	// CDN resource name.
	Name string `json:"name,nullable"`
	// List of options that can be configured for the CDN resource.
	//
	// In case of `null` value the option is not added to the CDN resource. Option may
	// inherit its value from the global account settings.
	Options CdnResourceOptions `json:"options"`
	// Origin group ID with which the CDN resource is associated.
	//
	// You can use either the `origin` or `originGroup` parameter in the request.
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
	OriginProtocol CdnResourceOriginProtocol `json:"originProtocol"`
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
	PrimaryResource int64 `json:"primary_resource,nullable"`
	// ID of the trusted CA certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslCa int64 `json:"proxy_ssl_ca,nullable"`
	// ID of the SSL certificate used to verify an origin.
	//
	// It can be used only with `"proxy_ssl_enabled": true`.
	ProxySslData int64 `json:"proxy_ssl_data,nullable"`
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
	ShieldDc string `json:"shield_dc,nullable"`
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
	ShieldRoutingMap int64 `json:"shield_routing_map,nullable"`
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
	SslData int64 `json:"sslData,nullable"`
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
	Status CdnResourceStatus `json:"status"`
	// Date when the CDN resource was suspended automatically if there is no traffic on
	// it for 90 days.
	//
	// Not specified if the resource was not stopped due to lack of traffic.
	SuspendDate string `json:"suspend_date,nullable"`
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
	WaapDomainID string `json:"waap_domain_id,nullable"`
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
func (r CdnResource) RawJSON() string { return r.JSON.raw }
func (r *CdnResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type CdnResourceOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CdnResourceOptionsAllowedHTTPMethods `json:"allowedHttpMethods,nullable"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CdnResourceOptionsBotProtection `json:"bot_protection,nullable"`
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
	BrotliCompression CdnResourceOptionsBrotliCompression `json:"brotli_compression,nullable"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CdnResourceOptionsBrowserCacheSettings `json:"browser_cache_settings,nullable"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CdnResourceOptionsCacheHTTPHeaders `json:"cache_http_headers,nullable"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CdnResourceOptionsCors `json:"cors,nullable"`
	// Enables control access to content for specified countries.
	CountryACL CdnResourceOptionsCountryACL `json:"country_acl,nullable"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CdnResourceOptionsDisableCache `json:"disable_cache,nullable"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CdnResourceOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,nullable"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CdnResourceOptionsEdgeCacheSettings `json:"edge_cache_settings,nullable"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CdnResourceOptionsFastedge `json:"fastedge,nullable"`
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
	FetchCompressed CdnResourceOptionsFetchCompressed `json:"fetch_compressed,nullable"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CdnResourceOptionsFollowOriginRedirect `json:"follow_origin_redirect,nullable"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CdnResourceOptionsForceReturn `json:"force_return,nullable"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CdnResourceOptionsForwardHostHeader `json:"forward_host_header,nullable"`
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
	GzipOn CdnResourceOptionsGzipOn `json:"gzipOn,nullable"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CdnResourceOptionsHostHeader `json:"hostHeader,nullable"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled CdnResourceOptionsHttp3Enabled `json:"http3_enabled,nullable"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CdnResourceOptionsIgnoreCookie `json:"ignore_cookie,nullable"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CdnResourceOptionsIgnoreQueryString `json:"ignoreQueryString,nullable"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CdnResourceOptionsImageStack `json:"image_stack,nullable"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance. We recommend you use a script
	// for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CdnResourceOptionsIPAddressACL `json:"ip_address_acl,nullable"`
	// Allows to control the download speed per connection.
	LimitBandwidth CdnResourceOptionsLimitBandwidth `json:"limit_bandwidth,nullable"`
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
	ProxyCacheKey CdnResourceOptionsProxyCacheKey `json:"proxy_cache_key,nullable"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CdnResourceOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,nullable"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CdnResourceOptionsProxyConnectTimeout `json:"proxy_connect_timeout,nullable"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CdnResourceOptionsProxyReadTimeout `json:"proxy_read_timeout,nullable"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CdnResourceOptionsQueryParamsBlacklist `json:"query_params_blacklist,nullable"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CdnResourceOptionsQueryParamsWhitelist `json:"query_params_whitelist,nullable"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CdnResourceOptionsQueryStringForwarding `json:"query_string_forwarding,nullable"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CdnResourceOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,nullable"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CdnResourceOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,nullable"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CdnResourceOptionsReferrerACL `json:"referrer_acl,nullable"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CdnResourceOptionsRequestLimiter `json:"request_limiter,nullable"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CdnResourceOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,nullable"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CdnResourceOptionsRewrite `json:"rewrite,nullable"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CdnResourceOptionsSecureKey `json:"secure_key,nullable"`
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
	Slice CdnResourceOptionsSlice `json:"slice,nullable"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CdnResourceOptionsSni `json:"sni,nullable"`
	// Serves stale cached content in case of origin unavailability.
	Stale CdnResourceOptionsStale `json:"stale,nullable"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CdnResourceOptionsStaticResponseHeaders `json:"static_response_headers,nullable"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CdnResourceOptionsStaticHeaders `json:"staticHeaders,nullable"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CdnResourceOptionsStaticRequestHeaders `json:"staticRequestHeaders,nullable"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions CdnResourceOptionsTlsVersions `json:"tls_versions,nullable"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain CdnResourceOptionsUseDefaultLeChain `json:"use_default_le_chain,nullable"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge CdnResourceOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,nullable"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert CdnResourceOptionsUseRsaLeCert `json:"use_rsa_le_cert,nullable"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CdnResourceOptionsUserAgentACL `json:"user_agent_acl,nullable"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CdnResourceOptionsWaap `json:"waap,nullable"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CdnResourceOptionsWebsockets `json:"websockets,nullable"`
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
func (r CdnResourceOptions) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
type CdnResourceOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsAllowedHTTPMethods) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
type CdnResourceOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CdnResourceOptionsBotProtectionBotChallenge `json:"bot_challenge,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BotChallenge respjson.Field
		Enabled      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsBotProtection) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CdnResourceOptionsBotProtectionBotChallenge struct {
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
func (r CdnResourceOptionsBotProtectionBotChallenge) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsBrotliCompression) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
type CdnResourceOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value,required" format:"nginx time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsBrowserCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
type CdnResourceOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled,required"`
	Value   []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsCacheHTTPHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
type CdnResourceOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,required"`
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
func (r CdnResourceOptionsCors) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
type CdnResourceOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,required" format:"country-code"`
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
	PolicyType string `json:"policy_type,required"`
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
func (r CdnResourceOptionsCountryACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
type CdnResourceOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsDisableCache) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
type CdnResourceOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsDisableProxyForceRanges) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
type CdnResourceOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
func (r CdnResourceOptionsEdgeCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
type CdnResourceOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CdnResourceOptionsFastedgeOnRequestBody `json:"on_request_body"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders CdnResourceOptionsFastedgeOnRequestHeaders `json:"on_request_headers"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CdnResourceOptionsFastedgeOnResponseBody `json:"on_response_body"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CdnResourceOptionsFastedgeOnResponseHeaders `json:"on_response_headers"`
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
func (r CdnResourceOptionsFastedge) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
type CdnResourceOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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
func (r CdnResourceOptionsFastedgeOnRequestBody) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
type CdnResourceOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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
func (r CdnResourceOptionsFastedgeOnRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
type CdnResourceOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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
func (r CdnResourceOptionsFastedgeOnResponseBody) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
type CdnResourceOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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
func (r CdnResourceOptionsFastedgeOnResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsFetchCompressed) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
type CdnResourceOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codes       respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsFollowOriginRedirect) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
type CdnResourceOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body,required"`
	// Status code value.
	Code int64 `json:"code,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval CdnResourceOptionsForceReturnTimeInterval `json:"time_interval,nullable"`
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
func (r CdnResourceOptionsForceReturn) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
type CdnResourceOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time,required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time,required"`
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
func (r CdnResourceOptionsForceReturnTimeInterval) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CdnResourceOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsForwardHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsGzipOn) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CdnResourceOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Host Header value.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
type CdnResourceOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsHttp3Enabled) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
type CdnResourceOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsIgnoreCookie) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsIgnoreQueryString) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
type CdnResourceOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
func (r CdnResourceOptionsImageStack) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance. We recommend you use a script
// for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
type CdnResourceOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in
	//     "`excepted_values`" field.
	//   - **deny** - Deny access to all IPs except IPs specified in "`excepted_values`"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,required"`
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
func (r CdnResourceOptionsIPAddressACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to control the download speed per connection.
type CdnResourceOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	LimitType string `json:"limit_type,required"`
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
func (r CdnResourceOptionsLimitBandwidth) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Key for caching.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsProxyCacheKey) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
type CdnResourceOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsProxyCacheMethodsSet) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
type CdnResourceOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsProxyConnectTimeout) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
type CdnResourceOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsProxyReadTimeout) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsQueryParamsBlacklist) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsQueryParamsWhitelist) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
type CdnResourceOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,required"`
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
func (r CdnResourceOptionsQueryStringForwarding) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CdnResourceOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsRedirectHTTPToHTTPS) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CdnResourceOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsRedirectHTTPSToHTTP) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
type CdnResourceOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,required" format:"domain or wildcard"`
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
	PolicyType string `json:"policy_type,required"`
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
func (r CdnResourceOptionsReferrerACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option allows to limit the amount of HTTP requests.
type CdnResourceOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Maximum request rate.
	Rate  int64 `json:"rate,required"`
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
func (r CdnResourceOptionsRequestLimiter) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hides HTTP headers from an origin server in the CDN response.
type CdnResourceOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	Excepted []string `json:"excepted,required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,required"`
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
func (r CdnResourceOptionsResponseHeadersHidingPolicy) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
type CdnResourceOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
func (r CdnResourceOptionsRewrite) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
type CdnResourceOptionsSecureKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Key generated on your side that will be used for URL signing.
	Key string `json:"key,required"`
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
func (r CdnResourceOptionsSecureKey) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsSecureKey) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsSlice) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CdnResourceOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname,required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
func (r CdnResourceOptionsSni) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serves stale cached content in case of origin unavailability.
type CdnResourceOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsStale) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
type CdnResourceOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                           `json:"enabled,required"`
	Value   []CdnResourceOptionsStaticResponseHeadersValue `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsStaticResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnResourceOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name,required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,required"`
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
func (r CdnResourceOptionsStaticResponseHeadersValue) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
type CdnResourceOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsStaticHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
type CdnResourceOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsStaticRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
type CdnResourceOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsTlsVersions) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
type CdnResourceOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsUseDefaultLeChain) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
type CdnResourceOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsUseDns01LeChallenge) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
type CdnResourceOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsUseRsaLeCert) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
type CdnResourceOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,required" format:"user_agent"`
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
	PolicyType string `json:"policy_type,required"`
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
func (r CdnResourceOptionsUserAgentACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to enable WAAP (Web Application and API Protection).
type CdnResourceOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsWaap) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
type CdnResourceOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceOptionsWebsockets) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type CdnResourceOriginProtocol string

const (
	CdnResourceOriginProtocolHTTP  CdnResourceOriginProtocol = "HTTP"
	CdnResourceOriginProtocolHTTPS CdnResourceOriginProtocol = "HTTPS"
	CdnResourceOriginProtocolMatch CdnResourceOriginProtocol = "MATCH"
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
type CdnResourceStatus string

const (
	CdnResourceStatusActive    CdnResourceStatus = "active"
	CdnResourceStatusSuspended CdnResourceStatus = "suspended"
	CdnResourceStatusProcessed CdnResourceStatus = "processed"
	CdnResourceStatusDeleted   CdnResourceStatus = "deleted"
)

type CdnResourceList []CdnResource

type ResourceNewParams struct {
	// Delivery domains that will be used for content delivery through a CDN.
	//
	// Delivery domains should be added to your DNS settings.
	Cname string `json:"cname,required"`
	// IP address or domain name of the origin and the port, if custom port is used.
	//
	// You can use either the `origin` or `originGroup` parameter in the request.
	Origin string `json:"origin,required"`
	// Origin group ID with which the CDN resource is associated.
	//
	// You can use either the `origin` or `originGroup` parameter in the request.
	OriginGroup int64 `json:"originGroup,required"`
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
	Options ResourceNewParamsOptions `json:"options,omitzero"`
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
	OriginProtocol ResourceNewParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r ResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type ResourceNewParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceNewParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceNewParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceNewParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceNewParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceNewParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceNewParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceNewParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceNewParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceNewParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceNewParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceNewParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceNewParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceNewParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceNewParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceNewParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceNewParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceNewParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled ResourceNewParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceNewParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceNewParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceNewParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance. We recommend you use a script
	// for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceNewParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceNewParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceNewParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceNewParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceNewParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceNewParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceNewParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceNewParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceNewParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceNewParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceNewParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceNewParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceNewParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceNewParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceNewParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceNewParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceNewParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceNewParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceNewParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceNewParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceNewParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceNewParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions ResourceNewParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain ResourceNewParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge ResourceNewParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert ResourceNewParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceNewParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceNewParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceNewParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceNewParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceNewParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceNewParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceNewParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceNewParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceNewParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value,required" format:"nginx time"`
	paramObj
}

func (r ResourceNewParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled,required"`
	Value   []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceNewParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceNewParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"country-code"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsCountryACL](
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
type ResourceNewParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceNewParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceNewParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceNewParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceNewParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceNewParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceNewParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceNewParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceNewParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceNewParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceNewParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceNewParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceNewParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceNewParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceNewParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceNewParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceNewParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceNewParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceNewParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceNewParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body,required"`
	// Status code value.
	Code int64 `json:"code,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval ResourceNewParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceNewParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceNewParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time,required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time,required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r ResourceNewParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Host Header value.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceNewParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceNewParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance. We recommend you use a script
// for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceNewParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in
	//     "`excepted_values`" field.
	//   - **deny** - Deny access to all IPs except IPs specified in "`excepted_values`"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceNewParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	LimitType string `json:"limit_type,omitzero,required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r ResourceNewParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsLimitBandwidth](
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
type ResourceNewParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Key for caching.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero,required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero,required"`
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

func (r ResourceNewParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceNewParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"domain or wildcard"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceNewParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Maximum request rate.
	Rate int64 `json:"rate,required"`
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

func (r ResourceNewParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceNewParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	Excepted []string `json:"excepted,omitzero,required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceNewParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceNewParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceNewParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceNewParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsSecureKey](
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
type ResourceNewParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname,required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceNewParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                 `json:"enabled,required"`
	Value   []ResourceNewParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceNewParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name,required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceNewParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceNewParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"user_agent"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceNewParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceNewParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceNewParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceNewParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceNewParamsOriginProtocol string

const (
	ResourceNewParamsOriginProtocolHTTP  ResourceNewParamsOriginProtocol = "HTTP"
	ResourceNewParamsOriginProtocolHTTPS ResourceNewParamsOriginProtocol = "HTTPS"
	ResourceNewParamsOriginProtocolMatch ResourceNewParamsOriginProtocol = "MATCH"
)

type ResourceUpdateParams struct {
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
	//
	// You can use either the `origin` or `originGroup` parameter in the request.
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
	Options ResourceUpdateParamsOptions `json:"options,omitzero"`
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
	OriginProtocol ResourceUpdateParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r ResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type ResourceUpdateParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceUpdateParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceUpdateParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceUpdateParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceUpdateParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceUpdateParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceUpdateParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceUpdateParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceUpdateParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceUpdateParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceUpdateParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceUpdateParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceUpdateParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceUpdateParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceUpdateParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceUpdateParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceUpdateParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceUpdateParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled ResourceUpdateParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceUpdateParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceUpdateParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceUpdateParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance. We recommend you use a script
	// for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceUpdateParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceUpdateParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceUpdateParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceUpdateParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceUpdateParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceUpdateParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceUpdateParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceUpdateParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceUpdateParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceUpdateParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceUpdateParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceUpdateParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceUpdateParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceUpdateParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceUpdateParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceUpdateParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceUpdateParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceUpdateParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceUpdateParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceUpdateParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceUpdateParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceUpdateParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions ResourceUpdateParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain ResourceUpdateParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge ResourceUpdateParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert ResourceUpdateParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceUpdateParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceUpdateParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceUpdateParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceUpdateParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceUpdateParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceUpdateParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceUpdateParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value,required" format:"nginx time"`
	paramObj
}

func (r ResourceUpdateParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled,required"`
	Value   []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceUpdateParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceUpdateParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"country-code"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsCountryACL](
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
type ResourceUpdateParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceUpdateParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceUpdateParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceUpdateParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceUpdateParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceUpdateParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceUpdateParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceUpdateParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceUpdateParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceUpdateParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceUpdateParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceUpdateParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceUpdateParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceUpdateParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceUpdateParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceUpdateParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceUpdateParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceUpdateParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceUpdateParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body,required"`
	// Status code value.
	Code int64 `json:"code,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval ResourceUpdateParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceUpdateParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceUpdateParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time,required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time,required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r ResourceUpdateParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Host Header value.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceUpdateParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceUpdateParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance. We recommend you use a script
// for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceUpdateParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in
	//     "`excepted_values`" field.
	//   - **deny** - Deny access to all IPs except IPs specified in "`excepted_values`"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceUpdateParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	LimitType string `json:"limit_type,omitzero,required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r ResourceUpdateParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsLimitBandwidth](
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
type ResourceUpdateParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Key for caching.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero,required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero,required"`
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

func (r ResourceUpdateParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceUpdateParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"domain or wildcard"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceUpdateParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Maximum request rate.
	Rate int64 `json:"rate,required"`
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

func (r ResourceUpdateParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceUpdateParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	Excepted []string `json:"excepted,omitzero,required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceUpdateParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceUpdateParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceUpdateParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceUpdateParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsSecureKey](
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
type ResourceUpdateParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname,required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceUpdateParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                    `json:"enabled,required"`
	Value   []ResourceUpdateParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceUpdateParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name,required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceUpdateParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceUpdateParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"user_agent"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceUpdateParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceUpdateParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceUpdateParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceUpdateParamsOriginProtocol string

const (
	ResourceUpdateParamsOriginProtocolHTTP  ResourceUpdateParamsOriginProtocol = "HTTP"
	ResourceUpdateParamsOriginProtocolHTTPS ResourceUpdateParamsOriginProtocol = "HTTPS"
	ResourceUpdateParamsOriginProtocolMatch ResourceUpdateParamsOriginProtocol = "MATCH"
)

type ResourceListParams struct {
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
	Status ResourceListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResourceListParams]'s query parameters as `url.Values`.
func (r ResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// CDN resource status.
type ResourceListParamsStatus string

const (
	ResourceListParamsStatusActive    ResourceListParamsStatus = "active"
	ResourceListParamsStatusProcessed ResourceListParamsStatus = "processed"
	ResourceListParamsStatusSuspended ResourceListParamsStatus = "suspended"
	ResourceListParamsStatusDeleted   ResourceListParamsStatus = "deleted"
)

type ResourcePrefetchParams struct {
	// Paths to files that should be pre-populated to the CDN.
	//
	// Paths to the files should be specified without a domain name.
	Paths []string `json:"paths,omitzero,required"`
	paramObj
}

func (r ResourcePrefetchParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourcePrefetchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourcePrefetchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourcePurgeParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfPurgeByURL *ResourcePurgeParamsBodyPurgeByURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPurgeByPattern *ResourcePurgeParamsBodyPurgeByPattern `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPurgeAllCache *ResourcePurgeParamsBodyPurgeAllCache `json:",inline"`

	paramObj
}

func (u ResourcePurgeParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPurgeByURL, u.OfPurgeByPattern, u.OfPurgeAllCache)
}
func (r *ResourcePurgeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourcePurgeParamsBodyPurgeByURL struct {
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
	//  4. "`query_params_blacklist`" is used. Only files with the listed in the option
	//     parameters will be cached as different objects. Files with other parameters
	//     will be cached as one object. In this case, specify the listed parameters in
	//     the Purge request. Don't specify other parameters.
	//  5. "`query_params_whitelist`" is used. Files with listed in the option
	//     parameters will be cached as one object. Files with other parameters will be
	//     cached as different objects. In this case, specify other parameters (if any)
	//     besides the ones listed in the purge request.
	URLs []string `json:"urls,omitzero"`
	paramObj
}

func (r ResourcePurgeParamsBodyPurgeByURL) MarshalJSON() (data []byte, err error) {
	type shadow ResourcePurgeParamsBodyPurgeByURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourcePurgeParamsBodyPurgeByURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourcePurgeParamsBodyPurgeByPattern struct {
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

func (r ResourcePurgeParamsBodyPurgeByPattern) MarshalJSON() (data []byte, err error) {
	type shadow ResourcePurgeParamsBodyPurgeByPattern
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourcePurgeParamsBodyPurgeByPattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourcePurgeParamsBodyPurgeAllCache struct {
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

func (r ResourcePurgeParamsBodyPurgeAllCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourcePurgeParamsBodyPurgeAllCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourcePurgeParamsBodyPurgeAllCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceReplaceParams struct {
	// Origin group ID with which the CDN resource is associated.
	//
	// You can use either the `origin` or `originGroup` parameter in the request.
	OriginGroup int64 `json:"originGroup,required"`
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
	Options ResourceReplaceParamsOptions `json:"options,omitzero"`
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
	OriginProtocol ResourceReplaceParamsOriginProtocol `json:"originProtocol,omitzero"`
	// Additional delivery domains (CNAMEs) that will be used to deliver content via
	// the CDN.
	//
	// Up to ten additional CNAMEs are possible.
	SecondaryHostnames []string `json:"secondaryHostnames,omitzero" format:"domain"`
	paramObj
}

func (r ResourceReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the CDN resource.
//
// In case of `null` value the option is not added to the CDN resource. Option may
// inherit its value from the global account settings.
type ResourceReplaceParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceReplaceParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceReplaceParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceReplaceParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceReplaceParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceReplaceParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceReplaceParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceReplaceParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceReplaceParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceReplaceParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceReplaceParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceReplaceParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceReplaceParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceReplaceParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceReplaceParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceReplaceParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceReplaceParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceReplaceParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Enables HTTP/3 protocol for content delivery.
	//
	// `http3_enabled` option works only with `"sslEnabled": true`.
	Http3Enabled ResourceReplaceParamsOptionsHttp3Enabled `json:"http3_enabled,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceReplaceParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceReplaceParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceReplaceParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance. We recommend you use a script
	// for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceReplaceParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceReplaceParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceReplaceParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceReplaceParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceReplaceParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceReplaceParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceReplaceParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceReplaceParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceReplaceParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceReplaceParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceReplaceParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceReplaceParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceReplaceParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceReplaceParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceReplaceParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceReplaceParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceReplaceParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceReplaceParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceReplaceParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceReplaceParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceReplaceParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceReplaceParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
	// to the domain.
	//
	// When the option is disabled, all protocols versions are allowed.
	TlsVersions ResourceReplaceParamsOptionsTlsVersions `json:"tls_versions,omitzero"`
	// Let's Encrypt certificate chain.
	//
	// The specified chain will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseDefaultLeChain ResourceReplaceParamsOptionsUseDefaultLeChain `json:"use_default_le_chain,omitzero"`
	// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
	//
	// DNS service should be activated to enable this option.
	UseDns01LeChallenge ResourceReplaceParamsOptionsUseDns01LeChallenge `json:"use_dns01_le_challenge,omitzero"`
	// RSA Let's Encrypt certificate type for the CDN resource.
	//
	// The specified value will be used during the next Let's Encrypt certificate issue
	// or renewal.
	UseRsaLeCert ResourceReplaceParamsOptionsUseRsaLeCert `json:"use_rsa_le_cert,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceReplaceParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceReplaceParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceReplaceParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceReplaceParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsAllowedHTTPMethods struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceReplaceParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceReplaceParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceReplaceParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceReplaceParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsBrotliCompression struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to select the content types you want to compress.
	//
	// `text/html` is a mandatory content type.
	//
	// Any of "application/javascript", "application/json",
	// "application/vnd.ms-fontobject", "application/wasm", "application/x-font-ttf",
	// "application/x-javascript", "application/xml", "application/xml+rss",
	// "image/svg+xml", "image/x-icon", "text/css", "text/html", "text/javascript",
	// "text/plain", "text/xml".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsBrowserCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Set the cache expiration time to '0s' to disable caching.
	//
	// The maximum duration is any equivalent to `1y`.
	Value string `json:"value,required" format:"nginx time"`
	paramObj
}

func (r ResourceReplaceParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsCacheHTTPHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool     `json:"enabled,required"`
	Value   []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsCors struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Value of the Access-Control-Allow-Origin header.
	//
	// Possible values:
	//
	//   - **Adds \* as the Access-Control-Allow-Origin header value** - Content will be
	//     uploaded for requests from any domain. `"value": ["*"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$`http_origin`" as the Access-Control-Allow-Origin header value** -
	//     Content will be uploaded for requests from any domain, and the domain from
	//     which the request was sent will be added to the "Access-Control-Allow-Origin"
	//     header in the response. `"value": ["$http_origin"]`
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceReplaceParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceReplaceParamsOptionsCountryACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of countries according to ISO-3166-1.
	//
	// The meaning of the parameter depends on `policy_type` value:
	//
	// - **allow** - List of countries for which access is prohibited.
	// - **deny** - List of countries for which access is allowed.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"country-code"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsCountryACL](
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
type ResourceReplaceParamsOptionsDisableCache struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - content caching is disabled.
	// - **false** - content caching is enabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsDisableProxyForceRanges struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceReplaceParamsOptionsEdgeCacheSettings struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceReplaceParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceReplaceParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceReplaceParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceReplaceParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceReplaceParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceReplaceParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceReplaceParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceReplaceParamsOptionsFastedgeOnRequestBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceReplaceParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceReplaceParamsOptionsFastedgeOnRequestHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceReplaceParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceReplaceParamsOptionsFastedgeOnResponseBody struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceReplaceParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceReplaceParamsOptionsFastedgeOnResponseHeaders struct {
	// The ID of the application in FastEdge.
	AppID string `json:"app_id,required"`
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

func (r ResourceReplaceParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsFetchCompressed struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceReplaceParamsOptionsFollowOriginRedirect struct {
	// Redirect status code that the origin server returns.
	//
	// To serve up to date content to end users, you will need to purge the cache after
	// managing the option.
	//
	// Any of 301, 302, 303, 307, 308.
	Codes []int64 `json:"codes,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceReplaceParamsOptionsForceReturn struct {
	// URL for redirection or text.
	Body string `json:"body,required"`
	// Status code value.
	Code int64 `json:"code,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Controls the time at which a custom HTTP response code should be applied. By
	// default, a custom HTTP response code is applied at any time.
	TimeInterval ResourceReplaceParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceReplaceParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceReplaceParamsOptionsForceReturnTimeInterval struct {
	// Time until which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	EndTime string `json:"end_time,required"`
	// Time from which a custom HTTP response code should be applied. Indicated in
	// 24-hour format.
	StartTime string `json:"start_time,required"`
	// Time zone used to calculate time.
	TimeZone param.Opt[string] `json:"time_zone,omitzero" format:"timezone"`
	paramObj
}

func (r ResourceReplaceParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsForwardHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsGzipOn struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsHostHeader struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Host Header value.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables HTTP/3 protocol for content delivery.
//
// `http3_enabled` option works only with `"sslEnabled": true`.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsHttp3Enabled struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsHttp3Enabled) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsHttp3Enabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsHttp3Enabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsIgnoreCookie struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Option is enabled, files with cookies are cached as one file.
	//   - **false** - Option is disabled, files with cookies are cached as different
	//     files.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsIgnoreQueryString struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceReplaceParamsOptionsImageStack struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceReplaceParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance. We recommend you use a script
// for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceReplaceParamsOptionsIPAddressACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"ipv4net or ipv6net"`
	// IP access policy type.
	//
	// Possible values:
	//
	//   - **allow** - Allow access to all IPs except IPs specified in
	//     "`excepted_values`" field.
	//   - **deny** - Deny access to all IPs except IPs specified in "`excepted_values`"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceReplaceParamsOptionsLimitBandwidth struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	LimitType string `json:"limit_type,omitzero,required"`
	// Amount of downloaded data after which the user will be rate limited.
	Buffer param.Opt[int64] `json:"buffer,omitzero"`
	// Maximum download speed per connection.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	paramObj
}

func (r ResourceReplaceParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsLimitBandwidth](
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
type ResourceReplaceParamsOptionsProxyCacheKey struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Key for caching.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsProxyCacheMethodsSet struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsProxyConnectTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 5s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsProxyReadTimeout struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Timeout value in seconds.
	//
	// Supported range: **1s - 30s**.
	Value string `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsQueryParamsBlacklist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsQueryParamsWhitelist struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of query parameters.
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsQueryStringForwarding struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// The `forward_from_files_types` field specifies the types of playlist files from
	// which parameters will be extracted and forwarded. This typically includes
	// formats that list multiple media chunk references, such as HLS and DASH
	// playlists. Parameters associated with these playlist files (like query strings
	// or headers) will be propagated to the chunks they reference.
	ForwardFromFileTypes []string `json:"forward_from_file_types,omitzero,required"`
	// The field specifies the types of media chunk files to which parameters,
	// extracted from playlist files, will be forwarded. These refer to the actual
	// segments of media content that are delivered to viewers. Ensuring the correct
	// parameters are forwarded to these files is crucial for maintaining the integrity
	// of the streaming session.
	ForwardToFileTypes []string `json:"forward_to_file_types,omitzero,required"`
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

func (r ResourceReplaceParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsRedirectHTTPToHTTPS struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsRedirectHTTPSToHTTP struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceReplaceParamsOptionsReferrerACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"domain or wildcard"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceReplaceParamsOptionsRequestLimiter struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Maximum request rate.
	Rate int64 `json:"rate,required"`
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

func (r ResourceReplaceParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceReplaceParamsOptionsResponseHeadersHidingPolicy struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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
	Excepted []string `json:"excepted,omitzero,required" format:"http_header"`
	// How HTTP headers are hidden from the response.
	//
	// Possible values:
	//
	//   - **show** - Hide only HTTP headers listed in the `excepted` field.
	//   - **hide** - Hide all HTTP headers except headers listed in the "excepted"
	//     field.
	//
	// Any of "hide", "show".
	Mode string `json:"mode,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceReplaceParamsOptionsRewrite struct {
	// Path for the Rewrite option.
	//
	// Example:
	//
	// - `/(.*) /media/$1`
	Body string `json:"body,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceReplaceParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceReplaceParamsOptionsSecureKey struct {
	// Key generated on your side that will be used for URL signing.
	Key param.Opt[string] `json:"key,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceReplaceParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsSecureKey](
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
type ResourceReplaceParamsOptionsSlice struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsSni struct {
	// Custom SNI hostname.
	//
	// It is required if `sni_type` is set to custom.
	CustomHostname string `json:"custom_hostname,required" format:"domain"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
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

func (r ResourceReplaceParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsStale struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Defines list of errors for which "Always online" option is applied.
	//
	// Any of "error", "http_403", "http_404", "http_429", "http_500", "http_502",
	// "http_503", "http_504", "invalid_header", "timeout", "updating".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                     `json:"enabled,required"`
	Value   []ResourceReplaceParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceReplaceParamsOptionsStaticResponseHeadersValue struct {
	// HTTP Header name.
	//
	// Restrictions:
	//
	// - Maximum 128 symbols.
	// - Latin letters (A-Z, a-z,) numbers (0-9,) dashes, and underscores only.
	Name string `json:"name,required"`
	// Header value.
	//
	// Restrictions:
	//
	//   - Maximum 512 symbols.
	//   - Letters (a-z), numbers (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+
	//     /|\";:?.,><{}[]).
	//   - Must start with a letter, number, asterisk or {.
	//   - Multiple values can be added.
	Value []string `json:"value,omitzero,required"`
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

func (r ResourceReplaceParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOptionsStaticHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 128 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value any `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsStaticRequestHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// A MAP for static headers in a format of `header_name: header_value`.
	//
	// Restrictions:
	//
	//   - **Header name** - Maximum 255 symbols, may contain Latin letters (A-Z, a-z),
	//     numbers (0-9), dashes, and underscores.
	//   - **Header value** - Maximum 512 symbols, may contain letters (a-z), numbers
	//     (0-9), spaces, and symbols (`~!@#%%^&\*()-\_=+ /|\";:?.,><{}[]). Must start
	//     with a letter, number, asterisk or {.
	Value map[string]string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of SSL/TLS protocol versions allowed for HTTPS connections from end users
// to the domain.
//
// When the option is disabled, all protocols versions are allowed.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsTlsVersions struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of SSL/TLS protocol versions (case sensitive).
	//
	// Any of "SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3".
	Value []string `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsTlsVersions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsTlsVersions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsTlsVersions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Let's Encrypt certificate chain.
//
// The specified chain will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsUseDefaultLeChain struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	//   - **true** - Default Let's Encrypt certificate chain. This is a deprecated
	//     version, use it only for compatibilities with Android devices 7.1.1 or lower.
	//   - **false** - Alternative Let's Encrypt certificate chain.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsUseDefaultLeChain) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsUseDefaultLeChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsUseDefaultLeChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS-01 challenge to issue a Let's Encrypt certificate for the resource.
//
// DNS service should be activated to enable this option.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsUseDns01LeChallenge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - DNS-01 challenge is used to issue Let's Encrypt certificate.
	// - **false** - HTTP-01 challenge is used to issue Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsUseDns01LeChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsUseDns01LeChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsUseDns01LeChallenge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RSA Let's Encrypt certificate type for the CDN resource.
//
// The specified value will be used during the next Let's Encrypt certificate issue
// or renewal.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsUseRsaLeCert struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - RSA Let's Encrypt certificate.
	// - **false** - ECDSA Let's Encrypt certificate.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsUseRsaLeCert) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsUseRsaLeCert
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsUseRsaLeCert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceReplaceParamsOptionsUserAgentACL struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// List of User-Agents that will be allowed/denied.
	//
	// The meaning of the parameter depends on `policy_type`:
	//
	// - **allow** - List of User-Agents for which access is prohibited.
	// - **deny** - List of User-Agents for which access is allowed.
	//
	// Use an empty string `""` to allow/deny access when the User-Agent header is
	// empty.
	ExceptedValues []string `json:"excepted_values,omitzero,required" format:"user_agent"`
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
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceReplaceParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsWaap struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceReplaceParamsOptionsWebsockets struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Value bool `json:"value,required"`
	paramObj
}

func (r ResourceReplaceParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceReplaceParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceReplaceParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceReplaceParamsOriginProtocol string

const (
	ResourceReplaceParamsOriginProtocolHTTP  ResourceReplaceParamsOriginProtocol = "HTTP"
	ResourceReplaceParamsOriginProtocolHTTPS ResourceReplaceParamsOriginProtocol = "HTTPS"
	ResourceReplaceParamsOriginProtocolMatch ResourceReplaceParamsOriginProtocol = "MATCH"
)
