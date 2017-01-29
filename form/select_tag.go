package form

import (
	"html/template"
	"reflect"
	"strings"

	"github.com/markbates/tags"
)

type SelectTag struct {
	*tags.BlockTag
	SelectedValue interface{}
	SelectOptions SelectOptions
}

func (s SelectTag) String() string {
	so := make([]string, len(s.SelectOptions))
	for _, x := range s.SelectOptions {
		x.SelectedValue = s.SelectedValue
		so = append(so, x.String())
	}
	s.BlockTag.Body = strings.Join(so, "")
	return s.BlockTag.String()
}

func (s SelectTag) HTML() template.HTML {
	return template.HTML(s.String())
}

func NewSelectTag(opts tags.Options) *SelectTag {
	so := parseSelectOptions(opts)
	selected := opts["selected"]
	delete(opts, "selected")

	st := &SelectTag{
		BlockTag:      tags.NewBlockTag("select", opts),
		SelectOptions: so,
		SelectedValue: selected,
	}
	return st
}

func (f Form) SelectTag(opts tags.Options) *SelectTag {
	return NewSelectTag(opts)
}

func parseSelectOptions(opts tags.Options) SelectOptions {
	if opts["options"] == nil {
		return SelectOptions{}
	}

	sopts := opts["options"]
	delete(opts, "options")

	if x, ok := sopts.(SelectOptions); ok {
		return x
	}

	rv := reflect.ValueOf(sopts)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	so := SelectOptions{}
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			x := rv.Index(i).Interface()
			so = append(so, SelectOption{Value: x, Label: x})
		}
	case reflect.Map:
		keys := rv.MapKeys()
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			so = append(so, SelectOption{
				Value: rv.MapIndex(k).Interface(),
				Label: k.Interface(),
			})
		}
	}
	return so
}
