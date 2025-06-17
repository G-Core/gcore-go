// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/policies/%s/toggle", body.DomainID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type DomainPolicyToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
