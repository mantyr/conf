package tmpl

import (
    "github.com/mantyr/conf"
    "net/http"
    "log"
)

func Render(w http.ResponseWriter, data interface{}, address string) {
    t, err := GetTemplate(conf.GetDirBin()+GetDir()+address)
    if err != nil {
        log.Printf("Bad template, %q\r\n", address)
        return
    }
    t.Execute(w, data)
}