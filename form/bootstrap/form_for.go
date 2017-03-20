package bootstrap

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
)

type FormFor struct {
	*form.FormFor
}

func (f FormFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {

	label := field
	if opts["label"] != nil {
		label = fmt.Sprint(opts["label"])
	}
	hl := opts["hide_label"]
	delete(opts, "label")

	if err := f.Errors.Get(field); err != nil {
		opts["errors"] = err
	}

	return divWrapper(opts, func(o tags.Options) tags.Body {
		if o["class"] != nil {
			cls := strings.Split(o["class"].(string), " ")
			ncls := make([]string, 0, len(cls))
			for _, c := range cls {
				if c != "form-control" {
					ncls = append(ncls, c)
				}
			}
			o["class"] = strings.Join(ncls, " ")
		}
		if label != "" {
			o["label"] = label
		}
		if hl != nil {
			o["hide_label"] = hl
		}
		return f.FormFor.CheckboxTag(field, o)
	})
}

func (f FormFor) InputTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.InputTag(field, opts)
	})
}

func (f FormFor) RadioButton(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.RadioButton(field, opts)
	})
}

func (f FormFor) SelectTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.SelectTag(field, opts)
	})
}

func (f FormFor) TextArea(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)
	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.TextArea(field, opts)
	})
}

func (f FormFor) buildOptions(field string, opts tags.Options) tags.Options {
	opts["tags-field"] = field

	if err := f.Errors.Get(field); err != nil {
		opts["errors"] = err
	}

	return opts
}

func NewFormFor(model interface{}, opts tags.Options) *FormFor {
	return &FormFor{form.NewFormFor(model, opts)}
}
