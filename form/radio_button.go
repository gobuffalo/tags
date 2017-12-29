package form

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags"
)

func (f Form) RadioButton(opts tags.Options) *tags.Tag {
	opts["type"] = "radio"

	var label string
	if opts["label"] != nil {
		label = fmt.Sprint(opts["label"])
		delete(opts, "label")
	}
	var ID string
	if opts["id"] != nil {
		ID = fmt.Sprint(opts["id"])
	}

	value := opts["value"]
	checked := opts["checked"]
	delete(opts, "checked")
	ct := f.InputTag(opts)
	ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
	labelOptions := tags.Options{
		"body": strings.Join([]string{ct.String(), label}, " "),
	}
	// If the ID is provided, give it to the label's for attribute
	if ID != "" {
		labelOptions["for"] = ID
	}
	tag := tags.New("label", labelOptions)
	return tag
}
