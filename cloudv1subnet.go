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

// CloudV1SubnetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SubnetService] method instead.
type CloudV1SubnetService struct {
	Options      []option.RequestOption
	Metadata     *CloudV1SubnetMetadataService
	MetadataItem *CloudV1SubnetMetadataItemService
}

// NewCloudV1SubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1SubnetService(opts ...option.RequestOption) (r *CloudV1SubnetService) {
	r = &CloudV1SubnetService{}
	r.Options = opts
	r.Metadata = NewCloudV1SubnetMetadataService(opts...)
	r.MetadataItem = NewCloudV1SubnetMetadataItemService(opts...)
	return
}

// Create subnet
func (r *CloudV1SubnetService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1SubnetNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get subnet
func (r *CloudV1SubnetService) Get(ctx context.Context, projectID int64, regionID int64, subnetID string, opts ...option.RequestOption) (res *Subnet, err error) {
	opts = append(r.Options[:], opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change subnet properties
func (r *CloudV1SubnetService) Update(ctx context.Context, projectID int64, regionID int64, subnetID string, body CloudV1SubnetUpdateParams, opts ...option.RequestOption) (res *Subnet, err error) {
	opts = append(r.Options[:], opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List subnets
func (r *CloudV1SubnetService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1SubnetListParams, opts ...option.RequestOption) (res *CloudV1SubnetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete subnet
func (r *CloudV1SubnetService) Delete(ctx context.Context, projectID int64, regionID int64, subnetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type IPVersion int64

const (
	IPVersion4 IPVersion = 4
	IPVersion6 IPVersion = 6
)

func (r IPVersion) IsKnown() bool {
	switch r {
	case IPVersion4, IPVersion6:
		return true
	}
	return false
}

type Subnet struct {
	// CIDR
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// Datetime when the subnet was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// True if DHCP should be enabled
	EnableDhcp bool `json:"enable_dhcp,required"`
	// IP version
	IPVersion IPVersion `json:"ip_version,required"`
	// Subnet name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Datetime when the subnet was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Subnet id.
	ID string `json:"id,nullable" format:"uuid4"`
	// Number of available ips in subnet
	AvailableIPs int64 `json:"available_ips,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// List IP addresses of a DNS resolver reachable from the network
	DNSNameservers []string `json:"dns_nameservers,nullable" format:"ipvanyaddress"`
	// Default GW IPv4 address, advertised in DHCP routes of this subnet. If null, no
	// gateway is advertised by this subnet.
	GatewayIP string `json:"gateway_ip,nullable" format:"ipvanyaddress"`
	// Subnet has router attached to it
	HasRouter bool `json:"has_router"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []NeutronRoute `json:"host_routes,nullable"`
	// Subnetwork metadata
	Metadata []DetailedMetadata `json:"metadata"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Total number of ips in subnet
	TotalIPs int64      `json:"total_ips,nullable"`
	JSON     subnetJSON `json:"-"`
}

// subnetJSON contains the JSON metadata for the struct [Subnet]
type subnetJSON struct {
	Cidr           apijson.Field
	CreatedAt      apijson.Field
	EnableDhcp     apijson.Field
	IPVersion      apijson.Field
	Name           apijson.Field
	NetworkID      apijson.Field
	ProjectID      apijson.Field
	Region         apijson.Field
	RegionID       apijson.Field
	UpdatedAt      apijson.Field
	ID             apijson.Field
	AvailableIPs   apijson.Field
	CreatorTaskID  apijson.Field
	DNSNameservers apijson.Field
	GatewayIP      apijson.Field
	HasRouter      apijson.Field
	HostRoutes     apijson.Field
	Metadata       apijson.Field
	TaskID         apijson.Field
	TotalIPs       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Subnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subnetJSON) RawJSON() string {
	return r.raw
}

type CloudV1SubnetListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Subnet                      `json:"results,required"`
	JSON    cloudV1SubnetListResponseJSON `json:"-"`
}

// cloudV1SubnetListResponseJSON contains the JSON metadata for the struct
// [CloudV1SubnetListResponse]
type cloudV1SubnetListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1SubnetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1SubnetListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1SubnetNewParams struct {
	// CIDR
	Cidr param.Field[string] `json:"cidr,required" format:"ipvanynetwork"`
	// Subnet name
	Name param.Field[string] `json:"name,required"`
	// Network ID
	NetworkID param.Field[string] `json:"network_id,required" format:"uuid4"`
	// True if the network's router should get a gateway in this subnet. Must be
	// explicitly 'false' when gateway_ip is null.
	ConnectToNetworkRouter param.Field[bool] `json:"connect_to_network_router"`
	// List IP addresses of DNS servers to advertise via DHCP.
	DNSNameservers param.Field[[]string] `json:"dns_nameservers" format:"ipvanyaddress"`
	// True if DHCP should be enabled
	EnableDhcp param.Field[bool] `json:"enable_dhcp"`
	// Default GW IPv4 address to advertise in DHCP routes in this subnet. Omit this
	// field to let the cloud backend allocate it automatically. Set to null if no
	// gateway must be advertised by this subnet's DHCP (useful when attaching
	// instances to multiple subnets in order to prevent default route conflicts).
	GatewayIP param.Field[string] `json:"gateway_ip" format:"ipvanyaddress"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes param.Field[[]NeutronRouteParam] `json:"host_routes"`
	// IP version
	IPVersion param.Field[IPVersion] `json:"ip_version"`
	// Create one or more metadata items for a subnet
	Metadata param.Field[interface{}] `json:"metadata"`
	// ID of the router to connect to. Requires `connect_to_network_router` set to
	// true. If not specified, attempts to find a router created during network
	// creation.
	RouterIDToConnect param.Field[string] `json:"router_id_to_connect" format:"uuid4"`
}

func (r CloudV1SubnetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1SubnetUpdateParams struct {
	// List IP addresses of DNS servers to advertise via DHCP.
	DNSNameservers param.Field[[]string] `json:"dns_nameservers" format:"ipvanyaddress"`
	// True if DHCP should be enabled
	EnableDhcp param.Field[bool] `json:"enable_dhcp"`
	// Default GW IPv4 address to advertise in DHCP routes in this subnet. Omit this
	// field to let the cloud backend allocate it automatically. Set to null if no
	// gateway must be advertised by this subnet's DHCP (useful when attaching
	// instances to multiple subnets in order to prevent default route conflicts).
	GatewayIP param.Field[string] `json:"gateway_ip" format:"ipvanyaddress"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes param.Field[[]NeutronRouteParam] `json:"host_routes"`
	// Name
	Name param.Field[string] `json:"name"`
}

func (r CloudV1SubnetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1SubnetListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Filter by metadata keys. ?metadata_k=key1&metadata_k=key2
	MetadataK param.Field[[]string] `query:"metadata_k"`
	// Optional. Filter by metadata key-value pairs. curl -G --data-urlencode
	// "metadata_kv={"key": "value"}" --url "https://example.com/cloud/v1/subnets/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Optional. Only list subnets of this network
	NetworkID param.Field[string] `query:"network_id" format:"uuid4"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Optional. Ordering subnets list result by `name`, `created_at`, `updated_at`,
	// `available_ips`, `total_ips`, and `cidr` (default) fields of the subnet and
	// directions (`name.asc`).
	OrderBy param.Field[[]CloudV1SubnetListParamsOrderBy] `query:"order_by"`
}

// URLQuery serializes [CloudV1SubnetListParams]'s query parameters as
// `url.Values`.
func (r CloudV1SubnetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1SubnetListParamsOrderBy string

const (
	CloudV1SubnetListParamsOrderByAvailableIPsAsc  CloudV1SubnetListParamsOrderBy = "available_ips.asc"
	CloudV1SubnetListParamsOrderByAvailableIPsDesc CloudV1SubnetListParamsOrderBy = "available_ips.desc"
	CloudV1SubnetListParamsOrderByCidrAsc          CloudV1SubnetListParamsOrderBy = "cidr.asc"
	CloudV1SubnetListParamsOrderByCidrDesc         CloudV1SubnetListParamsOrderBy = "cidr.desc"
	CloudV1SubnetListParamsOrderByCreatedAtAsc     CloudV1SubnetListParamsOrderBy = "created_at.asc"
	CloudV1SubnetListParamsOrderByCreatedAtDesc    CloudV1SubnetListParamsOrderBy = "created_at.desc"
	CloudV1SubnetListParamsOrderByNameAsc          CloudV1SubnetListParamsOrderBy = "name.asc"
	CloudV1SubnetListParamsOrderByNameDesc         CloudV1SubnetListParamsOrderBy = "name.desc"
	CloudV1SubnetListParamsOrderByTotalIPsAsc      CloudV1SubnetListParamsOrderBy = "total_ips.asc"
	CloudV1SubnetListParamsOrderByTotalIPsDesc     CloudV1SubnetListParamsOrderBy = "total_ips.desc"
	CloudV1SubnetListParamsOrderByUpdatedAtAsc     CloudV1SubnetListParamsOrderBy = "updated_at.asc"
	CloudV1SubnetListParamsOrderByUpdatedAtDesc    CloudV1SubnetListParamsOrderBy = "updated_at.desc"
)

func (r CloudV1SubnetListParamsOrderBy) IsKnown() bool {
	switch r {
	case CloudV1SubnetListParamsOrderByAvailableIPsAsc, CloudV1SubnetListParamsOrderByAvailableIPsDesc, CloudV1SubnetListParamsOrderByCidrAsc, CloudV1SubnetListParamsOrderByCidrDesc, CloudV1SubnetListParamsOrderByCreatedAtAsc, CloudV1SubnetListParamsOrderByCreatedAtDesc, CloudV1SubnetListParamsOrderByNameAsc, CloudV1SubnetListParamsOrderByNameDesc, CloudV1SubnetListParamsOrderByTotalIPsAsc, CloudV1SubnetListParamsOrderByTotalIPsDesc, CloudV1SubnetListParamsOrderByUpdatedAtAsc, CloudV1SubnetListParamsOrderByUpdatedAtDesc:
		return true
	}
	return false
}
