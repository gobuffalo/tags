package form

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/velvet"
)

func FormHelper(help velvet.HelperContext) (template.HTML, error) {
	return Helper(help, func(opts tags.Options) Helperable {
		return New(opts)
	})
}

func FormForHelper(model interface{}, help velvet.HelperContext) (template.HTML, error) {
	return Helper(help, func(opts tags.Options) Helperable {
		return NewFormFor(model, opts)
	})
}

type Helperable interface {
	SetAuthenticityToken(string)
	Append(...tags.Body)
	velvet.HTMLer
}

func Helper(help velvet.HelperContext, fn func(opts tags.Options) Helperable) (template.HTML, error) {
	opts := help.Context.Options()
	form := fn(opts)
	if help.Get("authenticity_token") != nil {
		form.SetAuthenticityToken(fmt.Sprint(help.Get("authenticity_token")))
	}
	ctx := help.Context.New()
	ctx.Set("f", form)
	s, err := help.BlockWith(ctx)
	if err != nil {
		return "", err
	}
	form.Append(s)
	return form.HTML(), nil
}
