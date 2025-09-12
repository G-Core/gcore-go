// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// PickerPresetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPickerPresetService] method instead.
type PickerPresetService struct {
	Options []option.RequestOption
}

// NewPickerPresetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPickerPresetService(opts ...option.RequestOption) (r PickerPresetService) {
	r = PickerPresetService{}
	r.Options = opts
	return
}

// Returns list of picker preset
func (r *PickerPresetService) List(ctx context.Context, opts ...option.RequestOption) (res *PickerPresetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "dns/v2/pickers/presets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PickerPresetListResponse map[string][]DNSLabelName
