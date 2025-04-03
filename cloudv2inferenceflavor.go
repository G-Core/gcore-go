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

// CloudV2InferenceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceFlavorService] method instead.
type CloudV2InferenceFlavorService struct {
	Options []option.RequestOption
}

// NewCloudV2InferenceFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2InferenceFlavorService(opts ...option.RequestOption) (r *CloudV2InferenceFlavorService) {
	r = &CloudV2InferenceFlavorService{}
	r.Options = opts
	return
}

// Deprecated. Get inference instance flavor Details
func (r *CloudV2InferenceFlavorService) Get(ctx context.Context, flavorID string, opts ...option.RequestOption) (res *FlavorInference, err error) {
	opts = append(r.Options[:], opts...)
	if flavorID == "" {
		err = errors.New("missing required flavor_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/flavors/%s", flavorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deprecated. List inference instance flavors
func (r *CloudV2InferenceFlavorService) List(ctx context.Context, query CloudV2InferenceFlavorListParams, opts ...option.RequestOption) (res *FlavorInference, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/flavors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FlavorInference struct {
	// Flavor ID.
	ID string `json:"id,required" format:"uuid"`
	// Number of CPU millicores.
	CPU int64 `json:"cpu,required"`
	// Description of the flavor.
	Description string `json:"description,required"`
	// Name of the flavor.
	Name string `json:"name,required"`
	// Amount of RAM in MiB.
	Ram int64 `json:"ram,required"`
	// Number of GPUs.
	GPUCount int64 `json:"gpu_count"`
	// GPU memory in MiB.
	GPUMemory int64 `json:"gpu_memory"`
	// Model of the GPU.
	GPUModel string              `json:"gpu_model"`
	JSON     flavorInferenceJSON `json:"-"`
}

// flavorInferenceJSON contains the JSON metadata for the struct [FlavorInference]
type flavorInferenceJSON struct {
	ID          apijson.Field
	CPU         apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Ram         apijson.Field
	GPUCount    apijson.Field
	GPUMemory   apijson.Field
	GPUModel    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlavorInference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorInferenceJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceFlavorListParams struct {
	// Limit the number of returned events. Defaults to 100. Limited by max limit value
	// of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV2InferenceFlavorListParams]'s query parameters as
// `url.Values`.
func (r CloudV2InferenceFlavorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
