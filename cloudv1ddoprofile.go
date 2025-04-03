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

// CloudV1DdoProfileService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DdoProfileService] method instead.
type CloudV1DdoProfileService struct {
	Options []option.RequestOption
}

// NewCloudV1DdoProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1DdoProfileService(opts ...option.RequestOption) (r *CloudV1DdoProfileService) {
	r = &CloudV1DdoProfileService{}
	r.Options = opts
	return
}

// Create DDoS protection profile
func (r *CloudV1DdoProfileService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1DdoProfileNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profiles/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update DDoS protection profile
func (r *CloudV1DdoProfileService) Update(ctx context.Context, projectID int64, regionID int64, profileID int64, body CloudV1DdoProfileUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profiles/%v/%v/%v", projectID, regionID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List active client DDoS protection profiles
func (r *CloudV1DdoProfileService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1DdoProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profiles/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete DDoS protection profile
func (r *CloudV1DdoProfileService) Delete(ctx context.Context, projectID int64, regionID int64, profileID int64, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profiles/%v/%v/%v", projectID, regionID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Activate or Deactivate DDoS protection profile
func (r *CloudV1DdoProfileService) ActivateDeactivate(ctx context.Context, projectID int64, regionID int64, profileID int64, body CloudV1DdoProfileActivateDeactivateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/ddos/profiles/%v/%v/%v/action", projectID, regionID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ClientProfileField struct {
	ID               int64                  `json:"id,required"`
	Default          interface{}            `json:"default,required"`
	Description      string                 `json:"description,required"`
	FieldValue       interface{}            `json:"field_value,required"`
	Name             string                 `json:"name,required"`
	BaseField        int64                  `json:"base_field,nullable"`
	FieldName        string                 `json:"field_name,nullable"`
	FieldType        string                 `json:"field_type,nullable"`
	Required         bool                   `json:"required,nullable"`
	ValidationSchema interface{}            `json:"validation_schema"`
	Value            string                 `json:"value,nullable"`
	JSON             clientProfileFieldJSON `json:"-"`
}

// clientProfileFieldJSON contains the JSON metadata for the struct
// [ClientProfileField]
type clientProfileFieldJSON struct {
	ID               apijson.Field
	Default          apijson.Field
	Description      apijson.Field
	FieldValue       apijson.Field
	Name             apijson.Field
	BaseField        apijson.Field
	FieldName        apijson.Field
	FieldType        apijson.Field
	Required         apijson.Field
	ValidationSchema apijson.Field
	Value            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ClientProfileField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientProfileFieldJSON) RawJSON() string {
	return r.raw
}

type ClientProfileTemplate struct {
	ID          int64                        `json:"id,required"`
	Name        string                       `json:"name,required"`
	Description string                       `json:"description,nullable"`
	Fields      []ClientProfileTemplateField `json:"fields"`
	JSON        clientProfileTemplateJSON    `json:"-"`
}

// clientProfileTemplateJSON contains the JSON metadata for the struct
// [ClientProfileTemplate]
type clientProfileTemplateJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Description apijson.Field
	Fields      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientProfileTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientProfileTemplateJSON) RawJSON() string {
	return r.raw
}

type CreateClientProfileParam struct {
	Fields param.Field[[]CreateClientProfileFieldParam] `json:"fields,required"`
	// IP address to be protected
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipv4"`
	// Advanced DDoS template ID
	ProfileTemplate param.Field[int64] `json:"profile_template,required"`
	// Deprecated. Use resource_id field
	BmInstanceID param.Field[string] `json:"bm_instance_id"`
	// DDoS profile template name
	ProfileTemplateName param.Field[string] `json:"profile_template_name"`
	// ID of resource (bare metal, load balancer, instance)
	ResourceID param.Field[string] `json:"resource_id"`
	// Resource type to be protected
	ResourceType param.Field[CreateClientProfileResourceType] `json:"resource_type"`
}

func (r CreateClientProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateClientProfileFieldParam struct {
	// ID of DDoS profile field
	BaseField param.Field[int64] `json:"base_field"`
	// Name of DDoS profile field
	FieldName param.Field[string] `json:"field_name"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue param.Field[CreateClientProfileFieldsFieldValueUnionParam] `json:"field_value"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Field[string] `json:"value"`
}

func (r CreateClientProfileFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Complex value. Only one of 'value' or 'field_value' must be specified.
//
// Satisfied by [CreateClientProfileFieldsFieldValueArrayParam], [shared.UnionInt],
// [shared.UnionString].
type CreateClientProfileFieldsFieldValueUnionParam interface {
	ImplementsCreateClientProfileFieldsFieldValueUnionParam()
}

type CreateClientProfileFieldsFieldValueArrayParam []interface{}

func (r CreateClientProfileFieldsFieldValueArrayParam) ImplementsCreateClientProfileFieldsFieldValueUnionParam() {
}

// Resource type to be protected
type CreateClientProfileResourceType string

const (
	CreateClientProfileResourceTypeInstance     CreateClientProfileResourceType = "instance"
	CreateClientProfileResourceTypeLoadbalancer CreateClientProfileResourceType = "loadbalancer"
)

func (r CreateClientProfileResourceType) IsKnown() bool {
	switch r {
	case CreateClientProfileResourceTypeInstance, CreateClientProfileResourceTypeLoadbalancer:
		return true
	}
	return false
}

type CloudV1DdoProfileListResponse struct {
	ID              int64                                   `json:"id,required"`
	Fields          []ClientProfileField                    `json:"fields,required"`
	IPAddress       string                                  `json:"ip_address,required,nullable"`
	Options         CloudV1DdoProfileListResponseOptions    `json:"options,required"`
	ProfileTemplate ClientProfileTemplate                   `json:"profile_template,required"`
	Protocols       []CloudV1DdoProfileListResponseProtocol `json:"protocols,required"`
	Site            string                                  `json:"site,required"`
	Status          CloudV1DdoProfileListResponseStatus     `json:"status,required"`
	JSON            cloudV1DdoProfileListResponseJSON       `json:"-"`
}

// cloudV1DdoProfileListResponseJSON contains the JSON metadata for the struct
// [CloudV1DdoProfileListResponse]
type cloudV1DdoProfileListResponseJSON struct {
	ID              apijson.Field
	Fields          apijson.Field
	IPAddress       apijson.Field
	Options         apijson.Field
	ProfileTemplate apijson.Field
	Protocols       apijson.Field
	Site            apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CloudV1DdoProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoProfileListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoProfileListResponseOptions struct {
	Active bool                                     `json:"active,required"`
	Bgp    bool                                     `json:"bgp,required"`
	JSON   cloudV1DdoProfileListResponseOptionsJSON `json:"-"`
}

// cloudV1DdoProfileListResponseOptionsJSON contains the JSON metadata for the
// struct [CloudV1DdoProfileListResponseOptions]
type cloudV1DdoProfileListResponseOptionsJSON struct {
	Active      apijson.Field
	Bgp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DdoProfileListResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoProfileListResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoProfileListResponseProtocol struct {
	Port      string                                    `json:"port,required"`
	Protocols []string                                  `json:"protocols,required"`
	JSON      cloudV1DdoProfileListResponseProtocolJSON `json:"-"`
}

// cloudV1DdoProfileListResponseProtocolJSON contains the JSON metadata for the
// struct [CloudV1DdoProfileListResponseProtocol]
type cloudV1DdoProfileListResponseProtocolJSON struct {
	Port        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DdoProfileListResponseProtocol) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoProfileListResponseProtocolJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoProfileListResponseStatus struct {
	Status           string                                  `json:"status,required"`
	ErrorDescription string                                  `json:"error_description"`
	JSON             cloudV1DdoProfileListResponseStatusJSON `json:"-"`
}

// cloudV1DdoProfileListResponseStatusJSON contains the JSON metadata for the
// struct [CloudV1DdoProfileListResponseStatus]
type cloudV1DdoProfileListResponseStatusJSON struct {
	Status           apijson.Field
	ErrorDescription apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudV1DdoProfileListResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DdoProfileListResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudV1DdoProfileNewParams struct {
	CreateClientProfile CreateClientProfileParam `json:"create_client_profile,required"`
}

func (r CloudV1DdoProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateClientProfile)
}

type CloudV1DdoProfileUpdateParams struct {
	CreateClientProfile CreateClientProfileParam `json:"create_client_profile,required"`
}

func (r CloudV1DdoProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateClientProfile)
}

type CloudV1DdoProfileActivateDeactivateParams struct {
	Active param.Field[bool] `json:"active,required"`
	Bgp    param.Field[bool] `json:"bgp,required"`
}

func (r CloudV1DdoProfileActivateDeactivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
