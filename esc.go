package tags

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/markbates/pop/nulls"
)

func esc(s interface{}) string {
	var st string
	switch t := s.(type) {
	case fmt.Stringer:
		st = t.String()
	case nulls.String:
		st = t.String
	case nulls.Int:
		st = strconv.Itoa(t.Int)
	case bool:
		st = fmt.Sprintf("%t", t)
	case string:
		st = t
	case int:
		st = strconv.Itoa(t)
	case nil:
		st = ""
	default:
		st = fmt.Sprintf("%s", s)
	}
	return template.HTMLEscapeString(st)
}
