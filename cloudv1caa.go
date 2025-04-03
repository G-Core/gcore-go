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

// CloudV1CaaService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1CaaService] method instead.
type CloudV1CaaService struct {
	Options    []option.RequestOption
	Keys       *CloudV1CaaKeyService
	Secrets    *CloudV1CaaSecretService
	Containers *CloudV1CaaContainerService
}

// NewCloudV1CaaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1CaaService(opts ...option.RequestOption) (r *CloudV1CaaService) {
	r = &CloudV1CaaService{}
	r.Options = opts
	r.Keys = NewCloudV1CaaKeyService(opts...)
	r.Secrets = NewCloudV1CaaSecretService(opts...)
	r.Containers = NewCloudV1CaaContainerService(opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified cluster will be calculated
func (r *CloudV1CaaService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1CaaCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get CaaS flavors
func (r *CloudV1CaaService) GetFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV1CaaGetFlavorsParams, opts ...option.RequestOption) (res *CloudV1CaaGetFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/flavors/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV1CaaGetFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1CaaGetFlavorsResponseResult `json:"results,required"`
	JSON    cloudV1CaaGetFlavorsResponseJSON     `json:"-"`
}

// cloudV1CaaGetFlavorsResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaGetFlavorsResponse]
type cloudV1CaaGetFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaGetFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaGetFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaGetFlavorsResponseResult struct {
	// Number of millicpu
	CPU int64 `json:"cpu,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Megabytes of RAM
	Ram  int64                                  `json:"ram,required"`
	JSON cloudV1CaaGetFlavorsResponseResultJSON `json:"-"`
}

// cloudV1CaaGetFlavorsResponseResultJSON contains the JSON metadata for the struct
// [CloudV1CaaGetFlavorsResponseResult]
type cloudV1CaaGetFlavorsResponseResultJSON struct {
	CPU         apijson.Field
	Name        apijson.Field
	Ram         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaGetFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaGetFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaCheckQuotaParams struct {
	// Container flavor
	Flavor param.Field[string] `json:"flavor,required"`
	// Maximum number of instances
	MaxReplicas param.Field[int64] `json:"max_replicas,required"`
}

func (r CloudV1CaaCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaGetFlavorsParams struct {
	// Flavor name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [CloudV1CaaGetFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1CaaGetFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
