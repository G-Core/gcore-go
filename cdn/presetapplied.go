// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Applied presets represent the association between a preset and a CDN resource or
// rule. Use them to apply a preset to an object, list the objects a preset is
// applied to, unapply it, and inspect which object fields a preset manages.
//
// PresetAppliedService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPresetAppliedService] method instead.
type PresetAppliedService struct {
	Options []option.RequestOption
}

// NewPresetAppliedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPresetAppliedService(opts ...option.RequestOption) (r PresetAppliedService) {
	r = PresetAppliedService{}
	r.Options = opts
	return
}

// Apply the preset to an object (CDN resource or rule, according to the preset
// `object_type`).
//
// The preset settings are applied to the object, and the options included in the
// preset can no longer be edited on the object until the preset is unapplied.
func (r *PresetAppliedService) Apply(ctx context.Context, presetID int64, body PresetAppliedApplyParams, opts ...option.RequestOption) (res *PresetAppliedApplyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/presets/%v/applied", presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the list of objects the preset is currently applied to.
//
// Non-staff users only see objects that belong to their account.
func (r *PresetAppliedService) GetObjects(ctx context.Context, presetID int64, opts ...option.RequestOption) (res *AppliedPreset, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/presets/%v/applied", presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the preset applied to a CDN resource and the list of resource fields managed
// by the preset.
func (r *PresetAppliedService) GetResourcePreset(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *AppliedPresetFields, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/preset", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the preset applied to a rule and the list of rule fields managed by the
// preset.
func (r *PresetAppliedService) GetRulePreset(ctx context.Context, ruleID int64, query PresetAppliedGetRulePresetParams, opts ...option.RequestOption) (res *AppliedPresetFields, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/rules/%v/preset", query.ResourceID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Unapply the preset from the object. The options managed by the preset are
// removed from the object.
func (r *PresetAppliedService) Unapply(ctx context.Context, objectID int64, body PresetAppliedUnapplyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/presets/%v/applied/%v", body.PresetID, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AppliedPreset struct {
	// IDs of the objects the preset is currently applied to. Empty when the preset is
	// not applied to anything.
	ObjectIDs []int64 `json:"object_ids" api:"required"`
	// Type of objects the preset is applied to.
	ObjectType string `json:"object_type" api:"required"`
	// Deprecated. Present only when `object_ids` is empty. Check `object_ids` instead.
	// This field is kept for backward compatibility and will be removed in a future
	// version.
	//
	// Deprecated: deprecated
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectIDs   respjson.Field
		ObjectType  respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppliedPreset) RawJSON() string { return r.JSON.raw }
func (r *AppliedPreset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preset applied to the object and the list of object fields that are managed by
// the preset.
type AppliedPresetFields struct {
	// ID of the applied preset.
	ID int64 `json:"id"`
	// Defines whether the applied preset can be unapplied from the object by the
	// current user.
	Deletable bool `json:"deletable"`
	// Object fields managed by the preset. Option fields are prefixed with `options.`.
	Fields []string `json:"fields"`
	// Name of the applied preset.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deletable   respjson.Field
		Fields      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppliedPresetFields) RawJSON() string { return r.JSON.raw }
func (r *AppliedPresetFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PresetAppliedApplyResponse struct {
	// ID of the object (CDN resource or rule, according to the preset `object_type`)
	// the preset is applied to.
	ObjectID int64 `json:"object_id" api:"required"`
	// ID of the preset that is applied to the object. Matches the `preset_id` path
	// parameter.
	PresetID int64 `json:"preset_id" api:"required"`
	// Deprecated. Use `preset_id` and `object_id` instead. This field is kept for
	// backward compatibility and will be removed in a future version.
	//
	// Deprecated: deprecated
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectID    respjson.Field
		PresetID    respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PresetAppliedApplyResponse) RawJSON() string { return r.JSON.raw }
func (r *PresetAppliedApplyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PresetAppliedApplyParams struct {
	// ID of the object (CDN resource or rule, according to the preset `object_type`)
	// to apply the preset to.
	ObjectID int64 `json:"object_id" api:"required"`
	paramObj
}

func (r PresetAppliedApplyParams) MarshalJSON() (data []byte, err error) {
	type shadow PresetAppliedApplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PresetAppliedApplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PresetAppliedGetRulePresetParams struct {
	ResourceID int64 `path:"resource_id" api:"required" json:"-"`
	paramObj
}

type PresetAppliedUnapplyParams struct {
	PresetID int64 `path:"preset_id" api:"required" json:"-"`
	paramObj
}
