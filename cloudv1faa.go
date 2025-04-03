// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1FaaService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FaaService] method instead.
type CloudV1FaaService struct {
	Options    []option.RequestOption
	Keys       *CloudV1FaaKeyService
	Namespaces *CloudV1FaaNamespaceService
}

// NewCloudV1FaaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1FaaService(opts ...option.RequestOption) (r *CloudV1FaaService) {
	r = &CloudV1FaaService{}
	r.Options = opts
	r.Keys = NewCloudV1FaaKeyService(opts...)
	r.Namespaces = NewCloudV1FaaNamespaceService(opts...)
	return
}

// Get FaaS flavors
func (r *CloudV1FaaService) GetFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV1FaaGetFlavorsParams, opts ...option.RequestOption) (res *CloudV1FaaGetFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/flavors/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get FaaS runtimes
func (r *CloudV1FaaService) GetRuntimes(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1FaaGetRuntimesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/faas/runtimes/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CloudV1FaaGetFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1FaaGetFlavorsResponseResult `json:"results,required"`
	JSON    cloudV1FaaGetFlavorsResponseJSON     `json:"-"`
}

// cloudV1FaaGetFlavorsResponseJSON contains the JSON metadata for the struct
// [CloudV1FaaGetFlavorsResponse]
type cloudV1FaaGetFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaGetFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaGetFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaGetFlavorsResponseResult struct {
	// The number of millicpus allocated to the flavor.
	CPU int64 `json:"cpu,required"`
	// The amount of RAM (in megabytes) allocated to the flavor.
	Ram int64 `json:"ram,required"`
	// The generated name of the flavor.
	Name string                                 `json:"name,nullable"`
	JSON cloudV1FaaGetFlavorsResponseResultJSON `json:"-"`
}

// cloudV1FaaGetFlavorsResponseResultJSON contains the JSON metadata for the struct
// [CloudV1FaaGetFlavorsResponseResult]
type cloudV1FaaGetFlavorsResponseResultJSON struct {
	CPU         apijson.Field
	Ram         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaGetFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaGetFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaGetRuntimesResponse struct {
	// FaaS runtimes.
	Runtimes map[string][]string               `json:"runtimes,required"`
	JSON     cloudV1FaaGetRuntimesResponseJSON `json:"-"`
}

// cloudV1FaaGetRuntimesResponseJSON contains the JSON metadata for the struct
// [CloudV1FaaGetRuntimesResponse]
type cloudV1FaaGetRuntimesResponseJSON struct {
	Runtimes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaGetRuntimesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaGetRuntimesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaGetFlavorsParams struct {
	// Function name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [CloudV1FaaGetFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1FaaGetFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
