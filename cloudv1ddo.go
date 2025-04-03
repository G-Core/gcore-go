// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1DdoService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DdoService] method instead.
type CloudV1DdoService struct {
	Options  []option.RequestOption
	Profiles *CloudV1DdoProfileService
}

// NewCloudV1DdoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1DdoService(opts ...option.RequestOption) (r *CloudV1DdoService) {
	r = &CloudV1DdoService{}
	r.Options = opts
	r.Profiles = NewCloudV1DdoProfileService(opts...)
	return
}

// Check if the provided region is covered by the Advanced DDoS protection feature
func (r *CloudV1DdoService) CheckRegionCoverage(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *CloudV1DdoCheckRegionCoverageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/region_coverage/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return DDoS Protection service access status
func (r *CloudV1DdoService) GetAccessibility(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *CloudV1DdoGetAccessibilityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/accessibility/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of available DDoS protection profile templates
func (r *CloudV1DdoService) ListProfileTemplates(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *CloudV1DdoListProfileTemplatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profile-templates/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientProfileTemplateField struct {
	ID               int64                          `json:"id,required"`
	Name             string                         `json:"name,required"`
	Default          string                         `json:"default,nullable"`
	Description      string                         `json:"description,nullable"`
	FieldType        string                         `json:"field_type,nullable"`
	Required         bool                           `json:"required,nullable"`
	ValidationSchema interface{}                    `json:"validation_schema"`
	JSON             clientProfileTemplateFieldJSON `json:"-"`
}

// clientProfileTemplateFieldJSON contains the JSON metadata for the struct
// [ClientProfileTemplateField]
type clientProfileTemplateFieldJSON struct {
	ID               apijson.Field
	Name             apijson.Field
	Default          apijson.Field
	Description      apijson.Field
	FieldType        apijson.Field
	Required         apijson.Field
	ValidationSchema apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ClientProfileTemplateField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientProfileTemplateFieldJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoCheckRegionCoverageResponse struct {
	// Is the current region covered by the Advanced DDoS protection features.
	IsCovered bool                                      `json:"is_covered,required"`
	JSON      cloudV1DdoCheckRegionCoverageResponseJSON `json:"-"`
}

// cloudV1DdoCheckRegionCoverageResponseJSON contains the JSON metadata for the
// struct [CloudV1DdoCheckRegionCoverageResponse]
type cloudV1DdoCheckRegionCoverageResponseJSON struct {
	IsCovered   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DdoCheckRegionCoverageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoCheckRegionCoverageResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoGetAccessibilityResponse struct {
	HTTPCode     int64                                  `json:"http_code,required"`
	IsAccessible bool                                   `json:"is_accessible,required"`
	Message      string                                 `json:"message,required"`
	JSON         cloudV1DdoGetAccessibilityResponseJSON `json:"-"`
}

// cloudV1DdoGetAccessibilityResponseJSON contains the JSON metadata for the struct
// [CloudV1DdoGetAccessibilityResponse]
type cloudV1DdoGetAccessibilityResponseJSON struct {
	HTTPCode     apijson.Field
	IsAccessible apijson.Field
	Message      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CloudV1DdoGetAccessibilityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoGetAccessibilityResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoListProfileTemplatesResponse struct {
	ID          int64                                      `json:"id,required"`
	Fields      []ClientProfileTemplateField               `json:"fields,required"`
	Name        string                                     `json:"name,required"`
	Description string                                     `json:"description,nullable"`
	JSON        cloudV1DdoListProfileTemplatesResponseJSON `json:"-"`
}

// cloudV1DdoListProfileTemplatesResponseJSON contains the JSON metadata for the
// struct [CloudV1DdoListProfileTemplatesResponse]
type cloudV1DdoListProfileTemplatesResponseJSON struct {
	ID          apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DdoListProfileTemplatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoListProfileTemplatesResponseJSON) RawJSON() string {
	return r.raw
}
