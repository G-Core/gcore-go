// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// CredentialService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	Options []option.RequestOption
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r CredentialService) {
	r = CredentialService{}
	r.Options = opts
	return
}

// Generates new access credentials for the storage (S3 keys for S3 storage, SFTP
// password for SFTP storage).
func (r *CredentialService) Recreate(ctx context.Context, storageID int64, body CredentialRecreateParams, opts ...option.RequestOption) (res *Storage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/credentials", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CredentialRecreateParams struct {
	DeleteSftpPassword   param.Opt[bool]   `json:"delete_sftp_password,omitzero"`
	GenerateS3Keys       param.Opt[bool]   `json:"generate_s3_keys,omitzero"`
	GenerateSftpPassword param.Opt[bool]   `json:"generate_sftp_password,omitzero"`
	ResetSftpKeys        param.Opt[bool]   `json:"reset_sftp_keys,omitzero"`
	SftpPassword         param.Opt[string] `json:"sftp_password,omitzero"`
	paramObj
}

func (r CredentialRecreateParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialRecreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialRecreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
