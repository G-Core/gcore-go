// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LocationService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r LocationService) {
	r = LocationService{}
	r.Options = opts
	return
}

// Returns storage locations where you can create new storages. Only locations
// currently accepting new storage creation are returned.
func (r *LocationService) List(ctx context.Context, query LocationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Location], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/v4/locations"
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

// Returns storage locations where you can create new storages. Only locations
// currently accepting new storage creation are returned.
func (r *LocationService) ListAutoPaging(ctx context.Context, query LocationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Location] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type Location struct {
	// Full hostname/address for accessing the storage endpoint.
	Address string `json:"address" api:"required"`
	// Human-readable display name for the location.
	Name string `json:"name" api:"required"`
	// Internal technical identifier for the location
	TechnicalName string `json:"technical_name" api:"required"`
	// Display title for the location (English). Null if no title is set.
	Title string `json:"title" api:"required"`
	// Storage type supported by this location
	//
	// Any of "s3_compatible", "sftp".
	Type LocationType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address       respjson.Field
		Name          respjson.Field
		TechnicalName respjson.Field
		Title         respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage type supported by this location
type LocationType string

const (
	LocationTypeS3Compatible LocationType = "s3_compatible"
	LocationTypeSftp         LocationType = "sftp"
)

type LocationListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip before beginning to return results
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LocationListParams]'s query parameters as `url.Values`.
func (r LocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
