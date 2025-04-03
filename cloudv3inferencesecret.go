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

// CloudV3InferenceSecretService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceSecretService] method instead.
type CloudV3InferenceSecretService struct {
	Options []option.RequestOption
}

// NewCloudV3InferenceSecretService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3InferenceSecretService(opts ...option.RequestOption) (r *CloudV3InferenceSecretService) {
	r = &CloudV3InferenceSecretService{}
	r.Options = opts
	return
}

// Create Inference Secret
func (r *CloudV3InferenceSecretService) New(ctx context.Context, projectID int64, body CloudV3InferenceSecretNewParams, opts ...option.RequestOption) (res *InferenceBoxSecrets, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Inference Secret
func (r *CloudV3InferenceSecretService) Get(ctx context.Context, projectID int64, secretName string, opts ...option.RequestOption) (res *InferenceBoxSecrets, err error) {
	opts = append(r.Options[:], opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", projectID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Inference Secret
func (r *CloudV3InferenceSecretService) Update(ctx context.Context, projectID int64, secretName string, body CloudV3InferenceSecretUpdateParams, opts ...option.RequestOption) (res *InferenceBoxSecrets, err error) {
	opts = append(r.Options[:], opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", projectID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Secrets for Inference
func (r *CloudV3InferenceSecretService) List(ctx context.Context, projectID int64, query CloudV3InferenceSecretListParams, opts ...option.RequestOption) (res *CloudV3InferenceSecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Inference Secret
func (r *CloudV3InferenceSecretService) Delete(ctx context.Context, projectID int64, secretName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", projectID, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AwsIamData struct {
	// AWS IAM key ID.
	AwsAccessKeyID string `json:"aws_access_key_id,required"`
	// AWS IAM secret key.
	AwsSecretAccessKey string         `json:"aws_secret_access_key,required"`
	JSON               awsIamDataJSON `json:"-"`
}

// awsIamDataJSON contains the JSON metadata for the struct [AwsIamData]
type awsIamDataJSON struct {
	AwsAccessKeyID     apijson.Field
	AwsSecretAccessKey apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AwsIamData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r awsIamDataJSON) RawJSON() string {
	return r.raw
}

type AwsIamDataParam struct {
	// AWS IAM key ID.
	AwsAccessKeyID param.Field[string] `json:"aws_access_key_id,required"`
	// AWS IAM secret key.
	AwsSecretAccessKey param.Field[string] `json:"aws_secret_access_key,required"`
}

func (r AwsIamDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InferenceBoxSecrets struct {
	// Secret data.
	Data AwsIamData `json:"data,required"`
	// Secret name.
	Name string `json:"name,required"`
	// Secret type.
	Type string                  `json:"type,required"`
	JSON inferenceBoxSecretsJSON `json:"-"`
}

// inferenceBoxSecretsJSON contains the JSON metadata for the struct
// [InferenceBoxSecrets]
type inferenceBoxSecretsJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferenceBoxSecrets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferenceBoxSecretsJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceSecretListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []InferenceBoxSecrets                  `json:"results,required"`
	JSON    cloudV3InferenceSecretListResponseJSON `json:"-"`
}

// cloudV3InferenceSecretListResponseJSON contains the JSON metadata for the struct
// [CloudV3InferenceSecretListResponse]
type cloudV3InferenceSecretListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceSecretListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceSecretNewParams struct {
	// Secret data.
	Data param.Field[AwsIamDataParam] `json:"data,required"`
	// Secret name.
	Name param.Field[string] `json:"name,required"`
	// Secret type. Currently only `aws-iam` is supported.
	Type param.Field[string] `json:"type,required"`
}

func (r CloudV3InferenceSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceSecretUpdateParams struct {
	// Secret data.
	Data param.Field[AwsIamDataParam] `json:"data,required"`
	// Secret type.
	Type param.Field[string] `json:"type,required"`
}

func (r CloudV3InferenceSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceSecretListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV3InferenceSecretListParams]'s query parameters as
// `url.Values`.
func (r CloudV3InferenceSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
