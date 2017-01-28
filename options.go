package tags

import (
	"fmt"
	"sort"
	"strings"
)

type Options map[string]interface{}

func (o Options) String() string {
	var out = make([]string, 0, len(o))
	var tmp = make([]string, 2)
	for k, v := range o {
		tmp[0] = esc(k)
		tmp[1] = fmt.Sprintf("\"%s\"", esc(v))
		out = append(out, strings.Join(tmp, "="))
	}
	sort.Strings(out)
	return strings.Join(out, " ")
}
