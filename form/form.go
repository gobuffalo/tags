package form

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/markbates/tags"
)

const authToken = "authenticity_token"

type Form struct {
	*tags.BlockTag
	subTags []*tags.Tag
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

func New(opts tags.Options) *Form {
	if opts["method"] == nil {
		opts["method"] = "POST"
	}
	form := &Form{
		BlockTag: tags.NewBlockTag("form", opts),
	}
	if form.Options[authToken] != nil {
		at := tags.New("input", tags.Options{
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
		mt := tags.New("input", tags.Options{
			"value": m,
			"type":  "hidden",
			"name":  "_method",
		})
		form.subTags = append(form.subTags, mt)
	}

	return form
}
