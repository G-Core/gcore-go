// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainPolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainPolicyService] method instead.
type DomainPolicyService struct {
	Options []option.RequestOption
}

// NewDomainPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainPolicyService(opts ...option.RequestOption) (r DomainPolicyService) {
	r = DomainPolicyService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of policies linked to a particular domain, optionally
// filtered by rule set
func (r *DomainPolicyService) List(ctx context.Context, domainID int64, query DomainPolicyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapDomainPolicy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v2/domains/%v/policies", domainID)
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

// Retrieve a paginated list of policies linked to a particular domain, optionally
// filtered by rule set
func (r *DomainPolicyService) ListAutoPaging(ctx context.Context, domainID int64, query DomainPolicyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapDomainPolicy] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Configure a security policy on a domain.
func (r *DomainPolicyService) Toggle(ctx context.Context, policyID string, params DomainPolicyToggleParams, opts ...option.RequestOption) (res *WaapDomainPolicySettings, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("waap/v2/domains/%v/policies/%s", params.DomainID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Represents a configurable WAAP security rule, also known as a policy.
type WaapDomainPolicy struct {
	// Unique identifier for the security rule
	ID string `json:"id" api:"required"`
	// Specifies the action taken by the WAAP upon rule activation
	//
	// Any of "Allow", "Block", "Captcha", "Gateway", "Handshake", "Monitor",
	// "Composite".
	Action WaapDomainPolicyAction `json:"action" api:"required"`
	// Detailed description of the security rule
	Description string `json:"description" api:"required"`
	// The rule set group name to which the rule belongs
	Group string `json:"group" api:"required"`
	// Indicates if the security rule is active
	Mode bool `json:"mode" api:"required"`
	// Name of the security rule
	Name string `json:"name" api:"required"`
	// Identifier of the rule set to which the rule belongs
	RuleSetID int64 `json:"rule_set_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Description respjson.Field
		Group       respjson.Field
		Mode        respjson.Field
		Name        respjson.Field
		RuleSetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainPolicy) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the action taken by the WAAP upon rule activation
type WaapDomainPolicyAction string

const (
	WaapDomainPolicyActionAllow     WaapDomainPolicyAction = "Allow"
	WaapDomainPolicyActionBlock     WaapDomainPolicyAction = "Block"
	WaapDomainPolicyActionCaptcha   WaapDomainPolicyAction = "Captcha"
	WaapDomainPolicyActionGateway   WaapDomainPolicyAction = "Gateway"
	WaapDomainPolicyActionHandshake WaapDomainPolicyAction = "Handshake"
	WaapDomainPolicyActionMonitor   WaapDomainPolicyAction = "Monitor"
	WaapDomainPolicyActionComposite WaapDomainPolicyAction = "Composite"
)

// Configurable settings of a security rule (a.k.a. policy) on a domain.
type WaapDomainPolicySettings struct {
	// Indicates if the security rule is active
	Mode bool `json:"mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainPolicySettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainPolicySettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapDomainPolicySettings to a
// WaapDomainPolicySettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapDomainPolicySettingsParam.Overrides()
func (r WaapDomainPolicySettings) ToParam() WaapDomainPolicySettingsParam {
	return param.Override[WaapDomainPolicySettingsParam](json.RawMessage(r.RawJSON()))
}

// Configurable settings of a security rule (a.k.a. policy) on a domain.
//
// The property Mode is required.
type WaapDomainPolicySettingsParam struct {
	// Indicates if the security rule is active
	Mode bool `json:"mode" api:"required"`
	paramObj
}

func (r WaapDomainPolicySettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapDomainPolicySettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapDomainPolicySettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainPolicyListParams struct {
	// Filter policies by rule set ID
	RuleSetID param.Opt[int64] `query:"rule_set_id,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainPolicyListParams]'s query parameters as `url.Values`.
func (r DomainPolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainPolicyToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	// Configurable settings of a security rule (a.k.a. policy) on a domain.
	WaapDomainPolicySettings WaapDomainPolicySettingsParam
	paramObj
}

func (r DomainPolicyToggleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WaapDomainPolicySettings)
}
func (r *DomainPolicyToggleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
