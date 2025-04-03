// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainAPIDiscoverySettingService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAPIDiscoverySettingService] method instead.
type WaapV1DomainAPIDiscoverySettingService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainAPIDiscoverySettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainAPIDiscoverySettingService(opts ...option.RequestOption) (r *WaapV1DomainAPIDiscoverySettingService) {
	r = &WaapV1DomainAPIDiscoverySettingService{}
	r.Options = opts
	return
}

// Retrieve the API discovery settings for a domain
func (r *WaapV1DomainAPIDiscoverySettingService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *APIDiscoverySettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the API discovery settings for a domain
func (r *WaapV1DomainAPIDiscoverySettingService) Update(ctx context.Context, domainID int64, body WaapV1DomainAPIDiscoverySettingUpdateParams, opts ...option.RequestOption) (res *APIDiscoverySettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response model for the API discovery settings
type APIDiscoverySettings struct {
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
	TrafficScanIntervalHours int64                    `json:"trafficScanIntervalHours,nullable"`
	JSON                     apiDiscoverySettingsJSON `json:"-"`
}

// apiDiscoverySettingsJSON contains the JSON metadata for the struct
// [APIDiscoverySettings]
type apiDiscoverySettingsJSON struct {
	DescriptionFileLocation          apijson.Field
	DescriptionFileScanEnabled       apijson.Field
	DescriptionFileScanIntervalHours apijson.Field
	TrafficScanEnabled               apijson.Field
	TrafficScanIntervalHours         apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *APIDiscoverySettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiDiscoverySettingsJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAPIDiscoverySettingUpdateParams struct {
	// The URL of the API description file. This will be periodically scanned if
	// `descriptionFileScanEnabled` is enabled. Supported formats are YAML and JSON,
	// and it must adhere to OpenAPI versions 2, 3, or 3.1.
	DescriptionFileLocation param.Field[string] `json:"descriptionFileLocation"`
	// Indicates if periodic scan of the description file is enabled
	DescriptionFileScanEnabled param.Field[bool] `json:"descriptionFileScanEnabled"`
	// The interval in hours for scanning the description file
	DescriptionFileScanIntervalHours param.Field[int64] `json:"descriptionFileScanIntervalHours"`
	// Indicates if traffic scan is enabled
	TrafficScanEnabled param.Field[bool] `json:"trafficScanEnabled"`
	// The interval in hours for scanning the traffic
	TrafficScanIntervalHours param.Field[int64] `json:"trafficScanIntervalHours"`
}

func (r WaapV1DomainAPIDiscoverySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
