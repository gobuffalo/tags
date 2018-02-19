package bootstrap

import (
	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
)

//Form is the bootstrap extension of Form
type Form struct {
	*form.Form
}

//New builds a new bootstrap form from passed options
func New(opts tags.Options) *Form {
	return &Form{form.New(opts)}
}

//CheckboxTag adds an input type=checkbox
func (f Form) CheckboxTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.CheckboxTag(o)
	})
}

//InputTag adds an input type=text by default
func (f Form) InputTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.InputTag(o)
	})
}

//FileTag adds an input type=file to the form
func (f Form) FileTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.FileTag(o)
	})
}

//RadioButton adds an input type=radio to the form
func (f Form) RadioButton(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.RadioButton(o)
	})
}

//SelectTag adds a select tag to the form
func (f Form) SelectTag(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.SelectTag(o)
	})
}

//TextArea adds a textarea tag to the form
func (f Form) TextArea(opts tags.Options) *tags.Tag {
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.Form.TextArea(o)
	})
}

//HiddenTag adds a hidden input to the form
func (f Form) HiddenTag(opts tags.Options) *tags.Tag {
	return f.Form.HiddenTag(opts)
}
