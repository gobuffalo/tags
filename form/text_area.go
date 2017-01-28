package form

import "github.com/markbates/tags"

func (f Form) TextArea(opts tags.Options) *tags.BlockTag {
	return tags.NewBlockTag("textarea", opts)
}
