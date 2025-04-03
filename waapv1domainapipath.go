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

// WaapV1DomainAPIPathService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAPIPathService] method instead.
type WaapV1DomainAPIPathService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainAPIPathService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1DomainAPIPathService(opts ...option.RequestOption) (r *WaapV1DomainAPIPathService) {
	r = &WaapV1DomainAPIPathService{}
	r.Options = opts
	return
}

// Create an API path for a domain
func (r *WaapV1DomainAPIPathService) New(ctx context.Context, domainID int64, body WaapV1DomainAPIPathNewParams, opts ...option.RequestOption) (res *PathResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific API path for a domain
func (r *WaapV1DomainAPIPathService) Get(ctx context.Context, domainID int64, pathID string, opts ...option.RequestOption) (res *PathResponse, err error) {
	opts = append(r.Options[:], opts...)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", domainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific API path for a domain
func (r *WaapV1DomainAPIPathService) Update(ctx context.Context, domainID int64, pathID string, body WaapV1DomainAPIPathUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", domainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of API paths for a specific domain
func (r *WaapV1DomainAPIPathService) List(ctx context.Context, domainID int64, query WaapV1DomainAPIPathListParams, opts ...option.RequestOption) (res *WaapV1DomainAPIPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific API path for a domain
func (r *WaapV1DomainAPIPathService) Delete(ctx context.Context, domainID int64, pathID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", domainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// The different statuses an API path can have
type APIPathStatus string

const (
	APIPathStatusConfirmedAPI APIPathStatus = "CONFIRMED_API"
	APIPathStatusPotentialAPI APIPathStatus = "POTENTIAL_API"
	APIPathStatusNotAPI       APIPathStatus = "NOT_API"
	APIPathStatusDelistedAPI  APIPathStatus = "DELISTED_API"
)

func (r APIPathStatus) IsKnown() bool {
	switch r {
	case APIPathStatusConfirmedAPI, APIPathStatusPotentialAPI, APIPathStatusNotAPI, APIPathStatusDelistedAPI:
		return true
	}
	return false
}

// The different HTTP schemes an API path can have
type HTTPScheme string

const (
	HTTPSchemeHTTP  HTTPScheme = "HTTP"
	HTTPSchemeHTTPS HTTPScheme = "HTTPS"
)

func (r HTTPScheme) IsKnown() bool {
	switch r {
	case HTTPSchemeHTTP, HTTPSchemeHTTPS:
		return true
	}
	return false
}

// The different methods an API path can have
type Method string

const (
	MethodGet     Method = "GET"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodTrace   Method = "TRACE"
	MethodHead    Method = "HEAD"
	MethodOptions Method = "OPTIONS"
)

func (r Method) IsKnown() bool {
	switch r {
	case MethodGet, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodTrace, MethodHead, MethodOptions:
		return true
	}
	return false
}

// Response model for the API path
type PathResponse struct {
	// The path ID
	ID string `json:"id,required" format:"uuid"`
	// An array of api groups associated with the API path
	APIGroups []string `json:"api_groups,required"`
	// The API version
	APIVersion string `json:"api_version,required"`
	// The date and time in ISO 8601 format the API path was first detected.
	FirstDetected time.Time `json:"first_detected,required" format:"date-time"`
	// The different HTTP schemes an API path can have
	HTTPScheme HTTPScheme `json:"http_scheme,required"`
	// The date and time in ISO 8601 format the API path was last detected.
	LastDetected time.Time `json:"last_detected,required" format:"date-time"`
	// The different methods an API path can have
	Method Method `json:"method,required"`
	// The API path, locations that are saved for resource IDs will be put in curly
	// brackets
	Path string `json:"path,required"`
	// The number of requests for this path in the last 24 hours
	RequestCount int64 `json:"request_count,required"`
	// The different sources an API path can have
	Source Source `json:"source,required"`
	// The different statuses an API path can have
	Status APIPathStatus `json:"status,required"`
	// An array of tags associated with the API path
	Tags []string         `json:"tags,required"`
	JSON pathResponseJSON `json:"-"`
}

// pathResponseJSON contains the JSON metadata for the struct [PathResponse]
type pathResponseJSON struct {
	ID            apijson.Field
	APIGroups     apijson.Field
	APIVersion    apijson.Field
	FirstDetected apijson.Field
	HTTPScheme    apijson.Field
	LastDetected  apijson.Field
	Method        apijson.Field
	Path          apijson.Field
	RequestCount  apijson.Field
	Source        apijson.Field
	Status        apijson.Field
	Tags          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pathResponseJSON) RawJSON() string {
	return r.raw
}

// The different sources an API path can have
type Source string

const (
	SourceAPIDescriptionFile Source = "API_DESCRIPTION_FILE"
	SourceTrafficScan        Source = "TRAFFIC_SCAN"
	SourceUserDefined        Source = "USER_DEFINED"
)

func (r Source) IsKnown() bool {
	switch r {
	case SourceAPIDescriptionFile, SourceTrafficScan, SourceUserDefined:
		return true
	}
	return false
}

type WaapV1DomainAPIPathListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []PathResponse                      `json:"results,required"`
	JSON    waapV1DomainAPIPathListResponseJSON `json:"-"`
}

// waapV1DomainAPIPathListResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainAPIPathListResponse]
type waapV1DomainAPIPathListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainAPIPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainAPIPathListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAPIPathNewParams struct {
	// The different HTTP schemes an API path can have
	HTTPScheme param.Field[HTTPScheme] `json:"http_scheme,required"`
	// The different methods an API path can have
	Method param.Field[Method] `json:"method,required"`
	// The API path, locations that are saved for resource IDs will be put in curly
	// brackets
	Path param.Field[string] `json:"path,required"`
	// An array of api groups associated with the API path
	APIGroups param.Field[[]string] `json:"api_groups"`
	// The API version
	APIVersion param.Field[string] `json:"api_version"`
	// An array of tags associated with the API path
	Tags param.Field[[]string] `json:"tags"`
}

func (r WaapV1DomainAPIPathNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainAPIPathUpdateParams struct {
	// An array of api groups associated with the API path
	APIGroups param.Field[[]string] `json:"api_groups"`
	// The updated API path. When updating the path, variables can be renamed, path
	// parts can be converted to variables and vice versa.
	Path param.Field[string] `json:"path"`
	// The different statuses an API path can have
	Status param.Field[APIPathStatus] `json:"status"`
	// An array of tags associated with the API path
	Tags param.Field[[]string] `json:"tags"`
}

func (r WaapV1DomainAPIPathUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainAPIPathListParams struct {
	// Filter by the API group associated with the API path
	APIGroup param.Field[string] `query:"api_group"`
	// Filter by the API version
	APIVersion param.Field[string] `query:"api_version"`
	// The different HTTP schemes an API path can have
	HTTPScheme param.Field[HTTPScheme] `query:"http_scheme"`
	// Filter by the path ID
	IDs param.Field[[]string] `query:"ids" format:"uuid"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// The different methods an API path can have
	Method param.Field[Method] `query:"method"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainAPIPathListParamsOrdering] `query:"ordering"`
	// Filter by the path. Supports '\*' as a wildcard character
	Path param.Field[string] `query:"path"`
	// The different sources an API path can have
	Source param.Field[Source] `query:"source"`
	// Filter by the status of the discovered API path
	Status param.Field[[]APIPathStatus] `query:"status"`
}

// URLQuery serializes [WaapV1DomainAPIPathListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainAPIPathListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainAPIPathListParamsOrdering string

const (
	WaapV1DomainAPIPathListParamsOrderingID                 WaapV1DomainAPIPathListParamsOrdering = "id"
	WaapV1DomainAPIPathListParamsOrderingPath               WaapV1DomainAPIPathListParamsOrdering = "path"
	WaapV1DomainAPIPathListParamsOrderingMethod             WaapV1DomainAPIPathListParamsOrdering = "method"
	WaapV1DomainAPIPathListParamsOrderingAPIVersion         WaapV1DomainAPIPathListParamsOrdering = "api_version"
	WaapV1DomainAPIPathListParamsOrderingHTTPScheme         WaapV1DomainAPIPathListParamsOrdering = "http_scheme"
	WaapV1DomainAPIPathListParamsOrderingFirstDetected      WaapV1DomainAPIPathListParamsOrdering = "first_detected"
	WaapV1DomainAPIPathListParamsOrderingLastDetected       WaapV1DomainAPIPathListParamsOrdering = "last_detected"
	WaapV1DomainAPIPathListParamsOrderingStatus             WaapV1DomainAPIPathListParamsOrdering = "status"
	WaapV1DomainAPIPathListParamsOrderingSource             WaapV1DomainAPIPathListParamsOrdering = "source"
	WaapV1DomainAPIPathListParamsOrderingMinusID            WaapV1DomainAPIPathListParamsOrdering = "-id"
	WaapV1DomainAPIPathListParamsOrderingMinusPath          WaapV1DomainAPIPathListParamsOrdering = "-path"
	WaapV1DomainAPIPathListParamsOrderingMinusMethod        WaapV1DomainAPIPathListParamsOrdering = "-method"
	WaapV1DomainAPIPathListParamsOrderingMinusAPIVersion    WaapV1DomainAPIPathListParamsOrdering = "-api_version"
	WaapV1DomainAPIPathListParamsOrderingMinusHTTPScheme    WaapV1DomainAPIPathListParamsOrdering = "-http_scheme"
	WaapV1DomainAPIPathListParamsOrderingMinusFirstDetected WaapV1DomainAPIPathListParamsOrdering = "-first_detected"
	WaapV1DomainAPIPathListParamsOrderingMinusLastDetected  WaapV1DomainAPIPathListParamsOrdering = "-last_detected"
	WaapV1DomainAPIPathListParamsOrderingMinusStatus        WaapV1DomainAPIPathListParamsOrdering = "-status"
	WaapV1DomainAPIPathListParamsOrderingMinusSource        WaapV1DomainAPIPathListParamsOrdering = "-source"
)

func (r WaapV1DomainAPIPathListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainAPIPathListParamsOrderingID, WaapV1DomainAPIPathListParamsOrderingPath, WaapV1DomainAPIPathListParamsOrderingMethod, WaapV1DomainAPIPathListParamsOrderingAPIVersion, WaapV1DomainAPIPathListParamsOrderingHTTPScheme, WaapV1DomainAPIPathListParamsOrderingFirstDetected, WaapV1DomainAPIPathListParamsOrderingLastDetected, WaapV1DomainAPIPathListParamsOrderingStatus, WaapV1DomainAPIPathListParamsOrderingSource, WaapV1DomainAPIPathListParamsOrderingMinusID, WaapV1DomainAPIPathListParamsOrderingMinusPath, WaapV1DomainAPIPathListParamsOrderingMinusMethod, WaapV1DomainAPIPathListParamsOrderingMinusAPIVersion, WaapV1DomainAPIPathListParamsOrderingMinusHTTPScheme, WaapV1DomainAPIPathListParamsOrderingMinusFirstDetected, WaapV1DomainAPIPathListParamsOrderingMinusLastDetected, WaapV1DomainAPIPathListParamsOrderingMinusStatus, WaapV1DomainAPIPathListParamsOrderingMinusSource:
		return true
	}
	return false
}
