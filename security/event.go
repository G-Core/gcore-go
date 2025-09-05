// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// EventService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Event Logs Clients View
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ClientView], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "security/notifier/v1/event_logs"
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

// Event Logs Clients View
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ClientView] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type ClientView struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "ddos_alert", "rtbh_alert".
	AlertType                  ClientViewAlertType `json:"alert_type,nullable"`
	AttackPowerBps             float64             `json:"attack_power_bps,nullable"`
	AttackPowerPps             float64             `json:"attack_power_pps,nullable"`
	AttackStartTime            time.Time           `json:"attack_start_time,nullable" format:"date-time"`
	ClientID                   int64               `json:"client_id,nullable"`
	NotificationType           string              `json:"notification_type,nullable"`
	NumberOfIPInvolvedInAttack int64               `json:"number_of_ip_involved_in_attack,nullable"`
	TargetedIPAddresses        string              `json:"targeted_ip_addresses,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AlertType                  respjson.Field
		AttackPowerBps             respjson.Field
		AttackPowerPps             respjson.Field
		AttackStartTime            respjson.Field
		ClientID                   respjson.Field
		NotificationType           respjson.Field
		NumberOfIPInvolvedInAttack respjson.Field
		TargetedIPAddresses        respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientView) RawJSON() string { return r.JSON.raw }
func (r *ClientView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientViewAlertType string

const (
	ClientViewAlertTypeDDOSAlert ClientViewAlertType = "ddos_alert"
	ClientViewAlertTypeRtbhAlert ClientViewAlertType = "rtbh_alert"
)

type EventListParams struct {
	TargetedIPAddresses param.Opt[string] `query:"targeted_ip_addresses,omitzero" json:"-"`
	Limit               param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset              param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	// Any of "ddos_alert", "rtbh_alert".
	AlertType EventListParamsAlertType     `query:"alert_type,omitzero" json:"-"`
	DateFrom  EventListParamsDateFromUnion `query:"date_from,omitzero" format:"date-time" json:"-"`
	DateTo    EventListParamsDateToUnion   `query:"date_to,omitzero" format:"date-time" json:"-"`
	// Any of "attack_start_time", "-attack_start_time", "attack_power_bps",
	// "-attack_power_bps", "attack_power_pps", "-attack_power_pps",
	// "number_of_ip_involved_in_attack", "-number_of_ip_involved_in_attack",
	// "alert_type", "-alert_type".
	Ordering EventListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EventListParamsAlertType string

const (
	EventListParamsAlertTypeDDOSAlert EventListParamsAlertType = "ddos_alert"
	EventListParamsAlertTypeRtbhAlert EventListParamsAlertType = "rtbh_alert"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventListParamsDateFromUnion struct {
	OfTime   param.Opt[time.Time] `query:",omitzero,inline"`
	OfString param.Opt[string]    `query:",omitzero,inline"`
	paramUnion
}

func (u *EventListParamsDateFromUnion) asAny() any {
	if !param.IsOmitted(u.OfTime) {
		return &u.OfTime.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventListParamsDateToUnion struct {
	OfTime   param.Opt[time.Time] `query:",omitzero,inline"`
	OfString param.Opt[string]    `query:",omitzero,inline"`
	paramUnion
}

func (u *EventListParamsDateToUnion) asAny() any {
	if !param.IsOmitted(u.OfTime) {
		return &u.OfTime.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type EventListParamsOrdering string

const (
	EventListParamsOrderingAttackStartTime                 EventListParamsOrdering = "attack_start_time"
	EventListParamsOrderingMinusAttackStartTime            EventListParamsOrdering = "-attack_start_time"
	EventListParamsOrderingAttackPowerBps                  EventListParamsOrdering = "attack_power_bps"
	EventListParamsOrderingMinusAttackPowerBps             EventListParamsOrdering = "-attack_power_bps"
	EventListParamsOrderingAttackPowerPps                  EventListParamsOrdering = "attack_power_pps"
	EventListParamsOrderingMinusAttackPowerPps             EventListParamsOrdering = "-attack_power_pps"
	EventListParamsOrderingNumberOfIPInvolvedInAttack      EventListParamsOrdering = "number_of_ip_involved_in_attack"
	EventListParamsOrderingMinusNumberOfIPInvolvedInAttack EventListParamsOrdering = "-number_of_ip_involved_in_attack"
	EventListParamsOrderingAlertType                       EventListParamsOrdering = "alert_type"
	EventListParamsOrderingMinusAlertType                  EventListParamsOrdering = "-alert_type"
)
