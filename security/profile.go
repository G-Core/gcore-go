// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ProfileService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	return
}

// Create protection profile. Protection is enabled at the same time as profile is
// created
func (r *ProfileService) New(ctx context.Context, body ProfileNewParams, opts ...option.RequestOption) (res *ClientProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "security/iaas/v2/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get list of protection profiles. Client receives only profiles created by him
func (r *ProfileService) List(ctx context.Context, query ProfileListParams, opts ...option.RequestOption) (res *[]ClientProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "security/iaas/v2/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete protection profile. Protection is disabled at the same time as profile is
// deleted
func (r *ProfileService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("security/iaas/v2/profiles/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get profile by id
func (r *ProfileService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ClientProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("security/iaas/v2/profiles/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Recreate profile with another profile template (for other cases use detail API)
func (r *ProfileService) Recreate(ctx context.Context, id int64, body ProfileRecreateParams, opts ...option.RequestOption) (res *ClientProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("security/iaas/v2/profiles/%v/recreate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Update profile. Protection policies are updated at the same time as profile
// updated
func (r *ProfileService) Replace(ctx context.Context, id int64, body ProfileReplaceParams, opts ...option.RequestOption) (res *ClientProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("security/iaas/v2/profiles/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ClientProfile struct {
	ID              int64                 `json:"id,required"`
	Fields          []ClientProfileField  `json:"fields,required"`
	Options         ClientProfileOptions  `json:"options,required"`
	Plan            string                `json:"plan,required"`
	ProfileTemplate ClientProfileTemplate `json:"profile_template,required"`
	Protocols       []map[string]any      `json:"protocols,required"`
	// Region where the protection profiles will be deployed
	Site      string         `json:"site,required"`
	Status    map[string]any `json:"status,required"`
	IPAddress string         `json:"ip_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Fields          respjson.Field
		Options         respjson.Field
		Plan            respjson.Field
		ProfileTemplate respjson.Field
		Protocols       respjson.Field
		Site            respjson.Field
		Status          respjson.Field
		IPAddress       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientProfile) RawJSON() string { return r.JSON.raw }
func (r *ClientProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientProfileField struct {
	ID               int64          `json:"id,required"`
	BaseField        int64          `json:"base_field,required"`
	Default          string         `json:"default,required"`
	Description      string         `json:"description,required"`
	FieldType        string         `json:"field_type,required"`
	Name             string         `json:"name,required"`
	Required         bool           `json:"required,required"`
	ValidationSchema map[string]any `json:"validation_schema,required"`
	FieldValue       any            `json:"field_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BaseField        respjson.Field
		Default          respjson.Field
		Description      respjson.Field
		FieldType        respjson.Field
		Name             respjson.Field
		Required         respjson.Field
		ValidationSchema respjson.Field
		FieldValue       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientProfileField) RawJSON() string { return r.JSON.raw }
func (r *ClientProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientProfileOptions struct {
	Active bool   `json:"active,required"`
	Bgp    bool   `json:"bgp,required"`
	Price  string `json:"price,required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Bgp         respjson.Field
		Price       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientProfileOptions) RawJSON() string { return r.JSON.raw }
func (r *ClientProfileOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewParams struct {
	Fields          []ProfileNewParamsField `json:"fields,omitzero,required"`
	ProfileTemplate int64                   `json:"profile_template,required"`
	// Region where the protection profiles will be deployed
	Site string `json:"site,required"`
	// Required for Universal template only. Optional for all others.
	IPAddress param.Opt[string] `json:"ip_address,omitzero"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, BaseField, Default, Description, FieldType, Name, Required,
// ValidationSchema are required.
type ProfileNewParamsField struct {
	BaseField  int64 `json:"base_field,required"`
	FieldValue any   `json:"field_value,omitzero"`
	paramObj
}

func (r ProfileNewParamsField) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListParams struct {
	ExcludeEmptyAddress param.Opt[bool]   `query:"exclude_empty_address,omitzero" json:"-"`
	IncludeDeleted      param.Opt[bool]   `query:"include_deleted,omitzero" json:"-"`
	IPAddress           param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	Site                param.Opt[string] `query:"site,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileListParams]'s query parameters as `url.Values`.
func (r ProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ProfileRecreateParams struct {
	Fields          []ProfileRecreateParamsField `json:"fields,omitzero,required"`
	ProfileTemplate int64                        `json:"profile_template,required"`
	// Required for Universal template only. Optional for all others.
	IPAddress param.Opt[string] `json:"ip_address,omitzero"`
	// Region where the protection profiles will be deployed
	Site param.Opt[string] `json:"site,omitzero"`
	paramObj
}

func (r ProfileRecreateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileRecreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileRecreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, BaseField, Default, Description, FieldType, Name, Required,
// ValidationSchema are required.
type ProfileRecreateParamsField struct {
	BaseField  int64 `json:"base_field,required"`
	FieldValue any   `json:"field_value,omitzero"`
	paramObj
}

func (r ProfileRecreateParamsField) MarshalJSON() (data []byte, err error) {
	type shadow ProfileRecreateParamsField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileRecreateParamsField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileReplaceParams struct {
	Fields          []ProfileReplaceParamsField `json:"fields,omitzero,required"`
	ProfileTemplate int64                       `json:"profile_template,required"`
	// Required for Universal template only. Optional for all others.
	IPAddress param.Opt[string] `json:"ip_address,omitzero"`
	// Region where the protection profiles will be deployed
	Site param.Opt[string] `json:"site,omitzero"`
	paramObj
}

func (r ProfileReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, BaseField, Default, Description, FieldType, Name, Required,
// ValidationSchema are required.
type ProfileReplaceParamsField struct {
	BaseField  int64 `json:"base_field,required"`
	FieldValue any   `json:"field_value,omitzero"`
	paramObj
}

func (r ProfileReplaceParamsField) MarshalJSON() (data []byte, err error) {
	type shadow ProfileReplaceParamsField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileReplaceParamsField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
