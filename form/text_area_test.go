package form_test

import (
	"testing"

	"github.com/markbates/tags"
	"github.com/markbates/tags/form"
	"github.com/stretchr/testify/require"
)

func Test_Form_TextArea(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ta := f.TextArea(tags.Options{
		"value": "hi",
	})
	r.Equal(`<textarea>hi</textarea>`, ta.String())
}
