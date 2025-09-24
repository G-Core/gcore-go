// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// AuditLogService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.Options = opts
	return
}

// Retrieve user action log for one client or a set of projects
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AuditLogEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cloud/v1/user_actions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieve user action log for one client or a set of projects
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AuditLogEntry] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type AuditLogEntry struct {
	// User action ID
	ID string `json:"id,required"`
	// User action log was successfully received by its subscriber in case there is one
	Acknowledged bool `json:"acknowledged,required"`
	// Additional information about the action
	ActionData any `json:"action_data,required"`
	// Action type
	//
	// Any of "activate", "attach", "change_logging_resources", "create",
	// "create_access_rule", "deactivate", "delete", "delete_access_rule",
	// "delete_metadata", "detach", "disable_logging", "disable_portsecurity",
	// "download", "enable_logging", "enable_portsecurity", "failover",
	// "put_into_servergroup", "reboot", "reboot_hard", "rebuild",
	// "regenerate_credentials", "remove_from_servergroup", "replace_metadata",
	// "resize", "resume", "retype", "revert", "scale_down", "scale_up", "start",
	// "stop", "suspend", "update", "update_metadata", "upgrade".
	ActionType AuditLogEntryActionType `json:"action_type,required"`
	// API group
	//
	// Any of "ai_cluster", "backup_service", "caas_container", "caas_key",
	// "caas_pull_secret", "dbaas_postgres", "ddos_profile", "faas_function",
	// "faas_key", "faas_namespace", "file_shares", "floating_ip", "image",
	// "inference_at_the_edge", "instance", "instance_isolation", "k8s_cluster",
	// "k8s_cluster_template", "k8s_pool", "laas", "laas_topic", "lb_health_monitor",
	// "lb_l7policy", "lb_l7rule", "lblistener", "lbpool", "lbpool_member",
	// "lifecycle_policy", "lifecycle_policy_volume_member", "loadbalancer", "network",
	// "port", "project", "quota_limit_request", "registry", "reservation",
	// "reserved_fixed_ip", "role", "router", "secret", "securitygroup",
	// "securitygrouprule", "servergroup", "shared_flavor", "shared_image",
	// "shared_network", "snapshot", "snapshot_schedule", "ssh_key", "subnet", "user",
	// "vip_ip_addresses", "volume".
	APIGroup AuditLogEntryAPIGroup `json:"api_group,required"`
	// Client ID of the user.
	ClientID int64 `json:"client_id,required"`
	// User email address
	Email string `json:"email,required" format:"email"`
	// User action was filled with all necessary information. If false, then something
	// went wrong during user action creation or update
	IsComplete bool `json:"is_complete,required"`
	// User ID who issued the token that made the request
	IssuedByUserID int64 `json:"issued_by_user_id,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Resources
	Resources []AuditLogEntryResource `json:"resources,required"`
	// Task ID
	TaskID string `json:"task_id,required"`
	// Datetime. Action timestamp
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Token ID
	TokenID int64 `json:"token_id,required"`
	// Total resource price VAT inclusive
	TotalPrice AuditLogEntryTotalPrice `json:"total_price,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Acknowledged   respjson.Field
		ActionData     respjson.Field
		ActionType     respjson.Field
		APIGroup       respjson.Field
		ClientID       respjson.Field
		Email          respjson.Field
		IsComplete     respjson.Field
		IssuedByUserID respjson.Field
		ProjectID      respjson.Field
		RegionID       respjson.Field
		Resources      respjson.Field
		TaskID         respjson.Field
		Timestamp      respjson.Field
		TokenID        respjson.Field
		TotalPrice     respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogEntry) RawJSON() string { return r.JSON.raw }
