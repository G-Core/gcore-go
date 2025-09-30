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
)

func TestRuleTemplateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.RuleTemplates.New(context.TODO(), cdn.RuleTemplateNewParams{
		Rule:     "/folder/images/*.png",
		RuleType: 0,
		Name:     gcore.String("All images template"),
		Options: cdn.RuleTemplateNewParamsOptions{
			AllowedHTTPMethods: cdn.RuleTemplateNewParamsOptionsAllowedHTTPMethods{
				Enabled: true,
				Value:   []string{"GET", "POST"},
			},
			BotProtection: cdn.RuleTemplateNewParamsOptionsBotProtection{
				BotChallenge: cdn.RuleTemplateNewParamsOptionsBotProtectionBotChallenge{
					Enabled: gcore.Bool(true),
				},
				Enabled: true,
			},
			BrotliCompression: cdn.RuleTemplateNewParamsOptionsBrotliCompression{
				Enabled: true,
				Value:   []string{"text/html", "text/plain"},
			},
			BrowserCacheSettings: cdn.RuleTemplateNewParamsOptionsBrowserCacheSettings{
				Enabled: true,
				Value:   "3600s",
			},
			CacheHTTPHeaders: cdn.RuleTemplateNewParamsOptionsCacheHTTPHeaders{
				Enabled: false,
				Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
			},
			Cors: cdn.RuleTemplateNewParamsOptionsCors{
				Enabled: true,
				Value:   []string{"domain.com", "domain2.com"},
				Always:  gcore.Bool(true),
			},
			CountryACL: cdn.RuleTemplateNewParamsOptionsCountryACL{
				Enabled:        true,
				ExceptedValues: []string{"GB", "DE"},
				PolicyType:     "allow",
			},
			DisableCache: cdn.RuleTemplateNewParamsOptionsDisableCache{
				Enabled: true,
				Value:   false,
			},
			DisableProxyForceRanges: cdn.RuleTemplateNewParamsOptionsDisableProxyForceRanges{
				Enabled: true,
				Value:   true,
			},
			EdgeCacheSettings: cdn.RuleTemplateNewParamsOptionsEdgeCacheSettings{
				Enabled: true,
				CustomValues: map[string]string{
					"100": "43200s",
				},
				Default: gcore.String("321669910225"),
				Value:   gcore.String("43200s"),
			},
			Fastedge: cdn.RuleTemplateNewParamsOptionsFastedge{
				Enabled: true,
				OnRequestBody: cdn.RuleTemplateNewParamsOptionsFastedgeOnRequestBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnRequestHeaders: cdn.RuleTemplateNewParamsOptionsFastedgeOnRequestHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseBody: cdn.RuleTemplateNewParamsOptionsFastedgeOnResponseBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseHeaders: cdn.RuleTemplateNewParamsOptionsFastedgeOnResponseHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
			},
			FetchCompressed: cdn.RuleTemplateNewParamsOptionsFetchCompressed{
				Enabled: true,
				Value:   false,
			},
			FollowOriginRedirect: cdn.RuleTemplateNewParamsOptionsFollowOriginRedirect{
				Codes:   []int64{302, 308},
				Enabled: true,
			},
			ForceReturn: cdn.RuleTemplateNewParamsOptionsForceReturn{
				Body:    "body",
				Code:    100,
				Enabled: true,
				TimeInterval: cdn.RuleTemplateNewParamsOptionsForceReturnTimeInterval{
					EndTime:   "18:11:19.117Z",
					StartTime: "18:11:19.117Z",
					TimeZone:  gcore.String("Europe/Luxembourg"),
				},
			},
			ForwardHostHeader: cdn.RuleTemplateNewParamsOptionsForwardHostHeader{
				Enabled: false,
				Value:   false,
			},
			GzipOn: cdn.RuleTemplateNewParamsOptionsGzipOn{
				Enabled: true,
				Value:   true,
			},
			HostHeader: cdn.RuleTemplateNewParamsOptionsHostHeader{
				Enabled: true,
				Value:   "host.com",
			},
			IgnoreCookie: cdn.RuleTemplateNewParamsOptionsIgnoreCookie{
				Enabled: true,
				Value:   true,
			},
			IgnoreQueryString: cdn.RuleTemplateNewParamsOptionsIgnoreQueryString{
				Enabled: true,
				Value:   false,
			},
			ImageStack: cdn.RuleTemplateNewParamsOptionsImageStack{
				Enabled:     true,
				AvifEnabled: gcore.Bool(true),
				PngLossless: gcore.Bool(true),
				Quality:     gcore.Int(80),
				WebpEnabled: gcore.Bool(false),
			},
			IPAddressACL: cdn.RuleTemplateNewParamsOptionsIPAddressACL{
				Enabled:        true,
				ExceptedValues: []string{"192.168.1.100/32"},
				PolicyType:     "deny",
			},
			LimitBandwidth: cdn.RuleTemplateNewParamsOptionsLimitBandwidth{
				Enabled:   true,
				LimitType: "static",
				Buffer:    gcore.Int(200),
				Speed:     gcore.Int(100),
			},
			ProxyCacheKey: cdn.RuleTemplateNewParamsOptionsProxyCacheKey{
				Enabled: true,
				Value:   "$scheme$uri",
			},
			ProxyCacheMethodsSet: cdn.RuleTemplateNewParamsOptionsProxyCacheMethodsSet{
				Enabled: true,
				Value:   false,
			},
			ProxyConnectTimeout: cdn.RuleTemplateNewParamsOptionsProxyConnectTimeout{
				Enabled: true,
				Value:   "4s",
			},
			ProxyReadTimeout: cdn.RuleTemplateNewParamsOptionsProxyReadTimeout{
				Enabled: true,
				Value:   "10s",
			},
			QueryParamsBlacklist: cdn.RuleTemplateNewParamsOptionsQueryParamsBlacklist{
				Enabled: true,
				Value:   []string{"some", "blacklisted", "query"},
			},
			QueryParamsWhitelist: cdn.RuleTemplateNewParamsOptionsQueryParamsWhitelist{
				Enabled: true,
				Value:   []string{"some", "whitelisted", "query"},
			},
			QueryStringForwarding: cdn.RuleTemplateNewParamsOptionsQueryStringForwarding{
				Enabled:              true,
				ForwardFromFileTypes: []string{"m3u8", "mpd"},
				ForwardToFileTypes:   []string{"ts", "mp4"},
			},
			RedirectHTTPToHTTPS: cdn.RuleTemplateNewParamsOptionsRedirectHTTPToHTTPS{
				Enabled: true,
				Value:   true,
			},
			RedirectHTTPSToHTTP: cdn.RuleTemplateNewParamsOptionsRedirectHTTPSToHTTP{
				Enabled: false,
				Value:   true,
			},
			ReferrerACL: cdn.RuleTemplateNewParamsOptionsReferrerACL{
				Enabled:        true,
				ExceptedValues: []string{"example.com", "*.example.net"},
				PolicyType:     "deny",
			},
			RequestLimiter: cdn.RuleTemplateNewParamsOptionsRequestLimiter{
				Enabled:  true,
				Rate:     5,
				RateUnit: "r/s",
			},
			ResponseHeadersHidingPolicy: cdn.RuleTemplateNewParamsOptionsResponseHeadersHidingPolicy{
				Enabled:  true,
				Excepted: []string{"my-header"},
				Mode:     "hide",
			},
			Rewrite: cdn.RuleTemplateNewParamsOptionsRewrite{
				Body:    "/(.*) /additional_path/$1",
				Enabled: true,
				Flag:    "break",
			},
			SecureKey: cdn.RuleTemplateNewParamsOptionsSecureKey{
				Enabled: true,
				Key:     gcore.String("secretkey"),
				Type:    2,
			},
			Slice: cdn.RuleTemplateNewParamsOptionsSlice{
				Enabled: true,
				Value:   true,
			},
			Sni: cdn.RuleTemplateNewParamsOptionsSni{
				CustomHostname: "custom.example.com",
				Enabled:        true,
				SniType:        "custom",
			},
			Stale: cdn.RuleTemplateNewParamsOptionsStale{
				Enabled: true,
				Value:   []string{"http_404", "http_500"},
			},
			StaticResponseHeaders: cdn.RuleTemplateNewParamsOptionsStaticResponseHeaders{
				Enabled: true,
				Value: []cdn.RuleTemplateNewParamsOptionsStaticResponseHeadersValue{{
					Name:   "X-Example",
					Value:  []string{"Value_1"},
					Always: gcore.Bool(true),
				}, {
					Name:   "X-Example-Multiple",
					Value:  []string{"Value_1", "Value_2", "Value_3"},
					Always: gcore.Bool(false),
				}},
			},
			StaticHeaders: cdn.RuleTemplateNewParamsOptionsStaticHeaders{
				Enabled: true,
				Value: map[string]string{
					"foo": "string",
				},
			},
			StaticRequestHeaders: cdn.RuleTemplateNewParamsOptionsStaticRequestHeaders{
				Enabled: true,
				Value: map[string]string{
					"Header-One": "Value 1",
					"Header-Two": "Value 2",
				},
			},
			UserAgentACL: cdn.RuleTemplateNewParamsOptionsUserAgentACL{
				Enabled:        true,
				ExceptedValues: []string{"UserAgent Value", ""},
				PolicyType:     "allow",
			},
			Waap: cdn.RuleTemplateNewParamsOptionsWaap{
				Enabled: true,
				Value:   true,
			},
			Websockets: cdn.RuleTemplateNewParamsOptionsWebsockets{
				Enabled: true,
				Value:   true,
			},
		},
		OverrideOriginProtocol: cdn.RuleTemplateNewParamsOverrideOriginProtocolHTTPS,
		Weight:                 gcore.Int(1),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.RuleTemplates.Update(
		context.TODO(),
		0,
		cdn.RuleTemplateUpdateParams{
			Name: gcore.String("All images template"),
			Options: cdn.RuleTemplateUpdateParamsOptions{
				AllowedHTTPMethods: cdn.RuleTemplateUpdateParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.RuleTemplateUpdateParamsOptionsBotProtection{
					BotChallenge: cdn.RuleTemplateUpdateParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.RuleTemplateUpdateParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.RuleTemplateUpdateParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.RuleTemplateUpdateParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.RuleTemplateUpdateParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.RuleTemplateUpdateParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.RuleTemplateUpdateParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.RuleTemplateUpdateParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.RuleTemplateUpdateParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.RuleTemplateUpdateParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.RuleTemplateUpdateParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.RuleTemplateUpdateParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.RuleTemplateUpdateParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.RuleTemplateUpdateParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.RuleTemplateUpdateParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.RuleTemplateUpdateParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.RuleTemplateUpdateParamsOptionsForceReturn{
					Body:    "body",
					Code:    100,
					Enabled: true,
					TimeInterval: cdn.RuleTemplateUpdateParamsOptionsForceReturnTimeInterval{
						EndTime:   "18:11:19.117Z",
						StartTime: "18:11:19.117Z",
						TimeZone:  gcore.String("Europe/Luxembourg"),
					},
				},
				ForwardHostHeader: cdn.RuleTemplateUpdateParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.RuleTemplateUpdateParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.RuleTemplateUpdateParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.RuleTemplateUpdateParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.RuleTemplateUpdateParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.RuleTemplateUpdateParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.RuleTemplateUpdateParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.RuleTemplateUpdateParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.RuleTemplateUpdateParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.RuleTemplateUpdateParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.RuleTemplateUpdateParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.RuleTemplateUpdateParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.RuleTemplateUpdateParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.RuleTemplateUpdateParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.RuleTemplateUpdateParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
				},
				RedirectHTTPToHTTPS: cdn.RuleTemplateUpdateParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.RuleTemplateUpdateParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.RuleTemplateUpdateParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.RuleTemplateUpdateParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.RuleTemplateUpdateParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.RuleTemplateUpdateParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.RuleTemplateUpdateParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.RuleTemplateUpdateParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.RuleTemplateUpdateParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.RuleTemplateUpdateParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.RuleTemplateUpdateParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.RuleTemplateUpdateParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.RuleTemplateUpdateParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]string{
						"foo": "string",
					},
				},
				StaticRequestHeaders: cdn.RuleTemplateUpdateParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.RuleTemplateUpdateParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.RuleTemplateUpdateParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.RuleTemplateUpdateParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OverrideOriginProtocol: cdn.RuleTemplateUpdateParamsOverrideOriginProtocolHTTPS,
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

func TestRuleTemplateList(t *testing.T) {
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
	_, err := client.Cdn.RuleTemplates.List(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleTemplateDelete(t *testing.T) {
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
	err := client.Cdn.RuleTemplates.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleTemplateGet(t *testing.T) {
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
	_, err := client.Cdn.RuleTemplates.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRuleTemplateReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.RuleTemplates.Replace(
		context.TODO(),
		0,
		cdn.RuleTemplateReplaceParams{
			Rule:     "/folder/images/*.png",
			RuleType: 0,
			Name:     gcore.String("All images template"),
			Options: cdn.RuleTemplateReplaceParamsOptions{
				AllowedHTTPMethods: cdn.RuleTemplateReplaceParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.RuleTemplateReplaceParamsOptionsBotProtection{
					BotChallenge: cdn.RuleTemplateReplaceParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.RuleTemplateReplaceParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.RuleTemplateReplaceParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.RuleTemplateReplaceParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.RuleTemplateReplaceParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.RuleTemplateReplaceParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.RuleTemplateReplaceParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.RuleTemplateReplaceParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.RuleTemplateReplaceParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.RuleTemplateReplaceParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.RuleTemplateReplaceParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.RuleTemplateReplaceParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.RuleTemplateReplaceParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.RuleTemplateReplaceParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.RuleTemplateReplaceParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.RuleTemplateReplaceParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.RuleTemplateReplaceParamsOptionsForceReturn{
					Body:    "body",
					Code:    100,
					Enabled: true,
					TimeInterval: cdn.RuleTemplateReplaceParamsOptionsForceReturnTimeInterval{
						EndTime:   "18:11:19.117Z",
						StartTime: "18:11:19.117Z",
						TimeZone:  gcore.String("Europe/Luxembourg"),
					},
				},
				ForwardHostHeader: cdn.RuleTemplateReplaceParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.RuleTemplateReplaceParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.RuleTemplateReplaceParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				IgnoreCookie: cdn.RuleTemplateReplaceParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.RuleTemplateReplaceParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.RuleTemplateReplaceParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.RuleTemplateReplaceParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.RuleTemplateReplaceParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.RuleTemplateReplaceParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.RuleTemplateReplaceParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.RuleTemplateReplaceParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.RuleTemplateReplaceParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.RuleTemplateReplaceParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.RuleTemplateReplaceParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.RuleTemplateReplaceParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
				},
				RedirectHTTPToHTTPS: cdn.RuleTemplateReplaceParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.RuleTemplateReplaceParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.RuleTemplateReplaceParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.RuleTemplateReplaceParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.RuleTemplateReplaceParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.RuleTemplateReplaceParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.RuleTemplateReplaceParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.RuleTemplateReplaceParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.RuleTemplateReplaceParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.RuleTemplateReplaceParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.RuleTemplateReplaceParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.RuleTemplateReplaceParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.RuleTemplateReplaceParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]string{
						"foo": "string",
					},
				},
				StaticRequestHeaders: cdn.RuleTemplateReplaceParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				UserAgentACL: cdn.RuleTemplateReplaceParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.RuleTemplateReplaceParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.RuleTemplateReplaceParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OverrideOriginProtocol: cdn.RuleTemplateReplaceParamsOverrideOriginProtocolHTTPS,
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
