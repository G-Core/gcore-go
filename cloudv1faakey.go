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

// CloudV1FaaKeyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FaaKeyService] method instead.
type CloudV1FaaKeyService struct {
	Options []option.RequestOption
}

// NewCloudV1FaaKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1FaaKeyService(opts ...option.RequestOption) (r *CloudV1FaaKeyService) {
	r = &CloudV1FaaKeyService{}
	r.Options = opts
	return
}

// Create api key
func (r *CloudV1FaaKeyService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1FaaKeyNewParams, opts ...option.RequestOption) (res *CloudV1FaaKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/keys/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get api key
func (r *CloudV1FaaKeyService) Get(ctx context.Context, projectID int64, regionID int64, keyName string, opts ...option.RequestOption) (res *FaasAPIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change api key
func (r *CloudV1FaaKeyService) Update(ctx context.Context, projectID int64, regionID int64, keyName string, body CloudV1FaaKeyUpdateParams, opts ...option.RequestOption) (res *FaasAPIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List api keys
func (r *CloudV1FaaKeyService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1FaaKeyListParams, opts ...option.RequestOption) (res *CloudV1FaaKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/keys/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete api key
func (r *CloudV1FaaKeyService) Delete(ctx context.Context, projectID int64, regionID int64, keyName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FaasAPIKey struct {
	// Key starts to work
	CreatedAt string `json:"created_at,required"`
	// List of functions using key
	Functions []NamespacedName `json:"functions,required"`
	// Key name
	Name string `json:"name,required"`
	// Key status
	Status string `json:"status,required"`
	// Optional key description
	Description string `json:"description,nullable"`
	// Key expires at
	Expire time.Time      `json:"expire,nullable" format:"date-time"`
	JSON   faasAPIKeyJSON `json:"-"`
}

// faasAPIKeyJSON contains the JSON metadata for the struct [FaasAPIKey]
type faasAPIKeyJSON struct {
	CreatedAt   apijson.Field
	Functions   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Description apijson.Field
	Expire      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FaasAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r faasAPIKeyJSON) RawJSON() string {
	return r.raw
}

type NamespacedName struct {
	// Function name
	Name string `json:"name,required"`
	// Namespace name
	Namespace string             `json:"namespace,required"`
	JSON      namespacedNameJSON `json:"-"`
}

// namespacedNameJSON contains the JSON metadata for the struct [NamespacedName]
type namespacedNameJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespacedName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespacedNameJSON) RawJSON() string {
	return r.raw
}

type NamespacedNameParam struct {
	// Function name
	Name param.Field[string] `json:"name,required"`
	// Namespace name
	Namespace param.Field[string] `json:"namespace,required"`
}

func (r NamespacedNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaKeyNewResponse struct {
	// Key starts to work
	CreatedAt string `json:"created_at,required"`
	// List of functions using key
	Functions []NamespacedName `json:"functions,required"`
	// Key name
	Name string `json:"name,required"`
	// Secret api key
	Secret string `json:"secret,required"`
	// Key status
	Status string `json:"status,required"`
	// Optional key description
	Description string `json:"description,nullable"`
	// Key expires at
	Expire time.Time                    `json:"expire,nullable" format:"date-time"`
	JSON   cloudV1FaaKeyNewResponseJSON `json:"-"`
}

// cloudV1FaaKeyNewResponseJSON contains the JSON metadata for the struct
// [CloudV1FaaKeyNewResponse]
type cloudV1FaaKeyNewResponseJSON struct {
	CreatedAt   apijson.Field
	Functions   apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Status      apijson.Field
	Description apijson.Field
	Expire      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaKeyListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []FaasAPIKey                  `json:"results,required"`
	JSON    cloudV1FaaKeyListResponseJSON `json:"-"`
}

// cloudV1FaaKeyListResponseJSON contains the JSON metadata for the struct
// [CloudV1FaaKeyListResponse]
type cloudV1FaaKeyListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaKeyNewParams struct {
	// List of functions using key
	Functions param.Field[[]NamespacedNameParam] `json:"functions,required"`
	// Key name
	Name param.Field[string] `json:"name,required"`
	// Optional key description
	Description param.Field[string] `json:"description"`
	// Key expires at
	Expire param.Field[time.Time] `json:"expire" format:"date-time"`
}

func (r CloudV1FaaKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaKeyUpdateParams struct {
	// List of functions using key
	Functions param.Field[[]NamespacedNameParam] `json:"functions,required"`
	// Optional key description
	Description param.Field[string] `json:"description"`
	// Key expires at
	Expire param.Field[time.Time] `json:"expire" format:"date-time"`
}

func (r CloudV1FaaKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaKeyListParams struct {
	// Days to expire
	ExpireIn param.Field[int64] `query:"expire_in"`
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Ordering fields and directions.
	OrderBy param.Field[[]CloudV1FaaKeyListParamsOrderBy] `query:"order_by"`
	// Search through which value
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [CloudV1FaaKeyListParams]'s query parameters as
// `url.Values`.
func (r CloudV1FaaKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1FaaKeyListParamsOrderBy string

const (
	CloudV1FaaKeyListParamsOrderByCreatedAtAsc  CloudV1FaaKeyListParamsOrderBy = "created_at.asc"
	CloudV1FaaKeyListParamsOrderByCreatedAtDesc CloudV1FaaKeyListParamsOrderBy = "created_at.desc"
	CloudV1FaaKeyListParamsOrderByExpireAsc     CloudV1FaaKeyListParamsOrderBy = "expire.asc"
	CloudV1FaaKeyListParamsOrderByExpireDesc    CloudV1FaaKeyListParamsOrderBy = "expire.desc"
	CloudV1FaaKeyListParamsOrderByNameAsc       CloudV1FaaKeyListParamsOrderBy = "name.asc"
	CloudV1FaaKeyListParamsOrderByNameDesc      CloudV1FaaKeyListParamsOrderBy = "name.desc"
)

func (r CloudV1FaaKeyListParamsOrderBy) IsKnown() bool {
	switch r {
	case CloudV1FaaKeyListParamsOrderByCreatedAtAsc, CloudV1FaaKeyListParamsOrderByCreatedAtDesc, CloudV1FaaKeyListParamsOrderByExpireAsc, CloudV1FaaKeyListParamsOrderByExpireDesc, CloudV1FaaKeyListParamsOrderByNameAsc, CloudV1FaaKeyListParamsOrderByNameDesc:
		return true
	}
	return false
}
