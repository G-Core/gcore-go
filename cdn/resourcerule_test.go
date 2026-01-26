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

func TestResourceRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.Rules.New(
		context.TODO(),
		0,
		cdn.ResourceRuleNewParams{
			Name:     "My first rule",
			Rule:     "/folder/images/*.png",
			RuleType: 0,
			Active:   gcore.Bool(true),
			Options: cdn.ResourceRuleNewParamsOptions{
				AllowedHTTPMethods: cdn.ResourceRuleNewParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.ResourceRuleNewParamsOptionsBotProtection{
					BotChallenge: cdn.ResourceRuleNewParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.ResourceRuleNewParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.ResourceRuleNewParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.ResourceRuleNewParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.ResourceRuleNewParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.ResourceRuleNewParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.ResourceRuleNewParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.ResourceRuleNewParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.ResourceRuleNewParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.ResourceRuleNewParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.ResourceRuleNewParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.ResourceRuleNewParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.ResourceRuleNewParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.ResourceRuleNewParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.ResourceRuleNewParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.ResourceRuleNewParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.ResourceRuleNewParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.ResourceRuleNewParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.ResourceRuleNewParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.ResourceRuleNewParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.ResourceRuleNewParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.ResourceRuleNewParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.ResourceRuleNewParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.ResourceRuleNewParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.ResourceRuleNewParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.ResourceRuleNewParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.ResourceRuleNewParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.ResourceRuleNewParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.ResourceRuleNewParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.ResourceRuleNewParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.ResourceRuleNewParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.ResourceRuleNewParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.ResourceRuleNewParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.ResourceRuleNewParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.ResourceRuleNewParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.ResourceRuleNewParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.ResourceRuleNewParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.ResourceRuleNewParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.ResourceRuleNewParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.ResourceRuleNewParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.ResourceRuleNewParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.ResourceRuleNewParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.ResourceRuleNewParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.ResourceRuleNewParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.ResourceRuleNewParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.ResourceRuleNewParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.ResourceRuleNewParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.ResourceRuleNewParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.ResourceRuleNewParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.ResourceRuleNewParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.ResourceRuleNewParamsOverrideOriginProtocolHTTPS,
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

func TestResourceRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.Rules.Update(
		context.TODO(),
		0,
		cdn.ResourceRuleUpdateParams{
			ResourceID: 0,
			Active:     gcore.Bool(true),
			Name:       gcore.String("My first rule"),
			Options: cdn.ResourceRuleUpdateParamsOptions{
				AllowedHTTPMethods: cdn.ResourceRuleUpdateParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.ResourceRuleUpdateParamsOptionsBotProtection{
					BotChallenge: cdn.ResourceRuleUpdateParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.ResourceRuleUpdateParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.ResourceRuleUpdateParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.ResourceRuleUpdateParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.ResourceRuleUpdateParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.ResourceRuleUpdateParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.ResourceRuleUpdateParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.ResourceRuleUpdateParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.ResourceRuleUpdateParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.ResourceRuleUpdateParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.ResourceRuleUpdateParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.ResourceRuleUpdateParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.ResourceRuleUpdateParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.ResourceRuleUpdateParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.ResourceRuleUpdateParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.ResourceRuleUpdateParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.ResourceRuleUpdateParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.ResourceRuleUpdateParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.ResourceRuleUpdateParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.ResourceRuleUpdateParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.ResourceRuleUpdateParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.ResourceRuleUpdateParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.ResourceRuleUpdateParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.ResourceRuleUpdateParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.ResourceRuleUpdateParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.ResourceRuleUpdateParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.ResourceRuleUpdateParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.ResourceRuleUpdateParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.ResourceRuleUpdateParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.ResourceRuleUpdateParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.ResourceRuleUpdateParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.ResourceRuleUpdateParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.ResourceRuleUpdateParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.ResourceRuleUpdateParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.ResourceRuleUpdateParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.ResourceRuleUpdateParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.ResourceRuleUpdateParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.ResourceRuleUpdateParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.ResourceRuleUpdateParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.ResourceRuleUpdateParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.ResourceRuleUpdateParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.ResourceRuleUpdateParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.ResourceRuleUpdateParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.ResourceRuleUpdateParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.ResourceRuleUpdateParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.ResourceRuleUpdateParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.ResourceRuleUpdateParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.ResourceRuleUpdateParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.ResourceRuleUpdateParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.ResourceRuleUpdateParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.ResourceRuleUpdateParamsOverrideOriginProtocolHTTPS,
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

func TestResourceRuleList(t *testing.T) {
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
	_, err := client.Cdn.Resources.Rules.List(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceRuleDelete(t *testing.T) {
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
	err := client.Cdn.Resources.Rules.Delete(
		context.TODO(),
		0,
		cdn.ResourceRuleDeleteParams{
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

func TestResourceRuleGet(t *testing.T) {
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
	_, err := client.Cdn.Resources.Rules.Get(
		context.TODO(),
		0,
		cdn.ResourceRuleGetParams{
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

func TestResourceRuleReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.Rules.Replace(
		context.TODO(),
		0,
		cdn.ResourceRuleReplaceParams{
			ResourceID: 0,
			Rule:       "/folder/images/*.png",
			RuleType:   0,
			Active:     gcore.Bool(true),
			Name:       gcore.String("My first rule"),
			Options: cdn.ResourceRuleReplaceParamsOptions{
				AllowedHTTPMethods: cdn.ResourceRuleReplaceParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.ResourceRuleReplaceParamsOptionsBotProtection{
					BotChallenge: cdn.ResourceRuleReplaceParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.ResourceRuleReplaceParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.ResourceRuleReplaceParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.ResourceRuleReplaceParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.ResourceRuleReplaceParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.ResourceRuleReplaceParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.ResourceRuleReplaceParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.ResourceRuleReplaceParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.ResourceRuleReplaceParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.ResourceRuleReplaceParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.ResourceRuleReplaceParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.ResourceRuleReplaceParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.ResourceRuleReplaceParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.ResourceRuleReplaceParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.ResourceRuleReplaceParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.ResourceRuleReplaceParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.ResourceRuleReplaceParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.ResourceRuleReplaceParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.ResourceRuleReplaceParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.ResourceRuleReplaceParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.ResourceRuleReplaceParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.ResourceRuleReplaceParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.ResourceRuleReplaceParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.ResourceRuleReplaceParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.ResourceRuleReplaceParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.ResourceRuleReplaceParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.ResourceRuleReplaceParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.ResourceRuleReplaceParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.ResourceRuleReplaceParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.ResourceRuleReplaceParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.ResourceRuleReplaceParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.ResourceRuleReplaceParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.ResourceRuleReplaceParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.ResourceRuleReplaceParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.ResourceRuleReplaceParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.ResourceRuleReplaceParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.ResourceRuleReplaceParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.ResourceRuleReplaceParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.ResourceRuleReplaceParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.ResourceRuleReplaceParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.ResourceRuleReplaceParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.ResourceRuleReplaceParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.ResourceRuleReplaceParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.ResourceRuleReplaceParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.ResourceRuleReplaceParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.ResourceRuleReplaceParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.ResourceRuleReplaceParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.ResourceRuleReplaceParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.ResourceRuleReplaceParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.ResourceRuleReplaceParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:            param.Null[int64](),
			OverrideOriginProtocol: cdn.ResourceRuleReplaceParamsOverrideOriginProtocolHTTPS,
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
