package tags_test

import (
	"testing"

	"github.com/markbates/tags"
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
	r.Equal(`<div>`, tag.String())
}

func Test_Tag_WithValue(t *testing.T) {
	r := require.New(t)
	tag := tags.New("input", tags.Options{
		"value": "Mark",
	})
	r.Equal(`<input value="Mark">`, tag.String())
}
