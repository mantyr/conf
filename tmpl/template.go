// beta version
package tmpl

import (
    "github.com/mantyr/conf"
    "html/template"
    "path/filepath"
    "sync"
)

var (
    subdir_sync sync.RWMutex
    subdir  string

    funcMap template.FuncMap = template.FuncMap {
        "humanInt": HumanInt,
        "isStatus": IsStatus,
    }
)

func InitDir(address string) {
    SetDir(address)
}

func GetTemplate(address string) (*template.Template, error) {
    return template.New(filepath.Base(address)).Delims("{", "}").Funcs(funcMap).ParseFiles(conf.GetDirBin()+GetDir()+address)
}

func SetDir(address string) {
    subdir_sync.Lock()
    defer subdir_sync.Unlock()

    subdir = address
}

func GetDir() string {
    subdir_sync.RLock()
    defer subdir_sync.RUnlock()

    return subdir
}
