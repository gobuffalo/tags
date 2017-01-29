package tags

import (
	"bytes"
	"fmt"
	"html/template"
)

type BlockTag struct {
	*Tag
	Body interface{}
}

func (b BlockTag) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString(b.Tag.String())
	if b.Body != nil {
		bb.WriteString(fmt.Sprintf("%s", b.Body))
	}
	bb.WriteString("</")
	bb.WriteString(b.Name)
	bb.WriteString(">")
	return bb.String()
}

func (b BlockTag) HTML() template.HTML {
	return template.HTML(b.String())
}

func NewBlockTag(name string, opts Options) *BlockTag {
	tag := &BlockTag{
		Tag:  New(name, opts),
		Body: opts["body"],
	}
	delete(tag.Options, "body")
	if tag.Body == nil {
		tag.Body = opts["value"]
		delete(tag.Options, "value")
	}

	return tag
}
