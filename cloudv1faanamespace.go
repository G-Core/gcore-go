// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1FaaNamespaceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FaaNamespaceService] method instead.
type CloudV1FaaNamespaceService struct {
	Options   []option.RequestOption
	Functions *CloudV1FaaNamespaceFunctionService
}

// NewCloudV1FaaNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1FaaNamespaceService(opts ...option.RequestOption) (r *CloudV1FaaNamespaceService) {
	r = &CloudV1FaaNamespaceService{}
	r.Options = opts
	r.Functions = NewCloudV1FaaNamespaceFunctionService(opts...)
	return
}

// Create namespace
func (r *CloudV1FaaNamespaceService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1FaaNamespaceNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get namespace
func (r *CloudV1FaaNamespaceService) Get(ctx context.Context, projectID int64, regionID int64, namespaceName string, opts ...option.RequestOption) (res *Namespace, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change namespace
func (r *CloudV1FaaNamespaceService) Update(ctx context.Context, projectID int64, regionID int64, namespaceName string, body CloudV1FaaNamespaceUpdateParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List namespaces
func (r *CloudV1FaaNamespaceService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1FaaNamespaceListParams, opts ...option.RequestOption) (res *CloudV1FaaNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete namespace
func (r *CloudV1FaaNamespaceService) Delete(ctx context.Context, projectID int64, regionID int64, namespaceName string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FunctionsOrderByChoices string

const (
	FunctionsOrderByChoicesCreatedAtAsc  FunctionsOrderByChoices = "created_at.asc"
	FunctionsOrderByChoicesCreatedAtDesc FunctionsOrderByChoices = "created_at.desc"
	FunctionsOrderByChoicesNameAsc       FunctionsOrderByChoices = "name.asc"
	FunctionsOrderByChoicesNameDesc      FunctionsOrderByChoices = "name.desc"
)

func (r FunctionsOrderByChoices) IsKnown() bool {
	switch r {
	case FunctionsOrderByChoicesCreatedAtAsc, FunctionsOrderByChoicesCreatedAtDesc, FunctionsOrderByChoicesNameAsc, FunctionsOrderByChoicesNameDesc:
		return true
	}
	return false
}

type Namespace struct {
	// Namespace creation date.
	CreatedAt string `json:"created_at,required"`
	// Namespace functions.
	Functions []Function `json:"functions,required"`
	// Deploy status of namespace functions.
	FunctionsDeployStatus map[string]int64 `json:"functions_deploy_status,required"`
	// Namespace name.
	Name string `json:"name,required"`
	// Namespace status.
	Status string `json:"status,required"`
	// Namespace description.
	Description string `json:"description,nullable"`
	// Namespace environment variables. Keys must match a specific regex pattern and be
	// 1-255 characters long, and values must be 0-255 characters long.
	Envs interface{}   `json:"envs,nullable"`
	JSON namespaceJSON `json:"-"`
}

// namespaceJSON contains the JSON metadata for the struct [Namespace]
type namespaceJSON struct {
	CreatedAt             apijson.Field
	Functions             apijson.Field
	FunctionsDeployStatus apijson.Field
	Name                  apijson.Field
	Status                apijson.Field
	Description           apijson.Field
	Envs                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Namespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaNamespaceListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Namespace                         `json:"results,required"`
	JSON    cloudV1FaaNamespaceListResponseJSON `json:"-"`
}

// cloudV1FaaNamespaceListResponseJSON contains the JSON metadata for the struct
// [CloudV1FaaNamespaceListResponse]
type cloudV1FaaNamespaceListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaNamespaceNewParams struct {
	// Namespace name.
	Name param.Field[string] `json:"name,required"`
	// Namespace description.
	Description param.Field[string] `json:"description"`
	// Namespace environment variables. Keys must match a specific regex pattern and be
	// 1-255 characters long, and values must be 0-255 characters long.
	Envs param.Field[interface{}] `json:"envs"`
}

func (r CloudV1FaaNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaNamespaceUpdateParams struct {
	// Namespace description.
	Description param.Field[string] `json:"description"`
	// Namespace environment variables. Keys must match a specific regex pattern and be
	// 1-255 characters long, and values must be 0-255 characters long.
	Envs param.Field[interface{}] `json:"envs"`
}

func (r CloudV1FaaNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaNamespaceListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Ordering fields and directions.
	OrderBy param.Field[[]FunctionsOrderByChoices] `query:"order_by"`
	// Search through which value
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [CloudV1FaaNamespaceListParams]'s query parameters as
// `url.Values`.
func (r CloudV1FaaNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
