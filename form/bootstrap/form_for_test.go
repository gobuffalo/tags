package bootstrap

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/tags/v3"
	"github.com/gobuffalo/validate/v3"
	"github.com/stretchr/testify/require"
)

func Test_InputFieldLabel(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Custom</label><input class="form-control" id="-Name" name="Name" type="text" value="" /></div>`, l.String())
}

func Test_Input_Floating(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.InputTag("Name", tags.Options{
		"bootstrap": map[string]interface{}{"form-group-class": "form-floating"},
	})
	r.Equal(`<div class="form-floating"><input class="form-control" id="-Name" name="Name" placeholder="Name" type="text" value="" /><label class="form-label" for="-Name">Name</label></div>`, l.String())
}

func Test_InputFieldLabelWithAchronym(t *testing.T) {
	cases := map[string]string{
		"URL":            "URL",
		"MyURL":          "My URL",
		"SimpleURIAdded": "Simple URI Added",
		"GaveAnExample":  "Gave An Example",
	}
	r := require.New(t)
	f := NewFormFor(struct{ URL string }{}, tags.Options{})

	for key, expectedLabel := range cases {
		l := f.InputTag(key, tags.Options{})
		r.Equal(`<div class="form-group"><label class="form-label" for="-`+key+`">`+expectedLabel+`</label><input class="form-control" id="-`+key+`" name="`+key+`" type="text" value="" /></div>`, l.String())
	}
}

func Test_InputFieldLabel_Humanized(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ LongName string }{}, tags.Options{})
	l := f.InputTag("LongName", tags.Options{})
	r.Equal(`<div class="form-group"><label class="form-label" for="-LongName">Long Name</label><input class="form-control" id="-LongName" name="LongName" type="text" value="" /></div>`, l.String())
}

func Test_InputFieldSchema(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct {
		Name string `schema:"notName"`
	}{}, tags.Options{})

	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-notName">Custom</label><input class="form-control" id="-notName" name="notName" type="text" value="" /></div>`, l.String())
}

func Test_InputFieldFormInsteadOfSchema(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct {
		Name string `form:"notName"`
	}{}, tags.Options{})

	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-notName">Custom</label><input class="form-control" id="-notName" name="notName" type="text" value="" /></div>`, l.String())
}

func Test_InputFieldFormAndSchema(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct {
		Name string `form:"notName" schema:"name"`
	}{}, tags.Options{})

	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-notName">Custom</label><input class="form-control" id="-notName" name="notName" type="text" value="" /></div>`, l.String())
}

func Test_InputFieldSchema_FieldNotPresent(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct {
		Name string `schema:"notName"`
	}{}, tags.Options{})

	l := f.InputTag("Other", tags.Options{})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Other">Other</label><input class="form-control" id="-Other" name="Other" type="text" value="" /></div>`, l.String())
}

func Test_InputFieldSchema_FieldDash(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct {
		Name string `schema:"-"`
	}{}, tags.Options{})

	l := f.InputTag("Name", tags.Options{})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Name</label><input class="form-control" id="-Name" name="Name" type="text" value="" /></div>`, l.String())
}

func Test_SelectLabel(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.SelectTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Custom</label><select class="form-control" id="-Name" name="Name"></select></div>`, l.String())
}

func Test_Select_With_String_As_BeforeTag_Opt(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})

	s := `<span>Test</span>`
	l := f.SelectTag("Name", tags.Options{"before_tag": s})

	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Name</label><span>Test</span><select class="form-control" id="-Name" name="Name"></select></div>`, l.String())
}

func Test_Select_With_Nested_Tag_As_BeforeTag_Opt(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})

	s := tags.New("span", tags.Options{"body": "Test"})
	l := f.SelectTag("Name", tags.Options{"before_tag": s})

	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Name</label><span>Test</span><select class="form-control" id="-Name" name="Name"></select></div>`, l.String())
}

