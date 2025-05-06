// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
	"github.com/tidwall/gjson"
)

// InstanceService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options    []option.RequestOption
	Flavors    InstanceFlavorService
	Interfaces InstanceInterfaceService
	Images     InstanceImageService
	Metrics    InstanceMetricService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r InstanceService) {
	r = InstanceService{}
	r.Options = opts
	r.Flavors = NewInstanceFlavorService(opts...)
	r.Interfaces = NewInstanceInterfaceService(opts...)
	r.Images = NewInstanceImageService(opts...)
	r.Metrics = NewInstanceMetricService(opts...)
	return
}

// For Linux,
//
//   - Use the `user_data` field to provide a
//     <a href=https://cloudinit.readthedocs.io/en/latest/reference/examples.html>cloud-init
//     script</a> in base64 to apply configurations to the instance.
//   - Specify the `username` and `password` to create a new user.
//   - When only `password` is provided, it is set as the password for the default
//     user of the image.
//   - The `user_data` is ignored when the `password` is specified.
//
// For Windows,
//
//   - Use the `user_data` field to provide a
//     <a href=https://cloudbase-init.readthedocs.io/en/latest/userdata.html#cloud-config>cloudbase-init
//     script</a> in base64 to create new users on Windows.
//   - Use the `password` field to set the password for the 'Admin' user on Windows.
//   - The password of the Admin user cannot be updated via `user_data`.
//   - The `username` cannot be specified in the request.
func (r *InstanceService) New(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v2/instances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename instance
func (r *InstanceService) Update(ctx context.Context, instanceID string, params InstanceUpdateParams, opts ...option.RequestOption) (res *Instance, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List instances
func (r *InstanceService) List(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Instance], err error) {
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
	path := fmt.Sprintf("cloud/v1/instances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List instances
func (r *InstanceService) ListAutoPaging(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Instance] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete instance
func (r *InstanceService) Delete(ctx context.Context, instanceID string, params InstanceDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// The action can be one of: start, stop, reboot, powercycle, suspend or resume.
// Suspend and resume are not available for baremetal instances.
func (r *InstanceService) Action(ctx context.Context, instanceID string, params InstanceActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/%s/action", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Put instance into the server group
func (r *InstanceService) AddToPlacementGroup(ctx context.Context, instanceID string, params InstanceAddToPlacementGroupParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/put_into_servergroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Assign the security group to the server. To assign multiple security groups to
// all ports, use the NULL value for the port_id field
func (r *InstanceService) AssignSecurityGroup(ctx context.Context, instanceID string, params InstanceAssignSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/addsecuritygroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Disable port security for instance interface
func (r *InstanceService) DisablePortSecurity(ctx context.Context, portID string, body InstanceDisablePortSecurityParams, opts ...option.RequestOption) (res *InstanceInterface, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/disable_port_security", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enable port security for instance interface
func (r *InstanceService) EnablePortSecurity(ctx context.Context, portID string, body InstanceEnablePortSecurityParams, opts ...option.RequestOption) (res *InstanceInterface, err error) {
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
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/enable_port_security", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// **Cookie Parameters**:
//
//   - `language` (str, optional): Language for the response content. Affects the
//     `ddos_profile` field. Supported values:
//   - `'en'` (default)
//   - `'de'`
//   - `'ru'`
func (r *InstanceService) Get(ctx context.Context, instanceID string, query InstanceGetParams, opts ...option.RequestOption) (res *Instance, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get instance console URL
func (r *InstanceService) GetConsole(ctx context.Context, instanceID string, params InstanceGetConsoleParams, opts ...option.RequestOption) (res *Console, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/get_console", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Remove instance from the server group
func (r *InstanceService) RemoveFromPlacementGroup(ctx context.Context, instanceID string, body InstanceRemoveFromPlacementGroupParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/remove_from_servergroup", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Change flavor of the instance
func (r *InstanceService) Resize(ctx context.Context, instanceID string, params InstanceResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/changeflavor", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Un-assign the security group to the server. To un-assign multiple security
// groups to all ports, use the NULL value for the port_id field
func (r *InstanceService) UnassignSecurityGroup(ctx context.Context, instanceID string, params InstanceUnassignSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/delsecuritygroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type InstanceInterface struct {
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs,required"`
	// Bodies of floating IPs that are NAT-ing IPs of this port
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// IP addresses assigned to this port
	IPAssignments []IPAssignment `json:"ip_assignments,required"`
	// Body of the network this port is attached to
	NetworkDetails NetworkDetails `json:"network_details,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// Port security status
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Interface name
	InterfaceName string `json:"interface_name,nullable"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		FloatingipDetails   respjson.Field
		IPAssignments       respjson.Field
		NetworkDetails      respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		PortSecurityEnabled respjson.Field
		InterfaceName       respjson.Field
		MacAddress          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterface) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The flavor of the instance.
	Flavor string `json:"flavor,required"`
	// A list of network interfaces for the instance. You can create one or more
	// interfaces - private, public, or both.
	Interfaces []InstanceNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// List of volumes that will be attached to the instance.
	Volumes []InstanceNewParamsVolumeUnion `json:"volumes,omitzero,required"`
	// Specifies the name of the SSH keypair, created via the
	// <a href="#operation/SSHKeyCollectionViewSet.post">/v1/ssh_keys endpoint</a>.
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// Set to `true` if creating the instance from an `apptemplate`. This allows
	// application ports in the security group for instances created from a marketplace
	// application template.
	AllowAppPorts param.Opt[bool] `json:"allow_app_ports,omitzero"`
	// Instance name.
	Name param.Opt[string] `json:"name,omitzero"`
	// If you want the instance name to be automatically generated based on IP
	// addresses, you can provide a name template instead of specifying the name
	// manually. The template should include a placeholder that will be replaced during
	// provisioning. Supported placeholders are: `{ip_octets}` (last 3 octets of the
	// IP), `{two_ip_octets}`, and `{one_ip_octet}`.
	NameTemplate param.Opt[string] `json:"name_template,omitzero"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// When only 'password' is provided, it is set as the password for the default user
	// of the image. For Windows instances, 'username' cannot be specified. Use the
	// 'password' field to set the password for the 'Admin' user on Windows. Use the
	// 'user_data' field to provide a script to create new users on Windows. The
	// password of the Admin user cannot be updated via 'user_data'.
	Password param.Opt[string] `json:"password,omitzero"`
	// Placement group ID for instance placement policy.
	//
	// Supported group types:
	//
	//   - `anti-affinity`: Ensures instances are placed on different hosts for high
	//     availability.
	//   - `affinity`: Places instances on the same host for low-latency communication.
	//   - `soft-anti-affinity`: Tries to place instances on different hosts but allows
	//     sharing if needed.
	ServergroupID param.Opt[string] `json:"servergroup_id,omitzero" format:"uuid4"`
	// String in base64 format. For Linux instances, 'user_data' is ignored when
	// 'password' field is provided. For Windows instances, Admin user password is set
	// by 'password' field and cannot be updated via 'user_data'. Examples of the
	// user_data: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// For Windows instances, 'username' cannot be specified. Use 'password' field to
	// set the password for the 'Admin' user on Windows.
	Username param.Opt[string] `json:"username,omitzero"`
	// Parameters for the application template if creating the instance from an
	// `apptemplate`.
	Configuration any `json:"configuration,omitzero"`
	// Specifies security group UUIDs to be applied to all instance network interfaces.
	SecurityGroups []InstanceNewParamsSecurityGroup `json:"security_groups,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceUnion struct {
	OfExternal        *InstanceNewParamsInterfaceExternal        `json:",omitzero,inline"`
	OfSubnet          *InstanceNewParamsInterfaceSubnet          `json:",omitzero,inline"`
	OfAnySubnet       *InstanceNewParamsInterfaceAnySubnet       `json:",omitzero,inline"`
	OfReservedFixedIP *InstanceNewParamsInterfaceReservedFixedIP `json:",omitzero,inline"`
	paramUnion
}

func (u InstanceNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceUnion](u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}

func (u *InstanceNewParamsInterfaceUnion) asAny() any {
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
func (u InstanceNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.Valid() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfReservedFixedIP; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetType() *string {
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
func (u InstanceNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfExternal; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfSubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetNetworkID() *string {
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
func (u InstanceNewParamsInterfaceUnion) GetSecurityGroups() (res instanceNewParamsInterfaceUnionSecurityGroups) {
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

// Can have the runtime types [_[]InstanceNewParamsInterfaceExternalSecurityGroup],
// [_[]InstanceNewParamsInterfaceSubnetSecurityGroup],
// [_[]InstanceNewParamsInterfaceAnySubnetSecurityGroup],
// [_[]InstanceNewParamsInterfaceReservedFixedIPSecurityGroup]
type instanceNewParamsInterfaceUnionSecurityGroups struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]cloud.InstanceNewParamsInterfaceExternalSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceSubnetSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceAnySubnetSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceReservedFixedIPSecurityGroup:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instanceNewParamsInterfaceUnionSecurityGroups) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u InstanceNewParamsInterfaceUnion) GetFloatingIP() (res instanceNewParamsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfReservedFixedIP; vt != nil {
		res.any = vt.FloatingIP.asAny()
	}
	return
}

// Can have the runtime types [*InstanceNewParamsInterfaceSubnetFloatingIPNew],
// [*InstanceNewParamsInterfaceSubnetFloatingIPExisting],
// [*InstanceNewParamsInterfaceAnySubnetFloatingIPNew],
// [*InstanceNewParamsInterfaceAnySubnetFloatingIPExisting],
// [*InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew],
// [*InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting]
type instanceNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.InstanceNewParamsInterfaceSubnetFloatingIPNew:
//	case *cloud.InstanceNewParamsInterfaceSubnetFloatingIPExisting:
//	case *cloud.InstanceNewParamsInterfaceAnySubnetFloatingIPNew:
//	case *cloud.InstanceNewParamsInterfaceAnySubnetFloatingIPExisting:
//	case *cloud.InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew:
//	case *cloud.InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instanceNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u instanceNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *InstanceNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetSource()
	case *InstanceNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetSource()
	case *InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u instanceNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *InstanceNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstanceNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceUnion](
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceExternal{}),
			DiscriminatorValue: "external",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceSubnet{}),
			DiscriminatorValue: "subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceAnySubnet{}),
			DiscriminatorValue: "any_subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceReservedFixedIP{}),
			DiscriminatorValue: "reserved_fixed_ip",
		},
	)
}

// Instance will be attached to default external network
//
// The property Type is required.
type InstanceNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Specifies security group UUIDs to be applied to the instance network interface.
	SecurityGroups []InstanceNewParamsInterfaceExternalSecurityGroup `json:"security_groups,omitzero"`
	// A public IP address will be assigned to the instance.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type InstanceNewParamsInterfaceExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsInterfaceExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The instance will get an IP address from the selected network. If you choose to
// add a floating IP, the instance will be reachable from the internet. Otherwise,
// it will only have a private IP within the network.
//
// The properties NetworkID, SubnetID, Type are required.
type InstanceNewParamsInterfaceSubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// The instance will get an IP address from this subnet.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstanceNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specifies security group UUIDs to be applied to the instance network interface.
	SecurityGroups []InstanceNewParamsInterfaceSubnetSecurityGroup `json:"security_groups,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceSubnetFloatingIPUnion struct {
	OfNew      *InstanceNewParamsInterfaceSubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstanceNewParamsInterfaceSubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstanceNewParamsInterfaceSubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceSubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *InstanceNewParamsInterfaceSubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceSubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceSubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceSubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceSubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceSubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewInstanceNewParamsInterfaceSubnetFloatingIPNew() InstanceNewParamsInterfaceSubnetFloatingIPNew {
	return InstanceNewParamsInterfaceSubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstanceNewParamsInterfaceSubnetFloatingIPNew].
type InstanceNewParamsInterfaceSubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceSubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceSubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceSubnetFloatingIPExisting struct {
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

func (r InstanceNewParamsInterfaceSubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceSubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type InstanceNewParamsInterfaceSubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsInterfaceSubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceSubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties NetworkID, Type are required.
type InstanceNewParamsInterfaceAnySubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstanceNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Specifies security group UUIDs to be applied to the instance network interface.
	SecurityGroups []InstanceNewParamsInterfaceAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Instance will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceAnySubnetFloatingIPUnion struct {
	OfNew      *InstanceNewParamsInterfaceAnySubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstanceNewParamsInterfaceAnySubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstanceNewParamsInterfaceAnySubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceAnySubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *InstanceNewParamsInterfaceAnySubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceAnySubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceAnySubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceAnySubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceAnySubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceAnySubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewInstanceNewParamsInterfaceAnySubnetFloatingIPNew() InstanceNewParamsInterfaceAnySubnetFloatingIPNew {
	return InstanceNewParamsInterfaceAnySubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstanceNewParamsInterfaceAnySubnetFloatingIPNew].
type InstanceNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceAnySubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceAnySubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceAnySubnetFloatingIPExisting struct {
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

func (r InstanceNewParamsInterfaceAnySubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceAnySubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type InstanceNewParamsInterfaceAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsInterfaceAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties PortID, Type are required.
type InstanceNewParamsInterfaceReservedFixedIP struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID string `json:"port_id,required"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specifies security group UUIDs to be applied to the instance network interface.
	SecurityGroups []InstanceNewParamsInterfaceReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// An existing available reserved fixed IP will be attached to the instance. If the
	// reserved IP is not public and you choose to add a floating IP, the instance will
	// be accessible from the internet.
	//
	// This field can be elided, and will marshal its zero value as
	// "reserved_fixed_ip".
	Type constant.ReservedFixedIP `json:"type,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion struct {
	OfNew      *InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceReservedFixedIPFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewInstanceNewParamsInterfaceReservedFixedIPFloatingIPNew() InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew {
	return InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstanceNewParamsInterfaceReservedFixedIPFloatingIPNew].
type InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceReservedFixedIPFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
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

func (r InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceReservedFixedIPFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type InstanceNewParamsInterfaceReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsInterfaceReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsVolumeUnion struct {
	OfNewVolume      *InstanceNewParamsVolumeNewVolume      `json:",omitzero,inline"`
	OfImage          *InstanceNewParamsVolumeImage          `json:",omitzero,inline"`
	OfSnapshot       *InstanceNewParamsVolumeSnapshot       `json:",omitzero,inline"`
	OfApptemplate    *InstanceNewParamsVolumeApptemplate    `json:",omitzero,inline"`
	OfExistingVolume *InstanceNewParamsVolumeExistingVolume `json:",omitzero,inline"`
	paramUnion
}

func (u InstanceNewParamsVolumeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsVolumeUnion](u.OfNewVolume,
		u.OfImage,
		u.OfSnapshot,
		u.OfApptemplate,
		u.OfExistingVolume)
}

func (u *InstanceNewParamsVolumeUnion) asAny() any {
	if !param.IsOmitted(u.OfNewVolume) {
		return u.OfNewVolume
	} else if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfSnapshot) {
		return u.OfSnapshot
	} else if !param.IsOmitted(u.OfApptemplate) {
		return u.OfApptemplate
	} else if !param.IsOmitted(u.OfExistingVolume) {
		return u.OfExistingVolume
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetImageID() *string {
	if vt := u.OfImage; vt != nil {
		return &vt.ImageID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetSnapshotID() *string {
	if vt := u.OfSnapshot; vt != nil {
		return &vt.SnapshotID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetApptemplateID() *string {
	if vt := u.OfApptemplate; vt != nil {
		return &vt.ApptemplateID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetVolumeID() *string {
	if vt := u.OfExistingVolume; vt != nil {
		return &vt.VolumeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetSize() *int64 {
	if vt := u.OfNewVolume; vt != nil {
		return (*int64)(&vt.Size)
	} else if vt := u.OfImage; vt != nil && vt.Size.Valid() {
		return &vt.Size.Value
	} else if vt := u.OfSnapshot; vt != nil {
		return (*int64)(&vt.Size)
	} else if vt := u.OfApptemplate; vt != nil && vt.Size.Valid() {
		return &vt.Size.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetSource() *string {
	if vt := u.OfNewVolume; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfSnapshot; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfApptemplate; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExistingVolume; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetAttachmentTag() *string {
	if vt := u.OfNewVolume; vt != nil && vt.AttachmentTag.Valid() {
		return &vt.AttachmentTag.Value
	} else if vt := u.OfImage; vt != nil && vt.AttachmentTag.Valid() {
		return &vt.AttachmentTag.Value
	} else if vt := u.OfSnapshot; vt != nil && vt.AttachmentTag.Valid() {
		return &vt.AttachmentTag.Value
	} else if vt := u.OfApptemplate; vt != nil && vt.AttachmentTag.Valid() {
		return &vt.AttachmentTag.Value
	} else if vt := u.OfExistingVolume; vt != nil && vt.AttachmentTag.Valid() {
		return &vt.AttachmentTag.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetDeleteOnTermination() *bool {
	if vt := u.OfNewVolume; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	} else if vt := u.OfImage; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	} else if vt := u.OfSnapshot; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	} else if vt := u.OfApptemplate; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	} else if vt := u.OfExistingVolume; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetName() *string {
	if vt := u.OfNewVolume; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfImage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSnapshot; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfApptemplate; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetTypeName() *string {
	if vt := u.OfNewVolume; vt != nil {
		return (*string)(&vt.TypeName)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.TypeName)
	} else if vt := u.OfSnapshot; vt != nil {
		return (*string)(&vt.TypeName)
	} else if vt := u.OfApptemplate; vt != nil {
		return (*string)(&vt.TypeName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsVolumeUnion) GetBootIndex() *int64 {
	if vt := u.OfImage; vt != nil && vt.BootIndex.Valid() {
		return &vt.BootIndex.Value
	} else if vt := u.OfSnapshot; vt != nil && vt.BootIndex.Valid() {
		return &vt.BootIndex.Value
	} else if vt := u.OfApptemplate; vt != nil && vt.BootIndex.Valid() {
		return &vt.BootIndex.Value
	} else if vt := u.OfExistingVolume; vt != nil && vt.BootIndex.Valid() {
		return &vt.BootIndex.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Tags property, if present.
func (u InstanceNewParamsVolumeUnion) GetTags() TagUpdateMap {
	if vt := u.OfNewVolume; vt != nil {
		return vt.Tags
	} else if vt := u.OfImage; vt != nil {
		return vt.Tags
	} else if vt := u.OfSnapshot; vt != nil {
		return vt.Tags
	} else if vt := u.OfApptemplate; vt != nil {
		return vt.Tags
	} else if vt := u.OfExistingVolume; vt != nil {
		return vt.Tags
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsVolumeUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsVolumeNewVolume{}),
			DiscriminatorValue: "new-volume",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsVolumeImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsVolumeSnapshot{}),
			DiscriminatorValue: "snapshot",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsVolumeApptemplate{}),
			DiscriminatorValue: "apptemplate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsVolumeExistingVolume{}),
			DiscriminatorValue: "existing-volume",
		},
	)
}

// The properties Size, Source are required.
type InstanceNewParamsVolumeNewVolume struct {
	// Volume size in GiB.
	Size int64 `json:"size,required"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// The name of the volume. If not specified, a name will be generated
	// automatically.
	Name param.Opt[string] `json:"name,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type name. Supported values:
	//
	//   - `standard` - Network SSD block storage offering stable performance with high
	//     random I/O and data reliability (6 IOPS per 1 GiB, 0.4 MB/s per 1 GiB). Max
	//     IOPS: 4500. Max bandwidth: 300 MB/s.
	//   - `ssd_hiiops` - High-performance SSD storage for latency-sensitive
	//     transactional workloads (60 IOPS per 1 GiB, 2.5 MB/s per 1 GiB). Max
	//     IOPS: 9000. Max bandwidth: 500 MB/s.
	//   - `ssd_lowlatency` - SSD storage optimized for low-latency and real-time
	//     processing. Max IOPS: 5000. Average latency: 300 µs. Snapshots and volume
	//     resizing are **not** supported for `ssd_lowlatency`.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// New volume will be created from scratch and attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "new-volume".
	Source constant.NewVolume `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsVolumeNewVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolumeNewVolume
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsVolumeNewVolume](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties ImageID, Source are required.
type InstanceNewParamsVolumeImage struct {
	// Image ID.
	ImageID string `json:"image_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// - `0` means that this is the primary boot device;
	// - A unique positive value is set for the secondary bootable devices;
	// - A negative number means that the boot is prohibited.
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// The name of the volume. If not specified, a name will be generated
	// automatically.
	Name param.Opt[string] `json:"name,omitzero"`
	// Volume size in GiB.
	//
	// - For instances: **specify the desired volume size explicitly**.
	// - For basic VMs: the size is set automatically based on the flavor.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type name. Supported values:
	//
	//   - `standard` - Network SSD block storage offering stable performance with high
	//     random I/O and data reliability (6 IOPS per 1 GiB, 0.4 MB/s per 1 GiB). Max
	//     IOPS: 4500. Max bandwidth: 300 MB/s.
	//   - `ssd_hiiops` - High-performance SSD storage for latency-sensitive
	//     transactional workloads (60 IOPS per 1 GiB, 2.5 MB/s per 1 GiB). Max
	//     IOPS: 9000. Max bandwidth: 500 MB/s.
	//   - `ssd_lowlatency` - SSD storage optimized for low-latency and real-time
	//     processing. Max IOPS: 5000. Average latency: 300 µs. Snapshots and volume
	//     resizing are **not** supported for `ssd_lowlatency`.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// New volume will be created from the image and attached to the instance. Specify
	// `boot_index=0` to boot from this volume.
	//
	// This field can be elided, and will marshal its zero value as "image".
	Source constant.Image `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsVolumeImage) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolumeImage
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsVolumeImage](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Size, SnapshotID, Source are required.
type InstanceNewParamsVolumeSnapshot struct {
	// Volume size in GiB.
	Size int64 `json:"size,required"`
	// Snapshot ID.
	SnapshotID string `json:"snapshot_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// - `0` means that this is the primary boot device;
	// - A unique positive value is set for the secondary bootable devices;
	// - A negative number means that the boot is prohibited.
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// The name of the volume. If not specified, a name will be generated
	// automatically.
	Name param.Opt[string] `json:"name,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Specifies the volume type. If omitted, the type from the source volume will be
	// used by default.
	//
	// Any of "ssd_hiiops", "standard".
	TypeName string `json:"type_name,omitzero"`
	// New volume will be created from the snapshot and attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "snapshot".
	Source constant.Snapshot `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsVolumeSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolumeSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsVolumeSnapshot](
		"TypeName", false, "ssd_hiiops", "standard",
	)
}

// The properties ApptemplateID, Source are required.
type InstanceNewParamsVolumeApptemplate struct {
	// App template ID.
	ApptemplateID string `json:"apptemplate_id,required"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// - `0` means that this is the primary boot device;
	// - A unique positive value is set for the secondary bootable devices;
	// - A negative number means that the boot is prohibited.
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// The name of the volume. If not specified, a name will be generated
	// automatically.
	Name param.Opt[string] `json:"name,omitzero"`
	// Volume size in GiB.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Volume type name. Supported values:
	//
	//   - `standard` - Network SSD block storage offering stable performance with high
	//     random I/O and data reliability (6 IOPS per 1 GiB, 0.4 MB/s per 1 GiB). Max
	//     IOPS: 4500. Max bandwidth: 300 MB/s.
	//   - `ssd_hiiops` - High-performance SSD storage for latency-sensitive
	//     transactional workloads (60 IOPS per 1 GiB, 2.5 MB/s per 1 GiB). Max
	//     IOPS: 9000. Max bandwidth: 500 MB/s.
	//   - `ssd_lowlatency` - SSD storage optimized for low-latency and real-time
	//     processing. Max IOPS: 5000. Average latency: 300 µs. Snapshots and volume
	//     resizing are **not** supported for `ssd_lowlatency`.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// New volume will be created from the app template and attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "apptemplate".
	Source constant.Apptemplate `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsVolumeApptemplate) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolumeApptemplate
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsVolumeApptemplate](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Source, VolumeID are required.
type InstanceNewParamsVolumeExistingVolume struct {
	// Volume ID.
	VolumeID string `json:"volume_id,required" format:"uuid4"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// - `0` means that this is the primary boot device;
	// - A unique positive value is set for the secondary bootable devices;
	// - A negative number means that the boot is prohibited.
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateMap `json:"tags,omitzero"`
	// Existing available volume will be attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "existing-volume".
	Source constant.ExistingVolume `json:"source,required"`
	paramObj
}

func (r InstanceNewParamsVolumeExistingVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolumeExistingVolume
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property ID is required.
type InstanceNewParamsSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name.
	Name string `json:"name,required"`
	paramObj
}

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Only show instances which are able to handle floating address
	AvailableFloating param.Opt[bool] `query:"available_floating,omitzero" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed.
	ChangesBefore param.Opt[time.Time] `query:"changes-before,omitzero" format:"date-time" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed
	// status.
	ChangesSince param.Opt[time.Time] `query:"changes-since,omitzero" format:"date-time" json:"-"`
	// Exclude instances with specified flavor prefix
	ExcludeFlavorPrefix param.Opt[string] `query:"exclude_flavor_prefix,omitzero" json:"-"`
	// Exclude instances with specified security group name
	ExcludeSecgroup param.Opt[string] `query:"exclude_secgroup,omitzero" json:"-"`
	// Filter out instances by flavor_id. Flavor id must match exactly.
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// Filter out instances by flavor_prefix.
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// Include GPU clusters' servers
	IncludeAI param.Opt[bool] `query:"include_ai,omitzero" json:"-"`
	// Include bare metal servers. Please, use `GET /v1/bminstances/` instead
	IncludeBaremetal param.Opt[bool] `query:"include_baremetal,omitzero" json:"-"`
	// Include managed k8s worker nodes
	IncludeK8s param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
	// An IPv4 address to filter results by. Note: partial matches are allowed. For
	// example, searching for 192.168.0.1 will return 192.168.0.1, 192.168.0.10,
	// 192.168.0.110, and so on.
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter instances by name. You can provide a full or partial name, instances with
	// matching names will be returned. For example, entering 'test' will return all
	// instances that contain 'test' in their name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Include only isolated instances
	OnlyIsolated param.Opt[bool] `query:"only_isolated,omitzero" json:"-"`
	// Return bare metals only with external fixed IP addresses.
	OnlyWithFixedExternalIP param.Opt[bool] `query:"only_with_fixed_external_ip,omitzero" json:"-"`
	// Filter result by ddos protection profile name. Effective only with with_ddos set
	// to true.
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs. curl -G --data-urlencode
	// "tag_key_value={"key": "value"}" --url
	// "https://example.com/cloud/v1/resource/1/1"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter the server list result by the UUID of the server. Allowed UUID part
	Uuid param.Opt[string] `query:"uuid,omitzero" json:"-"`
	// Include DDoS profile information in the response when set to `true`. Otherwise,
	// the `ddos_profile` field in the response is `null` by default.
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// Include `interface_name` in the addresses
	WithInterfacesName param.Opt[bool] `query:"with_interfaces_name,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "created.asc", "created.desc", "name.asc", "name.desc".
	OrderBy InstanceListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter result by DDoS protection_status. if parameter is provided. Effective
	// only with with_ddos set to true. (Active, Queued or Error)
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus InstanceListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// Filters instances by status.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "MIGRATING", "PAUSED",
	// "REBOOT", "REBUILD", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED",
	// "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE".
	Status InstanceListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional. Filter by tag values. ?tag_value=value1&tag_value=value2
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with with_ddos set to true. (advanced or basic)
	//
	// Any of "basic", "advanced".
	TypeDDOSProfile InstanceListParamsTypeDDOSProfile `query:"type_ddos_profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field and direction.
type InstanceListParamsOrderBy string

const (
	InstanceListParamsOrderByCreatedAsc  InstanceListParamsOrderBy = "created.asc"
	InstanceListParamsOrderByCreatedDesc InstanceListParamsOrderBy = "created.desc"
	InstanceListParamsOrderByNameAsc     InstanceListParamsOrderBy = "name.asc"
	InstanceListParamsOrderByNameDesc    InstanceListParamsOrderBy = "name.desc"
)

// Filter result by DDoS protection_status. if parameter is provided. Effective
// only with with_ddos set to true. (Active, Queued or Error)
type InstanceListParamsProtectionStatus string

const (
	InstanceListParamsProtectionStatusActive InstanceListParamsProtectionStatus = "Active"
	InstanceListParamsProtectionStatusQueued InstanceListParamsProtectionStatus = "Queued"
	InstanceListParamsProtectionStatusError  InstanceListParamsProtectionStatus = "Error"
)

// Filters instances by status.
type InstanceListParamsStatus string

const (
	InstanceListParamsStatusActive           InstanceListParamsStatus = "ACTIVE"
	InstanceListParamsStatusBuild            InstanceListParamsStatus = "BUILD"
	InstanceListParamsStatusError            InstanceListParamsStatus = "ERROR"
	InstanceListParamsStatusHardReboot       InstanceListParamsStatus = "HARD_REBOOT"
	InstanceListParamsStatusMigrating        InstanceListParamsStatus = "MIGRATING"
	InstanceListParamsStatusPaused           InstanceListParamsStatus = "PAUSED"
	InstanceListParamsStatusReboot           InstanceListParamsStatus = "REBOOT"
	InstanceListParamsStatusRebuild          InstanceListParamsStatus = "REBUILD"
	InstanceListParamsStatusResize           InstanceListParamsStatus = "RESIZE"
	InstanceListParamsStatusRevertResize     InstanceListParamsStatus = "REVERT_RESIZE"
	InstanceListParamsStatusShelved          InstanceListParamsStatus = "SHELVED"
	InstanceListParamsStatusShelvedOffloaded InstanceListParamsStatus = "SHELVED_OFFLOADED"
	InstanceListParamsStatusShutoff          InstanceListParamsStatus = "SHUTOFF"
	InstanceListParamsStatusSoftDeleted      InstanceListParamsStatus = "SOFT_DELETED"
	InstanceListParamsStatusSuspended        InstanceListParamsStatus = "SUSPENDED"
	InstanceListParamsStatusVerifyResize     InstanceListParamsStatus = "VERIFY_RESIZE"
)

// Return bare metals either only with advanced or only basic DDoS protection.
// Effective only with with_ddos set to true. (advanced or basic)
type InstanceListParamsTypeDDOSProfile string

const (
	InstanceListParamsTypeDDOSProfileBasic    InstanceListParamsTypeDDOSProfile = "basic"
	InstanceListParamsTypeDDOSProfileAdvanced InstanceListParamsTypeDDOSProfile = "advanced"
)

type InstanceDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// True if it is required to delete floating IPs assigned to the instance. Can't be
	// used with `floatings`.
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	// Comma separated list of floating ids that should be deleted. Can't be used with
	// `delete_floatings`.
	Floatings param.Opt[string] `query:"floatings,omitzero" json:"-"`
	// Comma separated list of port IDs to be deleted with the instance
	ReservedFixedIPs param.Opt[string] `query:"reserved_fixed_ips,omitzero" json:"-"`
	// Comma separated list of volume IDs to be deleted with the instance
	Volumes param.Opt[string] `query:"volumes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceDeleteParams]'s query parameters as `url.Values`.
func (r InstanceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceActionParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfStartActionInstanceSerializer *InstanceActionParamsBodyStartActionInstanceSerializer `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfBasicActionInstanceSerializer *InstanceActionParamsBodyBasicActionInstanceSerializer `json:",inline"`

	paramObj
}

func (u InstanceActionParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceActionParams](u.OfStartActionInstanceSerializer, u.OfBasicActionInstanceSerializer)
}

// The property Action is required.
type InstanceActionParamsBodyStartActionInstanceSerializer struct {
	// Used on start instance to activate Advanced DDoS profile
	ActivateProfile param.Opt[bool] `json:"activate_profile,omitzero"`
	// Instance action name
	//
	// This field can be elided, and will marshal its zero value as "start".
	Action constant.Start `json:"action,required"`
	paramObj
}

func (r InstanceActionParamsBodyStartActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceActionParamsBodyStartActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Action is required.
type InstanceActionParamsBodyBasicActionInstanceSerializer struct {
	// Instance action name
	//
	// Any of "reboot", "reboot_hard", "resume", "stop", "suspend".
	Action string `json:"action,omitzero,required"`
	paramObj
}

func (r InstanceActionParamsBodyBasicActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceActionParamsBodyBasicActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceActionParamsBodyBasicActionInstanceSerializer](
		"Action", false, "reboot", "reboot_hard", "resume", "stop", "suspend",
	)
}

type InstanceAddToPlacementGroupParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Anti-affinity or affinity or soft-anti-affinity server group ID.
	ServergroupID string `json:"servergroup_id,required"`
	paramObj
}

func (r InstanceAddToPlacementGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAddToPlacementGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceAssignSecurityGroupParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Security group name, applies to all ports
	Name param.Opt[string] `json:"name,omitzero"`
	// Port security groups mapping
	PortsSecurityGroupNames []InstanceAssignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

func (r InstanceAssignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAssignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Port security group names
//
// The properties PortID, SecurityGroupNames are required.
type InstanceAssignSecurityGroupParamsPortsSecurityGroupName struct {
	// Port ID. If None, security groups will be applied to all ports
	PortID param.Opt[string] `json:"port_id,omitzero,required"`
	// List of security group names
	SecurityGroupNames []string `json:"security_group_names,omitzero,required"`
	paramObj
}

func (r InstanceAssignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAssignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceDisablePortSecurityParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceEnablePortSecurityParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceGetConsoleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Console type
	ConsoleType param.Opt[string] `query:"console_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceGetConsoleParams]'s query parameters as
// `url.Values`.
func (r InstanceGetConsoleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceRemoveFromPlacementGroupParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type InstanceResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flavor ID
	FlavorID string `json:"flavor_id,required"`
	paramObj
}

func (r InstanceResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceUnassignSecurityGroupParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Security group name, applies to all ports
	Name param.Opt[string] `json:"name,omitzero"`
	// Port security groups mapping
	PortsSecurityGroupNames []InstanceUnassignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

func (r InstanceUnassignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUnassignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Port security group names
//
// The properties PortID, SecurityGroupNames are required.
type InstanceUnassignSecurityGroupParamsPortsSecurityGroupName struct {
	// Port ID. If None, security groups will be applied to all ports
	PortID param.Opt[string] `json:"port_id,omitzero,required"`
	// List of security group names
	SecurityGroupNames []string `json:"security_group_names,omitzero,required"`
	paramObj
}

func (r InstanceUnassignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUnassignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}
