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

// CloudV1CaaKeyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1CaaKeyService] method instead.
type CloudV1CaaKeyService struct {
	Options []option.RequestOption
}

// NewCloudV1CaaKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1CaaKeyService(opts ...option.RequestOption) (r *CloudV1CaaKeyService) {
	r = &CloudV1CaaKeyService{}
	r.Options = opts
	return
}

// Create api key
func (r *CloudV1CaaKeyService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1CaaKeyNewParams, opts ...option.RequestOption) (res *CloudV1CaaKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/keys/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get api key
func (r *CloudV1CaaKeyService) Get(ctx context.Context, projectID int64, regionID int64, keyName string, opts ...option.RequestOption) (res *CaasAPIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change api key
func (r *CloudV1CaaKeyService) Update(ctx context.Context, projectID int64, regionID int64, keyName string, body CloudV1CaaKeyUpdateParams, opts ...option.RequestOption) (res *CaasAPIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List api keys
func (r *CloudV1CaaKeyService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1CaaKeyListParams, opts ...option.RequestOption) (res *CloudV1CaaKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/keys/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete api key
func (r *CloudV1CaaKeyService) Delete(ctx context.Context, projectID int64, regionID int64, keyName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/keys/%v/%v/%s", projectID, regionID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CaasAPIKey struct {
	// List of containers using key
	Containers []CaasNamespacedName `json:"containers,required"`
	// Key starts to work
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Key description
	Description string `json:"description,required"`
	// Key expires at
	Expire time.Time `json:"expire,required,nullable" format:"date-time"`
	// Set to true if key is disabled
	IsDisabled bool `json:"is_disabled,required"`
	// Key name
	Name string `json:"name,required"`
	// Source of the apikey, can be 'cloud' or 'iate'
	Source string `json:"source,required"`
	// Key status
	Status string         `json:"status,required"`
	JSON   caasAPIKeyJSON `json:"-"`
}

// caasAPIKeyJSON contains the JSON metadata for the struct [CaasAPIKey]
type caasAPIKeyJSON struct {
	Containers  apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Expire      apijson.Field
	IsDisabled  apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaasAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caasAPIKeyJSON) RawJSON() string {
	return r.raw
}

type CaasNamespacedName struct {
	// Container name
	Name string `json:"name,required"`
	// Namespace name
	Namespace string                 `json:"namespace,required"`
	JSON      caasNamespacedNameJSON `json:"-"`
}

// caasNamespacedNameJSON contains the JSON metadata for the struct
// [CaasNamespacedName]
type caasNamespacedNameJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaasNamespacedName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caasNamespacedNameJSON) RawJSON() string {
	return r.raw
}

type CaasNamespacedNameParam struct {
	// Container name
	Name param.Field[string] `json:"name,required"`
	// Namespace name
	Namespace param.Field[string] `json:"namespace,required"`
}

func (r CaasNamespacedNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaKeyNewResponse struct {
	// List of containers using key
	Containers []CaasNamespacedName `json:"containers,required"`
	// Key starts to work
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Key description
	Description string `json:"description,required"`
	// Key expires at
	Expire time.Time `json:"expire,required,nullable" format:"date-time"`
	// Set to true if key is disabled
	IsDisabled bool `json:"is_disabled,required"`
	// Key name
	Name string `json:"name,required"`
	// Secret api key
	Secret string `json:"secret,required"`
	// Source of the apikey, can be 'cloud' or 'iate'
	Source string `json:"source,required"`
	// Key status
	Status string                       `json:"status,required"`
	JSON   cloudV1CaaKeyNewResponseJSON `json:"-"`
}

// cloudV1CaaKeyNewResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaKeyNewResponse]
type cloudV1CaaKeyNewResponseJSON struct {
	Containers  apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Expire      apijson.Field
	IsDisabled  apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaKeyListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CaasAPIKey                  `json:"results,required"`
	JSON    cloudV1CaaKeyListResponseJSON `json:"-"`
}

// cloudV1CaaKeyListResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaKeyListResponse]
type cloudV1CaaKeyListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaKeyNewParams struct {
	// List of containers using key
	Containers param.Field[[]CaasNamespacedNameParam] `json:"containers,required"`
	// Key name
	Name param.Field[string] `json:"name,required"`
	// Key description
	Description param.Field[string] `json:"description"`
	// Key expires at
	Expire param.Field[time.Time] `json:"expire" format:"date-time"`
	// Set to true if key is disabled
	IsDisabled param.Field[bool] `json:"is_disabled"`
}

func (r CloudV1CaaKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaKeyUpdateParams struct {
	// List of containers using key
	Containers param.Field[[]CaasNamespacedNameParam] `json:"containers"`
	// Key description
	Description param.Field[string] `json:"description"`
	// Key expires at
	Expire param.Field[time.Time] `json:"expire" format:"date-time"`
	// Set to true if key is disabled
	IsDisabled param.Field[bool] `json:"is_disabled"`
}

func (r CloudV1CaaKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaKeyListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1CaaKeyListParams]'s query parameters as
// `url.Values`.
func (r CloudV1CaaKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
