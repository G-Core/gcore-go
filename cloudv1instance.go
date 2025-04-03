// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/tidwall/gjson"
)

// CloudV1InstanceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1InstanceService] method instead.
type CloudV1InstanceService struct {
	Options          []option.RequestOption
	AvailableFlavors *CloudV1InstanceAvailableFlavorService
	Metadata         *CloudV1InstanceMetadataService
}

// NewCloudV1InstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1InstanceService(opts ...option.RequestOption) (r *CloudV1InstanceService) {
	r = &CloudV1InstanceService{}
	r.Options = opts
	r.AvailableFlavors = NewCloudV1InstanceAvailableFlavorService(opts...)
	r.Metadata = NewCloudV1InstanceMetadataService(opts...)
	return
}

// Get instance
func (r *CloudV1InstanceService) Get(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rename instance
func (r *CloudV1InstanceService) Update(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceUpdateParams, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// The date and time stamp format in changes-since or changes-before should be ISO
// 8601: CCYY-MM-DDThh:mm:ss±hh:mm For example, 2015-08-27T09:49:58-05:00. Values
// must be urlencoded. If the time zone is omitted, the UTC time zone is assumed.
// When both changes-since and changes-before are specified, the value of the
// changes-since must be earlier than or equal to the value of the changes-before.
func (r *CloudV1InstanceService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1InstanceListParams, opts ...option.RequestOption) (res *CloudV1InstanceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/instances/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete instance
func (r *CloudV1InstanceService) Delete(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Assign the security group to the server. To assign multiple security groups to
// all ports, use the NULL value for the port_id field
func (r *CloudV1InstanceService) AddSecurityGroup(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceAddSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/addsecuritygroup", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Attach interface to instance
func (r *CloudV1InstanceService) AttachInterface(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceAttachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/attach_interface", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change the flavor of the server instance
func (r *CloudV1InstanceService) ChangeFlavor(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceChangeFlavorParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/changeflavor", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach interface from instance
func (r *CloudV1InstanceService) DetachInterface(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceDetachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/detach_interface", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deprecated. Get instance naming restrictions that are applied to specified
// project and region.
func (r *CloudV1InstanceService) GetAvailableNames(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1InstanceGetAvailableNamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/available_names", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get instance console URL
func (r *CloudV1InstanceService) GetConsole(ctx context.Context, projectID int64, regionID int64, instanceID string, query CloudV1InstanceGetConsoleParams, opts ...option.RequestOption) (res *RemoteConsole, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/get_console", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get instance metrics, including cpu, memory, network and disk metrics
func (r *CloudV1InstanceService) GetMetrics(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceGetMetricsParams, opts ...option.RequestOption) (res *CloudV1InstanceGetMetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metrics", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List of server security groups
func (r *CloudV1InstanceService) GetSecurityGroups(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *CloudV1InstanceGetSecurityGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/securitygroups", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network interfaces attached to the instance
func (r *CloudV1InstanceService) ListInterfaces(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *PortList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/interfaces", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List network ports of instance
func (r *CloudV1InstanceService) ListPorts(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *PortSecurityGroupsList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/ports", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Powercycle (stop and start) the instance, aka hard reboot
func (r *CloudV1InstanceService) Powercycle(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/powercycle", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Put instance into the server group
func (r *CloudV1InstanceService) PutIntoServergroup(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstancePutIntoServergroupParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/put_into_servergroup", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deprecated. Reboot instance
func (r *CloudV1InstanceService) Reboot(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/reboot", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Remove instance from the server group
func (r *CloudV1InstanceService) RemoveFromServergroup(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/remove_from_servergroup", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Un-assign the security group to the server. To un-assign multiple security
// groups to all ports, use the NULL value for the port_id field
func (r *CloudV1InstanceService) RemoveSecurityGroup(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceRemoveSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/delsecuritygroup", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Resume a suspended server to an active state.
func (r *CloudV1InstanceService) Resume(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/resume", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Search for instance in all clients of the reseller
func (r *CloudV1InstanceService) Search(ctx context.Context, query CloudV1InstanceSearchParams, opts ...option.RequestOption) (res *CloudV1InstanceSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/instances/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Power on the server. Required VM state: Shutoff, Stopped
func (r *CloudV1InstanceService) Start(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceStartParams, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/start", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Power off the server. Required VM state: Active, Shutoff, Rescued
func (r *CloudV1InstanceService) Stop(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/stop", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// When you suspend a server, its VM state is stored on disk, all memory is written
// to disk, and the virtual machine is stopped. Suspending a server is similar to
// placing a device in hibernation, and its occupied resource will not be freed but
// rather kept for when the server is resumed. Required VM state: Active, Shutoff
func (r *CloudV1InstanceService) Suspend(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *DeprecatedInstance, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/suspend", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Name
type DeprecatedName struct {
	// Name.
	Name string             `json:"name,required"`
	JSON deprecatedNameJSON `json:"-"`
}

// deprecatedNameJSON contains the JSON metadata for the struct [DeprecatedName]
type deprecatedNameJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedNameJSON) RawJSON() string {
	return r.raw
}

// Name
type DeprecatedNameParam struct {
	// Name.
	Name param.Field[string] `json:"name,required"`
}

func (r DeprecatedNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// IdName schema
type IDName struct {
	// Name.
	Name string `json:"name,required"`
	// ID
	ID   string     `json:"id"`
	JSON idNameJSON `json:"-"`
}

// idNameJSON contains the JSON metadata for the struct [IDName]
type idNameJSON struct {
	Name        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idNameJSON) RawJSON() string {
	return r.raw
}

type InstanceMetricsTimeUnit string

const (
	InstanceMetricsTimeUnitDay  InstanceMetricsTimeUnit = "day"
	InstanceMetricsTimeUnitHour InstanceMetricsTimeUnit = "hour"
)

func (r InstanceMetricsTimeUnit) IsKnown() bool {
	switch r {
	case InstanceMetricsTimeUnitDay, InstanceMetricsTimeUnitHour:
		return true
	}
	return false
}

type CloudV1InstanceListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1InstanceListResponseResult `json:"results,required"`
	JSON    cloudV1InstanceListResponseJSON     `json:"-"`
}

// cloudV1InstanceListResponseJSON contains the JSON metadata for the struct
// [CloudV1InstanceListResponse]
type cloudV1InstanceListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceListResponseResult struct {
	// Map of network_name
	Addresses map[string][]CloudV1InstanceListResponseResultsAddress `json:"addresses,required"`
	// Flavor
	Flavor CloudV1InstanceListResponseResultsFlavor `json:"flavor,required"`
	// Datetime when instance was created
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// Instance ID
	InstanceID string `json:"instance_id,required"`
	// Instance name
	InstanceName string `json:"instance_name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Instance status
	Status string `json:"status,required"`
	// Virtual machine state (active)
	VmState string `json:"vm_state,required"`
	// List of volumes
	Volumes []CloudV1InstanceListResponseResultsVolume `json:"volumes,required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []CloudV1InstanceListResponseResultsBlackholePort `json:"blackhole_ports"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Bare metal advanced DDoS protection profile.
	DdosProfile GetClientProfile `json:"ddos_profile,nullable"`
	// Fixed IP assigned to instance
	FixedIPAssignments []CloudV1InstanceListResponseResultsFixedIPAssignment `json:"fixed_ip_assignments,nullable"`
	// Instance description
	InstanceDescription string `json:"instance_description,nullable"`
	// Instance isolation information
	InstanceIsolation CloudV1InstanceListResponseResultsInstanceIsolation `json:"instance_isolation,nullable"`
	// Key pair assigned to instance
	KeypairName string `json:"keypair_name,nullable"`
	// Metadata dictionary
	Metadata map[string]string `json:"metadata"`
	// Detailed VM metadata
	MetadataDetailed []DetailedMetadata `json:"metadata_detailed"`
	// Security groups
	SecurityGroups []NamePydantic `json:"security_groups"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,nullable"`
	// Task state
	TaskState string                                `json:"task_state,nullable"`
	JSON      cloudV1InstanceListResponseResultJSON `json:"-"`
}

// cloudV1InstanceListResponseResultJSON contains the JSON metadata for the struct
// [CloudV1InstanceListResponseResult]
type cloudV1InstanceListResponseResultJSON struct {
	Addresses           apijson.Field
	Flavor              apijson.Field
	InstanceCreated     apijson.Field
	InstanceID          apijson.Field
	InstanceName        apijson.Field
	ProjectID           apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Status              apijson.Field
	VmState             apijson.Field
	Volumes             apijson.Field
	BlackholePorts      apijson.Field
	CreatorTaskID       apijson.Field
	DdosProfile         apijson.Field
	FixedIPAssignments  apijson.Field
	InstanceDescription apijson.Field
	InstanceIsolation   apijson.Field
	KeypairName         apijson.Field
	Metadata            apijson.Field
	MetadataDetailed    apijson.Field
	SecurityGroups      apijson.Field
	TaskID              apijson.Field
	TaskState           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceListResponseResultsAddress struct {
	// Address
	Addr string `json:"addr,required"`
	// Type of the address
	Type string `json:"type,required"`
	// Interface name. The value will be null if `with_interfaces_name` is not set in
	// the request OR if the interface name was not provided.
	InterfaceName string `json:"interface_name,nullable"`
	// Subnet id
	SubnetID string `json:"subnet_id"`
	// Subnet name
	SubnetName string                                        `json:"subnet_name"`
	JSON       cloudV1InstanceListResponseResultsAddressJSON `json:"-"`
	union      CloudV1InstanceListResponseResultsAddressesUnion
}

// cloudV1InstanceListResponseResultsAddressJSON contains the JSON metadata for the
// struct [CloudV1InstanceListResponseResultsAddress]
type cloudV1InstanceListResponseResultsAddressJSON struct {
	Addr          apijson.Field
	Type          apijson.Field
	InterfaceName apijson.Field
	SubnetID      apijson.Field
	SubnetName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r cloudV1InstanceListResponseResultsAddressJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1InstanceListResponseResultsAddress) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1InstanceListResponseResultsAddress{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1InstanceListResponseResultsAddressesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer],
// [CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer],
// [CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer].
func (r CloudV1InstanceListResponseResultsAddress) AsUnion() CloudV1InstanceListResponseResultsAddressesUnion {
	return r.union
}

// Union satisfied by
// [CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer],
// [CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer] or
// [CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer].
type CloudV1InstanceListResponseResultsAddressesUnion interface {
	implementsCloudV1InstanceListResponseResultsAddress()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1InstanceListResponseResultsAddressesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer{}),
		},
	)
}

type CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer struct {
	// Address
	Addr string `json:"addr,required"`
	// Type of the address
	Type string                                                                 `json:"type,required"`
	JSON cloudV1InstanceListResponseResultsAddressesSimpleAddressSerializerJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsAddressesSimpleAddressSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer]
type cloudV1InstanceListResponseResultsAddressesSimpleAddressSerializerJSON struct {
	Addr        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsAddressesSimpleAddressSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1InstanceListResponseResultsAddressesSimpleAddressSerializer) implementsCloudV1InstanceListResponseResultsAddress() {
}

type CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer struct {
	// Address
	Addr string `json:"addr,required"`
	// Type of the address
	Type string `json:"type,required"`
	// Interface name. The value will be null if `with_interfaces_name` is not set in
	// the request OR if the interface name was not provided.
	InterfaceName string                                                                    `json:"interface_name,nullable"`
	JSON          cloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializerJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer]
type cloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializerJSON struct {
	Addr          apijson.Field
	Type          apijson.Field
	InterfaceName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1InstanceListResponseResultsAddressesAddressInterfaceSerializer) implementsCloudV1InstanceListResponseResultsAddress() {
}

type CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer struct {
	// Address
	Addr string `json:"addr,required"`
	// Subnet id
	SubnetID string `json:"subnet_id,required"`
	// Subnet name
	SubnetName string `json:"subnet_name,required"`
	// Type of the address
	Type string `json:"type,required"`
	// Interface name. The value will be null if `with_interfaces_name` is not set in
	// the request OR if the interface name was not provided.
	InterfaceName string                                                                   `json:"interface_name,nullable"`
	JSON          cloudV1InstanceListResponseResultsAddressesAddressDetailedSerializerJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsAddressesAddressDetailedSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer]
type cloudV1InstanceListResponseResultsAddressesAddressDetailedSerializerJSON struct {
	Addr          apijson.Field
	SubnetID      apijson.Field
	SubnetName    apijson.Field
	Type          apijson.Field
	InterfaceName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsAddressesAddressDetailedSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1InstanceListResponseResultsAddressesAddressDetailedSerializer) implementsCloudV1InstanceListResponseResultsAddress() {
}

// Flavor
type CloudV1InstanceListResponseResultsFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled,required"`
	// Short ID
	FlavorID string `json:"flavor_id,required"`
	// Human readable name
	FlavorName string `json:"flavor_name,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity,nullable"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// This field can have the runtime type of [map[string]string].
	HardwareDescription interface{} `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus `json:"price_status,nullable"`
	// Count of reserved but not used nodes. If a client don't have reservations for
	// the flavor, it's None.
	ReservedInStock int64 `json:"reserved_in_stock,nullable"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string                                       `json:"resource_class,nullable"`
	JSON          cloudV1InstanceListResponseResultsFlavorJSON `json:"-"`
	union         CloudV1InstanceListResponseResultsFlavorUnion
}

// cloudV1InstanceListResponseResultsFlavorJSON contains the JSON metadata for the
// struct [CloudV1InstanceListResponseResultsFlavor]
type cloudV1InstanceListResponseResultsFlavorJSON struct {
	Architecture        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	OsType              apijson.Field
	Ram                 apijson.Field
	Vcpus               apijson.Field
	Capacity            apijson.Field
	CurrencyCode        apijson.Field
	HardwareDescription apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	ReservedInStock     apijson.Field
	ResourceClass       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r cloudV1InstanceListResponseResultsFlavorJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1InstanceListResponseResultsFlavor) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1InstanceListResponseResultsFlavor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1InstanceListResponseResultsFlavorUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV1InstanceListResponseResultsFlavorFlavorSerializer],
// [CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer].
func (r CloudV1InstanceListResponseResultsFlavor) AsUnion() CloudV1InstanceListResponseResultsFlavorUnion {
	return r.union
}

// Flavor
//
// Union satisfied by [CloudV1InstanceListResponseResultsFlavorFlavorSerializer] or
// [CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer].
type CloudV1InstanceListResponseResultsFlavorUnion interface {
	implementsCloudV1InstanceListResponseResultsFlavor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1InstanceListResponseResultsFlavorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV1InstanceListResponseResultsFlavorFlavorSerializer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer{}),
		},
	)
}

type CloudV1InstanceListResponseResultsFlavorFlavorSerializer struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled,required"`
	// Short ID
	FlavorID string `json:"flavor_id,required"`
	// Human readable name
	FlavorName string `json:"flavor_name,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity,nullable"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus                                           `json:"price_status,nullable"`
	JSON        cloudV1InstanceListResponseResultsFlavorFlavorSerializerJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsFlavorFlavorSerializerJSON contains the JSON
// metadata for the struct
// [CloudV1InstanceListResponseResultsFlavorFlavorSerializer]
type cloudV1InstanceListResponseResultsFlavorFlavorSerializerJSON struct {
	Architecture        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	OsType              apijson.Field
	Ram                 apijson.Field
	Vcpus               apijson.Field
	Capacity            apijson.Field
	CurrencyCode        apijson.Field
	HardwareDescription apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsFlavorFlavorSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsFlavorFlavorSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1InstanceListResponseResultsFlavorFlavorSerializer) implementsCloudV1InstanceListResponseResultsFlavor() {
}

type CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled,required"`
	// Short ID
	FlavorID string `json:"flavor_id,required"`
	// Human readable name
	FlavorName string `json:"flavor_name,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity,nullable"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus `json:"price_status,nullable"`
	// Count of reserved but not used nodes. If a client don't have reservations for
	// the flavor, it's None.
	ReservedInStock int64 `json:"reserved_in_stock,nullable"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string                                                         `json:"resource_class,nullable"`
	JSON          cloudV1InstanceListResponseResultsFlavorBmFlavorSerializerJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsFlavorBmFlavorSerializerJSON contains the JSON
// metadata for the struct
// [CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer]
type cloudV1InstanceListResponseResultsFlavorBmFlavorSerializerJSON struct {
	Architecture        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	OsType              apijson.Field
	Ram                 apijson.Field
	Vcpus               apijson.Field
	Capacity            apijson.Field
	CurrencyCode        apijson.Field
	HardwareDescription apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	ReservedInStock     apijson.Field
	ResourceClass       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsFlavorBmFlavorSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1InstanceListResponseResultsFlavorBmFlavorSerializer) implementsCloudV1InstanceListResponseResultsFlavor() {
}

type CloudV1InstanceListResponseResultsVolume struct {
	// Volume ID
	ID string `json:"id,required"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination bool                                         `json:"delete_on_termination,required"`
	JSON                cloudV1InstanceListResponseResultsVolumeJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsVolumeJSON contains the JSON metadata for the
// struct [CloudV1InstanceListResponseResultsVolume]
type cloudV1InstanceListResponseResultsVolumeJSON struct {
	ID                  apijson.Field
	DeleteOnTermination apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsVolumeJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceListResponseResultsBlackholePort struct {
	// A date-time string giving the time that the alarm ended
	AlarmEnd time.Time `json:"AlarmEnd,required" format:"date-time"`
	// A date-time string giving the time that the alarm started
	AlarmStart time.Time `json:"AlarmStart,required" format:"date-time"`
	// Current state of alarm
	AlarmState CloudV1InstanceListResponseResultsBlackholePortsAlarmState `json:"AlarmState,required"`
	// Total alert duration
	AlertDuration string `json:"AlertDuration,required"`
	// Notification destination IP address
	DestinationIP string                                              `json:"DestinationIP,required"`
	ID            int64                                               `json:"ID,required"`
	JSON          cloudV1InstanceListResponseResultsBlackholePortJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsBlackholePortJSON contains the JSON metadata
// for the struct [CloudV1InstanceListResponseResultsBlackholePort]
type cloudV1InstanceListResponseResultsBlackholePortJSON struct {
	AlarmEnd      apijson.Field
	AlarmStart    apijson.Field
	AlarmState    apijson.Field
	AlertDuration apijson.Field
	DestinationIP apijson.Field
	ID            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsBlackholePort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsBlackholePortJSON) RawJSON() string {
	return r.raw
}

// Current state of alarm
type CloudV1InstanceListResponseResultsBlackholePortsAlarmState string

const (
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAckReqUppercase             CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "ACK_REQ"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAlarmUppercase              CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "ALARM"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateArchivedUppercase           CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "ARCHIVED"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearUppercase              CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "CLEAR"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingUppercase           CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "CLEARING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingFailUppercase       CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "CLEARING_FAIL"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndGraceUppercase           CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "END_GRACE"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndWaitUppercase            CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "END_WAIT"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearUppercase        CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_CLEAR"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingUppercase     CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_CLEARING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingFailUppercase CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_CLEARING_FAIL"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualMitigatingUppercase   CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_MITIGATING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingUppercase     CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_STARTING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingFailUppercase CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MANUAL_STARTING_FAIL"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateMitigatingUppercase         CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "MITIGATING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingUppercase           CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "STARTING"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingFailUppercase       CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "STARTING_FAIL"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartWaitUppercase          CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "START_WAIT"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAckReq                      CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "ack_req"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAlarm                       CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "alarm"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateArchived                    CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "archived"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClear                       CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "clear"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearing                    CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "clearing"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingFail                CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "clearing_fail"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndGrace                    CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "end_grace"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndWait                     CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "end_wait"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClear                 CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_clear"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearing              CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_clearing"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingFail          CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_clearing_fail"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualMitigating            CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_mitigating"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStarting              CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_starting"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingFail          CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "manual_starting_fail"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateMitigating                  CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "mitigating"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartWait                   CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "start_wait"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStarting                    CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "starting"
	CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingFail                CloudV1InstanceListResponseResultsBlackholePortsAlarmState = "starting_fail"
)

func (r CloudV1InstanceListResponseResultsBlackholePortsAlarmState) IsKnown() bool {
	switch r {
	case CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAckReqUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAlarmUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateArchivedUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingFailUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndGraceUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndWaitUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingFailUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualMitigatingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingFailUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateMitigatingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingFailUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartWaitUppercase, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAckReq, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateAlarm, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateArchived, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClear, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearing, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateClearingFail, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndGrace, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateEndWait, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClear, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearing, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualClearingFail, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualMitigating, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStarting, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateManualStartingFail, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateMitigating, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartWait, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStarting, CloudV1InstanceListResponseResultsBlackholePortsAlarmStateStartingFail:
		return true
	}
	return false
}

type CloudV1InstanceListResponseResultsFixedIPAssignment struct {
	// Is network external
	External bool `json:"external,required"`
	// Ip address
	IPAddress string `json:"ip_address,required"`
	// Interface subnet id
	SubnetID string                                                  `json:"subnet_id,required"`
	JSON     cloudV1InstanceListResponseResultsFixedIPAssignmentJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsFixedIPAssignmentJSON contains the JSON
// metadata for the struct [CloudV1InstanceListResponseResultsFixedIPAssignment]
type cloudV1InstanceListResponseResultsFixedIPAssignmentJSON struct {
	External    apijson.Field
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsFixedIPAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsFixedIPAssignmentJSON) RawJSON() string {
	return r.raw
}

// Instance isolation information
type CloudV1InstanceListResponseResultsInstanceIsolation struct {
	// The reason of instance isolation if it is isolated from external internet.
	Reason string                                                  `json:"reason,nullable"`
	JSON   cloudV1InstanceListResponseResultsInstanceIsolationJSON `json:"-"`
}

// cloudV1InstanceListResponseResultsInstanceIsolationJSON contains the JSON
// metadata for the struct [CloudV1InstanceListResponseResultsInstanceIsolation]
type cloudV1InstanceListResponseResultsInstanceIsolationJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceListResponseResultsInstanceIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceListResponseResultsInstanceIsolationJSON) RawJSON() string {
	return r.raw
}

// Instance naming restrictions.
type CloudV1InstanceGetAvailableNamesResponse struct {
	// If true, instances can be created using "names" field.
	CustomNameAllowed bool `json:"custom_name_allowed,required"`
	// If true, only specific strings are allowed in "name_templates" fields.
	NameTemplatesLimited bool `json:"name_templates_limited,required"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameTemplates []string `json:"allowed_bm_name_templates"`
	// Clients will only be able to use these strings as bare metal server
	// "bm_name_templates".
	AllowedBmNameWinTemplates []string `json:"allowed_bm_name_win_templates"`
	// If "name_templates_limited" is True, this is the list of allowed instance name
	// templates.
	AllowedNameTemplates []string `json:"allowed_name_templates"`
	// If "name_templates_limited" is True, this is the list of allowed windows
	// instance name templates.
	AllowedNameWinTemplates []string                                     `json:"allowed_name_win_templates"`
	JSON                    cloudV1InstanceGetAvailableNamesResponseJSON `json:"-"`
}

// cloudV1InstanceGetAvailableNamesResponseJSON contains the JSON metadata for the
// struct [CloudV1InstanceGetAvailableNamesResponse]
type cloudV1InstanceGetAvailableNamesResponseJSON struct {
	CustomNameAllowed         apijson.Field
	NameTemplatesLimited      apijson.Field
	AllowedBmNameTemplates    apijson.Field
	AllowedBmNameWinTemplates apijson.Field
	AllowedNameTemplates      apijson.Field
	AllowedNameWinTemplates   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CloudV1InstanceGetAvailableNamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceGetAvailableNamesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceGetMetricsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1InstanceGetMetricsResponseResult `json:"results,required"`
	JSON    cloudV1InstanceGetMetricsResponseJSON     `json:"-"`
}

// cloudV1InstanceGetMetricsResponseJSON contains the JSON metadata for the struct
// [CloudV1InstanceGetMetricsResponse]
type cloudV1InstanceGetMetricsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceGetMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceGetMetricsResponseJSON) RawJSON() string {
	return r.raw
}

// Instance metrics item
type CloudV1InstanceGetMetricsResponseResult struct {
	// Timestamp
	Time string `json:"time,required"`
	// CPU utilization, % (max 100% for multi-core)
	CPUUtil float64 `json:"cpu_util,nullable"`
	// Disks metrics for each of the disks attached
	Disks []CloudV1InstanceGetMetricsResponseResultsDisk `json:"disks,nullable"`
	// RAM utilization, %
	MemoryUtil float64 `json:"memory_util,nullable"`
	// Network out, bytes per second
	NetworkBpsEgress float64 `json:"network_Bps_egress,nullable"`
	// Network in, bytes per second
	NetworkBpsIngress float64 `json:"network_Bps_ingress,nullable"`
	// Network out, packets per second
	NetworkPpsEgress float64 `json:"network_pps_egress,nullable"`
	// Network in, packets per second
	NetworkPpsIngress float64                                     `json:"network_pps_ingress,nullable"`
	JSON              cloudV1InstanceGetMetricsResponseResultJSON `json:"-"`
}

// cloudV1InstanceGetMetricsResponseResultJSON contains the JSON metadata for the
// struct [CloudV1InstanceGetMetricsResponseResult]
type cloudV1InstanceGetMetricsResponseResultJSON struct {
	Time              apijson.Field
	CPUUtil           apijson.Field
	Disks             apijson.Field
	MemoryUtil        apijson.Field
	NetworkBpsEgress  apijson.Field
	NetworkBpsIngress apijson.Field
	NetworkPpsEgress  apijson.Field
	NetworkPpsIngress apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1InstanceGetMetricsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceGetMetricsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Disk metrics item
type CloudV1InstanceGetMetricsResponseResultsDisk struct {
	// Disk read, Bytes per second
	DiskBpsRead float64 `json:"disk_Bps_read,nullable"`
	// Disk write, Bytes per second
	DiskBpsWrite float64 `json:"disk_Bps_write,nullable"`
	// Disk read, iops
	DiskIopsRead float64 `json:"disk_iops_read,nullable"`
	// Disk write, iops
	DiskIopsWrite float64 `json:"disk_iops_write,nullable"`
	// Disk attached slot name
	DiskName string                                           `json:"disk_name,nullable"`
	JSON     cloudV1InstanceGetMetricsResponseResultsDiskJSON `json:"-"`
}

// cloudV1InstanceGetMetricsResponseResultsDiskJSON contains the JSON metadata for
// the struct [CloudV1InstanceGetMetricsResponseResultsDisk]
type cloudV1InstanceGetMetricsResponseResultsDiskJSON struct {
	DiskBpsRead   apijson.Field
	DiskBpsWrite  apijson.Field
	DiskIopsRead  apijson.Field
	DiskIopsWrite apijson.Field
	DiskName      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1InstanceGetMetricsResponseResultsDisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceGetMetricsResponseResultsDiskJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceGetSecurityGroupsResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []IDName                                     `json:"results"`
	JSON    cloudV1InstanceGetSecurityGroupsResponseJSON `json:"-"`
}

// cloudV1InstanceGetSecurityGroupsResponseJSON contains the JSON metadata for the
// struct [CloudV1InstanceGetSecurityGroupsResponse]
type cloudV1InstanceGetSecurityGroupsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceGetSecurityGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceGetSecurityGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceSearchResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []CloudV1InstanceSearchResponseResult `json:"results"`
	JSON    cloudV1InstanceSearchResponseJSON     `json:"-"`
}

// cloudV1InstanceSearchResponseJSON contains the JSON metadata for the struct
// [CloudV1InstanceSearchResponse]
type cloudV1InstanceSearchResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceSearchResponseJSON) RawJSON() string {
	return r.raw
}

// Instance location schema.
type CloudV1InstanceSearchResponseResult struct {
	// Instance id
	ID string `json:"id,required"`
	// Id of project's client
	ClientID int64 `json:"client_id,required"`
	// Instance name
	Name string `json:"name,required"`
	// Id of project where instance is deployed
	ProjectID string `json:"project_id,required"`
	// Name of project where instance is deployed
	ProjectName string `json:"project_name,required"`
	// Id of region where instance is deployed
	RegionID string `json:"region_id,required"`
	// Name of region where instance is deployed
	RegionName string                                  `json:"region_name,required"`
	JSON       cloudV1InstanceSearchResponseResultJSON `json:"-"`
}

// cloudV1InstanceSearchResponseResultJSON contains the JSON metadata for the
// struct [CloudV1InstanceSearchResponseResult]
type cloudV1InstanceSearchResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	ProjectName apijson.Field
	RegionID    apijson.Field
	RegionName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1InstanceSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1InstanceSearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1InstanceUpdateParams struct {
	// Name
	DeprecatedName DeprecatedNameParam `json:"deprecated_name,required"`
}

func (r CloudV1InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeprecatedName)
}

type CloudV1InstanceListParams struct {
	// Only show instances which are able to handle floating address
	AvailableFloating param.Field[bool] `query:"available_floating"`
	// Filters the instances by a date and time stamp when the instances last changed.
	// Those instances that changed before or equal to the specified date and time
	// stamp are returned.
	ChangesBefore param.Field[string] `query:"changes-before"`
	// Filters the instances by a date and time stamp when the instances last changed
	// status.
	ChangesSince param.Field[string] `query:"changes-since"`
	// Exclude instances with specified security group name
	ExcludeSecgroup param.Field[string] `query:"exclude_secgroup"`
	// Filter out instances by flavor_id. Flavor id must match exactly. Example:
	// "g1-standard-1-2"
	FlavorID param.Field[string] `query:"flavor_id"`
	// Filter out instances by flavor_prefix. Example: "g1-standard" or "g1-"
	FlavorPrefix param.Field[string] `query:"flavor_prefix"`
	// Include AI instances. default False.
	IncludeAI param.Field[bool] `query:"include_ai"`
	// Include bare metal servers.default False.
	IncludeBaremetal param.Field[bool] `query:"include_baremetal"`
	// Include k8s instances. default True.
	IncludeK8s param.Field[bool] `query:"include_k8s"`
	// An IPv4 address to filter results by. Regular expression allowed
	IP param.Field[string] `query:"ip"`
	// Limit the number of returned instances
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/instances/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Filter by metadata values. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_v=["value", "sense"]" --url "http://localhost:1111/v1/instances/1/1"
	MetadataV param.Field[string] `query:"metadata_v"`
	// Filter out instances by name. Use MySQL regular expression. Example:
	// "^.est*....*[0-9]$". Also, any substring can be used and instances will be
	// returned with names containing the substring. Example: "test".
	Name param.Field[string] `query:"name"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Include only isolated instances. default False.
	OnlyIsolated param.Field[bool] `query:"only_isolated"`
	// Return bare metals only with external fixed IP addresses.
	OnlyWithFixedExternalIP param.Field[bool] `query:"only_with_fixed_external_ip"`
	// Order instances by transmitted fields and directions (name.asc).
	OrderBy param.Field[string] `query:"order_by"`
	// Filter result by profile name. Effective only with with_ddos set to true.
	ProfileName param.Field[string] `query:"profile_name"`
	// Filter result by DDoS protection_status. if parameter is provided. Effective
	// only with with_ddos set to true. (Active, Queued or Error)
	ProtectionStatus param.Field[string] `query:"protection_status"`
	// Filters instances by a server status, as a string. Possible statuses are ACTIVE,
	// ERROR, SHUTOFF, REBOOT, PAUSED, etc.
	Status param.Field[string] `query:"status"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with with_ddos set to true. (advanced or basic)
	TypeDdosProfile param.Field[string] `query:"type_ddos_profile"`
	// Filter the server list result by the UUID of the server. Allowed UUID part
	Uuid param.Field[string] `query:"uuid"`
	// DDoS profile information will be included to Bare metal servers
	WithDdos param.Field[bool] `query:"with_ddos"`
	// Include interfaces name
	WithInterfacesName param.Field[bool] `query:"with_interfaces_name"`
}

// URLQuery serializes [CloudV1InstanceListParams]'s query parameters as
// `url.Values`.
func (r CloudV1InstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1InstanceDeleteParams struct {
	// True if it is required to delete floating IPs assigned to the instance. Can't be
	// used with floatings.
	DeleteFloatings param.Field[bool] `query:"delete_floatings"`
	// Comma separated list of floating ids that should be deleted. Can't be used with
	// delete_floatings.
	Floatings param.Field[string] `query:"floatings"`
	// Comma separated list of port IDs to be deleted with the instance
	ReservedFixedIPs param.Field[string] `query:"reserved_fixed_ips"`
	// Comma separated list of volume IDs to be deleted with the instance
	Volumes param.Field[string] `query:"volumes"`
}

// URLQuery serializes [CloudV1InstanceDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV1InstanceDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1InstanceAddSecurityGroupParams struct {
	// Instance ports security groups
	InstancePortsSecurityGroups InstancePortsSecurityGroupsParam `json:"instance_ports_security_groups"`
}

func (r CloudV1InstanceAddSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstancePortsSecurityGroups)
}

type CloudV1InstanceAttachInterfaceParams struct {
	// Instance will be attached to default external network
	Body CloudV1InstanceAttachInterfaceParamsBodyUnion `json:"body"`
}

func (r CloudV1InstanceAttachInterfaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Instance will be attached to default external network
type CloudV1InstanceAttachInterfaceParamsBody struct {
	// Advanced DDoS protection.
	DdosProfile param.Field[DeprecatedCreateDdosProfileParam] `json:"ddos_profile"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[CloudV1InstanceAttachInterfaceParamsBodyIPFamily] `json:"ip_family"`
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

func (r CloudV1InstanceAttachInterfaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1InstanceAttachInterfaceParamsBody) implementsCloudV1InstanceAttachInterfaceParamsBodyUnion() {
}

// Instance will be attached to default external network
//
// Satisfied by [NewInterfaceExternalExtendWithDdosParam],
// [NewInterfaceSpecificSubnetParam], [NewInterfaceAnySubnetParam],
// [NewInterfaceReservedFixedIPParam], [CloudV1InstanceAttachInterfaceParamsBody].
type CloudV1InstanceAttachInterfaceParamsBodyUnion interface {
	implementsCloudV1InstanceAttachInterfaceParamsBodyUnion()
}

// Which subnets should be selected: IPv4, IPv6 or use dual stack.
type CloudV1InstanceAttachInterfaceParamsBodyIPFamily string

const (
	CloudV1InstanceAttachInterfaceParamsBodyIPFamilyDual CloudV1InstanceAttachInterfaceParamsBodyIPFamily = "dual"
	CloudV1InstanceAttachInterfaceParamsBodyIPFamilyIpv4 CloudV1InstanceAttachInterfaceParamsBodyIPFamily = "ipv4"
	CloudV1InstanceAttachInterfaceParamsBodyIPFamilyIpv6 CloudV1InstanceAttachInterfaceParamsBodyIPFamily = "ipv6"
)

func (r CloudV1InstanceAttachInterfaceParamsBodyIPFamily) IsKnown() bool {
	switch r {
	case CloudV1InstanceAttachInterfaceParamsBodyIPFamilyDual, CloudV1InstanceAttachInterfaceParamsBodyIPFamilyIpv4, CloudV1InstanceAttachInterfaceParamsBodyIPFamilyIpv6:
		return true
	}
	return false
}

type CloudV1InstanceChangeFlavorParams struct {
	// Flavor ID
	FlavorID param.Field[string] `json:"flavor_id,required"`
}

func (r CloudV1InstanceChangeFlavorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1InstanceDetachInterfaceParams struct {
	// PortIdWithIpSchema schema
	PortIDWithIP PortIDWithIPParam `json:"port_id_with_ip,required"`
}

func (r CloudV1InstanceDetachInterfaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PortIDWithIP)
}

type CloudV1InstanceGetConsoleParams struct {
	// Console type
	ConsoleType param.Field[string] `query:"console_type"`
}

// URLQuery serializes [CloudV1InstanceGetConsoleParams]'s query parameters as
// `url.Values`.
func (r CloudV1InstanceGetConsoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1InstanceGetMetricsParams struct {
	// Time interval.
	TimeInterval param.Field[int64] `json:"time_interval,required"`
	// Time interval unit.
	TimeUnit param.Field[InstanceMetricsTimeUnit] `json:"time_unit,required"`
}

func (r CloudV1InstanceGetMetricsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1InstancePutIntoServergroupParams struct {
	// Anti-affinity or affinity or soft-anti-affinity server group ID.
	ServergroupID param.Field[string] `json:"servergroup_id,required"`
}

func (r CloudV1InstancePutIntoServergroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1InstanceRemoveSecurityGroupParams struct {
	// Instance ports security groups
	InstancePortsSecurityGroups InstancePortsSecurityGroupsParam `json:"instance_ports_security_groups"`
}

func (r CloudV1InstanceRemoveSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstancePortsSecurityGroups)
}

type CloudV1InstanceSearchParams struct {
	// Instance id pattern
	ID param.Field[string] `query:"id"`
	// Instance IP address
	InstanceIP param.Field[string] `query:"instance_ip"`
	// Instance name pattern
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [CloudV1InstanceSearchParams]'s query parameters as
// `url.Values`.
func (r CloudV1InstanceSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1InstanceStartParams struct {
	// Activate Advanced DDoS profile when instance is started
	ActivateProfile param.Field[bool] `json:"activate_profile"`
}

func (r CloudV1InstanceStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
