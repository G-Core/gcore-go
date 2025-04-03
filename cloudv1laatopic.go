// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LaaTopicService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LaaTopicService] method instead.
type CloudV1LaaTopicService struct {
	Options []option.RequestOption
}

// NewCloudV1LaaTopicService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1LaaTopicService(opts ...option.RequestOption) (r *CloudV1LaaTopicService) {
	r = &CloudV1LaaTopicService{}
	r.Options = opts
	return
}

// Create LaaS Kafka topic within client namespace
func (r *CloudV1LaaTopicService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1LaaTopicNewParams, opts ...option.RequestOption) (res *LaasKafkaTopic, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/topics", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List LaaS Kafka topics within client namespace
func (r *CloudV1LaaTopicService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *LaasKafkaTopic, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/topics", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete LaaS Kafka topic within client namespace
func (r *CloudV1LaaTopicService) Delete(ctx context.Context, projectID int64, regionID int64, topicName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if topicName == "" {
		err = errors.New("missing required topic_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/topics/%s", projectID, regionID, topicName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LaasIndexRetentionPolicyPydantic struct {
	// Duration of days for which logs must be kept.
	Period int64                                `json:"period,required,nullable"`
	JSON   laasIndexRetentionPolicyPydanticJSON `json:"-"`
}

// laasIndexRetentionPolicyPydanticJSON contains the JSON metadata for the struct
// [LaasIndexRetentionPolicyPydantic]
type laasIndexRetentionPolicyPydanticJSON struct {
	Period      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LaasIndexRetentionPolicyPydantic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r laasIndexRetentionPolicyPydanticJSON) RawJSON() string {
	return r.raw
}

type LaasIndexRetentionPolicyPydanticParam struct {
	// Duration of days for which logs must be kept.
	Period param.Field[int64] `json:"period,required"`
}

func (r LaasIndexRetentionPolicyPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LaasKafkaTopic struct {
	// Kafka topic name to be used in log shipper config to write logs to LaaS
	Name string `json:"name,required"`
	// Index Pattern data
	IndexPattern    LaasKafkaTopicIndexPattern       `json:"index_pattern,nullable"`
	RetentionPolicy LaasIndexRetentionPolicyPydantic `json:"retention_policy,nullable"`
	// Size of the topic in bytes
	SizeInBytes int64              `json:"size_in_bytes"`
	JSON        laasKafkaTopicJSON `json:"-"`
}

// laasKafkaTopicJSON contains the JSON metadata for the struct [LaasKafkaTopic]
type laasKafkaTopicJSON struct {
	Name            apijson.Field
	IndexPattern    apijson.Field
	RetentionPolicy apijson.Field
	SizeInBytes     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LaasKafkaTopic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r laasKafkaTopicJSON) RawJSON() string {
	return r.raw
}

// Index Pattern data
type LaasKafkaTopicIndexPattern struct {
	// Pattern ID
	ID   string                         `json:"id,required"`
	JSON laasKafkaTopicIndexPatternJSON `json:"-"`
}

// laasKafkaTopicIndexPatternJSON contains the JSON metadata for the struct
// [LaasKafkaTopicIndexPattern]
type laasKafkaTopicIndexPatternJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LaasKafkaTopicIndexPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r laasKafkaTopicIndexPatternJSON) RawJSON() string {
	return r.raw
}

type LaasKafkaTopicParam struct {
	// Kafka topic name to be used in log shipper config to write logs to LaaS
	Name param.Field[string] `json:"name,required"`
	// Index Pattern data
	IndexPattern    param.Field[LaasKafkaTopicIndexPatternParam]       `json:"index_pattern"`
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// Size of the topic in bytes
	SizeInBytes param.Field[int64] `json:"size_in_bytes"`
}

func (r LaasKafkaTopicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Index Pattern data
type LaasKafkaTopicIndexPatternParam struct {
	// Pattern ID
	ID param.Field[string] `json:"id,required"`
}

func (r LaasKafkaTopicIndexPatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LaaTopicNewParams struct {
	LaasKafkaTopic LaasKafkaTopicParam `json:"laas_kafka_topic,required"`
}

func (r CloudV1LaaTopicNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LaasKafkaTopic)
}
