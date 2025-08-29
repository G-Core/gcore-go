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

// DomainInsightSilenceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainInsightSilenceService] method instead.
type DomainInsightSilenceService struct {
	Options []option.RequestOption
}

// NewDomainInsightSilenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainInsightSilenceService(opts ...option.RequestOption) (r DomainInsightSilenceService) {
	r = DomainInsightSilenceService{}
	r.Options = opts
	return
}

// Create a new insight silence for a specified domain. Insight silences help in
// temporarily disabling certain insights based on specific criteria.
func (r *DomainInsightSilenceService) New(ctx context.Context, domainID int64, body DomainInsightSilenceNewParams, opts ...option.RequestOption) (res *WaapInsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an insight silence for a specific domain.
func (r *DomainInsightSilenceService) Update(ctx context.Context, silenceID string, params DomainInsightSilenceUpdateParams, opts ...option.RequestOption) (res *WaapInsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", params.DomainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve a list of insight silences for a specific domain
func (r *DomainInsightSilenceService) List(ctx context.Context, domainID int64, query DomainInsightSilenceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapInsightSilence], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences", domainID)
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

// Retrieve a list of insight silences for a specific domain
func (r *DomainInsightSilenceService) ListAutoPaging(ctx context.Context, domainID int64, query DomainInsightSilenceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapInsightSilence] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete an insight silence for a specific domain.
func (r *DomainInsightSilenceService) Delete(ctx context.Context, silenceID string, body DomainInsightSilenceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", body.DomainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a specific insight silence for a specific domain
func (r *DomainInsightSilenceService) Get(ctx context.Context, silenceID string, query DomainInsightSilenceGetParams, opts ...option.RequestOption) (res *WaapInsightSilence, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insight-silences/%s", query.DomainID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WaapInsightSilence struct {
	// A generated unique identifier for the silence
	ID string `json:"id,required" format:"uuid"`
	// The author of the silence
	Author string `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment string `json:"comment,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt time.Time `json:"expire_at,required" format:"date-time"`
	// The slug of the insight type
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels map[string]string `json:"labels,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Author      respjson.Field
		Comment     respjson.Field
		ExpireAt    respjson.Field
		InsightType respjson.Field
		Labels      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapInsightSilence) RawJSON() string { return r.JSON.raw }
func (r *WaapInsightSilence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainInsightSilenceNewParams struct {
	// The author of the silence
	Author string `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment string `json:"comment,required"`
	// The slug of the insight type
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels map[string]string `json:"labels,omitzero,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt param.Opt[time.Time] `json:"expire_at,omitzero" format:"date-time"`
	paramObj
}

func (r DomainInsightSilenceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainInsightSilenceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainInsightSilenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainInsightSilenceUpdateParams struct {
	// The date and time the silence expires in ISO 8601 format
	ExpireAt param.Opt[time.Time] `json:"expire_at,omitzero,required" format:"date-time"`
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The author of the silence
	Author string `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment string `json:"comment,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r DomainInsightSilenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainInsightSilenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainInsightSilenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainInsightSilenceListParams struct {
	// The author of the insight silence
	Author param.Opt[string] `query:"author,omitzero" json:"-"`
	// The comment of the insight silence
	Comment param.Opt[string] `query:"comment,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The ID of the insight silence
	ID []string `query:"id,omitzero" format:"uuid" json:"-"`
	// The type of the insight silence
	InsightType []string `query:"insight_type,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "-id", "insight_type", "-insight_type", "comment", "-comment",
	// "author", "-author", "expire_at", "-expire_at".
	Ordering DomainInsightSilenceListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainInsightSilenceListParams]'s query parameters as
// `url.Values`.
func (r DomainInsightSilenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainInsightSilenceListParamsOrdering string

const (
	DomainInsightSilenceListParamsOrderingID               DomainInsightSilenceListParamsOrdering = "id"
	DomainInsightSilenceListParamsOrderingMinusID          DomainInsightSilenceListParamsOrdering = "-id"
	DomainInsightSilenceListParamsOrderingInsightType      DomainInsightSilenceListParamsOrdering = "insight_type"
	DomainInsightSilenceListParamsOrderingMinusInsightType DomainInsightSilenceListParamsOrdering = "-insight_type"
	DomainInsightSilenceListParamsOrderingComment          DomainInsightSilenceListParamsOrdering = "comment"
	DomainInsightSilenceListParamsOrderingMinusComment     DomainInsightSilenceListParamsOrdering = "-comment"
	DomainInsightSilenceListParamsOrderingAuthor           DomainInsightSilenceListParamsOrdering = "author"
	DomainInsightSilenceListParamsOrderingMinusAuthor      DomainInsightSilenceListParamsOrdering = "-author"
	DomainInsightSilenceListParamsOrderingExpireAt         DomainInsightSilenceListParamsOrdering = "expire_at"
	DomainInsightSilenceListParamsOrderingMinusExpireAt    DomainInsightSilenceListParamsOrdering = "-expire_at"
)

type DomainInsightSilenceDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainInsightSilenceGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
