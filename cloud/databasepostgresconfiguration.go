// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DatabasePostgresConfigurationService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresConfigurationService] method instead.
type DatabasePostgresConfigurationService struct {
	Options []option.RequestOption
}

// NewDatabasePostgresConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDatabasePostgresConfigurationService(opts ...option.RequestOption) (r DatabasePostgresConfigurationService) {
	r = DatabasePostgresConfigurationService{}
	r.Options = opts
	return
}

// Get all available PostgreSQL configurations for the specified region.
func (r *DatabasePostgresConfigurationService) Get(ctx context.Context, query DatabasePostgresConfigurationGetParams, opts ...option.RequestOption) (res *PostgresConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/configuration/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PostgresConfiguration struct {
	Flavors        []PostgresConfigurationFlavor       `json:"flavors,required"`
	StorageClasses []PostgresConfigurationStorageClass `json:"storage_classes,required"`
	// Available versions
	Versions []string `json:"versions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flavors        respjson.Field
		StorageClasses respjson.Field
		Versions       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PostgresConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresConfigurationFlavor struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64  `json:"memory_gib,required"`
	PgConf    string `json:"pg_conf,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		MemoryGib   respjson.Field
		PgConf      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresConfigurationFlavor) RawJSON() string { return r.JSON.raw }
func (r *PostgresConfigurationFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresConfigurationStorageClass struct {
	// Storage type
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresConfigurationStorageClass) RawJSON() string { return r.JSON.raw }
func (r *PostgresConfigurationStorageClass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatabasePostgresConfigurationGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
