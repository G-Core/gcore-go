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

// CloudV1LbpoolHealthmonitorService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LbpoolHealthmonitorService] method instead.
type CloudV1LbpoolHealthmonitorService struct {
	Options []option.RequestOption
}

// NewCloudV1LbpoolHealthmonitorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1LbpoolHealthmonitorService(opts ...option.RequestOption) (r *CloudV1LbpoolHealthmonitorService) {
	r = &CloudV1LbpoolHealthmonitorService{}
	r.Options = opts
	return
}

// Create Load Balancer Pool Health Monitor
func (r *CloudV1LbpoolHealthmonitorService) New(ctx context.Context, projectID int64, regionID int64, poolID string, body CloudV1LbpoolHealthmonitorNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/healthmonitor", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete load balancer pool health monitor
func (r *CloudV1LbpoolHealthmonitorService) Delete(ctx context.Context, projectID int64, regionID int64, poolID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/healthmonitor", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CreateLbHealthMonitorParam struct {
	// The time, in seconds, between sending probes to members
	Delay param.Field[int64] `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries param.Field[int64] `json:"max_retries,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout param.Field[int64] `json:"timeout,required"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	Type          param.Field[HealthMonitorType] `json:"type,required"`
	ExpectedCodes param.Field[string]            `json:"expected_codes"`
	// HTTP method
	HTTPMethod param.Field[HTTPMethodHealthmonitor] `json:"http_method"`
	// Number of failures before the member is switched to ERROR state
	MaxRetriesDown param.Field[int64] `json:"max_retries_down"`
	// URL Path. Defaults to '/'
	URLPath param.Field[string] `json:"url_path"`
}

func (r CreateLbHealthMonitorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HealthMonitorType string

const (
	HealthMonitorTypeHTTP       HealthMonitorType = "HTTP"
	HealthMonitorTypeHTTPS      HealthMonitorType = "HTTPS"
	HealthMonitorTypeK8S        HealthMonitorType = "K8S"
	HealthMonitorTypePing       HealthMonitorType = "PING"
	HealthMonitorTypeTcp        HealthMonitorType = "TCP"
	HealthMonitorTypeTlsHello   HealthMonitorType = "TLS-HELLO"
	HealthMonitorTypeUdpConnect HealthMonitorType = "UDP-CONNECT"
)

func (r HealthMonitorType) IsKnown() bool {
	switch r {
	case HealthMonitorTypeHTTP, HealthMonitorTypeHTTPS, HealthMonitorTypeK8S, HealthMonitorTypePing, HealthMonitorTypeTcp, HealthMonitorTypeTlsHello, HealthMonitorTypeUdpConnect:
		return true
	}
	return false
}

type HTTPMethodHealthmonitor string

const (
	HTTPMethodHealthmonitorConnect HTTPMethodHealthmonitor = "CONNECT"
	HTTPMethodHealthmonitorDelete  HTTPMethodHealthmonitor = "DELETE"
	HTTPMethodHealthmonitorGet     HTTPMethodHealthmonitor = "GET"
	HTTPMethodHealthmonitorHead    HTTPMethodHealthmonitor = "HEAD"
	HTTPMethodHealthmonitorOptions HTTPMethodHealthmonitor = "OPTIONS"
	HTTPMethodHealthmonitorPatch   HTTPMethodHealthmonitor = "PATCH"
	HTTPMethodHealthmonitorPost    HTTPMethodHealthmonitor = "POST"
	HTTPMethodHealthmonitorPut     HTTPMethodHealthmonitor = "PUT"
	HTTPMethodHealthmonitorTrace   HTTPMethodHealthmonitor = "TRACE"
)

func (r HTTPMethodHealthmonitor) IsKnown() bool {
	switch r {
	case HTTPMethodHealthmonitorConnect, HTTPMethodHealthmonitorDelete, HTTPMethodHealthmonitorGet, HTTPMethodHealthmonitorHead, HTTPMethodHealthmonitorOptions, HTTPMethodHealthmonitorPatch, HTTPMethodHealthmonitorPost, HTTPMethodHealthmonitorPut, HTTPMethodHealthmonitorTrace:
		return true
	}
	return false
}

type CloudV1LbpoolHealthmonitorNewParams struct {
	CreateLbHealthMonitor CreateLbHealthMonitorParam `json:"create_lb_health_monitor,required"`
}

func (r CloudV1LbpoolHealthmonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLbHealthMonitor)
}
