// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// IPReputationService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPReputationService] method instead.
type IPReputationService struct {
	Options []option.RequestOption
}

// NewIPReputationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIPReputationService(opts ...option.RequestOption) (r IPReputationService) {
	r = IPReputationService{}
	r.Options = opts
	return
}

// Retrieve a list of tags associated with the IP address grouped by domains. This
// will return records for all domains which were accessed by the IP and received
// tags.
func (r *IPReputationService) List(ctx context.Context, query IPReputationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapLocalIPTagsRecord], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/ip-reputation"
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

// Retrieve a list of tags associated with the IP address grouped by domains. This
// will return records for all domains which were accessed by the IP and received
// tags.
func (r *IPReputationService) ListAutoPaging(ctx context.Context, query IPReputationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapLocalIPTagsRecord] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Clear all tags associated between the provided IP addresses and domains.
func (r *IPReputationService) ClearLocalTags(ctx context.Context, body IPReputationClearLocalTagsParams, opts ...option.RequestOption) (res *WaapClearLocalTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v2/ip-reputation/clear-local-tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WaapClearLocalTagsResponse struct {
	// The number of items affected by the action
	NumCleared int64 `json:"num_cleared"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumCleared  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapClearLocalTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *WaapClearLocalTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapLocalIPTagsRecord struct {
	// The timestamp of the last request from the IP to the domain
	Date   int64                       `json:"date" api:"required"`
	Domain WaapLocalIPTagsRecordDomain `json:"domain" api:"required"`
	// The tags associated with the IP globally
	GlobalTags []string `json:"global_tags" api:"required"`
	// The IP address
	IP string `json:"ip" api:"required"`
	// The tags associated with the IP in the context of the domain
	LocalTags []string `json:"local_tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Domain      respjson.Field
		GlobalTags  respjson.Field
		IP          respjson.Field
		LocalTags   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapLocalIPTagsRecord) RawJSON() string { return r.JSON.raw }
func (r *WaapLocalIPTagsRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapLocalIPTagsRecordDomain struct {
	// The unique identifier for the domain
	ID int64 `json:"id" api:"required"`
	// The domain name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapLocalIPTagsRecordDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapLocalIPTagsRecordDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPReputationListParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipvanyaddress" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "date", "-date", "domain_name", "-domain_name".
	Ordering IPReputationListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPReputationListParams]'s query parameters as `url.Values`.
func (r IPReputationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type IPReputationListParamsOrdering string

const (
	IPReputationListParamsOrderingDate            IPReputationListParamsOrdering = "date"
	IPReputationListParamsOrderingMinusDate       IPReputationListParamsOrdering = "-date"
	IPReputationListParamsOrderingDomainName      IPReputationListParamsOrdering = "domain_name"
	IPReputationListParamsOrderingMinusDomainName IPReputationListParamsOrdering = "-domain_name"
)

type IPReputationClearLocalTagsParams struct {
	Items []IPReputationClearLocalTagsParamsItem
	paramObj
}

func (r IPReputationClearLocalTagsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Items)
}
func (r *IPReputationClearLocalTagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request model for deleting local IP tags.
//
// The properties DomainID, IP are required.
type IPReputationClearLocalTagsParamsItem struct {
	// The domain ID to delete tags for
	DomainID int64 `json:"domain_id" api:"required"`
	// The IP address to delete tags for
	IP string `json:"ip" api:"required" format:"ipvanyaddress"`
	paramObj
}

func (r IPReputationClearLocalTagsParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow IPReputationClearLocalTagsParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPReputationClearLocalTagsParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
