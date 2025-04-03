// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ResellerNameTemplateService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ResellerNameTemplateService] method instead.
type CloudV1ResellerNameTemplateService struct {
	Options []option.RequestOption
}

// NewCloudV1ResellerNameTemplateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1ResellerNameTemplateService(opts ...option.RequestOption) (r *CloudV1ResellerNameTemplateService) {
	r = &CloudV1ResellerNameTemplateService{}
	r.Options = opts
	return
}

// Update or create reseller naming configuration.
func (r *CloudV1ResellerNameTemplateService) Update(ctx context.Context, body CloudV1ResellerNameTemplateUpdateParams, opts ...option.RequestOption) (res *ResellerNameTemplates, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reseller_name_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List name template configs of the reseller.
func (r *CloudV1ResellerNameTemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *CloudV1ResellerNameTemplateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reseller_name_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the reseller's name_templates setting and drop all naming restrictions on
// the region.
func (r *CloudV1ResellerNameTemplateService) Delete(ctx context.Context, resellerID int64, regionID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/reseller_name_templates/%v/%v", resellerID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get naming restrictions that are applied to specified project and region.
func (r *CloudV1ResellerNameTemplateService) GetAvailableNames(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1ResellerNameTemplateGetAvailableNamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reseller_name_templates/%v/%v/available_names", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResellerNameTemplates struct {
	// Clients will be able to use "names" field and completely custom names.
	AllowCustomName bool `json:"allow_custom_name,required"`
	// Clients will only be able to use these strings as bare metal server
	// "name_templates".
	AllowedBmNameTemplates []string `json:"allowed_bm_name_templates,required"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameWinTemplates []string `json:"allowed_bm_name_win_templates,required"`
	// Loadbalancer name restriction.
	AllowedLbNameRestriction string `json:"allowed_lb_name_restriction,required"`
	// Clients will only be able to use these strings as load balancer
	// "name_templates".
	AllowedLbNameTemplates []string `json:"allowed_lb_name_templates,required"`
	// Clients will only be able to use these strings as instance "name_templates".
	AllowedNameTemplates []string `json:"allowed_name_templates,required"`
	// If "name_templates_limited" is True, this is the list of allowed windows
	// instance name templates.
	AllowedNameWinTemplates []string `json:"allowed_name_win_templates,required"`
	// Restrictions are active in this region.
	RegionID int64 `json:"region_id,required"`
	// Clients with this reseller_id are governed by this restriction setting.
	ResellerID int64 `json:"reseller_id,required"`
	// Name of reseller.
	ResellerName string                    `json:"reseller_name,nullable"`
	JSON         resellerNameTemplatesJSON `json:"-"`
}

// resellerNameTemplatesJSON contains the JSON metadata for the struct
// [ResellerNameTemplates]
type resellerNameTemplatesJSON struct {
	AllowCustomName           apijson.Field
	AllowedBmNameTemplates    apijson.Field
	AllowedBmNameWinTemplates apijson.Field
	AllowedLbNameRestriction  apijson.Field
	AllowedLbNameTemplates    apijson.Field
	AllowedNameTemplates      apijson.Field
	AllowedNameWinTemplates   apijson.Field
	RegionID                  apijson.Field
	ResellerID                apijson.Field
	ResellerName              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ResellerNameTemplates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resellerNameTemplatesJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerNameTemplateListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ResellerNameTemplates                     `json:"results,required"`
	JSON    cloudV1ResellerNameTemplateListResponseJSON `json:"-"`
}

// cloudV1ResellerNameTemplateListResponseJSON contains the JSON metadata for the
// struct [CloudV1ResellerNameTemplateListResponse]
type cloudV1ResellerNameTemplateListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ResellerNameTemplateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ResellerNameTemplateListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerNameTemplateGetAvailableNamesResponse struct {
	// If true, instances can be created using "names" field.
	CustomNameAllowed bool `json:"custom_name_allowed,required"`
	// If true, only specific strings are allowed in "name_templates" fields.
	NameTemplatesLimited bool `json:"name_templates_limited,required"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameTemplates []string `json:"allowed_bm_name_templates,nullable"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameWinTemplates []string `json:"allowed_bm_name_win_templates,nullable"`
	// Loadbalancer name restriction.
	AllowedLbNameRestriction string `json:"allowed_lb_name_restriction,nullable"`
	// If "name_templates_limited" is True, this is the list of allowed name templates.
	AllowedLbNameTemplates []string `json:"allowed_lb_name_templates,nullable"`
	// If "name_templates_limited" is True, this is the list of allowed instance name
	// templates.
	AllowedNameTemplates []string `json:"allowed_name_templates,nullable"`
	// If "name_templates_limited" is True, this is the list of allowed windows
	// instance name templates.
	AllowedNameWinTemplates []string                                                 `json:"allowed_name_win_templates,nullable"`
	JSON                    cloudV1ResellerNameTemplateGetAvailableNamesResponseJSON `json:"-"`
}

// cloudV1ResellerNameTemplateGetAvailableNamesResponseJSON contains the JSON
// metadata for the struct [CloudV1ResellerNameTemplateGetAvailableNamesResponse]
type cloudV1ResellerNameTemplateGetAvailableNamesResponseJSON struct {
	CustomNameAllowed         apijson.Field
	NameTemplatesLimited      apijson.Field
	AllowedBmNameTemplates    apijson.Field
	AllowedBmNameWinTemplates apijson.Field
	AllowedLbNameRestriction  apijson.Field
	AllowedLbNameTemplates    apijson.Field
	AllowedNameTemplates      apijson.Field
	AllowedNameWinTemplates   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CloudV1ResellerNameTemplateGetAvailableNamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ResellerNameTemplateGetAvailableNamesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerNameTemplateUpdateParams struct {
	// Restrictions are active in this region.
	RegionID param.Field[int64] `json:"region_id,required"`
	// Clients with this reseller_id are governed by this restriction setting.
	ResellerID param.Field[int64] `json:"reseller_id,required"`
	// Clients will be able to use "names" field and completely custom names.
	AllowCustomName param.Field[bool] `json:"allow_custom_name"`
	// Clients will only be able to use these strings as bare metal server
	// "name_templates".
	AllowedBmNameTemplates param.Field[[]string] `json:"allowed_bm_name_templates"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameWinTemplates param.Field[[]string] `json:"allowed_bm_name_win_templates"`
	// Loadbalancer name restriction.
	AllowedLbNameRestriction param.Field[string] `json:"allowed_lb_name_restriction"`
	// Clients will only be able to use these strings as load balancer
	// "name_templates".
	AllowedLbNameTemplates param.Field[[]string] `json:"allowed_lb_name_templates"`
	// Clients will only be able to use these strings as instance "name_templates".
	AllowedNameTemplates param.Field[[]string] `json:"allowed_name_templates"`
	// If "name_templates_limited" is True, this is the list of allowed windows
	// instance name templates.
	AllowedNameWinTemplates param.Field[[]string] `json:"allowed_name_win_templates"`
}

func (r CloudV1ResellerNameTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
