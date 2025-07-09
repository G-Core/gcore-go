// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// GPUBaremetalClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterService] method instead.
type GPUBaremetalClusterService struct {
	Options    []option.RequestOption
	Interfaces GPUBaremetalClusterInterfaceService
	Servers    GPUBaremetalClusterServerService
	Flavors    GPUBaremetalClusterFlavorService
	Images     GPUBaremetalClusterImageService
	tasks      TaskService
}

// NewGPUBaremetalClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterService(opts ...option.RequestOption) (r GPUBaremetalClusterService) {
	r = GPUBaremetalClusterService{}
	r.Options = opts
	r.Interfaces = NewGPUBaremetalClusterInterfaceService(opts...)
	r.Servers = NewGPUBaremetalClusterServerService(opts...)
	r.Flavors = NewGPUBaremetalClusterFlavorService(opts...)
	r.Images = NewGPUBaremetalClusterImageService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new GPU cluster with specified configuration. The cluster can be
// created with one or more nodes.
func (r *GPUBaremetalClusterService) New(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List bare metal GPU clusters
func (r *GPUBaremetalClusterService) List(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GPUBaremetalCluster], err error) {
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
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List bare metal GPU clusters
func (r *GPUBaremetalClusterService) ListAutoPaging(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GPUBaremetalCluster] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete bare metal GPU cluster
func (r *GPUBaremetalClusterService) Delete(ctx context.Context, clusterID string, params GPUBaremetalClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Get bare metal GPU cluster
func (r *GPUBaremetalClusterService) Get(ctx context.Context, clusterID string, query GPUBaremetalClusterGetParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stops and then starts all cluster servers, effectively performing a hard reboot.
func (r *GPUBaremetalClusterService) PowercycleAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterPowercycleAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/powercycle", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot all bare metal GPU cluster servers
func (r *GPUBaremetalClusterService) RebootAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterRebootAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/reboot", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Rebuild one or more nodes in a GPU cluster. All cluster nodes must be specified
// to update the cluster image.
func (r *GPUBaremetalClusterService) Rebuild(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/rebuild", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the number of nodes in a GPU cluster. The cluster can be scaled up or
// down.
func (r *GPUBaremetalClusterService) Resize(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/resize", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new GPU bare metal cluster and polls for completion
func (r *GPUBaremetalClusterService) NewAndPoll(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.AIClusters) != 1 {
		return nil, errors.New("expected exactly one GPU bare metal cluster to be created in a task")
	}
	resourceID := task.CreatedResources.AIClusters[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// RebuildAndPoll rebuilds a GPU bare metal cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) RebuildAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.Rebuild(ctx, clusterID, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

// ResizeAndPoll resizes a GPU bare metal cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) ResizeAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.Resize(ctx, clusterID, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

type GPUBaremetalCluster struct {
	// GPU Cluster ID
	ClusterID string `json:"cluster_id,required" format:"uuid4"`
	// GPU Cluster Name
	ClusterName string `json:"cluster_name,required"`
	// GPU Cluster status
	//
	// Any of "ACTIVE", "ERROR", "PENDING", "SUSPENDED".
	ClusterStatus GPUBaremetalClusterClusterStatus `json:"cluster_status,required"`
	// Datetime when the cluster was created
	CreatedAt string `json:"created_at,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// Flavor ID is the same as the name
	Flavor string `json:"flavor,required"`
	// Image ID
	ImageID string `json:"image_id,required" format:"uuid4"`
	// Image name
	ImageName string `json:"image_name,required"`
	// Networks managed by user and associated with the cluster
	Interfaces []GPUBaremetalClusterInterface `json:"interfaces,required"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password string `json:"password,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// GPU cluster servers
	Servers []GPUBaremetalClusterServer `json:"servers,required"`
	// Keypair name to inject into new cluster(s)
	SSHKeyName string `json:"ssh_key_name,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Task ID associated with the cluster
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Task status
	//
	// Any of "CLUSTER_CLEAN_UP", "CLUSTER_RESIZE", "CLUSTER_RESUME",
	// "CLUSTER_SUSPEND", "ERROR", "FINISHED", "IPU_SERVERS", "NETWORK",
	// "POPLAR_SERVERS", "POST_DEPLOY_SETUP", "VIPU_CONTROLLER".
	TaskStatus GPUBaremetalClusterTaskStatus `json:"task_status,required"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData string `json:"user_data,required"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID     respjson.Field
		ClusterName   respjson.Field
		ClusterStatus respjson.Field
		CreatedAt     respjson.Field
		CreatorTaskID respjson.Field
		Flavor        respjson.Field
		ImageID       respjson.Field
		ImageName     respjson.Field
		Interfaces    respjson.Field
		Password      respjson.Field
		ProjectID     respjson.Field
		Region        respjson.Field
		RegionID      respjson.Field
		Servers       respjson.Field
		SSHKeyName    respjson.Field
		Tags          respjson.Field
		TaskID        respjson.Field
		TaskStatus    respjson.Field
		UserData      respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalCluster) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU Cluster status
type GPUBaremetalClusterClusterStatus string

const (
	GPUBaremetalClusterClusterStatusActive    GPUBaremetalClusterClusterStatus = "ACTIVE"
	GPUBaremetalClusterClusterStatusError     GPUBaremetalClusterClusterStatus = "ERROR"
	GPUBaremetalClusterClusterStatusPending   GPUBaremetalClusterClusterStatus = "PENDING"
	GPUBaremetalClusterClusterStatusSuspended GPUBaremetalClusterClusterStatus = "SUSPENDED"
)

type GPUBaremetalClusterInterface struct {
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	PortID string `json:"port_id,required" format:"uuid4"`
	// Port is assigned to IP address from the subnet
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Network type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkID   respjson.Field
		PortID      respjson.Field
		SubnetID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterInterface) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task status
type GPUBaremetalClusterTaskStatus string

const (
	GPUBaremetalClusterTaskStatusClusterCleanUp  GPUBaremetalClusterTaskStatus = "CLUSTER_CLEAN_UP"
	GPUBaremetalClusterTaskStatusClusterResize   GPUBaremetalClusterTaskStatus = "CLUSTER_RESIZE"
	GPUBaremetalClusterTaskStatusClusterResume   GPUBaremetalClusterTaskStatus = "CLUSTER_RESUME"
	GPUBaremetalClusterTaskStatusClusterSuspend  GPUBaremetalClusterTaskStatus = "CLUSTER_SUSPEND"
	GPUBaremetalClusterTaskStatusError           GPUBaremetalClusterTaskStatus = "ERROR"
	GPUBaremetalClusterTaskStatusFinished        GPUBaremetalClusterTaskStatus = "FINISHED"
	GPUBaremetalClusterTaskStatusIpuServers      GPUBaremetalClusterTaskStatus = "IPU_SERVERS"
	GPUBaremetalClusterTaskStatusNetwork         GPUBaremetalClusterTaskStatus = "NETWORK"
	GPUBaremetalClusterTaskStatusPoplarServers   GPUBaremetalClusterTaskStatus = "POPLAR_SERVERS"
	GPUBaremetalClusterTaskStatusPostDeploySetup GPUBaremetalClusterTaskStatus = "POST_DEPLOY_SETUP"
	GPUBaremetalClusterTaskStatusVipuController  GPUBaremetalClusterTaskStatus = "VIPU_CONTROLLER"
)

type GPUBaremetalClusterServer struct {
	// GPU server ID
	ID string `json:"id,required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]GPUBaremetalClusterServerAddressUnion `json:"addresses,required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []BlackholePort `json:"blackhole_ports,required"`
	// Datetime when GPU server was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Advanced DDoS protection profile. It is always `null` if query parameter
	// `with_ddos=true` is not set.
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// Fixed IP assigned to instance
	FixedIPAssignments []GPUBaremetalClusterServerFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// Flavor
	Flavor GPUBaremetalClusterServerFlavor `json:"flavor,required"`
	// Instance description
	InstanceDescription string `json:"instance_description,required"`
	// Instance isolation information
	InstanceIsolation InstanceIsolation `json:"instance_isolation,required"`
	// GPU server name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Security groups
	SecurityGroups []GPUBaremetalClusterServerSecurityGroup `json:"security_groups,required"`
	// SSH key name assigned to instance
	SSHKeyName string `json:"ssh_key_name,required"`
	// Instance status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status GPUBaremetalClusterServerStatus `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required"`
	// Task state
	TaskState string `json:"task_state,required"`
	// Virtual machine state (active)
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState GPUBaremetalClusterServerVmState `json:"vm_state,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Addresses           respjson.Field
		BlackholePorts      respjson.Field
		CreatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		DDOSProfile         respjson.Field
		FixedIPAssignments  respjson.Field
		Flavor              respjson.Field
		InstanceDescription respjson.Field
		InstanceIsolation   respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		SecurityGroups      respjson.Field
		SSHKeyName          respjson.Field
		Status              respjson.Field
		Tags                respjson.Field
		TaskID              respjson.Field
		TaskState           respjson.Field
		VmState             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServer) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalClusterServerAddressUnion contains all possible properties and
// values from [FloatingAddress], [FixedAddressShort], [FixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUBaremetalClusterServerAddressUnion struct {
	Addr          string `json:"addr"`
	Type          string `json:"type"`
	InterfaceName string `json:"interface_name"`
	// This field is from variant [FixedAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [FixedAddress].
	SubnetName string `json:"subnet_name"`
	JSON       struct {
		Addr          respjson.Field
		Type          respjson.Field
		InterfaceName respjson.Field
		SubnetID      respjson.Field
		SubnetName    respjson.Field
		raw           string
	} `json:"-"`
}

func (u GPUBaremetalClusterServerAddressUnion) AsFloatingIPAddress() (v FloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServerAddressUnion) AsFixedIPAddressShort() (v FixedAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServerAddressUnion) AsFixedIPAddress() (v FixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUBaremetalClusterServerAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUBaremetalClusterServerAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerFixedIPAssignment struct {
	// Is network external
	External bool `json:"external,required"`
	// Ip address
	IPAddress string `json:"ip_address,required"`
	// Interface subnet id
	SubnetID string `json:"subnet_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		External    respjson.Field
		IPAddress   respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor
type GPUBaremetalClusterServerFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription GPUBaremetalClusterServerFlavorHardwareDescription `json:"hardware_description,required"`
	// Operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		ResourceClass       respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerFlavor) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type GPUBaremetalClusterServerFlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// If the flavor is licensed, this field contains the license type
	License string `json:"license,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		License     respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerSecurityGroup struct {
	// Name.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance status
type GPUBaremetalClusterServerStatus string

const (
	GPUBaremetalClusterServerStatusActive           GPUBaremetalClusterServerStatus = "ACTIVE"
	GPUBaremetalClusterServerStatusBuild            GPUBaremetalClusterServerStatus = "BUILD"
	GPUBaremetalClusterServerStatusDeleted          GPUBaremetalClusterServerStatus = "DELETED"
	GPUBaremetalClusterServerStatusError            GPUBaremetalClusterServerStatus = "ERROR"
	GPUBaremetalClusterServerStatusHardReboot       GPUBaremetalClusterServerStatus = "HARD_REBOOT"
	GPUBaremetalClusterServerStatusMigrating        GPUBaremetalClusterServerStatus = "MIGRATING"
	GPUBaremetalClusterServerStatusPassword         GPUBaremetalClusterServerStatus = "PASSWORD"
	GPUBaremetalClusterServerStatusPaused           GPUBaremetalClusterServerStatus = "PAUSED"
	GPUBaremetalClusterServerStatusReboot           GPUBaremetalClusterServerStatus = "REBOOT"
	GPUBaremetalClusterServerStatusRebuild          GPUBaremetalClusterServerStatus = "REBUILD"
	GPUBaremetalClusterServerStatusRescue           GPUBaremetalClusterServerStatus = "RESCUE"
	GPUBaremetalClusterServerStatusResize           GPUBaremetalClusterServerStatus = "RESIZE"
	GPUBaremetalClusterServerStatusRevertResize     GPUBaremetalClusterServerStatus = "REVERT_RESIZE"
	GPUBaremetalClusterServerStatusShelved          GPUBaremetalClusterServerStatus = "SHELVED"
	GPUBaremetalClusterServerStatusShelvedOffloaded GPUBaremetalClusterServerStatus = "SHELVED_OFFLOADED"
	GPUBaremetalClusterServerStatusShutoff          GPUBaremetalClusterServerStatus = "SHUTOFF"
	GPUBaremetalClusterServerStatusSoftDeleted      GPUBaremetalClusterServerStatus = "SOFT_DELETED"
	GPUBaremetalClusterServerStatusSuspended        GPUBaremetalClusterServerStatus = "SUSPENDED"
	GPUBaremetalClusterServerStatusUnknown          GPUBaremetalClusterServerStatus = "UNKNOWN"
	GPUBaremetalClusterServerStatusVerifyResize     GPUBaremetalClusterServerStatus = "VERIFY_RESIZE"
)

// Virtual machine state (active)
type GPUBaremetalClusterServerVmState string

const (
	GPUBaremetalClusterServerVmStateActive           GPUBaremetalClusterServerVmState = "active"
	GPUBaremetalClusterServerVmStateBuilding         GPUBaremetalClusterServerVmState = "building"
	GPUBaremetalClusterServerVmStateDeleted          GPUBaremetalClusterServerVmState = "deleted"
	GPUBaremetalClusterServerVmStateError            GPUBaremetalClusterServerVmState = "error"
	GPUBaremetalClusterServerVmStatePaused           GPUBaremetalClusterServerVmState = "paused"
	GPUBaremetalClusterServerVmStateRescued          GPUBaremetalClusterServerVmState = "rescued"
	GPUBaremetalClusterServerVmStateResized          GPUBaremetalClusterServerVmState = "resized"
	GPUBaremetalClusterServerVmStateShelved          GPUBaremetalClusterServerVmState = "shelved"
	GPUBaremetalClusterServerVmStateShelvedOffloaded GPUBaremetalClusterServerVmState = "shelved_offloaded"
	GPUBaremetalClusterServerVmStateSoftDeleted      GPUBaremetalClusterServerVmState = "soft-deleted"
	GPUBaremetalClusterServerVmStateStopped          GPUBaremetalClusterServerVmState = "stopped"
	GPUBaremetalClusterServerVmStateSuspended        GPUBaremetalClusterServerVmState = "suspended"
)

type GPUBaremetalClusterServerList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUBaremetalClusterServer `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerList) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnion contains all possible properties and values from
// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice],
// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUBaremetalFlavorUnion struct {
	Architecture string `json:"architecture"`
	Capacity     int64  `json:"capacity"`
	Disabled     bool   `json:"disabled"`
	// This field is a union of
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription],
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription]
	HardwareDescription GPUBaremetalFlavorUnionHardwareDescription `json:"hardware_description"`
	// This field is a union of
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties],
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties]
	HardwareProperties GPUBaremetalFlavorUnionHardwareProperties `json:"hardware_properties"`
	Name               string                                    `json:"name"`
	// This field is from variant
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice].
	Price GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice `json:"price"`
	JSON  struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		Price               respjson.Field
		raw                 string
	} `json:"-"`
}

