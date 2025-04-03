// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2InferenceRegistryService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceRegistryService] method instead.
type CloudV2InferenceRegistryService struct {
	Options []option.RequestOption
}

// NewCloudV2InferenceRegistryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2InferenceRegistryService(opts ...option.RequestOption) (r *CloudV2InferenceRegistryService) {
	r = &CloudV2InferenceRegistryService{}
	r.Options = opts
	return
}

// Deprecated. Create Registry
func (r *CloudV2InferenceRegistryService) New(ctx context.Context, body CloudV2InferenceRegistryNewParams, opts ...option.RequestOption) (res *ImageRegistry, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/registries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deprecated. Get Registry
func (r *CloudV2InferenceRegistryService) Get(ctx context.Context, registryID string, opts ...option.RequestOption) (res *ImageRegistry, err error) {
	opts = append(r.Options[:], opts...)
	if registryID == "" {
		err = errors.New("missing required registry_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/registries/%s", registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deprecated. Update Registry
func (r *CloudV2InferenceRegistryService) Update(ctx context.Context, registryID string, body CloudV2InferenceRegistryUpdateParams, opts ...option.RequestOption) (res *ImageRegistry, err error) {
	opts = append(r.Options[:], opts...)
	if registryID == "" {
		err = errors.New("missing required registry_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/registries/%s", registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deprecated. List Registries
func (r *CloudV2InferenceRegistryService) List(ctx context.Context, query CloudV2InferenceRegistryListParams, opts ...option.RequestOption) (res *ImageRegistry, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/registries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deprecated. Delete Registry
func (r *CloudV2InferenceRegistryService) Delete(ctx context.Context, registryID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if registryID == "" {
		err = errors.New("missing required registry_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/registries/%s", registryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ImageRegistry struct {
	// Registry ID.
	ID string `json:"id,required" format:"uuid"`
	// Registry creation date in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Registry name.
	Name string `json:"name,required"`
	// Registry URL
	URL string `json:"url,required"`
	// Registry username.
	Username string            `json:"username,required"`
	JSON     imageRegistryJSON `json:"-"`
}

// imageRegistryJSON contains the JSON metadata for the struct [ImageRegistry]
type imageRegistryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageRegistry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageRegistryJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceRegistryNewParams struct {
	// Registry name.
	Name param.Field[string] `json:"name,required"`
	// Registry password
	Password param.Field[string] `json:"password,required"`
	// Project ID
	ProjectID param.Field[int64] `json:"project_id,required"`
	// Registry URL
	URL param.Field[string] `json:"url,required"`
	// Registry username.
	Username param.Field[string] `json:"username,required"`
}

func (r CloudV2InferenceRegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2InferenceRegistryUpdateParams struct {
	// Registry password.
	Password param.Field[string] `json:"password"`
	// Registry URL.
	URL param.Field[string] `json:"url"`
	// Registry username.
	Username param.Field[string] `json:"username"`
}

func (r CloudV2InferenceRegistryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2InferenceRegistryListParams struct {
	// Project ID
	ProjectID param.Field[int64] `query:"project_id"`
}

// URLQuery serializes [CloudV2InferenceRegistryListParams]'s query parameters as
// `url.Values`.
func (r CloudV2InferenceRegistryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
