// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainAPIDiscoveryScanResultService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIDiscoveryScanResultService] method instead.
type DomainAPIDiscoveryScanResultService struct {
	Options []option.RequestOption
}

// NewDomainAPIDiscoveryScanResultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoveryScanResultService(opts ...option.RequestOption) (r DomainAPIDiscoveryScanResultService) {
	r = DomainAPIDiscoveryScanResultService{}
	r.Options = opts
	return
}

// Get Scan Results
func (r *DomainAPIDiscoveryScanResultService) List(ctx context.Context, domainID int64, query DomainAPIDiscoveryScanResultListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapAPIScanResult], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *DomainAPIDiscoveryScanResultService) ListAutoPaging(ctx context.Context, domainID int64, query DomainAPIDiscoveryScanResultListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapAPIScanResult] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Get Scan Result
func (r *DomainAPIDiscoveryScanResultService) Get(ctx context.Context, scanID string, query DomainAPIDiscoveryScanResultGetParams, opts ...option.RequestOption) (res *WaapAPIScanResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan-results/%s", query.DomainID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The result of a scan
type WaapAPIScanResult struct {
	// The scan ID
	ID string `json:"id" api:"required" format:"uuid"`
	// The date and time the scan ended
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// The message associated with the scan
	Message string `json:"message" api:"required"`
	// The date and time the scan started
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// The status of the scan
	//
	// Any of "SUCCESS", "FAILURE", "IN_PROGRESS".
	Status WaapAPIScanResultStatus `json:"status" api:"required"`
	// The type of scan
	//
	// Any of "TRAFFIC_SCAN", "API_DESCRIPTION_FILE_SCAN".
	Type WaapAPIScanResultType `json:"type" api:"required"`
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

type DomainAPIDiscoveryScanResultListParams struct {
	// Filter by the message of the scan. Supports '\*' as a wildcard character
	Message param.Opt[string] `query:"message,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by the status of the scan
	//
	// Any of "SUCCESS", "FAILURE", "IN_PROGRESS".
	Status DomainAPIDiscoveryScanResultListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by the path of the scan type
	//
	// Any of "TRAFFIC_SCAN", "API_DESCRIPTION_FILE_SCAN".
	Type DomainAPIDiscoveryScanResultListParamsType `query:"type,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "type", "start_time", "end_time", "status", "message", "-id",
	// "-type", "-start_time", "-end_time", "-status", "-message".
	Ordering DomainAPIDiscoveryScanResultListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAPIDiscoveryScanResultListParams]'s query parameters
// as `url.Values`.
func (r DomainAPIDiscoveryScanResultListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainAPIDiscoveryScanResultListParamsOrdering string

const (
	DomainAPIDiscoveryScanResultListParamsOrderingID             DomainAPIDiscoveryScanResultListParamsOrdering = "id"
	DomainAPIDiscoveryScanResultListParamsOrderingType           DomainAPIDiscoveryScanResultListParamsOrdering = "type"
	DomainAPIDiscoveryScanResultListParamsOrderingStartTime      DomainAPIDiscoveryScanResultListParamsOrdering = "start_time"
	DomainAPIDiscoveryScanResultListParamsOrderingEndTime        DomainAPIDiscoveryScanResultListParamsOrdering = "end_time"
	DomainAPIDiscoveryScanResultListParamsOrderingStatus         DomainAPIDiscoveryScanResultListParamsOrdering = "status"
	DomainAPIDiscoveryScanResultListParamsOrderingMessage        DomainAPIDiscoveryScanResultListParamsOrdering = "message"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusID        DomainAPIDiscoveryScanResultListParamsOrdering = "-id"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusType      DomainAPIDiscoveryScanResultListParamsOrdering = "-type"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusStartTime DomainAPIDiscoveryScanResultListParamsOrdering = "-start_time"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusEndTime   DomainAPIDiscoveryScanResultListParamsOrdering = "-end_time"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusStatus    DomainAPIDiscoveryScanResultListParamsOrdering = "-status"
	DomainAPIDiscoveryScanResultListParamsOrderingMinusMessage   DomainAPIDiscoveryScanResultListParamsOrdering = "-message"
)

// Filter by the status of the scan
type DomainAPIDiscoveryScanResultListParamsStatus string

const (
	DomainAPIDiscoveryScanResultListParamsStatusSuccess    DomainAPIDiscoveryScanResultListParamsStatus = "SUCCESS"
	DomainAPIDiscoveryScanResultListParamsStatusFailure    DomainAPIDiscoveryScanResultListParamsStatus = "FAILURE"
	DomainAPIDiscoveryScanResultListParamsStatusInProgress DomainAPIDiscoveryScanResultListParamsStatus = "IN_PROGRESS"
)

// Filter by the path of the scan type
type DomainAPIDiscoveryScanResultListParamsType string

const (
	DomainAPIDiscoveryScanResultListParamsTypeTrafficScan            DomainAPIDiscoveryScanResultListParamsType = "TRAFFIC_SCAN"
	DomainAPIDiscoveryScanResultListParamsTypeAPIDescriptionFileScan DomainAPIDiscoveryScanResultListParamsType = "API_DESCRIPTION_FILE_SCAN"
)

type DomainAPIDiscoveryScanResultGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	paramObj
}
