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

// RegistryArtifactService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryArtifactService] method instead.
type RegistryArtifactService struct {
	Options []option.RequestOption
}

// NewRegistryArtifactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistryArtifactService(opts ...option.RequestOption) (r RegistryArtifactService) {
	r = RegistryArtifactService{}
	r.Options = opts
	return
}

// List artifacts
func (r *RegistryArtifactService) List(ctx context.Context, repositoryName string, query RegistryArtifactListParams, opts ...option.RequestOption) (res *RegistryArtifactList, err error) {
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
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts", query.ProjectID.Value, query.RegionID.Value, query.RegistryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an artifact
func (r *RegistryArtifactService) Delete(ctx context.Context, digest string, body RegistryArtifactDeleteParams, opts ...option.RequestOption) (err error) {
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
	if body.RepositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	if digest == "" {
		err = errors.New("missing required digest parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts/%s", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, body.RepositoryName, digest)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// '#/components/schemas/RegistryArtifactSerializer'
// "$.components.schemas.RegistryArtifactSerializer"
type RegistryArtifact struct {
	// '#/components/schemas/RegistryArtifactSerializer/properties/id'
	// "$.components.schemas.RegistryArtifactSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/digest'
	// "$.components.schemas.RegistryArtifactSerializer.properties.digest"
	Digest string `json:"digest,required"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/pulled_at'
	// "$.components.schemas.RegistryArtifactSerializer.properties.pulled_at"
	PulledAt time.Time `json:"pulled_at,required" format:"date-time"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/pushed_at'
	// "$.components.schemas.RegistryArtifactSerializer.properties.pushed_at"
	PushedAt time.Time `json:"pushed_at,required" format:"date-time"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/registry_id'
	// "$.components.schemas.RegistryArtifactSerializer.properties.registry_id"
	RegistryID int64 `json:"registry_id,required"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/repository_id'
	// "$.components.schemas.RegistryArtifactSerializer.properties.repository_id"
	RepositoryID int64 `json:"repository_id,required"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/size'
	// "$.components.schemas.RegistryArtifactSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/RegistryArtifactSerializer/properties/tags'
	// "$.components.schemas.RegistryArtifactSerializer.properties.tags"
	Tags []RegistryTag `json:"tags,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		Digest       resp.Field
		PulledAt     resp.Field
		PushedAt     resp.Field
		RegistryID   resp.Field
		RepositoryID resp.Field
		Size         resp.Field
		Tags         resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryArtifact) RawJSON() string { return r.JSON.raw }
func (r *RegistryArtifact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryArtifactCollectionSerializer'
// "$.components.schemas.RegistryArtifactCollectionSerializer"
type RegistryArtifactList struct {
	// '#/components/schemas/RegistryArtifactCollectionSerializer/properties/count'
	// "$.components.schemas.RegistryArtifactCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RegistryArtifactCollectionSerializer/properties/results'
	// "$.components.schemas.RegistryArtifactCollectionSerializer.properties.results"
	Results []RegistryArtifact `json:"results,required"`
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
func (r RegistryArtifactList) RawJSON() string { return r.JSON.raw }
func (r *RegistryArtifactList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryArtifactListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts/get/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts/get/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts/get/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts'].get.parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryArtifactListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryArtifactDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D/delete/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}']['delete'].parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D/delete/parameters/3/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}']['delete'].parameters[3].schema"
	RepositoryName string `path:"repository_name,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryArtifactDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
