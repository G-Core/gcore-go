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

// CloudV2K8ClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2K8ClusterService] method instead.
type CloudV2K8ClusterService struct {
	Options []option.RequestOption
	Pools   *CloudV2K8ClusterPoolService
}

// NewCloudV2K8ClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2K8ClusterService(opts ...option.RequestOption) (r *CloudV2K8ClusterService) {
	r = &CloudV2K8ClusterService{}
	r.Options = opts
	r.Pools = NewCloudV2K8ClusterPoolService(opts...)
	return
}

// Create k8s cluster
func (r *CloudV2K8ClusterService) New(ctx context.Context, projectID int64, regionID int64, body CloudV2K8ClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get k8s cluster
func (r *CloudV2K8ClusterService) Get(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update k8s cluster
func (r *CloudV2K8ClusterService) Update(ctx context.Context, projectID int64, regionID int64, clusterName string, body CloudV2K8ClusterUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List k8s clusters
func (r *CloudV2K8ClusterService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV2K8ClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete k8s cluster
func (r *CloudV2K8ClusterService) Delete(ctx context.Context, projectID int64, regionID int64, clusterName string, body CloudV2K8ClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified cluster will be calculated
func (r *CloudV2K8ClusterService) CheckLimits(ctx context.Context, projectID int64, regionID int64, body CloudV2K8ClusterCheckLimitsParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get k8s cluster CA certificate
func (r *CloudV2K8ClusterService) GetCertificates(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *CloudV2K8ClusterGetCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/certificates", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get k8s cluster kubeconfig
func (r *CloudV2K8ClusterService) GetConfig(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *CloudV2K8ClusterGetConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/config", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List available k8s cluster versions for upgrade
func (r *CloudV2K8ClusterService) GetUpgradeVersions(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *VersionList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/upgrade_versions", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List instances in k8s cluster
func (r *CloudV2K8ClusterService) ListInstances(ctx context.Context, projectID int64, regionID int64, clusterName string, query CloudV2K8ClusterListInstancesParams, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/instances", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upgrade k8s cluster
func (r *CloudV2K8ClusterService) Upgrade(ctx context.Context, projectID int64, regionID int64, clusterName string, body CloudV2K8ClusterUpgradeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/upgrade", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthenticationCreateParam struct {
	// OIDC authentication settings
	Oidc param.Field[OidcParam] `json:"oidc"`
}

func (r AuthenticationCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cilium struct {
	// Wireguard encryption
	Encryption bool `json:"encryption"`
	// Hubble Relay
	HubbleRelay bool `json:"hubble_relay"`
	// Hubble UI
	HubbleUi bool `json:"hubble_ui"`
	// LoadBalancer acceleration
	LbAcceleration bool `json:"lb_acceleration"`
	// LoadBalancer mode
	LbMode CiliumLbMode `json:"lb_mode"`
	// Mask size for IPv4
	MaskSize int64 `json:"mask_size"`
	// Mask size for IPv6
	MaskSizeV6 int64 `json:"mask_size_v6"`
	// Routing mode
	RoutingMode CiliumRoutingMode `json:"routing_mode"`
	// CNI provider
	Tunnel CiliumTunnel `json:"tunnel"`
	JSON   ciliumJSON   `json:"-"`
}

// ciliumJSON contains the JSON metadata for the struct [Cilium]
type ciliumJSON struct {
	Encryption     apijson.Field
	HubbleRelay    apijson.Field
	HubbleUi       apijson.Field
	LbAcceleration apijson.Field
	LbMode         apijson.Field
	MaskSize       apijson.Field
	MaskSizeV6     apijson.Field
	RoutingMode    apijson.Field
	Tunnel         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Cilium) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ciliumJSON) RawJSON() string {
	return r.raw
}

// LoadBalancer mode
type CiliumLbMode string

const (
	CiliumLbModeDsr    CiliumLbMode = "dsr"
	CiliumLbModeHybrid CiliumLbMode = "hybrid"
	CiliumLbModeSnat   CiliumLbMode = "snat"
)

func (r CiliumLbMode) IsKnown() bool {
	switch r {
	case CiliumLbModeDsr, CiliumLbModeHybrid, CiliumLbModeSnat:
		return true
	}
	return false
}

// Routing mode
type CiliumRoutingMode string

const (
	CiliumRoutingModeNative CiliumRoutingMode = "native"
	CiliumRoutingModeTunnel CiliumRoutingMode = "tunnel"
)

func (r CiliumRoutingMode) IsKnown() bool {
	switch r {
	case CiliumRoutingModeNative, CiliumRoutingModeTunnel:
		return true
	}
	return false
}

// CNI provider
type CiliumTunnel string

const (
	CiliumTunnelEmpty  CiliumTunnel = ""
	CiliumTunnelGeneve CiliumTunnel = "geneve"
	CiliumTunnelVxlan  CiliumTunnel = "vxlan"
)

func (r CiliumTunnel) IsKnown() bool {
	switch r {
	case CiliumTunnelEmpty, CiliumTunnelGeneve, CiliumTunnelVxlan:
		return true
	}
	return false
}

type CiliumParam struct {
	// Wireguard encryption
	Encryption param.Field[bool] `json:"encryption"`
	// Hubble Relay
	HubbleRelay param.Field[bool] `json:"hubble_relay"`
	// Hubble UI
	HubbleUi param.Field[bool] `json:"hubble_ui"`
	// LoadBalancer acceleration
	LbAcceleration param.Field[bool] `json:"lb_acceleration"`
	// LoadBalancer mode
	LbMode param.Field[CiliumLbMode] `json:"lb_mode"`
	// Mask size for IPv4
	MaskSize param.Field[int64] `json:"mask_size"`
	// Mask size for IPv6
	MaskSizeV6 param.Field[int64] `json:"mask_size_v6"`
	// Routing mode
	RoutingMode param.Field[CiliumRoutingMode] `json:"routing_mode"`
	// CNI provider
	Tunnel param.Field[CiliumTunnel] `json:"tunnel"`
}

func (r CiliumParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cluster struct {
	// Cluster pool uuid
	ID string `json:"id,required"`
	// Function creation date
	CreatedAt string `json:"created_at,required"`
	// Cluster is public
	IsPublic bool `json:"is_public,required"`
	// Keypair
	Keypair string `json:"keypair,required"`
	// Name
	Name string `json:"name,required"`
	// pools
	Pools []Pool `json:"pools,required"`
	// Status
	Status ClusterStatus `json:"status,required"`
	// K8s version
	Version string `json:"version,required"`
	// Cluster authentication settings
	Authentication ClusterAuthentication `json:"authentication,nullable"`
	// Cluster autoscaler configuration params
	AutoscalerConfig map[string]string `json:"autoscaler_config,nullable"`
	// Cluster CNI settings
	Cni ClusterCni `json:"cni,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Advanced DDoS Protection profile
	DdosProfile ClusterDdosProfile `json:"ddos_profile,nullable"`
	// Fixed network id
	FixedNetwork string `json:"fixed_network,nullable"`
	// Fixed subnet id
	FixedSubnet string `json:"fixed_subnet,nullable"`
	// Enable public v6 address
	IsIpv6 bool `json:"is_ipv6"`
	// Logging configuration
	Logging ClusterLogging `json:"logging,nullable"`
	// The IP pool for the pods
	PodsIPPool string `json:"pods_ip_pool,nullable"`
	// The IPv6 pool for the pods
	PodsIpv6Pool string `json:"pods_ipv6_pool,nullable"`
	// The IP pool for the services
	ServicesIPPool string `json:"services_ip_pool,nullable"`
	// The IPv6 pool for the services
	ServicesIpv6Pool string `json:"services_ipv6_pool,nullable"`
	// Active task
	TaskID string      `json:"task_id,nullable"`
	JSON   clusterJSON `json:"-"`
}

// clusterJSON contains the JSON metadata for the struct [Cluster]
type clusterJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	IsPublic         apijson.Field
	Keypair          apijson.Field
	Name             apijson.Field
	Pools            apijson.Field
	Status           apijson.Field
	Version          apijson.Field
	Authentication   apijson.Field
	AutoscalerConfig apijson.Field
	Cni              apijson.Field
	CreatorTaskID    apijson.Field
	DdosProfile      apijson.Field
	FixedNetwork     apijson.Field
	FixedSubnet      apijson.Field
	IsIpv6           apijson.Field
	Logging          apijson.Field
	PodsIPPool       apijson.Field
	PodsIpv6Pool     apijson.Field
	ServicesIPPool   apijson.Field
	ServicesIpv6Pool apijson.Field
	TaskID           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Cluster) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterJSON) RawJSON() string {
	return r.raw
}

// Status
type ClusterStatus string

const (
	ClusterStatusDeleting     ClusterStatus = "Deleting"
	ClusterStatusProvisioned  ClusterStatus = "Provisioned"
	ClusterStatusProvisioning ClusterStatus = "Provisioning"
)

func (r ClusterStatus) IsKnown() bool {
	switch r {
	case ClusterStatusDeleting, ClusterStatusProvisioned, ClusterStatusProvisioning:
		return true
	}
	return false
}

// Cluster authentication settings
type ClusterAuthentication struct {
	// Kubeconfig creation date
	KubeconfigCreatedAt time.Time `json:"kubeconfig_created_at,nullable" format:"date-time"`
	// Kubeconfig expiration date
	KubeconfigExpiresAt time.Time `json:"kubeconfig_expires_at,nullable" format:"date-time"`
	// OIDC authentication settings
	Oidc Oidc                      `json:"oidc,nullable"`
	JSON clusterAuthenticationJSON `json:"-"`
}

// clusterAuthenticationJSON contains the JSON metadata for the struct
// [ClusterAuthentication]
type clusterAuthenticationJSON struct {
	KubeconfigCreatedAt apijson.Field
	KubeconfigExpiresAt apijson.Field
	Oidc                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ClusterAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterAuthenticationJSON) RawJSON() string {
	return r.raw
}

// Cluster CNI settings
type ClusterCni struct {
	// Cilium settings
	Cilium Cilium `json:"cilium,nullable"`
	// CNI provider
	Provider Cni            `json:"provider"`
	JSON     clusterCniJSON `json:"-"`
}

// clusterCniJSON contains the JSON metadata for the struct [ClusterCni]
type clusterCniJSON struct {
	Cilium      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClusterCni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterCniJSON) RawJSON() string {
	return r.raw
}

// Advanced DDoS Protection profile
type ClusterDdosProfile struct {
	// Enable advanced DDoS protection
	Enabled bool `json:"enabled,required"`
	// DDoS profile parameters
	Fields []DdosProfile `json:"fields,nullable"`
	// DDoS profile template ID
	ProfileTemplate int64 `json:"profile_template,nullable"`
	// DDoS profile template name
	ProfileTemplateName string                 `json:"profile_template_name,nullable"`
	JSON                clusterDdosProfileJSON `json:"-"`
}

// clusterDdosProfileJSON contains the JSON metadata for the struct
// [ClusterDdosProfile]
type clusterDdosProfileJSON struct {
	Enabled             apijson.Field
	Fields              apijson.Field
	ProfileTemplate     apijson.Field
	ProfileTemplateName apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ClusterDdosProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterDdosProfileJSON) RawJSON() string {
	return r.raw
}

// Logging configuration
type ClusterLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID int64 `json:"destination_region_id,nullable"`
	// Enable/disable forwarding logs to LaaS
	Enabled bool `json:"enabled"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyPydantic `json:"retention_policy,nullable"`
	// The topic name to which the logs will be written
	TopicName string             `json:"topic_name,nullable"`
	JSON      clusterLoggingJSON `json:"-"`
}

// clusterLoggingJSON contains the JSON metadata for the struct [ClusterLogging]
type clusterLoggingJSON struct {
	DestinationRegionID apijson.Field
	Enabled             apijson.Field
	RetentionPolicy     apijson.Field
	TopicName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ClusterLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clusterLoggingJSON) RawJSON() string {
	return r.raw
}

type Cni string

const (
	CniCalico Cni = "calico"
	CniCilium Cni = "cilium"
)

func (r Cni) IsKnown() bool {
	switch r {
	case CniCalico, CniCilium:
		return true
	}
	return false
}

type CniCreateParam struct {
	// Cilium settings
	Cilium param.Field[CiliumParam] `json:"cilium"`
	// CNI provider
	Provider param.Field[Cni] `json:"provider"`
}

func (r CniCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DdosProfile struct {
	BaseField int64 `json:"base_field,required"`
	// Complex value. Only one of 'value' or 'field_value' must be specified
	FieldValue interface{} `json:"field_value"`
	// Basic value. Only one of 'value' or 'field_value' must be specified
	Value string          `json:"value,nullable"`
	JSON  ddosProfileJSON `json:"-"`
}

// ddosProfileJSON contains the JSON metadata for the struct [DdosProfile]
type ddosProfileJSON struct {
	BaseField   apijson.Field
	FieldValue  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DdosProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ddosProfileJSON) RawJSON() string {
	return r.raw
}

type DdosProfileParam struct {
	BaseField param.Field[int64] `json:"base_field,required"`
	// Complex value. Only one of 'value' or 'field_value' must be specified
	FieldValue param.Field[interface{}] `json:"field_value"`
	// Basic value. Only one of 'value' or 'field_value' must be specified
	Value param.Field[string] `json:"value"`
}

func (r DdosProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Oidc struct {
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
	SigningAlgs []OidcSigningAlg `json:"signing_algs,nullable"`
	// JWT claim to use as the user name
	UsernameClaim string `json:"username_claim,nullable"`
	// Prefix prepended to username claims to prevent clashes
	UsernamePrefix string   `json:"username_prefix,nullable"`
	JSON           oidcJSON `json:"-"`
}

// oidcJSON contains the JSON metadata for the struct [Oidc]
type oidcJSON struct {
	ClientID       apijson.Field
	GroupsClaim    apijson.Field
	GroupsPrefix   apijson.Field
	IssuerURL      apijson.Field
	RequiredClaims apijson.Field
	SigningAlgs    apijson.Field
	UsernameClaim  apijson.Field
	UsernamePrefix apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Oidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcJSON) RawJSON() string {
	return r.raw
}

type OidcSigningAlg string

const (
	OidcSigningAlgEs256 OidcSigningAlg = "ES256"
	OidcSigningAlgEs384 OidcSigningAlg = "ES384"
	OidcSigningAlgEs512 OidcSigningAlg = "ES512"
	OidcSigningAlgPs256 OidcSigningAlg = "PS256"
	OidcSigningAlgPs384 OidcSigningAlg = "PS384"
	OidcSigningAlgPs512 OidcSigningAlg = "PS512"
	OidcSigningAlgRs256 OidcSigningAlg = "RS256"
	OidcSigningAlgRs384 OidcSigningAlg = "RS384"
	OidcSigningAlgRs512 OidcSigningAlg = "RS512"
)

func (r OidcSigningAlg) IsKnown() bool {
	switch r {
	case OidcSigningAlgEs256, OidcSigningAlgEs384, OidcSigningAlgEs512, OidcSigningAlgPs256, OidcSigningAlgPs384, OidcSigningAlgPs512, OidcSigningAlgRs256, OidcSigningAlgRs384, OidcSigningAlgRs512:
		return true
	}
	return false
}

type OidcParam struct {
	// Client ID
	ClientID param.Field[string] `json:"client_id"`
	// JWT claim to use as the user's group
	GroupsClaim param.Field[string] `json:"groups_claim"`
	// Prefix prepended to group claims
	GroupsPrefix param.Field[string] `json:"groups_prefix"`
	// Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Key-value pairs that describe required claims in the token
	RequiredClaims param.Field[map[string]string] `json:"required_claims"`
	// Accepted signing algorithms
	SigningAlgs param.Field[[]OidcSigningAlg] `json:"signing_algs"`
	// JWT claim to use as the user name
	UsernameClaim param.Field[string] `json:"username_claim"`
	// Prefix prepended to username claims to prevent clashes
	UsernamePrefix param.Field[string] `json:"username_prefix"`
}

func (r OidcParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VersionList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []VersionListResult `json:"results,required"`
	JSON    versionListJSON     `json:"-"`
}

// versionListJSON contains the JSON metadata for the struct [VersionList]
type versionListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListJSON) RawJSON() string {
	return r.raw
}

type VersionListResult struct {
	// List of supported Kubernetes cluster versions
	Version string                `json:"version,required"`
	JSON    versionListResultJSON `json:"-"`
}

// versionListResultJSON contains the JSON metadata for the struct
// [VersionListResult]
type versionListResultJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResultJSON) RawJSON() string {
	return r.raw
}

type CloudV2K8ClusterListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Cluster                        `json:"results,required"`
	JSON    cloudV2K8ClusterListResponseJSON `json:"-"`
}

// cloudV2K8ClusterListResponseJSON contains the JSON metadata for the struct
// [CloudV2K8ClusterListResponse]
type cloudV2K8ClusterListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2K8ClusterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2K8ClusterListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2K8ClusterGetCertificatesResponse struct {
	// Cluster CA certificate
	Certificate string `json:"certificate,required"`
	// Cluster CA private key
	Key  string                                      `json:"key,required"`
	JSON cloudV2K8ClusterGetCertificatesResponseJSON `json:"-"`
}

// cloudV2K8ClusterGetCertificatesResponseJSON contains the JSON metadata for the
// struct [CloudV2K8ClusterGetCertificatesResponse]
type cloudV2K8ClusterGetCertificatesResponseJSON struct {
	Certificate apijson.Field
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2K8ClusterGetCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2K8ClusterGetCertificatesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2K8ClusterGetConfigResponse struct {
	// Cluster kubeconfig
	Config string `json:"config,required"`
	// Kubeconfig creation date
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Kubeconfig expiration date
	ExpiresAt time.Time                             `json:"expires_at,nullable" format:"date-time"`
	JSON      cloudV2K8ClusterGetConfigResponseJSON `json:"-"`
}

// cloudV2K8ClusterGetConfigResponseJSON contains the JSON metadata for the struct
// [CloudV2K8ClusterGetConfigResponse]
type cloudV2K8ClusterGetConfigResponseJSON struct {
	Config      apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2K8ClusterGetConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2K8ClusterGetConfigResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2K8ClusterNewParams struct {
	// The keypair of the cluster
	Keypair param.Field[string] `json:"keypair,required"`
	// The name of the cluster
	Name param.Field[string] `json:"name,required"`
	// The pools of the cluster
	Pools param.Field[[]PoolCreateParam] `json:"pools,required"`
	// The version of the k8s cluster
	Version param.Field[string] `json:"version,required"`
	// Authentication settings
	Authentication param.Field[AuthenticationCreateParam] `json:"authentication"`
	// Cluster autoscaler configuration params
	AutoscalerConfig param.Field[map[string]string] `json:"autoscaler_config"`
	// Cluster CNI settings
	Cni param.Field[CniCreateParam] `json:"cni"`
	// Advanced DDoS Protection profile
	DdosProfile param.Field[CloudV2K8ClusterNewParamsDdosProfile] `json:"ddos_profile"`
	// The network of the cluster
	FixedNetwork param.Field[string] `json:"fixed_network"`
	// The subnet of the cluster
	FixedSubnet param.Field[string] `json:"fixed_subnet"`
	// Enable public v6 address
	IsIpv6 param.Field[bool] `json:"is_ipv6"`
	// Logging configuration
	Logging param.Field[CloudV2K8ClusterNewParamsLogging] `json:"logging"`
	// The IP pool for the pods
	PodsIPPool param.Field[string] `json:"pods_ip_pool"`
	// The IPv6 pool for the pods
	PodsIpv6Pool param.Field[string] `json:"pods_ipv6_pool"`
	// The IP pool for the services
	ServicesIPPool param.Field[string] `json:"services_ip_pool"`
	// The IPv6 pool for the services
	ServicesIpv6Pool param.Field[string] `json:"services_ipv6_pool"`
}

func (r CloudV2K8ClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Advanced DDoS Protection profile
type CloudV2K8ClusterNewParamsDdosProfile struct {
	// Enable advanced DDoS protection
	Enabled param.Field[bool] `json:"enabled,required"`
	// DDoS profile parameters
	Fields param.Field[[]DdosProfileParam] `json:"fields"`
	// DDoS profile template ID
	ProfileTemplate param.Field[int64] `json:"profile_template"`
	// DDoS profile template name
	ProfileTemplateName param.Field[string] `json:"profile_template_name"`
}

func (r CloudV2K8ClusterNewParamsDdosProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging configuration
type CloudV2K8ClusterNewParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Field[bool] `json:"enabled"`
	// The logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to which the logs will be written
	TopicName param.Field[string] `json:"topic_name"`
}

func (r CloudV2K8ClusterNewParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterUpdateParams struct {
	// Authentication settings
	Authentication param.Field[AuthenticationCreateParam] `json:"authentication"`
	// Cluster autoscaler configuration params
	AutoscalerConfig param.Field[map[string]string] `json:"autoscaler_config"`
	// Cluster CNI settings
	Cni param.Field[CniCreateParam] `json:"cni"`
	// Advanced DDoS Protection profile
	DdosProfile param.Field[CloudV2K8ClusterUpdateParamsDdosProfile] `json:"ddos_profile"`
	// Logging configuration
	Logging param.Field[CloudV2K8ClusterUpdateParamsLogging] `json:"logging"`
}

func (r CloudV2K8ClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Advanced DDoS Protection profile
type CloudV2K8ClusterUpdateParamsDdosProfile struct {
	// Enable advanced DDoS protection
	Enabled param.Field[bool] `json:"enabled,required"`
	// DDoS profile parameters
	Fields param.Field[[]DdosProfileParam] `json:"fields"`
	// DDoS profile template ID
	ProfileTemplate param.Field[int64] `json:"profile_template"`
	// DDoS profile template name
	ProfileTemplateName param.Field[string] `json:"profile_template_name"`
}

func (r CloudV2K8ClusterUpdateParamsDdosProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging configuration
type CloudV2K8ClusterUpdateParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Field[bool] `json:"enabled"`
	// The logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to which the logs will be written
	TopicName param.Field[string] `json:"topic_name"`
}

func (r CloudV2K8ClusterUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterDeleteParams struct {
	// Comma separated list of volume IDs to be deleted with the cluster
	Volumes param.Field[string] `query:"volumes"`
}

// URLQuery serializes [CloudV2K8ClusterDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV2K8ClusterDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2K8ClusterCheckLimitsParams struct {
	// Logging configuration
	Logging param.Field[CloudV2K8ClusterCheckLimitsParamsLogging] `json:"logging"`
	// K8s pools to create
	Pools param.Field[[]CloudV2K8ClusterCheckLimitsParamsPool] `json:"pools"`
}

func (r CloudV2K8ClusterCheckLimitsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging configuration
type CloudV2K8ClusterCheckLimitsParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Field[bool] `json:"enabled"`
	// The logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to which the logs will be written
	TopicName param.Field[string] `json:"topic_name"`
}

func (r CloudV2K8ClusterCheckLimitsParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterCheckLimitsParamsPool struct {
	// Flavor ID
	FlavorID param.Field[string] `json:"flavor_id,required"`
	// Boot volume size
	BootVolumeSize param.Field[int64] `json:"boot_volume_size"`
	// Maximum node count
	MaxNodeCount param.Field[int64] `json:"max_node_count"`
	// Minimum node count
	MinNodeCount param.Field[int64] `json:"min_node_count"`
	// Name of the cluster pool
	Name param.Field[string] `json:"name"`
	// Maximum node count
	NodeCount param.Field[int64] `json:"node_count"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	ServergroupPolicy param.Field[ServerGroupPolicy] `json:"servergroup_policy"`
}

func (r CloudV2K8ClusterCheckLimitsParamsPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterListInstancesParams struct {
	// Include DDoS profile information if set to true. The default value is false.
	WithDdos param.Field[bool] `query:"with_ddos"`
}

// URLQuery serializes [CloudV2K8ClusterListInstancesParams]'s query parameters as
// `url.Values`.
func (r CloudV2K8ClusterListInstancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2K8ClusterUpgradeParams struct {
	// Target k8s cluster version
	Version param.Field[string] `json:"version,required"`
}

func (r CloudV2K8ClusterUpgradeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
