package tags

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
)

type Paginator struct {
	// Current page you're on
	Page int `json:"page"`
	// Number of results you want per page
	PerPage int `json:"per_page"`
	// Page * PerPage (ex: 2 * 20, Offset == 40)
	Offset int `json:"offset"`
	// Total potential records matching the query
	TotalEntriesSize int `json:"total_entries_size"`
	// Total records returns, will be <= PerPage
	CurrentEntriesSize int `json:"current_entries_size"`
	// Total pages
	TotalPages int `json:"total_pages"`
}

func (pagination Paginator) Tag(opts Options) (*Tag, error) {
	// return an empty div if there is only 1 page
	if pagination.TotalPages <= 1 {
		return New("div", Options{}), nil
	}
	var path string
	if p, ok := opts["path"]; ok {
		path = p.(string)
		delete(opts, "path")
	}
	if _, ok := opts["class"]; !ok {
		opts["class"] = ""
	}
	opts["class"] = strings.Join([]string{opts["class"].(string), "pagination"}, " ")
	t := New("ul", opts)

	showPrev := true
	if b, ok := opts["showPrev"].(bool); ok {
		showPrev = b
		delete(opts, "showPrev")
	}
	if showPrev {
		page := pagination.Page - 1
		li, err := pageLI("&laquo;", page, path, pagination)

		if err != nil {
			return t, errors.WithStack(err)
		}
		t.Append(li)
	}

	for i := 1; i < pagination.TotalPages+1; i++ {
		li, err := pageLI(strconv.Itoa(i), i, path, pagination)
		if err != nil {
			return t, errors.WithStack(err)
		}
		t.Append(li)
	}

	showNext := true
	if b, ok := opts["showNext"].(bool); ok {
		showNext = b
		delete(opts, "showNext")
	}
	if showNext {
		page := pagination.Page + 1
		li, err := pageLI("&raquo;", page, path, pagination)

		if err != nil {
			return t, errors.WithStack(err)
		}
		t.Append(li)
	}

	return t, nil
}

func Pagination(pagination interface{}, opts Options) (*Tag, error) {
	if p, ok := pagination.(Paginator); ok {
		return p.Tag(opts)
	}
	rv := reflect.ValueOf(pagination)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil, errors.Errorf("can't build a Paginator from %T", pagination)
	}

	s := structs.New(rv.Interface())

	p := Paginator{
		Page:    1,
		PerPage: 20,
	}

	if f, ok := s.FieldOk("Page"); ok {
		p.Page = f.Value().(int)
	}

	if f, ok := s.FieldOk("PerPage"); ok {
		p.PerPage = f.Value().(int)
	}

	if f, ok := s.FieldOk("Offset"); ok {
		p.Offset = f.Value().(int)
	}

	if f, ok := s.FieldOk("TotalEntriesSize"); ok {
		p.TotalEntriesSize = f.Value().(int)
	}

	if f, ok := s.FieldOk("TotalEntriesSize"); ok {
		p.TotalEntriesSize = f.Value().(int)
	}

	if f, ok := s.FieldOk("CurrentEntriesSize"); ok {
		p.CurrentEntriesSize = f.Value().(int)
	}

	if f, ok := s.FieldOk("TotalPages"); ok {
		p.TotalPages = f.Value().(int)
	}

	return p.Tag(opts)
}

func pageLI(text string, page int, path string, pagination Paginator) (*Tag, error) {

	lio := Options{}
	if page == pagination.Page {
		lio["class"] = "active"
	}
	li := New("li", lio)
	if page == 0 || page > pagination.TotalPages {
		li.Options["class"] = "disabled"
		li.Append(New("span", Options{
			"body": text,
		}))
		return li, nil
	}

	u, err := url.Parse(path)
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()
	ao := Options{
		"href": u.String(),
	}
	a := New("a", ao)
	a.Append(text)
	li.Append(a)
	if err != nil {
		return li, errors.WithStack(err)
	}
	return li, nil
}
