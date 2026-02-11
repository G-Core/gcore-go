// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CDNResourceRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCDNResourceRuleService] method instead.
type CDNResourceRuleService struct {
	Options []option.RequestOption
}

// NewCDNResourceRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCDNResourceRuleService(opts ...option.RequestOption) (r CDNResourceRuleService) {
	r = CDNResourceRuleService{}
	r.Options = opts
	return
}

// Create rule
func (r *CDNResourceRuleService) New(ctx context.Context, resourceID int64, body CDNResourceRuleNewParams, opts ...option.RequestOption) (res *CDNResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change rule
func (r *CDNResourceRuleService) Update(ctx context.Context, ruleID int64, params CDNResourceRuleUpdateParams, opts ...option.RequestOption) (res *CDNResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", params.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get rules list
func (r *CDNResourceRuleService) List(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *[]CDNResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the rule from the system permanently.
//
// Notes:
//
//   - **Deactivation Requirement**: Set the `active` attribute to `false` before
//     deletion.
//   - **Irreversibility**: This action is irreversible. Once deleted, the rule
//     cannot be recovered.
func (r *CDNResourceRuleService) Delete(ctx context.Context, ruleID int64, body CDNResourceRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", body.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get rule details
func (r *CDNResourceRuleService) Get(ctx context.Context, ruleID int64, query CDNResourceRuleGetParams, opts ...option.RequestOption) (res *CDNResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", query.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change rule
func (r *CDNResourceRuleService) Replace(ctx context.Context, ruleID int64, params CDNResourceRuleReplaceParams, opts ...option.RequestOption) (res *CDNResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", params.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CDNResourceRule struct {
	// Rule ID.
	ID int64 `json:"id"`
	// Enables or disables a rule.
	//
	// Possible values:
	//
	// - **true** - Rule is active, rule settings are applied.
	// - **false** - Rule is inactive, rule settings are not applied.
	Active bool `json:"active"`
	// Defines whether the rule has been deleted.
	//
	// Possible values:
	//
	// - **true** - Rule has been deleted.
	// - **false** - Rule has not been deleted.
	Deleted bool `json:"deleted"`
	// Rule name.
	Name string `json:"name"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options CDNResourceRuleOptions `json:"options"`
	// ID of the origin group to which the rule is applied.
	//
	// If the origin group is not specified, the rule is applied to the origin group
	// that the CDN resource is associated with.
	OriginGroup int64 `json:"originGroup,nullable"`
	// Protocol used by CDN servers to request content from an origin source.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
	//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
	//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
	//     on origin source should be available for the CDN both through HTTP and HTTPS
	//     protocols.
	//
	// Any of "HTTPS", "HTTP", "MATCH".
	OriginProtocol CDNResourceRuleOriginProtocol `json:"originProtocol"`
	// Sets a protocol other than the one specified in the CDN resource settings to
	// connect to the origin.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
	//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
	//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
	//     on origin source should be available for the CDN both through HTTP and HTTPS
	//     protocols.
	//   - **null** - `originProtocol` setting is inherited from the CDN resource
	//     settings.
	//
	// Any of "HTTPS", "HTTP", "MATCH".
	OverrideOriginProtocol CDNResourceRuleOverrideOriginProtocol `json:"overrideOriginProtocol,nullable"`
	// Defines whether the rule has an applied preset.
	//
	// Possible values:
	//
	// - **true** - Rule has a preset applied.
	// - **false** - Rule does not have a preset applied.
	//
	// If a preset is applied to the rule, the options included in the preset cannot be
	// edited for the rule.
	PresetApplied bool `json:"preset_applied"`
	// ID of the rule with which the current rule is synchronized within the CDN
	// resource shared cache zone feature.
	PrimaryRule int64 `json:"primary_rule,nullable"`
	// Path to the file or folder for which the rule will be applied.
	//
	// The rule is applied if the requested URI matches the rule path.
	//
	// We add a leading forward slash to any rule path. Specify a path without a
	// forward slash.
	Rule string `json:"rule"`
	// Rule type.
	//
	// Possible values:
	//
	//   - **Type 0** - Regular expression. Must start with '^/' or '/'.
	//   - **Type 1** - Regular expression. Note that for this rule type we automatically
	//     add / to each rule pattern before your regular expression. This type is
	//     **legacy**, please use Type 0.
	RuleType int64 `json:"ruleType"`
	// Rule execution order: from lowest (1) to highest.
	//
	// If requested URI matches multiple rules, the one higher in the order of the
	// rules will be applied.
	Weight int64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Active                 respjson.Field
		Deleted                respjson.Field
		Name                   respjson.Field
		Options                respjson.Field
		OriginGroup            respjson.Field
		OriginProtocol         respjson.Field
		OverrideOriginProtocol respjson.Field
		PresetApplied          respjson.Field
		PrimaryRule            respjson.Field
		Rule                   respjson.Field
		RuleType               respjson.Field
		Weight                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceRule) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type CDNResourceRuleOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceRuleOptionsAllowedHTTPMethods `json:"allowedHttpMethods,nullable"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceRuleOptionsBotProtection `json:"bot_protection,nullable"`
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
	BrotliCompression CDNResourceRuleOptionsBrotliCompression `json:"brotli_compression,nullable"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceRuleOptionsBrowserCacheSettings `json:"browser_cache_settings,nullable"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceRuleOptionsCacheHTTPHeaders `json:"cache_http_headers,nullable"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceRuleOptionsCors `json:"cors,nullable"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceRuleOptionsCountryACL `json:"country_acl,nullable"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceRuleOptionsDisableCache `json:"disable_cache,nullable"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceRuleOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,nullable"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceRuleOptionsEdgeCacheSettings `json:"edge_cache_settings,nullable"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceRuleOptionsFastedge `json:"fastedge,nullable"`
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
	FetchCompressed CDNResourceRuleOptionsFetchCompressed `json:"fetch_compressed,nullable"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceRuleOptionsFollowOriginRedirect `json:"follow_origin_redirect,nullable"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceRuleOptionsForceReturn `json:"force_return,nullable"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceRuleOptionsForwardHostHeader `json:"forward_host_header,nullable"`
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
	GzipOn CDNResourceRuleOptionsGzipOn `json:"gzipOn,nullable"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceRuleOptionsHostHeader `json:"hostHeader,nullable"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceRuleOptionsIgnoreCookie `json:"ignore_cookie,nullable"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceRuleOptionsIgnoreQueryString `json:"ignoreQueryString,nullable"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceRuleOptionsImageStack `json:"image_stack,nullable"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceRuleOptionsIPAddressACL `json:"ip_address_acl,nullable"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceRuleOptionsLimitBandwidth `json:"limit_bandwidth,nullable"`
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
	ProxyCacheKey CDNResourceRuleOptionsProxyCacheKey `json:"proxy_cache_key,nullable"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceRuleOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,nullable"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceRuleOptionsProxyConnectTimeout `json:"proxy_connect_timeout,nullable"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceRuleOptionsProxyReadTimeout `json:"proxy_read_timeout,nullable"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceRuleOptionsQueryParamsBlacklist `json:"query_params_blacklist,nullable"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceRuleOptionsQueryParamsWhitelist `json:"query_params_whitelist,nullable"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceRuleOptionsQueryStringForwarding `json:"query_string_forwarding,nullable"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceRuleOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,nullable"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceRuleOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,nullable"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceRuleOptionsReferrerACL `json:"referrer_acl,nullable"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceRuleOptionsRequestLimiter `json:"request_limiter,nullable"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceRuleOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,nullable"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceRuleOptionsRewrite `json:"rewrite,nullable"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceRuleOptionsSecureKey `json:"secure_key,nullable"`
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
	Slice CDNResourceRuleOptionsSlice `json:"slice,nullable"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceRuleOptionsSni `json:"sni,nullable"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceRuleOptionsStale `json:"stale,nullable"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceRuleOptionsStaticResponseHeaders `json:"static_response_headers,nullable"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceRuleOptionsStaticHeaders `json:"staticHeaders,nullable"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceRuleOptionsStaticRequestHeaders `json:"staticRequestHeaders,nullable"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceRuleOptionsUserAgentACL `json:"user_agent_acl,nullable"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceRuleOptionsWaap `json:"waap,nullable"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceRuleOptionsWebsockets `json:"websockets,nullable"`
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
		UserAgentACL                respjson.Field
		Waap                        respjson.Field
		Websockets                  respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceRuleOptions) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
type CDNResourceRuleOptionsAllowedHTTPMethods struct {
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
func (r CDNResourceRuleOptionsAllowedHTTPMethods) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
type CDNResourceRuleOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceRuleOptionsBotProtectionBotChallenge `json:"bot_challenge,required"`
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
func (r CDNResourceRuleOptionsBotProtection) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceRuleOptionsBotProtectionBotChallenge struct {
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
func (r CDNResourceRuleOptionsBotProtectionBotChallenge) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsBrotliCompression struct {
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
func (r CDNResourceRuleOptionsBrotliCompression) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
type CDNResourceRuleOptionsBrowserCacheSettings struct {
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
func (r CDNResourceRuleOptionsBrowserCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
type CDNResourceRuleOptionsCacheHTTPHeaders struct {
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
func (r CDNResourceRuleOptionsCacheHTTPHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
type CDNResourceRuleOptionsCors struct {
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
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
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
func (r CDNResourceRuleOptionsCors) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
type CDNResourceRuleOptionsCountryACL struct {
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
func (r CDNResourceRuleOptionsCountryACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
type CDNResourceRuleOptionsDisableCache struct {
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
func (r CDNResourceRuleOptionsDisableCache) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
type CDNResourceRuleOptionsDisableProxyForceRanges struct {
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
func (r CDNResourceRuleOptionsDisableProxyForceRanges) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
type CDNResourceRuleOptionsEdgeCacheSettings struct {
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
func (r CDNResourceRuleOptionsEdgeCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
type CDNResourceRuleOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceRuleOptionsFastedgeOnRequestBody `json:"on_request_body"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceRuleOptionsFastedgeOnRequestHeaders `json:"on_request_headers"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceRuleOptionsFastedgeOnResponseBody `json:"on_response_body"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceRuleOptionsFastedgeOnResponseHeaders `json:"on_response_headers"`
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
func (r CDNResourceRuleOptionsFastedge) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
type CDNResourceRuleOptionsFastedgeOnRequestBody struct {
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
func (r CDNResourceRuleOptionsFastedgeOnRequestBody) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
type CDNResourceRuleOptionsFastedgeOnRequestHeaders struct {
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
func (r CDNResourceRuleOptionsFastedgeOnRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
type CDNResourceRuleOptionsFastedgeOnResponseBody struct {
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
func (r CDNResourceRuleOptionsFastedgeOnResponseBody) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
type CDNResourceRuleOptionsFastedgeOnResponseHeaders struct {
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
func (r CDNResourceRuleOptionsFastedgeOnResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsFetchCompressed struct {
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
func (r CDNResourceRuleOptionsFetchCompressed) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
type CDNResourceRuleOptionsFollowOriginRedirect struct {
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
func (r CDNResourceRuleOptionsFollowOriginRedirect) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
type CDNResourceRuleOptionsForceReturn struct {
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
	TimeInterval CDNResourceRuleOptionsForceReturnTimeInterval `json:"time_interval,nullable"`
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
func (r CDNResourceRuleOptionsForceReturn) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
type CDNResourceRuleOptionsForceReturnTimeInterval struct {
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
func (r CDNResourceRuleOptionsForceReturnTimeInterval) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CDNResourceRuleOptionsForwardHostHeader struct {
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
func (r CDNResourceRuleOptionsForwardHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsGzipOn struct {
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
func (r CDNResourceRuleOptionsGzipOn) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CDNResourceRuleOptionsHostHeader struct {
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
func (r CDNResourceRuleOptionsHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
type CDNResourceRuleOptionsIgnoreCookie struct {
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
func (r CDNResourceRuleOptionsIgnoreCookie) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceRuleOptionsIgnoreQueryString struct {
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
func (r CDNResourceRuleOptionsIgnoreQueryString) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
type CDNResourceRuleOptionsImageStack struct {
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
func (r CDNResourceRuleOptionsImageStack) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
type CDNResourceRuleOptionsIPAddressACL struct {
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
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
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
func (r CDNResourceRuleOptionsIPAddressACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to control the download speed per connection.
type CDNResourceRuleOptionsLimitBandwidth struct {
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
func (r CDNResourceRuleOptionsLimitBandwidth) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsProxyCacheKey struct {
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
func (r CDNResourceRuleOptionsProxyCacheKey) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
type CDNResourceRuleOptionsProxyCacheMethodsSet struct {
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
func (r CDNResourceRuleOptionsProxyCacheMethodsSet) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
type CDNResourceRuleOptionsProxyConnectTimeout struct {
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
func (r CDNResourceRuleOptionsProxyConnectTimeout) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
type CDNResourceRuleOptionsProxyReadTimeout struct {
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
func (r CDNResourceRuleOptionsProxyReadTimeout) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceRuleOptionsQueryParamsBlacklist struct {
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
func (r CDNResourceRuleOptionsQueryParamsBlacklist) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CDNResourceRuleOptionsQueryParamsWhitelist struct {
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
func (r CDNResourceRuleOptionsQueryParamsWhitelist) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
type CDNResourceRuleOptionsQueryStringForwarding struct {
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
func (r CDNResourceRuleOptionsQueryStringForwarding) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CDNResourceRuleOptionsRedirectHTTPToHTTPS struct {
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
func (r CDNResourceRuleOptionsRedirectHTTPToHTTPS) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CDNResourceRuleOptionsRedirectHTTPSToHTTP struct {
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
func (r CDNResourceRuleOptionsRedirectHTTPSToHTTP) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
type CDNResourceRuleOptionsReferrerACL struct {
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
func (r CDNResourceRuleOptionsReferrerACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option allows to limit the amount of HTTP requests.
type CDNResourceRuleOptionsRequestLimiter struct {
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
func (r CDNResourceRuleOptionsRequestLimiter) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hides HTTP headers from an origin server in the CDN response.
type CDNResourceRuleOptionsResponseHeadersHidingPolicy struct {
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
func (r CDNResourceRuleOptionsResponseHeadersHidingPolicy) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
type CDNResourceRuleOptionsRewrite struct {
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
func (r CDNResourceRuleOptionsRewrite) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
type CDNResourceRuleOptionsSecureKey struct {
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
func (r CDNResourceRuleOptionsSecureKey) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsSecureKey) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsSlice struct {
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
func (r CDNResourceRuleOptionsSlice) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleOptionsSni struct {
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
func (r CDNResourceRuleOptionsSni) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serves stale cached content in case of origin unavailability.
type CDNResourceRuleOptionsStale struct {
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
func (r CDNResourceRuleOptionsStale) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
type CDNResourceRuleOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                               `json:"enabled,required"`
	Value   []CDNResourceRuleOptionsStaticResponseHeadersValue `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNResourceRuleOptionsStaticResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNResourceRuleOptionsStaticResponseHeadersValue struct {
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
func (r CDNResourceRuleOptionsStaticResponseHeadersValue) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
type CDNResourceRuleOptionsStaticHeaders struct {
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
func (r CDNResourceRuleOptionsStaticHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
type CDNResourceRuleOptionsStaticRequestHeaders struct {
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
func (r CDNResourceRuleOptionsStaticRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
type CDNResourceRuleOptionsUserAgentACL struct {
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
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
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
func (r CDNResourceRuleOptionsUserAgentACL) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to enable WAAP (Web Application and API Protection).
type CDNResourceRuleOptionsWaap struct {
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
func (r CDNResourceRuleOptionsWaap) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
type CDNResourceRuleOptionsWebsockets struct {
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
func (r CDNResourceRuleOptionsWebsockets) RawJSON() string { return r.JSON.raw }
func (r *CDNResourceRuleOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol used by CDN servers to request content from an origin source.
//
// Possible values:
//
//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
//     on origin source should be available for the CDN both through HTTP and HTTPS
//     protocols.
type CDNResourceRuleOriginProtocol string

const (
	CDNResourceRuleOriginProtocolHTTPS CDNResourceRuleOriginProtocol = "HTTPS"
	CDNResourceRuleOriginProtocolHTTP  CDNResourceRuleOriginProtocol = "HTTP"
	CDNResourceRuleOriginProtocolMatch CDNResourceRuleOriginProtocol = "MATCH"
)

// Sets a protocol other than the one specified in the CDN resource settings to
// connect to the origin.
//
// Possible values:
//
//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
//     on origin source should be available for the CDN both through HTTP and HTTPS
//     protocols.
//   - **null** - `originProtocol` setting is inherited from the CDN resource
//     settings.
type CDNResourceRuleOverrideOriginProtocol string

const (
	CDNResourceRuleOverrideOriginProtocolHTTPS CDNResourceRuleOverrideOriginProtocol = "HTTPS"
	CDNResourceRuleOverrideOriginProtocolHTTP  CDNResourceRuleOverrideOriginProtocol = "HTTP"
	CDNResourceRuleOverrideOriginProtocolMatch CDNResourceRuleOverrideOriginProtocol = "MATCH"
)

type CDNResourceRuleNewParams struct {
	// Rule name.
	Name string `json:"name,required"`
	// Path to the file or folder for which the rule will be applied.
	//
	// The rule is applied if the requested URI matches the rule path.
	//
	// We add a leading forward slash to any rule path. Specify a path without a
	// forward slash.
	Rule string `json:"rule,required"`
	// Rule type.
	//
	// Possible values:
	//
	//   - **Type 0** - Regular expression. Must start with '^/' or '/'.
	//   - **Type 1** - Regular expression. Note that for this rule type we automatically
	//     add / to each rule pattern before your regular expression. This type is
	//     **legacy**, please use Type 0.
	RuleType int64 `json:"ruleType,required"`
	// ID of the origin group to which the rule is applied.
	//
	// If the origin group is not specified, the rule is applied to the origin group
	// that the CDN resource is associated with.
	OriginGroup param.Opt[int64] `json:"originGroup,omitzero"`
	// Enables or disables a rule.
	//
	// Possible values:
	//
	// - **true** - Rule is active, rule settings are applied.
	// - **false** - Rule is inactive, rule settings are not applied.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Rule execution order: from lowest (1) to highest.
	//
	// If requested URI matches multiple rules, the one higher in the order of the
	// rules will be applied.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Sets a protocol other than the one specified in the CDN resource settings to
	// connect to the origin.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
	//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
	//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
	//     on origin source should be available for the CDN both through HTTP and HTTPS
	//     protocols.
	//   - **null** - `originProtocol` setting is inherited from the CDN resource
	//     settings.
	//
	// Any of "HTTPS", "HTTP", "MATCH".
	OverrideOriginProtocol CDNResourceRuleNewParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options CDNResourceRuleNewParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r CDNResourceRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type CDNResourceRuleNewParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceRuleNewParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceRuleNewParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression CDNResourceRuleNewParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceRuleNewParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceRuleNewParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceRuleNewParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceRuleNewParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceRuleNewParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceRuleNewParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceRuleNewParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceRuleNewParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed CDNResourceRuleNewParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceRuleNewParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceRuleNewParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceRuleNewParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn CDNResourceRuleNewParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceRuleNewParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceRuleNewParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceRuleNewParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceRuleNewParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceRuleNewParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceRuleNewParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey CDNResourceRuleNewParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceRuleNewParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceRuleNewParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceRuleNewParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceRuleNewParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceRuleNewParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceRuleNewParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceRuleNewParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceRuleNewParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceRuleNewParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice CDNResourceRuleNewParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceRuleNewParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceRuleNewParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceRuleNewParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceRuleNewParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceRuleNewParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceRuleNewParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceRuleNewParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceRuleNewParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsAllowedHTTPMethods struct {
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

func (r CDNResourceRuleNewParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceRuleNewParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsBrotliCompression struct {
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

func (r CDNResourceRuleNewParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsBrowserCacheSettings struct {
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

func (r CDNResourceRuleNewParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsCacheHTTPHeaders struct {
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

func (r CDNResourceRuleNewParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsCors struct {
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
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
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

func (r CDNResourceRuleNewParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleNewParamsOptionsCountryACL struct {
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

func (r CDNResourceRuleNewParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsCountryACL](
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
type CDNResourceRuleNewParamsOptionsDisableCache struct {
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

func (r CDNResourceRuleNewParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsDisableProxyForceRanges struct {
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

func (r CDNResourceRuleNewParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceRuleNewParamsOptionsEdgeCacheSettings struct {
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

func (r CDNResourceRuleNewParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceRuleNewParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody struct {
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

func (r CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody struct {
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

func (r CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsFetchCompressed struct {
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

func (r CDNResourceRuleNewParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceRuleNewParamsOptionsFollowOriginRedirect struct {
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

func (r CDNResourceRuleNewParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceRuleNewParamsOptionsForceReturn struct {
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
	TimeInterval CDNResourceRuleNewParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceRuleNewParamsOptionsForceReturnTimeInterval struct {
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

func (r CDNResourceRuleNewParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsForwardHostHeader struct {
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

func (r CDNResourceRuleNewParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsGzipOn struct {
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

func (r CDNResourceRuleNewParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsHostHeader struct {
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

func (r CDNResourceRuleNewParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsIgnoreCookie struct {
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

func (r CDNResourceRuleNewParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsIgnoreQueryString struct {
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

func (r CDNResourceRuleNewParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceRuleNewParamsOptionsImageStack struct {
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

func (r CDNResourceRuleNewParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsIPAddressACL struct {
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
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceRuleNewParamsOptionsLimitBandwidth struct {
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

func (r CDNResourceRuleNewParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsLimitBandwidth](
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
type CDNResourceRuleNewParamsOptionsProxyCacheKey struct {
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

func (r CDNResourceRuleNewParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet struct {
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

func (r CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsProxyConnectTimeout struct {
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

func (r CDNResourceRuleNewParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsProxyReadTimeout struct {
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

func (r CDNResourceRuleNewParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsQueryParamsBlacklist struct {
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

func (r CDNResourceRuleNewParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsQueryParamsWhitelist struct {
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

func (r CDNResourceRuleNewParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsQueryStringForwarding struct {
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

func (r CDNResourceRuleNewParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleNewParamsOptionsReferrerACL struct {
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

func (r CDNResourceRuleNewParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceRuleNewParamsOptionsRequestLimiter struct {
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

func (r CDNResourceRuleNewParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceRuleNewParamsOptionsRewrite struct {
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

func (r CDNResourceRuleNewParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceRuleNewParamsOptionsSecureKey struct {
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

func (r CDNResourceRuleNewParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsSecureKey](
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
type CDNResourceRuleNewParamsOptionsSlice struct {
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

func (r CDNResourceRuleNewParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsSni struct {
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

func (r CDNResourceRuleNewParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsStale struct {
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

func (r CDNResourceRuleNewParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                        `json:"enabled,required"`
	Value   []CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleNewParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue struct {
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

func (r CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleNewParamsOptionsStaticHeaders struct {
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

func (r CDNResourceRuleNewParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsStaticRequestHeaders struct {
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

func (r CDNResourceRuleNewParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleNewParamsOptionsUserAgentACL struct {
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
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
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

func (r CDNResourceRuleNewParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleNewParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsWaap struct {
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

func (r CDNResourceRuleNewParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceRuleNewParamsOptionsWebsockets struct {
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

func (r CDNResourceRuleNewParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleNewParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleNewParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets a protocol other than the one specified in the CDN resource settings to
// connect to the origin.
//
// Possible values:
//
//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
//     on origin source should be available for the CDN both through HTTP and HTTPS
//     protocols.
//   - **null** - `originProtocol` setting is inherited from the CDN resource
//     settings.
type CDNResourceRuleNewParamsOverrideOriginProtocol string

const (
	CDNResourceRuleNewParamsOverrideOriginProtocolHTTPS CDNResourceRuleNewParamsOverrideOriginProtocol = "HTTPS"
	CDNResourceRuleNewParamsOverrideOriginProtocolHTTP  CDNResourceRuleNewParamsOverrideOriginProtocol = "HTTP"
	CDNResourceRuleNewParamsOverrideOriginProtocolMatch CDNResourceRuleNewParamsOverrideOriginProtocol = "MATCH"
)

type CDNResourceRuleUpdateParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	// ID of the origin group to which the rule is applied.
	//
	// If the origin group is not specified, the rule is applied to the origin group
	// that the CDN resource is associated with.
	OriginGroup param.Opt[int64] `json:"originGroup,omitzero"`
	// Enables or disables a rule.
	//
	// Possible values:
	//
	// - **true** - Rule is active, rule settings are applied.
	// - **false** - Rule is inactive, rule settings are not applied.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Rule name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Path to the file or folder for which the rule will be applied.
	//
	// The rule is applied if the requested URI matches the rule path.
	//
	// We add a leading forward slash to any rule path. Specify a path without a
	// forward slash.
	Rule param.Opt[string] `json:"rule,omitzero"`
	// Rule type.
	//
	// Possible values:
	//
	//   - **Type 0** - Regular expression. Must start with '^/' or '/'.
	//   - **Type 1** - Regular expression. Note that for this rule type we automatically
	//     add / to each rule pattern before your regular expression. This type is
	//     **legacy**, please use Type 0.
	RuleType param.Opt[int64] `json:"ruleType,omitzero"`
	// Rule execution order: from lowest (1) to highest.
	//
	// If requested URI matches multiple rules, the one higher in the order of the
	// rules will be applied.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Sets a protocol other than the one specified in the CDN resource settings to
	// connect to the origin.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
	//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
	//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
	//     on origin source should be available for the CDN both through HTTP and HTTPS
	//     protocols.
	//   - **null** - `originProtocol` setting is inherited from the CDN resource
	//     settings.
	//
	// Any of "HTTPS", "HTTP", "MATCH".
	OverrideOriginProtocol CDNResourceRuleUpdateParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options CDNResourceRuleUpdateParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r CDNResourceRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type CDNResourceRuleUpdateParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceRuleUpdateParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression CDNResourceRuleUpdateParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceRuleUpdateParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceRuleUpdateParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceRuleUpdateParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceRuleUpdateParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed CDNResourceRuleUpdateParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceRuleUpdateParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceRuleUpdateParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn CDNResourceRuleUpdateParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceRuleUpdateParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceRuleUpdateParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceRuleUpdateParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceRuleUpdateParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceRuleUpdateParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceRuleUpdateParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey CDNResourceRuleUpdateParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceRuleUpdateParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceRuleUpdateParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceRuleUpdateParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceRuleUpdateParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceRuleUpdateParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceRuleUpdateParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice CDNResourceRuleUpdateParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceRuleUpdateParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceRuleUpdateParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceRuleUpdateParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceRuleUpdateParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceRuleUpdateParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceRuleUpdateParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods struct {
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

func (r CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceRuleUpdateParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsBrotliCompression struct {
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

func (r CDNResourceRuleUpdateParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings struct {
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

func (r CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders struct {
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

func (r CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsCors struct {
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
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
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

func (r CDNResourceRuleUpdateParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleUpdateParamsOptionsCountryACL struct {
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

func (r CDNResourceRuleUpdateParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsCountryACL](
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
type CDNResourceRuleUpdateParamsOptionsDisableCache struct {
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

func (r CDNResourceRuleUpdateParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges struct {
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

func (r CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings struct {
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

func (r CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceRuleUpdateParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsFetchCompressed struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect struct {
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

func (r CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceRuleUpdateParamsOptionsForceReturn struct {
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
	TimeInterval CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval struct {
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

func (r CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsForwardHostHeader struct {
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

func (r CDNResourceRuleUpdateParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsGzipOn struct {
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

func (r CDNResourceRuleUpdateParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsHostHeader struct {
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

func (r CDNResourceRuleUpdateParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsIgnoreCookie struct {
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

func (r CDNResourceRuleUpdateParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsIgnoreQueryString struct {
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

func (r CDNResourceRuleUpdateParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceRuleUpdateParamsOptionsImageStack struct {
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

func (r CDNResourceRuleUpdateParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsIPAddressACL struct {
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
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceRuleUpdateParamsOptionsLimitBandwidth struct {
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

func (r CDNResourceRuleUpdateParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsLimitBandwidth](
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
type CDNResourceRuleUpdateParamsOptionsProxyCacheKey struct {
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

func (r CDNResourceRuleUpdateParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet struct {
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

func (r CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout struct {
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

func (r CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsProxyReadTimeout struct {
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

func (r CDNResourceRuleUpdateParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist struct {
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

func (r CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist struct {
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

func (r CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsQueryStringForwarding struct {
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

func (r CDNResourceRuleUpdateParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleUpdateParamsOptionsReferrerACL struct {
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

func (r CDNResourceRuleUpdateParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceRuleUpdateParamsOptionsRequestLimiter struct {
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

func (r CDNResourceRuleUpdateParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceRuleUpdateParamsOptionsRewrite struct {
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

func (r CDNResourceRuleUpdateParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceRuleUpdateParamsOptionsSecureKey struct {
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

func (r CDNResourceRuleUpdateParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsSecureKey](
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
type CDNResourceRuleUpdateParamsOptionsSlice struct {
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

func (r CDNResourceRuleUpdateParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsSni struct {
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

func (r CDNResourceRuleUpdateParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsStale struct {
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

func (r CDNResourceRuleUpdateParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                           `json:"enabled,required"`
	Value   []CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue struct {
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

func (r CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleUpdateParamsOptionsStaticHeaders struct {
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

func (r CDNResourceRuleUpdateParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders struct {
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

func (r CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleUpdateParamsOptionsUserAgentACL struct {
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
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
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

func (r CDNResourceRuleUpdateParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleUpdateParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsWaap struct {
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

func (r CDNResourceRuleUpdateParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceRuleUpdateParamsOptionsWebsockets struct {
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

func (r CDNResourceRuleUpdateParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleUpdateParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleUpdateParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets a protocol other than the one specified in the CDN resource settings to
// connect to the origin.
//
// Possible values:
//
//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
//     on origin source should be available for the CDN both through HTTP and HTTPS
//     protocols.
//   - **null** - `originProtocol` setting is inherited from the CDN resource
//     settings.
type CDNResourceRuleUpdateParamsOverrideOriginProtocol string

const (
	CDNResourceRuleUpdateParamsOverrideOriginProtocolHTTPS CDNResourceRuleUpdateParamsOverrideOriginProtocol = "HTTPS"
	CDNResourceRuleUpdateParamsOverrideOriginProtocolHTTP  CDNResourceRuleUpdateParamsOverrideOriginProtocol = "HTTP"
	CDNResourceRuleUpdateParamsOverrideOriginProtocolMatch CDNResourceRuleUpdateParamsOverrideOriginProtocol = "MATCH"
)

type CDNResourceRuleDeleteParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	paramObj
}

type CDNResourceRuleGetParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	paramObj
}

type CDNResourceRuleReplaceParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	// Path to the file or folder for which the rule will be applied.
	//
	// The rule is applied if the requested URI matches the rule path.
	//
	// We add a leading forward slash to any rule path. Specify a path without a
	// forward slash.
	Rule string `json:"rule,required"`
	// Rule type.
	//
	// Possible values:
	//
	//   - **Type 0** - Regular expression. Must start with '^/' or '/'.
	//   - **Type 1** - Regular expression. Note that for this rule type we automatically
	//     add / to each rule pattern before your regular expression. This type is
	//     **legacy**, please use Type 0.
	RuleType int64 `json:"ruleType,required"`
	// ID of the origin group to which the rule is applied.
	//
	// If the origin group is not specified, the rule is applied to the origin group
	// that the CDN resource is associated with.
	OriginGroup param.Opt[int64] `json:"originGroup,omitzero"`
	// Enables or disables a rule.
	//
	// Possible values:
	//
	// - **true** - Rule is active, rule settings are applied.
	// - **false** - Rule is inactive, rule settings are not applied.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Rule name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Rule execution order: from lowest (1) to highest.
	//
	// If requested URI matches multiple rules, the one higher in the order of the
	// rules will be applied.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Sets a protocol other than the one specified in the CDN resource settings to
	// connect to the origin.
	//
	// Possible values:
	//
	//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
	//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
	//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
	//     on origin source should be available for the CDN both through HTTP and HTTPS
	//     protocols.
	//   - **null** - `originProtocol` setting is inherited from the CDN resource
	//     settings.
	//
	// Any of "HTTPS", "HTTP", "MATCH".
	OverrideOriginProtocol CDNResourceRuleReplaceParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options CDNResourceRuleReplaceParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r CDNResourceRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type CDNResourceRuleReplaceParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CDNResourceRuleReplaceParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression CDNResourceRuleReplaceParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CDNResourceRuleReplaceParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL CDNResourceRuleReplaceParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CDNResourceRuleReplaceParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CDNResourceRuleReplaceParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed CDNResourceRuleReplaceParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CDNResourceRuleReplaceParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CDNResourceRuleReplaceParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn CDNResourceRuleReplaceParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CDNResourceRuleReplaceParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CDNResourceRuleReplaceParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CDNResourceRuleReplaceParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CDNResourceRuleReplaceParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CDNResourceRuleReplaceParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth CDNResourceRuleReplaceParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey CDNResourceRuleReplaceParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CDNResourceRuleReplaceParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CDNResourceRuleReplaceParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CDNResourceRuleReplaceParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CDNResourceRuleReplaceParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CDNResourceRuleReplaceParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CDNResourceRuleReplaceParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice CDNResourceRuleReplaceParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CDNResourceRuleReplaceParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale CDNResourceRuleReplaceParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CDNResourceRuleReplaceParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CDNResourceRuleReplaceParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CDNResourceRuleReplaceParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CDNResourceRuleReplaceParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods struct {
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

func (r CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type CDNResourceRuleReplaceParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsBrotliCompression struct {
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

func (r CDNResourceRuleReplaceParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings struct {
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

func (r CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders struct {
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

func (r CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsCors struct {
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
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value if the
	//     origin matches one of the listed domains** - Content will be uploaded only for
	//     requests from the domains specified in the field.
	//     `"value": ["domain.com", "second.dom.com"]`
	//   - **Adds "$http_origin" as the Access-Control-Allow-Origin header value** -
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

func (r CDNResourceRuleReplaceParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleReplaceParamsOptionsCountryACL struct {
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

func (r CDNResourceRuleReplaceParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsCountryACL](
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
type CDNResourceRuleReplaceParamsOptionsDisableCache struct {
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

func (r CDNResourceRuleReplaceParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges struct {
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

func (r CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings struct {
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

func (r CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type CDNResourceRuleReplaceParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsFetchCompressed struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect struct {
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

func (r CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type CDNResourceRuleReplaceParamsOptionsForceReturn struct {
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
	TimeInterval CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval struct {
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

func (r CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsForwardHostHeader struct {
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

func (r CDNResourceRuleReplaceParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsGzipOn struct {
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

func (r CDNResourceRuleReplaceParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsHostHeader struct {
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

func (r CDNResourceRuleReplaceParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsIgnoreCookie struct {
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

func (r CDNResourceRuleReplaceParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsIgnoreQueryString struct {
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

func (r CDNResourceRuleReplaceParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type CDNResourceRuleReplaceParamsOptionsImageStack struct {
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

func (r CDNResourceRuleReplaceParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsIPAddressACL struct {
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
	//   - **allow** - Allow access to all IPs except IPs specified in "excepted_values"
	//     field.
	//   - **deny** - Deny access to all IPs except IPs specified in "excepted_values"
	//     field.
	//
	// Any of "allow", "deny".
	PolicyType string `json:"policy_type,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type CDNResourceRuleReplaceParamsOptionsLimitBandwidth struct {
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

func (r CDNResourceRuleReplaceParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsLimitBandwidth](
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
type CDNResourceRuleReplaceParamsOptionsProxyCacheKey struct {
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

func (r CDNResourceRuleReplaceParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet struct {
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

func (r CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout struct {
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

func (r CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsProxyReadTimeout struct {
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

func (r CDNResourceRuleReplaceParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist struct {
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

func (r CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist struct {
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

func (r CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsQueryStringForwarding struct {
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

func (r CDNResourceRuleReplaceParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleReplaceParamsOptionsReferrerACL struct {
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

func (r CDNResourceRuleReplaceParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type CDNResourceRuleReplaceParamsOptionsRequestLimiter struct {
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

func (r CDNResourceRuleReplaceParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type CDNResourceRuleReplaceParamsOptionsRewrite struct {
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

func (r CDNResourceRuleReplaceParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type CDNResourceRuleReplaceParamsOptionsSecureKey struct {
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

func (r CDNResourceRuleReplaceParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsSecureKey](
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
type CDNResourceRuleReplaceParamsOptionsSlice struct {
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

func (r CDNResourceRuleReplaceParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsSni struct {
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

func (r CDNResourceRuleReplaceParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsStale struct {
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

func (r CDNResourceRuleReplaceParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                            `json:"enabled,required"`
	Value   []CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue struct {
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

func (r CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type CDNResourceRuleReplaceParamsOptionsStaticHeaders struct {
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

func (r CDNResourceRuleReplaceParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders struct {
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

func (r CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type CDNResourceRuleReplaceParamsOptionsUserAgentACL struct {
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
	// You can provide exact User-Agent strings or regular expressions. Regular
	// expressions must start with `~` (case-sensitive) or `~*` (case-insensitive).
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

func (r CDNResourceRuleReplaceParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CDNResourceRuleReplaceParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsWaap struct {
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

func (r CDNResourceRuleReplaceParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type CDNResourceRuleReplaceParamsOptionsWebsockets struct {
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

func (r CDNResourceRuleReplaceParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow CDNResourceRuleReplaceParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CDNResourceRuleReplaceParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets a protocol other than the one specified in the CDN resource settings to
// connect to the origin.
//
// Possible values:
//
//   - **HTTPS** - CDN servers connect to origin via HTTPS protocol.
//   - **HTTP** - CDN servers connect to origin via HTTP protocol.
//   - **MATCH** - Connection protocol is chosen automatically; in this case, content
//     on origin source should be available for the CDN both through HTTP and HTTPS
//     protocols.
//   - **null** - `originProtocol` setting is inherited from the CDN resource
//     settings.
type CDNResourceRuleReplaceParamsOverrideOriginProtocol string

const (
	CDNResourceRuleReplaceParamsOverrideOriginProtocolHTTPS CDNResourceRuleReplaceParamsOverrideOriginProtocol = "HTTPS"
	CDNResourceRuleReplaceParamsOverrideOriginProtocolHTTP  CDNResourceRuleReplaceParamsOverrideOriginProtocol = "HTTP"
	CDNResourceRuleReplaceParamsOverrideOriginProtocolMatch CDNResourceRuleReplaceParamsOverrideOriginProtocol = "MATCH"
)