func Test_Select_With_String_As_AfterTag_Opt(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})

	b := `<button type="button">Button Name</button>`
	l := f.SelectTag("Name", tags.Options{"after_tag": b})

	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Name</label><select class="form-control" id="-Name" name="Name"></select><button type="button">Button Name</button></div>`, l.String())
}

func Test_Select_With_Nested_Tag_As_AfterTag_Opt(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})

	b := tags.New("button", tags.Options{
		"body": "Button Name",
		"type": "button",
	})
	l := f.SelectTag("Name", tags.Options{"after_tag": b})

	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Name</label><select class="form-control" id="-Name" name="Name"></select><button type="button">Button Name</button></div>`, l.String())
}

func Test_RadioButton(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.RadioButton("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Custom</label><label for="-Name"><input class="form-control" id="-Name" name="Name" type="radio" value="" /> </label></div>`, l.String())
}
func Test_TextArea(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.TextArea("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Custom</label><textarea class="form-control" id="-Name" name="Name"></textarea></div>`, l.String())
}

func Test_SubmitTag(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	i := f.SubmitTag("Send", tags.Options{"class": "btn btn-primary"})

	r.Equal(`<input class="btn btn-primary" type="submit" value="Send" />`, i.String())
}

func Test_CheckBox(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.CheckboxTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-check"><input class="form-check-input" id="-Name" name="Name" type="checkbox" value="true" /><label class="form-check-label" for="-Name">Custom</label></div>`, l.String())
}

func Test_CheckBox_Unchecked(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.CheckboxTag("Name", tags.Options{
		"label":     "Custom",
		"unchecked": "false",
	})
	r.Equal(`<div class="form-check"><input class="form-check-input" id="-Name" name="Name" type="checkbox" value="true" /><input name="Name" type="hidden" value="false" /><label class="form-check-label" for="-Name">Custom</label></div>`, l.String())
}

func Test_CheckBox_WithHelp(t *testing.T) {
	r := require.New(t)
	f := NewFormFor(struct{ Name string }{}, tags.Options{})
	l := f.CheckboxTag("Name", tags.Options{
		"label": "Custom",
		"help":  "Help message",
	})
	r.Equal(`<div class="form-check"><input class="form-check-input" id="-Name" name="Name" type="checkbox" value="true" /><label class="form-check-label" for="-Name">Custom</label><div class="form-text" id="Name-help">Help message</div></div>`, l.String())
}

type CustomError struct{}

func (ce CustomError) Get(key string) []string {
	return []string{"My Custom Error"}
}

func Test_InputError_CustomError(t *testing.T) {
	r := require.New(t)

	errors := CustomError{}

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group has-error"><label class="form-label" for="-Name">Custom</label><input class="form-control is-invalid" id="-Name" name="Name" type="text" value="" /><div class="invalid-feedback help-block">My Custom Error</div></div>`, l.String())
}

func Test_InputError(t *testing.T) {
	r := require.New(t)

	errors := validate.NewErrors()
	errors.Add("name", "Name shoud be AJ.")

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group has-error"><label class="form-label" for="-Name">Custom</label><input class="form-control is-invalid" id="-Name" name="Name" type="text" value="" /><div class="invalid-feedback help-block">Name shoud be AJ.</div></div>`, l.String())
}

func Test_InputHidden(t *testing.T) {
	r := require.New(t)

	errors := validate.NewErrors()
	errors.Add("name", "Name shoud be AJ.")

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"type": "hidden"})
	r.Equal(`<input errors="[Name shoud be AJ.]" id="-Name" name="Name" tags-field="Name" type="hidden" value="" />`, l.String())

	l = f.HiddenTag("Name", tags.Options{})
	r.Equal(`<input errors="[Name shoud be AJ.]" id="-Name" name="Name" tags-field="Name" type="hidden" value="" />`, l.String())
}

