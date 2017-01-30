package tags_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/stretchr/testify/require"
)

func Test_Tag(t *testing.T) {
	r := require.New(t)
	tag := tags.New("input", tags.Options{})
	r.Equal("input", tag.Name)
}

func Test_Tag_WithName(t *testing.T) {
	r := require.New(t)
	tag := tags.New("div", tags.Options{})
	r.Equal("div", tag.Name)
	r.Equal(`<div />`, tag.String())
}

func Test_Tag_WithValue(t *testing.T) {
	r := require.New(t)
	tag := tags.New("input", tags.Options{
		"value": "Mark",
	})
	r.Equal(`<input value="Mark" />`, tag.String())
}

func Test_Tag_WithBody(t *testing.T) {
	r := require.New(t)

	tag := tags.New("div", tags.Options{
		"body": "hi there!",
	})
	r.Equal(`<div>hi there!</div>`, tag.String())
	r.Nil(tag.Options["body"])
}

func Test_Tag_String(t *testing.T) {
	r := require.New(t)

	tag := tags.New("div", tags.Options{
		"body": "hi there!",
	})
	r.Equal(`<div>hi there!</div>`, tag.String())
}

func Test_Tag_String_WithOpts(t *testing.T) {
	r := require.New(t)

	tag := tags.New("div", tags.Options{
		"body":  "hi there!",
		"class": "foo bar baz",
	})
	r.Equal(`<div class="foo bar baz">hi there!</div>`, tag.String())
}

func Test_Tag_String_SubTag(t *testing.T) {
	r := require.New(t)

	tag := tags.New("div", tags.Options{
		"body": tags.New("p", tags.Options{
			"body": "hi!",
		}),
	})
	r.Equal(`<div><p>hi!</p></div>`, tag.String())
}
