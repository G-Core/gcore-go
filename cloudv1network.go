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

// CloudV1NetworkService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1NetworkService] method instead.
type CloudV1NetworkService struct {
	Options      []option.RequestOption
	Metadata     *CloudV1NetworkMetadataService
	MetadataItem *CloudV1NetworkMetadataItemService
}

// NewCloudV1NetworkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1NetworkService(opts ...option.RequestOption) (r *CloudV1NetworkService) {
	r = &CloudV1NetworkService{}
	r.Options = opts
	r.Metadata = NewCloudV1NetworkMetadataService(opts...)
	r.MetadataItem = NewCloudV1NetworkMetadataItemService(opts...)
	return
}

// Create network
func (r *CloudV1NetworkService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1NetworkNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get network
func (r *CloudV1NetworkService) Get(ctx context.Context, projectID int64, regionID int64, networkID string, opts ...option.RequestOption) (res *Network, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change network name
func (r *CloudV1NetworkService) Update(ctx context.Context, projectID int64, regionID int64, networkID string, body CloudV1NetworkUpdateParams, opts ...option.RequestOption) (res *Network, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List networks
func (r *CloudV1NetworkService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1NetworkListParams, opts ...option.RequestOption) (res *CloudV1NetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete network
func (r *CloudV1NetworkService) Delete(ctx context.Context, projectID int64, regionID int64, networkID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List of instance ports by network_id
func (r *CloudV1NetworkService) ListPorts(ctx context.Context, projectID int64, regionID int64, networkID string, opts ...option.RequestOption) (res *[]CloudV1NetworkListPortsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/ports", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Network struct {
	// Network ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// True if the network `router:external` attribute
	External bool `json:"external,required"`
	// Network name
	Name string `json:"name,required"`
	// Indicates port_security_enabled status of all newly created in the network
	// ports.
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// True when the network is shared with your project by external owner
	Shared bool `json:"shared,required"`
	// List of subnetworks
	Subnets []string `json:"subnets,required" format:"uuid4"`
	// Network type (vlan, vxlan)
	Type string `json:"type,required"`
	// Datetime when the network was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// True if network has is_default attribute
	Default bool `json:"default,nullable"`
	// Network metadata
	Metadata []DetailedMetadata `json:"metadata"`
	// MTU (maximum transmission unit). Default value is 1450
	Mtu int64 `json:"mtu"`
	// Project ID
	ProjectID int64 `json:"project_id,nullable"`
	// Id of network segment
	SegmentationID int64 `json:"segmentation_id,nullable"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string      `json:"task_id,nullable" format:"uuid4"`
	JSON   networkJSON `json:"-"`
}

// networkJSON contains the JSON metadata for the struct [Network]
type networkJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	External            apijson.Field
	Name                apijson.Field
	PortSecurityEnabled apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Shared              apijson.Field
	Subnets             apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	CreatorTaskID       apijson.Field
	Default             apijson.Field
	Metadata            apijson.Field
	Mtu                 apijson.Field
	ProjectID           apijson.Field
	SegmentationID      apijson.Field
	TaskID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Network) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkJSON) RawJSON() string {
	return r.raw
}

type CloudV1NetworkListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Network                      `json:"results,required"`
	JSON    cloudV1NetworkListResponseJSON `json:"-"`
}

// cloudV1NetworkListResponseJSON contains the JSON metadata for the struct
// [CloudV1NetworkListResponse]
type cloudV1NetworkListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1NetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1NetworkListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1NetworkListPortsResponse struct {
	// Port ID
	ID string `json:"id,required" format:"uuid4"`
	// Instance ID
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Instance name
	InstanceName string                              `json:"instance_name,required"`
	JSON         cloudV1NetworkListPortsResponseJSON `json:"-"`
}

// cloudV1NetworkListPortsResponseJSON contains the JSON metadata for the struct
// [CloudV1NetworkListPortsResponse]
type cloudV1NetworkListPortsResponseJSON struct {
	ID           apijson.Field
	InstanceID   apijson.Field
	InstanceName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CloudV1NetworkListPortsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1NetworkListPortsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1NetworkNewParams struct {
	// Network name
	Name param.Field[string] `json:"name,required"`
	// Defaults to True
	CreateRouter param.Field[bool] `json:"create_router"`
	// Create one or more metadata items for a network
	Metadata param.Field[interface{}] `json:"metadata"`
	// vlan or vxlan network type is allowed. Default value is vxlan
	Type param.Field[CloudV1NetworkNewParamsType] `json:"type"`
}

func (r CloudV1NetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// vlan or vxlan network type is allowed. Default value is vxlan
type CloudV1NetworkNewParamsType string

const (
	CloudV1NetworkNewParamsTypeVlan  CloudV1NetworkNewParamsType = "vlan"
	CloudV1NetworkNewParamsTypeVxlan CloudV1NetworkNewParamsType = "vxlan"
)

func (r CloudV1NetworkNewParamsType) IsKnown() bool {
	switch r {
	case CloudV1NetworkNewParamsTypeVlan, CloudV1NetworkNewParamsTypeVxlan:
		return true
	}
	return false
}

type CloudV1NetworkUpdateParams struct {
	NamePydantic NamePydanticParam `json:"name_pydantic,required"`
}

func (r CloudV1NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NamePydantic)
}

type CloudV1NetworkListParams struct {
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata keys. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_k=["value", "sense"]" --url "http://localhost:1111/v1/networks/1/1"
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/networks/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
	// Order networks by fields and directions (name.asc). Default is `created_at.asc`.
	OrderBy param.Field[string] `query:"order_by"`
}

// URLQuery serializes [CloudV1NetworkListParams]'s query parameters as
// `url.Values`.
func (r CloudV1NetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
