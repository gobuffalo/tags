package tags

import (
	"bytes"
	"html/template"
)

type Tag struct {
	Name     string
	Options  Options
	Selected bool
	Checked  bool
}

func (t Tag) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("<")
	bb.WriteString(t.Name)
	if len(t.Options) > 0 {
		bb.WriteString(" ")
		bb.WriteString(t.Options.String())
	}
	if t.Selected {
		bb.WriteString(" selected")
	}
	if t.Checked {
		bb.WriteString(" checked")
	}
	bb.WriteString(">")
	return bb.String()
}

func (t Tag) HTML() template.HTML {
	return template.HTML(t.String())
}

func New(name string, opts Options) *Tag {
	tag := &Tag{
		Name:    name,
		Options: opts,
	}

	if tag.Options["selected"] != nil {
		tag.Selected = template.HTMLEscaper(opts["value"]) == template.HTMLEscaper(opts["selected"])
		delete(tag.Options, "selected")
	}

	if tag.Options["checked"] != nil {
		tag.Checked = template.HTMLEscaper(opts["value"]) == template.HTMLEscaper(opts["checked"])
		delete(tag.Options, "checked")
	}

	return tag
}
