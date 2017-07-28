package form_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
	"github.com/stretchr/testify/require"
)

func Test_Form_SubmitTag(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	i := f.SubmitTag(tags.Options{"value": "Send"})

	r.Equal(`<input type="submit" value="Send" />`, i.String())
}
