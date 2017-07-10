package tags

import (
	"html/template"
	"testing"

	"github.com/markbates/pop"
	"github.com/stretchr/testify/require"
)

func Test_Pagination(t *testing.T) {
	r := require.New(t)

	tag, err := Pagination(&pop.Paginator{
		Page:       2,
		TotalPages: 3,
	}, Options{
		"path": "/foo",
	})
	r.NoError(err)

	r.Equal(template.HTML("<ul class=\" pagination\"><li><a href=\"/foo?page=1\">&laquo;</a></li><li><a href=\"/foo?page=1\">1</a></li><li class=\"active\"><a href=\"/foo?page=2\">2</a></li><li><a href=\"/foo?page=3\">3</a></li><li><a href=\"/foo?page=3\">&raquo;</a></li></ul>"), tag.HTML())
}

func Test_Pagination_Page1(t *testing.T) {
	r := require.New(t)

	tag, err := Pagination(&pop.Paginator{
		Page:       1,
		TotalPages: 3,
	}, Options{
		"path": "/foo",
	})
	r.NoError(err)

	r.Equal(template.HTML("<ul class=\" pagination\"><li class=\"disabled\"><span>&laquo;</span></li><li class=\"active\"><a href=\"/foo?page=1\">1</a></li><li><a href=\"/foo?page=2\">2</a></li><li><a href=\"/foo?page=3\">3</a></li><li><a href=\"/foo?page=2\">&raquo;</a></li></ul>"), tag.HTML())
}

func Test_Pagination_Page3(t *testing.T) {
	r := require.New(t)

	tag, err := Pagination(&pop.Paginator{
		Page:       3,
		TotalPages: 3,
	}, Options{
		"path": "/foo",
	})
	r.NoError(err)

	r.Equal(template.HTML("<ul class=\" pagination\"><li><a href=\"/foo?page=2\">&laquo;</a></li><li><a href=\"/foo?page=1\">1</a></li><li><a href=\"/foo?page=2\">2</a></li><li class=\"active\"><a href=\"/foo?page=3\">3</a></li><li class=\"disabled\"><span>&raquo;</span></li></ul>"), tag.HTML())
}
