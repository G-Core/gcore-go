// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
func (r *LocationService) List(ctx context.Context, opts ...option.RequestOption) (res *[]StorageLocation, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/provisioning/v1/location"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StorageLocation struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	// Indicates whether new storage can be created in this location: `allow` enables
	// storage creation, `deny` prevents it
	//
	// Any of "deny", "allow".
	AllowForNewStorage StorageLocationAllowForNewStorage `json:"allow_for_new_storage"`
	// Any of "s-ed1", "s-drc2", "s-sgc1", "s-nhn2", "s-darz", "s-ws1", "ams", "sin",
	// "fra", "mia".
	Name StorageLocationName `json:"name"`
	// Any of "s3", "sftp".
	Type StorageLocationType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		AllowForNewStorage respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageLocation) RawJSON() string { return r.JSON.raw }
func (r *StorageLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether new storage can be created in this location: `allow` enables
// storage creation, `deny` prevents it
type StorageLocationAllowForNewStorage string

const (
	StorageLocationAllowForNewStorageDeny  StorageLocationAllowForNewStorage = "deny"
	StorageLocationAllowForNewStorageAllow StorageLocationAllowForNewStorage = "allow"
)

type StorageLocationName string

const (
	StorageLocationNameSEd1  StorageLocationName = "s-ed1"
	StorageLocationNameSDrc2 StorageLocationName = "s-drc2"
	StorageLocationNameSSgc1 StorageLocationName = "s-sgc1"
	StorageLocationNameSNhn2 StorageLocationName = "s-nhn2"
	StorageLocationNameSDarz StorageLocationName = "s-darz"
	StorageLocationNameSWs1  StorageLocationName = "s-ws1"
	StorageLocationNameAms   StorageLocationName = "ams"
	StorageLocationNameSin   StorageLocationName = "sin"
	StorageLocationNameFra   StorageLocationName = "fra"
	StorageLocationNameMia   StorageLocationName = "mia"
)

type StorageLocationType string

const (
	StorageLocationTypeS3   StorageLocationType = "s3"
	StorageLocationTypeSftp StorageLocationType = "sftp"
)
