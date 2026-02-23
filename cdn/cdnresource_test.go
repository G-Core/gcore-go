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

func TestCDNResourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.CDNResources.New(context.TODO(), cdn.CDNResourceNewParams{
		Cname:       "cdn.site.com",
		Active:      gcore.Bool(true),
		Description: gcore.String("My resource"),
		Name:        gcore.String("Resource for images"),
		Options: cdn.CDNResourceNewParamsOptions{
			AllowedHTTPMethods: cdn.CDNResourceNewParamsOptionsAllowedHTTPMethods{
				Enabled: true,
				Value:   []string{"GET", "POST"},
			},
			BotProtection: cdn.CDNResourceNewParamsOptionsBotProtection{
				BotChallenge: cdn.CDNResourceNewParamsOptionsBotProtectionBotChallenge{
					Enabled: gcore.Bool(true),
				},
				Enabled: true,
			},
			BrotliCompression: cdn.CDNResourceNewParamsOptionsBrotliCompression{
				Enabled: true,
				Value:   []string{"text/html", "text/plain"},
			},
			BrowserCacheSettings: cdn.CDNResourceNewParamsOptionsBrowserCacheSettings{
				Enabled: true,
				Value:   "3600s",
			},
			CacheHTTPHeaders: cdn.CDNResourceNewParamsOptionsCacheHTTPHeaders{
				Enabled: false,
				Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
			},
			Cors: cdn.CDNResourceNewParamsOptionsCors{
				Enabled: true,
				Value:   []string{"domain.com", "domain2.com"},
				Always:  gcore.Bool(true),
			},
			CountryACL: cdn.CDNResourceNewParamsOptionsCountryACL{
				Enabled:        true,
				ExceptedValues: []string{"GB", "DE"},
				PolicyType:     "allow",
			},
			DisableCache: cdn.CDNResourceNewParamsOptionsDisableCache{
				Enabled: true,
				Value:   false,
			},
			DisableProxyForceRanges: cdn.CDNResourceNewParamsOptionsDisableProxyForceRanges{
				Enabled: true,
				Value:   true,
			},
			EdgeCacheSettings: cdn.CDNResourceNewParamsOptionsEdgeCacheSettings{
				Enabled: true,
				CustomValues: map[string]string{
					"100": "43200s",
				},
				Default: gcore.String("321669910225"),
				Value:   gcore.String("43200s"),
			},
			Fastedge: cdn.CDNResourceNewParamsOptionsFastedge{
				Enabled: true,
				OnRequestBody: cdn.CDNResourceNewParamsOptionsFastedgeOnRequestBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnRequestHeaders: cdn.CDNResourceNewParamsOptionsFastedgeOnRequestHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseBody: cdn.CDNResourceNewParamsOptionsFastedgeOnResponseBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseHeaders: cdn.CDNResourceNewParamsOptionsFastedgeOnResponseHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
			},
			FetchCompressed: cdn.CDNResourceNewParamsOptionsFetchCompressed{
				Enabled: true,
				Value:   false,
			},
			FollowOriginRedirect: cdn.CDNResourceNewParamsOptionsFollowOriginRedirect{
				Codes:   []int64{302, 308},
				Enabled: true,
			},
			ForceReturn: cdn.CDNResourceNewParamsOptionsForceReturn{
				Body:    "http://example.com/redirect_address",
				Code:    301,
				Enabled: true,
				TimeInterval: cdn.CDNResourceNewParamsOptionsForceReturnTimeInterval{
					EndTime:   "20:00",
					StartTime: "09:00",
					TimeZone:  gcore.String("CET"),
				},
			},
			ForwardHostHeader: cdn.CDNResourceNewParamsOptionsForwardHostHeader{
				Enabled: false,
				Value:   false,
			},
			GrpcPassthrough: cdn.CDNResourceNewParamsOptionsGrpcPassthrough{
				Enabled: true,
				Value:   true,
			},
			GzipOn: cdn.CDNResourceNewParamsOptionsGzipOn{
				Enabled: true,
				Value:   true,
			},
			HostHeader: cdn.CDNResourceNewParamsOptionsHostHeader{
				Enabled: true,
				Value:   "host.com",
			},
			Http3Enabled: cdn.CDNResourceNewParamsOptionsHttp3Enabled{
				Enabled: true,
				Value:   true,
			},
			IgnoreCookie: cdn.CDNResourceNewParamsOptionsIgnoreCookie{
				Enabled: true,
				Value:   true,
			},
			IgnoreQueryString: cdn.CDNResourceNewParamsOptionsIgnoreQueryString{
				Enabled: true,
				Value:   false,
			},
			ImageStack: cdn.CDNResourceNewParamsOptionsImageStack{
				Enabled:     true,
				AvifEnabled: gcore.Bool(true),
				PngLossless: gcore.Bool(true),
				Quality:     gcore.Int(80),
				WebpEnabled: gcore.Bool(false),
			},
			IPAddressACL: cdn.CDNResourceNewParamsOptionsIPAddressACL{
				Enabled:        true,
				ExceptedValues: []string{"192.168.1.100/32"},
				PolicyType:     "deny",
			},
			LimitBandwidth: cdn.CDNResourceNewParamsOptionsLimitBandwidth{
				Enabled:   true,
				LimitType: "static",
				Buffer:    gcore.Int(200),
				Speed:     gcore.Int(100),
			},
			ProxyCacheKey: cdn.CDNResourceNewParamsOptionsProxyCacheKey{
				Enabled: true,
				Value:   "$scheme$uri",
			},
			ProxyCacheMethodsSet: cdn.CDNResourceNewParamsOptionsProxyCacheMethodsSet{
				Enabled: true,
				Value:   false,
			},
			ProxyConnectTimeout: cdn.CDNResourceNewParamsOptionsProxyConnectTimeout{
				Enabled: true,
				Value:   "4s",
			},
			ProxyReadTimeout: cdn.CDNResourceNewParamsOptionsProxyReadTimeout{
				Enabled: true,
				Value:   "10s",
			},
			QueryParamsBlacklist: cdn.CDNResourceNewParamsOptionsQueryParamsBlacklist{
				Enabled: true,
				Value:   []string{"some", "blacklisted", "query"},
			},
			QueryParamsWhitelist: cdn.CDNResourceNewParamsOptionsQueryParamsWhitelist{
				Enabled: true,
				Value:   []string{"some", "whitelisted", "query"},
			},
			QueryStringForwarding: cdn.CDNResourceNewParamsOptionsQueryStringForwarding{
				Enabled:              true,
				ForwardFromFileTypes: []string{"m3u8", "mpd"},
				ForwardToFileTypes:   []string{"ts", "mp4"},
				ForwardExceptKeys:    []string{"debug_info"},
				ForwardOnlyKeys:      []string{"auth_token", "session_id"},
			},
			RedirectHTTPToHTTPS: cdn.CDNResourceNewParamsOptionsRedirectHTTPToHTTPS{
				Enabled: true,
				Value:   true,
			},
			RedirectHTTPSToHTTP: cdn.CDNResourceNewParamsOptionsRedirectHTTPSToHTTP{
				Enabled: false,
				Value:   true,
			},
			ReferrerACL: cdn.CDNResourceNewParamsOptionsReferrerACL{
				Enabled:        true,
				ExceptedValues: []string{"example.com", "*.example.net"},
				PolicyType:     "deny",
			},
			RequestLimiter: cdn.CDNResourceNewParamsOptionsRequestLimiter{
				Enabled:  true,
				Rate:     5,
				RateUnit: "r/s",
			},
			ResponseHeadersHidingPolicy: cdn.CDNResourceNewParamsOptionsResponseHeadersHidingPolicy{
				Enabled:  true,
				Excepted: []string{"my-header"},
				Mode:     "hide",
			},
			Rewrite: cdn.CDNResourceNewParamsOptionsRewrite{
				Body:    "/(.*) /additional_path/$1",
				Enabled: true,
				Flag:    "break",
			},
			SecureKey: cdn.CDNResourceNewParamsOptionsSecureKey{
				Enabled: true,
				Key:     gcore.String("secretkey"),
				Type:    2,
			},
			Slice: cdn.CDNResourceNewParamsOptionsSlice{
				Enabled: true,
				Value:   true,
			},
			Sni: cdn.CDNResourceNewParamsOptionsSni{
				CustomHostname: "custom.example.com",
				Enabled:        true,
				SniType:        "custom",
			},
			Stale: cdn.CDNResourceNewParamsOptionsStale{
				Enabled: true,
				Value:   []string{"http_404", "http_500"},
			},
			StaticResponseHeaders: cdn.CDNResourceNewParamsOptionsStaticResponseHeaders{
				Enabled: true,
				Value: []cdn.CDNResourceNewParamsOptionsStaticResponseHeadersValue{{
					Name:   "X-Example",
					Value:  []string{"Value_1"},
					Always: gcore.Bool(true),
				}, {
					Name:   "X-Example-Multiple",
					Value:  []string{"Value_1", "Value_2", "Value_3"},
					Always: gcore.Bool(false),
				}},
			},
			StaticHeaders: cdn.CDNResourceNewParamsOptionsStaticHeaders{
				Enabled: true,
				Value: map[string]any{
					"X-Example": "Value_1",
					"X-Example-Multiple": []string{
						"Value_2",
						"Value_3",
					},
				},
			},
			StaticRequestHeaders: cdn.CDNResourceNewParamsOptionsStaticRequestHeaders{
				Enabled: true,
				Value: map[string]string{
					"Header-One": "Value 1",
					"Header-Two": "Value 2",
				},
			},
			TlsVersions: cdn.CDNResourceNewParamsOptionsTlsVersions{
				Enabled: true,
				Value:   []string{"SSLv3", "TLSv1.3"},
			},
			UseDefaultLeChain: cdn.CDNResourceNewParamsOptionsUseDefaultLeChain{
				Enabled: true,
				Value:   true,
			},
			UseDns01LeChallenge: cdn.CDNResourceNewParamsOptionsUseDns01LeChallenge{
				Enabled: true,
				Value:   true,
			},
			UseRsaLeCert: cdn.CDNResourceNewParamsOptionsUseRsaLeCert{
				Enabled: true,
				Value:   true,
			},
			UserAgentACL: cdn.CDNResourceNewParamsOptionsUserAgentACL{
				Enabled:        true,
				ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
				PolicyType:     "allow",
			},
			Waap: cdn.CDNResourceNewParamsOptionsWaap{
				Enabled: true,
				Value:   true,
			},
			Websockets: cdn.CDNResourceNewParamsOptionsWebsockets{
				Enabled: true,
				Value:   true,
			},
		},
		Origin:               gcore.String("example.com"),
		OriginGroup:          gcore.Int(132),
		OriginProtocol:       cdn.CDNResourceNewParamsOriginProtocolHTTPS,
		PrimaryResource:      param.Null[int64](),
		ProxySslCa:           param.Null[int64](),
		ProxySslData:         param.Null[int64](),
		ProxySslEnabled:      gcore.Bool(false),
		SecondaryHostnames:   []string{"first.example.com", "second.example.com"},
		SslData:              gcore.Int(192),
		SslEnabled:           gcore.Bool(false),
		WaapAPIDomainEnabled: gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceUpdateWithOptionalParams(t *testing.T) {
	t.Skip("unexpected prism python test failures")
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
	_, err := client.CDN.CDNResources.Update(
		context.TODO(),
		0,
		cdn.CDNResourceUpdateParams{
			Active:      gcore.Bool(true),
			Description: gcore.String("My resource"),
			Name:        gcore.String("Resource for images"),
			Options: cdn.CDNResourceUpdateParamsOptions{
				AllowedHTTPMethods: cdn.CDNResourceUpdateParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.CDNResourceUpdateParamsOptionsBotProtection{
					BotChallenge: cdn.CDNResourceUpdateParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.CDNResourceUpdateParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.CDNResourceUpdateParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.CDNResourceUpdateParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.CDNResourceUpdateParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.CDNResourceUpdateParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.CDNResourceUpdateParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.CDNResourceUpdateParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.CDNResourceUpdateParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.CDNResourceUpdateParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.CDNResourceUpdateParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.CDNResourceUpdateParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.CDNResourceUpdateParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.CDNResourceUpdateParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.CDNResourceUpdateParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.CDNResourceUpdateParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.CDNResourceUpdateParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.CDNResourceUpdateParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.CDNResourceUpdateParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GrpcPassthrough: cdn.CDNResourceUpdateParamsOptionsGrpcPassthrough{
					Enabled: true,
					Value:   true,
				},
				GzipOn: cdn.CDNResourceUpdateParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.CDNResourceUpdateParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				Http3Enabled: cdn.CDNResourceUpdateParamsOptionsHttp3Enabled{
					Enabled: true,
					Value:   true,
				},
				IgnoreCookie: cdn.CDNResourceUpdateParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.CDNResourceUpdateParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.CDNResourceUpdateParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.CDNResourceUpdateParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.CDNResourceUpdateParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.CDNResourceUpdateParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.CDNResourceUpdateParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.CDNResourceUpdateParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.CDNResourceUpdateParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.CDNResourceUpdateParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.CDNResourceUpdateParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.CDNResourceUpdateParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.CDNResourceUpdateParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.CDNResourceUpdateParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.CDNResourceUpdateParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.CDNResourceUpdateParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.CDNResourceUpdateParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.CDNResourceUpdateParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.CDNResourceUpdateParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.CDNResourceUpdateParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.CDNResourceUpdateParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.CDNResourceUpdateParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.CDNResourceUpdateParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.CDNResourceUpdateParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.CDNResourceUpdateParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.CDNResourceUpdateParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				TlsVersions: cdn.CDNResourceUpdateParamsOptionsTlsVersions{
					Enabled: true,
					Value:   []string{"SSLv3", "TLSv1.3"},
				},
				UseDefaultLeChain: cdn.CDNResourceUpdateParamsOptionsUseDefaultLeChain{
					Enabled: true,
					Value:   true,
				},
				UseDns01LeChallenge: cdn.CDNResourceUpdateParamsOptionsUseDns01LeChallenge{
					Enabled: true,
					Value:   true,
				},
				UseRsaLeCert: cdn.CDNResourceUpdateParamsOptionsUseRsaLeCert{
					Enabled: true,
					Value:   true,
				},
				UserAgentACL: cdn.CDNResourceUpdateParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.CDNResourceUpdateParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.CDNResourceUpdateParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:        gcore.Int(132),
			OriginProtocol:     cdn.CDNResourceUpdateParamsOriginProtocolHTTPS,
			ProxySslCa:         param.Null[int64](),
			ProxySslData:       param.Null[int64](),
			ProxySslEnabled:    gcore.Bool(false),
			SecondaryHostnames: []string{"first.example.com", "second.example.com"},
			SslData:            gcore.Int(192),
			SslEnabled:         gcore.Bool(false),
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

func TestCDNResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.CDNResources.List(context.TODO(), cdn.CDNResourceListParams{
		Cname:              gcore.String("cname"),
		Deleted:            gcore.Bool(true),
		Enabled:            gcore.Bool(true),
		MaxCreated:         gcore.String("max_created"),
		MinCreated:         gcore.String("min_created"),
		OriginGroup:        gcore.Int(0),
		Rules:              gcore.String("rules"),
		SecondaryHostnames: gcore.String("secondaryHostnames"),
		ShieldDc:           gcore.String("shield_dc"),
		Shielded:           gcore.Bool(true),
		SslData:            gcore.Int(0),
		SslDataIn:          gcore.Int(0),
		SslEnabled:         gcore.Bool(true),
		Status:             cdn.CDNResourceListParamsStatusActive,
		Suspend:            gcore.Bool(true),
		VpEnabled:          gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceDelete(t *testing.T) {
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
	err := client.CDN.CDNResources.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourceGet(t *testing.T) {
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
	_, err := client.CDN.CDNResources.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourcePrefetch(t *testing.T) {
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
	err := client.CDN.CDNResources.Prefetch(
		context.TODO(),
		0,
		cdn.CDNResourcePrefetchParams{
			Paths: []string{"/test.jpg", "test1.jpg"},
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

func TestCDNResourcePrevalidateSslLeCertificate(t *testing.T) {
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
	err := client.CDN.CDNResources.PrevalidateSslLeCertificate(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCDNResourcePurgeWithOptionalParams(t *testing.T) {
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
	err := client.CDN.CDNResources.Purge(
		context.TODO(),
		0,
		cdn.CDNResourcePurgeParams{
			OfPurgeByURL: &cdn.CDNResourcePurgeParamsBodyPurgeByURL{
				URLs: []string{"/some-url.jpg", "/img/example.jpg"},
			},
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

func TestCDNResourceReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.CDNResources.Replace(
		context.TODO(),
		0,
		cdn.CDNResourceReplaceParams{
			OriginGroup: 132,
			Active:      gcore.Bool(true),
			Description: gcore.String("My resource"),
			Name:        gcore.String("Resource for images"),
			Options: cdn.CDNResourceReplaceParamsOptions{
				AllowedHTTPMethods: cdn.CDNResourceReplaceParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.CDNResourceReplaceParamsOptionsBotProtection{
					BotChallenge: cdn.CDNResourceReplaceParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.CDNResourceReplaceParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.CDNResourceReplaceParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.CDNResourceReplaceParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.CDNResourceReplaceParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.CDNResourceReplaceParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.CDNResourceReplaceParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.CDNResourceReplaceParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.CDNResourceReplaceParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.CDNResourceReplaceParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.CDNResourceReplaceParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.CDNResourceReplaceParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.CDNResourceReplaceParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.CDNResourceReplaceParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.CDNResourceReplaceParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.CDNResourceReplaceParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.CDNResourceReplaceParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.CDNResourceReplaceParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.CDNResourceReplaceParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GrpcPassthrough: cdn.CDNResourceReplaceParamsOptionsGrpcPassthrough{
					Enabled: true,
					Value:   true,
				},
				GzipOn: cdn.CDNResourceReplaceParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.CDNResourceReplaceParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				Http3Enabled: cdn.CDNResourceReplaceParamsOptionsHttp3Enabled{
					Enabled: true,
					Value:   true,
				},
				IgnoreCookie: cdn.CDNResourceReplaceParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.CDNResourceReplaceParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.CDNResourceReplaceParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.CDNResourceReplaceParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.CDNResourceReplaceParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.CDNResourceReplaceParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.CDNResourceReplaceParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.CDNResourceReplaceParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.CDNResourceReplaceParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.CDNResourceReplaceParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.CDNResourceReplaceParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.CDNResourceReplaceParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.CDNResourceReplaceParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.CDNResourceReplaceParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.CDNResourceReplaceParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.CDNResourceReplaceParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.CDNResourceReplaceParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.CDNResourceReplaceParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.CDNResourceReplaceParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.CDNResourceReplaceParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.CDNResourceReplaceParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.CDNResourceReplaceParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.CDNResourceReplaceParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.CDNResourceReplaceParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.CDNResourceReplaceParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]any{
						"X-Example": "Value_1",
						"X-Example-Multiple": []string{
							"Value_2",
							"Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.CDNResourceReplaceParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				TlsVersions: cdn.CDNResourceReplaceParamsOptionsTlsVersions{
					Enabled: true,
					Value:   []string{"SSLv3", "TLSv1.3"},
				},
				UseDefaultLeChain: cdn.CDNResourceReplaceParamsOptionsUseDefaultLeChain{
					Enabled: true,
					Value:   true,
				},
				UseDns01LeChallenge: cdn.CDNResourceReplaceParamsOptionsUseDns01LeChallenge{
					Enabled: true,
					Value:   true,
				},
				UseRsaLeCert: cdn.CDNResourceReplaceParamsOptionsUseRsaLeCert{
					Enabled: true,
					Value:   true,
				},
				UserAgentACL: cdn.CDNResourceReplaceParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", "~*.*bot.*", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.CDNResourceReplaceParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.CDNResourceReplaceParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginProtocol:       cdn.CDNResourceReplaceParamsOriginProtocolHTTPS,
			ProxySslCa:           param.Null[int64](),
			ProxySslData:         param.Null[int64](),
			ProxySslEnabled:      gcore.Bool(false),
			SecondaryHostnames:   []string{"first.example.com", "second.example.com"},
			SslData:              gcore.Int(192),
			SslEnabled:           gcore.Bool(false),
			WaapAPIDomainEnabled: gcore.Bool(true),
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
