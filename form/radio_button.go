package form

import (
	"fmt"
	"strings"

	"github.com/markbates/tags"
)

func (f Form) RadioButton(opts tags.Options) *tags.Tag {
	opts["type"] = "radio"

	var label string
	if opts["label"] != nil {
		label = fmt.Sprint(opts["label"])
		delete(opts, "label")
	}

	ct := f.InputTag(opts)
	tag := tags.New("label", tags.Options{
		"body": strings.Join([]string{ct.String(), label}, " "),
	})
	return tag
}
