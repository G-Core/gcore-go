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

// CloudV1RegistryUserService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1RegistryUserService] method instead.
type CloudV1RegistryUserService struct {
	Options []option.RequestOption
}

// NewCloudV1RegistryUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1RegistryUserService(opts ...option.RequestOption) (r *CloudV1RegistryUserService) {
	r = &CloudV1RegistryUserService{}
	r.Options = opts
	return
}

// Create a user
func (r *CloudV1RegistryUserService) New(ctx context.Context, projectID int64, regionID int64, registryID int64, body CloudV1RegistryUserNewParams, opts ...option.RequestOption) (res *RegistryUser, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a user
func (r *CloudV1RegistryUserService) Update(ctx context.Context, projectID int64, regionID int64, registryID int64, userID int64, body CloudV1RegistryUserUpdateParams, opts ...option.RequestOption) (res *RegistryUser, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", projectID, regionID, registryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get user list
func (r *CloudV1RegistryUserService) List(ctx context.Context, projectID int64, regionID int64, registryID int64, opts ...option.RequestOption) (res *RegistryUser, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a user
func (r *CloudV1RegistryUserService) Delete(ctx context.Context, projectID int64, regionID int64, registryID int64, userID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", projectID, regionID, registryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Batch create users
func (r *CloudV1RegistryUserService) BatchNew(ctx context.Context, projectID int64, regionID int64, registryID int64, body CloudV1RegistryUserBatchNewParams, opts ...option.RequestOption) (res *CloudV1RegistryUserBatchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/batch", projectID, regionID, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Refresh a secret
func (r *CloudV1RegistryUserService) RefreshSecret(ctx context.Context, projectID int64, regionID int64, registryID int64, userID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v/refresh_secret", projectID, regionID, registryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type RegistryUser struct {
	// User ID
	ID int64 `json:"id,required"`
	// User creation date-time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User account operating time, days
	Duration int64 `json:"duration,required"`
	// User operation end date-time
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// User name
	Name string `json:"name,required"`
	// Read-only user
	ReadOnly bool             `json:"read_only"`
	JSON     registryUserJSON `json:"-"`
}

// registryUserJSON contains the JSON metadata for the struct [RegistryUser]
type registryUserJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	Name        apijson.Field
	ReadOnly    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryUserJSON) RawJSON() string {
	return r.raw
}

type RegistryUserCreateParam struct {
	// User account operating time, days
	Duration param.Field[int64] `json:"duration,required"`
	// A name for the registry user.
	//
	// Should be in lowercase, consisting only of numbers and letters,
	//
	// with maximum length of 16 characters
	Name param.Field[string] `json:"name,required"`
	// Read-only user
	ReadOnly param.Field[bool] `json:"read_only"`
	// User secret
	Secret param.Field[string] `json:"secret"`
}

func (r RegistryUserCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RegistryUserBatchNewResponse struct {
	// User ID
	ID int64 `json:"id,required"`
	// User creation date-time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User account operating time, days
	Duration int64 `json:"duration,required"`
	// User operation end date-time
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// User name
	Name string `json:"name,required"`
	// Read-only user
	ReadOnly bool `json:"read_only"`
	// User secret
	Secret string                                  `json:"secret"`
	JSON   cloudV1RegistryUserBatchNewResponseJSON `json:"-"`
}

// cloudV1RegistryUserBatchNewResponseJSON contains the JSON metadata for the
// struct [CloudV1RegistryUserBatchNewResponse]
type cloudV1RegistryUserBatchNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	Name        apijson.Field
	ReadOnly    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1RegistryUserBatchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1RegistryUserBatchNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1RegistryUserNewParams struct {
	RegistryUserCreate RegistryUserCreateParam `json:"registry_user_create,required"`
}

func (r CloudV1RegistryUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RegistryUserCreate)
}

type CloudV1RegistryUserUpdateParams struct {
	// User account operating time, days
	Duration param.Field[int64] `json:"duration,required"`
	// Read-only user
	ReadOnly param.Field[bool] `json:"read_only"`
}

func (r CloudV1RegistryUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1RegistryUserBatchNewParams struct {
	// Set of users
	Users param.Field[[]RegistryUserCreateParam] `json:"users,required"`
}

func (r CloudV1RegistryUserBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
