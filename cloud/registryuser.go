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

// RegistryUserService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryUserService] method instead.
type RegistryUserService struct {
	Options []option.RequestOption
}

// NewRegistryUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryUserService(opts ...option.RequestOption) (r RegistryUserService) {
	r = RegistryUserService{}
	r.Options = opts
	return
}

// Create a new user for accessing the container registry.
func (r *RegistryUserService) New(ctx context.Context, registryID int64, params RegistryUserNewParams, opts ...option.RequestOption) (res *RegistryUserCreated, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the configuration of a specific registry user.
func (r *RegistryUserService) Update(ctx context.Context, userID int64, params RegistryUserUpdateParams, opts ...option.RequestOption) (res *RegistryUser, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", params.ProjectID.Value, params.RegionID.Value, params.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all users with access to the container registry.
func (r *RegistryUserService) List(ctx context.Context, registryID int64, query RegistryUserListParams, opts ...option.RequestOption) (res *RegistryUserList, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific user from the container registry.
func (r *RegistryUserService) Delete(ctx context.Context, userID int64, body RegistryUserDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create multiple users for accessing the container registry in a single request.
func (r *RegistryUserService) NewMultiple(ctx context.Context, registryID int64, params RegistryUserNewMultipleParams, opts ...option.RequestOption) (res *RegistryUserCreated, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/batch", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Generate a new secret for a specific registry user.
func (r *RegistryUserService) RefreshSecret(ctx context.Context, userID int64, body RegistryUserRefreshSecretParams, opts ...option.RequestOption) (res *RegistryUserRefreshSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v/refresh_secret", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type RegistryUser struct {
	// User ID
	ID int64 `json:"id" api:"required"`
	// User creation date-time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// User operation end date-time
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// User name
	Name string `json:"name" api:"required"`
	// Read-only user
	ReadOnly bool `json:"read_only" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Duration    respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		ReadOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUser) RawJSON() string { return r.JSON.raw }
func (r *RegistryUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserCreated struct {
	// User ID
	ID int64 `json:"id" api:"required"`
	// User creation date-time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// User operation end date-time
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// User name
	Name string `json:"name" api:"required"`
	// Read-only user
	ReadOnly bool `json:"read_only" api:"required"`
	// User secret
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Duration    respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		ReadOnly    respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUserCreated) RawJSON() string { return r.JSON.raw }
func (r *RegistryUserCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []RegistryUser `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUserList) RawJSON() string { return r.JSON.raw }
func (r *RegistryUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserRefreshSecretResponse struct {
	// User ID
	ID int64 `json:"id" api:"required"`
	// User creation date-time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// User operation end date-time
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// User name
	Name string `json:"name" api:"required"`
	// Read-only user
	ReadOnly bool `json:"read_only" api:"required"`
	// User secret
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Duration    respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		ReadOnly    respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUserRefreshSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryUserRefreshSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// A name for the registry user.
	//
	// Should be in lowercase, consisting only of numbers and letters,
	//
	// with maximum length of 16 characters
	Name string `json:"name" api:"required"`
	// Read-only user
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	// User secret
	Secret param.Opt[string] `json:"secret,omitzero"`
	paramObj
}

func (r RegistryUserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserUpdateParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	RegistryID int64            `path:"registry_id" api:"required" json:"-"`
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// Read-only user
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	paramObj
}

func (r RegistryUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type RegistryUserDeleteParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	RegistryID int64            `path:"registry_id" api:"required" json:"-"`
	paramObj
}

type RegistryUserNewMultipleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Set of users
	Users []RegistryUserNewMultipleParamsUser `json:"users,omitzero" api:"required"`
	paramObj
}

func (r RegistryUserNewMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryUserNewMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Duration, Name are required.
type RegistryUserNewMultipleParamsUser struct {
	// User account operating time, days
	Duration int64 `json:"duration" api:"required"`
	// A name for the registry user.
	//
	// Should be in lowercase, consisting only of numbers and letters,
	//
	// with maximum length of 16 characters
	Name string `json:"name" api:"required"`
	// Read-only user
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	// User secret
	Secret param.Opt[string] `json:"secret,omitzero"`
	paramObj
}

func (r RegistryUserNewMultipleParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewMultipleParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryUserNewMultipleParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserRefreshSecretParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	RegistryID int64            `path:"registry_id" api:"required" json:"-"`
	paramObj
}
