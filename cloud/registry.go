// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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

// Create a new container registry with the specified configuration.
func (r *RegistryService) New(ctx context.Context, params RegistryNewParams, opts ...option.RequestOption) (res *Registry, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all container registries in the specified project and region.
func (r *RegistryService) List(ctx context.Context, query RegistryListParams, opts ...option.RequestOption) (res *RegistryList, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific container registry and all its associated resources.
func (r *RegistryService) Delete(ctx context.Context, registryID int64, body RegistryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", body.ProjectID.Value, body.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get detailed information about a specific container registry.
func (r *RegistryService) Get(ctx context.Context, registryID int64, query RegistryGetParams, opts ...option.RequestOption) (res *Registry, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the size of a container registry.
func (r *RegistryService) Resize(ctx context.Context, registryID int64, params RegistryResizeParams, opts ...option.RequestOption) (res *Registry, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/resize", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type Registry struct {
	// Registry ID
	ID int64 `json:"id" api:"required"`
	// Registry creation date-time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Registry name
	Name string `json:"name" api:"required"`
	// Number of repositories in the registry
	RepoCount int64 `json:"repo_count" api:"required"`
	// Registry storage limit, GiB
	StorageLimit int64 `json:"storage_limit" api:"required"`
	// Registry storage used, bytes
	StorageUsed int64 `json:"storage_used" api:"required"`
	// Registry modification date-time
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Registry url
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		RepoCount    respjson.Field
		StorageLimit respjson.Field
		StorageUsed  respjson.Field
		UpdatedAt    respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Registry) RawJSON() string { return r.JSON.raw }
func (r *Registry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []Registry `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryList) RawJSON() string { return r.JSON.raw }
func (r *RegistryList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryTag struct {
	// Tag ID
	ID int64 `json:"id" api:"required"`
	// Artifact ID
	ArtifactID int64 `json:"artifact_id" api:"required"`
	// Tag name
	Name string `json:"name" api:"required"`
	// Tag last pull date-time
	PulledAt time.Time `json:"pulled_at" api:"required" format:"date-time"`
	// Tag push date-time
	PushedAt time.Time `json:"pushed_at" api:"required" format:"date-time"`
	// Repository ID
	RepositoryID int64 `json:"repository_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ArtifactID   respjson.Field
		Name         respjson.Field
		PulledAt     respjson.Field
		PushedAt     respjson.Field
		RepositoryID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryTag) RawJSON() string { return r.JSON.raw }
func (r *RegistryTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// A name for the container registry.
	//
	// Should be in lowercase, consisting only of numbers, letters and -,
	//
	// with maximum length of 24 characters
	Name string `json:"name" api:"required"`
	// Registry storage limit, GiB
	StorageLimit param.Opt[int64] `json:"storage_limit,omitzero"`
	paramObj
}

func (r RegistryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type RegistryDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type RegistryGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type RegistryResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Registry storage limit, GiB
	StorageLimit param.Opt[int64] `json:"storage_limit,omitzero"`
	paramObj
}

func (r RegistryResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
