// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainAPIDiscoveryOpenAPIService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIDiscoveryOpenAPIService] method instead.
type DomainAPIDiscoveryOpenAPIService struct {
	Options []option.RequestOption
}

// NewDomainAPIDiscoveryOpenAPIService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoveryOpenAPIService(opts ...option.RequestOption) (r DomainAPIDiscoveryOpenAPIService) {
	r = DomainAPIDiscoveryOpenAPIService{}
	r.Options = opts
	return
}

// Scan an API description file hosted online. The file must be in YAML or JSON
// format and adhere to the OpenAPI specification. The location of the API
// description file should be specified in the API discovery settings.
func (r *DomainAPIDiscoveryOpenAPIService) Scan(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapTaskID, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An API description file must adhere to the OpenAPI specification and be written
// in YAML or JSON format. The file name should be provided as the value for the
// `file_name` parameter. The contents of the file must be base64 encoded and
// supplied as the value for the `file_data` parameter.
func (r *DomainAPIDiscoveryOpenAPIService) Upload(ctx context.Context, domainID int64, body DomainAPIDiscoveryOpenAPIUploadParams, opts ...option.RequestOption) (res *WaapTaskID, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/upload", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model for the task result ID
type WaapTaskID struct {
	// The task ID
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTaskID) RawJSON() string { return r.JSON.raw }
func (r *WaapTaskID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainAPIDiscoveryOpenAPIUploadParams struct {
	// Base64 representation of the description file. Supported formats are YAML and
	// JSON, and it must adhere to OpenAPI versions 2, 3, or 3.1.
	FileData string `json:"file_data,required"`
	// The name of the file
	FileName string `json:"file_name,required"`
	paramObj
}

func (r DomainAPIDiscoveryOpenAPIUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAPIDiscoveryOpenAPIUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAPIDiscoveryOpenAPIUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
