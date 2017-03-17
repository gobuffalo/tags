package form_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
	"github.com/stretchr/testify/require"
)

type Talk struct {
}

func Test_NewFormFor(t *testing.T) {
	r := require.New(t)

	f := form.NewFormFor(Talk{}, tags.Options{
		"action": "/users/1",
	})
	r.Equal("form", f.Name)
	r.Equal(`<form action="/users/1" id="talk-form" method="POST" />`, f.String())
}

func Test_NewFormFor_With_AuthenticityToken(t *testing.T) {
	r := require.New(t)

	f := form.NewFormFor(Talk{}, tags.Options{
		"action": "/users/1",
	})
	f.SetAuthenticityToken("12345")
	r.Equal("form", f.Name)
	r.Equal(`<form action="/users/1" id="talk-form" method="POST"><input name="authenticity_token" type="hidden" value="12345" /></form>`, f.String())
}

func Test_NewFormFor_With_NotPostMethod(t *testing.T) {
	r := require.New(t)

	f := form.NewFormFor(Talk{}, tags.Options{
		"action": "/users/1",
		"method": "put",
	})
	r.Equal("form", f.Name)
	r.Equal(`<form action="/users/1" id="talk-form" method="POST"><input name="_method" type="hidden" value="PUT" /></form>`, f.String())
}

func Test_FormFor_Label(t *testing.T) {
	r := require.New(t)
	f := form.NewFormFor(Talk{}, tags.Options{})
	l := f.Label("Name", tags.Options{})
	r.Equal(`<label>Name</label>`, l.String())
}

func Test_FormFor_FieldDoesntExist(t *testing.T) {
	r := require.New(t)
	f := form.NewFormFor(Talk{}, tags.Options{})
	l := f.InputTag("IDontExist", tags.Options{})
	r.Equal(`<input id="talk-IDontExist" name="IDontExist" type="text" value="" />`, l.String())
}