func (u GPUBaremetalFlavorUnion) AsBareMetalGPUFlavorsChemaWithoutPrice() (v GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalFlavorUnion) AsBareMetalGPUFlavorsChemaWithPrice() (v GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUBaremetalFlavorUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUBaremetalFlavorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnionHardwareDescription is an implicit subunion of
// [GPUBaremetalFlavorUnion]. GPUBaremetalFlavorUnionHardwareDescription provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalFlavorUnion].
type GPUBaremetalFlavorUnionHardwareDescription struct {
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	GPU     string `json:"gpu"`
	Network string `json:"network"`
	Ram     string `json:"ram"`
	JSON    struct {
		CPU     respjson.Field
		Disk    respjson.Field
		GPU     respjson.Field
		Network respjson.Field
		Ram     respjson.Field
		raw     string
	} `json:"-"`
}

func (r *GPUBaremetalFlavorUnionHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnionHardwareProperties is an implicit subunion of
// [GPUBaremetalFlavorUnion]. GPUBaremetalFlavorUnionHardwareProperties provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalFlavorUnion].
type GPUBaremetalFlavorUnionHardwareProperties struct {
	GPUCount        int64  `json:"gpu_count"`
	GPUManufacturer string `json:"gpu_manufacturer"`
	GPUModel        string `json:"gpu_model"`
	JSON            struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		raw             string
	} `json:"-"`
}

func (r *GPUBaremetalFlavorUnionHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional bare metal hardware description
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware description
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware properties
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count,required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// GPU model
	GPUModel string `json:"gpu_model,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price
	Price GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice `json:"price,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		Price               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware description
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware properties
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count,required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// GPU model
	GPUModel string `json:"gpu_model,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor price
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice struct {
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code,required"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,required"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode  respjson.Field
		PricePerHour  respjson.Field
		PricePerMonth respjson.Field
		PriceStatus   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUBaremetalFlavorUnion `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorList) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flavor name
	Flavor string `json:"flavor,required"`
	// Image ID
	ImageID string `json:"image_id,required" format:"uuid4"`
	// A list of network interfaces for the server. You can create one or more
	// interfaces - private, public, or both.
	Interfaces []GPUBaremetalClusterNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// GPU Cluster name
	Name string `json:"name,required"`
	// Number of servers to create
	InstancesCount param.Opt[int64] `json:"instances_count,omitzero"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password param.Opt[string] `json:"password,omitzero"`
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username param.Opt[string] `json:"username,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsInterfaceUnion struct {
	OfExternal  *GPUBaremetalClusterNewParamsInterfaceExternal  `json:",omitzero,inline"`
	OfSubnet    *GPUBaremetalClusterNewParamsInterfaceSubnet    `json:",omitzero,inline"`
	OfAnySubnet *GPUBaremetalClusterNewParamsInterfaceAnySubnet `json:",omitzero,inline"`
	paramUnion
}

