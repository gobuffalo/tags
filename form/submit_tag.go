package form

import "github.com/gobuffalo/tags"

//SubmitTag generates an input tag with type "submit"
func (f Form) SubmitTag(opts tags.Options) *tags.Tag {
	opts["type"] = "submit"
	return tags.New("input", opts)
}
