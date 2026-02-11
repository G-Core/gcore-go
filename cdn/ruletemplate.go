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

// RuleTemplateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleTemplateService] method instead.
type RuleTemplateService struct {
	Options []option.RequestOption
}

// NewRuleTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleTemplateService(opts ...option.RequestOption) (r RuleTemplateService) {
	r = RuleTemplateService{}
	r.Options = opts
	return
}

// Create rule template
func (r *RuleTemplateService) New(ctx context.Context, body RuleTemplateNewParams, opts ...option.RequestOption) (res *RuleTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/resources/rule_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change rule template
func (r *RuleTemplateService) Update(ctx context.Context, ruleTemplateID int64, body RuleTemplateUpdateParams, opts ...option.RequestOption) (res *RuleTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/rule_templates/%v", ruleTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get rule templates list
func (r *RuleTemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *RuleTemplateList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/resources/rule_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete rule template
func (r *RuleTemplateService) Delete(ctx context.Context, ruleTemplateID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/resources/rule_templates/%v", ruleTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get rule template details
func (r *RuleTemplateService) Get(ctx context.Context, ruleTemplateID int64, opts ...option.RequestOption) (res *RuleTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/rule_templates/%v", ruleTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change rule template
func (r *RuleTemplateService) Replace(ctx context.Context, ruleTemplateID int64, body RuleTemplateReplaceParams, opts ...option.RequestOption) (res *RuleTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/rule_templates/%v", ruleTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type RuleTemplate struct {
	// Rule template ID.
	ID int64 `json:"id"`
	// Client ID
	Client int64 `json:"client"`
	// Defines whether the template is a system template developed for common cases.
	// System templates are available to all customers.
	//
	// Possible values:
	//
	// - **true** - Template is a system template and cannot be changed by a user.
	// - **false** - Template is a custom template and can be changed by a user.
	Default bool `json:"default"`
	// Defines whether the template has been deleted.
	//
	// Possible values:
	//
	// - **true** - Template has been deleted.
	// - **false** - Template has not been deleted.
	Deleted bool `json:"deleted"`
	// Rule template name.
	Name string `json:"name"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options RuleTemplateOptions `json:"options"`
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
	OverrideOriginProtocol RuleTemplateOverrideOriginProtocol `json:"overrideOriginProtocol,nullable"`
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
	// Determines whether the rule is a template.
	Template bool `json:"template"`
	// Rule execution order: from lowest (1) to highest.
	//
	// If requested URI matches multiple rules, the one higher in the order of the
	// rules will be applied.
	Weight int64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Client                 respjson.Field
		Default                respjson.Field
		Deleted                respjson.Field
		Name                   respjson.Field
		Options                respjson.Field
		OverrideOriginProtocol respjson.Field
		Rule                   respjson.Field
		RuleType               respjson.Field
		Template               respjson.Field
		Weight                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleTemplate) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type RuleTemplateOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods RuleTemplateOptionsAllowedHTTPMethods `json:"allowedHttpMethods,nullable"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection RuleTemplateOptionsBotProtection `json:"bot_protection,nullable"`
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
	BrotliCompression RuleTemplateOptionsBrotliCompression `json:"brotli_compression,nullable"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings RuleTemplateOptionsBrowserCacheSettings `json:"browser_cache_settings,nullable"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders RuleTemplateOptionsCacheHTTPHeaders `json:"cache_http_headers,nullable"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors RuleTemplateOptionsCors `json:"cors,nullable"`
	// Enables control access to content for specified countries.
	CountryACL RuleTemplateOptionsCountryACL `json:"country_acl,nullable"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache RuleTemplateOptionsDisableCache `json:"disable_cache,nullable"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges RuleTemplateOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,nullable"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings RuleTemplateOptionsEdgeCacheSettings `json:"edge_cache_settings,nullable"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge RuleTemplateOptionsFastedge `json:"fastedge,nullable"`
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
	FetchCompressed RuleTemplateOptionsFetchCompressed `json:"fetch_compressed,nullable"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect RuleTemplateOptionsFollowOriginRedirect `json:"follow_origin_redirect,nullable"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn RuleTemplateOptionsForceReturn `json:"force_return,nullable"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader RuleTemplateOptionsForwardHostHeader `json:"forward_host_header,nullable"`
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
	GzipOn RuleTemplateOptionsGzipOn `json:"gzipOn,nullable"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader RuleTemplateOptionsHostHeader `json:"hostHeader,nullable"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie RuleTemplateOptionsIgnoreCookie `json:"ignore_cookie,nullable"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString RuleTemplateOptionsIgnoreQueryString `json:"ignoreQueryString,nullable"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack RuleTemplateOptionsImageStack `json:"image_stack,nullable"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL RuleTemplateOptionsIPAddressACL `json:"ip_address_acl,nullable"`
	// Allows to control the download speed per connection.
	LimitBandwidth RuleTemplateOptionsLimitBandwidth `json:"limit_bandwidth,nullable"`
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
	ProxyCacheKey RuleTemplateOptionsProxyCacheKey `json:"proxy_cache_key,nullable"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet RuleTemplateOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,nullable"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout RuleTemplateOptionsProxyConnectTimeout `json:"proxy_connect_timeout,nullable"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout RuleTemplateOptionsProxyReadTimeout `json:"proxy_read_timeout,nullable"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist RuleTemplateOptionsQueryParamsBlacklist `json:"query_params_blacklist,nullable"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist RuleTemplateOptionsQueryParamsWhitelist `json:"query_params_whitelist,nullable"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding RuleTemplateOptionsQueryStringForwarding `json:"query_string_forwarding,nullable"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS RuleTemplateOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,nullable"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP RuleTemplateOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,nullable"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL RuleTemplateOptionsReferrerACL `json:"referrer_acl,nullable"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter RuleTemplateOptionsRequestLimiter `json:"request_limiter,nullable"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy RuleTemplateOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,nullable"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite RuleTemplateOptionsRewrite `json:"rewrite,nullable"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey RuleTemplateOptionsSecureKey `json:"secure_key,nullable"`
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
	Slice RuleTemplateOptionsSlice `json:"slice,nullable"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni RuleTemplateOptionsSni `json:"sni,nullable"`
	// Serves stale cached content in case of origin unavailability.
	Stale RuleTemplateOptionsStale `json:"stale,nullable"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders RuleTemplateOptionsStaticResponseHeaders `json:"static_response_headers,nullable"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders RuleTemplateOptionsStaticHeaders `json:"staticHeaders,nullable"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders RuleTemplateOptionsStaticRequestHeaders `json:"staticRequestHeaders,nullable"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL RuleTemplateOptionsUserAgentACL `json:"user_agent_acl,nullable"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap RuleTemplateOptionsWaap `json:"waap,nullable"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets RuleTemplateOptionsWebsockets `json:"websockets,nullable"`
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
func (r RuleTemplateOptions) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
type RuleTemplateOptionsAllowedHTTPMethods struct {
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
func (r RuleTemplateOptionsAllowedHTTPMethods) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
type RuleTemplateOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge RuleTemplateOptionsBotProtectionBotChallenge `json:"bot_challenge,required"`
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
func (r RuleTemplateOptionsBotProtection) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type RuleTemplateOptionsBotProtectionBotChallenge struct {
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
func (r RuleTemplateOptionsBotProtectionBotChallenge) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsBrotliCompression struct {
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
func (r RuleTemplateOptionsBrotliCompression) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for users browsers in seconds.
//
// Cache expiration time is applied to the following response codes: 200, 201, 204,
// 206, 301, 302, 303, 304, 307, 308.
//
// Responses with other codes will not be cached.
type RuleTemplateOptionsBrowserCacheSettings struct {
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
func (r RuleTemplateOptionsBrowserCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
type RuleTemplateOptionsCacheHTTPHeaders struct {
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
func (r RuleTemplateOptionsCacheHTTPHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
type RuleTemplateOptionsCors struct {
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
func (r RuleTemplateOptionsCors) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
type RuleTemplateOptionsCountryACL struct {
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
func (r RuleTemplateOptionsCountryACL) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `edge_cache_settings` option instead.
//
// Allows the complete disabling of content caching.
//
// Deprecated: deprecated
type RuleTemplateOptionsDisableCache struct {
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
func (r RuleTemplateOptionsDisableCache) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
type RuleTemplateOptionsDisableProxyForceRanges struct {
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
func (r RuleTemplateOptionsDisableProxyForceRanges) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
type RuleTemplateOptionsEdgeCacheSettings struct {
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
func (r RuleTemplateOptionsEdgeCacheSettings) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
type RuleTemplateOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody RuleTemplateOptionsFastedgeOnRequestBody `json:"on_request_body"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders RuleTemplateOptionsFastedgeOnRequestHeaders `json:"on_request_headers"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody RuleTemplateOptionsFastedgeOnResponseBody `json:"on_response_body"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders RuleTemplateOptionsFastedgeOnResponseHeaders `json:"on_response_headers"`
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
func (r RuleTemplateOptionsFastedge) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
type RuleTemplateOptionsFastedgeOnRequestBody struct {
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
func (r RuleTemplateOptionsFastedgeOnRequestBody) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
type RuleTemplateOptionsFastedgeOnRequestHeaders struct {
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
func (r RuleTemplateOptionsFastedgeOnRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
type RuleTemplateOptionsFastedgeOnResponseBody struct {
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
func (r RuleTemplateOptionsFastedgeOnResponseBody) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
type RuleTemplateOptionsFastedgeOnResponseHeaders struct {
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
func (r RuleTemplateOptionsFastedgeOnResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsFetchCompressed struct {
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
func (r RuleTemplateOptionsFetchCompressed) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
type RuleTemplateOptionsFollowOriginRedirect struct {
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
func (r RuleTemplateOptionsFollowOriginRedirect) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
type RuleTemplateOptionsForceReturn struct {
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
	TimeInterval RuleTemplateOptionsForceReturnTimeInterval `json:"time_interval,nullable"`
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
func (r RuleTemplateOptionsForceReturn) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
type RuleTemplateOptionsForceReturnTimeInterval struct {
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
func (r RuleTemplateOptionsForceReturnTimeInterval) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type RuleTemplateOptionsForwardHostHeader struct {
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
func (r RuleTemplateOptionsForwardHostHeader) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsGzipOn struct {
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
func (r RuleTemplateOptionsGzipOn) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsGzipOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the Host header that CDN servers use when request content from an origin
// server. Your server must be able to process requests with the chosen header.
//
// If the option is `null`, the Host Header value is equal to first CNAME.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
type RuleTemplateOptionsHostHeader struct {
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
func (r RuleTemplateOptionsHostHeader) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
type RuleTemplateOptionsIgnoreCookie struct {
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
func (r RuleTemplateOptionsIgnoreCookie) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type RuleTemplateOptionsIgnoreQueryString struct {
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
func (r RuleTemplateOptionsIgnoreQueryString) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
type RuleTemplateOptionsImageStack struct {
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
func (r RuleTemplateOptionsImageStack) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsImageStack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specific IP addresses.
//
// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
// you have to independently monitor their relevance.
//
// We recommend you use a script for automatically update IP ACL.
// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
type RuleTemplateOptionsIPAddressACL struct {
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
func (r RuleTemplateOptionsIPAddressACL) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to control the download speed per connection.
type RuleTemplateOptionsLimitBandwidth struct {
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
func (r RuleTemplateOptionsLimitBandwidth) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsProxyCacheKey struct {
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
func (r RuleTemplateOptionsProxyCacheKey) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
type RuleTemplateOptionsProxyCacheMethodsSet struct {
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
func (r RuleTemplateOptionsProxyCacheMethodsSet) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
type RuleTemplateOptionsProxyConnectTimeout struct {
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
func (r RuleTemplateOptionsProxyConnectTimeout) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
type RuleTemplateOptionsProxyReadTimeout struct {
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
func (r RuleTemplateOptionsProxyReadTimeout) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type RuleTemplateOptionsQueryParamsBlacklist struct {
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
func (r RuleTemplateOptionsQueryParamsBlacklist) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
type RuleTemplateOptionsQueryParamsWhitelist struct {
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
func (r RuleTemplateOptionsQueryParamsWhitelist) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Query String Forwarding feature allows for the seamless transfer of
// parameters embedded in playlist files to the corresponding media chunk files.
// This functionality ensures that specific attributes, such as authentication
// tokens or tracking information, are consistently passed along from the playlist
// manifest to the individual media segments. This is particularly useful for
// maintaining continuity in security, analytics, and any other parameter-based
// operations across the entire media delivery workflow.
type RuleTemplateOptionsQueryStringForwarding struct {
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
func (r RuleTemplateOptionsQueryStringForwarding) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type RuleTemplateOptionsRedirectHTTPToHTTPS struct {
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
func (r RuleTemplateOptionsRedirectHTTPToHTTPS) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
type RuleTemplateOptionsRedirectHTTPSToHTTP struct {
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
func (r RuleTemplateOptionsRedirectHTTPSToHTTP) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
type RuleTemplateOptionsReferrerACL struct {
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
func (r RuleTemplateOptionsReferrerACL) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Option allows to limit the amount of HTTP requests.
type RuleTemplateOptionsRequestLimiter struct {
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
func (r RuleTemplateOptionsRequestLimiter) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hides HTTP headers from an origin server in the CDN response.
type RuleTemplateOptionsResponseHeadersHidingPolicy struct {
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
func (r RuleTemplateOptionsResponseHeadersHidingPolicy) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
type RuleTemplateOptionsRewrite struct {
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
func (r RuleTemplateOptionsRewrite) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
type RuleTemplateOptionsSecureKey struct {
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
func (r RuleTemplateOptionsSecureKey) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsSecureKey) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsSlice struct {
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
func (r RuleTemplateOptionsSlice) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsSlice) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOptionsSni struct {
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
func (r RuleTemplateOptionsSni) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serves stale cached content in case of origin unavailability.
type RuleTemplateOptionsStale struct {
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
func (r RuleTemplateOptionsStale) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
type RuleTemplateOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                            `json:"enabled,required"`
	Value   []RuleTemplateOptionsStaticResponseHeadersValue `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleTemplateOptionsStaticResponseHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleTemplateOptionsStaticResponseHeadersValue struct {
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
func (r RuleTemplateOptionsStaticResponseHeadersValue) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `static_response_headers` option instead.
//
// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
// Headers can be specified. May contain a header with multiple values.
//
// Deprecated: deprecated
type RuleTemplateOptionsStaticHeaders struct {
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
func (r RuleTemplateOptionsStaticHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
type RuleTemplateOptionsStaticRequestHeaders struct {
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
func (r RuleTemplateOptionsStaticRequestHeaders) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
type RuleTemplateOptionsUserAgentACL struct {
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
func (r RuleTemplateOptionsUserAgentACL) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to enable WAAP (Web Application and API Protection).
type RuleTemplateOptionsWaap struct {
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
func (r RuleTemplateOptionsWaap) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
type RuleTemplateOptionsWebsockets struct {
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
func (r RuleTemplateOptionsWebsockets) RawJSON() string { return r.JSON.raw }
func (r *RuleTemplateOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type RuleTemplateOverrideOriginProtocol string

const (
	RuleTemplateOverrideOriginProtocolHTTPS RuleTemplateOverrideOriginProtocol = "HTTPS"
	RuleTemplateOverrideOriginProtocolHTTP  RuleTemplateOverrideOriginProtocol = "HTTP"
	RuleTemplateOverrideOriginProtocolMatch RuleTemplateOverrideOriginProtocol = "MATCH"
)

type RuleTemplateList []RuleTemplate

type RuleTemplateNewParams struct {
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
	// Rule template name.
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
	OverrideOriginProtocol RuleTemplateNewParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options RuleTemplateNewParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r RuleTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type RuleTemplateNewParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods RuleTemplateNewParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection RuleTemplateNewParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression RuleTemplateNewParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings RuleTemplateNewParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders RuleTemplateNewParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors RuleTemplateNewParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL RuleTemplateNewParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache RuleTemplateNewParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges RuleTemplateNewParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings RuleTemplateNewParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge RuleTemplateNewParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed RuleTemplateNewParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect RuleTemplateNewParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn RuleTemplateNewParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader RuleTemplateNewParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn RuleTemplateNewParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader RuleTemplateNewParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie RuleTemplateNewParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString RuleTemplateNewParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack RuleTemplateNewParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL RuleTemplateNewParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth RuleTemplateNewParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey RuleTemplateNewParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet RuleTemplateNewParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout RuleTemplateNewParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout RuleTemplateNewParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist RuleTemplateNewParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist RuleTemplateNewParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding RuleTemplateNewParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL RuleTemplateNewParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter RuleTemplateNewParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite RuleTemplateNewParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey RuleTemplateNewParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice RuleTemplateNewParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni RuleTemplateNewParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale RuleTemplateNewParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders RuleTemplateNewParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders RuleTemplateNewParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders RuleTemplateNewParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL RuleTemplateNewParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap RuleTemplateNewParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets RuleTemplateNewParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r RuleTemplateNewParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsAllowedHTTPMethods struct {
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

func (r RuleTemplateNewParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type RuleTemplateNewParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge RuleTemplateNewParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r RuleTemplateNewParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type RuleTemplateNewParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r RuleTemplateNewParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsBrotliCompression struct {
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

func (r RuleTemplateNewParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsBrowserCacheSettings struct {
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

func (r RuleTemplateNewParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsCacheHTTPHeaders struct {
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

func (r RuleTemplateNewParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsCors struct {
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

func (r RuleTemplateNewParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateNewParamsOptionsCountryACL struct {
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

func (r RuleTemplateNewParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsCountryACL](
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
type RuleTemplateNewParamsOptionsDisableCache struct {
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

func (r RuleTemplateNewParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsDisableProxyForceRanges struct {
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

func (r RuleTemplateNewParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type RuleTemplateNewParamsOptionsEdgeCacheSettings struct {
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

func (r RuleTemplateNewParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type RuleTemplateNewParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody RuleTemplateNewParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody RuleTemplateNewParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r RuleTemplateNewParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type RuleTemplateNewParamsOptionsFastedgeOnRequestBody struct {
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

func (r RuleTemplateNewParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateNewParamsOptionsFastedgeOnResponseBody struct {
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

func (r RuleTemplateNewParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsFetchCompressed struct {
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

func (r RuleTemplateNewParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type RuleTemplateNewParamsOptionsFollowOriginRedirect struct {
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

func (r RuleTemplateNewParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type RuleTemplateNewParamsOptionsForceReturn struct {
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
	TimeInterval RuleTemplateNewParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r RuleTemplateNewParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type RuleTemplateNewParamsOptionsForceReturnTimeInterval struct {
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

func (r RuleTemplateNewParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsForwardHostHeader struct {
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

func (r RuleTemplateNewParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsGzipOn struct {
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

func (r RuleTemplateNewParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsHostHeader struct {
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

func (r RuleTemplateNewParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsIgnoreCookie struct {
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

func (r RuleTemplateNewParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsIgnoreQueryString struct {
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

func (r RuleTemplateNewParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type RuleTemplateNewParamsOptionsImageStack struct {
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

func (r RuleTemplateNewParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsIPAddressACL struct {
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

func (r RuleTemplateNewParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type RuleTemplateNewParamsOptionsLimitBandwidth struct {
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

func (r RuleTemplateNewParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsLimitBandwidth](
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
type RuleTemplateNewParamsOptionsProxyCacheKey struct {
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

func (r RuleTemplateNewParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsProxyCacheMethodsSet struct {
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

func (r RuleTemplateNewParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsProxyConnectTimeout struct {
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

func (r RuleTemplateNewParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsProxyReadTimeout struct {
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

func (r RuleTemplateNewParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsQueryParamsBlacklist struct {
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

func (r RuleTemplateNewParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsQueryParamsWhitelist struct {
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

func (r RuleTemplateNewParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsQueryStringForwarding struct {
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

func (r RuleTemplateNewParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateNewParamsOptionsReferrerACL struct {
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

func (r RuleTemplateNewParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type RuleTemplateNewParamsOptionsRequestLimiter struct {
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

func (r RuleTemplateNewParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type RuleTemplateNewParamsOptionsRewrite struct {
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

func (r RuleTemplateNewParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type RuleTemplateNewParamsOptionsSecureKey struct {
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

func (r RuleTemplateNewParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsSecureKey](
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
type RuleTemplateNewParamsOptionsSlice struct {
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

func (r RuleTemplateNewParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsSni struct {
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

func (r RuleTemplateNewParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsStale struct {
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

func (r RuleTemplateNewParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                     `json:"enabled,required"`
	Value   []RuleTemplateNewParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r RuleTemplateNewParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type RuleTemplateNewParamsOptionsStaticResponseHeadersValue struct {
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

func (r RuleTemplateNewParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOptionsStaticHeaders struct {
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

func (r RuleTemplateNewParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsStaticRequestHeaders struct {
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

func (r RuleTemplateNewParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateNewParamsOptionsUserAgentACL struct {
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

func (r RuleTemplateNewParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateNewParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsWaap struct {
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

func (r RuleTemplateNewParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type RuleTemplateNewParamsOptionsWebsockets struct {
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

func (r RuleTemplateNewParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateNewParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateNewParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type RuleTemplateNewParamsOverrideOriginProtocol string

const (
	RuleTemplateNewParamsOverrideOriginProtocolHTTPS RuleTemplateNewParamsOverrideOriginProtocol = "HTTPS"
	RuleTemplateNewParamsOverrideOriginProtocolHTTP  RuleTemplateNewParamsOverrideOriginProtocol = "HTTP"
	RuleTemplateNewParamsOverrideOriginProtocolMatch RuleTemplateNewParamsOverrideOriginProtocol = "MATCH"
)

type RuleTemplateUpdateParams struct {
	// Rule template name.
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
	OverrideOriginProtocol RuleTemplateUpdateParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options RuleTemplateUpdateParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r RuleTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type RuleTemplateUpdateParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods RuleTemplateUpdateParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection RuleTemplateUpdateParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression RuleTemplateUpdateParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings RuleTemplateUpdateParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders RuleTemplateUpdateParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors RuleTemplateUpdateParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL RuleTemplateUpdateParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache RuleTemplateUpdateParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges RuleTemplateUpdateParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings RuleTemplateUpdateParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge RuleTemplateUpdateParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed RuleTemplateUpdateParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect RuleTemplateUpdateParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn RuleTemplateUpdateParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader RuleTemplateUpdateParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn RuleTemplateUpdateParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader RuleTemplateUpdateParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie RuleTemplateUpdateParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString RuleTemplateUpdateParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack RuleTemplateUpdateParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL RuleTemplateUpdateParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth RuleTemplateUpdateParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey RuleTemplateUpdateParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout RuleTemplateUpdateParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout RuleTemplateUpdateParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist RuleTemplateUpdateParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist RuleTemplateUpdateParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding RuleTemplateUpdateParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL RuleTemplateUpdateParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter RuleTemplateUpdateParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite RuleTemplateUpdateParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey RuleTemplateUpdateParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice RuleTemplateUpdateParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni RuleTemplateUpdateParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale RuleTemplateUpdateParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders RuleTemplateUpdateParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders RuleTemplateUpdateParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders RuleTemplateUpdateParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL RuleTemplateUpdateParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap RuleTemplateUpdateParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets RuleTemplateUpdateParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsAllowedHTTPMethods struct {
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

func (r RuleTemplateUpdateParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type RuleTemplateUpdateParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsBrotliCompression struct {
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

func (r RuleTemplateUpdateParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsBrowserCacheSettings struct {
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

func (r RuleTemplateUpdateParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsCacheHTTPHeaders struct {
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

func (r RuleTemplateUpdateParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsCors struct {
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

func (r RuleTemplateUpdateParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateUpdateParamsOptionsCountryACL struct {
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

func (r RuleTemplateUpdateParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsCountryACL](
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
type RuleTemplateUpdateParamsOptionsDisableCache struct {
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

func (r RuleTemplateUpdateParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsDisableProxyForceRanges struct {
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

func (r RuleTemplateUpdateParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type RuleTemplateUpdateParamsOptionsEdgeCacheSettings struct {
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

func (r RuleTemplateUpdateParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type RuleTemplateUpdateParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody struct {
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

func (r RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody struct {
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

func (r RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsFetchCompressed struct {
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

func (r RuleTemplateUpdateParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type RuleTemplateUpdateParamsOptionsFollowOriginRedirect struct {
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

func (r RuleTemplateUpdateParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type RuleTemplateUpdateParamsOptionsForceReturn struct {
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
	TimeInterval RuleTemplateUpdateParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type RuleTemplateUpdateParamsOptionsForceReturnTimeInterval struct {
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

func (r RuleTemplateUpdateParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsForwardHostHeader struct {
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

func (r RuleTemplateUpdateParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsGzipOn struct {
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

func (r RuleTemplateUpdateParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsHostHeader struct {
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

func (r RuleTemplateUpdateParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsIgnoreCookie struct {
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

func (r RuleTemplateUpdateParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsIgnoreQueryString struct {
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

func (r RuleTemplateUpdateParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type RuleTemplateUpdateParamsOptionsImageStack struct {
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

func (r RuleTemplateUpdateParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsIPAddressACL struct {
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

func (r RuleTemplateUpdateParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type RuleTemplateUpdateParamsOptionsLimitBandwidth struct {
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

func (r RuleTemplateUpdateParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsLimitBandwidth](
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
type RuleTemplateUpdateParamsOptionsProxyCacheKey struct {
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

func (r RuleTemplateUpdateParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet struct {
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

func (r RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsProxyConnectTimeout struct {
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

func (r RuleTemplateUpdateParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsProxyReadTimeout struct {
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

func (r RuleTemplateUpdateParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsQueryParamsBlacklist struct {
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

func (r RuleTemplateUpdateParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsQueryParamsWhitelist struct {
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

func (r RuleTemplateUpdateParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsQueryStringForwarding struct {
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

func (r RuleTemplateUpdateParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateUpdateParamsOptionsReferrerACL struct {
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

func (r RuleTemplateUpdateParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type RuleTemplateUpdateParamsOptionsRequestLimiter struct {
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

func (r RuleTemplateUpdateParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type RuleTemplateUpdateParamsOptionsRewrite struct {
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

func (r RuleTemplateUpdateParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type RuleTemplateUpdateParamsOptionsSecureKey struct {
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

func (r RuleTemplateUpdateParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsSecureKey](
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
type RuleTemplateUpdateParamsOptionsSlice struct {
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

func (r RuleTemplateUpdateParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsSni struct {
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

func (r RuleTemplateUpdateParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsStale struct {
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

func (r RuleTemplateUpdateParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                        `json:"enabled,required"`
	Value   []RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r RuleTemplateUpdateParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue struct {
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

func (r RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOptionsStaticHeaders struct {
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

func (r RuleTemplateUpdateParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsStaticRequestHeaders struct {
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

func (r RuleTemplateUpdateParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateUpdateParamsOptionsUserAgentACL struct {
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

func (r RuleTemplateUpdateParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateUpdateParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsWaap struct {
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

func (r RuleTemplateUpdateParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type RuleTemplateUpdateParamsOptionsWebsockets struct {
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

func (r RuleTemplateUpdateParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateUpdateParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateUpdateParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type RuleTemplateUpdateParamsOverrideOriginProtocol string

const (
	RuleTemplateUpdateParamsOverrideOriginProtocolHTTPS RuleTemplateUpdateParamsOverrideOriginProtocol = "HTTPS"
	RuleTemplateUpdateParamsOverrideOriginProtocolHTTP  RuleTemplateUpdateParamsOverrideOriginProtocol = "HTTP"
	RuleTemplateUpdateParamsOverrideOriginProtocolMatch RuleTemplateUpdateParamsOverrideOriginProtocol = "MATCH"
)

type RuleTemplateReplaceParams struct {
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
	// Rule template name.
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
	OverrideOriginProtocol RuleTemplateReplaceParamsOverrideOriginProtocol `json:"overrideOriginProtocol,omitzero"`
	// List of options that can be configured for the rule.
	//
	// In case of `null` value the option is not added to the rule. Option inherits its
	// value from the CDN resource settings.
	Options RuleTemplateReplaceParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r RuleTemplateReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of options that can be configured for the rule.
//
// In case of `null` value the option is not added to the rule. Option inherits its
// value from the CDN resource settings.
type RuleTemplateReplaceParamsOptions struct {
	// HTTP methods allowed for content requests from the CDN.
	AllowedHTTPMethods RuleTemplateReplaceParamsOptionsAllowedHTTPMethods `json:"allowedHttpMethods,omitzero"`
	// Allows to prevent online services from overloading and ensure your business
	// workflow running smoothly.
	BotProtection RuleTemplateReplaceParamsOptionsBotProtection `json:"bot_protection,omitzero"`
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
	BrotliCompression RuleTemplateReplaceParamsOptionsBrotliCompression `json:"brotli_compression,omitzero"`
	// Cache expiration time for users browsers in seconds.
	//
	// Cache expiration time is applied to the following response codes: 200, 201, 204,
	// 206, 301, 302, 303, 304, 307, 308.
	//
	// Responses with other codes will not be cached.
	BrowserCacheSettings RuleTemplateReplaceParamsOptionsBrowserCacheSettings `json:"browser_cache_settings,omitzero"`
	// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
	//
	// HTTP Headers that must be included in the response.
	//
	// Deprecated: deprecated
	CacheHTTPHeaders RuleTemplateReplaceParamsOptionsCacheHTTPHeaders `json:"cache_http_headers,omitzero"`
	// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
	//
	// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
	// to a response to a browser.
	Cors RuleTemplateReplaceParamsOptionsCors `json:"cors,omitzero"`
	// Enables control access to content for specified countries.
	CountryACL RuleTemplateReplaceParamsOptionsCountryACL `json:"country_acl,omitzero"`
	// **Legacy option**. Use the `edge_cache_settings` option instead.
	//
	// Allows the complete disabling of content caching.
	//
	// Deprecated: deprecated
	DisableCache RuleTemplateReplaceParamsOptionsDisableCache `json:"disable_cache,omitzero"`
	// Allows 206 responses regardless of the settings of an origin source.
	DisableProxyForceRanges RuleTemplateReplaceParamsOptionsDisableProxyForceRanges `json:"disable_proxy_force_ranges,omitzero"`
	// Cache expiration time for CDN servers.
	//
	// `value` and `default` fields cannot be used simultaneously.
	EdgeCacheSettings RuleTemplateReplaceParamsOptionsEdgeCacheSettings `json:"edge_cache_settings,omitzero"`
	// Allows to configure FastEdge app to be called on different request/response
	// phases.
	//
	// Note: At least one of `on_request_headers`, `on_request_body`,
	// `on_response_headers`, or `on_response_body` must be specified.
	Fastedge RuleTemplateReplaceParamsOptionsFastedge `json:"fastedge,omitzero"`
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
	FetchCompressed RuleTemplateReplaceParamsOptionsFetchCompressed `json:"fetch_compressed,omitzero"`
	// Enables redirection from origin. If the origin server returns a redirect, the
	// option allows the CDN to pull the requested content from the origin server that
	// was returned in the redirect.
	FollowOriginRedirect RuleTemplateReplaceParamsOptionsFollowOriginRedirect `json:"follow_origin_redirect,omitzero"`
	// Applies custom HTTP response codes for CDN content.
	//
	// The following codes are reserved by our system and cannot be specified in this
	// option: 408, 444, 477, 494, 495, 496, 497, 499.
	ForceReturn RuleTemplateReplaceParamsOptionsForceReturn `json:"force_return,omitzero"`
	// Forwards the Host header from a end-user request to an origin server.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	ForwardHostHeader RuleTemplateReplaceParamsOptionsForwardHostHeader `json:"forward_host_header,omitzero"`
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
	GzipOn RuleTemplateReplaceParamsOptionsGzipOn `json:"gzipOn,omitzero"`
	// Sets the Host header that CDN servers use when request content from an origin
	// server. Your server must be able to process requests with the chosen header.
	//
	// If the option is `null`, the Host Header value is equal to first CNAME.
	//
	// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
	HostHeader RuleTemplateReplaceParamsOptionsHostHeader `json:"hostHeader,omitzero"`
	// Defines whether the files with the Set-Cookies header are cached as one file or
	// as different ones.
	IgnoreCookie RuleTemplateReplaceParamsOptionsIgnoreCookie `json:"ignore_cookie,omitzero"`
	// How a file with different query strings is cached: either as one object (option
	// is enabled) or as different objects (option is disabled.)
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	IgnoreQueryString RuleTemplateReplaceParamsOptionsIgnoreQueryString `json:"ignoreQueryString,omitzero"`
	// Transforms JPG and PNG images (for example, resize or crop) and automatically
	// converts them to WebP or AVIF format.
	ImageStack RuleTemplateReplaceParamsOptionsImageStack `json:"image_stack,omitzero"`
	// Controls access to the CDN resource content for specific IP addresses.
	//
	// If you want to use IPs from our CDN servers IP list for IP ACL configuration,
	// you have to independently monitor their relevance.
	//
	// We recommend you use a script for automatically update IP ACL.
	// [Read more.](/docs/api-reference/cdn/ip-addresses-list/get-cdn-servers-ip-addresses)
	IPAddressACL RuleTemplateReplaceParamsOptionsIPAddressACL `json:"ip_address_acl,omitzero"`
	// Allows to control the download speed per connection.
	LimitBandwidth RuleTemplateReplaceParamsOptionsLimitBandwidth `json:"limit_bandwidth,omitzero"`
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
	ProxyCacheKey RuleTemplateReplaceParamsOptionsProxyCacheKey `json:"proxy_cache_key,omitzero"`
	// Caching for POST requests along with default GET and HEAD.
	ProxyCacheMethodsSet RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet `json:"proxy_cache_methods_set,omitzero"`
	// The time limit for establishing a connection with the origin.
	ProxyConnectTimeout RuleTemplateReplaceParamsOptionsProxyConnectTimeout `json:"proxy_connect_timeout,omitzero"`
	// The time limit for receiving a partial response from the origin. If no response
	// is received within this time, the connection will be closed.
	//
	// **Note:** When used with a WebSocket connection, this option supports values
	// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
	ProxyReadTimeout RuleTemplateReplaceParamsOptionsProxyReadTimeout `json:"proxy_read_timeout,omitzero"`
	// Files with the specified query parameters are cached as one object, files with
	// other parameters are cached as different objects.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsBlacklist RuleTemplateReplaceParamsOptionsQueryParamsBlacklist `json:"query_params_blacklist,omitzero"`
	// Files with the specified query parameters are cached as different objects, files
	// with other parameters are cached as one object.
	//
	// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
	// options cannot be enabled simultaneously.
	QueryParamsWhitelist RuleTemplateReplaceParamsOptionsQueryParamsWhitelist `json:"query_params_whitelist,omitzero"`
	// The Query String Forwarding feature allows for the seamless transfer of
	// parameters embedded in playlist files to the corresponding media chunk files.
	// This functionality ensures that specific attributes, such as authentication
	// tokens or tracking information, are consistently passed along from the playlist
	// manifest to the individual media segments. This is particularly useful for
	// maintaining continuity in security, analytics, and any other parameter-based
	// operations across the entire media delivery workflow.
	QueryStringForwarding RuleTemplateReplaceParamsOptionsQueryStringForwarding `json:"query_string_forwarding,omitzero"`
	// Enables redirect from HTTP to HTTPS.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPToHTTPS RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS `json:"redirect_http_to_https,omitzero"`
	// Enables redirect from HTTPS to HTTP.
	//
	// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
	// simultaneously.
	RedirectHTTPSToHTTP RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP `json:"redirect_https_to_http,omitzero"`
	// Controls access to the CDN resource content for specified domain names.
	ReferrerACL RuleTemplateReplaceParamsOptionsReferrerACL `json:"referrer_acl,omitzero"`
	// Option allows to limit the amount of HTTP requests.
	RequestLimiter RuleTemplateReplaceParamsOptionsRequestLimiter `json:"request_limiter,omitzero"`
	// Hides HTTP headers from an origin server in the CDN response.
	ResponseHeadersHidingPolicy RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy `json:"response_headers_hiding_policy,omitzero"`
	// Changes and redirects requests from the CDN to the origin. It operates according
	// to the
	// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
	// configuration.
	Rewrite RuleTemplateReplaceParamsOptionsRewrite `json:"rewrite,omitzero"`
	// Configures access with tokenized URLs. This makes impossible to access content
	// without a valid (unexpired) token.
	SecureKey RuleTemplateReplaceParamsOptionsSecureKey `json:"secure_key,omitzero"`
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
	Slice RuleTemplateReplaceParamsOptionsSlice `json:"slice,omitzero"`
	// The hostname that is added to SNI requests from CDN servers to the origin server
	// via HTTPS.
	//
	// SNI is generally only required if your origin uses shared hosting or does not
	// have a dedicated IP address. If the origin server presents multiple
	// certificates, SNI allows the origin server to know which certificate to use for
	// the connection.
	//
	// The option works only if `originProtocol` parameter is `HTTPS` or `MATCH`.
	Sni RuleTemplateReplaceParamsOptionsSni `json:"sni,omitzero"`
	// Serves stale cached content in case of origin unavailability.
	Stale RuleTemplateReplaceParamsOptionsStale `json:"stale,omitzero"`
	// Custom HTTP Headers that a CDN server adds to a response.
	StaticResponseHeaders RuleTemplateReplaceParamsOptionsStaticResponseHeaders `json:"static_response_headers,omitzero"`
	// **Legacy option**. Use the `static_response_headers` option instead.
	//
	// Custom HTTP Headers that a CDN server adds to response. Up to fifty custom HTTP
	// Headers can be specified. May contain a header with multiple values.
	//
	// Deprecated: deprecated
	StaticHeaders RuleTemplateReplaceParamsOptionsStaticHeaders `json:"staticHeaders,omitzero"`
	// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
	// Headers can be specified.
	StaticRequestHeaders RuleTemplateReplaceParamsOptionsStaticRequestHeaders `json:"staticRequestHeaders,omitzero"`
	// Controls access to the content for specified User-Agents.
	UserAgentACL RuleTemplateReplaceParamsOptionsUserAgentACL `json:"user_agent_acl,omitzero"`
	// Allows to enable WAAP (Web Application and API Protection).
	Waap RuleTemplateReplaceParamsOptionsWaap `json:"waap,omitzero"`
	// Enables or disables WebSockets connections to an origin server.
	Websockets RuleTemplateReplaceParamsOptionsWebsockets `json:"websockets,omitzero"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP methods allowed for content requests from the CDN.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsAllowedHTTPMethods struct {
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

func (r RuleTemplateReplaceParamsOptionsAllowedHTTPMethods) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsAllowedHTTPMethods
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsAllowedHTTPMethods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to prevent online services from overloading and ensure your business
// workflow running smoothly.
//
// The properties BotChallenge, Enabled are required.
type RuleTemplateReplaceParamsOptionsBotProtection struct {
	// Controls the bot challenge module state.
	BotChallenge RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge `json:"bot_challenge,omitzero,required"`
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptionsBotProtection) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsBotProtection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsBotProtection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the bot challenge module state.
type RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge struct {
	// Possible values:
	//
	// - **true** - Bot challenge is enabled.
	// - **false** - Bot challenge is disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsBrotliCompression struct {
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

func (r RuleTemplateReplaceParamsOptionsBrotliCompression) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsBrotliCompression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsBrotliCompression) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsBrowserCacheSettings struct {
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

func (r RuleTemplateReplaceParamsOptionsBrowserCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsBrowserCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsBrowserCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Legacy option**. Use the `response_headers_hiding_policy` option instead.
//
// HTTP Headers that must be included in the response.
//
// Deprecated: deprecated
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsCacheHTTPHeaders struct {
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

func (r RuleTemplateReplaceParamsOptionsCacheHTTPHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsCacheHTTPHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsCacheHTTPHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables CORS (Cross-Origin Resource Sharing) header support.
//
// CORS header support allows the CDN to add the Access-Control-Allow-Origin header
// to a response to a browser.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsCors struct {
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

func (r RuleTemplateReplaceParamsOptionsCors) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsCors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables control access to content for specified countries.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateReplaceParamsOptionsCountryACL struct {
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

func (r RuleTemplateReplaceParamsOptionsCountryACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsCountryACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsCountryACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsCountryACL](
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
type RuleTemplateReplaceParamsOptionsDisableCache struct {
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

func (r RuleTemplateReplaceParamsOptionsDisableCache) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsDisableCache
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsDisableCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows 206 responses regardless of the settings of an origin source.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsDisableProxyForceRanges struct {
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

func (r RuleTemplateReplaceParamsOptionsDisableProxyForceRanges) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsDisableProxyForceRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsDisableProxyForceRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cache expiration time for CDN servers.
//
// `value` and `default` fields cannot be used simultaneously.
//
// The property Enabled is required.
type RuleTemplateReplaceParamsOptionsEdgeCacheSettings struct {
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

func (r RuleTemplateReplaceParamsOptionsEdgeCacheSettings) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsEdgeCacheSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsEdgeCacheSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge app to be called on different request/response
// phases.
//
// Note: At least one of `on_request_headers`, `on_request_body`,
// `on_response_headers`, or `on_response_body` must be specified.
//
// The property Enabled is required.
type RuleTemplateReplaceParamsOptionsFastedge struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool `json:"enabled,required"`
	// Allows to configure FastEdge application that will be called to handle request
	// body as soon as CDN receives incoming HTTP request.
	OnRequestBody RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody `json:"on_request_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle request
	// headers as soon as CDN receives incoming HTTP request, **before cache**.
	OnRequestHeaders RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders `json:"on_request_headers,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// body before CDN sends the HTTP response.
	OnResponseBody RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody `json:"on_response_body,omitzero"`
	// Allows to configure FastEdge application that will be called to handle response
	// headers before CDN sends the HTTP response.
	OnResponseHeaders RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders `json:"on_response_headers,omitzero"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptionsFastedge) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFastedge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFastedge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// body as soon as CDN receives incoming HTTP request.
//
// The property AppID is required.
type RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody struct {
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

func (r RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle request
// headers as soon as CDN receives incoming HTTP request, **before cache**.
//
// The property AppID is required.
type RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders struct {
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

func (r RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// body before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody struct {
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

func (r RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to configure FastEdge application that will be called to handle response
// headers before CDN sends the HTTP response.
//
// The property AppID is required.
type RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders struct {
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

func (r RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsFetchCompressed struct {
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

func (r RuleTemplateReplaceParamsOptionsFetchCompressed) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFetchCompressed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFetchCompressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirection from origin. If the origin server returns a redirect, the
// option allows the CDN to pull the requested content from the origin server that
// was returned in the redirect.
//
// The properties Codes, Enabled are required.
type RuleTemplateReplaceParamsOptionsFollowOriginRedirect struct {
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

func (r RuleTemplateReplaceParamsOptionsFollowOriginRedirect) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsFollowOriginRedirect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsFollowOriginRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Applies custom HTTP response codes for CDN content.
//
// The following codes are reserved by our system and cannot be specified in this
// option: 408, 444, 477, 494, 495, 496, 497, 499.
//
// The properties Body, Code, Enabled are required.
type RuleTemplateReplaceParamsOptionsForceReturn struct {
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
	TimeInterval RuleTemplateReplaceParamsOptionsForceReturnTimeInterval `json:"time_interval,omitzero"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptionsForceReturn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsForceReturn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsForceReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the time at which a custom HTTP response code should be applied. By
// default, a custom HTTP response code is applied at any time.
//
// The properties EndTime, StartTime are required.
type RuleTemplateReplaceParamsOptionsForceReturnTimeInterval struct {
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

func (r RuleTemplateReplaceParamsOptionsForceReturnTimeInterval) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsForceReturnTimeInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsForceReturnTimeInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forwards the Host header from a end-user request to an origin server.
//
// `hostHeader` and `forward_host_header` options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsForwardHostHeader struct {
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

func (r RuleTemplateReplaceParamsOptionsForwardHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsForwardHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsForwardHostHeader) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsGzipOn struct {
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

func (r RuleTemplateReplaceParamsOptionsGzipOn) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsGzipOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsGzipOn) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsHostHeader struct {
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

func (r RuleTemplateReplaceParamsOptionsHostHeader) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsHostHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsHostHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether the files with the Set-Cookies header are cached as one file or
// as different ones.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsIgnoreCookie struct {
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

func (r RuleTemplateReplaceParamsOptionsIgnoreCookie) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsIgnoreCookie
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsIgnoreCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a file with different query strings is cached: either as one object (option
// is enabled) or as different objects (option is disabled.)
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsIgnoreQueryString struct {
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

func (r RuleTemplateReplaceParamsOptionsIgnoreQueryString) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsIgnoreQueryString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsIgnoreQueryString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transforms JPG and PNG images (for example, resize or crop) and automatically
// converts them to WebP or AVIF format.
//
// The property Enabled is required.
type RuleTemplateReplaceParamsOptionsImageStack struct {
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

func (r RuleTemplateReplaceParamsOptionsImageStack) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsImageStack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsImageStack) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsIPAddressACL struct {
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

func (r RuleTemplateReplaceParamsOptionsIPAddressACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsIPAddressACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsIPAddressACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsIPAddressACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to control the download speed per connection.
//
// The properties Enabled, LimitType are required.
type RuleTemplateReplaceParamsOptionsLimitBandwidth struct {
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

func (r RuleTemplateReplaceParamsOptionsLimitBandwidth) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsLimitBandwidth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsLimitBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsLimitBandwidth](
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
type RuleTemplateReplaceParamsOptionsProxyCacheKey struct {
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

func (r RuleTemplateReplaceParamsOptionsProxyCacheKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsProxyCacheKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsProxyCacheKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caching for POST requests along with default GET and HEAD.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet struct {
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

func (r RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for establishing a connection with the origin.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsProxyConnectTimeout struct {
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

func (r RuleTemplateReplaceParamsOptionsProxyConnectTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsProxyConnectTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsProxyConnectTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The time limit for receiving a partial response from the origin. If no response
// is received within this time, the connection will be closed.
//
// **Note:** When used with a WebSocket connection, this option supports values
// only in the range 1–20 seconds (instead of the usual 1–30 seconds).
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsProxyReadTimeout struct {
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

func (r RuleTemplateReplaceParamsOptionsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsProxyReadTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsProxyReadTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as one object, files with
// other parameters are cached as different objects.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsQueryParamsBlacklist struct {
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

func (r RuleTemplateReplaceParamsOptionsQueryParamsBlacklist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsQueryParamsBlacklist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsQueryParamsBlacklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files with the specified query parameters are cached as different objects, files
// with other parameters are cached as one object.
//
// `ignoreQueryString`, `query_params_whitelist` and `query_params_blacklist`
// options cannot be enabled simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsQueryParamsWhitelist struct {
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

func (r RuleTemplateReplaceParamsOptionsQueryParamsWhitelist) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsQueryParamsWhitelist
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsQueryParamsWhitelist) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsQueryStringForwarding struct {
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

func (r RuleTemplateReplaceParamsOptionsQueryStringForwarding) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsQueryStringForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsQueryStringForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTP to HTTPS.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS struct {
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

func (r RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables redirect from HTTPS to HTTP.
//
// `redirect_http_to_https` and `redirect_https_to_http` options cannot be enabled
// simultaneously.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP struct {
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

func (r RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the CDN resource content for specified domain names.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateReplaceParamsOptionsReferrerACL struct {
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

func (r RuleTemplateReplaceParamsOptionsReferrerACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsReferrerACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsReferrerACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsReferrerACL](
		"policy_type", "allow", "deny",
	)
}

// Option allows to limit the amount of HTTP requests.
//
// The properties Enabled, Rate are required.
type RuleTemplateReplaceParamsOptionsRequestLimiter struct {
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

func (r RuleTemplateReplaceParamsOptionsRequestLimiter) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsRequestLimiter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsRequestLimiter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsRequestLimiter](
		"rate_unit", "r/s", "r/m",
	)
}

// Hides HTTP headers from an origin server in the CDN response.
//
// The properties Enabled, Excepted, Mode are required.
type RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy struct {
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

func (r RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy](
		"mode", "hide", "show",
	)
}

// Changes and redirects requests from the CDN to the origin. It operates according
// to the
// [Nginx](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite)
// configuration.
//
// The properties Body, Enabled are required.
type RuleTemplateReplaceParamsOptionsRewrite struct {
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

func (r RuleTemplateReplaceParamsOptionsRewrite) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsRewrite
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsRewrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsRewrite](
		"flag", "break", "last", "redirect", "permanent",
	)
}

// Configures access with tokenized URLs. This makes impossible to access content
// without a valid (unexpired) token.
//
// The properties Enabled, Key are required.
type RuleTemplateReplaceParamsOptionsSecureKey struct {
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

func (r RuleTemplateReplaceParamsOptionsSecureKey) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsSecureKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsSecureKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsSecureKey](
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
type RuleTemplateReplaceParamsOptionsSlice struct {
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

func (r RuleTemplateReplaceParamsOptionsSlice) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsSlice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsSlice) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsSni struct {
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

func (r RuleTemplateReplaceParamsOptionsSni) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsSni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsSni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsSni](
		"sni_type", "dynamic", "custom",
	)
}

// Serves stale cached content in case of origin unavailability.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsStale struct {
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

func (r RuleTemplateReplaceParamsOptionsStale) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsStale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsStale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers that a CDN server adds to a response.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsStaticResponseHeaders struct {
	// Controls the option state.
	//
	// Possible values:
	//
	// - **true** - Option is enabled.
	// - **false** - Option is disabled.
	Enabled bool                                                         `json:"enabled,required"`
	Value   []RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue `json:"value,omitzero,required"`
	paramObj
}

func (r RuleTemplateReplaceParamsOptionsStaticResponseHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsStaticResponseHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsStaticResponseHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue struct {
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

func (r RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOptionsStaticHeaders struct {
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

func (r RuleTemplateReplaceParamsOptionsStaticHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsStaticHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsStaticHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom HTTP Headers for a CDN server to add to request. Up to fifty custom HTTP
// Headers can be specified.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsStaticRequestHeaders struct {
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

func (r RuleTemplateReplaceParamsOptionsStaticRequestHeaders) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsStaticRequestHeaders
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsStaticRequestHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls access to the content for specified User-Agents.
//
// The properties Enabled, ExceptedValues, PolicyType are required.
type RuleTemplateReplaceParamsOptionsUserAgentACL struct {
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

func (r RuleTemplateReplaceParamsOptionsUserAgentACL) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsUserAgentACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsUserAgentACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RuleTemplateReplaceParamsOptionsUserAgentACL](
		"policy_type", "allow", "deny",
	)
}

// Allows to enable WAAP (Web Application and API Protection).
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsWaap struct {
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

func (r RuleTemplateReplaceParamsOptionsWaap) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsWaap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsWaap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables WebSockets connections to an origin server.
//
// The properties Enabled, Value are required.
type RuleTemplateReplaceParamsOptionsWebsockets struct {
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

func (r RuleTemplateReplaceParamsOptionsWebsockets) MarshalJSON() (data []byte, err error) {
	type shadow RuleTemplateReplaceParamsOptionsWebsockets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleTemplateReplaceParamsOptionsWebsockets) UnmarshalJSON(data []byte) error {
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
type RuleTemplateReplaceParamsOverrideOriginProtocol string

const (
	RuleTemplateReplaceParamsOverrideOriginProtocolHTTPS RuleTemplateReplaceParamsOverrideOriginProtocol = "HTTPS"
	RuleTemplateReplaceParamsOverrideOriginProtocolHTTP  RuleTemplateReplaceParamsOverrideOriginProtocol = "HTTP"
	RuleTemplateReplaceParamsOverrideOriginProtocolMatch RuleTemplateReplaceParamsOverrideOriginProtocol = "MATCH"
)
