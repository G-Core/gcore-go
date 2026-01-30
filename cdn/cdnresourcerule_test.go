// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func TestCDNResourceRuleNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CDN.CDNResources.Rules.New(
		context.TODO(),
		0,
		cdn.CDNResourceRuleNewParams{
			Name:     "My first rule",
			Rule:     "/folder/images/*.png",
			RuleType: 0,
			Active:   gcore.Bool(true),
			Options: cdn.CDNResourceRuleNewParamsOptions{
				AllowedHTTPMethods: cdn.CDNResourceRuleNewParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.CDNResourceRuleNewParamsOptionsBotProtection{
					BotChallenge: cdn.CDNResourceRuleNewParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.CDNResourceRuleNewParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.CDNResourceRuleNewParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.CDNResourceRuleNewParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.CDNResourceRuleNewParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.CDNResourceRuleNewParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.CDNResourceRuleNewParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.CDNResourceRuleNewParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.CDNResourceRuleNewParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.CDNResourceRuleNewParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.CDNResourceRuleNewParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.CDNResourceRuleNewParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.CDNResourceRuleNewParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.CDNResourceRuleNewParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.CDNResourceRuleNewParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.CDNResourceRuleNewParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.CDNResourceRuleNewParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.CDNResourceRuleNewParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.CDNResourceRuleNewParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.CDNResourceRuleNewParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.CDNResourceRuleNewParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.CDNResourceRuleNewParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.CDNResourceRuleNewParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.CDNResourceRuleNewParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.CDNResourceRuleNewParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.CDNResourceRuleNewParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.CDNResourceRuleNewParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.CDNResourceRuleNewParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.CDNResourceRuleNewParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.CDNResourceRuleNewParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.CDNResourceRuleNewParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.CDNResourceRuleNewParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.CDNResourceRuleNewParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.CDNResourceRuleNewParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.CDNResourceRuleNewParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.CDNResourceRuleNewParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.CDNResourceRuleNewParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.CDNResourceRuleNewParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.CDNResourceRuleNewParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.CDNResourceRuleNewParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.CDNResourceRuleNewParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.CDNResourceRuleNewParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.CDNResourceRuleNewParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.CDNResourceRuleNewParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.CDNResourceRuleNewParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.CDNResourceRuleNewParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.CDNResourceRuleNewParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.CDNResourceRuleNewParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.CDNResourceRuleNewParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.CDNResourceRuleNewParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.CDNResourceRuleNewParamsOverrideOriginProtocolHTTPS,
			Weight:                 gcore.Int(1),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceRuleUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CDN.CDNResources.Rules.Update(
		context.TODO(),
		0,
		cdn.CDNResourceRuleUpdateParams{
			ResourceID: 0,
			Active:     gcore.Bool(true),
			Name:       gcore.String("My first rule"),
			Options: cdn.CDNResourceRuleUpdateParamsOptions{
				AllowedHTTPMethods: cdn.CDNResourceRuleUpdateParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.CDNResourceRuleUpdateParamsOptionsBotProtection{
					BotChallenge: cdn.CDNResourceRuleUpdateParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.CDNResourceRuleUpdateParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.CDNResourceRuleUpdateParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.CDNResourceRuleUpdateParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.CDNResourceRuleUpdateParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.CDNResourceRuleUpdateParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.CDNResourceRuleUpdateParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.CDNResourceRuleUpdateParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.CDNResourceRuleUpdateParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.CDNResourceRuleUpdateParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.CDNResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.CDNResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.CDNResourceRuleUpdateParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.CDNResourceRuleUpdateParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.CDNResourceRuleUpdateParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.CDNResourceRuleUpdateParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.CDNResourceRuleUpdateParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.CDNResourceRuleUpdateParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.CDNResourceRuleUpdateParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.CDNResourceRuleUpdateParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.CDNResourceRuleUpdateParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.CDNResourceRuleUpdateParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.CDNResourceRuleUpdateParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.CDNResourceRuleUpdateParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.CDNResourceRuleUpdateParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.CDNResourceRuleUpdateParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.CDNResourceRuleUpdateParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.CDNResourceRuleUpdateParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.CDNResourceRuleUpdateParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.CDNResourceRuleUpdateParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.CDNResourceRuleUpdateParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.CDNResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.CDNResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.CDNResourceRuleUpdateParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.CDNResourceRuleUpdateParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.CDNResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.CDNResourceRuleUpdateParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.CDNResourceRuleUpdateParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.CDNResourceRuleUpdateParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.CDNResourceRuleUpdateParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.CDNResourceRuleUpdateParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.CDNResourceRuleUpdateParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.CDNResourceRuleUpdateParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.CDNResourceRuleUpdateParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.CDNResourceRuleUpdateParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.CDNResourceRuleUpdateParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.CDNResourceRuleUpdateParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.CDNResourceRuleUpdateParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.CDNResourceRuleUpdateParamsOverrideOriginProtocolHTTPS,
			Rule:                   gcore.String("/folder/images/*.png"),
			RuleType:               gcore.Int(0),
			Weight:                 gcore.Int(1),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceRuleList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CDN.CDNResources.Rules.List(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceRuleDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.CDN.CDNResources.Rules.Delete(
		context.TODO(),
		0,
		cdn.CDNResourceRuleDeleteParams{
			ResourceID: 0,
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceRuleGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CDN.CDNResources.Rules.Get(
		context.TODO(),
		0,
		cdn.CDNResourceRuleGetParams{
			ResourceID: 0,
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceRuleReplaceWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CDN.CDNResources.Rules.Replace(
		context.TODO(),
		0,
		cdn.CDNResourceRuleReplaceParams{
			ResourceID: 0,
			Rule:       "/folder/images/*.png",
			RuleType:   0,
			Active:     gcore.Bool(true),
			Name:       gcore.String("My first rule"),
			Options: cdn.CDNResourceRuleReplaceParamsOptions{
				AllowedHTTPMethods: cdn.CDNResourceRuleReplaceParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.CDNResourceRuleReplaceParamsOptionsBotProtection{
					BotChallenge: cdn.CDNResourceRuleReplaceParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.CDNResourceRuleReplaceParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.CDNResourceRuleReplaceParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.CDNResourceRuleReplaceParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.CDNResourceRuleReplaceParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.CDNResourceRuleReplaceParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.CDNResourceRuleReplaceParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.CDNResourceRuleReplaceParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.CDNResourceRuleReplaceParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.CDNResourceRuleReplaceParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.CDNResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.CDNResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.CDNResourceRuleReplaceParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.CDNResourceRuleReplaceParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.CDNResourceRuleReplaceParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.CDNResourceRuleReplaceParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.CDNResourceRuleReplaceParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.CDNResourceRuleReplaceParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.CDNResourceRuleReplaceParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.CDNResourceRuleReplaceParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.CDNResourceRuleReplaceParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.CDNResourceRuleReplaceParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.CDNResourceRuleReplaceParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.CDNResourceRuleReplaceParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.CDNResourceRuleReplaceParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.CDNResourceRuleReplaceParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.CDNResourceRuleReplaceParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.CDNResourceRuleReplaceParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.CDNResourceRuleReplaceParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.CDNResourceRuleReplaceParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.CDNResourceRuleReplaceParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.CDNResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.CDNResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.CDNResourceRuleReplaceParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.CDNResourceRuleReplaceParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.CDNResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.CDNResourceRuleReplaceParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.CDNResourceRuleReplaceParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.CDNResourceRuleReplaceParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.CDNResourceRuleReplaceParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.CDNResourceRuleReplaceParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.CDNResourceRuleReplaceParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.CDNResourceRuleReplaceParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.CDNResourceRuleReplaceParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.CDNResourceRuleReplaceParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.CDNResourceRuleReplaceParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.CDNResourceRuleReplaceParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.CDNResourceRuleReplaceParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.CDNResourceRuleReplaceParamsOverrideOriginProtocolHTTPS,
			Weight:                 gcore.Int(1),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
