package bootstrap

import "github.com/markbates/tags"

func buildOptions(opts tags.Options) {
	if opts["class"] == nil {
		opts["class"] = ""
	}
	opts["class"] = opts["class"].(string) + " form-control"
	delete(opts, "hide_label")
}

func divWrapper(opts tags.Options, fn func(opts tags.Options) tags.Body) *tags.Tag {
	div := tags.New("div", tags.Options{
		"class": "form-group",
	})

	useLabel := opts["hide_label"] == nil
	if useLabel && opts["label"] != nil {
		div.Prepend(tags.New("label", tags.Options{
			"body": opts["label"],
		}))
	}

	buildOptions(opts)

	div.Append(fn(opts))

	return div
}
