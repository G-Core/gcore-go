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

// K8ClusterService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8ClusterService] method instead.
type K8ClusterService struct {
	Options []option.RequestOption
	Nodes   K8ClusterNodeService
	Pools   K8ClusterPoolService
}

// NewK8ClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewK8ClusterService(opts ...option.RequestOption) (r K8ClusterService) {
	r = K8ClusterService{}
	r.Options = opts
	r.Nodes = NewK8ClusterNodeService(opts...)
	r.Pools = NewK8ClusterPoolService(opts...)
	return
}

// Create k8s cluster
func (r *K8ClusterService) New(ctx context.Context, params K8ClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// Update k8s cluster
func (r *K8ClusterService) Update(ctx context.Context, clusterName string, params K8ClusterUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// List k8s clusters
func (r *K8ClusterService) List(ctx context.Context, query K8ClusterListParams, opts ...option.RequestOption) (res *K8sClusterList, err error) {
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
func (r *K8ClusterService) Delete(ctx context.Context, clusterName string, params K8ClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// Get k8s cluster
func (r *K8ClusterService) Get(ctx context.Context, clusterName string, query K8ClusterGetParams, opts ...option.RequestOption) (res *K8sCluster, err error) {
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
func (r *K8ClusterService) GetCertificate(ctx context.Context, clusterName string, query K8ClusterGetCertificateParams, opts ...option.RequestOption) (res *K8sClusterCertificate, err error) {
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
func (r *K8ClusterService) GetKubeconfig(ctx context.Context, clusterName string, query K8ClusterGetKubeconfigParams, opts ...option.RequestOption) (res *K8sClusterKubeconfig, err error) {
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
func (r *K8ClusterService) ListVersionsForUpgrade(ctx context.Context, clusterName string, query K8ClusterListVersionsForUpgradeParams, opts ...option.RequestOption) (res *K8sClusterVersionList, err error) {
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
func (r *K8ClusterService) Upgrade(ctx context.Context, clusterName string, params K8ClusterUpgradeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

type K8sCluster struct {
	// Cluster pool uuid
	ID string `json:"id,required"`
	// Function creation date
	CreatedAt string `json:"created_at,required"`
	// Cluster CSI settings
	Csi K8sClusterCsi `json:"csi,required"`
	// Cluster is public
	IsPublic bool `json:"is_public,required"`
	// Keypair
	Keypair string `json:"keypair,required"`
	// Logging configuration
	Logging Logging `json:"logging,required"`
	// Name
	Name string `json:"name,required"`
	// pools
	Pools []K8sClusterPool `json:"pools,required"`
	// Status
	//
	// Any of "Deleting", "Provisioned", "Provisioning".
	Status K8sClusterStatus `json:"status,required"`
	// K8s version
	Version string `json:"version,required"`
	// Cluster authentication settings
	Authentication K8sClusterAuthentication `json:"authentication,nullable"`
	// Cluster autoscaler configuration.
	//
	// It contains overrides to the default cluster-autoscaler parameters provided by
	// the platform.
	AutoscalerConfig map[string]string `json:"autoscaler_config,nullable"`
	// Cluster CNI settings
	Cni K8sClusterCni `json:"cni,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Advanced DDoS Protection profile
	DDOSProfile K8sClusterDDOSProfile `json:"ddos_profile,nullable"`
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
func (r K8sCluster) RawJSON() string { return r.JSON.raw }
func (r *K8sCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CSI settings
type K8sClusterCsi struct {
	// NFS settings
	Nfs K8sClusterCsiNfs `json:"nfs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Nfs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8sClusterCsi) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterCsi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NFS settings
type K8sClusterCsiNfs struct {
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
func (r K8sClusterCsiNfs) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterCsiNfs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status
type K8sClusterStatus string

const (
	K8sClusterStatusDeleting     K8sClusterStatus = "Deleting"
	K8sClusterStatusProvisioned  K8sClusterStatus = "Provisioned"
	K8sClusterStatusProvisioning K8sClusterStatus = "Provisioning"
)

// Cluster authentication settings
type K8sClusterAuthentication struct {
	// Kubeconfig creation date
	KubeconfigCreatedAt time.Time `json:"kubeconfig_created_at,nullable" format:"date-time"`
	// Kubeconfig expiration date
	KubeconfigExpiresAt time.Time `json:"kubeconfig_expires_at,nullable" format:"date-time"`
	// OIDC authentication settings
	Oidc K8sClusterAuthenticationOidc `json:"oidc,nullable"`
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
func (r K8sClusterAuthentication) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8sClusterAuthenticationOidc struct {
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
func (r K8sClusterAuthenticationOidc) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8sClusterCni struct {
	// Cilium settings
	Cilium K8sClusterCniCilium `json:"cilium,nullable"`
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
func (r K8sClusterCni) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cilium settings
type K8sClusterCniCilium struct {
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
func (r K8sClusterCniCilium) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS Protection profile
type K8sClusterDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile parameters
	Fields []K8sClusterDDOSProfileField `json:"fields,nullable"`
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
func (r K8sClusterDDOSProfile) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8sClusterDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified
	FieldValue any `json:"field_value"`
	// Basic value. Only one of 'value' or '`field_value`' must be specified
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
func (r K8sClusterDDOSProfileField) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8sClusterCertificate struct {
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
func (r K8sClusterCertificate) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8sClusterKubeconfig struct {
	// Cluster kubeconfig
	Config string `json:"config,required"`
	// Kubeconfig creation date
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Kubeconfig expiration date
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8sClusterKubeconfig) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterKubeconfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8sClusterList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []K8sCluster `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8sClusterList) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8ClusterNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The keypair of the cluster
	Keypair string `json:"keypair,required"`
	// The name of the cluster
	Name string `json:"name,required"`
	// The pools of the cluster
	Pools []K8ClusterNewParamsPool `json:"pools,omitzero,required"`
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
	Authentication K8ClusterNewParamsAuthentication `json:"authentication,omitzero"`
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
	Cni K8ClusterNewParamsCni `json:"cni,omitzero"`
	// Advanced DDoS Protection profile
	DDOSProfile K8ClusterNewParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// Logging configuration
	Logging K8ClusterNewParamsLogging `json:"logging,omitzero"`
	// Container Storage Interface (CSI) driver settings
	Csi K8ClusterNewParamsCsi `json:"csi,omitzero"`
	paramObj
}

func (r K8ClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FlavorID, MinNodeCount, Name are required.
type K8ClusterNewParamsPool struct {
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

func (r K8ClusterNewParamsPool) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8ClusterNewParamsPool](
		"boot_volume_type", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
	apijson.RegisterFieldValidator[K8ClusterNewParamsPool](
		"servergroup_policy", "affinity", "anti-affinity", "soft-anti-affinity",
	)
}

// Authentication settings
type K8ClusterNewParamsAuthentication struct {
	// OIDC authentication settings
	Oidc K8ClusterNewParamsAuthenticationOidc `json:"oidc,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsAuthentication) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsAuthentication
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8ClusterNewParamsAuthenticationOidc struct {
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

func (r K8ClusterNewParamsAuthenticationOidc) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsAuthenticationOidc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8ClusterNewParamsCni struct {
	// Cilium settings
	Cilium K8ClusterNewParamsCniCilium `json:"cilium,omitzero"`
	// CNI provider
	//
	// Any of "calico", "cilium".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsCni) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsCni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8ClusterNewParamsCni](
		"provider", "calico", "cilium",
	)
}

// Cilium settings
type K8ClusterNewParamsCniCilium struct {
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

func (r K8ClusterNewParamsCniCilium) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsCniCilium
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8ClusterNewParamsCniCilium](
		"lb_mode", "dsr", "hybrid", "snat",
	)
	apijson.RegisterFieldValidator[K8ClusterNewParamsCniCilium](
		"routing_mode", "native", "tunnel",
	)
	apijson.RegisterFieldValidator[K8ClusterNewParamsCniCilium](
		"tunnel", "", "geneve", "vxlan",
	)
}

// Container Storage Interface (CSI) driver settings
type K8ClusterNewParamsCsi struct {
	// NFS CSI driver settings
	Nfs K8ClusterNewParamsCsiNfs `json:"nfs,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsCsi) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsCsi
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsCsi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NFS CSI driver settings
type K8ClusterNewParamsCsiNfs struct {
	// Enable or disable VAST NFS integration. The default value is `false`. When set
	// to `true`, a dedicated StorageClass will be created in the cluster for each VAST
	// NFS file share defined in the cloud. All file shares created prior to cluster
	// creation will be available immediately, while those created afterward may take a
	// few minutes for to appear.
	VastEnabled param.Opt[bool] `json:"vast_enabled,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsCsiNfs) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsCsiNfs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsCsiNfs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS Protection profile
//
// The property Enabled is required.
type K8ClusterNewParamsDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile template ID
	ProfileTemplate param.Opt[int64] `json:"profile_template,omitzero"`
	// DDoS profile template name
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// DDoS profile parameters
	Fields []K8ClusterNewParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type K8ClusterNewParamsDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Basic value. Only one of 'value' or '`field_value`' must be specified
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r K8ClusterNewParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type K8ClusterNewParamsLogging struct {
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

func (r K8ClusterNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterNewParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8ClusterUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Authentication settings
	Authentication K8ClusterUpdateParamsAuthentication `json:"authentication,omitzero"`
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
	Cni K8ClusterUpdateParamsCni `json:"cni,omitzero"`
	// Advanced DDoS Protection profile
	DDOSProfile K8ClusterUpdateParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// Logging configuration
	Logging K8ClusterUpdateParamsLogging `json:"logging,omitzero"`
	paramObj
}

func (r K8ClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authentication settings
type K8ClusterUpdateParamsAuthentication struct {
	// OIDC authentication settings
	Oidc K8ClusterUpdateParamsAuthenticationOidc `json:"oidc,omitzero"`
	paramObj
}

func (r K8ClusterUpdateParamsAuthentication) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsAuthentication
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OIDC authentication settings
type K8ClusterUpdateParamsAuthenticationOidc struct {
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

func (r K8ClusterUpdateParamsAuthenticationOidc) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsAuthenticationOidc
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsAuthenticationOidc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster CNI settings
type K8ClusterUpdateParamsCni struct {
	// Cilium settings
	Cilium K8ClusterUpdateParamsCniCilium `json:"cilium,omitzero"`
	// CNI provider
	//
	// Any of "calico", "cilium".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r K8ClusterUpdateParamsCni) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsCni
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsCni) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8ClusterUpdateParamsCni](
		"provider", "calico", "cilium",
	)
}

// Cilium settings
type K8ClusterUpdateParamsCniCilium struct {
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

func (r K8ClusterUpdateParamsCniCilium) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsCniCilium
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsCniCilium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[K8ClusterUpdateParamsCniCilium](
		"lb_mode", "dsr", "hybrid", "snat",
	)
	apijson.RegisterFieldValidator[K8ClusterUpdateParamsCniCilium](
		"routing_mode", "native", "tunnel",
	)
	apijson.RegisterFieldValidator[K8ClusterUpdateParamsCniCilium](
		"tunnel", "", "geneve", "vxlan",
	)
}

// Advanced DDoS Protection profile
//
// The property Enabled is required.
type K8ClusterUpdateParamsDDOSProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile template ID
	ProfileTemplate param.Opt[int64] `json:"profile_template,omitzero"`
	// DDoS profile template name
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// DDoS profile parameters
	Fields []K8ClusterUpdateParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r K8ClusterUpdateParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type K8ClusterUpdateParamsDDOSProfileField struct {
	BaseField int64 `json:"base_field,required"`
	// Basic value. Only one of 'value' or '`field_value`' must be specified
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r K8ClusterUpdateParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type K8ClusterUpdateParamsLogging struct {
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

func (r K8ClusterUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpdateParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8ClusterListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Comma separated list of volume IDs to be deleted with the cluster
	Volumes param.Opt[string] `query:"volumes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [K8ClusterDeleteParams]'s query parameters as `url.Values`.
func (r K8ClusterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type K8ClusterGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterGetCertificateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterGetKubeconfigParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterListVersionsForUpgradeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterUpgradeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Target k8s cluster version
	Version string `json:"version,required"`
	paramObj
}

func (r K8ClusterUpgradeParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterUpgradeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterUpgradeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
