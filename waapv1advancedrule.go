// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1AdvancedRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1AdvancedRuleService] method instead.
type WaapV1AdvancedRuleService struct {
	Options []option.RequestOption
}

// NewWaapV1AdvancedRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1AdvancedRuleService(opts ...option.RequestOption) (r *WaapV1AdvancedRuleService) {
	r = &WaapV1AdvancedRuleService{}
	r.Options = opts
	return
}

// Retrieve an advanced rules descriptor
func (r *WaapV1AdvancedRuleService) GetDescriptor(ctx context.Context, opts ...option.RequestOption) (res *WaapV1AdvancedRuleGetDescriptorResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/advanced-rules/descriptor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A response from a request to retrieve an advanced rules descriptor
type WaapV1AdvancedRuleGetDescriptorResponse struct {
	// The descriptor's version
	Version string                                          `json:"version,required"`
	Objects []WaapV1AdvancedRuleGetDescriptorResponseObject `json:"objects,nullable"`
	JSON    waapV1AdvancedRuleGetDescriptorResponseJSON     `json:"-"`
}

// waapV1AdvancedRuleGetDescriptorResponseJSON contains the JSON metadata for the
// struct [WaapV1AdvancedRuleGetDescriptorResponse]
type waapV1AdvancedRuleGetDescriptorResponseJSON struct {
	Version     apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1AdvancedRuleGetDescriptorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1AdvancedRuleGetDescriptorResponseJSON) RawJSON() string {
	return r.raw
}

// Advanced rules descriptor object
type WaapV1AdvancedRuleGetDescriptorResponseObject struct {
	// The object's name
	Name string `json:"name,required"`
	// The object's type
	Type string `json:"type,required"`
	// The object's attributes list
	Attrs []WaapV1AdvancedRuleGetDescriptorResponseObjectsAttr `json:"attrs,nullable"`
	// The object's description
	Description string                                            `json:"description,nullable"`
	JSON        waapV1AdvancedRuleGetDescriptorResponseObjectJSON `json:"-"`
}

// waapV1AdvancedRuleGetDescriptorResponseObjectJSON contains the JSON metadata for
// the struct [WaapV1AdvancedRuleGetDescriptorResponseObject]
type waapV1AdvancedRuleGetDescriptorResponseObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Attrs       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1AdvancedRuleGetDescriptorResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1AdvancedRuleGetDescriptorResponseObjectJSON) RawJSON() string {
	return r.raw
}

// An attribute of a descriptor's object
type WaapV1AdvancedRuleGetDescriptorResponseObjectsAttr struct {
	// The attribute's name
	Name string `json:"name,required"`
	// The attribute's type
	Type string `json:"type,required"`
	// A list of arguments for the attribute
	Args []WaapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArg `json:"args,nullable"`
	// The attribute's description
	Description string `json:"description,nullable"`
	// The attribute's hint
	Hint string                                                 `json:"hint,nullable"`
	JSON waapV1AdvancedRuleGetDescriptorResponseObjectsAttrJSON `json:"-"`
}

// waapV1AdvancedRuleGetDescriptorResponseObjectsAttrJSON contains the JSON
// metadata for the struct [WaapV1AdvancedRuleGetDescriptorResponseObjectsAttr]
type waapV1AdvancedRuleGetDescriptorResponseObjectsAttrJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Args        apijson.Field
	Description apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1AdvancedRuleGetDescriptorResponseObjectsAttr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1AdvancedRuleGetDescriptorResponseObjectsAttrJSON) RawJSON() string {
	return r.raw
}

// An argument of a descriptor's object
type WaapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArg struct {
	// The argument's name
	Name string `json:"name,required"`
	// The argument's type
	Type string `json:"type,required"`
	// The argument's description
	Description string                                                     `json:"description,nullable"`
	JSON        waapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArgJSON `json:"-"`
}

// waapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArgJSON contains the JSON
// metadata for the struct [WaapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArg]
type waapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArgJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1AdvancedRuleGetDescriptorResponseObjectsAttrsArgJSON) RawJSON() string {
	return r.raw
}
