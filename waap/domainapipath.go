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

// DomainAPIPathService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIPathService] method instead.
type DomainAPIPathService struct {
	Options []option.RequestOption
}

// NewDomainAPIPathService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainAPIPathService(opts ...option.RequestOption) (r DomainAPIPathService) {
	r = DomainAPIPathService{}
	r.Options = opts
	return
}

// Create an API path for a domain
func (r *DomainAPIPathService) New(ctx context.Context, domainID int64, body DomainAPIPathNewParams, opts ...option.RequestOption) (res *WaapAPIPath, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a specific API path for a domain
func (r *DomainAPIPathService) Update(ctx context.Context, pathID string, params DomainAPIPathUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", params.DomainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Retrieve a list of API paths for a specific domain
func (r *DomainAPIPathService) List(ctx context.Context, domainID int64, query DomainAPIPathListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapAPIPath], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths", domainID)
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

// Retrieve a list of API paths for a specific domain
func (r *DomainAPIPathService) ListAutoPaging(ctx context.Context, domainID int64, query DomainAPIPathListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapAPIPath] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete a specific API path for a domain
func (r *DomainAPIPathService) Delete(ctx context.Context, pathID string, body DomainAPIPathDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", body.DomainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a specific API path for a domain
func (r *DomainAPIPathService) Get(ctx context.Context, pathID string, query DomainAPIPathGetParams, opts ...option.RequestOption) (res *WaapAPIPath, err error) {
	opts = slices.Concat(r.Options, opts)
	if pathID == "" {
		err = errors.New("missing required path_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/api-paths/%s", query.DomainID, pathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response model for the API path
type WaapAPIPath struct {
	// The path ID
	ID string `json:"id,required" format:"uuid"`
	// An array of api groups associated with the API path
	APIGroups []string `json:"api_groups,required"`
	// The API version
	APIVersion string `json:"api_version,required"`
	// The date and time in ISO 8601 format the API path was first detected.
	FirstDetected time.Time `json:"first_detected,required" format:"date-time"`
	// The HTTP version of the API path
	//
	// Any of "HTTP", "HTTPS".
	HTTPScheme WaapAPIPathHTTPScheme `json:"http_scheme,required"`
	// The date and time in ISO 8601 format the API path was last detected.
	LastDetected time.Time `json:"last_detected,required" format:"date-time"`
	// The API RESTful method
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE", "TRACE", "HEAD", "OPTIONS".
	Method WaapAPIPathMethod `json:"method,required"`
	// The API path, locations that are saved for resource IDs will be put in curly
	// brackets
	Path string `json:"path,required"`
	// The number of requests for this path in the last 24 hours
	RequestCount int64 `json:"request_count,required"`
	// The source of the discovered API
	//
	// Any of "API_DESCRIPTION_FILE", "TRAFFIC_SCAN", "USER_DEFINED".
	Source WaapAPIPathSource `json:"source,required"`
	// The status of the discovered API path
	//
	// Any of "CONFIRMED_API", "POTENTIAL_API", "NOT_API", "DELISTED_API".
	Status WaapAPIPathStatus `json:"status,required"`
	// An array of tags associated with the API path
	Tags []string `json:"tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		APIGroups     respjson.Field
		APIVersion    respjson.Field
		FirstDetected respjson.Field
		HTTPScheme    respjson.Field
		LastDetected  respjson.Field
		Method        respjson.Field
		Path          respjson.Field
		RequestCount  respjson.Field
		Source        respjson.Field
		Status        respjson.Field
		Tags          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAPIPath) RawJSON() string { return r.JSON.raw }
func (r *WaapAPIPath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP version of the API path
type WaapAPIPathHTTPScheme string

const (
	WaapAPIPathHTTPSchemeHTTP  WaapAPIPathHTTPScheme = "HTTP"
	WaapAPIPathHTTPSchemeHTTPS WaapAPIPathHTTPScheme = "HTTPS"
)

// The API RESTful method
type WaapAPIPathMethod string

const (
	WaapAPIPathMethodGet     WaapAPIPathMethod = "GET"
	WaapAPIPathMethodPost    WaapAPIPathMethod = "POST"
	WaapAPIPathMethodPut     WaapAPIPathMethod = "PUT"
	WaapAPIPathMethodPatch   WaapAPIPathMethod = "PATCH"
	WaapAPIPathMethodDelete  WaapAPIPathMethod = "DELETE"
	WaapAPIPathMethodTrace   WaapAPIPathMethod = "TRACE"
	WaapAPIPathMethodHead    WaapAPIPathMethod = "HEAD"
	WaapAPIPathMethodOptions WaapAPIPathMethod = "OPTIONS"
)

// The source of the discovered API
type WaapAPIPathSource string

const (
	WaapAPIPathSourceAPIDescriptionFile WaapAPIPathSource = "API_DESCRIPTION_FILE"
	WaapAPIPathSourceTrafficScan        WaapAPIPathSource = "TRAFFIC_SCAN"
	WaapAPIPathSourceUserDefined        WaapAPIPathSource = "USER_DEFINED"
)

// The status of the discovered API path
type WaapAPIPathStatus string

const (
	WaapAPIPathStatusConfirmedAPI WaapAPIPathStatus = "CONFIRMED_API"
	WaapAPIPathStatusPotentialAPI WaapAPIPathStatus = "POTENTIAL_API"
	WaapAPIPathStatusNotAPI       WaapAPIPathStatus = "NOT_API"
	WaapAPIPathStatusDelistedAPI  WaapAPIPathStatus = "DELISTED_API"
)

type DomainAPIPathNewParams struct {
	// The different HTTP schemes an API path can have
	//
	// Any of "HTTP", "HTTPS".
	HTTPScheme DomainAPIPathNewParamsHTTPScheme `json:"http_scheme,omitzero,required"`
	// The different methods an API path can have
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE", "TRACE", "HEAD", "OPTIONS".
	Method DomainAPIPathNewParamsMethod `json:"method,omitzero,required"`
	// The API path, locations that are saved for resource IDs will be put in curly
	// brackets
	Path       string            `json:"path,required"`
	APIVersion param.Opt[string] `json:"api_version,omitzero"`
	APIGroups  []string          `json:"api_groups,omitzero"`
	Tags       []string          `json:"tags,omitzero"`
	paramObj
}

func (r DomainAPIPathNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIPathNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIPathNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different HTTP schemes an API path can have
type DomainAPIPathNewParamsHTTPScheme string

const (
	DomainAPIPathNewParamsHTTPSchemeHTTP  DomainAPIPathNewParamsHTTPScheme = "HTTP"
	DomainAPIPathNewParamsHTTPSchemeHTTPS DomainAPIPathNewParamsHTTPScheme = "HTTPS"
)

// The different methods an API path can have
type DomainAPIPathNewParamsMethod string

const (
	DomainAPIPathNewParamsMethodGet     DomainAPIPathNewParamsMethod = "GET"
	DomainAPIPathNewParamsMethodPost    DomainAPIPathNewParamsMethod = "POST"
	DomainAPIPathNewParamsMethodPut     DomainAPIPathNewParamsMethod = "PUT"
	DomainAPIPathNewParamsMethodPatch   DomainAPIPathNewParamsMethod = "PATCH"
	DomainAPIPathNewParamsMethodDelete  DomainAPIPathNewParamsMethod = "DELETE"
	DomainAPIPathNewParamsMethodTrace   DomainAPIPathNewParamsMethod = "TRACE"
	DomainAPIPathNewParamsMethodHead    DomainAPIPathNewParamsMethod = "HEAD"
	DomainAPIPathNewParamsMethodOptions DomainAPIPathNewParamsMethod = "OPTIONS"
)

type DomainAPIPathUpdateParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The updated API path. When updating the path, variables can be renamed, path
	// parts can be converted to variables and vice versa.
	Path      param.Opt[string] `json:"path,omitzero"`
	APIGroups []string          `json:"api_groups,omitzero"`
	// The different statuses an API path can have
	//
	// Any of "CONFIRMED_API", "POTENTIAL_API", "NOT_API", "DELISTED_API".
	Status DomainAPIPathUpdateParamsStatus `json:"status,omitzero"`
	Tags   []string                        `json:"tags,omitzero"`
	paramObj
}

func (r DomainAPIPathUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIPathUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIPathUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses an API path can have
type DomainAPIPathUpdateParamsStatus string

const (
	DomainAPIPathUpdateParamsStatusConfirmedAPI DomainAPIPathUpdateParamsStatus = "CONFIRMED_API"
	DomainAPIPathUpdateParamsStatusPotentialAPI DomainAPIPathUpdateParamsStatus = "POTENTIAL_API"
	DomainAPIPathUpdateParamsStatusNotAPI       DomainAPIPathUpdateParamsStatus = "NOT_API"
	DomainAPIPathUpdateParamsStatusDelistedAPI  DomainAPIPathUpdateParamsStatus = "DELISTED_API"
)

type DomainAPIPathListParams struct {
	// Filter by the API group associated with the API path
	APIGroup param.Opt[string] `query:"api_group,omitzero" json:"-"`
	// Filter by the API version
	APIVersion param.Opt[string] `query:"api_version,omitzero" json:"-"`
	// Filter by the path. Supports '\*' as a wildcard character
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The different HTTP schemes an API path can have
	//
	// Any of "HTTP", "HTTPS".
	HTTPScheme DomainAPIPathListParamsHTTPScheme `query:"http_scheme,omitzero" json:"-"`
	// Filter by the path ID
	IDs []string `query:"ids,omitzero" format:"uuid" json:"-"`
	// The different methods an API path can have
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE", "TRACE", "HEAD", "OPTIONS".
	Method DomainAPIPathListParamsMethod `query:"method,omitzero" json:"-"`
	// The different sources an API path can have
	//
	// Any of "API_DESCRIPTION_FILE", "TRAFFIC_SCAN", "USER_DEFINED".
	Source DomainAPIPathListParamsSource `query:"source,omitzero" json:"-"`
	// Filter by the status of the discovered API path
	//
	// Any of "CONFIRMED_API", "POTENTIAL_API", "NOT_API", "DELISTED_API".
	Status []string `query:"status,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "path", "method", "api_version", "http_scheme", "first_detected",
	// "last_detected", "status", "source", "-id", "-path", "-method", "-api_version",
	// "-http_scheme", "-first_detected", "-last_detected", "-status", "-source".
	Ordering DomainAPIPathListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAPIPathListParams]'s query parameters as
// `url.Values`.
func (r DomainAPIPathListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The different HTTP schemes an API path can have
type DomainAPIPathListParamsHTTPScheme string

const (
	DomainAPIPathListParamsHTTPSchemeHTTP  DomainAPIPathListParamsHTTPScheme = "HTTP"
	DomainAPIPathListParamsHTTPSchemeHTTPS DomainAPIPathListParamsHTTPScheme = "HTTPS"
)

// The different methods an API path can have
type DomainAPIPathListParamsMethod string

const (
	DomainAPIPathListParamsMethodGet     DomainAPIPathListParamsMethod = "GET"
	DomainAPIPathListParamsMethodPost    DomainAPIPathListParamsMethod = "POST"
	DomainAPIPathListParamsMethodPut     DomainAPIPathListParamsMethod = "PUT"
	DomainAPIPathListParamsMethodPatch   DomainAPIPathListParamsMethod = "PATCH"
	DomainAPIPathListParamsMethodDelete  DomainAPIPathListParamsMethod = "DELETE"
	DomainAPIPathListParamsMethodTrace   DomainAPIPathListParamsMethod = "TRACE"
	DomainAPIPathListParamsMethodHead    DomainAPIPathListParamsMethod = "HEAD"
	DomainAPIPathListParamsMethodOptions DomainAPIPathListParamsMethod = "OPTIONS"
)

// Sort the response by given field.
type DomainAPIPathListParamsOrdering string

const (
	DomainAPIPathListParamsOrderingID                 DomainAPIPathListParamsOrdering = "id"
	DomainAPIPathListParamsOrderingPath               DomainAPIPathListParamsOrdering = "path"
	DomainAPIPathListParamsOrderingMethod             DomainAPIPathListParamsOrdering = "method"
	DomainAPIPathListParamsOrderingAPIVersion         DomainAPIPathListParamsOrdering = "api_version"
	DomainAPIPathListParamsOrderingHTTPScheme         DomainAPIPathListParamsOrdering = "http_scheme"
	DomainAPIPathListParamsOrderingFirstDetected      DomainAPIPathListParamsOrdering = "first_detected"
	DomainAPIPathListParamsOrderingLastDetected       DomainAPIPathListParamsOrdering = "last_detected"
	DomainAPIPathListParamsOrderingStatus             DomainAPIPathListParamsOrdering = "status"
	DomainAPIPathListParamsOrderingSource             DomainAPIPathListParamsOrdering = "source"
	DomainAPIPathListParamsOrderingMinusID            DomainAPIPathListParamsOrdering = "-id"
	DomainAPIPathListParamsOrderingMinusPath          DomainAPIPathListParamsOrdering = "-path"
	DomainAPIPathListParamsOrderingMinusMethod        DomainAPIPathListParamsOrdering = "-method"
	DomainAPIPathListParamsOrderingMinusAPIVersion    DomainAPIPathListParamsOrdering = "-api_version"
	DomainAPIPathListParamsOrderingMinusHTTPScheme    DomainAPIPathListParamsOrdering = "-http_scheme"
	DomainAPIPathListParamsOrderingMinusFirstDetected DomainAPIPathListParamsOrdering = "-first_detected"
	DomainAPIPathListParamsOrderingMinusLastDetected  DomainAPIPathListParamsOrdering = "-last_detected"
	DomainAPIPathListParamsOrderingMinusStatus        DomainAPIPathListParamsOrdering = "-status"
	DomainAPIPathListParamsOrderingMinusSource        DomainAPIPathListParamsOrdering = "-source"
)

// The different sources an API path can have
type DomainAPIPathListParamsSource string

const (
	DomainAPIPathListParamsSourceAPIDescriptionFile DomainAPIPathListParamsSource = "API_DESCRIPTION_FILE"
	DomainAPIPathListParamsSourceTrafficScan        DomainAPIPathListParamsSource = "TRAFFIC_SCAN"
	DomainAPIPathListParamsSourceUserDefined        DomainAPIPathListParamsSource = "USER_DEFINED"
)

type DomainAPIPathDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainAPIPathGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
