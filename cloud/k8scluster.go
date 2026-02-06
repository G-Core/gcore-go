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

// K8SClusterService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SClusterService] method instead.
type K8SClusterService struct {
	Options []option.RequestOption
	Nodes   K8SClusterNodeService
	Pools   K8SClusterPoolService
	tasks   TaskService
}

// NewK8SClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewK8SClusterService(opts ...option.RequestOption) (r K8SClusterService) {
	r = K8SClusterService{}
	r.Options = opts
	r.Nodes = NewK8SClusterNodeService(opts...)
	r.Pools = NewK8SClusterPoolService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create k8s cluster
func (r *K8SClusterService) New(ctx context.Context, params K8SClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new k8s cluster and polls for completion
func (r *K8SClusterService) NewAndPoll(ctx context.Context, params K8SClusterNewParams, opts ...option.RequestOption) (v *K8SCluster, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams K8SClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	// for k8s cluster creation the task.CreatedResources.K8SClusters only contains the cluster ID and not the cluster
	// name, which is the path parameter required to retrieve the created cluster. Therefore, we use the params.Name
	// instead.
	return r.Get(ctx, params.Name, getParams, opts...)
}

// Update k8s cluster
func (r *K8SClusterService) Update(ctx context.Context, clusterName string, params K8SClusterUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// UpdateAndPoll updates a k8s cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *K8SClusterService) UpdateAndPoll(ctx context.Context, clusterName string, params K8SClusterUpdateParams, opts ...option.RequestOption) (v *K8SCluster, err error) {
	resource, err := r.Update(ctx, clusterName, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams K8SClusterGetParams
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
		return nil, err
	}

	return r.Get(ctx, clusterName, getParams, opts...)
}

// List k8s clusters
func (r *K8SClusterService) List(ctx context.Context, query K8SClusterListParams, opts ...option.RequestOption) (res *K8SClusterList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete k8s cluster
func (r *K8SClusterService) Delete(ctx context.Context, clusterName string, params K8SClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// DeleteAndPoll deletes a k8s cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *K8SClusterService) DeleteAndPoll(ctx context.Context, clusterName string, params K8SClusterDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, clusterName, params, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Get k8s cluster
func (r *K8SClusterService) Get(ctx context.Context, clusterName string, query K8SClusterGetParams, opts ...option.RequestOption) (res *K8SCluster, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get k8s cluster CA certificate
func (r *K8SClusterService) GetCertificate(ctx context.Context, clusterName string, query K8SClusterGetCertificateParams, opts ...option.RequestOption) (res *K8SClusterCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/certificates", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get k8s cluster kubeconfig
func (r *K8SClusterService) GetKubeconfig(ctx context.Context, clusterName string, query K8SClusterGetKubeconfigParams, opts ...option.RequestOption) (res *K8SClusterKubeconfig, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/config", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List available k8s cluster versions for upgrade
func (r *K8SClusterService) ListVersionsForUpgrade(ctx context.Context, clusterName string, query K8SClusterListVersionsForUpgradeParams, opts ...option.RequestOption) (res *K8SClusterVersionList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/upgrade_versions", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upgrade k8s cluster
func (r *K8SClusterService) Upgrade(ctx context.Context, clusterName string, params K8SClusterUpgradeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/upgrade", params.ProjectID.Value, params.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpgradeAndPoll upgrades a k8s cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *K8SClusterService) UpgradeAndPoll(ctx context.Context, clusterName string, params K8SClusterUpgradeParams, opts ...option.RequestOption) (v *K8SCluster, err error) {
	resource, err := r.Upgrade(ctx, clusterName, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams K8SClusterGetParams
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
		return nil, err
	}

	return r.Get(ctx, clusterName, getParams, opts...)
}

type K8SCluster struct {
	// Cluster pool uuid
	ID string `json:"id,required"`
	// Cluster add-ons configuration
	AddOns K8SClusterAddOns `json:"add_ons,required"`
	// Function creation date
	CreatedAt string `json:"created_at,required"`
	// Cluster CSI settings
	Csi K8SClusterCsi `json:"csi,required"`
	// Cluster is public
	IsPublic bool `json:"is_public,required"`
	// Keypair
	Keypair string `json:"keypair,required"`
	// Logging configuration
	Logging Logging `json:"logging,required"`
	// Name
	Name string `json:"name,required"`
	// pools
	Pools []K8SClusterPool `json:"pools,required"`
	// Status
	//
	// Any of "Deleting", "Provisioned", "Provisioning".
	Status K8SClusterStatus `json:"status,required"`
	// K8s version
	Version string `json:"version,required"`
	// Cluster authentication settings
	Authentication K8SClusterAuthentication `json:"authentication,nullable"`
	// Cluster autoscaler configuration.
	//
	// It contains overrides to the default cluster-autoscaler parameters provided by
	// the platform.
	AutoscalerConfig map[string]string `json:"autoscaler_config,nullable"`
	// Cluster CNI settings
	Cni K8SClusterCni `json:"cni,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Advanced DDoS Protection profile
	DDOSProfile K8SClusterDDOSProfile `json:"ddos_profile,nullable"`
	// Fixed network id
	FixedNetwork string `json:"fixed_network,nullable"`
	// Fixed subnet id
	FixedSubnet string `json:"fixed_subnet,nullable"`
	// Enable public v6 address
	IsIpv6 bool `json:"is_ipv6"`
	// The IP pool for the pods
	PodsIPPool string `json:"pods_ip_pool,nullable"`
	// The IPv6 pool for the pods
	PodsIpv6Pool string `json:"pods_ipv6_pool,nullable"`
	// The IP pool for the services
	ServicesIPPool string `json:"services_ip_pool,nullable"`
	// The IPv6 pool for the services
	ServicesIpv6Pool string `json:"services_ipv6_pool,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AddOns           respjson.Field
		CreatedAt        respjson.Field
		Csi              respjson.Field
		IsPublic         respjson.Field
		Keypair          respjson.Field
		Logging          respjson.Field
		Name             respjson.Field
		Pools            respjson.Field
		Status           respjson.Field
		Version          respjson.Field
		Authentication   respjson.Field
		AutoscalerConfig respjson.Field
		Cni              respjson.Field
		CreatorTaskID    respjson.Field
		DDOSProfile      respjson.Field
		FixedNetwork     respjson.Field
		FixedSubnet      respjson.Field
		IsIpv6           respjson.Field
		PodsIPPool       respjson.Field
		PodsIpv6Pool     respjson.Field
		ServicesIPPool   respjson.Field
		ServicesIpv6Pool respjson.Field
		TaskID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SCluster) RawJSON() string { return r.JSON.raw }
func (r *K8SCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster add-ons configuration
type K8SClusterAddOns struct {
	// Slurm add-on configuration
	Slurm K8SClusterAddOnsSlurm `json:"slurm,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Slurm       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterAddOns) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterAddOns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Slurm add-on configuration
type K8SClusterAddOnsSlurm struct {
	// Indicates whether Slurm add-on is deployed in the cluster.
	//
	// This add-on is only supported in clusters running Kubernetes v1.31 and v1.32
	// with at least 1 GPU cluster pool.
	Enabled bool `json:"enabled,required"`
	// ID of a VAST file share used as Slurm storage.
	//
	// The Slurm add-on creates separate Persistent Volume Claims for different
	// purposes (controller spool, worker spool, jail) on that file share.
	FileShareID string `json:"file_share_id,required" format:"uuid4"`
	// IDs of SSH keys authorized for SSH connection to Slurm login nodes.
	SSHKeyIDs []string `json:"ssh_key_ids,required" format:"uuid4"`
	// Size of the worker pool, i.e. number of worker nodes.
	//
	// Each Slurm worker node is backed by a Pod scheduled on one of cluster's GPU
	// nodes.
	//
	// Note: Downscaling (reducing worker count) is not supported.
	WorkerCount int64 `json:"worker_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		FileShareID respjson.Field
		SSHKeyIDs   respjson.Field
		WorkerCount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterAddOnsSlurm) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterAddOnsSlurm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CSI settings
type K8SClusterCsi struct {
	// NFS settings
	Nfs K8SClusterCsiNfs `json:"nfs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Nfs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterCsi) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterCsi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NFS settings
type K8SClusterCsiNfs struct {
	// Indicates the status of VAST NFS integration
	VastEnabled bool `json:"vast_enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VastEnabled respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterCsiNfs) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterCsiNfs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status
type K8SClusterStatus string

const (
	K8SClusterStatusDeleting     K8SClusterStatus = "Deleting"
	K8SClusterStatusProvisioned  K8SClusterStatus = "Provisioned"
	K8SClusterStatusProvisioning K8SClusterStatus = "Provisioning"
)

// Cluster authentication settings
type K8SClusterAuthentication struct {
	// Kubeconfig creation date
	KubeconfigCreatedAt time.Time `json:"kubeconfig_created_at,nullable" format:"date-time"`
	// Kubeconfig expiration date
	KubeconfigExpiresAt time.Time `json:"kubeconfig_expires_at,nullable" format:"date-time"`
	// OIDC authentication settings
	Oidc K8SClusterAuthenticationOidc `json:"oidc,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KubeconfigCreatedAt respjson.Field
		KubeconfigExpiresAt respjson.Field
		Oidc                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterAuthentication) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8SClusterAuthenticationOidc struct {
	// Client ID
	ClientID string `json:"client_id,nullable"`
	// JWT claim to use as the user's group
	GroupsClaim string `json:"groups_claim,nullable"`
	// Prefix prepended to group claims
	GroupsPrefix string `json:"groups_prefix,nullable"`
	// Issuer URL
	IssuerURL string `json:"issuer_url,nullable"`
	// Key-value pairs that describe required claims in the token
	RequiredClaims map[string]string `json:"required_claims,nullable"`
	// Accepted signing algorithms
	//
	// Any of "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "RS256", "RS384",
	// "RS512".
	SigningAlgs []string `json:"signing_algs,nullable"`
	// JWT claim to use as the user name
	UsernameClaim string `json:"username_claim,nullable"`
	// Prefix prepended to username claims to prevent clashes
	UsernamePrefix string `json:"username_prefix,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID       respjson.Field
		GroupsClaim    respjson.Field
		GroupsPrefix   respjson.Field
		IssuerURL      respjson.Field
		RequiredClaims respjson.Field
		SigningAlgs    respjson.Field
		UsernameClaim  respjson.Field
		UsernamePrefix respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterAuthenticationOidc) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8SClusterCni struct {
	// Cilium settings
	Cilium K8SClusterCniCilium `json:"cilium,nullable"`
	// CNI provider
	//
	// Any of "calico", "cilium".
	Provider string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cilium      respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterCni) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cilium settings
type K8SClusterCniCilium struct {
	// Wireguard encryption
	Encryption bool `json:"encryption"`
	// Hubble Relay
	HubbleRelay bool `json:"hubble_relay"`
	// Hubble UI
	HubbleUi bool `json:"hubble_ui"`
	// LoadBalancer acceleration
	LbAcceleration bool `json:"lb_acceleration"`
	// LoadBalancer mode
	//
	// Any of "dsr", "hybrid", "snat".
	LbMode string `json:"lb_mode"`
	// Mask size for IPv4
	MaskSize int64 `json:"mask_size"`
	// Mask size for IPv6
	MaskSizeV6 int64 `json:"mask_size_v6"`
	// Routing mode
	//
	// Any of "native", "tunnel".
	RoutingMode string `json:"routing_mode"`
	// CNI provider
	//
	// Any of "", "geneve", "vxlan".
	Tunnel string `json:"tunnel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encryption     respjson.Field
		HubbleRelay    respjson.Field
		HubbleUi       respjson.Field
		LbAcceleration respjson.Field
		LbMode         respjson.Field
		MaskSize       respjson.Field
		MaskSizeV6     respjson.Field
		RoutingMode    respjson.Field
		Tunnel         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterCniCilium) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS Protection profile
type K8SClusterDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile parameters
	Fields []K8SClusterDDOSProfileField `json:"fields,nullable"`
	// DDoS profile template ID
	ProfileTemplate int64 `json:"profile_template,nullable"`
	// DDoS profile template name
	ProfileTemplateName string `json:"profile_template_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled             respjson.Field
		Fields              respjson.Field
		ProfileTemplate     respjson.Field
		ProfileTemplateName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterDDOSProfile) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Complex value. Only one of 'value' or 'field_value' must be specified
	FieldValue any `json:"field_value"`
	// Basic value. Only one of 'value' or 'field_value' must be specified
	Value string `json:"value,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseField   respjson.Field
		FieldValue  respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterDDOSProfileField) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterCertificate struct {
	// Cluster CA certificate
	Certificate string `json:"certificate,required"`
	// Cluster CA private key
	Key string `json:"key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Certificate respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterCertificate) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterKubeconfig struct {
	// String in base64 format. Cluster client certificate
	ClientCertificate string `json:"client_certificate,required"`
	// String in base64 format. Cluster client key
	ClientKey string `json:"client_key,required"`
	// String in base64 format. Cluster ca certificate
	ClusterCaCertificate string `json:"cluster_ca_certificate,required"`
	// Cluster kubeconfig
	Config string `json:"config,required"`
	// Kubeconfig creation date
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kubeconfig expiration date
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Cluster host
	Host string `json:"host,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientCertificate    respjson.Field
		ClientKey            respjson.Field
		ClusterCaCertificate respjson.Field
		Config               respjson.Field
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		Host                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterKubeconfig) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterKubeconfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []K8SCluster `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterList) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The keypair of the cluster
	Keypair string `json:"keypair,required"`
	// The name of the cluster
	Name string `json:"name,required"`
	// The pools of the cluster
	Pools []K8SClusterNewParamsPool `json:"pools,omitzero,required"`
	// The version of the k8s cluster
	Version string `json:"version,required"`
	// The network of the cluster
	FixedNetwork param.Opt[string] `json:"fixed_network,omitzero"`
	// The subnet of the cluster
	FixedSubnet param.Opt[string] `json:"fixed_subnet,omitzero"`
	// Enable public v6 address
	IsIpv6 param.Opt[bool] `json:"is_ipv6,omitzero"`
	// The IP pool for the pods
	PodsIPPool param.Opt[string] `json:"pods_ip_pool,omitzero"`
	// The IPv6 pool for the pods
	PodsIpv6Pool param.Opt[string] `json:"pods_ipv6_pool,omitzero"`
	// The IP pool for the services
	ServicesIPPool param.Opt[string] `json:"services_ip_pool,omitzero"`
	// The IPv6 pool for the services
	ServicesIpv6Pool param.Opt[string] `json:"services_ipv6_pool,omitzero"`
	// Authentication settings
	Authentication K8SClusterNewParamsAuthentication `json:"authentication,omitzero"`
	// Cluster autoscaler configuration.
	//
	// It allows you to override the default cluster-autoscaler parameters provided by
	// the platform with your preferred values.
	//
	// Supported parameters (in alphabetical order):
	//
	//   - balance-similar-node-groups (boolean: true/false) - Detect similar node groups
	//     and balance the number of nodes between them.
	//   - expander (string: random, most-pods, least-waste, price, priority, grpc) -
	//     Type of node group expander to be used in scale up. Specifying multiple values
	//     separated by commas will call the expanders in succession until there is only
	//     one option remaining.
	//   - expendable-pods-priority-cutoff (float) - Pods with priority below cutoff will
	//     be expendable. They can be killed without any consideration during scale down
	//     and they don't cause scale up. Pods with null priority (PodPriority disabled)
	//     are non expendable.
	//   - ignore-daemonsets-utilization (boolean: true/false) - Should CA ignore
	//     DaemonSet pods when calculating resource utilization for scaling down.
	//   - max-empty-bulk-delete (integer) - Maximum number of empty nodes that can be
	//     deleted at the same time.
	//   - max-graceful-termination-sec (integer) - Maximum number of seconds CA waits
	//     for pod termination when trying to scale down a node.
	//   - max-node-provision-time (duration: e.g., '15m') - The default maximum time CA
	//     waits for node to be provisioned - the value can be overridden per node group.
	//   - max-total-unready-percentage (float) - Maximum percentage of unready nodes in
	//     the cluster. After this is exceeded, CA halts operations.
	//   - new-pod-scale-up-delay (duration: e.g., '10s') - Pods less than this old will
	//     not be considered for scale-up. Can be increased for individual pods through
	//     annotation.
	//   - ok-total-unready-count (integer) - Number of allowed unready nodes,
	//     irrespective of max-total-unready-percentage.
	//   - scale-down-delay-after-add (duration: e.g., '10m') - How long after scale up
	//     that scale down evaluation resumes.
	//   - scale-down-delay-after-delete (duration: e.g., '10s') - How long after node
	//     deletion that scale down evaluation resumes.
	//   - scale-down-delay-after-failure (duration: e.g., '3m') - How long after scale
	//     down failure that scale down evaluation resumes.
	//   - scale-down-enabled (boolean: true/false) - Should CA scale down the cluster.
	//   - scale-down-unneeded-time (duration: e.g., '10m') - How long a node should be
	//     unneeded before it is eligible for scale down.
	//   - scale-down-unready-time (duration: e.g., '20m') - How long an unready node
	//     should be unneeded before it is eligible for scale down.
	//   - scale-down-utilization-threshold (float) - The maximum value between the sum
	//     of cpu requests and sum of memory requests of all pods running on the node
	//     divided by node's corresponding allocatable resource, below which a node can
	//     be considered for scale down.
	//   - scan-interval (duration: e.g., '10s') - How often cluster is reevaluated for
	//     scale up or down.
	//   - skip-nodes-with-custom-controller-pods (boolean: true/false) - If true cluster
	//     autoscaler will never delete nodes with pods owned by custom controllers.
	//   - skip-nodes-with-local-storage (boolean: true/false) - If true cluster
	//     autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir
	//     or HostPath.
	//   - skip-nodes-with-system-pods (boolean: true/false) - If true cluster autoscaler
	//     will never delete nodes with pods from kube-system (except for DaemonSet or
	//     mirror pods).
	AutoscalerConfig map[string]string `json:"autoscaler_config,omitzero"`
	// Cluster CNI settings
	Cni K8SClusterNewParamsCni `json:"cni,omitzero"`
	// Advanced DDoS Protection profile
	DDOSProfile K8SClusterNewParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// Logging configuration
	Logging K8SClusterNewParamsLogging `json:"logging,omitzero"`
	// Cluster add-ons configuration
	AddOns K8SClusterNewParamsAddOns `json:"add_ons,omitzero"`
	// Container Storage Interface (CSI) driver settings
	Csi K8SClusterNewParamsCsi `json:"csi,omitzero"`
	paramObj
}

func (r K8SClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FlavorID, MinNodeCount, Name are required.
type K8SClusterNewParamsPool struct {
	// Flavor ID
	FlavorID string `json:"flavor_id,required"`
	// Minimum node count
	MinNodeCount int64 `json:"min_node_count,required"`
	// Pool's name
	Name string `json:"name,required"`
	// Enable auto healing
	AutoHealingEnabled param.Opt[bool] `json:"auto_healing_enabled,omitzero"`
	// Boot volume size
	BootVolumeSize param.Opt[int64] `json:"boot_volume_size,omitzero"`
	// Enable public v4 address
	IsPublicIpv4 param.Opt[bool] `json:"is_public_ipv4,omitzero"`
	// Maximum node count
	MaxNodeCount param.Opt[int64] `json:"max_node_count,omitzero"`
	// Boot volume type
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	BootVolumeType string `json:"boot_volume_type,omitzero"`
	// Cri-o configuration for pool nodes
	CrioConfig map[string]string `json:"crio_config,omitzero"`
	// Kubelet configuration for pool nodes
	KubeletConfig map[string]string `json:"kubelet_config,omitzero"`
	// Labels applied to the cluster pool
	Labels map[string]string `json:"labels,omitzero"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	ServergroupPolicy string `json:"servergroup_policy,omitzero"`
	// Taints applied to the cluster pool
	Taints map[string]string `json:"taints,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsPool) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterNewParamsPool](
		"boot_volume_type", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
	apijson.RegisterFieldValidator[K8SClusterNewParamsPool](
		"servergroup_policy", "affinity", "anti-affinity", "soft-anti-affinity",
	)
}

// Cluster add-ons configuration
type K8SClusterNewParamsAddOns struct {
	// Slurm add-on configuration
	Slurm K8SClusterNewParamsAddOnsSlurm `json:"slurm,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsAddOns) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsAddOns
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsAddOns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Slurm add-on configuration
//
// The properties Enabled, FileShareID, SSHKeyIDs, WorkerCount are required.
type K8SClusterNewParamsAddOnsSlurm struct {
	// ID of a VAST file share to be used as Slurm storage.
	//
	// The Slurm add-on will create separate Persistent Volume Claims for different
	// purposes (controller spool, worker spool, jail) on that file share.
	//
	// The file share must have `root_squash` disabled, while `path_length` and
	// `allowed_characters` settings must be set to `NPL`.
	FileShareID string `json:"file_share_id,required" format:"uuid4"`
	// IDs of SSH keys to authorize for SSH connection to Slurm login nodes.
	SSHKeyIDs []string `json:"ssh_key_ids,omitzero,required" format:"uuid4"`
	// Size of the worker pool, i.e. the number of Slurm worker nodes.
	//
	// Each Slurm worker node will be backed by a Pod scheduled on one of cluster's GPU
	// nodes.
	//
	// Note: Downscaling (reducing worker count) is not supported.
	WorkerCount int64 `json:"worker_count,required"`
	// The Slurm add-on will be enabled in the cluster.
	//
	// This add-on is only supported in clusters running Kubernetes v1.31 and v1.32
	// with at least 1 GPU cluster pool and VAST NFS support enabled.
	//
	// This field can be elided, and will marshal its zero value as true.
	Enabled bool `json:"enabled,omitzero,required"`
	paramObj
}

func (r K8SClusterNewParamsAddOnsSlurm) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsAddOnsSlurm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsAddOnsSlurm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterNewParamsAddOnsSlurm](
		"enabled", true,
	)
}

// Authentication settings
type K8SClusterNewParamsAuthentication struct {
	// OIDC authentication settings
	Oidc K8SClusterNewParamsAuthenticationOidc `json:"oidc,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsAuthentication) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsAuthentication
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8SClusterNewParamsAuthenticationOidc struct {
	// Client ID
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// JWT claim to use as the user's group
	GroupsClaim param.Opt[string] `json:"groups_claim,omitzero"`
	// Prefix prepended to group claims
	GroupsPrefix param.Opt[string] `json:"groups_prefix,omitzero"`
	// Issuer URL
	IssuerURL param.Opt[string] `json:"issuer_url,omitzero"`
	// JWT claim to use as the user name
	UsernameClaim param.Opt[string] `json:"username_claim,omitzero"`
	// Prefix prepended to username claims to prevent clashes
	UsernamePrefix param.Opt[string] `json:"username_prefix,omitzero"`
	// Key-value pairs that describe required claims in the token
	RequiredClaims map[string]string `json:"required_claims,omitzero"`
	// Accepted signing algorithms
	//
	// Any of "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "RS256", "RS384",
	// "RS512".
	SigningAlgs []string `json:"signing_algs,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsAuthenticationOidc) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsAuthenticationOidc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8SClusterNewParamsCni struct {
	// Cilium settings
	Cilium K8SClusterNewParamsCniCilium `json:"cilium,omitzero"`
	// CNI provider
	//
	// Any of "calico", "cilium".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsCni) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsCni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterNewParamsCni](
		"provider", "calico", "cilium",
	)
}

// Cilium settings
type K8SClusterNewParamsCniCilium struct {
	// Wireguard encryption
	Encryption param.Opt[bool] `json:"encryption,omitzero"`
	// Hubble Relay
	HubbleRelay param.Opt[bool] `json:"hubble_relay,omitzero"`
	// Hubble UI
	HubbleUi param.Opt[bool] `json:"hubble_ui,omitzero"`
	// LoadBalancer acceleration
	LbAcceleration param.Opt[bool] `json:"lb_acceleration,omitzero"`
	// Mask size for IPv4
	MaskSize param.Opt[int64] `json:"mask_size,omitzero"`
	// Mask size for IPv6
	MaskSizeV6 param.Opt[int64] `json:"mask_size_v6,omitzero"`
	// LoadBalancer mode
	//
	// Any of "dsr", "hybrid", "snat".
	LbMode string `json:"lb_mode,omitzero"`
	// Routing mode
	//
	// Any of "native", "tunnel".
	RoutingMode string `json:"routing_mode,omitzero"`
	// CNI provider
	//
	// Any of "", "geneve", "vxlan".
	Tunnel string `json:"tunnel,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsCniCilium) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsCniCilium
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterNewParamsCniCilium](
		"lb_mode", "dsr", "hybrid", "snat",
	)
	apijson.RegisterFieldValidator[K8SClusterNewParamsCniCilium](
		"routing_mode", "native", "tunnel",
	)
	apijson.RegisterFieldValidator[K8SClusterNewParamsCniCilium](
		"tunnel", "", "geneve", "vxlan",
	)
}

// Container Storage Interface (CSI) driver settings
type K8SClusterNewParamsCsi struct {
	// NFS CSI driver settings
	Nfs K8SClusterNewParamsCsiNfs `json:"nfs,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsCsi) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsCsi
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsCsi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NFS CSI driver settings
type K8SClusterNewParamsCsiNfs struct {
	// Enable or disable VAST NFS integration. The default value is `false`. When set
	// to `true`, a dedicated StorageClass will be created in the cluster for each VAST
	// NFS file share defined in the cloud. All file shares created prior to cluster
	// creation will be available immediately, while those created afterward may take a
	// few minutes for to appear.
	VastEnabled param.Opt[bool] `json:"vast_enabled,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsCsiNfs) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsCsiNfs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsCsiNfs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS Protection profile
//
// The property Enabled is required.
type K8SClusterNewParamsDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile template ID
	ProfileTemplate param.Opt[int64] `json:"profile_template,omitzero"`
	// DDoS profile template name
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// DDoS profile parameters
	Fields []K8SClusterNewParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type K8SClusterNewParamsDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Basic value. Only one of 'value' or 'field_value' must be specified
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type K8SClusterNewParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to which the logs will be written
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r K8SClusterNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterNewParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Authentication settings
	Authentication K8SClusterUpdateParamsAuthentication `json:"authentication,omitzero"`
	// Cluster autoscaler configuration.
	//
	// It allows you to override the default cluster-autoscaler parameters provided by
	// the platform with your preferred values.
	//
	// Supported parameters (in alphabetical order):
	//
	//   - balance-similar-node-groups (boolean: true/false) - Detect similar node groups
	//     and balance the number of nodes between them.
	//   - expander (string: random, most-pods, least-waste, price, priority, grpc) -
	//     Type of node group expander to be used in scale up. Specifying multiple values
	//     separated by commas will call the expanders in succession until there is only
	//     one option remaining.
	//   - expendable-pods-priority-cutoff (float) - Pods with priority below cutoff will
	//     be expendable. They can be killed without any consideration during scale down
	//     and they don't cause scale up. Pods with null priority (PodPriority disabled)
	//     are non expendable.
	//   - ignore-daemonsets-utilization (boolean: true/false) - Should CA ignore
	//     DaemonSet pods when calculating resource utilization for scaling down.
	//   - max-empty-bulk-delete (integer) - Maximum number of empty nodes that can be
	//     deleted at the same time.
	//   - max-graceful-termination-sec (integer) - Maximum number of seconds CA waits
	//     for pod termination when trying to scale down a node.
	//   - max-node-provision-time (duration: e.g., '15m') - The default maximum time CA
	//     waits for node to be provisioned - the value can be overridden per node group.
	//   - max-total-unready-percentage (float) - Maximum percentage of unready nodes in
	//     the cluster. After this is exceeded, CA halts operations.
	//   - new-pod-scale-up-delay (duration: e.g., '10s') - Pods less than this old will
	//     not be considered for scale-up. Can be increased for individual pods through
	//     annotation.
	//   - ok-total-unready-count (integer) - Number of allowed unready nodes,
	//     irrespective of max-total-unready-percentage.
	//   - scale-down-delay-after-add (duration: e.g., '10m') - How long after scale up
	//     that scale down evaluation resumes.
	//   - scale-down-delay-after-delete (duration: e.g., '10s') - How long after node
	//     deletion that scale down evaluation resumes.
	//   - scale-down-delay-after-failure (duration: e.g., '3m') - How long after scale
	//     down failure that scale down evaluation resumes.
	//   - scale-down-enabled (boolean: true/false) - Should CA scale down the cluster.
	//   - scale-down-unneeded-time (duration: e.g., '10m') - How long a node should be
	//     unneeded before it is eligible for scale down.
	//   - scale-down-unready-time (duration: e.g., '20m') - How long an unready node
	//     should be unneeded before it is eligible for scale down.
	//   - scale-down-utilization-threshold (float) - The maximum value between the sum
	//     of cpu requests and sum of memory requests of all pods running on the node
	//     divided by node's corresponding allocatable resource, below which a node can
	//     be considered for scale down.
	//   - scan-interval (duration: e.g., '10s') - How often cluster is reevaluated for
	//     scale up or down.
	//   - skip-nodes-with-custom-controller-pods (boolean: true/false) - If true cluster
	//     autoscaler will never delete nodes with pods owned by custom controllers.
	//   - skip-nodes-with-local-storage (boolean: true/false) - If true cluster
	//     autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir
	//     or HostPath.
	//   - skip-nodes-with-system-pods (boolean: true/false) - If true cluster autoscaler
	//     will never delete nodes with pods from kube-system (except for DaemonSet or
	//     mirror pods).
	AutoscalerConfig map[string]string `json:"autoscaler_config,omitzero"`
	// Cluster CNI settings
	Cni K8SClusterUpdateParamsCni `json:"cni,omitzero"`
	// Advanced DDoS Protection profile
	DDOSProfile K8SClusterUpdateParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// Logging configuration
	Logging K8SClusterUpdateParamsLogging `json:"logging,omitzero"`
	// Cluster add-ons configuration
	AddOns K8SClusterUpdateParamsAddOns `json:"add_ons,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster add-ons configuration
type K8SClusterUpdateParamsAddOns struct {
	// Slurm add-on configuration
	Slurm K8SClusterUpdateParamsAddOnsSlurmUnion `json:"slurm,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsAddOns) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsAddOns
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsAddOns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type K8SClusterUpdateParamsAddOnsSlurmUnion struct {
	OfK8SClusterSlurmAddonEnableV2Serializer  *K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer  `json:",omitzero,inline"`
	OfK8SClusterSlurmAddonDisableV2Serializer *K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer `json:",omitzero,inline"`
	paramUnion
}

func (u K8SClusterUpdateParamsAddOnsSlurmUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfK8SClusterSlurmAddonEnableV2Serializer, u.OfK8SClusterSlurmAddonDisableV2Serializer)
}
func (u *K8SClusterUpdateParamsAddOnsSlurmUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *K8SClusterUpdateParamsAddOnsSlurmUnion) asAny() any {
	if !param.IsOmitted(u.OfK8SClusterSlurmAddonEnableV2Serializer) {
		return u.OfK8SClusterSlurmAddonEnableV2Serializer
	} else if !param.IsOmitted(u.OfK8SClusterSlurmAddonDisableV2Serializer) {
		return u.OfK8SClusterSlurmAddonDisableV2Serializer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u K8SClusterUpdateParamsAddOnsSlurmUnion) GetFileShareID() *string {
	if vt := u.OfK8SClusterSlurmAddonEnableV2Serializer; vt != nil {
		return &vt.FileShareID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u K8SClusterUpdateParamsAddOnsSlurmUnion) GetSSHKeyIDs() []string {
	if vt := u.OfK8SClusterSlurmAddonEnableV2Serializer; vt != nil {
		return vt.SSHKeyIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u K8SClusterUpdateParamsAddOnsSlurmUnion) GetWorkerCount() *int64 {
	if vt := u.OfK8SClusterSlurmAddonEnableV2Serializer; vt != nil {
		return &vt.WorkerCount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u K8SClusterUpdateParamsAddOnsSlurmUnion) GetEnabled() *bool {
	if vt := u.OfK8SClusterSlurmAddonEnableV2Serializer; vt != nil {
		return (*bool)(&vt.Enabled)
	} else if vt := u.OfK8SClusterSlurmAddonDisableV2Serializer; vt != nil {
		return (*bool)(&vt.Enabled)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[K8SClusterUpdateParamsAddOnsSlurmUnion](
		"enabled",
		apijson.Discriminator[K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer](true),
		apijson.Discriminator[K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer](false),
	)
}

// The properties Enabled, FileShareID, SSHKeyIDs, WorkerCount are required.
type K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer struct {
	// ID of a VAST file share to be used as Slurm storage.
	//
	// The Slurm add-on will create separate Persistent Volume Claims for different
	// purposes (controller spool, worker spool, jail) on that file share.
	//
	// The file share must have `root_squash` disabled, while `path_length` and
	// `allowed_characters` settings must be set to `NPL`.
	FileShareID string `json:"file_share_id,required" format:"uuid4"`
	// IDs of SSH keys to authorize for SSH connection to Slurm login nodes.
	SSHKeyIDs []string `json:"ssh_key_ids,omitzero,required" format:"uuid4"`
	// Size of the worker pool, i.e. the number of Slurm worker nodes.
	//
	// Each Slurm worker node will be backed by a Pod scheduled on one of cluster's GPU
	// nodes.
	//
	// Note: Downscaling (reducing worker count) is not supported.
	WorkerCount int64 `json:"worker_count,required"`
	// The Slurm add-on will be enabled in the cluster.
	//
	// This add-on is only supported in clusters running Kubernetes v1.31 and v1.32
	// with at least 1 GPU cluster pool and VAST NFS support enabled.
	//
	// This field can be elided, and will marshal its zero value as true.
	Enabled bool `json:"enabled,omitzero,required"`
	paramObj
}

func (r K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonEnableV2Serializer](
		"enabled", true,
	)
}

func NewK8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer() K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer {
	return K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer{
		Enabled: false,
	}
}

// This struct has a constant value, construct it with
// [NewK8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer].
type K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer struct {
	// The Slurm add-on will be disabled in the cluster.
	Enabled bool `json:"enabled,omitzero,required"`
	paramObj
}

func (r K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsAddOnsSlurmK8SClusterSlurmAddonDisableV2Serializer](
		"enabled", false,
	)
}

// Authentication settings
type K8SClusterUpdateParamsAuthentication struct {
	// OIDC authentication settings
	Oidc K8SClusterUpdateParamsAuthenticationOidc `json:"oidc,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsAuthentication) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsAuthentication
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8SClusterUpdateParamsAuthenticationOidc struct {
	// Client ID
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// JWT claim to use as the user's group
	GroupsClaim param.Opt[string] `json:"groups_claim,omitzero"`
	// Prefix prepended to group claims
	GroupsPrefix param.Opt[string] `json:"groups_prefix,omitzero"`
	// Issuer URL
	IssuerURL param.Opt[string] `json:"issuer_url,omitzero"`
	// JWT claim to use as the user name
	UsernameClaim param.Opt[string] `json:"username_claim,omitzero"`
	// Prefix prepended to username claims to prevent clashes
	UsernamePrefix param.Opt[string] `json:"username_prefix,omitzero"`
	// Key-value pairs that describe required claims in the token
	RequiredClaims map[string]string `json:"required_claims,omitzero"`
	// Accepted signing algorithms
	//
	// Any of "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "RS256", "RS384",
	// "RS512".
	SigningAlgs []string `json:"signing_algs,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsAuthenticationOidc) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsAuthenticationOidc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8SClusterUpdateParamsCni struct {
	// Cilium settings
	Cilium K8SClusterUpdateParamsCniCilium `json:"cilium,omitzero"`
	// CNI provider
	//
	// Any of "calico", "cilium".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsCni) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsCni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsCni](
		"provider", "calico", "cilium",
	)
}

// Cilium settings
type K8SClusterUpdateParamsCniCilium struct {
	// Wireguard encryption
	Encryption param.Opt[bool] `json:"encryption,omitzero"`
	// Hubble Relay
	HubbleRelay param.Opt[bool] `json:"hubble_relay,omitzero"`
	// Hubble UI
	HubbleUi param.Opt[bool] `json:"hubble_ui,omitzero"`
	// LoadBalancer acceleration
	LbAcceleration param.Opt[bool] `json:"lb_acceleration,omitzero"`
	// Mask size for IPv4
	MaskSize param.Opt[int64] `json:"mask_size,omitzero"`
	// Mask size for IPv6
	MaskSizeV6 param.Opt[int64] `json:"mask_size_v6,omitzero"`
	// LoadBalancer mode
	//
	// Any of "dsr", "hybrid", "snat".
	LbMode string `json:"lb_mode,omitzero"`
	// Routing mode
	//
	// Any of "native", "tunnel".
	RoutingMode string `json:"routing_mode,omitzero"`
	// CNI provider
	//
	// Any of "", "geneve", "vxlan".
	Tunnel string `json:"tunnel,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsCniCilium) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsCniCilium
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsCniCilium](
		"lb_mode", "dsr", "hybrid", "snat",
	)
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsCniCilium](
		"routing_mode", "native", "tunnel",
	)
	apijson.RegisterFieldValidator[K8SClusterUpdateParamsCniCilium](
		"tunnel", "", "geneve", "vxlan",
	)
}

// Advanced DDoS Protection profile
//
// The property Enabled is required.
type K8SClusterUpdateParamsDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile template ID
	ProfileTemplate param.Opt[int64] `json:"profile_template,omitzero"`
	// DDoS profile template name
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// DDoS profile parameters
	Fields []K8SClusterUpdateParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type K8SClusterUpdateParamsDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Basic value. Only one of 'value' or 'field_value' must be specified
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type K8SClusterUpdateParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to which the logs will be written
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r K8SClusterUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpdateParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Comma separated list of volume IDs to be deleted with the cluster
	Volumes param.Opt[string] `query:"volumes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [K8SClusterDeleteParams]'s query parameters as `url.Values`.
func (r K8SClusterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type K8SClusterGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterGetCertificateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterGetKubeconfigParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterListVersionsForUpgradeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterUpgradeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Target k8s cluster version
	Version string `json:"version,required"`
	paramObj
}

func (r K8SClusterUpgradeParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterUpgradeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterUpgradeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
