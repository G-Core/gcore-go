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

// WaapV1DomainInsightSilenceService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainInsightSilenceService] method instead.
type WaapV1DomainInsightSilenceService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainInsightSilenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainInsightSilenceService(opts ...option.RequestOption) (r *WaapV1DomainInsightSilenceService) {
	r = &WaapV1DomainInsightSilenceService{}
	r.Options = opts
	return
}

// Create a new insight silence for a specified domain. Insight silences help in
// temporarily disabling certain insights based on specific criteria.
func (r *WaapV1DomainInsightSilenceService) New(ctx context.Context, domainID int64, body WaapV1DomainInsightSilenceNewParams, opts ...option.RequestOption) (res *InsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific insight silence for a specific domain
func (r *WaapV1DomainInsightSilenceService) Get(ctx context.Context, domainID int64, silenceID string, opts ...option.RequestOption) (res *InsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", domainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an insight silence for a specific domain.
func (r *WaapV1DomainInsightSilenceService) Update(ctx context.Context, domainID int64, silenceID string, body WaapV1DomainInsightSilenceUpdateParams, opts ...option.RequestOption) (res *InsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", domainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a list of insight silences for a specific domain
func (r *WaapV1DomainInsightSilenceService) List(ctx context.Context, domainID int64, query WaapV1DomainInsightSilenceListParams, opts ...option.RequestOption) (res *WaapV1DomainInsightSilenceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an insight silence for a specific domain.
func (r *WaapV1DomainInsightSilenceService) Delete(ctx context.Context, domainID int64, silenceID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", domainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type InsightSilence struct {
	// A generated unique identifier for the silence
	ID string `json:"id,required" format:"uuid"`
	// The author of the silence
	Author string `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment string `json:"comment,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt time.Time `json:"expire_at,required,nullable" format:"date-time"`
	// The slug of the insight type
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels map[string]string  `json:"labels,required"`
	JSON   insightSilenceJSON `json:"-"`
}

// insightSilenceJSON contains the JSON metadata for the struct [InsightSilence]
type insightSilenceJSON struct {
	ID          apijson.Field
	Author      apijson.Field
	Comment     apijson.Field
	ExpireAt    apijson.Field
	InsightType apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightSilence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSilenceJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainInsightSilenceListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []InsightSilence                           `json:"results,required"`
	JSON    waapV1DomainInsightSilenceListResponseJSON `json:"-"`
}

// waapV1DomainInsightSilenceListResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainInsightSilenceListResponse]
type waapV1DomainInsightSilenceListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainInsightSilenceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainInsightSilenceListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainInsightSilenceNewParams struct {
	// The author of the silence
	Author param.Field[string] `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment param.Field[string] `json:"comment,required"`
	// The slug of the insight type
	InsightType param.Field[string] `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels param.Field[map[string]string] `json:"labels,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt param.Field[time.Time] `json:"expire_at" format:"date-time"`
}

func (r WaapV1DomainInsightSilenceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainInsightSilenceUpdateParams struct {
	// The author of the silence
	Author param.Field[string] `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment param.Field[string] `json:"comment,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt param.Field[time.Time] `json:"expire_at,required" format:"date-time"`
	// A hash table of label names and values that apply to the insight silence
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r WaapV1DomainInsightSilenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainInsightSilenceListParams struct {
	// The ID of the insight silence
	ID param.Field[[]string] `query:"id" format:"uuid"`
	// The author of the insight silence
	Author param.Field[string] `query:"author"`
	// The comment of the insight silence
	Comment param.Field[string] `query:"comment"`
	// The type of the insight silence
	InsightType param.Field[[]string] `query:"insight_type"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainInsightSilenceListParamsOrdering] `query:"ordering"`
}

// URLQuery serializes [WaapV1DomainInsightSilenceListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainInsightSilenceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainInsightSilenceListParamsOrdering string

const (
	WaapV1DomainInsightSilenceListParamsOrderingID               WaapV1DomainInsightSilenceListParamsOrdering = "id"
	WaapV1DomainInsightSilenceListParamsOrderingMinusID          WaapV1DomainInsightSilenceListParamsOrdering = "-id"
	WaapV1DomainInsightSilenceListParamsOrderingInsightType      WaapV1DomainInsightSilenceListParamsOrdering = "insight_type"
	WaapV1DomainInsightSilenceListParamsOrderingMinusInsightType WaapV1DomainInsightSilenceListParamsOrdering = "-insight_type"
	WaapV1DomainInsightSilenceListParamsOrderingComment          WaapV1DomainInsightSilenceListParamsOrdering = "comment"
	WaapV1DomainInsightSilenceListParamsOrderingMinusComment     WaapV1DomainInsightSilenceListParamsOrdering = "-comment"
	WaapV1DomainInsightSilenceListParamsOrderingAuthor           WaapV1DomainInsightSilenceListParamsOrdering = "author"
	WaapV1DomainInsightSilenceListParamsOrderingMinusAuthor      WaapV1DomainInsightSilenceListParamsOrdering = "-author"
	WaapV1DomainInsightSilenceListParamsOrderingExpireAt         WaapV1DomainInsightSilenceListParamsOrdering = "expire_at"
	WaapV1DomainInsightSilenceListParamsOrderingMinusExpireAt    WaapV1DomainInsightSilenceListParamsOrdering = "-expire_at"
)

func (r WaapV1DomainInsightSilenceListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainInsightSilenceListParamsOrderingID, WaapV1DomainInsightSilenceListParamsOrderingMinusID, WaapV1DomainInsightSilenceListParamsOrderingInsightType, WaapV1DomainInsightSilenceListParamsOrderingMinusInsightType, WaapV1DomainInsightSilenceListParamsOrderingComment, WaapV1DomainInsightSilenceListParamsOrderingMinusComment, WaapV1DomainInsightSilenceListParamsOrderingAuthor, WaapV1DomainInsightSilenceListParamsOrderingMinusAuthor, WaapV1DomainInsightSilenceListParamsOrderingExpireAt, WaapV1DomainInsightSilenceListParamsOrderingMinusExpireAt:
		return true
	}
	return false
}
