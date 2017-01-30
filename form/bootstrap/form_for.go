package bootstrap

import (
	"html/template"

	"github.com/gobuffalo/velvet"
	"github.com/markbates/tags"
	"github.com/markbates/tags/form"
)

type FormFor struct {
	*form.FormFor
}

func (f FormFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {
	opts["label"] = field
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.CheckboxTag(field, o)
	})
}

func (f FormFor) InputTag(field string, opts tags.Options) *tags.Tag {
	opts["label"] = field
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.InputTag(field, opts)
	})
}

func (f FormFor) RadioButton(field string, opts tags.Options) *tags.Tag {
	opts["label"] = field
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.RadioButton(field, opts)
	})
}

func (f FormFor) SelectTag(field string, opts tags.Options) *tags.Tag {
	opts["label"] = field
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.SelectTag(field, opts)
	})
}

func (f FormFor) TextArea(field string, opts tags.Options) *tags.Tag {
	opts["label"] = field
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.TextArea(field, opts)
	})
}

func NewFormFor(model interface{}, opts tags.Options) *FormFor {
	return &FormFor{form.NewFormFor(model, opts)}
}

func FormForHelper(model interface{}, help velvet.HelperContext) (template.HTML, error) {
	return form.Helper(help, func(opts tags.Options) form.Helperable {
		return NewFormFor(model, opts)
	})
}
