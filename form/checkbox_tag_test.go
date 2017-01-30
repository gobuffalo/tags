package form_test

import (
	"testing"

	"github.com/markbates/tags"
	"github.com/markbates/tags/form"
	"github.com/stretchr/testify/require"
)

func Test_Form_CheckboxTag(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{})
	r.Equal(`<label><input type="checkbox" /> </label>`, ct.String())
}

func Test_Form_CheckboxTag_WithValue(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value": 1,
	})
	r.Equal(`<label><input type="checkbox" value="1" /> </label>`, ct.String())
}

func Test_Form_CheckboxTag_WithValueSelected(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value":   1,
		"checked": "1",
	})
	r.Equal(`<label><input type="checkbox" value="1" checked /> </label>`, ct.String())
}

func Test_Form_CheckboxTag_WithLabel(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value": 1,
		"label": "check me",
	})
	r.Equal(`<label><input type="checkbox" value="1" /> check me</label>`, ct.String())
}
