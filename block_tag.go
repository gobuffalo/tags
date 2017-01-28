package tags

import (
	"bytes"
	"fmt"
)

type BlockTag struct {
	*Tag
	Body interface{}
}

func (b BlockTag) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString(b.Tag.String())
	bb.WriteString(fmt.Sprintf("%s", b.Body))
	bb.WriteString("</")
	bb.WriteString(b.Name)
	bb.WriteString(">")
	return bb.String()
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
