// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// FloatingIPService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFloatingIPService] method instead.
type FloatingIPService struct {
	Options []option.RequestOption
}

// NewFloatingIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFloatingIPService(opts ...option.RequestOption) (r FloatingIPService) {
	r = FloatingIPService{}
	r.Options = opts
	return
}

// Create floating IP
func (r *FloatingIPService) New(ctx context.Context, params FloatingIPNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List floating IPs
func (r *FloatingIPService) List(ctx context.Context, params FloatingIPListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FloatingIPDetailed], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List floating IPs
func (r *FloatingIPService) ListAutoPaging(ctx context.Context, params FloatingIPListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FloatingIPDetailed] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete floating IP
func (r *FloatingIPService) Delete(ctx context.Context, floatingIPID string, body FloatingIPDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Assign floating IP to instance or loadbalancer
func (r *FloatingIPService) Assign(ctx context.Context, floatingIPID string, params FloatingIPAssignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/assign", params.ProjectID.Value, params.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get floating IP
func (r *FloatingIPService) Get(ctx context.Context, floatingIPID string, query FloatingIPGetParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unassign floating IP from the instance
func (r *FloatingIPService) Unassign(ctx context.Context, floatingIPID string, body FloatingIPUnassignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/unassign", body.ProjectID.Value, body.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// '#/components/schemas/FloatingIPDetailedSerializer'
// "$.components.schemas.FloatingIPDetailedSerializer"
type FloatingIPDetailed struct {
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.id.anyOf[0]"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/created_at'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/dns_domain/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.dns_domain.anyOf[0]"
	//
	// Deprecated: deprecated
	DNSDomain string `json:"dns_domain,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/dns_name/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.dns_name.anyOf[0]"
	//
	// Deprecated: deprecated
	DNSName string `json:"dns_name,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/fixed_ip_address/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.fixed_ip_address.anyOf[0]"
	FixedIPAddress string `json:"fixed_ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/floating_ip_address/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.floating_ip_address.anyOf[0]"
	FloatingIPAddress string `json:"floating_ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/instance/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.instance.anyOf[0]"
	Instance FloatingIPDetailedInstance `json:"instance,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/loadbalancer/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.loadbalancer.anyOf[0]"
	Loadbalancer LoadBalancer `json:"loadbalancer,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/metadata'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.metadata"
	//
	// Deprecated: deprecated
	Metadata []Tag `json:"metadata,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/port_id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.port_id.anyOf[0]"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/project_id'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/region'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/region_id'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/router_id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.router_id.anyOf[0]"
	RouterID string `json:"router_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/status/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.status.anyOf[0]"
	//
	// Any of "ACTIVE", "DOWN", "ERROR".
	Status FloatingIPStatus `json:"status,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.subnet_id.anyOf[0]"
	//
	// Deprecated: deprecated
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/tags'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPDetailedSerializer/properties/updated_at'
	// "$.components.schemas.FloatingIPDetailedSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                resp.Field
		CreatedAt         resp.Field
		CreatorTaskID     resp.Field
		DNSDomain         resp.Field
		DNSName           resp.Field
		FixedIPAddress    resp.Field
		FloatingIPAddress resp.Field
		Instance          resp.Field
		Loadbalancer      resp.Field
		Metadata          resp.Field
		PortID            resp.Field
		ProjectID         resp.Field
		Region            resp.Field
		RegionID          resp.Field
		RouterID          resp.Field
		Status            resp.Field
		SubnetID          resp.Field
		Tags              resp.Field
		TaskID            resp.Field
		UpdatedAt         resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailed) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/FloatingIPDetailedSerializer/properties/instance/anyOf/0'
// "$.components.schemas.FloatingIPDetailedSerializer.properties.instance.anyOf[0]"
type FloatingIPDetailedInstance struct {
	// '#/components/schemas/InstanceInFloatingSerializer/properties/addresses'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.addresses"
	Addresses map[string][]FloatingIPDetailedInstanceAddressUnion `json:"addresses,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/creator_task_id'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.creator_task_id"
	CreatorTaskID string `json:"creator_task_id,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/flavor'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.flavor"
	Flavor FloatingIPDetailedInstanceFlavor `json:"flavor,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/instance_created'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.instance_created"
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/instance_description/anyOf/0'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.instance_description.anyOf[0]"
	InstanceDescription string `json:"instance_description,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/instance_id'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/instance_name'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/keypair_name/anyOf/0'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.keypair_name.anyOf[0]"
	KeypairName string `json:"keypair_name,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/metadata'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.metadata"
	//
	// Deprecated: deprecated
	Metadata map[string]string `json:"metadata,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/metadata_detailed'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.metadata_detailed"
	//
	// Deprecated: deprecated
	MetadataDetailed []Tag `json:"metadata_detailed,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/project_id'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/region'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/region_id'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/security_groups'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.security_groups"
	SecurityGroups []FloatingIPDetailedInstanceSecurityGroup `json:"security_groups,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/status'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.status"
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status string `json:"status,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/tags'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/task_state/anyOf/0'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.task_state.anyOf[0]"
	TaskState string `json:"task_state,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/vm_state'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.vm_state"
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState string `json:"vm_state,required"`
	// '#/components/schemas/InstanceInFloatingSerializer/properties/volumes'
	// "$.components.schemas.InstanceInFloatingSerializer.properties.volumes"
	Volumes []FloatingIPDetailedInstanceVolume `json:"volumes,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addresses           resp.Field
		CreatorTaskID       resp.Field
		Flavor              resp.Field
		InstanceCreated     resp.Field
		InstanceDescription resp.Field
		InstanceID          resp.Field
		InstanceName        resp.Field
		KeypairName         resp.Field
		Metadata            resp.Field
		MetadataDetailed    resp.Field
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		SecurityGroups      resp.Field
		Status              resp.Field
		Tags                resp.Field
		TaskID              resp.Field
		TaskState           resp.Field
		VmState             resp.Field
		Volumes             resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstance) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FloatingIPDetailedInstanceAddressUnion contains all possible properties and
// values from [FloatingIPDetailedInstanceAddressSimpleAddressSerializer],
// [FloatingIPDetailedInstanceAddressAddressInterfaceSerializer],
// [FloatingIPDetailedInstanceAddressAddressDetailedSerializer].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FloatingIPDetailedInstanceAddressUnion struct {
	Addr          string `json:"addr"`
	Type          string `json:"type"`
	InterfaceName string `json:"interface_name"`
	// This field is from variant
	// [FloatingIPDetailedInstanceAddressAddressDetailedSerializer].
	SubnetID string `json:"subnet_id"`
	// This field is from variant
	// [FloatingIPDetailedInstanceAddressAddressDetailedSerializer].
	SubnetName string `json:"subnet_name"`
	JSON       struct {
		Addr          resp.Field
		Type          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		raw           string
	} `json:"-"`
}

func (u FloatingIPDetailedInstanceAddressUnion) AsSimpleAddressSerializer() (v FloatingIPDetailedInstanceAddressSimpleAddressSerializer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FloatingIPDetailedInstanceAddressUnion) AsAddressInterfaceSerializer() (v FloatingIPDetailedInstanceAddressAddressInterfaceSerializer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FloatingIPDetailedInstanceAddressUnion) AsAddressDetailedSerializer() (v FloatingIPDetailedInstanceAddressAddressDetailedSerializer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FloatingIPDetailedInstanceAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *FloatingIPDetailedInstanceAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/addresses/additionalProperties/items/anyOf/0'
// "$.components.schemas.InstanceInFloatingSerializer.properties.addresses.additionalProperties.items.anyOf[0]"
type FloatingIPDetailedInstanceAddressSimpleAddressSerializer struct {
	// '#/components/schemas/SimpleAddressSerializer/properties/addr'
	// "$.components.schemas.SimpleAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/SimpleAddressSerializer/properties/type'
	// "$.components.schemas.SimpleAddressSerializer.properties.type"
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceAddressSimpleAddressSerializer) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceAddressSimpleAddressSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/addresses/additionalProperties/items/anyOf/1'
// "$.components.schemas.InstanceInFloatingSerializer.properties.addresses.additionalProperties.items.anyOf[1]"
type FloatingIPDetailedInstanceAddressAddressInterfaceSerializer struct {
	// '#/components/schemas/AddressInterfaceSerializer/properties/addr'
	// "$.components.schemas.AddressInterfaceSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/AddressInterfaceSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.AddressInterfaceSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/AddressInterfaceSerializer/properties/type'
	// "$.components.schemas.AddressInterfaceSerializer.properties.type"
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceAddressAddressInterfaceSerializer) RawJSON() string {
	return r.JSON.raw
}
func (r *FloatingIPDetailedInstanceAddressAddressInterfaceSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/addresses/additionalProperties/items/anyOf/2'
// "$.components.schemas.InstanceInFloatingSerializer.properties.addresses.additionalProperties.items.anyOf[2]"
type FloatingIPDetailedInstanceAddressAddressDetailedSerializer struct {
	// '#/components/schemas/AddressDetailedSerializer/properties/addr'
	// "$.components.schemas.AddressDetailedSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/AddressDetailedSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.AddressDetailedSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/AddressDetailedSerializer/properties/subnet_id'
	// "$.components.schemas.AddressDetailedSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// '#/components/schemas/AddressDetailedSerializer/properties/subnet_name'
	// "$.components.schemas.AddressDetailedSerializer.properties.subnet_name"
	SubnetName string `json:"subnet_name,required"`
	// '#/components/schemas/AddressDetailedSerializer/properties/type'
	// "$.components.schemas.AddressDetailedSerializer.properties.type"
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceAddressAddressDetailedSerializer) RawJSON() string {
	return r.JSON.raw
}
func (r *FloatingIPDetailedInstanceAddressAddressDetailedSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/flavor'
// "$.components.schemas.InstanceInFloatingSerializer.properties.flavor"
type FloatingIPDetailedInstanceFlavor struct {
	// '#/components/schemas/BaseFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.BaseFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/BaseFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.BaseFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/BaseFlavorSerializer/properties/ram'
	// "$.components.schemas.BaseFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/BaseFlavorSerializer/properties/vcpus'
	// "$.components.schemas.BaseFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FlavorID    resp.Field
		FlavorName  resp.Field
		Ram         resp.Field
		Vcpus       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/security_groups/items'
// "$.components.schemas.InstanceInFloatingSerializer.properties.security_groups.items"
type FloatingIPDetailedInstanceSecurityGroup struct {
	// '#/components/schemas/NameSerializerPydantic/properties/name'
	// "$.components.schemas.NameSerializerPydantic.properties.name"
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInFloatingSerializer/properties/volumes/items'
// "$.components.schemas.InstanceInFloatingSerializer.properties.volumes.items"
type FloatingIPDetailedInstanceVolume struct {
	// '#/components/schemas/InstanceVolumeSerializer/properties/id'
	// "$.components.schemas.InstanceVolumeSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/InstanceVolumeSerializer/properties/delete_on_termination'
	// "$.components.schemas.InstanceVolumeSerializer.properties.delete_on_termination"
	DeleteOnTermination bool `json:"delete_on_termination,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		DeleteOnTermination resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceVolume) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateFloatingIPSerializer/properties/fixed_ip_address/anyOf/0'
	// "$.components.schemas.CreateFloatingIPSerializer.properties.fixed_ip_address.anyOf[0]"
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipv4"`
	// '#/components/schemas/CreateFloatingIPSerializer/properties/port_id/anyOf/0'
	// "$.components.schemas.CreateFloatingIPSerializer.properties.port_id.anyOf[0]"
	PortID param.Opt[string] `json:"port_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateFloatingIPSerializer/properties/metadata'
	// "$.components.schemas.CreateFloatingIPSerializer.properties.metadata"
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FloatingIPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FloatingIPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type FloatingIPListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[4]"
	MetadataKv param.Opt[string] `query:"metadata_kv,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[5]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}'].get.parameters[3]"
	MetadataK []string `query:"metadata_k,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [FloatingIPListParams]'s query parameters as `url.Values`.
func (r FloatingIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FloatingIPDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type FloatingIPAssignParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D%2Fassign/post/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}/assign'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D%2Fassign/post/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}/assign'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/AssignFloatingIPSerializer/properties/port_id'
	// "$.components.schemas.AssignFloatingIPSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/AssignFloatingIPSerializer/properties/fixed_ip_address/anyOf/0'
	// "$.components.schemas.AssignFloatingIPSerializer.properties.fixed_ip_address.anyOf[0]"
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipvanyaddress"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPAssignParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FloatingIPAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow FloatingIPAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type FloatingIPGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type FloatingIPUnassignParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D%2Funassign/post/parameters/0/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}/unassign'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffloatingips%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfloating_ip_id%7D%2Funassign/post/parameters/1/schema'
	// "$.paths['/cloud/v1/floatingips/{project_id}/{region_id}/{floating_ip_id}/unassign'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FloatingIPUnassignParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
