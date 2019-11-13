package form_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
	"github.com/stretchr/testify/require"
)

func Test_Form_CheckboxTag(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{"name": "Chubby"})
	r.Equal(`<label><input name="Chubby" type="checkbox" value="true" /></label>`, ct.String())
}

func Test_Form_CheckboxTag_WithValue(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value":     1,
		"checked":   "1",
		"unchecked": "2",
		"name":      "Chubby",
	})
	r.Equal(`<label><input name="Chubby" type="checkbox" value="1" checked /><input name="Chubby" type="hidden" value="2" /></label>`, ct.String())
}

func Test_Form_CheckboxTag_WithLabel(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"label": "check me",
		"name":  "Chubby",
	})
	r.Equal(`<label><input name="Chubby" type="checkbox" value="true" /> check me</label>`, ct.String())
}

func Test_Form_CheckboxTag_With_Text_Value(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value": "Custom Value",
		"name":  "Chubby",
	})
	r.Equal(`<label><input name="Chubby" type="checkbox" value="Custom Value" /></label>`, ct.String())
}

func Test_Form_CheckboxTag_Tag_Only(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"tag_only": true,
		"name":     "Chubby",
	})
	r.Equal(`<input name="Chubby" type="checkbox" value="true" />`, ct.String())
}

func Test_Form_CheckboxTag_With_Empty_Text_Value(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"value": "",
		"name":  "Chubby",
	})
	r.Equal(`<label><input name="Chubby" type="checkbox" value="" /></label>`, ct.String())
}

func Test_Form_CheckboxTag_TagOnly_With_CustomValue(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	ct := f.CheckboxTag(tags.Options{
		"tag_only": true,
		"value":    "Cutsom Value",
		"name":     "Chubby",
	})
	r.Equal(`<input name="Chubby" type="checkbox" value="Cutsom Value" />`, ct.String())
}
