// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainAPIDiscoveryService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAPIDiscoveryService] method instead.
type WaapV1DomainAPIDiscoveryService struct {
	Options     []option.RequestOption
	Settings    *WaapV1DomainAPIDiscoverySettingService
	APIURLs     *WaapV1DomainAPIDiscoveryAPIURLService
	ScanResults *WaapV1DomainAPIDiscoveryScanResultService
}

// NewWaapV1DomainAPIDiscoveryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainAPIDiscoveryService(opts ...option.RequestOption) (r *WaapV1DomainAPIDiscoveryService) {
	r = &WaapV1DomainAPIDiscoveryService{}
	r.Options = opts
	r.Settings = NewWaapV1DomainAPIDiscoverySettingService(opts...)
	r.APIURLs = NewWaapV1DomainAPIDiscoveryAPIURLService(opts...)
	r.ScanResults = NewWaapV1DomainAPIDiscoveryScanResultService(opts...)
	return
}

// Scan an API description file hosted online. The file must be in YAML or JSON
// format and adhere to the OpenAPI specification. The location of the API
// description file should be specified in the API discovery settings.
func (r *WaapV1DomainAPIDiscoveryService) Scan(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *TaskResultID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/scan", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An API description file must adhere to the OpenAPI specification and be written
// in YAML or JSON format. The file name should be provided as the value for the
// `file_name` parameter. The contents of the file must be base64 encoded and
// supplied as the value for the `file_data` parameter.
func (r *WaapV1DomainAPIDiscoveryService) Upload(ctx context.Context, domainID int64, body WaapV1DomainAPIDiscoveryUploadParams, opts ...option.RequestOption) (res *TaskResultID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/upload", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model for the task result ID
type TaskResultID struct {
	// The task ID
	ID   string           `json:"id,required" format:"uuid"`
	JSON taskResultIDJSON `json:"-"`
}

// taskResultIDJSON contains the JSON metadata for the struct [TaskResultID]
type taskResultIDJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskResultID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskResultIDJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAPIDiscoveryUploadParams struct {
	// Base64 representation of the description file. Supported formats are YAML and
	// JSON, and it must adhere to OpenAPI versions 2, 3, or 3.1.
	FileData param.Field[string] `json:"file_data,required"`
	// The name of the file
	FileName param.Field[string] `json:"file_name,required"`
}

func (r WaapV1DomainAPIDiscoveryUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
