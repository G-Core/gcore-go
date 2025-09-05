// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
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
	Options []option.RequestOption
}

// NewDomainAPIDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoveryService(opts ...option.RequestOption) (r DomainAPIDiscoveryService) {
	r = DomainAPIDiscoveryService{}
	r.Options = opts
	return
}

// Get Scan Result
func (r *DomainAPIDiscoveryService) GetScanResult(ctx context.Context, scanID string, query DomainAPIDiscoveryGetScanResultParams, opts ...option.RequestOption) (res *WaapAPIScanResult, err error) {
	opts = append(r.Options[:], opts...)
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan-results/%s", query.DomainID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the API discovery settings for a domain
func (r *DomainAPIDiscoveryService) GetSettings(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapAPIDiscoverySettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Scan Results
func (r *DomainAPIDiscoveryService) ListScanResults(ctx context.Context, domainID int64, query DomainAPIDiscoveryListScanResultsParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapAPIScanResult], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan-results", domainID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get Scan Results
func (r *DomainAPIDiscoveryService) ListScanResultsAutoPaging(ctx context.Context, domainID int64, query DomainAPIDiscoveryListScanResultsParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapAPIScanResult] {
	return pagination.NewOffsetPageAutoPager(r.ListScanResults(ctx, domainID, query, opts...))
}

// Scan an API description file hosted online. The file must be in YAML or JSON
// format and adhere to the OpenAPI specification. The location of the API
// description file should be specified in the API discovery settings.
func (r *DomainAPIDiscoveryService) ScanOpenAPI(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapTaskID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update the API discovery settings for a domain
func (r *DomainAPIDiscoveryService) UpdateSettings(ctx context.Context, domainID int64, body DomainAPIDiscoveryUpdateSettingsParams, opts ...option.RequestOption) (res *WaapAPIDiscoverySettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// An API description file must adhere to the OpenAPI specification and be written
// in YAML or JSON format. The file name should be provided as the value for the
// `file_name` parameter. The contents of the file must be base64 encoded and
// supplied as the value for the `file_data` parameter.
func (r *DomainAPIDiscoveryService) UploadOpenAPI(ctx context.Context, domainID int64, body DomainAPIDiscoveryUploadOpenAPIParams, opts ...option.RequestOption) (res *WaapTaskID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/upload", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// The result of a scan
type WaapAPIScanResult struct {
	// The scan ID
	ID string `json:"id,required" format:"uuid"`
	// The date and time the scan ended
	EndTime time.Time `json:"end_time,required" format:"date-time"`
	// The message associated with the scan
	Message string `json:"message,required"`
	// The date and time the scan started
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// The status of the scan
	//
	// Any of "SUCCESS", "FAILURE", "IN_PROGRESS".
	Status WaapAPIScanResultStatus `json:"status,required"`
	// The type of scan
	//
	// Any of "TRAFFIC_SCAN", "API_DESCRIPTION_FILE_SCAN".
	Type WaapAPIScanResultType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EndTime     respjson.Field
		Message     respjson.Field
		StartTime   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAPIScanResult) RawJSON() string { return r.JSON.raw }
func (r *WaapAPIScanResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the scan
type WaapAPIScanResultStatus string

const (
	WaapAPIScanResultStatusSuccess    WaapAPIScanResultStatus = "SUCCESS"
	WaapAPIScanResultStatusFailure    WaapAPIScanResultStatus = "FAILURE"
	WaapAPIScanResultStatusInProgress WaapAPIScanResultStatus = "IN_PROGRESS"
)

// The type of scan
type WaapAPIScanResultType string

const (
	WaapAPIScanResultTypeTrafficScan            WaapAPIScanResultType = "TRAFFIC_SCAN"
	WaapAPIScanResultTypeAPIDescriptionFileScan WaapAPIScanResultType = "API_DESCRIPTION_FILE_SCAN"
)

// Response model for the task result ID
type WaapTaskID struct {
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
func (r WaapTaskID) RawJSON() string { return r.JSON.raw }
func (r *WaapTaskID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainAPIDiscoveryGetScanResultParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainAPIDiscoveryListScanResultsParams struct {
	// Filter by the message of the scan. Supports '\*' as a wildcard character
	Message param.Opt[string] `query:"message,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The different statuses a task result can have
	//
	// Any of "SUCCESS", "FAILURE", "IN_PROGRESS".
	Status DomainAPIDiscoveryListScanResultsParamsStatus `query:"status,omitzero" json:"-"`
	// The different types of scans that can be performed
	//
	// Any of "TRAFFIC_SCAN", "API_DESCRIPTION_FILE_SCAN".
	Type DomainAPIDiscoveryListScanResultsParamsType `query:"type,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "type", "start_time", "end_time", "status", "message", "-id",
	// "-type", "-start_time", "-end_time", "-status", "-message".
	Ordering DomainAPIDiscoveryListScanResultsParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAPIDiscoveryListScanResultsParams]'s query parameters
// as `url.Values`.
func (r DomainAPIDiscoveryListScanResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainAPIDiscoveryListScanResultsParamsOrdering string

const (
	DomainAPIDiscoveryListScanResultsParamsOrderingID             DomainAPIDiscoveryListScanResultsParamsOrdering = "id"
	DomainAPIDiscoveryListScanResultsParamsOrderingType           DomainAPIDiscoveryListScanResultsParamsOrdering = "type"
	DomainAPIDiscoveryListScanResultsParamsOrderingStartTime      DomainAPIDiscoveryListScanResultsParamsOrdering = "start_time"
	DomainAPIDiscoveryListScanResultsParamsOrderingEndTime        DomainAPIDiscoveryListScanResultsParamsOrdering = "end_time"
	DomainAPIDiscoveryListScanResultsParamsOrderingStatus         DomainAPIDiscoveryListScanResultsParamsOrdering = "status"
	DomainAPIDiscoveryListScanResultsParamsOrderingMessage        DomainAPIDiscoveryListScanResultsParamsOrdering = "message"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusID        DomainAPIDiscoveryListScanResultsParamsOrdering = "-id"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusType      DomainAPIDiscoveryListScanResultsParamsOrdering = "-type"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusStartTime DomainAPIDiscoveryListScanResultsParamsOrdering = "-start_time"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusEndTime   DomainAPIDiscoveryListScanResultsParamsOrdering = "-end_time"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusStatus    DomainAPIDiscoveryListScanResultsParamsOrdering = "-status"
	DomainAPIDiscoveryListScanResultsParamsOrderingMinusMessage   DomainAPIDiscoveryListScanResultsParamsOrdering = "-message"
)

// The different statuses a task result can have
type DomainAPIDiscoveryListScanResultsParamsStatus string

const (
	DomainAPIDiscoveryListScanResultsParamsStatusSuccess    DomainAPIDiscoveryListScanResultsParamsStatus = "SUCCESS"
	DomainAPIDiscoveryListScanResultsParamsStatusFailure    DomainAPIDiscoveryListScanResultsParamsStatus = "FAILURE"
	DomainAPIDiscoveryListScanResultsParamsStatusInProgress DomainAPIDiscoveryListScanResultsParamsStatus = "IN_PROGRESS"
)

// The different types of scans that can be performed
type DomainAPIDiscoveryListScanResultsParamsType string

const (
	DomainAPIDiscoveryListScanResultsParamsTypeTrafficScan            DomainAPIDiscoveryListScanResultsParamsType = "TRAFFIC_SCAN"
	DomainAPIDiscoveryListScanResultsParamsTypeAPIDescriptionFileScan DomainAPIDiscoveryListScanResultsParamsType = "API_DESCRIPTION_FILE_SCAN"
)

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
