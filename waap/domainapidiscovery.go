// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainAPIDiscoveryService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIDiscoveryService] method instead.
type DomainAPIDiscoveryService struct {
	Options     []option.RequestOption
	ScanResults DomainAPIDiscoveryScanResultService
}

// NewDomainAPIDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoveryService(opts ...option.RequestOption) (r DomainAPIDiscoveryService) {
	r = DomainAPIDiscoveryService{}
	r.Options = opts
	r.ScanResults = NewDomainAPIDiscoveryScanResultService(opts...)
	return
}

// Retrieve the API discovery settings for a domain
func (r *DomainAPIDiscoveryService) GetSettings(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *DomainAPIDiscoveryGetSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Scan an API description file hosted online. The file must be in YAML or JSON
// format and adhere to the OpenAPI specification. The location of the API
// description file should be specified in the API discovery settings.
func (r *DomainAPIDiscoveryService) ScanOpenAPI(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *DomainAPIDiscoveryScanOpenAPIResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update the API discovery settings for a domain
func (r *DomainAPIDiscoveryService) UpdateSettings(ctx context.Context, domainID int64, body DomainAPIDiscoveryUpdateSettingsParams, opts ...option.RequestOption) (res *DomainAPIDiscoveryUpdateSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// An API description file must adhere to the OpenAPI specification and be written
// in YAML or JSON format. The file name should be provided as the value for the
// `file_name` parameter. The contents of the file must be base64 encoded and
// supplied as the value for the `file_data` parameter.
func (r *DomainAPIDiscoveryService) UploadOpenAPI(ctx context.Context, domainID int64, body DomainAPIDiscoveryUploadOpenAPIParams, opts ...option.RequestOption) (res *DomainAPIDiscoveryUploadOpenAPIResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/upload", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model for the API discovery settings
type DomainAPIDiscoveryGetSettingsResponse struct {
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
func (r DomainAPIDiscoveryGetSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainAPIDiscoveryGetSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the task result ID
type DomainAPIDiscoveryScanOpenAPIResponse struct {
	// The task ID
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainAPIDiscoveryScanOpenAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainAPIDiscoveryScanOpenAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the API discovery settings
type DomainAPIDiscoveryUpdateSettingsResponse struct {
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
func (r DomainAPIDiscoveryUpdateSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainAPIDiscoveryUpdateSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the task result ID
type DomainAPIDiscoveryUploadOpenAPIResponse struct {
	// The task ID
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainAPIDiscoveryUploadOpenAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainAPIDiscoveryUploadOpenAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainAPIDiscoveryUpdateSettingsParams struct {
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

func (r DomainAPIDiscoveryUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIDiscoveryUpdateSettingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIDiscoveryUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainAPIDiscoveryUploadOpenAPIParams struct {
	// Base64 representation of the description file. Supported formats are YAML and
	// JSON, and it must adhere to OpenAPI versions 2, 3, or 3.1.
	FileData string `json:"file_data,required"`
	// The name of the file
	FileName string `json:"file_name,required"`
	paramObj
}

func (r DomainAPIDiscoveryUploadOpenAPIParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIDiscoveryUploadOpenAPIParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIDiscoveryUploadOpenAPIParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
