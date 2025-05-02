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

// '#/components/schemas/AIClusterSerializer'
// "$.components.schemas.AIClusterSerializer"
type GPUBaremetalCluster struct {
	// '#/components/schemas/AIClusterSerializer/properties/cluster_id'
	// "$.components.schemas.AIClusterSerializer.properties.cluster_id"
	ClusterID string `json:"cluster_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterSerializer/properties/cluster_name'
	// "$.components.schemas.AIClusterSerializer.properties.cluster_name"
	ClusterName string `json:"cluster_name,required"`
	// '#/components/schemas/AIClusterSerializer/properties/cluster_status'
	// "$.components.schemas.AIClusterSerializer.properties.cluster_status"
	//
	// Any of "ACTIVE", "ERROR", "PENDING", "SUSPENDED".
	ClusterStatus GPUBaremetalClusterClusterStatus `json:"cluster_status,required"`
	// '#/components/schemas/AIClusterSerializer/properties/created_at/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.created_at.anyOf[0]"
	CreatedAt string `json:"created_at,required"`
	// '#/components/schemas/AIClusterSerializer/properties/creator_task_id'
	// "$.components.schemas.AIClusterSerializer.properties.creator_task_id"
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterSerializer/properties/flavor'
	// "$.components.schemas.AIClusterSerializer.properties.flavor"
	Flavor string `json:"flavor,required"`
	// '#/components/schemas/AIClusterSerializer/properties/image_id'
	// "$.components.schemas.AIClusterSerializer.properties.image_id"
	ImageID string `json:"image_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterSerializer/properties/image_name'
	// "$.components.schemas.AIClusterSerializer.properties.image_name"
	ImageName string `json:"image_name,required"`
	// '#/components/schemas/AIClusterSerializer/properties/interfaces/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.interfaces.anyOf[0]"
	Interfaces []GPUBaremetalClusterInterface `json:"interfaces,required"`
	// '#/components/schemas/AIClusterSerializer/properties/keypair_name/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.keypair_name.anyOf[0]"
	KeypairName string `json:"keypair_name,required"`
	// '#/components/schemas/AIClusterSerializer/properties/password/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.password.anyOf[0]"
	Password string `json:"password,required"`
	// '#/components/schemas/AIClusterSerializer/properties/project_id'
	// "$.components.schemas.AIClusterSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/AIClusterSerializer/properties/region'
	// "$.components.schemas.AIClusterSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/AIClusterSerializer/properties/region_id'
	// "$.components.schemas.AIClusterSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/AIClusterSerializer/properties/servers'
	// "$.components.schemas.AIClusterSerializer.properties.servers"
	Servers []GPUClusterServer `json:"servers,required"`
	// '#/components/schemas/AIClusterSerializer/properties/tags'
	// "$.components.schemas.AIClusterSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/AIClusterSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterSerializer/properties/task_status'
	// "$.components.schemas.AIClusterSerializer.properties.task_status"
	//
	// Any of "CLUSTER_CLEAN_UP", "CLUSTER_RESIZE", "CLUSTER_RESUME",
	// "CLUSTER_SUSPEND", "ERROR", "FINISHED", "IPU_SERVERS", "NETWORK",
	// "POPLAR_SERVERS", "POST_DEPLOY_SETUP", "VIPU_CONTROLLER".
	TaskStatus GPUBaremetalClusterTaskStatus `json:"task_status,required"`
	// '#/components/schemas/AIClusterSerializer/properties/user_data/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.user_data.anyOf[0]"
	UserData string `json:"user_data,required"`
	// '#/components/schemas/AIClusterSerializer/properties/username/anyOf/0'
	// "$.components.schemas.AIClusterSerializer.properties.username.anyOf[0]"
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
		KeypairName   resp.Field
		Password      resp.Field
		ProjectID     resp.Field
		Region        resp.Field
		RegionID      resp.Field
		Servers       resp.Field
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

// '#/components/schemas/AIClusterSerializer/properties/cluster_status'
// "$.components.schemas.AIClusterSerializer.properties.cluster_status"
type GPUBaremetalClusterClusterStatus string

const (
	GPUBaremetalClusterClusterStatusActive    GPUBaremetalClusterClusterStatus = "ACTIVE"
	GPUBaremetalClusterClusterStatusError     GPUBaremetalClusterClusterStatus = "ERROR"
	GPUBaremetalClusterClusterStatusPending   GPUBaremetalClusterClusterStatus = "PENDING"
	GPUBaremetalClusterClusterStatusSuspended GPUBaremetalClusterClusterStatus = "SUSPENDED"
)

// '#/components/schemas/AIClusterSerializer/properties/interfaces/anyOf/0/items'
// "$.components.schemas.AIClusterSerializer.properties.interfaces.anyOf[0].items"
type GPUBaremetalClusterInterface struct {
	// '#/components/schemas/AIClusterNetworkSerializer/properties/network_id'
	// "$.components.schemas.AIClusterNetworkSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterNetworkSerializer/properties/port_id'
	// "$.components.schemas.AIClusterNetworkSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterNetworkSerializer/properties/subnet_id'
	// "$.components.schemas.AIClusterNetworkSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/AIClusterNetworkSerializer/properties/type'
	// "$.components.schemas.AIClusterNetworkSerializer.properties.type"
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

// '#/components/schemas/AIClusterSerializer/properties/task_status'
// "$.components.schemas.AIClusterSerializer.properties.task_status"
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

// '#/components/schemas/GpuBaremetalFlavorSerializer/anyOf/0'
// "$.components.schemas.GpuBaremetalFlavorSerializer.anyOf[0]"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice struct {
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/architecture/anyOf/0'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.architecture.anyOf[0]"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/capacity'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.capacity"
	Capacity int64 `json:"capacity,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/disabled'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.disabled"
	Disabled bool `json:"disabled,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/hardware_description'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.hardware_description"
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/hardware_properties'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.hardware_properties"
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties `json:"hardware_properties,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/name'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.name"
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

// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/hardware_description'
// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.hardware_description"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription struct {
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/cpu'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.cpu"
	CPU string `json:"cpu,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/disk'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.disk"
	Disk string `json:"disk,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/gpu'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.gpu"
	GPU string `json:"gpu,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/network'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.network"
	Network string `json:"network,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.ram"
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

// '#/components/schemas/GpuBaremetalFlavorSerializerWithoutPrice/properties/hardware_properties'
// "$.components.schemas.GpuBaremetalFlavorSerializerWithoutPrice.properties.hardware_properties"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties struct {
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_count/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_count.anyOf[0]"
	GPUCount int64 `json:"gpu_count,required"`
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_manufacturer/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_manufacturer.anyOf[0]"
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_model/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_model.anyOf[0]"
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

// '#/components/schemas/GpuBaremetalFlavorSerializer/anyOf/1'
// "$.components.schemas.GpuBaremetalFlavorSerializer.anyOf[1]"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice struct {
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/architecture/anyOf/0'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.architecture.anyOf[0]"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/capacity'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.capacity"
	Capacity int64 `json:"capacity,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/disabled'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.disabled"
	Disabled bool `json:"disabled,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/hardware_description'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.hardware_description"
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/hardware_properties'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.hardware_properties"
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties `json:"hardware_properties,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/name'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/price'
	// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.price"
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

// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/hardware_description'
// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.hardware_description"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription struct {
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/cpu'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.cpu"
	CPU string `json:"cpu,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/disk'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.disk"
	Disk string `json:"disk,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/gpu'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.gpu"
	GPU string `json:"gpu,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/network'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.network"
	Network string `json:"network,required"`
	// '#/components/schemas/GpuBaremetalFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.GpuBaremetalFlavorHardwareDescriptionSerializer.properties.ram"
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

// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/hardware_properties'
// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.hardware_properties"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties struct {
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_count/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_count.anyOf[0]"
	GPUCount int64 `json:"gpu_count,required"`
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_manufacturer/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_manufacturer.anyOf[0]"
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// '#/components/schemas/GpuFlavorHardwareProperties/properties/gpu_model/anyOf/0'
	// "$.components.schemas.GpuFlavorHardwareProperties.properties.gpu_model.anyOf[0]"
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

// '#/components/schemas/GpuBaremetalFlavorSerializerWithPrices/properties/price'
// "$.components.schemas.GpuBaremetalFlavorSerializerWithPrices.properties.price"
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice struct {
	// '#/components/schemas/FlavorPrice/properties/currency_code/anyOf/0'
	// "$.components.schemas.FlavorPrice.properties.currency_code.anyOf[0]"
	CurrencyCode string `json:"currency_code,required"`
	// '#/components/schemas/FlavorPrice/properties/price_per_hour/anyOf/0'
	// "$.components.schemas.FlavorPrice.properties.price_per_hour.anyOf[0]"
	PricePerHour float64 `json:"price_per_hour,required"`
	// '#/components/schemas/FlavorPrice/properties/price_per_month/anyOf/0'
	// "$.components.schemas.FlavorPrice.properties.price_per_month.anyOf[0]"
	PricePerMonth float64 `json:"price_per_month,required"`
	// '#/components/schemas/FlavorPrice/properties/price_status/anyOf/0'
	// "$.components.schemas.FlavorPrice.properties.price_status.anyOf[0]"
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

// '#/components/schemas/ListGpuBaremetalFlavorSerializer'
// "$.components.schemas.ListGpuBaremetalFlavorSerializer"
type GPUBaremetalFlavorList struct {
	// '#/components/schemas/ListGpuBaremetalFlavorSerializer/properties/count'
	// "$.components.schemas.ListGpuBaremetalFlavorSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ListGpuBaremetalFlavorSerializer/properties/results'
	// "$.components.schemas.ListGpuBaremetalFlavorSerializer.properties.results"
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
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/flavor'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.flavor"
	Flavor string `json:"flavor,required"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/image_id'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.image_id"
	ImageID string `json:"image_id,required" format:"uuid4"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/interfaces'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.interfaces"
	Interfaces []GPUBaremetalClusterNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/name'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/instances_count'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.instances_count"
	InstancesCount param.Opt[int64] `json:"instances_count,omitzero"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/keypair_name'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.keypair_name"
	KeypairName param.Opt[string] `json:"keypair_name,omitzero"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/password'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.password"
	Password param.Opt[string] `json:"password,omitzero"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/user_data'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.user_data"
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/username'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.username"
	Username param.Opt[string] `json:"username,omitzero"`
	// '#/components/schemas/CreateAIClusterGPUSerializer/properties/tags'
	// "$.components.schemas.CreateAIClusterGPUSerializer.properties.tags"
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
	// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/0'
	// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[0]"
	OfExternal *GPUBaremetalClusterNewParamsInterfaceExternal `json:",omitzero,inline"`
	// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/1'
	// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[1]"
	OfSubnet *GPUBaremetalClusterNewParamsInterfaceSubnet `json:",omitzero,inline"`
	// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/2'
	// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[2]"
	OfAnySubnet *GPUBaremetalClusterNewParamsInterfaceAnySubnet `json:",omitzero,inline"`
	// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/3'
	// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[3]"
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

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/0'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[0]"
//
// The property Type is required.
type GPUBaremetalClusterNewParamsInterfaceExternal struct {
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/interface_name'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/port_group'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/ip_family/anyOf/0'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.security_groups"
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.type"
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

// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceExternalSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
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

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/1'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[1]"
//
// The properties NetworkID, SubnetID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceSubnet struct {
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/network_id'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/subnet_id'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/interface_name'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/port_group'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip"
	FloatingIP GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.security_groups"
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.type"
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
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
	OfNew *GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceSubnetFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceSubnetSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
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

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/2'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[2]"
//
// The properties NetworkID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnet struct {
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/network_id'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/interface_name'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/ip_address'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.ip_address"
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/port_group'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip"
	FloatingIP GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/ip_family/anyOf/0'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.security_groups"
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.type"
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
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
	OfNew *GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceAnySubnetSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
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

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/3'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[3]"
//
// The properties PortID, Type are required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIP struct {
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/port_id'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.port_id"
	PortID string `json:"port_id,required"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/interface_name'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/port_group'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip"
	FloatingIP GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.security_groups"
	SecurityGroups []GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.type"
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
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/0'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[0]"
	OfNew *GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/1'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew].
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type GPUBaremetalClusterNewParamsInterfaceReservedFixedIPSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
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
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}'].get.parameters[3]"
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
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/delete/parameters/3'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}']['delete'].parameters[3]"
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/delete/parameters/4'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}']['delete'].parameters[4]"
	Floatings param.Opt[string] `query:"floatings,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/delete/parameters/5'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}']['delete'].parameters[5]"
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
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type GPUBaremetalClusterPowercycleAllServersParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fpowercycle/post/parameters/0/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}/powercycle'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fpowercycle/post/parameters/1/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}/powercycle'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterPowercycleAllServersParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterRebootAllServersParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Freboot/post/parameters/0/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}/reboot'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Freboot/post/parameters/1/schema'
	// "$.paths['/cloud/v2/ai/clusters/{project_id}/{region_id}/{cluster_id}/reboot'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterRebootAllServersParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type GPUBaremetalClusterRebuildParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Frebuild/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/rebuild'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Frebuild/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/rebuild'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RebuildClusterSerializer/properties/nodes'
	// "$.components.schemas.RebuildClusterSerializer.properties.nodes"
	Nodes []string `json:"nodes,omitzero,required"`
	// '#/components/schemas/RebuildClusterSerializer/properties/image_id/anyOf/0'
	// "$.components.schemas.RebuildClusterSerializer.properties.image_id.anyOf[0]"
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// '#/components/schemas/RebuildClusterSerializer/properties/user_data/anyOf/0'
	// "$.components.schemas.RebuildClusterSerializer.properties.user_data.anyOf[0]"
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
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fresize/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/resize'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2Fgpu%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Fresize/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/gpu/{project_id}/{region_id}/{cluster_id}/resize'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ResizeAIClusterGPUSerializerV1/properties/instances_count'
	// "$.components.schemas.ResizeAIClusterGPUSerializerV1.properties.instances_count"
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
