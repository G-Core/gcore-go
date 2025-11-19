// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CertificateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificateService] method instead.
type CertificateService struct {
	Options []option.RequestOption
}

// NewCertificateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificateService(opts ...option.RequestOption) (r CertificateService) {
	r = CertificateService{}
	r.Options = opts
	return
}

// Add an SSL certificate for content delivery over HTTPS protocol.
//
// Enter all strings of the certificate(s) and the private key into one string
// parameter. Each certificate and the private key in chain should be separated by
// the "\n" symbol, as shown in the example.
//
// Additionally, you can add a Let's Encrypt certificate. In this case, certificate
// and private key will be generated automatically after attaching this certificate
// to your CDN resource.
func (r *CertificateService) New(ctx context.Context, body CertificateNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cdn/sslData"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get information about SSL certificates.
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *SslDetailList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/sslData"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete SSL certificate
func (r *CertificateService) Delete(ctx context.Context, sslID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/sslData/%v", sslID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Force retry issuance of Let's Encrypt certificate if the previous attempt was
// failed.
func (r *CertificateService) ForceRetry(ctx context.Context, certID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/sslData/%v/force-retry", certID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Get SSL certificate details
func (r *CertificateService) Get(ctx context.Context, sslID int64, opts ...option.RequestOption) (res *SslDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/sslData/%v", sslID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about the latest Let's Encrypt certificate issuing attempt for the
// CDN resource. Returns attempts in all statuses.
func (r *CertificateService) GetStatus(ctx context.Context, certID int64, query CertificateGetStatusParams, opts ...option.RequestOption) (res *SslRequestStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/sslData/%v/status", certID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Renew free Let's Encrypt certificate for the CDN resource. It can take up to
// fifteen minutes.
func (r *CertificateService) Renew(ctx context.Context, certID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/sslData/%v/renew", certID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Change SSL certificate
func (r *CertificateService) Replace(ctx context.Context, sslID int64, body CertificateReplaceParams, opts ...option.RequestOption) (res *SslDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/sslData/%v", sslID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SslDetail struct {
	// SSL certificate ID.
	ID int64 `json:"id"`
	// How the SSL certificate was issued.
	//
	// Possible values:
	//
	// - **true** - Certificate was issued automatically.
	// - **false** - Certificate was added by a use.
	Automated bool `json:"automated"`
	// Name of the certification center issued the SSL certificate.
	CertIssuer string `json:"cert_issuer"`
	// Alternative domain names that the SSL certificate secures.
	CertSubjectAlt string `json:"cert_subject_alt"`
	// Domain name that the SSL certificate secures.
	CertSubjectCn string `json:"cert_subject_cn"`
	// Defines whether the certificate has been deleted. Parameter is **deprecated**.
	//
	// Possible values:
	//
	// - **true** - Certificate has been deleted.
	// - **false** - Certificate has not been deleted.
	Deleted bool `json:"deleted"`
	// Defines whether the SSL certificate is used by a CDN resource.
	//
	// Possible values:
	//
	// - **true** - Certificate is used by a CDN resource.
	// - **false** - Certificate is not used by a CDN resource.
	HasRelatedResources bool `json:"hasRelatedResources"`
	// SSL certificate name.
	Name string `json:"name"`
	// Parameter is **deprecated**.
	SslCertificateChain string `json:"sslCertificateChain"`
	// Date when certificate become untrusted (ISO 8601/RFC 3339 format, UTC.)
	ValidityNotAfter string `json:"validity_not_after"`
	// Date when certificate become valid (ISO 8601/RFC 3339 format, UTC.)
	ValidityNotBefore string `json:"validity_not_before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Automated           respjson.Field
		CertIssuer          respjson.Field
		CertSubjectAlt      respjson.Field
		CertSubjectCn       respjson.Field
		Deleted             respjson.Field
		HasRelatedResources respjson.Field
		Name                respjson.Field
		SslCertificateChain respjson.Field
		ValidityNotAfter    respjson.Field
		ValidityNotBefore   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslDetail) RawJSON() string { return r.JSON.raw }
func (r *SslDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SslDetailList []SslDetail

type SslRequestStatus struct {
	// ID of the attempt to issue a Let's Encrypt certificate.
	ID int64 `json:"id"`
	// Defines whether the Let's Encrypt certificate issuing process is active.
	//
	// Possible values:
	//
	// - **true** - Issuing process is active.
	// - **false** - Issuing process is completed.
	Active bool `json:"active"`
	// Number of attempts to issue the Let's Encrypt certificate.
	AttemptsCount int64 `json:"attempts_count"`
	// Date when the process of issuing a Let's Encrypt certificate was finished (ISO
	// 8601/RFC 3339 format, UTC).
	//
	// The field is **null** if the issuing process is not finished.
	Finished string `json:"finished"`
	// Detailed information about last attempt to issue a Let's Encrypt certificate.
	LatestStatus SslRequestStatusLatestStatus `json:"latest_status"`
	// Time of the next scheduled attempt to issue the Let's Encrypt certificate (ISO
	// 8601/RFC 3339 format, UTC).
	NextAttemptTime string `json:"next_attempt_time,nullable"`
	// CDN resource ID.
	Resource int64 `json:"resource"`
	// Date when the process of issuing a Let's Encrypt certificate was started (ISO
	// 8601/RFC 3339 format, UTC).
	Started string `json:"started"`
	// Detailed information about attempts to issue a Let's Encrypt certificate.
	Statuses []SslRequestStatusStatus `json:"statuses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Active          respjson.Field
		AttemptsCount   respjson.Field
		Finished        respjson.Field
		LatestStatus    respjson.Field
		NextAttemptTime respjson.Field
		Resource        respjson.Field
		Started         respjson.Field
		Statuses        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslRequestStatus) RawJSON() string { return r.JSON.raw }
func (r *SslRequestStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed information about last attempt to issue a Let's Encrypt certificate.
type SslRequestStatusLatestStatus struct {
	// ID of the attempt to issue the Let's Encrypt certificate.
	ID int64 `json:"id"`
	// Date and time when the issuing attempt status was created (ISO 8601/RFC 3339
	// format, UTC).
	Created string `json:"created"`
	// Detailed description of the error that occurred when trying to issue a Let's
	// Encrypt certificate.
	Details string `json:"details"`
	// Brief description of the error that occurred when trying to issue a Let's
	// Encrypt certificate.
	Error string `json:"error"`
	// Date indicating when the certificate issuance limit will be lifted (ISO 8601/RFC
	// 3339 format, UTC).
	//
	// It is filled in only if error = RateLimited.
	RetryAfter string `json:"retry_after"`
	// Status of the attempt to issue the Let's Encrypt certificate.
	//
	// Possible values:
	//
	// - **Done** - Attempt is successful. Let's Encrypt certificate was issued.
	// - **Failed** - Attempt failed. Let's Encrypt certificate was not issued.
	// - **Cancelled** - Attempt is canceled. Let's Encrypt certificate was not issued.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Details     respjson.Field
		Error       respjson.Field
		RetryAfter  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslRequestStatusLatestStatus) RawJSON() string { return r.JSON.raw }
func (r *SslRequestStatusLatestStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SslRequestStatusStatus struct {
	// ID of the attempt to issue the Let's Encrypt certificate.
	ID int64 `json:"id"`
	// Date and time when the issuing attempt status was created (ISO 8601/RFC 3339
	// format, UTC).
	Created string `json:"created"`
	// Detailed description of the error that occurred when trying to issue a Let's
	// Encrypt certificate.
	Details string `json:"details"`
	// Brief description of the error that occurred when trying to issue a Let's
	// Encrypt certificate.
	Error string `json:"error"`
	// Date indicating when the certificate issuance limit will be lifted (ISO 8601/RFC
	// 3339 format, UTC).
	//
	// It is filled in only if error = RateLimited.
	RetryAfter string `json:"retry_after"`
	// Status of the attempt to issue the Let's Encrypt certificate.
	//
	// Possible values:
	//
	// - **Done** - Attempt is successful. Let's Encrypt certificate was issued.
	// - **Failed** - Attempt failed. Let's Encrypt certificate was not issued.
	// - **Cancelled** - Attempt is canceled. Let's Encrypt certificate was not issued.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Details     respjson.Field
		Error       respjson.Field
		RetryAfter  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslRequestStatusStatus) RawJSON() string { return r.JSON.raw }
func (r *SslRequestStatusStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfOwnCertificate *CertificateNewParamsBodyOwnCertificate `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfLetSEncryptCertificate *CertificateNewParamsBodyLetSEncryptCertificate `json:",inline"`

	paramObj
}

func (u CertificateNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOwnCertificate, u.OfLetSEncryptCertificate)
}
func (r *CertificateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, SslCertificate, SslPrivateKey are required.
type CertificateNewParamsBodyOwnCertificate struct {
	// SSL certificate name.
	//
	// It must be unique.
	Name string `json:"name,required"`
	// Public part of the SSL certificate.
	//
	// All chain of the SSL certificate should be added.
	SslCertificate string `json:"sslCertificate,required"`
	// Private key of the SSL certificate.
	SslPrivateKey string `json:"sslPrivateKey,required"`
	// Defines whether to check the SSL certificate for a signature from a trusted
	// certificate authority.
	//
	// Possible values:
	//
	//   - **true** - SSL certificate must be verified to be signed by a trusted
	//     certificate authority.
	//   - **false** - SSL certificate will not be verified to be signed by a trusted
	//     certificate authority.
	ValidateRootCa param.Opt[bool] `json:"validate_root_ca,omitzero"`
	paramObj
}

func (r CertificateNewParamsBodyOwnCertificate) MarshalJSON() (data []byte, err error) {
	type shadow CertificateNewParamsBodyOwnCertificate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CertificateNewParamsBodyOwnCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Automated, Name are required.
type CertificateNewParamsBodyLetSEncryptCertificate struct {
	// Must be **true** to issue certificate automatically.
	Automated bool `json:"automated,required"`
	// SSL certificate name. It must be unique.
	Name string `json:"name,required"`
	paramObj
}

func (r CertificateNewParamsBodyLetSEncryptCertificate) MarshalJSON() (data []byte, err error) {
	type shadow CertificateNewParamsBodyLetSEncryptCertificate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CertificateNewParamsBodyLetSEncryptCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateListParams struct {
	// How the SSL certificate was issued.
	//
	// Possible values:
	//
	// - **true** – Certificate was issued automatically.
	// - **false** – Certificate was added by a user.
	Automated param.Opt[bool] `query:"automated,omitzero" json:"-"`
	// CDN resource ID for which certificates are requested.
	ResourceID param.Opt[int64] `query:"resource_id,omitzero" json:"-"`
	// Date and time when the certificate become untrusted (ISO 8601/RFC 3339 format,
	// UTC.)
	//
	// Response will contain only certificates valid until the specified time.
	ValidityNotAfterLte param.Opt[string] `query:"validity_not_after_lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CertificateListParams]'s query parameters as `url.Values`.
func (r CertificateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CertificateGetStatusParams struct {
	// Listed fields will be excluded from the response.
	Exclude []string `query:"exclude,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CertificateGetStatusParams]'s query parameters as
// `url.Values`.
func (r CertificateGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CertificateReplaceParams struct {
	// SSL certificate name.
	//
	// It must be unique.
	Name string `json:"name,required"`
	// Public part of the SSL certificate.
	//
	// All chain of the SSL certificate should be added.
	SslCertificate string `json:"sslCertificate,required"`
	// Private key of the SSL certificate.
	SslPrivateKey string `json:"sslPrivateKey,required"`
	// Defines whether to check the SSL certificate for a signature from a trusted
	// certificate authority.
	//
	// Possible values:
	//
	//   - **true** - SSL certificate must be verified to be signed by a trusted
	//     certificate authority.
	//   - **false** - SSL certificate will not be verified to be signed by a trusted
	//     certificate authority.
	ValidateRootCa param.Opt[bool] `json:"validate_root_ca,omitzero"`
	paramObj
}

func (r CertificateReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow CertificateReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CertificateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
