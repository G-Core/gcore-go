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

// CloudV2AIClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2AIClusterService] method instead.
type CloudV2AIClusterService struct {
	Options      []option.RequestOption
	Metadata     *CloudV2AIClusterMetadataService
	MetadataItem *CloudV2AIClusterMetadataItemService
}

// NewCloudV2AIClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2AIClusterService(opts ...option.RequestOption) (r *CloudV2AIClusterService) {
	r = &CloudV2AIClusterService{}
	r.Options = opts
	r.Metadata = NewCloudV2AIClusterMetadataService(opts...)
	r.MetadataItem = NewCloudV2AIClusterMetadataItemService(opts...)
	return
}

// Get AI cluster
func (r *CloudV2AIClusterService) Get(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *AICluster, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List AI clusters
func (r *CloudV2AIClusterService) List(ctx context.Context, projectID int64, regionID int64, query CloudV2AIClusterListParams, opts ...option.RequestOption) (res *CloudV2AIClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete AI cluster
func (r *CloudV2AIClusterService) Delete(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV2AIClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Powercycle (stop and start) all AI cluster nodes, aka hard reboot
func (r *CloudV2AIClusterService) Powercycle(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/powercycle", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot all AI cluster nodes
func (r *CloudV2AIClusterService) Reboot(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/reboot", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// AI cluster schema
type AICluster struct {
	// AI Cluster ID
	ClusterID string `json:"cluster_id,required"`
	// AI Cluster status
	ClusterStatus AIClusterClusterStatus `json:"cluster_status,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Task ID associated with the cluster
	TaskID string `json:"task_id,required"`
	// Cluster metadata (simple key-value pairs)
	ClusterMetadata interface{} `json:"cluster_metadata"`
	// Detailed cluster metadata
	ClusterMetadataDetailed interface{} `json:"cluster_metadata_detailed"`
	// AI Cluster Name
	ClusterName string `json:"cluster_name"`
	// Datetime when the cluster was created
	CreatedAt string `json:"created_at"`
	// Flavor ID (name)
	Flavor string `json:"flavor"`
	// Image ID
	ImageID string `json:"image_id"`
	// Image name
	ImageName string `json:"image_name"`
	// Networks managed by user and associated with the cluster
	Interfaces  []AIClusterInterface `json:"interfaces,nullable"`
	KeypairName string               `json:"keypair_name,nullable"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password string `json:"password"`
	// Poplar servers
	PoplarServers []DeprecatedInstance `json:"poplar_servers"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Security groups attached to the cluster
	SecurityGroups []AIClusterSecurityGroup `json:"security_groups,nullable"`
	// Task status
	TaskStatus AIClusterTaskStatus `json:"task_status"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData string `json:"user_data"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username string `json:"username"`
	// List of volumes attached to the cluster
	Volumes []DeprecatedVolume `json:"volumes,nullable"`
	JSON    aiClusterJSON      `json:"-"`
}

// aiClusterJSON contains the JSON metadata for the struct [AICluster]
type aiClusterJSON struct {
	ClusterID               apijson.Field
	ClusterStatus           apijson.Field
	CreatorTaskID           apijson.Field
	TaskID                  apijson.Field
	ClusterMetadata         apijson.Field
	ClusterMetadataDetailed apijson.Field
	ClusterName             apijson.Field
	CreatedAt               apijson.Field
	Flavor                  apijson.Field
	ImageID                 apijson.Field
	ImageName               apijson.Field
	Interfaces              apijson.Field
	KeypairName             apijson.Field
	Password                apijson.Field
	PoplarServers           apijson.Field
	ProjectID               apijson.Field
	Region                  apijson.Field
	RegionID                apijson.Field
	SecurityGroups          apijson.Field
	TaskStatus              apijson.Field
	UserData                apijson.Field
	Username                apijson.Field
	Volumes                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AICluster) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiClusterJSON) RawJSON() string {
	return r.raw
}

// AI Cluster status
type AIClusterClusterStatus string

const (
	AIClusterClusterStatusActive    AIClusterClusterStatus = "ACTIVE"
	AIClusterClusterStatusError     AIClusterClusterStatus = "ERROR"
	AIClusterClusterStatusPending   AIClusterClusterStatus = "PENDING"
	AIClusterClusterStatusSuspended AIClusterClusterStatus = "SUSPENDED"
)

func (r AIClusterClusterStatus) IsKnown() bool {
	switch r {
	case AIClusterClusterStatusActive, AIClusterClusterStatusError, AIClusterClusterStatusPending, AIClusterClusterStatusSuspended:
		return true
	}
	return false
}

type AIClusterInterface struct {
	// Network ID
	NetworkID string `json:"network_id"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID string `json:"port_id"`
	// Port is assigned to IP address from the subnet
	SubnetID string `json:"subnet_id"`
	// Network type
	Type string                 `json:"type"`
	JSON aiClusterInterfaceJSON `json:"-"`
}

// aiClusterInterfaceJSON contains the JSON metadata for the struct
// [AIClusterInterface]
type aiClusterInterfaceJSON struct {
	NetworkID   apijson.Field
	PortID      apijson.Field
	SubnetID    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIClusterInterface) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiClusterInterfaceJSON) RawJSON() string {
	return r.raw
}

// Security group settings for poplar servers in an AI cluster
type AIClusterSecurityGroup struct {
	// Network ID
	NetworkID string `json:"network_id"`
	// Port ID
	PortID string `json:"port_id"`
	// Security group ID list
	SecurityGroups []string                   `json:"security_groups"`
	JSON           aiClusterSecurityGroupJSON `json:"-"`
}

// aiClusterSecurityGroupJSON contains the JSON metadata for the struct
// [AIClusterSecurityGroup]
type aiClusterSecurityGroupJSON struct {
	NetworkID      apijson.Field
	PortID         apijson.Field
	SecurityGroups apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIClusterSecurityGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiClusterSecurityGroupJSON) RawJSON() string {
	return r.raw
}

// Task status
type AIClusterTaskStatus string

const (
	AIClusterTaskStatusClusterCleanUp  AIClusterTaskStatus = "CLUSTER_CLEAN_UP"
	AIClusterTaskStatusClusterResize   AIClusterTaskStatus = "CLUSTER_RESIZE"
	AIClusterTaskStatusClusterResume   AIClusterTaskStatus = "CLUSTER_RESUME"
	AIClusterTaskStatusClusterSuspend  AIClusterTaskStatus = "CLUSTER_SUSPEND"
	AIClusterTaskStatusError           AIClusterTaskStatus = "ERROR"
	AIClusterTaskStatusFinished        AIClusterTaskStatus = "FINISHED"
	AIClusterTaskStatusIpuServers      AIClusterTaskStatus = "IPU_SERVERS"
	AIClusterTaskStatusNetwork         AIClusterTaskStatus = "NETWORK"
	AIClusterTaskStatusPoplarServers   AIClusterTaskStatus = "POPLAR_SERVERS"
	AIClusterTaskStatusPostDeploySetup AIClusterTaskStatus = "POST_DEPLOY_SETUP"
	AIClusterTaskStatusVipuController  AIClusterTaskStatus = "VIPU_CONTROLLER"
)

func (r AIClusterTaskStatus) IsKnown() bool {
	switch r {
	case AIClusterTaskStatusClusterCleanUp, AIClusterTaskStatusClusterResize, AIClusterTaskStatusClusterResume, AIClusterTaskStatusClusterSuspend, AIClusterTaskStatusError, AIClusterTaskStatusFinished, AIClusterTaskStatusIpuServers, AIClusterTaskStatusNetwork, AIClusterTaskStatusPoplarServers, AIClusterTaskStatusPostDeploySetup, AIClusterTaskStatusVipuController:
		return true
	}
	return false
}

type CloudV2AIClusterListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []AICluster                      `json:"results"`
	JSON    cloudV2AIClusterListResponseJSON `json:"-"`
}

// cloudV2AIClusterListResponseJSON contains the JSON metadata for the struct
// [CloudV2AIClusterListResponse]
type cloudV2AIClusterListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2AIClusterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2AIClusterListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2AIClusterListParams struct {
	// Limit the number of returned clusters
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV2AIClusterListParams]'s query parameters as
// `url.Values`.
func (r CloudV2AIClusterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2AIClusterDeleteParams struct {
	// True if it is required to delete floating IPs assigned to the poplar(s). Can't
	// be used with floatings.
	DeleteFloatings param.Field[bool] `query:"delete_floatings"`
	// Comma separated list of floating ids that should be deleted. Can't be used with
	// delete_floatings.
	Floatings param.Field[string] `query:"floatings"`
	// Comma separated list of port IDs to be deleted with the poplar(s)
	ReservedFixedIPs param.Field[string] `query:"reserved_fixed_ips"`
}

// URLQuery serializes [CloudV2AIClusterDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV2AIClusterDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
