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

// WaapV1DomainInsightService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainInsightService] method instead.
type WaapV1DomainInsightService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1DomainInsightService(opts ...option.RequestOption) (r *WaapV1DomainInsightService) {
	r = &WaapV1DomainInsightService{}
	r.Options = opts
	return
}

// Retrieve a specific insight for a specific domain.
func (r *WaapV1DomainInsightService) Get(ctx context.Context, domainID int64, insightID string, opts ...option.RequestOption) (res *Insight, err error) {
	opts = append(r.Options[:], opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insights/%s", domainID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the status of an insight for a specific domain.
func (r *WaapV1DomainInsightService) Update(ctx context.Context, domainID int64, insightID string, body WaapV1DomainInsightUpdateParams, opts ...option.RequestOption) (res *Insight, err error) {
	opts = append(r.Options[:], opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insights/%s", domainID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of insights for a specific domain.
func (r *WaapV1DomainInsightService) List(ctx context.Context, domainID int64, query WaapV1DomainInsightListParams, opts ...option.RequestOption) (res *WaapV1DomainInsightListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insights", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Insight struct {
	// A generated unique identifier for the insight
	ID string `json:"id,required" format:"uuid"`
	// The description of the insight
	Description string `json:"description,required"`
	// The date and time the insight was first seen in ISO 8601 format
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// The type of the insight represented as a slug
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight
	Labels map[string]string `json:"labels,required"`
	// The date and time the insight was last seen in ISO 8601 format
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// The date and time the insight was last seen in ISO 8601 format
	LastStatusChange time.Time `json:"last_status_change,required" format:"date-time"`
	// The recommended action to perform to resolve the insight
	Recommendation string `json:"recommendation,required"`
	// The different statuses an insight can have
	Status InsightStatus `json:"status,required"`
	JSON   insightJSON   `json:"-"`
}

// insightJSON contains the JSON metadata for the struct [Insight]
type insightJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	FirstSeen        apijson.Field
	InsightType      apijson.Field
	Labels           apijson.Field
	LastSeen         apijson.Field
	LastStatusChange apijson.Field
	Recommendation   apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Insight) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightJSON) RawJSON() string {
	return r.raw
}

// The different statuses an insight can have
type InsightStatus string

const (
	InsightStatusOpen   InsightStatus = "OPEN"
	InsightStatusAcked  InsightStatus = "ACKED"
	InsightStatusClosed InsightStatus = "CLOSED"
)

func (r InsightStatus) IsKnown() bool {
	switch r {
	case InsightStatusOpen, InsightStatusAcked, InsightStatusClosed:
		return true
	}
	return false
}

type WaapV1DomainInsightListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []Insight                           `json:"results,required"`
	JSON    waapV1DomainInsightListResponseJSON `json:"-"`
}

// waapV1DomainInsightListResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainInsightListResponse]
type waapV1DomainInsightListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainInsightListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainInsightListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainInsightUpdateParams struct {
	// The different statuses an insight can have
	Status param.Field[InsightStatus] `json:"status,required"`
}

func (r WaapV1DomainInsightUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainInsightListParams struct {
	// The ID of the insight
	ID param.Field[[]string] `query:"id" format:"uuid"`
	// The description of the insight. Supports '\*' as a wildcard.
	Description param.Field[string] `query:"description"`
	// The type of the insight
	InsightType param.Field[[]string] `query:"insight_type"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainInsightListParamsOrdering] `query:"ordering"`
	// The status of the insight
	Status param.Field[[]InsightStatus] `query:"status"`
}

// URLQuery serializes [WaapV1DomainInsightListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainInsightListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainInsightListParamsOrdering string

const (
	WaapV1DomainInsightListParamsOrderingID                    WaapV1DomainInsightListParamsOrdering = "id"
	WaapV1DomainInsightListParamsOrderingMinusID               WaapV1DomainInsightListParamsOrdering = "-id"
	WaapV1DomainInsightListParamsOrderingInsightType           WaapV1DomainInsightListParamsOrdering = "insight_type"
	WaapV1DomainInsightListParamsOrderingMinusInsightType      WaapV1DomainInsightListParamsOrdering = "-insight_type"
	WaapV1DomainInsightListParamsOrderingFirstSeen             WaapV1DomainInsightListParamsOrdering = "first_seen"
	WaapV1DomainInsightListParamsOrderingMinusFirstSeen        WaapV1DomainInsightListParamsOrdering = "-first_seen"
	WaapV1DomainInsightListParamsOrderingLastSeen              WaapV1DomainInsightListParamsOrdering = "last_seen"
	WaapV1DomainInsightListParamsOrderingMinusLastSeen         WaapV1DomainInsightListParamsOrdering = "-last_seen"
	WaapV1DomainInsightListParamsOrderingLastStatusChange      WaapV1DomainInsightListParamsOrdering = "last_status_change"
	WaapV1DomainInsightListParamsOrderingMinusLastStatusChange WaapV1DomainInsightListParamsOrdering = "-last_status_change"
	WaapV1DomainInsightListParamsOrderingStatus                WaapV1DomainInsightListParamsOrdering = "status"
	WaapV1DomainInsightListParamsOrderingMinusStatus           WaapV1DomainInsightListParamsOrdering = "-status"
)

func (r WaapV1DomainInsightListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainInsightListParamsOrderingID, WaapV1DomainInsightListParamsOrderingMinusID, WaapV1DomainInsightListParamsOrderingInsightType, WaapV1DomainInsightListParamsOrderingMinusInsightType, WaapV1DomainInsightListParamsOrderingFirstSeen, WaapV1DomainInsightListParamsOrderingMinusFirstSeen, WaapV1DomainInsightListParamsOrderingLastSeen, WaapV1DomainInsightListParamsOrderingMinusLastSeen, WaapV1DomainInsightListParamsOrderingLastStatusChange, WaapV1DomainInsightListParamsOrderingMinusLastStatusChange, WaapV1DomainInsightListParamsOrderingStatus, WaapV1DomainInsightListParamsOrderingMinusStatus:
		return true
	}
	return false
}
