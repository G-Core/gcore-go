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

// ResourceRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceRuleService] method instead.
type ResourceRuleService struct {
	Options []option.RequestOption
}

// NewResourceRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceRuleService(opts ...option.RequestOption) (r ResourceRuleService) {
	r = ResourceRuleService{}
	r.Options = opts
	return
}

// Create rule
func (r *ResourceRuleService) New(ctx context.Context, resourceID int64, body ResourceRuleNewParams, opts ...option.RequestOption) (res *CdnResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change rule
func (r *ResourceRuleService) Update(ctx context.Context, ruleID int64, params ResourceRuleUpdateParams, opts ...option.RequestOption) (res *CdnResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", params.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get rules list
func (r *ResourceRuleService) List(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *[]CdnResourceRule, err error) {
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
func (r *ResourceRuleService) Delete(ctx context.Context, ruleID int64, body ResourceRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", body.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get rule details
func (r *ResourceRuleService) Get(ctx context.Context, ruleID int64, query ResourceRuleGetParams, opts ...option.RequestOption) (res *CdnResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", query.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change rule
func (r *ResourceRuleService) Replace(ctx context.Context, ruleID int64, params ResourceRuleReplaceParams, opts ...option.RequestOption) (res *CdnResourceRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v", params.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CdnResourceRule struct {
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
	Options CdnResourceRuleOptions `json:"options"`
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
	OriginProtocol CdnResourceRuleOriginProtocol `json:"originProtocol"`
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
	OverrideOriginProtocol CdnResourceRuleOverrideOriginProtocol `json:"overrideOriginProtocol,nullable"`
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
func (r CdnResourceRule) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type CdnResourceRuleOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods CdnResourceRuleOptionsAllowedHTTPMethods `json:"allowedHttpMethods,nullable"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection CdnResourceRuleOptionsBotProtection `json:"bot_protection,nullable"`
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
	BrotliCompression CdnResourceRuleOptionsBrotliCompression `json:"brotli_compression,nullable"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings CdnResourceRuleOptionsBrowserCacheSettings `json:"browser_cache_settings,nullable"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders CdnResourceRuleOptionsCacheHTTPHeaders `json:"cache_http_headers,nullable"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors CdnResourceRuleOptionsCors `json:"cors,nullable"`
	// Enables control access to content for specified countries.
	CountryACL CdnResourceRuleOptionsCountryACL `json:"country_acl,nullable"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache CdnResourceRuleOptionsDisableCache `json:"disable_cache,nullable"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges CdnResourceRuleOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,nullable"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings CdnResourceRuleOptionsEdgeCacheSettings `json:"edge_cache_settings,nullable"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge CdnResourceRuleOptionsFastedge `json:"fastedge,nullable"`
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
	FetchCompressed CdnResourceRuleOptionsFetchCompressed `json:"fetch_compressed,nullable"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect CdnResourceRuleOptionsFollowOriginRedirect `json:"follow_origin_redirect,nullable"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn CdnResourceRuleOptionsForceReturn `json:"force_return,nullable"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader CdnResourceRuleOptionsForwardHostHeader `json:"forward_host_header,nullable"`
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
	GzipOn CdnResourceRuleOptionsGzipOn `json:"gzipOn,nullable"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader CdnResourceRuleOptionsHostHeader `json:"hostHeader,nullable"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie CdnResourceRuleOptionsIgnoreCookie `json:"ignore_cookie,nullable"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString CdnResourceRuleOptionsIgnoreQueryString `json:"ignoreQueryString,nullable"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack CdnResourceRuleOptionsImageStack `json:"image_stack,nullable"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL CdnResourceRuleOptionsIPAddressACL `json:"ip_address_acl,nullable"`
	// Allows to control the download speed per connection.
	LimitBandwidth CdnResourceRuleOptionsLimitBandwidth `json:"limit_bandwidth,nullable"`
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
	ProxyCacheKey CdnResourceRuleOptionsProxyCacheKey `json:"proxy_cache_key,nullable"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet CdnResourceRuleOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,nullable"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout CdnResourceRuleOptionsProxyConnectTimeout `json:"proxy_connect_timeout,nullable"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout CdnResourceRuleOptionsProxyReadTimeout `json:"proxy_read_timeout,nullable"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist CdnResourceRuleOptionsQueryParamsBlacklist `json:"query_params_blacklist,nullable"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist CdnResourceRuleOptionsQueryParamsWhitelist `json:"query_params_whitelist,nullable"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding CdnResourceRuleOptionsQueryStringForwarding `json:"query_string_forwarding,nullable"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS CdnResourceRuleOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,nullable"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP CdnResourceRuleOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,nullable"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL CdnResourceRuleOptionsReferrerACL `json:"referrer_acl,nullable"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter CdnResourceRuleOptionsRequestLimiter `json:"request_limiter,nullable"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy CdnResourceRuleOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,nullable"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite CdnResourceRuleOptionsRewrite `json:"rewrite,nullable"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey CdnResourceRuleOptionsSecureKey `json:"secure_key,nullable"`
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
	Slice CdnResourceRuleOptionsSlice `json:"slice,nullable"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni CdnResourceRuleOptionsSni `json:"sni,nullable"`
	// Serves stale cached content in case of origin unavailability.
	Stale CdnResourceRuleOptionsStale `json:"stale,nullable"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders CdnResourceRuleOptionsStaticResponseHeaders `json:"static_response_headers,nullable"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders CdnResourceRuleOptionsStaticHeaders `json:"staticHeaders,nullable"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders CdnResourceRuleOptionsStaticRequestHeaders `json:"staticRequestHeaders,nullable"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL CdnResourceRuleOptionsUserAgentACL `json:"user_agent_acl,nullable"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap CdnResourceRuleOptionsWaap `json:"waap,nullable"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets CdnResourceRuleOptionsWebsockets `json:"websockets,nullable"`
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
func (r CdnResourceRuleOptions) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
type CdnResourceRuleOptionsAllowedHTTPMethods struct {
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
func (r CdnResourceRuleOptionsAllowedHTTPMethods) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
type CdnResourceRuleOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge CdnResourceRuleOptionsBotProtectionBotChallenge `json:"bot_challenge,required"`
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
func (r CdnResourceRuleOptionsBotProtection) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type CdnResourceRuleOptionsBotProtectionBotChallenge struct {
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
func (r CdnResourceRuleOptionsBotProtectionBotChallenge) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsBrotliCompression struct {
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
func (r CdnResourceRuleOptionsBrotliCompression) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
type CdnResourceRuleOptionsBrowserCacheSettings struct {
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
func (r CdnResourceRuleOptionsBrowserCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
type CdnResourceRuleOptionsCacheHTTPHeaders struct {
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
func (r CdnResourceRuleOptionsCacheHTTPHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
type CdnResourceRuleOptionsCors struct {
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
func (r CdnResourceRuleOptionsCors) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
type CdnResourceRuleOptionsCountryACL struct {
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
func (r CdnResourceRuleOptionsCountryACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
type CdnResourceRuleOptionsDisableCache struct {
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
func (r CdnResourceRuleOptionsDisableCache) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
type CdnResourceRuleOptionsDisableProxyForceRanges struct {
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
func (r CdnResourceRuleOptionsDisableProxyForceRanges) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
type CdnResourceRuleOptionsEdgeCacheSettings struct {
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
func (r CdnResourceRuleOptionsEdgeCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
type CdnResourceRuleOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody CdnResourceRuleOptionsFastedgeOnRequestBody `json:"on_request_body"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders CdnResourceRuleOptionsFastedgeOnRequestHeaders `json:"on_request_headers"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody CdnResourceRuleOptionsFastedgeOnResponseBody `json:"on_response_body"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders CdnResourceRuleOptionsFastedgeOnResponseHeaders `json:"on_response_headers"`
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
func (r CdnResourceRuleOptionsFastedge) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
type CdnResourceRuleOptionsFastedgeOnRequestBody struct {
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
func (r CdnResourceRuleOptionsFastedgeOnRequestBody) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
type CdnResourceRuleOptionsFastedgeOnRequestHeaders struct {
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
func (r CdnResourceRuleOptionsFastedgeOnRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
type CdnResourceRuleOptionsFastedgeOnResponseBody struct {
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
func (r CdnResourceRuleOptionsFastedgeOnResponseBody) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
type CdnResourceRuleOptionsFastedgeOnResponseHeaders struct {
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
func (r CdnResourceRuleOptionsFastedgeOnResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsFetchCompressed struct {
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
func (r CdnResourceRuleOptionsFetchCompressed) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
type CdnResourceRuleOptionsFollowOriginRedirect struct {
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
func (r CdnResourceRuleOptionsFollowOriginRedirect) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
type CdnResourceRuleOptionsForceReturn struct {
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
	TimeInterval CdnResourceRuleOptionsForceReturnTimeInterval `json:"time_interval,nullable"`
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
func (r CdnResourceRuleOptionsForceReturn) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
type CdnResourceRuleOptionsForceReturnTimeInterval struct {
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
func (r CdnResourceRuleOptionsForceReturnTimeInterval) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CdnResourceRuleOptionsForwardHostHeader struct {
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
func (r CdnResourceRuleOptionsForwardHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsGzipOn struct {
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
func (r CdnResourceRuleOptionsGzipOn) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type CdnResourceRuleOptionsHostHeader struct {
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
func (r CdnResourceRuleOptionsHostHeader) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
type CdnResourceRuleOptionsIgnoreCookie struct {
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
func (r CdnResourceRuleOptionsIgnoreCookie) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceRuleOptionsIgnoreQueryString struct {
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
func (r CdnResourceRuleOptionsIgnoreQueryString) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
type CdnResourceRuleOptionsImageStack struct {
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
func (r CdnResourceRuleOptionsImageStack) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
type CdnResourceRuleOptionsIPAddressACL struct {
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
func (r CdnResourceRuleOptionsIPAddressACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to control the download speed per connection.
type CdnResourceRuleOptionsLimitBandwidth struct {
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
func (r CdnResourceRuleOptionsLimitBandwidth) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsProxyCacheKey struct {
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
func (r CdnResourceRuleOptionsProxyCacheKey) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
type CdnResourceRuleOptionsProxyCacheMethodsSet struct {
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
func (r CdnResourceRuleOptionsProxyCacheMethodsSet) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
type CdnResourceRuleOptionsProxyConnectTimeout struct {
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
func (r CdnResourceRuleOptionsProxyConnectTimeout) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
type CdnResourceRuleOptionsProxyReadTimeout struct {
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
func (r CdnResourceRuleOptionsProxyReadTimeout) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceRuleOptionsQueryParamsBlacklist struct {
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
func (r CdnResourceRuleOptionsQueryParamsBlacklist) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type CdnResourceRuleOptionsQueryParamsWhitelist struct {
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
func (r CdnResourceRuleOptionsQueryParamsWhitelist) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
type CdnResourceRuleOptionsQueryStringForwarding struct {
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
func (r CdnResourceRuleOptionsQueryStringForwarding) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CdnResourceRuleOptionsRedirectHTTPToHTTPS struct {
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
func (r CdnResourceRuleOptionsRedirectHTTPToHTTPS) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type CdnResourceRuleOptionsRedirectHTTPSToHTTP struct {
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
func (r CdnResourceRuleOptionsRedirectHTTPSToHTTP) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
type CdnResourceRuleOptionsReferrerACL struct {
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
func (r CdnResourceRuleOptionsReferrerACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option allows to limit the amount of HTTP requests.
type CdnResourceRuleOptionsRequestLimiter struct {
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
func (r CdnResourceRuleOptionsRequestLimiter) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hides HTTP headers from an origin server in the CDN response.
type CdnResourceRuleOptionsResponseHeadersHidingPolicy struct {
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
func (r CdnResourceRuleOptionsResponseHeadersHidingPolicy) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
type CdnResourceRuleOptionsRewrite struct {
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
func (r CdnResourceRuleOptionsRewrite) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
type CdnResourceRuleOptionsSecureKey struct {
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
func (r CdnResourceRuleOptionsSecureKey) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsSecureKey) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsSlice struct {
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
func (r CdnResourceRuleOptionsSlice) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsSlice) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOptionsSni struct {
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
func (r CdnResourceRuleOptionsSni) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serves stale cached content in case of origin unavailability.
type CdnResourceRuleOptionsStale struct {
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
func (r CdnResourceRuleOptionsStale) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
type CdnResourceRuleOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                               `json:"enabled,required"`
	Value   []CdnResourceRuleOptionsStaticResponseHeadersValue `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnResourceRuleOptionsStaticResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnResourceRuleOptionsStaticResponseHeadersValue struct {
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
func (r CdnResourceRuleOptionsStaticResponseHeadersValue) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
type CdnResourceRuleOptionsStaticHeaders struct {
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
func (r CdnResourceRuleOptionsStaticHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
type CdnResourceRuleOptionsStaticRequestHeaders struct {
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
func (r CdnResourceRuleOptionsStaticRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
type CdnResourceRuleOptionsUserAgentACL struct {
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
func (r CdnResourceRuleOptionsUserAgentACL) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to enable WAAP (Web Application and API Protection).
type CdnResourceRuleOptionsWaap struct {
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
func (r CdnResourceRuleOptionsWaap) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
type CdnResourceRuleOptionsWebsockets struct {
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
func (r CdnResourceRuleOptionsWebsockets) RawJSON() string { return r.JSON.raw }
func (r *CdnResourceRuleOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type CdnResourceRuleOriginProtocol string

const (
	CdnResourceRuleOriginProtocolHTTPS CdnResourceRuleOriginProtocol = "HTTPS"
	CdnResourceRuleOriginProtocolHTTP  CdnResourceRuleOriginProtocol = "HTTP"
	CdnResourceRuleOriginProtocolMatch CdnResourceRuleOriginProtocol = "MATCH"
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
type CdnResourceRuleOverrideOriginProtocol string

const (
	CdnResourceRuleOverrideOriginProtocolHTTPS CdnResourceRuleOverrideOriginProtocol = "HTTPS"
	CdnResourceRuleOverrideOriginProtocolHTTP  CdnResourceRuleOverrideOriginProtocol = "HTTP"
	CdnResourceRuleOverrideOriginProtocolMatch CdnResourceRuleOverrideOriginProtocol = "MATCH"
)

type ResourceRuleNewParams struct {
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
	OverrideOriginProtocol ResourceRuleNewParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options ResourceRuleNewParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r ResourceRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type ResourceRuleNewParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceRuleNewParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceRuleNewParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceRuleNewParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceRuleNewParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceRuleNewParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceRuleNewParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceRuleNewParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceRuleNewParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceRuleNewParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceRuleNewParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceRuleNewParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceRuleNewParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceRuleNewParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceRuleNewParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceRuleNewParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceRuleNewParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceRuleNewParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceRuleNewParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceRuleNewParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceRuleNewParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceRuleNewParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceRuleNewParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceRuleNewParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceRuleNewParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceRuleNewParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceRuleNewParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceRuleNewParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceRuleNewParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceRuleNewParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceRuleNewParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceRuleNewParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceRuleNewParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceRuleNewParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceRuleNewParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceRuleNewParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceRuleNewParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceRuleNewParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceRuleNewParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceRuleNewParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceRuleNewParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceRuleNewParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceRuleNewParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceRuleNewParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsAllowedHTTPMethods struct {
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

func (r ResourceRuleNewParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceRuleNewParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceRuleNewParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceRuleNewParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceRuleNewParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceRuleNewParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsBrotliCompression struct {
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

func (r ResourceRuleNewParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsBrowserCacheSettings struct {
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

func (r ResourceRuleNewParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsCacheHTTPHeaders struct {
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

func (r ResourceRuleNewParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsCors struct {
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

func (r ResourceRuleNewParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleNewParamsOptionsCountryACL struct {
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

func (r ResourceRuleNewParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsCountryACL](
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
type ResourceRuleNewParamsOptionsDisableCache struct {
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

func (r ResourceRuleNewParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsDisableProxyForceRanges struct {
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

func (r ResourceRuleNewParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceRuleNewParamsOptionsEdgeCacheSettings struct {
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

func (r ResourceRuleNewParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceRuleNewParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceRuleNewParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceRuleNewParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceRuleNewParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleNewParamsOptionsFastedgeOnRequestBody struct {
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

func (r ResourceRuleNewParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleNewParamsOptionsFastedgeOnResponseBody struct {
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

func (r ResourceRuleNewParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsFetchCompressed struct {
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

func (r ResourceRuleNewParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceRuleNewParamsOptionsFollowOriginRedirect struct {
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

func (r ResourceRuleNewParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceRuleNewParamsOptionsForceReturn struct {
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
	TimeInterval ResourceRuleNewParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceRuleNewParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceRuleNewParamsOptionsForceReturnTimeInterval struct {
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

func (r ResourceRuleNewParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsForwardHostHeader struct {
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

func (r ResourceRuleNewParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsGzipOn struct {
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

func (r ResourceRuleNewParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsHostHeader struct {
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

func (r ResourceRuleNewParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsIgnoreCookie struct {
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

func (r ResourceRuleNewParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsIgnoreQueryString struct {
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

func (r ResourceRuleNewParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceRuleNewParamsOptionsImageStack struct {
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

func (r ResourceRuleNewParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsIPAddressACL struct {
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

func (r ResourceRuleNewParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceRuleNewParamsOptionsLimitBandwidth struct {
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

func (r ResourceRuleNewParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsLimitBandwidth](
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
type ResourceRuleNewParamsOptionsProxyCacheKey struct {
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

func (r ResourceRuleNewParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsProxyCacheMethodsSet struct {
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

func (r ResourceRuleNewParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsProxyConnectTimeout struct {
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

func (r ResourceRuleNewParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsProxyReadTimeout struct {
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

func (r ResourceRuleNewParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsQueryParamsBlacklist struct {
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

func (r ResourceRuleNewParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsQueryParamsWhitelist struct {
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

func (r ResourceRuleNewParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsQueryStringForwarding struct {
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

func (r ResourceRuleNewParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleNewParamsOptionsReferrerACL struct {
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

func (r ResourceRuleNewParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceRuleNewParamsOptionsRequestLimiter struct {
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

func (r ResourceRuleNewParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceRuleNewParamsOptionsRewrite struct {
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

func (r ResourceRuleNewParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceRuleNewParamsOptionsSecureKey struct {
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

func (r ResourceRuleNewParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsSecureKey](
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
type ResourceRuleNewParamsOptionsSlice struct {
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

func (r ResourceRuleNewParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsSni struct {
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

func (r ResourceRuleNewParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsStale struct {
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

func (r ResourceRuleNewParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                     `json:"enabled,required"`
	Value   []ResourceRuleNewParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceRuleNewParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceRuleNewParamsOptionsStaticResponseHeadersValue struct {
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

func (r ResourceRuleNewParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOptionsStaticHeaders struct {
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

func (r ResourceRuleNewParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsStaticRequestHeaders struct {
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

func (r ResourceRuleNewParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleNewParamsOptionsUserAgentACL struct {
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

func (r ResourceRuleNewParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleNewParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsWaap struct {
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

func (r ResourceRuleNewParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceRuleNewParamsOptionsWebsockets struct {
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

func (r ResourceRuleNewParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleNewParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleNewParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceRuleNewParamsOverrideOriginProtocol string

const (
	ResourceRuleNewParamsOverrideOriginProtocolHTTPS ResourceRuleNewParamsOverrideOriginProtocol = "HTTPS"
	ResourceRuleNewParamsOverrideOriginProtocolHTTP  ResourceRuleNewParamsOverrideOriginProtocol = "HTTP"
	ResourceRuleNewParamsOverrideOriginProtocolMatch ResourceRuleNewParamsOverrideOriginProtocol = "MATCH"
)

type ResourceRuleUpdateParams struct {
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
	OverrideOriginProtocol ResourceRuleUpdateParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options ResourceRuleUpdateParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r ResourceRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type ResourceRuleUpdateParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceRuleUpdateParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceRuleUpdateParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceRuleUpdateParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceRuleUpdateParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceRuleUpdateParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceRuleUpdateParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceRuleUpdateParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceRuleUpdateParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceRuleUpdateParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceRuleUpdateParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceRuleUpdateParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceRuleUpdateParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceRuleUpdateParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceRuleUpdateParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceRuleUpdateParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceRuleUpdateParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceRuleUpdateParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceRuleUpdateParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceRuleUpdateParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceRuleUpdateParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceRuleUpdateParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceRuleUpdateParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceRuleUpdateParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceRuleUpdateParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceRuleUpdateParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceRuleUpdateParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceRuleUpdateParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceRuleUpdateParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceRuleUpdateParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceRuleUpdateParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceRuleUpdateParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceRuleUpdateParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceRuleUpdateParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceRuleUpdateParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceRuleUpdateParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceRuleUpdateParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceRuleUpdateParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceRuleUpdateParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceRuleUpdateParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceRuleUpdateParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceRuleUpdateParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsAllowedHTTPMethods struct {
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

func (r ResourceRuleUpdateParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceRuleUpdateParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsBrotliCompression struct {
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

func (r ResourceRuleUpdateParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsBrowserCacheSettings struct {
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

func (r ResourceRuleUpdateParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsCacheHTTPHeaders struct {
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

func (r ResourceRuleUpdateParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsCors struct {
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

func (r ResourceRuleUpdateParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleUpdateParamsOptionsCountryACL struct {
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

func (r ResourceRuleUpdateParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsCountryACL](
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
type ResourceRuleUpdateParamsOptionsDisableCache struct {
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

func (r ResourceRuleUpdateParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsDisableProxyForceRanges struct {
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

func (r ResourceRuleUpdateParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceRuleUpdateParamsOptionsEdgeCacheSettings struct {
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

func (r ResourceRuleUpdateParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceRuleUpdateParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody struct {
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

func (r ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody struct {
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

func (r ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsFetchCompressed struct {
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

func (r ResourceRuleUpdateParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceRuleUpdateParamsOptionsFollowOriginRedirect struct {
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

func (r ResourceRuleUpdateParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceRuleUpdateParamsOptionsForceReturn struct {
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
	TimeInterval ResourceRuleUpdateParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceRuleUpdateParamsOptionsForceReturnTimeInterval struct {
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

func (r ResourceRuleUpdateParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsForwardHostHeader struct {
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

func (r ResourceRuleUpdateParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsGzipOn struct {
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

func (r ResourceRuleUpdateParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsHostHeader struct {
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

func (r ResourceRuleUpdateParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsIgnoreCookie struct {
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

func (r ResourceRuleUpdateParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsIgnoreQueryString struct {
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

func (r ResourceRuleUpdateParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceRuleUpdateParamsOptionsImageStack struct {
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

func (r ResourceRuleUpdateParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsIPAddressACL struct {
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

func (r ResourceRuleUpdateParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceRuleUpdateParamsOptionsLimitBandwidth struct {
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

func (r ResourceRuleUpdateParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsLimitBandwidth](
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
type ResourceRuleUpdateParamsOptionsProxyCacheKey struct {
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

func (r ResourceRuleUpdateParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet struct {
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

func (r ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsProxyConnectTimeout struct {
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

func (r ResourceRuleUpdateParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsProxyReadTimeout struct {
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

func (r ResourceRuleUpdateParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsQueryParamsBlacklist struct {
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

func (r ResourceRuleUpdateParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsQueryParamsWhitelist struct {
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

func (r ResourceRuleUpdateParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsQueryStringForwarding struct {
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

func (r ResourceRuleUpdateParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleUpdateParamsOptionsReferrerACL struct {
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

func (r ResourceRuleUpdateParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceRuleUpdateParamsOptionsRequestLimiter struct {
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

func (r ResourceRuleUpdateParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceRuleUpdateParamsOptionsRewrite struct {
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

func (r ResourceRuleUpdateParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceRuleUpdateParamsOptionsSecureKey struct {
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

func (r ResourceRuleUpdateParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsSecureKey](
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
type ResourceRuleUpdateParamsOptionsSlice struct {
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

func (r ResourceRuleUpdateParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsSni struct {
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

func (r ResourceRuleUpdateParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsStale struct {
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

func (r ResourceRuleUpdateParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                        `json:"enabled,required"`
	Value   []ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceRuleUpdateParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue struct {
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

func (r ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOptionsStaticHeaders struct {
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

func (r ResourceRuleUpdateParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsStaticRequestHeaders struct {
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

func (r ResourceRuleUpdateParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleUpdateParamsOptionsUserAgentACL struct {
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

func (r ResourceRuleUpdateParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleUpdateParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsWaap struct {
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

func (r ResourceRuleUpdateParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceRuleUpdateParamsOptionsWebsockets struct {
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

func (r ResourceRuleUpdateParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleUpdateParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleUpdateParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceRuleUpdateParamsOverrideOriginProtocol string

const (
	ResourceRuleUpdateParamsOverrideOriginProtocolHTTPS ResourceRuleUpdateParamsOverrideOriginProtocol = "HTTPS"
	ResourceRuleUpdateParamsOverrideOriginProtocolHTTP  ResourceRuleUpdateParamsOverrideOriginProtocol = "HTTP"
	ResourceRuleUpdateParamsOverrideOriginProtocolMatch ResourceRuleUpdateParamsOverrideOriginProtocol = "MATCH"
)

type ResourceRuleDeleteParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	paramObj
}

type ResourceRuleGetParams struct {
	ResourceID int64 `path:"resource_id,required" json:"-"`
	paramObj
}

type ResourceRuleReplaceParams struct {
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
	OverrideOriginProtocol ResourceRuleReplaceParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options ResourceRuleReplaceParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r ResourceRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type ResourceRuleReplaceParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods ResourceRuleReplaceParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection ResourceRuleReplaceParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression ResourceRuleReplaceParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings ResourceRuleReplaceParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders ResourceRuleReplaceParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors ResourceRuleReplaceParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL ResourceRuleReplaceParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache ResourceRuleReplaceParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges ResourceRuleReplaceParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings ResourceRuleReplaceParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge ResourceRuleReplaceParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed ResourceRuleReplaceParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect ResourceRuleReplaceParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn ResourceRuleReplaceParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader ResourceRuleReplaceParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn ResourceRuleReplaceParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader ResourceRuleReplaceParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie ResourceRuleReplaceParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString ResourceRuleReplaceParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack ResourceRuleReplaceParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL ResourceRuleReplaceParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth ResourceRuleReplaceParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey ResourceRuleReplaceParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout ResourceRuleReplaceParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout ResourceRuleReplaceParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist ResourceRuleReplaceParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist ResourceRuleReplaceParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding ResourceRuleReplaceParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL ResourceRuleReplaceParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter ResourceRuleReplaceParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite ResourceRuleReplaceParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey ResourceRuleReplaceParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice ResourceRuleReplaceParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni ResourceRuleReplaceParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale ResourceRuleReplaceParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders ResourceRuleReplaceParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders ResourceRuleReplaceParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders ResourceRuleReplaceParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL ResourceRuleReplaceParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap ResourceRuleReplaceParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets ResourceRuleReplaceParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsAllowedHTTPMethods struct {
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

func (r ResourceRuleReplaceParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type ResourceRuleReplaceParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsBrotliCompression struct {
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

func (r ResourceRuleReplaceParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsBrowserCacheSettings struct {
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

func (r ResourceRuleReplaceParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsCacheHTTPHeaders struct {
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

func (r ResourceRuleReplaceParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsCors struct {
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

func (r ResourceRuleReplaceParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleReplaceParamsOptionsCountryACL struct {
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

func (r ResourceRuleReplaceParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsCountryACL](
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
type ResourceRuleReplaceParamsOptionsDisableCache struct {
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

func (r ResourceRuleReplaceParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsDisableProxyForceRanges struct {
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

func (r ResourceRuleReplaceParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type ResourceRuleReplaceParamsOptionsEdgeCacheSettings struct {
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

func (r ResourceRuleReplaceParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type ResourceRuleReplaceParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request.
	OnRequestHeaders ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody struct {
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

func (r ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody struct {
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

func (r ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsFetchCompressed struct {
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

func (r ResourceRuleReplaceParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type ResourceRuleReplaceParamsOptionsFollowOriginRedirect struct {
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

func (r ResourceRuleReplaceParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type ResourceRuleReplaceParamsOptionsForceReturn struct {
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
	TimeInterval ResourceRuleReplaceParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type ResourceRuleReplaceParamsOptionsForceReturnTimeInterval struct {
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

func (r ResourceRuleReplaceParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsForwardHostHeader struct {
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

func (r ResourceRuleReplaceParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsGzipOn struct {
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

func (r ResourceRuleReplaceParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsHostHeader struct {
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

func (r ResourceRuleReplaceParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsIgnoreCookie struct {
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

func (r ResourceRuleReplaceParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsIgnoreQueryString struct {
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

func (r ResourceRuleReplaceParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type ResourceRuleReplaceParamsOptionsImageStack struct {
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

func (r ResourceRuleReplaceParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsIPAddressACL struct {
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

func (r ResourceRuleReplaceParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type ResourceRuleReplaceParamsOptionsLimitBandwidth struct {
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

func (r ResourceRuleReplaceParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsLimitBandwidth](
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
type ResourceRuleReplaceParamsOptionsProxyCacheKey struct {
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

func (r ResourceRuleReplaceParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet struct {
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

func (r ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsProxyConnectTimeout struct {
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

func (r ResourceRuleReplaceParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsProxyReadTimeout struct {
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

func (r ResourceRuleReplaceParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsQueryParamsBlacklist struct {
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

func (r ResourceRuleReplaceParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsQueryParamsWhitelist struct {
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

func (r ResourceRuleReplaceParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsQueryStringForwarding struct {
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

func (r ResourceRuleReplaceParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleReplaceParamsOptionsReferrerACL struct {
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

func (r ResourceRuleReplaceParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type ResourceRuleReplaceParamsOptionsRequestLimiter struct {
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

func (r ResourceRuleReplaceParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type ResourceRuleReplaceParamsOptionsRewrite struct {
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

func (r ResourceRuleReplaceParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type ResourceRuleReplaceParamsOptionsSecureKey struct {
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

func (r ResourceRuleReplaceParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsSecureKey](
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
type ResourceRuleReplaceParamsOptionsSlice struct {
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

func (r ResourceRuleReplaceParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsSni struct {
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

func (r ResourceRuleReplaceParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsStale struct {
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

func (r ResourceRuleReplaceParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                         `json:"enabled,required"`
	Value   []ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r ResourceRuleReplaceParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue struct {
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

func (r ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOptionsStaticHeaders struct {
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

func (r ResourceRuleReplaceParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsStaticRequestHeaders struct {
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

func (r ResourceRuleReplaceParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type ResourceRuleReplaceParamsOptionsUserAgentACL struct {
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

func (r ResourceRuleReplaceParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResourceRuleReplaceParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsWaap struct {
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

func (r ResourceRuleReplaceParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type ResourceRuleReplaceParamsOptionsWebsockets struct {
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

func (r ResourceRuleReplaceParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow ResourceRuleReplaceParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceRuleReplaceParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type ResourceRuleReplaceParamsOverrideOriginProtocol string

const (
	ResourceRuleReplaceParamsOverrideOriginProtocolHTTPS ResourceRuleReplaceParamsOverrideOriginProtocol = "HTTPS"
	ResourceRuleReplaceParamsOverrideOriginProtocolHTTP  ResourceRuleReplaceParamsOverrideOriginProtocol = "HTTP"
	ResourceRuleReplaceParamsOverrideOriginProtocolMatch ResourceRuleReplaceParamsOverrideOriginProtocol = "MATCH"
)
