// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// NetworkService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options []option.RequestOption
	Subnets NetworkSubnetService
	Routers NetworkRouterService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	r.Subnets = NewNetworkSubnetService(opts...)
	r.Routers = NewNetworkRouterService(opts...)
	return
}

// Create network
func (r *NetworkService) New(ctx context.Context, params NetworkNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change network name
func (r *NetworkService) Update(ctx context.Context, networkID string, params NetworkUpdateParams, opts ...option.RequestOption) (res *Network, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List networks
func (r *NetworkService) List(ctx context.Context, params NetworkListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Network], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List networks
func (r *NetworkService) ListAutoPaging(ctx context.Context, params NetworkListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Network] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete network
func (r *NetworkService) Delete(ctx context.Context, networkID string, body NetworkDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get network
func (r *NetworkService) Get(ctx context.Context, networkID string, query NetworkGetParams, opts ...option.RequestOption) (res *Network, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NetworkNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Network name
	Name string `json:"name,required"`
	// Defaults to True
	CreateRouter param.Opt[bool] `json:"create_router,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// vlan or vxlan network type is allowed. Default value is vxlan
	//
	// Any of "vlan", "vxlan".
	Type NetworkNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r NetworkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// vlan or vxlan network type is allowed. Default value is vxlan
type NetworkNewParamsType string

const (
	NetworkNewParamsTypeVlan  NetworkNewParamsType = "vlan"
	NetworkNewParamsTypeVxlan NetworkNewParamsType = "vxlan"
)

type NetworkUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name.
	Name string `json:"name,required"`
	paramObj
}

func (r NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit the number of returned limit request entities.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order networks by fields and directions (name.asc). Default is `created_at.asc`.
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "tag_key_value={"key": "value"}" --url
	// "http://localhost:1111/v1/networks/1/1"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by tag keys.
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListParams]'s query parameters as `url.Values`.
func (r NetworkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type NetworkGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
