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
// Choose how its domain is validated with `challenge_type`.
//
// To use your own certificate instead, pass its ID in `ssl_id`. The certificate
// must cover the alias hostname; a Let's Encrypt certificate issued for one of
// your CDN resources cannot be used.
//
// The hostname must be unique across all CDN resources, additional CNAMEs and
// aliases.
//
// The DNS record the hostname needs for Let's Encrypt domain validation is
// returned in `ssl_provisioning` of the alias details.
func (r *AliasService) New(ctx context.Context, body AliasNewParams, opts ...option.RequestOption) (res *Alias, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/aliases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change an alias. The hostname and the CDN resource of an alias cannot be
// changed; delete the alias and create a new one instead.
//
// Set `active` to `false` to stop serving the alias; certificate issuance stops,
// though an attempt already being processed completes. Set it back to `true` to
// resume. Resuming an alias without a valid Let's Encrypt certificate starts a new
// issuance attempt, which is rate-limited, so it cannot be combined with
// `automated` in one request. A renewal stopped by the pause is started again by
// the scheduled renewal after the alias is resumed.
//
// For an alias with your own certificate, pass a new certificate ID in `ssl_id` to
// replace it. The new certificate must cover the alias hostname.
//
// Pass `automated` to switch how the certificate is managed. Switching to your own
// certificate takes effect at once and the Let's Encrypt certificate issued for
// the alias is removed. Switching to a Let's Encrypt certificate starts an
// issuance attempt while your certificate keeps being served, and the alias starts
// serving the new certificate once it is issued; if the attempt fails,
// `ssl_status` becomes `failed` and your certificate keeps being served. The
// certificate mode cannot be changed while the alias is paused — resume it first.
//
// Pass `challenge_type` to change how the Let's Encrypt certificate's domain is
// validated. It cannot be changed while the alias is paused, or together with
// `active`. While no valid certificate is served, the change cancels the running
// issuance attempt and starts a new one; this is rate-limited like resuming, and
// is refused while that attempt is being processed. Otherwise the new method is
// used from the next attempt.
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

// Get details about the latest Let's Encrypt certificate issuance attempt for an
// alias, including the CNAME records to create for domain validation and the state
// of each domain validation challenge. Returns attempts in all statuses.
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
// none running. Use it after an attempt has failed and `ssl_status` is `failed`.
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
	// Whether you have enabled the alias.
	//
	// Possible values:
	//
	// - `true` – The alias is enabled.
	// - `false` – The alias is paused and is not served.
	Active bool `json:"active"`
	// How the alias certificate is managed. Can be changed while the alias is served.
	//
	// Possible values:
	//
	// - `true` – A Let's Encrypt certificate is issued and renewed automatically.
	// - `false` – The certificate referenced by `ssl_id`, if any, was added by you.
	Automated bool `json:"automated"`
	// Domain validation method used to issue the Let's Encrypt certificate for the
	// alias.
	//
	// Possible values:
	//
	//   - `dns_01` – Validation via a CNAME record for
	//     `_acme-challenge.<alias hostname>`, shown in
	//     `ssl_provisioning.acme_delegation`.
	//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
	//     hostname must point to your CDN CNAME, shown in
	//     `ssl_provisioning.http_validation`. The request is served with the settings of
	//     the alias's CDN resource, so they must not block or redirect
	//     `/.well-known/acme-challenge/`.
	//
	// Any of "dns_01", "http_01".
	ChallengeType AliasChallengeType `json:"challenge_type" api:"nullable"`
	// Alias hostname. Cannot be changed after creation.
	Cname string `json:"cname"`
	// Date and time when the alias was created (ISO 8601/RFC 3339 format, UTC).
	Created string `json:"created"`
	// Whether the alias is enabled by the system. Follows the state of the CDN
	// resource.
	//
	// Possible values:
	//
	// - `true` – The alias can be served.
	// - `false` – The alias is not served because its CDN resource is not active.
	Enabled bool `json:"enabled"`
	// ID of the CDN resource whose settings the alias is served with. Cannot be
	// changed after creation.
	ResourceID int64 `json:"resource_id"`
	// ID of the SSL certificate served for the alias.
	SslID int64 `json:"ssl_id" api:"nullable"`
	// Outcome of the certificate we manage for the alias.
	//
	// Possible values:
	//
	//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
	//     the alias and starts again after it is resumed.
	//   - `issued` – The latest attempt succeeded.
	//   - `failed` – The latest attempt ended without a certificate.
	//   - `null` – No certificate is managed for the alias: the alias uses a certificate
	//     you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus AliasSslStatus `json:"ssl_status" api:"nullable"`
	// Date and time when the served certificate expires (ISO 8601/RFC 3339 format,
	// UTC). `null` while no certificate is served.
	SslValidityNotAfter string `json:"ssl_validity_not_after" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//   - `pending` – The alias configuration has not reached the CDN yet.
	//   - `http_only` – The alias is served over HTTP only.
	//   - `active` – The alias is served over HTTP and HTTPS.
	//
	// Any of "inactive", "pending", "http_only", "active".
	Status AliasStatus `json:"status"`
	// Date and time when the alias was last changed (ISO 8601/RFC 3339 format, UTC).
	Updated string `json:"updated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Active              respjson.Field
		Automated           respjson.Field
		ChallengeType       respjson.Field
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

// Domain validation method used to issue the Let's Encrypt certificate for the
// alias.
//
// Possible values:
//
//   - `dns_01` – Validation via a CNAME record for
//     `_acme-challenge.<alias hostname>`, shown in
//     `ssl_provisioning.acme_delegation`.
//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
//     hostname must point to your CDN CNAME, shown in
//     `ssl_provisioning.http_validation`. The request is served with the settings of
//     the alias's CDN resource, so they must not block or redirect
//     `/.well-known/acme-challenge/`.
type AliasChallengeType string

const (
	AliasChallengeTypeDNS01  AliasChallengeType = "dns_01"
	AliasChallengeTypeHTTP01 AliasChallengeType = "http_01"
)

// Outcome of the certificate we manage for the alias.
//
// Possible values:
//
//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
//     the alias and starts again after it is resumed.
//   - `issued` – The latest attempt succeeded.
//   - `failed` – The latest attempt ended without a certificate.
//   - `null` – No certificate is managed for the alias: the alias uses a certificate
//     you provided, or is automated and has never opened an attempt.
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
//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
//     served.
//   - `pending` – The alias configuration has not reached the CDN yet.
//   - `http_only` – The alias is served over HTTP only.
//   - `active` – The alias is served over HTTP and HTTPS.
type AliasStatus string

const (
	AliasStatusInactive AliasStatus = "inactive"
	AliasStatusPending  AliasStatus = "pending"
	AliasStatusHTTPOnly AliasStatus = "http_only"
	AliasStatusActive   AliasStatus = "active"
)

type AliasCertificateStatus struct {
	// ID of the attempt to issue a Let's Encrypt certificate.
	ID int64 `json:"id"`
	// Defines whether the Let's Encrypt certificate issuing process is active.
	//
	// Possible values:
	//
	// - `true` – Issuing process is active.
	// - `false` – Issuing process is completed.
	Active bool `json:"active"`
	// Number of attempts to issue the Let's Encrypt certificate.
	AttemptsCount int64 `json:"attempts_count"`
	// The domain validation method this attempt uses.
	//
	// Possible values:
	//
	// - `dns_01` – Validation via a DNS record.
	// - `http_01` – Validation via an HTTP request.
	ChallengeType string `json:"challenge_type" api:"nullable"`
	// Domain validation challenges of the latest order placed with Let's Encrypt, one
	// per domain.
	Challenges []AliasCertificateStatusChallenge `json:"challenges"`
	// CNAME records to create so Let's Encrypt domain validation can complete for the
	// alias hostname. Empty when the attempt uses `http_01`.
	CnameRecords []AliasCertificateStatusCnameRecord `json:"cname_records"`
	// Date when the process of issuing a Let's Encrypt certificate was finished (ISO
	// 8601/RFC 3339 format, UTC).
	//
	// The field is `null` if the issuing process is not finished.
	Finished string `json:"finished" api:"nullable"`
	// Detailed information about last attempt to issue a Let's Encrypt certificate.
	//
	// The field is `null` until the first attempt of the issuing process has been
	// recorded.
	LatestStatus AliasCertificateStatusLatestStatus `json:"latest_status" api:"nullable"`
	// Time of the next scheduled attempt to issue the Let's Encrypt certificate (ISO
	// 8601/RFC 3339 format, UTC).
	NextAttemptTime string `json:"next_attempt_time" api:"nullable"`
	// Why this attempt was opened.
	//
	// Possible values:
	//
	// - `initial` – First attempt to issue a certificate.
	// - `renewal` – A scheduled renewal of an existing certificate.
	// - `manual` – Manually triggered, for example via a retry.
	// - `switch` – Switching from your own certificate to a Let's Encrypt certificate.
	// - `extend` – Extending the set of domains covered by the certificate.
	Reason string `json:"reason" api:"nullable"`
	// Domains this attempt requested a certificate for.
	RequestedDomains []string `json:"requested_domains" api:"nullable"`
	// Always `null` for an alias.
	Resource int64 `json:"resource" api:"nullable"`
	// Date when the process of issuing a Let's Encrypt certificate was started (ISO
	// 8601/RFC 3339 format, UTC).
	Started string `json:"started"`
	// Detailed information about attempts to issue a Let's Encrypt certificate.
	Statuses []AliasCertificateStatusStatus `json:"statuses"`
	// State of the domain validation for this attempt.
	//
	// Possible values:
	//
	// - `PENDING` – Validation has not started yet.
	// - `AWAITING_CNAME` – Waiting for the CNAME records below to be created.
	// - `VALIDATING` – Validation is in progress.
	// - `COMPLETED` – Validation succeeded.
	// - `FAILED` – Validation did not succeed.
	ValidationState string `json:"validation_state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Active           respjson.Field
		AttemptsCount    respjson.Field
		ChallengeType    respjson.Field
		Challenges       respjson.Field
		CnameRecords     respjson.Field
		Finished         respjson.Field
		LatestStatus     respjson.Field
		NextAttemptTime  respjson.Field
		Reason           respjson.Field
		RequestedDomains respjson.Field
		Resource         respjson.Field
		Started          respjson.Field
		Statuses         respjson.Field
		ValidationState  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasCertificateStatus) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasCertificateStatusChallenge struct {
	// Domain validation method of the challenge.
	//
	// Possible values:
	//
	// - `dns_01` – Validation via a DNS record.
	// - `http_01` – Validation via an HTTP request.
	ChallengeType string `json:"challenge_type"`
	// Domain the challenge validates.
	Domain string `json:"domain"`
	// Why the challenge failed. Empty unless `state` is `invalid`.
	Error string `json:"error"`
	// State of the challenge.
	//
	// Possible values:
	//
	// - `pending` – The challenge has not been checked yet.
	// - `processing` – Let's Encrypt is validating the challenge.
	// - `valid` – The challenge passed.
	// - `invalid` – The challenge failed; see `error`.
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChallengeType respjson.Field
		Domain        respjson.Field
		Error         respjson.Field
		State         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasCertificateStatusChallenge) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatusChallenge) UnmarshalJSON(data []byte) error {
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

// Detailed information about last attempt to issue a Let's Encrypt certificate.
//
// The field is `null` until the first attempt of the issuing process has been
// recorded.
type AliasCertificateStatusLatestStatus struct {
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
	//   - `DONE` – Attempt is successful. Let's Encrypt certificate was issued.
	//   - `FAILED` – Attempt failed. Let's Encrypt certificate was not issued.
	//   - `CANCELLED` – Attempt is cancelled. Let's Encrypt certificate was not issued.
	//   - `FORBIDDEN` – Issuing of Let's Encrypt certificates is disabled for your
	//     account. Let's Encrypt certificate was not issued.
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
func (r AliasCertificateStatusLatestStatus) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatusLatestStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasCertificateStatusStatus struct {
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
	//   - `DONE` – Attempt is successful. Let's Encrypt certificate was issued.
	//   - `FAILED` – Attempt failed. Let's Encrypt certificate was not issued.
	//   - `CANCELLED` – Attempt is cancelled. Let's Encrypt certificate was not issued.
	//   - `FORBIDDEN` – Issuing of Let's Encrypt certificates is disabled for your
	//     account. Let's Encrypt certificate was not issued.
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
func (r AliasCertificateStatusStatus) RawJSON() string { return r.JSON.raw }
func (r *AliasCertificateStatusStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AliasDetail struct {
	// DNS record needed for Let's Encrypt domain validation of the alias hostname. It
	// is `null` when `automated` is `false`, and when `challenge_type` is `dns_01`
	// until a delegation record has been prepared for the hostname.
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

// DNS record needed for Let's Encrypt domain validation of the alias hostname. It
// is `null` when `automated` is `false`, and when `challenge_type` is `dns_01`
// until a delegation record has been prepared for the hostname.
type AliasDetailSslProvisioning struct {
	// CNAME record to create so validation can be completed for the hostname. Returned
	// when `challenge_type` is `dns_01`.
	AcmeDelegation AliasDetailSslProvisioningAcmeDelegation `json:"acme_delegation"`
	// CNAME record pointing the alias hostname to your CDN CNAME. Returned when
	// `challenge_type` is `http_01`; `null` until a CDN CNAME is assigned to your
	// account.
	HTTPValidation AliasDetailSslProvisioningHTTPValidation `json:"http_validation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcmeDelegation respjson.Field
		HTTPValidation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AliasDetailSslProvisioning) RawJSON() string { return r.JSON.raw }
func (r *AliasDetailSslProvisioning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CNAME record to create so validation can be completed for the hostname. Returned
// when `challenge_type` is `dns_01`.
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

// CNAME record pointing the alias hostname to your CDN CNAME. Returned when
// `challenge_type` is `http_01`; `null` until a CDN CNAME is assigned to your
// account.
type AliasDetailSslProvisioningHTTPValidation struct {
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
func (r AliasDetailSslProvisioningHTTPValidation) RawJSON() string { return r.JSON.raw }
func (r *AliasDetailSslProvisioningHTTPValidation) UnmarshalJSON(data []byte) error {
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
	//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
	//     the alias and starts again after it is resumed.
	//   - `issued` – The latest attempt succeeded.
	//   - `failed` – The latest attempt ended without a certificate.
	//   - `null` – No certificate is managed for the alias: the alias uses a certificate
	//     you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus string `json:"ssl_status" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//   - `pending` – The alias configuration has not reached the CDN yet.
	//   - `http_only` – The alias is served over HTTP only.
	//   - `active` – The alias is served over HTTP and HTTPS.
	//
	// Any of "inactive", "pending", "http_only", "active".
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
	//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
	//     the alias and starts again after it is resumed.
	//   - `issued` – The latest attempt succeeded.
	//   - `failed` – The latest attempt ended without a certificate.
	//   - `null` – No certificate is managed for the alias: the alias uses a certificate
	//     you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus string `json:"ssl_status" api:"nullable"`
	// Alias status.
	//
	// Possible values:
	//
	//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//   - `pending` – The alias configuration has not reached the CDN yet.
	//   - `http_only` – The alias is served over HTTP only.
	//   - `active` – The alias is served over HTTP and HTTPS.
	//
	// Any of "inactive", "pending", "http_only", "active".
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
	// Whether the alias is enabled. Defaults to `true`.
	Active param.Opt[bool] `json:"active,omitzero"`
	// How the alias certificate is managed. Defaults to `true` when `ssl_id` is
	// omitted and to `false` when `ssl_id` is passed.
	//
	// Possible values:
	//
	//   - `true` – A Let's Encrypt certificate is issued and renewed automatically.
	//     `ssl_id` must be omitted.
	//   - `false` – Your own certificate referenced by `ssl_id`, if any, is served.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	// Domain validation method used to issue the Let's Encrypt certificate for the
	// alias.
	//
	// Possible values:
	//
	//   - `dns_01` – Validation via a CNAME record for
	//     `_acme-challenge.<alias hostname>`, shown in
	//     `ssl_provisioning.acme_delegation`.
	//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
	//     hostname must point to your CDN CNAME, shown in
	//     `ssl_provisioning.http_validation`. The request is served with the settings of
	//     the alias's CDN resource, so they must not block or redirect
	//     `/.well-known/acme-challenge/`.
	//
	// Any of "dns_01", "http_01".
	ChallengeType AliasNewParamsChallengeType `json:"challenge_type,omitzero"`
	paramObj
}

func (r AliasNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Domain validation method used to issue the Let's Encrypt certificate for the
// alias.
//
// Possible values:
//
//   - `dns_01` – Validation via a CNAME record for
//     `_acme-challenge.<alias hostname>`, shown in
//     `ssl_provisioning.acme_delegation`.
//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
//     hostname must point to your CDN CNAME, shown in
//     `ssl_provisioning.http_validation`. The request is served with the settings of
//     the alias's CDN resource, so they must not block or redirect
//     `/.well-known/acme-challenge/`.
type AliasNewParamsChallengeType string

const (
	AliasNewParamsChallengeTypeDNS01  AliasNewParamsChallengeType = "dns_01"
	AliasNewParamsChallengeTypeHTTP01 AliasNewParamsChallengeType = "http_01"
)

type AliasUpdateParams struct {
	// Whether the alias is enabled.
	//
	// Possible values:
	//
	//   - `true` – The alias is enabled.
	//   - `false` – The alias is paused and is not served. Certificate issuance stops,
	//     though an attempt already being processed completes.
	Active param.Opt[bool] `json:"active,omitzero"`
	// How the alias certificate is managed. Change it to switch the alias between a
	// Let's Encrypt certificate and your own. Cannot be changed while the alias is
	// paused.
	//
	// Possible values:
	//
	//   - `true` – Switch to a Let's Encrypt certificate. `ssl_id` must be omitted. The
	//     current certificate keeps being served until the new one is issued.
	//   - `false` – Switch to your own certificate. `ssl_id` is required and is served
	//     immediately.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	// ID of your own SSL certificate to serve for the alias instead of the current
	// one. The certificate must cover the alias hostname; a Let's Encrypt certificate
	// issued for one of your CDN resources cannot be used. Available for an alias with
	// `automated` set to `false`, and together with `automated` set to `false` to
	// switch to your own certificate.
	SslID param.Opt[int64] `json:"ssl_id,omitzero"`
	// Domain validation method used to issue the Let's Encrypt certificate for the
	// alias.
	//
	// Possible values:
	//
	//   - `dns_01` – Validation via a CNAME record for
	//     `_acme-challenge.<alias hostname>`, shown in
	//     `ssl_provisioning.acme_delegation`.
	//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
	//     hostname must point to your CDN CNAME, shown in
	//     `ssl_provisioning.http_validation`. The request is served with the settings of
	//     the alias's CDN resource, so they must not block or redirect
	//     `/.well-known/acme-challenge/`.
	//
	// Any of "dns_01", "http_01".
	ChallengeType AliasUpdateParamsChallengeType `json:"challenge_type,omitzero"`
	paramObj
}

func (r AliasUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AliasUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Domain validation method used to issue the Let's Encrypt certificate for the
// alias.
//
// Possible values:
//
//   - `dns_01` – Validation via a CNAME record for
//     `_acme-challenge.<alias hostname>`, shown in
//     `ssl_provisioning.acme_delegation`.
//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
//     hostname must point to your CDN CNAME, shown in
//     `ssl_provisioning.http_validation`. The request is served with the settings of
//     the alias's CDN resource, so they must not block or redirect
//     `/.well-known/acme-challenge/`.
type AliasUpdateParamsChallengeType string

const (
	AliasUpdateParamsChallengeTypeDNS01  AliasUpdateParamsChallengeType = "dns_01"
	AliasUpdateParamsChallengeTypeHTTP01 AliasUpdateParamsChallengeType = "http_01"
)

type AliasListParams struct {
	// How the alias certificate is managed.
	//
	// Possible values:
	//
	// - `true` – A Let's Encrypt certificate is issued and renewed automatically.
	// - `false` – The certificate was added by you.
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
	// Outcome of the certificate we manage for the alias.
	//
	// Possible values:
	//
	//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
	//     the alias and starts again after it is resumed.
	//   - `issued` – The latest attempt succeeded.
	//   - `failed` – The latest attempt ended without a certificate.
	//   - `null` – No certificate is managed for the alias: the alias uses a certificate
	//     you provided, or is automated and has never opened an attempt.
	//
	// Any of "pending", "issued", "failed".
	SslStatus AliasListParamsSslStatus `query:"ssl_status,omitzero" json:"-"`
	// Alias status.
	//
	// Possible values:
	//
	//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
	//     served.
	//   - `pending` – The alias configuration has not reached the CDN yet.
	//   - `http_only` – The alias is served over HTTP only.
	//   - `active` – The alias is served over HTTP and HTTPS.
	//
	// Any of "inactive", "pending", "http_only", "active".
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

// Outcome of the certificate we manage for the alias.
//
// Possible values:
//
//   - `pending` – A certificate attempt is in progress, or was stopped by pausing
//     the alias and starts again after it is resumed.
//   - `issued` – The latest attempt succeeded.
//   - `failed` – The latest attempt ended without a certificate.
//   - `null` – No certificate is managed for the alias: the alias uses a certificate
//     you provided, or is automated and has never opened an attempt.
type AliasListParamsSslStatus string

const (
	AliasListParamsSslStatusPending AliasListParamsSslStatus = "pending"
	AliasListParamsSslStatusIssued  AliasListParamsSslStatus = "issued"
	AliasListParamsSslStatusFailed  AliasListParamsSslStatus = "failed"
)

// Alias status.
//
// Possible values:
//
//   - `inactive` – The alias or its CDN resource is disabled; the alias is not
//     served.
//   - `pending` – The alias configuration has not reached the CDN yet.
//   - `http_only` – The alias is served over HTTP only.
//   - `active` – The alias is served over HTTP and HTTPS.
type AliasListParamsStatus string

const (
	AliasListParamsStatusInactive AliasListParamsStatus = "inactive"
	AliasListParamsStatusPending  AliasListParamsStatus = "pending"
	AliasListParamsStatusHTTPOnly AliasListParamsStatus = "http_only"
	AliasListParamsStatusActive   AliasListParamsStatus = "active"
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
	// How the alias certificate is managed. Defaults to `true` when `ssl_id` is
	// omitted and to `false` when `ssl_id` is passed.
	Automated param.Opt[bool] `json:"automated,omitzero"`
	// Domain validation method used to issue the Let's Encrypt certificate for the
	// alias.
	//
	// Possible values:
	//
	//   - `dns_01` – Validation via a CNAME record for
	//     `_acme-challenge.<alias hostname>`, shown in
	//     `ssl_provisioning.acme_delegation`.
	//   - `http_01` – Validation via an HTTP request to the alias hostname. The alias
	//     hostname must point to your CDN CNAME, shown in
	//     `ssl_provisioning.http_validation`. The request is served with the settings of
	//     the alias's CDN resource, so they must not block or redirect
	//     `/.well-known/acme-challenge/`.
	//
	// Any of "dns_01", "http_01".
	ChallengeType string `json:"challenge_type,omitzero"`
	paramObj
}

func (r AliasNewMultipleParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow AliasNewMultipleParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AliasNewMultipleParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AliasNewMultipleParamsItem](
		"challenge_type", "dns_01", "http_01",
	)
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
