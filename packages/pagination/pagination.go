// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/cjavdev/believe-go/internal/apijson"
	"github.com/cjavdev/believe-go/internal/requestconfig"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/believe-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type SkipLimitPage[T any] struct {
	Data  []T   `json:"data"`
	Total int64 `json:"total"`
	Skip  int64 `json:"skip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Total       respjson.Field
		Skip        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r SkipLimitPage[T]) RawJSON() string { return r.JSON.raw }
func (r *SkipLimitPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SkipLimitPage[T]) GetNextPage() (res *SkipLimitPage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	offset := r.Skip
	length := int64(len(r.Data))
	next := offset + length

	if next < r.Total && next != 0 {
		err = cfg.Apply(option.WithQuery("skip", strconv.FormatInt(next, 10)))
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

func (r *SkipLimitPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SkipLimitPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SkipLimitPageAutoPager[T any] struct {
	page *SkipLimitPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSkipLimitPageAutoPager[T any](page *SkipLimitPage[T], err error) *SkipLimitPageAutoPager[T] {
	return &SkipLimitPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SkipLimitPageAutoPager[T]) Next() bool {
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

func (r *SkipLimitPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SkipLimitPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SkipLimitPageAutoPager[T]) Index() int {
	return r.run
}
