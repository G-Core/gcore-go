// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// FastedgeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFastedgeService] method instead.
type FastedgeService struct {
	Options    []option.RequestOption
	Templates  TemplateService
	Secrets    SecretService
	Binaries   BinaryService
	Statistics StatisticService
	Apps       AppService
	KvStores   KvStoreService
}

// NewFastedgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFastedgeService(opts ...option.RequestOption) (r FastedgeService) {
	r = FastedgeService{}
	r.Options = opts
	r.Templates = NewTemplateService(opts...)
	r.Secrets = NewSecretService(opts...)
	r.Binaries = NewBinaryService(opts...)
	r.Statistics = NewStatisticService(opts...)
	r.Apps = NewAppService(opts...)
	r.KvStores = NewKvStoreService(opts...)
	return
}

// Get status and limits for the client
func (r *FastedgeService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *Client, err error) {
	opts = append(r.Options[:], opts...)
	path := "fastedge/v1/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Client struct {
	// Actual allowed number of apps
	AppCount int64 `json:"app_count,required"`
	// Max allowed number of apps
	AppLimit int64 `json:"app_limit,required"`
	// Actual number of calls for all apps during the current day (UTC)
	DailyConsumption int64 `json:"daily_consumption,required"`
	// Max allowed calls for all apps during a day (UTC)
	DailyLimit int64 `json:"daily_limit,required"`
	// Actual number of calls for all apps during the current hour
	HourlyConsumption int64 `json:"hourly_consumption,required"`
	// Max allowed calls for all apps during an hour
	HourlyLimit int64 `json:"hourly_limit,required"`
	// Actual number of calls for all apps during the current calendar month (UTC)
	MonthlyConsumption int64 `json:"monthly_consumption,required"`
	// List of enabled networks
	Networks []ClientNetwork `json:"networks,required"`
	// Plan ID
	PlanID int64 `json:"plan_id,required"`
	// Status code:
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status int64 `json:"status,required"`
	// Plan name
	Plan string `json:"plan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppCount           respjson.Field
		AppLimit           respjson.Field
		DailyConsumption   respjson.Field
		DailyLimit         respjson.Field
		HourlyConsumption  respjson.Field
		HourlyLimit        respjson.Field
		MonthlyConsumption respjson.Field
		Networks           respjson.Field
		PlanID             respjson.Field
		Status             respjson.Field
		Plan               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Client) RawJSON() string { return r.JSON.raw }
func (r *Client) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientNetwork struct {
	// Is network is default
	IsDefault bool `json:"is_default,required"`
	// Network name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsDefault   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientNetwork) RawJSON() string { return r.JSON.raw }
func (r *ClientNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
