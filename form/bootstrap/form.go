package bootstrap

import (
	"html/template"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
	"github.com/gobuffalo/velvet"
)

type Form struct {
	*form.Form
}

func New(opts tags.Options) *Form {
	return &Form{form.New(opts)}
}

func (f Form) CheckboxTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.CheckboxTag(o)
	})
}

func (f Form) InputTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.InputTag(o)
	})
}

func (f Form) RadioButton(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.RadioButton(o)
	})
}

func (f Form) SelectTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.SelectTag(o)
	})
}

func (f Form) TextArea(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.TextArea(o)
	})
}

func FormHelper(help velvet.HelperContext) (template.HTML, error) {
	return form.Helper(help, func(opts tags.Options) form.Helperable {
		return New(opts)
	})
}
