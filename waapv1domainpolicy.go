// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainPolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainPolicyService] method instead.
type WaapV1DomainPolicyService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1DomainPolicyService(opts ...option.RequestOption) (r *WaapV1DomainPolicyService) {
	r = &WaapV1DomainPolicyService{}
	r.Options = opts
	return
}

// Modify the activation state of a policy associated with a domain
func (r *WaapV1DomainPolicyService) Toggle(ctx context.Context, domainID int64, policyID string, opts ...option.RequestOption) (res *WaapV1DomainPolicyToggleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/policies/%s/toggle", domainID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Represents the mode of a security rule.
type WaapV1DomainPolicyToggleResponse struct {
	// Indicates if the security rule is active
	Mode bool                                 `json:"mode,required"`
	JSON waapV1DomainPolicyToggleResponseJSON `json:"-"`
}

// waapV1DomainPolicyToggleResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainPolicyToggleResponse]
type waapV1DomainPolicyToggleResponseJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainPolicyToggleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainPolicyToggleResponseJSON) RawJSON() string {
	return r.raw
}
