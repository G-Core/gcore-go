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

func TestResourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.New(context.TODO(), cdn.ResourceNewParams{
		Cname:       "cdn.site.com",
		Origin:      "example.com",
		OriginGroup: 132,
		Active:      gcore.Bool(true),
		Description: gcore.String("My resource"),
		Name:        gcore.String("Resource for images"),
		Options: cdn.ResourceNewParamsOptions{
			AllowedHTTPMethods: cdn.ResourceNewParamsOptionsAllowedHTTPMethods{
				Enabled: true,
				Value:   []string{"GET", "POST"},
			},
			BotProtection: cdn.ResourceNewParamsOptionsBotProtection{
				BotChallenge: cdn.ResourceNewParamsOptionsBotProtectionBotChallenge{
					Enabled: gcore.Bool(true),
				},
				Enabled: true,
			},
			BrotliCompression: cdn.ResourceNewParamsOptionsBrotliCompression{
				Enabled: true,
				Value:   []string{"text/html", "text/plain"},
			},
			BrowserCacheSettings: cdn.ResourceNewParamsOptionsBrowserCacheSettings{
				Enabled: true,
				Value:   "3600s",
			},
			CacheHTTPHeaders: cdn.ResourceNewParamsOptionsCacheHTTPHeaders{
				Enabled: false,
				Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
			},
			Cors: cdn.ResourceNewParamsOptionsCors{
				Enabled: true,
				Value:   []string{"domain.com", "domain2.com"},
				Always:  gcore.Bool(true),
			},
			CountryACL: cdn.ResourceNewParamsOptionsCountryACL{
				Enabled:        true,
				ExceptedValues: []string{"GB", "DE"},
				PolicyType:     "allow",
			},
			DisableCache: cdn.ResourceNewParamsOptionsDisableCache{
				Enabled: true,
				Value:   false,
			},
			DisableProxyForceRanges: cdn.ResourceNewParamsOptionsDisableProxyForceRanges{
				Enabled: true,
				Value:   true,
			},
			EdgeCacheSettings: cdn.ResourceNewParamsOptionsEdgeCacheSettings{
				Enabled: true,
				CustomValues: map[string]string{
					"100": "43200s",
				},
				Default: gcore.String("321669910225"),
				Value:   gcore.String("43200s"),
			},
			Fastedge: cdn.ResourceNewParamsOptionsFastedge{
				Enabled: true,
				OnRequestBody: cdn.ResourceNewParamsOptionsFastedgeOnRequestBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnRequestHeaders: cdn.ResourceNewParamsOptionsFastedgeOnRequestHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseBody: cdn.ResourceNewParamsOptionsFastedgeOnResponseBody{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
				OnResponseHeaders: cdn.ResourceNewParamsOptionsFastedgeOnResponseHeaders{
					AppID:            "1001",
					Enabled:          gcore.Bool(true),
					ExecuteOnEdge:    gcore.Bool(true),
					ExecuteOnShield:  gcore.Bool(false),
					InterruptOnError: gcore.Bool(true),
				},
			},
			FetchCompressed: cdn.ResourceNewParamsOptionsFetchCompressed{
				Enabled: true,
				Value:   false,
			},
			FollowOriginRedirect: cdn.ResourceNewParamsOptionsFollowOriginRedirect{
				Codes:   []int64{302, 308},
				Enabled: true,
			},
			ForceReturn: cdn.ResourceNewParamsOptionsForceReturn{
				Body:    "http://example.com/redirect_address",
				Code:    301,
				Enabled: true,
				TimeInterval: cdn.ResourceNewParamsOptionsForceReturnTimeInterval{
					EndTime:   "20:00",
					StartTime: "09:00",
					TimeZone:  gcore.String("CET"),
				},
			},
			ForwardHostHeader: cdn.ResourceNewParamsOptionsForwardHostHeader{
				Enabled: false,
				Value:   false,
			},
			GzipOn: cdn.ResourceNewParamsOptionsGzipOn{
				Enabled: true,
				Value:   true,
			},
			HostHeader: cdn.ResourceNewParamsOptionsHostHeader{
				Enabled: true,
				Value:   "host.com",
			},
			Http3Enabled: cdn.ResourceNewParamsOptionsHttp3Enabled{
				Enabled: true,
				Value:   true,
			},
			IgnoreCookie: cdn.ResourceNewParamsOptionsIgnoreCookie{
				Enabled: true,
				Value:   true,
			},
			IgnoreQueryString: cdn.ResourceNewParamsOptionsIgnoreQueryString{
				Enabled: true,
				Value:   false,
			},
			ImageStack: cdn.ResourceNewParamsOptionsImageStack{
				Enabled:     true,
				AvifEnabled: gcore.Bool(true),
				PngLossless: gcore.Bool(true),
				Quality:     gcore.Int(80),
				WebpEnabled: gcore.Bool(false),
			},
			IPAddressACL: cdn.ResourceNewParamsOptionsIPAddressACL{
				Enabled:        true,
				ExceptedValues: []string{"192.168.1.100/32"},
				PolicyType:     "deny",
			},
			LimitBandwidth: cdn.ResourceNewParamsOptionsLimitBandwidth{
				Enabled:   true,
				LimitType: "static",
				Buffer:    gcore.Int(200),
				Speed:     gcore.Int(100),
			},
			ProxyCacheKey: cdn.ResourceNewParamsOptionsProxyCacheKey{
				Enabled: true,
				Value:   "$scheme$uri",
			},
			ProxyCacheMethodsSet: cdn.ResourceNewParamsOptionsProxyCacheMethodsSet{
				Enabled: true,
				Value:   false,
			},
			ProxyConnectTimeout: cdn.ResourceNewParamsOptionsProxyConnectTimeout{
				Enabled: true,
				Value:   "4s",
			},
			ProxyReadTimeout: cdn.ResourceNewParamsOptionsProxyReadTimeout{
				Enabled: true,
				Value:   "10s",
			},
			QueryParamsBlacklist: cdn.ResourceNewParamsOptionsQueryParamsBlacklist{
				Enabled: true,
				Value:   []string{"some", "blacklisted", "query"},
			},
			QueryParamsWhitelist: cdn.ResourceNewParamsOptionsQueryParamsWhitelist{
				Enabled: true,
				Value:   []string{"some", "whitelisted", "query"},
			},
			QueryStringForwarding: cdn.ResourceNewParamsOptionsQueryStringForwarding{
				Enabled:              true,
				ForwardFromFileTypes: []string{"m3u8", "mpd"},
				ForwardToFileTypes:   []string{"ts", "mp4"},
				ForwardExceptKeys:    []string{"debug_info"},
				ForwardOnlyKeys:      []string{"auth_token", "session_id"},
			},
			RedirectHTTPToHTTPS: cdn.ResourceNewParamsOptionsRedirectHTTPToHTTPS{
				Enabled: true,
				Value:   true,
			},
			RedirectHTTPSToHTTP: cdn.ResourceNewParamsOptionsRedirectHTTPSToHTTP{
				Enabled: false,
				Value:   true,
			},
			ReferrerACL: cdn.ResourceNewParamsOptionsReferrerACL{
				Enabled:        true,
				ExceptedValues: []string{"example.com", "*.example.net"},
				PolicyType:     "deny",
			},
			RequestLimiter: cdn.ResourceNewParamsOptionsRequestLimiter{
				Enabled:  true,
				Rate:     5,
				RateUnit: "r/s",
			},
			ResponseHeadersHidingPolicy: cdn.ResourceNewParamsOptionsResponseHeadersHidingPolicy{
				Enabled:  true,
				Excepted: []string{"my-header"},
				Mode:     "hide",
			},
			Rewrite: cdn.ResourceNewParamsOptionsRewrite{
				Body:    "/(.*) /additional_path/$1",
				Enabled: true,
				Flag:    "break",
			},
			SecureKey: cdn.ResourceNewParamsOptionsSecureKey{
				Enabled: true,
				Key:     gcore.String("secretkey"),
				Type:    2,
			},
			Slice: cdn.ResourceNewParamsOptionsSlice{
				Enabled: true,
				Value:   true,
			},
			Sni: cdn.ResourceNewParamsOptionsSni{
				CustomHostname: "custom.example.com",
				Enabled:        true,
				SniType:        "custom",
			},
			Stale: cdn.ResourceNewParamsOptionsStale{
				Enabled: true,
				Value:   []string{"http_404", "http_500"},
			},
			StaticResponseHeaders: cdn.ResourceNewParamsOptionsStaticResponseHeaders{
				Enabled: true,
				Value: []cdn.ResourceNewParamsOptionsStaticResponseHeadersValue{{
					Name:   "X-Example",
					Value:  []string{"Value_1"},
					Always: gcore.Bool(true),
				}, {
					Name:   "X-Example-Multiple",
					Value:  []string{"Value_1", "Value_2", "Value_3"},
					Always: gcore.Bool(false),
				}},
			},
			StaticHeaders: cdn.ResourceNewParamsOptionsStaticHeaders{
				Enabled: true,
				Value: map[string]interface{}{
					"X-Example": "Value_1",
					"X-Example-Multiple": map[string]interface{}{
						"0": "Value_2",
						"1": "Value_3",
					},
				},
			},
			StaticRequestHeaders: cdn.ResourceNewParamsOptionsStaticRequestHeaders{
				Enabled: true,
				Value: map[string]string{
					"Header-One": "Value 1",
					"Header-Two": "Value 2",
				},
			},
			TlsVersions: cdn.ResourceNewParamsOptionsTlsVersions{
				Enabled: true,
				Value:   []string{"SSLv3", "TLSv1.3"},
			},
			UseDefaultLeChain: cdn.ResourceNewParamsOptionsUseDefaultLeChain{
				Enabled: true,
				Value:   true,
			},
			UseDns01LeChallenge: cdn.ResourceNewParamsOptionsUseDns01LeChallenge{
				Enabled: true,
				Value:   true,
			},
			UseRsaLeCert: cdn.ResourceNewParamsOptionsUseRsaLeCert{
				Enabled: true,
				Value:   true,
			},
			UserAgentACL: cdn.ResourceNewParamsOptionsUserAgentACL{
				Enabled:        true,
				ExceptedValues: []string{"UserAgent Value", ""},
				PolicyType:     "allow",
			},
			Waap: cdn.ResourceNewParamsOptionsWaap{
				Enabled: true,
				Value:   true,
			},
			Websockets: cdn.ResourceNewParamsOptionsWebsockets{
				Enabled: true,
				Value:   true,
			},
		},
		OriginProtocol:       cdn.ResourceNewParamsOriginProtocolHTTPS,
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

func TestResourceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.Update(
		context.TODO(),
		0,
		cdn.ResourceUpdateParams{
			Active:      gcore.Bool(true),
			Description: gcore.String("My resource"),
			Name:        gcore.String("Resource for images"),
			Options: cdn.ResourceUpdateParamsOptions{
				AllowedHTTPMethods: cdn.ResourceUpdateParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.ResourceUpdateParamsOptionsBotProtection{
					BotChallenge: cdn.ResourceUpdateParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.ResourceUpdateParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.ResourceUpdateParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.ResourceUpdateParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.ResourceUpdateParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.ResourceUpdateParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.ResourceUpdateParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.ResourceUpdateParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.ResourceUpdateParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.ResourceUpdateParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.ResourceUpdateParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.ResourceUpdateParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.ResourceUpdateParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.ResourceUpdateParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.ResourceUpdateParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.ResourceUpdateParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.ResourceUpdateParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.ResourceUpdateParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.ResourceUpdateParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.ResourceUpdateParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.ResourceUpdateParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				Http3Enabled: cdn.ResourceUpdateParamsOptionsHttp3Enabled{
					Enabled: true,
					Value:   true,
				},
				IgnoreCookie: cdn.ResourceUpdateParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.ResourceUpdateParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.ResourceUpdateParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.ResourceUpdateParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.ResourceUpdateParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.ResourceUpdateParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.ResourceUpdateParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.ResourceUpdateParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.ResourceUpdateParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.ResourceUpdateParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.ResourceUpdateParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.ResourceUpdateParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.ResourceUpdateParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.ResourceUpdateParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.ResourceUpdateParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.ResourceUpdateParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.ResourceUpdateParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.ResourceUpdateParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.ResourceUpdateParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.ResourceUpdateParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.ResourceUpdateParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.ResourceUpdateParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.ResourceUpdateParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.ResourceUpdateParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.ResourceUpdateParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]interface{}{
						"X-Example": "Value_1",
						"X-Example-Multiple": map[string]interface{}{
							"0": "Value_2",
							"1": "Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.ResourceUpdateParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				TlsVersions: cdn.ResourceUpdateParamsOptionsTlsVersions{
					Enabled: true,
					Value:   []string{"SSLv3", "TLSv1.3"},
				},
				UseDefaultLeChain: cdn.ResourceUpdateParamsOptionsUseDefaultLeChain{
					Enabled: true,
					Value:   true,
				},
				UseDns01LeChallenge: cdn.ResourceUpdateParamsOptionsUseDns01LeChallenge{
					Enabled: true,
					Value:   true,
				},
				UseRsaLeCert: cdn.ResourceUpdateParamsOptionsUseRsaLeCert{
					Enabled: true,
					Value:   true,
				},
				UserAgentACL: cdn.ResourceUpdateParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.ResourceUpdateParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.ResourceUpdateParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginGroup:        gcore.Int(132),
			OriginProtocol:     cdn.ResourceUpdateParamsOriginProtocolHTTPS,
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

func TestResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.List(context.TODO(), cdn.ResourceListParams{
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
		Status:             cdn.ResourceListParamsStatusActive,
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

func TestResourceDelete(t *testing.T) {
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
	err := client.Cdn.Resources.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceGet(t *testing.T) {
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
	_, err := client.Cdn.Resources.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourcePrefetch(t *testing.T) {
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
	err := client.Cdn.Resources.Prefetch(
		context.TODO(),
		0,
		cdn.ResourcePrefetchParams{
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

func TestResourcePrevalidateSslLeCertificate(t *testing.T) {
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
	err := client.Cdn.Resources.PrevalidateSslLeCertificate(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourcePurgeWithOptionalParams(t *testing.T) {
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
	err := client.Cdn.Resources.Purge(
		context.TODO(),
		0,
		cdn.ResourcePurgeParams{
			OfPurgeByURL: &cdn.ResourcePurgeParamsBodyPurgeByURL{
				URLs: []string{"string"},
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

func TestResourceReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Cdn.Resources.Replace(
		context.TODO(),
		0,
		cdn.ResourceReplaceParams{
			OriginGroup: 132,
			Active:      gcore.Bool(true),
			Description: gcore.String("My resource"),
			Name:        gcore.String("Resource for images"),
			Options: cdn.ResourceReplaceParamsOptions{
				AllowedHTTPMethods: cdn.ResourceReplaceParamsOptionsAllowedHTTPMethods{
					Enabled: true,
					Value:   []string{"GET", "POST"},
				},
				BotProtection: cdn.ResourceReplaceParamsOptionsBotProtection{
					BotChallenge: cdn.ResourceReplaceParamsOptionsBotProtectionBotChallenge{
						Enabled: gcore.Bool(true),
					},
					Enabled: true,
				},
				BrotliCompression: cdn.ResourceReplaceParamsOptionsBrotliCompression{
					Enabled: true,
					Value:   []string{"text/html", "text/plain"},
				},
				BrowserCacheSettings: cdn.ResourceReplaceParamsOptionsBrowserCacheSettings{
					Enabled: true,
					Value:   "3600s",
				},
				CacheHTTPHeaders: cdn.ResourceReplaceParamsOptionsCacheHTTPHeaders{
					Enabled: false,
					Value:   []string{"vary", "content-length", "last-modified", "connection", "accept-ranges", "content-type", "content-encoding", "etag", "cache-control", "expires", "keep-alive", "server"},
				},
				Cors: cdn.ResourceReplaceParamsOptionsCors{
					Enabled: true,
					Value:   []string{"domain.com", "domain2.com"},
					Always:  gcore.Bool(true),
				},
				CountryACL: cdn.ResourceReplaceParamsOptionsCountryACL{
					Enabled:        true,
					ExceptedValues: []string{"GB", "DE"},
					PolicyType:     "allow",
				},
				DisableCache: cdn.ResourceReplaceParamsOptionsDisableCache{
					Enabled: true,
					Value:   false,
				},
				DisableProxyForceRanges: cdn.ResourceReplaceParamsOptionsDisableProxyForceRanges{
					Enabled: true,
					Value:   true,
				},
				EdgeCacheSettings: cdn.ResourceReplaceParamsOptionsEdgeCacheSettings{
					Enabled: true,
					CustomValues: map[string]string{
						"100": "43200s",
					},
					Default: gcore.String("321669910225"),
					Value:   gcore.String("43200s"),
				},
				Fastedge: cdn.ResourceReplaceParamsOptionsFastedge{
					Enabled: true,
					OnRequestBody: cdn.ResourceReplaceParamsOptionsFastedgeOnRequestBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnRequestHeaders: cdn.ResourceReplaceParamsOptionsFastedgeOnRequestHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseBody: cdn.ResourceReplaceParamsOptionsFastedgeOnResponseBody{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
					OnResponseHeaders: cdn.ResourceReplaceParamsOptionsFastedgeOnResponseHeaders{
						AppID:            "1001",
						Enabled:          gcore.Bool(true),
						ExecuteOnEdge:    gcore.Bool(true),
						ExecuteOnShield:  gcore.Bool(false),
						InterruptOnError: gcore.Bool(true),
					},
				},
				FetchCompressed: cdn.ResourceReplaceParamsOptionsFetchCompressed{
					Enabled: true,
					Value:   false,
				},
				FollowOriginRedirect: cdn.ResourceReplaceParamsOptionsFollowOriginRedirect{
					Codes:   []int64{302, 308},
					Enabled: true,
				},
				ForceReturn: cdn.ResourceReplaceParamsOptionsForceReturn{
					Body:    "http://example.com/redirect_address",
					Code:    301,
					Enabled: true,
					TimeInterval: cdn.ResourceReplaceParamsOptionsForceReturnTimeInterval{
						EndTime:   "20:00",
						StartTime: "09:00",
						TimeZone:  gcore.String("CET"),
					},
				},
				ForwardHostHeader: cdn.ResourceReplaceParamsOptionsForwardHostHeader{
					Enabled: false,
					Value:   false,
				},
				GzipOn: cdn.ResourceReplaceParamsOptionsGzipOn{
					Enabled: true,
					Value:   true,
				},
				HostHeader: cdn.ResourceReplaceParamsOptionsHostHeader{
					Enabled: true,
					Value:   "host.com",
				},
				Http3Enabled: cdn.ResourceReplaceParamsOptionsHttp3Enabled{
					Enabled: true,
					Value:   true,
				},
				IgnoreCookie: cdn.ResourceReplaceParamsOptionsIgnoreCookie{
					Enabled: true,
					Value:   true,
				},
				IgnoreQueryString: cdn.ResourceReplaceParamsOptionsIgnoreQueryString{
					Enabled: true,
					Value:   false,
				},
				ImageStack: cdn.ResourceReplaceParamsOptionsImageStack{
					Enabled:     true,
					AvifEnabled: gcore.Bool(true),
					PngLossless: gcore.Bool(true),
					Quality:     gcore.Int(80),
					WebpEnabled: gcore.Bool(false),
				},
				IPAddressACL: cdn.ResourceReplaceParamsOptionsIPAddressACL{
					Enabled:        true,
					ExceptedValues: []string{"192.168.1.100/32"},
					PolicyType:     "deny",
				},
				LimitBandwidth: cdn.ResourceReplaceParamsOptionsLimitBandwidth{
					Enabled:   true,
					LimitType: "static",
					Buffer:    gcore.Int(200),
					Speed:     gcore.Int(100),
				},
				ProxyCacheKey: cdn.ResourceReplaceParamsOptionsProxyCacheKey{
					Enabled: true,
					Value:   "$scheme$uri",
				},
				ProxyCacheMethodsSet: cdn.ResourceReplaceParamsOptionsProxyCacheMethodsSet{
					Enabled: true,
					Value:   false,
				},
				ProxyConnectTimeout: cdn.ResourceReplaceParamsOptionsProxyConnectTimeout{
					Enabled: true,
					Value:   "4s",
				},
				ProxyReadTimeout: cdn.ResourceReplaceParamsOptionsProxyReadTimeout{
					Enabled: true,
					Value:   "10s",
				},
				QueryParamsBlacklist: cdn.ResourceReplaceParamsOptionsQueryParamsBlacklist{
					Enabled: true,
					Value:   []string{"some", "blacklisted", "query"},
				},
				QueryParamsWhitelist: cdn.ResourceReplaceParamsOptionsQueryParamsWhitelist{
					Enabled: true,
					Value:   []string{"some", "whitelisted", "query"},
				},
				QueryStringForwarding: cdn.ResourceReplaceParamsOptionsQueryStringForwarding{
					Enabled:              true,
					ForwardFromFileTypes: []string{"m3u8", "mpd"},
					ForwardToFileTypes:   []string{"ts", "mp4"},
					ForwardExceptKeys:    []string{"debug_info"},
					ForwardOnlyKeys:      []string{"auth_token", "session_id"},
				},
				RedirectHTTPToHTTPS: cdn.ResourceReplaceParamsOptionsRedirectHTTPToHTTPS{
					Enabled: true,
					Value:   true,
				},
				RedirectHTTPSToHTTP: cdn.ResourceReplaceParamsOptionsRedirectHTTPSToHTTP{
					Enabled: false,
					Value:   true,
				},
				ReferrerACL: cdn.ResourceReplaceParamsOptionsReferrerACL{
					Enabled:        true,
					ExceptedValues: []string{"example.com", "*.example.net"},
					PolicyType:     "deny",
				},
				RequestLimiter: cdn.ResourceReplaceParamsOptionsRequestLimiter{
					Enabled:  true,
					Rate:     5,
					RateUnit: "r/s",
				},
				ResponseHeadersHidingPolicy: cdn.ResourceReplaceParamsOptionsResponseHeadersHidingPolicy{
					Enabled:  true,
					Excepted: []string{"my-header"},
					Mode:     "hide",
				},
				Rewrite: cdn.ResourceReplaceParamsOptionsRewrite{
					Body:    "/(.*) /additional_path/$1",
					Enabled: true,
					Flag:    "break",
				},
				SecureKey: cdn.ResourceReplaceParamsOptionsSecureKey{
					Enabled: true,
					Key:     gcore.String("secretkey"),
					Type:    2,
				},
				Slice: cdn.ResourceReplaceParamsOptionsSlice{
					Enabled: true,
					Value:   true,
				},
				Sni: cdn.ResourceReplaceParamsOptionsSni{
					CustomHostname: "custom.example.com",
					Enabled:        true,
					SniType:        "custom",
				},
				Stale: cdn.ResourceReplaceParamsOptionsStale{
					Enabled: true,
					Value:   []string{"http_404", "http_500"},
				},
				StaticResponseHeaders: cdn.ResourceReplaceParamsOptionsStaticResponseHeaders{
					Enabled: true,
					Value: []cdn.ResourceReplaceParamsOptionsStaticResponseHeadersValue{{
						Name:   "X-Example",
						Value:  []string{"Value_1"},
						Always: gcore.Bool(true),
					}, {
						Name:   "X-Example-Multiple",
						Value:  []string{"Value_1", "Value_2", "Value_3"},
						Always: gcore.Bool(false),
					}},
				},
				StaticHeaders: cdn.ResourceReplaceParamsOptionsStaticHeaders{
					Enabled: true,
					Value: map[string]interface{}{
						"X-Example": "Value_1",
						"X-Example-Multiple": map[string]interface{}{
							"0": "Value_2",
							"1": "Value_3",
						},
					},
				},
				StaticRequestHeaders: cdn.ResourceReplaceParamsOptionsStaticRequestHeaders{
					Enabled: true,
					Value: map[string]string{
						"Header-One": "Value 1",
						"Header-Two": "Value 2",
					},
				},
				TlsVersions: cdn.ResourceReplaceParamsOptionsTlsVersions{
					Enabled: true,
					Value:   []string{"SSLv3", "TLSv1.3"},
				},
				UseDefaultLeChain: cdn.ResourceReplaceParamsOptionsUseDefaultLeChain{
					Enabled: true,
					Value:   true,
				},
				UseDns01LeChallenge: cdn.ResourceReplaceParamsOptionsUseDns01LeChallenge{
					Enabled: true,
					Value:   true,
				},
				UseRsaLeCert: cdn.ResourceReplaceParamsOptionsUseRsaLeCert{
					Enabled: true,
					Value:   true,
				},
				UserAgentACL: cdn.ResourceReplaceParamsOptionsUserAgentACL{
					Enabled:        true,
					ExceptedValues: []string{"UserAgent Value", ""},
					PolicyType:     "allow",
				},
				Waap: cdn.ResourceReplaceParamsOptionsWaap{
					Enabled: true,
					Value:   true,
				},
				Websockets: cdn.ResourceReplaceParamsOptionsWebsockets{
					Enabled: true,
					Value:   true,
				},
			},
			OriginProtocol:       cdn.ResourceReplaceParamsOriginProtocolHTTPS,
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
