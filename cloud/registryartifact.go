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

// List all artifacts in a specific repository.
func (r *RegistryArtifactService) List(ctx context.Context, repositoryName string, query RegistryArtifactListParams, opts ...option.RequestOption) (res *RegistryArtifactList, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts", query.ProjectID.Value, query.RegionID.Value, query.RegistryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific artifact from a repository.
func (r *RegistryArtifactService) Delete(ctx context.Context, digest string, body RegistryArtifactDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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

type RegistryArtifact struct {
	// Repository ID
	ID int64 `json:"id,required"`
	// Artifact digest
	Digest string `json:"digest,required"`
	// Artifact last pull date-time
	PulledAt time.Time `json:"pulled_at,required" format:"date-time"`
	// Artifact push date-time
	PushedAt time.Time `json:"pushed_at,required" format:"date-time"`
	// Artifact registry ID
	RegistryID int64 `json:"registry_id,required"`
	// Artifact repository ID
	RepositoryID int64 `json:"repository_id,required"`
	// Artifact size, bytes
	Size int64 `json:"size,required"`
	// Artifact tags
	Tags []RegistryTag `json:"tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Digest       respjson.Field
		PulledAt     respjson.Field
		PushedAt     respjson.Field
		RegistryID   respjson.Field
		RepositoryID respjson.Field
		Size         respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryArtifact) RawJSON() string { return r.JSON.raw }
func (r *RegistryArtifact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryArtifactList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []RegistryArtifact `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryArtifactList) RawJSON() string { return r.JSON.raw }
func (r *RegistryArtifactList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryArtifactListParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	RegistryID int64            `path:"registry_id,required" json:"-"`
	paramObj
}

type RegistryArtifactDeleteParams struct {
	ProjectID      param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID       param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	RegistryID     int64            `path:"registry_id,required" json:"-"`
	RepositoryName string           `path:"repository_name,required" json:"-"`
	paramObj
}
