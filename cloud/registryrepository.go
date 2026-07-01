// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// RegistryRepositoryService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryRepositoryService] method instead.
type RegistryRepositoryService struct {
	Options []option.RequestOption
}

// NewRegistryRepositoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistryRepositoryService(opts ...option.RequestOption) (r RegistryRepositoryService) {
	r = RegistryRepositoryService{}
	r.Options = opts
	return
}

// List all repositories in the container registry.
func (r *RegistryRepositoryService) List(ctx context.Context, registryID int64, params RegistryRepositoryListParams, opts ...option.RequestOption) (res *RegistryRepositoryList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a specific repository from the container registry.
func (r *RegistryRepositoryService) Delete(ctx context.Context, repositoryName string, body RegistryRepositoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return err
	}
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return err
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type RegistryRepository struct {
	// Repository ID
	ID int64 `json:"id" api:"required"`
	// Number of artifacts in the repository
	ArtifactCount int64 `json:"artifact_count" api:"required"`
	// Repository creation date-time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Repository name
	Name string `json:"name" api:"required"`
	// Number of pools from the repository
	PullCount int64 `json:"pull_count" api:"required"`
	// Repository registry ID
	RegistryID int64 `json:"registry_id" api:"required"`
	// Repository modification date-time
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArtifactCount respjson.Field
		CreatedAt     respjson.Field
		Name          respjson.Field
		PullCount     respjson.Field
		RegistryID    respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryRepository) RawJSON() string { return r.JSON.raw }
func (r *RegistryRepository) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryRepositoryList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []RegistryRepository `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryRepositoryList) RawJSON() string { return r.JSON.raw }
func (r *RegistryRepositoryList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryRepositoryListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegistryRepositoryListParams]'s query parameters as
// `url.Values`.
func (r RegistryRepositoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegistryRepositoryDeleteParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	RegistryID int64            `path:"registry_id" api:"required" json:"-"`
	paramObj
}
