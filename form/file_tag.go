package form

import "github.com/gobuffalo/tags"

func (f Form) FileTag(opts tags.Options) *tags.Tag {
	if opts["type"] == nil {
		opts["type"] = "file"
	}
	if opts["type"] == "file" {
		f.Options["enctype"] = "multipart/form-data"
	}
	return f.InputTag(opts)
}
