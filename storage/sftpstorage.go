// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/polling"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/tidwall/sjson"
)

// SFTP storages provide file transfer protocol access for securely uploading,
// downloading, and managing files over SSH.
//
// SftpStorageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSftpStorageService] method instead.
type SftpStorageService struct {
	Options []option.RequestOption
}

// NewSftpStorageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSftpStorageService(opts ...option.RequestOption) (r SftpStorageService) {
	r = SftpStorageService{}
	r.Options = opts
	return
}

// Creates a new SFTP storage instance in the specified location and returns the
// storage details including credentials.
func (r *SftpStorageService) New(ctx context.Context, body SftpStorageNewParams, opts ...option.RequestOption) (res *SftpStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/v4/sftp_storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// NewAndPoll creates a new SFTP storage and polls until provisioning is complete,
// returning the original create response (with the one-time Password preserved
// when password_mode is "auto" or "set") and the provisioning status promoted to
// "active". Polling reuses the polling_interval_seconds and
// polling_timeout_seconds client options.
//
// JSON.raw on the returned value is the original POST body, so RawJSON() will
// still report provisioning_status="creating". Use the typed fields for
// post-provisioning state.
func (r *SftpStorageService) NewAndPoll(ctx context.Context, body SftpStorageNewParams, opts ...option.RequestOption) (res *SftpStorage, err error) {
	var raw *http.Response
	actionOpts := slices.Concat(opts, []option.RequestOption{option.WithResponseInto(&raw)})

	created, err := r.New(ctx, body, actionOpts...)
	if err != nil {
		return nil, err
	}

	var rawBytes []byte
	if created == nil {
		// Caller's WithResponseBodyInto overrode the default deserialization
		// target. Recover the typed struct from the raw response so polling can
		// proceed; supports **http.Response and *[]byte body shapes.
		created = &SftpStorage{}
		rawBytes, err = polling.RecoverActionBody(raw, created, opts...)
		if err != nil {
			return nil, fmt.Errorf("sftp NewAndPoll: %w", err)
		}
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, created.ID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusActive {
			// Promote the polled status onto the create response (which still
			// carries the one-time Password). All other fields are populated
			// synchronously on the POST response and don't change during
			// provisioning.
			created.ProvisioningStatus = s.ProvisioningStatus
			if rawBytes != nil && raw != nil {
				enriched, sjErr := sjson.SetBytes(rawBytes, "provisioning_status", string(s.ProvisioningStatus))
				if sjErr != nil {
					return nil, fmt.Errorf("failed to enrich sftp create response with polled provisioning_status: %w", sjErr)
				}
				raw.Body = io.NopCloser(bytes.NewReader(enriched))
				polling.WriteResponseBodyInto(opts, enriched)
			}
			return created, nil
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleting ||
			s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("sftp storage %d entered terminal state %q during creation", created.ID, s.ProvisioningStatus)
		}

		select {
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// Updates SFTP storage configuration and/or credentials including password and SSH
// key management. Supports JSON merge patch semantics: "password": null deletes
// the password, "ssh_key_ids": [] clears all keys.
func (r *SftpStorageService) Update(ctx context.Context, storageID int64, body SftpStorageUpdateParams, opts ...option.RequestOption) (res *SftpStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/sftp_storages/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// UpdateAndPoll updates SFTP storage configuration and polls until provisioning
// is complete, returning the original update response (with a regenerated
// Password preserved when password_mode is "auto") and the provisioning status
// promoted to "active". Polling reuses the polling_interval_seconds and
// polling_timeout_seconds client options.
func (r *SftpStorageService) UpdateAndPoll(ctx context.Context, storageID int64, body SftpStorageUpdateParams, opts ...option.RequestOption) (res *SftpStorage, err error) {
	var raw *http.Response
	actionOpts := slices.Concat(opts, []option.RequestOption{option.WithResponseInto(&raw)})

	updated, err := r.Update(ctx, storageID, body, actionOpts...)
	if err != nil {
		return nil, err
	}

	var rawBytes []byte
	if updated == nil {
		updated = &SftpStorage{}
		rawBytes, err = polling.RecoverActionBody(raw, updated, opts...)
		if err != nil {
			return nil, fmt.Errorf("sftp UpdateAndPoll: %w", err)
		}
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusActive {
			// Promote the polled status onto the PATCH response (which still
			// carries any regenerated Password). All other fields are already
			// authoritative on the PATCH response.
			updated.ProvisioningStatus = s.ProvisioningStatus
			if rawBytes != nil && raw != nil {
				enriched, sjErr := sjson.SetBytes(rawBytes, "provisioning_status", string(s.ProvisioningStatus))
				if sjErr != nil {
					return nil, fmt.Errorf("failed to enrich sftp update response with polled provisioning_status: %w", sjErr)
				}
				raw.Body = io.NopCloser(bytes.NewReader(enriched))
				polling.WriteResponseBodyInto(opts, enriched)
			}
			return updated, nil
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleting ||
			s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("sftp storage %d entered terminal state %q during update", storageID, s.ProvisioningStatus)
		}

		select {
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// Returns a paginated list of SFTP storage instances for the authenticated client.
func (r *SftpStorageService) List(ctx context.Context, query SftpStorageListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SftpStorage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/v4/sftp_storages"
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

// Returns a paginated list of SFTP storage instances for the authenticated client.
func (r *SftpStorageService) ListAutoPaging(ctx context.Context, query SftpStorageListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SftpStorage] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an SFTP storage instance. This operation cannot be undone.
func (r *SftpStorageService) Delete(ctx context.Context, storageID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("storage/v4/sftp_storages/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// DeleteAndPoll deletes an SFTP storage and polls until the resource is fully
// removed. Polling completes when Get returns 404 or the polled
// provisioning_status reaches "deleted". Polling reuses the
// polling_interval_seconds and polling_timeout_seconds client options.
func (r *SftpStorageService) DeleteAndPoll(ctx context.Context, storageID int64, opts ...option.RequestOption) error {
	if err := r.Delete(ctx, storageID, opts...); err != nil {
		return err
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			var apiErr *Error
			if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
				return nil
			}
			return fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil
		}

		select {
		case <-pollingCtx.Done():
			return pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// Returns details of a specific SFTP storage instance (without credentials).
func (r *SftpStorageService) Get(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *SftpStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/sftp_storages/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SftpStorage struct {
	// Unique identifier for the storage instance
	ID int64 `json:"id" api:"required"`
	// Full hostname/address for accessing the storage endpoint
	Address string `json:"address" api:"required"`
	// ISO 8601 timestamp when the storage was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Duration when the storage will expire. Null if no expiration is set.
	Expires string `json:"expires" api:"required"`
	// Read-only internal full name of the storage, composed as "{`client_id`}-{name}".
	// Used by the SFTP backend as the login username. Clients should use this value
	// when connecting but should continue to identify the storage by `name` in their
	// own configuration.
	FullName string `json:"full_name" api:"required"`
	// Whether this storage uses a custom configuration file
	HasCustomConfigFile bool `json:"has_custom_config_file" api:"required"`
	// Whether password authentication is configured for this storage
	HasPassword bool `json:"has_password" api:"required"`
	// Whether HTTP access is disabled for this storage (HTTPS only)
	IsHTTPDisabled bool `json:"is_http_disabled" api:"required"`
	// Geographic location code where the storage is provisioned
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance, as supplied at creation time.
	Name string `json:"name" api:"required"`
	// Lifecycle status of the storage. Use this to check readiness before operations.
	//
	// Any of "creating", "active", "updating", "deleting", "deleted".
	ProvisioningStatus SftpStorageProvisioningStatus `json:"provisioning_status" api:"required"`
	// Custom domain alias for accessing the storage. Null if no alias is configured.
	ServerAlias string `json:"server_alias" api:"required"`
	// IDs of SSH keys associated with this SFTP storage
	SSHKeyIDs []int64 `json:"ssh_key_ids" api:"required"`
	// SFTP password. Only returned when newly generated or set (create/patch). Omitted
	// in GET/list responses.
	Password string `json:"password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Address             respjson.Field
		CreatedAt           respjson.Field
		Expires             respjson.Field
		FullName            respjson.Field
		HasCustomConfigFile respjson.Field
		HasPassword         respjson.Field
		IsHTTPDisabled      respjson.Field
		LocationName        respjson.Field
		Name                respjson.Field
		ProvisioningStatus  respjson.Field
		ServerAlias         respjson.Field
		SSHKeyIDs           respjson.Field
		Password            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SftpStorage) RawJSON() string { return r.JSON.raw }
func (r *SftpStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the storage. Use this to check readiness before operations.
type SftpStorageProvisioningStatus string

const (
	SftpStorageProvisioningStatusCreating SftpStorageProvisioningStatus = "creating"
	SftpStorageProvisioningStatusActive   SftpStorageProvisioningStatus = "active"
	SftpStorageProvisioningStatusUpdating SftpStorageProvisioningStatus = "updating"
	SftpStorageProvisioningStatusDeleting SftpStorageProvisioningStatus = "deleting"
	SftpStorageProvisioningStatusDeleted  SftpStorageProvisioningStatus = "deleted"
)

type SftpStorageNewParams struct {
	// Location code where the storage should be created
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance
	Name string `json:"name" api:"required"`
	// Password handling mode for SFTP access: 'auto': generate a random password
	// (returned in the response) 'set': use the password provided in `sftp_password`
	// 'none': no password (SSH-key-only access)
	//
	// Any of "auto", "set", "none".
	PasswordMode SftpStorageNewParamsPasswordMode `json:"password_mode,omitzero" api:"required"`
	// Duration when the storage should expire (e.g., "2 years 6 months"). Omit for no
	// expiration.
	Expires param.Opt[string] `json:"expires,omitzero"`
	// Whether this storage should use a custom configuration file
	HasCustomConfigFile param.Opt[bool] `json:"has_custom_config_file,omitzero"`
	// Whether HTTP access should be disabled (HTTPS only)
	IsHTTPDisabled param.Opt[bool] `json:"is_http_disabled,omitzero"`
	// Custom domain alias for accessing the storage. Omit for no alias.
	ServerAlias param.Opt[string] `json:"server_alias,omitzero"`
	// SFTP password (8-63 chars). Required when `password_mode` is 'set'. Must be
	// omitted when `password_mode` is 'auto' or 'none'.
	SftpPassword param.Opt[string] `json:"sftp_password,omitzero"`
	// SSH key IDs to associate with this storage at creation time. If omitted, no keys
	// are linked.
	SSHKeyIDs []int64 `json:"ssh_key_ids,omitzero"`
	paramObj
}

func (r SftpStorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SftpStorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SftpStorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Password handling mode for SFTP access: 'auto': generate a random password
// (returned in the response) 'set': use the password provided in `sftp_password`
// 'none': no password (SSH-key-only access)
type SftpStorageNewParamsPasswordMode string

const (
	SftpStorageNewParamsPasswordModeAuto SftpStorageNewParamsPasswordMode = "auto"
	SftpStorageNewParamsPasswordModeSet  SftpStorageNewParamsPasswordMode = "set"
	SftpStorageNewParamsPasswordModeNone SftpStorageNewParamsPasswordMode = "none"
)

type SftpStorageUpdateParams struct {
	// Duration when the storage should expire (e.g., "2 years 6 months"). Empty string
	// to remove.
	Expires param.Opt[string] `json:"expires,omitzero"`
	// Whether this storage should use a custom configuration file
	HasCustomConfigFile param.Opt[bool] `json:"has_custom_config_file,omitzero"`
	// Whether HTTP access should be disabled (HTTPS only)
	IsHTTPDisabled param.Opt[bool] `json:"is_http_disabled,omitzero"`
	// Custom domain alias for accessing the storage. Empty string to remove.
	ServerAlias param.Opt[string] `json:"server_alias,omitzero"`
	// Password handling mode. Omit to leave password unchanged. 'auto': regenerate
	// password (returned in response) 'none': remove password Note: 'set' is not
	// allowed in PATCH.
	//
	// Any of "auto", "none".
	PasswordMode SftpStorageUpdateParamsPasswordMode `json:"password_mode,omitzero"`
	// SSH key IDs to associate with this storage. Replaces all existing keys. If
	// omitted, existing keys are unchanged. If empty array, all keys are removed.
	SSHKeyIDs []int64 `json:"ssh_key_ids,omitzero"`
	paramObj
}

func (r SftpStorageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SftpStorageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SftpStorageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Password handling mode. Omit to leave password unchanged. 'auto': regenerate
// password (returned in response) 'none': remove password Note: 'set' is not
// allowed in PATCH.
type SftpStorageUpdateParamsPasswordMode string

const (
	SftpStorageUpdateParamsPasswordModeAuto SftpStorageUpdateParamsPasswordMode = "auto"
	SftpStorageUpdateParamsPasswordModeNone SftpStorageUpdateParamsPasswordMode = "none"
)

type SftpStorageListParams struct {
	// Filter by storage ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by storage location/region
	LocationName param.Opt[string] `query:"location_name,omitzero" json:"-"`
	// Filter by storage name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of records to skip
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Include deleted storages
	ShowDeleted param.Opt[bool] `query:"show_deleted,omitzero" json:"-"`
	// Filter by provisioning status
	//
	// Any of "active", "creating", "updating", "deleting", "deleted".
	ProvisioningStatus SftpStorageListParamsProvisioningStatus `query:"provisioning_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SftpStorageListParams]'s query parameters as `url.Values`.
func (r SftpStorageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by provisioning status
type SftpStorageListParamsProvisioningStatus string

const (
	SftpStorageListParamsProvisioningStatusActive   SftpStorageListParamsProvisioningStatus = "active"
	SftpStorageListParamsProvisioningStatusCreating SftpStorageListParamsProvisioningStatus = "creating"
	SftpStorageListParamsProvisioningStatusUpdating SftpStorageListParamsProvisioningStatus = "updating"
	SftpStorageListParamsProvisioningStatusDeleting SftpStorageListParamsProvisioningStatus = "deleting"
	SftpStorageListParamsProvisioningStatusDeleted  SftpStorageListParamsProvisioningStatus = "deleted"
)
