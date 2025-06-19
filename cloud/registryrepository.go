// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
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
func (r *RegistryRepositoryService) List(ctx context.Context, registryID int64, query RegistryRepositoryListParams, opts ...option.RequestOption) (res *RegistryRepositoryList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific repository from the container registry.
func (r *RegistryRepositoryService) Delete(ctx context.Context, repositoryName string, body RegistryRepositoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RegistryRepository struct {
	// Repository ID
	ID int64 `json:"id,required"`
	// Number of artifacts in the repository
	ArtifactCount int64 `json:"artifact_count,required"`
	// Repository creation date-time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Repository name
	Name string `json:"name,required"`
	// Number of pools from the repository
	PullCount int64 `json:"pull_count,required"`
	// Repository registry ID
	RegistryID int64 `json:"registry_id,required"`
	// Repository modification date-time
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
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
	Count int64 `json:"count,required"`
	// Objects
	Results []RegistryRepository `json:"results,required"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type RegistryRepositoryDeleteParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	RegistryID int64            `path:"registry_id,required" json:"-"`
	paramObj
}
