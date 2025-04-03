// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2KeypairService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2KeypairService] method instead.
type CloudV2KeypairService struct {
	Options []option.RequestOption
}

// NewCloudV2KeypairService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2KeypairService(opts ...option.RequestOption) (r *CloudV2KeypairService) {
	r = &CloudV2KeypairService{}
	r.Options = opts
	return
}

// To generate a keypair do not use the public_key parameter in the request body
func (r *CloudV2KeypairService) New(ctx context.Context, body CloudV2KeypairNewParams, opts ...option.RequestOption) (res *Keypair, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/keypairs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get keypair
func (r *CloudV2KeypairService) Get(ctx context.Context, keypairID string, opts ...option.RequestOption) (res *Keypair, err error) {
	opts = append(r.Options[:], opts...)
	if keypairID == "" {
		err = errors.New("missing required keypair_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/keypairs/%s", keypairID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List keypairs
func (r *CloudV2KeypairService) List(ctx context.Context, query CloudV2KeypairListParams, opts ...option.RequestOption) (res *ListKeypair, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/keypairs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete keypair
func (r *CloudV2KeypairService) Delete(ctx context.Context, keypairID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keypairID == "" {
		err = errors.New("missing required keypair_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/keypairs/%s", keypairID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CloudV2KeypairNewParams struct {
	// Project ID
	ProjectID param.Field[int64] `json:"project_id,required"`
	// Keypair name
	SshkeyName param.Field[string] `json:"sshkey_name,required"`
	// Public part of the key. To generate public and private keys in the platform, do
	// not specify the parameter in the request body.
	PublicKey param.Field[string] `json:"public_key"`
	// Keypair is shared for all users in the project. Default True
	SharedInProject param.Field[bool] `json:"shared_in_project"`
}

func (r CloudV2KeypairNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2KeypairListParams struct {
	// Project ID. It's required with rights lower than a cloud administrator.
	ProjectID param.Field[int64] `query:"project_id"`
	// User ID. To use it, you need set project id or use cloud administrator rights.
	UserID param.Field[int64] `query:"user_id"`
}

// URLQuery serializes [CloudV2KeypairListParams]'s query parameters as
// `url.Values`.
func (r CloudV2KeypairListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
