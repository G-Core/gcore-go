// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/respjson"
)

// InferenceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceFlavorService] method instead.
type InferenceFlavorService struct {
	Options []option.RequestOption
}

// NewInferenceFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceFlavorService(opts ...option.RequestOption) (r InferenceFlavorService) {
	r = InferenceFlavorService{}
	r.Options = opts
	return
}

// List inference flavors
func (r *InferenceFlavorService) List(ctx context.Context, query InferenceFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceFlavor], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v3/inference/flavors"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List inference flavors
func (r *InferenceFlavorService) ListAutoPaging(ctx context.Context, query InferenceFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get inference flavor
func (r *InferenceFlavorService) Get(ctx context.Context, flavorName string, opts ...option.RequestOption) (res *InferenceFlavor, err error) {
	opts = append(r.Options[:], opts...)
	if flavorName == "" {
		err = errors.New("missing required flavor_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/flavors/%s", flavorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InferenceFlavor struct {
	// Inference flavor cpu count.
	CPU float64 `json:"cpu,required"`
	// Inference flavor description.
	Description string `json:"description,required"`
	// Inference flavor gpu count.
	GPU int64 `json:"gpu,required"`
	// Inference flavor gpu compute capability.
	GPUComputeCapability string `json:"gpu_compute_capability,required"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU                  respjson.Field
		Description          respjson.Field
		GPU                  respjson.Field
		GPUComputeCapability respjson.Field
		GPUMemory            respjson.Field
		GPUModel             respjson.Field
		IsGPUShared          respjson.Field
		Memory               respjson.Field
		Name                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InferenceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceFlavorListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceFlavorListParams]'s query parameters as
// `url.Values`.
func (r InferenceFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
