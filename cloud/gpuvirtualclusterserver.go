// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// GPUVirtualClusterServerService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualClusterServerService] method instead.
type GPUVirtualClusterServerService struct {
	Options []option.RequestOption
}

// NewGPUVirtualClusterServerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUVirtualClusterServerService(opts ...option.RequestOption) (r GPUVirtualClusterServerService) {
	r = GPUVirtualClusterServerService{}
	r.Options = opts
	return
}

// List all servers in a virtual GPU cluster.
func (r *GPUVirtualClusterServerService) List(ctx context.Context, clusterID string, params GPUVirtualClusterServerListParams, opts ...option.RequestOption) (res *GPUVirtualClusterServerList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s/servers", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a server from a virtual GPU cluster and its associated resources.
func (r *GPUVirtualClusterServerService) Delete(ctx context.Context, serverID string, params GPUVirtualClusterServerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s/servers/%s", params.ProjectID.Value, params.RegionID.Value, params.ClusterID, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

type GPUVirtualClusterServer struct {
	// Server unique identifier
	ID string `json:"id" api:"required" format:"uuid4"`
	// Server creation date and time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique flavor identifier
	Flavor string `json:"flavor" api:"required"`
	// Server's image UUID
	ImageID string `json:"image_id" api:"required" format:"uuid4"`
	// List of IP addresses
	IPAddresses []string `json:"ip_addresses" api:"required"`
	// Server's name generated using cluster's name
	Name string `json:"name" api:"required"`
	// Security groups
	SecurityGroups []GPUVirtualClusterServerSecurityGroup `json:"security_groups" api:"required"`
	// SSH key pair assigned to the server
	SSHKeyName string `json:"ssh_key_name" api:"required"`
	// Current server status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status GPUVirtualClusterServerStatus `json:"status" api:"required"`
	// User defined tags
	Tags []Tag `json:"tags" api:"required"`
	// Identifier of the task currently modifying the GPU cluster
	TaskID string `json:"task_id" api:"required" format:"uuid4"`
	// Server update date and time
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Flavor         respjson.Field
		ImageID        respjson.Field
		IPAddresses    respjson.Field
		Name           respjson.Field
		SecurityGroups respjson.Field
		SSHKeyName     respjson.Field
		Status         respjson.Field
		Tags           respjson.Field
		TaskID         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServer) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServerSecurityGroup struct {
	// Security group ID
	ID string `json:"id" api:"required"`
	// Security group name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServerSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServerSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current server status
type GPUVirtualClusterServerStatus string

const (
	GPUVirtualClusterServerStatusActive           GPUVirtualClusterServerStatus = "ACTIVE"
	GPUVirtualClusterServerStatusBuild            GPUVirtualClusterServerStatus = "BUILD"
	GPUVirtualClusterServerStatusDeleted          GPUVirtualClusterServerStatus = "DELETED"
	GPUVirtualClusterServerStatusError            GPUVirtualClusterServerStatus = "ERROR"
	GPUVirtualClusterServerStatusHardReboot       GPUVirtualClusterServerStatus = "HARD_REBOOT"
	GPUVirtualClusterServerStatusMigrating        GPUVirtualClusterServerStatus = "MIGRATING"
	GPUVirtualClusterServerStatusPassword         GPUVirtualClusterServerStatus = "PASSWORD"
	GPUVirtualClusterServerStatusPaused           GPUVirtualClusterServerStatus = "PAUSED"
	GPUVirtualClusterServerStatusReboot           GPUVirtualClusterServerStatus = "REBOOT"
	GPUVirtualClusterServerStatusRebuild          GPUVirtualClusterServerStatus = "REBUILD"
	GPUVirtualClusterServerStatusRescue           GPUVirtualClusterServerStatus = "RESCUE"
	GPUVirtualClusterServerStatusResize           GPUVirtualClusterServerStatus = "RESIZE"
	GPUVirtualClusterServerStatusRevertResize     GPUVirtualClusterServerStatus = "REVERT_RESIZE"
	GPUVirtualClusterServerStatusShelved          GPUVirtualClusterServerStatus = "SHELVED"
	GPUVirtualClusterServerStatusShelvedOffloaded GPUVirtualClusterServerStatus = "SHELVED_OFFLOADED"
	GPUVirtualClusterServerStatusShutoff          GPUVirtualClusterServerStatus = "SHUTOFF"
	GPUVirtualClusterServerStatusSoftDeleted      GPUVirtualClusterServerStatus = "SOFT_DELETED"
	GPUVirtualClusterServerStatusSuspended        GPUVirtualClusterServerStatus = "SUSPENDED"
	GPUVirtualClusterServerStatusUnknown          GPUVirtualClusterServerStatus = "UNKNOWN"
	GPUVirtualClusterServerStatusVerifyResize     GPUVirtualClusterServerStatus = "VERIFY_RESIZE"
)

type GPUVirtualClusterServerList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []GPUVirtualClusterServer `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServerList) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServerListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Filters the results to include only servers whose last change timestamp is less
	// than the specified datetime. Format: ISO 8601.
	ChangedBefore param.Opt[time.Time] `query:"changed_before,omitzero" format:"date-time" json:"-"`
	// Filters the results to include only servers whose last change timestamp is
	// greater than or equal to the specified datetime. Format: ISO 8601.
	ChangedSince param.Opt[time.Time] `query:"changed_since,omitzero" format:"date-time" json:"-"`
	// Filter servers by ip address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter servers by name. You can provide a full or partial name, servers with
	// matching names will be returned. For example, entering 'test' will return all
	// servers that contain 'test' in their name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order field
	//
	// Any of "created_at.asc", "created_at.desc", "status.asc", "status.desc".
	OrderBy GPUVirtualClusterServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filters servers by status.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "MIGRATING", "PAUSED",
	// "REBOOT", "REBUILD", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED",
	// "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE".
	Status GPUVirtualClusterServerListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter servers by uuid.
	Uuids []string `query:"uuids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUVirtualClusterServerListParams]'s query parameters as
// `url.Values`.
func (r GPUVirtualClusterServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order field
type GPUVirtualClusterServerListParamsOrderBy string

const (
	GPUVirtualClusterServerListParamsOrderByCreatedAtAsc  GPUVirtualClusterServerListParamsOrderBy = "created_at.asc"
	GPUVirtualClusterServerListParamsOrderByCreatedAtDesc GPUVirtualClusterServerListParamsOrderBy = "created_at.desc"
	GPUVirtualClusterServerListParamsOrderByStatusAsc     GPUVirtualClusterServerListParamsOrderBy = "status.asc"
	GPUVirtualClusterServerListParamsOrderByStatusDesc    GPUVirtualClusterServerListParamsOrderBy = "status.desc"
)

// Filters servers by status.
type GPUVirtualClusterServerListParamsStatus string

const (
	GPUVirtualClusterServerListParamsStatusActive           GPUVirtualClusterServerListParamsStatus = "ACTIVE"
	GPUVirtualClusterServerListParamsStatusBuild            GPUVirtualClusterServerListParamsStatus = "BUILD"
	GPUVirtualClusterServerListParamsStatusError            GPUVirtualClusterServerListParamsStatus = "ERROR"
	GPUVirtualClusterServerListParamsStatusHardReboot       GPUVirtualClusterServerListParamsStatus = "HARD_REBOOT"
	GPUVirtualClusterServerListParamsStatusMigrating        GPUVirtualClusterServerListParamsStatus = "MIGRATING"
	GPUVirtualClusterServerListParamsStatusPaused           GPUVirtualClusterServerListParamsStatus = "PAUSED"
	GPUVirtualClusterServerListParamsStatusReboot           GPUVirtualClusterServerListParamsStatus = "REBOOT"
	GPUVirtualClusterServerListParamsStatusRebuild          GPUVirtualClusterServerListParamsStatus = "REBUILD"
	GPUVirtualClusterServerListParamsStatusResize           GPUVirtualClusterServerListParamsStatus = "RESIZE"
	GPUVirtualClusterServerListParamsStatusRevertResize     GPUVirtualClusterServerListParamsStatus = "REVERT_RESIZE"
	GPUVirtualClusterServerListParamsStatusShelved          GPUVirtualClusterServerListParamsStatus = "SHELVED"
	GPUVirtualClusterServerListParamsStatusShelvedOffloaded GPUVirtualClusterServerListParamsStatus = "SHELVED_OFFLOADED"
	GPUVirtualClusterServerListParamsStatusShutoff          GPUVirtualClusterServerListParamsStatus = "SHUTOFF"
	GPUVirtualClusterServerListParamsStatusSoftDeleted      GPUVirtualClusterServerListParamsStatus = "SOFT_DELETED"
	GPUVirtualClusterServerListParamsStatusSuspended        GPUVirtualClusterServerListParamsStatus = "SUSPENDED"
	GPUVirtualClusterServerListParamsStatusVerifyResize     GPUVirtualClusterServerListParamsStatus = "VERIFY_RESIZE"
)

type GPUVirtualClusterServerDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Cluster unique identifier
	ClusterID string `path:"cluster_id" api:"required" format:"uuid4" json:"-"`
	// Flag indicating whether the floating ips associated with server / cluster are
	// deleted
	AllFloatingIPs param.Opt[bool] `query:"all_floating_ips,omitzero" json:"-"`
	// Flag indicating whether the reserved fixed ips associated with server / cluster
	// are deleted
	AllReservedFixedIPs param.Opt[bool] `query:"all_reserved_fixed_ips,omitzero" json:"-"`
	// Flag indicating whether all attached volumes are deleted
	AllVolumes param.Opt[bool] `query:"all_volumes,omitzero" json:"-"`
	// Optional list of floating ips to be deleted
	FloatingIPIDs []string `query:"floating_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Optional list of reserved fixed ips to be deleted
	ReservedFixedIPIDs []string `query:"reserved_fixed_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Optional list of volumes to be deleted
	VolumeIDs []string `query:"volume_ids,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [GPUVirtualClusterServerDeleteParams]'s query parameters as
// `url.Values`.
func (r GPUVirtualClusterServerDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
