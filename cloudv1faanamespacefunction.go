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

// CloudV1FaaNamespaceFunctionService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FaaNamespaceFunctionService] method instead.
type CloudV1FaaNamespaceFunctionService struct {
	Options []option.RequestOption
}

// NewCloudV1FaaNamespaceFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1FaaNamespaceFunctionService(opts ...option.RequestOption) (r *CloudV1FaaNamespaceFunctionService) {
	r = &CloudV1FaaNamespaceFunctionService{}
	r.Options = opts
	return
}

// Create function
func (r *CloudV1FaaNamespaceFunctionService) New(ctx context.Context, projectID int64, regionID int64, namespaceName string, body CloudV1FaaNamespaceFunctionNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get function
func (r *CloudV1FaaNamespaceFunctionService) Get(ctx context.Context, projectID int64, regionID int64, namespaceName string, functionName string, opts ...option.RequestOption) (res *Function, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	if functionName == "" {
		err = errors.New("missing required function_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions/%s", projectID, regionID, namespaceName, functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change function
func (r *CloudV1FaaNamespaceFunctionService) Update(ctx context.Context, projectID int64, regionID int64, namespaceName string, functionName string, body CloudV1FaaNamespaceFunctionUpdateParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	if functionName == "" {
		err = errors.New("missing required function_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions/%s", projectID, regionID, namespaceName, functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List functions
func (r *CloudV1FaaNamespaceFunctionService) List(ctx context.Context, projectID int64, regionID int64, namespaceName string, query CloudV1FaaNamespaceFunctionListParams, opts ...option.RequestOption) (res *CloudV1FaaNamespaceFunctionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete function
func (r *CloudV1FaaNamespaceFunctionService) Delete(ctx context.Context, projectID int64, regionID int64, namespaceName string, functionName string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	if functionName == "" {
		err = errors.New("missing required function_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions/%s", projectID, regionID, namespaceName, functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Path Parameter: project_id (int) -- Project ID region_id (int) -- Region ID
// namespace_name (str) -- Namespace name function_name (str) -- Function name
//
// Query Parameter: limit (int) -- Optional. Log line limit
//
// HTTP 200 Response: logs (LogsSchema) -- Function logs
//
// Error Response: 404 -- Resource not found
//
// Tags: Function-as-a-Service
//
// Security: Bearer
func (r *CloudV1FaaNamespaceFunctionService) GetLogs(ctx context.Context, projectID int64, regionID int64, namespaceName string, functionName string, query CloudV1FaaNamespaceFunctionGetLogsParams, opts ...option.RequestOption) (res *CloudV1FaaNamespaceFunctionGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	if functionName == "" {
		err = errors.New("missing required function_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/faas/namespaces/%v/%v/%s/functions/%s/logs", projectID, regionID, namespaceName, functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FaasAutoscaling struct {
	// The maximum number of instances to run.
	MaxInstances int64 `json:"max_instances,required"`
	// The minimum number of instances to run.
	MinInstances int64               `json:"min_instances,required"`
	JSON         faasAutoscalingJSON `json:"-"`
}

// faasAutoscalingJSON contains the JSON metadata for the struct [FaasAutoscaling]
type faasAutoscalingJSON struct {
	MaxInstances apijson.Field
	MinInstances apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FaasAutoscaling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r faasAutoscalingJSON) RawJSON() string {
	return r.raw
}

type FaasAutoscalingParam struct {
	// The maximum number of instances to run.
	MaxInstances param.Field[int64] `json:"max_instances,required"`
	// The minimum number of instances to run.
	MinInstances param.Field[int64] `json:"min_instances,required"`
}

func (r FaasAutoscalingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Function struct {
	// Function ID.
	ID string `json:"id,required"`
	// Autoscaling configuration for the function. Keys must be 'min_instances' or
	// 'max_instances', and values must be integers between 0 and 50.
	Autoscaling FaasAutoscaling `json:"autoscaling,required"`
	// Function build message.
	BuildMessage string `json:"build_message,required"`
	// Function build status.
	BuildStatus string `json:"build_status,required"`
	// Function code text.
	CodeText string `json:"code_text,required"`
	// Function creation date.
	CreatedAt string `json:"created_at,required"`
	// Dependencies for the function to install.
	Dependencies string `json:"dependencies,required"`
	// Function deploy status.
	DeployStatus map[string]int64 `json:"deploy_status,required"`
	// Description of the function.
	Description string `json:"description,required"`
	// Function endpoint.
	Endpoint string `json:"endpoint,required"`
	// The name of the flavor associated with the function.
	Flavor string `json:"flavor,required"`
	// The main startup method name.
	MainMethod string `json:"main_method,required"`
	// Function name.
	Name string `json:"name,required"`
	// Function runtime.
	Runtime string `json:"runtime,required"`
	// Function status.
	Status string `json:"status,required"`
	// Function timeout in seconds.
	Timeout int64 `json:"timeout,required"`
	// Enable or disable API key authentication.
	EnableAPIKey bool `json:"enable_api_key,nullable"`
	// Environment variables for the function. Keys must match a specific regex pattern
	// and be 1-255 characters long, and values must be 0-255 characters long.
	Envs interface{}  `json:"envs,nullable"`
	JSON functionJSON `json:"-"`
}

// functionJSON contains the JSON metadata for the struct [Function]
type functionJSON struct {
	ID           apijson.Field
	Autoscaling  apijson.Field
	BuildMessage apijson.Field
	BuildStatus  apijson.Field
	CodeText     apijson.Field
	CreatedAt    apijson.Field
	Dependencies apijson.Field
	DeployStatus apijson.Field
	Description  apijson.Field
	Endpoint     apijson.Field
	Flavor       apijson.Field
	MainMethod   apijson.Field
	Name         apijson.Field
	Runtime      apijson.Field
	Status       apijson.Field
	Timeout      apijson.Field
	EnableAPIKey apijson.Field
	Envs         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaNamespaceFunctionListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Function                                  `json:"results,required"`
	JSON    cloudV1FaaNamespaceFunctionListResponseJSON `json:"-"`
}

// cloudV1FaaNamespaceFunctionListResponseJSON contains the JSON metadata for the
// struct [CloudV1FaaNamespaceFunctionListResponse]
type cloudV1FaaNamespaceFunctionListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaNamespaceFunctionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaNamespaceFunctionListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaNamespaceFunctionGetLogsResponse struct {
	// Function logs.
	Logs []map[string]string                            `json:"logs,required"`
	JSON cloudV1FaaNamespaceFunctionGetLogsResponseJSON `json:"-"`
}

// cloudV1FaaNamespaceFunctionGetLogsResponseJSON contains the JSON metadata for
// the struct [CloudV1FaaNamespaceFunctionGetLogsResponse]
type cloudV1FaaNamespaceFunctionGetLogsResponseJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FaaNamespaceFunctionGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FaaNamespaceFunctionGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FaaNamespaceFunctionNewParams struct {
	// Autoscaling configuration for the function. Keys must be 'min_instances' or
	// 'max_instances', and values must be integers between 0 and 50.
	Autoscaling param.Field[FaasAutoscalingParam] `json:"autoscaling,required"`
	// Function code text.
	CodeText param.Field[string] `json:"code_text,required"`
	// The name of the flavor associated with the function.
	Flavor param.Field[string] `json:"flavor,required"`
	// The main startup method name.
	MainMethod param.Field[string] `json:"main_method,required"`
	// Function name.
	Name param.Field[string] `json:"name,required"`
	// Function runtime.
	Runtime param.Field[string] `json:"runtime,required"`
	// Function timeout in seconds.
	Timeout param.Field[int64] `json:"timeout,required"`
	// Dependencies for the function to install.
	Dependencies param.Field[string] `json:"dependencies"`
	// Description of the function.
	Description param.Field[string] `json:"description"`
	// Enable or disable API key authentication.
	EnableAPIKey param.Field[bool] `json:"enable_api_key"`
	// Environment variables for the function. Keys must match a specific regex pattern
	// and be 1-255 characters long, and values must be 0-255 characters long.
	Envs param.Field[interface{}] `json:"envs"`
}

func (r CloudV1FaaNamespaceFunctionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaNamespaceFunctionUpdateParams struct {
	// Autoscaling configuration for the function. Keys must be 'min_instances' or
	// 'max_instances', and values must be integers between 0 and 50.
	Autoscaling param.Field[FaasAutoscalingParam] `json:"autoscaling"`
	// Function code text.
	CodeText param.Field[string] `json:"code_text"`
	// Dependencies for the function to install.
	Dependencies param.Field[string] `json:"dependencies"`
	// Description of the function.
	Description param.Field[string] `json:"description"`
	// Set to true if the function is disabled.
	Disabled param.Field[bool] `json:"disabled"`
	// Enable or disable API key authentication.
	EnableAPIKey param.Field[bool] `json:"enable_api_key"`
	// Environment variables for the function. Keys must match a specific regex pattern
	// and be 1-255 characters long, and values must be 0-255 characters long.
	Envs param.Field[interface{}] `json:"envs"`
	// The name of the flavor associated with the function.
	Flavor param.Field[string] `json:"flavor"`
	// List of used authentication API keys, matching a specific regex pattern.
	Keys param.Field[[]string] `json:"keys"`
	// The main startup method name.
	MainMethod param.Field[string] `json:"main_method"`
	// Function timeout in seconds.
	Timeout param.Field[int64] `json:"timeout"`
}

func (r CloudV1FaaNamespaceFunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FaaNamespaceFunctionListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Ordering fields and directions.
	OrderBy param.Field[[]FunctionsOrderByChoices] `query:"order_by"`
	// Search through which value
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [CloudV1FaaNamespaceFunctionListParams]'s query parameters
// as `url.Values`.
func (r CloudV1FaaNamespaceFunctionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1FaaNamespaceFunctionGetLogsParams struct {
	// Limit the number of logs lines
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CloudV1FaaNamespaceFunctionGetLogsParams]'s query
// parameters as `url.Values`.
func (r CloudV1FaaNamespaceFunctionGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
