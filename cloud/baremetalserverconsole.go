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

// BaremetalServerConsoleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalServerConsoleService] method instead.
type BaremetalServerConsoleService struct {
	Options []option.RequestOption
}

// NewBaremetalServerConsoleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBaremetalServerConsoleService(opts ...option.RequestOption) (r BaremetalServerConsoleService) {
	r = BaremetalServerConsoleService{}
	r.Options = opts
	return
}

// Returns a URL that gives access to the remote console of a bare metal server.
func (r *BaremetalServerConsoleService) Get(ctx context.Context, serverID string, params BaremetalServerConsoleGetParams, opts ...option.RequestOption) (res *Console, err error) {
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
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers/%s/console", params.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type BaremetalServerConsoleGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Console type you expect the server to give. The server hardware sets the
	// available types. A value that the hardware does not give causes an error. Omit
	// this parameter to accept the type of the server.
	//
	// Any of "serial", "vnc".
	ConsoleType BaremetalServerConsoleGetParamsConsoleType `query:"console_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerConsoleGetParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerConsoleGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Console type you expect the server to give. The server hardware sets the
// available types. A value that the hardware does not give causes an error. Omit
// this parameter to accept the type of the server.
type BaremetalServerConsoleGetParamsConsoleType string

const (
	BaremetalServerConsoleGetParamsConsoleTypeSerial BaremetalServerConsoleGetParamsConsoleType = "serial"
	BaremetalServerConsoleGetParamsConsoleTypeVnc    BaremetalServerConsoleGetParamsConsoleType = "vnc"
)
