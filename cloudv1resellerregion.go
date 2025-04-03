// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ResellerRegionService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ResellerRegionService] method instead.
type CloudV1ResellerRegionService struct {
	Options []option.RequestOption
}

// NewCloudV1ResellerRegionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ResellerRegionService(opts ...option.RequestOption) (r *CloudV1ResellerRegionService) {
	r = &CloudV1ResellerRegionService{}
	r.Options = opts
	return
}

// Region access management. Resellers can open or close access to specific 'core'
// or 'edge' regions for all their clients either individually or collectively.
// Cloud admins can do it for any clients. By default, only 'core' regions are
// available to clients.
func (r *CloudV1ResellerRegionService) Update(ctx context.Context, body CloudV1ResellerRegionUpdateParams, opts ...option.RequestOption) (res *RegionAccess, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reseller_region"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a list of region access records for a client or a reseller
func (r *CloudV1ResellerRegionService) List(ctx context.Context, query CloudV1ResellerRegionListParams, opts ...option.RequestOption) (res *CloudV1ResellerRegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reseller_region"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete reseller's region access records
func (r *CloudV1ResellerRegionService) Delete(ctx context.Context, resellerID int64, body CloudV1ResellerRegionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/reseller_region/%v", resellerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type RegionAccess struct {
	// region access ID
	ID int64 `json:"id,required"`
	// Name of client
	ClientName string `json:"client_name,required,nullable"`
	// Name of reseller
	ResellerName string `json:"reseller_name,required,nullable"`
	// If true, allow access to all edge regions, regardless of the content of
	// region_ids array.
	AccessAllEdgeRegions bool `json:"access_all_edge_regions"`
	// Client ID. If you want to set limits to all reseller clients, skip this field.
	// The client_id has priority over the reseller_id. If both client_id and
	// reseller_id are specified, reseller_id must be the correct reseller_id of the
	// client.
	ClientID int64 `json:"client_id,nullable"`
	// List of available region ids
	RegionIDs []int64 `json:"region_ids,nullable"`
	// Reseller ID. It's null when a client doesn't have a reseller.
	ResellerID int64            `json:"reseller_id,nullable"`
	JSON       regionAccessJSON `json:"-"`
}

// regionAccessJSON contains the JSON metadata for the struct [RegionAccess]
type regionAccessJSON struct {
	ID                   apijson.Field
	ClientName           apijson.Field
	ResellerName         apijson.Field
	AccessAllEdgeRegions apijson.Field
	ClientID             apijson.Field
	RegionIDs            apijson.Field
	ResellerID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RegionAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionAccessJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerRegionListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []RegionAccess                        `json:"results,required"`
	JSON    cloudV1ResellerRegionListResponseJSON `json:"-"`
}

// cloudV1ResellerRegionListResponseJSON contains the JSON metadata for the struct
// [CloudV1ResellerRegionListResponse]
type cloudV1ResellerRegionListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ResellerRegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ResellerRegionListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerRegionUpdateParams struct {
	// If true, allow access to all edge regions, regardless of the content of
	// region_ids array.
	AccessAllEdgeRegions param.Field[bool] `json:"access_all_edge_regions"`
	// Client ID. If you want to set limits to all reseller clients, skip this field.
	// The client_id has priority over the reseller_id. If both client_id and
	// reseller_id are specified, reseller_id must be the correct reseller_id of the
	// client.
	ClientID param.Field[int64] `json:"client_id"`
	// List of available region ids
	RegionIDs param.Field[[]int64] `json:"region_ids"`
	// Reseller ID. It's null when a client doesn't have a reseller.
	ResellerID param.Field[int64] `json:"reseller_id"`
}

func (r CloudV1ResellerRegionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ResellerRegionListParams struct {
	// Client ID. It is assumed to be null when unspecified.
	ClientID param.Field[int64] `query:"client_id"`
	// Reseller ID
	ResellerID param.Field[int64] `query:"reseller_id"`
}

// URLQuery serializes [CloudV1ResellerRegionListParams]'s query parameters as
// `url.Values`.
func (r CloudV1ResellerRegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1ResellerRegionDeleteParams struct {
	// Delete all reseller's records. Defaults to false.
	AllRecords param.Field[bool] `query:"all_records"`
	// Client ID. It is assumed to be null when unspecified.
	ClientID param.Field[int64] `query:"client_id"`
}

// URLQuery serializes [CloudV1ResellerRegionDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV1ResellerRegionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
