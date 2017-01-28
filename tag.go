package tags

import "bytes"

type Tag struct {
	Name    string
	Options Options
}

func (t Tag) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("<")
	bb.WriteString(t.Name)
	if len(t.Options) > 0 {
		bb.WriteString(" ")
		bb.WriteString(t.Options.String())
	}
	bb.WriteString(">")
	return bb.String()
}

func New(name string, opts Options) *Tag {
	tag := &Tag{
		Name:    name,
		Options: opts,
	}

	return tag
}
