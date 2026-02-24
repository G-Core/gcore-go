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

// DatabasePostgresClusterUserCredentialService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresClusterUserCredentialService] method instead.
type DatabasePostgresClusterUserCredentialService struct {
	Options []option.RequestOption
}

// NewDatabasePostgresClusterUserCredentialService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDatabasePostgresClusterUserCredentialService(opts ...option.RequestOption) (r DatabasePostgresClusterUserCredentialService) {
	r = DatabasePostgresClusterUserCredentialService{}
	r.Options = opts
	return
}

// Get the credentials for a specific user in a PostgreSQL cluster. This endpoint
// can only be used once per user.
func (r *DatabasePostgresClusterUserCredentialService) Get(ctx context.Context, username string, query DatabasePostgresClusterUserCredentialGetParams, opts ...option.RequestOption) (res *PostgresUserCredentials, err error) {
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
	if query.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s/users/%s/credentials", query.ProjectID.Value, query.RegionID.Value, query.ClusterName, username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate new credentials for a specific user in a PostgreSQL cluster.
func (r *DatabasePostgresClusterUserCredentialService) Regenerate(ctx context.Context, username string, body DatabasePostgresClusterUserCredentialRegenerateParams, opts ...option.RequestOption) (res *PostgresUserCredentials, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if body.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s/users/%s/credentials", body.ProjectID.Value, body.RegionID.Value, body.ClusterName, username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type PostgresUserCredentials struct {
	// Password
	Password string `json:"password" api:"required"`
	// Username
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Password    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresUserCredentials) RawJSON() string { return r.JSON.raw }
func (r *PostgresUserCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatabasePostgresClusterUserCredentialGetParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	ClusterName string           `path:"cluster_name" api:"required" json:"-"`
	paramObj
}

type DatabasePostgresClusterUserCredentialRegenerateParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	ClusterName string           `path:"cluster_name" api:"required" json:"-"`
	paramObj
}
