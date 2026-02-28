// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Binaries are WebAssembly executables that are actually executed when app is ran.
//
// BinaryService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBinaryService] method instead.
type BinaryService struct {
	Options []option.RequestOption
}

// NewBinaryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBinaryService(opts ...option.RequestOption) (r BinaryService) {
	r = BinaryService{}
	r.Options = opts
	return
}

// Store compiled WASM binary
func (r *BinaryService) New(ctx context.Context, body io.Reader, opts ...option.RequestOption) (res *BinaryShort, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := "fastedge/v1/binaries/raw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List binaries
func (r *BinaryService) List(ctx context.Context, opts ...option.RequestOption) (res *BinaryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/binaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a binary
func (r *BinaryService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("fastedge/v1/binaries/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get binary
func (r *BinaryService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Binary, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/binaries/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Binary struct {
	// Binary ID
	ID int64 `json:"id" api:"required"`
	// Wasm API type
	APIType string `json:"api_type" api:"required"`
	// Source language:
	// 0 - unknown
	// 1 - Rust
	// 2 - JavaScript
	Source int64 `json:"source" api:"required"`
	// Status code:
	// 0 - pending
	// 1 - compiled
	// 2 - compilation failed (errors available)
	// 3 - compilation failed (errors not available)
	// 4 - resulting binary exceeded the limit
	// 5 - unsupported source language
	Status int64 `json:"status" api:"required"`
	// MD5 hash of the binary
	Checksum string `json:"checksum"`
	// Not used since (UTC)
	UnrefSince string `json:"unref_since"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIType     respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Checksum    respjson.Field
		UnrefSince  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Binary) RawJSON() string { return r.JSON.raw }
func (r *Binary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BinaryShort struct {
	// Binary ID
	ID int64 `json:"id" api:"required"`
	// Wasm API type
	APIType string `json:"api_type" api:"required"`
	// Status code:
	// 0 - pending
	// 1 - compiled
	// 2 - compilation failed (errors available)
	// 3 - compilation failed (errors not available)
	// 4 - resulting binary exceeded the limit
	// 5 - unsupported source language
	Status int64 `json:"status" api:"required"`
	// MD5 hash of the binary
	Checksum string `json:"checksum"`
	// Not used since (UTC)
	UnrefSince string `json:"unref_since"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIType     respjson.Field
		Status      respjson.Field
		Checksum    respjson.Field
		UnrefSince  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BinaryShort) RawJSON() string { return r.JSON.raw }
func (r *BinaryShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BinaryListResponse struct {
	Binaries []BinaryShort `json:"binaries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binaries    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BinaryListResponse) RawJSON() string { return r.JSON.raw }
func (r *BinaryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
