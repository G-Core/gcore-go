// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
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

// Trusted CA certificates verify the authenticity of CDN origin servers during
// HTTPS connections.
//
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
	path := "cdn/sslCertificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get list of trusted CA certificates used to verify an origin.
func (r *TrustedCaCertificateService) List(ctx context.Context, query TrustedCaCertificateListParams, opts ...option.RequestOption) (res *CaCertificateListUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/sslCertificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete trusted CA certificate
func (r *TrustedCaCertificateService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get trusted CA certificate details
func (r *TrustedCaCertificateService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *CaCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Change trusted CA certificate
func (r *TrustedCaCertificateService) Replace(ctx context.Context, id int64, body TrustedCaCertificateReplaceParams, opts ...option.RequestOption) (res *CaCertificate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/sslCertificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
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

// CaCertificateListUnion contains all possible properties and values from
// [[]CaCertificate], [CaCertificateListPaginatedList].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPlainList]
type CaCertificateListUnion struct {
	// This field will be present if the value is a [[]CaCertificate] instead of an
	// object.
	OfPlainList []CaCertificate `json:",inline"`
	// This field is from variant [CaCertificateListPaginatedList].
	Count int64 `json:"count"`
	// This field is from variant [CaCertificateListPaginatedList].
	Next string `json:"next"`
	// This field is from variant [CaCertificateListPaginatedList].
	Previous string `json:"previous"`
	// This field is from variant [CaCertificateListPaginatedList].
	Results []CaCertificate `json:"results"`
	JSON    struct {
		OfPlainList respjson.Field
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		raw         string
	} `json:"-"`
}

func (u CaCertificateListUnion) AsPlainList() (v []CaCertificate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CaCertificateListUnion) AsPaginatedList() (v CaCertificateListPaginatedList) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CaCertificateListUnion) RawJSON() string { return u.JSON.raw }

func (r *CaCertificateListUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaCertificateListPaginatedList struct {
	// Total number of items.
	Count int64 `json:"count" api:"required"`
	// URL to the next page of results. Null if current page is the last one.
	Next string `json:"next" api:"required"`
	// URL to the previous page of results. Null if current page is the first one.
	Previous string          `json:"previous" api:"required"`
	Results  []CaCertificate `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaCertificateListPaginatedList) RawJSON() string { return r.JSON.raw }
func (r *CaCertificateListPaginatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrustedCaCertificateNewParams struct {
	// CA certificate name.
	//
	// It must be unique.
	Name string `json:"name" api:"required"`
	// Public part of the CA certificate.
	//
	// It must be in the PEM format.
	SslCertificate string `json:"sslCertificate" api:"required"`
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
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
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
	Name string `json:"name" api:"required"`
	paramObj
}

func (r TrustedCaCertificateReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow TrustedCaCertificateReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrustedCaCertificateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
