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
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// Instances are cloud virtual machines with configurable CPU, memory, storage, and
// networking, supporting various operating systems and workloads.
//
// InstancesRegionalService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstancesRegionalService] method instead.
type InstancesRegionalService struct {
	Options []option.RequestOption
	// Instance flavors define available CPU, memory, and disk configurations for
	// creating cloud instances.
	Flavors    InstancesRegionalFlavorService
	Interfaces InstancesRegionalInterfaceService
}

// NewInstancesRegionalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstancesRegionalService(opts ...option.RequestOption) (r InstancesRegionalService) {
	r = InstancesRegionalService{}
	r.Options = opts
	r.Flavors = NewInstancesRegionalFlavorService(opts...)
	r.Interfaces = NewInstancesRegionalInterfaceService(opts...)
	return
}

// Use
// [POST /v1/projects/{`project_id`}/instances](#operation/GlobalInstanceListView.post)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Create an instance with specified configuration.
//
// How to get access:
//
// For Linux,
//
//   - Use the `user_data` field to provide a
//     [cloud-init script](https://cloudinit.readthedocs.io/en/latest/reference/examples.html)
//     in base64 to apply configurations to the instance.
//   - Specify the `username` and `password` to create a new user.
//   - When only `password` is provided, it is set as the password for the default
//     user of the image.
//   - The `user_data` is ignored when the `password` is specified.
//
// For Windows,
//
//   - Use the `user_data` field to provide a
//     [cloudbase-init script](https://cloudbase-init.readthedocs.io/en/latest/userdata.html#cloud-config)
//     in base64 to create new users on Windows.
//   - Use the `password` field to set the password for the 'Admin' user on Windows.
//   - The password of the Admin user cannot be updated via `user_data`.
//   - The `username` cannot be specified in the request.
func (r *InstancesRegionalService) New(ctx context.Context, params InstancesRegionalNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Rename instance or update tags
func (r *InstancesRegionalService) Update(ctx context.Context, instanceID string, params InstancesRegionalUpdateParams, opts ...option.RequestOption) (res *InstanceRegional, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Use
// [GET /v1/projects/{`project_id`}/instances](#operation/GlobalInstanceListView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// List all instances in the specified project and region. Results can be filtered
// by various parameters like name, status, and IP address.
func (r *InstancesRegionalService) List(ctx context.Context, params InstancesRegionalListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InstanceRegional], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
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

// Use
// [GET /v1/projects/{`project_id`}/instances](#operation/GlobalInstanceListView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// List all instances in the specified project and region. Results can be filtered
// by various parameters like name, status, and IP address.
func (r *InstancesRegionalService) ListAutoPaging(ctx context.Context, params InstancesRegionalListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InstanceRegional] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Use
// [DELETE /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.delete)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Delete the instance and its ports. A reserved fixed IP stays, unless
// `reserved_fixed_ips` names it. A port that you
// [create as a port](#operation/GlobalPortListView.post) is not a reserved fixed
// IP, so this delete removes it.
//
// To keep the ports, use
// [`DELETE /v1/projects/{project_id}/instances/{instance_id}`](#operation/GlobalInstanceDetailView.delete).
// That delete keeps every port by default.
func (r *InstancesRegionalService) Delete(ctx context.Context, instanceID string, params InstancesRegionalDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Use
// [POST /v1/projects/{`project_id`}/instances/{`instance_id`}/action](#operation/GlobalInstanceActionView.post)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// The action can be one of: start, stop, reboot, `reboot_hard`, suspend or resume.
// Suspend and resume are not available for bare metal instances.
func (r *InstancesRegionalService) Action(ctx context.Context, instanceID string, params InstancesRegionalActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/%s/action", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. Send `placement_group_id`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Add an instance to a placement group. The instance must not already be in a
// placement group. Bare metal servers do not support placement groups.
func (r *InstancesRegionalService) AddToPlacementGroup(ctx context.Context, instanceID string, params InstancesRegionalAddToPlacementGroupParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/put_into_servergroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/ports/{`port_id`}](#operation/GlobalPortDetailView.patch)
// instead. Send `security_group_ids`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Assign the security group to the server. To assign multiple security groups to
// all ports, use the NULL value for the `port_id` field
func (r *InstancesRegionalService) AssignSecurityGroup(ctx context.Context, instanceID string, params InstancesRegionalAssignSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/addsecuritygroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Deprecated. Use PATCH /v1/ports/{`project_id`}/{`region_id`}/{`port_id`}
// instead.
//
// Deprecated: deprecated
func (r *InstancesRegionalService) DisablePortSecurity(ctx context.Context, portID string, body InstancesRegionalDisablePortSecurityParams, opts ...option.RequestOption) (res *InstanceRegionalInterface, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/disable_port_security", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Deprecated. Use PATCH /v1/ports/{`project_id`}/{`region_id`}/{`port_id`}
// instead.
//
// Deprecated: deprecated
func (r *InstancesRegionalService) EnablePortSecurity(ctx context.Context, portID string, body InstancesRegionalEnablePortSecurityParams, opts ...option.RequestOption) (res *InstanceRegionalInterface, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/enable_port_security", body.ProjectID.Value, body.RegionID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Use
// [GET /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Retrieve detailed information about a specific instance. The response content
// language for `ddos_profile` can be controlled via the 'language' cookie
// parameter.
//
// **Cookie Parameters**:
//
//   - `language` (str, optional): Language for the response content. Affects the
//     `ddos_profile` field. Supported values:
//   - `'en'` (default)
//   - `'de'`
func (r *InstancesRegionalService) Get(ctx context.Context, instanceID string, query InstancesRegionalGetParams, opts ...option.RequestOption) (res *InstanceRegional, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Use
// [GET /v1/projects/{`project_id`}/instances/{`instance_id`}/console](#operation/GlobalInstanceConsoleView.get)
// instead. For a bare metal server, use
// [Get bare metal server console URL](#operation/GlobalBaremetalConsoleView.get).
// For a basic VM, use
// [Get basic VM console URL](#operation/GlobalBasicVmConsoleView.get). This route
// continues to work. Later, we will deprecate it and then remove it. We will tell
// you in advance.
//
// Returns a URL to access the remote console for an instance. Supported console
// types depend on instance type: novnc and spice for virtual instances, vnc and
// serial for bare metal instances.
func (r *InstancesRegionalService) GetConsole(ctx context.Context, instanceID string, params InstancesRegionalGetConsoleParams, opts ...option.RequestOption) (res *Console, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/get_console", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. Send `placement_group_id`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Remove an instance from its current placement group. The instance must be in a
// placement group to be removed. Bare metal servers do not support placement
// groups.
func (r *InstancesRegionalService) RemoveFromPlacementGroup(ctx context.Context, instanceID string, body InstancesRegionalRemoveFromPlacementGroupParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/remove_from_servergroup", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/instances/{`instance_id`}](#operation/GlobalInstanceDetailView.patch)
// instead. Send `flavor_name`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Change flavor of the instance.
func (r *InstancesRegionalService) Resize(ctx context.Context, instanceID string, params InstancesRegionalResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/changeflavor", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/ports/{`port_id`}](#operation/GlobalPortDetailView.patch)
// instead. Send `security_group_ids`. This route continues to work. Later, we will
// deprecate it and then remove it. We will tell you in advance.
//
// Un-assign the security group to the server. To un-assign multiple security
// groups to all ports, use the NULL value for the `port_id` field
func (r *InstancesRegionalService) UnassignSecurityGroup(ctx context.Context, instanceID string, params InstancesRegionalUnassignSecurityGroupParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return err
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/delsecuritygroup", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type InstanceRegionalInterface struct {
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs" api:"required"`
	// Bodies of floating IPs that are NAT-ing IPs of this port
	FloatingipDetails []FloatingIP `json:"floatingip_details" api:"required"`
	// Interface name. Null when the interface has no name set.
	InterfaceName string `json:"interface_name" api:"required"`
	// IP addresses assigned to this port
	IPAssignments []IPAssignment `json:"ip_assignments" api:"required"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address" api:"required"`
	// Body of the network this port is attached to
	NetworkDetails NetworkDetails `json:"network_details" api:"required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Port security status
	PortSecurityEnabled bool `json:"port_security_enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		FloatingipDetails   respjson.Field
		InterfaceName       respjson.Field
		IPAssignments       respjson.Field
		MacAddress          respjson.Field
		NetworkDetails      respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		PortSecurityEnabled respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceRegionalInterface) RawJSON() string { return r.JSON.raw }
func (r *InstanceRegionalInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// The flavor of the instance.
	Flavor string `json:"flavor" api:"required"`
	// A list of network interfaces for the instance. You can create one or more
	// interfaces - private, public, or both.
	Interfaces []InstancesRegionalNewParamsInterfaceUnion `json:"interfaces,omitzero" api:"required"`
	// List of volumes that will be attached to the instance.
	Volumes []InstancesRegionalNewParamsVolumeUnion `json:"volumes,omitzero" api:"required"`
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
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
	// `user_data`: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// For Windows instances, 'username' cannot be specified. Use 'password' field to
	// set the password for the 'Admin' user on Windows.
	Username param.Opt[string] `json:"username,omitzero"`
	// Parameters for the application template if creating the instance from an
	// `apptemplate`.
	Configuration map[string]any `json:"configuration,omitzero"`
	// Deprecated. Use per-interface `security_groups` inside `interfaces[]` instead.
	// Cannot be combined with per-interface `security_groups`. If omitted everywhere,
	// the project's default security group is applied.
	SecurityGroups []InstancesRegionalNewParamsSecurityGroup `json:"security_groups,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r InstancesRegionalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstancesRegionalNewParamsInterfaceUnion struct {
	OfExternal        *InstancesRegionalNewParamsInterfaceExternal        `json:",omitzero,inline"`
	OfSubnet          *InstancesRegionalNewParamsInterfaceSubnet          `json:",omitzero,inline"`
	OfAnySubnet       *InstancesRegionalNewParamsInterfaceAnySubnet       `json:",omitzero,inline"`
	OfReservedFixedIP *InstancesRegionalNewParamsInterfaceReservedFixedIP `json:",omitzero,inline"`
	paramUnion
}

func (u InstancesRegionalNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}
func (u *InstancesRegionalNewParamsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InstancesRegionalNewParamsInterfaceUnion) asAny() any {
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
func (u InstancesRegionalNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.Valid() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfReservedFixedIP; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceUnion) GetType() *string {
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
func (u InstancesRegionalNewParamsInterfaceUnion) GetInterfaceName() *string {
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
func (u InstancesRegionalNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceUnion) GetNetworkID() *string {
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
func (u InstancesRegionalNewParamsInterfaceUnion) GetSecurityGroups() (res instancesRegionalNewParamsInterfaceUnionSecurityGroups) {
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
// [_[]InstancesRegionalNewParamsInterfaceExternalSecurityGroup],
// [_[]InstancesRegionalNewParamsInterfaceSubnetSecurityGroup],
// [_[]InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup],
// [_[]InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup]
type instancesRegionalNewParamsInterfaceUnionSecurityGroups struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]cloud.InstancesRegionalNewParamsInterfaceExternalSecurityGroup:
//	case *[]cloud.InstancesRegionalNewParamsInterfaceSubnetSecurityGroup:
//	case *[]cloud.InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup:
//	case *[]cloud.InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instancesRegionalNewParamsInterfaceUnionSecurityGroups) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u InstancesRegionalNewParamsInterfaceUnion) GetFloatingIP() (res instancesRegionalNewParamsInterfaceUnionFloatingIP) {
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
// [*InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew],
// [*InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting],
// [*InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew],
// [*InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting],
// [*InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew],
// [*InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting]
type instancesRegionalNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew:
//	case *cloud.InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting:
//	case *cloud.InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew:
//	case *cloud.InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting:
//	case *cloud.InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew:
//	case *cloud.InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instancesRegionalNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u instancesRegionalNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetSource()
	case *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetSource()
	case *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u instancesRegionalNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstancesRegionalNewParamsInterfaceUnion](
		"type",
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceExternal]("external"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceSubnet]("subnet"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceAnySubnet]("any_subnet"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceReservedFixedIP]("reserved_fixed_ip"),
	)
}

// Instance will be attached to default external network
//
// The property Type is required.
type InstancesRegionalNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []InstancesRegionalNewParamsInterfaceExternalSecurityGroup `json:"security_groups,omitzero"`
	// A public IP address will be assigned to the instance.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type" default:"external"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalNewParamsInterfaceExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The instance will get an IP address from the selected network. If you choose to
// add a floating IP, the instance will be reachable from the internet. Otherwise,
// it will only have a private IP within the network.
//
// The properties NetworkID, SubnetID, Type are required.
type InstancesRegionalNewParamsInterfaceSubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// The instance will get an IP address from this subnet.
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []InstancesRegionalNewParamsInterfaceSubnetSecurityGroup `json:"security_groups,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type" default:"subnet"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion struct {
	OfNew      *InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstancesRegionalNewParamsInterfaceSubnetFloatingIPUnion](
		"source",
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew]("new"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting]("existing"),
	)
}

func NewInstancesRegionalNewParamsInterfaceSubnetFloatingIPNew() InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew {
	return InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstancesRegionalNewParamsInterfaceSubnetFloatingIPNew].
type InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceSubnetFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceSubnetFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalNewParamsInterfaceSubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceSubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceSubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceSubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type InstancesRegionalNewParamsInterfaceAnySubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// Instance will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type" default:"any_subnet"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion struct {
	OfNew      *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPUnion](
		"source",
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew]("new"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting]("existing"),
	)
}

func NewInstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew() InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew {
	return InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew].
type InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceAnySubnetFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortID, Type are required.
type InstancesRegionalNewParamsInterfaceReservedFixedIP struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID string `json:"port_id" api:"required"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup `json:"security_groups,omitzero"`
	// An existing available reserved fixed IP will be attached to the instance. If the
	// reserved IP is not public and you choose to add a floating IP, the instance will
	// be accessible from the internet.
	//
	// This field can be elided, and will marshal its zero value as
	// "reserved_fixed_ip".
	Type constant.ReservedFixedIP `json:"type" default:"reserved_fixed_ip"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion struct {
	OfNew      *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPUnion](
		"source",
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew]("new"),
		apijson.Discriminator[InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting]("existing"),
	)
}

func NewInstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew() InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew {
	return InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewInstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew].
type InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceReservedFixedIPFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsInterfaceReservedFixedIPSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstancesRegionalNewParamsVolumeUnion struct {
	OfNewVolume      *InstancesRegionalNewParamsVolumeNewVolume      `json:",omitzero,inline"`
	OfImage          *InstancesRegionalNewParamsVolumeImage          `json:",omitzero,inline"`
	OfSnapshot       *InstancesRegionalNewParamsVolumeSnapshot       `json:",omitzero,inline"`
	OfApptemplate    *InstancesRegionalNewParamsVolumeApptemplate    `json:",omitzero,inline"`
	OfExistingVolume *InstancesRegionalNewParamsVolumeExistingVolume `json:",omitzero,inline"`
	paramUnion
}

func (u InstancesRegionalNewParamsVolumeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNewVolume,
		u.OfImage,
		u.OfSnapshot,
		u.OfApptemplate,
		u.OfExistingVolume)
}
func (u *InstancesRegionalNewParamsVolumeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InstancesRegionalNewParamsVolumeUnion) asAny() any {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetImageID() *string {
	if vt := u.OfImage; vt != nil {
		return &vt.ImageID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsVolumeUnion) GetSnapshotID() *string {
	if vt := u.OfSnapshot; vt != nil {
		return &vt.SnapshotID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsVolumeUnion) GetApptemplateID() *string {
	if vt := u.OfApptemplate; vt != nil {
		return &vt.ApptemplateID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsVolumeUnion) GetVolumeID() *string {
	if vt := u.OfExistingVolume; vt != nil {
		return &vt.VolumeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstancesRegionalNewParamsVolumeUnion) GetSize() *int64 {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetSource() *string {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetAttachmentTag() *string {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetDeleteOnTermination() *bool {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetName() *string {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetTypeName() *string {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetBootIndex() *int64 {
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
func (u InstancesRegionalNewParamsVolumeUnion) GetTags() map[string]string {
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
	apijson.RegisterUnion[InstancesRegionalNewParamsVolumeUnion](
		"source",
		apijson.Discriminator[InstancesRegionalNewParamsVolumeNewVolume]("new-volume"),
		apijson.Discriminator[InstancesRegionalNewParamsVolumeImage]("image"),
		apijson.Discriminator[InstancesRegionalNewParamsVolumeSnapshot]("snapshot"),
		apijson.Discriminator[InstancesRegionalNewParamsVolumeApptemplate]("apptemplate"),
		apijson.Discriminator[InstancesRegionalNewParamsVolumeExistingVolume]("existing-volume"),
	)
}

// The properties Size, Source are required.
type InstancesRegionalNewParamsVolumeNewVolume struct {
	// Volume size in GiB.
	Size int64 `json:"size" api:"required"`
	// Block device attachment tag (not exposed in the normal tags)
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// Set to `true` to automatically delete the volume when the instance is deleted.
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// The name of the volume. If not specified, a name will be generated
	// automatically.
	Name param.Opt[string] `json:"name,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
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
	Source constant.NewVolume `json:"source" default:"new-volume"`
	paramObj
}

func (r InstancesRegionalNewParamsVolumeNewVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsVolumeNewVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsVolumeNewVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalNewParamsVolumeNewVolume](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties ImageID, Source are required.
type InstancesRegionalNewParamsVolumeImage struct {
	// Image ID.
	ImageID string `json:"image_id" api:"required" format:"uuid4"`
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
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
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
	Source constant.Image `json:"source" default:"image"`
	paramObj
}

func (r InstancesRegionalNewParamsVolumeImage) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsVolumeImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsVolumeImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalNewParamsVolumeImage](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Size, SnapshotID, Source are required.
type InstancesRegionalNewParamsVolumeSnapshot struct {
	// Volume size in GiB.
	Size int64 `json:"size" api:"required"`
	// Snapshot ID.
	SnapshotID string `json:"snapshot_id" api:"required" format:"uuid4"`
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
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// Specifies the volume type. If omitted, the type from the source volume will be
	// used by default.
	//
	// Any of "ssd_hiiops", "standard".
	TypeName string `json:"type_name,omitzero"`
	// New volume will be created from the snapshot and attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "snapshot".
	Source constant.Snapshot `json:"source" default:"snapshot"`
	paramObj
}

func (r InstancesRegionalNewParamsVolumeSnapshot) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsVolumeSnapshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsVolumeSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalNewParamsVolumeSnapshot](
		"type_name", "ssd_hiiops", "standard",
	)
}

// The properties ApptemplateID, Source are required.
type InstancesRegionalNewParamsVolumeApptemplate struct {
	// App template ID.
	ApptemplateID string `json:"apptemplate_id" api:"required"`
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
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
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
	Source constant.Apptemplate `json:"source" default:"apptemplate"`
	paramObj
}

func (r InstancesRegionalNewParamsVolumeApptemplate) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsVolumeApptemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsVolumeApptemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalNewParamsVolumeApptemplate](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties Source, VolumeID are required.
type InstancesRegionalNewParamsVolumeExistingVolume struct {
	// Volume ID.
	VolumeID string `json:"volume_id" api:"required" format:"uuid4"`
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
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// Existing available volume will be attached to the instance.
	//
	// This field can be elided, and will marshal its zero value as "existing-volume".
	Source constant.ExistingVolume `json:"source" default:"existing-volume"`
	paramObj
}

func (r InstancesRegionalNewParamsVolumeExistingVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsVolumeExistingVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsVolumeExistingVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InstancesRegionalNewParamsSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstancesRegionalNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalNewParamsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalNewParamsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7396). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//   - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and 'cost_center' while removing 'deprecated_tag',
	//     preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r InstancesRegionalUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
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
	// Filter out instances by `flavor_id`. Flavor id must match exactly.
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// Filter out instances by `flavor_prefix`.
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// Include GPU clusters' servers
	IncludeAI param.Opt[bool] `query:"include_ai,omitzero" json:"-"`
	// Include bare metal servers. Please, use `GET /v1/bminstances/` instead
	IncludeBaremetal param.Opt[bool] `query:"include_baremetal,omitzero" json:"-"`
	// Include managed k8s worker nodes
	IncludeK8S param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
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
	// Filter result by ddos protection profile name. Effective only with `with_ddos`
	// set to true.
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
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
	// Any of "created.asc", "created.desc", "name.asc", "name.desc", "status.asc",
	// "status.desc".
	OrderBy InstancesRegionalListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter result by DDoS `protection_status`. if parameter is provided. Effective
	// only with `with_ddos` set to true. (Active, Queued or Error)
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus InstancesRegionalListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// Filters instances by status.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "MIGRATING", "PAUSED",
	// "REBOOT", "REBUILD", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED",
	// "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE".
	Status InstancesRegionalListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional. Filter by tag values. ?`tag_value`=value1&`tag_value`=value2
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with `with_ddos` set to true. (advanced or basic)
	//
	// Any of "basic", "advanced".
	TypeDDOSProfile InstancesRegionalListParamsTypeDDOSProfile `query:"type_ddos_profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstancesRegionalListParams]'s query parameters as
// `url.Values`.
func (r InstancesRegionalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order by field and direction.
type InstancesRegionalListParamsOrderBy string

const (
	InstancesRegionalListParamsOrderByCreatedAsc  InstancesRegionalListParamsOrderBy = "created.asc"
	InstancesRegionalListParamsOrderByCreatedDesc InstancesRegionalListParamsOrderBy = "created.desc"
	InstancesRegionalListParamsOrderByNameAsc     InstancesRegionalListParamsOrderBy = "name.asc"
	InstancesRegionalListParamsOrderByNameDesc    InstancesRegionalListParamsOrderBy = "name.desc"
	InstancesRegionalListParamsOrderByStatusAsc   InstancesRegionalListParamsOrderBy = "status.asc"
	InstancesRegionalListParamsOrderByStatusDesc  InstancesRegionalListParamsOrderBy = "status.desc"
)

// Filter result by DDoS `protection_status`. if parameter is provided. Effective
// only with `with_ddos` set to true. (Active, Queued or Error)
type InstancesRegionalListParamsProtectionStatus string

const (
	InstancesRegionalListParamsProtectionStatusActive InstancesRegionalListParamsProtectionStatus = "Active"
	InstancesRegionalListParamsProtectionStatusQueued InstancesRegionalListParamsProtectionStatus = "Queued"
	InstancesRegionalListParamsProtectionStatusError  InstancesRegionalListParamsProtectionStatus = "Error"
)

// Filters instances by status.
type InstancesRegionalListParamsStatus string

const (
	InstancesRegionalListParamsStatusActive           InstancesRegionalListParamsStatus = "ACTIVE"
	InstancesRegionalListParamsStatusBuild            InstancesRegionalListParamsStatus = "BUILD"
	InstancesRegionalListParamsStatusError            InstancesRegionalListParamsStatus = "ERROR"
	InstancesRegionalListParamsStatusHardReboot       InstancesRegionalListParamsStatus = "HARD_REBOOT"
	InstancesRegionalListParamsStatusMigrating        InstancesRegionalListParamsStatus = "MIGRATING"
	InstancesRegionalListParamsStatusPaused           InstancesRegionalListParamsStatus = "PAUSED"
	InstancesRegionalListParamsStatusReboot           InstancesRegionalListParamsStatus = "REBOOT"
	InstancesRegionalListParamsStatusRebuild          InstancesRegionalListParamsStatus = "REBUILD"
	InstancesRegionalListParamsStatusResize           InstancesRegionalListParamsStatus = "RESIZE"
	InstancesRegionalListParamsStatusRevertResize     InstancesRegionalListParamsStatus = "REVERT_RESIZE"
	InstancesRegionalListParamsStatusShelved          InstancesRegionalListParamsStatus = "SHELVED"
	InstancesRegionalListParamsStatusShelvedOffloaded InstancesRegionalListParamsStatus = "SHELVED_OFFLOADED"
	InstancesRegionalListParamsStatusShutoff          InstancesRegionalListParamsStatus = "SHUTOFF"
	InstancesRegionalListParamsStatusSoftDeleted      InstancesRegionalListParamsStatus = "SOFT_DELETED"
	InstancesRegionalListParamsStatusSuspended        InstancesRegionalListParamsStatus = "SUSPENDED"
	InstancesRegionalListParamsStatusVerifyResize     InstancesRegionalListParamsStatus = "VERIFY_RESIZE"
)

// Return bare metals either only with advanced or only basic DDoS protection.
// Effective only with `with_ddos` set to true. (advanced or basic)
type InstancesRegionalListParamsTypeDDOSProfile string

const (
	InstancesRegionalListParamsTypeDDOSProfileBasic    InstancesRegionalListParamsTypeDDOSProfile = "basic"
	InstancesRegionalListParamsTypeDDOSProfileAdvanced InstancesRegionalListParamsTypeDDOSProfile = "advanced"
)

type InstancesRegionalDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
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

// URLQuery serializes [InstancesRegionalDeleteParams]'s query parameters as
// `url.Values`.
func (r InstancesRegionalDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstancesRegionalActionParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfStartActionInstanceSerializer *InstancesRegionalActionParamsBodyStartActionInstanceSerializer `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfBasicActionInstanceSerializer *InstancesRegionalActionParamsBodyBasicActionInstanceSerializer `json:",inline"`

	paramObj
}

func (u InstancesRegionalActionParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartActionInstanceSerializer, u.OfBasicActionInstanceSerializer)
}
func (r *InstancesRegionalActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type InstancesRegionalActionParamsBodyStartActionInstanceSerializer struct {
	// Used on start instance to activate Advanced DDoS profile
	ActivateProfile param.Opt[bool] `json:"activate_profile,omitzero"`
	// Instance action name
	//
	// This field can be elided, and will marshal its zero value as "start".
	Action constant.Start `json:"action" default:"start"`
	paramObj
}

func (r InstancesRegionalActionParamsBodyStartActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalActionParamsBodyStartActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalActionParamsBodyStartActionInstanceSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type InstancesRegionalActionParamsBodyBasicActionInstanceSerializer struct {
	// Instance action name
	//
	// Any of "reboot", "reboot_hard", "resume", "stop", "suspend".
	Action string `json:"action,omitzero" api:"required"`
	paramObj
}

func (r InstancesRegionalActionParamsBodyBasicActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalActionParamsBodyBasicActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalActionParamsBodyBasicActionInstanceSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstancesRegionalActionParamsBodyBasicActionInstanceSerializer](
		"action", "reboot", "reboot_hard", "resume", "stop", "suspend",
	)
}

type InstancesRegionalAddToPlacementGroupParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Anti-affinity or affinity or soft-anti-affinity placement group ID.
	ServergroupID string `json:"servergroup_id" api:"required"`
	paramObj
}

func (r InstancesRegionalAddToPlacementGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalAddToPlacementGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalAddToPlacementGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalAssignSecurityGroupParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Security group name, applies to all ports
	Name param.Opt[string] `json:"name,omitzero"`
	// Port security groups mapping
	PortsSecurityGroupNames []InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

func (r InstancesRegionalAssignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalAssignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalAssignSecurityGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortID, SecurityGroupNames are required.
type InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName struct {
	// Port ID. If None, security groups will be applied to all ports
	PortID param.Opt[string] `json:"port_id,omitzero" api:"required"`
	// List of security group names
	SecurityGroupNames []string `json:"security_group_names,omitzero" api:"required"`
	paramObj
}

func (r InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalAssignSecurityGroupParamsPortsSecurityGroupName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalDisablePortSecurityParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstancesRegionalEnablePortSecurityParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstancesRegionalGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstancesRegionalGetConsoleParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Console type
	//
	// Any of "novnc", "serial", "spice", "vnc".
	ConsoleType InstancesRegionalGetConsoleParamsConsoleType `query:"console_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstancesRegionalGetConsoleParams]'s query parameters as
// `url.Values`.
func (r InstancesRegionalGetConsoleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Console type
type InstancesRegionalGetConsoleParamsConsoleType string

const (
	InstancesRegionalGetConsoleParamsConsoleTypeNovnc  InstancesRegionalGetConsoleParamsConsoleType = "novnc"
	InstancesRegionalGetConsoleParamsConsoleTypeSerial InstancesRegionalGetConsoleParamsConsoleType = "serial"
	InstancesRegionalGetConsoleParamsConsoleTypeSpice  InstancesRegionalGetConsoleParamsConsoleType = "spice"
	InstancesRegionalGetConsoleParamsConsoleTypeVnc    InstancesRegionalGetConsoleParamsConsoleType = "vnc"
)

type InstancesRegionalRemoveFromPlacementGroupParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstancesRegionalResizeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Flavor ID
	FlavorID string `json:"flavor_id" api:"required"`
	paramObj
}

func (r InstancesRegionalResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalUnassignSecurityGroupParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Security group name, applies to all ports
	Name param.Opt[string] `json:"name,omitzero"`
	// Port security groups mapping
	PortsSecurityGroupNames []InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

func (r InstancesRegionalUnassignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalUnassignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalUnassignSecurityGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortID, SecurityGroupNames are required.
type InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName struct {
	// Port ID. If None, security groups will be applied to all ports
	PortID param.Opt[string] `json:"port_id,omitzero" api:"required"`
	// List of security group names
	SecurityGroupNames []string `json:"security_group_names,omitzero" api:"required"`
	paramObj
}

func (r InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstancesRegionalUnassignSecurityGroupParamsPortsSecurityGroupName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
