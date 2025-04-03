// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1SecretService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SecretService] method instead.
type CloudV1SecretService struct {
	Options []option.RequestOption
}

// NewCloudV1SecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1SecretService(opts ...option.RequestOption) (r *CloudV1SecretService) {
	r = &CloudV1SecretService{}
	r.Options = opts
	return
}

// Create secret
func (r *CloudV1SecretService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1SecretNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get secret
func (r *CloudV1SecretService) Get(ctx context.Context, projectID int64, regionID int64, secretID string, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v/%s", projectID, regionID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List secrets
func (r *CloudV1SecretService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1SecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete secret
func (r *CloudV1SecretService) Delete(ctx context.Context, projectID int64, regionID int64, secretID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/secrets/%v/%v/%s", projectID, regionID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	SecretType SecretTypes `json:"secret_type,required"`
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
	Mode string     `json:"mode,nullable"`
	JSON secretJSON `json:"-"`
}

// secretJSON contains the JSON metadata for the struct [Secret]
type secretJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	SecretType   apijson.Field
	Status       apijson.Field
	Algorithm    apijson.Field
	BitLength    apijson.Field
	ContentTypes apijson.Field
	Created      apijson.Field
	Expiration   apijson.Field
	Mode         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Secret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretJSON) RawJSON() string {
	return r.raw
}

type SecretTypes string

const (
	SecretTypesCertificate SecretTypes = "certificate"
	SecretTypesOpaque      SecretTypes = "opaque"
	SecretTypesPassphrase  SecretTypes = "passphrase"
	SecretTypesPrivate     SecretTypes = "private"
	SecretTypesPublic      SecretTypes = "public"
	SecretTypesSymmetric   SecretTypes = "symmetric"
)

func (r SecretTypes) IsKnown() bool {
	switch r {
	case SecretTypesCertificate, SecretTypesOpaque, SecretTypesPassphrase, SecretTypesPrivate, SecretTypesPublic, SecretTypesSymmetric:
		return true
	}
	return false
}

type CloudV1SecretListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Secret                      `json:"results,required"`
	JSON    cloudV1SecretListResponseJSON `json:"-"`
}

// cloudV1SecretListResponseJSON contains the JSON metadata for the struct
// [CloudV1SecretListResponse]
type cloudV1SecretListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1SecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1SecretListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1SecretNewParams struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret payload. For HTTPS-terminated load balancing, provide base64 encoded
	// conents of a PKCS12 file. The PKCS12 file is the combined TLS certificate, key,
	// and intermediate certificate chain obtained from an external certificate
	// authority. The file can be created via openssl, e.g.'openssl pkcs12 -export
	// -inkey server.key -in server.crt -certfile ca-chain.crt -passout pass: -out
	// server.p12'The key and certificate should be PEM-encoded, and the intermediate
	// certificate chain should be multiple PEM-encoded certs concatenated together
	Payload param.Field[string] `json:"payload,required" format:"password"`
	// The encoding used for the payload to be able to include it in the JSON request.
	// Currently only base64 is supported
	PayloadContentEncoding param.Field[CloudV1SecretNewParamsPayloadContentEncoding] `json:"payload_content_encoding,required"`
	// The media type for the content of the payload
	PayloadContentType param.Field[string] `json:"payload_content_type,required"`
	// Secret type. symmetric - Used for storing byte arrays such as keys suitable for
	// symmetric encryption; public - Used for storing the public key of an asymmetric
	// keypair; private - Used for storing the private key of an asymmetric keypair;
	// passphrase - Used for storing plain text passphrases; certificate - Used for
	// storing cryptographic certificates such as X.509 certificates; opaque - Used for
	// backwards compatibility with previous versions of the API
	SecretType param.Field[SecretTypes] `json:"secret_type,required"`
	// Metadata provided by a user or system for informational purposes.
	Algorithm param.Field[string] `json:"algorithm"`
	// Metadata provided by a user or system for informational purposes. Value must be
	// greater than zero.
	BitLength param.Field[int64] `json:"bit_length"`
	// Datetime when the secret will expire.
	Expiration param.Field[string] `json:"expiration"`
	// Metadata provided by a user or system for informational purposes.
	Mode param.Field[string] `json:"mode"`
}

func (r CloudV1SecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The encoding used for the payload to be able to include it in the JSON request.
// Currently only base64 is supported
type CloudV1SecretNewParamsPayloadContentEncoding string

const (
	CloudV1SecretNewParamsPayloadContentEncodingBase64 CloudV1SecretNewParamsPayloadContentEncoding = "base64"
)

func (r CloudV1SecretNewParamsPayloadContentEncoding) IsKnown() bool {
	switch r {
	case CloudV1SecretNewParamsPayloadContentEncodingBase64:
		return true
	}
	return false
}
