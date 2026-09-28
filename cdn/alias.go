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
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CDN aliases are hostnames you own that are served with the settings of one of
// your CDN resources, each with its own SSL certificate.
//
// AliasService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAliasService] method instead.
type AliasService struct {
	Options []option.RequestOption
}

// NewAliasService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAliasService(opts ...option.RequestOption) (r AliasService) {
	r = AliasService{}
	r.Options = opts
	return
}

// Create an alias: a hostname you own that is served with the settings of one of
// your CDN resources and with its own SSL certificate.
//
// By default, a Let's Encrypt certificate is issued for the alias automatically.
// Until the certificate is issued the alias stays in the `pending` status.
//
// To use your own certificate instead, pass its ID in `ssl_id`. The certificate
// must cover the alias hostname; a Let's Encrypt certificate issued for one of
// your CDN resources cannot be used. Such an alias becomes `active` immediately.
//
// The hostname must be unique across all CDN resources, additional CNAMEs and
// aliases.
//
// If Let's Encrypt validation for the hostname requires a delegation record, it is
// reported as `ssl_provisioning` when you get the alias.
func (r *AliasService) New(ctx context.Context, body AliasNewParams, opts ...option.RequestOption) (res *Alias, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/aliases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change an alias. The hostname and the CDN resource of an alias cannot be
// changed; delete the alias and create a new one instead.
//
// Set `active` to **false** to stop serving the alias; set it back to **true** to
// resume. Resuming an alias that has no valid certificate starts a new issuance
// attempt, which is rate-limited, so it cannot be combined with `automated` in one
// request.
//
// For an alias with your own certificate, pass a new certificate ID in `ssl_id` to
// replace it. The new certificate must cover the alias hostname.
//
// Pass `automated` to switch how the certificate is managed. Switching to your own
// certificate takes effect at once and the Let's Encrypt certificate issued for
// the alias is removed. Switching to a Let's Encrypt certificate starts an
// issuance attempt while your certificate keeps being served, and the alias starts
// serving the new certificate once it is issued; if the attempt fails, `status`
// becomes **`ssl_error`** and your certificate keeps being served. The certificate
// mode cannot be changed while the alias is paused — resume it first.
func (r *AliasService) Update(ctx context.Context, aliasID int64, body AliasUpdateParams, opts ...option.RequestOption) (res *AliasDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/aliases/%v", aliasID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get information about aliases.
//
// The response is always paginated.
func (r *AliasService) List(ctx context.Context, query AliasListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Alias], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/aliases"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get information about aliases.
//
// The response is always paginated.
func (r *AliasService) ListAutoPaging(ctx context.Context, query AliasListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Alias] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an alias. The hostname stops being served and can be used again. A
// certificate you added yourself is not affected.
func (r *AliasService) Delete(ctx context.Context, aliasID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/aliases/%v", aliasID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create up to 1000 aliases for one CDN resource in a single request.
//
// The request is atomic: if any item is invalid, nothing is created and the errors
// are returned per item, in the same order as the request items. An item whose
// hostname is already an alias of the same CDN resource is reported in `exists`
// and left unchanged.
//
// The aliases limit of your account is checked against the items that would be
// created.
func (r *AliasService) NewMultiple(ctx context.Context, body AliasNewMultipleParams, opts ...option.RequestOption) (res *AliasNewMultipleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/aliases/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete up to 1000 aliases in a single request, selected either by ID or by
// hostname.
//
// Aliases that are not found are reported in `not_found` and do not fail the
// request.
func (r *AliasService) DeleteMultiple(ctx context.Context, body AliasDeleteMultipleParams, opts ...option.RequestOption) (res *AliasDeleteMultipleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/aliases/remove"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get alias details
func (r *AliasService) Get(ctx context.Context, aliasID int64, opts ...option.RequestOption) (res *AliasDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/aliases/%v", aliasID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get details about the latest Let's Encrypt certificate issuing attempt for an
// alias, including the CNAME records to create for domain validation. Returns
// attempts in all statuses.
//
// While the alias is switching to a Let's Encrypt certificate, the attempt for the
// new certificate is returned.
func (r *AliasService) GetCertificateStatus(ctx context.Context, aliasID int64, opts ...option.RequestOption) (res *AliasCertificateStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/aliases/%v/status", aliasID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Start a new Let's Encrypt certificate issuance attempt for an alias that has
// none running. Use it after an attempt has failed and `status` is
// **`ssl_error`**.
//
// The number of attempts you can start per hour is limited.
func (r *AliasService) RetryCertificate(ctx context.Context, aliasID int64, opts ...option.RequestOption) (res *AliasDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/aliases/%v/retry", aliasID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Alias struct {
	// Alias ID.
	ID int64 `json:"id"`
	// Whether you have enabled the alias. Enabling it does not by itself make it live:
	// the alias is served only while `status` is **active**, **`ssl_issuing`** or
	// **`ssl_error`**.
	//
	// Possible values:
	//
	// - **true** – The alias is enabled and is served once its certificate is ready.
	// - **false** – The alias is paused and is not served.
	Active bool `json:"active"`
	// How the alias certificate is managed. Can be changed while the alias is served.
	//
	// Possible values:
	//
	// - **true** – A Let's Encrypt certificate is issued and renewed automatically.
	// - **false** – The certificate referenced by `ssl_id` was added by you.
	Automated bool `json:"automated"`
	// Alias hostname. Cannot be changed after creation.
	Cname string `json:"cname"`
	// Date and time when the alias was created (ISO 8601/RFC 3339 format, UTC.)
	Created string `json:"created"`
	// Whether the alias is enabled by the system. Follows the state of the CDN
	// resource.
	//
	// Possible values:
	//
	// - **true** – The alias can be served.
	// - **false** – The alias is not served because its CDN resource is not active.
	Enabled bool `json:"enabled"`
	// ID of the CDN resource whose settings the alias is served with. Cannot be
	// changed after creation.
	ResourceID int64 `json:"resource_id"`
	// ID of the SSL certificate served for the alias. Null until a certificate is
	// issued.
	SslID int64 `json:"ssl_id" api:"nullable"`
	// Outcome of the certificate we manage for the alias.
	//
	// Possible values:
	//
	//   - **pending** – A certificate attempt is in progress.
	//   - **issued** – The latest attempt succeeded.
	//   - **failed** – The latest attempt ended without a certificate.
	//   - **null** – No certificate is managed for the alias: the alias uses a
	//     certificate you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus AliasSslStatus `json:"ssl_status" api:"nullable"`
	// Date and time when the served certificate expires (ISO 8601/RFC 3339 format,
	// UTC). Null while no certificate is served yet.
	SslValidityNotAfter string `json:"ssl_validity_not_after" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - **pending** – The certificate has not been issued yet; the alias is not
	//     served.
	//   - **active** – The alias is served with its certificate.
	//   - **`ssl_issuing`** – A new certificate is being issued; the current one keeps
	//     being served.
	//   - **`ssl_error`** – Certificate issuance failed; a previously issued certificate
	//     keeps being served.
	//   - **inactive** – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//
	// Any of "pending", "active", "ssl_issuing", "ssl_error", "inactive".
	Status AliasStatus `json:"status"`
	// Date and time when the alias was last changed (ISO 8601/RFC 3339 format, UTC.)
	Updated string `json:"updated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Active              respjson.Field
		Automated           respjson.Field
		Cname               respjson.Field
		Created             respjson.Field
		Enabled             respjson.Field
		ResourceID          respjson.Field
		SslID               respjson.Field
		SslStatus           respjson.Field
		SslValidityNotAfter respjson.Field
		Status              respjson.Field
		Updated             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Alias) RawJSON() string { return r.JSON.raw }
func (r *Alias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outcome of the certificate we manage for the alias.
//
// Possible values:
//
//   - **pending** – A certificate attempt is in progress.
//   - **issued** – The latest attempt succeeded.
//   - **failed** – The latest attempt ended without a certificate.
//   - **null** – No certificate is managed for the alias: the alias uses a
//     certificate you provided, or is automated and has never opened an attempt.
type AliasSslStatus string

const (
	AliasSslStatusPending AliasSslStatus = "pending"
	AliasSslStatusIssued  AliasSslStatus = "issued"
	AliasSslStatusFailed  AliasSslStatus = "failed"
)

// Alias status.
//
// Possible values:
//
//   - **pending** – The certificate has not been issued yet; the alias is not
//     served.
//   - **active** – The alias is served with its certificate.
//   - **`ssl_issuing`** – A new certificate is being issued; the current one keeps
//     being served.
//   - **`ssl_error`** – Certificate issuance failed; a previously issued certificate
//     keeps being served.
//   - **inactive** – The alias or its CDN resource is disabled; the alias is not
//     served.
type AliasStatus string

const (
	AliasStatusPending    AliasStatus = "pending"
	AliasStatusActive     AliasStatus = "active"
	AliasStatusSslIssuing AliasStatus = "ssl_issuing"
	AliasStatusSslError   AliasStatus = "ssl_error"
	AliasStatusInactive   AliasStatus = "inactive"
)

type AliasCertificateStatus struct {
	// The domain validation method this attempt uses.
	//
	// Possible values:
	//
	// - **`dns_01`** - Validation via a DNS record.
	// - **`http_01`** - Validation via an HTTP request.
	ChallengeType string `json:"challenge_type" api:"nullable"`
	// CNAME records to create so Let's Encrypt domain validation can complete for the
	// alias hostname.
	CnameRecords []AliasCertificateStatusCnameRecord `json:"cname_records"`
	// Why this attempt was opened.
	//
	// Possible values:
	//
	//   - **initial** - First attempt to issue a certificate.
	//   - **renewal** - A scheduled renewal of an existing certificate.
	//   - **manual** - Manually triggered, for example via a retry.
	//   - **switch** - Switching from your own certificate to a Let's Encrypt
	//     certificate.
	//   - **extend** - Extending the set of domains covered by the certificate.
	Reason string `json:"reason" api:"nullable"`
	// Domains this attempt requested a certificate for.
	RequestedDomains []string `json:"requested_domains" api:"nullable"`
	// State of the domain validation for this attempt.
	//
	// Possible values:
	//
	// - **PENDING** - Validation has not started yet.
	// - **`AWAITING_CNAME`** - Waiting for the CNAME records below to be created.
	// - **VALIDATING** - Validation is in progress.
	// - **COMPLETED** - Validation succeeded.
	// - **FAILED** - Validation did not succeed.
	ValidationState string `json:"validation_state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChallengeType    respjson.Field
		CnameRecords     respjson.Field
		Reason           respjson.Field
		RequestedDomains respjson.Field
		ValidationState  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
	SslRequestStatus
}

// Returns the unmodified JSON received from the API
func (r AliasCertificateStatus) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasCertificateStatusCnameRecord struct {
	// Hostname to create the CNAME record for.
	CnameName string `json:"cname_name"`
	// Value the CNAME record must point to.
	CnameTarget string `json:"cname_target"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameName   respjson.Field
		CnameTarget respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasCertificateStatusCnameRecord) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatusCnameRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasDetail struct {
	// DNS record delegating Let's Encrypt domain validation for the alias hostname. It
	// is `null` for an alias serving a certificate you added yourself, and until a
	// delegation record has been prepared for the hostname.
	SslProvisioning AliasDetailSslProvisioning `json:"ssl_provisioning" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SslProvisioning respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Alias
}

// Returns the unmodified JSON received from the API
func (r AliasDetail) RawJSON() string { return r.JSON.raw }
func (r *AliasDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS record delegating Let's Encrypt domain validation for the alias hostname. It
// is `null` for an alias serving a certificate you added yourself, and until a
// delegation record has been prepared for the hostname.
type AliasDetailSslProvisioning struct {
	// CNAME record to create so validation can be completed for the hostname.
	AcmeDelegation AliasDetailSslProvisioningAcmeDelegation `json:"acme_delegation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcmeDelegation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasDetailSslProvisioning) RawJSON() string { return r.JSON.raw }
func (r *AliasDetailSslProvisioning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CNAME record to create so validation can be completed for the hostname.
type AliasDetailSslProvisioningAcmeDelegation struct {
	// Hostname to create the CNAME record for.
	CnameName string `json:"cname_name"`
	// Value the CNAME record must point to.
	CnameTarget string `json:"cname_target"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameName   respjson.Field
		CnameTarget respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasDetailSslProvisioningAcmeDelegation) RawJSON() string { return r.JSON.raw }
func (r *AliasDetailSslProvisioningAcmeDelegation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasNewMultipleResponse struct {
	// Aliases created by this request.
	Created []AliasNewMultipleResponseCreated `json:"created"`
	// Items whose hostname was already an alias of the CDN resource. Nothing was
	// changed for them.
	Exists []AliasNewMultipleResponseExist `json:"exists"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Exists      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasNewMultipleResponse) RawJSON() string { return r.JSON.raw }
func (r *AliasNewMultipleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasNewMultipleResponseCreated struct {
	// Alias ID.
	ID int64 `json:"id"`
	// Alias hostname.
	Cname string `json:"cname"`
	// Outcome of the certificate we manage for the alias.
	//
	// Possible values:
	//
	//   - **pending** – A certificate attempt is in progress.
	//   - **issued** – The latest attempt succeeded.
	//   - **failed** – The latest attempt ended without a certificate.
	//   - **null** – No certificate is managed for the alias: the alias uses a
	//     certificate you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus string `json:"ssl_status" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - **pending** – The certificate has not been issued yet; the alias is not
	//     served.
	//   - **active** – The alias is served with its certificate.
	//   - **`ssl_issuing`** – A new certificate is being issued; the current one keeps
	//     being served.
	//   - **`ssl_error`** – Certificate issuance failed; a previously issued certificate
	//     keeps being served.
	//   - **inactive** – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//
	// Any of "pending", "active", "ssl_issuing", "ssl_error", "inactive".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cname       respjson.Field
		SslStatus   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasNewMultipleResponseCreated) RawJSON() string { return r.JSON.raw }
func (r *AliasNewMultipleResponseCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasNewMultipleResponseExist struct {
	// Alias ID.
	ID int64 `json:"id"`
	// Alias hostname.
	Cname string `json:"cname"`
	// Outcome of the certificate we manage for the alias.
	//
	// Possible values:
	//
	//   - **pending** – A certificate attempt is in progress.
	//   - **issued** – The latest attempt succeeded.
	//   - **failed** – The latest attempt ended without a certificate.
	//   - **null** – No certificate is managed for the alias: the alias uses a
	//     certificate you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus string `json:"ssl_status" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - **pending** – The certificate has not been issued yet; the alias is not
	//     served.
	//   - **active** – The alias is served with its certificate.
	//   - **`ssl_issuing`** – A new certificate is being issued; the current one keeps
	//     being served.
	//   - **`ssl_error`** – Certificate issuance failed; a previously issued certificate
	//     keeps being served.
	//   - **inactive** – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//
	// Any of "pending", "active", "ssl_issuing", "ssl_error", "inactive".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cname       respjson.Field
		SslStatus   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasNewMultipleResponseExist) RawJSON() string { return r.JSON.raw }
func (r *AliasNewMultipleResponseExist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasDeleteMultipleResponse struct {
	// Requested IDs or hostnames that did not match an alias.
	NotFound []AliasDeleteMultipleResponseNotFoundUnion `json:"not_found"`
	// Aliases deleted by this request.
	Removed []AliasDeleteMultipleResponseRemoved `json:"removed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NotFound    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasDeleteMultipleResponse) RawJSON() string { return r.JSON.raw }
func (r *AliasDeleteMultipleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AliasDeleteMultipleResponseNotFoundUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AliasDeleteMultipleResponseNotFoundUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AliasDeleteMultipleResponseNotFoundUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AliasDeleteMultipleResponseNotFoundUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AliasDeleteMultipleResponseNotFoundUnion) RawJSON() string { return u.JSON.raw }

func (r *AliasDeleteMultipleResponseNotFoundUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasDeleteMultipleResponseRemoved struct {
	ID    int64  `json:"id"`
	Cname string `json:"cname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cname       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasDeleteMultipleResponseRemoved) RawJSON() string { return r.JSON.raw }
func (r *AliasDeleteMultipleResponseRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasNewParams struct {
	// Alias hostname. Wildcard hostnames are not supported.
	Cname string `json:"cname" api:"required"`
	// ID of the CDN resource whose settings the alias is served with.
	ResourceID int64 `json:"resource_id" api:"required"`
	// ID of your own SSL certificate to serve for the alias. The certificate must
	// cover the alias hostname. A Let's Encrypt certificate issued for one of your CDN
	// resources cannot be used.
	//
	// Omit it to have a Let's Encrypt certificate issued automatically.
	SslID param.Opt[int64] `json:"ssl_id,omitzero"`
	// Whether the alias is enabled. Defaults to **true**. The alias is served once its
	// certificate is ready.
	Active param.Opt[bool] `json:"active,omitzero"`
	// How the alias certificate is managed. Defaults to **true** when `ssl_id` is
	// omitted and to **false** when `ssl_id` is passed.
	//
	// Possible values:
	//
	//   - **true** – A Let's Encrypt certificate is issued and renewed automatically.
	//     `ssl_id` must be omitted.
	//   - **false** – Your own certificate is served. `ssl_id` is required.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	paramObj
}

func (r AliasNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasUpdateParams struct {
	// Whether the alias is enabled.
	//
	// Possible values:
	//
	// - **true** – The alias is enabled and is served once its certificate is ready.
	// - **false** – The alias is paused and is not served.
	Active param.Opt[bool] `json:"active,omitzero"`
	// How the alias certificate is managed. Change it to switch the alias between a
	// Let's Encrypt certificate and your own. Cannot be changed while the alias is
	// paused.
	//
	// Possible values:
	//
	//   - **true** – Switch to a Let's Encrypt certificate. `ssl_id` must be omitted.
	//     The current certificate keeps being served until the new one is issued.
	//   - **false** – Switch to your own certificate. `ssl_id` is required and is served
	//     immediately.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	// ID of your own SSL certificate to serve for the alias instead of the current
	// one. The certificate must cover the alias hostname; a Let's Encrypt certificate
	// issued for one of your CDN resources cannot be used. Available for an alias with
	// `automated` set to **false**, and together with `automated` set to **false** to
	// switch to your own certificate.
	SslID param.Opt[int64] `json:"ssl_id,omitzero"`
	paramObj
}

func (r AliasUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasListParams struct {
	// How the alias certificate is managed.
	//
	// Possible values:
	//
	// - **true** – Certificate is issued and renewed automatically.
	// - **false** – Certificate was added by a user.
	Automated param.Opt[bool] `query:"automated,omitzero" json:"-"`
	// Hostname substring, case-insensitive.
	Cname param.Opt[string] `query:"cname,omitzero" json:"-"`
	// Hostname suffix, case-insensitive.
	CnameEndswith param.Opt[string] `query:"cname_endswith,omitzero" json:"-"`
	// Hostname prefix, case-insensitive.
	CnameStartswith param.Opt[string] `query:"cname_startswith,omitzero" json:"-"`
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Field to sort by. Prefix with `-` for descending order.
	//
	// Possible values: `id`, `cname`, `status`, `created`, `updated`, `resource_id`.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// CDN resource ID. Only aliases of this CDN resource are returned.
	ResourceID param.Opt[int64] `query:"resource_id,omitzero" json:"-"`
	// Comma-separated list of CDN resource IDs.
	ResourceIDIn param.Opt[string] `query:"resource_id__in,omitzero" json:"-"`
	// Search by alias ID or hostname.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// SSL certificate ID. Only aliases linked to this certificate are returned.
	SslID param.Opt[int64] `query:"ssl_id,omitzero" json:"-"`
	// Comma-separated list of certificate outcomes.
	SslStatusIn param.Opt[string] `query:"ssl_status__in,omitzero" json:"-"`
	// Only aliases whose served certificate expires at or after this date and time
	// (ISO 8601/RFC 3339 format, UTC).
	SslValidityNotAfterGte param.Opt[string] `query:"ssl_validity_not_after_gte,omitzero" json:"-"`
	// Only aliases whose served certificate expires at or before this date and time
	// (ISO 8601/RFC 3339 format, UTC).
	SslValidityNotAfterLte param.Opt[string] `query:"ssl_validity_not_after_lte,omitzero" json:"-"`
	// Comma-separated list of alias statuses.
	StatusIn param.Opt[string] `query:"status__in,omitzero" json:"-"`
	// Certificate outcome.
	//
	// Any of "pending", "issued", "failed".
	SslStatus AliasListParamsSslStatus `query:"ssl_status,omitzero" json:"-"`
	// Alias status.
	//
	// Any of "pending", "active", "ssl_issuing", "ssl_error", "inactive".
	Status AliasListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AliasListParams]'s query parameters as `url.Values`.
func (r AliasListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Certificate outcome.
type AliasListParamsSslStatus string

const (
	AliasListParamsSslStatusPending AliasListParamsSslStatus = "pending"
	AliasListParamsSslStatusIssued  AliasListParamsSslStatus = "issued"
	AliasListParamsSslStatusFailed  AliasListParamsSslStatus = "failed"
)

// Alias status.
type AliasListParamsStatus string

const (
	AliasListParamsStatusPending    AliasListParamsStatus = "pending"
	AliasListParamsStatusActive     AliasListParamsStatus = "active"
	AliasListParamsStatusSslIssuing AliasListParamsStatus = "ssl_issuing"
	AliasListParamsStatusSslError   AliasListParamsStatus = "ssl_error"
	AliasListParamsStatusInactive   AliasListParamsStatus = "inactive"
)

type AliasNewMultipleParams struct {
	// Aliases to create.
	Items []AliasNewMultipleParamsItem `json:"items,omitzero" api:"required"`
	// ID of the CDN resource whose settings the aliases are served with.
	ResourceID int64 `json:"resource_id" api:"required"`
	paramObj
}

func (r AliasNewMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasNewMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasNewMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Cname is required.
type AliasNewMultipleParamsItem struct {
	// Alias hostname.
	Cname string `json:"cname" api:"required"`
	// ID of your own SSL certificate to serve for the alias. A Let's Encrypt
	// certificate issued for one of your CDN resources cannot be used. Omit it to have
	// a Let's Encrypt certificate issued automatically.
	SslID param.Opt[int64] `json:"ssl_id,omitzero"`
	// How the alias certificate is managed. Defaults to **true** when `ssl_id` is
	// omitted and to **false** when `ssl_id` is passed.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	paramObj
}

func (r AliasNewMultipleParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow AliasNewMultipleParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasNewMultipleParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasDeleteMultipleParams struct {
	// Hostnames of the aliases to delete.
	Cnames []string `json:"cnames,omitzero"`
	// IDs of the aliases to delete.
	IDs []int64 `json:"ids,omitzero"`
	paramObj
}

func (r AliasDeleteMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasDeleteMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasDeleteMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
