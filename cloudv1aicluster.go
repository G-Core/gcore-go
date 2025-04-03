// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1AIClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIClusterService] method instead.
type CloudV1AIClusterService struct {
	Options      []option.RequestOption
	MetadataItem *CloudV1AIClusterMetadataItemService
	GPU          *CloudV1AIClusterGPUService
	Metadata     *CloudV1AIClusterMetadataService
}

// NewCloudV1AIClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1AIClusterService(opts ...option.RequestOption) (r *CloudV1AIClusterService) {
	r = &CloudV1AIClusterService{}
	r.Options = opts
	r.MetadataItem = NewCloudV1AIClusterMetadataItemService(opts...)
	r.GPU = NewCloudV1AIClusterGPUService(opts...)
	r.Metadata = NewCloudV1AIClusterMetadataService(opts...)
	return
}

// Assign the security group to the AI cluster. To assign multiple security groups
// to all ports, use the NULL value for the port_id field
func (r *CloudV1AIClusterService) AddSecurityGroup(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV1AIClusterAddSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/addsecuritygroup", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Attach interface to AI cluster node
func (r *CloudV1AIClusterService) AttachInterface(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1AIClusterAttachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/attach_interface", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified AI cluster will be calculated
func (r *CloudV1AIClusterService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1AIClusterCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach interface from AI cluster node
func (r *CloudV1AIClusterService) DetachInterface(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1AIClusterDetachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/detach_interface", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get AI cluster node console URL
func (r *CloudV1AIClusterService) GetConsole(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *RemoteConsole, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/get_console", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network interfaces attached to AI cluster poplar servers
func (r *CloudV1AIClusterService) ListInterfaces(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *PortList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/interfaces", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network ports of AI cluster poplar servers
func (r *CloudV1AIClusterService) ListPorts(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *PortSecurityGroupsList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/ports", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Powercycle (stop and start) one AI cluster node, aka hard reboot
func (r *CloudV1AIClusterService) Powercycle(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/powercycle", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot one AI cluster node
func (r *CloudV1AIClusterService) Reboot(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/reboot", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Un-assign the security group from the AI cluster. To un-assign multiple security
// groups to all ports, use the NULL value for the port_id field
func (r *CloudV1AIClusterService) RemoveSecurityGroup(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV1AIClusterRemoveSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/delsecuritygroup", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get DDoS Protection service client profile
type DeprecatedGetClientProfile struct {
	// DDoS protection profile ID
	ID     int64                             `json:"id"`
	Fields []DeprecatedGetClientProfileField `json:"fields"`
	// Activate or deactivate exist DDoS protection profile.
	Options DeprecatedGetClientProfileOptions `json:"options"`
	// DDoS Protection service profile template
	ProfileTemplate DeprecatedGetClientProfileProfileTemplate `json:"profile_template"`
	// DDoS profile template description
	ProfileTemplateDescription string `json:"profile_template_description,nullable"`
	// List of protocols
	Protocols []interface{}                    `json:"protocols"`
	Site      string                           `json:"site,nullable"`
	Status    DeprecatedGetClientProfileStatus `json:"status"`
	JSON      deprecatedGetClientProfileJSON   `json:"-"`
}

// deprecatedGetClientProfileJSON contains the JSON metadata for the struct
// [DeprecatedGetClientProfile]
type deprecatedGetClientProfileJSON struct {
	ID                         apijson.Field
	Fields                     apijson.Field
	Options                    apijson.Field
	ProfileTemplate            apijson.Field
	ProfileTemplateDescription apijson.Field
	Protocols                  apijson.Field
	Site                       apijson.Field
	Status                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DeprecatedGetClientProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileJSON) RawJSON() string {
	return r.raw
}

type DeprecatedGetClientProfileField struct {
	ID        int64 `json:"id,required"`
	BaseField int64 `json:"base_field,required"`
	// Field name
	Name    string `json:"name,required"`
	Default string `json:"default,nullable"`
	// Field description
	Description string `json:"description"`
	// Field type
	FieldType DeprecatedGetClientProfileFieldsFieldType `json:"field_type,nullable"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue interface{} `json:"field_value"`
	Required   bool        `json:"required"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value string                              `json:"value,nullable"`
	JSON  deprecatedGetClientProfileFieldJSON `json:"-"`
}

// deprecatedGetClientProfileFieldJSON contains the JSON metadata for the struct
// [DeprecatedGetClientProfileField]
type deprecatedGetClientProfileFieldJSON struct {
	ID          apijson.Field
	BaseField   apijson.Field
	Name        apijson.Field
	Default     apijson.Field
	Description apijson.Field
	FieldType   apijson.Field
	FieldValue  apijson.Field
	Required    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedGetClientProfileField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileFieldJSON) RawJSON() string {
	return r.raw
}

// Field type
type DeprecatedGetClientProfileFieldsFieldType string

const (
	DeprecatedGetClientProfileFieldsFieldTypeBool DeprecatedGetClientProfileFieldsFieldType = "bool"
	DeprecatedGetClientProfileFieldsFieldTypeInt  DeprecatedGetClientProfileFieldsFieldType = "int"
	DeprecatedGetClientProfileFieldsFieldTypeStr  DeprecatedGetClientProfileFieldsFieldType = "str"
)

func (r DeprecatedGetClientProfileFieldsFieldType) IsKnown() bool {
	switch r {
	case DeprecatedGetClientProfileFieldsFieldTypeBool, DeprecatedGetClientProfileFieldsFieldTypeInt, DeprecatedGetClientProfileFieldsFieldTypeStr:
		return true
	}
	return false
}

// Activate or deactivate exist DDoS protection profile.
type DeprecatedGetClientProfileOptions struct {
	// Activate profile.
	Active bool `json:"active"`
	// Activate BGP protocol
	Bgp  bool                                  `json:"bgp"`
	JSON deprecatedGetClientProfileOptionsJSON `json:"-"`
}

// deprecatedGetClientProfileOptionsJSON contains the JSON metadata for the struct
// [DeprecatedGetClientProfileOptions]
type deprecatedGetClientProfileOptionsJSON struct {
	Active      apijson.Field
	Bgp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedGetClientProfileOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileOptionsJSON) RawJSON() string {
	return r.raw
}

// DDoS Protection service profile template
type DeprecatedGetClientProfileProfileTemplate struct {
	ID int64 `json:"id,required"`
	// Additional fields
	Fields []DeprecatedGetClientProfileProfileTemplateField `json:"fields,required"`
	// Template name
	Name string `json:"name,required"`
	// Template description
	Description string                                        `json:"description"`
	JSON        deprecatedGetClientProfileProfileTemplateJSON `json:"-"`
}

// deprecatedGetClientProfileProfileTemplateJSON contains the JSON metadata for the
// struct [DeprecatedGetClientProfileProfileTemplate]
type deprecatedGetClientProfileProfileTemplateJSON struct {
	ID          apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedGetClientProfileProfileTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileProfileTemplateJSON) RawJSON() string {
	return r.raw
}

type DeprecatedGetClientProfileProfileTemplateField struct {
	ID int64 `json:"id,required"`
	// Field name
	Name    string `json:"name,required"`
	Default string `json:"default,nullable"`
	// Field description
	Description string `json:"description"`
	// Field type
	FieldType DeprecatedGetClientProfileProfileTemplateFieldsFieldType `json:"field_type,nullable"`
	Required  bool                                                     `json:"required"`
	// Json schema to validate field_values
	ValidationSchema interface{}                                        `json:"validation_schema"`
	JSON             deprecatedGetClientProfileProfileTemplateFieldJSON `json:"-"`
}

// deprecatedGetClientProfileProfileTemplateFieldJSON contains the JSON metadata
// for the struct [DeprecatedGetClientProfileProfileTemplateField]
type deprecatedGetClientProfileProfileTemplateFieldJSON struct {
	ID               apijson.Field
	Name             apijson.Field
	Default          apijson.Field
	Description      apijson.Field
	FieldType        apijson.Field
	Required         apijson.Field
	ValidationSchema apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeprecatedGetClientProfileProfileTemplateField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileProfileTemplateFieldJSON) RawJSON() string {
	return r.raw
}

// Field type
type DeprecatedGetClientProfileProfileTemplateFieldsFieldType string

const (
	DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeBool DeprecatedGetClientProfileProfileTemplateFieldsFieldType = "bool"
	DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeInt  DeprecatedGetClientProfileProfileTemplateFieldsFieldType = "int"
	DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeStr  DeprecatedGetClientProfileProfileTemplateFieldsFieldType = "str"
)

func (r DeprecatedGetClientProfileProfileTemplateFieldsFieldType) IsKnown() bool {
	switch r {
	case DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeBool, DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeInt, DeprecatedGetClientProfileProfileTemplateFieldsFieldTypeStr:
		return true
	}
	return false
}

type DeprecatedGetClientProfileStatus struct {
	// Description of the error, if it exists
	ErrorDescription string `json:"error_description"`
	// Profile status
	Status string                               `json:"status"`
	JSON   deprecatedGetClientProfileStatusJSON `json:"-"`
}

// deprecatedGetClientProfileStatusJSON contains the JSON metadata for the struct
// [DeprecatedGetClientProfileStatus]
type deprecatedGetClientProfileStatusJSON struct {
	ErrorDescription apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeprecatedGetClientProfileStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedGetClientProfileStatusJSON) RawJSON() string {
	return r.raw
}

// Instance schema
type DeprecatedInstance struct {
	// Map of network_name → [AddressSchema]
	Addresses map[string][]DeprecatedInstanceAddress `json:"addresses"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []DeprecatedInstanceBlackholePort `json:"blackhole_ports"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	// Get DDoS Protection service client profile
	DdosProfile        DeprecatedGetClientProfile            `json:"ddos_profile"`
	FixedIPAssignments []DeprecatedInstanceFixedIPAssignment `json:"fixed_ip_assignments"`
	// Flavor
	Flavor   DeprecatedInstanceFlavor `json:"flavor"`
	FlavorID string                   `json:"flavor_id"`
	// Datetime when instance was created
	InstanceCreated interface{} `json:"instance_created"`
	// Instance description
	InstanceDescription string `json:"instance_description"`
	// Instance ID
	InstanceID string `json:"instance_id"`
	// Instance isolation information
	InstanceIsolation DeprecatedInstanceInstanceIsolation `json:"instance_isolation"`
	// Instance name
	InstanceName string `json:"instance_name"`
	KeypairName  string `json:"keypair_name,nullable"`
	// VM metadata
	Metadata interface{} `json:"metadata"`
	// Detailed VM metadata
	MetadataDetailed interface{} `json:"metadata_detailed"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID       int64            `json:"region_id"`
	SecurityGroups []DeprecatedName `json:"security_groups"`
	// VM status
	Status string `json:"status"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Task state
	TaskState string `json:"task_state,nullable"`
	// Virtual machine state (active)
	VmState string `json:"vm_state"`
	// List of volumes
	Volumes []DeprecatedInstanceVolume `json:"volumes"`
	JSON    deprecatedInstanceJSON     `json:"-"`
}

// deprecatedInstanceJSON contains the JSON metadata for the struct
// [DeprecatedInstance]
type deprecatedInstanceJSON struct {
	Addresses           apijson.Field
	BlackholePorts      apijson.Field
	CreatorTaskID       apijson.Field
	DdosProfile         apijson.Field
	FixedIPAssignments  apijson.Field
	Flavor              apijson.Field
	FlavorID            apijson.Field
	InstanceCreated     apijson.Field
	InstanceDescription apijson.Field
	InstanceID          apijson.Field
	InstanceIsolation   apijson.Field
	InstanceName        apijson.Field
	KeypairName         apijson.Field
	Metadata            apijson.Field
	MetadataDetailed    apijson.Field
	ProjectID           apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	SecurityGroups      apijson.Field
	Status              apijson.Field
	TaskID              apijson.Field
	TaskState           apijson.Field
	VmState             apijson.Field
	Volumes             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeprecatedInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceJSON) RawJSON() string {
	return r.raw
}

// Address schema
type DeprecatedInstanceAddress struct {
	Addr string `json:"addr"`
	// Interface name
	InterfaceName string                        `json:"interface_name"`
	SubnetID      string                        `json:"subnet_id"`
	SubnetName    string                        `json:"subnet_name"`
	Type          string                        `json:"type"`
	JSON          deprecatedInstanceAddressJSON `json:"-"`
}

// deprecatedInstanceAddressJSON contains the JSON metadata for the struct
// [DeprecatedInstanceAddress]
type deprecatedInstanceAddressJSON struct {
	Addr          apijson.Field
	InterfaceName apijson.Field
	SubnetID      apijson.Field
	SubnetName    apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeprecatedInstanceAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceAddressJSON) RawJSON() string {
	return r.raw
}

// Blocking notification
type DeprecatedInstanceBlackholePort struct {
	// A date-time string giving the time that the alarm ended. If not yet ended, time
	// will be given as 0001-01-01T00:00:00Z
	AlarmEnd time.Time `json:"AlarmEnd,required" format:"date-time"`
	// A date-time string giving the time that the alarm started
	AlarmStart time.Time `json:"AlarmStart,required" format:"date-time"`
	// Current state of alarm
	AlarmState DeprecatedInstanceBlackholePortsAlarmState `json:"AlarmState,required"`
	// An array of the dimensions that make up the key used to evaluate traffic for
	// this alert policy
	AlertKey []interface{} `json:"AlertKey,required"`
	// The name of the alert policy generating this notification
	AlertPolicyName DeprecatedInstanceBlackholePortsAlertPolicyName `json:"AlertPolicyName,required"`
	// Total alert duration
	AlertDuration string `json:"AlertDuration"`
	// Notification destination IP address
	DestinationIP string                              `json:"DestinationIP"`
	ID            int64                               `json:"ID"`
	JSON          deprecatedInstanceBlackholePortJSON `json:"-"`
}

// deprecatedInstanceBlackholePortJSON contains the JSON metadata for the struct
// [DeprecatedInstanceBlackholePort]
type deprecatedInstanceBlackholePortJSON struct {
	AlarmEnd        apijson.Field
	AlarmStart      apijson.Field
	AlarmState      apijson.Field
	AlertKey        apijson.Field
	AlertPolicyName apijson.Field
	AlertDuration   apijson.Field
	DestinationIP   apijson.Field
	ID              apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeprecatedInstanceBlackholePort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceBlackholePortJSON) RawJSON() string {
	return r.raw
}

// Current state of alarm
type DeprecatedInstanceBlackholePortsAlarmState string

const (
	DeprecatedInstanceBlackholePortsAlarmStateAckReqUppercase             DeprecatedInstanceBlackholePortsAlarmState = "ACK_REQ"
	DeprecatedInstanceBlackholePortsAlarmStateAlarmUppercase              DeprecatedInstanceBlackholePortsAlarmState = "ALARM"
	DeprecatedInstanceBlackholePortsAlarmStateArchivedUppercase           DeprecatedInstanceBlackholePortsAlarmState = "ARCHIVED"
	DeprecatedInstanceBlackholePortsAlarmStateClearUppercase              DeprecatedInstanceBlackholePortsAlarmState = "CLEAR"
	DeprecatedInstanceBlackholePortsAlarmStateClearingUppercase           DeprecatedInstanceBlackholePortsAlarmState = "CLEARING"
	DeprecatedInstanceBlackholePortsAlarmStateClearingFailUppercase       DeprecatedInstanceBlackholePortsAlarmState = "CLEARING_FAIL"
	DeprecatedInstanceBlackholePortsAlarmStateEndGraceUppercase           DeprecatedInstanceBlackholePortsAlarmState = "END_GRACE"
	DeprecatedInstanceBlackholePortsAlarmStateEndWaitUppercase            DeprecatedInstanceBlackholePortsAlarmState = "END_WAIT"
	DeprecatedInstanceBlackholePortsAlarmStateManualClearUppercase        DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_CLEAR"
	DeprecatedInstanceBlackholePortsAlarmStateManualClearingUppercase     DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_CLEARING"
	DeprecatedInstanceBlackholePortsAlarmStateManualClearingFailUppercase DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_CLEARING_FAIL"
	DeprecatedInstanceBlackholePortsAlarmStateManualMitigatingUppercase   DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_MITIGATING"
	DeprecatedInstanceBlackholePortsAlarmStateManualStartingUppercase     DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_STARTING"
	DeprecatedInstanceBlackholePortsAlarmStateManualStartingFailUppercase DeprecatedInstanceBlackholePortsAlarmState = "MANUAL_STARTING_FAIL"
	DeprecatedInstanceBlackholePortsAlarmStateMitigatingUppercase         DeprecatedInstanceBlackholePortsAlarmState = "MITIGATING"
	DeprecatedInstanceBlackholePortsAlarmStateStartingUppercase           DeprecatedInstanceBlackholePortsAlarmState = "STARTING"
	DeprecatedInstanceBlackholePortsAlarmStateStartingFailUppercase       DeprecatedInstanceBlackholePortsAlarmState = "STARTING_FAIL"
	DeprecatedInstanceBlackholePortsAlarmStateStartWaitUppercase          DeprecatedInstanceBlackholePortsAlarmState = "START_WAIT"
	DeprecatedInstanceBlackholePortsAlarmStateAckReq                      DeprecatedInstanceBlackholePortsAlarmState = "ack_req"
	DeprecatedInstanceBlackholePortsAlarmStateAlarm                       DeprecatedInstanceBlackholePortsAlarmState = "alarm"
	DeprecatedInstanceBlackholePortsAlarmStateArchived                    DeprecatedInstanceBlackholePortsAlarmState = "archived"
	DeprecatedInstanceBlackholePortsAlarmStateClear                       DeprecatedInstanceBlackholePortsAlarmState = "clear"
	DeprecatedInstanceBlackholePortsAlarmStateClearing                    DeprecatedInstanceBlackholePortsAlarmState = "clearing"
	DeprecatedInstanceBlackholePortsAlarmStateClearingFail                DeprecatedInstanceBlackholePortsAlarmState = "clearing_fail"
	DeprecatedInstanceBlackholePortsAlarmStateEndGrace                    DeprecatedInstanceBlackholePortsAlarmState = "end_grace"
	DeprecatedInstanceBlackholePortsAlarmStateEndWait                     DeprecatedInstanceBlackholePortsAlarmState = "end_wait"
	DeprecatedInstanceBlackholePortsAlarmStateManualClear                 DeprecatedInstanceBlackholePortsAlarmState = "manual_clear"
	DeprecatedInstanceBlackholePortsAlarmStateManualClearing              DeprecatedInstanceBlackholePortsAlarmState = "manual_clearing"
	DeprecatedInstanceBlackholePortsAlarmStateManualClearingFail          DeprecatedInstanceBlackholePortsAlarmState = "manual_clearing_fail"
	DeprecatedInstanceBlackholePortsAlarmStateManualMitigating            DeprecatedInstanceBlackholePortsAlarmState = "manual_mitigating"
	DeprecatedInstanceBlackholePortsAlarmStateManualStarting              DeprecatedInstanceBlackholePortsAlarmState = "manual_starting"
	DeprecatedInstanceBlackholePortsAlarmStateManualStartingFail          DeprecatedInstanceBlackholePortsAlarmState = "manual_starting_fail"
	DeprecatedInstanceBlackholePortsAlarmStateMitigating                  DeprecatedInstanceBlackholePortsAlarmState = "mitigating"
	DeprecatedInstanceBlackholePortsAlarmStateStartWait                   DeprecatedInstanceBlackholePortsAlarmState = "start_wait"
	DeprecatedInstanceBlackholePortsAlarmStateStarting                    DeprecatedInstanceBlackholePortsAlarmState = "starting"
	DeprecatedInstanceBlackholePortsAlarmStateStartingFail                DeprecatedInstanceBlackholePortsAlarmState = "starting_fail"
)

func (r DeprecatedInstanceBlackholePortsAlarmState) IsKnown() bool {
	switch r {
	case DeprecatedInstanceBlackholePortsAlarmStateAckReqUppercase, DeprecatedInstanceBlackholePortsAlarmStateAlarmUppercase, DeprecatedInstanceBlackholePortsAlarmStateArchivedUppercase, DeprecatedInstanceBlackholePortsAlarmStateClearUppercase, DeprecatedInstanceBlackholePortsAlarmStateClearingUppercase, DeprecatedInstanceBlackholePortsAlarmStateClearingFailUppercase, DeprecatedInstanceBlackholePortsAlarmStateEndGraceUppercase, DeprecatedInstanceBlackholePortsAlarmStateEndWaitUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualClearUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualClearingUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualClearingFailUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualMitigatingUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualStartingUppercase, DeprecatedInstanceBlackholePortsAlarmStateManualStartingFailUppercase, DeprecatedInstanceBlackholePortsAlarmStateMitigatingUppercase, DeprecatedInstanceBlackholePortsAlarmStateStartingUppercase, DeprecatedInstanceBlackholePortsAlarmStateStartingFailUppercase, DeprecatedInstanceBlackholePortsAlarmStateStartWaitUppercase, DeprecatedInstanceBlackholePortsAlarmStateAckReq, DeprecatedInstanceBlackholePortsAlarmStateAlarm, DeprecatedInstanceBlackholePortsAlarmStateArchived, DeprecatedInstanceBlackholePortsAlarmStateClear, DeprecatedInstanceBlackholePortsAlarmStateClearing, DeprecatedInstanceBlackholePortsAlarmStateClearingFail, DeprecatedInstanceBlackholePortsAlarmStateEndGrace, DeprecatedInstanceBlackholePortsAlarmStateEndWait, DeprecatedInstanceBlackholePortsAlarmStateManualClear, DeprecatedInstanceBlackholePortsAlarmStateManualClearing, DeprecatedInstanceBlackholePortsAlarmStateManualClearingFail, DeprecatedInstanceBlackholePortsAlarmStateManualMitigating, DeprecatedInstanceBlackholePortsAlarmStateManualStarting, DeprecatedInstanceBlackholePortsAlarmStateManualStartingFail, DeprecatedInstanceBlackholePortsAlarmStateMitigating, DeprecatedInstanceBlackholePortsAlarmStateStartWait, DeprecatedInstanceBlackholePortsAlarmStateStarting, DeprecatedInstanceBlackholePortsAlarmStateStartingFail:
		return true
	}
	return false
}

// The name of the alert policy generating this notification
type DeprecatedInstanceBlackholePortsAlertPolicyName string

const (
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients1000Gbps     DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_1000Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients120Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_120Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients160Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_160Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients2000Gbps     DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_2000Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients200Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_200Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients200kpps5Gbps DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_200kpps_5Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients240Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_240Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients300Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_300Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients400Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_400Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients40Gbps       DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_40Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients500Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_500Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients600Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_600Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients700Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_700Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients800Gbps      DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_800Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients80Gbps       DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_80Gbps"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClientsGlobal       DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_GLOBAL"
	DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClientsIspGrace2    DeprecatedInstanceBlackholePortsAlertPolicyName = "RTBH_Clients_ISP_grace2"
)

func (r DeprecatedInstanceBlackholePortsAlertPolicyName) IsKnown() bool {
	switch r {
	case DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients1000Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients120Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients160Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients2000Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients200Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients200kpps5Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients240Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients300Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients400Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients40Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients500Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients600Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients700Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients800Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClients80Gbps, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClientsGlobal, DeprecatedInstanceBlackholePortsAlertPolicyNameRtbhClientsIspGrace2:
		return true
	}
	return false
}

type DeprecatedInstanceFixedIPAssignment struct {
	// Is network external
	External bool `json:"external"`
	// Ip address
	IPAddress string `json:"ip_address"`
	// Interface subnet id
	SubnetID string                                  `json:"subnet_id"`
	JSON     deprecatedInstanceFixedIPAssignmentJSON `json:"-"`
}

// deprecatedInstanceFixedIPAssignmentJSON contains the JSON metadata for the
// struct [DeprecatedInstanceFixedIPAssignment]
type deprecatedInstanceFixedIPAssignmentJSON struct {
	External    apijson.Field
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedInstanceFixedIPAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceFixedIPAssignmentJSON) RawJSON() string {
	return r.raw
}

// Flavor
type DeprecatedInstanceFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture"`
	FlavorID     string `json:"flavor_id"`
	FlavorName   string `json:"flavor_name"`
	// Various hardware hints for flavor.
	HardwareDescription DeprecatedFlavorHardwareDescription `json:"hardware_description"`
	// Flavor operating system
	OsType string `json:"os_type"`
	Ram    int64  `json:"ram"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string                       `json:"resource_class"`
	Vcpus         int64                        `json:"vcpus"`
	JSON          deprecatedInstanceFlavorJSON `json:"-"`
}

// deprecatedInstanceFlavorJSON contains the JSON metadata for the struct
// [DeprecatedInstanceFlavor]
type deprecatedInstanceFlavorJSON struct {
	Architecture        apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	OsType              apijson.Field
	Ram                 apijson.Field
	ResourceClass       apijson.Field
	Vcpus               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeprecatedInstanceFlavor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceFlavorJSON) RawJSON() string {
	return r.raw
}

// Instance isolation information
type DeprecatedInstanceInstanceIsolation struct {
	// The reason of instance isolation if it is isolated from external internet.
	Reason string                                  `json:"reason,nullable"`
	JSON   deprecatedInstanceInstanceIsolationJSON `json:"-"`
}

// deprecatedInstanceInstanceIsolationJSON contains the JSON metadata for the
// struct [DeprecatedInstanceInstanceIsolation]
type deprecatedInstanceInstanceIsolationJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedInstanceInstanceIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceInstanceIsolationJSON) RawJSON() string {
	return r.raw
}

// Instance volume
type DeprecatedInstanceVolume struct {
	// Volume ID
	ID string `json:"id"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination bool                         `json:"delete_on_termination"`
	JSON                deprecatedInstanceVolumeJSON `json:"-"`
}

// deprecatedInstanceVolumeJSON contains the JSON metadata for the struct
// [DeprecatedInstanceVolume]
type deprecatedInstanceVolumeJSON struct {
	ID                  apijson.Field
	DeleteOnTermination apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeprecatedInstanceVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedInstanceVolumeJSON) RawJSON() string {
	return r.raw
}

// Instance ports security groups
type InstancePortsSecurityGroupsParam struct {
	// Security group name, applies to all ports
	Name param.Field[string] `json:"name"`
	// Port security groups mapping
	PortsSecurityGroupNames param.Field[[]InstancePortsSecurityGroupsPortsSecurityGroupNameParam] `json:"ports_security_group_names"`
}

func (r InstancePortsSecurityGroupsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Port security group names
type InstancePortsSecurityGroupsPortsSecurityGroupNameParam struct {
	// Port ID. If None, security groups will be applied to all ports
	PortID param.Field[string] `json:"port_id,required"`
	// List of security group names
	SecurityGroupNames param.Field[[]string] `json:"security_group_names,required"`
}

func (r InstancePortsSecurityGroupsPortsSecurityGroupNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance will be attached to the network subnet with the largest count of
// available ips
type NewInterfaceAnySubnetParam struct {
	// Port will get an IP address in this network subnet
	NetworkID param.Field[string] `json:"network_id,required"`
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[NewInterfaceAnySubnetIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Must be 'any_subnet'
	Type param.Field[string] `json:"type"`
}

func (r NewInterfaceAnySubnetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewInterfaceAnySubnetParam) implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion() {}

func (r NewInterfaceAnySubnetParam) implementsCloudV1InstanceAttachInterfaceParamsBodyUnion() {}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type NewInterfaceAnySubnetIPFamily string

const (
	NewInterfaceAnySubnetIPFamilyDual NewInterfaceAnySubnetIPFamily = "dual"
	NewInterfaceAnySubnetIPFamilyIpv4 NewInterfaceAnySubnetIPFamily = "ipv4"
	NewInterfaceAnySubnetIPFamilyIpv6 NewInterfaceAnySubnetIPFamily = "ipv6"
)

func (r NewInterfaceAnySubnetIPFamily) IsKnown() bool {
	switch r {
	case NewInterfaceAnySubnetIPFamilyDual, NewInterfaceAnySubnetIPFamilyIpv4, NewInterfaceAnySubnetIPFamilyIpv6:
		return true
	}
	return false
}

// Instance will be attached to default external network
type NewInterfaceExternalExtendWithDdosParam struct {
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[NewInterfaceExternalExtendWithDdosIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r NewInterfaceExternalExtendWithDdosParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewInterfaceExternalExtendWithDdosParam) implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion() {
}

func (r NewInterfaceExternalExtendWithDdosParam) implementsCloudV1InstanceAttachInterfaceParamsBodyUnion() {
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type NewInterfaceExternalExtendWithDdosIPFamily string

const (
	NewInterfaceExternalExtendWithDdosIPFamilyDual NewInterfaceExternalExtendWithDdosIPFamily = "dual"
	NewInterfaceExternalExtendWithDdosIPFamilyIpv4 NewInterfaceExternalExtendWithDdosIPFamily = "ipv4"
	NewInterfaceExternalExtendWithDdosIPFamilyIpv6 NewInterfaceExternalExtendWithDdosIPFamily = "ipv6"
)

func (r NewInterfaceExternalExtendWithDdosIPFamily) IsKnown() bool {
	switch r {
	case NewInterfaceExternalExtendWithDdosIPFamilyDual, NewInterfaceExternalExtendWithDdosIPFamilyIpv4, NewInterfaceExternalExtendWithDdosIPFamilyIpv6:
		return true
	}
	return false
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
type NewInterfaceReservedFixedIPParam struct {
	// Port ID
	PortID param.Field[string] `json:"port_id,required"`
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Must be 'reserved_fixed_ip'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r NewInterfaceReservedFixedIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewInterfaceReservedFixedIPParam) implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion() {
}

func (r NewInterfaceReservedFixedIPParam) implementsCloudV1InstanceAttachInterfaceParamsBodyUnion() {}

// Instance will be attached to specified subnet
type NewInterfaceSpecificSubnetParam struct {
	// Port will get an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id,required"`
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// List of security group IDs
	SecurityGroups param.Field[[]MandatoryIDParam] `json:"security_groups"`
	// Must be 'subnet'
	Type param.Field[string] `json:"type"`
}

func (r NewInterfaceSpecificSubnetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewInterfaceSpecificSubnetParam) implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion() {}

func (r NewInterfaceSpecificSubnetParam) implementsCloudV1InstanceAttachInterfaceParamsBodyUnion() {}

// PortIdWithIpSchema schema
type PortIDWithIPParam struct {
	// IP address
	IPAddress param.Field[string] `json:"ip_address,required"`
	// ID of the port
	PortID param.Field[string] `json:"port_id,required"`
}

func (r PortIDWithIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PortList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []PortListResult `json:"results"`
	JSON    portListJSON     `json:"-"`
}

// portListJSON contains the JSON metadata for the struct [PortList]
type portListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PortList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portListJSON) RawJSON() string {
	return r.raw
}

// PortSchema schema
type PortListResult struct {
	// IP addresses assigned to this port
	IPAssignments []PortIP `json:"ip_assignments,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required"`
	// MAC address of the virtual port
	MacAddress string             `json:"mac_address"`
	JSON       portListResultJSON `json:"-"`
}

// portListResultJSON contains the JSON metadata for the struct [PortListResult]
type portListResultJSON struct {
	IPAssignments apijson.Field
	NetworkID     apijson.Field
	PortID        apijson.Field
	MacAddress    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PortListResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portListResultJSON) RawJSON() string {
	return r.raw
}

type PortSecurityGroupsList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []PortSecurityGroupsListResult `json:"results"`
	JSON    portSecurityGroupsListJSON     `json:"-"`
}

// portSecurityGroupsListJSON contains the JSON metadata for the struct
// [PortSecurityGroupsList]
type portSecurityGroupsListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PortSecurityGroupsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portSecurityGroupsListJSON) RawJSON() string {
	return r.raw
}

// Port object
type PortSecurityGroupsListResult struct {
	// Port ID
	ID string `json:"id"`
	// Port name
	Name string `json:"name"`
	// Security groups applied to port
	SecurityGroups []IDName                         `json:"security_groups"`
	JSON           portSecurityGroupsListResultJSON `json:"-"`
}

// portSecurityGroupsListResultJSON contains the JSON metadata for the struct
// [PortSecurityGroupsListResult]
type portSecurityGroupsListResultJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	SecurityGroups apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PortSecurityGroupsListResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portSecurityGroupsListResultJSON) RawJSON() string {
	return r.raw
}

type RemoteConsole struct {
	// Remote console information
	RemoteConsole RemoteConsoleRemoteConsole `json:"remote_console,required"`
	JSON          remoteConsoleJSON          `json:"-"`
}

// remoteConsoleJSON contains the JSON metadata for the struct [RemoteConsole]
type remoteConsoleJSON struct {
	RemoteConsole apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RemoteConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r remoteConsoleJSON) RawJSON() string {
	return r.raw
}

// Remote console information
type RemoteConsoleRemoteConsole struct {
	Protocol string                         `json:"protocol,required"`
	Type     string                         `json:"type,required"`
	URL      string                         `json:"url,required" format:"uri"`
	JSON     remoteConsoleRemoteConsoleJSON `json:"-"`
}

// remoteConsoleRemoteConsoleJSON contains the JSON metadata for the struct
// [RemoteConsoleRemoteConsole]
type remoteConsoleRemoteConsoleJSON struct {
	Protocol    apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RemoteConsoleRemoteConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r remoteConsoleRemoteConsoleJSON) RawJSON() string {
	return r.raw
}

type CloudV1AIClusterAddSecurityGroupParams struct {
	// Instance ports security groups
	InstancePortsSecurityGroups InstancePortsSecurityGroupsParam `json:"instance_ports_security_groups"`
}

func (r CloudV1AIClusterAddSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstancePortsSecurityGroups)
}

type CloudV1AIClusterAttachInterfaceParams struct {
	// Instance will be attached to default external network
	Body CloudV1AIClusterAttachInterfaceParamsBodyUnion `json:"body"`
}

func (r CloudV1AIClusterAttachInterfaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Instance will be attached to default external network
type CloudV1AIClusterAttachInterfaceParamsBody struct {
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV1AIClusterAttachInterfaceParamsBodyIPFamily] `json:"ip_family"`
	// Port will get an IP address in this network subnet
	NetworkID param.Field[string] `json:"network_id"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Port ID
	PortID         param.Field[string]      `json:"port_id"`
	SecurityGroups param.Field[interface{}] `json:"security_groups"`
	// Port will get an IP address from this subnet
	SubnetID param.Field[string] `json:"subnet_id"`
	// Must be 'external'. Union tag
	Type param.Field[string] `json:"type"`
}

func (r CloudV1AIClusterAttachInterfaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1AIClusterAttachInterfaceParamsBody) implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion() {
}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalExtendWithDdosParam],
// [NewInterfaceSpecificSubnetParam], [NewInterfaceAnySubnetParam],
// [NewInterfaceReservedFixedIPParam], [CloudV1AIClusterAttachInterfaceParamsBody].
type CloudV1AIClusterAttachInterfaceParamsBodyUnion interface {
	implementsCloudV1AIClusterAttachInterfaceParamsBodyUnion()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV1AIClusterAttachInterfaceParamsBodyIPFamily string

const (
	CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyDual CloudV1AIClusterAttachInterfaceParamsBodyIPFamily = "dual"
	CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyIpv4 CloudV1AIClusterAttachInterfaceParamsBodyIPFamily = "ipv4"
	CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyIpv6 CloudV1AIClusterAttachInterfaceParamsBodyIPFamily = "ipv6"
)

func (r CloudV1AIClusterAttachInterfaceParamsBodyIPFamily) IsKnown() bool {
	switch r {
	case CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyDual, CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyIpv4, CloudV1AIClusterAttachInterfaceParamsBodyIPFamilyIpv6:
		return true
	}
	return false
}

type CloudV1AIClusterCheckQuotaParams struct {
	// Check quota before bare metal servers will be created.
	CheckQuotaBeforeBmCreation CheckQuotaBeforeBmCreationParam `json:"check_quota_before_bm_creation"`
}

func (r CloudV1AIClusterCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckQuotaBeforeBmCreation)
}

type CloudV1AIClusterDetachInterfaceParams struct {
	// PortIdWithIpSchema schema
	PortIDWithIP PortIDWithIPParam `json:"port_id_with_ip,required"`
}

func (r CloudV1AIClusterDetachInterfaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PortIDWithIP)
}

type CloudV1AIClusterRemoveSecurityGroupParams struct {
	// Instance ports security groups
	InstancePortsSecurityGroups InstancePortsSecurityGroupsParam `json:"instance_ports_security_groups"`
}

func (r CloudV1AIClusterRemoveSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstancePortsSecurityGroups)
}
