// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
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
