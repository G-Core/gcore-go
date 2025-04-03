// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainAPIDiscoveryScanResultService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAPIDiscoveryScanResultService] method instead.
type WaapV1DomainAPIDiscoveryScanResultService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainAPIDiscoveryScanResultService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWaapV1DomainAPIDiscoveryScanResultService(opts ...option.RequestOption) (r *WaapV1DomainAPIDiscoveryScanResultService) {
	r = &WaapV1DomainAPIDiscoveryScanResultService{}
	r.Options = opts
	return
}

// Get Scan Result
func (r *WaapV1DomainAPIDiscoveryScanResultService) Get(ctx context.Context, domainID int64, scanID string, opts ...option.RequestOption) (res *APIScanResult, err error) {
	opts = append(r.Options[:], opts...)
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan-results/%s", domainID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Scan Results
func (r *WaapV1DomainAPIDiscoveryScanResultService) List(ctx context.Context, domainID int64, query WaapV1DomainAPIDiscoveryScanResultListParams, opts ...option.RequestOption) (res *WaapV1DomainAPIDiscoveryScanResultListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan-results", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The result of a scan
type APIScanResult struct {
	// The scan ID
	ID string `json:"id,required" format:"uuid"`
	// The date and time the scan ended
	EndTime time.Time `json:"end_time,required,nullable" format:"date-time"`
	// The message associated with the scan
	Message string `json:"message,required"`
	// The date and time the scan started
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// The different statuses a task result can have
	Status TaskResultStatus `json:"status,required"`
	// The different types of scans that can be performed
	Type APIScanType       `json:"type,required"`
	JSON apiScanResultJSON `json:"-"`
}

// apiScanResultJSON contains the JSON metadata for the struct [APIScanResult]
type apiScanResultJSON struct {
	ID          apijson.Field
	EndTime     apijson.Field
	Message     apijson.Field
	StartTime   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIScanResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiScanResultJSON) RawJSON() string {
	return r.raw
}

// The different types of scans that can be performed
type APIScanType string

const (
	APIScanTypeTrafficScan            APIScanType = "TRAFFIC_SCAN"
	APIScanTypeAPIDescriptionFileScan APIScanType = "API_DESCRIPTION_FILE_SCAN"
)

func (r APIScanType) IsKnown() bool {
	switch r {
	case APIScanTypeTrafficScan, APIScanTypeAPIDescriptionFileScan:
		return true
	}
	return false
}

// The different statuses a task result can have
type TaskResultStatus string

const (
	TaskResultStatusSuccess    TaskResultStatus = "SUCCESS"
	TaskResultStatusFailure    TaskResultStatus = "FAILURE"
	TaskResultStatusInProgress TaskResultStatus = "IN_PROGRESS"
)

func (r TaskResultStatus) IsKnown() bool {
	switch r {
	case TaskResultStatusSuccess, TaskResultStatusFailure, TaskResultStatusInProgress:
		return true
	}
	return false
}

type WaapV1DomainAPIDiscoveryScanResultListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []APIScanResult                                    `json:"results,required"`
	JSON    waapV1DomainAPIDiscoveryScanResultListResponseJSON `json:"-"`
}

// waapV1DomainAPIDiscoveryScanResultListResponseJSON contains the JSON metadata
// for the struct [WaapV1DomainAPIDiscoveryScanResultListResponse]
type waapV1DomainAPIDiscoveryScanResultListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainAPIDiscoveryScanResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainAPIDiscoveryScanResultListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAPIDiscoveryScanResultListParams struct {
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by the message of the scan. Supports '\*' as a wildcard character
	Message param.Field[string] `query:"message"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainAPIDiscoveryScanResultListParamsOrdering] `query:"ordering"`
	// The different statuses a task result can have
	Status param.Field[TaskResultStatus] `query:"status"`
	// The different types of scans that can be performed
	Type param.Field[APIScanType] `query:"type"`
}

// URLQuery serializes [WaapV1DomainAPIDiscoveryScanResultListParams]'s query
// parameters as `url.Values`.
func (r WaapV1DomainAPIDiscoveryScanResultListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainAPIDiscoveryScanResultListParamsOrdering string

const (
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingID             WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "id"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingType           WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "type"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingStartTime      WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "start_time"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingEndTime        WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "end_time"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingStatus         WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "status"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMessage        WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "message"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusID        WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-id"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusType      WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-type"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusStartTime WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-start_time"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusEndTime   WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-end_time"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusStatus    WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-status"
	WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusMessage   WaapV1DomainAPIDiscoveryScanResultListParamsOrdering = "-message"
)

func (r WaapV1DomainAPIDiscoveryScanResultListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainAPIDiscoveryScanResultListParamsOrderingID, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingType, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingStartTime, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingEndTime, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingStatus, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMessage, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusID, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusType, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusStartTime, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusEndTime, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusStatus, WaapV1DomainAPIDiscoveryScanResultListParamsOrderingMinusMessage:
		return true
	}
	return false
}
