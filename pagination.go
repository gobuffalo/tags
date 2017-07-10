package tags

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

func Pagination(pagination *pop.Paginator, opts Options) (*Tag, error) {
	// return an empty div if there is only 1 page
	if pagination.TotalPages == 1 {
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

func pageLI(text string, page int, path string, pagination *pop.Paginator) (*Tag, error) {

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
