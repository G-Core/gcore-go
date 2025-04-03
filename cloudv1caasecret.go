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

// CloudV1CaaSecretService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1CaaSecretService] method instead.
type CloudV1CaaSecretService struct {
	Options []option.RequestOption
}

// NewCloudV1CaaSecretService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1CaaSecretService(opts ...option.RequestOption) (r *CloudV1CaaSecretService) {
	r = &CloudV1CaaSecretService{}
	r.Options = opts
	return
}

// Create pull secret
func (r *CloudV1CaaSecretService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1CaaSecretNewParams, opts ...option.RequestOption) (res *PullSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/secrets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get pull secret
func (r *CloudV1CaaSecretService) Get(ctx context.Context, projectID int64, regionID int64, secretName string, opts ...option.RequestOption) (res *PullSecret, err error) {
	opts = append(r.Options[:], opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/secrets/%v/%v/%s", projectID, regionID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change pull secret
func (r *CloudV1CaaSecretService) Update(ctx context.Context, projectID int64, regionID int64, secretName string, body CloudV1CaaSecretUpdateParams, opts ...option.RequestOption) (res *PullSecret, err error) {
	opts = append(r.Options[:], opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/secrets/%v/%v/%s", projectID, regionID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List pull secrets
func (r *CloudV1CaaSecretService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1CaaSecretListParams, opts ...option.RequestOption) (res *CloudV1CaaSecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/secrets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete pull secret
func (r *CloudV1CaaSecretService) Delete(ctx context.Context, projectID int64, regionID int64, secretName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/secrets/%v/%v/%s", projectID, regionID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PullSecret struct {
	// Pull secret login
	Login string `json:"login,required"`
	// Pull secret name
	Name string `json:"name,required"`
	// Pull secret registry
	Registry string `json:"registry,required"`
	// Pull secret updated at
	UpdatedAt string         `json:"updated_at,required"`
	JSON      pullSecretJSON `json:"-"`
}

// pullSecretJSON contains the JSON metadata for the struct [PullSecret]
type pullSecretJSON struct {
	Login       apijson.Field
	Name        apijson.Field
	Registry    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PullSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pullSecretJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaSecretListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []PullSecret                     `json:"results,required"`
	JSON    cloudV1CaaSecretListResponseJSON `json:"-"`
}

// cloudV1CaaSecretListResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaSecretListResponse]
type cloudV1CaaSecretListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaSecretListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaSecretNewParams struct {
	// Pull secret login
	Login param.Field[string] `json:"login,required"`
	// Pull secret name
	Name param.Field[string] `json:"name,required"`
	// Pull secret password
	Password param.Field[string] `json:"password,required" format:"password"`
	// Pull secret registry
	Registry param.Field[string] `json:"registry,required"`
}

func (r CloudV1CaaSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaSecretUpdateParams struct {
	// Pull secret login
	Login param.Field[string] `json:"login"`
	// Pull secret password
	Password param.Field[string] `json:"password" format:"password"`
	// Pull secret registry
	Registry param.Field[string] `json:"registry"`
}

func (r CloudV1CaaSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaSecretListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1CaaSecretListParams]'s query parameters as
// `url.Values`.
func (r CloudV1CaaSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
