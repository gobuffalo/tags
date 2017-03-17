package tags

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/velvet"
	"github.com/markbates/pop/nulls"
)

type Body interface{}

type Tag struct {
	Name     string
	Options  Options
	Selected bool
	Checked  bool
	Body     []Body
}

func (t *Tag) Append(b ...Body) {
	t.Body = append(t.Body, b...)
}

func (t *Tag) Prepend(b ...Body) {
	t.Body = append(b, t.Body...)
}

type interfacer interface {
	Interface() interface{}
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
	if len(t.Body) > 0 {
		bb.WriteString(">")
		for _, b := range t.Body {
			switch tb := b.(type) {
			case velvet.HTMLer:
				bb.Write([]byte(tb.HTML()))
			case nulls.String:
				body := b.(nulls.String)
				bb.WriteString(body.String)
			case fmt.Stringer:
				bb.WriteString(tb.String())
			case interfacer:
				bb.WriteString(fmt.Sprint(tb.Interface()))
			default:
				bb.WriteString(fmt.Sprint(tb))
			}
		}
		bb.WriteString("</")
		bb.WriteString(t.Name)
		bb.WriteString(">")
		return bb.String()
	}
	bb.WriteString(" />")
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
	if tag.Options["body"] != nil {
		tag.Body = []Body{tag.Options["body"]}
		delete(tag.Options, "body")
	}

	if tag.Options["selected"] != nil {
		tag.Selected = template.HTMLEscaper(opts["value"]) == template.HTMLEscaper(opts["selected"])
		delete(tag.Options, "selected")
	}

	return tag
}
