// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// SecretService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r SecretService) {
	r = SecretService{}
	r.Options = opts
	return
}

// Create secret
func (r *SecretService) New(ctx context.Context, params SecretNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List secrets
func (r *SecretService) List(ctx context.Context, query SecretListParams, opts ...option.RequestOption) (res *SecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete secret
func (r *SecretService) Delete(ctx context.Context, secretID string, body SecretDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get secret
func (r *SecretService) Get(ctx context.Context, secretID string, query SecretGetParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create secret
func (r *SecretService) UploadTlsCertificate(ctx context.Context, params SecretUploadTlsCertificateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/secrets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// '#/components/schemas/SecretSerializer' "$.components.schemas.SecretSerializer"
type Secret struct {
	// '#/components/schemas/SecretSerializer/properties/id'
	// "$.components.schemas.SecretSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/SecretSerializer/properties/name'
	// "$.components.schemas.SecretSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/SecretSerializer/properties/secret_type'
	// "$.components.schemas.SecretSerializer.properties.secret_type"
	//
	// Any of "certificate", "opaque", "passphrase", "private", "public", "symmetric".
	SecretType SecretSecretType `json:"secret_type,required"`
	// '#/components/schemas/SecretSerializer/properties/status'
	// "$.components.schemas.SecretSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/SecretSerializer/properties/algorithm/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.algorithm.anyOf[0]"
	Algorithm string `json:"algorithm,nullable"`
	// '#/components/schemas/SecretSerializer/properties/bit_length/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.bit_length.anyOf[0]"
	BitLength int64 `json:"bit_length,nullable"`
	// '#/components/schemas/SecretSerializer/properties/content_types/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.content_types.anyOf[0]"
	ContentTypes map[string]string `json:"content_types,nullable"`
	// '#/components/schemas/SecretSerializer/properties/created/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.created.anyOf[0]"
	Created time.Time `json:"created,nullable" format:"date-time"`
	// '#/components/schemas/SecretSerializer/properties/expiration/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.expiration.anyOf[0]"
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// '#/components/schemas/SecretSerializer/properties/mode/anyOf/0'
	// "$.components.schemas.SecretSerializer.properties.mode.anyOf[0]"
	Mode string `json:"mode,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		Name         resp.Field
		SecretType   resp.Field
		Status       resp.Field
		Algorithm    resp.Field
		BitLength    resp.Field
		ContentTypes resp.Field
		Created      resp.Field
		Expiration   resp.Field
		Mode         resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SecretSerializer/properties/secret_type'
// "$.components.schemas.SecretSerializer.properties.secret_type"
type SecretSecretType string

const (
	SecretSecretTypeCertificate SecretSecretType = "certificate"
	SecretSecretTypeOpaque      SecretSecretType = "opaque"
	SecretSecretTypePassphrase  SecretSecretType = "passphrase"
	SecretSecretTypePrivate     SecretSecretType = "private"
	SecretSecretTypePublic      SecretSecretType = "public"
	SecretSecretTypeSymmetric   SecretSecretType = "symmetric"
)

// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}'].get.responses[200].content['application/json'].schema"
type SecretListResponse struct {
	// '#/components/schemas/SecretSerializerList/properties/count'
	// "$.components.schemas.SecretSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/SecretSerializerList/properties/results'
	// "$.components.schemas.SecretSerializerList.properties.results"
	Results []Secret `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *SecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSecretSerializer/properties/name'
	// "$.components.schemas.CreateSecretSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateSecretSerializer/properties/payload'
	// "$.components.schemas.CreateSecretSerializer.properties.payload"
	Payload string `json:"payload,required" format:"password"`
	// '#/components/schemas/CreateSecretSerializer/properties/payload_content_encoding'
	// "$.components.schemas.CreateSecretSerializer.properties.payload_content_encoding"
	//
	// Any of "base64".
	PayloadContentEncoding SecretNewParamsPayloadContentEncoding `json:"payload_content_encoding,omitzero,required"`
	// '#/components/schemas/CreateSecretSerializer/properties/payload_content_type'
	// "$.components.schemas.CreateSecretSerializer.properties.payload_content_type"
	PayloadContentType string `json:"payload_content_type,required"`
	// '#/components/schemas/CreateSecretSerializer/properties/secret_type'
	// "$.components.schemas.CreateSecretSerializer.properties.secret_type"
	//
	// Any of "certificate", "opaque", "passphrase", "private", "public", "symmetric".
	SecretType SecretNewParamsSecretType `json:"secret_type,omitzero,required"`
	// '#/components/schemas/CreateSecretSerializer/properties/algorithm/anyOf/0'
	// "$.components.schemas.CreateSecretSerializer.properties.algorithm.anyOf[0]"
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// '#/components/schemas/CreateSecretSerializer/properties/bit_length/anyOf/0'
	// "$.components.schemas.CreateSecretSerializer.properties.bit_length.anyOf[0]"
	BitLength param.Opt[int64] `json:"bit_length,omitzero"`
	// '#/components/schemas/CreateSecretSerializer/properties/expiration/anyOf/0'
	// "$.components.schemas.CreateSecretSerializer.properties.expiration.anyOf[0]"
	Expiration param.Opt[string] `json:"expiration,omitzero"`
	// '#/components/schemas/CreateSecretSerializer/properties/mode/anyOf/0'
	// "$.components.schemas.CreateSecretSerializer.properties.mode.anyOf[0]"
	Mode param.Opt[string] `json:"mode,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateSecretSerializer/properties/payload_content_encoding'
// "$.components.schemas.CreateSecretSerializer.properties.payload_content_encoding"
type SecretNewParamsPayloadContentEncoding string

const (
	SecretNewParamsPayloadContentEncodingBase64 SecretNewParamsPayloadContentEncoding = "base64"
)

// '#/components/schemas/CreateSecretSerializer/properties/secret_type'
// "$.components.schemas.CreateSecretSerializer.properties.secret_type"
type SecretNewParamsSecretType string

const (
	SecretNewParamsSecretTypeCertificate SecretNewParamsSecretType = "certificate"
	SecretNewParamsSecretTypeOpaque      SecretNewParamsSecretType = "opaque"
	SecretNewParamsSecretTypePassphrase  SecretNewParamsSecretType = "passphrase"
	SecretNewParamsSecretTypePrivate     SecretNewParamsSecretType = "private"
	SecretNewParamsSecretTypePublic      SecretNewParamsSecretType = "public"
	SecretNewParamsSecretTypeSymmetric   SecretNewParamsSecretType = "symmetric"
)

type SecretListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecretDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsecret_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}/{secret_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsecret_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}/{secret_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecretGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsecret_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}/{secret_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bsecret_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/secrets/{project_id}/{region_id}/{secret_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecretUploadTlsCertificateParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v2/secrets/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fsecrets%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v2/secrets/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSecretSerializerV2/properties/name'
	// "$.components.schemas.CreateSecretSerializerV2.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateSecretSerializerV2/properties/payload'
	// "$.components.schemas.CreateSecretSerializerV2.properties.payload"
	Payload SecretUploadTlsCertificateParamsPayload `json:"payload,omitzero,required"`
	// '#/components/schemas/CreateSecretSerializerV2/properties/expiration/anyOf/0'
	// "$.components.schemas.CreateSecretSerializerV2.properties.expiration.anyOf[0]"
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretUploadTlsCertificateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecretUploadTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecretUploadTlsCertificateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateSecretSerializerV2/properties/payload'
// "$.components.schemas.CreateSecretSerializerV2.properties.payload"
//
// The properties Certificate, CertificateChain, PrivateKey are required.
type SecretUploadTlsCertificateParamsPayload struct {
	// '#/components/schemas/CreateSecretPayloadSerializer/properties/certificate'
	// "$.components.schemas.CreateSecretPayloadSerializer.properties.certificate"
	Certificate string `json:"certificate,required" format:"password"`
	// '#/components/schemas/CreateSecretPayloadSerializer/properties/certificate_chain'
	// "$.components.schemas.CreateSecretPayloadSerializer.properties.certificate_chain"
	CertificateChain string `json:"certificate_chain,required" format:"password"`
	// '#/components/schemas/CreateSecretPayloadSerializer/properties/private_key'
	// "$.components.schemas.CreateSecretPayloadSerializer.properties.private_key"
	PrivateKey string `json:"private_key,required" format:"password"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecretUploadTlsCertificateParamsPayload) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r SecretUploadTlsCertificateParamsPayload) MarshalJSON() (data []byte, err error) {
	type shadow SecretUploadTlsCertificateParamsPayload
	return param.MarshalObject(r, (*shadow)(&r))
}
