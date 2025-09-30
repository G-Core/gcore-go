// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LogsUploaderService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogsUploaderService] method instead.
type LogsUploaderService struct {
	Options  []option.RequestOption
	Policies LogsUploaderPolicyService
	Targets  LogsUploaderTargetService
	Configs  LogsUploaderConfigService
}

// NewLogsUploaderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogsUploaderService(opts ...option.RequestOption) (r LogsUploaderService) {
	r = LogsUploaderService{}
	r.Options = opts
	r.Policies = NewLogsUploaderPolicyService(opts...)
	r.Targets = NewLogsUploaderTargetService(opts...)
	r.Configs = NewLogsUploaderConfigService(opts...)
	return
}

type LogsUploaderValidation struct {
	// Error code indicating the type of validation error.
	Code int64 `json:"code"`
	// Error message if the validation failed.
	Details string `json:"details"`
	// Status of the validation.
	//
	// Any of "in_progress", "successful", "failed".
	Status LogsUploaderValidationStatus `json:"status"`
	// Time when the validation status was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		Status      respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderValidation) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the validation.
type LogsUploaderValidationStatus string

const (
	LogsUploaderValidationStatusInProgress LogsUploaderValidationStatus = "in_progress"
	LogsUploaderValidationStatusSuccessful LogsUploaderValidationStatus = "successful"
	LogsUploaderValidationStatusFailed     LogsUploaderValidationStatus = "failed"
)
