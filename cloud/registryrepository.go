// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
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

// List repositories
func (r *RegistryRepositoryService) List(ctx context.Context, registryID int64, query RegistryRepositoryListParams, opts ...option.RequestOption) (res *RegistryRepositoryList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a repository
func (r *RegistryRepositoryService) Delete(ctx context.Context, repositoryName string, body RegistryRepositoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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

// '#/components/schemas/RegistryRepositorySerializer'
// "$.components.schemas.RegistryRepositorySerializer"
type RegistryRepository struct {
	// '#/components/schemas/RegistryRepositorySerializer/properties/id'
	// "$.components.schemas.RegistryRepositorySerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/artifact_count'
	// "$.components.schemas.RegistryRepositorySerializer.properties.artifact_count"
	ArtifactCount int64 `json:"artifact_count,required"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/created_at'
	// "$.components.schemas.RegistryRepositorySerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/name'
	// "$.components.schemas.RegistryRepositorySerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/pull_count'
	// "$.components.schemas.RegistryRepositorySerializer.properties.pull_count"
	PullCount int64 `json:"pull_count,required"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/registry_id'
	// "$.components.schemas.RegistryRepositorySerializer.properties.registry_id"
	RegistryID int64 `json:"registry_id,required"`
	// '#/components/schemas/RegistryRepositorySerializer/properties/updated_at'
	// "$.components.schemas.RegistryRepositorySerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID            resp.Field
		ArtifactCount resp.Field
		CreatedAt     resp.Field
		Name          resp.Field
		PullCount     resp.Field
		RegistryID    resp.Field
		UpdatedAt     resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryRepository) RawJSON() string { return r.JSON.raw }
func (r *RegistryRepository) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryRepositoryCollectionSerializer'
// "$.components.schemas.RegistryRepositoryCollectionSerializer"
type RegistryRepositoryList struct {
	// '#/components/schemas/RegistryRepositoryCollectionSerializer/properties/count'
	// "$.components.schemas.RegistryRepositoryCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RegistryRepositoryCollectionSerializer/properties/results'
	// "$.components.schemas.RegistryRepositoryCollectionSerializer.properties.results"
	Results []RegistryRepository `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryRepositoryList) RawJSON() string { return r.JSON.raw }
func (r *RegistryRepositoryList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryRepositoryListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories/get/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories/get/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryRepositoryListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryRepositoryDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D/delete/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}']['delete'].parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryRepositoryDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
