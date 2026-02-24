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

// DatabasePostgresCustomConfigurationService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresCustomConfigurationService] method instead.
type DatabasePostgresCustomConfigurationService struct {
	Options []option.RequestOption
}

// NewDatabasePostgresCustomConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDatabasePostgresCustomConfigurationService(opts ...option.RequestOption) (r DatabasePostgresCustomConfigurationService) {
	r = DatabasePostgresCustomConfigurationService{}
	r.Options = opts
	return
}

// Validate a custom PostgreSQL configuration file.
func (r *DatabasePostgresCustomConfigurationService) Validate(ctx context.Context, params DatabasePostgresCustomConfigurationValidateParams, opts ...option.RequestOption) (res *PgConfValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/validate_pg_conf/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PgConfValidation struct {
	// Errors list
	Errors []string `json:"errors" api:"required"`
	// Validity of pg.conf file
	IsValid bool `json:"is_valid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		IsValid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PgConfValidation) RawJSON() string { return r.JSON.raw }
func (r *PgConfValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatabasePostgresCustomConfigurationValidateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// PostgreSQL configuration
	PgConf string `json:"pg_conf" api:"required"`
	// PostgreSQL version
	Version string `json:"version" api:"required"`
	paramObj
}

func (r DatabasePostgresCustomConfigurationValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresCustomConfigurationValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresCustomConfigurationValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
