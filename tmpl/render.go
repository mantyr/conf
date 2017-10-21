package tmpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mantyr/debug"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func RenderHTML(data interface{}, address string) (r template.HTML, err error) {
	var v string
	v, err = RenderString(data, address)
	r = template.HTML(v)
	return
}

func RenderString(data interface{}, address string) (r string, err error) {
	var t *template.Template
	t, err = GetTemplate(address)
	if err != nil {
		log.Printf("Bad template, %q, %q\r\n", address, err)
		return
	}
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, data)
	if err != nil {
		debug.Printf("Bad data in template, %q, %q\r\n", address, data)
	}

	var byt []byte
	byt, err = ioutil.ReadAll(buf)
	r = string(byt)
	return
}

func Render(w http.ResponseWriter, data interface{}, address string) {
	t, err := GetTemplate(address)
	if err != nil {
		log.Printf("Bad template, %q, %q\r\n", address, err)
		return
	}
	t.Execute(w, data)
}

func RenderText(w http.ResponseWriter, data interface{}, address string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Render(w, data, address)
}

func Json(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s\r\n", Json_error_data_en)
		return
	}
	fmt.Fprintf(w, "%s\r\n", b)
}

func JsonString(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	fmt.Fprintf(w, "%s\r\n", data)
}

func MessageError(w http.ResponseWriter, message ...string) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	data := Json_error_data{
		Message: strings.Join(message, " "),
		Status:  "Error",
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s\r\n", Json_error_data_en)
		return
	}
	fmt.Fprintf(w, "%s\r\n", b)
}

func MessageOK(w http.ResponseWriter, message ...string) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	data := Json_error_data{
		Message: strings.Join(message, " "),
		Status:  "OK",
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s\r\n", Json_error_data_en)
		return
	}
	fmt.Fprintf(w, "%s\r\n", b)
}
