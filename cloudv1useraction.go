// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1UserActionService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1UserActionService] method instead.
type CloudV1UserActionService struct {
	Options []option.RequestOption
}

// NewCloudV1UserActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1UserActionService(opts ...option.RequestOption) (r *CloudV1UserActionService) {
	r = &CloudV1UserActionService{}
	r.Options = opts
	return
}

// Retrieve user action log for one client or a set of projects
func (r *CloudV1UserActionService) List(ctx context.Context, query CloudV1UserActionListParams, opts ...option.RequestOption) (res *CloudV1UserActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/user_actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get client subscriptions list
func (r *CloudV1UserActionService) GetSubscriptionsList(ctx context.Context, query CloudV1UserActionGetSubscriptionsListParams, opts ...option.RequestOption) (res *CloudV1UserActionGetSubscriptionsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/user_actions/subscriptions_list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Subscribe to the user action log. Subscription is created for the client_id that
// is taken from the JWT token. See how to get the token in the Account / Login
// section of this documentation
func (r *CloudV1UserActionService) Subscribe(ctx context.Context, body CloudV1UserActionSubscribeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/user_actions/subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Subscribe to the user action log over AMQP. Subscription is created for the
// client_id that is taken from the JWT token. See how to get the token in the
// Account / Login section of this documentation
func (r *CloudV1UserActionService) SubscribeAmqp(ctx context.Context, body CloudV1UserActionSubscribeAmqpParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/user_actions/subscribe_amqp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Unsubscribe from the user action log
func (r *CloudV1UserActionService) Unsubscribe(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/user_actions/unsubscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Unsubscribe from the user action log over AMQP
func (r *CloudV1UserActionService) UnsubscribeAmqp(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/user_actions/unsubscribe_amqp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type CloudV1UserActionListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1UserActionListResponseResult `json:"results,required"`
	JSON    cloudV1UserActionListResponseJSON     `json:"-"`
}

// cloudV1UserActionListResponseJSON contains the JSON metadata for the struct
// [CloudV1UserActionListResponse]
type cloudV1UserActionListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserActionListResponseResult struct {
	// User action ID
	ID string `json:"id,required"`
	// Action type
	ActionType CloudV1UserActionListResponseResultsActionType `json:"action_type,required"`
	// API group
	APIGroup CloudV1UserActionListResponseResultsAPIGroup `json:"api_group,required"`
	// User action was filled with all necessary information. If false, then something
	// went wrong during user action creation or update
	IsComplete bool `json:"is_complete,required"`
	// Datetime. Action timestamp
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// User action log was successfully received by its subscriber in case there is one
	Acknowledged bool `json:"acknowledged"`
	// Additional information about the action
	ActionData interface{} `json:"action_data,nullable"`
	// Client ID of the user.
	ClientID int64 `json:"client_id,nullable"`
	// User email address
	Email string `json:"email,nullable" format:"email"`
	// User ID who issued the token that made the request
	IssuedByUserID int64 `json:"issued_by_user_id,nullable"`
	// Project ID
	ProjectID int64 `json:"project_id,nullable"`
	// Region ID
	RegionID int64 `json:"region_id,nullable"`
	// Resources
	Resources []CloudV1UserActionListResponseResultsResource `json:"resources"`
	// Task ID
	TaskID string `json:"task_id,nullable"`
	// Token ID
	TokenID int64 `json:"token_id,nullable"`
	// Total resource price VAT inclusive
	TotalPrice CloudV1UserActionListResponseResultsTotalPrice `json:"total_price,nullable"`
	JSON       cloudV1UserActionListResponseResultJSON        `json:"-"`
}

// cloudV1UserActionListResponseResultJSON contains the JSON metadata for the
// struct [CloudV1UserActionListResponseResult]
type cloudV1UserActionListResponseResultJSON struct {
	ID             apijson.Field
	ActionType     apijson.Field
	APIGroup       apijson.Field
	IsComplete     apijson.Field
	Timestamp      apijson.Field
	UserID         apijson.Field
	Acknowledged   apijson.Field
	ActionData     apijson.Field
	ClientID       apijson.Field
	Email          apijson.Field
	IssuedByUserID apijson.Field
	ProjectID      apijson.Field
	RegionID       apijson.Field
	Resources      apijson.Field
	TaskID         apijson.Field
	TokenID        apijson.Field
	TotalPrice     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CloudV1UserActionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Action type
type CloudV1UserActionListResponseResultsActionType string

const (
	CloudV1UserActionListResponseResultsActionTypeActivate               CloudV1UserActionListResponseResultsActionType = "activate"
	CloudV1UserActionListResponseResultsActionTypeAttach                 CloudV1UserActionListResponseResultsActionType = "attach"
	CloudV1UserActionListResponseResultsActionTypeChangeLoggingResources CloudV1UserActionListResponseResultsActionType = "change_logging_resources"
	CloudV1UserActionListResponseResultsActionTypeCreate                 CloudV1UserActionListResponseResultsActionType = "create"
	CloudV1UserActionListResponseResultsActionTypeCreateAccessRule       CloudV1UserActionListResponseResultsActionType = "create_access_rule"
	CloudV1UserActionListResponseResultsActionTypeDeactivate             CloudV1UserActionListResponseResultsActionType = "deactivate"
	CloudV1UserActionListResponseResultsActionTypeDelete                 CloudV1UserActionListResponseResultsActionType = "delete"
	CloudV1UserActionListResponseResultsActionTypeDeleteAccessRule       CloudV1UserActionListResponseResultsActionType = "delete_access_rule"
	CloudV1UserActionListResponseResultsActionTypeDeleteMetadata         CloudV1UserActionListResponseResultsActionType = "delete_metadata"
	CloudV1UserActionListResponseResultsActionTypeDetach                 CloudV1UserActionListResponseResultsActionType = "detach"
	CloudV1UserActionListResponseResultsActionTypeDisableLogging         CloudV1UserActionListResponseResultsActionType = "disable_logging"
	CloudV1UserActionListResponseResultsActionTypeDisablePortsecurity    CloudV1UserActionListResponseResultsActionType = "disable_portsecurity"
	CloudV1UserActionListResponseResultsActionTypeDownload               CloudV1UserActionListResponseResultsActionType = "download"
	CloudV1UserActionListResponseResultsActionTypeEnableLogging          CloudV1UserActionListResponseResultsActionType = "enable_logging"
	CloudV1UserActionListResponseResultsActionTypeEnablePortsecurity     CloudV1UserActionListResponseResultsActionType = "enable_portsecurity"
	CloudV1UserActionListResponseResultsActionTypeFailover               CloudV1UserActionListResponseResultsActionType = "failover"
	CloudV1UserActionListResponseResultsActionTypePutIntoServergroup     CloudV1UserActionListResponseResultsActionType = "put_into_servergroup"
	CloudV1UserActionListResponseResultsActionTypeReboot                 CloudV1UserActionListResponseResultsActionType = "reboot"
	CloudV1UserActionListResponseResultsActionTypeRebootHard             CloudV1UserActionListResponseResultsActionType = "reboot_hard"
	CloudV1UserActionListResponseResultsActionTypeRebuild                CloudV1UserActionListResponseResultsActionType = "rebuild"
	CloudV1UserActionListResponseResultsActionTypeRegenerateCredentials  CloudV1UserActionListResponseResultsActionType = "regenerate_credentials"
	CloudV1UserActionListResponseResultsActionTypeRemoveFromServergroup  CloudV1UserActionListResponseResultsActionType = "remove_from_servergroup"
	CloudV1UserActionListResponseResultsActionTypeReplaceMetadata        CloudV1UserActionListResponseResultsActionType = "replace_metadata"
	CloudV1UserActionListResponseResultsActionTypeResize                 CloudV1UserActionListResponseResultsActionType = "resize"
	CloudV1UserActionListResponseResultsActionTypeResume                 CloudV1UserActionListResponseResultsActionType = "resume"
	CloudV1UserActionListResponseResultsActionTypeRetype                 CloudV1UserActionListResponseResultsActionType = "retype"
	CloudV1UserActionListResponseResultsActionTypeRevert                 CloudV1UserActionListResponseResultsActionType = "revert"
	CloudV1UserActionListResponseResultsActionTypeScaleDown              CloudV1UserActionListResponseResultsActionType = "scale_down"
	CloudV1UserActionListResponseResultsActionTypeScaleUp                CloudV1UserActionListResponseResultsActionType = "scale_up"
	CloudV1UserActionListResponseResultsActionTypeStart                  CloudV1UserActionListResponseResultsActionType = "start"
	CloudV1UserActionListResponseResultsActionTypeStop                   CloudV1UserActionListResponseResultsActionType = "stop"
	CloudV1UserActionListResponseResultsActionTypeSuspend                CloudV1UserActionListResponseResultsActionType = "suspend"
	CloudV1UserActionListResponseResultsActionTypeUpdate                 CloudV1UserActionListResponseResultsActionType = "update"
	CloudV1UserActionListResponseResultsActionTypeUpdateMetadata         CloudV1UserActionListResponseResultsActionType = "update_metadata"
	CloudV1UserActionListResponseResultsActionTypeUpgrade                CloudV1UserActionListResponseResultsActionType = "upgrade"
)

func (r CloudV1UserActionListResponseResultsActionType) IsKnown() bool {
	switch r {
	case CloudV1UserActionListResponseResultsActionTypeActivate, CloudV1UserActionListResponseResultsActionTypeAttach, CloudV1UserActionListResponseResultsActionTypeChangeLoggingResources, CloudV1UserActionListResponseResultsActionTypeCreate, CloudV1UserActionListResponseResultsActionTypeCreateAccessRule, CloudV1UserActionListResponseResultsActionTypeDeactivate, CloudV1UserActionListResponseResultsActionTypeDelete, CloudV1UserActionListResponseResultsActionTypeDeleteAccessRule, CloudV1UserActionListResponseResultsActionTypeDeleteMetadata, CloudV1UserActionListResponseResultsActionTypeDetach, CloudV1UserActionListResponseResultsActionTypeDisableLogging, CloudV1UserActionListResponseResultsActionTypeDisablePortsecurity, CloudV1UserActionListResponseResultsActionTypeDownload, CloudV1UserActionListResponseResultsActionTypeEnableLogging, CloudV1UserActionListResponseResultsActionTypeEnablePortsecurity, CloudV1UserActionListResponseResultsActionTypeFailover, CloudV1UserActionListResponseResultsActionTypePutIntoServergroup, CloudV1UserActionListResponseResultsActionTypeReboot, CloudV1UserActionListResponseResultsActionTypeRebootHard, CloudV1UserActionListResponseResultsActionTypeRebuild, CloudV1UserActionListResponseResultsActionTypeRegenerateCredentials, CloudV1UserActionListResponseResultsActionTypeRemoveFromServergroup, CloudV1UserActionListResponseResultsActionTypeReplaceMetadata, CloudV1UserActionListResponseResultsActionTypeResize, CloudV1UserActionListResponseResultsActionTypeResume, CloudV1UserActionListResponseResultsActionTypeRetype, CloudV1UserActionListResponseResultsActionTypeRevert, CloudV1UserActionListResponseResultsActionTypeScaleDown, CloudV1UserActionListResponseResultsActionTypeScaleUp, CloudV1UserActionListResponseResultsActionTypeStart, CloudV1UserActionListResponseResultsActionTypeStop, CloudV1UserActionListResponseResultsActionTypeSuspend, CloudV1UserActionListResponseResultsActionTypeUpdate, CloudV1UserActionListResponseResultsActionTypeUpdateMetadata, CloudV1UserActionListResponseResultsActionTypeUpgrade:
		return true
	}
	return false
}

// API group
type CloudV1UserActionListResponseResultsAPIGroup string

const (
	CloudV1UserActionListResponseResultsAPIGroupAICluster                   CloudV1UserActionListResponseResultsAPIGroup = "ai_cluster"
	CloudV1UserActionListResponseResultsAPIGroupBackupService               CloudV1UserActionListResponseResultsAPIGroup = "backup_service"
	CloudV1UserActionListResponseResultsAPIGroupCaasContainer               CloudV1UserActionListResponseResultsAPIGroup = "caas_container"
	CloudV1UserActionListResponseResultsAPIGroupCaasKey                     CloudV1UserActionListResponseResultsAPIGroup = "caas_key"
	CloudV1UserActionListResponseResultsAPIGroupCaasPullSecret              CloudV1UserActionListResponseResultsAPIGroup = "caas_pull_secret"
	CloudV1UserActionListResponseResultsAPIGroupDbaasPostgres               CloudV1UserActionListResponseResultsAPIGroup = "dbaas_postgres"
	CloudV1UserActionListResponseResultsAPIGroupDdosProfile                 CloudV1UserActionListResponseResultsAPIGroup = "ddos_profile"
	CloudV1UserActionListResponseResultsAPIGroupFaasFunction                CloudV1UserActionListResponseResultsAPIGroup = "faas_function"
	CloudV1UserActionListResponseResultsAPIGroupFaasKey                     CloudV1UserActionListResponseResultsAPIGroup = "faas_key"
	CloudV1UserActionListResponseResultsAPIGroupFaasNamespace               CloudV1UserActionListResponseResultsAPIGroup = "faas_namespace"
	CloudV1UserActionListResponseResultsAPIGroupFileShares                  CloudV1UserActionListResponseResultsAPIGroup = "file_shares"
	CloudV1UserActionListResponseResultsAPIGroupFloatingIP                  CloudV1UserActionListResponseResultsAPIGroup = "floating_ip"
	CloudV1UserActionListResponseResultsAPIGroupHeat                        CloudV1UserActionListResponseResultsAPIGroup = "heat"
	CloudV1UserActionListResponseResultsAPIGroupImage                       CloudV1UserActionListResponseResultsAPIGroup = "image"
	CloudV1UserActionListResponseResultsAPIGroupInferenceAtTheEdge          CloudV1UserActionListResponseResultsAPIGroup = "inference_at_the_edge"
	CloudV1UserActionListResponseResultsAPIGroupInstance                    CloudV1UserActionListResponseResultsAPIGroup = "instance"
	CloudV1UserActionListResponseResultsAPIGroupInstanceIsolation           CloudV1UserActionListResponseResultsAPIGroup = "instance_isolation"
	CloudV1UserActionListResponseResultsAPIGroupK8sCluster                  CloudV1UserActionListResponseResultsAPIGroup = "k8s_cluster"
	CloudV1UserActionListResponseResultsAPIGroupK8sClusterTemplate          CloudV1UserActionListResponseResultsAPIGroup = "k8s_cluster_template"
	CloudV1UserActionListResponseResultsAPIGroupK8sPool                     CloudV1UserActionListResponseResultsAPIGroup = "k8s_pool"
	CloudV1UserActionListResponseResultsAPIGroupLaas                        CloudV1UserActionListResponseResultsAPIGroup = "laas"
	CloudV1UserActionListResponseResultsAPIGroupLaasTopic                   CloudV1UserActionListResponseResultsAPIGroup = "laas_topic"
	CloudV1UserActionListResponseResultsAPIGroupLbHealthMonitor             CloudV1UserActionListResponseResultsAPIGroup = "lb_health_monitor"
	CloudV1UserActionListResponseResultsAPIGroupLbL7policy                  CloudV1UserActionListResponseResultsAPIGroup = "lb_l7policy"
	CloudV1UserActionListResponseResultsAPIGroupLbL7rule                    CloudV1UserActionListResponseResultsAPIGroup = "lb_l7rule"
	CloudV1UserActionListResponseResultsAPIGroupLblistener                  CloudV1UserActionListResponseResultsAPIGroup = "lblistener"
	CloudV1UserActionListResponseResultsAPIGroupLbpool                      CloudV1UserActionListResponseResultsAPIGroup = "lbpool"
	CloudV1UserActionListResponseResultsAPIGroupLbpoolMember                CloudV1UserActionListResponseResultsAPIGroup = "lbpool_member"
	CloudV1UserActionListResponseResultsAPIGroupLifecyclePolicy             CloudV1UserActionListResponseResultsAPIGroup = "lifecycle_policy"
	CloudV1UserActionListResponseResultsAPIGroupLifecyclePolicyVolumeMember CloudV1UserActionListResponseResultsAPIGroup = "lifecycle_policy_volume_member"
	CloudV1UserActionListResponseResultsAPIGroupLoadbalancer                CloudV1UserActionListResponseResultsAPIGroup = "loadbalancer"
	CloudV1UserActionListResponseResultsAPIGroupNetwork                     CloudV1UserActionListResponseResultsAPIGroup = "network"
	CloudV1UserActionListResponseResultsAPIGroupPort                        CloudV1UserActionListResponseResultsAPIGroup = "port"
	CloudV1UserActionListResponseResultsAPIGroupProject                     CloudV1UserActionListResponseResultsAPIGroup = "project"
	CloudV1UserActionListResponseResultsAPIGroupQuotaLimitRequest           CloudV1UserActionListResponseResultsAPIGroup = "quota_limit_request"
	CloudV1UserActionListResponseResultsAPIGroupRegistry                    CloudV1UserActionListResponseResultsAPIGroup = "registry"
	CloudV1UserActionListResponseResultsAPIGroupReservation                 CloudV1UserActionListResponseResultsAPIGroup = "reservation"
	CloudV1UserActionListResponseResultsAPIGroupReservedFixedIP             CloudV1UserActionListResponseResultsAPIGroup = "reserved_fixed_ip"
	CloudV1UserActionListResponseResultsAPIGroupRole                        CloudV1UserActionListResponseResultsAPIGroup = "role"
	CloudV1UserActionListResponseResultsAPIGroupRouter                      CloudV1UserActionListResponseResultsAPIGroup = "router"
	CloudV1UserActionListResponseResultsAPIGroupSecret                      CloudV1UserActionListResponseResultsAPIGroup = "secret"
	CloudV1UserActionListResponseResultsAPIGroupSecuritygroup               CloudV1UserActionListResponseResultsAPIGroup = "securitygroup"
	CloudV1UserActionListResponseResultsAPIGroupSecuritygrouprule           CloudV1UserActionListResponseResultsAPIGroup = "securitygrouprule"
	CloudV1UserActionListResponseResultsAPIGroupServergroup                 CloudV1UserActionListResponseResultsAPIGroup = "servergroup"
	CloudV1UserActionListResponseResultsAPIGroupSharedFlavor                CloudV1UserActionListResponseResultsAPIGroup = "shared_flavor"
	CloudV1UserActionListResponseResultsAPIGroupSharedImage                 CloudV1UserActionListResponseResultsAPIGroup = "shared_image"
	CloudV1UserActionListResponseResultsAPIGroupSharedNetwork               CloudV1UserActionListResponseResultsAPIGroup = "shared_network"
	CloudV1UserActionListResponseResultsAPIGroupSnapshot                    CloudV1UserActionListResponseResultsAPIGroup = "snapshot"
	CloudV1UserActionListResponseResultsAPIGroupSnapshotSchedule            CloudV1UserActionListResponseResultsAPIGroup = "snapshot_schedule"
	CloudV1UserActionListResponseResultsAPIGroupSSHKey                      CloudV1UserActionListResponseResultsAPIGroup = "ssh_key"
	CloudV1UserActionListResponseResultsAPIGroupSubnet                      CloudV1UserActionListResponseResultsAPIGroup = "subnet"
	CloudV1UserActionListResponseResultsAPIGroupUser                        CloudV1UserActionListResponseResultsAPIGroup = "user"
	CloudV1UserActionListResponseResultsAPIGroupVipIPAddresses              CloudV1UserActionListResponseResultsAPIGroup = "vip_ip_addresses"
	CloudV1UserActionListResponseResultsAPIGroupVolume                      CloudV1UserActionListResponseResultsAPIGroup = "volume"
)

func (r CloudV1UserActionListResponseResultsAPIGroup) IsKnown() bool {
	switch r {
	case CloudV1UserActionListResponseResultsAPIGroupAICluster, CloudV1UserActionListResponseResultsAPIGroupBackupService, CloudV1UserActionListResponseResultsAPIGroupCaasContainer, CloudV1UserActionListResponseResultsAPIGroupCaasKey, CloudV1UserActionListResponseResultsAPIGroupCaasPullSecret, CloudV1UserActionListResponseResultsAPIGroupDbaasPostgres, CloudV1UserActionListResponseResultsAPIGroupDdosProfile, CloudV1UserActionListResponseResultsAPIGroupFaasFunction, CloudV1UserActionListResponseResultsAPIGroupFaasKey, CloudV1UserActionListResponseResultsAPIGroupFaasNamespace, CloudV1UserActionListResponseResultsAPIGroupFileShares, CloudV1UserActionListResponseResultsAPIGroupFloatingIP, CloudV1UserActionListResponseResultsAPIGroupHeat, CloudV1UserActionListResponseResultsAPIGroupImage, CloudV1UserActionListResponseResultsAPIGroupInferenceAtTheEdge, CloudV1UserActionListResponseResultsAPIGroupInstance, CloudV1UserActionListResponseResultsAPIGroupInstanceIsolation, CloudV1UserActionListResponseResultsAPIGroupK8sCluster, CloudV1UserActionListResponseResultsAPIGroupK8sClusterTemplate, CloudV1UserActionListResponseResultsAPIGroupK8sPool, CloudV1UserActionListResponseResultsAPIGroupLaas, CloudV1UserActionListResponseResultsAPIGroupLaasTopic, CloudV1UserActionListResponseResultsAPIGroupLbHealthMonitor, CloudV1UserActionListResponseResultsAPIGroupLbL7policy, CloudV1UserActionListResponseResultsAPIGroupLbL7rule, CloudV1UserActionListResponseResultsAPIGroupLblistener, CloudV1UserActionListResponseResultsAPIGroupLbpool, CloudV1UserActionListResponseResultsAPIGroupLbpoolMember, CloudV1UserActionListResponseResultsAPIGroupLifecyclePolicy, CloudV1UserActionListResponseResultsAPIGroupLifecyclePolicyVolumeMember, CloudV1UserActionListResponseResultsAPIGroupLoadbalancer, CloudV1UserActionListResponseResultsAPIGroupNetwork, CloudV1UserActionListResponseResultsAPIGroupPort, CloudV1UserActionListResponseResultsAPIGroupProject, CloudV1UserActionListResponseResultsAPIGroupQuotaLimitRequest, CloudV1UserActionListResponseResultsAPIGroupRegistry, CloudV1UserActionListResponseResultsAPIGroupReservation, CloudV1UserActionListResponseResultsAPIGroupReservedFixedIP, CloudV1UserActionListResponseResultsAPIGroupRole, CloudV1UserActionListResponseResultsAPIGroupRouter, CloudV1UserActionListResponseResultsAPIGroupSecret, CloudV1UserActionListResponseResultsAPIGroupSecuritygroup, CloudV1UserActionListResponseResultsAPIGroupSecuritygrouprule, CloudV1UserActionListResponseResultsAPIGroupServergroup, CloudV1UserActionListResponseResultsAPIGroupSharedFlavor, CloudV1UserActionListResponseResultsAPIGroupSharedImage, CloudV1UserActionListResponseResultsAPIGroupSharedNetwork, CloudV1UserActionListResponseResultsAPIGroupSnapshot, CloudV1UserActionListResponseResultsAPIGroupSnapshotSchedule, CloudV1UserActionListResponseResultsAPIGroupSSHKey, CloudV1UserActionListResponseResultsAPIGroupSubnet, CloudV1UserActionListResponseResultsAPIGroupUser, CloudV1UserActionListResponseResultsAPIGroupVipIPAddresses, CloudV1UserActionListResponseResultsAPIGroupVolume:
		return true
	}
	return false
}

type CloudV1UserActionListResponseResultsResource struct {
	// Resource ID
	ResourceID string `json:"resource_id,required"`
	// Resource type
	ResourceType CloudV1UserActionListResponseResultsResourcesResourceType `json:"resource_type,required"`
	// Free-form object, resource body.
	ResourceBody interface{} `json:"resource_body,nullable"`
	// Often used property for filtering actions. It can be a name, IP address, or
	// other property, depending on the resource_type
	SearchField string                                           `json:"search_field,nullable"`
	JSON        cloudV1UserActionListResponseResultsResourceJSON `json:"-"`
}

// cloudV1UserActionListResponseResultsResourceJSON contains the JSON metadata for
// the struct [CloudV1UserActionListResponseResultsResource]
type cloudV1UserActionListResponseResultsResourceJSON struct {
	ResourceID   apijson.Field
	ResourceType apijson.Field
	ResourceBody apijson.Field
	SearchField  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CloudV1UserActionListResponseResultsResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionListResponseResultsResourceJSON) RawJSON() string {
	return r.raw
}

// Resource type
type CloudV1UserActionListResponseResultsResourcesResourceType string

const (
	CloudV1UserActionListResponseResultsResourcesResourceTypeCaasContainer                CloudV1UserActionListResponseResultsResourcesResourceType = "caas_container"
	CloudV1UserActionListResponseResultsResourcesResourceTypeCaasKey                      CloudV1UserActionListResponseResultsResourcesResourceType = "caas_key"
	CloudV1UserActionListResponseResultsResourcesResourceTypeCaasPullSecret               CloudV1UserActionListResponseResultsResourcesResourceType = "caas_pull_secret"
	CloudV1UserActionListResponseResultsResourcesResourceTypeDbaasPostgres                CloudV1UserActionListResponseResultsResourcesResourceType = "dbaas_postgres"
	CloudV1UserActionListResponseResultsResourcesResourceTypeDdosProfile                  CloudV1UserActionListResponseResultsResourcesResourceType = "ddos_profile"
	CloudV1UserActionListResponseResultsResourcesResourceTypeFaasFunction                 CloudV1UserActionListResponseResultsResourcesResourceType = "faas_function"
	CloudV1UserActionListResponseResultsResourcesResourceTypeFaasKey                      CloudV1UserActionListResponseResultsResourcesResourceType = "faas_key"
	CloudV1UserActionListResponseResultsResourcesResourceTypeFaasNamespace                CloudV1UserActionListResponseResultsResourcesResourceType = "faas_namespace"
	CloudV1UserActionListResponseResultsResourcesResourceTypeFileShares                   CloudV1UserActionListResponseResultsResourcesResourceType = "file_shares"
	CloudV1UserActionListResponseResultsResourcesResourceTypeFloatingIP                   CloudV1UserActionListResponseResultsResourcesResourceType = "floating_ip"
	CloudV1UserActionListResponseResultsResourcesResourceTypeGpuaiCluster                 CloudV1UserActionListResponseResultsResourcesResourceType = "gpuai_cluster"
	CloudV1UserActionListResponseResultsResourcesResourceTypeHeat                         CloudV1UserActionListResponseResultsResourcesResourceType = "heat"
	CloudV1UserActionListResponseResultsResourcesResourceTypeImage                        CloudV1UserActionListResponseResultsResourcesResourceType = "image"
	CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceInstance            CloudV1UserActionListResponseResultsResourcesResourceType = "inference_instance"
	CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceRegistryCredentials CloudV1UserActionListResponseResultsResourcesResourceType = "inference_registry_credentials"
	CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceSecret              CloudV1UserActionListResponseResultsResourcesResourceType = "inference_secret"
	CloudV1UserActionListResponseResultsResourcesResourceTypeInstance                     CloudV1UserActionListResponseResultsResourcesResourceType = "instance"
	CloudV1UserActionListResponseResultsResourcesResourceTypeIpuCluster                   CloudV1UserActionListResponseResultsResourcesResourceType = "ipu_cluster"
	CloudV1UserActionListResponseResultsResourcesResourceTypeK8sCluster                   CloudV1UserActionListResponseResultsResourcesResourceType = "k8s_cluster"
	CloudV1UserActionListResponseResultsResourcesResourceTypeK8sClusterTemplate           CloudV1UserActionListResponseResultsResourcesResourceType = "k8s_cluster_template"
	CloudV1UserActionListResponseResultsResourcesResourceTypeK8sPool                      CloudV1UserActionListResponseResultsResourcesResourceType = "k8s_pool"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLaas                         CloudV1UserActionListResponseResultsResourcesResourceType = "laas"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLaasTopic                    CloudV1UserActionListResponseResultsResourcesResourceType = "laas_topic"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLbHealthMonitor              CloudV1UserActionListResponseResultsResourcesResourceType = "lb_health_monitor"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLbL7policy                   CloudV1UserActionListResponseResultsResourcesResourceType = "lb_l7policy"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLbL7rule                     CloudV1UserActionListResponseResultsResourcesResourceType = "lb_l7rule"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLblistener                   CloudV1UserActionListResponseResultsResourcesResourceType = "lblistener"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLbpool                       CloudV1UserActionListResponseResultsResourcesResourceType = "lbpool"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLbpoolMember                 CloudV1UserActionListResponseResultsResourcesResourceType = "lbpool_member"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLifecyclePolicy              CloudV1UserActionListResponseResultsResourcesResourceType = "lifecycle_policy"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLifecyclePolicyVolumeMember  CloudV1UserActionListResponseResultsResourcesResourceType = "lifecycle_policy_volume_member"
	CloudV1UserActionListResponseResultsResourcesResourceTypeLoadbalancer                 CloudV1UserActionListResponseResultsResourcesResourceType = "loadbalancer"
	CloudV1UserActionListResponseResultsResourcesResourceTypeNetwork                      CloudV1UserActionListResponseResultsResourcesResourceType = "network"
	CloudV1UserActionListResponseResultsResourcesResourceTypePort                         CloudV1UserActionListResponseResultsResourcesResourceType = "port"
	CloudV1UserActionListResponseResultsResourcesResourceTypeProject                      CloudV1UserActionListResponseResultsResourcesResourceType = "project"
	CloudV1UserActionListResponseResultsResourcesResourceTypeQuotaLimitRequest            CloudV1UserActionListResponseResultsResourcesResourceType = "quota_limit_request"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistry                     CloudV1UserActionListResponseResultsResourcesResourceType = "registry"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepository           CloudV1UserActionListResponseResultsResourcesResourceType = "registry_repository"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepositoryArtifact   CloudV1UserActionListResponseResultsResourcesResourceType = "registry_repository_artifact"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepositoryTag        CloudV1UserActionListResponseResultsResourcesResourceType = "registry_repository_tag"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryUser                 CloudV1UserActionListResponseResultsResourcesResourceType = "registry_user"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryUserSercret          CloudV1UserActionListResponseResultsResourcesResourceType = "registry_user_sercret"
	CloudV1UserActionListResponseResultsResourcesResourceTypeReservation                  CloudV1UserActionListResponseResultsResourcesResourceType = "reservation"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRole                         CloudV1UserActionListResponseResultsResourcesResourceType = "role"
	CloudV1UserActionListResponseResultsResourcesResourceTypeRouter                       CloudV1UserActionListResponseResultsResourcesResourceType = "router"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSecret                       CloudV1UserActionListResponseResultsResourcesResourceType = "secret"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSecuritygroup                CloudV1UserActionListResponseResultsResourcesResourceType = "securitygroup"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSecuritygrouprule            CloudV1UserActionListResponseResultsResourcesResourceType = "securitygrouprule"
	CloudV1UserActionListResponseResultsResourcesResourceTypeServergroup                  CloudV1UserActionListResponseResultsResourcesResourceType = "servergroup"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSharedFlavor                 CloudV1UserActionListResponseResultsResourcesResourceType = "shared_flavor"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSharedImage                  CloudV1UserActionListResponseResultsResourcesResourceType = "shared_image"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSharedNetwork                CloudV1UserActionListResponseResultsResourcesResourceType = "shared_network"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSnapshot                     CloudV1UserActionListResponseResultsResourcesResourceType = "snapshot"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSnapshotSchedule             CloudV1UserActionListResponseResultsResourcesResourceType = "snapshot_schedule"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSSHKey                       CloudV1UserActionListResponseResultsResourcesResourceType = "ssh_key"
	CloudV1UserActionListResponseResultsResourcesResourceTypeSubnet                       CloudV1UserActionListResponseResultsResourcesResourceType = "subnet"
	CloudV1UserActionListResponseResultsResourcesResourceTypeToken                        CloudV1UserActionListResponseResultsResourcesResourceType = "token"
	CloudV1UserActionListResponseResultsResourcesResourceTypeUser                         CloudV1UserActionListResponseResultsResourcesResourceType = "user"
	CloudV1UserActionListResponseResultsResourcesResourceTypeVirtualGpuaiCluster          CloudV1UserActionListResponseResultsResourcesResourceType = "virtual_gpuai_cluster"
	CloudV1UserActionListResponseResultsResourcesResourceTypeVolume                       CloudV1UserActionListResponseResultsResourcesResourceType = "volume"
)

func (r CloudV1UserActionListResponseResultsResourcesResourceType) IsKnown() bool {
	switch r {
	case CloudV1UserActionListResponseResultsResourcesResourceTypeCaasContainer, CloudV1UserActionListResponseResultsResourcesResourceTypeCaasKey, CloudV1UserActionListResponseResultsResourcesResourceTypeCaasPullSecret, CloudV1UserActionListResponseResultsResourcesResourceTypeDbaasPostgres, CloudV1UserActionListResponseResultsResourcesResourceTypeDdosProfile, CloudV1UserActionListResponseResultsResourcesResourceTypeFaasFunction, CloudV1UserActionListResponseResultsResourcesResourceTypeFaasKey, CloudV1UserActionListResponseResultsResourcesResourceTypeFaasNamespace, CloudV1UserActionListResponseResultsResourcesResourceTypeFileShares, CloudV1UserActionListResponseResultsResourcesResourceTypeFloatingIP, CloudV1UserActionListResponseResultsResourcesResourceTypeGpuaiCluster, CloudV1UserActionListResponseResultsResourcesResourceTypeHeat, CloudV1UserActionListResponseResultsResourcesResourceTypeImage, CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceInstance, CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceRegistryCredentials, CloudV1UserActionListResponseResultsResourcesResourceTypeInferenceSecret, CloudV1UserActionListResponseResultsResourcesResourceTypeInstance, CloudV1UserActionListResponseResultsResourcesResourceTypeIpuCluster, CloudV1UserActionListResponseResultsResourcesResourceTypeK8sCluster, CloudV1UserActionListResponseResultsResourcesResourceTypeK8sClusterTemplate, CloudV1UserActionListResponseResultsResourcesResourceTypeK8sPool, CloudV1UserActionListResponseResultsResourcesResourceTypeLaas, CloudV1UserActionListResponseResultsResourcesResourceTypeLaasTopic, CloudV1UserActionListResponseResultsResourcesResourceTypeLbHealthMonitor, CloudV1UserActionListResponseResultsResourcesResourceTypeLbL7policy, CloudV1UserActionListResponseResultsResourcesResourceTypeLbL7rule, CloudV1UserActionListResponseResultsResourcesResourceTypeLblistener, CloudV1UserActionListResponseResultsResourcesResourceTypeLbpool, CloudV1UserActionListResponseResultsResourcesResourceTypeLbpoolMember, CloudV1UserActionListResponseResultsResourcesResourceTypeLifecyclePolicy, CloudV1UserActionListResponseResultsResourcesResourceTypeLifecyclePolicyVolumeMember, CloudV1UserActionListResponseResultsResourcesResourceTypeLoadbalancer, CloudV1UserActionListResponseResultsResourcesResourceTypeNetwork, CloudV1UserActionListResponseResultsResourcesResourceTypePort, CloudV1UserActionListResponseResultsResourcesResourceTypeProject, CloudV1UserActionListResponseResultsResourcesResourceTypeQuotaLimitRequest, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistry, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepository, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepositoryArtifact, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryRepositoryTag, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryUser, CloudV1UserActionListResponseResultsResourcesResourceTypeRegistryUserSercret, CloudV1UserActionListResponseResultsResourcesResourceTypeReservation, CloudV1UserActionListResponseResultsResourcesResourceTypeRole, CloudV1UserActionListResponseResultsResourcesResourceTypeRouter, CloudV1UserActionListResponseResultsResourcesResourceTypeSecret, CloudV1UserActionListResponseResultsResourcesResourceTypeSecuritygroup, CloudV1UserActionListResponseResultsResourcesResourceTypeSecuritygrouprule, CloudV1UserActionListResponseResultsResourcesResourceTypeServergroup, CloudV1UserActionListResponseResultsResourcesResourceTypeSharedFlavor, CloudV1UserActionListResponseResultsResourcesResourceTypeSharedImage, CloudV1UserActionListResponseResultsResourcesResourceTypeSharedNetwork, CloudV1UserActionListResponseResultsResourcesResourceTypeSnapshot, CloudV1UserActionListResponseResultsResourcesResourceTypeSnapshotSchedule, CloudV1UserActionListResponseResultsResourcesResourceTypeSSHKey, CloudV1UserActionListResponseResultsResourcesResourceTypeSubnet, CloudV1UserActionListResponseResultsResourcesResourceTypeToken, CloudV1UserActionListResponseResultsResourcesResourceTypeUser, CloudV1UserActionListResponseResultsResourcesResourceTypeVirtualGpuaiCluster, CloudV1UserActionListResponseResultsResourcesResourceTypeVolume:
		return true
	}
	return false
}

// Total resource price VAT inclusive
type CloudV1UserActionListResponseResultsTotalPrice struct {
	// Currency code (3 letter code per ISO 4217)
	CurrencyCode string `json:"currency_code,required,nullable"`
	// Total price VAT inclusive per hour
	PricePerHour float64 `json:"price_per_hour,required,nullable"`
	// Total price VAT inclusive per month (30 days)
	PricePerMonth float64 `json:"price_per_month,required,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus                                 `json:"price_status,required"`
	JSON        cloudV1UserActionListResponseResultsTotalPriceJSON `json:"-"`
}

// cloudV1UserActionListResponseResultsTotalPriceJSON contains the JSON metadata
// for the struct [CloudV1UserActionListResponseResultsTotalPrice]
type cloudV1UserActionListResponseResultsTotalPriceJSON struct {
	CurrencyCode  apijson.Field
	PricePerHour  apijson.Field
	PricePerMonth apijson.Field
	PriceStatus   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1UserActionListResponseResultsTotalPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionListResponseResultsTotalPriceJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserActionGetSubscriptionsListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1UserActionGetSubscriptionsListResponseResult `json:"results,required"`
	JSON    cloudV1UserActionGetSubscriptionsListResponseJSON     `json:"-"`
}

// cloudV1UserActionGetSubscriptionsListResponseJSON contains the JSON metadata for
// the struct [CloudV1UserActionGetSubscriptionsListResponse]
type cloudV1UserActionGetSubscriptionsListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserActionGetSubscriptionsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionGetSubscriptionsListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserActionGetSubscriptionsListResponseResult struct {
	// Subscription ID
	ID int64 `json:"id,required"`
	// Authorization header name
	AuthHeaderName string `json:"auth_header_name,required"`
	// Authorization header value
	AuthHeaderValue string `json:"auth_header_value,required"`
	// URL to send user action logs for the specified client
	URL  string                                                  `json:"url,required"`
	JSON cloudV1UserActionGetSubscriptionsListResponseResultJSON `json:"-"`
}

// cloudV1UserActionGetSubscriptionsListResponseResultJSON contains the JSON
// metadata for the struct [CloudV1UserActionGetSubscriptionsListResponseResult]
type cloudV1UserActionGetSubscriptionsListResponseResultJSON struct {
	ID              apijson.Field
	AuthHeaderName  apijson.Field
	AuthHeaderValue apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CloudV1UserActionGetSubscriptionsListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserActionGetSubscriptionsListResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserActionListParams struct {
	// User action type. Several options can be specified.
	ActionType param.Field[CloudV1UserActionListParamsActionType] `query:"action_type"`
	// API group that requested action belongs to. Several options can be specified.
	APIGroup param.Field[CloudV1UserActionListParamsAPIGroup] `query:"api_group"`
	// ISO formatted datetime string. Starting timestamp from which user actions are
	// requested
	FromTimestamp param.Field[string] `query:"from_timestamp"`
	// Limit the number of returned user actions. Falls back to default if not
	// specified. Limited by max limit value
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Project ID. Several options can be specified.
	ProjectID param.Field[int64] `query:"project_id"`
	// Resource ID. Several options can be specified.
	ResourceID param.Field[string] `query:"resource_id"`
	// Name or IP address of the resource
	SearchField param.Field[string] `query:"search_field"`
	// Sorting by date. Oldest first, or most recent first
	Sorting param.Field[CloudV1UserActionListParamsSorting] `query:"sorting"`
	// ISO formatted datetime string. Ending timestamp until which user actions are
	// requested
	ToTimestamp param.Field[string] `query:"to_timestamp"`
}

// URLQuery serializes [CloudV1UserActionListParams]'s query parameters as
// `url.Values`.
func (r CloudV1UserActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// User action type. Several options can be specified.
type CloudV1UserActionListParamsActionType string

const (
	CloudV1UserActionListParamsActionTypeActivate               CloudV1UserActionListParamsActionType = "activate"
	CloudV1UserActionListParamsActionTypeAttach                 CloudV1UserActionListParamsActionType = "attach"
	CloudV1UserActionListParamsActionTypeChangeLoggingResources CloudV1UserActionListParamsActionType = "change_logging_resources"
	CloudV1UserActionListParamsActionTypeCreate                 CloudV1UserActionListParamsActionType = "create"
	CloudV1UserActionListParamsActionTypeCreateAccessRule       CloudV1UserActionListParamsActionType = "create_access_rule"
	CloudV1UserActionListParamsActionTypeDeactivate             CloudV1UserActionListParamsActionType = "deactivate"
	CloudV1UserActionListParamsActionTypeDelete                 CloudV1UserActionListParamsActionType = "delete"
	CloudV1UserActionListParamsActionTypeDeleteAccessRule       CloudV1UserActionListParamsActionType = "delete_access_rule"
	CloudV1UserActionListParamsActionTypeDeleteMetadata         CloudV1UserActionListParamsActionType = "delete_metadata"
	CloudV1UserActionListParamsActionTypeDetach                 CloudV1UserActionListParamsActionType = "detach"
	CloudV1UserActionListParamsActionTypeDisableLogging         CloudV1UserActionListParamsActionType = "disable_logging"
	CloudV1UserActionListParamsActionTypeDisablePortsecurity    CloudV1UserActionListParamsActionType = "disable_portsecurity"
	CloudV1UserActionListParamsActionTypeDownload               CloudV1UserActionListParamsActionType = "download"
	CloudV1UserActionListParamsActionTypeEnableLogging          CloudV1UserActionListParamsActionType = "enable_logging"
	CloudV1UserActionListParamsActionTypeEnablePortsecurity     CloudV1UserActionListParamsActionType = "enable_portsecurity"
	CloudV1UserActionListParamsActionTypeFailover               CloudV1UserActionListParamsActionType = "failover"
	CloudV1UserActionListParamsActionTypePutIntoServergroup     CloudV1UserActionListParamsActionType = "put_into_servergroup"
	CloudV1UserActionListParamsActionTypeReboot                 CloudV1UserActionListParamsActionType = "reboot"
	CloudV1UserActionListParamsActionTypeRebootHard             CloudV1UserActionListParamsActionType = "reboot_hard"
	CloudV1UserActionListParamsActionTypeRebuild                CloudV1UserActionListParamsActionType = "rebuild"
	CloudV1UserActionListParamsActionTypeRegenerateCredentials  CloudV1UserActionListParamsActionType = "regenerate_credentials"
	CloudV1UserActionListParamsActionTypeRemoveFromServergroup  CloudV1UserActionListParamsActionType = "remove_from_servergroup"
	CloudV1UserActionListParamsActionTypeReplaceMetadata        CloudV1UserActionListParamsActionType = "replace_metadata"
	CloudV1UserActionListParamsActionTypeResize                 CloudV1UserActionListParamsActionType = "resize"
	CloudV1UserActionListParamsActionTypeResume                 CloudV1UserActionListParamsActionType = "resume"
	CloudV1UserActionListParamsActionTypeRetype                 CloudV1UserActionListParamsActionType = "retype"
	CloudV1UserActionListParamsActionTypeRevert                 CloudV1UserActionListParamsActionType = "revert"
	CloudV1UserActionListParamsActionTypeScaleDown              CloudV1UserActionListParamsActionType = "scale_down"
	CloudV1UserActionListParamsActionTypeScaleUp                CloudV1UserActionListParamsActionType = "scale_up"
	CloudV1UserActionListParamsActionTypeStart                  CloudV1UserActionListParamsActionType = "start"
	CloudV1UserActionListParamsActionTypeStop                   CloudV1UserActionListParamsActionType = "stop"
	CloudV1UserActionListParamsActionTypeSuspend                CloudV1UserActionListParamsActionType = "suspend"
	CloudV1UserActionListParamsActionTypeUpdate                 CloudV1UserActionListParamsActionType = "update"
	CloudV1UserActionListParamsActionTypeUpdateMetadata         CloudV1UserActionListParamsActionType = "update_metadata"
	CloudV1UserActionListParamsActionTypeUpgrade                CloudV1UserActionListParamsActionType = "upgrade"
)

func (r CloudV1UserActionListParamsActionType) IsKnown() bool {
	switch r {
	case CloudV1UserActionListParamsActionTypeActivate, CloudV1UserActionListParamsActionTypeAttach, CloudV1UserActionListParamsActionTypeChangeLoggingResources, CloudV1UserActionListParamsActionTypeCreate, CloudV1UserActionListParamsActionTypeCreateAccessRule, CloudV1UserActionListParamsActionTypeDeactivate, CloudV1UserActionListParamsActionTypeDelete, CloudV1UserActionListParamsActionTypeDeleteAccessRule, CloudV1UserActionListParamsActionTypeDeleteMetadata, CloudV1UserActionListParamsActionTypeDetach, CloudV1UserActionListParamsActionTypeDisableLogging, CloudV1UserActionListParamsActionTypeDisablePortsecurity, CloudV1UserActionListParamsActionTypeDownload, CloudV1UserActionListParamsActionTypeEnableLogging, CloudV1UserActionListParamsActionTypeEnablePortsecurity, CloudV1UserActionListParamsActionTypeFailover, CloudV1UserActionListParamsActionTypePutIntoServergroup, CloudV1UserActionListParamsActionTypeReboot, CloudV1UserActionListParamsActionTypeRebootHard, CloudV1UserActionListParamsActionTypeRebuild, CloudV1UserActionListParamsActionTypeRegenerateCredentials, CloudV1UserActionListParamsActionTypeRemoveFromServergroup, CloudV1UserActionListParamsActionTypeReplaceMetadata, CloudV1UserActionListParamsActionTypeResize, CloudV1UserActionListParamsActionTypeResume, CloudV1UserActionListParamsActionTypeRetype, CloudV1UserActionListParamsActionTypeRevert, CloudV1UserActionListParamsActionTypeScaleDown, CloudV1UserActionListParamsActionTypeScaleUp, CloudV1UserActionListParamsActionTypeStart, CloudV1UserActionListParamsActionTypeStop, CloudV1UserActionListParamsActionTypeSuspend, CloudV1UserActionListParamsActionTypeUpdate, CloudV1UserActionListParamsActionTypeUpdateMetadata, CloudV1UserActionListParamsActionTypeUpgrade:
		return true
	}
	return false
}

// API group that requested action belongs to. Several options can be specified.
type CloudV1UserActionListParamsAPIGroup string

const (
	CloudV1UserActionListParamsAPIGroupAICluster                   CloudV1UserActionListParamsAPIGroup = "ai_cluster"
	CloudV1UserActionListParamsAPIGroupBackupService               CloudV1UserActionListParamsAPIGroup = "backup_service"
	CloudV1UserActionListParamsAPIGroupCaasContainer               CloudV1UserActionListParamsAPIGroup = "caas_container"
	CloudV1UserActionListParamsAPIGroupCaasKey                     CloudV1UserActionListParamsAPIGroup = "caas_key"
	CloudV1UserActionListParamsAPIGroupCaasPullSecret              CloudV1UserActionListParamsAPIGroup = "caas_pull_secret"
	CloudV1UserActionListParamsAPIGroupDbaasPostgres               CloudV1UserActionListParamsAPIGroup = "dbaas_postgres"
	CloudV1UserActionListParamsAPIGroupDdosProfile                 CloudV1UserActionListParamsAPIGroup = "ddos_profile"
	CloudV1UserActionListParamsAPIGroupFaasFunction                CloudV1UserActionListParamsAPIGroup = "faas_function"
	CloudV1UserActionListParamsAPIGroupFaasKey                     CloudV1UserActionListParamsAPIGroup = "faas_key"
	CloudV1UserActionListParamsAPIGroupFaasNamespace               CloudV1UserActionListParamsAPIGroup = "faas_namespace"
	CloudV1UserActionListParamsAPIGroupFileShares                  CloudV1UserActionListParamsAPIGroup = "file_shares"
	CloudV1UserActionListParamsAPIGroupFloatingIP                  CloudV1UserActionListParamsAPIGroup = "floating_ip"
	CloudV1UserActionListParamsAPIGroupHeat                        CloudV1UserActionListParamsAPIGroup = "heat"
	CloudV1UserActionListParamsAPIGroupImage                       CloudV1UserActionListParamsAPIGroup = "image"
	CloudV1UserActionListParamsAPIGroupInferenceAtTheEdge          CloudV1UserActionListParamsAPIGroup = "inference_at_the_edge"
	CloudV1UserActionListParamsAPIGroupInstance                    CloudV1UserActionListParamsAPIGroup = "instance"
	CloudV1UserActionListParamsAPIGroupInstanceIsolation           CloudV1UserActionListParamsAPIGroup = "instance_isolation"
	CloudV1UserActionListParamsAPIGroupK8sCluster                  CloudV1UserActionListParamsAPIGroup = "k8s_cluster"
	CloudV1UserActionListParamsAPIGroupK8sClusterTemplate          CloudV1UserActionListParamsAPIGroup = "k8s_cluster_template"
	CloudV1UserActionListParamsAPIGroupK8sPool                     CloudV1UserActionListParamsAPIGroup = "k8s_pool"
	CloudV1UserActionListParamsAPIGroupLaas                        CloudV1UserActionListParamsAPIGroup = "laas"
	CloudV1UserActionListParamsAPIGroupLaasTopic                   CloudV1UserActionListParamsAPIGroup = "laas_topic"
	CloudV1UserActionListParamsAPIGroupLbHealthMonitor             CloudV1UserActionListParamsAPIGroup = "lb_health_monitor"
	CloudV1UserActionListParamsAPIGroupLbL7policy                  CloudV1UserActionListParamsAPIGroup = "lb_l7policy"
	CloudV1UserActionListParamsAPIGroupLbL7rule                    CloudV1UserActionListParamsAPIGroup = "lb_l7rule"
	CloudV1UserActionListParamsAPIGroupLblistener                  CloudV1UserActionListParamsAPIGroup = "lblistener"
	CloudV1UserActionListParamsAPIGroupLbpool                      CloudV1UserActionListParamsAPIGroup = "lbpool"
	CloudV1UserActionListParamsAPIGroupLbpoolMember                CloudV1UserActionListParamsAPIGroup = "lbpool_member"
	CloudV1UserActionListParamsAPIGroupLifecyclePolicy             CloudV1UserActionListParamsAPIGroup = "lifecycle_policy"
	CloudV1UserActionListParamsAPIGroupLifecyclePolicyVolumeMember CloudV1UserActionListParamsAPIGroup = "lifecycle_policy_volume_member"
	CloudV1UserActionListParamsAPIGroupLoadbalancer                CloudV1UserActionListParamsAPIGroup = "loadbalancer"
	CloudV1UserActionListParamsAPIGroupNetwork                     CloudV1UserActionListParamsAPIGroup = "network"
	CloudV1UserActionListParamsAPIGroupPort                        CloudV1UserActionListParamsAPIGroup = "port"
	CloudV1UserActionListParamsAPIGroupProject                     CloudV1UserActionListParamsAPIGroup = "project"
	CloudV1UserActionListParamsAPIGroupQuotaLimitRequest           CloudV1UserActionListParamsAPIGroup = "quota_limit_request"
	CloudV1UserActionListParamsAPIGroupRegistry                    CloudV1UserActionListParamsAPIGroup = "registry"
	CloudV1UserActionListParamsAPIGroupReservation                 CloudV1UserActionListParamsAPIGroup = "reservation"
	CloudV1UserActionListParamsAPIGroupReservedFixedIP             CloudV1UserActionListParamsAPIGroup = "reserved_fixed_ip"
	CloudV1UserActionListParamsAPIGroupRole                        CloudV1UserActionListParamsAPIGroup = "role"
	CloudV1UserActionListParamsAPIGroupRouter                      CloudV1UserActionListParamsAPIGroup = "router"
	CloudV1UserActionListParamsAPIGroupSecret                      CloudV1UserActionListParamsAPIGroup = "secret"
	CloudV1UserActionListParamsAPIGroupSecuritygroup               CloudV1UserActionListParamsAPIGroup = "securitygroup"
	CloudV1UserActionListParamsAPIGroupSecuritygrouprule           CloudV1UserActionListParamsAPIGroup = "securitygrouprule"
	CloudV1UserActionListParamsAPIGroupServergroup                 CloudV1UserActionListParamsAPIGroup = "servergroup"
	CloudV1UserActionListParamsAPIGroupSharedFlavor                CloudV1UserActionListParamsAPIGroup = "shared_flavor"
	CloudV1UserActionListParamsAPIGroupSharedImage                 CloudV1UserActionListParamsAPIGroup = "shared_image"
	CloudV1UserActionListParamsAPIGroupSharedNetwork               CloudV1UserActionListParamsAPIGroup = "shared_network"
	CloudV1UserActionListParamsAPIGroupSnapshot                    CloudV1UserActionListParamsAPIGroup = "snapshot"
	CloudV1UserActionListParamsAPIGroupSnapshotSchedule            CloudV1UserActionListParamsAPIGroup = "snapshot_schedule"
	CloudV1UserActionListParamsAPIGroupSSHKey                      CloudV1UserActionListParamsAPIGroup = "ssh_key"
	CloudV1UserActionListParamsAPIGroupSubnet                      CloudV1UserActionListParamsAPIGroup = "subnet"
	CloudV1UserActionListParamsAPIGroupUser                        CloudV1UserActionListParamsAPIGroup = "user"
	CloudV1UserActionListParamsAPIGroupVipIPAddresses              CloudV1UserActionListParamsAPIGroup = "vip_ip_addresses"
	CloudV1UserActionListParamsAPIGroupVolume                      CloudV1UserActionListParamsAPIGroup = "volume"
)

func (r CloudV1UserActionListParamsAPIGroup) IsKnown() bool {
	switch r {
	case CloudV1UserActionListParamsAPIGroupAICluster, CloudV1UserActionListParamsAPIGroupBackupService, CloudV1UserActionListParamsAPIGroupCaasContainer, CloudV1UserActionListParamsAPIGroupCaasKey, CloudV1UserActionListParamsAPIGroupCaasPullSecret, CloudV1UserActionListParamsAPIGroupDbaasPostgres, CloudV1UserActionListParamsAPIGroupDdosProfile, CloudV1UserActionListParamsAPIGroupFaasFunction, CloudV1UserActionListParamsAPIGroupFaasKey, CloudV1UserActionListParamsAPIGroupFaasNamespace, CloudV1UserActionListParamsAPIGroupFileShares, CloudV1UserActionListParamsAPIGroupFloatingIP, CloudV1UserActionListParamsAPIGroupHeat, CloudV1UserActionListParamsAPIGroupImage, CloudV1UserActionListParamsAPIGroupInferenceAtTheEdge, CloudV1UserActionListParamsAPIGroupInstance, CloudV1UserActionListParamsAPIGroupInstanceIsolation, CloudV1UserActionListParamsAPIGroupK8sCluster, CloudV1UserActionListParamsAPIGroupK8sClusterTemplate, CloudV1UserActionListParamsAPIGroupK8sPool, CloudV1UserActionListParamsAPIGroupLaas, CloudV1UserActionListParamsAPIGroupLaasTopic, CloudV1UserActionListParamsAPIGroupLbHealthMonitor, CloudV1UserActionListParamsAPIGroupLbL7policy, CloudV1UserActionListParamsAPIGroupLbL7rule, CloudV1UserActionListParamsAPIGroupLblistener, CloudV1UserActionListParamsAPIGroupLbpool, CloudV1UserActionListParamsAPIGroupLbpoolMember, CloudV1UserActionListParamsAPIGroupLifecyclePolicy, CloudV1UserActionListParamsAPIGroupLifecyclePolicyVolumeMember, CloudV1UserActionListParamsAPIGroupLoadbalancer, CloudV1UserActionListParamsAPIGroupNetwork, CloudV1UserActionListParamsAPIGroupPort, CloudV1UserActionListParamsAPIGroupProject, CloudV1UserActionListParamsAPIGroupQuotaLimitRequest, CloudV1UserActionListParamsAPIGroupRegistry, CloudV1UserActionListParamsAPIGroupReservation, CloudV1UserActionListParamsAPIGroupReservedFixedIP, CloudV1UserActionListParamsAPIGroupRole, CloudV1UserActionListParamsAPIGroupRouter, CloudV1UserActionListParamsAPIGroupSecret, CloudV1UserActionListParamsAPIGroupSecuritygroup, CloudV1UserActionListParamsAPIGroupSecuritygrouprule, CloudV1UserActionListParamsAPIGroupServergroup, CloudV1UserActionListParamsAPIGroupSharedFlavor, CloudV1UserActionListParamsAPIGroupSharedImage, CloudV1UserActionListParamsAPIGroupSharedNetwork, CloudV1UserActionListParamsAPIGroupSnapshot, CloudV1UserActionListParamsAPIGroupSnapshotSchedule, CloudV1UserActionListParamsAPIGroupSSHKey, CloudV1UserActionListParamsAPIGroupSubnet, CloudV1UserActionListParamsAPIGroupUser, CloudV1UserActionListParamsAPIGroupVipIPAddresses, CloudV1UserActionListParamsAPIGroupVolume:
		return true
	}
	return false
}

// Sorting by date. Oldest first, or most recent first
type CloudV1UserActionListParamsSorting string

const (
	CloudV1UserActionListParamsSortingAsc  CloudV1UserActionListParamsSorting = "asc"
	CloudV1UserActionListParamsSortingDesc CloudV1UserActionListParamsSorting = "desc"
)

func (r CloudV1UserActionListParamsSorting) IsKnown() bool {
	switch r {
	case CloudV1UserActionListParamsSortingAsc, CloudV1UserActionListParamsSortingDesc:
		return true
	}
	return false
}

type CloudV1UserActionGetSubscriptionsListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1UserActionGetSubscriptionsListParams]'s query
// parameters as `url.Values`.
func (r CloudV1UserActionGetSubscriptionsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1UserActionSubscribeParams struct {
	// Authorization header name
	AuthHeaderName param.Field[string] `json:"auth_header_name,required"`
	// Authorization header value
	AuthHeaderValue param.Field[string] `json:"auth_header_value,required"`
	// URL to send user action logs for the specified client
	URL param.Field[string] `json:"url,required"`
}

func (r CloudV1UserActionSubscribeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1UserActionSubscribeAmqpParams struct {
	// Connection string of the following structure
	// "scheme://username:password@host:port/virtual_host"
	ConnectionString param.Field[string] `json:"connection_string,required"`
	// Exchange name
	Exchange param.Field[string] `json:"exchange"`
	// Set to true if you would like to receive user action logs of all clients with
	// the reseller_id matching the current client_id. Defaults to false
	ReceiveChildClientEvents param.Field[bool] `json:"receive_child_client_events"`
	// Routing key
	RoutingKey param.Field[string] `json:"routing_key"`
}

func (r CloudV1UserActionSubscribeAmqpParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
