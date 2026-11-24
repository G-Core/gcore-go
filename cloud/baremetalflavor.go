// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Bare metal flavors define the CPU, memory, disks, and network cards of a bare
// metal server, in all regions of the project.
//
// BaremetalFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalFlavorService] method instead.
type BaremetalFlavorService struct {
	Options []option.RequestOption
}

// NewBaremetalFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalFlavorService(opts ...option.RequestOption) (r BaremetalFlavorService) {
	r = BaremetalFlavorService{}
	r.Options = opts
	return
}

// Bare metal flavors available to the client in all regions, with buildable
// capacity and port support.
func (r *BaremetalFlavorService) List(ctx context.Context, params BaremetalFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BaremetalFlavor], err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/flavors", params.ProjectID.Value)
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

// Bare metal flavors available to the client in all regions, with buildable
// capacity and port support.
func (r *BaremetalFlavorService) ListAutoPaging(ctx context.Context, params BaremetalFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BaremetalFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

type BaremetalFlavor struct {
	// Flavor architecture type.
	//
	// Any of "aarch64", "unknown", "x86_64".
	Architecture BaremetalFlavorArchitecture `json:"architecture" api:"required"`
	// How many servers of this flavor the client can build in this region now, after
	// reservations. 0 also means the count is unavailable.
	Capacity int64 `json:"capacity" api:"required"`
	// CPU thread count. The server is dedicated, so nothing is oversubscribed.
	CPUThreads int64 `json:"cpu_threads" api:"required"`
	// Hardware of the flavor, as text to show to a person.
	//
	// `disk` gives the disk layout. `license` marks a Windows flavor. The value is
	// null for a flavor with no hardware text.
	HardwareDescription BaremetalFlavorHardwareDescription `json:"hardware_description" api:"required"`
	// Flavor name.
	Name string `json:"name" api:"required"`
	// What ports attached to a server of this flavor support.
	NetworkingCapabilities BaremetalFlavorNetworkingCapabilities `json:"networking_capabilities" api:"required"`
	// Flavor operating system.
	//
	// Any of "linux", "unknown", "windows".
	OsType BaremetalFlavorOsType `json:"os_type" api:"required"`
	// RAM size in MiB.
	RamMib int64 `json:"ram_mib" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// How many of `capacity` exist only because of the client's own reservations.
	ReservedCapacity int64 `json:"reserved_capacity" api:"required"`
	// Hardware class the flavor is served from. The value is null if the flavor has no
	// resource class.
	ResourceClass string `json:"resource_class" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture           respjson.Field
		Capacity               respjson.Field
		CPUThreads             respjson.Field
		HardwareDescription    respjson.Field
		Name                   respjson.Field
		NetworkingCapabilities respjson.Field
		OsType                 respjson.Field
		RamMib                 respjson.Field
		RegionID               respjson.Field
		ReservedCapacity       respjson.Field
		ResourceClass          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalFlavor) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor architecture type.
type BaremetalFlavorArchitecture string

const (
	BaremetalFlavorArchitectureAarch64 BaremetalFlavorArchitecture = "aarch64"
	BaremetalFlavorArchitectureUnknown BaremetalFlavorArchitecture = "unknown"
	BaremetalFlavorArchitectureX86_64  BaremetalFlavorArchitecture = "x86_64"
)

// Hardware of the flavor, as text to show to a person.
//
// `disk` gives the disk layout. `license` marks a Windows flavor. The value is
// null for a flavor with no hardware text.
type BaremetalFlavorHardwareDescription struct {
	// CPU model and core count.
	CPU string `json:"cpu" api:"required"`
	// Disk layout, including count and RAID level.
	Disk string `json:"disk" api:"required"`
	// Bundled license, empty for the Linux variants.
	License string `json:"license" api:"required"`
	// Network interfaces.
	Network string `json:"network" api:"required"`
	// Memory size.
	Ram string `json:"ram" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		License     respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What ports attached to a server of this flavor support.
type BaremetalFlavorNetworkingCapabilities struct {
	// Whether ports on this flavor accept allowed address pairs.
	AllowedAddressPairs bool `json:"allowed_address_pairs" api:"required"`
	// Whether ports on this flavor accept port security.
	PortSecurity bool `json:"port_security" api:"required"`
	// Whether ports on this flavor accept security groups.
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
func (r BaremetalFlavorNetworkingCapabilities) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFlavorNetworkingCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor operating system.
type BaremetalFlavorOsType string

const (
	BaremetalFlavorOsTypeLinux   BaremetalFlavorOsType = "linux"
	BaremetalFlavorOsTypeUnknown BaremetalFlavorOsType = "unknown"
	BaremetalFlavorOsTypeWindows BaremetalFlavorOsType = "windows"
)

type BaremetalFlavorListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Return only the flavors that can boot this image: its region, OS type and CPU
	// architecture, minus the flavors registered as incompatible with its OS release.
	// An unknown image matches no flavor.
	CompatibleWithImageID param.Opt[string] `query:"compatible_with_image_id,omitzero" format:"uuid4" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by flavor architecture type(s).
	//
	// Any of "aarch64", "unknown", "x86_64".
	Architectures []string `query:"architectures,omitzero" json:"-"`
	// Range filter on CPU thread count.
	CPUThreads BaremetalFlavorListParamsCPUThreads `query:"cpu_threads,omitzero" json:"-"`
	// Operator flavor name filter, e.g. name[prefix]=bm2-.
	Name BaremetalFlavorListParamsName `query:"name,omitzero" json:"-"`
	// Sort field and direction.
	//
	// Any of "cpu_threads.asc", "cpu_threads.desc", "name.asc", "name.desc",
	// "ram_mib.asc", "ram_mib.desc", "region_id.asc", "region_id.desc".
	OrderBy BaremetalFlavorListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by flavor operating system(s).
	//
	// Any of "linux", "unknown", "windows".
	OsTypes []string `query:"os_types,omitzero" json:"-"`
	// Range filter on RAM size in MiB.
	RamMib BaremetalFlavorListParamsRamMib `query:"ram_mib,omitzero" json:"-"`
	// Filter by region ID(s). Omit = all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalFlavorListParams]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on CPU thread count.
type BaremetalFlavorListParamsCPUThreads struct {
	// Strictly greater than, e.g. `[gt]=1`.
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Greater than or equal, e.g. `[gte]=2`.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Strictly less than, e.g. `[lt]=8`.
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Less than or equal, e.g. `[lte]=4`.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalFlavorListParamsCPUThreads]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParamsCPUThreads) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operator flavor name filter, e.g. name[prefix]=bm2-.
type BaremetalFlavorListParamsName struct {
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

// URLQuery serializes [BaremetalFlavorListParamsName]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type BaremetalFlavorListParamsOrderBy string

const (
	BaremetalFlavorListParamsOrderByCPUThreadsAsc  BaremetalFlavorListParamsOrderBy = "cpu_threads.asc"
	BaremetalFlavorListParamsOrderByCPUThreadsDesc BaremetalFlavorListParamsOrderBy = "cpu_threads.desc"
	BaremetalFlavorListParamsOrderByNameAsc        BaremetalFlavorListParamsOrderBy = "name.asc"
	BaremetalFlavorListParamsOrderByNameDesc       BaremetalFlavorListParamsOrderBy = "name.desc"
	BaremetalFlavorListParamsOrderByRamMibAsc      BaremetalFlavorListParamsOrderBy = "ram_mib.asc"
	BaremetalFlavorListParamsOrderByRamMibDesc     BaremetalFlavorListParamsOrderBy = "ram_mib.desc"
	BaremetalFlavorListParamsOrderByRegionIDAsc    BaremetalFlavorListParamsOrderBy = "region_id.asc"
	BaremetalFlavorListParamsOrderByRegionIDDesc   BaremetalFlavorListParamsOrderBy = "region_id.desc"
)

// Range filter on RAM size in MiB.
type BaremetalFlavorListParamsRamMib struct {
	// Strictly greater than, e.g. `[gt]=1`.
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Greater than or equal, e.g. `[gte]=2`.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Strictly less than, e.g. `[lt]=8`.
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Less than or equal, e.g. `[lte]=4`.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalFlavorListParamsRamMib]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParamsRamMib) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
