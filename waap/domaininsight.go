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

// DomainInsightService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainInsightService] method instead.
type DomainInsightService struct {
	Options []option.RequestOption
}

// NewDomainInsightService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainInsightService(opts ...option.RequestOption) (r DomainInsightService) {
	r = DomainInsightService{}
	r.Options = opts
	return
}

// Retrieve a list of insights for a specific domain.
func (r *DomainInsightService) List(ctx context.Context, domainID int64, query DomainInsightListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapInsight], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insights", domainID)
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

// Retrieve a list of insights for a specific domain.
func (r *DomainInsightService) ListAutoPaging(ctx context.Context, domainID int64, query DomainInsightListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapInsight] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Retrieve a specific insight for a specific domain.
func (r *DomainInsightService) Get(ctx context.Context, insightID string, query DomainInsightGetParams, opts ...option.RequestOption) (res *WaapInsight, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insights/%s", query.DomainID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the status of an insight for a specific domain.
func (r *DomainInsightService) Replace(ctx context.Context, insightID string, params DomainInsightReplaceParams, opts ...option.RequestOption) (res *WaapInsight, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insights/%s", params.DomainID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type WaapInsight struct {
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
	// The status of the insight
	//
	// Any of "OPEN", "ACKED", "CLOSED".
	Status WaapInsightStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Description      respjson.Field
		FirstSeen        respjson.Field
		InsightType      respjson.Field
		Labels           respjson.Field
		LastSeen         respjson.Field
		LastStatusChange respjson.Field
		Recommendation   respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapInsight) RawJSON() string { return r.JSON.raw }
func (r *WaapInsight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the insight
type WaapInsightStatus string

const (
	WaapInsightStatusOpen   WaapInsightStatus = "OPEN"
	WaapInsightStatusAcked  WaapInsightStatus = "ACKED"
	WaapInsightStatusClosed WaapInsightStatus = "CLOSED"
)

type DomainInsightListParams struct {
	// The description of the insight. Supports '\*' as a wildcard.
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The ID of the insight
	ID []string `query:"id,omitzero" format:"uuid" json:"-"`
	// The type of the insight
	InsightType []string `query:"insight_type,omitzero" json:"-"`
	// The status of the insight
	//
	// Any of "OPEN", "ACKED", "CLOSED".
	Status []string `query:"status,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "-id", "insight_type", "-insight_type", "first_seen",
	// "-first_seen", "last_seen", "-last_seen", "last_status_change",
	// "-last_status_change", "status", "-status".
	Ordering DomainInsightListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainInsightListParams]'s query parameters as
// `url.Values`.
func (r DomainInsightListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainInsightListParamsOrdering string

const (
	DomainInsightListParamsOrderingID                    DomainInsightListParamsOrdering = "id"
	DomainInsightListParamsOrderingMinusID               DomainInsightListParamsOrdering = "-id"
	DomainInsightListParamsOrderingInsightType           DomainInsightListParamsOrdering = "insight_type"
	DomainInsightListParamsOrderingMinusInsightType      DomainInsightListParamsOrdering = "-insight_type"
	DomainInsightListParamsOrderingFirstSeen             DomainInsightListParamsOrdering = "first_seen"
	DomainInsightListParamsOrderingMinusFirstSeen        DomainInsightListParamsOrdering = "-first_seen"
	DomainInsightListParamsOrderingLastSeen              DomainInsightListParamsOrdering = "last_seen"
	DomainInsightListParamsOrderingMinusLastSeen         DomainInsightListParamsOrdering = "-last_seen"
	DomainInsightListParamsOrderingLastStatusChange      DomainInsightListParamsOrdering = "last_status_change"
	DomainInsightListParamsOrderingMinusLastStatusChange DomainInsightListParamsOrdering = "-last_status_change"
	DomainInsightListParamsOrderingStatus                DomainInsightListParamsOrdering = "status"
	DomainInsightListParamsOrderingMinusStatus           DomainInsightListParamsOrdering = "-status"
)

type DomainInsightGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainInsightReplaceParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The status of the insight
	//
	// Any of "OPEN", "ACKED", "CLOSED".
	Status DomainInsightReplaceParamsStatus `json:"status,omitzero,required"`
	paramObj
}

func (r DomainInsightReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainInsightReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainInsightReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the insight
type DomainInsightReplaceParamsStatus string

const (
	DomainInsightReplaceParamsStatusOpen   DomainInsightReplaceParamsStatus = "OPEN"
	DomainInsightReplaceParamsStatusAcked  DomainInsightReplaceParamsStatus = "ACKED"
	DomainInsightReplaceParamsStatusClosed DomainInsightReplaceParamsStatus = "CLOSED"
)
