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

// TrustedCaCertificateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrustedCaCertificateService] method instead.
type TrustedCaCertificateService struct {
	Options []option.RequestOption
}

// NewTrustedCaCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrustedCaCertificateService(opts ...option.RequestOption) (r TrustedCaCertificateService) {
	r = TrustedCaCertificateService{}
	r.Options = opts
	return
}

// Add a trusted CA certificate to verify an origin.
//
// Enter all strings of the certificate in one string parameter. Each string should
// be separated by the "\n" symbol.
func (r *TrustedCaCertificateService) New(ctx context.Context, body TrustedCaCertificateNewParams, opts ...option.RequestOption) (res *CaCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/sslCertificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get list of trusted CA certificates used to verify an origin.
func (r *TrustedCaCertificateService) List(ctx context.Context, query TrustedCaCertificateListParams, opts ...option.RequestOption) (res *CaCertificateList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/sslCertificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete trusted CA certificate
func (r *TrustedCaCertificateService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get trusted CA certificate details
func (r *TrustedCaCertificateService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *CaCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change trusted CA certificate
func (r *TrustedCaCertificateService) Replace(ctx context.Context, id int64, body TrustedCaCertificateReplaceParams, opts ...option.RequestOption) (res *CaCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CaCertificate struct {
	// CA certificate ID.
	ID int64 `json:"id"`
	// Name of the certification center that issued the CA certificate.
	CertIssuer string `json:"cert_issuer"`
	// Alternative domain names that the CA certificate secures.
	CertSubjectAlt string `json:"cert_subject_alt"`
	// Domain name that the CA certificate secures.
	CertSubjectCn string `json:"cert_subject_cn"`
	// Defines whether the certificate has been deleted. Parameter is **deprecated**.
	//
	// Possible values:
	//
	// - **true** - Certificate has been deleted.
	// - **false** - Certificate has not been deleted.
	Deleted bool `json:"deleted"`
	// Defines whether the CA certificate is used by a CDN resource.
	//
	// Possible values:
	//
	// - **true** - Certificate is used by a CDN resource.
	// - **false** - Certificate is not used by a CDN resource.
	HasRelatedResources bool `json:"hasRelatedResources"`
	// CA certificate name.
	Name string `json:"name"`
	// Parameter is **deprecated**.
	SslCertificateChain string `json:"sslCertificateChain"`
	// Date when the CA certificate become untrusted (ISO 8601/RFC 3339 format, UTC.)
	ValidityNotAfter string `json:"validity_not_after"`
	// Date when the CA certificate become valid (ISO 8601/RFC 3339 format, UTC.)
	ValidityNotBefore string `json:"validity_not_before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
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
func (r CaCertificate) RawJSON() string { return r.JSON.raw }
func (r *CaCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaCertificateList []CaCertificate

type TrustedCaCertificateNewParams struct {
	// CA certificate name.
	//
	// It must be unique.
	Name string `json:"name,required"`
	// Public part of the CA certificate.
	//
	// It must be in the PEM format.
	SslCertificate string `json:"sslCertificate,required"`
	paramObj
}

func (r TrustedCaCertificateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TrustedCaCertificateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrustedCaCertificateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrustedCaCertificateListParams struct {
	// How the certificate was issued.
	//
	// Possible values:
	//
	// - **true** – Certificate was issued automatically.
	// - **false** – Certificate was added by a user.
	Automated param.Opt[bool] `query:"automated,omitzero" json:"-"`
	// CDN resource ID for which the certificates are requested.
	ResourceID param.Opt[int64] `query:"resource_id,omitzero" json:"-"`
	// Date and time when the certificate become untrusted (ISO 8601/RFC 3339 format,
	// UTC.)
	//
	// Response will contain certificates valid until the specified time.
	ValidityNotAfterLte param.Opt[string] `query:"validity_not_after_lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrustedCaCertificateListParams]'s query parameters as
// `url.Values`.
func (r TrustedCaCertificateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TrustedCaCertificateReplaceParams struct {
	// CA certificate name.
	//
	// It must be unique.
	Name string `json:"name,required"`
	paramObj
}

func (r TrustedCaCertificateReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow TrustedCaCertificateReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrustedCaCertificateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
