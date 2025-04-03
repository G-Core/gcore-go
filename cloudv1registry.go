// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1RegistryService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RegistryService] method instead.
type CloudV1RegistryService struct {
	Options      []option.RequestOption
	Repositories *CloudV1RegistryRepositoryService
	Users        *CloudV1RegistryUserService
}

// NewCloudV1RegistryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1RegistryService(opts ...option.RequestOption) (r *CloudV1RegistryService) {
	r = &CloudV1RegistryService{}
	r.Options = opts
	r.Repositories = NewCloudV1RegistryRepositoryService(opts...)
	r.Users = NewCloudV1RegistryUserService(opts...)
	return
}

// Create a registry
func (r *CloudV1RegistryService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1RegistryNewParams, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a registry
func (r *CloudV1RegistryService) Get(ctx context.Context, projectID int64, regionID int64, registryID int64, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get registry list
func (r *CloudV1RegistryService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a registry
func (r *CloudV1RegistryService) Delete(ctx context.Context, projectID int64, regionID int64, registryID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Check the quota required to create a registry
func (r *CloudV1RegistryService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1RegistryCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resize a registry
func (r *CloudV1RegistryService) Resize(ctx context.Context, projectID int64, regionID int64, registryID int64, body CloudV1RegistryResizeParams, opts ...option.RequestOption) (res *Registry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/resize", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Registry struct {
	// Registry ID
	ID int64 `json:"id,required"`
	// Registry creation date-time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Registry name
	Name string `json:"name,required"`
	// Number of repositories in the registry
	RepoCount int64 `json:"repo_count,required"`
	// Registry storage limit, GiB
	StorageLimit int64 `json:"storage_limit,required"`
	// Registry storage used, bytes
	StorageUsed int64 `json:"storage_used,required"`
	// Registry modification date-time
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Registry url
	URL  string       `json:"url,required"`
	JSON registryJSON `json:"-"`
}

// registryJSON contains the JSON metadata for the struct [Registry]
type registryJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	RepoCount    apijson.Field
	StorageLimit apijson.Field
	StorageUsed  apijson.Field
	UpdatedAt    apijson.Field
	URL          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Registry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryJSON) RawJSON() string {
	return r.raw
}

type CloudV1RegistryNewParams struct {
	// A name for the container registry.
	//
	// Should be in lowercase, consisting only of numbers, letters and -,
	//
	// with maximum length of 24 characters
	Name param.Field[string] `json:"name,required"`
	// Registry storage limit, GiB
	StorageLimit param.Field[int64] `json:"storage_limit"`
}

func (r CloudV1RegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RegistryCheckQuotaParams struct {
	// Registry storage limit, GiB
	StorageLimit param.Field[int64] `json:"storage_limit"`
}

func (r CloudV1RegistryCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RegistryResizeParams struct {
	// Registry storage limit, GiB
	StorageLimit param.Field[int64] `json:"storage_limit"`
}

func (r CloudV1RegistryResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
