// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2PortService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2PortService] method instead.
type CloudV2PortService struct {
	Options []option.RequestOption
}

// NewCloudV2PortService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2PortService(opts ...option.RequestOption) (r *CloudV2PortService) {
	r = &CloudV2PortService{}
	r.Options = opts
	return
}

// Assign allowed address pairs for instance port
func (r *CloudV2PortService) AllowAddressPairs(ctx context.Context, projectID int64, regionID int64, portID string, body CloudV2PortAllowAddressPairsParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ports/%v/%v/%s/allow_address_pairs", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CloudV2PortAllowAddressPairsParams struct {
	// A set of zero or more allowed port address pair and/or subnet masks
	AllowedAddressPairs param.Field[[]AllowedAddressPairsParam] `json:"allowed_address_pairs"`
}

func (r CloudV2PortAllowAddressPairsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
