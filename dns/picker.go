// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// PickerService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPickerService] method instead.
type PickerService struct {
	Options []option.RequestOption
	Presets PickerPresetService
}

// NewPickerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPickerService(opts ...option.RequestOption) (r PickerService) {
	r = PickerService{}
	r.Options = opts
	r.Presets = NewPickerPresetService(opts...)
	return
}

// Returns list of picker
func (r *PickerService) List(ctx context.Context, opts ...option.RequestOption) (res *[]DNSLabelName, err error) {
	opts = append(r.Options[:], opts...)
	path := "dns/v2/pickers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSLabelName struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSLabelName) RawJSON() string { return r.JSON.raw }
func (r *DNSLabelName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
