package form

import (
	"fmt"
	"reflect"

	"github.com/gobuffalo/tags"
	"github.com/markbates/inflect"
	"github.com/markbates/pop/nulls"
)

type FormFor struct {
	*Form
	Model      interface{}
	name       string
	dashedName string
	reflection reflect.Value
}

func NewFormFor(model interface{}, opts tags.Options) *FormFor {
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	name := rv.Type().Name()
	dashedName := inflect.Dasherize(name)

	if opts["id"] == nil {
		opts["id"] = fmt.Sprintf("%s-form", dashedName)
	}

	return &FormFor{
		Form:       New(opts),
		Model:      model,
		name:       name,
		dashedName: dashedName,
		reflection: rv,
	}
}

func (f FormFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.Form.CheckboxTag(opts)
}

func (f FormFor) InputTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.Form.InputTag(opts)
}

func (f FormFor) RadioButton(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.Form.RadioButton(opts)
}

func (f FormFor) SelectTag(field string, opts tags.Options) *SelectTag {
	f.buildOptions(field, opts)
	return f.Form.SelectTag(opts)
}

func (f FormFor) TextArea(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.Form.TextArea(opts)
}

func (f FormFor) buildOptions(field string, opts tags.Options) {
	opts["value"] = f.value(field)
	if opts["name"] == nil {
		opts["name"] = field
	}
	if opts["id"] == nil {
		opts["id"] = fmt.Sprintf("%s-%s", f.dashedName, field)
	}
}

func (f FormFor) value(field string) interface{} {
	fn := f.reflection.FieldByName(field)

	if fn.IsValid() == false {
		return ""
	}

	i := fn.Interface()
	if dv, ok := i.(nulls.String); ok {
		return dv.String
	}
	return i
}
