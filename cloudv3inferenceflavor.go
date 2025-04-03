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

// CloudV3InferenceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceFlavorService] method instead.
type CloudV3InferenceFlavorService struct {
	Options []option.RequestOption
}

// NewCloudV3InferenceFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3InferenceFlavorService(opts ...option.RequestOption) (r *CloudV3InferenceFlavorService) {
	r = &CloudV3InferenceFlavorService{}
	r.Options = opts
	return
}

// Get inference flavor
func (r *CloudV3InferenceFlavorService) Get(ctx context.Context, flavorName string, opts ...option.RequestOption) (res *InferenceFlavorOut, err error) {
	opts = append(r.Options[:], opts...)
	if flavorName == "" {
		err = errors.New("missing required flavor_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/flavors/%s", flavorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List inference flavors
func (r *CloudV3InferenceFlavorService) List(ctx context.Context, query CloudV3InferenceFlavorListParams, opts ...option.RequestOption) (res *CloudV3InferenceFlavorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v3/inference/flavors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InferenceFlavorOut struct {
	// Inference flavor cpu count.
	CPU float64 `json:"cpu,required"`
	// Inference flavor description.
	Description string `json:"description,required"`
	// Inference flavor gpu count.
	GPU int64 `json:"gpu,required"`
	// Inference flavor gpu memory in Gi.
	GPUMemory float64 `json:"gpu_memory,required"`
	// Inference flavor gpu model.
	GPUModel string `json:"gpu_model,required"`
	// Inference flavor is gpu shared.
	IsGPUShared bool `json:"is_gpu_shared,required"`
	// Inference flavor memory in Gi.
	Memory float64 `json:"memory,required"`
	// Inference flavor name.
	Name string `json:"name,required"`
	// Inference flavor gpu compute capability.
	GPUComputeCapability string                 `json:"gpu_compute_capability"`
	JSON                 inferenceFlavorOutJSON `json:"-"`
}

// inferenceFlavorOutJSON contains the JSON metadata for the struct
// [InferenceFlavorOut]
type inferenceFlavorOutJSON struct {
	CPU                  apijson.Field
	Description          apijson.Field
	GPU                  apijson.Field
	GPUMemory            apijson.Field
	GPUModel             apijson.Field
	IsGPUShared          apijson.Field
	Memory               apijson.Field
	Name                 apijson.Field
	GPUComputeCapability apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InferenceFlavorOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferenceFlavorOutJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceFlavorListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []InferenceFlavorOut                   `json:"results,required"`
	JSON    cloudV3InferenceFlavorListResponseJSON `json:"-"`
}

// cloudV3InferenceFlavorListResponseJSON contains the JSON metadata for the struct
// [CloudV3InferenceFlavorListResponse]
type cloudV3InferenceFlavorListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceFlavorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceFlavorListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceFlavorListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV3InferenceFlavorListParams]'s query parameters as
// `url.Values`.
func (r CloudV3InferenceFlavorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
