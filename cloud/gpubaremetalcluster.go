// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
	"github.com/tidwall/gjson"
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
	return
}

// Create a new GPU cluster.
func (r *GPUBaremetalClusterService) New(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List GPU clusters
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// List GPU clusters
func (r *GPUBaremetalClusterService) ListAutoPaging(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GPUBaremetalCluster] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete GPU cluster
func (r *GPUBaremetalClusterService) Delete(ctx context.Context, clusterID string, params GPUBaremetalClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// Get GPU cluster
func (r *GPUBaremetalClusterService) Get(ctx context.Context, clusterID string, query GPUBaremetalClusterGetParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
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

// Powercycle (stop and start) all GPU cluster nodes, aka hard reboot
func (r *GPUBaremetalClusterService) PowercycleAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterPowercycleAllServersParams, opts ...option.RequestOption) (res *GPUClusterServerList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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

// Reboot all GPU cluster nodes
func (r *GPUBaremetalClusterService) RebootAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterRebootAllServersParams, opts ...option.RequestOption) (res *GPUClusterServerList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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

// Rebuild one or many nodes from GPU cluster. All cluster nodes need to be
// provided to change the cluster image.
func (r *GPUBaremetalClusterService) Rebuild(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// Resize an existing AI GPU cluster.
func (r *GPUBaremetalClusterService) Resize(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	Servers []GPUClusterServer `json:"servers,required"`
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
	// 'password'. Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData string `json:"user_data,required"`
	// A name of a new user in the Linux instance. It may be passed with a 'password'
	// parameter
	Username string `json:"username,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ClusterID     resp.Field
		ClusterName   resp.Field
		ClusterStatus resp.Field
		CreatedAt     resp.Field
		CreatorTaskID resp.Field
		Flavor        resp.Field
		ImageID       resp.Field
		ImageName     resp.Field
		Interfaces    resp.Field
		Password      resp.Field
		ProjectID     resp.Field
		Region        resp.Field
		RegionID      resp.Field
		Servers       resp.Field
		SSHKeyName    resp.Field
		Tags          resp.Field
		TaskID        resp.Field
		TaskStatus    resp.Field
		UserData      resp.Field
		Username      resp.Field
		ExtraFields   map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		NetworkID   resp.Field
		PortID      resp.Field
		SubnetID    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
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
		Architecture        resp.Field
		Capacity            resp.Field
		Disabled            resp.Field
		HardwareDescription resp.Field
		HardwareProperties  resp.Field
		Name                resp.Field
		Price               resp.Field
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
		CPU     resp.Field
		Disk    resp.Field
		GPU     resp.Field
		Network resp.Field
		Ram     resp.Field
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
		GPUCount        resp.Field
		GPUManufacturer resp.Field
		GPUModel        resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		Capacity            resp.Field
		Disabled            resp.Field
		HardwareDescription resp.Field
		HardwareProperties  resp.Field
		Name                resp.Field
		ExtraFields         map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		GPU         resp.Field
		Network     resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GPUCount        resp.Field
		GPUManufacturer resp.Field
		GPUModel        resp.Field
		ExtraFields     map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		Capacity            resp.Field
		Disabled            resp.Field
		HardwareDescription resp.Field
		HardwareProperties  resp.Field
		Name                resp.Field
		Price               resp.Field
		ExtraFields         map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		GPU         resp.Field
		Network     resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GPUCount        resp.Field
		GPUManufacturer resp.Field
		GPUModel        resp.Field
		ExtraFields     map[string]resp.Field
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
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,required"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,required"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CurrencyCode  resp.Field
		PricePerHour  resp.Field
		PricePerMonth resp.Field
		PriceStatus   resp.Field
		ExtraFields   map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
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
	// Subnet IPs and floating IPs
	Interfaces []GPUBaremetalClusterNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// GPU Cluster name
	Name string `json:"name,required"`
	// Number of servers to create
	InstancesCount param.Opt[int64] `json:"instances_count,omitzero"`
	// A password for a bare metal server. This parameter is used to set a password for
	// the "Admin" user on a Windows instance, a default user or a new user on a Linux
	// instance
	Password param.Opt[string] `json:"password,omitzero"`
	// Specifies the name of the SSH keypair, created via the `/v1/ssh_keys` endpoint.
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the user_data:
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
	Tags TagUpdateList `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GPUBaremetalClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsInterfaceUnion struct {
	OfExternal        *GPUBaremetalClusterNewParamsInterfaceExternal        `json:",omitzero,inline"`
	OfSubnet          *GPUBaremetalClusterNewParamsInterfaceSubnet          `json:",omitzero,inline"`
	OfAnySubnet       *GPUBaremetalClusterNewParamsInterfaceAnySubnet       `json:",omitzero,inline"`
	OfReservedFixedIP *GPUBaremetalClusterNewParamsInterfaceReservedFixedIP `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u GPUBaremetalClusterNewParamsInterfaceUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u GPUBaremetalClusterNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GPUBaremetalClusterNewParamsInterfaceUnion](u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}

func (u *GPUBaremetalClusterNewParamsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfExternal) {
		return u.OfExternal
	} else if !param.IsOmitted(u.OfSubnet) {
		return u.OfSubnet
	} else if !param.IsOmitted(u.OfAnySubnet) {
		return u.OfAnySubnet
	} else if !param.IsOmitted(u.OfReservedFixedIP) {
		return u.OfReservedFixedIP
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
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.IsPresent() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfReservedFixedIP; vt != nil {
		return &vt.PortID
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
	} else if vt := u.OfReservedFixedIP; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfExternal; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfSubnet; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.InterfaceName.IsPresent() {
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
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetPortGroup() *int64 {
	if vt := u.OfExternal; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfSubnet; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
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
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetSecurityGroups() (res gpuBaremetalClusterNewParamsInterfaceUnionSecurityGroups) {
	if vt := u.OfExternal; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfSubnet; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfReservedFixedIP; vt != nil {
		res.any = &vt.SecurityGroups
	}
	return
}

// Can have the runtime types
// [_[]GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup],
// [_[]GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup],
// [_[]GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup],
// [_[]GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup]
type gpuBaremetalClusterNewParamsInterfaceUnionSecurityGroups struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]cloud.GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup:
//	case *[]cloud.GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup:
//	case *[]cloud.GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup:
//	case *[]cloud.GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuBaremetalClusterNewParamsInterfaceUnionSecurityGroups) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u GPUBaremetalClusterNewParamsInterfaceUnion) GetFloatingIP() (res gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfReservedFixedIP; vt != nil {
		res.any = vt.FloatingIP.asAny()
	}
	return
}

// Can have the runtime types
// [*GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew],
// [*GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting],
// [*GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew],
// [*GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting],
// [*GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew],
// [*GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting]
type gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew:
//	case *cloud.GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetSource()
	case *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetSource()
	case *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u gpuBaremetalClusterNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsInterfaceUnion](
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceExternal{}),
			DiscriminatorValue: "external",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceSubnet{}),
			DiscriminatorValue: "subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceAnySubnet{}),
			DiscriminatorValue: "any_subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceReservedFixedIP{}),
			DiscriminatorValue: "reserved_fixed_ip",
		},
	)
}

// Instance will be attached to default external network
//
// The property Type is required.
type GPUBaremetalClusterNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Applies only to instances and is ignored for bare metal. Specifies security
	// group UUIDs to be applied to the instance network interface.
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup `json:"security_groups,omitzero"`
	// A public IP address will be assigned to the instance.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceExternal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The instance will get an IP address from the selected network. If you choose to
// add a floating IP, the instance will be reachable from the internet. Otherwise,
// it will only have a private IP within the network.
//
// The properties NetworkID, SubnetID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceSubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// The instance will get an IP address from this subnet.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Applies only to instances and is ignored for bare metal. Specifies security
	// group UUIDs to be applied to the instance network interface.
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup `json:"security_groups,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceSubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion struct {
	OfNew      *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewGPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew() GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew {
	return GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties NetworkID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Applies only to instances and is ignored for bare metal. Specifies security
	// group UUIDs to be applied to the instance network interface.
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Instance will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceAnySubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion struct {
	OfNew      *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewGPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew() GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew {
	return GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties PortID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIP struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID string `json:"port_id,required"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// Applies only to instances and is ignored for bare metal. Specifies security
	// group UUIDs to be applied to the instance network interface.
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// An existing available reserved fixed IP will be attached to the instance. If the
	// reserved IP is not public and you choose to add a floating IP, the instance will
	// be accessible from the internet.
	//
	// This field can be elided, and will marshal its zero value as
	// "reserved_fixed_ip".
	Type constant.ReservedFixedIP `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceReservedFixedIP) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion struct {
	OfNew      *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewGPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew() GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew {
	return GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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
	// delete_floatings.
	Floatings param.Opt[string] `query:"floatings,omitzero" json:"-"`
	// Comma separated list of port IDs to be deleted with the servers
	ReservedFixedIPs param.Opt[string] `query:"reserved_fixed_ips,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type GPUBaremetalClusterPowercycleAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterPowercycleAllServersParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterRebootAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterRebootAllServersParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterRebuildParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of nodes uuids to be rebuild
	Nodes []string `json:"nodes,omitzero,required"`
	// AI GPU image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// String in base64 format.Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterRebuildParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GPUBaremetalClusterRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type GPUBaremetalClusterResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Resized (total) number of instances
	InstancesCount int64 `json:"instances_count,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GPUBaremetalClusterResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
