// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// NetworkMappingService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkMappingService] method instead.
type NetworkMappingService struct {
	Options []option.RequestOption
}

// NewNetworkMappingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkMappingService(opts ...option.RequestOption) (r NetworkMappingService) {
	r = NetworkMappingService{}
	r.Options = opts
	return
}

// Create new network mapping.
//
// Example of request:
//
// ```
// curl --location --request POST 'https://api.gcore.com/dns/v2/network-mappings' \
// --header 'Authorization: Bearer ...' \
// --header 'Content-Type: application/json' \
//
//	--data-raw '{
//		"name": "test",
//		"mapping": [
//			{
//				"tags": [
//					"tag1"
//				],
//				"cidr4": [
//					"192.0.2.0/24",
//					"198.0.100.0/24"
//				]
//			},
//			{
//				"tags": [
//					"tag2",
//					"tag3"
//				],
//				"cidr4": [
//					"192.1.2.0/24",
//					"198.1.100.0/24"
//				],
//				"cidr6": [
//					"aa:10::/64"
//				]
//			}
//		]
//	}'
//
// ```
func (r *NetworkMappingService) New(ctx context.Context, body NetworkMappingNewParams, opts ...option.RequestOption) (res *NetworkMappingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/network-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List of network mappings.
//
// Example of request:
//
// ```
//
//	curl --location --request GET 'https://api.gcore.com/dns/v2/network-mappings' \
//	--header 'Authorization: Bearer ...'
//
// ```
func (r *NetworkMappingService) List(ctx context.Context, query NetworkMappingListParams, opts ...option.RequestOption) (res *NetworkMappingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/network-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete network mapping.
//
// Example of request:
//
// ```
// curl --location --request DELETE 'https://api.gcore.com/dns/v2/network-mappings/123' \
// --header 'Authorization: Bearer ...'
// ```
func (r *NetworkMappingService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *NetworkMappingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("dns/v2/network-mappings/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Particular network mapping item info
//
// Example of request:
//
// ```
// curl --location --request GET 'https://api.gcore.com/dns/v2/network-mappings/123' \
// --header 'Authorization: Bearer ...'
// ```
func (r *NetworkMappingService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *DNSNetworkMapping, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("dns/v2/network-mappings/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get network mapping by name.
//
// # Particular network mapping item info
//
// Example of request:
//
// ```
// curl --location --request GET 'https://api.gcore.com/dns/v2/network-mappings/test-mapping' \
// --header 'Authorization: Bearer ...'
// ```
func (r *NetworkMappingService) GetByName(ctx context.Context, name string, opts ...option.RequestOption) (res *DNSNetworkMapping, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/network-mappings/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Import network mapping from YAML file.
//
// Note: A YAML file use spaces as indentation, tabs are not allowed. Example of
// input file:
//
// ```
//
//	name: mapping_rule_1
//	mapping:
//	    - tags:
//	        - tag_name_1
//	      cidr4:
//	        - 127.0.2.0/24
//	    - tags:
//	        - tag_name_2
//	        - tag_name_3
//	      cidr4:
//	        - 128.0.1.0/24
//	        - 128.0.2.0/24
//	        - 128.0.3.0/24
//	      cidr6:
//	        - ac:20::0/64
//
// ---
//
//	name: mapping_rule_2
//	mapping:
//	    - tags:
//	        - my_network
//	      cidr4:
//	        - 129.0.2.0/24
//	      cidr6:
//	        - ac:20::0/64
//
// ```
//
// Example of request:
//
// ```
// curl --location --request POST 'https://api.gcore.com/dns/v2/network-mappings/import' \
// --header 'Authorization: Bearer ...' \
// --header 'Content-Type: text/plain' \
// --data-raw 'name: mapping_rule_1
// mapping:
//   - tags:
//   - tag_name_1
//     cidr4:
//   - 127.0.2.0/24
//   - tags:
//   - tag_name_2
//   - tag_name_3
//     cidr4:
//   - 128.0.1.0/24
//   - 128.0.2.0/24
//   - 128.0.3.0/24
//     cidr6:
//   - aa:10::/64
//
// ---
// name: mapping_rule_2
// mapping:
//   - tags:
//   - my_network
//     cidr4:
//   - 129.0.2.0/24
//     cidr6:
//   - ac:20::0/64'
//
// ```
func (r *NetworkMappingService) Import(ctx context.Context, opts ...option.RequestOption) (res *NetworkMappingImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/network-mappings/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update network mapping (Note: name of network mapping cannot be changed)
//
// Example of request:
//
// ```
// curl --location --request PUT 'https://api.gcore.com/dns/v2/network-mappings/123' \
// --header 'Authorization: Bearer ...' \
// --header 'Content-Type: application/json' \
//
//	--data-raw '{
//		"name": "test-mapping",
//		"mapping": [
//			{
//				"tags": [
//					"tag1"
//				],
//				"cidr4": [
//					"192.0.2.0/24"
//				]
//			},
//			{
//				"tags": [
//					"tag2",
//					"tag3"
//				],
//				"cidr4": [
//					"192.1.2.0/24"
//				],
//				"cidr6": [
//					"aa:10::/64"
//				]
//			}
//		]
//	}'
//
// ```
func (r *NetworkMappingService) Replace(ctx context.Context, id int64, body NetworkMappingReplaceParams, opts ...option.RequestOption) (res *NetworkMappingReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("dns/v2/network-mappings/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type DNSMappingEntry struct {
	Cidr4 []any    `json:"cidr4"`
	Cidr6 []any    `json:"cidr6"`
	Tags  []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cidr4       respjson.Field
		Cidr6       respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSMappingEntry) RawJSON() string { return r.JSON.raw }
func (r *DNSMappingEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DNSMappingEntry to a DNSMappingEntryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DNSMappingEntryParam.Overrides()
func (r DNSMappingEntry) ToParam() DNSMappingEntryParam {
	return param.Override[DNSMappingEntryParam](json.RawMessage(r.RawJSON()))
}

type DNSMappingEntryParam struct {
	Cidr4 []any    `json:"cidr4,omitzero"`
	Cidr6 []any    `json:"cidr6,omitzero"`
	Tags  []string `json:"tags,omitzero"`
	paramObj
}

func (r DNSMappingEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow DNSMappingEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSMappingEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSNetworkMapping struct {
	ID      int64             `json:"id"`
	Mapping []DNSMappingEntry `json:"mapping"`
	Name    string            `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Mapping     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSNetworkMapping) RawJSON() string { return r.JSON.raw }
func (r *DNSNetworkMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DNSNetworkMapping to a DNSNetworkMappingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DNSNetworkMappingParam.Overrides()
func (r DNSNetworkMapping) ToParam() DNSNetworkMappingParam {
	return param.Override[DNSNetworkMappingParam](json.RawMessage(r.RawJSON()))
}

type DNSNetworkMappingParam struct {
	ID      param.Opt[int64]       `json:"id,omitzero"`
	Name    param.Opt[string]      `json:"name,omitzero"`
	Mapping []DNSMappingEntryParam `json:"mapping,omitzero"`
	paramObj
}

func (r DNSNetworkMappingParam) MarshalJSON() (data []byte, err error) {
	type shadow DNSNetworkMappingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSNetworkMappingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkMappingNewResponse struct {
	ID int64 `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkMappingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkMappingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkMappingListResponse struct {
	NetworkMappings []DNSNetworkMapping `json:"network_mappings"`
	TotalAmount     int64               `json:"total_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkMappings respjson.Field
		TotalAmount     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkMappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkMappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkMappingDeleteResponse = any

type NetworkMappingImportResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkMappingImportResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkMappingImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkMappingReplaceResponse = any

type NetworkMappingNewParams struct {
	DNSNetworkMapping DNSNetworkMappingParam
	paramObj
}

func (r NetworkMappingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DNSNetworkMapping)
}
func (r *NetworkMappingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DNSNetworkMapping)
}

type NetworkMappingListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Amount of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Field name to sort by
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Ascending or descending order
	//
	// Any of "asc", "desc".
	OrderDirection NetworkMappingListParamsOrderDirection `query:"order_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkMappingListParams]'s query parameters as
// `url.Values`.
func (r NetworkMappingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ascending or descending order
type NetworkMappingListParamsOrderDirection string

const (
	NetworkMappingListParamsOrderDirectionAsc  NetworkMappingListParamsOrderDirection = "asc"
	NetworkMappingListParamsOrderDirectionDesc NetworkMappingListParamsOrderDirection = "desc"
)

type NetworkMappingReplaceParams struct {
	DNSNetworkMapping DNSNetworkMappingParam
	paramObj
}

func (r NetworkMappingReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DNSNetworkMapping)
}
func (r *NetworkMappingReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DNSNetworkMapping)
}
