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

// CloudV3InferenceRegistryCredentialService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceRegistryCredentialService] method instead.
type CloudV3InferenceRegistryCredentialService struct {
	Options []option.RequestOption
}

// NewCloudV3InferenceRegistryCredentialService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCloudV3InferenceRegistryCredentialService(opts ...option.RequestOption) (r *CloudV3InferenceRegistryCredentialService) {
	r = &CloudV3InferenceRegistryCredentialService{}
	r.Options = opts
	return
}

// Create inference registry credential
func (r *CloudV3InferenceRegistryCredentialService) New(ctx context.Context, projectID int64, body CloudV3InferenceRegistryCredentialNewParams, opts ...option.RequestOption) (res *CloudV3InferenceRegistryCredentialNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get inference registry credential
func (r *CloudV3InferenceRegistryCredentialService) Get(ctx context.Context, projectID int64, credentialName string, opts ...option.RequestOption) (res *RegistryCredentialOut, err error) {
	opts = append(r.Options[:], opts...)
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", projectID, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update inference registry credential
func (r *CloudV3InferenceRegistryCredentialService) Update(ctx context.Context, projectID int64, credentialName string, body CloudV3InferenceRegistryCredentialUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", projectID, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// List inference registry credentials
func (r *CloudV3InferenceRegistryCredentialService) List(ctx context.Context, projectID int64, query CloudV3InferenceRegistryCredentialListParams, opts ...option.RequestOption) (res *CloudV3InferenceRegistryCredentialListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete inference registry credential
func (r *CloudV3InferenceRegistryCredentialService) Delete(ctx context.Context, projectID int64, credentialName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", projectID, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RegistryCredentialOut struct {
	// Registry credential name.
	Name string `json:"name,required"`
	// Project ID to which the inference registry credentials belongs.
	ProjectID int64 `json:"project_id,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string                    `json:"username,required"`
	JSON     registryCredentialOutJSON `json:"-"`
}

// registryCredentialOutJSON contains the JSON metadata for the struct
// [RegistryCredentialOut]
type registryCredentialOutJSON struct {
	Name        apijson.Field
	ProjectID   apijson.Field
	RegistryURL apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistryCredentialOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registryCredentialOutJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceRegistryCredentialNewResponse struct {
	// Registry credential name.
	Name string `json:"name,required"`
	// Registry password.
	Password string `json:"password,required"`
	// Project ID to which the inference registry credentials belongs.
	ProjectID int64 `json:"project_id,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string                                            `json:"username,required"`
	JSON     cloudV3InferenceRegistryCredentialNewResponseJSON `json:"-"`
}

// cloudV3InferenceRegistryCredentialNewResponseJSON contains the JSON metadata for
// the struct [CloudV3InferenceRegistryCredentialNewResponse]
type cloudV3InferenceRegistryCredentialNewResponseJSON struct {
	Name        apijson.Field
	Password    apijson.Field
	ProjectID   apijson.Field
	RegistryURL apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceRegistryCredentialNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceRegistryCredentialNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceRegistryCredentialListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []RegistryCredentialOut                            `json:"results,required"`
	JSON    cloudV3InferenceRegistryCredentialListResponseJSON `json:"-"`
}

// cloudV3InferenceRegistryCredentialListResponseJSON contains the JSON metadata
// for the struct [CloudV3InferenceRegistryCredentialListResponse]
type cloudV3InferenceRegistryCredentialListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceRegistryCredentialListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceRegistryCredentialListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceRegistryCredentialNewParams struct {
	// Registry credential name.
	Name param.Field[string] `json:"name,required"`
	// Registry password.
	Password param.Field[string] `json:"password,required"`
	// Registry URL.
	RegistryURL param.Field[string] `json:"registry_url,required"`
	// Registry username.
	Username param.Field[string] `json:"username,required"`
}

func (r CloudV3InferenceRegistryCredentialNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceRegistryCredentialUpdateParams struct {
	// Registry password.
	Password param.Field[string] `json:"password,required"`
	// Registry URL.
	RegistryURL param.Field[string] `json:"registry_url,required"`
	// Registry username.
	Username param.Field[string] `json:"username,required"`
}

func (r CloudV3InferenceRegistryCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceRegistryCredentialListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV3InferenceRegistryCredentialListParams]'s query
// parameters as `url.Values`.
func (r CloudV3InferenceRegistryCredentialListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
