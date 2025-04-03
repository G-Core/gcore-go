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

// CloudV1RegistryRepositoryService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RegistryRepositoryService] method instead.
type CloudV1RegistryRepositoryService struct {
	Options   []option.RequestOption
	Artifacts *CloudV1RegistryRepositoryArtifactService
}

// NewCloudV1RegistryRepositoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1RegistryRepositoryService(opts ...option.RequestOption) (r *CloudV1RegistryRepositoryService) {
	r = &CloudV1RegistryRepositoryService{}
	r.Options = opts
	r.Artifacts = NewCloudV1RegistryRepositoryArtifactService(opts...)
	return
}

// List repositories
func (r *CloudV1RegistryRepositoryService) List(ctx context.Context, projectID int64, regionID int64, registryID int64, opts ...option.RequestOption) (res *CloudV1RegistryRepositoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a repository
func (r *CloudV1RegistryRepositoryService) Delete(ctx context.Context, projectID int64, regionID int64, registryID int64, repositoryName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if repositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s", projectID, regionID, registryID, repositoryName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CloudV1RegistryRepositoryListResponse struct {
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
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	JSON      cloudV1RegistryRepositoryListResponseJSON `json:"-"`
}

// cloudV1RegistryRepositoryListResponseJSON contains the JSON metadata for the
// struct [CloudV1RegistryRepositoryListResponse]
type cloudV1RegistryRepositoryListResponseJSON struct {
	ID            apijson.Field
	ArtifactCount apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	PullCount     apijson.Field
	RegistryID    apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1RegistryRepositoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RegistryRepositoryListResponseJSON) RawJSON() string {
	return r.raw
}
