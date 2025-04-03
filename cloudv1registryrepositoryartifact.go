// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1RegistryRepositoryArtifactService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RegistryRepositoryArtifactService] method instead.
type CloudV1RegistryRepositoryArtifactService struct {
	Options []option.RequestOption
}

// NewCloudV1RegistryRepositoryArtifactService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1RegistryRepositoryArtifactService(opts ...option.RequestOption) (r *CloudV1RegistryRepositoryArtifactService) {
	r = &CloudV1RegistryRepositoryArtifactService{}
	r.Options = opts
	return
}

// List artifacts
func (r *CloudV1RegistryRepositoryArtifactService) List(ctx context.Context, projectID int64, regionID int64, registryID int64, repositoryName string, opts ...option.RequestOption) (res *CloudV1RegistryRepositoryArtifactListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts", projectID, regionID, registryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an artifact
func (r *CloudV1RegistryRepositoryArtifactService) Delete(ctx context.Context, projectID int64, regionID int64, registryID int64, repositoryName string, digest string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	if digest == "" {
		err = errors.New("missing required digest parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts/%s", projectID, regionID, registryID, repositoryName, digest)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a tag
func (r *CloudV1RegistryRepositoryArtifactService) DeleteTag(ctx context.Context, projectID int64, regionID int64, registryID int64, repositoryName string, digest string, tagName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	if digest == "" {
		err = errors.New("missing required digest parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts/%s/tags/%s", projectID, regionID, registryID, repositoryName, digest, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CloudV1RegistryRepositoryArtifactListResponse struct {
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
	Tags []CloudV1RegistryRepositoryArtifactListResponseTag `json:"tags,required"`
	JSON cloudV1RegistryRepositoryArtifactListResponseJSON  `json:"-"`
}

// cloudV1RegistryRepositoryArtifactListResponseJSON contains the JSON metadata for
// the struct [CloudV1RegistryRepositoryArtifactListResponse]
type cloudV1RegistryRepositoryArtifactListResponseJSON struct {
	ID           apijson.Field
	Digest       apijson.Field
	PulledAt     apijson.Field
	PushedAt     apijson.Field
	RegistryID   apijson.Field
	RepositoryID apijson.Field
	Size         apijson.Field
	Tags         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CloudV1RegistryRepositoryArtifactListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RegistryRepositoryArtifactListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1RegistryRepositoryArtifactListResponseTag struct {
	// Tag ID
	ID int64 `json:"id,required"`
	// Artifact ID
	ArtifactID int64 `json:"artifact_id,required"`
	// Tag name
	Name string `json:"name,required"`
	// Tag last pull date-time
	PulledAt time.Time `json:"pulled_at,required" format:"date-time"`
	// Tag push date-time
	PushedAt time.Time `json:"pushed_at,required" format:"date-time"`
	// Repository ID
	RepositoryID int64                                                `json:"repository_id,required"`
	JSON         cloudV1RegistryRepositoryArtifactListResponseTagJSON `json:"-"`
}

// cloudV1RegistryRepositoryArtifactListResponseTagJSON contains the JSON metadata
// for the struct [CloudV1RegistryRepositoryArtifactListResponseTag]
type cloudV1RegistryRepositoryArtifactListResponseTagJSON struct {
	ID           apijson.Field
	ArtifactID   apijson.Field
	Name         apijson.Field
	PulledAt     apijson.Field
	PushedAt     apijson.Field
	RepositoryID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CloudV1RegistryRepositoryArtifactListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RegistryRepositoryArtifactListResponseTagJSON) RawJSON() string {
	return r.raw
}
