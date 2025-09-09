// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// SecretService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
	task    TaskService
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r SecretService) {
	r = SecretService{}
	r.Options = opts
	r.task = NewTaskService(opts...)
	return
}

// List secrets
func (r *SecretService) List(ctx context.Context, params SecretListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Secret], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List secrets
func (r *SecretService) ListAutoPaging(ctx context.Context, params SecretListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Secret] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
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

// Create secret and poll for the result
func (r *SecretService) UploadTlsCertificateAndPoll(ctx context.Context, params SecretUploadTlsCertificateParams, opts ...option.RequestOption) (v *Secret, err error) {
	resource, err := r.UploadTlsCertificate(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams SecretGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.task.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Secrets) != 1 {
		return nil, errors.New("expected exactly one secret to be created in a task")
	}
	resourceID := task.CreatedResources.Secrets[0]

	return r.Get(ctx, resourceID, getParams, opts...)
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

type SecretListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecretListParams]'s query parameters as `url.Values`.
func (r SecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SecretDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecretGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecretUploadTlsCertificateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
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
func (r *SecretUploadTlsCertificateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
func (r *SecretUploadTlsCertificateParamsPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
