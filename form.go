package tags

import (
	"bytes"
	"fmt"
	"strings"
)

const authToken = "authenticity_token"

type Form struct {
	*BlockTag
	subTags []*Tag
}

func (f Form) String() string {
	bb := &bytes.Buffer{}
	for _, t := range f.subTags {
		bb.WriteString(t.String())
	}
	if f.Body != nil {
		bb.WriteString(fmt.Sprintf("%s", f.Body))
	}
	f.Body = bb.String()
	return f.BlockTag.String()
}

func NewForm(opts Options) *Form {
	if opts["method"] == nil {
		opts["method"] = "POST"
	}
	form := &Form{
		BlockTag: NewBlockTag("form", opts),
	}
	if form.Options[authToken] != nil {
		at := New("input", Options{
			"value": form.Options[authToken],
			"type":  "hidden",
			"name":  authToken,
		})
		form.subTags = append(form.subTags, at)
		delete(form.Options, authToken)
	}

	m := strings.ToUpper(form.Options["method"].(string))
	if m != "POST" && m != "GET" {
		form.Options["method"] = "POST"
		mt := New("input", Options{
			"value": m,
			"type":  "hidden",
			"name":  "_method",
		})
		form.subTags = append(form.subTags, mt)
	}

	return form
}
