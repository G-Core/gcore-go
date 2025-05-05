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
	"github.com/stainless-sdks/gcore-go/packages/respjson"
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
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
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/secrets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Secret struct {
	// Secret uuid
	ID string `json:"id,required"`
	// Secret name
	Name string `json:"name,required"`
	// Secret type, base64 encoded. symmetric - Used for storing byte arrays such as
	// keys suitable for symmetric encryption; public - Used for storing the public key
	// of an asymmetric keypair; private - Used for storing the private key of an
	// asymmetric keypair; passphrase - Used for storing plain text passphrases;
	// certificate - Used for storing cryptographic certificates such as X.509
	// certificates; opaque - Used for backwards compatibility with previous versions
	// of the API
	//
	// Any of "certificate", "opaque", "passphrase", "private", "public", "symmetric".
	SecretType SecretSecretType `json:"secret_type,required"`
	// Status
	Status string `json:"status,required"`
	// Metadata provided by a user or system for informational purposes. Defaults to
	// None
	Algorithm string `json:"algorithm,nullable"`
	// Metadata provided by a user or system for informational purposes. Value must be
	// greater than zero. Defaults to None
	BitLength int64 `json:"bit_length,nullable"`
	// Describes the content-types that can be used to retrieve the payload. The
	// content-type used with symmetric secrets is application/octet-stream
	ContentTypes map[string]string `json:"content_types,nullable"`
	// Datetime when the secret was created. The format is 2020-01-01T12:00:00+00:00
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Datetime when the secret will expire. The format is 2020-01-01T12:00:00+00:00.
	// Defaults to None
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Metadata provided by a user or system for informational purposes. Defaults to
	// None
	Mode string `json:"mode,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		SecretType   respjson.Field
		Status       respjson.Field
		Algorithm    respjson.Field
		BitLength    respjson.Field
		ContentTypes respjson.Field
		Created      respjson.Field
		Expiration   respjson.Field
		Mode         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret type, base64 encoded. symmetric - Used for storing byte arrays such as
// keys suitable for symmetric encryption; public - Used for storing the public key
// of an asymmetric keypair; private - Used for storing the private key of an
// asymmetric keypair; passphrase - Used for storing plain text passphrases;
// certificate - Used for storing cryptographic certificates such as X.509
// certificates; opaque - Used for backwards compatibility with previous versions
// of the API
type SecretSecretType string

const (
	SecretSecretTypeCertificate SecretSecretType = "certificate"
	SecretSecretTypeOpaque      SecretSecretType = "opaque"
	SecretSecretTypePassphrase  SecretSecretType = "passphrase"
	SecretSecretTypePrivate     SecretSecretType = "private"
	SecretSecretTypePublic      SecretSecretType = "public"
	SecretSecretTypeSymmetric   SecretSecretType = "symmetric"
)

type SecretListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Secret `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *SecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Secret name
	Name string `json:"name,required"`
	// Secret payload. For HTTPS-terminated load balancing, provide base64 encoded
	// conents of a PKCS12 file. The PKCS12 file is the combined TLS certificate, key,
	// and intermediate certificate chain obtained from an external certificate
	// authority. The file can be created via openssl, e.g.'openssl pkcs12 -export
	// -inkey server.key -in server.crt -certfile ca-chain.crt -passout pass: -out
	// server.p12'The key and certificate should be PEM-encoded, and the intermediate
	// certificate chain should be multiple PEM-encoded certs concatenated together
	Payload string `json:"payload,required" format:"password"`
	// The encoding used for the payload to be able to include it in the JSON request.
	// Currently only base64 is supported
	//
	// Any of "base64".
	PayloadContentEncoding SecretNewParamsPayloadContentEncoding `json:"payload_content_encoding,omitzero,required"`
	// The media type for the content of the payload
	PayloadContentType string `json:"payload_content_type,required"`
	// Secret type. symmetric - Used for storing byte arrays such as keys suitable for
	// symmetric encryption; public - Used for storing the public key of an asymmetric
	// keypair; private - Used for storing the private key of an asymmetric keypair;
	// passphrase - Used for storing plain text passphrases; certificate - Used for
	// storing cryptographic certificates such as X.509 certificates; opaque - Used for
	// backwards compatibility with previous versions of the API
	//
	// Any of "certificate", "opaque", "passphrase", "private", "public", "symmetric".
	SecretType SecretNewParamsSecretType `json:"secret_type,omitzero,required"`
	// Metadata provided by a user or system for informational purposes.
	Algorithm param.Opt[string] `json:"algorithm,omitzero"`
	// Metadata provided by a user or system for informational purposes. Value must be
	// greater than zero.
	BitLength param.Opt[int64] `json:"bit_length,omitzero"`
	// Datetime when the secret will expire.
	Expiration param.Opt[string] `json:"expiration,omitzero"`
	// Metadata provided by a user or system for informational purposes.
	Mode param.Opt[string] `json:"mode,omitzero"`
	paramObj
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The encoding used for the payload to be able to include it in the JSON request.
// Currently only base64 is supported
type SecretNewParamsPayloadContentEncoding string

const (
	SecretNewParamsPayloadContentEncodingBase64 SecretNewParamsPayloadContentEncoding = "base64"
)

// Secret type. symmetric - Used for storing byte arrays such as keys suitable for
// symmetric encryption; public - Used for storing the public key of an asymmetric
// keypair; private - Used for storing the private key of an asymmetric keypair;
// passphrase - Used for storing plain text passphrases; certificate - Used for
// storing cryptographic certificates such as X.509 certificates; opaque - Used for
// backwards compatibility with previous versions of the API
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecretDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecretGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecretUploadTlsCertificateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Secret name
	Name string `json:"name,required"`
	// Secret payload.
	Payload SecretUploadTlsCertificateParamsPayload `json:"payload,omitzero,required"`
	// Datetime when the secret will expire. Defaults to None
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	paramObj
}

func (r SecretUploadTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecretUploadTlsCertificateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Secret payload.
//
// The properties Certificate, CertificateChain, PrivateKey are required.
type SecretUploadTlsCertificateParamsPayload struct {
	// SSL certificate in PEM format.
	Certificate string `json:"certificate,required" format:"password"`
	// SSL certificate chain of intermediates and root certificates in PEM format.
	CertificateChain string `json:"certificate_chain,required" format:"password"`
	// SSL private key in PEM format.
	PrivateKey string `json:"private_key,required" format:"password"`
	paramObj
}

func (r SecretUploadTlsCertificateParamsPayload) MarshalJSON() (data []byte, err error) {
	type shadow SecretUploadTlsCertificateParamsPayload
	return param.MarshalObject(r, (*shadow)(&r))
}