func Test_InputError_Map(t *testing.T) {
	r := require.New(t)

	errors := map[string][]string{
		"name": {"Name shoud be AJ."},
	}

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group has-error"><label class="form-label" for="-Name">Custom</label><input class="form-control is-invalid" id="-Name" name="Name" type="text" value="" /><div class="invalid-feedback help-block">Name shoud be AJ.</div></div>`, l.String())
}

func Test_InputError_InvalidMap(t *testing.T) {
	r := require.New(t)

	errors := map[string]string{
		"name": "Name shoud be AJ.",
	}

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group"><label class="form-label" for="-Name">Custom</label><input class="form-control" id="-Name" name="Name" type="text" value="" /></div>`, l.String())
}

func Test_InputMultipleError(t *testing.T) {
	r := require.New(t)

	errors := validate.NewErrors()
	errors.Add("name", "Name shoud be AJ.")
	errors.Add("name", "Name shoud start with A.")

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.InputTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-group has-error"><label class="form-label" for="-Name">Custom</label><input class="form-control is-invalid" id="-Name" name="Name" type="text" value="" /><div class="invalid-feedback help-block">Name shoud be AJ.</div><div class="invalid-feedback help-block">Name shoud start with A.</div></div>`, l.String())
}

func Test_CheckBoxError(t *testing.T) {
	r := require.New(t)

	errors := validate.NewErrors()
	errors.Add("name", "Name shoud be AJ.")

	f := NewFormFor(struct{ Name string }{}, tags.Options{"errors": errors})
	l := f.CheckboxTag("Name", tags.Options{"label": "Custom"})
	r.Equal(`<div class="form-check has-error"><input class="form-check-input is-invalid" id="-Name" name="Name" type="checkbox" value="true" /><label class="form-check-label" for="-Name">Custom</label><div class="invalid-feedback help-block">Name shoud be AJ.</div></div>`, l.String())
}

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	City  string
	State string
}

func Test_FormFor_Nested_Struct(t *testing.T) {
	r := require.New(t)
	p := Person{
		Name: "Mark",
		Address: Address{
			City:  "Boston",
			State: "MA",
		},
	}

	f := NewFormFor(p, tags.Options{})
	tag := f.InputTag("Address.State", tags.Options{})

	exp := `<div class="form-group"><label class="form-label" for="person-Address.State">Address State</label><input class="form-control" id="person-Address.State" name="Address.State" type="text" value="MA" /></div>`
	r.Equal(exp, tag.String())
}

func Test_Field_TagOnly(t *testing.T) {
	f := NewFormFor(struct {
		Name string `schema:"-"`
	}{}, tags.Options{})

	cases := []struct {
		f      func(field string, opt tags.Options) *tags.Tag
		name   string
		opts   tags.Options
		output string
	}{
		{
			f:    f.InputTag,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
				"class":    "custom-input",
			},
			output: `<input class="custom-input" id="-Name" name="Name" type="text" value="" />`,
		},

		{
			f:    f.TextArea,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
			},
			output: `<textarea class="" id="-Name" name="Name"></textarea>`,
		},

		{
			f:    f.RadioButton,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
			},
			output: `<input class="" id="-Name" name="Name" type="radio" value="" />`,
		},

		{
			f:    f.CheckboxTag,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
				"class":    "custom-input",
			},
			output: `<input class="custom-input" id="-Name" name="Name" type="checkbox" value="true" />`,
		},

		{
			f:    f.SelectTag,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
			},
			output: `<select class="" id="-Name" name="Name"></select>`,
		},

		{
			f:    f.FileTag,
			name: "Name",
			opts: tags.Options{
				"tag_only": true,
			},
			output: `<input class="" id="-Name" name="Name" type="file" value="" />`,
		},
	}

	for index, tcase := range cases {
		t.Run(fmt.Sprintf("%v", index), func(tt *testing.T) {
			r := require.New(tt)
			l := tcase.f(tcase.name, tcase.opts)
			r.Equal(tcase.output, l.String())
		})
	}
}
