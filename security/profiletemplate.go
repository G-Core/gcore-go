// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"context"
	"net/http"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ProfileTemplateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileTemplateService] method instead.
type ProfileTemplateService struct {
	Options []option.RequestOption
}

// NewProfileTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileTemplateService(opts ...option.RequestOption) (r ProfileTemplateService) {
	r = ProfileTemplateService{}
	r.Options = opts
	return
}

// Get list of profile templates. Profile template is used as a template to create
// profile. Client receives only common and created for him profile templates.
func (r *ProfileTemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ClientProfileTemplate, err error) {
	opts = append(r.Options[:], opts...)
	path := "security/iaas/profile-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientProfileTemplate struct {
	ID             int64                        `json:"id,required"`
	Created        time.Time                    `json:"created,required" format:"date-time"`
	Fields         []ClientProfileTemplateField `json:"fields,required"`
	Name           string                       `json:"name,required"`
	Version        string                       `json:"version,required" format:"uuid"`
	BaseTemplate   int64                        `json:"base_template,nullable"`
	Description    string                       `json:"description"`
	TemplateSifter string                       `json:"template_sifter,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Created        respjson.Field
		Fields         respjson.Field
		Name           respjson.Field
		Version        respjson.Field
		BaseTemplate   respjson.Field
		Description    respjson.Field
		TemplateSifter respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientProfileTemplate) RawJSON() string { return r.JSON.raw }
func (r *ClientProfileTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientProfileTemplateField struct {
	ID          int64  `json:"id,required"`
	Name        string `json:"name,required"`
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	// Any of "int", "bool", "str".
	FieldType        ClientProfileTemplateFieldFieldType `json:"field_type,nullable"`
	Required         bool                                `json:"required"`
	ValidationSchema map[string]any                      `json:"validation_schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Name             respjson.Field
		Default          respjson.Field
		Description      respjson.Field
		FieldType        respjson.Field
		Required         respjson.Field
		ValidationSchema respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientProfileTemplateField) RawJSON() string { return r.JSON.raw }
func (r *ClientProfileTemplateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientProfileTemplateFieldFieldType string

const (
	ClientProfileTemplateFieldFieldTypeInt  ClientProfileTemplateFieldFieldType = "int"
	ClientProfileTemplateFieldFieldTypeBool ClientProfileTemplateFieldFieldType = "bool"
	ClientProfileTemplateFieldFieldTypeStr  ClientProfileTemplateFieldFieldType = "str"
)
