// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// AdvancedRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedRuleService] method instead.
type AdvancedRuleService struct {
	Options []option.RequestOption
}

// NewAdvancedRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdvancedRuleService(opts ...option.RequestOption) (r AdvancedRuleService) {
	r = AdvancedRuleService{}
	r.Options = opts
	return
}

// Retrieve an advanced rules descriptor
func (r *AdvancedRuleService) List(ctx context.Context, opts ...option.RequestOption) (res *WaapAdvancedRuleDescriptorList, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/advanced-rules/descriptor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Advanced rules descriptor object
type WaapAdvancedRuleDescriptor struct {
	// The object's name
	Name string `json:"name,required"`
	// The object's type
	Type string `json:"type,required"`
	// The object's attributes list
	Attrs []WaapAdvancedRuleDescriptorAttr `json:"attrs,nullable"`
	// The object's description
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Attrs       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptor) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An attribute of a descriptor's object
type WaapAdvancedRuleDescriptorAttr struct {
	// The attribute's name
	Name string `json:"name,required"`
	// The attribute's type
	Type string `json:"type,required"`
	// A list of arguments for the attribute
	Args []WaapAdvancedRuleDescriptorAttrArg `json:"args,nullable"`
	// The attribute's description
	Description string `json:"description,nullable"`
	// The attribute's hint
	Hint string `json:"hint,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Args        respjson.Field
		Description respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorAttr) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorAttr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An argument of a descriptor's object
type WaapAdvancedRuleDescriptorAttrArg struct {
	// The argument's name
	Name string `json:"name,required"`
	// The argument's type
	Type string `json:"type,required"`
	// The argument's description
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorAttrArg) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorAttrArg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response from a request to retrieve an advanced rules descriptor
type WaapAdvancedRuleDescriptorList struct {
	// The descriptor's version
	Version string                       `json:"version,required"`
	Objects []WaapAdvancedRuleDescriptor `json:"objects,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorList) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
