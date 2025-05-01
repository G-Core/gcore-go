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

// Create a user
func (r *RegistryUserService) New(ctx context.Context, registryID int64, params RegistryUserNewParams, opts ...option.RequestOption) (res *RegistryUser, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a user
func (r *RegistryUserService) Update(ctx context.Context, userID int64, params RegistryUserUpdateParams, opts ...option.RequestOption) (res *RegistryUser, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", params.ProjectID.Value, params.RegionID.Value, params.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get user list
func (r *RegistryUserService) List(ctx context.Context, registryID int64, query RegistryUserListParams, opts ...option.RequestOption) (res *RegistryUserList, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users", query.ProjectID.Value, query.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a user
func (r *RegistryUserService) Delete(ctx context.Context, userID int64, body RegistryUserDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Batch create users
func (r *RegistryUserService) NewMultiple(ctx context.Context, registryID int64, params RegistryUserNewMultipleParams, opts ...option.RequestOption) (res *RegistryUserCreated, err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/batch", params.ProjectID.Value, params.RegionID.Value, registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Refresh a secret
func (r *RegistryUserService) RefreshSecret(ctx context.Context, userID int64, body RegistryUserRefreshSecretParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/users/%v/refresh_secret", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// '#/components/schemas/RegistryUserSerializer'
// "$.components.schemas.RegistryUserSerializer"
type RegistryUser struct {
	// '#/components/schemas/RegistryUserSerializer/properties/id'
	// "$.components.schemas.RegistryUserSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistryUserSerializer/properties/created_at'
	// "$.components.schemas.RegistryUserSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RegistryUserSerializer/properties/duration'
	// "$.components.schemas.RegistryUserSerializer.properties.duration"
	Duration int64 `json:"duration,required"`
	// '#/components/schemas/RegistryUserSerializer/properties/expires_at'
	// "$.components.schemas.RegistryUserSerializer.properties.expires_at"
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// '#/components/schemas/RegistryUserSerializer/properties/name'
	// "$.components.schemas.RegistryUserSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryUserSerializer/properties/read_only'
	// "$.components.schemas.RegistryUserSerializer.properties.read_only"
	ReadOnly bool `json:"read_only"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		CreatedAt   resp.Field
		Duration    resp.Field
		ExpiresAt   resp.Field
		Name        resp.Field
		ReadOnly    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUser) RawJSON() string { return r.JSON.raw }
func (r *RegistryUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryUserCreateResponseSerializer'
// "$.components.schemas.RegistryUserCreateResponseSerializer"
type RegistryUserCreated struct {
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/id'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/created_at'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/duration'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.duration"
	Duration int64 `json:"duration,required"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/expires_at'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.expires_at"
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/name'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/read_only'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.read_only"
	ReadOnly bool `json:"read_only"`
	// '#/components/schemas/RegistryUserCreateResponseSerializer/properties/secret'
	// "$.components.schemas.RegistryUserCreateResponseSerializer.properties.secret"
	Secret string `json:"secret"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		CreatedAt   resp.Field
		Duration    resp.Field
		ExpiresAt   resp.Field
		Name        resp.Field
		ReadOnly    resp.Field
		Secret      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryUserCreated) RawJSON() string { return r.JSON.raw }
func (r *RegistryUserCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegistryUserCollectionSerializer'
// "$.components.schemas.RegistryUserCollectionSerializer"
type RegistryUserList struct {
	// '#/components/schemas/RegistryUserCollectionSerializer/properties/count'
	// "$.components.schemas.RegistryUserCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RegistryUserCollectionSerializer/properties/results'
	// "$.components.schemas.RegistryUserCollectionSerializer.properties.results"
	Results []RegistryUser `json:"results,required"`
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
func (r RegistryUserList) RawJSON() string { return r.JSON.raw }
func (r *RegistryUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryUserNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers/post/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers/post/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/duration'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.duration"
	Duration int64 `json:"duration,required"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/name'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/read_only'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.read_only"
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/secret'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.secret"
	Secret param.Opt[string] `json:"secret,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RegistryUserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type RegistryUserUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/patch/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}'].patch.parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	// '#/components/schemas/RegistryUserUpdateSerializer/properties/duration'
	// "$.components.schemas.RegistryUserUpdateSerializer.properties.duration"
	Duration int64 `json:"duration,required"`
	// '#/components/schemas/RegistryUserUpdateSerializer/properties/read_only'
	// "$.components.schemas.RegistryUserUpdateSerializer.properties.read_only"
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RegistryUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type RegistryUserListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers/get/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers/get/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryUserDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D/delete/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}']['delete'].parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type RegistryUserNewMultipleParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2Fbatch/post/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/batch'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2Fbatch/post/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/batch'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RegistryBatchUsersCreateSerializer/properties/users'
	// "$.components.schemas.RegistryBatchUsersCreateSerializer.properties.users"
	Users []RegistryUserNewMultipleParamsUser `json:"users,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserNewMultipleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RegistryUserNewMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/RegistryBatchUsersCreateSerializer/properties/users/items'
// "$.components.schemas.RegistryBatchUsersCreateSerializer.properties.users.items"
//
// The properties Duration, Name are required.
type RegistryUserNewMultipleParamsUser struct {
	// '#/components/schemas/RegistryUserCreateSerializer/properties/duration'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.duration"
	Duration int64 `json:"duration,required"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/name'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/read_only'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.read_only"
	ReadOnly param.Opt[bool] `json:"read_only,omitzero"`
	// '#/components/schemas/RegistryUserCreateSerializer/properties/secret'
	// "$.components.schemas.RegistryUserCreateSerializer.properties.secret"
	Secret param.Opt[string] `json:"secret,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserNewMultipleParamsUser) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RegistryUserNewMultipleParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow RegistryUserNewMultipleParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}

type RegistryUserRefreshSecretParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D%2Frefresh_secret/post/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}/refresh_secret'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D%2Frefresh_secret/post/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}/refresh_secret'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Fusers%2F%7Buser_id%7D%2Frefresh_secret/post/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/users/{user_id}/refresh_secret'].post.parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryUserRefreshSecretParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
