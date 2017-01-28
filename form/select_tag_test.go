package form_test

import (
	"testing"

	"github.com/markbates/tags"
	"github.com/markbates/tags/form"
	"github.com/stretchr/testify/require"
)

func Test_SelectTag(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{})
	r.Equal(`<select></select>`, s.String())
}

func Test_SelectTag_WithSelectOptions(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options": form.SelectOptions{
			{Value: 1, Label: "one"},
			{Value: 2, Label: "two"},
		},
	})
	r.Equal(`<select><option value="1">one</option><option value="2">two</option></select>`, s.String())
}

func Test_SelectTag_WithSelectOptions_Selected(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options": form.SelectOptions{
			{Value: 3, Label: "three"},
			{Value: 2, Label: "two"},
		},
		"selected": "3",
	})
	r.Equal(`<select><option value="3" selected>three</option><option value="2">two</option></select>`, s.String())
}

func Test_SelectTag_WithMap(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options": map[string]interface{}{
			"one": 1,
			"two": 2,
		},
	})
	r.Equal(`<select><option value="1">one</option><option value="2">two</option></select>`, s.String())
}

func Test_SelectTag_WithMap_Selected(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options": map[string]interface{}{
			"three": 3,
			"two":   2,
		},
		"selected": 3,
	})
	r.Equal(`<select><option value="3" selected>three</option><option value="2">two</option></select>`, s.String())
}

func Test_SelectTag_WithSlice(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options": []string{"one", "two"},
	})
	r.Equal(`<select><option value="one">one</option><option value="two">two</option></select>`, s.String())
}

func Test_SelectTag_WithSlice_Selected(t *testing.T) {
	r := require.New(t)
	f := form.New(tags.Options{})
	s := f.SelectTag(tags.Options{
		"options":  []string{"one", "two"},
		"selected": "two",
	})
	r.Equal(`<select><option value="one">one</option><option value="two" selected>two</option></select>`, s.String())
}
