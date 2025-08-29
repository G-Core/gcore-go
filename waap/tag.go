// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// TagService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

// Tags are shortcuts for the rules used in WAAP policies for the creation of more
// complex WAAP rules
func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapTag], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "waap/v1/tags"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Tags are shortcuts for the rules used in WAAP policies for the creation of more
// complex WAAP rules
func (r *TagService) ListAutoPaging(ctx context.Context, query TagListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapTag] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Tags provide shortcuts for the rules used in WAAP policies for the creation of
// more complex WAAP rules.
type WaapTag struct {
	// A tag's human readable description
	Description string `json:"description,required"`
	// The name of a tag that should be used in a WAAP rule condition
	Name string `json:"name,required"`
	// The display name of the tag
	ReadableName string `json:"readable_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		Name         respjson.Field
		ReadableName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTag) RawJSON() string { return r.JSON.raw }
func (r *WaapTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter tags by their name. Supports '\*' as a wildcard character.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter tags by their readable name. Supports '\*' as a wildcard character.
	ReadableName param.Opt[string] `query:"readable_name,omitzero" json:"-"`
	// Filter to include only reserved tags.
	Reserved param.Opt[bool] `query:"reserved,omitzero" json:"-"`
	// Determine the field to order results by
	//
	// Any of "name", "readable_name", "reserved", "-name", "-readable_name",
	// "-reserved".
	Ordering TagListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Determine the field to order results by
type TagListParamsOrdering string

const (
	TagListParamsOrderingName              TagListParamsOrdering = "name"
	TagListParamsOrderingReadableName      TagListParamsOrdering = "readable_name"
	TagListParamsOrderingReserved          TagListParamsOrdering = "reserved"
	TagListParamsOrderingMinusName         TagListParamsOrdering = "-name"
	TagListParamsOrderingMinusReadableName TagListParamsOrdering = "-readable_name"
	TagListParamsOrderingMinusReserved     TagListParamsOrdering = "-reserved"
)
