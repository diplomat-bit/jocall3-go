// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/diplomat-bit/jocall3-go/internal/apijson"
	"github.com/diplomat-bit/jocall3-go/internal/requestconfig"
	"github.com/diplomat-bit/jocall3-go/option"
)

type Page[T any] struct {
	Items      interface{} `json:"items"`
	Total      interface{} `json:"total"`
	NextOffset interface{} `json:"next_offset"`
	JSON       pageJSON    `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// pageJSON contains the JSON metadata for the struct [Page[T]]
type pageJSON struct {
	Items       apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	if len(r.data) == 0 {
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

func (r *Page[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Page[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}
