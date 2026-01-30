// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CDNResourceShieldService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCDNResourceShieldService] method instead.
type CDNResourceShieldService struct {
	Options []option.RequestOption
}

// NewCDNResourceShieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCDNResourceShieldService(opts ...option.RequestOption) (r CDNResourceShieldService) {
	r = CDNResourceShieldService{}
	r.Options = opts
	return
}

// Get information about origin shielding.
func (r *CDNResourceShieldService) Get(ctx context.Context, resourceID int64, opts ...option.RequestOption) (res *OriginShielding, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/shielding_v2", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change origin shielding settings or disabled origin shielding.
func (r *CDNResourceShieldService) Replace(ctx context.Context, resourceID int64, body CDNResourceShieldReplaceParams, opts ...option.RequestOption) (res *OriginShieldingReplaced, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/resources/%v/shielding_v2", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OriginShielding struct {
	// Shielding location ID.
	//
	// If origin shielding is disabled, the parameter value is **null**.
	ShieldingPop int64 `json:"shielding_pop,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ShieldingPop respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginShielding) RawJSON() string { return r.JSON.raw }
func (r *OriginShielding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OriginShielding to a OriginShieldingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OriginShieldingParam.Overrides()
func (r OriginShielding) ToParam() OriginShieldingParam {
	return param.Override[OriginShieldingParam](json.RawMessage(r.RawJSON()))
}

type OriginShieldingParam struct {
	// Shielding location ID.
	//
	// If origin shielding is disabled, the parameter value is **null**.
	ShieldingPop param.Opt[int64] `json:"shielding_pop,omitzero"`
	paramObj
}

func (r OriginShieldingParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginShieldingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginShieldingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginShieldingReplaced = any

type CDNResourceShieldReplaceParams struct {
	OriginShielding OriginShieldingParam
	paramObj
}

func (r CDNResourceShieldReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.OriginShielding)
}
func (r *CDNResourceShieldReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.OriginShielding)
}
