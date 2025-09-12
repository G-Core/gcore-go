// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"
	"net/url"

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

// Returns available storage locations where you can create storages. Each location
// represents a geographic region with specific data center facilities.
func (r *LocationService) List(ctx context.Context, query LocationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Location], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "storage/provisioning/v2/locations"
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

// Returns available storage locations where you can create storages. Each location
// represents a geographic region with specific data center facilities.
func (r *LocationService) ListAutoPaging(ctx context.Context, query LocationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Location] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// LocationV2 represents location data for v2 API where title is a string
type Location struct {
	// Full hostname/address for accessing the storage endpoint in this location
	Address string `json:"address,required"`
	// Indicates whether new storage can be created in this location
	//
	// Any of "deny", "allow".
	AllowForNewStorage LocationAllowForNewStorage `json:"allow_for_new_storage,required"`
	// Location code (region identifier)
	Name string `json:"name,required"`
	// Human-readable title for the location
	Title string `json:"title,required"`
	// Storage protocol type supported in this location
	//
	// Any of "s3", "sftp".
	Type LocationType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		AllowForNewStorage respjson.Field
		Name               respjson.Field
		Title              respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether new storage can be created in this location
type LocationAllowForNewStorage string

const (
	LocationAllowForNewStorageDeny  LocationAllowForNewStorage = "deny"
	LocationAllowForNewStorageAllow LocationAllowForNewStorage = "allow"
)

// Storage protocol type supported in this location
type LocationType string

const (
	LocationTypeS3   LocationType = "s3"
	LocationTypeSftp LocationType = "sftp"
)

type LocationListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LocationListParams]'s query parameters as `url.Values`.
func (r LocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
