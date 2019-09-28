package bootstrap_test

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form/bootstrap"
	"github.com/stretchr/testify/require"
)

func Test_BootstrapFormGroupClass(t *testing.T) {
	r := require.New(t)
	f := bootstrap.NewFormFor(struct{ Name string }{}, tags.Options{})

	tcases := []struct {
		options  tags.Options
		expected string
	}{
		{
			expected: `<div class="form-group row"><label>Name</label><input class=" form-control" id="-Name" name="Name" type="text" value="" /></div>`,
			options: tags.Options{
				"bootstrap": map[string]interface{}{
					"form-group-class": "form-group row",
				},
			},
		},

		{
			expected: `<div class="form-group"><label>Name</label><input class=" form-control" id="-Name" name="Name" type="text" value="" /></div>`,
			options:  tags.Options{},
		},
	}

	for _, tcase := range tcases {
		l := f.InputTag("Name", tcase.options)
		r.Equal(tcase.expected, l.String())
	}

}
