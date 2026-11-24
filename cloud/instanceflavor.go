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

// Instance flavors define the CPU, memory, and OS type of an instance, in all
// regions of the project.
//
// InstanceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceFlavorService] method instead.
type InstanceFlavorService struct {
	Options []option.RequestOption
}

// NewInstanceFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceFlavorService(opts ...option.RequestOption) (r InstanceFlavorService) {
	r = InstanceFlavorService{}
	r.Options = opts
	return
}

// List the virtual instance flavors available to the client, across all regions.
func (r *InstanceFlavorService) List(ctx context.Context, params InstanceFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InstanceFlavor], err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/flavors", params.ProjectID.Value)
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

// List the virtual instance flavors available to the client, across all regions.
func (r *InstanceFlavorService) ListAutoPaging(ctx context.Context, params InstanceFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InstanceFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

type InstanceFlavor struct {
	// Flavor architecture type.
	//
	// Any of "aarch64", "unknown", "x86_64".
	Architecture InstanceFlavorArchitecture `json:"architecture" api:"required"`
	// Whether a new server of this flavor can be created in this region now. The value
	// is false when the flavor is disabled in the region or is out of stock there.
	Available bool `json:"available" api:"required"`
	// Flavor name.
	Name string `json:"name" api:"required"`
	// Flavor operating system.
	//
	// Any of "linux", "unknown", "windows".
	OsType InstanceFlavorOsType `json:"os_type" api:"required"`
	// RAM size in MiB.
	RamMib int64 `json:"ram_mib" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// Virtual CPU count.
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture respjson.Field
		Available    respjson.Field
		Name         respjson.Field
		OsType       respjson.Field
		RamMib       respjson.Field
		RegionID     respjson.Field
		Vcpus        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor architecture type.
type InstanceFlavorArchitecture string

const (
	InstanceFlavorArchitectureAarch64 InstanceFlavorArchitecture = "aarch64"
	InstanceFlavorArchitectureUnknown InstanceFlavorArchitecture = "unknown"
	InstanceFlavorArchitectureX86_64  InstanceFlavorArchitecture = "x86_64"
)

// Flavor operating system.
type InstanceFlavorOsType string

const (
	InstanceFlavorOsTypeLinux   InstanceFlavorOsType = "linux"
	InstanceFlavorOsTypeUnknown InstanceFlavorOsType = "unknown"
	InstanceFlavorOsTypeWindows InstanceFlavorOsType = "windows"
)

type InstanceFlavorListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Set to true to return only the available flavors. Set to false to return only
	// the flavors that are disabled in their region or out of stock there. Omit it to
	// return all flavors.
	Available param.Opt[bool] `query:"available,omitzero" json:"-"`
	// Return only the flavors that can boot from this volume, from its image metadata.
	// An unknown volume matches no flavor. A non-bootable volume or one without image
	// metadata is a 400. At most one `compatible_with_`\* parameter per request.
	CompatibleWithBootVolumeID param.Opt[string] `query:"compatible_with_boot_volume_id,omitzero" format:"uuid4" json:"-"`
	// Return only the flavors that can boot this image: its region, OS type and CPU
	// architecture. An unknown image matches no flavor. At most one
	// `compatible_with_`\* parameter per request.
	CompatibleWithImageID param.Opt[string] `query:"compatible_with_image_id,omitzero" format:"uuid4" json:"-"`
	// Return only the flavors that can boot from this snapshot, resolved through its
	// parent volume's image metadata. An unknown snapshot matches no flavor. A parent
	// volume with no image metadata is a 400. At most one `compatible_with_`\*
	// parameter per request.
	CompatibleWithSnapshotID param.Opt[string] `query:"compatible_with_snapshot_id,omitzero" format:"uuid4" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by flavor architecture type(s).
	//
	// Any of "aarch64", "unknown", "x86_64".
	Architectures []string `query:"architectures,omitzero" json:"-"`
	// Operator flavor name filter, e.g. name[prefix]=g2-.
	Name InstanceFlavorListParamsName `query:"name,omitzero" json:"-"`
	// Sort field and direction.
	//
	// Any of "name.asc", "name.desc", "ram_mib.asc", "ram_mib.desc", "region_id.asc",
	// "region_id.desc", "vcpus.asc", "vcpus.desc".
	OrderBy InstanceFlavorListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by flavor operating system(s).
	//
	// Any of "linux", "unknown", "windows".
	OsTypes []string `query:"os_types,omitzero" json:"-"`
	// Range filter on RAM size in MiB.
	RamMib InstanceFlavorListParamsRamMib `query:"ram_mib,omitzero" json:"-"`
	// Filter by region ID(s). Omit = all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	// Range filter on virtual CPU count.
	Vcpus InstanceFlavorListParamsVcpus `query:"vcpus,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceFlavorListParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operator flavor name filter, e.g. name[prefix]=g2-.
type InstanceFlavorListParamsName struct {
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

// URLQuery serializes [InstanceFlavorListParamsName]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type InstanceFlavorListParamsOrderBy string

const (
	InstanceFlavorListParamsOrderByNameAsc      InstanceFlavorListParamsOrderBy = "name.asc"
	InstanceFlavorListParamsOrderByNameDesc     InstanceFlavorListParamsOrderBy = "name.desc"
	InstanceFlavorListParamsOrderByRamMibAsc    InstanceFlavorListParamsOrderBy = "ram_mib.asc"
	InstanceFlavorListParamsOrderByRamMibDesc   InstanceFlavorListParamsOrderBy = "ram_mib.desc"
	InstanceFlavorListParamsOrderByRegionIDAsc  InstanceFlavorListParamsOrderBy = "region_id.asc"
	InstanceFlavorListParamsOrderByRegionIDDesc InstanceFlavorListParamsOrderBy = "region_id.desc"
	InstanceFlavorListParamsOrderByVcpusAsc     InstanceFlavorListParamsOrderBy = "vcpus.asc"
	InstanceFlavorListParamsOrderByVcpusDesc    InstanceFlavorListParamsOrderBy = "vcpus.desc"
)

// Range filter on RAM size in MiB.
type InstanceFlavorListParamsRamMib struct {
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

// URLQuery serializes [InstanceFlavorListParamsRamMib]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParamsRamMib) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on virtual CPU count.
type InstanceFlavorListParamsVcpus struct {
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

// URLQuery serializes [InstanceFlavorListParamsVcpus]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParamsVcpus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
