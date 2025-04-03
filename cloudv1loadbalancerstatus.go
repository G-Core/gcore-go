// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LoadbalancerStatusService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LoadbalancerStatusService] method instead.
type CloudV1LoadbalancerStatusService struct {
	Options []option.RequestOption
}

// NewCloudV1LoadbalancerStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1LoadbalancerStatusService(opts ...option.RequestOption) (r *CloudV1LoadbalancerStatusService) {
	r = &CloudV1LoadbalancerStatusService{}
	r.Options = opts
	return
}

// Get load balancer status
func (r *CloudV1LoadbalancerStatusService) Get(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, opts ...option.RequestOption) (res *LoadBalancerStatus, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/status", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Listeners of the Load Balancer
	Listeners []LoadBalancerStatusListener `json:"listeners,required"`
	// Name of the load balancer
	Name string `json:"name,required"`
	// Operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Metadata of the Load Balancer
	Metadata []DetailedMetadata     `json:"metadata"`
	JSON     loadBalancerStatusJSON `json:"-"`
}

// loadBalancerStatusJSON contains the JSON metadata for the struct
// [LoadBalancerStatus]
type loadBalancerStatusJSON struct {
	ID                 apijson.Field
	Listeners          apijson.Field
	Name               apijson.Field
	OperatingStatus    apijson.Field
	ProvisioningStatus apijson.Field
	Metadata           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerStatusJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerStatusListener struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Name of the load balancer listener
	Name string `json:"name,required"`
	// Operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Pools of the Listeners
	Pools []LoadBalancerStatusListenersPool `json:"pools,required"`
	// Provisioning status of the entity
	ProvisioningStatus ProvisioningStatus             `json:"provisioning_status,required"`
	JSON               loadBalancerStatusListenerJSON `json:"-"`
}

// loadBalancerStatusListenerJSON contains the JSON metadata for the struct
// [LoadBalancerStatusListener]
type loadBalancerStatusListenerJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	OperatingStatus    apijson.Field
	Pools              apijson.Field
	ProvisioningStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerStatusListener) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerStatusListenerJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerStatusListenersPool struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Members (servers) of the pool
	Members []LoadBalancerStatusListenersPoolsMember `json:"members,required"`
	// Name of the load balancer pool
	Name string `json:"name,required"`
	// Operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Health Monitor of the Pool
	HealthMonitor LoadBalancerStatusListenersPoolsHealthMonitor `json:"health_monitor"`
	JSON          loadBalancerStatusListenersPoolJSON           `json:"-"`
}

// loadBalancerStatusListenersPoolJSON contains the JSON metadata for the struct
// [LoadBalancerStatusListenersPool]
type loadBalancerStatusListenersPoolJSON struct {
	ID                 apijson.Field
	Members            apijson.Field
	Name               apijson.Field
	OperatingStatus    apijson.Field
	ProvisioningStatus apijson.Field
	HealthMonitor      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerStatusListenersPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerStatusListenersPoolJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerStatusListenersPoolsMember struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Address of the member (server)
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Port of the member (server)
	ProtocolPort int64 `json:"protocol_port,required"`
	// Provisioning status of the entity
	ProvisioningStatus ProvisioningStatus                         `json:"provisioning_status,required"`
	JSON               loadBalancerStatusListenersPoolsMemberJSON `json:"-"`
}

// loadBalancerStatusListenersPoolsMemberJSON contains the JSON metadata for the
// struct [LoadBalancerStatusListenersPoolsMember]
type loadBalancerStatusListenersPoolsMemberJSON struct {
	ID                 apijson.Field
	Address            apijson.Field
	OperatingStatus    apijson.Field
	ProtocolPort       apijson.Field
	ProvisioningStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerStatusListenersPoolsMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerStatusListenersPoolsMemberJSON) RawJSON() string {
	return r.raw
}

// Health Monitor of the Pool
type LoadBalancerStatusListenersPoolsHealthMonitor struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Type of the Health Monitor
	Type HealthMonitorType                                 `json:"type,required"`
	JSON loadBalancerStatusListenersPoolsHealthMonitorJSON `json:"-"`
}

// loadBalancerStatusListenersPoolsHealthMonitorJSON contains the JSON metadata for
// the struct [LoadBalancerStatusListenersPoolsHealthMonitor]
type loadBalancerStatusListenersPoolsHealthMonitorJSON struct {
	ID                 apijson.Field
	OperatingStatus    apijson.Field
	ProvisioningStatus apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerStatusListenersPoolsHealthMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerStatusListenersPoolsHealthMonitorJSON) RawJSON() string {
	return r.raw
}