func (r *AuditLogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action type
type AuditLogEntryActionType string

const (
	AuditLogEntryActionTypeActivate               AuditLogEntryActionType = "activate"
	AuditLogEntryActionTypeAttach                 AuditLogEntryActionType = "attach"
	AuditLogEntryActionTypeChangeLoggingResources AuditLogEntryActionType = "change_logging_resources"
	AuditLogEntryActionTypeCreate                 AuditLogEntryActionType = "create"
	AuditLogEntryActionTypeCreateAccessRule       AuditLogEntryActionType = "create_access_rule"
	AuditLogEntryActionTypeDeactivate             AuditLogEntryActionType = "deactivate"
	AuditLogEntryActionTypeDelete                 AuditLogEntryActionType = "delete"
	AuditLogEntryActionTypeDeleteAccessRule       AuditLogEntryActionType = "delete_access_rule"
	AuditLogEntryActionTypeDeleteMetadata         AuditLogEntryActionType = "delete_metadata"
	AuditLogEntryActionTypeDetach                 AuditLogEntryActionType = "detach"
	AuditLogEntryActionTypeDisableLogging         AuditLogEntryActionType = "disable_logging"
	AuditLogEntryActionTypeDisablePortsecurity    AuditLogEntryActionType = "disable_portsecurity"
	AuditLogEntryActionTypeDownload               AuditLogEntryActionType = "download"
	AuditLogEntryActionTypeEnableLogging          AuditLogEntryActionType = "enable_logging"
	AuditLogEntryActionTypeEnablePortsecurity     AuditLogEntryActionType = "enable_portsecurity"
	AuditLogEntryActionTypeFailover               AuditLogEntryActionType = "failover"
	AuditLogEntryActionTypePutIntoServergroup     AuditLogEntryActionType = "put_into_servergroup"
	AuditLogEntryActionTypeReboot                 AuditLogEntryActionType = "reboot"
	AuditLogEntryActionTypeRebootHard             AuditLogEntryActionType = "reboot_hard"
	AuditLogEntryActionTypeRebuild                AuditLogEntryActionType = "rebuild"
	AuditLogEntryActionTypeRegenerateCredentials  AuditLogEntryActionType = "regenerate_credentials"
	AuditLogEntryActionTypeRemoveFromServergroup  AuditLogEntryActionType = "remove_from_servergroup"
	AuditLogEntryActionTypeReplaceMetadata        AuditLogEntryActionType = "replace_metadata"
	AuditLogEntryActionTypeResize                 AuditLogEntryActionType = "resize"
	AuditLogEntryActionTypeResume                 AuditLogEntryActionType = "resume"
	AuditLogEntryActionTypeRetype                 AuditLogEntryActionType = "retype"
	AuditLogEntryActionTypeRevert                 AuditLogEntryActionType = "revert"
	AuditLogEntryActionTypeScaleDown              AuditLogEntryActionType = "scale_down"
	AuditLogEntryActionTypeScaleUp                AuditLogEntryActionType = "scale_up"
	AuditLogEntryActionTypeStart                  AuditLogEntryActionType = "start"
	AuditLogEntryActionTypeStop                   AuditLogEntryActionType = "stop"
	AuditLogEntryActionTypeSuspend                AuditLogEntryActionType = "suspend"
	AuditLogEntryActionTypeUpdate                 AuditLogEntryActionType = "update"
	AuditLogEntryActionTypeUpdateMetadata         AuditLogEntryActionType = "update_metadata"
	AuditLogEntryActionTypeUpgrade                AuditLogEntryActionType = "upgrade"
)

// API group
type AuditLogEntryAPIGroup string

const (
	AuditLogEntryAPIGroupAICluster                   AuditLogEntryAPIGroup = "ai_cluster"
	AuditLogEntryAPIGroupBackupService               AuditLogEntryAPIGroup = "backup_service"
	AuditLogEntryAPIGroupCaasContainer               AuditLogEntryAPIGroup = "caas_container"
	AuditLogEntryAPIGroupCaasKey                     AuditLogEntryAPIGroup = "caas_key"
	AuditLogEntryAPIGroupCaasPullSecret              AuditLogEntryAPIGroup = "caas_pull_secret"
	AuditLogEntryAPIGroupDbaasPostgres               AuditLogEntryAPIGroup = "dbaas_postgres"
	AuditLogEntryAPIGroupDDOSProfile                 AuditLogEntryAPIGroup = "ddos_profile"
	AuditLogEntryAPIGroupFaasFunction                AuditLogEntryAPIGroup = "faas_function"
	AuditLogEntryAPIGroupFaasKey                     AuditLogEntryAPIGroup = "faas_key"
	AuditLogEntryAPIGroupFaasNamespace               AuditLogEntryAPIGroup = "faas_namespace"
	AuditLogEntryAPIGroupFileShares                  AuditLogEntryAPIGroup = "file_shares"
	AuditLogEntryAPIGroupFloatingIP                  AuditLogEntryAPIGroup = "floating_ip"
	AuditLogEntryAPIGroupImage                       AuditLogEntryAPIGroup = "image"
	AuditLogEntryAPIGroupInferenceAtTheEdge          AuditLogEntryAPIGroup = "inference_at_the_edge"
	AuditLogEntryAPIGroupInstance                    AuditLogEntryAPIGroup = "instance"
	AuditLogEntryAPIGroupInstanceIsolation           AuditLogEntryAPIGroup = "instance_isolation"
	AuditLogEntryAPIGroupK8sCluster                  AuditLogEntryAPIGroup = "k8s_cluster"
	AuditLogEntryAPIGroupK8sClusterTemplate          AuditLogEntryAPIGroup = "k8s_cluster_template"
	AuditLogEntryAPIGroupK8sPool                     AuditLogEntryAPIGroup = "k8s_pool"
	AuditLogEntryAPIGroupLaas                        AuditLogEntryAPIGroup = "laas"
	AuditLogEntryAPIGroupLaasTopic                   AuditLogEntryAPIGroup = "laas_topic"
	AuditLogEntryAPIGroupLbHealthMonitor             AuditLogEntryAPIGroup = "lb_health_monitor"
	AuditLogEntryAPIGroupLbL7policy                  AuditLogEntryAPIGroup = "lb_l7policy"
	AuditLogEntryAPIGroupLbL7rule                    AuditLogEntryAPIGroup = "lb_l7rule"
	AuditLogEntryAPIGroupLblistener                  AuditLogEntryAPIGroup = "lblistener"
	AuditLogEntryAPIGroupLbpool                      AuditLogEntryAPIGroup = "lbpool"
	AuditLogEntryAPIGroupLbpoolMember                AuditLogEntryAPIGroup = "lbpool_member"
	AuditLogEntryAPIGroupLifecyclePolicy             AuditLogEntryAPIGroup = "lifecycle_policy"
	AuditLogEntryAPIGroupLifecyclePolicyVolumeMember AuditLogEntryAPIGroup = "lifecycle_policy_volume_member"
	AuditLogEntryAPIGroupLoadbalancer                AuditLogEntryAPIGroup = "loadbalancer"
	AuditLogEntryAPIGroupNetwork                     AuditLogEntryAPIGroup = "network"
	AuditLogEntryAPIGroupPort                        AuditLogEntryAPIGroup = "port"
	AuditLogEntryAPIGroupProject                     AuditLogEntryAPIGroup = "project"
	AuditLogEntryAPIGroupQuotaLimitRequest           AuditLogEntryAPIGroup = "quota_limit_request"
	AuditLogEntryAPIGroupRegistry                    AuditLogEntryAPIGroup = "registry"
	AuditLogEntryAPIGroupReservation                 AuditLogEntryAPIGroup = "reservation"
	AuditLogEntryAPIGroupReservedFixedIP             AuditLogEntryAPIGroup = "reserved_fixed_ip"
	AuditLogEntryAPIGroupRole                        AuditLogEntryAPIGroup = "role"
	AuditLogEntryAPIGroupRouter                      AuditLogEntryAPIGroup = "router"
	AuditLogEntryAPIGroupSecret                      AuditLogEntryAPIGroup = "secret"
	AuditLogEntryAPIGroupSecuritygroup               AuditLogEntryAPIGroup = "securitygroup"
	AuditLogEntryAPIGroupSecuritygrouprule           AuditLogEntryAPIGroup = "securitygrouprule"
	AuditLogEntryAPIGroupServergroup                 AuditLogEntryAPIGroup = "servergroup"
	AuditLogEntryAPIGroupSharedFlavor                AuditLogEntryAPIGroup = "shared_flavor"
	AuditLogEntryAPIGroupSharedImage                 AuditLogEntryAPIGroup = "shared_image"
	AuditLogEntryAPIGroupSharedNetwork               AuditLogEntryAPIGroup = "shared_network"
	AuditLogEntryAPIGroupSnapshot                    AuditLogEntryAPIGroup = "snapshot"
	AuditLogEntryAPIGroupSnapshotSchedule            AuditLogEntryAPIGroup = "snapshot_schedule"
	AuditLogEntryAPIGroupSSHKey                      AuditLogEntryAPIGroup = "ssh_key"
	AuditLogEntryAPIGroupSubnet                      AuditLogEntryAPIGroup = "subnet"
	AuditLogEntryAPIGroupUser                        AuditLogEntryAPIGroup = "user"
	AuditLogEntryAPIGroupVipIPAddresses              AuditLogEntryAPIGroup = "vip_ip_addresses"
	AuditLogEntryAPIGroupVolume                      AuditLogEntryAPIGroup = "volume"
)

type AuditLogEntryResource struct {
	// Resource ID
	ResourceID string `json:"resource_id,required"`
	// Resource type
	//
	// Any of "caas_container", "caas_key", "caas_pull_secret", "dbaas_postgres",
	// "ddos_profile", "external_ip", "faas_function", "faas_key", "faas_namespace",
	// "file_shares", "floating_ip", "gpu_baremetal_server", "gpu_virtual_server",
	// "gpuai_cluster", "image", "inference_api_key", "inference_application",
	// "inference_instance", "inference_registry_credentials", "inference_secret",
	// "instance", "ipu_cluster", "k8s_cluster", "k8s_cluster_template", "k8s_pool",
	// "laas", "laas_topic", "lb_health_monitor", "lb_l7policy", "lb_l7rule",
	// "lblistener", "lbpool", "lbpool_member", "lifecycle_policy",
	// "lifecycle_policy_volume_member", "loadbalancer", "network", "port", "project",
	// "quota_limit_request", "registry", "registry_repository",
	// "registry_repository_artifact", "registry_repository_tag", "registry_user",
	// "registry_user_sercret", "reservation", "role", "router", "secret",
	// "securitygroup", "securitygrouprule", "servergroup", "shared_flavor",
	// "shared_image", "shared_network", "snapshot", "snapshot_schedule", "ssh_key",
	// "subnet", "token", "user", "virtual_gpu_cluster", "volume".
	ResourceType string `json:"resource_type,required"`
	// Free-form object, resource body.
	ResourceBody any `json:"resource_body,nullable"`
	// Often used property for filtering actions. It can be a name, IP address, or
	// other property, depending on the `resource_type`
	SearchField string `json:"search_field,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceType respjson.Field
		ResourceBody respjson.Field
		SearchField  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogEntryResource) RawJSON() string { return r.JSON.raw }
func (r *AuditLogEntryResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Total resource price VAT inclusive
type AuditLogEntryTotalPrice struct {
	// Currency code (3 letter code per ISO 4217)
	CurrencyCode string `json:"currency_code,required"`
	// Total price VAT inclusive per hour
	PricePerHour float64 `json:"price_per_hour,required"`
	// Total price VAT inclusive per month (30 days)
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
func (r AuditLogEntryTotalPrice) RawJSON() string { return r.JSON.raw }
func (r *AuditLogEntryTotalPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListParams struct {
	// ISO formatted datetime string. Starting timestamp from which user actions are
	// requested
	FromTimestamp param.Opt[time.Time] `query:"from_timestamp,omitzero" format:"date-time" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Extra search field for common object properties such as name, IP address, or
	// other, depending on the `resource_type`
	SearchField param.Opt[string] `query:"search_field,omitzero" json:"-"`
	// ISO formatted datetime string. Ending timestamp until which user actions are
	// requested
	ToTimestamp param.Opt[time.Time] `query:"to_timestamp,omitzero" format:"date-time" json:"-"`
	// User action type. Several options can be specified.
	//
	// Any of "activate", "attach", "change_logging_resources", "create",
	// "create_access_rule", "deactivate", "delete", "delete_access_rule",
	// "delete_metadata", "detach", "disable_logging", "disable_portsecurity",
	// "download", "enable_logging", "enable_portsecurity", "failover",
	// "put_into_servergroup", "reboot", "reboot_hard", "rebuild",
	// "regenerate_credentials", "remove_from_servergroup", "replace_metadata",
	// "resize", "resume", "retype", "revert", "scale_down", "scale_up", "start",
	// "stop", "suspend", "update", "update_metadata", "upgrade".
	ActionType []string `query:"action_type,omitzero" json:"-"`
	// API group that requested action belongs to. Several options can be specified.
	//
	// Any of "ai_cluster", "backup_service", "caas_container", "caas_key",
	// "caas_pull_secret", "dbaas_postgres", "ddos_profile", "faas_function",
	// "faas_key", "faas_namespace", "file_shares", "floating_ip", "image",
	// "inference_at_the_edge", "instance", "instance_isolation", "k8s_cluster",
	// "k8s_cluster_template", "k8s_pool", "laas", "laas_topic", "lb_health_monitor",
	// "lb_l7policy", "lb_l7rule", "lblistener", "lbpool", "lbpool_member",
	// "lifecycle_policy", "lifecycle_policy_volume_member", "loadbalancer", "network",
	// "port", "project", "quota_limit_request", "registry", "reservation",
	// "reserved_fixed_ip", "role", "router", "secret", "securitygroup",
	// "securitygrouprule", "servergroup", "shared_flavor", "shared_image",
	// "shared_network", "snapshot", "snapshot_schedule", "ssh_key", "subnet", "user",
	// "vip_ip_addresses", "volume".
	APIGroup []string `query:"api_group,omitzero" json:"-"`
	// Sorting by timestamp. Oldest first, or most recent first
	//
	// Any of "asc", "desc".
	OrderBy AuditLogListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Project ID. Several options can be specified.
	ProjectID []int64 `query:"project_id,omitzero" json:"-"`
	// Region ID. Several options can be specified.
	RegionID []int64 `query:"region_id,omitzero" json:"-"`
	// Resource ID. Several options can be specified.
	ResourceID []string `query:"resource_id,omitzero" json:"-"`
	// (DEPRECATED Use '`order_by`' instead) Sorting by timestamp. Oldest first, or
	// most recent first
	//
	// Any of "asc", "desc".
	Sorting AuditLogListParamsSorting `query:"sorting,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sorting by timestamp. Oldest first, or most recent first
type AuditLogListParamsOrderBy string

const (
	AuditLogListParamsOrderByAsc  AuditLogListParamsOrderBy = "asc"
	AuditLogListParamsOrderByDesc AuditLogListParamsOrderBy = "desc"
)

// (DEPRECATED Use '`order_by`' instead) Sorting by timestamp. Oldest first, or
// most recent first
type AuditLogListParamsSorting string

const (
	AuditLogListParamsSortingAsc  AuditLogListParamsSorting = "asc"
	AuditLogListParamsSortingDesc AuditLogListParamsSorting = "desc"
)