func (u GPUBaremetalClusterNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet)
}
func (u *GPUBaremetalClusterNewParamsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GPUBaremetalClusterNewParamsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfExternal) {
		return u.OfExternal
	} else if !param.IsOmitted(u.OfSubnet) {
		return u.OfSubnet
	} else if !param.IsOmitted(u.OfAnySubnet) {
		return u.OfAnySubnet
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.Valid() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetType() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSubnet; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfExternal; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfSubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfSubnet; vt != nil {
		return (*string)(&vt.NetworkID)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.NetworkID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetFloatingIP() (res gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = &vt.FloatingIP
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = &vt.FloatingIP
	}
	return
}

// Can have the runtime types
// [*GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP],
// [*GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP]
type gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP:
		return (*string)(&vt.Source)
	case *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP:
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsInterfaceUnion](
		"type",
		apijson.Discriminator[GPUBaremetalClusterNewParamsInterfaceExternal]("external"),
		apijson.Discriminator[GPUBaremetalClusterNewParamsInterfaceSubnet]("subnet"),
		apijson.Discriminator[GPUBaremetalClusterNewParamsInterfaceAnySubnet]("any_subnet"),
	)
}

// The property Type is required.
type GPUBaremetalClusterNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// A public IP address will be assigned to the server.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, SubnetID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceSubnet struct {
	// The network where the server will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// The server will get an IP address from this subnet.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP `json:"floating_ip,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP() GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP {
	return GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP{
		Source: "new",
	}
}

// Floating IP config for this subnet attachment
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP].
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnet struct {
	// The network where the server will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Allows the server to have a public IP that can be reached from the internet.
	FloatingIP GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Server will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP() GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP {
	return GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP{
		Source: "new",
	}
}

// Allows the server to have a public IP that can be reached from the internet.
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP].
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit the number of returned clusters
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterListParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// True if it is required to delete floating IPs assigned to the servers. Can't be
	// used with floatings.
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	// Comma separated list of floating ids that should be deleted. Can't be used with
	// `delete_floatings`.
	Floatings param.Opt[string] `query:"floatings,omitzero" json:"-"`
	// Comma separated list of port IDs to be deleted with the servers
	ReservedFixedIPs param.Opt[string] `query:"reserved_fixed_ips,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterDeleteParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterPowercycleAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebootAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebuildParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of nodes uuids to be rebuild
	Nodes []string `json:"nodes,omitzero,required"`
	// AI GPU image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// String in base64 format.Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterRebuildParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Resized (total) number of instances
	InstancesCount int64 `json:"instances_count,required"`
	paramObj
}

func (r GPUBaremetalClusterResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
