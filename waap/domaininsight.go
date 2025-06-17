// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
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
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/insights/%s", params.DomainID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

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
	Status []WaapInsightStatus `query:"status,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "-id", "insight_type", "-insight_type", "first_seen",
	// "-first_seen", "last_seen", "-last_seen", "last_status_change",
	// "-last_status_change", "status", "-status".
	Ordering WaapInsightSortBy `query:"ordering,omitzero" json:"-"`
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

type DomainInsightGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainInsightReplaceParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The different statuses an insight can have
	//
	// Any of "OPEN", "ACKED", "CLOSED".
	Status WaapInsightStatus `json:"status,omitzero,required"`
	paramObj
}

func (r DomainInsightReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainInsightReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainInsightReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
