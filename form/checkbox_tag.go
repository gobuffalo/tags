package form

import (
	"fmt"
	"strings"

	"github.com/markbates/tags"
)

func (f Form) CheckboxTag(opts tags.Options) *tags.BlockTag {
	opts["type"] = "checkbox"

	var label string
	if opts["label"] != nil {
		label = fmt.Sprint(opts["label"])
		delete(opts, "label")
	}

	ct := f.InputTag(opts)
	tag := tags.NewBlockTag("label", tags.Options{
		"body": strings.Join([]string{ct.String(), label}, " "),
	})
	return tag
}
