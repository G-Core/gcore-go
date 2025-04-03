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

// CloudV2InferenceAPIKeyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceAPIKeyService] method instead.
type CloudV2InferenceAPIKeyService struct {
	Options []option.RequestOption
}

// NewCloudV2InferenceAPIKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2InferenceAPIKeyService(opts ...option.RequestOption) (r *CloudV2InferenceAPIKeyService) {
	r = &CloudV2InferenceAPIKeyService{}
	r.Options = opts
	return
}

// Deprecated. Create API Key
func (r *CloudV2InferenceAPIKeyService) New(ctx context.Context, body CloudV2InferenceAPIKeyNewParams, opts ...option.RequestOption) (res *CloudV2InferenceAPIKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deprecated. Get API Key
func (r *CloudV2InferenceAPIKeyService) Get(ctx context.Context, keyID string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/api_keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deprecated. Update API Key
func (r *CloudV2InferenceAPIKeyService) Update(ctx context.Context, keyID string, body CloudV2InferenceAPIKeyUpdateParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/api_keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deprecated. List API Keys
func (r *CloudV2InferenceAPIKeyService) List(ctx context.Context, query CloudV2InferenceAPIKeyListParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deprecated. Delete API Key
func (r *CloudV2InferenceAPIKeyService) Delete(ctx context.Context, keyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/api_keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type APIKey struct {
	// API key creation date in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// API key ID.
	KeyID string `json:"key_id,required" format:"uuid"`
	// API key name.
	Name string `json:"name,required"`
	// API key status.
	Status string `json:"status,required"`
	// API key description.
	Description string `json:"description,nullable"`
	// Expiration date for the API key in ISO 8601 format.
	ExpiresAt string `json:"expires_at,nullable"`
	// Inference instance IDs to bind to the API key.
	InferenceInstanceIDs []string `json:"inference_instance_ids"`
	// Project ID.
	ProjectID int64      `json:"project_id,nullable"`
	JSON      apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	CreatedAt            apijson.Field
	KeyID                apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	Description          apijson.Field
	ExpiresAt            apijson.Field
	InferenceInstanceIDs apijson.Field
	ProjectID            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceAPIKeyNewResponse struct {
	// API key creation date in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// API key ID.
	KeyID string `json:"key_id,required" format:"uuid"`
	// API key name.
	Name string `json:"name,required"`
	// API key secret.
	Secret string `json:"secret,required"`
	// API key status.
	Status string `json:"status,required"`
	// API key description.
	Description string `json:"description,nullable"`
	// Expiration date for the API key in ISO 8601 format.
	ExpiresAt string `json:"expires_at,nullable"`
	// Inference instance IDs to bind to the API key.
	InferenceInstanceIDs []string `json:"inference_instance_ids"`
	// Project ID.
	ProjectID int64                                 `json:"project_id,nullable"`
	JSON      cloudV2InferenceAPIKeyNewResponseJSON `json:"-"`
}

// cloudV2InferenceAPIKeyNewResponseJSON contains the JSON metadata for the struct
// [CloudV2InferenceAPIKeyNewResponse]
type cloudV2InferenceAPIKeyNewResponseJSON struct {
	CreatedAt            apijson.Field
	KeyID                apijson.Field
	Name                 apijson.Field
	Secret               apijson.Field
	Status               apijson.Field
	Description          apijson.Field
	ExpiresAt            apijson.Field
	InferenceInstanceIDs apijson.Field
	ProjectID            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CloudV2InferenceAPIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceAPIKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceAPIKeyNewParams struct {
	// API key name.
	Name param.Field[string] `json:"name,required"`
	// API key description.
	Description param.Field[string] `json:"description"`
	// Expiration date for the API key in ISO 8601 format.
	ExpiresAt param.Field[string] `json:"expires_at"`
	// Inference instance IDs to bind to the API key.
	InferenceInstanceIDs param.Field[[]string] `json:"inference_instance_ids"`
	// Project ID.
	ProjectID param.Field[int64] `json:"project_id"`
}

func (r CloudV2InferenceAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2InferenceAPIKeyUpdateParams struct {
	// API key description.
	Description param.Field[string] `json:"description"`
	// Expiration date for the API key in ISO 8601 format.
	ExpiresAt param.Field[string] `json:"expires_at"`
	// Inference instance IDs to bind to the API key.
	InferenceInstanceIDs param.Field[[]string] `json:"inference_instance_ids"`
	// Project ID.
	ProjectID param.Field[int64] `json:"project_id"`
}

func (r CloudV2InferenceAPIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2InferenceAPIKeyListParams struct {
	// Limit the number of returned instances. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Order instances by transmitted fields and directions (name.asc)
	OrderBy param.Field[string] `query:"order_by"`
	// Project ID. If not provided, will use the default project
	ProjectID param.Field[int64] `query:"project_id"`
}

// URLQuery serializes [CloudV2InferenceAPIKeyListParams]'s query parameters as
// `url.Values`.
func (r CloudV2InferenceAPIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
