// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type OffsetPage[T any] struct {
	Results []T   `json:"results"`
	Count   int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPage[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPage[T]) GetNextPage() (res *OffsetPage[T], err error) {
	if len(r.Results) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Results))
	next := offset + length

	if next < r.Count && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageAutoPager[T any] struct {
	page *OffsetPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageAutoPager[T any](page *OffsetPage[T], err error) *OffsetPageAutoPager[T] {
	return &OffsetPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Results) == 0 {
		return false
	}
	if r.idx >= len(r.page.Results) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Results) == 0 {
			return false
		}
	}
	r.cur = r.page.Results[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageAutoPager[T]) Index() int {
	return r.run
}

type OffsetPageFastedgeApps[T any] struct {
	Apps  []T   `json:"apps"`
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Apps        respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPageFastedgeApps[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageFastedgeApps[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPageFastedgeApps[T]) GetNextPage() (res *OffsetPageFastedgeApps[T], err error) {
	if len(r.Apps) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Apps))
	next := offset + length

	if next < r.Count && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPageFastedgeApps[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPageFastedgeApps[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageFastedgeAppsAutoPager[T any] struct {
	page *OffsetPageFastedgeApps[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageFastedgeAppsAutoPager[T any](page *OffsetPageFastedgeApps[T], err error) *OffsetPageFastedgeAppsAutoPager[T] {
	return &OffsetPageFastedgeAppsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageFastedgeAppsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Apps) == 0 {
		return false
	}
	if r.idx >= len(r.page.Apps) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Apps) == 0 {
			return false
		}
	}
	r.cur = r.page.Apps[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageFastedgeAppsAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageFastedgeAppsAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageFastedgeAppsAutoPager[T]) Index() int {
	return r.run
}

type OffsetPageFastedgeTemplates[T any] struct {
	Templates []T   `json:"templates"`
	Count     int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Templates   respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPageFastedgeTemplates[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageFastedgeTemplates[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPageFastedgeTemplates[T]) GetNextPage() (res *OffsetPageFastedgeTemplates[T], err error) {
	if len(r.Templates) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Templates))
	next := offset + length

	if next < r.Count && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPageFastedgeTemplates[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPageFastedgeTemplates[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageFastedgeTemplatesAutoPager[T any] struct {
	page *OffsetPageFastedgeTemplates[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageFastedgeTemplatesAutoPager[T any](page *OffsetPageFastedgeTemplates[T], err error) *OffsetPageFastedgeTemplatesAutoPager[T] {
	return &OffsetPageFastedgeTemplatesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageFastedgeTemplatesAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Templates) == 0 {
		return false
	}
	if r.idx >= len(r.page.Templates) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Templates) == 0 {
			return false
		}
	}
	r.cur = r.page.Templates[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageFastedgeTemplatesAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageFastedgeTemplatesAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageFastedgeTemplatesAutoPager[T]) Index() int {
	return r.run
}

type OffsetPageFastedgeAppLogs[T any] struct {
	Logs       []T   `json:"logs"`
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logs        respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPageFastedgeAppLogs[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageFastedgeAppLogs[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPageFastedgeAppLogs[T]) GetNextPage() (res *OffsetPageFastedgeAppLogs[T], err error) {
	if len(r.Logs) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Logs))
	next := offset + length

	if next < r.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPageFastedgeAppLogs[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPageFastedgeAppLogs[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageFastedgeAppLogsAutoPager[T any] struct {
	page *OffsetPageFastedgeAppLogs[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageFastedgeAppLogsAutoPager[T any](page *OffsetPageFastedgeAppLogs[T], err error) *OffsetPageFastedgeAppLogsAutoPager[T] {
	return &OffsetPageFastedgeAppLogsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageFastedgeAppLogsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Logs) == 0 {
		return false
	}
	if r.idx >= len(r.page.Logs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Logs) == 0 {
			return false
		}
	}
	r.cur = r.page.Logs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageFastedgeAppLogsAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageFastedgeAppLogsAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageFastedgeAppLogsAutoPager[T]) Index() int {
	return r.run
}

type PageStreamingAI[T any] struct {
	Results []T `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageStreamingAI[T]) RawJSON() string { return r.JSON.raw }
func (r *PageStreamingAI[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageStreamingAI[T]) GetNextPage() (res *PageStreamingAI[T], err error) {
	if len(r.Results) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageStreamingAI[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageStreamingAI[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageStreamingAIAutoPager[T any] struct {
	page *PageStreamingAI[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageStreamingAIAutoPager[T any](page *PageStreamingAI[T], err error) *PageStreamingAIAutoPager[T] {
	return &PageStreamingAIAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageStreamingAIAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Results) == 0 {
		return false
	}
	if r.idx >= len(r.page.Results) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Results) == 0 {
			return false
		}
	}
	r.cur = r.page.Results[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageStreamingAIAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageStreamingAIAutoPager[T]) Err() error {
	return r.err
}

func (r *PageStreamingAIAutoPager[T]) Index() int {
	return r.run
}

type PageStreaming[T any] struct {
	Items []T `json:",inline"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageStreaming[T]) RawJSON() string { return r.JSON.raw }
func (r *PageStreaming[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageStreaming[T]) GetNextPage() (res *PageStreaming[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageStreaming[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageStreaming[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageStreamingAutoPager[T any] struct {
	page *PageStreaming[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageStreamingAutoPager[T any](page *PageStreaming[T], err error) *PageStreamingAutoPager[T] {
	return &PageStreamingAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageStreamingAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageStreamingAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageStreamingAutoPager[T]) Err() error {
	return r.err
}

func (r *PageStreamingAutoPager[T]) Index() int {
	return r.run
}

type OffsetPageCDN[T any] struct {
	Items []T `json:",inline"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPageCDN[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageCDN[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPageCDN[T]) GetNextPage() (res *OffsetPageCDN[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Items))
	next := offset + length

	if length > 0 && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPageCDN[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPageCDN[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageCDNAutoPager[T any] struct {
	page *OffsetPageCDN[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageCDNAutoPager[T any](page *OffsetPageCDN[T], err error) *OffsetPageCDNAutoPager[T] {
	return &OffsetPageCDNAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageCDNAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageCDNAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageCDNAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageCDNAutoPager[T]) Index() int {
	return r.run
}

type OffsetPageCDNLogsMeta struct {
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OffsetPageCDNLogsMeta) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageCDNLogsMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OffsetPageCDNLogs[T any] struct {
	Data []T                   `json:"data"`
	Meta OffsetPageCDNLogsMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPageCDNLogs[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPageCDNLogs[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPageCDNLogs[T]) GetNextPage() (res *OffsetPageCDNLogs[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Data))
	next := offset + length

	if next < r.Meta.Count && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPageCDNLogs[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPageCDNLogs[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageCDNLogsAutoPager[T any] struct {
	page *OffsetPageCDNLogs[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageCDNLogsAutoPager[T any](page *OffsetPageCDNLogs[T], err error) *OffsetPageCDNLogsAutoPager[T] {
	return &OffsetPageCDNLogsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageCDNLogsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageCDNLogsAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageCDNLogsAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageCDNLogsAutoPager[T]) Index() int {
	return r.run
}
