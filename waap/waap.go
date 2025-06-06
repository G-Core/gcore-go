// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// WaapService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapService] method instead.
type WaapService struct {
	Options []option.RequestOption
	Domains DomainService
}

// NewWaapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWaapService(opts ...option.RequestOption) (r WaapService) {
	r = WaapService{}
	r.Options = opts
	r.Domains = NewDomainService(opts...)
	return
}

// API settings of a domain
type WaapAPIURLs struct {
	// The API URLs for a domain. If your domain has a common base URL for all API
	// paths, it can be set here
	APIURLs []string `json:"api_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIURLs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAPIURLs) RawJSON() string { return r.JSON.raw }
func (r *WaapAPIURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a WAAP domain, serving as a singular unit within the WAAP service.
//
// Each domain functions autonomously, possessing its own set of rules and
// configurations to manage web application firewall settings and behaviors.
type WaapDetailedDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDomainStatus `json:"status,required"`
	// Domain level quotas
	Quotas map[string]WaapDetailedDomainQuota `json:"quotas,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		Quotas        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDetailedDomainQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64 `json:"current,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomainQuota) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomainQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DDoS settings for a domain.
type WaapDomainDDOSSettings struct {
	// The burst threshold detects sudden rises in traffic. If it is met and the number
	// of requests is at least five times the last 2-second interval, DDoS protection
	// will activate. Default is 1000.
	BurstThreshold int64 `json:"burst_threshold"`
	// The global threshold is responsible for identifying DDoS attacks with a slow
	// rise in traffic. If the threshold is met and the current number of requests is
	// at least double that of the previous 10-second window, DDoS protection will
	// activate. Default is 5000.
	GlobalThreshold int64 `json:"global_threshold"`
	// The sub-second threshold protects WAAP servers against attacks from traffic
	// bursts. When this threshold is reached, the DDoS mode will activate on the
	// affected WAAP server, not the whole WAAP cluster. Default is 50.
	SubSecondThreshold int64 `json:"sub_second_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BurstThreshold     respjson.Field
		GlobalThreshold    respjson.Field
		SubSecondThreshold respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainDDOSSettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainDDOSSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for a domain.
type WaapDomainSettings struct {
	// API settings of a domain
	API WaapAPIURLs `json:"api,required"`
	// DDoS settings for a domain.
	DDOS WaapDomainDDOSSettings `json:"ddos,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API         respjson.Field
		DDOS        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainSettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses a domain can have
type WaapDomainStatus string

const (
	WaapDomainStatusActive  WaapDomainStatus = "active"
	WaapDomainStatusBypass  WaapDomainStatus = "bypass"
	WaapDomainStatusMonitor WaapDomainStatus = "monitor"
	WaapDomainStatusLocked  WaapDomainStatus = "locked"
)

// Represents a WAAP domain when getting a list of domains.
type WaapSummaryDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDomainStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSummaryDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapSummaryDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
