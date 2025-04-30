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

// RegistryService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryService] method instead.
type RegistryService struct {
	Options      []option.RequestOption
	Repositories RegistryRepositoryService
	Artifacts    RegistryArtifactService
	Tags         RegistryTagService
	Users        RegistryUserService
}

// NewRegistryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryService(opts ...option.RequestOption) (r RegistryService) {
	r = RegistryService{}
	r.Options = opts
	r.Repositories = NewRegistryRepositoryService(opts...)
	r.Artifacts = NewRegistryArtifactService(opts...)
	r.Tags = NewRegistryTagService(opts...)
	r.Users = NewRegistryUserService(opts...)
	return
}

// Create a registry
func (r *RegistryService) New(ctx context.Context, params RegistryNewParams, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get registry list
func (r *RegistryService) List(ctx context.Context, query RegistryListParams, opts ...option.RequestOption) (res *RegistryList, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a registry
func (r *RegistryService) Delete(ctx context.Context, registryID int64, body RegistryDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", body.ProjectID.Value, body.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a registry
func (r *RegistryService) Get(ctx context.Context, registryID int64, query RegistryGetParams, opts ...option.RequestOption) (res *Registry, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resize a registry
func (r *RegistryService) Resize(ctx context.Context, registryID int64, params RegistryResizeParams, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/resize", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// '#/components/schemas/RegistrySerializer'
// "$.components.schemas.RegistrySerializer"
type Registry struct {
	// '#/components/schemas/RegistrySerializer/properties/id'
	// "$.components.schemas.RegistrySerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistrySerializer/properties/created_at'
	// "$.components.schemas.RegistrySerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RegistrySerializer/properties/name'
	// "$.components.schemas.RegistrySerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistrySerializer/properties/repo_count'
	// "$.components.schemas.RegistrySerializer.properties.repo_count"
	RepoCount int64 `json:"repo_count,required"`
	// '#/components/schemas/RegistrySerializer/properties/storage_limit'
	// "$.components.schemas.RegistrySerializer.properties.storage_limit"
	StorageLimit int64 `json:"storage_limit,required"`
	// '#/components/schemas/RegistrySerializer/properties/storage_used'
	// "$.components.schemas.RegistrySerializer.properties.storage_used"
	StorageUsed int64 `json:"storage_used,required"`
	// '#/components/schemas/RegistrySerializer/properties/updated_at'
	// "$.components.schemas.RegistrySerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/RegistrySerializer/properties/url'
	// "$.components.schemas.RegistrySerializer.properties.url"
	URL string `json:"url,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		CreatedAt    resp.Field
		Name         resp.Field
		RepoCount    resp.Field
		StorageLimit resp.Field
		StorageUsed  resp.Field
		UpdatedAt    resp.Field
		URL          resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Registry) RawJSON() string { return r.JSON.raw }
func (r *Registry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryCollectionSerializer'
// "$.components.schemas.RegistryCollectionSerializer"
type RegistryList struct {
	// '#/components/schemas/RegistryCollectionSerializer/properties/count'
	// "$.components.schemas.RegistryCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RegistryCollectionSerializer/properties/results'
	// "$.components.schemas.RegistryCollectionSerializer.properties.results"
	Results []Registry `json:"results,required"`
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
func (r RegistryList) RawJSON() string { return r.JSON.raw }
func (r *RegistryList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryTagSerializer'
// "$.components.schemas.RegistryTagSerializer"
type RegistryTag struct {
	// '#/components/schemas/RegistryTagSerializer/properties/id'
	// "$.components.schemas.RegistryTagSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistryTagSerializer/properties/artifact_id'
	// "$.components.schemas.RegistryTagSerializer.properties.artifact_id"
	ArtifactID int64 `json:"artifact_id,required"`
	// '#/components/schemas/RegistryTagSerializer/properties/name'
	// "$.components.schemas.RegistryTagSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryTagSerializer/properties/pulled_at'
	// "$.components.schemas.RegistryTagSerializer.properties.pulled_at"
	PulledAt time.Time `json:"pulled_at,required" format:"date-time"`
	// '#/components/schemas/RegistryTagSerializer/properties/pushed_at'
	// "$.components.schemas.RegistryTagSerializer.properties.pushed_at"
	PushedAt time.Time `json:"pushed_at,required" format:"date-time"`
	// '#/components/schemas/RegistryTagSerializer/properties/repository_id'
	// "$.components.schemas.RegistryTagSerializer.properties.repository_id"
	RepositoryID int64 `json:"repository_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		ArtifactID   resp.Field
		Name         resp.Field
		PulledAt     resp.Field
		PushedAt     resp.Field
		RepositoryID resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryTag) RawJSON() string { return r.JSON.raw }
func (r *RegistryTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RegistryCreateSerializer/properties/name'
	// "$.components.schemas.RegistryCreateSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryCreateSerializer/properties/storage_limit'
	// "$.components.schemas.RegistryCreateSerializer.properties.storage_limit"
	StorageLimit param.Opt[int64] `json:"storage_limit,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RegistryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type RegistryListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fresize/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/resize'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fresize/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/resize'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RegistryResizeSerializer/properties/storage_limit'
	// "$.components.schemas.RegistryResizeSerializer.properties.storage_limit"
	StorageLimit param.Opt[int64] `json:"storage_limit,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RegistryResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
