// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

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

// DomainAPIDiscoverySettingService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIDiscoverySettingService] method instead.
type DomainAPIDiscoverySettingService struct {
	Options []option.RequestOption
}

// NewDomainAPIDiscoverySettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoverySettingService(opts ...option.RequestOption) (r DomainAPIDiscoverySettingService) {
	r = DomainAPIDiscoverySettingService{}
	r.Options = opts
	return
}

// Update the API discovery settings for a domain
func (r *DomainAPIDiscoverySettingService) Update(ctx context.Context, domainID int64, body DomainAPIDiscoverySettingUpdateParams, opts ...option.RequestOption) (res *WaapAPIDiscoverySettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve the API discovery settings for a domain
func (r *DomainAPIDiscoverySettingService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapAPIDiscoverySettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response model for the API discovery settings
type WaapAPIDiscoverySettings struct {
	// The URL of the API description file. This will be periodically scanned if
	// `descriptionFileScanEnabled` is enabled. Supported formats are YAML and JSON,
	// and it must adhere to OpenAPI versions 2, 3, or 3.1.
	DescriptionFileLocation string `json:"descriptionFileLocation,nullable"`
	// Indicates if periodic scan of the description file is enabled
	DescriptionFileScanEnabled bool `json:"descriptionFileScanEnabled,nullable"`
	// The interval in hours for scanning the description file
	DescriptionFileScanIntervalHours int64 `json:"descriptionFileScanIntervalHours,nullable"`
	// Indicates if traffic scan is enabled. Traffic scan is used to discover
	// undocumented APIs
	TrafficScanEnabled bool `json:"trafficScanEnabled,nullable"`
	// The interval in hours for scanning the traffic
	TrafficScanIntervalHours int64 `json:"trafficScanIntervalHours,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DescriptionFileLocation          respjson.Field
		DescriptionFileScanEnabled       respjson.Field
		DescriptionFileScanIntervalHours respjson.Field
		TrafficScanEnabled               respjson.Field
		TrafficScanIntervalHours         respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAPIDiscoverySettings) RawJSON() string { return r.JSON.raw }
func (r *WaapAPIDiscoverySettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainAPIDiscoverySettingUpdateParams struct {
	// The URL of the API description file. This will be periodically scanned if
	// `descriptionFileScanEnabled` is enabled. Supported formats are YAML and JSON,
	// and it must adhere to OpenAPI versions 2, 3, or 3.1.
	DescriptionFileLocation param.Opt[string] `json:"descriptionFileLocation,omitzero"`
	// Indicates if periodic scan of the description file is enabled
	DescriptionFileScanEnabled param.Opt[bool] `json:"descriptionFileScanEnabled,omitzero"`
	// The interval in hours for scanning the description file
	DescriptionFileScanIntervalHours param.Opt[int64] `json:"descriptionFileScanIntervalHours,omitzero"`
	// Indicates if traffic scan is enabled
	TrafficScanEnabled param.Opt[bool] `json:"trafficScanEnabled,omitzero"`
	// The interval in hours for scanning the traffic
	TrafficScanIntervalHours param.Opt[int64] `json:"trafficScanIntervalHours,omitzero"`
	paramObj
}

func (r DomainAPIDiscoverySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIDiscoverySettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIDiscoverySettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
