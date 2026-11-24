// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// InstanceConsoleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceConsoleService] method instead.
type InstanceConsoleService struct {
	Options []option.RequestOption
}

// NewInstanceConsoleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceConsoleService(opts ...option.RequestOption) (r InstanceConsoleService) {
	r = InstanceConsoleService{}
	r.Options = opts
	return
}

// Returns a URL that gives access to the remote console of an instance. The API
// finds the region of the instance from its ID.
func (r *InstanceConsoleService) Get(ctx context.Context, instanceID string, params InstanceConsoleGetParams, opts ...option.RequestOption) (res *Console, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s/console", params.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type InstanceConsoleGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Console type that you expect. If you omit it, you get the first available type
	// in this order: `novnc`, `spice`.
	//
	// Any of "novnc", "spice".
	ConsoleType InstanceConsoleGetParamsConsoleType `query:"console_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceConsoleGetParams]'s query parameters as
// `url.Values`.
func (r InstanceConsoleGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Console type that you expect. If you omit it, you get the first available type
// in this order: `novnc`, `spice`.
type InstanceConsoleGetParamsConsoleType string

const (
	InstanceConsoleGetParamsConsoleTypeNovnc InstanceConsoleGetParamsConsoleType = "novnc"
	InstanceConsoleGetParamsConsoleTypeSpice InstanceConsoleGetParamsConsoleType = "spice"
)
