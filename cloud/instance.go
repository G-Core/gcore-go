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
	"time"

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

// Create one or many instances or basic VMs. For Linux instances, use the
// 'username' and 'password' to create a new user. When only 'password' is
// provided, it is set as the password for the default user of the image. The
// 'user_data' is ignored when the 'password' is specified. Use the 'user_data'
// field to provide a cloud-init script in base64 to apply configurations to the
// instance. For Windows instances, the 'username' cannot be specified in the
// request. Use the 'password' field to set the password for the 'Admin' user on
// Windows. Use the 'user_data' field to provide a cloudbase-init script in base64
// to create new users on Windows. The password of the Admin user cannot be updated
// via 'user_data'.
func (r *InstanceService) New(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// '#/components/schemas/InstanceInterfaceSerializer'
// "$.components.schemas.InstanceInterfaceSerializer"
type InstanceInterface struct {
	// '#/components/schemas/InstanceInterfaceSerializer/properties/allowed_address_pairs'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.allowed_address_pairs"
	AllowedAddressPairs []InstanceInterfaceAllowedAddressPair `json:"allowed_address_pairs,required"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/floatingip_details'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.floatingip_details"
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/ip_assignments'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.ip_assignments"
	IPAssignments []InstanceInterfaceIPAssignment `json:"ip_assignments,required"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/network_details'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.network_details"
	NetworkDetails InstanceInterfaceNetworkDetails `json:"network_details,required"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/network_id'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/port_id'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/port_security_enabled'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,nullable"`
	// '#/components/schemas/InstanceInterfaceSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.InstanceInterfaceSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AllowedAddressPairs resp.Field
		FloatingipDetails   resp.Field
		IPAssignments       resp.Field
		NetworkDetails      resp.Field
		NetworkID           resp.Field
		PortID              resp.Field
		PortSecurityEnabled resp.Field
		InterfaceName       resp.Field
		MacAddress          resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterface) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSerializer/properties/allowed_address_pairs/items'
// "$.components.schemas.InstanceInterfaceSerializer.properties.allowed_address_pairs.items"
type InstanceInterfaceAllowedAddressPair struct {
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		MacAddress  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterfaceAllowedAddressPair) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterfaceAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSerializer/properties/ip_assignments/items'
// "$.components.schemas.InstanceInterfaceSerializer.properties.ip_assignments.items"
type InstanceInterfaceIPAssignment struct {
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/ip_address'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSerializer/properties/network_details'
// "$.components.schemas.InstanceInterfaceSerializer.properties.network_details"
type InstanceInterfaceNetworkDetails struct {
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/created_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/external'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/name'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/port_security_enabled'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region_id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/shared'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.shared"
	Shared bool `json:"shared,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/tags'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/type'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.type"
	Type string `json:"type,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/updated_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/default/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties['default'].anyOf[0]"
	Default bool `json:"default,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/mtu'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.mtu"
	Mtu int64 `json:"mtu"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/segmentation_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.segmentation_id.anyOf[0]"
	SegmentationID int64 `json:"segmentation_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/subnets/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.subnets.anyOf[0]"
	Subnets []Subnet `json:"subnets,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		CreatedAt           resp.Field
		External            resp.Field
		Name                resp.Field
		PortSecurityEnabled resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Shared              resp.Field
		Tags                resp.Field
		Type                resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		Default             resp.Field
		Mtu                 resp.Field
		ProjectID           resp.Field
		SegmentationID      resp.Field
		Subnets             resp.Field
		TaskID              resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterfaceNetworkDetails) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterfaceNetworkDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceNewParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v2/instances/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v2/instances/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/flavor'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.flavor"
	Flavor string `json:"flavor,required"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/interfaces'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.interfaces"
	Interfaces []InstanceNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/volumes'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.volumes"
	Volumes []InstanceNewParamsVolume `json:"volumes,omitzero,required"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/keypair_name/anyOf/0'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.keypair_name.anyOf[0]"
	KeypairName param.Opt[string] `json:"keypair_name,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/allow_app_ports'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.allow_app_ports"
	AllowAppPorts param.Opt[bool] `json:"allow_app_ports,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/password'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.password"
	Password param.Opt[string] `json:"password,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/servergroup_id'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.servergroup_id"
	ServergroupID param.Opt[string] `json:"servergroup_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/user_data'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.user_data"
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/username'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.username"
	Username param.Opt[string] `json:"username,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/configuration/anyOf/0'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.configuration.anyOf[0]"
	Configuration any `json:"configuration,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/name_templates'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.name_templates"
	NameTemplates []string `json:"name_templates,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/names'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.names"
	Names []string `json:"names,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/security_groups'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.security_groups"
	SecurityGroups []InstanceNewParamsSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/CreateInstanceSerializerV2/properties/tags'
	// "$.components.schemas.CreateInstanceSerializerV2.properties.tags"
	Tags TagUpdateList `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceUnion struct {
	OfNewInterfaceExternalSerializerPydantic           *InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic           `json:",omitzero,inline"`
	OfNewInterfaceSpecificSubnetFipSerializerPydantic  *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic  `json:",omitzero,inline"`
	OfNewInterfaceAnySubnetFipSerializerPydantic       *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic       `json:",omitzero,inline"`
	OfNewInterfaceReservedFixedIPFipSerializerPydantic *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u InstanceNewParamsInterfaceUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u InstanceNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceUnion](u.OfNewInterfaceExternalSerializerPydantic, u.OfNewInterfaceSpecificSubnetFipSerializerPydantic, u.OfNewInterfaceAnySubnetFipSerializerPydantic, u.OfNewInterfaceReservedFixedIPFipSerializerPydantic)
}

func (u *InstanceNewParamsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfNewInterfaceExternalSerializerPydantic) {
		return u.OfNewInterfaceExternalSerializerPydantic
	} else if !param.IsOmitted(u.OfNewInterfaceSpecificSubnetFipSerializerPydantic) {
		return u.OfNewInterfaceSpecificSubnetFipSerializerPydantic
	} else if !param.IsOmitted(u.OfNewInterfaceAnySubnetFipSerializerPydantic) {
		return u.OfNewInterfaceAnySubnetFipSerializerPydantic
	} else if !param.IsOmitted(u.OfNewInterfaceReservedFixedIPFipSerializerPydantic) {
		return u.OfNewInterfaceReservedFixedIPFipSerializerPydantic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil && vt.IPAddress.IsPresent() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetType() *string {
	if vt := u.OfNewInterfaceExternalSerializerPydantic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfNewInterfaceExternalSerializerPydantic; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfNewInterfaceExternalSerializerPydantic; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetPortGroup() *int64 {
	if vt := u.OfNewInterfaceExternalSerializerPydantic; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil {
		return (*string)(&vt.NetworkID)
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil {
		return (*string)(&vt.NetworkID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u InstanceNewParamsInterfaceUnion) GetSecurityGroups() (res instanceNewParamsInterfaceUnionSecurityGroups) {
	if vt := u.OfNewInterfaceExternalSerializerPydantic; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil {
		res.any = &vt.SecurityGroups
	}
	return
}

// Can have the runtime types
// [_[]InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup],
// [_[]InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup],
// [_[]InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup],
// [_[]InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup]
type instanceNewParamsInterfaceUnionSecurityGroups struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]cloud.InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup:
//	case *[]cloud.InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instanceNewParamsInterfaceUnionSecurityGroups) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u InstanceNewParamsInterfaceUnion) GetFloatingIP() (res instanceNewParamsInterfaceUnionFloatingIP) {
	if vt := u.OfNewInterfaceSpecificSubnetFipSerializerPydantic; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfNewInterfaceAnySubnetFipSerializerPydantic; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfNewInterfaceReservedFixedIPFipSerializerPydantic; vt != nil {
		res.any = vt.FloatingIP.asAny()
	}
	return
}

// Can have the runtime types
// [*InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer],
// [*InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer],
// [*InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer],
// [*InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer],
// [*InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer],
// [*InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer]
type instanceNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer:
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer:
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer:
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer:
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer:
//	case *cloud.InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u instanceNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u instanceNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion:
		return vt.GetSource()
	case *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion:
		return vt.GetSource()
	case *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u instanceNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceUnion](
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic{}),
			DiscriminatorValue: "external",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic{}),
			DiscriminatorValue: "subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic{}),
			DiscriminatorValue: "any_subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic{}),
			DiscriminatorValue: "reserved_fixed_ip",
		},
	)
}

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/0'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[0]"
//
// The property Type is required.
type InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic struct {
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
	SecurityGroups []InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydantic
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceExternalSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceExternalSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceExternalSerializerPydanticSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/1'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[1]"
//
// The properties NetworkID, SubnetID, Type are required.
type InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic struct {
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
	FloatingIP InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.security_groups"
	SecurityGroups []InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydantic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion struct {
	OfNewFloatingIP      *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer      `json:",omitzero,inline"`
	OfExistingFloatingIP *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion](u.OfNewFloatingIP, u.OfExistingFloatingIP)
}

func (u *InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNewFloatingIP) {
		return u.OfNewFloatingIP
	} else if !param.IsOmitted(u.OfExistingFloatingIP) {
		return u.OfExistingFloatingIP
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExistingFloatingIP; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion) GetSource() *string {
	if vt := u.OfNewFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExistingFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "existing",
		},
	)
}

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// The property Source is required.
type InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "new".
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer struct {
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
func (f InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceSpecificSubnetFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceSpecificSubnetFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceSpecificSubnetFipSerializerPydanticSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/2'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[2]"
//
// The properties NetworkID, Type are required.
type InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic struct {
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
	FloatingIP InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/ip_family/anyOf/0'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.security_groups"
	SecurityGroups []InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup `json:"security_groups,omitzero"`
	// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/type'
	// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.type"
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydantic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion struct {
	OfNewFloatingIP      *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer      `json:",omitzero,inline"`
	OfExistingFloatingIP *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion](u.OfNewFloatingIP, u.OfExistingFloatingIP)
}

func (u *InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNewFloatingIP) {
		return u.OfNewFloatingIP
	} else if !param.IsOmitted(u.OfExistingFloatingIP) {
		return u.OfExistingFloatingIP
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExistingFloatingIP; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion) GetSource() *string {
	if vt := u.OfNewFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExistingFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "existing",
		},
	)
}

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// The property Source is required.
type InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "new".
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer struct {
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
func (f InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceAnySubnetFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceAnySubnetFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceAnySubnetFipSerializerPydanticSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewVmInterfaceSerializersPydantic/anyOf/3'
// "$.components.schemas.NewVmInterfaceSerializersPydantic.anyOf[3]"
//
// The properties PortID, Type are required.
type InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic struct {
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
	FloatingIP InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/security_groups'
	// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.security_groups"
	SecurityGroups []InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup `json:"security_groups,omitzero"`
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
func (f InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydantic
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion struct {
	OfNewFloatingIP      *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer      `json:",omitzero,inline"`
	OfExistingFloatingIP *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion](u.OfNewFloatingIP, u.OfExistingFloatingIP)
}

func (u *InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNewFloatingIP) {
		return u.OfNewFloatingIP
	} else if !param.IsOmitted(u.OfExistingFloatingIP) {
		return u.OfExistingFloatingIP
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExistingFloatingIP; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion) GetSource() *string {
	if vt := u.OfNewFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExistingFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "existing",
		},
	)
}

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/0'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[0]"
//
// The property Source is required.
type InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "new".
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPNewInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/floating_ip/anyOf/1'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer struct {
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
func (f InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticFloatingIPExistingInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/NewInterfaceReservedFixedIpFipSerializerPydantic/properties/security_groups/items'
// "$.components.schemas.NewInterfaceReservedFixedIpFipSerializerPydantic.properties.security_groups.items"
//
// The property ID is required.
type InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterfaceNewInterfaceReservedFixedIPFipSerializerPydanticSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateInstanceSerializerV2/properties/volumes/items'
// "$.components.schemas.CreateInstanceSerializerV2.properties.volumes.items"
//
// The property Source is required.
type InstanceNewParamsVolume struct {
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/source'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.source"
	//
	// Any of "apptemplate", "existing-volume", "image", "new-volume", "snapshot".
	Source string `json:"source,omitzero,required"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/apptemplate_id'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.apptemplate_id"
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/attachment_tag'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.attachment_tag"
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/boot_index'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.boot_index"
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/delete_on_termination'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.delete_on_termination"
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/image_id'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.image_id"
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/name'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/size'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.size"
	Size param.Opt[int64] `json:"size,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/snapshot_id'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.snapshot_id"
	SnapshotID param.Opt[string] `json:"snapshot_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/volume_id'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.volume_id"
	VolumeID param.Opt[string] `json:"volume_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/tags'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.tags"
	Tags TagUpdateList `json:"tags,omitzero"`
	// '#/components/schemas/CreateInstanceVolumeSerializerPydantic/properties/type_name'
	// "$.components.schemas.CreateInstanceVolumeSerializerPydantic.properties.type_name"
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsVolume) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InstanceNewParamsVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsVolume
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsVolume](
		"Source", false, "apptemplate", "existing-volume", "image", "new-volume", "snapshot",
	)
	apijson.RegisterFieldValidator[InstanceNewParamsVolume](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// '#/components/schemas/CreateInstanceSerializerV2/properties/security_groups/items'
// "$.components.schemas.CreateInstanceSerializerV2.properties.security_groups.items"
//
// The property ID is required.
type InstanceNewParamsSecurityGroup struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceNewParamsSecurityGroup) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InstanceNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameSerializer/properties/name'
	// "$.components.schemas.NameSerializer.properties.name"
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[2]"
	AvailableFloating param.Opt[bool] `query:"available_floating,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[3]"
	ChangesBefore param.Opt[time.Time] `query:"changes-before,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[4]"
	ChangesSince param.Opt[time.Time] `query:"changes-since,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[5]"
	ExcludeFlavorPrefix param.Opt[string] `query:"exclude_flavor_prefix,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[6]"
	ExcludeSecgroup param.Opt[string] `query:"exclude_secgroup,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[7]"
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/8'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[8]"
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/9'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[9]"
	IncludeAI param.Opt[bool] `query:"include_ai,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/10'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[10]"
	IncludeBaremetal param.Opt[bool] `query:"include_baremetal,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/11'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[11]"
	IncludeK8s param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/12'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[12]"
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/13'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[13]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/14'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[14]"
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/15'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[15]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/16'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[16]"
	OnlyIsolated param.Opt[bool] `query:"only_isolated,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/17'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[17]"
	OnlyWithFixedExternalIP param.Opt[bool] `query:"only_with_fixed_external_ip,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/19'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[19]"
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/22'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[22]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/25'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[25]"
	Uuid param.Opt[string] `query:"uuid,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/26'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[26]"
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/27'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[27]"
	WithInterfacesName param.Opt[bool] `query:"with_interfaces_name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/18'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[18]"
	//
	// Any of "created.asc", "created.desc", "name.asc", "name.desc".
	OrderBy InstanceListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/20'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[20]"
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus InstanceListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/21'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[21]"
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "MIGRATING", "PAUSED",
	// "REBOOT", "REBUILD", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED",
	// "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE".
	Status InstanceListParamsStatus `query:"status,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/23'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[23]"
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/24'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[24]"
	//
	// Any of "basic", "advanced".
	TypeDDOSProfile InstanceListParamsTypeDDOSProfile `query:"type_ddos_profile,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/18'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[18]"
type InstanceListParamsOrderBy string

const (
	InstanceListParamsOrderByCreatedAsc  InstanceListParamsOrderBy = "created.asc"
	InstanceListParamsOrderByCreatedDesc InstanceListParamsOrderBy = "created.desc"
	InstanceListParamsOrderByNameAsc     InstanceListParamsOrderBy = "name.asc"
	InstanceListParamsOrderByNameDesc    InstanceListParamsOrderBy = "name.desc"
)

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/20'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[20]"
type InstanceListParamsProtectionStatus string

const (
	InstanceListParamsProtectionStatusActive InstanceListParamsProtectionStatus = "Active"
	InstanceListParamsProtectionStatusQueued InstanceListParamsProtectionStatus = "Queued"
	InstanceListParamsProtectionStatusError  InstanceListParamsProtectionStatus = "Error"
)

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/21'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[21]"
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

// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/24'
// "$.paths['/cloud/v1/instances/{project_id}/{region_id}'].get.parameters[24]"
type InstanceListParamsTypeDDOSProfile string

const (
	InstanceListParamsTypeDDOSProfileBasic    InstanceListParamsTypeDDOSProfile = "basic"
	InstanceListParamsTypeDDOSProfileAdvanced InstanceListParamsTypeDDOSProfile = "advanced"
)

type InstanceDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/3'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[3]"
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/4'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[4]"
	Floatings param.Opt[string] `query:"floatings,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/5'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[5]"
	ReservedFixedIPs param.Opt[string] `query:"reserved_fixed_ips,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/delete/parameters/6'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}']['delete'].parameters[6]"
	Volumes param.Opt[string] `query:"volumes,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceDeleteParams]'s query parameters as `url.Values`.
func (r InstanceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceActionParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faction/post/parameters/0/schema'
	// "$.paths['/cloud/v2/instances/{project_id}/{region_id}/{instance_id}/action'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faction/post/parameters/1/schema'
	// "$.paths['/cloud/v2/instances/{project_id}/{region_id}/{instance_id}/action'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faction/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v2/instances/{project_id}/{region_id}/{instance_id}/action'].post.requestBody.content['application/json'].schema"
	Body InstanceActionParamsBodyUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceActionParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceActionParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// '#/paths/%2Fcloud%2Fv2%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faction/post/requestBody/content/application%2Fjson/schema'
// "$.paths['/cloud/v2/instances/{project_id}/{region_id}/{instance_id}/action'].post.requestBody.content['application/json'].schema"
//
// Satisfied by [InstanceActionParamsBodyStartActionInstanceSerializer] and
// [InstanceActionParamsBodyBasicActionInstanceSerializer]
type InstanceActionParamsBodyUnion interface {
	implInstanceActionParamsBodyUnion()
}

func (InstanceActionParamsBodyStartActionInstanceSerializer) implInstanceActionParamsBodyUnion() {}
func (InstanceActionParamsBodyBasicActionInstanceSerializer) implInstanceActionParamsBodyUnion() {}

// '#/components/schemas/InstanceActionSerializerPydantic/anyOf/0'
// "$.components.schemas.InstanceActionSerializerPydantic.anyOf[0]"
//
// The property Action is required.
type InstanceActionParamsBodyStartActionInstanceSerializer struct {
	// '#/components/schemas/StartActionInstanceSerializer/properties/activate_profile/anyOf/0'
	// "$.components.schemas.StartActionInstanceSerializer.properties.activate_profile.anyOf[0]"
	ActivateProfile param.Opt[bool] `json:"activate_profile,omitzero"`
	// '#/components/schemas/StartActionInstanceSerializer/properties/action'
	// "$.components.schemas.StartActionInstanceSerializer.properties.action"
	//
	// This field can be elided, and will marshal its zero value as "start".
	Action constant.Start `json:"action,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceActionParamsBodyStartActionInstanceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceActionParamsBodyStartActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow InstanceActionParamsBodyStartActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InstanceActionSerializerPydantic/anyOf/1'
// "$.components.schemas.InstanceActionSerializerPydantic.anyOf[1]"
//
// The property Action is required.
type InstanceActionParamsBodyBasicActionInstanceSerializer struct {
	// '#/components/schemas/BasicActionInstanceSerializer/properties/action'
	// "$.components.schemas.BasicActionInstanceSerializer.properties.action"
	//
	// Any of "reboot", "reboot_hard", "resume", "stop", "suspend".
	Action string `json:"action,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceActionParamsBodyBasicActionInstanceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fput_into_servergroup/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/put_into_servergroup'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fput_into_servergroup/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/put_into_servergroup'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/InstancePutServerGroupSchema/properties/servergroup_id'
	// "$.components.schemas.InstancePutServerGroupSchema.properties.servergroup_id"
	ServergroupID string `json:"servergroup_id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceAddToPlacementGroupParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r InstanceAddToPlacementGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAddToPlacementGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceAssignSecurityGroupParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faddsecuritygroup/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/addsecuritygroup'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Faddsecuritygroup/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/addsecuritygroup'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/name'
	// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/ports_security_group_names'
	// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.ports_security_group_names"
	PortsSecurityGroupNames []InstanceAssignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceAssignSecurityGroupParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r InstanceAssignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAssignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/ports_security_group_names/items'
// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.ports_security_group_names.items"
//
// The properties PortID, SecurityGroupNames are required.
type InstanceAssignSecurityGroupParamsPortsSecurityGroupName struct {
	// '#/components/schemas/PortSecurityGroupNamesSchema/properties/port_id'
	// "$.components.schemas.PortSecurityGroupNamesSchema.properties.port_id"
	PortID param.Opt[string] `json:"port_id,omitzero,required"`
	// '#/components/schemas/PortSecurityGroupNamesSchema/properties/security_group_names'
	// "$.components.schemas.PortSecurityGroupNamesSchema.properties.security_group_names"
	SecurityGroupNames []string `json:"security_group_names,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceAssignSecurityGroupParamsPortsSecurityGroupName) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceAssignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstanceAssignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceDisablePortSecurityParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fports%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fdisable_port_security/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ports/{project_id}/{region_id}/{port_id}/disable_port_security'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fports%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fdisable_port_security/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ports/{project_id}/{region_id}/{port_id}/disable_port_security'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceDisablePortSecurityParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type InstanceEnablePortSecurityParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fports%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fenable_port_security/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ports/{project_id}/{region_id}/{port_id}/enable_port_security'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fports%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bport_id%7D%2Fenable_port_security/post/parameters/1/schema'
	// "$.paths['/cloud/v1/ports/{project_id}/{region_id}/{port_id}/enable_port_security'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceEnablePortSecurityParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InstanceGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InstanceGetConsoleParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fget_console/get/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/get_console'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fget_console/get/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/get_console'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fget_console/get/parameters/3'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/get_console'].get.parameters[3]"
	ConsoleType param.Opt[string] `query:"console_type,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceGetConsoleParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceGetConsoleParams]'s query parameters as
// `url.Values`.
func (r InstanceGetConsoleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceRemoveFromPlacementGroupParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fremove_from_servergroup/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/remove_from_servergroup'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fremove_from_servergroup/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/remove_from_servergroup'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceRemoveFromPlacementGroupParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type InstanceResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fchangeflavor/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/changeflavor'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fchangeflavor/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/changeflavor'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/FlavorIdSchema/properties/flavor_id'
	// "$.components.schemas.FlavorIdSchema.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InstanceUnassignSecurityGroupParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdelsecuritygroup/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/delsecuritygroup'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Fdelsecuritygroup/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/delsecuritygroup'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/name'
	// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/ports_security_group_names'
	// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.ports_security_group_names"
	PortsSecurityGroupNames []InstanceUnassignSecurityGroupParamsPortsSecurityGroupName `json:"ports_security_group_names,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceUnassignSecurityGroupParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r InstanceUnassignSecurityGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUnassignSecurityGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InstancePortsSecurityGroupsSchema/properties/ports_security_group_names/items'
// "$.components.schemas.InstancePortsSecurityGroupsSchema.properties.ports_security_group_names.items"
//
// The properties PortID, SecurityGroupNames are required.
type InstanceUnassignSecurityGroupParamsPortsSecurityGroupName struct {
	// '#/components/schemas/PortSecurityGroupNamesSchema/properties/port_id'
	// "$.components.schemas.PortSecurityGroupNamesSchema.properties.port_id"
	PortID param.Opt[string] `json:"port_id,omitzero,required"`
	// '#/components/schemas/PortSecurityGroupNamesSchema/properties/security_group_names'
	// "$.components.schemas.PortSecurityGroupNamesSchema.properties.security_group_names"
	SecurityGroupNames []string `json:"security_group_names,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceUnassignSecurityGroupParamsPortsSecurityGroupName) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceUnassignSecurityGroupParamsPortsSecurityGroupName) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUnassignSecurityGroupParamsPortsSecurityGroupName
	return param.MarshalObject(r, (*shadow)(&r))
}
