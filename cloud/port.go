// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
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

// Ports reserve IP addresses on a network and act as network interfaces of compute
// resources, with their own security groups and allowed address pairs.
//
// PortService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortService] method instead.
type PortService struct {
	Options []option.RequestOption
}

// NewPortService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPortService(opts ...option.RequestOption) (r PortService) {
	r = PortService{}
	r.Options = opts
	return
}

// A port reserves addresses on a network. The port can exist on its own. A compute
// resource claims the port and then uses it as a network interface. Port with one
// public address Send only the region. The port gets one IPv4 address from the
// region's public network.
//
// ```json
// { "region_id": 1 }
// ```
//
// Port in your own network Name only the network. The port gets one IPv4 address
// from any IPv4 subnet of that network.
//
// ```json
// { "region_id": 1, "network_id": "<network_id>" }
// ```
//
// Port in a chosen subnet Name the subnet in `ip_assignments`. An IPv6-only
// network always needs this form. Without `network_id`, the port goes to the
// network of the subnet.
//
// ```json
//
//	{
//	  "region_id": 1,
//	  "network_id": "<network_id>",
//	  "ip_assignments": [{ "subnet_id": "<subnet_id>" }]
//	}
//
// ```
//
// Port with IPv4 and IPv6 Name one IPv4 subnet and one IPv6 subnet of the same
// network.
//
// ```json
//
//	{
//	  "region_id": 1,
//	  "network_id": "<network_id>",
//	  "ip_assignments": [
//	    { "subnet_id": "<ipv4_subnet_id>" },
//	    { "subnet_id": "<ipv6_subnet_id>" }
//	  ]
//	}
//
// ```
//
// Port with a chosen address Add `ip_address` to the assignment. The address must
// be free and inside the subnet.
//
// ```json
//
//	{
//	  "region_id": 1,
//	  "network_id": "<network_id>",
//	  "ip_assignments": [
//	    { "subnet_id": "<subnet_id>", "ip_address": "192.168.123.20" }
//	  ]
//	}
//
// ```
//
// Port that holds a failover address Put the shared address in
// `allowed_address_pairs`. Every port that can hold the address needs this. Create
// the port that owns the address first.
//
// ```json
//
//	{
//	  "region_id": 1,
//	  "network_id": "<network_id>",
//	  "allowed_address_pairs": [{ "ip_address": "192.168.123.100" }]
//	}
//
// ```
//
// # Rules
//
//   - `network_id` defaults to the region's public network. The subnets in
//     `ip_assignments` do not change it.
//   - Every subnet in `ip_assignments` must belong to `network_id`.
//   - Without `ip_assignments` the network must have an IPv4 subnet.
//   - One entry in `ip_assignments` gives one address. Repeat a subnet to get more
//     addresses from it.
//   - Each `allowed_address_pairs` address must already belong to a port in the same
//     network.
//   - A port in the region's public network refuses `port_security_enabled`.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real create.
func (r *PortService) New(ctx context.Context, params PortNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates port security, security groups, allowed address pairs, tags and IP
// assignments.
//
// **Behavior:**
//
//   - Undefined fields remain unchanged; list fields replace the whole set, so send
//     the full desired list (`[]` to clear, `ip_assignments` must keep at least one
//     entry)
//   - Tags are merged per key; a `null` value removes the key, and `tags: null`
//     removes all user tags
//   - If no changes are detected, no task is created and an empty task list is
//     returned
//
// **Supported ports:** all compute types (instances, basic VMs, bare metal
// servers, GPU virtual servers, GPU bare metal servers), as long as the port
// `capabilities` report `true` for the field being changed. Ports that the
// platform manages (DHCP, router, load balancer) are rejected. You can update a
// port that is not attached to a server. The server checks the settings when you
// attach the port.
//
// **Typical uses:**
//
//   - Per-port firewall: `security_group_ids` applies to this port only, so ports on
//     the same server can differ. While port security is enabled, an empty list
//     permits no traffic
//   - VIP failover (VRRP/keepalived): add the VIP address to `allowed_address_pairs`
//     of every port allowed to claim it, all on the VIP's network
//   - Routing and NAT appliances: `port_security_enabled: false` permits forwarding
//     traffic with foreign source addresses; clear security groups and allowed
//     address pairs in the same request
//   - Additional addresses: `ip_assignments` adds or releases IPs (dual-stack, extra
//     service IPs) without detaching the port; the guest OS must be configured for
//     the extra ones. Send an entry with only `subnet_id` to get any free address of
//     that subnet, which is the only option on a shared subnet whose free addresses
//     the caller does not know. Such an entry keeps an address the port already has
//     in that subnet, in the order the port reports them, so a shorter list releases
//     the addresses at the end of that order. A request can replace all current
//     addresses. A port an instance holds takes its addresses from one subnet per IP
//     version, because the guest reaches one address range per version
//
// **Public network:** port security and allowed address pairs cannot be changed on
// ports in the region's public network. IP assignments can. To add an address
// there, send an entry with only `subnet_id`. A new entry with `ip_address` is not
// allowed. An entry can still specify an address that the port already holds. Each
// IPv4 address there counts against the external IP quota, so a request that
// allocates one fails when the limit is full, and a request that releases one
// gives the quota back. IPv6 addresses in that network are not counted.
//
// **Network that other clients share:** a network of another project can hold
// ports of other clients. On such a network, port security stays enabled. Each
// allowed address pair must name one address and the MAC address of a port of the
// same project, and a range is not allowed. A network that no other client uses
// keeps the rules of an own network.
func (r *PortService) Update(ctx context.Context, portID string, params PortUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports/%s", params.ProjectID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List a project's ports across all regions.
func (r *PortService) List(ctx context.Context, params PortListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Port], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports", params.ProjectID.Value)
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

// List a project's ports across all regions.
func (r *PortService) ListAutoPaging(ctx context.Context, params PortListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Port] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a port that no compute resource uses. The request returns a task. The
// task deletes the port. Effects
//
//   - A floating IP of the port stays in the project. The task removes it from the
//     port.
//   - Each IPv4 address on the region's public network goes back to the external IP
//     quota. Rules
//   - The port must not be attached. Detach it from the compute resource first.
//   - The platform manages ports of routers, load balancers and DHCP. You cannot
//     delete these ports.
//   - A port that an attach or detach operation holds cannot be deleted. Try again
//     when the operation ends.
//   - An address of the port must not be in `allowed_address_pairs` of other ports.
//     Remove these entries first.
func (r *PortService) Delete(ctx context.Context, portID string, body PortDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports/%s", body.ProjectID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Converts a port to a floating IP with the same address. The task deletes the
// port and creates the floating IP. The port must be unattached, be on a public
// network and have exactly one IPv4 address.
//
// The external IP quota decreases by one and the floating IP quota increases by
// one.
//
// If the floating IP create fails, the task creates the port again with the same
// address. The new port gets a new ID. It keeps only the network, name,
// description, security groups, port security, allowed address pairs and tags.
// Other attributes get their default values.
//
// A concurrent network operation can cause the conversion to fail. In rare cases
// the address is then lost, and you must create a new one.
func (r *PortService) Convert(ctx context.Context, portID string, params PortConvertParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports/%s/convert", params.ProjectID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get port
func (r *PortService) Get(ctx context.Context, portID string, query PortGetParams, opts ...option.RequestOption) (res *Port, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/ports/%s", query.ProjectID.Value, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Port struct {
	// Port ID
	ID string `json:"id" api:"required" format:"uuid4"`
	// List of allowed address pairs. Sorted by `ip_address` for stable client
	// refreshes.
	AllowedAddressPairs []PortAllowedAddressPair `json:"allowed_address_pairs" api:"required"`
	// The compute resource this port is attached to, including trunk role for trunk
	// parents and sub-ports. null when the port is standalone (unattached).
	AttachedTo PortAttachedTo `json:"attached_to" api:"required"`
	// What the driver enforces. null when the port is unattached or managed, or when
	// the project sees the port only because it owns the network of the port. null
	// also when this API version does not know the manager type. Such a port is
	// read-only: a change or a delete returns permission-denied.
	Capabilities PortCapabilities `json:"capabilities" api:"required"`
	// Datetime when the port was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// List of IP addresses assigned to the port.
	IPAssignments []PortIPAssignment `json:"ip_assignments" api:"required"`
	// MAC address of the port
	MacAddress string `json:"mac_address" api:"required"`
	// Service managing this port. null if customer-operated. null also when this API
	// version does not know the manager type. Such a port is read-only: a change or a
	// delete returns permission-denied.
	ManagedBy PortManagedBy `json:"managed_by" api:"required"`
	// ID of the network the port belongs to
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// When true, the driver enforces port security on this port (anti-spoofing checks
	// and security-group filtering apply). When false, all traffic is allowed and
	// security groups are not applied.
	PortSecurityEnabled bool `json:"port_security_enabled" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// List of security group IDs applied to the port. Sorted lexicographically for
	// stable client refreshes.
	SecurityGroupIDs []string `json:"security_group_ids" api:"required" format:"uuid4"`
	// Read-only key-value tags on the port, set by the platform and not modifiable by
	// the user
	SystemTags map[string]string `json:"system_tags" api:"required"`
	// Key-value tags on the port
	Tags map[string]string `json:"tags" api:"required"`
	// Datetime when the port was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AllowedAddressPairs respjson.Field
		AttachedTo          respjson.Field
		Capabilities        respjson.Field
		CreatedAt           respjson.Field
		IPAssignments       respjson.Field
		MacAddress          respjson.Field
		ManagedBy           respjson.Field
		NetworkID           respjson.Field
		PortSecurityEnabled respjson.Field
		RegionID            respjson.Field
		SecurityGroupIDs    respjson.Field
		SystemTags          respjson.Field
		Tags                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Port) RawJSON() string { return r.JSON.raw }
func (r *Port) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortAllowedAddressPair struct {
	// Subnet mask or IP address of the allowed address pair
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortAllowedAddressPair) RawJSON() string { return r.JSON.raw }
func (r *PortAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The compute resource this port is attached to, including trunk role for trunk
// parents and sub-ports. null when the port is standalone (unattached).
type PortAttachedTo struct {
	// ID of the compute resource the port is attached to. null for DHCP attachments.
	ResourceID string `json:"resource_id" api:"required" format:"uuid4"`
	// Type of resource the port is attached to. null when the compute resource's type
	// is not yet known.
	//
	// Any of "baremetal_server", "basic_vm", "dhcp", "gpu_baremetal_server",
	// "gpu_virtual_server", "instance", "load_balancer", "router".
	ResourceType string `json:"resource_type" api:"required"`
	// Trunk info: present when port is a trunk parent or `sub_port`. null otherwise.
	Trunk PortAttachedToTrunkUnion `json:"trunk" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceType respjson.Field
		Trunk        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortAttachedTo) RawJSON() string { return r.JSON.raw }
func (r *PortAttachedTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortAttachedToTrunkUnion contains all possible properties and values from
// [PortAttachedToTrunkParent], [PortAttachedToTrunkSubPort].
//
// Use the [PortAttachedToTrunkUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortAttachedToTrunkUnion struct {
	// Any of "parent", "sub_port".
	Role string `json:"role"`
	// This field is from variant [PortAttachedToTrunkSubPort].
	ParentPortID string `json:"parent_port_id"`
	JSON         struct {
		Role         respjson.Field
		ParentPortID respjson.Field
		raw          string
	} `json:"-"`
}

// anyPortAttachedToTrunk is implemented by each variant of
// [PortAttachedToTrunkUnion] to add type safety for the return type of
// [PortAttachedToTrunkUnion.AsAny]
type anyPortAttachedToTrunk interface {
	implPortAttachedToTrunkUnion()
}

func (PortAttachedToTrunkParent) implPortAttachedToTrunkUnion()  {}
func (PortAttachedToTrunkSubPort) implPortAttachedToTrunkUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PortAttachedToTrunkUnion.AsAny().(type) {
//	case cloud.PortAttachedToTrunkParent:
//	case cloud.PortAttachedToTrunkSubPort:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PortAttachedToTrunkUnion) AsAny() anyPortAttachedToTrunk {
	switch u.Role {
	case "parent":
		return u.AsParent()
	case "sub_port":
		return u.AsSubPort()
	}
	return nil
}

func (u PortAttachedToTrunkUnion) AsParent() (v PortAttachedToTrunkParent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortAttachedToTrunkUnion) AsSubPort() (v PortAttachedToTrunkSubPort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortAttachedToTrunkUnion) RawJSON() string { return u.JSON.raw }

func (r *PortAttachedToTrunkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortAttachedToTrunkParent struct {
	// Trunk role. 'parent' means this port is a trunk parent.
	Role constant.Parent `json:"role" default:"parent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortAttachedToTrunkParent) RawJSON() string { return r.JSON.raw }
func (r *PortAttachedToTrunkParent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortAttachedToTrunkSubPort struct {
	// ID of the trunk parent port this sub-port belongs to.
	ParentPortID string `json:"parent_port_id" api:"required" format:"uuid4"`
	// Trunk role. 'sub_port' means this port is a sub-port of a trunk.
	Role constant.SubPort `json:"role" default:"sub_port"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParentPortID respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortAttachedToTrunkSubPort) RawJSON() string { return r.JSON.raw }
func (r *PortAttachedToTrunkSubPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What the driver enforces. null when the port is unattached or managed, or when
// the project sees the port only because it owns the network of the port. null
// also when this API version does not know the manager type. Such a port is
// read-only: a change or a delete returns permission-denied.
type PortCapabilities struct {
	// Whether the driver enforces allowed address pairs on this port
	AllowedAddressPairs bool `json:"allowed_address_pairs" api:"required"`
	// Whether the driver enforces port security on this port
	PortSecurity bool `json:"port_security" api:"required"`
	// Whether the driver enforces security groups on this port
	SecurityGroups bool `json:"security_groups" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		PortSecurity        respjson.Field
		SecurityGroups      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortCapabilities) RawJSON() string { return r.JSON.raw }
func (r *PortCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortIPAssignment struct {
	// IP address assigned to the port
	IPAddress string `json:"ip_address" api:"required" format:"ipv4"`
	// ID of the subnet
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *PortIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service managing this port. null if customer-operated. null also when this API
// version does not know the manager type. Such a port is read-only: a change or a
// delete returns permission-denied.
type PortManagedBy struct {
	// ID of the managing resource. Null for system.
	ResourceID string `json:"resource_id" api:"required"`
	// Name of the managing resource. Null for system.
	ResourceName string `json:"resource_name" api:"required"`
	// Type of the managing service.
	//
	// Any of "file_share", "gpu_baremetal_cluster", "gpu_virtual_cluster",
	// "k8s_cluster", "load_balancer", "network", "router", "system".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceName respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortManagedBy) RawJSON() string { return r.JSON.raw }
func (r *PortManagedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID where the port is created
	RegionID int64 `json:"region_id" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Network to create the port in. When omitted, it is the network of the subnets in
	// `ip_assignments`, or the region's public network if `ip_assignments` is omitted.
	// Every subnet in `ip_assignments` must belong to it.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid4"`
	// When true, the network driver enforces port security on this port (anti-spoofing
	// checks and security-group filtering apply). Defaults to the network's setting.
	PortSecurityEnabled param.Opt[bool] `json:"port_security_enabled,omitzero"`
	// A set of zero or more allowed port address pair and/or subnet masks.
	AllowedAddressPairs []PortNewParamsAllowedAddressPair `json:"allowed_address_pairs,omitzero"`
	// IP addresses to assign. All subnets must belong to the same network. Repeat a
	// subnet to get several addresses from it. Omit the field to get one address from
	// an IPv4 subnet.
	IPAssignments []IPAssignmentRequestParam `json:"ip_assignments,omitzero"`
	// List of security group IDs to attach to the port.
	SecurityGroupIDs []string `json:"security_group_ids,omitzero" format:"uuid4"`
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

func (r PortNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PortNewParams]'s query parameters as `url.Values`.
func (r PortNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property IPAddress is required.
type PortNewParamsAllowedAddressPair struct {
	// IP address or subnet of the allowed address pair
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	paramObj
}

func (r PortNewParamsAllowedAddressPair) MarshalJSON() (data []byte, err error) {
	type shadow PortNewParamsAllowedAddressPair
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortNewParamsAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// When true, the network driver enforces port security on this port (anti-spoofing
	// checks and security-group filtering apply). When false, all traffic is allowed
	// and security groups are not applied.
	PortSecurityEnabled param.Opt[bool] `json:"port_security_enabled,omitzero"`
	// A set of zero or more allowed port address pair and/or subnet masks. Replaces
	// the full set of allowed address pairs. Send `[]` to clear; omit the field to
	// leave existing pairs unchanged.
	AllowedAddressPairs []PortUpdateParamsAllowedAddressPair `json:"allowed_address_pairs,omitzero"`
	// IP assignments on the port. Replaces the full set. Subnets must belong to the
	// port's network, and a port an instance holds supports one subnet per IP version.
	// An entry without `ip_address` keeps an address the port already has in that
	// subnet, in the order the port reports them, or takes a free one when the port
	// has none. Addresses that no entry keeps are released.
	IPAssignments []IPAssignmentRequestParam `json:"ip_assignments,omitzero"`
	// List of security group IDs. Replaces the full set of security groups attached to
	// the port. Send `[]` to detach all groups; omit the field to leave existing
	// groups unchanged.
	SecurityGroupIDs []string `json:"security_group_ids,omitzero" format:"uuid4"`
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

func (r PortUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property IPAddress is required.
type PortUpdateParamsAllowedAddressPair struct {
	// IP address or subnet of the allowed address pair
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	paramObj
}

func (r PortUpdateParamsAllowedAddressPair) MarshalJSON() (data []byte, err error) {
	type shadow PortUpdateParamsAllowedAddressPair
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortUpdateParamsAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Filter attached (true) or standalone (false) ports.
	Attached param.Opt[bool] `query:"attached,omitzero" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by managed status. true = managed by a service, false =
	// customer-operated. A port with a manager type that this API version does not
	// know counts as managed, but its `managed_by` is null. Cannot be combined with
	// `managed_by` when false.
	Managed param.Opt[bool] `query:"managed,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter ports on networks that give publicly routed address space (true) or not
	// (false).
	Public param.Opt[bool] `query:"public,omitzero" json:"-"`
	// Filter by attached resource ID(s). Matches ports attached to any of the given
	// resources.
	AttachedToResourceIDs []string `query:"attached_to_resource_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by attachment resource type(s). A port matches if its type is any of the
	// given values.
	//
	// Any of "baremetal_server", "basic_vm", "dhcp", "gpu_baremetal_server",
	// "gpu_virtual_server", "instance", "load_balancer", "router".
	AttachedToResourceTypes []string `query:"attached_to_resource_types,omitzero" json:"-"`
	// Range filter on creation time.
	CreatedAt PortListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Filter by port UUID(s).
	IDs []string `query:"ids,omitzero" format:"uuid4" json:"-"`
	// Filter by fixed IP address.
	IPAddress PortListParamsIPAddress `query:"ip_address,omitzero" json:"-"`
	// Filter by manager type, e.g. system or `k8s_cluster`.
	//
	// Any of "file_share", "gpu_baremetal_cluster", "gpu_virtual_cluster",
	// "k8s_cluster", "load_balancer", "network", "router", "system".
	ManagedBy PortListParamsManagedBy `query:"managed_by,omitzero" json:"-"`
	// Filter by managing resource ID(s).
	ManagedByResourceIDs []string `query:"managed_by_resource_ids,omitzero" json:"-"`
	// Filter by managing resource name(s).
	ManagedByResourceNames []string `query:"managed_by_resource_names,omitzero" json:"-"`
	// Filter by network ID(s).
	NetworkIDs []string `query:"network_ids,omitzero" format:"uuid4" json:"-"`
	// Sort field and direction. `network_grouping.asc` keeps the ports of one network
	// together. In each region, the networks shared with the project come first, in
	// the order of the share. If no shared network is public, the default public
	// network of the region comes next. Then come other public networks, then all
	// other networks.
	//
	// Any of "created_at.asc", "created_at.desc", "network_grouping.asc",
	// "region_id.asc", "region_id.desc", "updated_at.asc", "updated_at.desc".
	OrderBy PortListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by region ID(s). Omit = all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	// Filter by security group ID(s). A port matches if it has any of the given
	// security groups.
	SecurityGroupIDs []string `query:"security_group_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by subnet ID(s). A port matches if it has a fixed IP in any of the given
	// subnets.
	SubnetIDs []string `query:"subnet_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
	TagKey PortListParamsTagKey `query:"tag_key,omitzero" json:"-"`
	// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
	TagValue PortListParamsTagValue `query:"tag_value,omitzero" json:"-"`
	// Filter by exact tag key-value pairs, e.g. `tags[env]=prod&tags[team]=core`.
	// Repeat a key to match any of several values for it, e.g.
	// `tags[env]=prod&tags[env]=dev`. Values for one key are ORed, different keys are
	// ANDed; values match case-insensitively.
	Tags map[string][]string `query:"tags,omitzero" json:"-"`
	// Range filter on last-update time.
	UpdatedAt PortListParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParams]'s query parameters as `url.Values`.
func (r PortListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on creation time.
type PortListParamsCreatedAt struct {
	// Strictly after this timestamp, e.g. `[gt]=2026-01-01T00:00:00Z`.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// At or after this timestamp (inclusive), e.g. `[gte]=2026-01-01T00:00:00Z`.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Strictly before this timestamp, e.g. `[lt]=2026-02-01T00:00:00Z`.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// At or before this timestamp (inclusive), e.g. `[lte]=2026-02-01T00:00:00Z`.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r PortListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by fixed IP address.
type PortListParamsIPAddress struct {
	// Substring match on IP address, e.g. `[contains]=192.168`. Repeat the key to
	// match any substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Exact IP match, e.g. `[exact]=192.168.1.10`. Repeat the key to match any of
	// several.
	Exact []string `query:"exact,omitzero" format:"ipvanyaddress" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParamsIPAddress]'s query parameters as
// `url.Values`.
func (r PortListParamsIPAddress) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by manager type, e.g. system or `k8s_cluster`.
type PortListParamsManagedBy string

const (
	PortListParamsManagedByFileShare           PortListParamsManagedBy = "file_share"
	PortListParamsManagedByGPUBaremetalCluster PortListParamsManagedBy = "gpu_baremetal_cluster"
	PortListParamsManagedByGPUVirtualCluster   PortListParamsManagedBy = "gpu_virtual_cluster"
	PortListParamsManagedByK8SCluster          PortListParamsManagedBy = "k8s_cluster"
	PortListParamsManagedByLoadBalancer        PortListParamsManagedBy = "load_balancer"
	PortListParamsManagedByNetwork             PortListParamsManagedBy = "network"
	PortListParamsManagedByRouter              PortListParamsManagedBy = "router"
	PortListParamsManagedBySystem              PortListParamsManagedBy = "system"
)

// Sort field and direction. `network_grouping.asc` keeps the ports of one network
// together. In each region, the networks shared with the project come first, in
// the order of the share. If no shared network is public, the default public
// network of the region comes next. Then come other public networks, then all
// other networks.
type PortListParamsOrderBy string

const (
	PortListParamsOrderByCreatedAtAsc       PortListParamsOrderBy = "created_at.asc"
	PortListParamsOrderByCreatedAtDesc      PortListParamsOrderBy = "created_at.desc"
	PortListParamsOrderByNetworkGroupingAsc PortListParamsOrderBy = "network_grouping.asc"
	PortListParamsOrderByRegionIDAsc        PortListParamsOrderBy = "region_id.asc"
	PortListParamsOrderByRegionIDDesc       PortListParamsOrderBy = "region_id.desc"
	PortListParamsOrderByUpdatedAtAsc       PortListParamsOrderBy = "updated_at.asc"
	PortListParamsOrderByUpdatedAtDesc      PortListParamsOrderBy = "updated_at.desc"
)

// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
type PortListParamsTagKey struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParamsTagKey]'s query parameters as `url.Values`.
func (r PortListParamsTagKey) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
type PortListParamsTagValue struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParamsTagValue]'s query parameters as `url.Values`.
func (r PortListParamsTagValue) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on last-update time.
type PortListParamsUpdatedAt struct {
	// Strictly after this timestamp, e.g. `[gt]=2026-01-01T00:00:00Z`.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// At or after this timestamp (inclusive), e.g. `[gte]=2026-01-01T00:00:00Z`.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Strictly before this timestamp, e.g. `[lt]=2026-02-01T00:00:00Z`.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// At or before this timestamp (inclusive), e.g. `[lte]=2026-02-01T00:00:00Z`.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortListParamsUpdatedAt]'s query parameters as
// `url.Values`.
func (r PortListParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type PortConvertParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Resource type to convert the port to
	//
	// Any of "floating_ip".
	TargetType PortConvertParamsTargetType `json:"target_type,omitzero" api:"required"`
	paramObj
}

func (r PortConvertParams) MarshalJSON() (data []byte, err error) {
	type shadow PortConvertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortConvertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type to convert the port to
type PortConvertParamsTargetType string

const (
	PortConvertParamsTargetTypeFloatingIP PortConvertParamsTargetType = "floating_ip"
)

type PortGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}
