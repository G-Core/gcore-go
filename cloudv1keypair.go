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

// CloudV1KeypairService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1KeypairService] method instead.
type CloudV1KeypairService struct {
	Options []option.RequestOption
}

// NewCloudV1KeypairService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1KeypairService(opts ...option.RequestOption) (r *CloudV1KeypairService) {
	r = &CloudV1KeypairService{}
	r.Options = opts
	return
}

// To generate a keypair do not use the public_key parameter in the request body
func (r *CloudV1KeypairService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1KeypairNewParams, opts ...option.RequestOption) (res *Keypair, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/keypairs/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get keypair
func (r *CloudV1KeypairService) Get(ctx context.Context, projectID int64, regionID int64, keypairID string, opts ...option.RequestOption) (res *Keypair, err error) {
	opts = append(r.Options[:], opts...)
	if keypairID == "" {
		err = errors.New("missing required keypair_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/keypairs/%v/%v/%s", projectID, regionID, keypairID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List keypairs
func (r *CloudV1KeypairService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *ListKeypair, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/keypairs/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete keypair
func (r *CloudV1KeypairService) Delete(ctx context.Context, projectID int64, regionID int64, keypairID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keypairID == "" {
		err = errors.New("missing required keypair_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/keypairs/%v/%v/%s", projectID, regionID, keypairID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Share keypair to view for all users in project
func (r *CloudV1KeypairService) Share(ctx context.Context, projectID int64, regionID int64, keypairID string, body CloudV1KeypairShareParams, opts ...option.RequestOption) (res *Keypair, err error) {
	opts = append(r.Options[:], opts...)
	if keypairID == "" {
		err = errors.New("missing required keypair_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/keypairs/%v/%v/%s/share", projectID, regionID, keypairID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Keypair struct {
	// Key ID, equal to sshkey_name
	ID string `json:"id,required"`
	// Keypair name
	Name string `json:"name,required"`
	// Public part of the key. To generate public and private keys in the platform, do
	// not specify the parameter in the request body.
	PublicKey string `json:"public_key,required"`
	// Keypair is shared for all users in the project
	SharedInProject bool `json:"shared_in_project,required"`
	// Keypair state
	State KeypairState `json:"state,required"`
	// Keypair creation datetime
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Key fingerprint
	Fingerprint string `json:"fingerprint,nullable"`
	// Private part of the key
	PrivateKey string `json:"private_key,nullable"`
	// Project ID
	ProjectID int64       `json:"project_id,nullable"`
	JSON      keypairJSON `json:"-"`
}

// keypairJSON contains the JSON metadata for the struct [Keypair]
type keypairJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	PublicKey       apijson.Field
	SharedInProject apijson.Field
	State           apijson.Field
	CreatedAt       apijson.Field
	Fingerprint     apijson.Field
	PrivateKey      apijson.Field
	ProjectID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Keypair) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keypairJSON) RawJSON() string {
	return r.raw
}

// Keypair state
type KeypairState string

const (
	KeypairStateActive   KeypairState = "ACTIVE"
	KeypairStateDeleting KeypairState = "DELETING"
)

func (r KeypairState) IsKnown() bool {
	switch r {
	case KeypairStateActive, KeypairStateDeleting:
		return true
	}
	return false
}

type ListKeypair struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Keypair       `json:"results,required"`
	JSON    listKeypairJSON `json:"-"`
}

// listKeypairJSON contains the JSON metadata for the struct [ListKeypair]
type listKeypairJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListKeypair) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listKeypairJSON) RawJSON() string {
	return r.raw
}

type CloudV1KeypairNewParams struct {
	// Keypair name
	SshkeyName param.Field[string] `json:"sshkey_name,required"`
	// Public part of the key. To generate public and private keys in the platform, do
	// not specify the parameter in the request body.
	PublicKey param.Field[string] `json:"public_key"`
	// Keypair is shared for all users in the project. Default True
	SharedInProject param.Field[bool] `json:"shared_in_project"`
}

func (r CloudV1KeypairNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1KeypairShareParams struct {
	// Data for share(True)/unshare(False) Keypair between all users in project
	SharedInProject param.Field[bool] `json:"shared_in_project,required"`
}

func (r CloudV1KeypairShareParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
