// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ZoneDnssecService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDnssecService] method instead.
type ZoneDnssecService struct {
	Options []option.RequestOption
}

// NewZoneDnssecService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDnssecService(opts ...option.RequestOption) (r ZoneDnssecService) {
	r = ZoneDnssecService{}
	r.Options = opts
	return
}

// Enable or disable DNSSEC for a DNS zone.
func (r *ZoneDnssecService) Update(ctx context.Context, name string, body ZoneDnssecUpdateParams, opts ...option.RequestOption) (res *ZoneDnssecUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/dnssec", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get DNSSEC DS for a DNS zone.
func (r *ZoneDnssecService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneDnssecGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/dnssec", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneDnssecUpdateResponse struct {
	// Specifies the algorithm used for the key.
	Algorithm string `json:"algorithm"`
	// Represents the hashed value of the DS record.
	Digest string `json:"digest"`
	// Specifies the algorithm used to generate the digest.
	DigestAlgorithm string `json:"digest_algorithm"`
	// Specifies the type of the digest algorithm used.
	DigestType string `json:"digest_type"`
	// Represents the complete DS record.
	Ds string `json:"ds"`
	// Represents the flag for DNSSEC record.
	Flags int64 `json:"flags"`
	// Represents the identifier of the DNSKEY record.
	KeyTag int64 `json:"key_tag"`
	// Specifies the type of the key used in the algorithm.
	KeyType string `json:"key_type"`
	Message string `json:"message"`
	// Represents the public key used in the DS record.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Algorithm       respjson.Field
		Digest          respjson.Field
		DigestAlgorithm respjson.Field
		DigestType      respjson.Field
		Ds              respjson.Field
		Flags           respjson.Field
		KeyTag          respjson.Field
		KeyType         respjson.Field
		Message         respjson.Field
		PublicKey       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneDnssecUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneDnssecUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecGetResponse struct {
	// Specifies the algorithm used for the key.
	Algorithm string `json:"algorithm"`
	// Represents the hashed value of the DS record.
	Digest string `json:"digest"`
	// Specifies the algorithm used to generate the digest.
	DigestAlgorithm string `json:"digest_algorithm"`
	// Specifies the type of the digest algorithm used.
	DigestType string `json:"digest_type"`
	// Represents the complete DS record.
	Ds string `json:"ds"`
	// Represents the flag for DNSSEC record.
	Flags int64 `json:"flags"`
	// Represents the identifier of the DNSKEY record.
	KeyTag int64 `json:"key_tag"`
	// Specifies the type of the key used in the algorithm.
	KeyType string `json:"key_type"`
	// Represents the public key used in the DS record.
	PublicKey string `json:"public_key"`
	Uuid      string `json:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Algorithm       respjson.Field
		Digest          respjson.Field
		DigestAlgorithm respjson.Field
		DigestType      respjson.Field
		Ds              respjson.Field
		Flags           respjson.Field
		KeyTag          respjson.Field
		KeyType         respjson.Field
		PublicKey       respjson.Field
		Uuid            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneDnssecGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneDnssecGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecUpdateParams struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ZoneDnssecUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneDnssecUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneDnssecUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
