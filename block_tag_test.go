package tags_test

import (
	"testing"

	"github.com/markbates/tags"
	"github.com/stretchr/testify/require"
)

func Test_NewBlockTag_WithBody(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{
		"body": "hi there!",
	})
	r.Equal(tag.Body, "hi there!")
	r.Nil(tag.Options["body"])
}

func Test_NewBlockTag_WithValue(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{
		"value": "hi there!",
	})
	r.Equal(tag.Body, "hi there!")
	r.Nil(tag.Options["value"])
}

func Test_BlockTag_String(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{
		"body": "hi there!",
	})
	r.Equal(`<div>hi there!</div>`, tag.String())
}

func Test_BlockTag_String_NoBody(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{})
	r.Equal(`<div></div>`, tag.String())
}

func Test_BlockTag_String_WithOpts(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{
		"body":  "hi there!",
		"class": "foo bar baz",
	})
	r.Equal(`<div class="foo bar baz">hi there!</div>`, tag.String())
}

func Test_BlockTag_String_SubTag(t *testing.T) {
	r := require.New(t)

	tag := tags.NewBlockTag("div", tags.Options{
		"body": tags.NewBlockTag("p", tags.Options{
			"body": "hi!",
		}),
	})
	r.Equal(`<div><p>hi!</p></div>`, tag.String())
}
