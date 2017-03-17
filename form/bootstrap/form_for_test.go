package bootstrap_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form/bootstrap"
	"github.com/stretchr/testify/require"
)

func Test_InputFieldLabel(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label>Custom</label><input class=" form-control" id="-Name" name="Name" type="text" value="" /></div>`, l.String())
}

func Test_SelectLabel(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.SelectTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label>Custom</label><select class=" form-control" id="-Name" name="Name" /></div>`, l.String())
}

func Test_RadioButton(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.RadioButton("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label>Custom</label><label><input class=" form-control" id="-Name" name="Name" type="radio" value="" /> </label></div>`, l.String())
}
func Test_TextArea(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.TextArea("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label>Custom</label><textarea class=" form-control" id="-Name" name="Name"></textarea></div>`, l.String())
}

func Test_CheckBox(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.CheckboxTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label><input name="Name" type="hidden" value="false" /><input class="" id="-Name" name="Name" type="checkbox" value="true" />Custom</label></div>`, l.String())
}
