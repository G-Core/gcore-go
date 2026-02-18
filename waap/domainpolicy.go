// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Modify the activation state of a policy associated with a domain
func (r *DomainPolicyService) Toggle(ctx context.Context, policyID string, body DomainPolicyToggleParams, opts ...option.RequestOption) (res *WaapPolicyMode, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/policies/%s/toggle", body.DomainID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Represents the mode of a security rule.
type WaapPolicyMode struct {
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyMode) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainPolicyToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
