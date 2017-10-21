package tmpl

import (
	"encoding/json"
	"html/template"
)

func TMPL_json(v interface{}) template.JS {
	s, err := json.Marshal(v)
	if err == nil {
		return template.JS(s)
	}
	return template.JS("")
}
