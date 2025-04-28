// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
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
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateNetworkSerializer/properties/name'
	// "$.components.schemas.CreateNetworkSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateNetworkSerializer/properties/create_router'
	// "$.components.schemas.CreateNetworkSerializer.properties.create_router"
	CreateRouter param.Opt[bool] `json:"create_router,omitzero"`
	// '#/components/schemas/CreateNetworkSerializer/properties/metadata/anyOf/0'
	// "$.components.schemas.CreateNetworkSerializer.properties.metadata.anyOf[0]"
	Metadata map[string]string `json:"metadata,omitzero"`
	// '#/components/schemas/CreateNetworkSerializer/properties/type'
	// "$.components.schemas.CreateNetworkSerializer.properties.type"
	//
	// Any of "vlan", "vxlan".
	Type NetworkNewParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateNetworkSerializer/properties/type'
// "$.components.schemas.CreateNetworkSerializer.properties.type"
type NetworkNewParamsType string

const (
	NetworkNewParamsTypeVlan  NetworkNewParamsType = "vlan"
	NetworkNewParamsTypeVxlan NetworkNewParamsType = "vxlan"
)

type NetworkUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameSerializerPydantic/properties/name'
	// "$.components.schemas.NameSerializerPydantic.properties.name"
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[3]"
	MetadataK param.Opt[string] `query:"metadata_k,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[4]"
	MetadataKv param.Opt[string] `query:"metadata_kv,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[5]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}'].get.parameters[6]"
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [NetworkListParams]'s query parameters as `url.Values`.
func (r NetworkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type NetworkGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fnetworks%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bnetwork_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/networks/{project_id}/{region_id}/{network_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f NetworkGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
